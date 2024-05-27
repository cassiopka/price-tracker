package parser

import (
	"context"
	"net/http"

	"parser/internal/models"
)

// Parser определяет интерфейс для парсеров магазинов.
type Parser interface {
	// ParseProducts парсит страницу категории и возвращает список товаров.
	ParseProducts(ctx context.Context, url string) ([]models.Product, error)

	// GetClient возвращает HTTP-клиент, настроенный для работы с магазином.
	GetClient() *http.Client
}
