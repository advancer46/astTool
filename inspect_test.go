package astTool

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHappyAst_FindFuncDeclNode(t *testing.T) {
	srcode := `package miclient 
var  ppx int
func testFoo(){}`
	h := ParseFromCode(srcode)

	wantedFuncName := "testFoo"
	fpos := h.FindFuncDeclNode(wantedFuncName)
	fnode := h.NodeByPos(fpos)
	gotName := fnode.(*ast.FuncDecl).Name.Name
	if gotName != wantedFuncName {
		t.Errorf("got %q;wanted %q", gotName, wantedFuncName)
	}
}

func TestHappyAst_FindStructDeclNode(t *testing.T) {
	srcode := `package miclient
type Microservice struct {
ActivityHost           string ` + "`" + `json:"activity_service_host"` + "`" + `
}`
	h := ParseFromCode(srcode)

	pos := h.FindStructDeclNode("Microservice")
	if pos == token.NoPos || h.Position(pos).String() != "2:6" {
		t.Errorf("got %q;wanted %q", h.Position(pos), "2:6")
	}
}

func TestHappyAst_FindStructFieldFromNode(t *testing.T) {
	srcode := `package miclient
type Microservice struct {
ActivityHost           string ` + "`" + `json:"activity_service_host"` + "`" + `
}`
	h := ParseFromCode(srcode)

	structPos := h.FindStructDeclNode("Microservice")
	structNode := h.NodeByPos(structPos)
	pos := h.FindStructFieldFromNode(structNode, "ActivityHost")
	if pos == token.NoPos || h.Position(pos).String() != "3:1" {
		t.Errorf("got %q;wanted %q", h.Position(pos), "3:1")
	}
}

func TestHappyAst_Print(t *testing.T) {
	srcCode := `package service

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
	h := ParseFromCode(srcCode)

	h.Print()
}
