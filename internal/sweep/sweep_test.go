//go:build sweep
// +build sweep

package sweep_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/accessanalyzer"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/acm"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/acmpca"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/amplify"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/apigateway"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/apigatewayv2"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/appconfig"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/appmesh"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/apprunner"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/appstream"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/appsync"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/autoscaling"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/autoscalingplans"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/backup"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/batch"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/budgets"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cloud9"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cloudformation"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cloudfront"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cloudhsmv2"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cloudsearch"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cloudtrail"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cloudwatch"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cloudwatchlogs"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/codeartifact"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/codebuild"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/codedeploy"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/codepipeline"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cognitoidp"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/configservice"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/connect"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/cur"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/datasync"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/dax"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/directconnect"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/dms"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/ds"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/dynamodb"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/ecr"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/ecrpublic"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/ecs"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/efs"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/eks"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/elasticache"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/elasticbeanstalk"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/elasticsearch"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/elb"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/elbv2"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/emr"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/events"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/firehose"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/fsx"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/gamelift"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/glacier"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/globalaccelerator"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/glue"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/guardduty"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/iam"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/imagebuilder"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/iot"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/kafka"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/kafkaconnect"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/kinesis"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/kinesisanalytics"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/kinesisanalyticsv2"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/kms"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/lambda"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/lexmodels"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/licensemanager"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/lightsail"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/memorydb"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/mq"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/mwaa"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/neptune"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/networkfirewall"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/opsworks"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/pinpoint"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/qldb"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/quicksight"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/rds"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/redshift"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/route53"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/route53recoverycontrolconfig"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/route53resolver"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/s3"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/s3control"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/sagemaker"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/schemas"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/secretsmanager"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/servicecatalog"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/servicediscovery"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/ses"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/sns"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/sqs"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/ssm"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/ssoadmin"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/storagegateway"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/synthetics"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/timestreamwrite"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/transfer"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/waf"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/wafregional"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/wafv2"
	_ "github.com/hashicorp/terraform-provider-aws/internal/service/workspaces"
	"github.com/hashicorp/terraform-provider-aws/internal/sweep"
)

func TestMain(m *testing.M) {
	sweep.SweeperClients = make(map[string]interface{})
	resource.TestMain(m)
}
