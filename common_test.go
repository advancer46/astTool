package astTool

import (
	"go/ast"
	"testing"
)

func TestNewFieldOfFuncType(t *testing.T) {
	var input = `package miclient
type svc interface {
}`
	var expect = `package miclient

type svc interface {
	RoleGet()
}
`

	_, _ = input, expect
	h := ParseFromCode(input)
	//search
	searcher := Searcher{Root: h.ast}
	resultNode := searcher.FindTypeDecl("svc")
	if resultNode == nil {
		t.Logf("can not find typeDecl(%s)", "svc")
	}
	typeSpec := resultNode.Specs[0].(*ast.TypeSpec)
	interfaceType := typeSpec.Type.(*ast.InterfaceType)

	newfunctype := NewFuncType("", nil, nil)
	newField := NewFieldOfFuncType([]string{"RoleGet"}, newfunctype, nil)

	fieldList := NewFieldList(newField)

	//replace
	interfaceType.Methods = fieldList

	got := h.Output()
	if got != expect {
		t.Errorf("\n got:%q,\n exp:%q", got, expect)
	}
}
