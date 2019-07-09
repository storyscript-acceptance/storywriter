package acceptance_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Write Endpoint", func() {
	var (
		session *gexec.Session

		template string
		words    []string

		code int
		body string
	)

	BeforeEach(func() {
		template = ""
		words = []string{}

		storywriterCmd := exec.Command(storywriterPath)
		session = execBin(storywriterCmd)

		Eventually(ping).Should(Succeed())
	})

	AfterEach(func() {
		session.Kill().Wait()
	})

	When("a well-formed request is made", func() {
		JustBeforeEach(func() {
			code, body = writeStory(template, words)
		})

		When("a template is provided with no interpolation", func() {
			BeforeEach(func() {
				template = "This is a template with no interpolation"
			})

			It("responds with 200 OK and the template", func() {
				Expect(code).To(Equal(200))
				Expect(body).To(Equal(template))
			})
		})

		When("a template is provided with words to interpolate", func() {
			BeforeEach(func() {
				template = "This is a &1 with &2"
				words = []string{"template", "some interpolation"}
			})

			It("responds with 200 OK and the interpolated template", func() {
				Expect(code).To(Equal(200))
				Expect(body).To(Equal("This is a template with some interpolation"))
			})
		})
	})

	When("a malformed body is provided", func() {
		BeforeEach(func() {
			malformedBody := []byte("malformed")
			code, body = makeRequest(malformedBody)
		})

		It("returns an 400 Bad Request with an informative error", func() {
			Expect(code).To(Equal(400))
			Expect(body).To(Equal("please provide a well-formed body"))
		})
	})
})

func writeStory(template string, words []string) (int, string) {
	draft := draft{
		Template: template,
		Words:    words,
	}
	requestBody, err := json.Marshal(draft)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	return makeRequest(requestBody)
}

func makeRequest(requestBody []byte) (int, string) {
	resp, err := http.Post("http://localhost:9000/write", "application/json", bytes.NewBuffer(requestBody))
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	return resp.StatusCode, string(responseBody)
}

type draft struct {
	Template string   `json:"template"`
	Words    []string `json:"words"`
}
