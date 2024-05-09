package simplesessions

// Store represents store interface. This interface can be
// implemented to create various backend stores for session.
type Store interface {
	// Create creates new session in store and returns the cookie value.
	Create() (cookieValue string, err error)

	// Get gets a value for given key from session.
	Get(cookieValue, key string) (value interface{}, err error)

	// GetMulti gets a maps of multiple values for given keys.
	GetMulti(cookieValue string, keys ...string) (values map[string]interface{}, err error)

	// GetAll gets all key and value from session,
	GetAll(cookieValue string) (values map[string]interface{}, err error)

	// Set sets an value for a field in session.
	// Its up to store to either store it in session right after set or after commit.
	Set(cookieValue, key string, value interface{}) error

	// Commit commits all the previously set values to store.
	Commit(cookieValue string) error

	// Delete a field from session.
	Delete(cookieValue string, key string) error

	// Clear clears the session key from backend if exists.
	Clear(cookieValue string) error

	// Helper method for typecasting/asserting.
	Int(interface{}, error) (int, error)
	Int64(interface{}, error) (int64, error)
	UInt64(interface{}, error) (uint64, error)
	Float64(interface{}, error) (float64, error)
	String(interface{}, error) (string, error)
	Bytes(interface{}, error) ([]byte, error)
	Bool(interface{}, error) (bool, error)
}
