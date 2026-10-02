---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_detection_rule."
xcsh_docs: {"aliases": ["bot detection rule"], "body_bytes": 10405, "body_sha256": "sha256:5141bb8b7e1792a0d823827bc903a6bc1ca8a09f279d00fd5170ed402aaa73b3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_detection_rule:reference", "parent_id": "xcsh-docs:data-sources:bot_detection_rule:fundamentals", "path": "documentation/data-sources/bot_detection_rule/properties/index.md", "product": "distributed-cloud", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1300000332213130-3231303130200131-1302120313303120-2230111331211100-2123003010122200-3002313113113213-0321013203213000-2322013230021313", "registry_path": "docs/guides/data-sources--bot_detection_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["bot detection rule configs per bot infras"], "anchor": "section", "description": "Rule configurations per Bot-Infras. Rule or policy definition", "document_id": "xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["bot_detection_rule_configs_per_bot_infras"], "syntax": "attribute", "type": "object"}, {"aliases": ["classification"], "anchor": "schema-classification", "description": "Classification. Classification or category data", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["classification"], "syntax": "attribute", "type": "list"}, {"aliases": ["cluster groups"], "anchor": "schema-cluster_groups", "description": "Cluster Groups. Cluster or grouping configuration", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_groups"], "syntax": "attribute", "type": "list"}, {"aliases": ["created at"], "anchor": "schema-created_at", "description": "Created At. Created at as per SAPI's database.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["created_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["last modified at"], "anchor": "schema-last_modified_at", "description": "Last modified at as per SAPI's database.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["last_modified_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["last modified by"], "anchor": "schema-last_modified_by", "description": "Last Modified By as per SAPI's database.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["last_modified_by"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the BotDetectionRule to look up.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the BotDetectionRule.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["rule name"], "anchor": "schema-rule_name", "description": "Rule Name. Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rule type"], "anchor": "schema-rule_type", "description": "Bot Detection Rule Type. Possible values are `BOT_DETECTION_RULE_TYPE_UNKNOWN`, `BOT_DETECTION_RULE_TYPE_ENFORCED_BLOCKING`, `BOT_DETECTION_RULE_TYPE_CONTROL_BLOCKING`. Defaults to `BOT_DETECTION_RULE_TYPE_UNKNOWN`.", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["traffic type"], "anchor": "schema-traffic_type", "description": "The type of traffic that is routed to and processed by this infrastructure (Web or Mobile). Only web traffic, including browser-based traffic from mobile devices, is routed through this Bot Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense SDK are routed.. Possible values are", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["traffic_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["version"], "anchor": "schema-version", "description": "Version. Version number or identifier", "document_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["version"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_detection_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_detection_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [bot_detection_rule_configs_per_bot_infras](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/): complete subsection reference.

<a id="schema-classification"></a>

### classification property

Type: `["list", "string"]`. Computed.

Classification. Classification or category data

<a id="schema-cluster_groups"></a>

### cluster_groups property

Type: `["list", "string"]`. Computed.

Cluster Groups. Cluster or grouping configuration

<a id="schema-created_at"></a>

### created_at property

Type: `"string"`. Computed.

Created At. Created at as per SAPI's database.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-last_modified_at"></a>

### last_modified_at property

Type: `"string"`. Computed.

Last modified at as per SAPI's database.

<a id="schema-last_modified_by"></a>

### last_modified_by property

Type: `"string"`. Computed.

Last Modified By as per SAPI's database.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the BotDetectionRule to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the BotDetectionRule.

<a id="schema-rule_name"></a>

### rule_name property

Type: `"string"`. Computed.

Rule Name. Human-readable name for the resource

<a id="schema-rule_type"></a>

### rule_type property

Type: `"string"`. Computed.

\[Enum:
BOT\_DETECTION\_RULE\_TYPE\_UNKNOWN|BOT\_DETECTION\_RULE\_TYPE\_ENFORCED\_BLOCKING|BOT\_DETECTION\_RULE\_TYPE\_CONTROL\_BLOCKING\]
Bot Detection Rule Type. Possible values are \`BOT\_DETECTION\_RULE\_TYPE\_UNKNOWN\`,
\`BOT\_DETECTION\_RULE\_TYPE\_ENFORCED\_BLOCKING\`,
\`BOT\_DETECTION\_RULE\_TYPE\_CONTROL\_BLOCKING\`. Defaults to
\`BOT\_DETECTION\_RULE\_TYPE\_UNKNOWN\`.

<a id="schema-traffic_type"></a>

### traffic_type property

Type: `"string"`. Computed.

\[Enum: WEB|MOBILE\] The type of traffic that is routed to and processed by this infrastructure (Web
or Mobile). Only web traffic, including browser-based traffic from mobile devices, is routed through
this Bot Defense infrastructure. Only mobile traffic from native mobile apps with the Bot Defense
SDK are routed.. Possible values are \`WEB\`, \`MOBILE\`. Defaults to \`WEB\`.

<a id="schema-version"></a>

### version property

Type: `"number"`. Computed.

Version. Version number or identifier

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-annotations) |
| `bot_detection_rule_configs_per_bot_infras` | [bot_detection_rule_configs_per_bot_infras](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/#section) |
| `bot_detection_rule_configs_per_bot_infras.bot_detection_rule_config` | [bot_detection_rule_configs_per_bot_infras.bot_detection_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/bot_detection_rule_config/#section) |
| `bot_detection_rule_configs_per_bot_infras.bot_detection_rule_config.mitigation` | [bot_detection_rule_configs_per_bot_infras.bot_detection_rule_config.mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/bot_detection_rule_config/#schema-bot_detection_rule_configs_per_bot_infras--bot_detection_rule_config--mitigation) |
| `bot_detection_rule_configs_per_bot_infras.bot_infra_id` | [bot_detection_rule_configs_per_bot_infras.bot_infra_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/#schema-bot_detection_rule_configs_per_bot_infras--bot_infra_id) |
| `bot_detection_rule_configs_per_bot_infras.bot_infra_name` | [bot_detection_rule_configs_per_bot_infras.bot_infra_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/#schema-bot_detection_rule_configs_per_bot_infras--bot_infra_name) |
| `bot_detection_rule_configs_per_bot_infras.bot_infrastructure_type` | [bot_detection_rule_configs_per_bot_infras.bot_infrastructure_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/#schema-bot_detection_rule_configs_per_bot_infras--bot_infrastructure_type) |
| `bot_detection_rule_configs_per_bot_infras.environment_type` | [bot_detection_rule_configs_per_bot_infras.environment_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/#schema-bot_detection_rule_configs_per_bot_infras--environment_type) |
| `bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config` | [bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/k8s_cluster_rule_config/#section) |
| `bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.in_used_k8s_cluster_rule_config` | [bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.in_used_k8s_cluster_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/k8s_cluster_rule_config/in_used_k8s_cluster_rule_config/#section) |
| `bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.sync_status_per_k8s_cluster` | [bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.sync_status_per_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/k8s_cluster_rule_config/sync_status_per_k8s_cluster/#section) |
| `bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.target_k8s_cluster_rule_config` | [bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.target_k8s_cluster_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/k8s_cluster_rule_config/target_k8s_cluster_rule_config/#section) |
| `bot_detection_rule_configs_per_bot_infras.region_rule_config` | [bot_detection_rule_configs_per_bot_infras.region_rule_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/region_rule_config/#section) |
| `bot_detection_rule_configs_per_bot_infras.region_rule_config.region_config` | [bot_detection_rule_configs_per_bot_infras.region_rule_config.region_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/region_rule_config/region_config/#section) |
| `classification` | [classification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-classification) |
| `cluster_groups` | [cluster_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-cluster_groups) |
| `created_at` | [created_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-created_at) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-labels) |
| `last_modified_at` | [last_modified_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-last_modified_at) |
| `last_modified_by` | [last_modified_by](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-last_modified_by) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-namespace) |
| `rule_name` | [rule_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-rule_name) |
| `rule_type` | [rule_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-rule_type) |
| `traffic_type` | [traffic_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-traffic_type) |
| `version` | [version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/#schema-version) |

## Next pages

- [bot_detection_rule_configs_per_bot_infras](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/bot_detection_rule_configs_per_bot_infras/)
- [xcsh_bot_detection_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/)
