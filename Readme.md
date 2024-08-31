# Go-Redis

Go-Redis is an in-memory database implementation inspired by Redis, written in Go. It utilizes the RESP (REdis Serialization Protocol) structure to enable clients to interact with the database by getting and setting data.

## Features

- In-memory database: Go-Redis stores data in memory, allowing for fast and efficient data retrieval.
- RESP structure: The database uses the RESP structure, making it compatible with existing Redis clients.
- Get and set operations: Clients can easily retrieve and update data using the familiar get and set operations.

## Installation

To install Go-Redis, follow these steps:

1. Clone the repository: `git clone https://github.com/rishikesh-suvarna/go-redis.git`
2. Navigate to the project directory: `cd go-redis`
3. Build the project: `go build`
4. Run the executable: `./go-redis`

## Usage

To interact with Go-Redis, you can use any Redis client that supports the RESP structure. Here's an example using the `redis-cli` command-line client:

1. Start the Go-Redis server: `./go-redis`
2. Open a new terminal window and run `redis-cli`.
3. Connect to the Go-Redis server: `redis-cli -h localhost -p 6379`
4. Now you can use Redis commands like `GET`, `SET`, and more to interact with the database.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License

Go-Redis is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
