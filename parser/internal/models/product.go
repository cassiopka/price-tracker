package models

// Product представляет собой товар из магазина.
type Product struct {
	Store                string  `json:"store"`                  // Название магазина
	Category             string  `json:"category"`               // Категория товара
	URL                  string  `json:"url,omitempty"`          // URL товара
	Name                 string  `json:"name"`                   // Название товара
	ImageURLSmall        string  `json:"image_url_small"`        // Ссылка на маленькое изображение товара
	ImageURLLarge        string  `json:"image_url_large"`        // Ссылка на большое изображение товара
	Price                string  `json:"price"`                  // Цена товара
	PriceWithoutDiscount float64 `json:"price_without_discount"` // Цена товара без скидки
}
