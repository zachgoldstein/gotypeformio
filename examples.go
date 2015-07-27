package typeform
import "os"

func main () {
	typeformKey := os.Getenv("TYPEFORM_KEY")
	api := NewApi(typeformKey)

	formSubmission := NewFormSubmission()

	api.CreateForm()
}