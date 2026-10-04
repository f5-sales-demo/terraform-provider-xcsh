---
page_title: "default_pool.advanced_options.enable_subsets.default_subset"
subcategory: "Load Balancing"
description: "Default Subset definition."
xcsh_docs: {"aliases": ["default pool advanced options enable subsets default subset"], "body_bytes": 2212, "body_sha256": "sha256:3e1940aeef9459de1b7c32af51df70c4b6d7f268a3e753f7a8fba2f10845f4d8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset:default_subset"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "path": "documentation/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0311121031003212-1023312103112111-1223133333320310-2000132222222210-0111223230023131-1033323031033322-0303122132222203-2121101033013301", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "advanced_options", "enable_subsets", "default_subset"], "schema_version": 1, "sections": [{"aliases": ["default pool advanced options enable subsets default subset default subset"], "anchor": "section", "description": "List of key-value pairs that define default subset. Which gets used when route specifies no metadata or no subset matching the metadata exists.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset:default_subset", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "advanced_options", "enable_subsets", "default_subset", "default_subset"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Default Subset definition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.enable_subsets.default_subset

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/)
- [default_pool.advanced_options.enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/)
- default_pool.advanced_options.enable_subsets.default_subset

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default subset.

Upstream description:

Default Subset definition.

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
default_subset {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/default_subset/): complete subsection reference.

## Next pages

- [default_pool.advanced_options.enable_subsets.default_subset.default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/default_subset/)
- [default_pool.advanced_options.enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
