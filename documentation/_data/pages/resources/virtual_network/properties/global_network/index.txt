---
page_title: "global_network"
subcategory: "Networking"
description: "Select the global virtual-network scope for connectivity across participating sites."
xcsh_docs: {"aliases": ["global network"], "body_bytes": 1822, "body_sha256": "sha256:67f6e2d3dfdb1a8be3f9f476fadb8f796e19c2fbc8a6a0c418794c80bc640488", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:properties:global_network", "parent_id": "xcsh-docs:resources:virtual_network:reference", "path": "documentation/resources/virtual_network/properties/global_network/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2022100300112003-1220202113121001-1131003021302131-3131310112012213-0021220111013333-0303333111203120-3031131131222030-2012313201111120", "registry_path": "docs/guides/resources--virtual_network--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["global_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/properties/global_network/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select the global virtual-network scope for connectivity across participating sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# global_network

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/)
- global_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: global\_network, site\_local\_inside\_network, site\_local\_network\] Select the global
virtual-network scope for connectivity across participating sites.

Upstream description:

Select the global virtual-network scope for connectivity across participating sites.

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

OneOf alternatives in this subsection:

- [global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/global_network/#section)
- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/site_local_inside_network/#section)
- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/site_local_network/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
global_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/)
- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
