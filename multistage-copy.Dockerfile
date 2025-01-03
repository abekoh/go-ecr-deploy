FROM golang:1.23-bullseye AS builder

WORKDIR /go/src/github.com/micin-jp/chicken-api

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /dist/app .

FROM public.ecr.aws/lambda/provided:al2023 AS runner

COPY --from=builder /dist/app ./app

ENTRYPOINT ["./app"]
