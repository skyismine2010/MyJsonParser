package JsonParser

import "errors"

type JsonValueType int32
const (
	JSON_TYPE_NIL JsonValueType = 0
	JSON_TYPE_FALSE JsonValueType = 1
	JSON_TYPE_TRUE  JsonValueType = 2
	JSON_TYPE_STRING JsonValueType = 3
	JSON_TYPE_NUMBER JsonValueType = 4
	JSON_TYPE_ARRAY JsonValueType = 5

)

type JsonParserCtx struct {
	content *string
	pos  int
	errMsg string
}

func InitJsonParseCtx(jsonStr *string) *JsonParserCtx {
	var ctx JsonParserCtx
	ctx.content = jsonStr
	ctx.pos = 0
	return &ctx
}


func (self *JsonParserCtx) IsCtxEnd() bool {
	return self.pos >= len(*self.content)
}

func (self *JsonParserCtx) isCurrentCharBlank() bool {
	c := (*self.content)[self.pos]
	if c == ' ' || c == '\r' || c == '\n' || c == '\t' {
		return true
	}
	return false
}

func (self *JsonParserCtx) IgnoreBlankSpace() error {
	for ;!self.IsCtxEnd(); {
		if self.isCurrentCharBlank() {
			self.pos++
			continue
		}
		return nil
	}
	return errors.New("EOF")
}

func (self *JsonParserCtx) GetOneToken() (*JsonToken, error){

	return nil, nil
}




type JsonKey struct {
	keyName string
}

type JsonValue struct {
	valueType JsonValueType
	value interface{}
}

type TokensMap struct {
	tokensMap map[JsonKey]JsonValue
}

func InitTokensMap() *TokensMap {
	var jsonMap TokensMap
	jsonMap.tokensMap = make(map[JsonKey]JsonValue)
	return &jsonMap
}

// Obj --> { "string" :
func ParseJsonObject(ctx *JsonParserCtx) *TokensMap {
	t,e := ctx.GetOneToken()
	if e !=nil || t.tokenType != JSON_LEFT_CURLY_BRACKET {

	}
}


func ParseJson(jsonStr *string) *TokensMap {
	jsonMap := InitTokensMap()
	ctx := InitJsonParseCtx(jsonStr)
	for ;!ctx.IsCtxEnd(); {

	}


	return jsonMap


}



