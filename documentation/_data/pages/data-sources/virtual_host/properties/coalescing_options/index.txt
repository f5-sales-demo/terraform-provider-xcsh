---
page_title: "coalescing_options"
subcategory: ""
description: "coalescing_options for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1843, "body_sha256": "sha256:fdfc57e653a2b150e704d0276311759f3cee5ce62d66c20e5cc42852d3b1105d", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:coalescing_options:default_coalescing", "xcsh-docs:data-sources:virtual_host:properties:coalescing_options:strict_coalescing"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:coalescing_options", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "documentation/data-sources/virtual_host/properties/coalescing_options/index.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["coalescing_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/coalescing_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "coalescing_options for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# coalescing_options

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- coalescing_options

<a id="section"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

## Direct properties

- [default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/coalescing_options/default_coalescing/): complete subsection reference.

- [strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/coalescing_options/strict_coalescing/): complete subsection reference.

## Next pages

- [coalescing_options.default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/coalescing_options/default_coalescing/)
- [coalescing_options.strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/coalescing_options/strict_coalescing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
