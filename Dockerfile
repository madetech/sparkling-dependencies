FROM golang AS builder
WORKDIR /go/src/sparkling-dependencies
ADD . ./
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o ./compiled github.com/madetech/sparkling-dependencies/cmd/action

FROM scratch
COPY --from=builder /go/src/sparkling-dependencies/compiled/action /action
ENTRYPOINT ["/action"]
