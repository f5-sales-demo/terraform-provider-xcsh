---
page_title: "site_local_network"
subcategory: "Networking"
description: "Select a site-local virtual network when connectivity must remain within one site."
xcsh_docs: {"aliases": ["site local network"], "body_bytes": 1097, "body_sha256": "sha256:78dc16b9e9ca1cbd74b3b71379dc18e29c4961d9c56dafb0de6e0062957808e3", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:properties:site_local_network", "parent_id": "xcsh-docs:resources:virtual_network:reference", "path": "documentation/resources/virtual_network/properties/site_local_network/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2221221203312202-2002322102302002-1023111120132301-3331003331013221-1031132000201122-3011230312213030-3013302022320022-3112321211102311", "registry_path": "docs/guides/resources--virtual_network--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_local_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/properties/site_local_network/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select a site-local virtual network when connectivity must remain within one site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_local_network

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/)
- site_local_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

Select a site-local virtual network when connectivity must remain within one site.

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
site_local_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/)
- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
