---
page_title: "records"
subcategory: ""
description: "records for xcsh_device_intelligence_device_history."
xcsh_docs: {"aliases": [], "body_bytes": 2069, "body_sha256": "sha256:70c173c3ee6fb43c6eaf1b7bcba17cc022698cde451c0acd0dcc4a0cfe1b18de", "canonical_id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:properties:records", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:reference", "path": "docs/guides/data-sources--device_intelligence_device_history--properties--records.md", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["records"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/properties/records/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "records for xcsh_device_intelligence_device_history.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# records

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md)
- [Property reference](data-sources--device_intelligence_device_history--reference.md)
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

- [Property reference](data-sources--device_intelligence_device_history--reference.md)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md)
