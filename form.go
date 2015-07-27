package typeform

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bytes"
)

type Form struct {

}

func (f *Form) Create(authToken string, formSubmission *FormSubmission) (resp *FormSubmissionResponse, err error) {
	rawSubmission, err := json.Marshal(formSubmission)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	buf := bytes.NewBuffer(rawSubmission)

	req, err := http.NewRequest("POST", TYPEFORM_API, buf)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	req.Header.Add("x-api-token", authToken)

	res, err := http.DefaultClient.Do(req)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	var resp FormSubmissionResponse

	err = json.Unmarshal(body, &resp)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	return resp, nil
}

func (f *Form) Get() {

}

func NewFormSubmission (title string, webhook string, fields []Field) *FormSubmission {
	return &FormSubmission{
		Fields : fields,
		WebhookSubmitURL : webhook,
		Title : title,
	}
}

type FormSubmission struct {
	Fields []Field `json:"fields"`
	Title            string `json:"title"`
	WebhookSubmitURL string `json:"webhook_submit_url"`
}

func NewField(choices []Choice, description string, question string, required bool, fieldType string) *Field {
	return &Field{
		Choices : choices,
		Description : description,
		Question : question,
		Required : required,
		Type : fieldType,
	}
}

type Field struct {
	Choices []Choice `json:"choices"`
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
}

func NewChoice(imageId int, label string) *Choice {
	return &Choice{
		ImageID: imageId,
		Label: label,
	}
}

type Choice struct {
	ImageID int    `json:"image_id"`
	Label   string `json:"label"`
}

type FormSubmissionResponse struct {
	Fields []struct {
		AllowMultipleSelections bool `json:"allow_multiple_selections"`
		Choices                 []struct {
			Filename string `json:"filename"`
			Height   int    `json:"height"`
			ImageID  int    `json:"image_id"`
			Label    string `json:"label"`
			Width    int    `json:"width"`
		} `json:"choices"`
		Description string `json:"description"`
		ID          int    `json:"id"`
		Labels      bool   `json:"labels"`
		Question    string `json:"question"`
		Required    bool   `json:"required"`
		Type        string `json:"type"`
	} `json:"fields"`
	ID    string `json:"id"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Title            string `json:"title"`
	WebhookSubmitURL string `json:"webhook_submit_url"`
}
