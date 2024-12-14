FROM public.ecr.aws/lambda/provided:al2023 AS runner

COPY  /dist/app ./app
# set datadog extension
# see: https://docs.datadoghq.com/ja/serverless/aws_lambda/installation/go/?tab=containerimage
COPY --from=public.ecr.aws/datadog/lambda-extension:66 /opt/. /opt/

ENTRYPOINT ["./app"]
