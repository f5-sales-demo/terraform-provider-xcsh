---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_flow_anomaly."
xcsh_docs: {"aliases": [], "body_bytes": 2984, "body_sha256": "sha256:e8d00f7fc72b56ab4ed83713332394edecea35e6c3199bf95295489f9f03b76f", "child_ids": [], "collection_id": "xcsh-docs:data-sources:flow_anomaly:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:flow_anomaly:reference", "parent_id": "xcsh-docs:data-sources:flow_anomaly:fundamentals", "path": "documentation/data-sources/flow_anomaly/properties/index.md", "provider_name": "flow_anomaly", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/flow_anomaly/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_flow_anomaly.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_flow_anomaly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

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

<a id="schema-last_enabled_time"></a>

### last_enabled_time property

Type: `"string"`. Computed.

Last enabled time for flow anomaly add on service.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the FlowAnomaly to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the FlowAnomaly.

<a id="schema-service_state"></a>

### service_state property

Type: `"string"`. Computed.

\[Enum: AS\_NONE|AS\_PENDING|AS\_SUBSCRIBED|AS\_ERROR\] State of a service default state initiated
subscription request and is pending to activate (requested). Successfully subscribed service
subscription request ended up in error state. Possible values are \`AS\_NONE\`, \`AS\_PENDING\`,
\`AS\_SUBSCRIBED\`, \`AS\_ERROR\`. Defaults to \`AS\_NONE\`.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/#schema-labels) |
| `last_enabled_time` | [last_enabled_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/#schema-last_enabled_time) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/#schema-namespace) |
| `service_state` | [service_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/#schema-service_state) |

## Next pages

- [xcsh_flow_anomaly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/)
