package astTool

import (
	"errors"
	"go/ast"
	"go/token"
)

const (
	HEAD = 0
	TAIL = -1
)

/*
	替换节点
*/
func (h *HappyAst) ReplaceNode(pos token.Pos, newNode ast.Node) error {
	if pos == token.NoPos {
		return errors.New("pos not exist")
	}
	ast.Inspect(h.Ast, func(node ast.Node) bool {
		switch node.(type) {
		case ast.Node:
			if !node.Pos().IsValid() {
				return true
			}
			if node.Pos() == pos {
				node = newNode
				return false
			}
		}
		return true
	})
	return nil
}

/*
	添加语句
	@deprecated
*/
func (h *HappyAst) AddStmt(bmt *ast.BlockStmt, location int, stmt ast.Stmt) {

	switch location {
	case HEAD:
		tempStmtList := make([]ast.Stmt, 0)
		tempStmtList = append(tempStmtList, stmt)
		for _, v := range bmt.List {
			tempStmtList = append(tempStmtList, v)
		}
		bmt.List = tempStmtList

	case TAIL:
		tempStmtList := make([]ast.Stmt, 0)
		tempStmtList = append(tempStmtList, stmt)
		bmt.List = tempStmtList
	default:
		tempStmtList := make([]ast.Stmt, 0)
		tempStmtList = append(tempStmtList, bmt.List[:location]...)
		tempStmtList = append(tempStmtList, stmt)
		tempStmtList = append(tempStmtList, bmt.List[location:]...)
		bmt.List = tempStmtList
	}

}

/*
	添加赋值语句
	bmt 语句块节点
	location 在bmt.List数组中的位置
	stmt 待放置的赋值语句节点
*/
func (h *HappyAst) AddAssignStmt(bmt *ast.BlockStmt, location int, stmt ast.Stmt) {
	tempStmtList := make([]ast.Stmt, 0)

	switch location {
	case HEAD:
		tempStmtList = append(tempStmtList, stmt)
		tempStmtList = append(tempStmtList, bmt.List[:]...)
	case TAIL:
		tempStmtList = append(tempStmtList, bmt.List[:]...)
		tempStmtList = append(tempStmtList, stmt)
	default:
		tempStmtList = append(tempStmtList, bmt.List[:location]...)
		tempStmtList = append(tempStmtList, stmt)
		tempStmtList = append(tempStmtList, bmt.List[location:]...)
	}

	bmt.List = tempStmtList
}

/*
	添加函数声明
	location 在root.Decls数组中的位置
	stmt 待放置的赋值语句节点
*/
func (h *HappyAst) AppendFundDecl(location int, decl ast.Decl) {
	tempDeclList := make([]ast.Decl, 0)

	switch location {
	case HEAD:
		tempDeclList = append(tempDeclList, decl)
		tempDeclList = append(tempDeclList, h.Ast.Decls[:]...)
	case TAIL:
		tempDeclList = append(tempDeclList, h.Ast.Decls[:]...)
		tempDeclList = append(tempDeclList, decl)
	default:
		tempDeclList = append(tempDeclList, h.Ast.Decls[:location]...)
		tempDeclList = append(tempDeclList, decl)
		tempDeclList = append(tempDeclList, h.Ast.Decls[location:]...)
	}
	h.Ast.Decls = tempDeclList
}

/*
	添加声明
	@deprecated
*/
func (h *HappyAst) AddDecl(location int, appends ...ast.Decl) {

	tempDecls := make([]ast.Decl, 0)
	switch location {
	case HEAD:
		for _, v := range appends {
			tempDecls = append(tempDecls, v)
		}
		for _, val := range h.Ast.Decls {
			tempDecls = append(tempDecls, val)
		}
	case TAIL:
		tempDecls = h.Ast.Decls
		for _, v := range appends {
			tempDecls = append(tempDecls, v)
		}
	default:
		tempDecls = h.Ast.Decls
		for _, v := range appends {
			tempDecls = append(tempDecls, v)
		}
	}

	h.Ast.Decls = tempDecls
}

/*
type svc interface {
  UserGet()
}

=>

type svc interface {
  UserGet()
  RoleGet()
}
*/
func (h *HappyAst) AddFieldOfFuncType(bmt *ast.FieldList, location int, field *ast.Field) {
	tempFieldList := make([]*ast.Field, 0)

	switch location {
	case HEAD:
		tempFieldList = append(tempFieldList, field)
		tempFieldList = append(tempFieldList, bmt.List[:]...)
	case TAIL:
		tempFieldList = append(tempFieldList, bmt.List[:]...)
		tempFieldList = append(tempFieldList, field)
	default:
		tempFieldList = append(tempFieldList, bmt.List[:location]...)
		tempFieldList = append(tempFieldList, field)
		tempFieldList = append(tempFieldList, bmt.List[location:]...)
	}

	bmt.List = tempFieldList
}

/*
参数:
bmt,被插入的fieldList
location,在bmt中的顺序. -1尾部添加  0头部添加
field,插入的field

示例:
type PartnerSvcEndpoints struct {
	modelFetchEndpoint  kitendpoint.Endpoint
}

=>

type PartnerSvcEndpoints struct {
	modelFetchEndpoint  kitendpoint.Endpoint
	gameFetchEndpoint  kitendpoint.Endpoint
}
*/
func (h *HappyAst) AddField(bmt *ast.FieldList, location int, field *ast.Field) {
	tempFieldList := make([]*ast.Field, 0)

	switch location {
	case HEAD:
		tempFieldList = append(tempFieldList, field)
		tempFieldList = append(tempFieldList, bmt.List[:]...)
	case TAIL:
		tempFieldList = append(tempFieldList, bmt.List[:]...)
		tempFieldList = append(tempFieldList, field)
	default:
		tempFieldList = append(tempFieldList, bmt.List[:location]...)
		tempFieldList = append(tempFieldList, field)
		tempFieldList = append(tempFieldList, bmt.List[location:]...)
	}

	bmt.List = tempFieldList
}
