package main

import (
	"fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "os"
	)

	// fmt.Println(google.TransformString("abcd", "a"))
	// fmt.Println(IsValidSudoku([][]byte{{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}}))

// func day1(){
	
// }

/* 
1. Create an AST by parsing src in the neetcode
2. Walk the AST and look for all the function declarations
3. For each function declaration, create a new Run object
4. Run the Run object by nameing the function and passing the arguments
*/

func main() {
    fmt.Println('5'-'0')
	subPackage := "neetcode"
    set := token.NewFileSet()
    packs, err := parser.ParseDir(set, subPackage, nil, parser.ParseComments)
    if err != nil {
        fmt.Println("Failed to parse package:", err)
        os.Exit(1)
    }

    funcs := []*ast.FuncDecl{}
    for _, pack := range packs {
        for _, f := range pack.Files {
            for _, d := range f.Decls {
                if fn, isFn := d.(*ast.FuncDecl); isFn {
                    funcs = append(funcs, fn)
                }
            }
        }
    }

    fmt.Printf("all funcs: %+v\n", funcs)

}