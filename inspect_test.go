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
	fnode := h.FindNodeByPos(fpos)
	gotName := (*fnode).(*ast.FuncDecl).Name.Name
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
	structNode := h.FindNodeByPos(structPos)
	pos := h.FindStructFieldFromNode(*structNode, "ActivityHost")
	if pos == token.NoPos || h.Position(pos).String() != "3:1" {
		t.Errorf("got %q;wanted %q", h.Position(pos), "3:1")
	}
}
