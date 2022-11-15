FROM golang AS builder
WORKDIR /go/src/sparkling-dependencies
ADD go.mod ./
ADD go.sum ./
RUN go mod download
ADD . ./

FROM builder AS test

FROM builder AS compiled
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o ./compiled github.com/madetech/sparkling-dependencies/cmd/action

FROM scratch
COPY --from=compiled /go/src/sparkling-dependencies/compiled/action /action
ENTRYPOINT ["/action"]
