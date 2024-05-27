package parser

import (
	"net/http"
	"net/url"
	"time"
)

// BaseParser предоставляет базовые методы для всех парсеров.
type BaseParser struct {
	client *http.Client
}

// NewBaseParser создает новый базовый парсер с настроенным HTTP-клиентом.
func NewBaseParser(proxyURL *url.URL) *BaseParser {
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
	return &BaseParser{
		client: client,
	}
}

// GetClient возвращает HTTP-клиент.
func (p *BaseParser) GetClient() *http.Client {
	return p.client
}
