# Instructions
## Running the application
Run all services. The command below will start the webserver in the port 8000,
the graphql service in the port 8080 and grpc in the port 50052::
```
docker-compose up --build
```

## Testing webserver
Use the files located in /api directory to run requests against the webserver.
## Testing graphql service
Once started, the graphql service can be accessed in
`http://localhost:8080`. Currently, only two features are available:
```graphql
query orders {
    orders {
        id
        Price
        Tax
        FinalPrice
    }
}
```

```graphql
mutation createOrder {
    createOrder(
        input:{
            id: "unique_id",
            Price: 15.0,
            Tax: 1
        }
    )
    {
        id
        Price
        Tax
    }
}
```
## Testing grpc service
The grpc server can be accessed through the following command:
```
evans -r repl -p 50052
```
Once logged in, only the `CreateOrder` and `ListOrder` services are available 
for testing.

