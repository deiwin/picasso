package picasso_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPicasso(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Picasso Suite")
}
