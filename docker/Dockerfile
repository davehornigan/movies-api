FROM golang:1.21-alpine as build
LABEL authors="Dave Hornigan<dave@hornigan.com>"

WORKDIR /app

COPY . /app/
RUN go mod download

RUN go build -o /app/build/api /app/cmd/api/main.go

FROM scratch as api

WORKDIR /app
COPY --from=build /app/build/api /app/

CMD ["/app/api"]