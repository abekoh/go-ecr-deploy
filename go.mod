module github.com/abekoh/go-ecr-deploy

go 1.23.4

require (
	github.com/aws/aws-lambda-go v1.46.0
	github.com/aws/aws-sdk-go-v2/service/amplify v1.27.5
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.24.7
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.36.1
	github.com/aws/aws-sdk-go-v2/service/appconfigdata v1.18.7
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.29.7
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.32.7
	github.com/aws/aws-sdk-go-v2/service/athena v1.49.0
	github.com/aws/aws-sdk-go-v2/service/batch v1.48.2
	github.com/aws/aws-sdk-go-v2/service/bedrock v1.25.1
	github.com/aws/aws-sdk-go-v2/service/billing v1.0.2
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.56.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.43.1
	github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore v1.8.7
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.43.3
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.27.7
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.45.0
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.33.7
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.49.3
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.27.7
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.29.7
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.27.8
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.48.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.38.0
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.24.9
	github.com/aws/aws-sdk-go-v2/service/ebs v1.27.7
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.197.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.36.7
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.27.7
	github.com/aws/aws-sdk-go-v2/service/ecs v1.52.2
	github.com/aws/aws-sdk-go-v2/service/efs v1.34.1
	github.com/aws/aws-sdk-go-v2/service/eks v1.54.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.44.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.43.1
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.32.7
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.36.0
	github.com/aws/aws-sdk-go-v2/service/glue v1.104.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.52.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.38.2
	github.com/aws/aws-sdk-go-v2/service/kafka v1.38.7
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.21.7
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.32.7
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.31.8
	github.com/aws/aws-sdk-go-v2/service/lambda v1.69.1
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.42.7
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.63.1
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.45.0
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.17.4
	github.com/aws/aws-sdk-go-v2/service/rds v1.92.0
	github.com/aws/aws-sdk-go-v2/service/rdsdata v1.26.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.53.0
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.31.4
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.25.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.46.3
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.28.0
	github.com/aws/aws-sdk-go-v2/service/route53profiles v1.4.7
	github.com/aws/aws-sdk-go-v2/service/route53recoverycluster v1.23.7
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.25.7
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.21.7
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.34.2
	github.com/aws/aws-sdk-go-v2/service/s3 v1.71.0
	github.com/aws/aws-sdk-go-v2/service/s3control v1.52.0
	github.com/aws/aws-sdk-go-v2/service/s3outposts v1.28.7
	github.com/aws/aws-sdk-go-v2/service/s3tables v1.0.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.169.0
	github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime v1.27.7
	github.com/aws/aws-sdk-go-v2/service/sagemakeredge v1.25.7
	github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime v1.29.7
	github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial v1.14.7
	github.com/aws/aws-sdk-go-v2/service/sagemakermetrics v1.12.7
	github.com/aws/aws-sdk-go-v2/service/sagemakerruntime v1.32.7
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.34.7
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.40.0
	github.com/aws/aws-sdk-go-v2/service/sms v1.24.7
	github.com/aws/aws-sdk-go-v2/service/sns v1.33.7
	github.com/aws/aws-sdk-go-v2/service/sqs v1.37.2
	github.com/aws/aws-sdk-go-v2/service/ssm v1.56.1
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.55.6
)

require (
	github.com/aws/aws-sdk-go-v2 v1.32.6 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.4.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.6 // indirect
	github.com/aws/smithy-go v1.22.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
