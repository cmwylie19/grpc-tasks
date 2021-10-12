# Dev

**Build Protos**
```
protoc api/tasks.proto --go_out=. --go-grpc_out=. 
```

**Run Local**
```
MONGO_URL=mongodb://localhost:27017 go run main.go
```

**List gRPC Services**
```
grpcurl -plaintext localhost:8080 list
```

**List gRPC Functions**
```
grpcurl -plaintext localhost:8080 describe todos.Todo
```

**Create a task**
```
grpcurl -plaintext -d '{"task":{"name":"Dentist @ 5:30","author":"casewylie@gmail.com"}}' localhost:8080 todos.Todo/CreateTodo 
```

**Get All Tasks**
```
grpcurl -plaintext localhost:8080 todos.Todo/GetTodos
```

**Get a specific task**
```
grpcurl -plaintext -d '{"_id":"611c1c1c53284802e9ef2c34"}'  localhost:8080 todos.Todo/GetTodo
```

**Get all tasks by a user**
```
grpcurl -plaintext -d '{"author":"casewylie@gmail.com"}'  localhost:8080 todos.Todo/GetTodosByUser
```

**Update todo**
```
grpcurl -plaintext -d '{"id":"611c1edc81c061c3c2606c6b","complete":"true"}'  localhost:8080 todos.Todo/UpdateTodo 
```

**Delete todos by user**
```
grpcurl -plaintext -d '{"author":"casewylie@gmail.com"}'  localhost:8080 todos.Todo/DeleteTodosByUser
```

**Delete a task by ID**
```
grpcurl -plaintext -d '{"_id":"611c12c747b644ceb5c3809b"}' localhost:8080 todos.Todo/DeleteTodo 
**Get a task**