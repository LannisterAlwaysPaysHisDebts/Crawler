package engine

type ParserFunc func(contents []byte, url string) ParserResult

// 网页解析器
type Parser interface {
	Parse(contents []byte, url string) ParserResult // 解析函数
	Serialize() (name string, args interface{})     // 数据 rpc序列化
}

// 单个采集item
type Request struct {
	Url    string // 采集连接
	Parser Parser // 解析器
}

// 网页解析器 解析结果
type ParserResult struct {
	Requests []Request // 写入当前页面解析出来的其他需要采集的地址
	Items    []Item    // 解析出来的数据
}

// 页面采集最后的解析数据
type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

// 这是一个Parser， 实现了Parse与Serialize方法
type FuncParser struct {
	parser ParserFunc // 解析器
	name   string     // 解析器名
}

func (f *FuncParser) Parse(contents []byte, url string) ParserResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// 获取一个Parser
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
