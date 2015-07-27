package typeform

var TYPEFORM_API string = "https://api.typeform.io/v0.3/"

type API struct {
	AuthToken string
}

func NewApi(authToken string) *API {
	return &API{
		AuthToken: authToken,
	}
}

func (a *API) CreateForm(formSubmission *FormSubmission) *FormSubmissionResponse {
	return Form.Create(a.AuthToken, formSubmission)
}

func (a *API) GetForm() {
	Form.Get(a.AuthToken)
}

func (a *API) CreateImage() {

}

func (a *API) CreateDesign() {

}
