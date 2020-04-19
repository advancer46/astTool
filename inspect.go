package astTool

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/printer"
	"go/token"
	"io"
	"os"
	"reflect"
)

type IdentInfo struct {
	Name string    `description:"ident 名字"`
	Pos  token.Pos `description:"在代码文件中的位置"`
}

var dummyData []IdentInfo //存放Parse()中的缓存数据,以除去冗余ast.Node
var fset *token.FileSet

func init() {
	dummyData = make([]IdentInfo, 0)
}

func (h HappyAst) Position(tp token.Pos) token.Position {
	return h.fileSet.PositionFor(tp, true)
}

func (h HappyAst) FindNodeByPos(pos token.Pos) (ret *ast.Node) {
	if pos == token.NoPos {
		return nil
	}
	ast.Inspect(h.ast, func(node ast.Node) bool {
		switch node.(type) {
		case ast.Node:
			if !node.Pos().IsValid() {
				return true
			}
			if node.Pos() == pos {
				ret = &node
				return false
			}
		}
		return true
	})
	return ret
}

// find function by name,
// return token.Pos
func (h HappyAst) FindIdent(ident string) (pos []token.Pos) {
	ast.Inspect(h.ast, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.Ident:
			if x.Name == ident {
				pos = append(pos, x.Pos())
				return true
			}
		}
		return true
	})
	return pos
}

// find value spec node by name,
// return token.Pos
func (h HappyAst) FindValueSpec(ident string) (pos []token.Pos) {
	ast.Inspect(h.ast, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.ValueSpec:
			for _, v := range x.Names {
				if v.Name == ident {
					pos = append(pos, x.Pos())
					return true
				}
			}
		}
		return true
	})
	return pos
}

// find value spec node by name,
// return token.Pos
func (h HappyAst) FindValueSpecFromNode(n ast.Node, ident string) (pos []token.Pos) {
	ast.Inspect(n, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.ValueSpec:
			for _, v := range x.Names {
				if v.Name == ident {
					pos = append(pos, x.Pos())
					return true
				}
			}
		}
		return true
	})
	return pos
}

// find function by name,
// return token.Pos
func (h HappyAst) FindFuncDeclNode(wantName string) (pos token.Pos) {
	ast.Inspect(h.ast, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.FuncDecl:
			if x.Name.Name == wantName {
				pos = x.Pos()
				return false
			}
		}
		return true
	})

	return
}

// find Struct Decl by name,
// return token.Pos
func (h HappyAst) FindStructDeclNode(wantName string) (pos token.Pos) {
	ast.Inspect(h.ast, func(node ast.Node) bool {
		switch g := node.(type) {
		case *ast.GenDecl:
			if g.Tok == token.TYPE {
				for _, v := range g.Specs {
					switch x := v.(type) {
					case *ast.TypeSpec:
						switch x.Type.(type) {
						case *ast.StructType:
							if x.Name.Name == wantName {
								pos = x.Pos()
								return false
							}
						}
					}
				}
			}
		}
		return true
	})
	return
}

// find Struct Field by name from given Node,
// return token.Pos
func (h HappyAst) FindStructFieldFromNode(n ast.Node, wantName string) (pos token.Pos) {
	ast.Inspect(n, func(node ast.Node) bool {
		switch g := node.(type) {
		case *ast.StructType:
			for _, v := range g.Fields.List {
				for _, val := range v.Names {
					if val.Name == wantName {
						pos = val.Pos()
						return false
					}
				}
			}
		}
		return true
	})
	return
}

// find function call from given Node,
// return token.Pos
func (h HappyAst) FindCallExprFromNode(n ast.Node, funcName string) (pos token.Pos) {
	ast.Inspect(n, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.CallExpr:
			switch i := x.Fun.(type) {
			case *ast.Ident:
				if i.Name == funcName {
					pos = x.Pos()
					return false
				}
			}
		}
		return true
	})

	return
}

// print HappyAst into string
func (h HappyAst) Output() string {
	var buf bytes.Buffer
	printConfig := &printer.Config{Mode: printer.RawFormat, Tabwidth: 4}
	err := printConfig.Fprint(&buf, h.fileSet, h.ast)
	if err != nil {
		panic(err)
	}
	out := buf.Bytes()
	out, err = format.Source(out)
	if err != nil {
		panic(err)
	}
	return string(out)
}

//print
func (h HappyAst) Print() {
	ast.Fprint(os.Stdout, h.fileSet, h.ast, nil)
}

//visitor
func (h HappyAst) Visit() {
	//visitor
	ast.Walk(HappyVisitor(printNode), h.ast)
}

// node Visitor
type HappyVisitor func(ast.Node) string

func (v HappyVisitor) Visit(node ast.Node) ast.Visitor {
	view := v(node)
	//_=view
	fmt.Print(view)
	return v
}

// print node
func printNode(node ast.Node) string {
	if node == nil {
		return ""
	}
	//save
	ptr := node.Pos()
	if line, exists := p.ptrmap[ptr]; exists {
		_ = line
		//printf("(obj @ %d)", line)
		return ""
	} else {
		p.ptrmap[node.Pos()] = node.Pos()
	}
	switch n := node.(type) {
	case *ast.Comment:
		return ""
		//code = n.Text
		//return printLine(code)
	case *ast.CommentGroup:
		return ""
		//for _,c := range n.List{
		//	code += printNode(c)
		//}
		//return printLine(code)
	case *ast.Field:
		return ""
		//for _,v := range n.Names{
		//	code+=printNode(v)
		//}
		//return code + string(tab)
	case *ast.FieldList:
		return ""
		//for _,v := range n.List{
		//	code += printNode(v)
		//}
		//return code + string(tab)
	case *ast.BadExpr:
	case *ast.Ident:
		inner := ""
		if n.Obj != nil {
			inner += n.Obj.Kind.String()
		}
		inner += n.Name
		//return ""
		return inner + string(blank)
	case *ast.BasicLit:
		inner := ""
		inner += n.Kind.String()
		inner += n.Value
		return inner
		//return n.Value+ string(blank)
	case *ast.Ellipsis:
		return ""
		//return printNode(n)+ string(blank)
	case *ast.FuncLit:
		return ""
		//return printNode(n.Body)+ string(newline)
	case *ast.CompositeLit:
		return ""
		//return printNode(n.Type)+ string(blank)
	case *ast.ParenExpr:
		return ""
		//return printNode(n)+ string(blank)
	case *ast.SelectorExpr:
		inner := ""
		//inner += printNode(n.X)
		inner += "."
		//inner += printNode(n.Sel)
		return inner
	case *ast.IndexExpr:
		inner := "*"
		inner += printNode(n.X)
		return inner
	case *ast.SliceExpr:
		inner := "*"
		inner += printNode(n.X)
		return inner
	case *ast.TypeAssertExpr:
		inner := "*"
		//inner += printNode(n.Type)
		//inner += printNode(n.X)
		return inner
	case *ast.CallExpr:
	case *ast.StarExpr:
		inner := "*"
		inner += printNode(n.X)
		return inner
	case *ast.UnaryExpr:
		inner := ""
		inner += printNode(n.X)
		return inner
	case *ast.BinaryExpr:
		inner := ""
		inner += printNode(n.X)
		inner += printNode(n.Y)
		return inner
	case *ast.KeyValueExpr:
	case *ast.ArrayType:
	case *ast.StructType:
	case *ast.FuncType:
	case *ast.InterfaceType:
	case *ast.MapType:
	case *ast.ChanType:
	case *ast.BadStmt:
	case *ast.DeclStmt:
	case *ast.EmptyStmt:
	case *ast.LabeledStmt:
	case *ast.ExprStmt:
	case *ast.SendStmt:
	case *ast.IncDecStmt:
	case *ast.AssignStmt:
	case *ast.GoStmt:
	case *ast.DeferStmt:
	case *ast.ReturnStmt:
	case *ast.BranchStmt:
	case *ast.BlockStmt:
	case *ast.IfStmt:
	case *ast.CaseClause:
	case *ast.SwitchStmt:
	case *ast.TypeSwitchStmt:
	case *ast.CommClause:
	case *ast.SelectStmt:
	case *ast.ForStmt:
	case *ast.RangeStmt:
	case *ast.ImportSpec:

		return n.Path.Value + string(newline)
	case *ast.ValueSpec:
		inner := ""
		for _, v := range n.Names {
			inner += v.Name + " "
		}
		return inner + string(newline)
	case *ast.TypeSpec:

	case *ast.BadDecl:
	case *ast.GenDecl:
		return inspectGenDecl(n)
	case *ast.FuncDecl:
		return inspectFuncDecl(n)
	case *ast.File:
		//packageName := n.Name.Name
		inner := ""
		inner += printNode(n.Name)
		return "package " + inner + string(newline)
	case *ast.Package:
		return ""

	default:
		_ = n
	}
	return ""
}

type mprinter struct {
	output io.Writer
	fset   *token.FileSet
	filter ast.FieldFilter
	ptrmap map[interface{}]interface{} // *T -> line number
	//ptrmap map[interface{}]int // *T -> line number
	indent int  // current indentation level
	last   byte // the last byte processed by Write
	line   int  // current line number
}

//var p = mprinter{os.Stdout,fset,nil,0,byte(0),0}
var p = mprinter{
	output: os.Stdout,
	fset:   fset,
	filter: nil,
	ptrmap: make(map[interface{}]interface{}),
	last:   '\n', // force printing of line number on first line
}

func print(x reflect.Value) {

	switch x.Kind() {
	case reflect.Interface:
		print(x.Elem())

	case reflect.Map:
		printf("%s (len = %d) {", x.Type(), x.Len())
		if x.Len() > 0 {
			p.indent++
			printf("\n")
			for _, key := range x.MapKeys() {
				print(key)
				printf(": ")
				print(x.MapIndex(key))
				printf("\n")
			}
			p.indent--
		}
		printf("}")

	case reflect.Ptr:
		printf("*")
		// type-checked ASTs may contain cycles - use ptrmap
		// to keep track of objects that have been printed
		// already and print the respective line number instead
		ptr := x.Interface()
		if line, exists := p.ptrmap[ptr]; exists {
			printf("(obj @ %d)", line)
		} else {
			p.ptrmap[ptr] = p.line
			print(x.Elem())
		}

	case reflect.Array:
		printf("%s {", x.Type())
		if x.Len() > 0 {
			p.indent++
			printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				printf("%d: ", i)
				print(x.Index(i))
				printf("\n")
			}
			p.indent--
		}
		printf("}")

	case reflect.Slice:
		if s, ok := x.Interface().([]byte); ok {
			printf("%#q", s)
			return
		}
		printf("%s (len = %d) {", x.Type(), x.Len())
		if x.Len() > 0 {
			p.indent++
			printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				printf("%d: ", i)
				print(x.Index(i))
				printf("\n")
			}
			p.indent--
		}
		printf("}")

	case reflect.Struct:
		t := x.Type()

		printf("%s {", t)
		p.indent++
		first := true
		for i, n := 0, t.NumField(); i < n; i++ {
			// exclude non-exported fields because their
			// values cannot be accessed via reflection
			if name := t.Field(i).Name; ast.IsExported(name) {
				value := x.Field(i)
				if first {
					printf("\n")
					first = false
				}
				printf("%s: ", name)
				print(value)
				printf("\n")
			}
		}
		p.indent--
		printf("}")
		////test
		//xxxy := x.Interface().(ast.Node)
		//switch xxxy.(type){
		//case *ast.File:
		//	fmt.Println("=================================file")
		//default:
		//
		//}
		////test end

	default:
		v := x.Interface()
		switch v := v.(type) {
		case string:
			// print strings in quotes
			printf("%q", v)
			return
		case token.Pos:
			// position values can be printed nicely if we have a file set
			if p.fset != nil {
				printf("%s", p.fset.Position(v))
				return
			}
		}
		// default
		printf("%v", v)
	}
}

// printf is a convenience wrapper that takes care of print errors.
func printf(format string, args ...interface{}) {
	if _, err := fmt.Fprintf(p.output, format, args...); err != nil {
		panic(err)
	}
}

//从ast.Node中根据函数名查找指定函数
func findFuncByName(node ast.Node) bool {
	//ast.Print(fset,node)
	//ast.Fprint(os.Stdout,fset,node,nil)
	var wantName string
	wantName = "registerBrandcustomerClientConn"
	switch x := node.(type) {
	case *ast.FuncDecl:
		for _, v := range x.Body.List {
			switch decl := v.(type) {
			case *ast.DeclStmt:
				findFuncByName(decl.Decl)
			case *ast.ExprStmt:
				switch exp := decl.X.(type) {
				case *ast.CallExpr:
					funcName, pos := parseCallExpr(exp.Fun)
					if funcName == wantName {
						dummyData = append(dummyData, IdentInfo{
							Name: funcName,
							Pos:  pos,
						})
					}
				}
			}
		}
	}

	return true
}

func parseCallExpr(cexpr ast.Expr) (string, token.Pos) {
	switch x := cexpr.(type) {
	case *ast.FuncLit:
	case *ast.Ident:
		return x.Name, x.Pos()
	}
	return "", cexpr.Pos()
}

// 分析函数声明
func inspectFuncDecl(n *ast.FuncDecl) (declStr string) {
	funcName := fmt.Sprintf("%s", n.Name)
	if n.Type.Params != nil {
		paramList := n.Type.Params.List
		declStr = fmt.Sprintf("Func: %s\n", funcName)
		for k, v := range paramList {
			for _, val := range v.Names {
				declStr += fmt.Sprintf("param%d %s %s\n", k, val.Obj.Name, val.Obj.Decl.(*ast.Field).Type)
			}
		}
	}
	if n.Type.Results != nil {
		resultList := n.Type.Results.List
		for k, v := range resultList {
			for _, val := range v.Names {
				declStr += fmt.Sprintf("return%d %s %s\n", k, val.Obj.Name, val.Obj.Decl.(*ast.Field).Type)
			}
		}
	}

	return
}

// 分析通用声明
func inspectGenDecl(n *ast.GenDecl) (declStr string) {
	declStr += n.Tok.String() + " "
	for _, v := range n.Specs {
		declStr += printNode(v)
	}

	return
}

func inspectNameOfSpec(n ast.Spec) []*ast.Ident {

	switch x := n.(type) {
	case *ast.ValueSpec:
		return x.Names
	case *ast.TypeSpec:
		ret := make([]*ast.Ident, 0)
		ret = append(ret, x.Name)
		return ret
	}
	return make([]*ast.Ident, 0)
}
