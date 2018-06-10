package main

type shortlink struct{}

func (s *shortlink) New() string {
	return "asdf"
}
