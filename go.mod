module github.com/satimoto/go-api

go 1.16

require (
	github.com/99designs/gqlgen v0.14.0
	github.com/aws/aws-lambda-go v1.27.0
	github.com/aws/aws-sdk-go v1.42.17
	github.com/awslabs/aws-lambda-go-api-proxy v0.11.0
	github.com/fiatjaf/go-lnurl v1.9.2
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-chi/cors v1.2.0
	github.com/go-chi/render v1.0.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.4.0
	github.com/lightningnetwork/lnd v0.14.2-beta.rc1 // indirect
	github.com/satimoto/go-datastore v0.1.2-0.20220325212000-2f869873391c
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/vektah/gqlparser/v2 v2.2.0
)

replace git.schwanenlied.me/yawning/bsaes.git => github.com/Yawning/bsaes v0.0.0-20180720073208-c0276d75487e
