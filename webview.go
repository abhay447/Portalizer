package main

import (
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Portalizer")
	window.Connect("destroy", gtk.MainQuit)
	
	notebook := gtk.NewNotebook()
	for i:=0;i<4;i++{
		t := NewTab()
		notebook.Add(t.page)
	}
	window.Add(notebook)
	window.SetSizeRequest(600, 600)
	window.ShowAll()

	gtk.Main()
}