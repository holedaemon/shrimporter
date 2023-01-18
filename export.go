package main

import "time"

type export struct {
	Messages []*exportedMessage `json:"messages"`
}

type exportedMessageAuthor struct {
	Name          string `json:"name,omitempty"`
	Discriminator string `json:"discriminator,omitempty"`
	AvatarURL     string `json:"avatarUrl,omitempty"`
}

type exportedMessageEmbedFooter struct {
	Text string `json:"text,omitempty"`
}

type exportedMessageEmbedField struct {
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
	IsInline bool   `json:"isInline,omitempty"`
}

type exportedMessageEmbedAuthor struct {
	Name    string `json:"name,omitempty"`
	IconURL string `json:"iconUrl,omitempty"`
}

type exportedMessageEmbed struct {
	Title       string                       `json:"title,omitempty"`
	Description string                       `json:"description,omitempty"`
	Timestamp   time.Time                    `json:"timestamp,omitempty"`
	Color       string                       `json:"color,omitempty"`
	Footer      *exportedMessageEmbedFooter  `json:"footer,omitempty"`
	Fields      []*exportedMessageEmbedField `json:"fields,omitempty"`
	Image       *exportedMessageEmbedImage   `json:"image,omitempty"`
	Author      *exportedMessageEmbedAuthor  `json:"author,omitempty"`
}

type exportedMessageEmbedImage struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type exportedMessage struct {
	Content string                  `json:"content,omitempty"`
	Embeds  []*exportedMessageEmbed `json:"embeds,omitempty"`
	Author  *exportedMessageAuthor  `json:"author,omitempty"`
}
