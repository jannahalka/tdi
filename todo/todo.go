package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Item struct {
	Id       int
	Text     string
	Priority int
	Done     bool
}

func (i Item) String() string {
	var doneText string

	if i.Done == false {
		doneText = "uncomplete"
	} else {
		doneText = "complete"
	}
	return fmt.Sprint(strconv.Itoa(i.Id) + "\t" + strconv.Itoa(i.Priority) + "\t" + i.Text + "\t" + doneText)

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

func (i *Item) SetToDone() {
	i.Done = true
}

func (i Item) ToRow() []string {

	row := []string{strconv.Itoa(i.Id), i.Text, strconv.Itoa(i.Priority), strconv.FormatBool(i.Done)}
	return row
}

func SaveItems(filename string, items []Item) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, i := range items {
		if err := w.Write(i.ToRow()); err != nil {
			return err
		}
	}

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	var items []Item

	file, err := os.Open(filename)
	if err != nil {
		return []Item{}, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	data, err := r.ReadAll()

	if err != nil {
		return []Item{}, nil
	}

	for _, row := range data {
		text := row[1]

		priority, err := strconv.Atoi(row[2])
		if err != nil {
			return []Item{}, err
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			return []Item{}, err
		}

		done, err := strconv.ParseBool(row[3])

		if err != nil {
			return []Item{}, err
		}

		items = append(items, Item{Text: text, Priority: priority, Id: id, Done: done})
	}

	return items, nil
}

func FindItem(items []Item, itemId int) *Item {
	idx := slices.IndexFunc(items, func(i Item) bool { return i.Id == itemId })

	if idx == -1 {
		return nil
	}

	return &items[idx]
}
