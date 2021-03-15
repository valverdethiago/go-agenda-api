FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /env
COPY env/docker.env ./app.env

WORKDIR /src

COPY src .

RUN go mod download

RUN go build -o /bin/agenda

EXPOSE ${PORT}

CMD ["/bin/agenda"]
