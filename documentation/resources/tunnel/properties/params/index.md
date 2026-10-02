---
page_title: "params"
subcategory: ""
description: "Tunnel configuration parameters for supported encapsulation 1. IPsec is supported with PSK for which PSK can be configured."
xcsh_docs: {"aliases": ["params"], "body_bytes": 1492, "body_sha256": "sha256:32ace857029c50e31e5e48dc1a5912fbc601d0fe24211c303e694c430247fed4", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:params:ipsec"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:params", "parent_id": "xcsh-docs:resources:tunnel:reference", "path": "documentation/resources/tunnel/properties/params/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3112221221213033-1303222023230230-0222322021220323-2332312322123003-2200133123232101-3021120033320130-3200322033133310-2131310310120213", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["params"], "schema_version": 1, "sections": [{"aliases": ["ipsec"], "anchor": "section", "description": "Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE.", "document_id": "xcsh-docs:resources:tunnel:properties:params:ipsec", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["params", "ipsec"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/params/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Tunnel configuration parameters for supported encapsulation 1. IPsec is supported with PSK for which PSK can be configured.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Tunnel configuration parameters for supported encapsulation 1. IPsec is supported with PSK for which
PSK can be configured.

Upstream description:

Tunnel configuration parameters for supported encapsulation &#8203;1. IPsec is supported with PSK
for which PSK can be configured.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"ipsec\"]"
}
```

Terraform syntax:

```terraform
params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/): complete subsection reference.

## Next pages

- [params.ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
