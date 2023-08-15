# GoKu (Golang Kubernetes)

What a nice name for a pretty useless application.

GoKu is a Gin Go API to deploy Pod in a Kubernetes cluster.

# Build and Usage

```bash
go build main.go
./main
```

OR

```bash
docker build -t goku .
docker run -p8080:8080 goku
```