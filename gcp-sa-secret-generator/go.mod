module github.com/jrthrawny/kustomize-plugins/gcp-sa-secret-generator

go 1.12

require (
	github.com/evanphx/json-patch v4.5.0+incompatible
	github.com/go-openapi/spec v0.19.2
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.2
	github.com/spf13/pflag v1.0.3
	google.golang.org/api v0.9.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apimachinery v0.0.0-20190313205120-d7deff9243b1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20190603182131-db7b694dc208
	sigs.k8s.io/kustomize/v3 v3.1.0 // indirect
	sigs.k8s.io/yaml v1.1.0
)
