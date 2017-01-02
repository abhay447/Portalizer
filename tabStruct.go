package main
import (
    "os"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-webkit/webkit"
    . "./queryHandler"
)

type Tab struct{
    page *gtk.Frame
    /**/	        
}

func NewTab() *Tab{
    t := new(Tab)
    t.page = gtk.NewFrame("demo")
    vbox := gtk.NewVBox(false, 1)
    actionBar := gtk.NewHBox(false, 1)
    
    vbox.PackStart(actionBar, false, false, 0)
    
    //Action Bar Compenents
    Spindicator := gtk.NewSpinner()
	actionBar.PackStart(Spindicator, false, false, 30)
        
    //Query entry box
	entry := gtk.NewEntry()
	entry.SetText("http://google.com/")
    entry.SetWidthChars(150)
	actionBar.PackStart(entry, false, false, 0)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)

	webview := webkit.NewWebView()
	webview.Connect("load-committed", func() {
		entry.SetText(webview.GetUri())
	})
    webview.Connect("load-started", func() {
        Spindicator.Start()
	})
    webview.Connect("load-finished", func() {
        Spindicator.Stop()
	})
	swin.Add(webview)

	vbox.Add(swin)

	entry.Connect("activate", func() {
        webview.LoadUri(ActOnQuery(entry.GetText()))
	})
    t.page.Add(vbox)
    entry.Emit("activate")
    
    proxy := os.Getenv("HTTP_PROXY")
	if len(proxy) > 0 {
		soup_uri := webkit.SoupUri(proxy)
		webkit.GetDefaultSession().Set("proxy-uri", soup_uri)
		soup_uri.Free()
	}
    
    return t
}