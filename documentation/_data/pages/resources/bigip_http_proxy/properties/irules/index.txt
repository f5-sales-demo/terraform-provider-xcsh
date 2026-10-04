---
page_title: "irules"
subcategory: ""
description: "IRules Configuration for downstream connections."
xcsh_docs: {"aliases": ["irules"], "body_bytes": 1303, "body_sha256": "sha256:467c2b0a0e6639d1ed78392a2888fc2f2d6ee89bd1d4277cac493c4dbc8ffe45", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:irules:irules"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:irules", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "documentation/resources/bigip_http_proxy/properties/irules/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1120320133323032-2322013030113130-0032022100212120-1002211211233300-1000313031332132-3203200330320300-1100023210101120-0131020312332202", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["irules"], "schema_version": 1, "sections": [{"aliases": ["irules irules"], "anchor": "section", "description": "OPTIONS for attaching iRules to BIG-IP HTTP Proxy.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:irules:irules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-irules--irules--name", "enforcement": "provider-schema", "group": "irules.irules:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:irules:irules", "type": "requires"}], "schema_path": ["irules", "irules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/irules/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "IRules Configuration for downstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# irules

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- irules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IRules Configuration for downstream connections.

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
irules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/irules/irules/): complete subsection reference.

## Next pages

- [irules.irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/irules/irules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
