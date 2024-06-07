package main

import (
	"fmt"
	// "os"

	"repl-experiment/parser"

	"github.com/peterh/liner"
	"github.com/antlr4-go/antlr/v4"
)

type exprListener struct {
	*parser.BaseExprListener
}

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range []string{"help", "exit", "print"} {
			if len(line) <= len(n) && n[:len(line)] == line {
				c = append(c, n)
			}
		}
		return
	})

	for {
		if input, err := line.Prompt("> "); err == nil {
			line.AppendHistory(input)

            is := antlr.NewInputStream(input)

            // Create the Lexer
            lexer := parser.NewExprLexer(is)
            stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

            // Create the Parser
            p := parser.NewExprParser(stream)

            // Finally parse the expression
            antlr.ParseTreeWalkerDefault.Walk(&exprListener{}, p.Prog())


            //if input == "exit" {
			// 	fmt.Println("Bye!")
			// 	break
			// } else if input == "help" {
			// 	fmt.Println("Commands: help, exit, print")
			// } else if input == "print" {
			// 	fmt.Println("You typed: print")
			// } else {
			// 	fmt.Printf("You typed: %s\n", input)
			// }
		} else if err == liner.ErrPromptAborted {
			fmt.Println("Aborted")
			break
		} else {
			fmt.Println("Error reading line: ", err)
			break
		}
	}
}
