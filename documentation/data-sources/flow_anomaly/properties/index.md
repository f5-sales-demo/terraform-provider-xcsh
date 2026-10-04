---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_flow_anomaly."
xcsh_docs: {"aliases": ["flow anomaly"], "body_bytes": 2984, "body_sha256": "sha256:e8d00f7fc72b56ab4ed83713332394edecea35e6c3199bf95295489f9f03b76f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:flow_anomaly:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:flow_anomaly:reference", "parent_id": "xcsh-docs:data-sources:flow_anomaly:fundamentals", "path": "documentation/data-sources/flow_anomaly/properties/index.md", "product": "distributed-cloud", "provider_name": "flow_anomaly", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1001030133000121-1133022331020113-1013220120013233-1103310331100231-0302200102302313-0033012100301122-3101322032301330-3222201320301020", "registry_path": "docs/guides/data-sources--flow_anomaly--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:flow_anomaly:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:flow_anomaly:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:flow_anomaly:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:flow_anomaly:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["last enabled time"], "anchor": "schema-last_enabled_time", "description": "Last enabled time for flow anomaly add on service.", "document_id": "xcsh-docs:data-sources:flow_anomaly:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["last_enabled_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the FlowAnomaly to look up.", "document_id": "xcsh-docs:data-sources:flow_anomaly:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the FlowAnomaly.", "document_id": "xcsh-docs:data-sources:flow_anomaly:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["service state", "succeeded", "success", "successful"], "anchor": "schema-service_state", "description": "State of a service default state initiated subscription request and is pending to activate (requested). Successfully subscribed service subscription request ended up in error state. Possible values are `AS_NONE`, `AS_PENDING`, `AS_SUBSCRIBED`, `AS_ERROR`. Defaults to `AS_NONE`.", "document_id": "xcsh-docs:data-sources:flow_anomaly:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service_state"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/flow_anomaly/properties/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Property reference for xcsh_flow_anomaly.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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
