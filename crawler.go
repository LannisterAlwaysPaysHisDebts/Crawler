package main

import (
	"myGit/Crawler/engine"
	"myGit/Crawler/source/zhenAi"
)

func main()  {
	e := engine.SimpleEngine{}

	e.Run(zhenAi.IndexRequest())
}
