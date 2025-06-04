# smolBasket

⚡️ “Fast, simple, lightweight KV caching service written in Go — like Redis, but simpler and pluggable. Zero setup, and perfect for local dev or edge compute."

![main_image](doc/main_image.png)

---

smolBasket reads lock free and writes isolated. Optimised for fast read access. Great for small tasks


## Features
- **Lightweight**: Minimal resource usage, perfect for edge devices.
- **Multi Pool Connect** : Automatic connection pooling.
- **Pluggable**: Easily extend functionality with custom handlers.
- **Zero Setup**: Start the server with a single command.
- **Multi-Core Support**: Optimized for high-performance workloads.
- **Beginner Friendly** : Simple and intuitive CLI commands. 


## Learn smolBasket?

// Add link to documentation

// Checkout our developer guide

// Join our Discord server - Add server link

[Configuration Guide](doc/configuration-guide.md)

[BAE Protocol](doc/BAE.md)


## How to use ?

### Running the binaries in-house
Clone the repository.
```sh
git clone https://github.com/KiritoCyanPine/smolBasket.git
cd smolBasket
```
#### Running the server :
```sh
go build -o smolBasket ./cmd/server/main.go
./smolBasket
```

The output log is usually. 
```
Starting KV Server on tcp://:9001...
KV Server started on :9001 [multi-core: false]
```

And this means that your caching server is working just fine and is available to connect on port 9001. The configuration guide will provide information regarding changeing the port and other settings.

#### Running cli tool:
smolBasket comes with its own cli tool to communicate with the server directly with commands. 

Use the built-in CLI client to interact with the server:

```sh
go build -o smol-cli ./cmd/client
./smol-cli
```

This will open up an interavtive cli to work with smolBasket and start pushing and pulling items. It looks something like this. 

```
smolBasket CLI (with history). Type 'exit' to quit.
>  
```

As for your first command you can try to ping the server and wait for the response. 
```
smolBasket CLI (with history). Type 'exit' to quit.
> PING
Decoded response: [PONG] 1
>  
```

Please look at 