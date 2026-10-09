---
page_title: "coalescing_options.strict_coalescing"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["coalescing options strict coalescing"], "body_bytes": 994, "body_sha256": "sha256:e8c348dea761ecdd24d03f1d91545a3b49d2d2ca4d43fc35cb5b757cddb78c68", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:coalescing_options:strict_coalescing", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:coalescing_options", "path": "documentation/data-sources/virtual_host/properties/coalescing_options/strict_coalescing/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1033223010220211-3000021102020113-3003031222230013-2132001300002312-0030131223331001-0211113233311210-1011102112032202-0323301122011323", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["coalescing_options", "strict_coalescing"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/coalescing_options/strict_coalescing/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# coalescing_options.strict_coalescing

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/coalescing_options/)
- coalescing_options.strict_coalescing

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.
