---
page_title: "advanced_options.enable_subsets"
subcategory: "Load Balancing"
description: "Configure subset OPTIONS for origin pool."
xcsh_docs: {"aliases": ["advanced options enable subsets"], "body_bytes": 1698, "body_sha256": "sha256:2d3cb814227572d5182a336ccef91783fc0f5dd449e864b55fc2db9b496a8026", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:fail_request"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "path": "documentation/data-sources/origin_pool/properties/advanced_options/enable_subsets/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0110032023000030-2133333102332220-1113122213100231-1312233300221201-0111112301331212-3313313311200001-0030202030202212-2003333122112321", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "enable_subsets"], "schema_version": 1, "sections": [{"aliases": ["advanced options enable subsets any endpoint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "any_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced options enable subsets default subset"], "anchor": "section", "description": "Default Subset definition.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "default_subset"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced options enable subsets endpoint subsets"], "anchor": "section", "description": "List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset within the class.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "endpoint_subsets"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced options enable subsets fail request"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:fail_request", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "fail_request"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/enable_subsets/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configure subset OPTIONS for origin pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["origin_poolCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
