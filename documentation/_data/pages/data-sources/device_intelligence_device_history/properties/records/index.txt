---
page_title: "records"
subcategory: ""
description: "Records. List of activity records."
xcsh_docs: {"aliases": ["records"], "body_bytes": 2130, "body_sha256": "sha256:4f3d0cf6093e6fb68eb15379a050063ab87094f1f90db56e8605c900c8dc5775", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "path": "documentation/data-sources/device_intelligence_device_history/properties/records/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0103123322112321-1002002023332310-2322312220320002-3121012100220233-0020102331021200-0002012222320112-3133111021102023-1323303003123020", "registry_path": "docs/guides/data-sources--device_intelligence_device_history--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["records"], "schema_version": 1, "sections": [{"aliases": ["records action taken"], "anchor": "schema-records--action_taken", "description": "Action taken in response to the risk evaluation for this event.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "action_taken"], "syntax": "attribute", "type": "string"}, {"aliases": ["records detected signals"], "anchor": "schema-records--detected_signals", "description": "Risk or integrity signals observed for this event.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "detected_signals"], "syntax": "attribute", "type": "list"}, {"aliases": ["records endpoint label"], "anchor": "schema-records--endpoint_label", "description": "Human-readable label for the endpoint or action.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "endpoint_label"], "syntax": "attribute", "type": "string"}, {"aliases": ["records risk score"], "anchor": "schema-records--risk_score", "description": "Risk score assigned to this event (0–100).", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "risk_score"], "syntax": "attribute", "type": "number"}, {"aliases": ["records timestamp"], "anchor": "schema-records--timestamp", "description": "Event timestamp in RFC 3339 format with milliseconds (UTC).", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["records txn id"], "anchor": "schema-records--txn_id", "description": "Identifier of the associated transaction.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "txn_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["records url"], "anchor": "schema-records--url", "description": "Endpoint URL or host associated with the event.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/properties/records/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Records. List of activity records.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# records

Breadcrumbs:

- [xcsh_device_intelligence_device_history](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/)
- records

<a id="section"></a>

Type: `"list"`. Computed.

Records. List of activity records.

## Direct properties

<a id="schema-records--action_taken"></a>

### action_taken property

Type: `"string"`. Computed.

Action taken in response to the risk evaluation for this event.

<a id="schema-records--detected_signals"></a>

### detected_signals property

Type: `["list", "string"]`. Computed.

Risk or integrity signals observed for this event.

<a id="schema-records--endpoint_label"></a>

### endpoint_label property

Type: `"string"`. Computed.

Human-readable label for the endpoint or action.

<a id="schema-records--risk_score"></a>

### risk_score property

Type: `"number"`. Computed.

Risk score assigned to this event (0–100).

<a id="schema-records--timestamp"></a>

### timestamp property

Type: `"string"`. Computed.

Event timestamp in RFC 3339 format with milliseconds (UTC).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="schema-records--txn_id"></a>

### txn_id property

Type: `"string"`. Computed.

Identifier of the associated transaction.

<a id="schema-records--url"></a>

### url property

Type: `"string"`. Computed.

Endpoint URL or host associated with the event.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`),
    ""),
}
```
