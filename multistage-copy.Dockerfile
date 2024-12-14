FROM golang:1.23-bullseye AS builder

WORKDIR /go/src/github.com/micin-jp/chicken-api

COPY go.mod go.sum ./
RUN go mod download -x

COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /dist/app .

FROM public.ecr.aws/lambda/provided:al2023 AS runner

# set datadog extension
# see: https://docs.datadoghq.com/ja/serverless/aws_lambda/installation/go/?tab=containerimage
COPY --from=public.ecr.aws/datadog/lambda-extension:66 /opt/. /opt/
COPY --from=builder /dist/app ./app

ENTRYPOINT ["./app"]
