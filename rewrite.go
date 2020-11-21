package astTool

/*
	rewrite node
*/

import (
	"errors"
	"go/ast"
	"go/token"
	"log"
)

const (
	HEAD = 0
	TAIL = -1
)

///*
//	替换节点
//*/
//func (h *HappyAst) ReplaceNode(pos token.Pos, newNode ast.Node) error {
//	if pos == token.NoPos {
//		return errors.New("pos not exist")
//	}
//	ast.Inspect(h.Ast, func(node ast.Node) bool {
//		switch node.(type) {
//		case ast.Node:
//			if !node.Pos().IsValid() {
//				return true
//			}
//			if node.Pos() == pos {
//				node = newNode
//				return false
//			}
//		}
//		return true
//	})
//	return nil
//}

//
///*
//	替换节点
//*/
//func (h *HappyAst) ReplaceNode(pos token.Pos, newNode ast.Node) error {
//	if pos == token.NoPos {
//		return errors.New("pos not exist")
//	}
//	ast.Inspect(h.Ast, func(node ast.Node) bool {
//		if node.Pos().IsValid() && node.Pos() == pos {
//			node = newNode
//			return false
//		}
//
//		return true
//	})
//	return nil
//}

/*
	替换节点
*/
func (h *HappyAst) ReplaceNode(pos token.Pos, newNode ast.Node) error {
	if pos == token.NoPos {
		return errors.New("pos not exist")
	}
	ast.Inspect(h.Ast, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.Ident:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.Ident)
				return false
			}
		case *ast.BasicLit:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.BasicLit)
				return false
			}
		case *ast.File:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.File)
				return false
			}
		case *ast.ImportSpec:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.ImportSpec)
				return false
			}
		case *ast.Package:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.Package)
				return false
			}
		case *ast.AssignStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.AssignStmt)
				return false
			}
		case *ast.RangeStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.RangeStmt)
				return false
			}
		case *ast.BranchStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.BranchStmt)
				return false
			}
		case *ast.BinaryExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.BinaryExpr)
				return false
			}
		case *ast.IncDecStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.IncDecStmt)
				return false
			}
		case *ast.ForStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.ForStmt)
				return false
			}
		case *ast.SelectorExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.SelectorExpr)
				return false
			}
		case *ast.CallExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.CallExpr)
				return false
			}
		case *ast.ReturnStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.ReturnStmt)
				return false
			}
		case *ast.BlockStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.BlockStmt)
				return false
			}
		case *ast.IfStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.IfStmt)
				return false
			}
		case *ast.InterfaceType:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.InterfaceType)
				return false
			}
		case *ast.FuncType:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.FuncType)
				return false
			}
		case *ast.ChanType:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.ChanType)
				return false
			}
		case *ast.MapType:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.MapType)
				return false
			}
		case *ast.ArrayType:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.ArrayType)
				return false
			}
		case *ast.TypeSpec:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.TypeSpec)
				return false
			}
		case *ast.GenDecl:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.GenDecl)
				return false
			}
		case *ast.UnaryExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.UnaryExpr)
				return false
			}
		case *ast.ValueSpec:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.ValueSpec)
				return false
			}
		case *ast.DeclStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.DeclStmt)
				return false
			}
		case *ast.FuncDecl:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.FuncDecl)
				return false
			}
		case *ast.KeyValueExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.KeyValueExpr)
				return false
			}
		case *ast.BadDecl:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.BadDecl)
				return false
			}
		case *ast.BadExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.BadExpr)
				return false
			}
		case *ast.CaseClause:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.CaseClause)
				return false
			}
		case *ast.CommClause:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.CommClause)
				return false
			}
		case *ast.CompositeLit:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.CompositeLit)
				return false
			}
		case *ast.DeferStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.DeferStmt)
				return false
			}
		case *ast.Ellipsis:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.Ellipsis)
				return false
			}
		case *ast.EmptyStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.EmptyStmt)
				return false
			}
		case *ast.ExprStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.ExprStmt)
				return false
			}
		case *ast.Field:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.Field)
				return false
			}
		case *ast.FieldList:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.FieldList)
				return false
			}
		case *ast.SwitchStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.SwitchStmt)
				return false
			}
		case *ast.TypeAssertExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.TypeAssertExpr)
				return false
			}
		case *ast.TypeSwitchStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.TypeSwitchStmt)
				return false
			}
		case *ast.StarExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.StarExpr)
				return false
			}
		case *ast.SelectStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.SelectStmt)
				return false
			}
		case *ast.IndexExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.IndexExpr)
				return false
			}
		case *ast.LabeledStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.LabeledStmt)
				return false
			}
		case *ast.GoStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.GoStmt)
				return false
			}
		case *ast.SendStmt:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.SendStmt)
				return false
			}
		case *ast.ParenExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.ParenExpr)
				return false
			}
		case *ast.SliceExpr:
			if node.Pos().IsValid() && node.Pos() == pos {
				*x = *newNode.(*ast.SliceExpr)
				return false
			}
		default:

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
		for _, v := range bmt.List {
			tempStmtList = append(tempStmtList, v)
		}
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

func (h *HappyAst) FreshPosInfoOfCommentGroup() error {

	searcher := Searcher{h.Ast}

	commentGroups := searcher.FindCommentGroups()

	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			err := h.ReplaceNode(comment.Pos(), &ast.Comment{
				Slash: token.NoPos,
				Text:  comment.Text,
			})
			if err != nil {
				log.Printf("err :%s[FreshPosInfoOfCommentGroup %s]", err.Error(), comment.Text)
			}
		}
	}
	return nil
}
