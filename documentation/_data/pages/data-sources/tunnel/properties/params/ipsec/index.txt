---
page_title: "params.ipsec"
subcategory: ""
description: "Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE."
xcsh_docs: {"aliases": ["params ipsec"], "body_bytes": 1424, "body_sha256": "sha256:77e2a85e46e0d0a5e21a885dfffba9db6e84b9480d5452aca3f87ae64f1e6f35", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec", "parent_id": "xcsh-docs:data-sources:tunnel:properties:params", "path": "documentation/data-sources/tunnel/properties/params/ipsec/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0030323213022232-2312233320100221-2220112221301302-0010102320222200-0302013300201022-0133213203111210-1313323212113030-1102320330331201", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["params", "ipsec"], "schema_version": 1, "sections": [{"aliases": ["ipsec psk"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["params", "ipsec", "ipsec_psk"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/params/ipsec/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Configuration for IPsec encapsulation are: &#8203;1. PSK - pre shared key to be used by IKE.

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

## Next pages

- [params.ipsec.ipsec_psk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/ipsec_psk/)
- [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
