package JsonParser

type JsonTokenType int32
const (
	JSON_TOKEN_STRING               JsonTokenType = 0
	JSON_TOKEN_NUMBER               JsonTokenType = 1
	JSON_TOKEN_QUOT                 JsonTokenType = 2 // "
	JSON_TOKEN_LEFT_SQUARE_BRACKET  JsonTokenType = 3 // [
	JSON_TOKEN_RIGHT_SQUARE_BRACKET JsonTokenType = 4 // ]
	JSON_TOKEN_LEFT_CURLY_BRACKET   JsonTokenType = 5 // {
	JSON_TOKEN_RIGHT_CURLY_BRACKET  JsonTokenType = 6 // }
	JSON_TOKEN_COLON                JsonTokenType = 7 // :
)

type JsonToken struct {
	tokenType JsonTokenType
	tokenValue interface{}
}


type TokensMap struct {
	tokensMap map[JsonKey]JsonValue
}

func InitTokensMap() *TokensMap {
	var jsonMap TokensMap
	jsonMap.tokensMap = make(map[JsonKey]JsonValue)
	return &jsonMap
}






