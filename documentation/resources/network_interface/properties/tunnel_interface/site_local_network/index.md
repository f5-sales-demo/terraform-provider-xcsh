---
page_title: "tunnel_interface.site_local_network"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tunnel interface site local network"], "body_bytes": 1311, "body_sha256": "sha256:9d8a633d93925eec99bcae8561e337d4fa8bd4bfc8c2883d82daaede4face814", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_network", "parent_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface", "path": "documentation/resources/network_interface/properties/tunnel_interface/site_local_network/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3212113202201030-2202333120032003-3031302320112300-2311223300103221-1313002220323323-2010130030213203-1131231031332120-0030230321133032", "registry_path": "docs/guides/resources--network_interface--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tunnel_interface", "site_local_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/tunnel_interface/site_local_network/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tunnel_interface.site_local_network

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/)
- tunnel_interface.site_local_network

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
site_local_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/tunnel_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
