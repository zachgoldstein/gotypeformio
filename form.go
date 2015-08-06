package gotypeformio

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func CreateForm(authToken string, debug bool, formSubmission *FormSubmission) (*FormSubmissionResponse, error) {
	rawSubmission, err := json.Marshal(formSubmission)
	if err != nil {
		return &FormSubmissionResponse{}, err
	}

	buf := bytes.NewBuffer(rawSubmission)

	req, err := http.NewRequest("POST", TYPEFORM_API+"forms", buf)
	if err != nil {
		return &FormSubmissionResponse{}, err
	}

	req.Header.Add("x-api-token", authToken)

	if debug {
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Printf("DEBUG - Could not dump request: \n %v \n", err)
		} else {
			log.Printf("DEBUG - Full request: \n %v \n", string(dump))
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &FormSubmissionResponse{}, err
	}

	if debug {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Printf("DEBUG - Could not dump response: \n %v \n", err)
		} else {
			log.Printf("DEBUG - Full response: \n %v \n", string(dump))
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &FormSubmissionResponse{}, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		err = NewRequestError(res.StatusCode, res.Request.URL.String(), string(body))
		return &FormSubmissionResponse{}, err
	}

	var resp FormSubmissionResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return &FormSubmissionResponse{}, err
	}

	return &resp, nil
}

func GetForm(authToken string, debug bool, id string) (*FormSubmissionResponse, error) {
	req, err := http.NewRequest("Get", TYPEFORM_API+"forms/"+id, nil)
	if err != nil {
		return &FormSubmissionResponse{}, err
	}

	req.Header.Add("x-api-token", authToken)

	if debug {
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Printf("DEBUG - Could not dump request: \n %v \n", err)
		} else {
			log.Printf("DEBUG - Full request: \n %v \n", string(dump))
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &FormSubmissionResponse{}, err
	}

	if debug {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Printf("DEBUG - Could not dump response: \n %v \n", err)
		} else {
			log.Printf("DEBUG - Full response: \n %v \n", string(dump))
		}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &FormSubmissionResponse{}, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		err = NewRequestError(res.StatusCode, res.Request.URL.String(), string(body))
		return &FormSubmissionResponse{}, err
	}

	var resp FormSubmissionResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return &FormSubmissionResponse{}, err
	}

	return &resp, nil

}

func NewFormSubmission(title string, webhook string, queryStrings map[string]string, designId int, fields []FieldWrapper) *FormSubmission {
	//TODO encode querystrings here

	formSubmission := &FormSubmission{
		Fields:           fields,
		WebhookSubmitURL: webhook,
		Title:            title,
	}

	if designId > -1 {
		formSubmission.DesignId = designId
	}

	if len(queryStrings) > 0 {

	}

	return formSubmission
}

type FormSubmission struct {
	Fields           []FieldWrapper `json:"fields"`
	Title            string         `json:"title"`
	DesignId         int            `json:"design_id,omitempty"`
	WebhookSubmitURL string         `json:"webhook_submit_url"`
}

func NewChoice(imageId int, label string) *Choice {
	choice := &Choice{Label: label}
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
	FieldsRaw        []json.RawMessage `json:"fields"`
	Fields           []FieldWrapper    `json:"-"`
	ID               string            `json:"id"`
	Links            []Link            `json:"links"`
	Title            string            `json:"title"`
	WebhookSubmitURL string            `json:"webhook_submit_url"`
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
			if err == nil {
				t.Fields = append(t.Fields, txtField)
				continue
			} else {
				return err
			}
		case "statment":
			var statementField StatementField
			err = json.Unmarshal(rawMessage, &statementField)
			if err == nil {
				t.Fields = append(t.Fields, statementField)
				continue
			} else {
				return err
			}
		case "multiple_choice":
			var mcChoiceField MultipleChoiceField
			err = json.Unmarshal(rawMessage, &mcChoiceField)
			if err == nil {
				t.Fields = append(t.Fields, mcChoiceField)
				continue
			} else {
				return err
			}
		case "picture_choice":
			var picChoiceField PictureChoiceField
			err = json.Unmarshal(rawMessage, &picChoiceField)
			if err == nil {
				t.Fields = append(t.Fields, picChoiceField)
				continue
			} else {
				return err
			}
		case "dropdown":
			var dropdownField DropdownField
			err = json.Unmarshal(rawMessage, &dropdownField)
			if err == nil {
				t.Fields = append(t.Fields, dropdownField)
				continue
			} else {
				return err
			}
		case "yes_no":
			var yesnoField YesNoField
			err = json.Unmarshal(rawMessage, &yesnoField)
			if err == nil {
				t.Fields = append(t.Fields, yesnoField)
				continue
			} else {
				return err
			}
		case "number":
			var numField NumberField
			err = json.Unmarshal(rawMessage, &numField)
			if err == nil {
				t.Fields = append(t.Fields, numField)
				continue
			} else {
				return err
			}
		case "rating":
			var rateField RatingField
			err = json.Unmarshal(rawMessage, &rateField)
			if err == nil {
				t.Fields = append(t.Fields, rateField)
				continue
			} else {
				return err
			}
		case "opinion_scale":
			var opinionField OpinionField
			err = json.Unmarshal(rawMessage, &opinionField)
			if err == nil {
				t.Fields = append(t.Fields, opinionField)
				continue
			} else {
				return err
			}
		case "email":
			var emailField EmailField
			err = json.Unmarshal(rawMessage, &emailField)
			if err == nil {
				t.Fields = append(t.Fields, emailField)
				continue
			} else {
				return err
			}
		case "website":
			var websiteField WebsiteField
			err = json.Unmarshal(rawMessage, &websiteField)
			if err == nil {
				t.Fields = append(t.Fields, websiteField)
				continue
			} else {
				return err
			}
		case "legal":
			var legalField LegalField
			err = json.Unmarshal(rawMessage, &legalField)
			if err == nil {
				t.Fields = append(t.Fields, legalField)
				continue
			} else {
				return err
			}
		}
	}

	return nil
}
