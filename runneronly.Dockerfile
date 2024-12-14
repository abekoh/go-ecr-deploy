FROM public.ecr.aws/lambda/provided:al2023 AS runner

# set datadog extension
# see: https://docs.datadoghq.com/ja/serverless/aws_lambda/installation/go/?tab=containerimage
COPY --from=public.ecr.aws/datadog/lambda-extension:66 /opt/. /opt/
COPY  /dist/app ./app

ENTRYPOINT ["./app"]
