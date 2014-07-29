package home

import (
    "fmt"
    "testing"
)

func beforeEach() {
    fmt.Println("Starting test..")
}

func TestSerializeTreeOneNodeZeroDepth(t *testing.T) {
    beforeEach()

    root := &treeNode{Parent: nil, Children: nil, Node: Condition{Text: "age eq 5", Type: "equality", Field: "age", Operator: "eq", Value: "5"}}
    expectedOut := []Condition{Condition{Text: "age eq 5", Type: "equality", Field: "age", Operator: "eq", Value: "5"}}
    var expectedOutErr error

    conditionsReturned, errorsReturned := serializeTree(root)

    if !matchesArray(conditionsReturned, expectedOut) {
        t.Errorf("serializeTree(%v) - got %v, want %v", root, conditionsReturned, expectedOut)
    }

    if errorsReturned != expectedOutErr {
        t.Errorf("serializeTree(%v) errorsReturned - got %v, want %v", root, errorsReturned, expectedOutErr)
    }
}

func TestSerializeTreeThreeNodeOneDepth(t *testing.T) {
    beforeEach()

    root := &treeNode{Parent: nil, Children: nil, Node: Condition{Text: "AND", Type: "logic", Operator: "AND"}}
    child1 := treeNode{Parent: root, Children: nil, Node: Condition{Text: "age eq 8", Type: "equality", Field: "age", Operator: "eq", Value: "8"}}
    child2 := treeNode{Parent: root, Children: nil, Node: Condition{Text: "age eq 2", Type: "equality", Field: "age", Operator: "eq", Value: "2"}}
    root.Children = []*treeNode{&child1, &child2}

    expectedOut := []Condition{
        Condition{Text: "age eq 8", Type: "equality", Field: "age", Operator: "eq", Value: "8"},
        Condition{Text: "AND", Type: "logic", Operator: "AND"},
        Condition{Text: "age eq 2", Type: "equality", Field: "age", Operator: "eq", Value: "2"},
    }
    var expectedOutErr error

    conditionsReturned, errorsReturned := serializeTree(root)

    if !matchesArray(conditionsReturned, expectedOut) {
        t.Errorf("serializeTree(%v) conditionsReturned - got %v, want %v", root, conditionsReturned, expectedOut)
    }

    if errorsReturned != expectedOutErr {
        t.Errorf("serializeTree(%v) errorsReturned - got %v, want %v", root, errorsReturned, expectedOutErr)
    }
}

func TestSerializeTreeWithEqualityBranch(t *testing.T) {
    beforeEach()

    root := &treeNode{Parent: nil, Children: nil, Node: Condition{Text: "age eq 4", Type: "equality", Field: "age", Operator: "eq", Value: "4"}}
    child1 := treeNode{Parent: root, Children: nil, Node: Condition{Text: "age eq 8", Type: "equality", Field: "age", Operator: "eq", Value: "8"}}
    child2 := treeNode{Parent: root, Children: nil, Node: Condition{Text: "age eq 2", Type: "equality", Field: "age", Operator: "eq", Value: "2"}}
    root.Children = []*treeNode{&child1, &child2}

    expectedOutErr := "ERROR: This tree has an equality condition as a branch. Quitting."

    conditionsReturned, errorsReturned := serializeTree(root) 

    if conditionsReturned != nil {
        t.Errorf("serializeTree(%v) - got %v, want %v", root, conditionsReturned, nil)
    }

    if errorsReturned.Error() != expectedOutErr {
        t.Errorf("serializeTree(%v) errorsReturned - got %v, want %v", root, errorsReturned, expectedOutErr)
    }
}

func matchesArray(conditionsA []Condition, conditionsB []Condition) bool {
    var truth bool

    if conditionsA == nil || len(conditionsA) != len(conditionsB) {
        return false
    }

    for _, valA := range conditionsA {
        truth = false

        for _, valB := range conditionsB {
            if valA.matches(valB) {
                truth = true
            }
        }

        if !truth {
            return false
        }
    }

    return true
}

func (conditionA Condition) matches(conditionB Condition) bool {
    if conditionA.Text != conditionB.Text {
        return false
    }

    if conditionA.Type != conditionB.Type {
        return false
    }

    if conditionA.Field != conditionB.Field {
        return false
    }

    if conditionA.Operator != conditionB.Operator {
        return false
    }

    if conditionA.Value != conditionB.Value {
        return false
    }

    return true
}

