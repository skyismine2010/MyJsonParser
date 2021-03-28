package JsonParser

import (
	"errors"
	"strconv"
)

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
	keyName *string
}

type JsonValue struct {
	valueType JsonValueType
	value interface{}
}


//
func ParseJsonKey(ctx *JsonParserCtx)(*JsonKey, error) {
	keyStr, err :=ParseJsonString(ctx)
	if err == nil {
		return &JsonKey{keyStr}, nil
	}
	return nil, err
}

func ParsePostJsonString(ctx *JsonParserCtx)(*string, error) {
	var key string
	t, e := ctx.GetOneToken()
	if e != nil || t.tokenType != JSON_TOKEN_STRING {
		return  nil, errors.New("Bad Format, Line=%d, column=%d")  //todo 错误处理
	}

	key = t.tokenValue.(string)

	t, e = ctx.GetOneToken()
	if e != nil || t.tokenType != JSON_TOKEN_QUOT {
		return  nil, errors.New("Bad Format, Line=%d, column=%d")  //todo 错误处理
	}
	return &key, nil
}

func ParseJsonString(ctx *JsonParserCtx)(*string, error) {
	t, e := ctx.GetOneToken()
	if e!= nil || t.tokenType != JSON_TOKEN_QUOT {
		return  nil, errors.New("Bad Format, Line=%d, column=%d")  //todo 错误处理
	}
	return ParsePostJsonString(ctx)
}

func ParseJsonNumber(ctx *JsonParserCtx)(int, error) {
	t, e := ctx.GetOneToken()
	if e != nil || t.tokenType != JSON_TOKEN_NUMBER {
		return 0, errors.New("Not a number")
	}
	return strconv.Atoi(t.tokenValue.(string))
}


// Obj --> { "string" :
func ParseJsonObject(ctx *JsonParserCtx) (*TokensMap, error) {
	var key *JsonKey
	var value JsonValue

	t,e := ctx.GetOneToken()
	if e !=nil || t.tokenType != JSON_LEFT_CURLY_BRACKET {
		return  nil, errors.New("Bad Format, Line=%d, column=%d")  //todo 错误处理
	}

	key, err :=ParseJsonKey(ctx)
	if err != nil {
		return nil, err
	}

	t, e= ctx.GetOneToken()
	if e!=nil {
		if t.tokenType == JSON_TOKEN_QUOT {
			if keyStr,e := ParsePostJsonString(ctx); e== nil {
				value.valueType = JSON_TYPE_STRING
				value.value  = keyStr
			} else if t.tokenType == JSON_TOKEN_NUMBER {
				if keyNum,e := ParseJsonNumber(ctx); e== nil {
					value.valueType = JSON_TYPE_NUMBER
					value.value = keyNum
				}
			} else if t.tokenType == JSON_LEFT_CURLY_BRACKET; e == nil {

			}

		}
	}

}


func ParseJson(jsonStr *string) *TokensMap {
	jsonMap := InitTokensMap()
	ctx := InitJsonParseCtx(jsonStr)
	for ;!ctx.IsCtxEnd(); {

	}


	return jsonMap


}



