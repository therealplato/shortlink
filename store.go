package main

import "errors"

type store struct{}

func (s store) Lookup(slug string) shortlink {
	return shortlink{
		slug: "asdf",
		link: "hjkl",
		base: "http://localhost:8000",
	}
}

func (s store) Save(sl shortlink) error {
	return errors.New("unimplemented")
}
