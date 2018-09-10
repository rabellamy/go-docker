# Go Docker

This project assumes the below directory structure where the Dockerfile is a sibiling
of `main.go`.

```bash
.
├── Dockerfile
├── README.md
└── main.go
```

To build the Dockerfile run:

```bash
$ docker build -t [repository]/[name]:[tag] [path-to-Dockerfile]
```

To run the container:

```bash
$ docker run -p 127.0.0.1:8080:8080  go-docker:latest
```

To test the container in a new terminal run:

```bash
$ curl localhost:8080/healthcheck
{"status":"OK"}
```
