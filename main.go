package main

import (
	"embed"
	"io"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//go:embed resume.tmpl
var resume embed.FS

type info struct {
	Name      string
	Certs     string
	Skills    string
	Education string
	Jobs      string
	About     string
	Phone     string
	Email     string
	Final     string
}

func generateTemplate(data info) {
	tmpl := template.Must(template.ParseFS(resume, "resume.tmpl"))
	dirName, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to get users home directory")
	}
	t := time.Now().Day()
	filePath := dirName + "Documents/newResume" + strconv.Itoa(t) + ".txt"
	newDoc, openErr := os.Create(filePath)
	if openErr != nil {
		log.Fatal("Failed to create new resume, ", openErr.Error())
	}
	writer := io.Writer(newDoc)
	tmplErr := tmpl.Execute(writer, data)
	if tmplErr != nil {
		log.Fatal("Failed to execute template, ", tmplErr.Error())
	}
}

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()
	nm := tview.NewInputField().SetLabel("Name: ").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)
	crt := tview.NewInputField().SetLabel("Certs: ").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)
	skl := tview.NewInputField().SetLabel("Skills: ").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)
	edu := tview.NewInputField().SetLabel("Education: ").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)
	jbs := tview.NewInputField().SetLabel("Work history: ").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)
	abt := tview.NewInputField().SetLabel("About: ").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)
	phn := tview.NewInputField().SetLabel("Phone number: ").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)
	eml := tview.NewInputField().SetLabel("Email: ").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)
	fnl := tview.NewInputField().SetLabel("Final remarks: ").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)

	cmplt := tview.NewModal().SetText("Complete!").SetTextColor(tcell.ColorBlack).AddButtons([]string{"done"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonIndex == 0 {
			app.Stop()
		}
	})
	cmplt.SetBackgroundColor(tcell.ColorDarkViolet).SetBorder(true)
	//  pages.AddPage("cmplt", cmplt, true, true)

	form := tview.NewForm().
		AddFormItem(nm).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray).
		AddFormItem(abt).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray).
		AddFormItem(crt).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray).
		AddFormItem(skl).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray).
		AddFormItem(edu).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray).
		AddFormItem(jbs).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray).
		AddFormItem(phn).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray).
		AddFormItem(eml).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray).
		AddFormItem(fnl).SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorGray)

	pages.AddPage("form", form, true, true)
	pages.AddPage("cmplt", cmplt, true, false)

	form.AddButton("Process", func() {
		inf := info{
			Name:      nm.GetText(),
			Certs:     crt.GetText(),
			Skills:    skl.GetText(),
			Education: edu.GetText(),
			Jobs:      jbs.GetText(),
			About:     abt.GetText(),
			Phone:     phn.GetText(),
			Email:     eml.GetText(),
			Final:     fnl.GetText(),
		}
		generateTemplate(inf)
		pages.SwitchToPage("cmplt")
	}).AddButton("Quit", func() {
		app.Stop()
	})
	form.SetBackgroundColor(tcell.ColorDarkViolet).SetBorder(true).SetBorderColor(tcell.ColorDarkBlue)

	if err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).Run(); err != nil {
		log.Fatal("Fatal error running application ", err.Error())
	}
}
