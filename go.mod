module github.com/bmutziu/autobucket-operator

go 1.13

require (
	cloud.google.com/go/storage v1.12.0
	github.com/go-logr/logr v0.1.0
	github.com/joho/godotenv v1.3.0
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/stretchr/testify v1.6.1
	google.golang.org/api v0.32.0
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	sigs.k8s.io/controller-runtime v0.6.3
)
