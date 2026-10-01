---
page_title: "records"
subcategory: ""
description: "records for xcsh_device_intelligence_device_history."
xcsh_docs: {"aliases": [], "body_bytes": 2376, "body_sha256": "sha256:d717cdee117653f52d9f556ae1577dea0639d587a0e8937d06e3033fa6d27f76", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "path": "documentation/data-sources/device_intelligence_device_history/properties/records/index.md", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["records"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/properties/records/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "records for xcsh_device_intelligence_device_history.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
