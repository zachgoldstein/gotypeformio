package main

import (
	"github.com/zachgoldstein/go-typeform"
	"log"
	"os"
)

func main() {
	typeformKey := os.Getenv("TYPEFORM_KEY")
	api := typeform.NewApi(typeformKey)
	api.DebugMode = true

	colors := typeform.NewColors("#3D3D3D", "#4FB0AE", "#4FB0AE", "#FFFFFF")
	designBackground, err := api.CreateImage("http://i.imgur.com/EVOFpNF.png")
	design := typeform.NewDesign(colors, "Bangers", designBackground.Id)
	createdDesign, err := api.CreateDesign(design)
	if err != nil {
		log.Printf("Could not create design err: %#v", err)
		return
	}

	shortTextField := typeform.NewShortTextField(140, "A Simple Question", "What is your name?", true)
	longTextField := typeform.NewLongTextField(140, "A Longer Question", "Tell me about your sassy self", true)

	fightAHorseDuck := typeform.NewChoice(-1, "A horse sized duck")
	fightADuckHorse := typeform.NewChoice(-1, "A thousand duck sized horses")
	fightNothing := typeform.NewChoice(-1, "Nothing, we're all pals now and it's amazing")
	mcChoiceField := typeform.NewMultipleChoiceField("A Question of self", "What would you fight?", true, []typeform.Choice{*fightAHorseDuck, *fightADuckHorse, *fightNothing})

	horseDuckImg, err := api.CreateImage("http://i.imgur.com/t5IQH2o.jpg")
	fightAHorseDuck = typeform.NewChoice(horseDuckImg.Id, "A horse sized duck")
	duckHorseImg, err := api.CreateImage("http://i.imgur.com/NYqW38B.jpg")
	fightADuckHorse = typeform.NewChoice(duckHorseImg.Id, "A thousand duck sized horses")
	nothingImg, err := api.CreateImage("http://i.imgur.com/FWt2Y.jpg")
	fightNothing = typeform.NewChoice(nothingImg.Id, "Nothing, we're all pals now and it's amazing")
	picChoiceField := typeform.NewPictureChoiceField("A visual question of self", "What would you seriously fight?", true, []typeform.Choice{*fightAHorseDuck, *fightADuckHorse, *fightNothing})

	statementField := typeform.NewStatementField("Please take this super seriously", "They are life and death questions", true)

	holyGrailChoice := typeform.NewChoice(-1, "To seek the holy grail")
	swallowChoice := typeform.NewChoice(-1, "To seek an unladed swallow")
	dropdownField := typeform.NewDropdownField("What is your quest?", "we all have one", true, []typeform.Choice{*holyGrailChoice, *swallowChoice})

	yesNoField := typeform.NewYesNoField("Would you ride a t-rex around if you could?", "Super fun, super dangerous....", true)

	numberField := typeform.NewNumberField("What is the wind speed velocity of an unladen swallow?", "It's quite fast", true)

	opinionField := typeform.NewOpinionField("On a scale of 1-bf, how well do we know each other?", "Aw come on", true, 5)

	ratingField := typeform.NewRatingField("How awesome is this little library?", "Pretty good right?", true, 10)

	emailField := typeform.NewEmailField("What email should I use to send you virtual glitter?", "ya it's a thing", true)

	websiteField := typeform.NewWebsiteField("Any website you'd like to send strangers to?", "it's cool no worries", true)

	legalField := typeform.NewLegalField("Say yes and I own your soul", "promise to do good things with it ;)", true)

	fields := []typeform.FieldWrapper{
		shortTextField,
		longTextField,
		mcChoiceField,
		picChoiceField,
		statementField,
		dropdownField,
		yesNoField,
		numberField,
		opinionField,
		ratingField,
		emailField,
		websiteField,
		legalField,
	}

	formSubmission := typeform.NewFormSubmission("A simple form", "123", map[string]string{}, createdDesign.Id, fields)

	resp, err := api.CreateForm(formSubmission)
	if err != nil {
		log.Printf("Could not create form err: %#v", err)
		return
	}

	test, ok := resp.Fields[5].(typeform.YesNoField) //cast fields in the resp like so
	log.Println("form field 5 test ", test, " ok: ", ok)

	existingForm, err := api.GetForm(resp.ID)
	if err != nil {
		log.Printf("Could not create form err: %#v", err)
		return
	}
	log.Printf("\n retrieved form %v", existingForm)


}
