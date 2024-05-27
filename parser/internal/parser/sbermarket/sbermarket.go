package sbermarket

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"net/url"
	"time"

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
	var products []models.Product

	// Создание нового контекста с отменой по таймауту
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Запуск Chrome в headless-режиме
	chromeCtx, cancelChrome := chromedp.NewContext(ctx)
	defer cancelChrome()

	// Навигация на страницу категории
	err := chromedp.Run(chromeCtx,
		chromedp.Navigate(url),
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка навигации на страницу: %w", err)
	}

	// Ожидание загрузки необходимых элементов страницы
	err = chromedp.Run(chromeCtx,
		chromedp.WaitVisible(`.CategoryFilters_linksContainer__muFrR`, chromedp.ByQuery),
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка ожидания загрузки страницы: %w", err)
	}

	// Извлечение данных о товарах
	var nodes []*cdp.Node
	err = chromedp.Run(chromeCtx,
		chromedp.Nodes(`.ProductCard_root__K6IZK`, &nodes, chromedp.ByQueryAll, chromedp.NodeVisible),
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка извлечения данных о товарах: %w", err)
	}

	// Заполнение данных о товарах
	for i := range nodes {
		chromedp.Run(chromeCtx,
			chromedp.Text(`.ProductCard_title__iNsaD`, &products[i].Name, chromedp.ByQuery, chromedp.NodeVisible),
			chromedp.Text(`.ProductCard_price__aRuGG`, &products[i].Price, chromedp.ByQuery, chromedp.NodeVisible),
			// ... (извлечение остальных данных о товаре)
		)
		products[i].Store = "SberMarket"
		products[i].Category = "Тестовая категория" // TODO: Извлекать актуальную категорию
	}

	return products, nil
}
