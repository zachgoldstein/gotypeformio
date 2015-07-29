package typeform

import "fmt"

var TYPEFORM_API string = "https://api.typeform.io/v0.3/"

type API struct {
	AuthToken string
	DebugMode bool
}

func NewApi(authToken string) *API {
	return &API{
		AuthToken: authToken,
	}
}

type RequestError struct {
	StatusCode   int
	Urlstr       string
	ErrorMessage string
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("A error with status code %v occurred when issueing a request to %v: %v ", e.StatusCode, e.Urlstr, e.ErrorMessage)
}

func NewRequestError(statusCode int, urlStr string, errorMsg string) *RequestError {
	return &RequestError{
		StatusCode:   statusCode,
		Urlstr:       urlStr,
		ErrorMessage: errorMsg,
	}
}

func (a *API) CreateForm(formSubmission *FormSubmission) (resp *FormSubmissionResponse, err error) {
	return CreateForm(a.AuthToken, a.DebugMode, formSubmission)
}

func (a *API) GetForm() {
	GetForm(a.AuthToken)
}

func (a *API) CreateImage(urlstr string) (img *Image, err error) {
	return CreateImage(a.AuthToken, a.DebugMode, urlstr)

}

func (a *API) CreateDesign(design *Design) (resp *Design, err error) {
	return CreateDesign(a.AuthToken, a.DebugMode, design)
}
