---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_detection_rule."
xcsh_docs: {"aliases": [], "body_bytes": 8755, "body_sha256": "sha256:484995568c7a2fca10d03bb46134bb537ae01c645380cb8119fefd749bafe22b", "canonical_id": "xcsh-docs:data-sources:bot_detection_rule:reference", "child_ids": ["xcsh-docs:data-sources:bot_detection_rule:properties:bot_detection_rule_configs_per_bot_infras"], "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_detection_rule:reference", "parent_id": "xcsh-docs:data-sources:bot_detection_rule:fundamentals", "path": "docs/guides/data-sources--bot_detection_rule--reference.md", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_detection_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [bot_detection_rule_configs_per_bot_infras](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras.md): complete subsection reference.

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
| `annotations` | [annotations](data-sources--bot_detection_rule--reference.md#schema-annotations) |
| `bot_detection_rule_configs_per_bot_infras` | [bot_detection_rule_configs_per_bot_infras](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras.md#section) |
| `bot_detection_rule_configs_per_bot_infras.bot_detection_rule_config` | [bot_detection_rule_configs_per_bot_infras.bot_detection_rule_config](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras--bot_detection_rule_config.md#section) |
| `bot_detection_rule_configs_per_bot_infras.bot_detection_rule_config.mitigation` | [bot_detection_rule_configs_per_bot_infras.bot_detection_rule_config.mitigation](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras--bot_detection_rule_config.md#schema-bot_detection_rule_configs_per_bot_infras--bot_detection_rule_config--mitigation) |
| `bot_detection_rule_configs_per_bot_infras.bot_infra_id` | [bot_detection_rule_configs_per_bot_infras.bot_infra_id](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras.md#schema-bot_detection_rule_configs_per_bot_infras--bot_infra_id) |
| `bot_detection_rule_configs_per_bot_infras.bot_infra_name` | [bot_detection_rule_configs_per_bot_infras.bot_infra_name](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras.md#schema-bot_detection_rule_configs_per_bot_infras--bot_infra_name) |
| `bot_detection_rule_configs_per_bot_infras.bot_infrastructure_type` | [bot_detection_rule_configs_per_bot_infras.bot_infrastructure_type](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras.md#schema-bot_detection_rule_configs_per_bot_infras--bot_infrastructure_type) |
| `bot_detection_rule_configs_per_bot_infras.environment_type` | [bot_detection_rule_configs_per_bot_infras.environment_type](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras.md#schema-bot_detection_rule_configs_per_bot_infras--environment_type) |
| `bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config` | [bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras--k8s_cluster_rule_config.md#section) |
| `bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.in_used_k8s_cluster_rule_config` | [bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.in_used_k8s_cluster_rule_config](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras--k8s_cluster_rule_config--in_used_k8s_cluster_rule_config.md#section) |
| `bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.sync_status_per_k8s_cluster` | [bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.sync_status_per_k8s_cluster](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras--k8s_cluster_rule_config--sync_status_per_k8s_cluster.md#section) |
| `bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.target_k8s_cluster_rule_config` | [bot_detection_rule_configs_per_bot_infras.k8s_cluster_rule_config.target_k8s_cluster_rule_config](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras--k8s_cluster_rule_config--target_k8s_cluster_rule_config.md#section) |
| `bot_detection_rule_configs_per_bot_infras.region_rule_config` | [bot_detection_rule_configs_per_bot_infras.region_rule_config](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras--region_rule_config.md#section) |
| `bot_detection_rule_configs_per_bot_infras.region_rule_config.region_config` | [bot_detection_rule_configs_per_bot_infras.region_rule_config.region_config](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras--region_rule_config--region_config.md#section) |
| `classification` | [classification](data-sources--bot_detection_rule--reference.md#schema-classification) |
| `cluster_groups` | [cluster_groups](data-sources--bot_detection_rule--reference.md#schema-cluster_groups) |
| `created_at` | [created_at](data-sources--bot_detection_rule--reference.md#schema-created_at) |
| `description` | [description](data-sources--bot_detection_rule--reference.md#schema-description) |
| `id` | [id](data-sources--bot_detection_rule--reference.md#schema-id) |
| `labels` | [labels](data-sources--bot_detection_rule--reference.md#schema-labels) |
| `last_modified_at` | [last_modified_at](data-sources--bot_detection_rule--reference.md#schema-last_modified_at) |
| `last_modified_by` | [last_modified_by](data-sources--bot_detection_rule--reference.md#schema-last_modified_by) |
| `name` | [name](data-sources--bot_detection_rule--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--bot_detection_rule--reference.md#schema-namespace) |
| `rule_name` | [rule_name](data-sources--bot_detection_rule--reference.md#schema-rule_name) |
| `rule_type` | [rule_type](data-sources--bot_detection_rule--reference.md#schema-rule_type) |
| `traffic_type` | [traffic_type](data-sources--bot_detection_rule--reference.md#schema-traffic_type) |
| `version` | [version](data-sources--bot_detection_rule--reference.md#schema-version) |

## Next pages

- [bot_detection_rule_configs_per_bot_infras](data-sources--bot_detection_rule--properties--bot_detection_rule_configs_per_bot_infras.md)
- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md)
