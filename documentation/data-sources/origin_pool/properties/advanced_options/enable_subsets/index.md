---
page_title: "advanced_options.enable_subsets"
subcategory: "Load Balancing"
description: "Configure subset OPTIONS for origin pool."
xcsh_docs: {"aliases": ["advanced options enable subsets", "backend servers", "origin servers", "upstream servers"], "body_bytes": 2714, "body_sha256": "sha256:4214a7234240a10ed71a4f16c3aea12a8d4e7d891881fb2ff6d37fa827a0b00c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:fail_request"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "path": "documentation/data-sources/origin_pool/properties/advanced_options/enable_subsets/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "enable_subsets"], "schema_version": 1, "sections": [{"aliases": ["any endpoint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "any_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["default subset"], "anchor": "section", "description": "Default Subset definition.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "default_subset"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint subsets"], "anchor": "section", "description": "List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset within the class.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "endpoint_subsets"], "syntax": "attribute", "type": "object"}, {"aliases": ["fail request"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:fail_request", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "fail_request"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/enable_subsets/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure subset OPTIONS for origin pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.enable_subsets

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/)
- advanced_options.enable_subsets

<a id="section"></a>

Type: `"single"`. Computed.

Configure subset OPTIONS for origin pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fallback_policy_choice": "[\"any_endpoint\",\"default_subset\",\"fail_request\"]"
}
```

## Direct properties

- [any_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/any_endpoint/): complete subsection reference.

- [default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/default_subset/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/endpoint_subsets/): complete subsection reference.

- [fail_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/fail_request/): complete subsection reference.

## Next pages

- [advanced_options.enable_subsets.any_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/any_endpoint/)
- [advanced_options.enable_subsets.default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/default_subset/)
- [advanced_options.enable_subsets.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/endpoint_subsets/)
- [advanced_options.enable_subsets.fail_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/fail_request/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
