package main

import (
	"fmt"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
)

var qs = []*survey.Question{
	{
		Name:     "ticket_title",
		Prompt:   &survey.Input{Message: "Ticket title"},
		Validate: survey.Required,
	},
	{
		Name:     "ticket_url",
		Prompt:   &survey.Input{Message: "Ticket URL"},
		Validate: survey.Required,
	},
	{
		Name: "concept",
		Prompt: &survey.Select{

			Message: "Choose one:",
			Options: []string{"Feat", "Fix"},
			Default: "Feat",
		},
	},
}

// the answers will be written to this struct

func main() {
	answers := struct {
		TicketTitle string `survey:"ticket_title"`
		TicketUrl   string `survey:"ticket_url"`
		Concept     string
	}{}
	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r, _ := regexp.Compile(`TYBA-\d+`)
	ticketNumber := r.FindString(answers.TicketUrl)
	fmt.Printf("- %s: %s [[%s](https://starkmvp.atlassian.net/browse/%s)]\n", answers.Concept, answers.TicketTitle, ticketNumber, ticketNumber)
}
