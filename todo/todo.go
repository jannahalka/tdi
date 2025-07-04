package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
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
		doneText = "\u274C"
	} else {
		doneText = "\u2705"
	}
	return fmt.Sprint(strconv.Itoa(i.Id) + "\t" + i.Text + "\t" + doneText + "\t" + strconv.Itoa(i.Priority))
}

func (i *Item) SetPriority(pri int) {
	i.Priority = pri
}

func (i *Item) SetToDone() {
	i.Done = true
}

func (i Item) ToRow() []string {

	row := []string{strconv.Itoa(i.Id), i.Text, strconv.Itoa(i.Priority), strconv.FormatBool(i.Done)}
	return row
}

func DisplayTodos(items []Item) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	defer w.Flush()

	fmt.Fprintln(w, "id\ttask\tstatus\tpriority")
	fmt.Fprintln(w, "--\t-------------------------------------\t---\t---")

	for _, i := range items {
		fmt.Fprintln(w, i)
	}
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

type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].Id < s[j].Id
	}
	return s[i].Priority < s[j].Priority
}
