module github.com/KatherineGuseva/kubePlayGround

go 1.15

require (
	github.com/IBM-Cloud/bluemix-go v0.0.0-20210319111107-fd88f7966d1c
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.18.0
	k8s.io/apimachinery v0.20.5
	k8s.io/client-go v0.0.0-00010101000000-000000000000

)

replace k8s.io/client-go => k8s.io/client-go v0.18.0
