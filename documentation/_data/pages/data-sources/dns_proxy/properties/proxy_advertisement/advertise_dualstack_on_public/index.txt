---
page_title: "proxy_advertisement.advertise_dualstack_on_public"
subcategory: ""
description: "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available."
xcsh_docs: {"aliases": ["proxy advertisement advertise dualstack on public"], "body_bytes": 1844, "body_sha256": "sha256:afa93ef236a77cf9c9e0150c4b38c0cbea4ac576cac065620d1d2dbfdc05c766", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_dualstack_on_public:public_ip"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_dualstack_on_public", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement", "path": "documentation/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0201333130332013-0112020301131201-2021022313321201-3013233130311223-3022300202201021-2131301231003023-3121323331203223-2002102000332121", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_dualstack_on_public"], "schema_version": 1, "sections": [{"aliases": ["proxy advertisement advertise dualstack on public public ip"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_dualstack_on_public:public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_dualstack_on_public", "public_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_dualstack_on_public

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/)
- proxy_advertisement.advertise_dualstack_on_public

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/public_ip/): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_dualstack_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/public_ip/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
