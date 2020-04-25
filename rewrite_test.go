package astTool

import (
	"go/ast"
	"go/token"
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
	if h.Output() != wantedCodes[0] {
		t.Errorf("case %d got %q;wanted %q", 0, h.Output(), wantedCodes[0])
	}

	// 1, add tail decl
	expStmt := NewExpStmt(NewCallExpr("callFoo1"))
	blkStmt := NewBlockStmt(expStmt)
	recvField := NewField([]string{"h"}, "Receive", true, nil)
	_ = recvField
	field := NewField([]string{"a", "b"}, "string", false, nil)
	commentgroup := NewCommentGroup(NewComment("// this is comment"))
	funcDecl := NewFuncDecl("testFoo", blkStmt, NewFieldList(recvField), NewFieldList(field), commentgroup)
	h.AddDecl(TAIL, funcDecl)
	if h.Output() != wantedCodes[1] {
		t.Errorf("case %d got %q;wanted %q", 1, h.Output(), wantedCodes[1])
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
	fnode := h.FindNodeByPos(fpos)
	fbody := (*fnode).(*ast.FuncDecl).Body

	stmt := NewExpStmt(NewCallExpr("CALLfoo"))
	h.AddStmt(fbody, TAIL, stmt)
	gotCode := h.Output()
	if gotCode != wantedCode {
		t.Errorf("got %q;wanted %q", gotCode, wantedCode)
	}

}

func TestHappyAst_ReplaceNode(t *testing.T) {
	srcode := `package miclient
type Microservice struct {
ActivityHost string ` + "`" + `json:"activity_service_host"` + "`" + `
}`
	caseCode := `package miclient

type Microservice struct {
	ActivityHost string ` + "`" + `json:"activity_service_host"` + "`" + `
	BrandHost    string ` + "`" + `json:"brandcustomer_service_host"` + "`" + `
}
`
	h := ParseFromCode(srcode)

	tag := NewBasicLit(token.STRING, "`json:\"brandcustomer_service_host\"`")
	field := NewField([]string{"BrandHost"}, "string", false, tag)
	gpos := h.FindStructDeclNode("Microservice")
	gnode := h.FindNodeByPos(gpos)
	structNode := (*gnode).(*ast.TypeSpec).Type.(*ast.StructType)
	structNode.Fields.List = append(structNode.Fields.List, field)
	h.ReplaceNode(gpos, *gnode)
	updatedCode := h.Output()
	if updatedCode != caseCode {
		t.Errorf("got \n%q\nwanted \n%q", updatedCode, caseCode)
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

	searcher := Searcher{Root: h.ast}
	funcNode := searcher.FindFuncDecl("testFoo")
	h.AddAssignStmt(funcNode.Body, 0, assignStmt1)
	h.AddAssignStmt(funcNode.Body, 1, assignStmt2)
	newcode := h.Output()

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
	expStmt := NewExpStmt(NewCallExpr("callFoo"))
	blkStmt := NewBlockStmt(expStmt)
	recvField := NewField([]string{"h"}, "Receive", true, nil)
	field := NewField([]string{"a", "b"}, "string", false, nil)
	commentgroup := NewCommentGroup(NewComment("// this is comment"))
	funcDecl1 := NewFuncDecl("testFoo1", blkStmt, NewFieldList(recvField), NewFieldList(field), commentgroup)
	h.AddFundDecl(1, funcDecl1)

	// 1, append decl 2
	expStmt = NewExpStmt(NewCallExpr("callFoo"))
	blkStmt = NewBlockStmt(expStmt)
	recvField = NewField([]string{"h"}, "Receive", true, nil)
	field = NewField([]string{"a", "b"}, "string", false, nil)
	commentgroup = NewCommentGroup(NewComment("// this is comment"))
	funcDecl2 := NewFuncDecl("testFoo2", blkStmt, NewFieldList(recvField), NewFieldList(field), commentgroup)
	h.AddFundDecl(2, funcDecl2)

	if h.Output() != expect {
		t.Errorf("\ngot: %q \nexp: %q", h.Output(), expect)
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
	searcher := Searcher{Root: h.ast}
	resultNode := searcher.FindTypeDecl("svc")
	if resultNode == nil {
		t.Logf("can not find typeDecl(%s)", "svc")
	}
	typeSpec := resultNode.Specs[0].(*ast.TypeSpec)
	interfaceType := typeSpec.Type.(*ast.InterfaceType)
	fieldList := interfaceType.Methods
	_ = fieldList

	// add
	newfunctype := NewFuncType("", nil, nil)
	newField := NewFieldOfFuncType([]string{"RoleGet"}, newfunctype, nil)
	h.AddFieldOfFuncType(fieldList, 1, newField)

	got := h.Output()
	if got != expect {
		t.Errorf("\n got:%q, \n exp:%q", got, expect)
	}
}
