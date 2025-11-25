package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}
func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	for i, _ := range items {
		items[i].position = i + 1
	}
	return items, nil
}
func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

type ByPriority []Item

func (a ByPriority) Len() int      { return len(a) }
func (a ByPriority) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool {
	if a[i].Done != a[j].Done {
		return a[i].Done
	}
	if a[i].Priority != a[j].Priority {
		return a[i].Priority < a[j].Priority
	}
	return a[i].position < a[j].position
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}
