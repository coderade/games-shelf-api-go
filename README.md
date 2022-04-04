# games-shelf-api-go

Games Shelf API developed in GO.

## Database
#### TODO: Update doc

### commands

    CREATE DATABASE games_shelf
    CREATE ROLE admin;
    ALTER ROLE admin login;

    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public to admin;
    GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public to admin;
    GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public to admin;
      
