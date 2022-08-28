# HK - 🍎 Health Kit
Set of utilities to parse analyze Apple HealthKit exported dataset.

```shell
# Parsing a export
hk -file=./data/at/export.xml

# Update all dependencies
go get -u ./...
go mod tidy
go get github.com/google/uuid # Get and install a specific package
go mod download # Download all dependencies

# Build a docker image
docker build -t hk:v1 .
```

Required environment variable:
```shell
USENSE_OAUTH2_CLIENT_ID= "" # OAuth2 client ID from console.google.com 
USENSE_OAUTH2_CLIENT_SECRET= "" # OAuth2 client Secret from console.google.com 
USENSE_ENVIRONMENT= "DEVELOPMENT" # or "PROD"
``` 

## Google Authentication
The repo uses Google OAuth2 to authenticate users. Some helpful links are:
 - [Google Sign-In for server-side apps](https://developers.google.com/identity/sign-in/web/server-side-flow)
 - [Getting Started with OAuth2 in Go](https://itnext.io/getting-started-with-oauth2-in-go-1c692420e03)
 - [Upload an object with HTML forms](https://cloud.google.com/storage/docs/xml-api/post-object-forms)
 - [OAuth 2.0 Scopes for Google APIs](https://cloud.google.com/docs/authentication/production#cloud-console)