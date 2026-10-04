---
page_title: "bgp_parameters.local_address"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["bgp parameters local address"], "body_bytes": 1186, "body_sha256": "sha256:fb58a6b8ca4882719836bcc3c98feac72ffe5f021d8875d9218c98399e341f42", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:bgp_parameters:local_address", "parent_id": "xcsh-docs:resources:bgp:properties:bgp_parameters", "path": "documentation/resources/bgp/properties/bgp_parameters/local_address/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0322212333332301-3100201110222010-3210102120330212-2321123232001100-2011033130030100-3032303131220330-2123012131323112-3321230120321320", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bgp_parameters", "local_address"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/bgp_parameters/local_address/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["bgpCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bgp_parameters.local_address

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [bgp_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/bgp_parameters/)
- bgp_parameters.local_address

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
local_address = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bgp_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/bgp_parameters/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
