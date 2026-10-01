---
page_title: "time_series_results.time_series"
subcategory: ""
description: "time_series_results.time_series for xcsh_device_intelligence_high_risk_transactions."
xcsh_docs: {"aliases": [], "body_bytes": 1472, "body_sha256": "sha256:3b3dc54da04db8244933c8a06da0cb46a5b721dffd2ebfb5c6698d34a25c6cab", "canonical_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results:time_series", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results:time_series", "parent_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results", "path": "docs/guides/data-sources--device_intelligence_high_risk_transactions--properties--time_series_results--time_series.md", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["time_series_results", "time_series"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/time_series/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "time_series_results.time_series for xcsh_device_intelligence_high_risk_transactions.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# time_series_results.time_series

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference.md)
- [time_series_results](data-sources--device_intelligence_high_risk_transactions--properties--time_series_results.md)
- time_series_results.time_series

<a id="section"></a>

Type: `"list"`. Computed.

Sequence of timestamped values for this series.

## Direct properties

<a id="schema-time_series_results--time_series--timestamp"></a>

### timestamp property

Type: `"string"`. Computed.

Timestamp (epoch seconds). Unix epoch timestamp in seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="schema-time_series_results--time_series--value"></a>

### value property

Type: `"string"`. Computed.

Value. Value observed at the given timestamp.

## Next pages

- [time_series_results](data-sources--device_intelligence_high_risk_transactions--properties--time_series_results.md)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md)
