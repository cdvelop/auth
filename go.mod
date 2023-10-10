module github.com/cdvelop/auth

go 1.20

require (
	github.com/cdvelop/model v0.0.51
	golang.org/x/oauth2 v0.9.0
)

require (
	cloud.google.com/go/compute v1.20.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/s2a-go v0.1.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.5 // indirect
	github.com/googleapis/gax-go/v2 v2.11.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629192749-1f95444afc44 // indirect
	google.golang.org/grpc v1.56.1 // indirect
)

require (
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/cdvelop/input v0.0.26
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.11.0 // indirect
	google.golang.org/api v0.129.0
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/input => ../input
