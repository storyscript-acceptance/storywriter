package acceptance_test

import (
	"fmt"
	"net/http"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var storywriterPath string

var _ = BeforeSuite(func() {
	var err error
	storywriterPath, err = gexec.Build("github.com/williammartin/storywriter")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(time.Second * 5)
	RunSpecs(t, "Storywriter Server Acceptance Suite")
}

func execBin(cmd *exec.Cmd) *gexec.Session {
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}

func ping() error {
	resp, err := http.Get("http://localhost:9000/ping")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	return fmt.Errorf("expected status code 200 but got %d", resp.StatusCode)
}
