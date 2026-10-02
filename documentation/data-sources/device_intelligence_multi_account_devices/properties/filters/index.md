---
page_title: "filters"
subcategory: ""
description: "Global Filters. Query Global Filters."
xcsh_docs: {"aliases": ["filters"], "body_bytes": 1764, "body_sha256": "sha256:2a5e7f34879a99889cb91aa3d3ef3a50584e85d26d35b6b88d12fd80f729b177", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters:global_filters"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "path": "documentation/data-sources/device_intelligence_multi_account_devices/properties/filters/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1102032323221001-0221233322123221-3230303132011023-0033322310323010-2232202301320220-1331223132023300-0222030001102313-3113222101022233", "registry_path": "docs/guides/data-sources--device_intelligence_multi_account_devices--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["filters"], "schema_version": 1, "sections": [{"aliases": ["global filters"], "anchor": "section", "description": "Global Filters. List of global filters.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters:global_filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["filters", "global_filters"], "syntax": "attribute", "type": "object"}, {"aliases": ["region filter"], "anchor": "schema-filters--region_filter", "description": "Defines a selection for Bot Defense region - US: US United States of America - EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are `US`, `EU`, `ASIA`, `CA`. Defaults to `US`.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:filters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["filters", "region_filter"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/properties/filters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Global Filters. Query Global Filters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# filters

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/)
- filters

<a id="section"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

## Direct properties

- [global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/global_filters/): complete subsection reference.

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

- [filters.global_filters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/filters/global_filters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/)
- [xcsh_device_intelligence_multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/)
