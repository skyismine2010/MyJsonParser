package JsonParser

type JsonTokenType int32
const (
	JSON_TOKEN_STRING JsonTokenType = 0
	JSON_TOKEN_NUMBER JsonTokenType = 1
	JSON_TOKEN_QUOT   JsonTokenType = 2 // "
	JSON_LEFT_SQUARE_BRACKET JsonTokenType = 3 //[
	JSON_RIGHT_SQUARE_BRACKET JsonTokenType = 4 //]
	JSON_LEFT_CURLY_BRACKET JsonTokenType = 5 //{
	JSON_RIGHT_CURLY_BRACKET JsonTokenType = 6 //}
)

type JsonToken struct {
	tokenType JsonTokenType
	tokenValue interface{}
}








