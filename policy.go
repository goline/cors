package cors

type Policy struct {
	AllowOrigin      string
	AllowHeaders     string
	AllowMethods     string
	AllowCredentials string
	Origin           string
	RequestMethod    string
	RequestHeaders   string
	ExposeHeaders    string
	MaxAge           int
}
