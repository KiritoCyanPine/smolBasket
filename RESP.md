# RESP 

RESP is the protocol Redis uses to communicate between client and server. It's text-based, simple, and human-readable (but also compact).
Every command and its arguments are serialized in a specific format, which includes:
```
*<count> for arrays

$<len> for bulk strings

+ for simple strings

- for errors

: for integers
```


---

## Resp definitions

For a request going from client to server:


### Ping the Server for keep alive
`#\nPING`

`#\nPONG`

### Creating a Basket
`=\nCREATE\n[Basket-Name]`

`+\nOK`

### Clear the Basket
`=\nCLEAR\n[Basket-Name]`

`+\nOK`

`-\nBASKET NOT FOUND`

### Deleting the Basket
`=\nDROP\n[Basket-Name]`

`+\nOK`

`-\nBASKET NOT FOUND`

### Putting Value into a Basket
`*\nSET\n[Basket-Name]\n[Key]\n[Value]`

`+\nOK`

### Getting Value from a Basket
`*\nGET\n[Basket-Name]\n[Key]`

`:\n[Value]`

`-\nKEY NOT FOUND`

### Checking Key's Existance
`*\nEXISTS\n[Basket-Name]\n[Key]`

`+\n1`

`+\n-1`

### Deleting Value from a Basket
`*\nDEL\n[Basket-Name]\n[Key]`

`+\nOK`

`-\nKEY NOT FOUND`

### Get list of keys matching the pattern
`*\nKEYS\n[Basket-Name]\n[prefix_*]`

`$\nKEY_1\nKEY_2\nKEY_3`

`-\nPATTERN NOT FOUND`

---

## Symbol Table

| Symbol    | Meaning |
| -------- | ------- |
| `*`  | For any String value to be communicated  |
| `+` | For all response status     |
| `$`    | For List of Values    |
| `-`    | For all Errors    |
