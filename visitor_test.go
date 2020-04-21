package astTool

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestSearcher_FindFuncDecl(t *testing.T) {

	srccode := `package main

import "fmt"

const TIPS = "jack";

func PrintTips() {
    fmt.Println(TIPS)
}
func Tet(){}
`
	fset = token.NewFileSet()
	astNode, err := parser.ParseFile(fset, "", srccode, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	nodes := make([]ast.Node, 0)
	nodes = append(nodes, astNode)
	searcher := Searcher{
		Root: astNode}
	resultNode := searcher.FindFuncDecl("PrintTips")
	if resultNode != nil {
		if resultNode.Name.Name != "PrintTips" {
			t.Errorf("got %s,expect %s", resultNode.Name.Name, "PrintTips")
		}
	}
}

func TestSearcher_FindValueSpec_case1(t *testing.T) {

	srccode := `package main

import "fmt"

const TIPS = "jack";
var x string
func PrintTips() {
    fmt.Println(TIPS)
	x= "tom and jack"
}
func Tet(){}
`
	fset = token.NewFileSet()
	astNode, err := parser.ParseFile(fset, "", srccode, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	nodes := make([]ast.Node, 0)
	nodes = append(nodes, astNode)
	searcher := Searcher{
		Root: astNode}
	resultNode := searcher.FindValueSpec("TIPS")
	if resultNode != nil {
		if resultNode.Names[0].Name != "TIPS" {
			t.Errorf("got %s,expect %s", resultNode.Names[0].Name, "TIPS")
		}
	}
}

func TestSearcher_FindValueSpec_case2(t *testing.T) {

	srccode := `package main

import "fmt"

const TIPS = "jack";
var x string
func PrintTips() {
	var y string
	y="3"
	_=y
    fmt.Println(TIPS)
	x= "tom and jack"
}
func Tet(){}
`
	fset = token.NewFileSet()
	astNode, err := parser.ParseFile(fset, "", srccode, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	nodes := make([]ast.Node, 0)
	nodes = append(nodes, astNode)
	searcher := Searcher{Root: astNode}
	resultNode := searcher.FindFuncDecl("PrintTips")
	if resultNode != nil {
		if resultNode.Name.Name != "PrintTips" {
			t.Errorf("got %s,expect %s", resultNode.Name.Name, "PrintTips")
		}

		searcher2 := Searcher{Root: resultNode}
		resultNode2 := searcher2.FindValueSpec("y")
		if resultNode2 != nil {
			if resultNode2.Names[0].Name != "y" {
				t.Errorf("got %s,expect %s", resultNode2.Names[0].Name, "y")
			}
		}
	}
}
