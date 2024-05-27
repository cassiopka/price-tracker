package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewParseCmd() *cobra.Command {
	var parseCmd = &cobra.Command{
		Use:   "parse",
		Short: "Запуск парсинга товаров",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Парсинг запущен...")
		},
	}
	return parseCmd
}
