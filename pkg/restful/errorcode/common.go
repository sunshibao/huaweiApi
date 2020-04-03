package errorcode

const (
	CommonError            = 500000
	ParameterError         = 400000
	JSONParseError         = 406001
	AlreadyExistError      = 406002
	NotExistError          = 406003
	NullDataError          = 406004
	EmailAlreadyExistError = 406005
)

var statusText = map[int]string{
	CommonError:            "Common Error",
	ParameterError:         "Parameter Error",
	JSONParseError:         "JSON Parse Error",
	AlreadyExistError:      "Already Exist Error",
	NotExistError:          "Not Exist Error",
	NullDataError:          "Null Data Error",
	EmailAlreadyExistError: "Email Already Exist Error",
}

func StatusText(code int) string {
	return statusText[code]
}
