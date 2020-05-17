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
	searcher := Searcher{Root: h.Ast}
	resultNode := searcher.FindTypeDecl("svc")
	if resultNode == nil {
		t.Logf("can not find typeDecl(%s)", "svc")
	}
	typeSpec := resultNode.Specs[0].(*ast.TypeSpec)
	interfaceType := typeSpec.Type.(*ast.InterfaceType)

	newfunctype := NewFuncType(nil, nil)
	newField := NewFieldOfFuncType([]string{"RoleGet"}, newfunctype, nil)

	fieldList := NewFieldList(newField)

	//replace
	interfaceType.Methods = fieldList

	got := h.Output()
	if got != expect {
		t.Errorf("\n got:%q,\n exp:%q", got, expect)
	}
}

func TestNewField(t *testing.T) {
	var input = `package miclient
type svc interface {
}`
	var expect = `package miclient

type svc interface {
	RoleGet(ctx context.Context, reqproto *partnerproto.StaffAuthFetchReqProto) (*partnerproto.StaffAuthFetchRespProto, error)
}
`

	_, _ = input, expect
	h := ParseFromCode(input)

	//search
	searcher := Searcher{Root: h.Ast}
	resultNode := searcher.FindTypeDecl("svc")
	if resultNode == nil {
		t.Logf("can not find typeDecl(%s)", "svc")
	}
	typeSpec := resultNode.Specs[0].(*ast.TypeSpec)
	interfaceType := typeSpec.Type.(*ast.InterfaceType)

	selectExp := NewSelectExp(NewIdent("context"), NewIdent("Context"))
	paramFieldCtx := NewField([]string{"ctx"}, selectExp, ExprTypeSelectorExpr, nil)

	selectExp1 := NewSelectExp(NewIdent("partnerproto"), NewIdent("StaffAuthFetchReqProto"))
	startExp1 := NewStarExp(selectExp1)
	paramFieldReqproto := NewField([]string{"reqproto"}, startExp1, ExprTypeStartExpr, nil)
	params := NewFieldList(paramFieldCtx, paramFieldReqproto)

	selectExp2 := NewSelectExp(NewIdent("partnerproto"), NewIdent("StaffAuthFetchRespProto"))
	startExp2 := NewStarExp(selectExp2)
	resultField1 := NewField([]string{""}, startExp2, ExprTypeStartExpr, nil)
	resultField2 := NewField([]string{""}, NewIdent("error"), ExprTypeIdent, nil)
	results := NewFieldList(resultField1, resultField2)
	newfunctype := NewFuncType(params, results)
	newField := NewFieldOfFuncType([]string{"RoleGet"}, newfunctype, nil)
	fieldList := NewFieldList(newField)

	//replace
	interfaceType.Methods = fieldList

	got := h.Output()
	if got != expect {
		t.Errorf("\n got:%q,\n exp:%q", got, expect)
	}
}
