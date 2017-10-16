package main

import (
	"os"
	"fmt"
	"strconv"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

var templates = template.Must(template.ParseGlob("tmpl/*"))

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

type Filedata struct {
  Name string
  Size int64
  Modtime time.Time
}

type Dir struct {
  Files []Filedata
}

func (d *Dir) getDatafileInfo() (dd Dir, err error) {
  f := Filedata{}
  filenames, err := filepath.Glob("data/*")
  if err != nil {
    return Dir{}, err
  }
  for i := range filenames {
    Fileinfo, err := os.Stat(filenames[i])
    if err != nil {
      return Dir{}, err
    }

    // isolate name without directory or extension
    name := strings.SplitAfter(filenames[i],"/")
		name = strings.Split(name[1],".")
    f.Name = name[0]
    f.Size = Fileinfo.Size()
    f.Modtime = Fileinfo.ModTime()
    d.Files = append(d.Files, f)
  }
  return Dir{Files:d.Files}, err
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
		err := templates.ExecuteTemplate(w, tmpl, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	dd := Dir{}
	d, err := dd.getDatafileInfo()

	//d, err := getDatafileInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = templates.ExecuteTemplate(w, "list", d)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("body")
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/delete/"):]

	err := os.Remove("data/" + title + ".txt")
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/list/", http.StatusFound)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/list/", http.StatusFound)
}

func main() {
	var hostport string = ":8080"
	if len(os.Args) > 1 {
		_, err := strconv.ParseUint(os.Args[1], 10, 16)
		if err != nil {
			panic(err)
		} else {
			hostport = ":" + os.Args[1]
		}
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/list/", listHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/create/", createHandler)
	http.ListenAndServe(hostport, nil)
	fmt.Println("Host listening on port", hostport)
}
