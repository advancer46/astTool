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
	resultNode := searcher.FindFuncDeclGlobal("PrintTips")
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
	resultNode := searcher.FindValueSpecGlobal("TIPS")
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
	resultNode := searcher.FindFuncDeclGlobal("PrintTips")
	if resultNode != nil {
		if resultNode.Name.Name != "PrintTips" {
			t.Errorf("got %s,expect %s", resultNode.Name.Name, "PrintTips")
		}

		searcher2 := Searcher{Root: resultNode}
		resultNode2 := searcher2.FindValueSpecGlobal("y")
		if resultNode2 != nil {
			if resultNode2.Names[0].Name != "y" {
				t.Errorf("got %s,expect %s", resultNode2.Names[0].Name, "y")
			}
		}
	}
}

func TestSearcher_FindTypeDecl_case1(t *testing.T) {
	var input = `package miclient
type s struct{
 	a int32
}
`
	var expect = `type s struct {
	a int32
}`

	h := ParseFromCode(input)

	searcher := Searcher{Root: h.Ast}

	resultNode := searcher.FindTypeDecl("s")
	if resultNode != nil {
		got := h.OutputNode(resultNode)

		if got != expect {
			t.Errorf("\n got %q, \n exp %q", got, expect)
		}
	} else {
		t.Logf("got nil,expect %q", expect)
	}
}

func TestSearcher_FindTypeDecl_case2(t *testing.T) {
	var input = `package miclient
type s interface{
 	UserGet()
}
`
	var expect = `type s interface {
	UserGet()
}`

	h := ParseFromCode(input)

	searcher := Searcher{Root: h.Ast}

	resultNode := searcher.FindTypeDecl("s")
	if resultNode != nil {
		got := h.OutputNode(resultNode)

		if got != expect {
			t.Errorf("\n got %q, \n exp %q", got, expect)
		}
	} else {
		t.Logf("got nil,expect %q", expect)
	}
}

func TestSearcher_FindFuncDecl_case1(t *testing.T) {
	var input = `package miclient

func NewPartnerSvcEndpoints(service service.PartnerSvcService) int {
	return 0
}
`
	var expect = `func NewPartnerSvcEndpoints(service service.PartnerSvcService) int {
	return 0
}`

	h := ParseFromCode(input)

	searcher := Searcher{Root: h.Ast}

	resultNode := searcher.FindFuncDecl("NewPartnerSvcEndpoints")
	if resultNode != nil {
		got := h.OutputNode(resultNode)

		if got != expect {
			t.Errorf("\n got %q, \n exp %q", got, expect)
		}
	} else {
		t.Logf("got nil,expect %q", expect)
	}
}

func TestSearcher_FindField(t *testing.T) {
	var input = `package service

type PrivilegeSvcService interface {
	PrivilegeFetch()
}
`
	var expect = `PrivilegeFetch`

	h := ParseFromCode(input)

	searcher := Searcher{Root: h.Ast}

	resultNode := searcher.FindField("PrivilegeFetch")
	if resultNode != nil {
		//got := h.OutputNode(resultNode)

		if resultNode.Names[0].Name != expect {
			t.Errorf("\n got %q, \n exp %q", resultNode.Names[0].Name, expect)
		}
	} else {
		t.Logf("got nil,expect %q", expect)
	}
}
