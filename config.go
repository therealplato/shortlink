package main

import (
	"log"
	"net/url"
	"os"
	"regexp"
)

type config struct {
	HealthcheckAddr string
	ShortlinkAddr   string
	BaseURL         string
	PostgresURL     string
}

func MustLoadConfig() config {
	cfg := config{
		HealthcheckAddr: os.Getenv("HEALTHCHECK_LISTEN_ADDR"),
		ShortlinkAddr:   os.Getenv("SHORTLINK_LISTEN_ADDR"),
		BaseURL:         os.Getenv("BASE_URL"),
		PostgresURL:     os.Getenv("POSTGRES_URL"),
	}
	if cfg.HealthcheckAddr == "" {
		log.Fatal("HEALTHCHECK_LISTEN_ADDR must be set")
	}
	if cfg.ShortlinkAddr == "" {
		log.Fatal("SHORTLINK_LISTEN_ADDR must be set")
	}
	if cfg.PostgresURL == "" {
		log.Fatal("POSTGRES_URL must be set")
	}
	trailingSlash := regexp.MustCompile("/$")
	_, err := url.Parse(cfg.BaseURL)
	if cfg.BaseURL == "" || err != nil || !trailingSlash.MatchString(cfg.BaseURL) {
		log.Fatal("BASE_URL must be set, valid, and have a trailing slash")
	}
	return cfg
}
