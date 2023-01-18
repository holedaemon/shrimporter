package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type payload struct {
	Content   string   `json:"content"`
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Embeds    []*embed `json:"embeds"`
}

func embedToEmbed(me *exportedMessageEmbed) *embed {
	em := &embed{}

	if len(me.Title) > 0 {
		em.Title = me.Title
	}

	if len(me.Description) > 0 {
		em.Description = me.Description
	}

	if me.Color != "" {
		clean := strings.TrimPrefix(me.Color, "#")
		hex, _ := strconv.ParseInt(clean, 16, 64)
		em.Color = int(hex)
	}

	if !me.Timestamp.IsZero() {
		em.Timestamp = me.Timestamp.Format(time.RFC3339)
	}

	if me.Image != nil && me.Image.URL != "" {
		img := &image{
			URL: me.Image.URL,
		}

		if me.Image.Height > 0 {
			img.Height = me.Image.Height
		}

		if me.Image.Width > 0 {
			img.Width = me.Image.Width
		}

		em.Image = img
	}

	if me.Author != nil {
		a := &author{}

		if me.Author.Name != "" {
			a.Name = me.Author.Name
		} else {
			a.Name = "Original Author Unknown"
		}

		if me.Author.IconURL != "" {
			a.IconURL = me.Author.IconURL
		}

		em.Author = a
	}

	if me.Footer != nil && me.Footer.Text != "" {
		em.Footer = &footer{
			Text: me.Footer.Text,
		}
	}

	if len(me.Fields) > 0 {
		em.Fields = make([]*field, len(me.Fields))

		for i := 0; i < len(me.Fields); i++ {
			em.Fields[i] = fieldToField(me.Fields[i])
		}
	}

	return em
}

func fieldToField(f *exportedMessageEmbedField) *field {
	nf := &field{}

	if f.Name != "" {
		nf.Name = f.Name
	}

	if f.Value != "" {
		nf.Value = f.Value
	}

	nf.Inline = f.IsInline
	return nf
}

func exportToPayload(em *exportedMessage) *payload {
	pl := &payload{}

	if em.Author != nil {
		if len(em.Author.Name) > 0 {
			pl.Username = em.Author.Name
		}

		if len(em.Author.AvatarURL) > 0 {
			pl.AvatarURL = em.Author.AvatarURL
		}
	} else {
		pl.Username = "Original Author Unknown"
	}

	if len(em.Content) > 0 {
		pl.Content = em.Content
	}

	if len(em.Embeds) > 0 {
		pl.Embeds = make([]*embed, len(em.Embeds))
		for i := 0; i < len(em.Embeds); i++ {
			pl.Embeds[i] = embedToEmbed(em.Embeds[i])
		}
	}

	return pl
}

var (
	webhookURL = flag.String("w", "", "Discord Webhook URL")
	exportFile = flag.String("f", "export.json", "JSON file containing exported messages")
)

func main() {
	flag.Parse()

	if *webhookURL == "" {
		log.Fatalln("A webhook URL is required.")
	}

	file, err := os.Open(*exportFile)
	if err != nil {
		log.Fatalf("Error opening export file: %s\n", err.Error())
	}

	var export *export
	if err := json.NewDecoder(file).Decode(&export); err != nil {
		log.Fatalf("Error decoding export file: %s\n", err.Error())
	}

	var buf bytes.Buffer
	for i := 0; i < len(export.Messages); i++ {
		pay := exportToPayload(export.Messages[i])

		buf.Reset()
		err := json.NewEncoder(&buf).Encode(pay)
		if err != nil {
			log.Fatalf("Error encoding webhook payload: %s\n", err.Error())
		}

		res, err := http.Post(*webhookURL, "application/json", &buf)
		if err != nil {
			log.Fatalf("Error POSTing webhook: %s\n", err.Error())
		}

		defer res.Body.Close()

		left := res.Header.Get("X-RateLimit-Remaining")
		if left == "0" {
			ru := res.Header.Get("X-RateLimit-Reset")
			reset, err := strconv.ParseInt(ru, 10, 64)
			if err != nil {
				log.Fatalf("Error converting X-RateLimit-Reset header into an integer: %s\n", err.Error())
			}

			rt := time.Unix(reset, 0)
			until := time.Until(rt)
			time.Sleep(until)
		}
	}
}
