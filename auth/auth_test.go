package auth

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var (
		mockCtrl   *gomock.Controller
		httpClient *MockHTTPClient
		authClient *Client
		apiUrl     = "http://foo.bar"
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		httpClient = NewMockHTTPClient(mockCtrl)
		authClient = NewClient(
			httpClient,
			apiUrl,
			"azerty",
			"foo",
			"foo:bar:zoo",
		)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("generates headers", func() {
		headers := authClient.buildHeaders()
		Expect(headers["Accept"]).To(Equal("application/json"))
		Expect(headers["Content-Type"]).To(Equal("application/x-www-form-urlencoded"))
	})

	Context("When private key is valid", func() {
		It("generates request body", func() {
			body := authClient.buildParams()
			Expect(body["client_id"]).ToNot(BeEmpty())
			Expect(body["client_secret"]).ToNot(BeEmpty())
			Expect(body["scope"]).ToNot(BeEmpty())
			Expect(body["grant_type"]).ToNot(BeEmpty())
		})
	})
})
