package astTool

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func TestSearcher_FindFuncDecl(t *testing.T) {

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

func TestSearcher_FindKeyValExpr(t *testing.T) {
	var input = `package miclient
type s struct{
 	a int32
}
func NewPartnerSvcEndpoints(service service.PartnerSvcService) PartnerSvcEndpoints {
    return PartnerSvcEndpoints{
		modelFetchEndpoint:MakemodelFetchEndpoint(service),
		modelCreateEndpoint:MakemodelCreateEndpoint(service),
	}
}
`
	var expect = `modelFetchEndpoint:
MakemodelFetchEndpoint(service)`

	h := ParseFromCode(input)
	searcher := Searcher{Root: h.Ast}
	funcDeclName := "NewPartnerSvcEndpoints"
	funcDeclNode := searcher.FindFuncDecl(funcDeclName)
	if funcDeclNode == nil {
		log.Printf("func decl(%s) not exsit", funcDeclName)
		return
	}
	funcBodyNode := funcDeclNode.Body
	returnStmtNode := funcBodyNode.List[0].(*ast.ReturnStmt)
	compLitNode := returnStmtNode.Results[0].(*ast.CompositeLit)

	kvsearcher := Searcher{Root: compLitNode}
	resultNode := kvsearcher.FindKeyValExpr("modelFetchEndpoint")
	if resultNode != nil {
		got := h.OutputNode(resultNode)
		if got != expect {
			t.Errorf("\n got %q, \n exp %q", got, expect)
		}
	} else {
		t.Logf("got nil,expect %q", expect)
	}
}

func TestSearcher_FindAssignAssign(t *testing.T) {
	var input = `package service

func NewGRPCEndpoint(client kitconsul.Client, guesssvcport int) (guessendpoint.GuessSvcEndpoints, *tlog.LogError) {
 	modelFetchFactory := createGuessSvcFactory(Modelendpoint.MakeModelFetchEndpoint, guesssvcport)
	return endpoints, nil
}
`
	var expect = `modelFetchFactory := createGuessSvcFactory(Modelendpoint.MakeModelFetchEndpoint, guesssvcport)`

	h := ParseFromCode(input)
	searcher := Searcher{Root: h.Ast}
	funcDeclName := "NewGRPCEndpoint"
	funcDeclNode := searcher.FindFuncDecl(funcDeclName)
	if funcDeclNode == nil {
		log.Printf("func decl(%s) not exsit", funcDeclName)
		return
	}

	assignStmtSearcher := Searcher{Root: funcDeclNode.Body}
	resultNode := assignStmtSearcher.FindAssignStmt("modelFetchFactory")
	if resultNode != nil {
		got := h.OutputNode(resultNode)
		if got != expect {
			t.Errorf("\n got %q, \n exp %q", got, expect)
		}
	} else {
		t.Logf("got nil,expect %q", expect)
	}
}

func TestSearcher_FindCommentGroup(t *testing.T) {
	var input = `package service

import (
	"context"

	"svcGenerator/data/proto/v1"
)

//{{template1}}
type CommonSvcService interface {
	//{{template2}}
	IdentifyFetch(ctx context.Context, reqproto *commonproto.IdentifyFetchReqProto) (*commonproto.IdentifyFetchRespProto, error)

	//{{template9}}

}
`
	var expect = `{{template1}}
{{template2}}
{{template9}}
`

	h := ParseFromCode(input)
	searcher := Searcher{Root: h.Ast}
	commentGroups := searcher.FindCommentGroups()
	if len(commentGroups) <= 0 {
		log.Printf("commentGroup not exsit")
		return
	}

	got := ""
	for _, commentGroup := range commentGroups {
		got += commentGroup.Text()
	}

	if got != expect {
		t.Errorf("\n got %q, \n exp %q", got, expect)
	}
}
