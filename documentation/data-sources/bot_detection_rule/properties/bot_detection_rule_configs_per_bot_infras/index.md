---
page_title: "bot_detection_rule_configs_per_bot_infras"
subcategory: ""
description: "Rule configurations per Bot-Infras. Rule or policy definition"
xcsh_docs: {"aliases": ["bot detection rule configs per bot infras"], "body_bytes": 3894, "body_sha256": "sha256:b23ebbbbe74551d02283c9ba7892b4f502642cdf466d6f1bb5f21744b9144fe6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras:bot_detection_rule_config", "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras:k8s_cluster_rule_config", "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras:region_rule_config"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras", "parent_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "path": "documentation/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/index.md", "product": "distributed-cloud", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3301101230033201-2203210111313013-1330230003003131-1221203300220202-0022120320333020-0222312221030022-3133302320311110-0301113011312221", "registry_path": "docs/guides/data-sources--bot_detection_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_detection_rule_configs_per_bot_infras"], "schema_version": 1, "sections": [{"aliases": ["bot detection rule config"], "anchor": "section", "description": "Rule configuration. Rule configuration.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras:bot_detection_rule_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_detection_rule_configs_per_bot_infras", "bot_detection_rule_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot infra id"], "anchor": "schema-bot_detection_rule_configs_per_bot_infras--bot_infra_id", "description": "Bot Infrastructure ID as per SAPI's database.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_detection_rule_configs_per_bot_infras", "bot_infra_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot infra name"], "anchor": "schema-bot_detection_rule_configs_per_bot_infras--bot_infra_name", "description": "Name of the bot infrastructure as per SAPI's database.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_detection_rule_configs_per_bot_infras", "bot_infra_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot infrastructure type"], "anchor": "schema-bot_detection_rule_configs_per_bot_infras--bot_infrastructure_type", "description": "Type of the Bot Infrastructure - BOT_INFRA_TYPE_UNKNOWN: Unknown - BOT_INFRA_TYPE_CLOUD_HOSTED: F5 Cloud Hosted - BOT_INFRA_TYPE_HOSTED: F5 Hosted - BOT_INFRA_TYPE_ON_PREM: F5 On Premises - BOT_INFRA_TYPE_K8S_CLUSTER: Kubernetes Cluster. Possible values are `BOT_INFRA_TYPE_UNKNOWN`, `BOT_INFRA_TYPE_CLOUD_HOSTED`,", "document_id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_detection_rule_configs_per_bot_infras", "bot_infrastructure_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["environment type"], "anchor": "schema-bot_detection_rule_configs_per_bot_infras--environment_type", "description": "Identifies the environment as either Production or Testing. Production environments have two infrastructure regions in an Active-Active configuration where traffic is routed equally between the two regions. Test environments have a single infrastructure region. Possible values are `PRODUCTION`, `TESTING`. Defaults to", "document_id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_detection_rule_configs_per_bot_infras", "environment_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["k8s cluster rule config"], "anchor": "section", "description": "Kubernetes Cluster Rule Config. Kubernetes cluster rule config.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras:k8s_cluster_rule_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_detection_rule_configs_per_bot_infras", "k8s_cluster_rule_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["region rule config"], "anchor": "section", "description": "Rule Config Per Region. Rule config per region.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras:region_rule_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_detection_rule_configs_per_bot_infras", "region_rule_config"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Rule configurations per Bot-Infras. Rule or policy definition", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
