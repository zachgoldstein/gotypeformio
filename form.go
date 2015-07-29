package typeform

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bytes"
	"log"
	"net/http/httputil"
)

func CreateForm(authToken string, debug bool, formSubmission *FormSubmission) (*FormSubmissionResponse, error) {
	rawSubmission, err := json.Marshal(formSubmission)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	buf := bytes.NewBuffer(rawSubmission)

	req, err := http.NewRequest("POST", TYPEFORM_API+"forms", buf)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	req.Header.Add("x-api-token", authToken)

	if (debug) {
		dump, err := httputil.DumpRequest(req, true)
		if (err != nil) {
			log.Printf("DEBUG - Could not dump request: \n %v \n", err)
		} else {
			log.Printf("DEBUG - Full request: \n %v \n", string(dump))
		}
	}

	res, err := http.DefaultClient.Do(req)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	if (debug) {
		dump, err := httputil.DumpResponse(res, true)
		if (err != nil) {
			log.Printf("DEBUG - Could not dump response: \n %v \n", err)
		} else {
			log.Printf("DEBUG - Full response: \n %v \n", string(dump))
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	if (res.StatusCode != 200 && res.StatusCode != 201) {
		err = NewRequestError(res.StatusCode, res.Request.URL.String(), string(body))
		return &FormSubmissionResponse{}, err
	}

	var resp FormSubmissionResponse
	err = json.Unmarshal(body, &resp)
	if (err != nil) {
		return &FormSubmissionResponse{}, err
	}

	return &resp, nil
}

func GetForm(authToken string) {

}

func NewFormSubmission (title string, webhook string, queryStrings map[string]string, fields []FieldWrapper) *FormSubmission {
	//TODO encode querystrings here

	return &FormSubmission{
		Fields : fields,
		WebhookSubmitURL : webhook,
		Title : title,
	}
}

type FormSubmission struct {
	Fields []FieldWrapper `json:"fields"`
	Title            string `json:"title"`
	WebhookSubmitURL string `json:"webhook_submit_url"`
}

//ShortTextField
//LongTextField
//MultipleChoiceField
//PictureChoiceField
//StatementField
//DropdownField
//YesNoField
//NumberField
//RatingField
//OpinionScaleField
//EmailField
//WebsiteField
//LegalField


func NewChoice(imageId int, label string) *Choice {
	choice := &Choice{Label: label,}
	if imageId > 0 {
		choice.ImageID = imageId
	}
	return choice
}

type Choice struct {
	ImageID int    `json:"image_id,omitempty"`
	Label   string `json:"label"`
}

type FormSubmissionResponse struct {
	FieldsRaw []json.RawMessage `json:"fields"`
	Fields []FieldWrapper `json:"-"`
	ID    string `json:"id"`
	Links []Link `json:"links"`
	Title            string `json:"title"`
	WebhookSubmitURL string `json:"webhook_submit_url"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type formResp FormSubmissionResponse

func (t *FormSubmissionResponse) UnmarshalJSON(data []byte) error {

	resp := formResp{}
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	*t = FormSubmissionResponse(resp)

	for _, rawMessage := range t.FieldsRaw {
		var field Field
		err = json.Unmarshal(rawMessage, &field)
		switch field.Type {
			case "long_text", "short_text":
			var txtField TextField
			err = json.Unmarshal(rawMessage, &txtField)
			if (err == nil) {
				t.Fields = append(t.Fields, txtField)
				continue
			} else {
				return err
			}
		}
	}

	return nil
}