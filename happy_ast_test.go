
package astTool

import (
	"testing"
	"go/token"
	"go/ast"
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
	spec := NewValueSpec("name","string")
	decl := NewGenDecl(token.VAR,spec)
	h.AddDecl(HEAD,decl)
	if h.Output() != wantedCodes[0]{
		t.Errorf("case %d got %q;wanted %q",0,h.Output(),wantedCodes[0])
	}

	// 1, add tail decl
	expStmt := NewExpStmt(NewCallExpr("callFoo1"))
	blkStmt := NewBlockStmt(expStmt)
	funcDecl := NewFuncDecl("testFoo",blkStmt)
	h.AddDecl(TAIL,funcDecl)
	if h.Output() != wantedCodes[1]{
		t.Errorf("case %d got %q;wanted %q",1,h.Output(),wantedCodes[1])
	}
}


func TestHappyAst_FindFuncDeclNode(t *testing.T) {
	srcode := `package miclient 
var  ppx int
func testFoo(){}`
	h := ParseFromCode(srcode)

	wantedFuncName := "testFoo"
	fpos := h.FindFuncDeclNode(wantedFuncName)
	fnode := h.FindNodeByPos(fpos)
	gotName := (*fnode).(*ast.FuncDecl).Name.Name
	if  gotName != wantedFuncName{
		t.Errorf("got %q;wanted %q",gotName,wantedFuncName)
	}
}

func TestHappyAst_FindStructDeclNode(t *testing.T) {
	srcode := `package miclient
type Microservice struct {
ActivityHost           string `+"`"+`json:"activity_service_host"`+"`"+`
}`
	h := ParseFromCode(srcode)

	pos := h.FindStructDeclNode("Microservice")
	if pos == token.NoPos || h.Position(pos).String() != "2:6" {
		t.Errorf("got %q;wanted %q",h.Position(pos),"2:6")
	}
}

func TestHappyAst_FindStructFieldFromNode(t *testing.T) {
	srcode := `package miclient
type Microservice struct {
ActivityHost           string `+"`"+`json:"activity_service_host"`+"`"+`
}`
	h := ParseFromCode(srcode)

	structPos := h.FindStructDeclNode("Microservice")
	structNode := h.FindNodeByPos(structPos)
	pos := h.FindStructFieldFromNode(*structNode,"ActivityHost")
	if pos == token.NoPos || h.Position(pos).String() != "3:1" {
		t.Errorf("got %q;wanted %q",h.Position(pos),"3:1")
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
	h.AddStmt(fbody,TAIL,stmt)
	gotCode := h.Output()
	if gotCode != wantedCode{
		t.Errorf("got %q;wanted %q",gotCode, wantedCode)
	}

}


func TestHappyAst_ReplaceNode(t *testing.T) {
	srcode := `package miclient
type Microservice struct {
ActivityHost string `+"`"+`json:"activity_service_host"`+"`"+`
}`
	caseCode := `package miclient

type Microservice struct {
	ActivityHost string `+"`"+`json:"activity_service_host"`+"`"+`
	BrandHost    string `+"`"+`json:"brandcustomer_service_host"`+"`"+`
}
`
	h := ParseFromCode(srcode)

	tag := NewBasicLit(token.STRING,"`json:\"brandcustomer_service_host\"`")
	field := NewField([]string{"BrandHost"},"string",false,tag)
	gpos := h.FindStructDeclNode("Microservice")
	gnode := h.FindNodeByPos(gpos)
	structNode := (*gnode).(*ast.TypeSpec).Type.(*ast.StructType)
	structNode.Fields.List = append(structNode.Fields.List,field)
	h.ReplaceNode(gpos,*gnode)
	updatedCode := h.Output()
	if updatedCode != caseCode{
		t.Errorf("got \n%q\nwanted \n%q",updatedCode,caseCode)
	}
}
