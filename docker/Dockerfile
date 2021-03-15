FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

ARG PORT 
ARG DB_SERVER
ARG DB_NAME
ARG SERVER_ADDRESS
ARG PROMETHEUS_SERVER

ENV PORT ${PORT}
ENV DB_SERVER ${DB_SERVER}
ENV DB_NAME ${DB_NAME}
ENV SERVER_ADDRESS ${SERVER_ADDRESS}
ENV PROMETHEUS_SERVER ${PROMETHEUS_SERVER}

# Export necessary port
EXPOSE ${PORT}}

RUN env

# Command to run when starting the container
CMD ["/dist/main"]