package engine

type ParserFunc func(contents []byte, url string) ParserResult

type Parser interface {
	Parse(contents []byte, url string) ParserResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url        string
	ParserFunc ParserFunc
	Parser     Parser
}

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParserResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
