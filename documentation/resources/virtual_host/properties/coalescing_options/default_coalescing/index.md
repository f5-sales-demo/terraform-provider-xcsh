---
page_title: "coalescing_options.default_coalescing"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["coalescing options default coalescing"], "body_bytes": 1317, "body_sha256": "sha256:14bbd8d790cd1b52622b10fe5084b2c8723db6494d1bc3d8867b7e01432a1888", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:coalescing_options:default_coalescing", "parent_id": "xcsh-docs:resources:virtual_host:properties:coalescing_options", "path": "documentation/resources/virtual_host/properties/coalescing_options/default_coalescing/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2012121230122021-0023032213020033-1233303132320033-3122023311001213-2010021213012031-1222000331003313-1020200120230101-2012231101333203", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["coalescing_options", "default_coalescing"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/coalescing_options/default_coalescing/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# coalescing_options.default_coalescing

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/coalescing_options/)
- coalescing_options.default_coalescing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

Upstream description:

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

Terraform syntax:

```terraform
default_coalescing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/coalescing_options/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
