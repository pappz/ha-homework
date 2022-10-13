package middleware

type Json interface {
	Validate() error
}
