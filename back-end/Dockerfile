FROM golang:1.18-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
ENV APP_HOME /go/src/autoexam
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /back-end-build

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
# CMD ["./run.sh"]
CMD [ "/back-end-build" ]

