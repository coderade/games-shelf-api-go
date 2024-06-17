
# games-shelf-api (Go)

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


## Project Structure

Here is an overview of the project's directory structure:

      games-shelf-api-go/
      ├── cmd/
      │   ├── main_test.go
      │   └── main.go
      ├── docs/
      ├── internal/
      │   ├── api/
      │   │   ├── handlers/
      │   │   │   ├── auth_handler.go
      │   │   │   ├── auth_handler_test.go
      │   │   │   ├── games_handler.go
      │   │   │   ├── games_handler_test.go
      │   │   │   ├── genres_handler.go
      │   │   │   ├── genres_handler_test.go
      │   │   │   ├── graphql_handler.go
      │   │   │   ├── graphql_handler_test.go
      │   │   │   ├── platforms_handler.go
      │   │   │   ├── platforms_handler_test.go
      │   │   │   ├── status_handler.go
      │   │   │   └── status_handler_test.go
      │   │   ├── middleware.go
      │   │   ├── middleware_test.go
      │   │   ├── routes.go
      │   │   ├── server.go
      │   │   └── server_test.go
      │   ├── config/
      │   │   └── config.go
      │   ├── db/
      │   │   ├── database_helper.go
      │   │   ├── database_interface.go
      │   │   └── database_mock.go
      │   ├── logger/
      │   │   └── logger.go
      │   ├── mocks/
      │   │   ├── rawgservice_mock.go
      │   │   └── shelf_mock.go
      │   ├── models/
      │   │   ├── game.go
      │   │   ├── genre.go
      │   │   ├── platform.go
      │   │   └── user.go
      │   ├── repository/
      │   │   ├── shelf_interface.go
      │   │   └── shelf.go
      │   ├── service/
      │   │   ├── graphql/
      │   │   │   └── schema.go
      │   │   ├── rawg_service.go
      │   │   └── rawg_service_interface.go
      │   └── utils/
      │       └── json.go
      ├── .gitignore
      ├── docker-compose.yml
      ├── Dockerfile
      ├── go.mod
      ├── go.sum
      └── README.md


## Running the Application 


### Running the application locally

#### Environment Variables

The following environment variables are required to run this application:

- `PORT`: Server port to listen on (default: **4000**)
- `ENV`: Application environment (development|production) (default: **development**)
- `DB_DATA_SOURCE`: Database data source (default: **postgres://admin@localhost/games_shelf?sslmode=disable**)
- `APP_SECRET`: Application secret, used to create JWT token (default: **games-shelf-api-secret**)
- `RAWG_API_KEY`: API key used to connect to the RAWG API. To get an API Key, check [here](https://rawg.io/apidocs).

#### Database Setup

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
psql -h localhost -d games_shelf -U admin -p 5432 -a -q -f docs/db/games_shelf.sql
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

### Run the Application using Docker

 **Build and run using Docker Compose**:
   ```sh
   docker-compose up --build
   ```

 **Access the application**:
   The API will be available at `http://localhost:4000`.



### Running Tests

To run the tests for the application, use the following command:

```sh
go test ./...
```
