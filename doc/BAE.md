# Basket Encoding Protocol (Bae Protocol)

Bae protocol, simple to understand and catchs up to speed. 
It is very similar to the RESP protocol that is used by redis for communicating withing client and server.

## Commands

### Service Level Commands
These commands are used for server-level operations.

| Command       | Description                     | Example         | Response       |
|---------------|---------------------------------|-----------------|----------------|
| `PING`        | Check if the server is alive.  | `PING`          | `PONG`         |
| `CLOSE-CONN`  | Close the current connection.  | `CLOSE-CONN`    | `CLOSE-CONN`   |

---

### Database Level Commands
These commands manage baskets (databases).

| Command         | Description                              | Example                     | Response                     |
|-----------------|------------------------------------------|-----------------------------|-----------------------------|
| `CREATE`        | Create a new basket.                    | `CREATE myBasket`           | `OK`                        |
| `DROP`          | Delete an existing basket.              | `DROP myBasket`             | `OK`                        |
| `BASKET-INFO`   | Get information about a specific basket.| `BASKET-INFO myBasket`      | `<basket-info>`             |
| `LIST`          | List all available baskets.             | `LIST`                      | `<basket1> <basket2> ...`   |

---

### Basket Level Commands
These commands operate on keys within a specific basket.

| Command         | Description                              | Example                     | Response                     |
|-----------------|------------------------------------------|-----------------------------|-----------------------------|
| `GET`           | Retrieve the value of a key in a basket.| `GET myBasket myKey`        | `<value>`                   |
| `SET`           | Set a key-value pair in a basket.       | `SET myBasket myKey myValue`| `OK`                        |
| `DEL`           | Delete a key from a basket.             | `DEL myBasket myKey`        | `OK`                        |
| `CLEAR`         | Clear all keys in a basket.             | `CLEAR myBasket`            | `OK`                        |
| `EXISTS`        | Check if a key exists in a basket.      | `EXISTS myBasket myKey`     | `+1` (exists) or `-1` (not exists) |
| `KEYS`          | Retrieve all keys matching a pattern.   | `KEYS myBasket prefix_*`    | `<key1> <key2> ...`         |

---