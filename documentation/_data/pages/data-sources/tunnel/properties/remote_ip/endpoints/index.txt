---
page_title: "remote_ip.endpoints"
subcategory: ""
description: "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel."
xcsh_docs: {"aliases": ["remote ip endpoints"], "body_bytes": 1544, "body_sha256": "sha256:7f6d36fdc78be6461e83586cb00b4e56c33ebbb60096f1ef7220a63bcb3cde6d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints:endpoints"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "parent_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip", "path": "documentation/data-sources/tunnel/properties/remote_ip/endpoints/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3003001003012231-2020031310333033-1120113023030332-2131213220100033-3331220330322202-1332221101013030-1121210230102320-0332331033123313", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip", "endpoints"], "schema_version": 1, "sections": [{"aliases": ["remote ip endpoints endpoints"], "anchor": "section", "description": "Map of remote attributes to which tunnel will be established on per site node basis Every node can have a different attributes and IP address to connect to Key is ver node name and value is Remote node attributes.", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints:endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "endpoints", "endpoints"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/endpoints/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.endpoints

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/)
- remote_ip.endpoints

<a id="section"></a>

Type: `"single"`. Computed.

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

Upstream description:

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

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

## Direct properties

- [endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/endpoints/endpoints/): complete subsection reference.

## Next pages

- [remote_ip.endpoints.endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/endpoints/endpoints/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
