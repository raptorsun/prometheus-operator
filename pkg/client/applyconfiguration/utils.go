// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	v1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	v1beta1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1beta1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1alpha1"
	monitoringv1beta1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=monitoring.coreos.com, Version=v1
	case v1.SchemeGroupVersion.WithKind("AlertingSpec"):
		return &monitoringv1.AlertingSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Alertmanager"):
		return &monitoringv1.AlertmanagerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AlertmanagerConfigMatcherStrategy"):
		return &monitoringv1.AlertmanagerConfigMatcherStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AlertmanagerConfiguration"):
		return &monitoringv1.AlertmanagerConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AlertmanagerEndpoints"):
		return &monitoringv1.AlertmanagerEndpointsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AlertmanagerGlobalConfig"):
		return &monitoringv1.AlertmanagerGlobalConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AlertmanagerSpec"):
		return &monitoringv1.AlertmanagerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AlertmanagerStatus"):
		return &monitoringv1.AlertmanagerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AlertmanagerWebSpec"):
		return &monitoringv1.AlertmanagerWebSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("APIServerConfig"):
		return &monitoringv1.APIServerConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ArbitraryFSAccessThroughSMsConfig"):
		return &monitoringv1.ArbitraryFSAccessThroughSMsConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Argument"):
		return &monitoringv1.ArgumentApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AttachMetadata"):
		return &monitoringv1.AttachMetadataApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Authorization"):
		return &monitoringv1.AuthorizationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AzureAD"):
		return &monitoringv1.AzureADApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AzureOAuth"):
		return &monitoringv1.AzureOAuthApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BasicAuth"):
		return &monitoringv1.BasicAuthApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CommonPrometheusFields"):
		return &monitoringv1.CommonPrometheusFieldsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Condition"):
		return &monitoringv1.ConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CoreV1TopologySpreadConstraint"):
		return &monitoringv1.CoreV1TopologySpreadConstraintApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EmbeddedObjectMetadata"):
		return &monitoringv1.EmbeddedObjectMetadataApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EmbeddedPersistentVolumeClaim"):
		return &monitoringv1.EmbeddedPersistentVolumeClaimApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Endpoint"):
		return &monitoringv1.EndpointApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Exemplars"):
		return &monitoringv1.ExemplarsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GlobalSMTPConfig"):
		return &monitoringv1.GlobalSMTPConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HostAlias"):
		return &monitoringv1.HostAliasApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HostPort"):
		return &monitoringv1.HostPortApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HTTPConfig"):
		return &monitoringv1.HTTPConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ManagedIdentity"):
		return &monitoringv1.ManagedIdentityApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MetadataConfig"):
		return &monitoringv1.MetadataConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NamespaceSelector"):
		return &monitoringv1.NamespaceSelectorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OAuth2"):
		return &monitoringv1.OAuth2ApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ObjectReference"):
		return &monitoringv1.ObjectReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodMetricsEndpoint"):
		return &monitoringv1.PodMetricsEndpointApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodMetricsEndpointTLSConfig"):
		return &monitoringv1.PodMetricsEndpointTLSConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodMonitor"):
		return &monitoringv1.PodMonitorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodMonitorSpec"):
		return &monitoringv1.PodMonitorSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Probe"):
		return &monitoringv1.ProbeApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProberSpec"):
		return &monitoringv1.ProberSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProbeSpec"):
		return &monitoringv1.ProbeSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProbeTargetIngress"):
		return &monitoringv1.ProbeTargetIngressApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProbeTargets"):
		return &monitoringv1.ProbeTargetsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProbeTargetStaticConfig"):
		return &monitoringv1.ProbeTargetStaticConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ProbeTLSConfig"):
		return &monitoringv1.ProbeTLSConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Prometheus"):
		return &monitoringv1.PrometheusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PrometheusRule"):
		return &monitoringv1.PrometheusRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PrometheusRuleExcludeConfig"):
		return &monitoringv1.PrometheusRuleExcludeConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PrometheusRuleSpec"):
		return &monitoringv1.PrometheusRuleSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PrometheusSpec"):
		return &monitoringv1.PrometheusSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PrometheusStatus"):
		return &monitoringv1.PrometheusStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PrometheusTracingConfig"):
		return &monitoringv1.PrometheusTracingConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PrometheusWebSpec"):
		return &monitoringv1.PrometheusWebSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("QuerySpec"):
		return &monitoringv1.QuerySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("QueueConfig"):
		return &monitoringv1.QueueConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RelabelConfig"):
		return &monitoringv1.RelabelConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RemoteReadSpec"):
		return &monitoringv1.RemoteReadSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RemoteWriteSpec"):
		return &monitoringv1.RemoteWriteSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Rule"):
		return &monitoringv1.RuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RuleGroup"):
		return &monitoringv1.RuleGroupApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Rules"):
		return &monitoringv1.RulesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RulesAlert"):
		return &monitoringv1.RulesAlertApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SafeAuthorization"):
		return &monitoringv1.SafeAuthorizationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SafeTLSConfig"):
		return &monitoringv1.SafeTLSConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SecretOrConfigMap"):
		return &monitoringv1.SecretOrConfigMapApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceMonitor"):
		return &monitoringv1.ServiceMonitorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceMonitorSpec"):
		return &monitoringv1.ServiceMonitorSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ShardStatus"):
		return &monitoringv1.ShardStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Sigv4"):
		return &monitoringv1.Sigv4ApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StorageSpec"):
		return &monitoringv1.StorageSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ThanosRuler"):
		return &monitoringv1.ThanosRulerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ThanosRulerSpec"):
		return &monitoringv1.ThanosRulerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ThanosRulerStatus"):
		return &monitoringv1.ThanosRulerStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ThanosSpec"):
		return &monitoringv1.ThanosSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TLSConfig"):
		return &monitoringv1.TLSConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TopologySpreadConstraint"):
		return &monitoringv1.TopologySpreadConstraintApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TSDBSpec"):
		return &monitoringv1.TSDBSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WebConfigFileFields"):
		return &monitoringv1.WebConfigFileFieldsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WebHTTPConfig"):
		return &monitoringv1.WebHTTPConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WebHTTPHeaders"):
		return &monitoringv1.WebHTTPHeadersApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WebTLSConfig"):
		return &monitoringv1.WebTLSConfigApplyConfiguration{}

		// Group=monitoring.coreos.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("AlertmanagerConfig"):
		return &monitoringv1alpha1.AlertmanagerConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AlertmanagerConfigSpec"):
		return &monitoringv1alpha1.AlertmanagerConfigSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AttachMetadata"):
		return &monitoringv1alpha1.AttachMetadataApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("AzureSDConfig"):
		return &monitoringv1alpha1.AzureSDConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ConsulSDConfig"):
		return &monitoringv1alpha1.ConsulSDConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DayOfMonthRange"):
		return &monitoringv1alpha1.DayOfMonthRangeApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DiscordConfig"):
		return &monitoringv1alpha1.DiscordConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DNSSDConfig"):
		return &monitoringv1alpha1.DNSSDConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EC2Filter"):
		return &monitoringv1alpha1.EC2FilterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EC2SDConfig"):
		return &monitoringv1alpha1.EC2SDConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EmailConfig"):
		return &monitoringv1alpha1.EmailConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FileSDConfig"):
		return &monitoringv1alpha1.FileSDConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GCESDConfig"):
		return &monitoringv1alpha1.GCESDConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HTTPConfig"):
		return &monitoringv1alpha1.HTTPConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HTTPSDConfig"):
		return &monitoringv1alpha1.HTTPSDConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("InhibitRule"):
		return &monitoringv1alpha1.InhibitRuleApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("K8SSelectorConfig"):
		return &monitoringv1alpha1.K8SSelectorConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KeyValue"):
		return &monitoringv1alpha1.KeyValueApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KubernetesSDConfig"):
		return &monitoringv1alpha1.KubernetesSDConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Matcher"):
		return &monitoringv1alpha1.MatcherApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("MSTeamsConfig"):
		return &monitoringv1alpha1.MSTeamsConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("MuteTimeInterval"):
		return &monitoringv1alpha1.MuteTimeIntervalApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("NamespaceDiscovery"):
		return &monitoringv1alpha1.NamespaceDiscoveryApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OpenStackSDConfig"):
		return &monitoringv1alpha1.OpenStackSDConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OpsGenieConfig"):
		return &monitoringv1alpha1.OpsGenieConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OpsGenieConfigResponder"):
		return &monitoringv1alpha1.OpsGenieConfigResponderApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PagerDutyConfig"):
		return &monitoringv1alpha1.PagerDutyConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PagerDutyImageConfig"):
		return &monitoringv1alpha1.PagerDutyImageConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PagerDutyLinkConfig"):
		return &monitoringv1alpha1.PagerDutyLinkConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PrometheusAgent"):
		return &monitoringv1alpha1.PrometheusAgentApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PrometheusAgentSpec"):
		return &monitoringv1alpha1.PrometheusAgentSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ProxyConfig"):
		return &monitoringv1alpha1.ProxyConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PushoverConfig"):
		return &monitoringv1alpha1.PushoverConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Receiver"):
		return &monitoringv1alpha1.ReceiverApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Route"):
		return &monitoringv1alpha1.RouteApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ScrapeConfig"):
		return &monitoringv1alpha1.ScrapeConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ScrapeConfigSpec"):
		return &monitoringv1alpha1.ScrapeConfigSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SlackAction"):
		return &monitoringv1alpha1.SlackActionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SlackConfig"):
		return &monitoringv1alpha1.SlackConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SlackConfirmationField"):
		return &monitoringv1alpha1.SlackConfirmationFieldApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SlackField"):
		return &monitoringv1alpha1.SlackFieldApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SNSConfig"):
		return &monitoringv1alpha1.SNSConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("StaticConfig"):
		return &monitoringv1alpha1.StaticConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TelegramConfig"):
		return &monitoringv1alpha1.TelegramConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TimeInterval"):
		return &monitoringv1alpha1.TimeIntervalApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TimeRange"):
		return &monitoringv1alpha1.TimeRangeApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("VictorOpsConfig"):
		return &monitoringv1alpha1.VictorOpsConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("WebexConfig"):
		return &monitoringv1alpha1.WebexConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("WebhookConfig"):
		return &monitoringv1alpha1.WebhookConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("WeChatConfig"):
		return &monitoringv1alpha1.WeChatConfigApplyConfiguration{}

		// Group=monitoring.coreos.com, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("AlertmanagerConfig"):
		return &monitoringv1beta1.AlertmanagerConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AlertmanagerConfigSpec"):
		return &monitoringv1beta1.AlertmanagerConfigSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("DayOfMonthRange"):
		return &monitoringv1beta1.DayOfMonthRangeApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("DiscordConfig"):
		return &monitoringv1beta1.DiscordConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("EmailConfig"):
		return &monitoringv1beta1.EmailConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("HTTPConfig"):
		return &monitoringv1beta1.HTTPConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("InhibitRule"):
		return &monitoringv1beta1.InhibitRuleApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KeyValue"):
		return &monitoringv1beta1.KeyValueApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Matcher"):
		return &monitoringv1beta1.MatcherApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MSTeamsConfig"):
		return &monitoringv1beta1.MSTeamsConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("OpsGenieConfig"):
		return &monitoringv1beta1.OpsGenieConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("OpsGenieConfigResponder"):
		return &monitoringv1beta1.OpsGenieConfigResponderApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PagerDutyConfig"):
		return &monitoringv1beta1.PagerDutyConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PagerDutyImageConfig"):
		return &monitoringv1beta1.PagerDutyImageConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PagerDutyLinkConfig"):
		return &monitoringv1beta1.PagerDutyLinkConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PushoverConfig"):
		return &monitoringv1beta1.PushoverConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Receiver"):
		return &monitoringv1beta1.ReceiverApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Route"):
		return &monitoringv1beta1.RouteApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("SecretKeySelector"):
		return &monitoringv1beta1.SecretKeySelectorApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("SlackAction"):
		return &monitoringv1beta1.SlackActionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("SlackConfig"):
		return &monitoringv1beta1.SlackConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("SlackConfirmationField"):
		return &monitoringv1beta1.SlackConfirmationFieldApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("SlackField"):
		return &monitoringv1beta1.SlackFieldApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("SNSConfig"):
		return &monitoringv1beta1.SNSConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("TelegramConfig"):
		return &monitoringv1beta1.TelegramConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("TimeInterval"):
		return &monitoringv1beta1.TimeIntervalApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("TimePeriod"):
		return &monitoringv1beta1.TimePeriodApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("TimeRange"):
		return &monitoringv1beta1.TimeRangeApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("VictorOpsConfig"):
		return &monitoringv1beta1.VictorOpsConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WebexConfig"):
		return &monitoringv1beta1.WebexConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WebhookConfig"):
		return &monitoringv1beta1.WebhookConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WeChatConfig"):
		return &monitoringv1beta1.WeChatConfigApplyConfiguration{}

	}
	return nil
}
