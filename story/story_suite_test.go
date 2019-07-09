package story_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Story Suite")
}
