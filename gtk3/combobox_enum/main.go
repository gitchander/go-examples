package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	w, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	w.SetTitle("Example ComboBox")
	w.Connect("destroy", func() {
		gtk.MainQuit()
	})

	ps := []enumPair{
		{id: 100, name: "alpha"},
		{id: 200, name: "beta"},
		{id: 300, name: "gamma"},
		{id: 400, name: "zetta"},
	}

	cb, err := makeComboBoxByEnums(ps)
	checkError(err)
	w.Add(cb)

	w.ShowAll()

	gtk.Main()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type enumPair struct {
	id   int
	name string
}

func makeComboBoxByEnums(ps []enumPair) (*gtk.ComboBox, error) {

	ls, err := gtk.ListStoreNew(glib.TYPE_INT, glib.TYPE_STRING)
	if err != nil {
		return nil, err
	}

	columns := []int{0, 1}
	for _, p := range ps {
		iter := ls.Append()
		values := []interface{}{
			p.id,
			p.name,
		}
		err = ls.Set(iter, columns, values)
		if err != nil {
			return nil, err
		}
	}

	cb, err := gtk.ComboBoxNew()
	if err != nil {
		return nil, err
	}

	cb.SetModel(ls)
	cb.SetActive(0)

	cell, err := gtk.CellRendererTextNew()
	if err != nil {
		return nil, err
	}

	cb.PackStart(cell, true)
	cb.AddAttribute(cell, "text", 1) // column index = 1

	cb.Connect("changed", func() {

		iter, err := cb.GetActiveIter()
		checkError(err)

		gv, err := ls.GetValue(iter, 0)
		checkError(err)

		v, err := gv.GoValue()
		checkError(err)

		enumID, ok := v.(int)
		if ok {
			fmt.Println("id:", enumID)
		}
	})

	return cb, nil
}
