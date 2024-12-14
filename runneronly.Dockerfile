FROM public.ecr.aws/lambda/provided:al2023 AS runner

COPY  ./dist/app ./app

ENTRYPOINT ["./app"]
