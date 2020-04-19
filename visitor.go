package astTool

import (
	"go/ast"
)

/**
提供方法查看或者搜素ast节点
*/
type MyVisitor struct {
	Result []ast.Node
	Name   string
	Type   string
}

func (v MyVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return v
	}
	return v
}

func (v *MyVisitor) Inspecter(node ast.Node) bool {
	switch x := node.(type) {
	case *ast.FuncDecl:
		if v.Name == x.Name.Name && v.Type == "FuncDecl" {
			v.Result = append(v.Result, x)
			return false
		}
	case *ast.ValueSpec:
		for _, item := range x.Names {
			if v.Name == item.Name && v.Type == "ValueSpec" {
				v.Result = append(v.Result, x)
				return false
			}
		}
	}
	return true
}

/**
搜索器
按照节点名，节点类型
*/
type Searcher struct {
	Root ast.Node
}

func (v Searcher) FindFuncDecl(name string) *ast.FuncDecl {
	visitor := MyVisitor{Result: make([]ast.Node, 0), Name: name, Type: "FuncDecl"}

	ast.Inspect(v.Root, visitor.Inspecter)
	if len(visitor.Result) > 0 {
		return visitor.Result[0].(*ast.FuncDecl)
	} else {
		return nil
	}
}

func (v Searcher) FindValueSpec(name string) *ast.ValueSpec {
	visitor := MyVisitor{Result: make([]ast.Node, 0), Name: name, Type: "ValueSpec"}

	ast.Inspect(v.Root, visitor.Inspecter)
	if len(visitor.Result) > 0 {
		return visitor.Result[0].(*ast.ValueSpec)
	} else {
		return nil
	}
}

func (v Searcher) FindValueSpec(name string) *ast.ValueSpec {
	visitor := MyVisitor{Result: make([]ast.Node, 0), Name: name, Type: "ValueSpec"}

	ast.Inspect(v.Root, visitor.Inspecter)
	if len(visitor.Result) > 0 {
		return visitor.Result[0].(*ast.ValueSpec)
	} else {
		return nil
	}
}
