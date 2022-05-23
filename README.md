
# Games Shelf API (Go)

Games Shelf REST API developed in Go.

This application allows you to:

- Login using JWT tokens (user and password management is handled on the backend)
- Perform CRUD operations on games, platforms, and genres
- Connect to a PostgreSQL database
- Use GraphQL to search for games
- Integrate with the RAWG API, the largest video game database, to retrieve game images and ratings

## Technologies Used

- Go 1.18.1
- PostgreSQL
- JWT
- Crypto
- GraphQL

Client application developed in React is available [here](https://github.com/coderade/games-shelf-client-react).

## Running the Application 

### Environment Variables

The following environment variables are required to run this application:

- `PORT`: Server port to listen on (default: **4000**)
- `ENV`: Application environment (development|production) (default: **development**)
- `DB_DATA_SOURCE`: Database data source (default: **postgres://admin@localhost/games_shelf?sslmode=disable**)
- `APP_SECRET`: Application secret, used to create JWT token (default: **games-shelf-api-secret**)
- `RAWG_API_KEY`: API key used to connect to the RAWG API. To get an API Key, check [here](https://rawg.io/apidocs).

### Database Setup

To create the database, roles, and grant the needed privileges, run the following SQL commands:

```sql
CREATE DATABASE games_shelf;
CREATE ROLE admin;
ALTER ROLE admin login;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO admin;
```

To create all the tables, sequences, constraints, and content needed for the application, run the following command:

```sh
psql -h localhost -d games_shelf -U admin -p 5432 -a -q -f db/games_shelf.sql
```

### Running the Application

After setting the required environment variables and creating the database, you can run the application in the following ways:

#### Build and Execute via `./main`

```sh
go build cmd/*.go
./main
```

#### Run with `go run`

```sh
go run cmd/*.go
```

### Steps to Run the Application

1. **Build and run using Docker Compose**:
   ```sh
   docker-compose up --build
   ```

2. **Access the application**:
   The API will be available at `http://localhost:4000`.



### Running Tests

To run the tests for the application, use the following command:

```sh
go test ./...
```
