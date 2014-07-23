package home

import (
    "fmt"
    "strconv"
    "net/http"
    "html/template"

    "git-misc/logic-tree/app/common"
)

type Condition struct {
    Field string
    Operator string
    Value int
}

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
        Conditions []Condition
    }
    
    p := Page{
        Title: "home",
        Conditions: getConditions(),
    }

    common.Templates = template.Must(template.ParseFiles("templates/home/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}

func SaveForm(rw http.ResponseWriter, req *http.Request) {
    field := req.FormValue("field")
    operator := req.FormValue("operator")
    value, err := strconv.Atoi(req.FormValue("value"))
    common.CheckError(err, 2)

    _, err = common.DB.Query(fmt.Sprintf("INSERT INTO logictree.conditions(field, operator, value) VALUES ('%s', '%s', %d)", field, operator, value))
    common.CheckError(err, 2)

    GetHomePage(rw, req)
}

func getConditions() []Condition {
    conditions := make([]Condition, 0)

    rows, err := common.DB.Query("SELECT field, operator, value FROM logictree.conditions")
    common.CheckError(err, 2)

    var field, operator, value string
    var valueInt int

    for rows.Next() {
        rows.Scan(&field, &operator, &value)
        valueInt, err = strconv.Atoi(value)
        common.CheckError(err, 2)

        conditions = append(conditions, Condition{Field: field, Operator: operator, Value: valueInt})
    }

    return conditions
}