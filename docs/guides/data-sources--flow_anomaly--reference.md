---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_flow_anomaly."
xcsh_docs: {"aliases": [], "body_bytes": 2458, "body_sha256": "sha256:2e6c3e3a04ec922ed723bd106691ac75fa4c4cd12d33fb56edc74ba8998acdf3", "canonical_id": "xcsh-docs:data-sources:flow_anomaly:reference", "child_ids": [], "collection_id": "xcsh-docs:data-sources:flow_anomaly:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:flow_anomaly:reference", "parent_id": "xcsh-docs:data-sources:flow_anomaly:fundamentals", "path": "docs/guides/data-sources--flow_anomaly--reference.md", "provider_name": "flow_anomaly", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/flow_anomaly/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_flow_anomaly.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_flow_anomaly](../data-sources/flow_anomaly.md)
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
| `annotations` | [annotations](data-sources--flow_anomaly--reference.md#schema-annotations) |
| `description` | [description](data-sources--flow_anomaly--reference.md#schema-description) |
| `id` | [id](data-sources--flow_anomaly--reference.md#schema-id) |
| `labels` | [labels](data-sources--flow_anomaly--reference.md#schema-labels) |
| `last_enabled_time` | [last_enabled_time](data-sources--flow_anomaly--reference.md#schema-last_enabled_time) |
| `name` | [name](data-sources--flow_anomaly--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--flow_anomaly--reference.md#schema-namespace) |
| `service_state` | [service_state](data-sources--flow_anomaly--reference.md#schema-service_state) |

## Next pages

- [xcsh_flow_anomaly](../data-sources/flow_anomaly.md)
