package main

import (
	"Crawler/engine"
	"Crawler/source/zhenAi"
)

func main() {
	e := engine.SimpleEngine{}

	e.Run(zhenAi.IndexRequest())
}
