package main

type image struct {
	URL    string `json:"url"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type author struct {
	Name    string `json:"name"`
	IconURL string `json:"icon_url,omitempty"`
}

type field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type footer struct {
	Text string `json:"footer,omitempty"`
}

type embed struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Timestamp   string   `json:"timestamp,omitempty"`
	Color       int      `json:"color,omitempty"`
	Image       *image   `json:"image,omitempty"`
	Author      *author  `json:"author,omitempty"`
	Fields      []*field `json:"fields,omitempty"`
	Footer      *footer  `json:"footer,omitempty"`
}
