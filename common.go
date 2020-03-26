
package astTool

import (
	"go/ast"
	"go/token"
)

func NewIdent(name string) *ast.Ident{
	return &ast.Ident{
		Name: name,
		Obj:nil,
		NamePos:token.NoPos,
	}
}

func NewComment(text string) *ast.Comment{
	return &ast.Comment{
		Slash:token.NoPos,
		Text: text,
	}
}
func NewCommentGroup(cmt *ast.Comment) *ast.CommentGroup{
	ret := &ast.CommentGroup{}
	ret.List = make([]*ast.Comment,0)
	ret.List = append(ret.List,cmt)
	return ret
}
//new valueSpec
func  NewValueSpec(ident string, expr ast.Expr) *ast.ValueSpec{
	return &ast.ValueSpec{
		Doc: nil,
		Names: []*ast.Ident{NewIdent(ident)},
		Type: expr,
		//Type: NewIdent(typeName),
		Values:nil,
		Comment: nil,
	}
}

func NewCallExpr(funcName string) *ast.CallExpr{
	return &ast.CallExpr{
		Fun: NewIdent(funcName),
		Args: []ast.Expr{
			//NewBasicLit(token.INT,"42"),
		},
	}
}
func NewExpStmt(x ast.Expr) *ast.ExprStmt{
	return &ast.ExprStmt{
		X: x,
	}
}
func NewBlockStmt(expStmt ...*ast.ExprStmt) ( *ast.BlockStmt){
	ret := &ast.BlockStmt{}
	ret.List = make([]ast.Stmt,0)
	for _,v :=range expStmt{
		if expStmt == nil{
			continue
		}
		ret.List = append(ret.List,v)
	}
	return ret
}

func NewSelectExp(exp ast.Expr,ident *ast.Ident) *ast.SelectorExpr{
	return &ast.SelectorExpr{
		X:exp,
		Sel:ident,
		//X:NewIdent(ident),
	}
}

func NewStarExp(exp ast.Expr) *ast.StarExpr{
	return &ast.StarExpr{
		X:exp,
		//X:NewIdent(ident),
	}
}
func NewField(fieldNames []string,typeName string,isStartExpr bool,tag *ast.BasicLit) *ast.Field{
	names := make([]*ast.Ident,0)
	for _,v := range fieldNames{
		names = append(names,NewIdent(v))
	}
	ret := &ast.Field{
		Doc:nil,
		Names:names,
		//Tag:NewBasicLit(token.STRING,"`json:\"brandcustomer_service_host\"`"),
		Tag:tag,
	}
	if isStartExpr {
		ret.Type = NewStarExp(NewIdent(typeName))
	}else{
		ret.Type = NewIdent(typeName)
	}

	return ret

}
func NewFieldList(fields ...*ast.Field) *ast.FieldList{
	fieldList := make([]*ast.Field,0)
	for _,v := range fields{
		fieldList = append(fieldList, v)
	}
	return &ast.FieldList{
		List: fieldList,
	}
}


//new funcDecl
func NewFuncDecl(funcName string,blkStmt *ast.BlockStmt) *ast.FuncDecl{
	recvField :=  NewField([]string{"h",},"Receive",true,nil)
	_=recvField
	field := NewField([]string{"a","b"},"string",false,nil)
	return &ast.FuncDecl{
		Name: &ast.Ident{Name: funcName},
		Body: blkStmt,
		Type: &ast.FuncType{Func:token.NoPos,
			Params:NewFieldList(field),
			Results:nil,
		},
		Recv:NewFieldList(recvField),
		Doc: NewCommentGroup(NewComment("// this is comment")),

	}
}

//new GenDecl
func NewGenDecl(tok token.Token, specs ...ast.Spec) *ast.GenDecl{
	return &ast.GenDecl{
		Doc:nil,
		TokPos: token.NoPos,
		Tok: tok,
		Specs:specs,
	}
}

//new basicLit
func NewBasicLit( kind token.Token, value string) *ast.BasicLit{
	return &ast.BasicLit{
		ValuePos: token.NoPos,
		Kind: kind,
		Value: value,
	}
}

//new import spec
func NewImportSpec(name string, path string) *ast.ImportSpec{
	return &ast.ImportSpec{
		Doc:nil,
		Path: NewBasicLit(token.STRING,path),
		Name: NewIdent(name),
	}
}