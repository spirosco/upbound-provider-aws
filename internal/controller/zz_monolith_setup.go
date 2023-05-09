/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	analyzer "github.com/spirosco/upbound-provider-aws/internal/controller/accessanalyzer/analyzer"
	archiverule "github.com/spirosco/upbound-provider-aws/internal/controller/accessanalyzer/archiverule"
	alternatecontact "github.com/spirosco/upbound-provider-aws/internal/controller/account/alternatecontact"
	certificate "github.com/spirosco/upbound-provider-aws/internal/controller/acm/certificate"
	certificatevalidation "github.com/spirosco/upbound-provider-aws/internal/controller/acm/certificatevalidation"
	certificateacmpca "github.com/spirosco/upbound-provider-aws/internal/controller/acmpca/certificate"
	certificateauthority "github.com/spirosco/upbound-provider-aws/internal/controller/acmpca/certificateauthority"
	certificateauthoritycertificate "github.com/spirosco/upbound-provider-aws/internal/controller/acmpca/certificateauthoritycertificate"
	permission "github.com/spirosco/upbound-provider-aws/internal/controller/acmpca/permission"
	policy "github.com/spirosco/upbound-provider-aws/internal/controller/acmpca/policy"
	alertmanagerdefinition "github.com/spirosco/upbound-provider-aws/internal/controller/amp/alertmanagerdefinition"
	rulegroupnamespace "github.com/spirosco/upbound-provider-aws/internal/controller/amp/rulegroupnamespace"
	workspace "github.com/spirosco/upbound-provider-aws/internal/controller/amp/workspace"
	app "github.com/spirosco/upbound-provider-aws/internal/controller/amplify/app"
	backendenvironment "github.com/spirosco/upbound-provider-aws/internal/controller/amplify/backendenvironment"
	branch "github.com/spirosco/upbound-provider-aws/internal/controller/amplify/branch"
	webhook "github.com/spirosco/upbound-provider-aws/internal/controller/amplify/webhook"
	account "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/account"
	apikey "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/apikey"
	authorizer "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/authorizer"
	basepathmapping "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/basepathmapping"
	clientcertificate "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/clientcertificate"
	deployment "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/deployment"
	documentationpart "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/documentationpart"
	documentationversion "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/documentationversion"
	domainname "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/domainname"
	gatewayresponse "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/gatewayresponse"
	integration "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/integration"
	integrationresponse "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/integrationresponse"
	method "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/method"
	methodresponse "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/methodresponse"
	methodsettings "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/methodsettings"
	model "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/model"
	requestvalidator "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/requestvalidator"
	resource "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/resource"
	restapi "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/restapi"
	restapipolicy "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/restapipolicy"
	stage "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/stage"
	usageplan "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/usageplan"
	usageplankey "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/usageplankey"
	vpclink "github.com/spirosco/upbound-provider-aws/internal/controller/apigateway/vpclink"
	api "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/api"
	apimapping "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/apimapping"
	authorizerapigatewayv2 "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/authorizer"
	deploymentapigatewayv2 "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/deployment"
	domainnameapigatewayv2 "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/domainname"
	integrationapigatewayv2 "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/integration"
	integrationresponseapigatewayv2 "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/integrationresponse"
	modelapigatewayv2 "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/model"
	route "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/route"
	routeresponse "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/routeresponse"
	stageapigatewayv2 "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/stage"
	vpclinkapigatewayv2 "github.com/spirosco/upbound-provider-aws/internal/controller/apigatewayv2/vpclink"
	policyappautoscaling "github.com/spirosco/upbound-provider-aws/internal/controller/appautoscaling/policy"
	scheduledaction "github.com/spirosco/upbound-provider-aws/internal/controller/appautoscaling/scheduledaction"
	target "github.com/spirosco/upbound-provider-aws/internal/controller/appautoscaling/target"
	application "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/application"
	configurationprofile "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/configurationprofile"
	deploymentappconfig "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/deployment"
	deploymentstrategy "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/deploymentstrategy"
	environment "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/environment"
	extension "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/extension"
	extensionassociation "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/extensionassociation"
	hostedconfigurationversion "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/hostedconfigurationversion"
	flow "github.com/spirosco/upbound-provider-aws/internal/controller/appflow/flow"
	eventintegration "github.com/spirosco/upbound-provider-aws/internal/controller/appintegrations/eventintegration"
	applicationapplicationinsights "github.com/spirosco/upbound-provider-aws/internal/controller/applicationinsights/application"
	gatewayroute "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/gatewayroute"
	mesh "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/mesh"
	routeappmesh "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/route"
	virtualgateway "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/virtualgateway"
	virtualnode "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/virtualnode"
	virtualrouter "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/virtualrouter"
	virtualservice "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/virtualservice"
	autoscalingconfigurationversion "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/autoscalingconfigurationversion"
	connection "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/connection"
	observabilityconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/observabilityconfiguration"
	service "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/service"
	vpcconnector "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/vpcconnector"
	directoryconfig "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/directoryconfig"
	fleet "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/fleet"
	fleetstackassociation "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/fleetstackassociation"
	imagebuilder "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/imagebuilder"
	stack "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/stack"
	user "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/user"
	userstackassociation "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/userstackassociation"
	apicache "github.com/spirosco/upbound-provider-aws/internal/controller/appsync/apicache"
	apikeyappsync "github.com/spirosco/upbound-provider-aws/internal/controller/appsync/apikey"
	datasource "github.com/spirosco/upbound-provider-aws/internal/controller/appsync/datasource"
	function "github.com/spirosco/upbound-provider-aws/internal/controller/appsync/function"
	graphqlapi "github.com/spirosco/upbound-provider-aws/internal/controller/appsync/graphqlapi"
	resolver "github.com/spirosco/upbound-provider-aws/internal/controller/appsync/resolver"
	database "github.com/spirosco/upbound-provider-aws/internal/controller/athena/database"
	datacatalog "github.com/spirosco/upbound-provider-aws/internal/controller/athena/datacatalog"
	namedquery "github.com/spirosco/upbound-provider-aws/internal/controller/athena/namedquery"
	workgroup "github.com/spirosco/upbound-provider-aws/internal/controller/athena/workgroup"
	attachment "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/autoscalinggroup"
	grouptag "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/grouptag"
	launchconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/launchconfiguration"
	lifecyclehook "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/lifecyclehook"
	notification "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/notification"
	policyautoscaling "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/policy"
	schedule "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/schedule"
	scalingplan "github.com/spirosco/upbound-provider-aws/internal/controller/autoscalingplans/scalingplan"
	framework "github.com/spirosco/upbound-provider-aws/internal/controller/backup/framework"
	globalsettings "github.com/spirosco/upbound-provider-aws/internal/controller/backup/globalsettings"
	plan "github.com/spirosco/upbound-provider-aws/internal/controller/backup/plan"
	regionsettings "github.com/spirosco/upbound-provider-aws/internal/controller/backup/regionsettings"
	reportplan "github.com/spirosco/upbound-provider-aws/internal/controller/backup/reportplan"
	selection "github.com/spirosco/upbound-provider-aws/internal/controller/backup/selection"
	vault "github.com/spirosco/upbound-provider-aws/internal/controller/backup/vault"
	vaultlockconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/backup/vaultlockconfiguration"
	vaultnotifications "github.com/spirosco/upbound-provider-aws/internal/controller/backup/vaultnotifications"
	vaultpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/backup/vaultpolicy"
	schedulingpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/batch/schedulingpolicy"
	budget "github.com/spirosco/upbound-provider-aws/internal/controller/budgets/budget"
	budgetaction "github.com/spirosco/upbound-provider-aws/internal/controller/budgets/budgetaction"
	anomalymonitor "github.com/spirosco/upbound-provider-aws/internal/controller/ce/anomalymonitor"
	voiceconnector "github.com/spirosco/upbound-provider-aws/internal/controller/chime/voiceconnector"
	voiceconnectorgroup "github.com/spirosco/upbound-provider-aws/internal/controller/chime/voiceconnectorgroup"
	voiceconnectorlogging "github.com/spirosco/upbound-provider-aws/internal/controller/chime/voiceconnectorlogging"
	voiceconnectororigination "github.com/spirosco/upbound-provider-aws/internal/controller/chime/voiceconnectororigination"
	voiceconnectorstreaming "github.com/spirosco/upbound-provider-aws/internal/controller/chime/voiceconnectorstreaming"
	voiceconnectortermination "github.com/spirosco/upbound-provider-aws/internal/controller/chime/voiceconnectortermination"
	voiceconnectorterminationcredentials "github.com/spirosco/upbound-provider-aws/internal/controller/chime/voiceconnectorterminationcredentials"
	environmentec2 "github.com/spirosco/upbound-provider-aws/internal/controller/cloud9/environmentec2"
	environmentmembership "github.com/spirosco/upbound-provider-aws/internal/controller/cloud9/environmentmembership"
	resourcecloudcontrol "github.com/spirosco/upbound-provider-aws/internal/controller/cloudcontrol/resource"
	stackcloudformation "github.com/spirosco/upbound-provider-aws/internal/controller/cloudformation/stack"
	stackset "github.com/spirosco/upbound-provider-aws/internal/controller/cloudformation/stackset"
	cachepolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/cachepolicy"
	distribution "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/distribution"
	fieldlevelencryptionconfig "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/fieldlevelencryptionconfig"
	fieldlevelencryptionprofile "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/fieldlevelencryptionprofile"
	functioncloudfront "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/function"
	keygroup "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/keygroup"
	monitoringsubscription "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/monitoringsubscription"
	originaccesscontrol "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/originaccesscontrol"
	originaccessidentity "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/originaccessidentity"
	originrequestpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/originrequestpolicy"
	publickey "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/publickey"
	realtimelogconfig "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/realtimelogconfig"
	responseheaderspolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudfront/responseheaderspolicy"
	domain "github.com/spirosco/upbound-provider-aws/internal/controller/cloudsearch/domain"
	domainserviceaccesspolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudsearch/domainserviceaccesspolicy"
	eventdatastore "github.com/spirosco/upbound-provider-aws/internal/controller/cloudtrail/eventdatastore"
	trail "github.com/spirosco/upbound-provider-aws/internal/controller/cloudtrail/trail"
	compositealarm "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatch/compositealarm"
	dashboard "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatch/dashboard"
	metricalarm "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatch/metricalarm"
	metricstream "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatch/metricstream"
	apidestination "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/apidestination"
	archive "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/archive"
	bus "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/bus"
	buspolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/buspolicy"
	connectioncloudwatchevents "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/connection"
	permissioncloudwatchevents "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/permission"
	rule "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/rule"
	targetcloudwatchevents "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/target"
	definition "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/definition"
	destination "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/destination"
	destinationpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/destinationpolicy"
	group "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/group"
	metricfilter "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/resourcepolicy"
	stream "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/stream"
	subscriptionfilter "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/subscriptionfilter"
	approvalruletemplate "github.com/spirosco/upbound-provider-aws/internal/controller/codecommit/approvalruletemplate"
	approvalruletemplateassociation "github.com/spirosco/upbound-provider-aws/internal/controller/codecommit/approvalruletemplateassociation"
	repository "github.com/spirosco/upbound-provider-aws/internal/controller/codecommit/repository"
	trigger "github.com/spirosco/upbound-provider-aws/internal/controller/codecommit/trigger"
	codepipeline "github.com/spirosco/upbound-provider-aws/internal/controller/codepipeline/codepipeline"
	customactiontype "github.com/spirosco/upbound-provider-aws/internal/controller/codepipeline/customactiontype"
	webhookcodepipeline "github.com/spirosco/upbound-provider-aws/internal/controller/codepipeline/webhook"
	connectioncodestarconnections "github.com/spirosco/upbound-provider-aws/internal/controller/codestarconnections/connection"
	host "github.com/spirosco/upbound-provider-aws/internal/controller/codestarconnections/host"
	notificationrule "github.com/spirosco/upbound-provider-aws/internal/controller/codestarnotifications/notificationrule"
	cognitoidentitypoolproviderprincipaltag "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidentity/cognitoidentitypoolproviderprincipaltag"
	pool "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidentity/pool"
	poolrolesattachment "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidentity/poolrolesattachment"
	identityprovider "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/identityprovider"
	resourceserver "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/resourceserver"
	riskconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/riskconfiguration"
	usercognitoidp "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/user"
	usergroup "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/usergroup"
	useringroup "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/useringroup"
	userpool "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/userpool"
	userpoolclient "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/userpoolclient"
	userpooldomain "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/userpooldomain"
	userpooluicustomization "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidp/userpooluicustomization"
	awsconfigurationrecorderstatus "github.com/spirosco/upbound-provider-aws/internal/controller/configservice/awsconfigurationrecorderstatus"
	configrule "github.com/spirosco/upbound-provider-aws/internal/controller/configservice/configrule"
	configurationaggregator "github.com/spirosco/upbound-provider-aws/internal/controller/configservice/configurationaggregator"
	configurationrecorder "github.com/spirosco/upbound-provider-aws/internal/controller/configservice/configurationrecorder"
	conformancepack "github.com/spirosco/upbound-provider-aws/internal/controller/configservice/conformancepack"
	deliverychannel "github.com/spirosco/upbound-provider-aws/internal/controller/configservice/deliverychannel"
	remediationconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/configservice/remediationconfiguration"
	botassociation "github.com/spirosco/upbound-provider-aws/internal/controller/connect/botassociation"
	contactflow "github.com/spirosco/upbound-provider-aws/internal/controller/connect/contactflow"
	contactflowmodule "github.com/spirosco/upbound-provider-aws/internal/controller/connect/contactflowmodule"
	hoursofoperation "github.com/spirosco/upbound-provider-aws/internal/controller/connect/hoursofoperation"
	instance "github.com/spirosco/upbound-provider-aws/internal/controller/connect/instance"
	instancestorageconfig "github.com/spirosco/upbound-provider-aws/internal/controller/connect/instancestorageconfig"
	lambdafunctionassociation "github.com/spirosco/upbound-provider-aws/internal/controller/connect/lambdafunctionassociation"
	phonenumber "github.com/spirosco/upbound-provider-aws/internal/controller/connect/phonenumber"
	queue "github.com/spirosco/upbound-provider-aws/internal/controller/connect/queue"
	quickconnect "github.com/spirosco/upbound-provider-aws/internal/controller/connect/quickconnect"
	routingprofile "github.com/spirosco/upbound-provider-aws/internal/controller/connect/routingprofile"
	securityprofile "github.com/spirosco/upbound-provider-aws/internal/controller/connect/securityprofile"
	userconnect "github.com/spirosco/upbound-provider-aws/internal/controller/connect/user"
	userhierarchystructure "github.com/spirosco/upbound-provider-aws/internal/controller/connect/userhierarchystructure"
	vocabulary "github.com/spirosco/upbound-provider-aws/internal/controller/connect/vocabulary"
	reportdefinition "github.com/spirosco/upbound-provider-aws/internal/controller/cur/reportdefinition"
	dataset "github.com/spirosco/upbound-provider-aws/internal/controller/dataexchange/dataset"
	revision "github.com/spirosco/upbound-provider-aws/internal/controller/dataexchange/revision"
	pipeline "github.com/spirosco/upbound-provider-aws/internal/controller/datapipeline/pipeline"
	cluster "github.com/spirosco/upbound-provider-aws/internal/controller/dax/cluster"
	parametergroup "github.com/spirosco/upbound-provider-aws/internal/controller/dax/parametergroup"
	subnetgroup "github.com/spirosco/upbound-provider-aws/internal/controller/dax/subnetgroup"
	appdeploy "github.com/spirosco/upbound-provider-aws/internal/controller/deploy/app"
	deploymentconfig "github.com/spirosco/upbound-provider-aws/internal/controller/deploy/deploymentconfig"
	deploymentgroup "github.com/spirosco/upbound-provider-aws/internal/controller/deploy/deploymentgroup"
	graph "github.com/spirosco/upbound-provider-aws/internal/controller/detective/graph"
	invitationaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/detective/invitationaccepter"
	member "github.com/spirosco/upbound-provider-aws/internal/controller/detective/member"
	devicepool "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/devicepool"
	instanceprofile "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/instanceprofile"
	networkprofile "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/networkprofile"
	project "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/project"
	testgridproject "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/testgridproject"
	upload "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/upload"
	bgppeer "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/bgppeer"
	connectiondirectconnect "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/connection"
	connectionassociation "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/connectionassociation"
	gateway "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/gateway"
	gatewayassociation "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/gatewayassociation"
	gatewayassociationproposal "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/lag"
	privatevirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/privatevirtualinterface"
	publicvirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/publicvirtualinterface"
	transitvirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/transitvirtualinterface"
	lifecyclepolicy "github.com/spirosco/upbound-provider-aws/internal/controller/dlm/lifecyclepolicy"
	certificatedms "github.com/spirosco/upbound-provider-aws/internal/controller/dms/certificate"
	endpoint "github.com/spirosco/upbound-provider-aws/internal/controller/dms/endpoint"
	eventsubscription "github.com/spirosco/upbound-provider-aws/internal/controller/dms/eventsubscription"
	replicationinstance "github.com/spirosco/upbound-provider-aws/internal/controller/dms/replicationinstance"
	replicationsubnetgroup "github.com/spirosco/upbound-provider-aws/internal/controller/dms/replicationsubnetgroup"
	replicationtask "github.com/spirosco/upbound-provider-aws/internal/controller/dms/replicationtask"
	s3endpoint "github.com/spirosco/upbound-provider-aws/internal/controller/dms/s3endpoint"
	clusterdocdb "github.com/spirosco/upbound-provider-aws/internal/controller/docdb/cluster"
	clusterinstance "github.com/spirosco/upbound-provider-aws/internal/controller/docdb/clusterinstance"
	clusterparametergroup "github.com/spirosco/upbound-provider-aws/internal/controller/docdb/clusterparametergroup"
	clustersnapshot "github.com/spirosco/upbound-provider-aws/internal/controller/docdb/clustersnapshot"
	eventsubscriptiondocdb "github.com/spirosco/upbound-provider-aws/internal/controller/docdb/eventsubscription"
	globalcluster "github.com/spirosco/upbound-provider-aws/internal/controller/docdb/globalcluster"
	subnetgroupdocdb "github.com/spirosco/upbound-provider-aws/internal/controller/docdb/subnetgroup"
	conditionalforwarder "github.com/spirosco/upbound-provider-aws/internal/controller/ds/conditionalforwarder"
	directory "github.com/spirosco/upbound-provider-aws/internal/controller/ds/directory"
	shareddirectory "github.com/spirosco/upbound-provider-aws/internal/controller/ds/shareddirectory"
	contributorinsights "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/contributorinsights"
	globaltable "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/table"
	tableitem "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/tableitem"
	tablereplica "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/tablereplica"
	tag "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/tag"
	ami "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/ami"
	amicopy "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/amicopy"
	amilaunchpermission "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/amilaunchpermission"
	availabilityzonegroup "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/availabilityzonegroup"
	capacityreservation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/capacityreservation"
	carriergateway "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/carriergateway"
	customergateway "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/customergateway"
	defaultnetworkacl "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/defaultnetworkacl"
	defaultroutetable "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/defaultroutetable"
	defaultsecuritygroup "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/defaultsecuritygroup"
	defaultsubnet "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/defaultsubnet"
	defaultvpc "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/defaultvpc"
	defaultvpcdhcpoptions "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/defaultvpcdhcpoptions"
	ebsdefaultkmskey "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/ebsdefaultkmskey"
	ebsencryptionbydefault "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/ebsencryptionbydefault"
	ebssnapshot "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/ebssnapshot"
	ebssnapshotcopy "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/ebssnapshotcopy"
	ebssnapshotimport "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/ebssnapshotimport"
	ebsvolume "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/ebsvolume"
	egressonlyinternetgateway "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/egressonlyinternetgateway"
	eip "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/eip"
	eipassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/eipassociation"
	flowlog "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/flowlog"
	hostec2 "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/host"
	instanceec2 "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/instance"
	instancestate "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/instancestate"
	internetgateway "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/internetgateway"
	keypair "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/keypair"
	launchtemplate "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/launchtemplate"
	mainroutetableassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/mainroutetableassociation"
	managedprefixlist "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/managedprefixlist"
	managedprefixlistentry "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/managedprefixlistentry"
	natgateway "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/natgateway"
	networkacl "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/networkacl"
	networkaclrule "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/networkaclrule"
	networkinsightsanalysis "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/networkinsightsanalysis"
	networkinsightspath "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/networkinsightspath"
	networkinterface "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/networkinterface"
	networkinterfaceattachment "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/networkinterfaceattachment"
	networkinterfacesgattachment "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/networkinterfacesgattachment"
	placementgroup "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/placementgroup"
	routeec2 "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/route"
	routetable "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/routetable"
	routetableassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/routetableassociation"
	securitygroup "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/securitygroup"
	securitygrouprule "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/securitygrouprule"
	serialconsoleaccess "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/serialconsoleaccess"
	snapshotcreatevolumepermission "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/snapshotcreatevolumepermission"
	spotdatafeedsubscription "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/spotdatafeedsubscription"
	spotfleetrequest "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/spotfleetrequest"
	spotinstancerequest "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/spotinstancerequest"
	subnet "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/subnet"
	subnetcidrreservation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/subnetcidrreservation"
	tagec2 "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/tag"
	trafficmirrorfilter "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/trafficmirrorfilterrule"
	transitgateway "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgateway"
	transitgatewayconnect "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewayconnect"
	transitgatewayconnectpeer "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewayconnectpeer"
	transitgatewaymulticastdomain "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewaymulticastdomain"
	transitgatewaymulticastdomainassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewaymulticastdomainassociation"
	transitgatewaymulticastgroupmember "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewaymulticastgroupmember"
	transitgatewaymulticastgroupsource "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewaymulticastgroupsource"
	transitgatewaypeeringattachment "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewaypeeringattachment"
	transitgatewaypeeringattachmentaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewaypeeringattachmentaccepter"
	transitgatewaypolicytable "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewaypolicytable"
	transitgatewayprefixlistreference "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewayprefixlistreference"
	transitgatewayroute "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewayroute"
	transitgatewayroutetable "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewayroutetable"
	transitgatewayroutetableassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewayroutetableassociation"
	transitgatewayroutetablepropagation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewayroutetablepropagation"
	transitgatewayvpcattachment "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewayvpcattachment"
	transitgatewayvpcattachmentaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/transitgatewayvpcattachmentaccepter"
	volumeattachment "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/volumeattachment"
	vpc "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpc"
	vpcdhcpoptions "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcdhcpoptions"
	vpcdhcpoptionsassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcdhcpoptionsassociation"
	vpcendpoint "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcendpoint"
	vpcendpointconnectionnotification "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcendpointconnectionnotification"
	vpcendpointroutetableassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcendpointroutetableassociation"
	vpcendpointsecuritygroupassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcendpointsecuritygroupassociation"
	vpcendpointservice "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcendpointservice"
	vpcendpointserviceallowedprincipal "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcendpointserviceallowedprincipal"
	vpcendpointsubnetassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcendpointsubnetassociation"
	vpcipam "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcipam"
	vpcipampool "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcipampool"
	vpcipampoolcidr "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcipampoolcidr"
	vpcipampoolcidrallocation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcipampoolcidrallocation"
	vpcipamscope "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcipamscope"
	vpcipv4cidrblockassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcipv4cidrblockassociation"
	vpcpeeringconnection "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcpeeringconnection"
	vpcpeeringconnectionaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcpeeringconnectionaccepter"
	vpcpeeringconnectionoptions "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpcpeeringconnectionoptions"
	vpnconnection "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpnconnection"
	vpnconnectionroute "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpnconnectionroute"
	vpngateway "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpngateway"
	vpngatewayattachment "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpngatewayattachment"
	vpngatewayroutepropagation "github.com/spirosco/upbound-provider-aws/internal/controller/ec2/vpngatewayroutepropagation"
	lifecyclepolicyecr "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/lifecyclepolicy"
	pullthroughcacherule "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/pullthroughcacherule"
	registrypolicy "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/registrypolicy"
	registryscanningconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/registryscanningconfiguration"
	replicationconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/replicationconfiguration"
	repositoryecr "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/repositorypolicy"
	repositoryecrpublic "github.com/spirosco/upbound-provider-aws/internal/controller/ecrpublic/repository"
	repositorypolicyecrpublic "github.com/spirosco/upbound-provider-aws/internal/controller/ecrpublic/repositorypolicy"
	accountsettingdefault "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/accountsettingdefault"
	capacityprovider "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/capacityprovider"
	clusterecs "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/cluster"
	clustercapacityproviders "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/clustercapacityproviders"
	serviceecs "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/service"
	taskdefinition "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/taskdefinition"
	accesspoint "github.com/spirosco/upbound-provider-aws/internal/controller/efs/accesspoint"
	backuppolicy "github.com/spirosco/upbound-provider-aws/internal/controller/efs/backuppolicy"
	filesystem "github.com/spirosco/upbound-provider-aws/internal/controller/efs/filesystem"
	filesystempolicy "github.com/spirosco/upbound-provider-aws/internal/controller/efs/filesystempolicy"
	mounttarget "github.com/spirosco/upbound-provider-aws/internal/controller/efs/mounttarget"
	replicationconfigurationefs "github.com/spirosco/upbound-provider-aws/internal/controller/efs/replicationconfiguration"
	addon "github.com/spirosco/upbound-provider-aws/internal/controller/eks/addon"
	clustereks "github.com/spirosco/upbound-provider-aws/internal/controller/eks/cluster"
	clusterauth "github.com/spirosco/upbound-provider-aws/internal/controller/eks/clusterauth"
	fargateprofile "github.com/spirosco/upbound-provider-aws/internal/controller/eks/fargateprofile"
	identityproviderconfig "github.com/spirosco/upbound-provider-aws/internal/controller/eks/identityproviderconfig"
	nodegroup "github.com/spirosco/upbound-provider-aws/internal/controller/eks/nodegroup"
	clusterelasticache "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/cluster"
	parametergroupelasticache "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/replicationgroup"
	subnetgroupelasticache "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/subnetgroup"
	userelasticache "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/user"
	usergroupelasticache "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/usergroup"
	applicationelasticbeanstalk "github.com/spirosco/upbound-provider-aws/internal/controller/elasticbeanstalk/application"
	applicationversion "github.com/spirosco/upbound-provider-aws/internal/controller/elasticbeanstalk/applicationversion"
	configurationtemplate "github.com/spirosco/upbound-provider-aws/internal/controller/elasticbeanstalk/configurationtemplate"
	domainelasticsearch "github.com/spirosco/upbound-provider-aws/internal/controller/elasticsearch/domain"
	domainpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/elasticsearch/domainpolicy"
	domainsamloptions "github.com/spirosco/upbound-provider-aws/internal/controller/elasticsearch/domainsamloptions"
	pipelineelastictranscoder "github.com/spirosco/upbound-provider-aws/internal/controller/elastictranscoder/pipeline"
	preset "github.com/spirosco/upbound-provider-aws/internal/controller/elastictranscoder/preset"
	appcookiestickinesspolicy "github.com/spirosco/upbound-provider-aws/internal/controller/elb/appcookiestickinesspolicy"
	attachmentelb "github.com/spirosco/upbound-provider-aws/internal/controller/elb/attachment"
	backendserverpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/elb/backendserverpolicy"
	elb "github.com/spirosco/upbound-provider-aws/internal/controller/elb/elb"
	lbcookiestickinesspolicy "github.com/spirosco/upbound-provider-aws/internal/controller/elb/lbcookiestickinesspolicy"
	lbsslnegotiationpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/elb/lbsslnegotiationpolicy"
	listenerpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/elb/listenerpolicy"
	policyelb "github.com/spirosco/upbound-provider-aws/internal/controller/elb/policy"
	proxyprotocolpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/elb/proxyprotocolpolicy"
	lb "github.com/spirosco/upbound-provider-aws/internal/controller/elbv2/lb"
	lblistener "github.com/spirosco/upbound-provider-aws/internal/controller/elbv2/lblistener"
	lblistenerrule "github.com/spirosco/upbound-provider-aws/internal/controller/elbv2/lblistenerrule"
	lbtargetgroup "github.com/spirosco/upbound-provider-aws/internal/controller/elbv2/lbtargetgroup"
	lbtargetgroupattachment "github.com/spirosco/upbound-provider-aws/internal/controller/elbv2/lbtargetgroupattachment"
	securityconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/emr/securityconfiguration"
	applicationemrserverless "github.com/spirosco/upbound-provider-aws/internal/controller/emrserverless/application"
	feature "github.com/spirosco/upbound-provider-aws/internal/controller/evidently/feature"
	projectevidently "github.com/spirosco/upbound-provider-aws/internal/controller/evidently/project"
	segment "github.com/spirosco/upbound-provider-aws/internal/controller/evidently/segment"
	deliverystream "github.com/spirosco/upbound-provider-aws/internal/controller/firehose/deliverystream"
	experimenttemplate "github.com/spirosco/upbound-provider-aws/internal/controller/fis/experimenttemplate"
	backup "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/backup"
	datarepositoryassociation "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/datarepositoryassociation"
	lustrefilesystem "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/lustrefilesystem"
	ontapfilesystem "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/ontapfilesystem"
	ontapstoragevirtualmachine "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/ontapstoragevirtualmachine"
	windowsfilesystem "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/windowsfilesystem"
	alias "github.com/spirosco/upbound-provider-aws/internal/controller/gamelift/alias"
	build "github.com/spirosco/upbound-provider-aws/internal/controller/gamelift/build"
	fleetgamelift "github.com/spirosco/upbound-provider-aws/internal/controller/gamelift/fleet"
	gamesessionqueue "github.com/spirosco/upbound-provider-aws/internal/controller/gamelift/gamesessionqueue"
	script "github.com/spirosco/upbound-provider-aws/internal/controller/gamelift/script"
	vaultglacier "github.com/spirosco/upbound-provider-aws/internal/controller/glacier/vault"
	vaultlock "github.com/spirosco/upbound-provider-aws/internal/controller/glacier/vaultlock"
	accelerator "github.com/spirosco/upbound-provider-aws/internal/controller/globalaccelerator/accelerator"
	endpointgroup "github.com/spirosco/upbound-provider-aws/internal/controller/globalaccelerator/endpointgroup"
	listener "github.com/spirosco/upbound-provider-aws/internal/controller/globalaccelerator/listener"
	catalogdatabase "github.com/spirosco/upbound-provider-aws/internal/controller/glue/catalogdatabase"
	catalogtable "github.com/spirosco/upbound-provider-aws/internal/controller/glue/catalogtable"
	classifier "github.com/spirosco/upbound-provider-aws/internal/controller/glue/classifier"
	connectionglue "github.com/spirosco/upbound-provider-aws/internal/controller/glue/connection"
	crawler "github.com/spirosco/upbound-provider-aws/internal/controller/glue/crawler"
	datacatalogencryptionsettings "github.com/spirosco/upbound-provider-aws/internal/controller/glue/datacatalogencryptionsettings"
	job "github.com/spirosco/upbound-provider-aws/internal/controller/glue/job"
	registry "github.com/spirosco/upbound-provider-aws/internal/controller/glue/registry"
	resourcepolicyglue "github.com/spirosco/upbound-provider-aws/internal/controller/glue/resourcepolicy"
	schema "github.com/spirosco/upbound-provider-aws/internal/controller/glue/schema"
	securityconfigurationglue "github.com/spirosco/upbound-provider-aws/internal/controller/glue/securityconfiguration"
	triggerglue "github.com/spirosco/upbound-provider-aws/internal/controller/glue/trigger"
	userdefinedfunction "github.com/spirosco/upbound-provider-aws/internal/controller/glue/userdefinedfunction"
	workflow "github.com/spirosco/upbound-provider-aws/internal/controller/glue/workflow"
	licenseassociation "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/licenseassociation"
	roleassociation "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/roleassociation"
	workspacegrafana "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/workspace"
	workspaceapikey "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/workspaceapikey"
	workspacesamlconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/workspacesamlconfiguration"
	detector "github.com/spirosco/upbound-provider-aws/internal/controller/guardduty/detector"
	filter "github.com/spirosco/upbound-provider-aws/internal/controller/guardduty/filter"
	memberguardduty "github.com/spirosco/upbound-provider-aws/internal/controller/guardduty/member"
	accesskey "github.com/spirosco/upbound-provider-aws/internal/controller/iam/accesskey"
	accountalias "github.com/spirosco/upbound-provider-aws/internal/controller/iam/accountalias"
	accountpasswordpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/iam/accountpasswordpolicy"
	groupiam "github.com/spirosco/upbound-provider-aws/internal/controller/iam/group"
	groupmembership "github.com/spirosco/upbound-provider-aws/internal/controller/iam/groupmembership"
	grouppolicyattachment "github.com/spirosco/upbound-provider-aws/internal/controller/iam/grouppolicyattachment"
	instanceprofileiam "github.com/spirosco/upbound-provider-aws/internal/controller/iam/instanceprofile"
	openidconnectprovider "github.com/spirosco/upbound-provider-aws/internal/controller/iam/openidconnectprovider"
	policyiam "github.com/spirosco/upbound-provider-aws/internal/controller/iam/policy"
	role "github.com/spirosco/upbound-provider-aws/internal/controller/iam/role"
	rolepolicyattachment "github.com/spirosco/upbound-provider-aws/internal/controller/iam/rolepolicyattachment"
	samlprovider "github.com/spirosco/upbound-provider-aws/internal/controller/iam/samlprovider"
	servercertificate "github.com/spirosco/upbound-provider-aws/internal/controller/iam/servercertificate"
	servicelinkedrole "github.com/spirosco/upbound-provider-aws/internal/controller/iam/servicelinkedrole"
	servicespecificcredential "github.com/spirosco/upbound-provider-aws/internal/controller/iam/servicespecificcredential"
	signingcertificate "github.com/spirosco/upbound-provider-aws/internal/controller/iam/signingcertificate"
	useriam "github.com/spirosco/upbound-provider-aws/internal/controller/iam/user"
	usergroupmembership "github.com/spirosco/upbound-provider-aws/internal/controller/iam/usergroupmembership"
	userloginprofile "github.com/spirosco/upbound-provider-aws/internal/controller/iam/userloginprofile"
	userpolicyattachment "github.com/spirosco/upbound-provider-aws/internal/controller/iam/userpolicyattachment"
	usersshkey "github.com/spirosco/upbound-provider-aws/internal/controller/iam/usersshkey"
	virtualmfadevice "github.com/spirosco/upbound-provider-aws/internal/controller/iam/virtualmfadevice"
	component "github.com/spirosco/upbound-provider-aws/internal/controller/imagebuilder/component"
	containerrecipe "github.com/spirosco/upbound-provider-aws/internal/controller/imagebuilder/containerrecipe"
	distributionconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/imagebuilder/distributionconfiguration"
	image "github.com/spirosco/upbound-provider-aws/internal/controller/imagebuilder/image"
	imagepipeline "github.com/spirosco/upbound-provider-aws/internal/controller/imagebuilder/imagepipeline"
	imagerecipe "github.com/spirosco/upbound-provider-aws/internal/controller/imagebuilder/imagerecipe"
	infrastructureconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/imagebuilder/infrastructureconfiguration"
	assessmenttarget "github.com/spirosco/upbound-provider-aws/internal/controller/inspector/assessmenttarget"
	assessmenttemplate "github.com/spirosco/upbound-provider-aws/internal/controller/inspector/assessmenttemplate"
	resourcegroup "github.com/spirosco/upbound-provider-aws/internal/controller/inspector/resourcegroup"
	enabler "github.com/spirosco/upbound-provider-aws/internal/controller/inspector2/enabler"
	certificateiot "github.com/spirosco/upbound-provider-aws/internal/controller/iot/certificate"
	indexingconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/iot/indexingconfiguration"
	loggingoptions "github.com/spirosco/upbound-provider-aws/internal/controller/iot/loggingoptions"
	policyiot "github.com/spirosco/upbound-provider-aws/internal/controller/iot/policy"
	policyattachment "github.com/spirosco/upbound-provider-aws/internal/controller/iot/policyattachment"
	provisioningtemplate "github.com/spirosco/upbound-provider-aws/internal/controller/iot/provisioningtemplate"
	rolealias "github.com/spirosco/upbound-provider-aws/internal/controller/iot/rolealias"
	thing "github.com/spirosco/upbound-provider-aws/internal/controller/iot/thing"
	thinggroup "github.com/spirosco/upbound-provider-aws/internal/controller/iot/thinggroup"
	thinggroupmembership "github.com/spirosco/upbound-provider-aws/internal/controller/iot/thinggroupmembership"
	thingprincipalattachment "github.com/spirosco/upbound-provider-aws/internal/controller/iot/thingprincipalattachment"
	thingtype "github.com/spirosco/upbound-provider-aws/internal/controller/iot/thingtype"
	topicrule "github.com/spirosco/upbound-provider-aws/internal/controller/iot/topicrule"
	channel "github.com/spirosco/upbound-provider-aws/internal/controller/ivs/channel"
	recordingconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/ivs/recordingconfiguration"
	clusterkafka "github.com/spirosco/upbound-provider-aws/internal/controller/kafka/cluster"
	configuration "github.com/spirosco/upbound-provider-aws/internal/controller/kafka/configuration"
	datasourcekendra "github.com/spirosco/upbound-provider-aws/internal/controller/kendra/datasource"
	experience "github.com/spirosco/upbound-provider-aws/internal/controller/kendra/experience"
	index "github.com/spirosco/upbound-provider-aws/internal/controller/kendra/index"
	querysuggestionsblocklist "github.com/spirosco/upbound-provider-aws/internal/controller/kendra/querysuggestionsblocklist"
	thesaurus "github.com/spirosco/upbound-provider-aws/internal/controller/kendra/thesaurus"
	keyspace "github.com/spirosco/upbound-provider-aws/internal/controller/keyspaces/keyspace"
	tablekeyspaces "github.com/spirosco/upbound-provider-aws/internal/controller/keyspaces/table"
	streamkinesis "github.com/spirosco/upbound-provider-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/spirosco/upbound-provider-aws/internal/controller/kinesis/streamconsumer"
	applicationkinesisanalytics "github.com/spirosco/upbound-provider-aws/internal/controller/kinesisanalytics/application"
	applicationkinesisanalyticsv2 "github.com/spirosco/upbound-provider-aws/internal/controller/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/spirosco/upbound-provider-aws/internal/controller/kinesisanalyticsv2/applicationsnapshot"
	streamkinesisvideo "github.com/spirosco/upbound-provider-aws/internal/controller/kinesisvideo/stream"
	aliaskms "github.com/spirosco/upbound-provider-aws/internal/controller/kms/alias"
	ciphertext "github.com/spirosco/upbound-provider-aws/internal/controller/kms/ciphertext"
	externalkey "github.com/spirosco/upbound-provider-aws/internal/controller/kms/externalkey"
	grant "github.com/spirosco/upbound-provider-aws/internal/controller/kms/grant"
	key "github.com/spirosco/upbound-provider-aws/internal/controller/kms/key"
	replicaexternalkey "github.com/spirosco/upbound-provider-aws/internal/controller/kms/replicaexternalkey"
	replicakey "github.com/spirosco/upbound-provider-aws/internal/controller/kms/replicakey"
	datalakesettings "github.com/spirosco/upbound-provider-aws/internal/controller/lakeformation/datalakesettings"
	permissions "github.com/spirosco/upbound-provider-aws/internal/controller/lakeformation/permissions"
	resourcelakeformation "github.com/spirosco/upbound-provider-aws/internal/controller/lakeformation/resource"
	aliaslambda "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/alias"
	codesigningconfig "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/codesigningconfig"
	eventsourcemapping "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/eventsourcemapping"
	functionlambda "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/function"
	functioneventinvokeconfig "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/functioneventinvokeconfig"
	functionurl "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/functionurl"
	invocation "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/invocation"
	layerversion "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/layerversion"
	layerversionpermission "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/layerversionpermission"
	permissionlambda "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/permission"
	provisionedconcurrencyconfig "github.com/spirosco/upbound-provider-aws/internal/controller/lambda/provisionedconcurrencyconfig"
	bot "github.com/spirosco/upbound-provider-aws/internal/controller/lexmodels/bot"
	botalias "github.com/spirosco/upbound-provider-aws/internal/controller/lexmodels/botalias"
	intent "github.com/spirosco/upbound-provider-aws/internal/controller/lexmodels/intent"
	slottype "github.com/spirosco/upbound-provider-aws/internal/controller/lexmodels/slottype"
	association "github.com/spirosco/upbound-provider-aws/internal/controller/licensemanager/association"
	licenseconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/licensemanager/licenseconfiguration"
	bucket "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/bucket"
	certificatelightsail "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/certificate"
	containerservice "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/containerservice"
	disk "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/disk"
	diskattachment "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/diskattachment"
	domainlightsail "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/domain"
	domainentry "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/domainentry"
	instancelightsail "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/instance"
	instancepublicports "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/instancepublicports"
	keypairlightsail "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/keypair"
	lblightsail "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/lb"
	lbattachment "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/lbattachment"
	lbcertificate "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/lbcertificate"
	lbstickinesspolicy "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/lbstickinesspolicy"
	staticip "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/staticip"
	staticipattachment "github.com/spirosco/upbound-provider-aws/internal/controller/lightsail/staticipattachment"
	geofencecollection "github.com/spirosco/upbound-provider-aws/internal/controller/location/geofencecollection"
	placeindex "github.com/spirosco/upbound-provider-aws/internal/controller/location/placeindex"
	routecalculator "github.com/spirosco/upbound-provider-aws/internal/controller/location/routecalculator"
	tracker "github.com/spirosco/upbound-provider-aws/internal/controller/location/tracker"
	trackerassociation "github.com/spirosco/upbound-provider-aws/internal/controller/location/trackerassociation"
	accountmacie2 "github.com/spirosco/upbound-provider-aws/internal/controller/macie2/account"
	classificationjob "github.com/spirosco/upbound-provider-aws/internal/controller/macie2/classificationjob"
	customdataidentifier "github.com/spirosco/upbound-provider-aws/internal/controller/macie2/customdataidentifier"
	findingsfilter "github.com/spirosco/upbound-provider-aws/internal/controller/macie2/findingsfilter"
	invitationacceptermacie2 "github.com/spirosco/upbound-provider-aws/internal/controller/macie2/invitationaccepter"
	membermacie2 "github.com/spirosco/upbound-provider-aws/internal/controller/macie2/member"
	queuemediaconvert "github.com/spirosco/upbound-provider-aws/internal/controller/mediaconvert/queue"
	channelmedialive "github.com/spirosco/upbound-provider-aws/internal/controller/medialive/channel"
	input "github.com/spirosco/upbound-provider-aws/internal/controller/medialive/input"
	inputsecuritygroup "github.com/spirosco/upbound-provider-aws/internal/controller/medialive/inputsecuritygroup"
	multiplex "github.com/spirosco/upbound-provider-aws/internal/controller/medialive/multiplex"
	channelmediapackage "github.com/spirosco/upbound-provider-aws/internal/controller/mediapackage/channel"
	container "github.com/spirosco/upbound-provider-aws/internal/controller/mediastore/container"
	containerpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/mediastore/containerpolicy"
	acl "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/acl"
	clustermemorydb "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/cluster"
	parametergroupmemorydb "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/parametergroup"
	snapshot "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/snapshot"
	subnetgroupmemorydb "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/subnetgroup"
	broker "github.com/spirosco/upbound-provider-aws/internal/controller/mq/broker"
	configurationmq "github.com/spirosco/upbound-provider-aws/internal/controller/mq/configuration"
	clusterneptune "github.com/spirosco/upbound-provider-aws/internal/controller/neptune/cluster"
	clusterendpoint "github.com/spirosco/upbound-provider-aws/internal/controller/neptune/clusterendpoint"
	clusterinstanceneptune "github.com/spirosco/upbound-provider-aws/internal/controller/neptune/clusterinstance"
	clusterparametergroupneptune "github.com/spirosco/upbound-provider-aws/internal/controller/neptune/clusterparametergroup"
	clustersnapshotneptune "github.com/spirosco/upbound-provider-aws/internal/controller/neptune/clustersnapshot"
	eventsubscriptionneptune "github.com/spirosco/upbound-provider-aws/internal/controller/neptune/eventsubscription"
	globalclusterneptune "github.com/spirosco/upbound-provider-aws/internal/controller/neptune/globalcluster"
	parametergroupneptune "github.com/spirosco/upbound-provider-aws/internal/controller/neptune/parametergroup"
	subnetgroupneptune "github.com/spirosco/upbound-provider-aws/internal/controller/neptune/subnetgroup"
	firewall "github.com/spirosco/upbound-provider-aws/internal/controller/networkfirewall/firewall"
	firewallpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/networkfirewall/firewallpolicy"
	loggingconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/networkfirewall/loggingconfiguration"
	rulegroup "github.com/spirosco/upbound-provider-aws/internal/controller/networkfirewall/rulegroup"
	attachmentaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/attachmentaccepter"
	connectattachment "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/connectattachment"
	connectionnetworkmanager "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/connection"
	corenetwork "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/corenetwork"
	customergatewayassociation "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/customergatewayassociation"
	device "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/device"
	globalnetwork "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/globalnetwork"
	link "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/link"
	linkassociation "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/linkassociation"
	site "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/site"
	transitgatewayconnectpeerassociation "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/transitgatewayconnectpeerassociation"
	transitgatewayregistration "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/transitgatewayregistration"
	vpcattachment "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/vpcattachment"
	domainopensearch "github.com/spirosco/upbound-provider-aws/internal/controller/opensearch/domain"
	domainpolicyopensearch "github.com/spirosco/upbound-provider-aws/internal/controller/opensearch/domainpolicy"
	domainsamloptionsopensearch "github.com/spirosco/upbound-provider-aws/internal/controller/opensearch/domainsamloptions"
	applicationopsworks "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/application"
	customlayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/customlayer"
	ecsclusterlayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/ecsclusterlayer"
	ganglialayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/ganglialayer"
	haproxylayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/haproxylayer"
	instanceopsworks "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/instance"
	javaapplayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/javaapplayer"
	memcachedlayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/memcachedlayer"
	mysqllayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/mysqllayer"
	nodejsapplayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/nodejsapplayer"
	permissionopsworks "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/permission"
	phpapplayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/phpapplayer"
	railsapplayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/railsapplayer"
	rdsdbinstance "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/rdsdbinstance"
	stackopsworks "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/stack"
	staticweblayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/staticweblayer"
	userprofile "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/userprofile"
	accountorganizations "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/account"
	delegatedadministrator "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/delegatedadministrator"
	organization "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/organization"
	organizationalunit "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/organizationalunit"
	policyorganizations "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/policy"
	policyattachmentorganizations "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/policyattachment"
	apppinpoint "github.com/spirosco/upbound-provider-aws/internal/controller/pinpoint/app"
	smschannel "github.com/spirosco/upbound-provider-aws/internal/controller/pinpoint/smschannel"
	providerconfig "github.com/spirosco/upbound-provider-aws/internal/controller/providerconfig"
	ledger "github.com/spirosco/upbound-provider-aws/internal/controller/qldb/ledger"
	streamqldb "github.com/spirosco/upbound-provider-aws/internal/controller/qldb/stream"
	groupquicksight "github.com/spirosco/upbound-provider-aws/internal/controller/quicksight/group"
	userquicksight "github.com/spirosco/upbound-provider-aws/internal/controller/quicksight/user"
	resourceassociation "github.com/spirosco/upbound-provider-aws/internal/controller/ram/resourceassociation"
	resourceshare "github.com/spirosco/upbound-provider-aws/internal/controller/ram/resourceshare"
	clusterrds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/cluster"
	clusteractivitystream "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusteractivitystream"
	clusterendpointrds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusterendpoint"
	clusterinstancerds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusterinstance"
	clusterparametergrouprds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshotrds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "github.com/spirosco/upbound-provider-aws/internal/controller/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "github.com/spirosco/upbound-provider-aws/internal/controller/rds/dbsnapshotcopy"
	eventsubscriptionrds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/eventsubscription"
	globalclusterrds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/globalcluster"
	instancerds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/instance"
	instanceroleassociation "github.com/spirosco/upbound-provider-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "github.com/spirosco/upbound-provider-aws/internal/controller/rds/optiongroup"
	parametergrouprds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/parametergroup"
	proxy "github.com/spirosco/upbound-provider-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/spirosco/upbound-provider-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/spirosco/upbound-provider-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/spirosco/upbound-provider-aws/internal/controller/rds/proxytarget"
	securitygrouprds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/securitygroup"
	snapshotrds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/snapshot"
	subnetgrouprds "github.com/spirosco/upbound-provider-aws/internal/controller/rds/subnetgroup"
	authenticationprofile "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/authenticationprofile"
	clusterredshift "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/cluster"
	eventsubscriptionredshift "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/eventsubscription"
	hsmclientcertificate "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/hsmclientcertificate"
	hsmconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/hsmconfiguration"
	parametergroupredshift "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/parametergroup"
	scheduledactionredshift "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/scheduledaction"
	snapshotcopygrant "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/snapshotcopygrant"
	snapshotschedule "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/snapshotschedule"
	snapshotscheduleassociation "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/snapshotscheduleassociation"
	subnetgroupredshift "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/subnetgroup"
	usagelimit "github.com/spirosco/upbound-provider-aws/internal/controller/redshift/usagelimit"
	groupresourcegroups "github.com/spirosco/upbound-provider-aws/internal/controller/resourcegroups/group"
	profile "github.com/spirosco/upbound-provider-aws/internal/controller/rolesanywhere/profile"
	delegationset "github.com/spirosco/upbound-provider-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/spirosco/upbound-provider-aws/internal/controller/route53/healthcheck"
	hostedzonednssec "github.com/spirosco/upbound-provider-aws/internal/controller/route53/hostedzonednssec"
	record "github.com/spirosco/upbound-provider-aws/internal/controller/route53/record"
	resolverconfig "github.com/spirosco/upbound-provider-aws/internal/controller/route53/resolverconfig"
	trafficpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/route53/trafficpolicy"
	trafficpolicyinstance "github.com/spirosco/upbound-provider-aws/internal/controller/route53/trafficpolicyinstance"
	vpcassociationauthorization "github.com/spirosco/upbound-provider-aws/internal/controller/route53/vpcassociationauthorization"
	zone "github.com/spirosco/upbound-provider-aws/internal/controller/route53/zone"
	clusterroute53recoverycontrolconfig "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoverycontrolconfig/cluster"
	controlpanel "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoverycontrolconfig/controlpanel"
	routingcontrol "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoverycontrolconfig/routingcontrol"
	safetyrule "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoverycontrolconfig/safetyrule"
	cell "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoveryreadiness/cell"
	readinesscheck "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoveryreadiness/readinesscheck"
	recoverygroup "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoveryreadiness/recoverygroup"
	resourceset "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoveryreadiness/resourceset"
	endpointroute53resolver "github.com/spirosco/upbound-provider-aws/internal/controller/route53resolver/endpoint"
	ruleroute53resolver "github.com/spirosco/upbound-provider-aws/internal/controller/route53resolver/rule"
	ruleassociation "github.com/spirosco/upbound-provider-aws/internal/controller/route53resolver/ruleassociation"
	appmonitor "github.com/spirosco/upbound-provider-aws/internal/controller/rum/appmonitor"
	metricsdestination "github.com/spirosco/upbound-provider-aws/internal/controller/rum/metricsdestination"
	buckets3 "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucket"
	bucketaccelerateconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketaccelerateconfiguration"
	bucketacl "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketacl"
	bucketanalyticsconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketanalyticsconfiguration"
	bucketcorsconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketcorsconfiguration"
	bucketintelligenttieringconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketintelligenttieringconfiguration"
	bucketinventory "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketinventory"
	bucketlifecycleconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketlifecycleconfiguration"
	bucketlogging "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketlogging"
	bucketmetric "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketmetric"
	bucketnotification "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketnotification"
	bucketobject "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketobject"
	bucketobjectlockconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketobjectlockconfiguration"
	bucketownershipcontrols "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketownershipcontrols"
	bucketpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketpublicaccessblock"
	bucketreplicationconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketreplicationconfiguration"
	bucketrequestpaymentconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketrequestpaymentconfiguration"
	bucketserversideencryptionconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3/bucketwebsiteconfiguration"
	object "github.com/spirosco/upbound-provider-aws/internal/controller/s3/object"
	objectcopy "github.com/spirosco/upbound-provider-aws/internal/controller/s3/objectcopy"
	accesspoints3control "github.com/spirosco/upbound-provider-aws/internal/controller/s3control/accesspoint"
	accesspointpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/s3control/accesspointpolicy"
	accountpublicaccessblock "github.com/spirosco/upbound-provider-aws/internal/controller/s3control/accountpublicaccessblock"
	multiregionaccesspoint "github.com/spirosco/upbound-provider-aws/internal/controller/s3control/multiregionaccesspoint"
	multiregionaccesspointpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/s3control/multiregionaccesspointpolicy"
	objectlambdaaccesspoint "github.com/spirosco/upbound-provider-aws/internal/controller/s3control/objectlambdaaccesspoint"
	objectlambdaaccesspointpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/s3control/objectlambdaaccesspointpolicy"
	storagelensconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/s3control/storagelensconfiguration"
	appsagemaker "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/app"
	appimageconfig "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/appimageconfig"
	coderepository "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/coderepository"
	devicesagemaker "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/device"
	devicefleet "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/devicefleet"
	domainsagemaker "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/domain"
	endpointconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/endpointconfiguration"
	featuregroup "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/featuregroup"
	imagesagemaker "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/image"
	imageversion "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/imageversion"
	modelsagemaker "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/model"
	modelpackagegroup "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/modelpackagegroup"
	modelpackagegrouppolicy "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/modelpackagegrouppolicy"
	notebookinstance "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/notebookinstance"
	notebookinstancelifecycleconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/notebookinstancelifecycleconfiguration"
	servicecatalogportfoliostatus "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/servicecatalogportfoliostatus"
	space "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/space"
	studiolifecycleconfig "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/studiolifecycleconfig"
	userprofilesagemaker "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/userprofile"
	workforce "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/workforce"
	workteam "github.com/spirosco/upbound-provider-aws/internal/controller/sagemaker/workteam"
	schedulescheduler "github.com/spirosco/upbound-provider-aws/internal/controller/scheduler/schedule"
	schedulegroup "github.com/spirosco/upbound-provider-aws/internal/controller/scheduler/schedulegroup"
	discoverer "github.com/spirosco/upbound-provider-aws/internal/controller/schemas/discoverer"
	registryschemas "github.com/spirosco/upbound-provider-aws/internal/controller/schemas/registry"
	schemaschemas "github.com/spirosco/upbound-provider-aws/internal/controller/schemas/schema"
	secret "github.com/spirosco/upbound-provider-aws/internal/controller/secretsmanager/secret"
	secretpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/secretsmanager/secretpolicy"
	secretrotation "github.com/spirosco/upbound-provider-aws/internal/controller/secretsmanager/secretrotation"
	secretversion "github.com/spirosco/upbound-provider-aws/internal/controller/secretsmanager/secretversion"
	accountsecurityhub "github.com/spirosco/upbound-provider-aws/internal/controller/securityhub/account"
	actiontarget "github.com/spirosco/upbound-provider-aws/internal/controller/securityhub/actiontarget"
	findingaggregator "github.com/spirosco/upbound-provider-aws/internal/controller/securityhub/findingaggregator"
	insight "github.com/spirosco/upbound-provider-aws/internal/controller/securityhub/insight"
	inviteaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/securityhub/inviteaccepter"
	membersecurityhub "github.com/spirosco/upbound-provider-aws/internal/controller/securityhub/member"
	productsubscription "github.com/spirosco/upbound-provider-aws/internal/controller/securityhub/productsubscription"
	standardssubscription "github.com/spirosco/upbound-provider-aws/internal/controller/securityhub/standardssubscription"
	cloudformationstack "github.com/spirosco/upbound-provider-aws/internal/controller/serverlessrepo/cloudformationstack"
	budgetresourceassociation "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/budgetresourceassociation"
	constraint "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/constraint"
	portfolio "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/portfolio"
	portfolioshare "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/portfolioshare"
	principalportfolioassociation "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/principalportfolioassociation"
	product "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/product"
	productportfolioassociation "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/productportfolioassociation"
	provisioningartifact "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/provisioningartifact"
	serviceaction "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/serviceaction"
	tagoption "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/tagoption"
	tagoptionresourceassociation "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/tagoptionresourceassociation"
	httpnamespace "github.com/spirosco/upbound-provider-aws/internal/controller/servicediscovery/httpnamespace"
	privatednsnamespace "github.com/spirosco/upbound-provider-aws/internal/controller/servicediscovery/privatednsnamespace"
	publicdnsnamespace "github.com/spirosco/upbound-provider-aws/internal/controller/servicediscovery/publicdnsnamespace"
	serviceservicediscovery "github.com/spirosco/upbound-provider-aws/internal/controller/servicediscovery/service"
	servicequota "github.com/spirosco/upbound-provider-aws/internal/controller/servicequotas/servicequota"
	activereceiptruleset "github.com/spirosco/upbound-provider-aws/internal/controller/ses/activereceiptruleset"
	configurationset "github.com/spirosco/upbound-provider-aws/internal/controller/ses/configurationset"
	domaindkim "github.com/spirosco/upbound-provider-aws/internal/controller/ses/domaindkim"
	domainidentity "github.com/spirosco/upbound-provider-aws/internal/controller/ses/domainidentity"
	domainmailfrom "github.com/spirosco/upbound-provider-aws/internal/controller/ses/domainmailfrom"
	emailidentity "github.com/spirosco/upbound-provider-aws/internal/controller/ses/emailidentity"
	eventdestination "github.com/spirosco/upbound-provider-aws/internal/controller/ses/eventdestination"
	identitynotificationtopic "github.com/spirosco/upbound-provider-aws/internal/controller/ses/identitynotificationtopic"
	identitypolicy "github.com/spirosco/upbound-provider-aws/internal/controller/ses/identitypolicy"
	receiptfilter "github.com/spirosco/upbound-provider-aws/internal/controller/ses/receiptfilter"
	receiptrule "github.com/spirosco/upbound-provider-aws/internal/controller/ses/receiptrule"
	receiptruleset "github.com/spirosco/upbound-provider-aws/internal/controller/ses/receiptruleset"
	template "github.com/spirosco/upbound-provider-aws/internal/controller/ses/template"
	configurationsetsesv2 "github.com/spirosco/upbound-provider-aws/internal/controller/sesv2/configurationset"
	configurationseteventdestination "github.com/spirosco/upbound-provider-aws/internal/controller/sesv2/configurationseteventdestination"
	dedicatedippool "github.com/spirosco/upbound-provider-aws/internal/controller/sesv2/dedicatedippool"
	emailidentitysesv2 "github.com/spirosco/upbound-provider-aws/internal/controller/sesv2/emailidentity"
	emailidentityfeedbackattributes "github.com/spirosco/upbound-provider-aws/internal/controller/sesv2/emailidentityfeedbackattributes"
	emailidentitymailfromattributes "github.com/spirosco/upbound-provider-aws/internal/controller/sesv2/emailidentitymailfromattributes"
	activity "github.com/spirosco/upbound-provider-aws/internal/controller/sfn/activity"
	statemachine "github.com/spirosco/upbound-provider-aws/internal/controller/sfn/statemachine"
	signingjob "github.com/spirosco/upbound-provider-aws/internal/controller/signer/signingjob"
	signingprofile "github.com/spirosco/upbound-provider-aws/internal/controller/signer/signingprofile"
	signingprofilepermission "github.com/spirosco/upbound-provider-aws/internal/controller/signer/signingprofilepermission"
	domainsimpledb "github.com/spirosco/upbound-provider-aws/internal/controller/simpledb/domain"
	platformapplication "github.com/spirosco/upbound-provider-aws/internal/controller/sns/platformapplication"
	smspreferences "github.com/spirosco/upbound-provider-aws/internal/controller/sns/smspreferences"
	topic "github.com/spirosco/upbound-provider-aws/internal/controller/sns/topic"
	topicpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/sns/topicpolicy"
	topicsubscription "github.com/spirosco/upbound-provider-aws/internal/controller/sns/topicsubscription"
	queuesqs "github.com/spirosco/upbound-provider-aws/internal/controller/sqs/queue"
	queuepolicy "github.com/spirosco/upbound-provider-aws/internal/controller/sqs/queuepolicy"
	queueredriveallowpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/sqs/queueredriveallowpolicy"
	queueredrivepolicy "github.com/spirosco/upbound-provider-aws/internal/controller/sqs/queueredrivepolicy"
	activation "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/activation"
	associationssm "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/association"
	defaultpatchbaseline "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/defaultpatchbaseline"
	document "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/document"
	maintenancewindow "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/maintenancewindow"
	maintenancewindowtarget "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/maintenancewindowtarget"
	maintenancewindowtask "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/maintenancewindowtask"
	parameter "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/parameter"
	patchbaseline "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/patchbaseline"
	patchgroup "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/patchgroup"
	resourcedatasync "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/resourcedatasync"
	servicesetting "github.com/spirosco/upbound-provider-aws/internal/controller/ssm/servicesetting"
	accountassignment "github.com/spirosco/upbound-provider-aws/internal/controller/ssoadmin/accountassignment"
	managedpolicyattachment "github.com/spirosco/upbound-provider-aws/internal/controller/ssoadmin/managedpolicyattachment"
	permissionset "github.com/spirosco/upbound-provider-aws/internal/controller/ssoadmin/permissionset"
	permissionsetinlinepolicy "github.com/spirosco/upbound-provider-aws/internal/controller/ssoadmin/permissionsetinlinepolicy"
	domainswf "github.com/spirosco/upbound-provider-aws/internal/controller/swf/domain"
	databasetimestreamwrite "github.com/spirosco/upbound-provider-aws/internal/controller/timestreamwrite/database"
	tabletimestreamwrite "github.com/spirosco/upbound-provider-aws/internal/controller/timestreamwrite/table"
	languagemodel "github.com/spirosco/upbound-provider-aws/internal/controller/transcribe/languagemodel"
	vocabularytranscribe "github.com/spirosco/upbound-provider-aws/internal/controller/transcribe/vocabulary"
	vocabularyfilter "github.com/spirosco/upbound-provider-aws/internal/controller/transcribe/vocabularyfilter"
	server "github.com/spirosco/upbound-provider-aws/internal/controller/transfer/server"
	sshkey "github.com/spirosco/upbound-provider-aws/internal/controller/transfer/sshkey"
	tagtransfer "github.com/spirosco/upbound-provider-aws/internal/controller/transfer/tag"
	usertransfer "github.com/spirosco/upbound-provider-aws/internal/controller/transfer/user"
	workflowtransfer "github.com/spirosco/upbound-provider-aws/internal/controller/transfer/workflow"
	networkperformancemetricsubscription "github.com/spirosco/upbound-provider-aws/internal/controller/vpc/networkperformancemetricsubscription"
	bytematchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/bytematchset"
	geomatchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/geomatchset"
	ipset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/ipset"
	ratebasedrule "github.com/spirosco/upbound-provider-aws/internal/controller/waf/ratebasedrule"
	regexmatchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/regexmatchset"
	regexpatternset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/regexpatternset"
	rulewaf "github.com/spirosco/upbound-provider-aws/internal/controller/waf/rule"
	sizeconstraintset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/sizeconstraintset"
	sqlinjectionmatchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/sqlinjectionmatchset"
	webacl "github.com/spirosco/upbound-provider-aws/internal/controller/waf/webacl"
	xssmatchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/xssmatchset"
	bytematchsetwafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/bytematchset"
	geomatchsetwafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/geomatchset"
	ipsetwafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/ipset"
	ratebasedrulewafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/ratebasedrule"
	regexmatchsetwafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/regexmatchset"
	regexpatternsetwafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/regexpatternset"
	rulewafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/rule"
	sizeconstraintsetwafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/sizeconstraintset"
	sqlinjectionmatchsetwafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/sqlinjectionmatchset"
	webaclwafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/webacl"
	xssmatchsetwafregional "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/xssmatchset"
	ipsetwafv2 "github.com/spirosco/upbound-provider-aws/internal/controller/wafv2/ipset"
	regexpatternsetwafv2 "github.com/spirosco/upbound-provider-aws/internal/controller/wafv2/regexpatternset"
	directoryworkspaces "github.com/spirosco/upbound-provider-aws/internal/controller/workspaces/directory"
	ipgroup "github.com/spirosco/upbound-provider-aws/internal/controller/workspaces/ipgroup"
	encryptionconfig "github.com/spirosco/upbound-provider-aws/internal/controller/xray/encryptionconfig"
	groupxray "github.com/spirosco/upbound-provider-aws/internal/controller/xray/group"
	samplingrule "github.com/spirosco/upbound-provider-aws/internal/controller/xray/samplingrule"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzer.Setup,
		archiverule.Setup,
		alternatecontact.Setup,
		certificate.Setup,
		certificatevalidation.Setup,
		certificateacmpca.Setup,
		certificateauthority.Setup,
		certificateauthoritycertificate.Setup,
		permission.Setup,
		policy.Setup,
		alertmanagerdefinition.Setup,
		rulegroupnamespace.Setup,
		workspace.Setup,
		app.Setup,
		backendenvironment.Setup,
		branch.Setup,
		webhook.Setup,
		account.Setup,
		apikey.Setup,
		authorizer.Setup,
		basepathmapping.Setup,
		clientcertificate.Setup,
		deployment.Setup,
		documentationpart.Setup,
		documentationversion.Setup,
		domainname.Setup,
		gatewayresponse.Setup,
		integration.Setup,
		integrationresponse.Setup,
		method.Setup,
		methodresponse.Setup,
		methodsettings.Setup,
		model.Setup,
		requestvalidator.Setup,
		resource.Setup,
		restapi.Setup,
		restapipolicy.Setup,
		stage.Setup,
		usageplan.Setup,
		usageplankey.Setup,
		vpclink.Setup,
		api.Setup,
		apimapping.Setup,
		authorizerapigatewayv2.Setup,
		deploymentapigatewayv2.Setup,
		domainnameapigatewayv2.Setup,
		integrationapigatewayv2.Setup,
		integrationresponseapigatewayv2.Setup,
		modelapigatewayv2.Setup,
		route.Setup,
		routeresponse.Setup,
		stageapigatewayv2.Setup,
		vpclinkapigatewayv2.Setup,
		policyappautoscaling.Setup,
		scheduledaction.Setup,
		target.Setup,
		application.Setup,
		configurationprofile.Setup,
		deploymentappconfig.Setup,
		deploymentstrategy.Setup,
		environment.Setup,
		extension.Setup,
		extensionassociation.Setup,
		hostedconfigurationversion.Setup,
		flow.Setup,
		eventintegration.Setup,
		applicationapplicationinsights.Setup,
		gatewayroute.Setup,
		mesh.Setup,
		routeappmesh.Setup,
		virtualgateway.Setup,
		virtualnode.Setup,
		virtualrouter.Setup,
		virtualservice.Setup,
		autoscalingconfigurationversion.Setup,
		connection.Setup,
		observabilityconfiguration.Setup,
		service.Setup,
		vpcconnector.Setup,
		directoryconfig.Setup,
		fleet.Setup,
		fleetstackassociation.Setup,
		imagebuilder.Setup,
		stack.Setup,
		user.Setup,
		userstackassociation.Setup,
		apicache.Setup,
		apikeyappsync.Setup,
		datasource.Setup,
		function.Setup,
		graphqlapi.Setup,
		resolver.Setup,
		database.Setup,
		datacatalog.Setup,
		namedquery.Setup,
		workgroup.Setup,
		attachment.Setup,
		autoscalinggroup.Setup,
		grouptag.Setup,
		launchconfiguration.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		policyautoscaling.Setup,
		schedule.Setup,
		scalingplan.Setup,
		framework.Setup,
		globalsettings.Setup,
		plan.Setup,
		regionsettings.Setup,
		reportplan.Setup,
		selection.Setup,
		vault.Setup,
		vaultlockconfiguration.Setup,
		vaultnotifications.Setup,
		vaultpolicy.Setup,
		schedulingpolicy.Setup,
		budget.Setup,
		budgetaction.Setup,
		anomalymonitor.Setup,
		voiceconnector.Setup,
		voiceconnectorgroup.Setup,
		voiceconnectorlogging.Setup,
		voiceconnectororigination.Setup,
		voiceconnectorstreaming.Setup,
		voiceconnectortermination.Setup,
		voiceconnectorterminationcredentials.Setup,
		environmentec2.Setup,
		environmentmembership.Setup,
		resourcecloudcontrol.Setup,
		stackcloudformation.Setup,
		stackset.Setup,
		cachepolicy.Setup,
		distribution.Setup,
		fieldlevelencryptionconfig.Setup,
		fieldlevelencryptionprofile.Setup,
		functioncloudfront.Setup,
		keygroup.Setup,
		monitoringsubscription.Setup,
		originaccesscontrol.Setup,
		originaccessidentity.Setup,
		originrequestpolicy.Setup,
		publickey.Setup,
		realtimelogconfig.Setup,
		responseheaderspolicy.Setup,
		domain.Setup,
		domainserviceaccesspolicy.Setup,
		eventdatastore.Setup,
		trail.Setup,
		compositealarm.Setup,
		dashboard.Setup,
		metricalarm.Setup,
		metricstream.Setup,
		apidestination.Setup,
		archive.Setup,
		bus.Setup,
		buspolicy.Setup,
		connectioncloudwatchevents.Setup,
		permissioncloudwatchevents.Setup,
		rule.Setup,
		targetcloudwatchevents.Setup,
		definition.Setup,
		destination.Setup,
		destinationpolicy.Setup,
		group.Setup,
		metricfilter.Setup,
		resourcepolicy.Setup,
		stream.Setup,
		subscriptionfilter.Setup,
		approvalruletemplate.Setup,
		approvalruletemplateassociation.Setup,
		repository.Setup,
		trigger.Setup,
		codepipeline.Setup,
		customactiontype.Setup,
		webhookcodepipeline.Setup,
		connectioncodestarconnections.Setup,
		host.Setup,
		notificationrule.Setup,
		cognitoidentitypoolproviderprincipaltag.Setup,
		pool.Setup,
		poolrolesattachment.Setup,
		identityprovider.Setup,
		resourceserver.Setup,
		riskconfiguration.Setup,
		usercognitoidp.Setup,
		usergroup.Setup,
		useringroup.Setup,
		userpool.Setup,
		userpoolclient.Setup,
		userpooldomain.Setup,
		userpooluicustomization.Setup,
		awsconfigurationrecorderstatus.Setup,
		configrule.Setup,
		configurationaggregator.Setup,
		configurationrecorder.Setup,
		conformancepack.Setup,
		deliverychannel.Setup,
		remediationconfiguration.Setup,
		botassociation.Setup,
		contactflow.Setup,
		contactflowmodule.Setup,
		hoursofoperation.Setup,
		instance.Setup,
		instancestorageconfig.Setup,
		lambdafunctionassociation.Setup,
		phonenumber.Setup,
		queue.Setup,
		quickconnect.Setup,
		routingprofile.Setup,
		securityprofile.Setup,
		userconnect.Setup,
		userhierarchystructure.Setup,
		vocabulary.Setup,
		reportdefinition.Setup,
		dataset.Setup,
		revision.Setup,
		pipeline.Setup,
		cluster.Setup,
		parametergroup.Setup,
		subnetgroup.Setup,
		appdeploy.Setup,
		deploymentconfig.Setup,
		deploymentgroup.Setup,
		graph.Setup,
		invitationaccepter.Setup,
		member.Setup,
		devicepool.Setup,
		instanceprofile.Setup,
		networkprofile.Setup,
		project.Setup,
		testgridproject.Setup,
		upload.Setup,
		bgppeer.Setup,
		connectiondirectconnect.Setup,
		connectionassociation.Setup,
		gateway.Setup,
		gatewayassociation.Setup,
		gatewayassociationproposal.Setup,
		hostedprivatevirtualinterface.Setup,
		hostedprivatevirtualinterfaceaccepter.Setup,
		hostedpublicvirtualinterface.Setup,
		hostedpublicvirtualinterfaceaccepter.Setup,
		hostedtransitvirtualinterface.Setup,
		hostedtransitvirtualinterfaceaccepter.Setup,
		lag.Setup,
		privatevirtualinterface.Setup,
		publicvirtualinterface.Setup,
		transitvirtualinterface.Setup,
		lifecyclepolicy.Setup,
		certificatedms.Setup,
		endpoint.Setup,
		eventsubscription.Setup,
		replicationinstance.Setup,
		replicationsubnetgroup.Setup,
		replicationtask.Setup,
		s3endpoint.Setup,
		clusterdocdb.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clustersnapshot.Setup,
		eventsubscriptiondocdb.Setup,
		globalcluster.Setup,
		subnetgroupdocdb.Setup,
		conditionalforwarder.Setup,
		directory.Setup,
		shareddirectory.Setup,
		contributorinsights.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		table.Setup,
		tableitem.Setup,
		tablereplica.Setup,
		tag.Setup,
		ami.Setup,
		amicopy.Setup,
		amilaunchpermission.Setup,
		availabilityzonegroup.Setup,
		capacityreservation.Setup,
		carriergateway.Setup,
		customergateway.Setup,
		defaultnetworkacl.Setup,
		defaultroutetable.Setup,
		defaultsecuritygroup.Setup,
		defaultsubnet.Setup,
		defaultvpc.Setup,
		defaultvpcdhcpoptions.Setup,
		ebsdefaultkmskey.Setup,
		ebsencryptionbydefault.Setup,
		ebssnapshot.Setup,
		ebssnapshotcopy.Setup,
		ebssnapshotimport.Setup,
		ebsvolume.Setup,
		egressonlyinternetgateway.Setup,
		eip.Setup,
		eipassociation.Setup,
		flowlog.Setup,
		hostec2.Setup,
		instanceec2.Setup,
		instancestate.Setup,
		internetgateway.Setup,
		keypair.Setup,
		launchtemplate.Setup,
		mainroutetableassociation.Setup,
		managedprefixlist.Setup,
		managedprefixlistentry.Setup,
		natgateway.Setup,
		networkacl.Setup,
		networkaclrule.Setup,
		networkinsightsanalysis.Setup,
		networkinsightspath.Setup,
		networkinterface.Setup,
		networkinterfaceattachment.Setup,
		networkinterfacesgattachment.Setup,
		placementgroup.Setup,
		routeec2.Setup,
		routetable.Setup,
		routetableassociation.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		serialconsoleaccess.Setup,
		snapshotcreatevolumepermission.Setup,
		spotdatafeedsubscription.Setup,
		spotfleetrequest.Setup,
		spotinstancerequest.Setup,
		subnet.Setup,
		subnetcidrreservation.Setup,
		tagec2.Setup,
		trafficmirrorfilter.Setup,
		trafficmirrorfilterrule.Setup,
		transitgateway.Setup,
		transitgatewayconnect.Setup,
		transitgatewayconnectpeer.Setup,
		transitgatewaymulticastdomain.Setup,
		transitgatewaymulticastdomainassociation.Setup,
		transitgatewaymulticastgroupmember.Setup,
		transitgatewaymulticastgroupsource.Setup,
		transitgatewaypeeringattachment.Setup,
		transitgatewaypeeringattachmentaccepter.Setup,
		transitgatewaypolicytable.Setup,
		transitgatewayprefixlistreference.Setup,
		transitgatewayroute.Setup,
		transitgatewayroutetable.Setup,
		transitgatewayroutetableassociation.Setup,
		transitgatewayroutetablepropagation.Setup,
		transitgatewayvpcattachment.Setup,
		transitgatewayvpcattachmentaccepter.Setup,
		volumeattachment.Setup,
		vpc.Setup,
		vpcdhcpoptions.Setup,
		vpcdhcpoptionsassociation.Setup,
		vpcendpoint.Setup,
		vpcendpointconnectionnotification.Setup,
		vpcendpointroutetableassociation.Setup,
		vpcendpointsecuritygroupassociation.Setup,
		vpcendpointservice.Setup,
		vpcendpointserviceallowedprincipal.Setup,
		vpcendpointsubnetassociation.Setup,
		vpcipam.Setup,
		vpcipampool.Setup,
		vpcipampoolcidr.Setup,
		vpcipampoolcidrallocation.Setup,
		vpcipamscope.Setup,
		vpcipv4cidrblockassociation.Setup,
		vpcpeeringconnection.Setup,
		vpcpeeringconnectionaccepter.Setup,
		vpcpeeringconnectionoptions.Setup,
		vpnconnection.Setup,
		vpnconnectionroute.Setup,
		vpngateway.Setup,
		vpngatewayattachment.Setup,
		vpngatewayroutepropagation.Setup,
		lifecyclepolicyecr.Setup,
		pullthroughcacherule.Setup,
		registrypolicy.Setup,
		registryscanningconfiguration.Setup,
		replicationconfiguration.Setup,
		repositoryecr.Setup,
		repositorypolicy.Setup,
		repositoryecrpublic.Setup,
		repositorypolicyecrpublic.Setup,
		accountsettingdefault.Setup,
		capacityprovider.Setup,
		clusterecs.Setup,
		clustercapacityproviders.Setup,
		serviceecs.Setup,
		taskdefinition.Setup,
		accesspoint.Setup,
		backuppolicy.Setup,
		filesystem.Setup,
		filesystempolicy.Setup,
		mounttarget.Setup,
		replicationconfigurationefs.Setup,
		addon.Setup,
		clustereks.Setup,
		clusterauth.Setup,
		fargateprofile.Setup,
		identityproviderconfig.Setup,
		nodegroup.Setup,
		clusterelasticache.Setup,
		parametergroupelasticache.Setup,
		replicationgroup.Setup,
		subnetgroupelasticache.Setup,
		userelasticache.Setup,
		usergroupelasticache.Setup,
		applicationelasticbeanstalk.Setup,
		applicationversion.Setup,
		configurationtemplate.Setup,
		domainelasticsearch.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
		pipelineelastictranscoder.Setup,
		preset.Setup,
		appcookiestickinesspolicy.Setup,
		attachmentelb.Setup,
		backendserverpolicy.Setup,
		elb.Setup,
		lbcookiestickinesspolicy.Setup,
		lbsslnegotiationpolicy.Setup,
		listenerpolicy.Setup,
		policyelb.Setup,
		proxyprotocolpolicy.Setup,
		lb.Setup,
		lblistener.Setup,
		lblistenerrule.Setup,
		lbtargetgroup.Setup,
		lbtargetgroupattachment.Setup,
		securityconfiguration.Setup,
		applicationemrserverless.Setup,
		feature.Setup,
		projectevidently.Setup,
		segment.Setup,
		deliverystream.Setup,
		experimenttemplate.Setup,
		backup.Setup,
		datarepositoryassociation.Setup,
		lustrefilesystem.Setup,
		ontapfilesystem.Setup,
		ontapstoragevirtualmachine.Setup,
		windowsfilesystem.Setup,
		alias.Setup,
		build.Setup,
		fleetgamelift.Setup,
		gamesessionqueue.Setup,
		script.Setup,
		vaultglacier.Setup,
		vaultlock.Setup,
		accelerator.Setup,
		endpointgroup.Setup,
		listener.Setup,
		catalogdatabase.Setup,
		catalogtable.Setup,
		classifier.Setup,
		connectionglue.Setup,
		crawler.Setup,
		datacatalogencryptionsettings.Setup,
		job.Setup,
		registry.Setup,
		resourcepolicyglue.Setup,
		schema.Setup,
		securityconfigurationglue.Setup,
		triggerglue.Setup,
		userdefinedfunction.Setup,
		workflow.Setup,
		licenseassociation.Setup,
		roleassociation.Setup,
		workspacegrafana.Setup,
		workspaceapikey.Setup,
		workspacesamlconfiguration.Setup,
		detector.Setup,
		filter.Setup,
		memberguardduty.Setup,
		accesskey.Setup,
		accountalias.Setup,
		accountpasswordpolicy.Setup,
		groupiam.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		instanceprofileiam.Setup,
		openidconnectprovider.Setup,
		policyiam.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		samlprovider.Setup,
		servercertificate.Setup,
		servicelinkedrole.Setup,
		servicespecificcredential.Setup,
		signingcertificate.Setup,
		useriam.Setup,
		usergroupmembership.Setup,
		userloginprofile.Setup,
		userpolicyattachment.Setup,
		usersshkey.Setup,
		virtualmfadevice.Setup,
		component.Setup,
		containerrecipe.Setup,
		distributionconfiguration.Setup,
		image.Setup,
		imagepipeline.Setup,
		imagerecipe.Setup,
		infrastructureconfiguration.Setup,
		assessmenttarget.Setup,
		assessmenttemplate.Setup,
		resourcegroup.Setup,
		enabler.Setup,
		certificateiot.Setup,
		indexingconfiguration.Setup,
		loggingoptions.Setup,
		policyiot.Setup,
		policyattachment.Setup,
		provisioningtemplate.Setup,
		rolealias.Setup,
		thing.Setup,
		thinggroup.Setup,
		thinggroupmembership.Setup,
		thingprincipalattachment.Setup,
		thingtype.Setup,
		topicrule.Setup,
		channel.Setup,
		recordingconfiguration.Setup,
		clusterkafka.Setup,
		configuration.Setup,
		datasourcekendra.Setup,
		experience.Setup,
		index.Setup,
		querysuggestionsblocklist.Setup,
		thesaurus.Setup,
		keyspace.Setup,
		tablekeyspaces.Setup,
		streamkinesis.Setup,
		streamconsumer.Setup,
		applicationkinesisanalytics.Setup,
		applicationkinesisanalyticsv2.Setup,
		applicationsnapshot.Setup,
		streamkinesisvideo.Setup,
		aliaskms.Setup,
		ciphertext.Setup,
		externalkey.Setup,
		grant.Setup,
		key.Setup,
		replicaexternalkey.Setup,
		replicakey.Setup,
		datalakesettings.Setup,
		permissions.Setup,
		resourcelakeformation.Setup,
		aliaslambda.Setup,
		codesigningconfig.Setup,
		eventsourcemapping.Setup,
		functionlambda.Setup,
		functioneventinvokeconfig.Setup,
		functionurl.Setup,
		invocation.Setup,
		layerversion.Setup,
		layerversionpermission.Setup,
		permissionlambda.Setup,
		provisionedconcurrencyconfig.Setup,
		bot.Setup,
		botalias.Setup,
		intent.Setup,
		slottype.Setup,
		association.Setup,
		licenseconfiguration.Setup,
		bucket.Setup,
		certificatelightsail.Setup,
		containerservice.Setup,
		disk.Setup,
		diskattachment.Setup,
		domainlightsail.Setup,
		domainentry.Setup,
		instancelightsail.Setup,
		instancepublicports.Setup,
		keypairlightsail.Setup,
		lblightsail.Setup,
		lbattachment.Setup,
		lbcertificate.Setup,
		lbstickinesspolicy.Setup,
		staticip.Setup,
		staticipattachment.Setup,
		geofencecollection.Setup,
		placeindex.Setup,
		routecalculator.Setup,
		tracker.Setup,
		trackerassociation.Setup,
		accountmacie2.Setup,
		classificationjob.Setup,
		customdataidentifier.Setup,
		findingsfilter.Setup,
		invitationacceptermacie2.Setup,
		membermacie2.Setup,
		queuemediaconvert.Setup,
		channelmedialive.Setup,
		input.Setup,
		inputsecuritygroup.Setup,
		multiplex.Setup,
		channelmediapackage.Setup,
		container.Setup,
		containerpolicy.Setup,
		acl.Setup,
		clustermemorydb.Setup,
		parametergroupmemorydb.Setup,
		snapshot.Setup,
		subnetgroupmemorydb.Setup,
		broker.Setup,
		configurationmq.Setup,
		clusterneptune.Setup,
		clusterendpoint.Setup,
		clusterinstanceneptune.Setup,
		clusterparametergroupneptune.Setup,
		clustersnapshotneptune.Setup,
		eventsubscriptionneptune.Setup,
		globalclusterneptune.Setup,
		parametergroupneptune.Setup,
		subnetgroupneptune.Setup,
		firewall.Setup,
		firewallpolicy.Setup,
		loggingconfiguration.Setup,
		rulegroup.Setup,
		attachmentaccepter.Setup,
		connectattachment.Setup,
		connectionnetworkmanager.Setup,
		corenetwork.Setup,
		customergatewayassociation.Setup,
		device.Setup,
		globalnetwork.Setup,
		link.Setup,
		linkassociation.Setup,
		site.Setup,
		transitgatewayconnectpeerassociation.Setup,
		transitgatewayregistration.Setup,
		vpcattachment.Setup,
		domainopensearch.Setup,
		domainpolicyopensearch.Setup,
		domainsamloptionsopensearch.Setup,
		applicationopsworks.Setup,
		customlayer.Setup,
		ecsclusterlayer.Setup,
		ganglialayer.Setup,
		haproxylayer.Setup,
		instanceopsworks.Setup,
		javaapplayer.Setup,
		memcachedlayer.Setup,
		mysqllayer.Setup,
		nodejsapplayer.Setup,
		permissionopsworks.Setup,
		phpapplayer.Setup,
		railsapplayer.Setup,
		rdsdbinstance.Setup,
		stackopsworks.Setup,
		staticweblayer.Setup,
		userprofile.Setup,
		accountorganizations.Setup,
		delegatedadministrator.Setup,
		organization.Setup,
		organizationalunit.Setup,
		policyorganizations.Setup,
		policyattachmentorganizations.Setup,
		apppinpoint.Setup,
		smschannel.Setup,
		providerconfig.Setup,
		ledger.Setup,
		streamqldb.Setup,
		groupquicksight.Setup,
		userquicksight.Setup,
		resourceassociation.Setup,
		resourceshare.Setup,
		clusterrds.Setup,
		clusteractivitystream.Setup,
		clusterendpointrds.Setup,
		clusterinstancerds.Setup,
		clusterparametergrouprds.Setup,
		clusterroleassociation.Setup,
		clustersnapshotrds.Setup,
		dbinstanceautomatedbackupsreplication.Setup,
		dbsnapshotcopy.Setup,
		eventsubscriptionrds.Setup,
		globalclusterrds.Setup,
		instancerds.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergrouprds.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		securitygrouprds.Setup,
		snapshotrds.Setup,
		subnetgrouprds.Setup,
		authenticationprofile.Setup,
		clusterredshift.Setup,
		eventsubscriptionredshift.Setup,
		hsmclientcertificate.Setup,
		hsmconfiguration.Setup,
		parametergroupredshift.Setup,
		scheduledactionredshift.Setup,
		snapshotcopygrant.Setup,
		snapshotschedule.Setup,
		snapshotscheduleassociation.Setup,
		subnetgroupredshift.Setup,
		usagelimit.Setup,
		groupresourcegroups.Setup,
		profile.Setup,
		delegationset.Setup,
		healthcheck.Setup,
		hostedzonednssec.Setup,
		record.Setup,
		resolverconfig.Setup,
		trafficpolicy.Setup,
		trafficpolicyinstance.Setup,
		vpcassociationauthorization.Setup,
		zone.Setup,
		clusterroute53recoverycontrolconfig.Setup,
		controlpanel.Setup,
		routingcontrol.Setup,
		safetyrule.Setup,
		cell.Setup,
		readinesscheck.Setup,
		recoverygroup.Setup,
		resourceset.Setup,
		endpointroute53resolver.Setup,
		ruleroute53resolver.Setup,
		ruleassociation.Setup,
		appmonitor.Setup,
		metricsdestination.Setup,
		buckets3.Setup,
		bucketaccelerateconfiguration.Setup,
		bucketacl.Setup,
		bucketanalyticsconfiguration.Setup,
		bucketcorsconfiguration.Setup,
		bucketintelligenttieringconfiguration.Setup,
		bucketinventory.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketlogging.Setup,
		bucketmetric.Setup,
		bucketnotification.Setup,
		bucketobject.Setup,
		bucketobjectlockconfiguration.Setup,
		bucketownershipcontrols.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		bucketreplicationconfiguration.Setup,
		bucketrequestpaymentconfiguration.Setup,
		bucketserversideencryptionconfiguration.Setup,
		bucketversioning.Setup,
		bucketwebsiteconfiguration.Setup,
		object.Setup,
		objectcopy.Setup,
		accesspoints3control.Setup,
		accesspointpolicy.Setup,
		accountpublicaccessblock.Setup,
		multiregionaccesspoint.Setup,
		multiregionaccesspointpolicy.Setup,
		objectlambdaaccesspoint.Setup,
		objectlambdaaccesspointpolicy.Setup,
		storagelensconfiguration.Setup,
		appsagemaker.Setup,
		appimageconfig.Setup,
		coderepository.Setup,
		devicesagemaker.Setup,
		devicefleet.Setup,
		domainsagemaker.Setup,
		endpointconfiguration.Setup,
		featuregroup.Setup,
		imagesagemaker.Setup,
		imageversion.Setup,
		modelsagemaker.Setup,
		modelpackagegroup.Setup,
		modelpackagegrouppolicy.Setup,
		notebookinstance.Setup,
		notebookinstancelifecycleconfiguration.Setup,
		servicecatalogportfoliostatus.Setup,
		space.Setup,
		studiolifecycleconfig.Setup,
		userprofilesagemaker.Setup,
		workforce.Setup,
		workteam.Setup,
		schedulescheduler.Setup,
		schedulegroup.Setup,
		discoverer.Setup,
		registryschemas.Setup,
		schemaschemas.Setup,
		secret.Setup,
		secretpolicy.Setup,
		secretrotation.Setup,
		secretversion.Setup,
		accountsecurityhub.Setup,
		actiontarget.Setup,
		findingaggregator.Setup,
		insight.Setup,
		inviteaccepter.Setup,
		membersecurityhub.Setup,
		productsubscription.Setup,
		standardssubscription.Setup,
		cloudformationstack.Setup,
		budgetresourceassociation.Setup,
		constraint.Setup,
		portfolio.Setup,
		portfolioshare.Setup,
		principalportfolioassociation.Setup,
		product.Setup,
		productportfolioassociation.Setup,
		provisioningartifact.Setup,
		serviceaction.Setup,
		tagoption.Setup,
		tagoptionresourceassociation.Setup,
		httpnamespace.Setup,
		privatednsnamespace.Setup,
		publicdnsnamespace.Setup,
		serviceservicediscovery.Setup,
		servicequota.Setup,
		activereceiptruleset.Setup,
		configurationset.Setup,
		domaindkim.Setup,
		domainidentity.Setup,
		domainmailfrom.Setup,
		emailidentity.Setup,
		eventdestination.Setup,
		identitynotificationtopic.Setup,
		identitypolicy.Setup,
		receiptfilter.Setup,
		receiptrule.Setup,
		receiptruleset.Setup,
		template.Setup,
		configurationsetsesv2.Setup,
		configurationseteventdestination.Setup,
		dedicatedippool.Setup,
		emailidentitysesv2.Setup,
		emailidentityfeedbackattributes.Setup,
		emailidentitymailfromattributes.Setup,
		activity.Setup,
		statemachine.Setup,
		signingjob.Setup,
		signingprofile.Setup,
		signingprofilepermission.Setup,
		domainsimpledb.Setup,
		platformapplication.Setup,
		smspreferences.Setup,
		topic.Setup,
		topicpolicy.Setup,
		topicsubscription.Setup,
		queuesqs.Setup,
		queuepolicy.Setup,
		queueredriveallowpolicy.Setup,
		queueredrivepolicy.Setup,
		activation.Setup,
		associationssm.Setup,
		defaultpatchbaseline.Setup,
		document.Setup,
		maintenancewindow.Setup,
		maintenancewindowtarget.Setup,
		maintenancewindowtask.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		patchgroup.Setup,
		resourcedatasync.Setup,
		servicesetting.Setup,
		accountassignment.Setup,
		managedpolicyattachment.Setup,
		permissionset.Setup,
		permissionsetinlinepolicy.Setup,
		domainswf.Setup,
		databasetimestreamwrite.Setup,
		tabletimestreamwrite.Setup,
		languagemodel.Setup,
		vocabularytranscribe.Setup,
		vocabularyfilter.Setup,
		server.Setup,
		sshkey.Setup,
		tagtransfer.Setup,
		usertransfer.Setup,
		workflowtransfer.Setup,
		networkperformancemetricsubscription.Setup,
		bytematchset.Setup,
		geomatchset.Setup,
		ipset.Setup,
		ratebasedrule.Setup,
		regexmatchset.Setup,
		regexpatternset.Setup,
		rulewaf.Setup,
		sizeconstraintset.Setup,
		sqlinjectionmatchset.Setup,
		webacl.Setup,
		xssmatchset.Setup,
		bytematchsetwafregional.Setup,
		geomatchsetwafregional.Setup,
		ipsetwafregional.Setup,
		ratebasedrulewafregional.Setup,
		regexmatchsetwafregional.Setup,
		regexpatternsetwafregional.Setup,
		rulewafregional.Setup,
		sizeconstraintsetwafregional.Setup,
		sqlinjectionmatchsetwafregional.Setup,
		webaclwafregional.Setup,
		xssmatchsetwafregional.Setup,
		ipsetwafv2.Setup,
		regexpatternsetwafv2.Setup,
		directoryworkspaces.Setup,
		ipgroup.Setup,
		encryptionconfig.Setup,
		groupxray.Setup,
		samplingrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
