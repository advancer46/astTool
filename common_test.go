package astTool

import (
	"go/ast"
	"go/token"
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

	got := h.Output(nil)
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

	got := h.Output(nil)
	if got != expect {
		t.Errorf("\n got:%q,\n exp:%q", got, expect)
	}
}

// todo
func TestNewEmptyStmt(t *testing.T) {
	srcode := `package miclient
var  ppx int
func testFoo(){}`
	h := ParseFromCode(srcode)

	wantedCode := `package miclient

var ppx int

func testFoo() { CALLfoo() }
`
	//2, add stmt into func body
	searcher := Searcher{Root: h.Ast}
	funcDecl := searcher.FindFuncDecl("testFoo")

	//fpos := h.FindFuncDeclNode("testFoo")
	//fnode := h.FindFuncDeclNode(fpos)
	//fbody := (*fnode).(*ast.FuncDecl).Body

	stmt := NewExpStmt(NewCallExpr(NewIdent("CALLfoo"), nil))
	h.AddStmt(funcDecl.Body, TAIL, stmt)

	pos := funcDecl.Pos()
	emptystmt := NewEmptyStmt(pos)
	h.AddStmt(funcDecl.Body, TAIL, emptystmt)

	gotCode := h.Output(nil)
	if gotCode != wantedCode {
		t.Errorf("\n got: %q \n exp: %q", gotCode, wantedCode)
	}

}

func TestNewDeclStmt(t *testing.T) {
	srcode := `package miclient
var  ppx int
func testFoo(){}`
	h := ParseFromCode(srcode)

	wantedCode := `package miclient

var ppx int

func testFoo() { var guessCreateEndpoint kitendpoint.Endpoint }
`
	//2, add stmt into func body
	searcher := Searcher{Root: h.Ast}
	funcDecl := searcher.FindFuncDecl("testFoo")

	specs := make([]ast.Spec, 0)
	specs = append(specs, NewValueSpec("guessCreateEndpoint", NewSelectExp(NewIdent("kitendpoint"), NewIdent("Endpoint"))))
	genDecl := NewGenDecl(token.VAR, specs...)
	declStmt := NewDeclStmt(genDecl)
	h.AddStmt(funcDecl.Body, TAIL, declStmt)

	gotCode := h.Output(nil)
	if gotCode != wantedCode {
		t.Errorf("\n got: %q \n exp: %q", gotCode, wantedCode)
	}

}
