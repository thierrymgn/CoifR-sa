# Étape de construction
FROM golang:1.18 as BUILDER

# Active le comportement de module indépendant
ENV GO111MODULE=on

# Désactive cgo pour une build multiplateforme
ENV CGO_ENABLED=0

WORKDIR /go_app
COPY ./go_app .
RUN go mod download \
    && go mod verify \
    && go get -u github.com/golang-migrate/migrate/v4/cmd/migrate \
    && go build -o /build/main ./main

# Étape finale
FROM scratch as FINAL

WORKDIR /main
COPY --from=BUILDER /build/main .

ENTRYPOINT ["./main"]