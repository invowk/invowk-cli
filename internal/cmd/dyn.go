package cmd

import (
	"context"
	"fmt"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"github.com/urfave/cli/v2"
	"os"
	"reflect"
)

func RunDynCommand() *cli.Command {
	return &cli.Command{
		Name:  "rundyn",
		Usage: "A dynamic command",
		Action: func(c *cli.Context) error {
			fmt.Println("rundyn called")

			content, err := os.ReadFile("D:\\dyn.go")
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				return err
			}
			fileContent := string(content)
			Eval(fileContent)

			return nil
		},
	}
}

func Eval(content string) {
	fmt.Printf("Evaluating: %s\n", content)

	i := interp.New(interp.Options{Unrestricted: true})

	i.Use(stdlib.Symbols)

	i.Use(interp.Exports{
		"sdk/sdk": {
			"SdkPrint": reflect.ValueOf(SdkPrint),
		},
	})

	i.ImportUsed()

	//buffer := bytes.NewBuffer(nil)
	//
	//extractor := extract.Extractor{Dest: "sdk"}
	//_, err := extractor.Extract("github.com/invowk/invowk-cli/internal/cmd", "github.com/invowk/invowk-cli/internal/cmd", buffer)
	//if err != nil {
	//	return
	//}

	ctx := GetContext()

	v, err := i.Eval(content)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
		panic(err)
	}

	fmt.Printf("Before evaling dynfunc\n")

	//p, err2 := i.Compile(content)
	//if err2 != nil {
	//	panic(err2)
	//}
	//p.PackageName()

	v, err = i.Eval("cmd.CtxCommand")

	if err != nil {
		fmt.Printf(err.Error() + "\n")
		panic(err)
	}

	fmt.Printf("Before calling dynfunc 1\n")

	dynfunc := v.Interface().(func(context.Context) string)

	fmt.Printf("Before calling dynfunc 2\n")

	dynfunc(ctx)

	fmt.Printf("After calling dynfunc\n")
}

func GetContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "SdkPrint", SdkPrint)
	return ctx
}

func SdkPrint() {
	fmt.Println("sdkPrint called")
}
