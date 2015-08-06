## gotypeformio

A go wrapper for the typeform io api

` go get github.com/zachgoldstein/gotypeformio `

### Usage:

See  for full docs.

[![GoDoc](https://godoc.org/github.com/zachgoldstein/gotypeformio?status.png)](https://godoc.org/github.com/zachgoldstein/gotypeformio)

Use the `API` object for a convenience object to issue all requests. Set the typeformio auth key once in the constructor, and then call all other functions.

```golang
typeformKey := os.Getenv("TYPEFORM_KEY")
api := gotypeformio.NewApi(typeformKey)

longTextField := gotypeformio.NewLongTextField(140, "A long Question", "Tell me about your sassy self", true)

formSubmission := gotypeformio.NewFormSubmission("A simple form", "123", map[string]string{}, createdDesign.Id, []gotypeformio.FieldWrapper{longTextField} )

```

There are convenience methods for creating all fields

```golang
NewShortTextField
NewLongTextField
NewMultipleChoiceField
NewPictureChoiceField
NewStatementField
NewDropdownField
NewYesNoField
NewNumberField
NewOpinionField
NewRatingField
NewEmailField
NewWebsiteField
NewLegalField
```

Some fields (like NewPictureChoiceField) have a variety of choices, which can be created via
```golang
//Image id is 125
horseDuckChoice := gotypeformio.NewChoice(125, "A horse sized duck")
```

When dealing with a created form, you can cast the fields back to their respective types:

```golang
formSubmission := gotypeformio.NewFormSubmission("A simple form", "123", map[string]string{}, createdDesign.Id, []gotypeformio.FieldWrapper{longTextField} )
txtField, ok := formSubmission.Fields[0].(gotypeformio.TextField)
```

### Tips:

For testing, you can use the debug flag on the API type to trigger complete req/resp logging via `httputil.DumpRequest` and `httputil.DumpResponse`

```golang
typeformKey := os.Getenv("TYPEFORM_KEY")
api := gotypeformio.NewApi(typeformKey)
api.DebugMode = true
img, err := api.CreateImage("http://i.imgur.com/cQVzmdw.jpg")
```

see [examples/example.go](https://github.com/zachgoldstein/gotypeformio/blob/master/examples/examples.go) for more information

