---
page_title: "filters"
subcategory: ""
description: "Global Filters. Query Global Filters."
xcsh_docs: {"aliases": ["filters"], "body_bytes": 1652, "body_sha256": "sha256:f09a6d6a949e17491726516f38b53f984af9e011e0622c2f9ec59ec777bb7d1e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_summary:properties:filters:global_filters"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_summary:properties:filters", "parent_id": "xcsh-docs:data-sources:device_intelligence_summary:reference", "path": "documentation/data-sources/device_intelligence_summary/properties/filters/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_summary", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0310002002331232-2312333120230003-0000332312002113-1212213332001111-1322320003331030-0033011301031002-0011120310231111-1300301022011032", "registry_path": "docs/guides/data-sources--device_intelligence_summary--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["filters"], "schema_version": 1, "sections": [{"aliases": ["global filters"], "anchor": "section", "description": "Global Filters. List of global filters.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:properties:filters:global_filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["filters", "global_filters"], "syntax": "attribute", "type": "object"}, {"aliases": ["region filter"], "anchor": "schema-filters--region_filter", "description": "Defines a selection for Bot Defense region - US: US United States of America - EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are `US`, `EU`, `ASIA`, `CA`. Defaults to `US`.", "document_id": "xcsh-docs:data-sources:device_intelligence_summary:properties:filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filters", "region_filter"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_summary/properties/filters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Global Filters. Query Global Filters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filters

Breadcrumbs:

- [xcsh_device_intelligence_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/)
- filters

<a id="section"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

## Direct properties

- [global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/global_filters/): complete subsection reference.

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

- [filters.global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/filters/global_filters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/properties/)
- [xcsh_device_intelligence_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/)
