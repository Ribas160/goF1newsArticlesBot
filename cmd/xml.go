package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type Articles struct {
	Articles []Article `xml:"channel>item"`
}

type Article struct {
	Title       string `xml:"title" json:"title"`
	Description string `xml:"description" json:"description"`
	Link        string `xml:"link" json:"link"`
	PubDate     string `xml:"pubDate" json:"pubDate"`
}

func (a *App) UpdateLocalXml() *Articles {
	f1newsXml := a.getXml()

	if _, err := os.Stat(path.Join(appDirectory, LOCAL_XML_FILE)); os.IsNotExist(err) {
		a.writeLocalXml(f1newsXml)

		return &Articles{}
	}

	fp, err := os.Open(LOCAL_XML_FILE)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	defer fp.Close()

	localXml, _ := ioutil.ReadAll(fp)

	newArticles := a.compareXml(f1newsXml, localXml)

	a.writeLocalXml(f1newsXml)

	return newArticles
}

func (a *App) getXml() []byte {
	resp, err := http.Get(XML_URL)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	xmlDocument, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	return xmlDocument
}

func (a *App) writeLocalXml(f1newsXml []byte) {
	err := os.WriteFile(path.Join(appDirectory, LOCAL_XML_FILE), f1newsXml, 0755)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}
}

func (a *App) compareXml(f1newsXml []byte, localXml []byte) *Articles {
	var f1newsArticles Articles
	var localArticles Articles
	var newArticles Articles

	xml.Unmarshal(f1newsXml, &f1newsArticles)
	xml.Unmarshal(localXml, &localArticles)

	for i := 0; i < len(f1newsArticles.Articles); i++ {
		new := true

		for j := 0; j < len(localArticles.Articles); j++ {
			if f1newsArticles.Articles[i].Link == localArticles.Articles[j].Link {
				new = false
			}
		}

		if new == true {
			newArticles.Articles = append(newArticles.Articles, f1newsArticles.Articles[i])
		}
	}

	return &newArticles
}
