## Considerations
This repo is an attempt to learn about hexagonal architecture while while getting to grips with a programming language that is relatively new to me.

I'll be honest, there are a few things that are different for me, especially when comparing how to do things in Python and Java. However, so far the experience is very good, Go seems like a cool language.

Probably it will take a while to fully grasp hexagonal architecture and Go, but it's a start. :grin:

## How run application?

- Run the command: `docker-compose up -d --build`

### Create table in sqlite3 db.
After this, make sure to create the table in the sqlite3.
- Run the command: `docker exec -it appproduct`
- Inside the container, run: `touch sqlite.db` then `sqlite3 sqlite.db`.

Now create the table:
- `CREATE TABLE products (
    id string PRIMARY KEY,
    name string,
    price float,
    status string
);`

### Create another container to start the webserver:

- Create another container with: `docker exec -it appproduct`
- Start the webserver: `go run main.go http`





