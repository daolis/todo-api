# todo-api example

Simple web api to create, get todos, get the list of todos, and set them to done.

If you run it locally with... 

```http request
go run .
```
... you can access the api at `localhost:8080`

## Endpoints

For a detailed API description see [api-spec.yaml](api-spec.yaml)

### Get all ToDo items
Request:

```http
GET http://localhost:8080/api/toDoItems
```

### Add a new ToDo item
Request:


```http
POST http://localhost:8080/api/toDoItems
```
**Content-Type: text/plain**

### Get a single ToDo item

```http
GET http://localhost:8080/api/toDoItems/MXdzY2yO5rnZoZeg
```
### Set ToDo item to done

```http
POST http://localhost:8080/api/toDoItems/MXdzY2yO5rnZoZeg/setDone
```

## Start with docker

```shell
docker run -it --rm -p8090:8080 daolis/todo-api:latest
```

Then you can access the api at `localhost:8090/`


