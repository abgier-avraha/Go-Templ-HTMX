package main

import (
	"context"
	"os"
)

type Something struct {
	Name string
}

func main() {
	component := hello(&Something{
		Name: "Shit",
	})
	component.Render(context.Background(), os.Stdout)
}
