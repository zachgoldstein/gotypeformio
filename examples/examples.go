package main
import "os"

func main () {
	typeformKey := os.Getenv("TYPEFORM_KEY")
	api := typeform.NewApi()

	multipleChoiceField := NewShortTextField(140, "A Simple Question", "What is your name?", true)

	formSubmission := NewFormSubmission("A simple form", "", []Field{multipleChoiceField} )

	api.CreateForm( formSubmission )
}