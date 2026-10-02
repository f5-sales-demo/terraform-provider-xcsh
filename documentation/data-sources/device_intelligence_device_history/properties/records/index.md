---
page_title: "records"
subcategory: ""
description: "Records. List of activity records."
xcsh_docs: {"aliases": ["records"], "body_bytes": 2376, "body_sha256": "sha256:d717cdee117653f52d9f556ae1577dea0639d587a0e8937d06e3033fa6d27f76", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "path": "documentation/data-sources/device_intelligence_device_history/properties/records/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0103123322112321-1002002023332310-2322312220320002-3121012100220233-0020102331021200-0002012222320112-3133111021102023-1323303003123020", "registry_path": "docs/guides/data-sources--device_intelligence_device_history--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["records"], "schema_version": 1, "sections": [{"aliases": ["action taken"], "anchor": "schema-records--action_taken", "description": "Action taken in response to the risk evaluation for this event.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "action_taken"], "syntax": "attribute", "type": "string"}, {"aliases": ["detected signals"], "anchor": "schema-records--detected_signals", "description": "Risk or integrity signals observed for this event.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "detected_signals"], "syntax": "attribute", "type": "list"}, {"aliases": ["endpoint label"], "anchor": "schema-records--endpoint_label", "description": "Human-readable label for the endpoint or action.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "endpoint_label"], "syntax": "attribute", "type": "string"}, {"aliases": ["risk score"], "anchor": "schema-records--risk_score", "description": "Risk score assigned to this event (0–100).", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "risk_score"], "syntax": "attribute", "type": "number"}, {"aliases": ["timestamp"], "anchor": "schema-records--timestamp", "description": "Event timestamp in RFC 3339 format with milliseconds (UTC).", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["txn id"], "anchor": "schema-records--txn_id", "description": "Identifier of the associated transaction.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "txn_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-records--url", "description": "Endpoint URL or host associated with the event.", "document_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["records", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/properties/records/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Records. List of activity records.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`),
    ""),
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/)
- [xcsh_device_intelligence_device_history](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/)
