---
page_title: "bot_detection_rule_configs_per_bot_infras"
subcategory: ""
description: "bot_detection_rule_configs_per_bot_infras for xcsh_bot_detection_rule."
xcsh_docs: {"aliases": [], "body_bytes": 3795, "body_sha256": "sha256:1d134a451c0719bee119ad1bed4655ed9214946f97aff0ca81f71efe189f0b87", "child_ids": ["xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras:bot_detection_rule_config", "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras:k8s_cluster_rule_config", "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras:region_rule_config"], "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras", "parent_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "path": "documentation/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/index.md", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["bot_detection_rule_configs_per_bot_infras"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_detection_rule_configs_per_bot_infras for xcsh_bot_detection_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_detection_rule_configs_per_bot_infras

Breadcrumbs:

- [xcsh_bot_detection_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/)
- bot_detection_rule_configs_per_bot_infras

<a id="section"></a>

Type: `"list"`. Computed.

Rule configurations per Bot-Infras. Rule or policy definition

## Direct properties

- [bot_detection_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/bot_detection_rule_config/): complete subsection reference.

<a id="schema-bot_detection_rule_configs_per_bot_infras--bot_infra_id"></a>

### bot_infra_id property

Type: `"string"`. Computed.

Bot Infrastructure ID as per SAPI's database.

<a id="schema-bot_detection_rule_configs_per_bot_infras--bot_infra_name"></a>

### bot_infra_name property

Type: `"string"`. Computed.

Name of the bot infrastructure as per SAPI's database.

<a id="schema-bot_detection_rule_configs_per_bot_infras--bot_infrastructure_type"></a>

### bot_infrastructure_type property

Type: `"string"`. Computed.

\[Enum:
BOT\_INFRA\_TYPE\_UNKNOWN|BOT\_INFRA\_TYPE\_CLOUD\_HOSTED|BOT\_INFRA\_TYPE\_HOSTED|BOT\_INFRA\_TYPE\_ON\_PREM|BOT\_INFRA\_TYPE\_K8S\_CLUSTER\]
Type of the Bot Infrastructure - BOT\_INFRA\_TYPE\_UNKNOWN: Unknown -
BOT\_INFRA\_TYPE\_CLOUD\_HOSTED: F5 Cloud Hosted - BOT\_INFRA\_TYPE\_HOSTED: F5 Hosted -
BOT\_INFRA\_TYPE\_ON\_PREM: F5 On Premises - BOT\_INFRA\_TYPE\_K8S\_CLUSTER: Kubernetes Cluster.
Possible values are \`BOT\_INFRA\_TYPE\_UNKNOWN\`, \`BOT\_INFRA\_TYPE\_CLOUD\_HOSTED\`,
\`BOT\_INFRA\_TYPE\_HOSTED\`, \`BOT\_INFRA\_TYPE\_ON\_PREM\`, \`BOT\_INFRA\_TYPE\_K8S\_CLUSTER\`.
Defaults to \`BOT\_INFRA\_TYPE\_UNKNOWN\`.

<a id="schema-bot_detection_rule_configs_per_bot_infras--environment_type"></a>

### environment_type property

Type: `"string"`. Computed.

\[Enum: PRODUCTION|TESTING\] Identifies the environment as either Production or Testing. Production
environments have two infrastructure regions in an Active-Active configuration where traffic is
routed equally between the two regions. Test environments have a single infrastructure region.
Possible values are \`PRODUCTION\`, \`TESTING\`. Defaults to \`PRODUCTION\`.

- [k8s_cluster_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/k8s_cluster_rule_config/): complete subsection reference.

- [region_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/region_rule_config/): complete subsection reference.

## Next pages

- [bot_detection_rule_configs_per_bot_infras.bot_detection_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/bot_detection_rule_config/)
- [bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/k8s_cluster_rule_config/)
- [bot_detection_rule_configs_per_bot_infras.region_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/region_rule_config/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/)
- [xcsh_bot_detection_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/)
