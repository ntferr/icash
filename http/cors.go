package http

var Cors struct {
	AllowOrigins string
	AllowMethods string
	AllowHeaders string
}

func init() {
	Cors.AllowHeaders = "Origin, Content-Type, Accept"
	Cors.AllowMethods = "GET,POST,HEAD,PUT,DELETE,PATCH"
	Cors.AllowOrigins = ""
}
