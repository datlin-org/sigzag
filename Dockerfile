FROM golang:latest AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

WORKDIR /src/cmd/

RUN go build main.go

RUN mv main sigzag

FROM ubuntu:latest

WORKDIR /opt/app/

COPY --from=build /src/cmd/sigzag /opt/app/sigzag

CMD ["./sigzag"]