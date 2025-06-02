# Configuration guide for the smolBasket server

To change the settings of smolBasket and the server configurations, change the .env file. 

This document outlines the configuration options available for the `smolBasket` server. The configuration can also be set using environment variables or defaults provided in the code.

---

## Configuration Options

| Environment Variable | Default Value | Description                                                                 |
|-----------------------|---------------|-----------------------------------------------------------------------------|
| `TCP_PORT`            | `9000`        | The port on which the server listens for incoming connections.              |
| `REUSE_PORT`          | `false`       | Enables or disables port reuse.                                            |
| `MULTICORE_MODE`      | `false`       | Enables or disables multi-core mode for the server.                        |
| `LOAD_BALANCING`      | `0`           | Specifies the load balancing strategy. See the **Load Balancing Options**. |

---

## Load Balancing Options

The `LOAD_BALANCING` environment variable determines the load balancing strategy used by the server. The following values are supported:

| Value | Strategy             | Description                                                                 |
|-------|-----------------------|-----------------------------------------------------------------------------|
| `0`   | `RoundRobin`         | Distributes connections evenly across all available workers.               |
| `1`   | `LeastConnections`   | Directs new connections to the worker with the fewest active connections.  |
| `2`   | `SourceAddrHash`     | Routes connections based on the hash of the client's source address.       |

---

## Default Configuration

If no environment variables are provided, the server uses the following default configuration:

```json
{
  "port": "9000",
  "reuse_port": false,
  "multicore": false,
  "load_balancing": "RoundRobin"
}
```

---

## Environment Variable Details

### `TCP_PORT`
- **Description**: Specifies the port on which the server listens for incoming connections.
- **Default**: `9000`
- **Example**:
  ```bash
  export TCP_PORT=8080
  ```

### `REUSE_PORT`
- **Description**: Enables or disables port reuse.
- **Default**: `false`
- **Accepted Values**: `true` or `false`
- **Example**:
  ```bash
  export REUSE_PORT=true
  ```

### `MULTICORE_MODE`
- **Description**: Enables or disables multi-core mode for the server.
- **Default**: `false`
- **Accepted Values**: `true` or `false`
- **Example**:
  ```bash
  export MULTICORE_MODE=true
  ```

### `LOAD_BALANCING`
- **Description**: Specifies the load balancing strategy.
- **Default**: `0` (RoundRobin)
- **Accepted Values**:
  - `0`: RoundRobin
  - `1`: LeastConnections
  - `2`: SourceAddrHash
- **Example**:
  ```bash
  export LOAD_BALANCING=1
  ```

---

## How Configuration is Loaded

1. **Environment Variables**:
   - The server first attempts to load configuration values from environment variables.
   - If a `.env` file is present, it is loaded using the `godotenv` package.

2. **Defaults**:
   - If an environment variable is not set, the server falls back to the default values defined in the code.

---

## Example `.env` File

You can create a `.env` file in the root of your project to define configuration values:

```env
TCP_PORT=9000
REUSE_PORT=true
MULTICORE_MODE=true
LOAD_BALANCING=1
```

---

## Notes

- Invalid values for `REUSE_PORT`, `MULTICORE_MODE`, or `LOAD_BALANCING` will cause the server to fall back to their respective default values.
- Ensure that the `.env` file is properly formatted and placed in the root directory of the project.

---

## Usage in Code

The configuration is loaded using the `GetConfiguraation` function:

```go
import "github.com/KiritoCyanPine/smolBasket/configuration"

func main() {
    config := configuration.GetConfiguraation()
    fmt.Printf("Server running on port: %s\n", config.Port)
}
```

This function automatically loads the configuration from environment variables or defaults.
