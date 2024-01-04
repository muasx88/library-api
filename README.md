# Library API

Project very basic for Library using Domain Driven Design (DDD) Architecture

This project has 2 user access :

- Admin
- User

#### Admin

- can access login endpoint: `/auth/login`. `emai: admin@mail.com and password: Admin123`
- can add book
- can acces user transaction history
- can acces list of books

#### User

- can access login endpoint: `/auth/login`.
- can access register endpoint: `/auth/register`.
- can acces list of books
- can borrow book
- can return book

#### Run the Applications following these steps:

1. clone the project
2. copy file .env-example to .env

```bash
$ cp ./config/config_template.yaml ./config/config.yaml
```

3. run the sql in `db.sql` file to your database
4. install modules

```bash
$ go mod tidy

#or
$ go mod download
```

5. run the application into development mode

```bash
$ make dev
```

#### Build process:

```bash
# run build command
$ make build

# execute
$ ./target/library-api
```

TODO:

- migration
