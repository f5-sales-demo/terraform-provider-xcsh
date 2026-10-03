---
page_title: "proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public"
subcategory: ""
description: "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available."
xcsh_docs: {"aliases": ["proxy advertisement advertise custom advertise where advertise dualstack on public"], "body_bytes": 2558, "body_sha256": "sha256:5e799056c836bc1462385f6beb5dd78c6341b2d9d00aaa9c071009abc48dc3d0", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_dualstack_on_public:public_ip"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_dualstack_on_public", "parent_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "path": "documentation/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1210302231200223-3200232113022130-0220223102302232-3000323111130131-1101121322121130-1100201211301122-3202031033020112-3030021320132313", "registry_path": "docs/guides/resources--dns_proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_dualstack_on_public"], "schema_version": 1, "sections": [{"aliases": ["proxy advertisement advertise custom advertise where advertise dualstack on public public ip"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_dualstack_on_public:public_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--name", "enforcement": "provider-schema", "group": "proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_dualstack_on_public:public_ip", "type": "requires"}], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_dualstack_on_public", "public_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/)
- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/)
- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

## Direct properties

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/)
- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
