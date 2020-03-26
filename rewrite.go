
package astTool

import (
	"go/ast"
	"go/token"
	"errors"
)
const (
	HEAD = "head"
	TAIL = "tail"
)
func (h *HappyAst) ReplaceNode(pos token.Pos,newNode ast.Node)error{
	if pos == token.NoPos {
		return errors.New("pos not exist")
	}
	ast.Inspect(h.ast, func(node ast.Node)bool{
		switch node.(type){
		case ast.Node:
			if !node.Pos().IsValid() {
				return true
			}
			if  node.Pos() == pos{
				node = newNode
				return false
			}
		}
		return true
	})
	return nil
}

// add stmt
func (h *HappyAst) AddStmt(bmt *ast.BlockStmt,location string,stmt ast.Stmt){
	if location == HEAD{
		tempStmtList := make([]ast.Stmt,0)
		tempStmtList = append(tempStmtList,stmt)
		for _,v := range bmt.List{
			tempStmtList = append(tempStmtList,v)
		}
	}else{
		tempStmtList := bmt.List
		tempStmtList = append(tempStmtList,stmt)
		bmt.List = tempStmtList
	}
}
//add decl
func (h *HappyAst) AddDecl(location string, appends ...ast.Decl) {
	if location == HEAD{
		tempDecls := make([]ast.Decl,0)
		for _,v := range appends{
			tempDecls = append(tempDecls,v)
		}
		for _,val := range h.ast.Decls{
			tempDecls = append(tempDecls,val)
		}
		h.ast.Decls = tempDecls
		return
	}
	if location == TAIL{
		tempDecls := make([]ast.Decl,0)
		tempDecls = h.ast.Decls
		for _,v := range appends{
			tempDecls = append(tempDecls, v)
		}
		h.ast.Decls = tempDecls
		return
	}
}

