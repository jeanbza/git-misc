package home

import (
    "fmt"
    "strconv"
    "net/http"
    "html/template"

    "git-misc/logic-tree/app/common"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }
    
    p := Page{
        Title: "home",
    }

    common.Templates = template.Must(template.ParseFiles("templates/home/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}

func SaveForm(rw http.ResponseWriter, req *http.Request) {
    field := req.FormValue("field")
    operator := req.FormValue("operator")
    value, _ := strconv.Atoi(req.FormValue("value"))

    _, err := common.DB.Query(fmt.Sprintf("INSERT INTO logictree.conditions(field, operator, value) VALUES ('%s', '%s', %d)", field, operator, value))

    GetHomePage(rw, req)
}