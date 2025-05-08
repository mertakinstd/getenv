# GetEnv

[![Go Report Card](https://goreportcard.com/badge/github.com/mertakinstd/getenv)](https://goreportcard.com/report/github.com/mertakinstd/getenv)

A simple environment variables loader for Go that loads environment variables from `.env`, `.env.development`, and `.env.production` files.

## Features

- Default `.env` file support
- Development environment support with `.env.development`
- Production environment support with `.env.production`
- Automatic variable reference resolution (`${VAR}`)
- Preserves existing environment variables
- Comment line support (`#`)

## Installation

```bash
go get github.com/mertakinstd/getenv
```

## Usage

```go
package main

import "github.com/mertakinstd/getenv"

func main() {
    // Load default .env file
    getenv.Load().Default()

    // or

    // For development environment
    getenv.Load().Development()

    // For production environment
    getenv.Load().Production()

    // For getting variables

    port := os.Getenv("PORT")
}
```

## Example .env File

```env
PORT=8080
IP=10.0.0.1
HOST=${IP}:${PORT}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
