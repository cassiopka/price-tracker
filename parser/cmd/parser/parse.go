package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewParseCmd() *cobra.Command {
	var (
		store    string
		category string
		format   string
		proxy    string
	)

	var parseCmd = &cobra.Command{
		Use:   "parse",
		Short: "Запуск парсинга товаров",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Парсинг запущен с параметрами:")
			fmt.Println("Магазин:", store)
			fmt.Println("Категория:", category)
			fmt.Println("Формат:", format)
			fmt.Println("Прокси:", proxy)
		},
	}

	// Флаги
	parseCmd.Flags().StringVarP(&store, "store", "s", "", "Магазин (sbermarket, okeydostavka, samokat, bristol)")
	parseCmd.Flags().StringVarP(&category, "category", "c", "", "Категория товаров")
	parseCmd.Flags().StringVarP(&format, "format", "f", "csv", "Формат вывода (csv, json, txt)")
	parseCmd.Flags().StringVarP(&proxy, "proxy", "p", "", "Адрес прокси-сервера")

	// Обязательные флаги
	parseCmd.MarkFlagRequired("store")
	parseCmd.MarkFlagRequired("category")

	return parseCmd
}
