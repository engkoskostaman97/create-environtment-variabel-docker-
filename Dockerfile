FROM golang:1.9

COPY web.go /app/main.go
CMD [ "go","run","/app/main.go" ]