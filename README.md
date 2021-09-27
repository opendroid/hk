# HK - 🍎 Health Kit
Set of utilities to parse analyze Apple HealthKit exported dataset.

```shell
# Parsing a export
hk -file=/path/to/apple_health_export/export.xml

# Build a docker image
docker build -t hk:v1 .
```

## Google Authentication
The repo uses Google OAuth2 to authenticate users. Some helpful links are:
 - [Google Sign-In for server-side apps](https://developers.google.com/identity/sign-in/web/server-side-flow)
 - [Getting Started with OAuth2 in Go](https://itnext.io/getting-started-with-oauth2-in-go-1c692420e03)
 - [Upload an object with HTML forms](https://cloud.google.com/storage/docs/xml-api/post-object-forms)
 - [OAuth 2.0 Scopes for Google APIs](https://cloud.google.com/docs/authentication/production#cloud-console)