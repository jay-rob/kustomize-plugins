module github.com/jrthrawny/kustomize-plugins/gcp-sa-secret-generator

go 1.12

require (
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.9.0
	k8s.io/apimachinery v0.15.7 // indirect
	sigs.k8s.io/kustomize/v3 v3.1.0
	sigs.k8s.io/yaml v1.1.0
)
