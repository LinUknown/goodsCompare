package engine

import "imooc.com/model"

type Request struct {
	URL string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct{
	Requests []Request
	Items []model.Goods
}

// NilParser return nil
func NilParser([]byte) ParseResult {
	return ParseResult{}
}