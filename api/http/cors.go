package http

var Cors struct {
	AllowOrigins string
	AllowMethods string
	AllowHeaders string
}

func init() {
	Cors.AllowHeaders = "*"
	Cors.AllowMethods = "*"
	Cors.AllowOrigins = "*"
}
