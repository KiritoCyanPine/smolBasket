package handler

const (
	helpText = `smolBasket CLI Help
===================
Available Commands:

1. Service Level Commands:
   - PING
     Description: Check if the server is alive.
     Example: PING
     Response: PONG

   - CLOSE-CONN
     Description: Close the current connection.
     Example: CLOSE-CONN
     Response: CLOSE-CONN

2. Database Level Commands:
   - CREATE <basket-name>
     Description: Create a new basket.
     Example: CREATE myBasket
     Response: OK

   - DROP <basket-name>
     Description: Delete an existing basket.
     Example: DROP myBasket
     Response: OK

   - BASKET-INFO <basket-name>
     Description: Get information about a specific basket.
     Example: BASKET-INFO myBasket
     Response: <basket-info>

   - LIST
     Description: List all available baskets.
     Example: LIST
     Response: <basket1> <basket2> ...

3. Basket Level Commands:
   - GET <basket-name> <key>
     Description: Retrieve the value of a key in a basket.
     Example: GET myBasket myKey
     Response: <value>

   - SET <basket-name> <key> <value>
     Description: Set a key-value pair in a basket.
     Example: SET myBasket myKey myValue
     Response: OK

   - DEL <basket-name> <key>
     Description: Delete a key from a basket.
     Example: DEL myBasket myKey
     Response: OK

   - CLEAR <basket-name>
     Description: Clear all keys in a basket.
     Example: CLEAR myBasket
     Response: OK

   - EXISTS <basket-name> <key>
     Description: Check if a key exists in a basket.
     Example: EXISTS myBasket myKey
     Response: +1 (exists) or -1 (does not exist)

   - KEYS <basket-name> <pattern>
     Description: Retrieve all keys in a basket matching a pattern.
     Example: KEYS myBasket prefix_*
     Response: <key1> <key2> ...

---

Usage:
- Type a command followed by its arguments.
- Use "exit" to quit the CLI.
- Use "help" to display this help message.`
)

func GetHelpText() string {
	return helpText
}
