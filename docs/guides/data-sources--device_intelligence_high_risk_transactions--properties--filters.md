---
page_title: "filters"
subcategory: ""
description: "filters for xcsh_device_intelligence_high_risk_transactions."
xcsh_docs: {"aliases": [], "body_bytes": 1464, "body_sha256": "sha256:84af871fd96400247c4516e7ff37db51743fc5e46215b5efbb16a3f66818fb13", "canonical_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:filters", "child_ids": ["xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:filters:global_filters"], "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:filters", "parent_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:reference", "path": "docs/guides/data-sources--device_intelligence_high_risk_transactions--properties--filters.md", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["filters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/properties/filters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "filters for xcsh_device_intelligence_high_risk_transactions.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filters

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference.md)
- filters

<a id="section"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

## Direct properties

- [global_filters](data-sources--device_intelligence_high_risk_transactions--properties--filters--global_filters.md): complete subsection reference.

<a id="schema-filters--region_filter"></a>

### region_filter property

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA",
    "CA"),
}
```

## Next pages

- [filters.global_filters](data-sources--device_intelligence_high_risk_transactions--properties--filters--global_filters.md)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference.md)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md)
