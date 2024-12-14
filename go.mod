module github.com/abekoh/go-ecr-deploy

go 1.23.4

require (
	github.com/DataDog/datadog-lambda-go v1.20.0
	github.com/aws/aws-lambda-go v1.47.0
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.36.2
	github.com/aws/aws-sdk-go-v2/service/account v1.21.7
	github.com/aws/aws-sdk-go-v2/service/acm v1.30.7
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.37.8
	github.com/aws/aws-sdk-go-v2/service/amp v1.30.4
	github.com/aws/aws-sdk-go-v2/service/amplify v1.27.5
	github.com/aws/aws-sdk-go-v2/service/amplifybackend v1.27.7
	github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder v1.23.7
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.28.1
	github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi v1.23.7
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.24.7
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.36.1
	github.com/aws/aws-sdk-go-v2/service/appconfigdata v1.18.7
	github.com/aws/aws-sdk-go-v2/service/appfabric v1.11.7
	github.com/aws/aws-sdk-go-v2/service/appflow v1.45.8
	github.com/aws/aws-sdk-go-v2/service/appintegrations v1.30.7
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.34.2
	github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler v1.21.7
	github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice v1.29.1
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.29.5
	github.com/aws/aws-sdk-go-v2/service/applicationsignals v1.7.2
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.29.7
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.32.7
	github.com/aws/aws-sdk-go-v2/service/appstream v1.41.7
	github.com/aws/aws-sdk-go-v2/service/appsync v1.40.2
	github.com/aws/aws-sdk-go-v2/service/apptest v1.4.7
	github.com/aws/aws-sdk-go-v2/service/arczonalshift v1.14.7
	github.com/aws/aws-sdk-go-v2/service/artifact v1.7.0
	github.com/aws/aws-sdk-go-v2/service/athena v1.49.0
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.37.7
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.51.1
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.24.7
	github.com/aws/aws-sdk-go-v2/service/b2bi v1.0.0-preview.55
	github.com/aws/aws-sdk-go-v2/service/backup v1.39.8
	github.com/aws/aws-sdk-go-v2/service/backupgateway v1.20.7
	github.com/aws/aws-sdk-go-v2/service/batch v1.48.2
	github.com/aws/aws-sdk-go-v2/service/bcmdataexports v1.7.7
	github.com/aws/aws-sdk-go-v2/service/bcmpricingcalculator v1.0.1
	github.com/aws/aws-sdk-go-v2/service/bedrock v1.25.1
	github.com/aws/aws-sdk-go-v2/service/bedrockagent v1.32.0
	github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime v1.30.0
	github.com/aws/aws-sdk-go-v2/service/bedrockdataautomation v1.0.0
	github.com/aws/aws-sdk-go-v2/service/bedrockdataautomationruntime v1.0.0
	github.com/aws/aws-sdk-go-v2/service/bedrockruntime v1.23.0
	github.com/aws/aws-sdk-go-v2/service/billing v1.0.2
	github.com/aws/aws-sdk-go-v2/service/billingconductor v1.20.7
	github.com/aws/aws-sdk-go-v2/service/braket v1.31.8
	github.com/aws/aws-sdk-go-v2/service/budgets v1.28.8
	github.com/aws/aws-sdk-go-v2/service/chatbot v1.9.1
	github.com/aws/aws-sdk-go-v2/service/chime v1.34.7
	github.com/aws/aws-sdk-go-v2/service/chimesdkidentity v1.22.7
	github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines v1.21.2
	github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings v1.27.7
	github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging v1.26.7
	github.com/aws/aws-sdk-go-v2/service/chimesdkvoice v1.20.0
	github.com/aws/aws-sdk-go-v2/service/cleanrooms v1.21.0
	github.com/aws/aws-sdk-go-v2/service/cleanroomsml v1.10.2
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.28.7
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.23.2
	github.com/aws/aws-sdk-go-v2/service/clouddirectory v1.24.7
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.56.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.43.1
	github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore v1.8.7
	github.com/aws/aws-sdk-go-v2/service/cloudhsm v1.24.7
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.28.0
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.26.6
	github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain v1.23.7
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.46.3
	github.com/aws/aws-sdk-go-v2/service/cloudtraildata v1.11.7
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.43.3
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.27.7
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.45.0
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.33.7
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.49.3
	github.com/aws/aws-sdk-go-v2/service/codecatalyst v1.17.7
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.27.7
	github.com/aws/aws-sdk-go-v2/service/codeconnections v1.5.7
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.29.7
	github.com/aws/aws-sdk-go-v2/service/codeguruprofiler v1.24.7
	github.com/aws/aws-sdk-go-v2/service/codegurureviewer v1.29.7
	github.com/aws/aws-sdk-go-v2/service/codegurusecurity v1.12.7
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.37.1
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.29.7
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.26.7
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.27.8
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.48.1
	github.com/aws/aws-sdk-go-v2/service/cognitosync v1.23.7
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.35.7
	github.com/aws/aws-sdk-go-v2/service/comprehendmedical v1.26.7
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.40.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.51.1
	github.com/aws/aws-sdk-go-v2/service/connect v1.121.0
	github.com/aws/aws-sdk-go-v2/service/connectcampaigns v1.15.8
	github.com/aws/aws-sdk-go-v2/service/connectcampaignsv2 v1.1.0
	github.com/aws/aws-sdk-go-v2/service/connectcases v1.21.7
	github.com/aws/aws-sdk-go-v2/service/connectcontactlens v1.26.7
	github.com/aws/aws-sdk-go-v2/service/connectparticipant v1.27.7
	github.com/aws/aws-sdk-go-v2/service/controlcatalog v1.6.3
	github.com/aws/aws-sdk-go-v2/service/controltower v1.20.1
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v1.28.7
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.45.1
	github.com/aws/aws-sdk-go-v2/service/costoptimizationhub v1.11.1
	github.com/aws/aws-sdk-go-v2/service/customerprofiles v1.44.0
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.45.0
	github.com/aws/aws-sdk-go-v2/service/databrew v1.33.7
	github.com/aws/aws-sdk-go-v2/service/dataexchange v1.33.5
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.25.7
	github.com/aws/aws-sdk-go-v2/service/datasync v1.43.5
	github.com/aws/aws-sdk-go-v2/service/datazone v1.25.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.23.7
	github.com/aws/aws-sdk-go-v2/service/deadline v1.7.2
	github.com/aws/aws-sdk-go-v2/service/detective v1.31.7
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.28.7
	github.com/aws/aws-sdk-go-v2/service/devopsguru v1.34.7
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.30.1
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.30.8
	github.com/aws/aws-sdk-go-v2/service/directoryservicedata v1.2.7
	github.com/aws/aws-sdk-go-v2/service/dlm v1.28.10
	github.com/aws/aws-sdk-go-v2/service/docdb v1.39.6
	github.com/aws/aws-sdk-go-v2/service/docdbelastic v1.14.4
	github.com/aws/aws-sdk-go-v2/service/drs v1.30.7
	github.com/aws/aws-sdk-go-v2/service/dsql v1.0.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.38.0
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.24.9
	github.com/aws/aws-sdk-go-v2/service/ebs v1.27.7
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.197.0
	github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect v1.27.7
	github.com/aws/aws-sdk-go-v2/service/ecr v1.36.7
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.27.7
	github.com/aws/aws-sdk-go-v2/service/ecs v1.52.2
	github.com/aws/aws-sdk-go-v2/service/efs v1.34.1
	github.com/aws/aws-sdk-go-v2/service/eks v1.54.0
	github.com/aws/aws-sdk-go-v2/service/eksauth v1.7.7
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.44.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.28.7
	github.com/aws/aws-sdk-go-v2/service/elasticinference v1.23.8
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.28.6
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.43.1
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.32.7
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.27.7
	github.com/aws/aws-sdk-go-v2/service/emr v1.47.1
	github.com/aws/aws-sdk-go-v2/service/emrcontainers v1.33.8
	github.com/aws/aws-sdk-go-v2/service/emrserverless v1.27.0
	github.com/aws/aws-sdk-go-v2/service/entityresolution v1.15.7
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.36.0
	github.com/aws/aws-sdk-go-v2/service/evidently v1.23.7
	github.com/aws/aws-sdk-go-v2/service/finspace v1.28.7
	github.com/aws/aws-sdk-go-v2/service/finspacedata v1.28.7
	github.com/aws/aws-sdk-go-v2/service/firehose v1.35.2
	github.com/aws/aws-sdk-go-v2/service/fis v1.31.2
	github.com/aws/aws-sdk-go-v2/service/fms v1.38.5
	github.com/aws/aws-sdk-go-v2/service/forecast v1.36.7
	github.com/aws/aws-sdk-go-v2/service/forecastquery v1.24.7
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.35.7
	github.com/aws/aws-sdk-go-v2/service/freetier v1.7.7
	github.com/aws/aws-sdk-go-v2/service/fsx v1.51.0
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.37.2
	github.com/aws/aws-sdk-go-v2/service/geomaps v1.0.4
	github.com/aws/aws-sdk-go-v2/service/geoplaces v1.0.4
	github.com/aws/aws-sdk-go-v2/service/georoutes v1.0.4
	github.com/aws/aws-sdk-go-v2/service/glacier v1.26.7
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.29.7
	github.com/aws/aws-sdk-go-v2/service/glue v1.104.0
	github.com/aws/aws-sdk-go-v2/service/grafana v1.26.7
	github.com/aws/aws-sdk-go-v2/service/greengrass v1.27.7
	github.com/aws/aws-sdk-go-v2/service/greengrassv2 v1.35.7
	github.com/aws/aws-sdk-go-v2/service/groundstation v1.31.7
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.52.1
	github.com/aws/aws-sdk-go-v2/service/health v1.29.1
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.28.7
	github.com/aws/aws-sdk-go-v2/service/iam v1.38.2
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.27.7
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.39.0
	github.com/aws/aws-sdk-go-v2/service/inspector v1.25.7
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.34.1
	github.com/aws/aws-sdk-go-v2/service/inspectorscan v1.7.7
	github.com/aws/aws-sdk-go-v2/service/internetmonitor v1.20.2
	github.com/aws/aws-sdk-go-v2/service/invoicing v1.0.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.61.1
	github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice v1.23.7
	github.com/aws/aws-sdk-go-v2/service/iot1clickprojects v1.23.7
	github.com/aws/aws-sdk-go-v2/service/iotanalytics v1.26.7
	github.com/aws/aws-sdk-go-v2/service/iotdataplane v1.26.7
	github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor v1.31.7
	github.com/aws/aws-sdk-go-v2/service/iotevents v1.27.7
	github.com/aws/aws-sdk-go-v2/service/ioteventsdata v1.24.7
	github.com/aws/aws-sdk-go-v2/service/iotfleethub v1.24.7
	github.com/aws/aws-sdk-go-v2/service/iotfleetwise v1.22.1
	github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane v1.24.1
	github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling v1.27.7
	github.com/aws/aws-sdk-go-v2/service/iotsitewise v1.44.1
	github.com/aws/aws-sdk-go-v2/service/iotthingsgraph v1.25.7
	github.com/aws/aws-sdk-go-v2/service/iottwinmaker v1.24.7
	github.com/aws/aws-sdk-go-v2/service/iotwireless v1.45.2
	github.com/aws/aws-sdk-go-v2/service/ivs v1.42.2
	github.com/aws/aws-sdk-go-v2/service/ivschat v1.16.7
	github.com/aws/aws-sdk-go-v2/service/ivsrealtime v1.21.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.38.7
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.21.7
	github.com/aws/aws-sdk-go-v2/service/kendra v1.55.0
	github.com/aws/aws-sdk-go-v2/service/kendraranking v1.11.7
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.16.2
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.32.7
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.25.8
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.31.8
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.27.7
	github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia v1.27.7
	github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia v1.22.7
	github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling v1.23.7
	github.com/aws/aws-sdk-go-v2/service/kinesisvideowebrtcstorage v1.14.7
	github.com/aws/aws-sdk-go-v2/service/kms v1.37.7
	github.com/aws/aws-sdk-go-v2/service/lakeformation v1.39.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.69.1
	github.com/aws/aws-sdk-go-v2/service/launchwizard v1.8.7
	github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice v1.28.7
	github.com/aws/aws-sdk-go-v2/service/lexmodelsv2 v1.49.7
	github.com/aws/aws-sdk-go-v2/service/lexruntimeservice v1.24.7
	github.com/aws/aws-sdk-go-v2/service/lexruntimev2 v1.29.7
	github.com/aws/aws-sdk-go-v2/service/licensemanager v1.29.7
	github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions v1.14.8
	github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions v1.14.2
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.42.7
	github.com/aws/aws-sdk-go-v2/service/location v1.42.7
	github.com/aws/aws-sdk-go-v2/service/lookoutequipment v1.30.7
	github.com/aws/aws-sdk-go-v2/service/lookoutmetrics v1.31.7
	github.com/aws/aws-sdk-go-v2/service/lookoutvision v1.27.7
	github.com/aws/aws-sdk-go-v2/service/m2 v1.18.5
	github.com/aws/aws-sdk-go-v2/service/machinelearning v1.28.7
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.43.7
	github.com/aws/aws-sdk-go-v2/service/mailmanager v1.7.1
	github.com/aws/aws-sdk-go-v2/service/managedblockchain v1.26.7
	github.com/aws/aws-sdk-go-v2/service/managedblockchainquery v1.16.7
	github.com/aws/aws-sdk-go-v2/service/marketplaceagreement v1.6.8
	github.com/aws/aws-sdk-go-v2/service/marketplacecatalog v1.30.7
	github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics v1.24.7
	github.com/aws/aws-sdk-go-v2/service/marketplacedeployment v1.6.7
	github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice v1.25.7
	github.com/aws/aws-sdk-go-v2/service/marketplacemetering v1.25.7
	github.com/aws/aws-sdk-go-v2/service/marketplacereporting v1.1.7
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.36.0
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.63.1
	github.com/aws/aws-sdk-go-v2/service/medialive v1.63.0
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.34.7
	github.com/aws/aws-sdk-go-v2/service/mediapackagev2 v1.20.1
	github.com/aws/aws-sdk-go-v2/service/mediapackagevod v1.34.8
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.24.7
	github.com/aws/aws-sdk-go-v2/service/mediastoredata v1.24.7
	github.com/aws/aws-sdk-go-v2/service/mediatailor v1.42.7
	github.com/aws/aws-sdk-go-v2/service/medicalimaging v1.14.7
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.25.0
	github.com/aws/aws-sdk-go-v2/service/mgn v1.32.7
	github.com/aws/aws-sdk-go-v2/service/migrationhub v1.25.0
	github.com/aws/aws-sdk-go-v2/service/migrationhubconfig v1.25.7
	github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator v1.13.7
	github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces v1.20.7
	github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy v1.21.7
	github.com/aws/aws-sdk-go-v2/service/mq v1.27.8
	github.com/aws/aws-sdk-go-v2/service/mturk v1.25.7
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.33.1
	github.com/aws/aws-sdk-go-v2/service/neptune v1.35.6
	github.com/aws/aws-sdk-go-v2/service/neptunedata v1.9.6
	github.com/aws/aws-sdk-go-v2/service/neptunegraph v1.15.1
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.44.5
	github.com/aws/aws-sdk-go-v2/service/networkflowmonitor v1.0.0
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.32.2
	github.com/aws/aws-sdk-go-v2/service/networkmonitor v1.7.7
	github.com/aws/aws-sdk-go-v2/service/notifications v1.0.1
	github.com/aws/aws-sdk-go-v2/service/notificationscontacts v1.0.1
	github.com/aws/aws-sdk-go-v2/service/oam v1.15.7
	github.com/aws/aws-sdk-go-v2/service/observabilityadmin v1.0.1
	github.com/aws/aws-sdk-go-v2/service/omics v1.28.1
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.45.0
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.17.4
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.26.7
	github.com/aws/aws-sdk-go-v2/service/opsworkscm v1.27.7
	github.com/aws/aws-sdk-go-v2/service/organizations v1.36.1
	github.com/aws/aws-sdk-go-v2/service/osis v1.14.7
	github.com/aws/aws-sdk-go-v2/service/outposts v1.47.2
	github.com/aws/aws-sdk-go-v2/service/panorama v1.22.7
	github.com/aws/aws-sdk-go-v2/service/partnercentralselling v1.2.0
	github.com/aws/aws-sdk-go-v2/service/paymentcryptography v1.16.2
	github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata v1.16.5
	github.com/aws/aws-sdk-go-v2/service/pcaconnectorad v1.9.7
	github.com/aws/aws-sdk-go-v2/service/pcaconnectorscep v1.4.7
	github.com/aws/aws-sdk-go-v2/service/pcs v1.2.8
	github.com/aws/aws-sdk-go-v2/service/personalize v1.39.7
	github.com/aws/aws-sdk-go-v2/service/personalizeevents v1.25.7
	github.com/aws/aws-sdk-go-v2/service/personalizeruntime v1.27.7
	github.com/aws/aws-sdk-go-v2/service/pi v1.29.7
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.34.7
	github.com/aws/aws-sdk-go-v2/service/pinpointemail v1.23.7
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice v1.23.7
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2 v1.18.2
	github.com/aws/aws-sdk-go-v2/service/pipes v1.18.5
	github.com/aws/aws-sdk-go-v2/service/polly v1.45.8
	github.com/aws/aws-sdk-go-v2/service/pricing v1.32.7
	github.com/aws/aws-sdk-go-v2/service/privatenetworks v1.13.7
	github.com/aws/aws-sdk-go-v2/service/proton v1.33.7
	github.com/aws/aws-sdk-go-v2/service/qapps v1.6.0
	github.com/aws/aws-sdk-go-v2/service/qbusiness v1.19.0
	github.com/aws/aws-sdk-go-v2/service/qconnect v1.14.0
	github.com/aws/aws-sdk-go-v2/service/qldb v1.25.7
	github.com/aws/aws-sdk-go-v2/service/qldbsession v1.25.7
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.81.0
	github.com/aws/aws-sdk-go-v2/service/ram v1.29.7
	github.com/aws/aws-sdk-go-v2/service/rbin v1.21.1
	github.com/aws/aws-sdk-go-v2/service/rds v1.92.0
	github.com/aws/aws-sdk-go-v2/service/rdsdata v1.26.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.53.0
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.31.4
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.25.0
	github.com/aws/aws-sdk-go-v2/service/rekognition v1.45.8
	github.com/aws/aws-sdk-go-v2/service/repostspace v1.8.5
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.28.1
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.16.2
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.27.7
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.25.7
	github.com/aws/aws-sdk-go-v2/service/robomaker v1.30.8
	github.com/aws/aws-sdk-go-v2/service/rolesanywhere v1.16.7
	github.com/aws/aws-sdk-go-v2/service/route53 v1.46.3
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.28.0
	github.com/aws/aws-sdk-go-v2/service/route53profiles v1.4.7
	github.com/aws/aws-sdk-go-v2/service/route53recoverycluster v1.23.7
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.25.7
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.21.7
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.34.2
	github.com/aws/aws-sdk-go-v2/service/rum v1.21.7
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
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.23.7
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.12.8
	github.com/aws/aws-sdk-go-v2/service/schemas v1.28.8
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.34.7
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.55.0
	github.com/aws/aws-sdk-go-v2/service/securityir v1.0.0
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.19.5
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.24.7
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.32.7
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.30.7
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.34.0
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.25.7
	github.com/aws/aws-sdk-go-v2/service/ses v1.29.1
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.40.0
	github.com/aws/aws-sdk-go-v2/service/sfn v1.34.1
	github.com/aws/aws-sdk-go-v2/service/shield v1.29.7
	github.com/aws/aws-sdk-go-v2/service/signer v1.26.7
	github.com/aws/aws-sdk-go-v2/service/simspaceweaver v1.14.7
	github.com/aws/aws-sdk-go-v2/service/sms v1.24.7
	github.com/aws/aws-sdk-go-v2/service/snowball v1.30.7
	github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement v1.20.7
	github.com/aws/aws-sdk-go-v2/service/sns v1.33.7
	github.com/aws/aws-sdk-go-v2/service/socialmessaging v1.1.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.37.2
	github.com/aws/aws-sdk-go-v2/service/ssm v1.56.1
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.26.7
	github.com/aws/aws-sdk-go-v2/service/ssmincidents v1.34.7
	github.com/aws/aws-sdk-go-v2/service/ssmquicksetup v1.3.1
	github.com/aws/aws-sdk-go-v2/service/ssmsap v1.18.7
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.7
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.29.7
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.6
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.34.7
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.2
	github.com/aws/aws-sdk-go-v2/service/supplychain v1.10.5
	github.com/aws/aws-sdk-go-v2/service/support v1.26.7
	github.com/aws/aws-sdk-go-v2/service/supportapp v1.13.7
	github.com/aws/aws-sdk-go-v2/service/swf v1.27.8
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.30.2
	github.com/aws/aws-sdk-go-v2/service/taxsettings v1.7.1
	github.com/aws/aws-sdk-go-v2/service/textract v1.34.9
	github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb v1.7.0
	github.com/aws/aws-sdk-go-v2/service/timestreamquery v1.29.1
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.29.8
	github.com/aws/aws-sdk-go-v2/service/tnb v1.13.7
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.41.7
	github.com/aws/aws-sdk-go-v2/service/transcribestreaming v1.22.5
	github.com/aws/aws-sdk-go-v2/service/transfer v1.54.0
	github.com/aws/aws-sdk-go-v2/service/translate v1.28.7
	github.com/aws/aws-sdk-go-v2/service/trustedadvisor v1.8.8
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.20.3
	github.com/aws/aws-sdk-go-v2/service/voiceid v1.24.7
	github.com/aws/aws-sdk-go-v2/service/vpclattice v1.13.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.25.7
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.25.7
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.55.6
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.34.7
	github.com/aws/aws-sdk-go-v2/service/wisdom v1.27.8
	github.com/aws/aws-sdk-go-v2/service/workdocs v1.25.7
	github.com/aws/aws-sdk-go-v2/service/workmail v1.30.4
	github.com/aws/aws-sdk-go-v2/service/workmailmessageflow v1.23.7
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.50.3
	github.com/aws/aws-sdk-go-v2/service/workspacesthinclient v1.10.7
	github.com/aws/aws-sdk-go-v2/service/workspacesweb v1.25.1
	github.com/aws/aws-sdk-go-v2/service/xray v1.30.1
)

require (
	github.com/DataDog/appsec-internal-go v1.6.0 // indirect
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.50.2 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.50.2 // indirect
	github.com/DataDog/datadog-go/v5 v5.5.0 // indirect
	github.com/DataDog/go-libddwaf/v3 v3.2.1 // indirect
	github.com/DataDog/go-sqllexer v0.0.10 // indirect
	github.com/DataDog/go-tuf v1.0.2-0.5.2 // indirect
	github.com/DataDog/sketches-go v1.4.5 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/aws/aws-sdk-go v1.50.9 // indirect
	github.com/aws/aws-sdk-go-v2 v1.32.6 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.7 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.28.6 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.47 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.4.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.6 // indirect
	github.com/aws/aws-xray-sdk-go v1.8.3 // indirect
	github.com/aws/smithy-go v1.22.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/ebitengine/purego v0.6.0-alpha.5 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.17.5 // indirect
	github.com/outcaste-io/ristretto v0.2.3 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.8.0 // indirect
	github.com/sony/gobreaker v0.5.0 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe // indirect
	google.golang.org/grpc v1.61.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.65.1 // indirect
)
