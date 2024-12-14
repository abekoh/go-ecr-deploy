package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	ddlambda "github.com/DataDog/datadog-lambda-go"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	_ "github.com/aws/aws-sdk-go-v2/service/account"
	_ "github.com/aws/aws-sdk-go-v2/service/acm"
	_ "github.com/aws/aws-sdk-go-v2/service/acmpca"
	_ "github.com/aws/aws-sdk-go-v2/service/amp"
	_ "github.com/aws/aws-sdk-go-v2/service/amplify"
	_ "github.com/aws/aws-sdk-go-v2/service/amplifybackend"
	_ "github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder"
	_ "github.com/aws/aws-sdk-go-v2/service/apigateway"
	_ "github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	_ "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	_ "github.com/aws/aws-sdk-go-v2/service/appconfig"
	_ "github.com/aws/aws-sdk-go-v2/service/appconfigdata"
	_ "github.com/aws/aws-sdk-go-v2/service/appfabric"
	_ "github.com/aws/aws-sdk-go-v2/service/appflow"
	_ "github.com/aws/aws-sdk-go-v2/service/appintegrations"
	_ "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	_ "github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler"
	_ "github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
	_ "github.com/aws/aws-sdk-go-v2/service/applicationinsights"
	_ "github.com/aws/aws-sdk-go-v2/service/applicationsignals"
	_ "github.com/aws/aws-sdk-go-v2/service/appmesh"
	_ "github.com/aws/aws-sdk-go-v2/service/apprunner"
	_ "github.com/aws/aws-sdk-go-v2/service/appstream"
	_ "github.com/aws/aws-sdk-go-v2/service/appsync"
	_ "github.com/aws/aws-sdk-go-v2/service/apptest"
	_ "github.com/aws/aws-sdk-go-v2/service/arczonalshift"
	_ "github.com/aws/aws-sdk-go-v2/service/artifact"
	_ "github.com/aws/aws-sdk-go-v2/service/athena"
	_ "github.com/aws/aws-sdk-go-v2/service/auditmanager"
	_ "github.com/aws/aws-sdk-go-v2/service/autoscaling"
	_ "github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	_ "github.com/aws/aws-sdk-go-v2/service/b2bi"
	_ "github.com/aws/aws-sdk-go-v2/service/backup"
	_ "github.com/aws/aws-sdk-go-v2/service/backupgateway"
	_ "github.com/aws/aws-sdk-go-v2/service/batch"
	_ "github.com/aws/aws-sdk-go-v2/service/bcmdataexports"
	_ "github.com/aws/aws-sdk-go-v2/service/bcmpricingcalculator"
	_ "github.com/aws/aws-sdk-go-v2/service/bedrock"
	_ "github.com/aws/aws-sdk-go-v2/service/bedrockagent"
	_ "github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
	_ "github.com/aws/aws-sdk-go-v2/service/bedrockdataautomation"
	_ "github.com/aws/aws-sdk-go-v2/service/bedrockdataautomationruntime"
	_ "github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	_ "github.com/aws/aws-sdk-go-v2/service/billing"
	_ "github.com/aws/aws-sdk-go-v2/service/billingconductor"
	_ "github.com/aws/aws-sdk-go-v2/service/braket"
	_ "github.com/aws/aws-sdk-go-v2/service/budgets"
	_ "github.com/aws/aws-sdk-go-v2/service/chatbot"
	_ "github.com/aws/aws-sdk-go-v2/service/chime"
	_ "github.com/aws/aws-sdk-go-v2/service/chimesdkidentity"
	_ "github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines"
	_ "github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings"
	_ "github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging"
	_ "github.com/aws/aws-sdk-go-v2/service/chimesdkvoice"
	_ "github.com/aws/aws-sdk-go-v2/service/cleanrooms"
	_ "github.com/aws/aws-sdk-go-v2/service/cleanroomsml"
	_ "github.com/aws/aws-sdk-go-v2/service/cloud9"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	_ "github.com/aws/aws-sdk-go-v2/service/clouddirectory"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudhsm"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudtraildata"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	_ "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	_ "github.com/aws/aws-sdk-go-v2/service/codeartifact"
	_ "github.com/aws/aws-sdk-go-v2/service/codebuild"
	_ "github.com/aws/aws-sdk-go-v2/service/codecatalyst"
	_ "github.com/aws/aws-sdk-go-v2/service/codecommit"
	_ "github.com/aws/aws-sdk-go-v2/service/codeconnections"
	_ "github.com/aws/aws-sdk-go-v2/service/codedeploy"
	_ "github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
	_ "github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
	_ "github.com/aws/aws-sdk-go-v2/service/codegurusecurity"
	_ "github.com/aws/aws-sdk-go-v2/service/codepipeline"
	_ "github.com/aws/aws-sdk-go-v2/service/codestarconnections"
	_ "github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	_ "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	_ "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	_ "github.com/aws/aws-sdk-go-v2/service/cognitosync"
	_ "github.com/aws/aws-sdk-go-v2/service/comprehend"
	_ "github.com/aws/aws-sdk-go-v2/service/comprehendmedical"
	_ "github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	_ "github.com/aws/aws-sdk-go-v2/service/configservice"
	_ "github.com/aws/aws-sdk-go-v2/service/connect"
	_ "github.com/aws/aws-sdk-go-v2/service/connectcampaigns"
	_ "github.com/aws/aws-sdk-go-v2/service/connectcampaignsv2"
	_ "github.com/aws/aws-sdk-go-v2/service/connectcases"
	_ "github.com/aws/aws-sdk-go-v2/service/connectcontactlens"
	_ "github.com/aws/aws-sdk-go-v2/service/connectparticipant"
	_ "github.com/aws/aws-sdk-go-v2/service/controlcatalog"
	_ "github.com/aws/aws-sdk-go-v2/service/controltower"
	_ "github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	_ "github.com/aws/aws-sdk-go-v2/service/costexplorer"
	_ "github.com/aws/aws-sdk-go-v2/service/costoptimizationhub"
	_ "github.com/aws/aws-sdk-go-v2/service/customerprofiles"
	_ "github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	_ "github.com/aws/aws-sdk-go-v2/service/databrew"
	_ "github.com/aws/aws-sdk-go-v2/service/dataexchange"
	_ "github.com/aws/aws-sdk-go-v2/service/datapipeline"
	_ "github.com/aws/aws-sdk-go-v2/service/datasync"
	_ "github.com/aws/aws-sdk-go-v2/service/datazone"
	_ "github.com/aws/aws-sdk-go-v2/service/dax"
	_ "github.com/aws/aws-sdk-go-v2/service/deadline"
	_ "github.com/aws/aws-sdk-go-v2/service/detective"
	_ "github.com/aws/aws-sdk-go-v2/service/devicefarm"
	_ "github.com/aws/aws-sdk-go-v2/service/devopsguru"
	_ "github.com/aws/aws-sdk-go-v2/service/directconnect"
	_ "github.com/aws/aws-sdk-go-v2/service/directoryservice"
	_ "github.com/aws/aws-sdk-go-v2/service/directoryservicedata"
	_ "github.com/aws/aws-sdk-go-v2/service/dlm"
	_ "github.com/aws/aws-sdk-go-v2/service/docdb"
	_ "github.com/aws/aws-sdk-go-v2/service/docdbelastic"
	_ "github.com/aws/aws-sdk-go-v2/service/drs"
	_ "github.com/aws/aws-sdk-go-v2/service/dsql"
	_ "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	_ "github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	_ "github.com/aws/aws-sdk-go-v2/service/ebs"
	_ "github.com/aws/aws-sdk-go-v2/service/ec2"
	_ "github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	_ "github.com/aws/aws-sdk-go-v2/service/ecr"
	_ "github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	_ "github.com/aws/aws-sdk-go-v2/service/ecs"
	_ "github.com/aws/aws-sdk-go-v2/service/efs"
	_ "github.com/aws/aws-sdk-go-v2/service/eks"
	_ "github.com/aws/aws-sdk-go-v2/service/eksauth"
	_ "github.com/aws/aws-sdk-go-v2/service/elasticache"
	_ "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	_ "github.com/aws/aws-sdk-go-v2/service/elasticinference"
	_ "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	_ "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	_ "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	_ "github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	_ "github.com/aws/aws-sdk-go-v2/service/emr"
	_ "github.com/aws/aws-sdk-go-v2/service/emrcontainers"
	_ "github.com/aws/aws-sdk-go-v2/service/emrserverless"
	_ "github.com/aws/aws-sdk-go-v2/service/entityresolution"
	_ "github.com/aws/aws-sdk-go-v2/service/eventbridge"
	_ "github.com/aws/aws-sdk-go-v2/service/evidently"
	_ "github.com/aws/aws-sdk-go-v2/service/finspace"
	_ "github.com/aws/aws-sdk-go-v2/service/finspacedata"
	_ "github.com/aws/aws-sdk-go-v2/service/firehose"
	_ "github.com/aws/aws-sdk-go-v2/service/fis"
	_ "github.com/aws/aws-sdk-go-v2/service/fms"
	_ "github.com/aws/aws-sdk-go-v2/service/forecast"
	_ "github.com/aws/aws-sdk-go-v2/service/forecastquery"
	_ "github.com/aws/aws-sdk-go-v2/service/frauddetector"
	_ "github.com/aws/aws-sdk-go-v2/service/freetier"
	_ "github.com/aws/aws-sdk-go-v2/service/fsx"
	_ "github.com/aws/aws-sdk-go-v2/service/gamelift"
	_ "github.com/aws/aws-sdk-go-v2/service/geomaps"
	_ "github.com/aws/aws-sdk-go-v2/service/geoplaces"
	_ "github.com/aws/aws-sdk-go-v2/service/georoutes"
	_ "github.com/aws/aws-sdk-go-v2/service/glacier"
	_ "github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	_ "github.com/aws/aws-sdk-go-v2/service/glue"
	_ "github.com/aws/aws-sdk-go-v2/service/grafana"
	_ "github.com/aws/aws-sdk-go-v2/service/greengrass"
	_ "github.com/aws/aws-sdk-go-v2/service/greengrassv2"
	_ "github.com/aws/aws-sdk-go-v2/service/groundstation"
	_ "github.com/aws/aws-sdk-go-v2/service/guardduty"
	_ "github.com/aws/aws-sdk-go-v2/service/health"
	_ "github.com/aws/aws-sdk-go-v2/service/healthlake"
	_ "github.com/aws/aws-sdk-go-v2/service/iam"
	_ "github.com/aws/aws-sdk-go-v2/service/identitystore"
	_ "github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	_ "github.com/aws/aws-sdk-go-v2/service/inspector"
	_ "github.com/aws/aws-sdk-go-v2/service/inspector2"
	_ "github.com/aws/aws-sdk-go-v2/service/inspectorscan"
	_ "github.com/aws/aws-sdk-go-v2/service/internetmonitor"
	_ "github.com/aws/aws-sdk-go-v2/service/invoicing"
	_ "github.com/aws/aws-sdk-go-v2/service/iot"
	_ "github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
	_ "github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
	_ "github.com/aws/aws-sdk-go-v2/service/iotanalytics"
	_ "github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	_ "github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor"
	_ "github.com/aws/aws-sdk-go-v2/service/iotevents"
	_ "github.com/aws/aws-sdk-go-v2/service/ioteventsdata"
	_ "github.com/aws/aws-sdk-go-v2/service/iotfleethub"
	_ "github.com/aws/aws-sdk-go-v2/service/iotfleetwise"
	_ "github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
	_ "github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling"
	_ "github.com/aws/aws-sdk-go-v2/service/iotsitewise"
	_ "github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
	_ "github.com/aws/aws-sdk-go-v2/service/iottwinmaker"
	_ "github.com/aws/aws-sdk-go-v2/service/iotwireless"
	_ "github.com/aws/aws-sdk-go-v2/service/ivs"
	_ "github.com/aws/aws-sdk-go-v2/service/ivschat"
	_ "github.com/aws/aws-sdk-go-v2/service/ivsrealtime"
	_ "github.com/aws/aws-sdk-go-v2/service/kafka"
	_ "github.com/aws/aws-sdk-go-v2/service/kafkaconnect"
	_ "github.com/aws/aws-sdk-go-v2/service/kendra"
	_ "github.com/aws/aws-sdk-go-v2/service/kendraranking"
	_ "github.com/aws/aws-sdk-go-v2/service/keyspaces"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesis"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling"
	_ "github.com/aws/aws-sdk-go-v2/service/kinesisvideowebrtcstorage"
	_ "github.com/aws/aws-sdk-go-v2/service/kms"
	_ "github.com/aws/aws-sdk-go-v2/service/lakeformation"
	_ "github.com/aws/aws-sdk-go-v2/service/lambda"
	_ "github.com/aws/aws-sdk-go-v2/service/launchwizard"
	_ "github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	_ "github.com/aws/aws-sdk-go-v2/service/lexmodelsv2"
	_ "github.com/aws/aws-sdk-go-v2/service/lexruntimeservice"
	_ "github.com/aws/aws-sdk-go-v2/service/lexruntimev2"
	_ "github.com/aws/aws-sdk-go-v2/service/licensemanager"
	_ "github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions"
	_ "github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions"
	_ "github.com/aws/aws-sdk-go-v2/service/lightsail"
	_ "github.com/aws/aws-sdk-go-v2/service/location"
	_ "github.com/aws/aws-sdk-go-v2/service/lookoutequipment"
	_ "github.com/aws/aws-sdk-go-v2/service/lookoutmetrics"
	_ "github.com/aws/aws-sdk-go-v2/service/lookoutvision"
	_ "github.com/aws/aws-sdk-go-v2/service/m2"
	_ "github.com/aws/aws-sdk-go-v2/service/machinelearning"
	_ "github.com/aws/aws-sdk-go-v2/service/macie2"
	_ "github.com/aws/aws-sdk-go-v2/service/mailmanager"
	_ "github.com/aws/aws-sdk-go-v2/service/managedblockchain"
	_ "github.com/aws/aws-sdk-go-v2/service/managedblockchainquery"
	_ "github.com/aws/aws-sdk-go-v2/service/marketplaceagreement"
	_ "github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
	_ "github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
	_ "github.com/aws/aws-sdk-go-v2/service/marketplacedeployment"
	_ "github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
	_ "github.com/aws/aws-sdk-go-v2/service/marketplacemetering"
	_ "github.com/aws/aws-sdk-go-v2/service/marketplacereporting"
	_ "github.com/aws/aws-sdk-go-v2/service/mediaconnect"
	_ "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	_ "github.com/aws/aws-sdk-go-v2/service/medialive"
	_ "github.com/aws/aws-sdk-go-v2/service/mediapackage"
	_ "github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
	_ "github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
	_ "github.com/aws/aws-sdk-go-v2/service/mediastore"
	_ "github.com/aws/aws-sdk-go-v2/service/mediastoredata"
	_ "github.com/aws/aws-sdk-go-v2/service/mediatailor"
	_ "github.com/aws/aws-sdk-go-v2/service/medicalimaging"
	_ "github.com/aws/aws-sdk-go-v2/service/memorydb"
	_ "github.com/aws/aws-sdk-go-v2/service/mgn"
	_ "github.com/aws/aws-sdk-go-v2/service/migrationhub"
	_ "github.com/aws/aws-sdk-go-v2/service/migrationhubconfig"
	_ "github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator"
	_ "github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces"
	_ "github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy"
	_ "github.com/aws/aws-sdk-go-v2/service/mq"
	_ "github.com/aws/aws-sdk-go-v2/service/mturk"
	_ "github.com/aws/aws-sdk-go-v2/service/mwaa"
	_ "github.com/aws/aws-sdk-go-v2/service/neptune"
	_ "github.com/aws/aws-sdk-go-v2/service/neptunedata"
	_ "github.com/aws/aws-sdk-go-v2/service/neptunegraph"
	_ "github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	_ "github.com/aws/aws-sdk-go-v2/service/networkflowmonitor"
	_ "github.com/aws/aws-sdk-go-v2/service/networkmanager"
	_ "github.com/aws/aws-sdk-go-v2/service/networkmonitor"
	_ "github.com/aws/aws-sdk-go-v2/service/notifications"
	_ "github.com/aws/aws-sdk-go-v2/service/notificationscontacts"
	_ "github.com/aws/aws-sdk-go-v2/service/oam"
	_ "github.com/aws/aws-sdk-go-v2/service/observabilityadmin"
	_ "github.com/aws/aws-sdk-go-v2/service/omics"
	_ "github.com/aws/aws-sdk-go-v2/service/opensearch"
	_ "github.com/aws/aws-sdk-go-v2/service/opensearchserverless"
	_ "github.com/aws/aws-sdk-go-v2/service/opsworks"
	_ "github.com/aws/aws-sdk-go-v2/service/opsworkscm"
	_ "github.com/aws/aws-sdk-go-v2/service/organizations"
	_ "github.com/aws/aws-sdk-go-v2/service/osis"
	_ "github.com/aws/aws-sdk-go-v2/service/outposts"
	_ "github.com/aws/aws-sdk-go-v2/service/panorama"
	_ "github.com/aws/aws-sdk-go-v2/service/partnercentralselling"
	_ "github.com/aws/aws-sdk-go-v2/service/paymentcryptography"
	_ "github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata"
	_ "github.com/aws/aws-sdk-go-v2/service/pcaconnectorad"
	_ "github.com/aws/aws-sdk-go-v2/service/pcaconnectorscep"
	_ "github.com/aws/aws-sdk-go-v2/service/pcs"
	_ "github.com/aws/aws-sdk-go-v2/service/personalize"
	_ "github.com/aws/aws-sdk-go-v2/service/personalizeevents"
	_ "github.com/aws/aws-sdk-go-v2/service/personalizeruntime"
	_ "github.com/aws/aws-sdk-go-v2/service/pi"
	_ "github.com/aws/aws-sdk-go-v2/service/pinpoint"
	_ "github.com/aws/aws-sdk-go-v2/service/pinpointemail"
	_ "github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice"
	_ "github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2"
	_ "github.com/aws/aws-sdk-go-v2/service/pipes"
	_ "github.com/aws/aws-sdk-go-v2/service/polly"
	_ "github.com/aws/aws-sdk-go-v2/service/pricing"
	_ "github.com/aws/aws-sdk-go-v2/service/privatenetworks"
	_ "github.com/aws/aws-sdk-go-v2/service/proton"
	_ "github.com/aws/aws-sdk-go-v2/service/qapps"
	_ "github.com/aws/aws-sdk-go-v2/service/qbusiness"
	_ "github.com/aws/aws-sdk-go-v2/service/qconnect"
	_ "github.com/aws/aws-sdk-go-v2/service/qldb"
	_ "github.com/aws/aws-sdk-go-v2/service/qldbsession"
	_ "github.com/aws/aws-sdk-go-v2/service/quicksight"
	_ "github.com/aws/aws-sdk-go-v2/service/ram"
	_ "github.com/aws/aws-sdk-go-v2/service/rbin"
	_ "github.com/aws/aws-sdk-go-v2/service/rds"
	_ "github.com/aws/aws-sdk-go-v2/service/rdsdata"
	_ "github.com/aws/aws-sdk-go-v2/service/redshift"
	_ "github.com/aws/aws-sdk-go-v2/service/redshiftdata"
	_ "github.com/aws/aws-sdk-go-v2/service/redshiftserverless"
	_ "github.com/aws/aws-sdk-go-v2/service/rekognition"
	_ "github.com/aws/aws-sdk-go-v2/service/repostspace"
	_ "github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	_ "github.com/aws/aws-sdk-go-v2/service/resourceexplorer2"
	_ "github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	_ "github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	_ "github.com/aws/aws-sdk-go-v2/service/robomaker"
	_ "github.com/aws/aws-sdk-go-v2/service/rolesanywhere"
	_ "github.com/aws/aws-sdk-go-v2/service/route53"
	_ "github.com/aws/aws-sdk-go-v2/service/route53domains"
	_ "github.com/aws/aws-sdk-go-v2/service/route53profiles"
	_ "github.com/aws/aws-sdk-go-v2/service/route53recoverycluster"
	_ "github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	_ "github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	_ "github.com/aws/aws-sdk-go-v2/service/route53resolver"
	_ "github.com/aws/aws-sdk-go-v2/service/rum"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	_ "github.com/aws/aws-sdk-go-v2/service/savingsplans"
	_ "github.com/aws/aws-sdk-go-v2/service/scheduler"
	_ "github.com/aws/aws-sdk-go-v2/service/schemas"
	_ "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/aws/aws-sdk-go-v2/service/securityhub"
	_ "github.com/aws/aws-sdk-go-v2/service/securityir"
	_ "github.com/aws/aws-sdk-go-v2/service/securitylake"
	_ "github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
	_ "github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	_ "github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	_ "github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	_ "github.com/aws/aws-sdk-go-v2/service/servicequotas"
	_ "github.com/aws/aws-sdk-go-v2/service/ses"
	_ "github.com/aws/aws-sdk-go-v2/service/sesv2"
	_ "github.com/aws/aws-sdk-go-v2/service/sfn"
	_ "github.com/aws/aws-sdk-go-v2/service/shield"
	_ "github.com/aws/aws-sdk-go-v2/service/signer"
	_ "github.com/aws/aws-sdk-go-v2/service/simspaceweaver"
	_ "github.com/aws/aws-sdk-go-v2/service/sms"
	_ "github.com/aws/aws-sdk-go-v2/service/snowball"
	_ "github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement"
	_ "github.com/aws/aws-sdk-go-v2/service/sns"
	_ "github.com/aws/aws-sdk-go-v2/service/socialmessaging"
	_ "github.com/aws/aws-sdk-go-v2/service/sqs"
	_ "github.com/aws/aws-sdk-go-v2/service/ssm"
	_ "github.com/aws/aws-sdk-go-v2/service/ssmcontacts"
	_ "github.com/aws/aws-sdk-go-v2/service/ssmincidents"
	_ "github.com/aws/aws-sdk-go-v2/service/ssmquicksetup"
	_ "github.com/aws/aws-sdk-go-v2/service/ssmsap"
	_ "github.com/aws/aws-sdk-go-v2/service/sso"
	_ "github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	_ "github.com/aws/aws-sdk-go-v2/service/ssooidc"
	_ "github.com/aws/aws-sdk-go-v2/service/storagegateway"
	_ "github.com/aws/aws-sdk-go-v2/service/sts"
	_ "github.com/aws/aws-sdk-go-v2/service/supplychain"
	_ "github.com/aws/aws-sdk-go-v2/service/support"
	_ "github.com/aws/aws-sdk-go-v2/service/supportapp"
	_ "github.com/aws/aws-sdk-go-v2/service/swf"
	_ "github.com/aws/aws-sdk-go-v2/service/synthetics"
	_ "github.com/aws/aws-sdk-go-v2/service/taxsettings"
	_ "github.com/aws/aws-sdk-go-v2/service/textract"
	_ "github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb"
	_ "github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	_ "github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	_ "github.com/aws/aws-sdk-go-v2/service/tnb"
	_ "github.com/aws/aws-sdk-go-v2/service/transcribe"
	_ "github.com/aws/aws-sdk-go-v2/service/transcribestreaming"
	_ "github.com/aws/aws-sdk-go-v2/service/transfer"
	_ "github.com/aws/aws-sdk-go-v2/service/translate"
	_ "github.com/aws/aws-sdk-go-v2/service/trustedadvisor"
	_ "github.com/aws/aws-sdk-go-v2/service/verifiedpermissions"
	_ "github.com/aws/aws-sdk-go-v2/service/voiceid"
	_ "github.com/aws/aws-sdk-go-v2/service/vpclattice"
	_ "github.com/aws/aws-sdk-go-v2/service/waf"
	_ "github.com/aws/aws-sdk-go-v2/service/wafregional"
	_ "github.com/aws/aws-sdk-go-v2/service/wafv2"
	_ "github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	_ "github.com/aws/aws-sdk-go-v2/service/wisdom"
	_ "github.com/aws/aws-sdk-go-v2/service/workdocs"
	_ "github.com/aws/aws-sdk-go-v2/service/workmail"
	_ "github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
	_ "github.com/aws/aws-sdk-go-v2/service/workspaces"
	_ "github.com/aws/aws-sdk-go-v2/service/workspacesthinclient"
	_ "github.com/aws/aws-sdk-go-v2/service/workspacesweb"
	_ "github.com/aws/aws-sdk-go-v2/service/xray"
)

type Order struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Item    string  `json:"item"`
}

var (
	s3Client *s3.Client
)

func init() {
	// Initialize the S3 client outside of the handler, during the init phase
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3Client = s3.NewFromConfig(cfg)
}

func uploadReceiptToS3(ctx context.Context, bucketName, key, receiptContent string) error {
	_, err := s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &key,
		Body:   strings.NewReader(receiptContent),
	})
	if err != nil {
		log.Printf("Failed to upload receipt to S3: %v", err)
		return err
	}
	return nil
}

func handleRequest(ctx context.Context, event json.RawMessage) error {
	// Parse the input event
	var order Order
	if err := json.Unmarshal(event, &order); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return err
	}

	// Access environment variables
	bucketName := os.Getenv("RECEIPT_BUCKET")
	if bucketName == "" {
		log.Printf("RECEIPT_BUCKET environment variable is not set")
		return fmt.Errorf("missing required environment variable RECEIPT_BUCKET")
	}

	// Create the receipt content and key destination
	receiptContent := fmt.Sprintf("OrderID: %s\nAmount: $%.2f\nItem: %s",
		order.OrderID, order.Amount, order.Item)
	key := "receipts/" + order.OrderID + ".txt"

	// Upload the receipt to S3 using the helper method
	if err := uploadReceiptToS3(ctx, bucketName, key, receiptContent); err != nil {
		return err
	}

	log.Printf("Successfully processed order %s and stored receipt in S3 bucket %s", order.OrderID, bucketName)
	return nil
}

func main() {
	lambda.Start(ddlambda.WrapFunction(handleRequest, nil))
}
