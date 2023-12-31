package server

import (
	_ "git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/server/middleware"
	_ "git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

// login godoc
//
//	@Summary		Log in to the server.
//	@Description	log in to the server
//	@Tags			Auth
//	@Param			request	body	entities.AuthRequest	true	"Authentication request"
//	@Produce		json
//	@Success		200	{object}	middleware.LoginResponse
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/login [post]
func (s *Server) login(f gin.HandlerFunc) gin.HandlerFunc {
	return f
}

// refresh godoc
//
//	@Summary		Refresh user authentication token.
//	@Description	refresh user authentication token
//	@Tags			Auth
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string	true	"Authorization"
//	@Success		200				{object}	middleware.LoginResponse
//	@Failure		400				{object}	string
//	@Failure		401				{object}	string
//	@Failure		500				{object}	string
//	@Router			/refresh [get]
func (s *Server) refresh(f gin.HandlerFunc) gin.HandlerFunc {
	return f
}

// logout godoc
//
//	@Summary		Log out from the server.
//	@Description	log out from the server
//	@Tags			Auth
//	@Security		ApiKeyAuth
//	@Param			Authorization	header	string	true	"Authorization"
//	@Success		200
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/logout [delete]
func (s *Server) logout(f gin.HandlerFunc) gin.HandlerFunc {
	return f
}
