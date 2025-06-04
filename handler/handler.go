package handler

import (
	"github.com/KiritoCyanPine/smolBasket/encoder"
	"github.com/KiritoCyanPine/smolBasket/storage"
)

func Handler(enc encoder.Encoder, sm storage.Manager, commad []string) ([]byte, error) {
	var err error
	var reply []byte

	route := commad[0]
	switch GetCommandLevel(route) {
	case LEVEL_SERVICE:
		// Handle service level commands
		reply, err = handleServiceCommand(enc, commad)
	case LEVEL_DATABASE:
		// Handle database level commands
		reply, err = handleDatabaseCommand(enc, sm, commad)
	case LEVEL_BASKET:
		// Handle basket level commands
		reply, err = handleBasketCommand(enc, sm, commad)
	case LEVEL_NONE:
		err = ErrInvalidCommand
	}

	return reply, err
}

func handleServiceCommand(enc encoder.Encoder, command []string) ([]byte, error) {
	// Implement service level command handling
	commandName := command[0]

	switch commandName {
	case "PING":
		return enc.EncodeBAECommand("PONG"), nil
	case "CLOSE-CONN":
		return enc.EncodeBAECommand("CLOSE-CONN"), ErrConnectionClosed
	}
	return nil, nil
}

func handleDatabaseCommand(enc encoder.Encoder, sm storage.Manager, command []string) ([]byte, error) {
	// Implement database level command handling
	commandName := command[0]

	switch commandName {
	case "CREATE":
		if len(command) != 2 {
			return nil, ErrInvalidCommand
		}

		basketName := command[1]
		if err := sm.Create(basketName); err != nil {
			return nil, err
		}

		return enc.EncodeBAECommand("OK", "Basket created successfully"), nil
	case "DROP":
		if len(command) != 2 {
			return nil, ErrInvalidCommand
		}

		basketName := command[1]
		if err := sm.Drop(basketName); err != nil {
			return nil, err
		}

		return enc.EncodeBAECommand("OK"), nil
	case "BASKET-INFO":
		if len(command) != 2 {
			return nil, ErrInvalidCommand
		}

		basketName := command[1]
		info, err := sm.Info(basketName)
		if err != nil {
			return nil, err
		}

		return enc.EncodeBAECommand(info), nil
	case "LIST":
		if len(command) != 2 {
			return nil, ErrInvalidCommand
		}

		baskets, err := sm.List()
		if err != nil {
			return nil, err
		}

		if len(baskets) == 0 {
			return nil, ErrNoBasketFound
		}

		return enc.EncodeBAECommand(baskets...), nil
	}

	return nil, nil
}

func handleBasketCommand(enc encoder.Encoder, sm storage.Manager, command []string) ([]byte, error) {
	// Implement basket level command handling
	if len(command) < 2 {
		return nil, ErrInvalidCommand
	}

	basketName := command[1]
	db, err := sm.GetBasket(basketName)
	if err != nil {
		return nil, err
	}

	switch command[0] {
	case "GET":
		if len(command) != 3 {
			return nil, ErrInvalidCommand
		}
		key := command[2]
		value, ok := db.Get(key)
		if !ok {
			return nil, ErrKeyNotFound
		}

		return enc.EncodeBAECommand(value), nil
	case "SET":
		if len(command) < 4 {
			return nil, ErrInvalidCommand
		}
		key := command[2]
		value := command[3]
		db.Set(key, value)

		return enc.EncodeBAECommand("OK"), nil
	case "DEL":
		if len(command) != 3 {
			return nil, ErrInvalidCommand
		}

		key := command[2]
		db.Delete(key)

		return enc.EncodeBAECommand("OK"), nil
	case "CLEAR":

		if len(command) != 3 {
			return nil, ErrInvalidCommand
		}

		db.Clear()
		return enc.EncodeBAECommand("OK"), nil
	case "EXISTS":
		if len(command) != 3 {
			return nil, ErrInvalidCommand
		}

		key := command[2]
		exists := db.Exists(key)

		if exists {
			return enc.EncodeBAECommand("+1"), nil
		}
		return enc.EncodeBAECommand("-1"), nil
	case "KEYS":
		if len(command) != 3 {
			return nil, ErrInvalidCommand
		}
		pattern := command[2]
		keys, err := db.Keys(pattern)
		if err != nil {
			return nil, err
		}
		if len(keys) == 0 {
			return enc.EncodeBAECommand("nil"), nil
		}
		return enc.EncodeBAECommand(keys...), nil
	}
	// If command does not match any known basket commands
	// return an error indicating the command is invalid
	return nil, ErrInvalidCommand
}
