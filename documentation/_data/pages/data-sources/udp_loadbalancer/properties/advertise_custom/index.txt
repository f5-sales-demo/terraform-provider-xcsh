---
page_title: "advertise_custom"
subcategory: ""
description: "This defines a way to advertise a VIP on specific sites."
xcsh_docs: {"aliases": ["advertise custom"], "body_bytes": 2253, "body_sha256": "sha256:88433887d94a50f5e07f79c92ab2bd76a3e1f6b6f6fbcc1bfdcf535ad945cab8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom", "parent_id": "xcsh-docs:data-sources:udp_loadbalancer:reference", "path": "documentation/data-sources/udp_loadbalancer/properties/advertise_custom/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112", "registry_path": "docs/guides/data-sources--udp_loadbalancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_custom"], "schema_version": 1, "sections": [{"aliases": ["advertise where"], "anchor": "section", "description": "Where should this load balancer be available.", "document_id": "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom:advertise_where", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["advertise_custom", "advertise_where"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/properties/advertise_custom/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a way to advertise a VIP on specific sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom

Breadcrumbs:

- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/)
- advertise_custom

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: advertise\_custom, advertise\_on\_public, advertise\_on\_public\_default\_vip,
do\_not\_advertise; Default: advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP
on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

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

OneOf alternatives in this subsection:

- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/#section)
- [advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_on_public/#section)
- [advertise_on_public_default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_on_public_default_vip/#section)
- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/do_not_advertise/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/advertise_custom/advertise_where/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/)
- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/)
