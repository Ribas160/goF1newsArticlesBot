package main

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

const (
	XML_URL        string = "https://www.f1news.ru/export/news.xml"
	LOCAL_XML_FILE string = "articles.xml"
)

var (
	appDirectory = ""
)

type App struct {
	ErrorLog log.Logger
}

func main() {
	app := &App{
		ErrorLog: *log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	app.fillVars()

	f, err := os.OpenFile(path.Join(appDirectory, "runtime.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	defer f.Close()

	app.ErrorLog.SetOutput(f)

	if err := godotenv.Load(path.Join(appDirectory, ".env")); err != nil {
		app.ErrorLog.Fatal("error loading env variables: %s", err.Error())
	}

	newArticles := app.UpdateLocalXml()
	app.SubmitNewArticles(*newArticles)
}

func (a *App) fillVars() {
	_, filename, _, ok := runtime.Caller(1)
	if ok != true {
		a.ErrorLog.Fatal("No caller information")
	}

	appDirectory = path.Join(path.Dir(filename), "../")
}
