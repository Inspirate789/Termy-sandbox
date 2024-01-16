package middleware

import (
	"crypto/md5"
	"encoding/base64"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/migrations"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/test"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/services/storage"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"os"
	"strconv"
	"testing"
	"time"
)

func Test2FA(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Suite")
}

var (
	dbContainer   testcontainers.Container
	db            *sqlx.DB
	migration     *migrations.Migration
	store         storage.Storage
	authenticator *jwt.GinJWTMiddleware
)

var _ = BeforeSuite(func() {
	Context("Setup test database", func() {
		var err error
		dbContainer, db, err = test.SetupTestDatabase()
		Expect(err).To(Succeed())
		Expect(dbContainer).NotTo(BeNil())
		Expect(db).NotTo(BeNil())
		migration = migrations.NewMigration(db)
		Expect(migration.UpTest()).To(Succeed())
	})
})

var _ = AfterSuite(func() {
	Context("Cleanup after test execution", func() {
		store = nil
		authenticator = nil
		Expect(migration.DownTest()).To(Succeed())
	})
	Context("Teardown test database", func() {
		Expect(test.TeardownTestDatabase(dbContainer, db)).To(Succeed())
	})
})

var _ = Describe("Two-factor authentication", func() {
	setupAuthenticator := func() {
		Context("Prepare for test execution", func() {
			store = postgres_storage.NewPostgresStorage(db)
			var err error
			authenticator, err = MakeGinJWTMiddlewareWith2FA([]string{
				entities.AdminRole,
			}, store.GetUser, os.Getenv("TELEGRAM_API_TOKEN"))
			Expect(authenticator).NotTo(BeNil())
			Expect(err).To(Succeed())
		})
	}

	id, _ := strconv.ParseInt(os.Getenv("TEST_USER_ID"), 10, 64)
	validTestUser := &entities.User{
		ID:       int(id),
		Name:     os.Getenv("TEST_USER_NAME"),
		Password: os.Getenv("TEST_USER_PASSWORD"),
		Role:     entities.AdminRole,
	}
	validTestUser.SetRegDate(time.Now())
	invalidTestUser := &entities.User{}

	var validClaims jwt.MapClaims
	When("I log in to my account", func() {
		Context("and enter the valid username and password", func() {
			BeforeEach(setupAuthenticator)
			Specify("the claims are not empty", func() {
				validClaims = authenticator.PayloadFunc(validTestUser)
				Expect(validClaims).NotTo(BeEmpty())
			})
		})
		Context("and enter an invalid username and password", func() {
			BeforeEach(setupAuthenticator)
			Specify("the claims are empty", func() {
				invalidClaims := authenticator.PayloadFunc(invalidTestUser)
				Expect(invalidClaims).To(BeEmpty())
			})
		})
	})

	When("I authenticate", func() {
		Context("and give the valid authentication code", func() {
			BeforeEach(setupAuthenticator)
			It("accept authentication", func() {
				validCtx := new(gin.Context)
				validCtx.Set("JWT_PAYLOAD", validClaims)
				hasher := md5.New()
				hasher.Write([]byte(strconv.Itoa(validTestUser.ID)))
				authCode := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
				validCtx.Set(factorAuthKey, authCode)
				identity := authenticator.IdentityHandler(validCtx)
				authenticated := authenticator.Authorizator(identity, validCtx)
				Expect(authenticated).To(BeTrue())
			})
		})
		Context("and give an invalid authentication code", func() {
			BeforeEach(setupAuthenticator)
			invalidCtx := new(gin.Context)
			It("reject authentication", func() {
				identity := authenticator.IdentityHandler(invalidCtx)
				authenticated := authenticator.Authorizator(identity, invalidCtx)
				Expect(authenticated).To(BeFalse())
			})
		})
	})
})
