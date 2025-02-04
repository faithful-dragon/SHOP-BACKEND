package utils

type codes struct {
	StatusCode  int
	Message     string
	Description string
}

const (
	StatusOK                  = 200
	StatusBadRequest          = 400
	StatusNotFound            = 404
	StatusRequestTimeout      = 408
	StatusInternalServerError = 500
)

var codeMap map[string]codes

func SetCodeMap() {
	codeMap = make(map[string]codes)
	codeMap["4001"] = codes{StatusBadRequest, "4001", "Bad Request"}
	codeMap["4002"] = codes{StatusNotFound, "4002", "Not Found"}
	codeMap["4003"] = codes{StatusRequestTimeout, "4003", "Request Timeout"}
	codeMap["4004"] = codes{StatusInternalServerError, "4004", "Internal Server Error"}
}
