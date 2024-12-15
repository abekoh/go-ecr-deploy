package main

import (
	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/aws/aws-sdk-go-v2/service/amplify"
	_ "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	_ "github.com/aws/aws-sdk-go-v2/service/appconfig"
	_ "github.com/aws/aws-sdk-go-v2/service/appconfigdata"
	_ "github.com/aws/aws-sdk-go-v2/service/appmesh"
	_ "github.com/aws/aws-sdk-go-v2/service/apprunner"
	_ "github.com/aws/aws-sdk-go-v2/service/athena"
	_ "github.com/aws/aws-sdk-go-v2/service/batch"
	_ "github.com/aws/aws-sdk-go-v2/service/bedrock"
	_ "github.com/aws/aws-sdk-go-v2/service/billing"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	_ "github.com/aws/aws-sdk-go-v2/service/codeartifact"
	_ "github.com/aws/aws-sdk-go-v2/service/codebuild"
	_ "github.com/aws/aws-sdk-go-v2/service/codecommit"
	_ "github.com/aws/aws-sdk-go-v2/service/codedeploy"
	_ "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	_ "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	_ "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	_ "github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	_ "github.com/aws/aws-sdk-go-v2/service/ebs"
	_ "github.com/aws/aws-sdk-go-v2/service/ec2"
	_ "github.com/aws/aws-sdk-go-v2/service/ecr"
	_ "github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	_ "github.com/aws/aws-sdk-go-v2/service/ecs"
	_ "github.com/aws/aws-sdk-go-v2/service/efs"
	_ "github.com/aws/aws-sdk-go-v2/service/eks"
	_ "github.com/aws/aws-sdk-go-v2/service/elasticache"
	_ "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	_ "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	_ "github.com/aws/aws-sdk-go-v2/service/eventbridge"
	_ "github.com/aws/aws-sdk-go-v2/service/glue"
	_ "github.com/aws/aws-sdk-go-v2/service/guardduty"
	_ "github.com/aws/aws-sdk-go-v2/service/iam"
	_ "github.com/aws/aws-sdk-go-v2/service/kafka"
	_ "github.com/aws/aws-sdk-go-v2/service/kafkaconnect"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesis"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	_ "github.com/aws/aws-sdk-go-v2/service/lambda"
	_ "github.com/aws/aws-sdk-go-v2/service/lightsail"
	_ "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	_ "github.com/aws/aws-sdk-go-v2/service/opensearch"
	_ "github.com/aws/aws-sdk-go-v2/service/opensearchserverless"
	_ "github.com/aws/aws-sdk-go-v2/service/rds"
	_ "github.com/aws/aws-sdk-go-v2/service/rdsdata"
	_ "github.com/aws/aws-sdk-go-v2/service/redshift"
	_ "github.com/aws/aws-sdk-go-v2/service/redshiftdata"
	_ "github.com/aws/aws-sdk-go-v2/service/redshiftserverless"
	_ "github.com/aws/aws-sdk-go-v2/service/route53"
	_ "github.com/aws/aws-sdk-go-v2/service/route53domains"
	_ "github.com/aws/aws-sdk-go-v2/service/route53profiles"
	_ "github.com/aws/aws-sdk-go-v2/service/route53recoverycluster"
	_ "github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	_ "github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	_ "github.com/aws/aws-sdk-go-v2/service/route53resolver"
	_ "github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/aws/aws-sdk-go-v2/service/s3control"
	_ "github.com/aws/aws-sdk-go-v2/service/s3outposts"
	_ "github.com/aws/aws-sdk-go-v2/service/s3tables"
	_ "github.com/aws/aws-sdk-go-v2/service/sagemaker"
	_ "github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime"
	_ "github.com/aws/aws-sdk-go-v2/service/sagemakeredge"
	_ "github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime"
	_ "github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial"
	_ "github.com/aws/aws-sdk-go-v2/service/sagemakermetrics"
	_ "github.com/aws/aws-sdk-go-v2/service/sagemakerruntime"
	_ "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/aws/aws-sdk-go-v2/service/sesv2"
	_ "github.com/aws/aws-sdk-go-v2/service/sms"
	_ "github.com/aws/aws-sdk-go-v2/service/sns"
	_ "github.com/aws/aws-sdk-go-v2/service/sqs"
	_ "github.com/aws/aws-sdk-go-v2/service/ssm"
	_ "github.com/aws/aws-sdk-go-v2/service/wafv2"
)

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	decodedBody, err := base64.StdEncoding.DecodeString(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to decode body",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(decodedBody),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
