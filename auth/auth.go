package auth

// Controller is an interface for authentication controllers.
type Controller interface {

	// Auth authenticates a user on CONNECT and returns true if a user is
	// allowed to join the server.
	Auth(user string, password string) bool

	// ACL returns true if a user has read or write access to a given topic.
	ACL(user string, topic string, write bool) bool
}