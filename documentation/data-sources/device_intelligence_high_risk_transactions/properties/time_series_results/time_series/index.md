---
page_title: "time_series_results.time_series"
subcategory: ""
description: "Sequence of timestamped values for this series."
xcsh_docs: {"aliases": ["time series results time series"], "body_bytes": 1729, "body_sha256": "sha256:a7c6f04ede359ae6c7fed2a410382cdca36517b679627fa5ac28570d5c319a18", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results:time_series", "parent_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results", "path": "documentation/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/time_series/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1103330333220221-1330301211031103-2300031012001033-1321100232023320-3102010201013032-3301230000320123-0013032102112133-0131033330033331", "registry_path": "docs/guides/data-sources--device_intelligence_high_risk_transactions--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["time_series_results", "time_series"], "schema_version": 1, "sections": [{"aliases": ["timestamp"], "anchor": "schema-time_series_results--time_series--timestamp", "description": "Timestamp (epoch seconds). Unix epoch timestamp in seconds.", "document_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results:time_series", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["time_series_results", "time_series", "timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["value"], "anchor": "schema-time_series_results--time_series--value", "description": "Value. Value observed at the given timestamp.", "document_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results:time_series", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["time_series_results", "time_series", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/time_series/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Sequence of timestamped values for this series.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# time_series_results.time_series

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/properties/)
- [time_series_results](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/)
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

- [time_series_results](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/)
- [xcsh_device_intelligence_high_risk_transactions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/)
