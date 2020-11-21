package astTool

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

const (
	TYPE_IDENT     = "IDENT"
	TYPE_BASICLIT  = "BASICLIT"
	TYPE_TYPESPEC  = "TYPESPEC"
	TYPE_VALUESPEC = "VALUESPEC"

	KIND_VAR    = "var"    // Ident in Ident
	KIND_TYPE   = "type"   // Ident in TypeSpec
	KIND_STRING = "STRING" //basiclit
)

type HappyAst struct {
	FileSet *token.FileSet //file info
	Ast     *ast.File      //collection of ast.Node
}

func DemoInspect(srcode string) {

	h := ParseFromCode(srcode)

	h.Print()
	//h.Visit()
	//fmt.Println(h.Output())
	pos := h.FindStructDeclNode("Microservice")
	fmt.Println(h.Position(pos))
}

func DemoRewrite(codeFile string) {
	srcode := `package miclient
type Microservice struct {
ActivityHost           string ` + "`" + `json:"activity_service_host"` + "`" + `
}`
	h := ParseFromCode(srcode)
	if codeFile != "" {
		h = ParseFromFile(codeFile)
	}
	tag := NewBasicLit(token.STRING, "`json:\"brandcustomer_service_host\"`")
	field := NewField([]string{"BrandHost"}, NewIdent("string"), ExprTypeIdent, tag)
	gpos := h.FindStructDeclNode("Microservice")
	gnode := h.NodeByPos(gpos)
	structNode := gnode.(*ast.TypeSpec).Type.(*ast.StructType)
	structNode.Fields.List = append(structNode.Fields.List, field)
	h.ReplaceNode(gpos, gnode)
	fmt.Println(h.Output(nil))
}

func ParseFromFile(codeFile string) *HappyAst {
	fset = token.NewFileSet()
	//astNodes, err := parser.ParseFile(fset, codeFile, nil,parser.ParseComments)
	astNodes, err := parser.ParseFile(fset, codeFile, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	h := &HappyAst{
		FileSet: fset,
		Ast:     astNodes,
	}
	return h
}

func ParseFromCode(srccode string) *HappyAst {
	fset = token.NewFileSet()
	//astNodes, err := parser.ParseFile(fset, codeFile, nil,parser.ParseComments)
	astNodes, err := parser.ParseFile(fset, "", srccode, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	h := &HappyAst{
		FileSet: fset,
		Ast:     astNodes,
	}
	return h
}

// 根据参数查找节点位置
// name:　节点名称,不支持模糊查询
// tokenType:　token.go规定的宏变量,　ast.Ident,ast.BasicLit等等
// tokenKind:　节点种类,var,int,string等
// lineStart:　位置范围查询的开始,-1表示忽略
// lineEnd:　位置返回查询的截止,-1表示忽略
func (h HappyAst) PosOfNode(name string, tokenType string, tokenKind string, lineStart int, lineEnd int) (pos []token.Pos) {
	ast.Inspect(h.Ast, func(node ast.Node) bool {

		xname := ""
		xtype := ""
		xline := 0
		xkind := ""

		switch x := node.(type) {
		case *ast.Ident:
			xname = x.Name
			xtype = TYPE_IDENT
			xline = h.Position(x.Pos()).Line
			xkind = ""
			if x.Obj != nil {
				xkind = x.Obj.Kind.String()
			}
		case *ast.BasicLit:
			xname = x.Value
			xtype = TYPE_BASICLIT
			xline = h.Position(x.Pos()).Line
			xkind = x.Kind.String()
		}

		// filter
		if name != "" && xname != name {
			return true
		}

		if tokenType != "" && xtype != tokenType {
			return true
		}
		if tokenKind != "" && xkind != tokenKind {
			return true
		}
		if lineStart >= 0 && xline < lineStart {
			return true
		}
		if lineEnd >= 0 && xline > lineEnd {
			return true
		}
		pos = append(pos, node.Pos())
		return true
	})
	return pos
}

func (h HappyAst) NodeByPos(pos token.Pos) (ret ast.Node) {
	if pos == token.NoPos {
		return nil
	}
	ast.Inspect(h.Ast, func(node ast.Node) bool {

		switch x := node.(type) {
		case *ast.Ident:
			if !node.Pos().IsValid() {
				return true
			}
			if node.Pos() == pos {
				ret = x
				return false
			}
		case *ast.BasicLit:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.File:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.ImportSpec:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.Package:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.AssignStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.RangeStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.BranchStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.BinaryExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.IncDecStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.ForStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.SelectorExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.CallExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.ReturnStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.BlockStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.IfStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.InterfaceType:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.FuncType:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.ChanType:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.MapType:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.StructType:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.ArrayType:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.TypeSpec:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.GenDecl:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.UnaryExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.ValueSpec:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.DeclStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.FuncDecl:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.KeyValueExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.BadDecl:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.BadExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.CaseClause:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.CommClause:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.CompositeLit:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.DeferStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.Ellipsis:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.EmptyStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.ExprStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.Field:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.FieldList:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.SwitchStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.TypeAssertExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.TypeSwitchStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.StarExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.SelectStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.IndexExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.LabeledStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.GoStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.SendStmt:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.ParenExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		case *ast.SliceExpr:
			if !node.Pos().IsValid() {
				return true
			}
			if x.Pos() == pos {
				ret = x
				return false
			}
		default:
			return true
		}
		return true
	})
	return ret
}
