/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cobra-cli/todo"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete that shi",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: delete,
}

func delete(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalln("Silinecek notun dizinini (indeksini) belirtmelisiniz.")
	}

	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalln("Notlar okunurken bir hata oluştu:", err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "geçerli bir dizin (indeks) değil. Lütfen bir sayı girin.", err)
	}

	if i <= 0 || i > len(items) {
		log.Fatalln(i, "geçerli bir dizin değil. Listenizde", len(items), "adet not var.")
	}

	items = append(items[:i-1], items[i:]...)

	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Fatalln("Notlar kaydedilirken bir hata oluştu:", err)
	}

	fmt.Printf("Not #%d silindi.\n", i)
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
