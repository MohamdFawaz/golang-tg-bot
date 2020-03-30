package main

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
		token     = os.Getenv("TOKEN")      // you must add it to your config vars
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/help", func(m *tb.Message) {
		b.Send(m.Sender, "To Say Hello use /hello \n To pick your time of the day use /pick_time\n")
	})

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "You entered "+m.Text)
	})

	moonBtn := tb.InlineButton{
		Unique: "moon",
		Text:   "Moon 🌚",
	}

	sunBtn := tb.InlineButton{
		Unique: "sun",
		Text:   "Sun 🌞",
	}

	b.Handle(&moonBtn, func(c *tb.Callback) {
		// Required for proper work
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		// Send messages here
		b.Send(c.Sender, "Moon says 'Hi'!")
	})

	b.Handle(&sunBtn, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Sun says 'Hi'!")
	})

	inlineKeys := [][]tb.InlineButton{
		[]tb.InlineButton{sunBtn, moonBtn},
	}

	b.Handle("/pick_time", func(m *tb.Message) {
		b.Send(
			m.Sender,
			"Pick your time of the day",
			&tb.ReplyMarkup{InlineKeyboard: inlineKeys})
	})

	b.Start()
}
