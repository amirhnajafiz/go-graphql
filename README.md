# go-graphql

Creating a GraphQL server with **Golang**.

## What is the project?
In this project I implemented a simple library system database,
where you can query to database with _GraphQL_.

It uses _sqlite3_ to create the main database, and then it will
create the library schema for books and authors.

After that you can submit your queries in _GraphQL_ type and get the
response from the server.

One important thing to note is that GraphQL is not a query language like our
traditional SQL. 
It is an abstraction that sits in-front of our APIs and is not tied to 
any specific database or storage engine.

This is actually really cool. 
We can stand up a GraphQL server that interacts with existing services 
and then build around this new GraphQL server instead of having 
to worry about modifying existing REST APIs.

## How to use?
First set up the server by the following command:
```shell
docker-compose up -d
```

After that you can check the server status in _**localhost:5000**_

If you get the guid page, you are good to go. You can submit your queries
by the following request:
```shell
curl -X POST -H "Content-Type: application/json" \
    -d '{"query": "[GraphQL query]"' \
    https://localhost:5000/query
```

Just insert your GraphQL query into request.

## Configs
First copy the configs file:
```shell
cp ./configs/example.yaml ./config.yaml
```

Now you can set the configs to set up the server:
```yaml
port: "[server port]"
proxy: "[proxy IP]"
...
```