package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	Text string
	Priority int
	Done bool
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

func (i Item) ToRow() []string {
	row := []string{i.Text, strconv.Itoa(i.Priority)}
	return row
}

func SaveItems(filename string, items []Item) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	for _, i := range items {
		if err := w.Write(i.ToRow()); err != nil {
			return err
		}
	}
	defer w.Flush()
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
		fmt.Println(row)
		text := row[0]
		priority, err := strconv.Atoi(row[1])
		if err != nil {
			return []Item{}, err
		}

		items = append(items, Item{Text:text, Priority: priority})
	}

	return items, nil
}
