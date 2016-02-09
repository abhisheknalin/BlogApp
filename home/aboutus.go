package home

import (
    "net/http"
    "html/template"

    "github.com/abhisheknalin/blogapp/common"
)

func GetAboutusPage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }
    
    p := Page{
        Title: "Aboutus",
    }

    common.Templates = template.Must(template.ParseFiles("templates/home/aboutus.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}