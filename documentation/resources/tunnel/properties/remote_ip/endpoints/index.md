---
page_title: "remote_ip.endpoints"
subcategory: ""
description: "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel."
xcsh_docs: {"aliases": ["remote ip endpoints"], "body_bytes": 1641, "body_sha256": "sha256:4ceb56a5a2843022aab81cf05aacbe73903f503db8285bec347760a388158d8a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:remote_ip:endpoints:endpoints"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints", "parent_id": "xcsh-docs:resources:tunnel:properties:remote_ip", "path": "documentation/resources/tunnel/properties/remote_ip/endpoints/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3131101033002030-1000022121022120-0013113323333100-0321112100201110-2032211213331201-1111000223013220-3300302303113101-1020222312210221", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip", "endpoints"], "schema_version": 1, "sections": [{"aliases": ["endpoints"], "anchor": "section", "description": "Map of remote attributes to which tunnel will be established on per site node basis Every node can have a different attributes and IP address to connect to Key is ver node name and value is Remote node attributes.", "document_id": "xcsh-docs:resources:tunnel:properties:remote_ip:endpoints:endpoints", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "endpoints", "endpoints"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/endpoints/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.endpoints

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/)
- remote_ip.endpoints

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
endpoints {
  # Configure direct properties listed below.
}
```

## Direct properties

- [endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/endpoints/endpoints/): complete subsection reference.

## Next pages

- [remote_ip.endpoints.endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/endpoints/endpoints/)
- [remote_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/remote_ip/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
