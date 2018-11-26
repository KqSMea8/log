package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

type Post struct {
	Title   string
	Date    string
	Summary string
	Body    string
	File    string
}

func handlerequest(w http.ResponseWriter, r *http.Request) {
	posts := getPosts()
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")
	t.Execute(w, posts)

	//tt := template.New("new")
	//tt,_ = tt.Parse(posts[0].Body)
	//
	//tt.Execute(w,nil)
}

func getPosts() []Post {
	a := []Post{}
	files, _ := filepath.Glob("posts/*")
	for _, f := range files {
		fmt.Println("f:", f)
		file := strings.Replace(f, "posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileread, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		date := string(lines[1])
		summary := string(lines[2])
		body := strings.Join(lines[3:len(lines)], "\n")
		body = string(blackfriday.MarkdownCommon([]byte(body)))
		a = append(a, Post{title, date, summary, body, file})
	}
	return a
}

func main() {
	http.HandleFunc("/", handlerequest)
	http.ListenAndServe(":8000", nil)
}

//
//func MarkdownToHtml(con string) string {
//	renderer := blackfriday.HtmlRenderer(MarkdownToHtmlCommonHtmlFlags, "", "")
//	unsafe := blackfriday.Markdown([]byte(con), renderer, MarkdownToHtmlCommonExtensions)
//	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
//
//	return string(html)
//}