## How to Run

There are two ways to run the application. The first is to setup all dependencies manually.
The second is to use docker images and docker compose.

### Manual

- Create `.env` file

  You can copy the `env.sample` and change its values.

    ```
    $ cp env.sample .env
    ```
- Generate OAuth Key

    ```
    $ make key.generate
    ```

- Run the migration

    ```
    $ make all-db-migrate url="postgres://YOUR_POSTGRES_USER:YOUR_POSTGRES_PASSWORD@YOUR_POSTGRES_HOST:YOUR_POSTGRES_PORT/YOUR_DATABASE_NAME"
    ```

- Run the initial user seeder

    ```
    $ make seed
    ```

- Run the application

    ```
    $ go run main.go
    ```

### Docker

- Install [Docker Compose](https://docs.docker.com/compose/).

- Download the dependencies

    ```
    $ make tidy
    ```

- Compile the backend binary

    ```
    $ make compile-server
    ```

- Build backend-server image

    ```
    $ make docker-build-server
    ```

- Run docker compose

    ```
    $ docker-compose up
    ```