// Package services defines AWS console services supported by Granted.
package services

// Service is an AWS console destination exposed by Granted's -s flag.
type Service struct {
	UID      string
	Title    string
	Arg      string
	Icon     string
	Keywords string
}

// All returns every canonical Granted console service with aliases for fuzzy matching.
// Service args and aliases are sourced from fwdcloudsec/granted pkg/console/service_map.go.
func All() []Service {
	return []Service{
		{UID: "dashboard", Title: "Dashboard", Arg: "", Icon: "services/ec2.svg", Keywords: "dashboard aws console home"},
		{UID: "acm", Title: "ACM - Certificate Manager", Arg: "acm", Icon: "services/acm.svg", Keywords: "acm certificate manager ssl tls cert"},
		{UID: "aos", Title: "OpenSearch", Arg: "aos", Icon: "services/aos.svg", Keywords: "aos opensearch elasticsearch search"},
		{UID: "apigateway", Title: "API Gateway", Arg: "apigateway", Icon: "services/apigateway.svg", Keywords: "apigateway apigw api gateway rest graphql"},
		{UID: "appsync", Title: "AppSync", Arg: "appsync", Icon: "services/appsync.svg", Keywords: "appsync graphql api realtime"},
		{UID: "athena", Title: "Athena", Arg: "athena", Icon: "services/athena.svg", Keywords: "athena sql query analytics s3"},
		{UID: "backup", Title: "AWS Backup", Arg: "backup", Icon: "services/backup.svg", Keywords: "backup aws backup recovery restore"},
		{UID: "bedrock", Title: "Bedrock", Arg: "bedrock", Icon: "services/bedrock.svg", Keywords: "bedrock ai ml foundation model"},
		{UID: "billing", Title: "Billing", Arg: "billing", Icon: "services/billing.svg", Keywords: "billing cost invoice payment"},
		{UID: "cloud9", Title: "Cloud9", Arg: "c9", Icon: "services/cloud9.svg", Keywords: "cloud9 c9 ide editor development"},
		{UID: "cloudformation", Title: "CloudFormation", Arg: "cloudformation", Icon: "services/cloudformation.svg", Keywords: "cloudformation cfn cf infrastructure iac template"},
		{UID: "cloudfront", Title: "CloudFront", Arg: "cf", Icon: "services/cloudfront.svg", Keywords: "cloudfront cf cdn distribution"},
		{UID: "cloudmap", Title: "Cloud Map", Arg: "cloudmap", Icon: "services/cloudmap.svg", Keywords: "cloudmap service discovery dns"},
		{UID: "cloudtrail", Title: "CloudTrail", Arg: "ct", Icon: "services/cloudtrail.svg", Keywords: "cloudtrail ct audit logging trail"},
		{UID: "cloudwatch", Title: "CloudWatch", Arg: "cloudwatch", Icon: "services/cloudwatch.svg", Keywords: "cloudwatch cw monitoring metrics logs"},
		{UID: "codeartifact", Title: "CodeArtifact", Arg: "codeartifact", Icon: "services/codeartifact.svg", Keywords: "codeartifact artifact package npm maven"},
		{UID: "codecommit", Title: "CodeCommit", Arg: "codecommit", Icon: "services/codecommit.svg", Keywords: "codecommit git repository source"},
		{UID: "codedeploy", Title: "CodeDeploy", Arg: "codedeploy", Icon: "services/codedeploy.svg", Keywords: "codedeploy deployment blue green"},
		{UID: "codepipeline", Title: "CodePipeline", Arg: "codepipeline", Icon: "services/codepipeline.svg", Keywords: "codepipeline ci cd pipeline"},
		{UID: "codesuite", Title: "CodeSuite", Arg: "codesuite", Icon: "services/codesuite.svg", Keywords: "codesuite devops ci cd"},
		{UID: "cognito", Title: "Cognito", Arg: "cognito", Icon: "services/cognito.svg", Keywords: "cognito authentication auth user pool"},
		{UID: "config", Title: "AWS Config", Arg: "config", Icon: "services/config.svg", Keywords: "config aws config compliance audit"},
		{UID: "controltower", Title: "Control Tower", Arg: "controltower", Icon: "services/controltower.svg", Keywords: "controltower governance multi account"},
		{UID: "costmanagement", Title: "Cost Management", Arg: "ce", Icon: "services/costmanagement.svg", Keywords: "costmanagement ce cost explorer billing"},
		{UID: "dms", Title: "Database Migration Service", Arg: "dms", Icon: "services/dms.svg", Keywords: "dms database migration service"},
		{UID: "directconnect", Title: "Direct Connect", Arg: "dx", Icon: "services/directconnect.svg", Keywords: "directconnect dx network vpn"},
		{UID: "dynamodb", Title: "DynamoDB", Arg: "ddb", Icon: "services/dynamodb.svg", Keywords: "dynamodb ddb nosql database"},
		{UID: "ec2", Title: "EC2", Arg: "ec2", Icon: "services/ec2.svg", Keywords: "ec2 compute instance virtual machine"},
		{UID: "ecr", Title: "ECR", Arg: "ecr", Icon: "services/ecr.svg", Keywords: "ecr container registry docker image"},
		{UID: "ecs", Title: "ECS", Arg: "ecs", Icon: "services/ecs.svg", Keywords: "ecs container service fargate"},
		{UID: "eks", Title: "EKS", Arg: "eks", Icon: "services/eks.svg", Keywords: "eks kubernetes k8s container"},
		{UID: "elasticache", Title: "ElastiCache", Arg: "elasticache", Icon: "services/elasticache.svg", Keywords: "elasticache redis memcached cache"},
		{UID: "elasticbeanstalk", Title: "Elastic Beanstalk", Arg: "eb", Icon: "services/elasticbeanstalk.svg", Keywords: "elasticbeanstalk eb ebs deployment platform"},
		{UID: "eventbridge", Title: "EventBridge", Arg: "eventbridge", Icon: "services/eventbridge.svg", Keywords: "eventbridge events event bus"},
		{UID: "globalaccelerator", Title: "Global Accelerator", Arg: "globalaccelerator", Icon: "services/globalaccelerator.svg", Keywords: "globalaccelerator ga network performance"},
		{UID: "grafana", Title: "Grafana", Arg: "grafana", Icon: "services/grafana.svg", Keywords: "grafana monitoring observability dashboard"},
		{UID: "guardduty", Title: "GuardDuty", Arg: "gd", Icon: "services/guardduty.svg", Keywords: "guardduty gd threat detection security"},
		{UID: "iam", Title: "IAM", Arg: "iam", Icon: "services/iam.svg", Keywords: "iam identity access management user role"},
		{UID: "kms", Title: "KMS", Arg: "kms", Icon: "services/kms.svg", Keywords: "kms key management encryption"},
		{UID: "lambda", Title: "Lambda", Arg: "lambda", Icon: "services/lambda.svg", Keywords: "lambda l serverless function compute"},
		{UID: "mwaa", Title: "MWAA - Managed Workflows", Arg: "mwaa", Icon: "services/mwaa.svg", Keywords: "mwaa airflow workflow orchestration"},
		{UID: "organizations", Title: "Organizations", Arg: "organizations", Icon: "services/organizations.svg", Keywords: "organizations orgs org multi account"},
		{UID: "paramstore", Title: "Parameter Store - SSM", Arg: "param", Icon: "services/systemsmanager.svg", Keywords: "paramstore parameter store ssm config param"},
		{UID: "ram", Title: "RAM - Resource Access Manager", Arg: "ram", Icon: "services/ram.svg", Keywords: "ram resource access manager sharing"},
		{UID: "rds", Title: "RDS", Arg: "rds", Icon: "services/rds.svg", Keywords: "rds database mysql postgresql"},
		{UID: "redshift", Title: "Redshift", Arg: "redshift", Icon: "services/redshift.svg", Keywords: "redshift data warehouse analytics"},
		{UID: "route53", Title: "Route53", Arg: "route53", Icon: "services/route53.svg", Keywords: "route53 r53 dns domain name"},
		{UID: "s3", Title: "S3", Arg: "s3", Icon: "services/s3.svg", Keywords: "s3 storage bucket object"},
		{UID: "sagemaker", Title: "SageMaker", Arg: "sagemaker", Icon: "services/sagemaker.svg", Keywords: "sagemaker ml machine learning model"},
		{UID: "secretsmanager", Title: "Secrets Manager", Arg: "secretsmanager", Icon: "services/secretsmanager.svg", Keywords: "secretsmanager scrm sm secret password credential"},
		{UID: "securityhub", Title: "Security Hub", Arg: "securityhub", Icon: "services/securityhub.svg", Keywords: "securityhub scrh security findings compliance"},
		{UID: "ses", Title: "SES", Arg: "ses", Icon: "services/ses.svg", Keywords: "ses email smtp mail"},
		{UID: "singlesignon", Title: "SSO - Single Sign On", Arg: "sso", Icon: "services/singlesignon.svg", Keywords: "singlesignon sso identity federation"},
		{UID: "sns", Title: "SNS", Arg: "sns", Icon: "services/sns.svg", Keywords: "sns notification pub sub messaging"},
		{UID: "sqs", Title: "SQS", Arg: "sqs", Icon: "services/sqs.svg", Keywords: "sqs queue messaging decoupling"},
		{UID: "ssm", Title: "Systems Manager - SSM", Arg: "ssm", Icon: "services/ssm.svg", Keywords: "ssm systems manager parameter patch"},
		{UID: "states", Title: "Step Functions", Arg: "states", Icon: "services/states.svg", Keywords: "states sfn stepfn step functions workflow orchestration"},
		{UID: "trustedadvisor", Title: "Trusted Advisor", Arg: "trustedadvisor", Icon: "services/trustedadvisor.svg", Keywords: "trustedadvisor tra advisor recommendation best practice"},
		{UID: "vpc", Title: "VPC", Arg: "vpc", Icon: "services/vpc.svg", Keywords: "vpc virtual private cloud network"},
		{UID: "waf", Title: "AWS WAF", Arg: "wafv2", Icon: "services/waf.svg", Keywords: "waf wafv2 web application firewall security"},
	}
}
