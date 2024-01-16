### Vault setup tutorial

```shell
vault server -dev
```

```shell
vault login <ROOT_TOKEN>
vault auth enable jwt
vault write auth/jwt/config \
  oidc_discovery_url="https://git.iu7.bmstu.ru" \
  bound_issuer="git.iu7.bmstu.ru"
vault policy write testapp - <<EOF
path "secret/*" {             
  capabilities = ["read"]
}
EOF
vault write auth/jwt/role/testapp - <<EOF
{
  "role_type": "jwt",
  "policies": ["testapp"],
  "token_explicit_max_ttl": 120,
  "user_claim": "user_email",
  "bound_claims": {
    "project_id": "5998" 
  }
}
EOF
```

```shell
vault kv put secret/test/telegram_api_token telegram_api_token=<TG_API_TOKEN>
vault kv put secret/test/user_id            user_id=<TG_USER_ID>
vault kv put secret/test/user_name          user_name=Inspirate789
vault kv put secret/test/user_password      user_password=12345

vault kv get -field=telegram_api_token secret/test/telegram_api_token
vault kv get -field=user_id            secret/test/user_id
vault kv get -field=user_name          secret/test/user_name
vault kv get -field=user_password      secret/test/user_password
```
