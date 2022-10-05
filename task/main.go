package main

import (
	"os"

	"pprof-task/fast"
)

const filePath string = "./data/users.txt"

func main() {
	fast.Search(filePath, os.Stdout)
}
