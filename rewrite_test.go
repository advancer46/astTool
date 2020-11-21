package astTool

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"testing"
)

func TestHappyAst_AddDecl(t *testing.T) {
	srcode := `package miclient
var ppx int`
	h := ParseFromCode(srcode)

	wantedCodes := []string{
		`package miclient

var name string
var ppx int
`,
		`package miclient

var name string
var ppx int

// this is comment
func (h *Receive) testFoo(a, b string) {
	callFoo1()
}
`,
	}

	// 0,add head decl
	spec := NewValueSpec("name", nil)
	decl := NewGenDecl(token.VAR, spec)
	h.AddDecl(HEAD, decl)
	if h.Output(nil) != wantedCodes[0] {
		t.Errorf("case %d got %q;wanted %q", 0, h.Output(nil), wantedCodes[0])
	}

	// 1, add tail decl
	expStmt := NewExpStmt(NewCallExpr(NewIdent("callFoo1"), nil))
	blkStmt := NewBlockStmt(expStmt)
	recvField := NewField([]string{"h"}, NewIdent("Receive"), ExprTypeIdent, nil)
	_ = recvField
	field := NewField([]string{"a", "b"}, NewIdent("string"), ExprTypeIdent, nil)
	commentgroup := NewCommentGroup(NewComment("// this is comment"))
	funcDecl := NewFuncDecl("testFoo", blkStmt, NewFieldList(recvField), NewFieldList(field), commentgroup)
	h.AddDecl(TAIL, funcDecl)
	if h.Output(nil) != wantedCodes[1] {
		t.Errorf("case %d got %q;wanted %q", 1, h.Output(nil), wantedCodes[1])
	}
}

func TestHappyAst_AddStmt(t *testing.T) {
	srcode := `package miclient
var  ppx int
func testFoo(){}`
	h := ParseFromCode(srcode)

	wantedCode := `package miclient

var ppx int

func testFoo() { CALLfoo() }
`
	//2, add stmt into func body
	fpos := h.FindFuncDeclNode("testFoo")
	fnode := h.NodeByPos(fpos)
	fbody := fnode.(*ast.FuncDecl).Body

	stmt := NewExpStmt(NewCallExpr(NewIdent("CALLfoo"), nil))
	h.AddStmt(fbody, TAIL, stmt)
	gotCode := h.Output(nil)
	if gotCode != wantedCode {
		t.Errorf("got %q;wanted %q", gotCode, wantedCode)
	}

}

func TestHappyAst_ReplaceNode(t *testing.T) {
	var input = `package miclient
func testFoo() {}
`
	var expect = `package miclient

func testFoo2() {
}
`
	h := ParseFromCode(input)
	pos := h.FindFuncDeclNode("testFoo")

	newNode := NewFuncDecl("testFoo2", NewBlockStmt(), nil, nil, nil)

	err := h.ReplaceNode(pos, newNode)
	if err != nil {
		log.Print(err.Error())
	}

	got := h.Output(nil)
	if got != expect {
		t.Errorf("got \n%q\nwanted \n%q", got, expect)
	}
}

func TestHappyAst_AddAssignStmt(t *testing.T) {
	var input = `package miclient
func testFoo() {}
`
	var expect = `package miclient

func testFoo() { x := 3; _ = x }
`
	h := ParseFromCode(input)

	lhs := make([]ast.Expr, 0)
	lhs = append(lhs, NewIdent("x"))
	rhs := make([]ast.Expr, 0)
	rhs = append(rhs, NewBasicLit(token.INT, "3"))
	assignStmt1 := NewShortAssignStmt(lhs, rhs)

	lhs = make([]ast.Expr, 0)
	lhs = append(lhs, NewIdent("_"))
	rhs = make([]ast.Expr, 0)
	rhs = append(rhs, NewIdent("x"))
	assignStmt2 := NewAssignStmt(lhs, rhs)

	searcher := Searcher{Root: h.Ast}
	funcNode := searcher.FindFuncDecl("testFoo")
	h.AddAssignStmt(funcNode.Body, 0, assignStmt1)
	h.AddAssignStmt(funcNode.Body, 1, assignStmt2)
	newcode := h.Output(nil)

	if newcode != expect {
		t.Errorf("\ngot:%q \nexp:%q", newcode, expect)
	}
}

func TestHappyAst_AddFundDecl(t *testing.T) {
	var input = `package miclient
func testFoo() {}
`
	var expect = `package miclient

func testFoo() {}

// this is comment
func (h *Receive) testFoo1(a, b string) {
	callFoo()
}

// this is comment
func (h *Receive) testFoo2(a, b string) {
	callFoo()
}
`
	h := ParseFromCode(input)

	// 1, append decl 1
	expStmt := NewExpStmt(NewCallExpr(NewIdent("callFoo"), nil))
	blkStmt := NewBlockStmt(expStmt)
	recvField := NewField([]string{"h"}, NewIdent("Receive"), ExprTypeIdent, nil)
	field := NewField([]string{"a", "b"}, NewIdent("string"), ExprTypeIdent, nil)
	commentgroup := NewCommentGroup(NewComment("// this is comment"))
	funcDecl1 := NewFuncDecl("testFoo1", blkStmt, NewFieldList(recvField), NewFieldList(field), commentgroup)
	h.AppendFundDecl(1, funcDecl1)

	// 1, append decl 2
	expStmt = NewExpStmt(NewCallExpr(NewIdent("callFoo"), nil))
	blkStmt = NewBlockStmt(expStmt)
	recvField = NewField([]string{"h"}, NewIdent("Receive"), ExprTypeIdent, nil)
	field = NewField([]string{"a", "b"}, NewIdent("string"), ExprTypeIdent, nil)
	commentgroup = NewCommentGroup(NewComment("// this is comment"))
	funcDecl2 := NewFuncDecl("testFoo2", blkStmt, NewFieldList(recvField), NewFieldList(field), commentgroup)
	h.AppendFundDecl(2, funcDecl2)

	if h.Output(nil) != expect {
		t.Errorf("\ngot: %q \nexp: %q", h.Output(nil), expect)
	}
}

func TestHappyAst_AddFieldOfFuncType(t *testing.T) {
	var input = `package miclient
type svc interface {
				UserGet()
}`
	var expect = `package miclient

type svc interface {
	UserGet()
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
	fieldList := interfaceType.Methods
	_ = fieldList

	// add
	newfunctype := NewFuncType(nil, nil)
	newField := NewFieldOfFuncType([]string{"RoleGet"}, newfunctype, nil)
	h.AddFieldOfFuncType(fieldList, 1, newField)

	got := h.Output(nil)
	if got != expect {
		t.Errorf("\n got:%q, \n exp:%q", got, expect)
	}
}

func TestHappyAst_AddField(t *testing.T) {
	var input = `package miclient
type PartnerSvcEndpoints struct {
	modelFetchEndpoint kitendpoint.Endpoint
}`
	var expect = `package miclient

type PartnerSvcEndpoints struct {
	modelFetchEndpoint kitendpoint.Endpoint
	gameFetchEndpoint  kitendpoint.Endpoint
}
`

	_, _ = input, expect

	h := ParseFromCode(input)

	//search
	searcher := Searcher{Root: h.Ast}
	declName := "PartnerSvcEndpoints"
	resultNode := searcher.FindTypeDecl(declName)
	if resultNode == nil {
		t.Logf("can not find typeDecl(%s)", declName)
	}
	typeSpec := resultNode.Specs[0].(*ast.TypeSpec)
	interfaceType := typeSpec.Type.(*ast.StructType)
	fieldList := interfaceType.Fields
	_ = fieldList

	// add
	selectExp := NewSelectExp(NewIdent("kitendpoint"), NewIdent("Endpoint"))
	paramField := NewField([]string{"gameFetchEndpoint"}, selectExp, ExprTypeSelectorExpr, nil)

	h.AddField(fieldList, 1, paramField)

	got := h.Output(nil)
	if got != expect {
		t.Errorf("\n got:%q, \n exp:%q", got, expect)
	}
}

// todo 测试未通过
func TestHappyAst_FreshPosInfoOfCommentGroup(t *testing.T) {
	var input = `package miclient

// comment 1
type PartnerSvcEndpoints struct {

	// comment 2
	modelFetchEndpoint kitendpoint.Endpoint

	// comment 3	

}`
	var expect = `package miclient

type PartnerSvcEndpoints struct {
	modelFetchEndpoint kitendpoint.Endpoint
	gameFetchEndpoint  kitendpoint.Endpoint
}
`
	_ = expect

	h := ParseFromCode(input)

	// fresh position info of comment group
	err := h.FreshPosInfoOfCommentGroup()
	if err != nil {
		log.Print(err.Error())
	}

	// 更新节点
	searcher := Searcher{Root: h.Ast}
	declName := "PartnerSvcEndpoints"
	resultNode := searcher.FindTypeDecl(declName)
	if resultNode == nil {
		t.Logf("can not find typeDecl(%s)", declName)
	}
	typeSpec := resultNode.Specs[0].(*ast.TypeSpec)
	interfaceType := typeSpec.Type.(*ast.StructType)
	fieldList := interfaceType.Fields
	_ = fieldList

	// add
	selectExp := NewSelectExp(NewIdent("kitendpoint"), NewIdent("Endpoint"))
	paramField := NewField([]string{"gameFetchEndpoint"}, selectExp, ExprTypeSelectorExpr, nil)

	h.AddField(fieldList, 1, paramField)

	s := h.Output(nil)
	fmt.Print(s)
}
