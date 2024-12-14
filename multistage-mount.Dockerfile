FROM golang:1.23-bullseye AS builder

WORKDIR /go/src/github.com/micin-jp/chicken-api

RUN --mount=type=cache,target=/go/pkg/mod,sharing=locked \
    --mount=type=cache,target=/root/.cache/go-build,sharing=locked \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,source=.,target=. \
    go build -o /dist/app .

FROM public.ecr.aws/lambda/provided:al2023 AS runner

COPY --from=builder /dist/app ./app

ENTRYPOINT ["./app"]
