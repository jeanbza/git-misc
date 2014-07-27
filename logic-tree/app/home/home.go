package home

import (
    "fmt"
    "strconv"
    "net/http"
    "html/template"

    "git-misc/logic-tree/app/common"
)

type EqualityCondition struct {
    Field string
    Operator string
    Value int
}

type LogicalCondition struct {
    Operator string
}

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
        Equality []EqualityCondition
        Logic []LogicalCondition
    }
    
    p := Page{
        Title: "home",
        Equality: getEqualityConditions(),
        Logic: getLogicConditions(),
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

    _, err = common.DB.Query(fmt.Sprintf("INSERT INTO logictree.equality(field, operator, value) VALUES ('%s', '%s', %d)", field, operator, value))
    common.CheckError(err, 2)

    _, err = common.DB.Query("INSERT INTO logictree.logic(operator) VALUES ('AND')")
    common.CheckError(err, 2)

    GetHomePage(rw, req)
}

func TruncateEquality(rw http.ResponseWriter, req *http.Request) {
    _, err := common.DB.Query("TRUNCATE TABLE logictree.equality")
    common.CheckError(err, 2)

    GetHomePage(rw, req)
}

func TruncateLogic(rw http.ResponseWriter, req *http.Request) {
    _, err := common.DB.Query("TRUNCATE TABLE logictree.logic")
    common.CheckError(err, 2)

    GetHomePage(rw, req)
}

func getLogicConditions() []LogicalCondition {
    conditions := make([]LogicalCondition, 0)

    rows, err := common.DB.Query("SELECT operator FROM logictree.logic")
    common.CheckError(err, 2)

    var operator string

    for rows.Next() {
        rows.Scan(&operator)

        conditions = append(conditions, LogicalCondition{Operator: operator})
    }

    return conditions
}

func getEqualityConditions() []EqualityCondition {
    conditions := make([]EqualityCondition, 0)

    rows, err := common.DB.Query("SELECT field, operator, value FROM logictree.equality")
    common.CheckError(err, 2)

    var field, operator, value string
    var valueInt int

    for rows.Next() {
        rows.Scan(&field, &operator, &value)
        valueInt, err = strconv.Atoi(value)
        common.CheckError(err, 2)

        conditions = append(conditions, EqualityCondition{Field: field, Operator: operator, Value: valueInt})
    }

    return conditions
}