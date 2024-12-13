# Watchd

A simple file system watcher that executes commands or runs executables when changes are detected in the working directory.

## Supported Platforms

- macOS (Darwin/amd64)

## Installation

### Build from source

```bash
GOOS=darwin GOARCH=amd64 go build -o ./bin/watchd main.go
```

### Install to system

```bash
sudo cp ./bin/watchd /usr/local/bin/
```

## Usage

Watchd supports two modes of operation:

1. Execute a shell command
2. Run an executable file

### Command mode

```bash
watchd --command "go test ./..."
```

This will watch the current directory and its subdirectories for changes and execute the specified command when changes are detected.

### Executable mode

```bash
watchd --path "/path/to/executable"
```

This will watch the current directory and its subdirectories for changes and run the specified executable when changes are detected.

## Flags

- `--command`: Specifies a shell command to execute
- `--path`: Specifies the path to an executable file

Note: You must specify exactly one flag. Using both or neither flags will result in an error.

## License

[MIT](LICENSE)
