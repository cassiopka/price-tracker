package sbermarket

import (
	"context"
	"fmt"
	"net/url"

	"parser/internal/models"
	"parser/internal/parser"
)

// SberMarketParser представляет собой парсер для магазина SberMarket.
type SberMarketParser struct {
	parser.BaseParser
}

// NewSberMarketParser создает новый экземпляр SberMarketParser.
func NewSberMarketParser(proxyURL *url.URL) *SberMarketParser {
	return &SberMarketParser{
		BaseParser: *parser.NewBaseParser(proxyURL),
	}
}

// ParseProducts парсит страницу категории и возвращает список товаров.
func (p *SberMarketParser) ParseProducts(ctx context.Context, url string) ([]models.Product, error) {
	// TODO: Реализация парсинга продуктов SberMarket
	return nil, fmt.Errorf("парсинг для SberMarket еще не реализован")
}
