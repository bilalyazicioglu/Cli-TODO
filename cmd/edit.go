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
	"strings"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Editing the existing note",
	Long:  `Editing the existing note over an index`,
	Run:   editRun,
}

func editRun(cmd *cobra.Command, args []string) {
	indexStr := args[0]
	i, err := strconv.Atoi(indexStr)
	if err != nil {
		log.Fatalln(indexStr, "geçerli bir dizin (indeks) değil. Lütfen bir sayı girin.")
	}

	newText := strings.Join(args[1:], " ")

	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalln("Notlar okunurken bir hata oluştu:", err)
	}

	if i <= 0 || i > len(items) {
		log.Fatalln(i, "geçerli bir dizin değil. Listenizde", len(items), "adet not var.")
	}

	itemIndex := i - 1

	oldText := items[itemIndex].Text
	oldPriority := items[itemIndex].Priority

	items[itemIndex].Text = newText

	newPriority := oldPriority
	if cmd.Flags().Changed("priority") {
		items[itemIndex].SetPriority(priority)
		newPriority = priority
	}

	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Fatalln("Notlar kaydedilirken bir hata oluştu:", err)
	}

	fmt.Printf("Not #%d başarıyla düzenlendi.\n", i)
	fmt.Printf("  Eski Metin: \"%s\"\n", oldText)
	fmt.Printf("  Yeni Metin: \"%s\"\n", newText)
	if cmd.Flags().Changed("priority") {
		fmt.Printf("  Eski Öncelik: %d, Yeni Öncelik: %d\n", oldPriority, newPriority)
	}
}
func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
