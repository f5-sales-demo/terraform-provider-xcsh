---
page_title: "params.ipsec"
subcategory: ""
description: "Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE."
xcsh_docs: {"aliases": ["params ipsec"], "body_bytes": 951, "body_sha256": "sha256:e3fa5ee8223257188dfe3a90f7d2475b2a83173396bd042806d012d545c10f24", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec", "parent_id": "xcsh-docs:data-sources:tunnel:properties:params", "path": "documentation/data-sources/tunnel/properties/params/ipsec/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0030323213022232-2312233320100221-2220112221301302-0010102320222200-0302013300201022-0133213203111210-1313323212113030-1102320330331201", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["params", "ipsec"], "schema_version": 1, "sections": [{"aliases": ["params ipsec ipsec psk"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["params", "ipsec", "ipsec_psk"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/params/ipsec/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params.ipsec

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/)
- params.ipsec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE.

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

- [ipsec_psk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/ipsec_psk/): complete subsection reference.
