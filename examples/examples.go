package main

import (
	"log"
	"os"
	"github.com/zachgoldstein/gotypeformio"
)

func main() {
	typeformKey := os.Getenv("TYPEFORM_KEY")
	api := gotypeformio.NewApi(typeformKey)
	api.DebugMode = true

	colors := gotypeformio.NewColors("#3D3D3D", "#4FB0AE", "#4FB0AE", "#FFFFFF")
	designBackground, err := api.CreateImage("http://i.imgur.com/EVOFpNF.png")
	design := gotypeformio.NewDesign(colors, "Bangers", designBackground.Id)
	createdDesign, err := api.CreateDesign(design)
	if err != nil {
		log.Printf("Could not create design err: %#v", err)
		return
	}

	shortTextField := gotypeformio.NewShortTextField(140, "A Simple Question", "What is your name?", true)
	longTextField := gotypeformio.NewLongTextField(140, "A Longer Question", "Tell me about your sassy self", true)

	fightAHorseDuck := gotypeformio.NewChoice(-1, "A horse sized duck")
	fightADuckHorse := gotypeformio.NewChoice(-1, "A thousand duck sized horses")
	fightNothing := gotypeformio.NewChoice(-1, "Nothing, we're all pals now and it's amazing")
	mcChoiceField := gotypeformio.NewMultipleChoiceField("A Question of self", "What would you fight?", true, []gotypeformio.Choice{*fightAHorseDuck, *fightADuckHorse, *fightNothing})

	horseDuckImg, err := api.CreateImage("http://i.imgur.com/t5IQH2o.jpg")
	fightAHorseDuck = gotypeformio.NewChoice(horseDuckImg.Id, "A horse sized duck")
	duckHorseImg, err := api.CreateImage("http://i.imgur.com/NYqW38B.jpg")
	fightADuckHorse = gotypeformio.NewChoice(duckHorseImg.Id, "A thousand duck sized horses")
	nothingImg, err := api.CreateImage("http://i.imgur.com/FWt2Y.jpg")
	fightNothing = gotypeformio.NewChoice(nothingImg.Id, "Nothing, we're all pals now and it's amazing")
	picChoiceField := gotypeformio.NewPictureChoiceField("A visual question of self", "What would you seriously fight?", true, []gotypeformio.Choice{*fightAHorseDuck, *fightADuckHorse, *fightNothing})

	statementField := gotypeformio.NewStatementField("Please take this super seriously", "They are life and death questions", true)

	holyGrailChoice := gotypeformio.NewChoice(-1, "To seek the holy grail")
	swallowChoice := gotypeformio.NewChoice(-1, "To seek an unladed swallow")
	dropdownField := gotypeformio.NewDropdownField("What is your quest?", "we all have one", true, []gotypeformio.Choice{*holyGrailChoice, *swallowChoice})

	yesNoField := gotypeformio.NewYesNoField("Would you ride a t-rex around if you could?", "Super fun, super dangerous....", true)

	numberField := gotypeformio.NewNumberField("What is the wind speed velocity of an unladen swallow?", "It's quite fast", true)

	opinionField := gotypeformio.NewOpinionField("On a scale of 1-bf, how well do we know each other?", "Aw come on", true, 5)

	ratingField := gotypeformio.NewRatingField("How awesome is this little library?", "Pretty good right?", true, 10)

	emailField := gotypeformio.NewEmailField("What email should I use to send you virtual glitter?", "ya it's a thing", true)

	websiteField := gotypeformio.NewWebsiteField("Any website you'd like to send strangers to?", "it's cool no worries", true)

	legalField := gotypeformio.NewLegalField("Say yes and I own your soul", "promise to do good things with it ;)", true)

	fields := []gotypeformio.FieldWrapper{
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

	formSubmission := gotypeformio.NewFormSubmission("A simple form", "123", map[string]string{}, createdDesign.Id, fields)

	resp, err := api.CreateForm(formSubmission)
	if err != nil {
		log.Printf("Could not create form err: %#v", err)
		return
	}

	test, ok := resp.Fields[5].(gotypeformio.YesNoField) //cast fields in the resp like so
	log.Println("form field 5 test ", test, " ok: ", ok)

	existingForm, err := api.GetForm(resp.ID)
	if err != nil {
		log.Printf("Could not create form err: %#v", err)
		return
	}
	log.Printf("\n retrieved form %v", existingForm)

}
