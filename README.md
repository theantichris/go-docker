# go-docker

An example of using Docker for a Go project.

## Running the project

```bash
docker build -t go-docker-image .
docker run -v ~/PATH/TO/PROJECT/go-docker:/app -p 8080:8080 go-docker-image
```
