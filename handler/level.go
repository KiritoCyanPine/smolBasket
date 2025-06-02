package handler

const (
	// LEVEL_NONE indicates no transaction level
	LEVEL_NONE = "none"
	// LEVEL_SERVICE represents commands in the service level.
	LEVEL_SERVICE = "service"
	// LEVEL_DATABASE represents commands in the database level.
	LEVEL_DATABASE = "database"
	// LEVEL_BASKET represents commands in the cache level.
	LEVEL_BASKET = "basket"
)

func GetCommandLevel(command string) string {
	switch command {
	case "PING", "CLOSE-CONN":
		return LEVEL_SERVICE
	case "CREATE", "DROP", "BASKET-INFO", "LIST":
		return LEVEL_DATABASE
	case "GET", "SET", "DEL", "CLEAR", "EXISTS", "KEYS":
		return LEVEL_BASKET
	default:
		return LEVEL_NONE
	}
}
