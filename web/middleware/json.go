package middleware

// Json general interface what represent the request parameters
type Json interface {
	Validate() error
}
