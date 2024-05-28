package service

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
)

type Info struct {
	CurrentLang string // En
	ToLang      string // Es
	Text        string
}

//https://translate.google.com/?hl=es&sl=es&tl=en&text=
func TranslatorService(info Info) (string, error) {
	url := fmt.Sprintf("https://translate.google.com/?hl=%s&sl=%s&tl=%s&text=%s",
		info.CurrentLang, info.CurrentLang, info.ToLang, info.Text)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	return "", err
	//}
	//fmt.Println(string(body))

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(doc)
	var f func(*html.Node) string
	f = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "ryNqvb" {
					//fmt.Println(n.FirstChild.Data)
					return n.FirstChild.Data
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
		return ""
	}

	return f(doc), nil
}
