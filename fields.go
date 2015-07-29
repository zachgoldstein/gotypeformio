package typeform

import ()

type FieldWrapper interface{}

func NewField(question string, description string, required bool, fieldType string) *Field {
	return &Field{
		Description: description,
		Question:    question,
		Required:    required,
		Type:        fieldType,
	}
}

type Field struct {
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
}

type TextField struct {
	Description   string `json:"description"`
	Question      string `json:"question"`
	Required      bool   `json:"required"`
	Type          string `json:"type"`
	Id            int    `json:"id,omitempty"`
	MaxCharacters int    `json:"max_characters"`
}

func NewTextField(maxCharacters int, question string, description string, required bool, fieldType string) *TextField {
	return &TextField{
		MaxCharacters: maxCharacters,
		Description:   description,
		Question:      question,
		Required:      required,
		Type:          fieldType,
	}
}

func NewLongTextField(maxCharacters int, question string, description string, required bool) *TextField {
	return &TextField{
		MaxCharacters: maxCharacters,
		Description:   description,
		Question:      question,
		Required:      required,
		Type:          "long_text",
	}

}

func NewShortTextField(maxCharacters int, question string, description string, required bool) *TextField {
	return &TextField{
		MaxCharacters: maxCharacters,
		Description:   description,
		Question:      question,
		Required:      required,
		Type:          "short_text",
	}
}

type StatementField struct {
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
}

func NewStatementField(question string, description string, required bool) *StatementField {
	return &StatementField{
		Description: description,
		Question:    question,
		Required:    required,
		Type:        "statement",
	}
}

type MultipleChoiceField struct {
	Description string   `json:"description"`
	Question    string   `json:"question"`
	Required    bool     `json:"required"`
	Id          int      `json:"id,omitempty"`
	Type        string   `json:"type"`
	Choices     []Choice `json:"choices"`
}

func NewMultipleChoiceField(question string, description string, required bool, choices []Choice) *MultipleChoiceField {
	return &MultipleChoiceField{
		Description: description,
		Question:    question,
		Required:    required,
		Choices:     choices,
		Type:        "multiple_choice",
	}
}

type PictureChoiceField struct {
	Description string   `json:"description"`
	Question    string   `json:"question"`
	Required    bool     `json:"required"`
	Id          int      `json:"id,omitempty"`
	Type        string   `json:"type"`
	Choices     []Choice `json:"choices"`
}

func NewPictureChoiceField(question string, description string, required bool, choices []Choice) *PictureChoiceField {
	return &PictureChoiceField{
		Description: description,
		Question:    question,
		Required:    required,
		Choices:     choices,
		Type:        "picture_choice",
	}
}

type DropdownField struct {
	Description string   `json:"description"`
	Question    string   `json:"question"`
	Required    bool     `json:"required"`
	Id          int      `json:"id,omitempty"`
	Type        string   `json:"type"`
	Choices     []Choice `json:"choices"`
}

func NewDropdownField(question string, description string, required bool, choices []Choice) *DropdownField {
	return &DropdownField{
		Description: description,
		Question:    question,
		Required:    required,
		Choices:     choices,
		Type:        "dropdown",
	}
}

type YesNoField struct {
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
}

func NewYesNoField(question string, description string, required bool) *YesNoField {
	return &YesNoField{
		Description: description,
		Question:    question,
		Required:    required,
		Type:        "yes_no",
	}
}

type NumberField struct {
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
}

func NewNumberField(question string, description string, required bool) *NumberField {
	return &NumberField{
		Description: description,
		Question:    question,
		Required:    required,
		Type:        "number",
	}
}

type RatingField struct {
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Steps       int    `json:"-"`
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
}

func NewRatingField(question string, description string, required bool, steps int) *RatingField {
	//TODO: add validation for steps 3-10?

	return &RatingField{
		Description: description,
		Question:    question,
		Required:    required,
		Steps:       steps,
		Type:        "rating",
	}
}

type OpinionField struct {
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Id          int    `json:"id,omitempty"`
	Steps       int    `json:"-"`
	Type        string `json:"type"`
}

func NewOpinionField(question string, description string, required bool, steps int) *OpinionField {
	//TODO: add validation for steps 5-11?

	return &OpinionField{
		Description: description,
		Question:    question,
		Required:    required,
		Steps:       steps,
		Type:        "opinion_scale",
	}
}

type EmailField struct {
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
}

func NewEmailField(question string, description string, required bool) *EmailField {
	return &EmailField{
		Description: description,
		Question:    question,
		Required:    required,
		Type:        "email",
	}
}

type WebsiteField struct {
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
}

func NewWebsiteField(question string, description string, required bool) *WebsiteField {
	return &WebsiteField{
		Description: description,
		Question:    question,
		Required:    required,
		Type:        "website",
	}
}

type LegalField struct {
	Description string `json:"description"`
	Question    string `json:"question"`
	Required    bool   `json:"required"`
	Id          int    `json:"id,omitempty"`
	Type        string `json:"type"`
}

func NewLegalField(question string, description string, required bool) *LegalField {
	return &LegalField{
		Description: description,
		Question:    question,
		Required:    required,
		Type:        "legal",
	}
}
