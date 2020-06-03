package astTool

import (
	"go/ast"
	"go/token"
)

type ExprType string

const (
	ExprTypeIdent        ExprType = "Ident"
	ExprTypeStartExpr             = "StartExpr"
	ExprTypeSelectorExpr          = "SelectorExpr"
)

func NewIdent(name string) *ast.Ident {
	return &ast.Ident{
		Name:    name,
		Obj:     nil,
		NamePos: token.NoPos,
	}
}

/*
type svc interface {
  UserGet()
}
*/
func NewFuncType(params, results *ast.FieldList) *ast.FuncType {
	return &ast.FuncType{
		Results: results,
		Params:  params,
		Func:    token.NoPos,
	}
}

func NewComment(text string) *ast.Comment {
	return &ast.Comment{
		Slash: token.NoPos,
		Text:  text,
	}
}
func NewCommentGroup(cmt *ast.Comment) *ast.CommentGroup {
	ret := &ast.CommentGroup{}
	ret.List = make([]*ast.Comment, 0)
	ret.List = append(ret.List, cmt)
	return ret
}

//new valueSpec
func NewValueSpec(ident string, expr ast.Expr) *ast.ValueSpec {
	return &ast.ValueSpec{
		Doc:   nil,
		Names: []*ast.Ident{NewIdent(ident)},
		Type:  expr,
		//Type: NewIdent(typeName),
		Values:  nil,
		Comment: nil,
	}
}

func NewCallExpr(f ast.Expr, args []ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:      f,
		Args:     args,
		Lparen:   token.NoPos,
		Ellipsis: token.NoPos,
		Rparen:   token.NoPos,
	}
}

func NewSelectExp(x ast.Expr, sel *ast.Ident) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   x,
		Sel: sel,
	}
}

func NewStarExp(exp ast.Expr) *ast.StarExpr {
	return &ast.StarExpr{
		X: exp,
		//X:NewIdent(ident),
	}
}

func NewKeyValueExp(key, val ast.Expr) *ast.KeyValueExpr {
	return &ast.KeyValueExpr{
		Key:   key,
		Value: val,
		Colon: token.NoPos,
	}
}

/*
example ====>
type svc interface {
  UserGet(ctx context.Context, reqproto *partnerproto.StaffAuthFetchReqProto)(*partnerproto.StaffAuthFetchRespProto, error)
}

====>
typeVals: context.Context
typeName: StartExpr
fieldNames: ctx
tag: nil
*/
func NewField(fieldNames []string, typeVal ast.Expr, typeName ExprType, tag *ast.BasicLit) *ast.Field {
	names := make([]*ast.Ident, 0)
	for _, v := range fieldNames {
		names = append(names, NewIdent(v))
	}
	ret := &ast.Field{
		Doc:   nil,
		Names: names,
		Tag:   tag,
	}
	switch typeName {
	case ExprTypeStartExpr:
		ret.Type = typeVal
	case ExprTypeIdent:
		ret.Type = typeVal
	case ExprTypeSelectorExpr:
		//ret.Type = NewSelectExp(NewIdent(typeVals[0]), NewIdent(typeVals[1]))
		ret.Type = typeVal
	//todo:    other expr
	default:
		ret.Type = typeVal
	}

	return ret

}

func NewFieldOfFuncType(fieldNames []string, functype *ast.FuncType, tag *ast.BasicLit) *ast.Field {
	names := make([]*ast.Ident, 0)
	for _, v := range fieldNames {
		names = append(names, NewIdent(v))
	}
	ret := &ast.Field{
		Doc:   nil,
		Names: names,
		Type:  functype,
		Tag:   tag,
	}

	return ret
}
func NewFieldList(fields ...*ast.Field) *ast.FieldList {
	fieldList := make([]*ast.Field, 0)
	for _, v := range fields {
		fieldList = append(fieldList, v)
	}
	return &ast.FieldList{
		List: fieldList,
	}
}

//new funcDecl
func NewFuncDecl(funcName string, blkStmt *ast.BlockStmt, recv, params *ast.FieldList, commentgroup *ast.CommentGroup) *ast.FuncDecl {
	//recvField := NewField([]string{"h"}, "Receive", true, nil)
	//_ = recvField
	//field := NewField([]string{"a", "b"}, "string", false, nil)
	//commentgroup := NewCommentGroup(NewComment("// this is comment"))
	return &ast.FuncDecl{
		Name: &ast.Ident{Name: funcName},
		Body: blkStmt,
		Type: &ast.FuncType{Func: token.NoPos,
			Params:  params,
			Results: nil,
		},
		Recv: recv,
		Doc:  commentgroup,
	}
}

//new GenDecl
func NewGenDecl(tok token.Token, specs ...ast.Spec) *ast.GenDecl {
	return &ast.GenDecl{
		Doc:    nil,
		TokPos: token.NoPos,
		Tok:    tok,
		Specs:  specs,
	}
}

//new basicLit
func NewBasicLit(kind token.Token, value string) *ast.BasicLit {
	return &ast.BasicLit{
		ValuePos: token.NoPos,
		Kind:     kind,
		Value:    value,
	}
}

//new compositeLit
func NewCompositeLit(typ ast.Expr, elts []ast.Expr) *ast.CompositeLit {
	return &ast.CompositeLit{
		Type:   typ,
		Lbrace: token.NoPos,
		Elts:   elts,
		Rbrace: token.NoPos,
	}
}

func NewAssignStmt(lhs, rhs []ast.Expr) *ast.AssignStmt {

	return &ast.AssignStmt{
		Lhs:    lhs,
		Rhs:    rhs,
		TokPos: token.NoPos,
		Tok:    token.ASSIGN,
	}
}

func NewShortAssignStmt(lhs, rhs []ast.Expr) *ast.AssignStmt {

	return &ast.AssignStmt{
		Lhs:    lhs,
		Rhs:    rhs,
		TokPos: token.NoPos,
		Tok:    token.DEFINE,
	}
}

/*
实例:
package main
func demo(){
	x := 3
}

========>
package main
func demo(){
	x := 3
	var guessCreateEndpoint kitendpoint.Endpoint
}
*/

func NewDeclStmt(decl ast.Decl) *ast.DeclStmt {
	return &ast.DeclStmt{
		Decl: decl,
	}
}

func NewReturnStmt(results []ast.Expr) *ast.ReturnStmt {

	return &ast.ReturnStmt{
		Results: results,
		Return:  token.NoPos,
	}
}

func NewExpStmt(x ast.Expr) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: x,
	}
}
func NewBlockStmt(expStmt ...*ast.ExprStmt) *ast.BlockStmt {
	ret := &ast.BlockStmt{}
	ret.List = make([]ast.Stmt, 0)
	for _, v := range expStmt {
		if expStmt == nil {
			continue
		}
		ret.List = append(ret.List, v)
	}
	return ret
}

func NewEmptyStmt(pos token.Pos) *ast.EmptyStmt {
	ret := &ast.EmptyStmt{}
	ret.Implicit = false
	ret.Semicolon = pos

	return ret
}

//new import spec
func NewImportSpec(name string, path string) *ast.ImportSpec {
	return &ast.ImportSpec{
		Doc:  nil,
		Path: NewBasicLit(token.STRING, path),
		Name: NewIdent(name),
	}
}

func NewInterface(fieldList *ast.FieldList) *ast.InterfaceType {
	return &ast.InterfaceType{
		Interface:  token.NoPos,
		Methods:    fieldList,
		Incomplete: false,
	}
}
