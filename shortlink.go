package main

import "strings"

type shortlink struct {
	slug string
	link string
	base string
}

func New(link string) shortlink {
	return shortlink{
		slug: "asdf",
		link: link,
		base: "https://base/",
	}
}

func (s shortlink) String() string {
	return strings.TrimSuffix(s.base, "/") + "/" + s.slug
}

func (s shortlink) Target() string {
	return s.link
}
