package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to print the AST.
	src := `
package main
func main() {
	println("Hello, World!")
}
`
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}

/*output:

   0  *ast.File {
   1  .  Package: 2:1
   2  .  Name: *ast.Ident {
   3  .  .  NamePos: 2:9
   4  .  .  Name: "main"
   5  .  }
   6  .  Decls: []ast.Decl (len = 1) {
   7  .  .  0: *ast.FuncDecl {
   8  .  .  .  Name: *ast.Ident {
   9  .  .  .  .  NamePos: 3:6
  10  .  .  .  .  Name: "main"
  11  .  .  .  .  Obj: *ast.Object {
  12  .  .  .  .  .  Kind: func
  13  .  .  .  .  .  Name: "main"
  14  .  .  .  .  .  Decl: *(obj @ 7)
  15  .  .  .  .  }
  16  .  .  .  }
  17  .  .  .  Type: *ast.FuncType {
  18  .  .  .  .  Func: 3:1
  19  .  .  .  .  Params: *ast.FieldList {
  20  .  .  .  .  .  Opening: 3:10
  21  .  .  .  .  .  Closing: 3:11
  22  .  .  .  .  }
  23  .  .  .  }
  24  .  .  .  Body: *ast.BlockStmt {
  25  .  .  .  .  Lbrace: 3:13
  26  .  .  .  .  List: []ast.Stmt (len = 1) {
  27  .  .  .  .  .  0: *ast.ExprStmt {
  28  .  .  .  .  .  .  X: *ast.CallExpr {
  29  .  .  .  .  .  .  .  Fun: *ast.Ident {
  30  .  .  .  .  .  .  .  .  NamePos: 4:2
  31  .  .  .  .  .  .  .  .  Name: "println"
  32  .  .  .  .  .  .  .  }
  33  .  .  .  .  .  .  .  Lparen: 4:9
  34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
  35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
  36  .  .  .  .  .  .  .  .  .  ValuePos: 4:10
  37  .  .  .  .  .  .  .  .  .  Kind: STRING
  38  .  .  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
  39  .  .  .  .  .  .  .  .  }
  40  .  .  .  .  .  .  .  }
  41  .  .  .  .  .  .  .  Ellipsis: -
  42  .  .  .  .  .  .  .  Rparen: 4:25
  43  .  .  .  .  .  .  }
  44  .  .  .  .  .  }
  45  .  .  .  .  }
  46  .  .  .  .  Rbrace: 5:1
  47  .  .  .  }
  48  .  .  }
  49  .  }
  50  .  Scope: *ast.Scope {
  51  .  .  Objects: map[string]*ast.Object (len = 1) {
  52  .  .  .  "main": *(obj @ 11)
  53  .  .  }
  54  .  }
  55  .  Unresolved: []*ast.Ident (len = 1) {
  56  .  .  0: *(obj @ 29)
  57  .  }
  58  }


  remark:
Package: 2:1代表Go解析出package这个词在第二行的第一个
main是一个ast.Ident标识符，它的位置在第二行的第9个
此处func main被解析成ast.FuncDecl（function declaration）,而函数的参数（Params）和函数体（Body）
自然也在这个FuncDecl中。Params对应的是*ast.FieldList，顾名思义就是项列表；
而由大括号“｛｝”组成的函数体对应的是ast.BlockStmt（block statement） 而对于main函数的函数体中，
我们可以看到调用了println函数，在ast中对应的是ExprStmt（Express Statement），
调用函数的表达式对应的是CallExpr(Call Expression)，调用的参数自然不能错过，因为参数只有字符串，所以go把它归为ast.BasicLis (a literal of basic type)。
最后，我们可以看出ast还解析出了函数的作用域，以及作用域对应的对象。
Go将所有可以识别的token抽象成Node，通过interface方式组织在一起。
在这里说到token我们需要说一下词法分析，token是词法分析的结果，即将字符序列转换为标记(token)的过程，
这个操作由词法分析器完成。这里的标记是一个字符串，是构成源代码的最小单位。


*/
