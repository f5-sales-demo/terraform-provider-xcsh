---
page_title: "remote_ip"
subcategory: ""
description: "Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map of IP address on per ver node basis."
xcsh_docs: {"aliases": ["remote ip"], "body_bytes": 1923, "body_sha256": "sha256:30ad099dd54b3d1c46599fd152f307d9d22eba5fd166b010c0bf58ec3380d204", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip", "parent_id": "xcsh-docs:data-sources:tunnel:reference", "path": "documentation/data-sources/tunnel/properties/remote_ip/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3132312023230012-3013110002122031-3132111221301201-3022331010323223-1232301300132031-2222202103210231-2231122010211301-2302103301333120", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["remote_ip"], "schema_version": 1, "sections": [{"aliases": ["endpoints"], "anchor": "section", "description": "Provides a map of ver node name to remote node attributes Ver node should use these attributes to configure as remote tunnel.", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["ip"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["remote_ip", "ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map of IP address on per ver node basis.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- remote_ip

<a id="section"></a>

Type: `"single"`. Computed.

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP
Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map
of IP address on per ver node basis.

Upstream description:

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - &#8203;1.
IP Address - Specifies the remote IP to which tunnel has to be connected &#8203;2. Remote endpoint -
Is a map of IP address on per ver node basis.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"endpoints\",\"ip\"]"
}
```

## Direct properties

- [endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/endpoints/): complete subsection reference.

- [ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/): complete subsection reference.

## Next pages

- [remote_ip.endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/endpoints/)
- [remote_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/remote_ip/ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
