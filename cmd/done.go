// TODO maddelerini tamamlandı olarak işaretlemek için `done` komutu.
package cmd

import (
	"cobra-cli/todo"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"sort"
	"strconv"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark item as done",
	Run:     donerun,
}

func donerun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label", err)
	}
	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked as done")

		sort.Sort(todo.ByPriority(items))
		err := todo.SaveItems(dataFile, items)
		if err != nil {
			return
		}
	} else {
		log.Println(i, "doesnt match any item")
	}

}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
