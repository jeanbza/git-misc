package home

import (
    "fmt"
    "encoding/json"
    "strconv"
    "net/http"
    "html/template"
    "errors"

    "git-misc/logic-tree/app/common"
)

type Condition struct {
    Text string
    Type string
    Field string
    Operator string
    Value string
}

type treeNode struct {
    Parent *treeNode
    Children []*treeNode
    Node Condition
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

    common.Templates = template.Must(template.New("asd").ParseFiles("templates/home/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}

func SaveState(rw http.ResponseWriter, req *http.Request) {
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

func Truncate(rw http.ResponseWriter, req *http.Request) {
    _, err := common.DB.Query("TRUNCATE TABLE logictree.equality")
    common.CheckError(err, 2)

    _, err = common.DB.Query("TRUNCATE TABLE logictree.logic")
    common.CheckError(err, 2)

    GetHomePage(rw, req)
}

func getConditions() []Condition {
    conditions := make([]Condition, 0)

    rows, err := common.DB.Query("SELECT field, operator, value FROM logictree.equality")
    common.CheckError(err, 2)

    var field, operator, value string

    i := 0

    for rows.Next() {
        rows.Scan(&field, &operator, &value)
        common.CheckError(err, 2)

        if i != 0 {
            conditions = append(conditions, Condition{
                Text: "AND",
                Type: "logic",
            })
        }

        conditions = append(conditions, Condition{
            Text: fmt.Sprintf("%s %s %s", field, operator, value),
            Type: "equality",
            Field: field,
            Operator: operator,
            Value: value,
        })

        i++
    }

    return conditions
}

func parseJSON(conditionsString string) ([]Condition, error) {
    var conditionsSlice []Condition
    
    err := json.Unmarshal([]byte(conditionsString), &conditionsSlice)
    if err != nil {
        return nil, err
    }
    
    return conditionsSlice, nil
}

func unserializeTree(conditions []Condition) (*treeNode, error) {
    depth := 0
    var root treeNode

    for _, condition := range conditions {
        switch condition.Type {
        case "scope":
            if condition.Operator == "(" {
                depth++
            }

            if condition.Operator == ")" {
                depth--
            }
        case "logic":
            root.Node = condition
        case "equality":
            if depth > 0 {
                root.Children = append(root.Children, &treeNode{Parent: &root, Node: condition})
            } else {
                root.Node = condition
            }
        }
    }

    return &root, nil
}

func serializeTree(node *treeNode) ([]Condition, error) {
    if node.Children == nil || len(node.Children) == 0 {
        // Has no children - should be equality

        if node.Node.Type != "equality" {
            return nil, errors.New("ERROR: This tree has a logic condition as a leaf. Quitting.")
        }

        return []Condition{node.Node}, nil
    } else {
        // Has children - should be logic

        if node.Node.Type != "logic" {
            return nil, errors.New("ERROR: This tree has an equality condition as a branch. Quitting.")
        }
    }

    linearConditions := []Condition{Condition{Text: "(", Type: "scope", Operator: "("}}

    for key, child := range node.Children {
        if key != 0 {
            linearConditions = append(linearConditions, node.Node)
        }

        serializedChild, err := serializeTree(child)

        if err != nil {
            return nil, err
        }

        linearConditions = append(linearConditions, serializedChild...)
    }

    linearConditions = append(linearConditions, Condition{Text: ")", Type: "scope", Operator: ")"})

    return linearConditions, nil
}

func (t *treeNode) print() string {
    var s string

    for _, child := range t.Children {
        s += child.print()
    }

    return s + " :: " + fmt.Sprintf("%v", t.Node)
}

func simplifyConditions(conditions []Condition) string {
    var t string

    for k, c := range conditions {
        if k != 0 {
            t += " "
        }

        t += c.Text
    }

    return t
}


