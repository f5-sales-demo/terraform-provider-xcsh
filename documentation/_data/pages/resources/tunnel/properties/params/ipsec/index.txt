---
page_title: "params.ipsec"
subcategory: ""
description: "Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE."
xcsh_docs: {"aliases": ["params ipsec"], "body_bytes": 1517, "body_sha256": "sha256:f08fb6750b6b8a0604c47df891c6728ed1f821e61f55b2978d76b8869ea1cc26", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:params:ipsec", "parent_id": "xcsh-docs:resources:tunnel:properties:params", "path": "documentation/resources/tunnel/properties/params/ipsec/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1333220212012230-1111003122111213-2201303031102303-3212013001210022-0333212002213133-3030003012100230-3000310333301101-1011203123230131", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["params", "ipsec"], "schema_version": 1, "sections": [{"aliases": ["params ipsec ipsec psk"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "params.ipsec.ipsec_psk:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "params.ipsec.ipsec_psk:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info", "type": "conflicts"}], "schema_path": ["params", "ipsec", "ipsec_psk"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/params/ipsec/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params.ipsec

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/)
- params.ipsec

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipsec {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipsec_psk](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/): complete subsection reference.

## Next pages

- [params.ipsec.ipsec_psk](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/)
- [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
