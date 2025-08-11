package main

import "github.com/dr8co/palettes/palettes"

func main() {
	registry := palettes.InitAllSchemes()
	registry.ShowAll()
}
