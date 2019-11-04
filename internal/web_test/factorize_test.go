package web_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/yetialex/factorization/internal/evaluate"
	"github.com/yetialex/factorization/internal/web"

	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Check service handlers", func() {

	var r *mux.Router
	var w *httptest.ResponseRecorder

	handlerFactorize := factorizeHandler()

	pathPost := "/factorize"
	pathGet := "/factorize/{number}"

	BeforeEach(func() {
		r = mux.NewRouter()
	})
	Describe("Post numbers", func() {

		Context("Provide a valid object", func() {
			It("should send via POST valid JSON and get HTTP Status: 200", func() {
				commandJson := `{"number":"5555"}`
				reqUrl := "/factorize"
				testCommand(r, w, handlerFactorize, "POST", pathPost, reqUrl, commandJson, 200, true, "5555 = 5 * 11 * 101")
			})

			It("should send valid GET parameter and get HTTP Status: 200", func() {
				commandJson := ``
				reqUrl := "/factorize/5555"
				testCommand(r, w, handlerFactorize, "GET", pathGet, reqUrl, commandJson, 200, true, "5555 = 5 * 11 * 101")
			})
		})

		Context("Provide invalid numbers via GET", func() {
			It("should send invalid GET parameter and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/factorize/0"
				testCommand(r, w, handlerFactorize, "GET", pathGet, reqUrl, commandJson, 400, false, evaluate.ErrWrongNumber.Error())
			})

			It("should send invalid GET parameter and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/factorize/-10"
				testCommand(r, w, handlerFactorize, "GET", pathGet, reqUrl, commandJson, 400, false, evaluate.ErrWrongNumber.Error())
			})

			It("should send invalid GET parameter and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/factorize/5.5"
				testCommand(r, w, handlerFactorize, "GET", pathGet, reqUrl, commandJson, 400, false, evaluate.ErrConvert.Error())
			})

			It("should send invalid GET parameter and get HTTP Status: 400", func() {
				commandJson := ``
				reqUrl := "/factorize/abcd"
				testCommand(r, w, handlerFactorize, "GET", pathGet, reqUrl, commandJson, 400, false, evaluate.ErrConvert.Error())
			})
		})

		Context("Provide invalid numbers via POST", func() {
			It("should send via POST empty JSON and get HTTP Status: 400", func() {
				commandJson := `{}`
				reqUrl := "/factorize"
				testCommand(r, w, handlerFactorize, "POST", pathPost, reqUrl, commandJson, 400, false, web.ErrInvalidRequestParams.Error())
			})

			It("should send via POST invalid JSON and get HTTP Status: 400", func() {
				commandJson := `{rtrytr`
				reqUrl := "/factorize"
				testCommand(r, w, handlerFactorize, "POST", pathPost, reqUrl, commandJson, 400, false, web.ErrInvalidRequestParams.Error())
			})

			It("should send via POST invalid JSON and get HTTP Status: 400", func() {
				commandJson := `{"number":567}`
				reqUrl := "/factorize"
				testCommand(r, w, handlerFactorize, "POST", pathPost, reqUrl, commandJson, 400, false, web.ErrInvalidRequestParams.Error())
			})

			It("should send via POST invalid JSON and get HTTP Status: 400", func() {
				commandJson := `{"number":"567.567"}`
				reqUrl := "/factorize"
				testCommand(r, w, handlerFactorize, "POST", pathPost, reqUrl, commandJson, 400, false, evaluate.ErrConvert.Error())
			})

			It("should send via POST invalid JSON and get HTTP Status: 400", func() {
				commandJson := `{"number":"0"}`
				reqUrl := "/factorize"
				testCommand(r, w, handlerFactorize, "POST", pathPost, reqUrl, commandJson, 400, false, evaluate.ErrWrongNumber.Error())
			})
		})

	})
})

func factorizeHandler() http.Handler {
	return http.HandlerFunc(web.FactorizeHandler)
}
