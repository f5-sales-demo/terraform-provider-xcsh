---
page_title: "proxy_advertisement"
subcategory: ""
description: "Proxy Advertisement Type."
xcsh_docs: {"aliases": ["proxy advertisement"], "body_bytes": 4399, "body_sha256": "sha256:258d58aa2d6fade25584226e1da0be5c647176932780e58f3b5255906a093f82", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_dualstack_on_public", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_dualstack_vip", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_ipv6_vip", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_vip", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:do_not_advertise"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement", "parent_id": "xcsh-docs:data-sources:dns_proxy:reference", "path": "documentation/data-sources/dns_proxy/properties/proxy_advertisement/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement"], "schema_version": 1, "sections": [{"aliases": ["proxy advertisement advertise custom"], "anchor": "section", "description": "This defines a way to advertise a VIP on specific sites.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise dualstack on public"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_dualstack_on_public", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_dualstack_on_public"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise on public"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_on_public"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise on public default dualstack vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_dualstack_vip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_on_public_default_dualstack_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise on public default ipv6 vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_ipv6_vip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_on_public_default_ipv6_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise on public default vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_vip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_on_public_default_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise v6 on public"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_v6_on_public"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement do not advertise"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:do_not_advertise", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "do_not_advertise"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/proxy_advertisement/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Proxy Advertisement Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- proxy_advertisement

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for proxy advertisement.

Upstream description:

Proxy Advertisement Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_dualstack_on_public\",\"advertise_on_public\",\"advertise_on_public_default_dualstack_vip\",\"advertise_on_public_default_ipv6_vip\",\"advertise_on_public_default_vip\",\"advertise_v6_on_public\",\"do_not_advertise\"]"
}
```

## Direct properties

- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/): complete subsection reference.

- [advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/): complete subsection reference.

- [advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public/): complete subsection reference.

- [advertise_on_public_default_dualstack_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_dualstack_vip/): complete subsection reference.

- [advertise_on_public_default_ipv6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_ipv6_vip/): complete subsection reference.

- [advertise_on_public_default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_vip/): complete subsection reference.

- [advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/): complete subsection reference.

- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/do_not_advertise/): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/)
- [proxy_advertisement.advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/)
- [proxy_advertisement.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public/)
- [proxy_advertisement.advertise_on_public_default_dualstack_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_dualstack_vip/)
- [proxy_advertisement.advertise_on_public_default_ipv6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_ipv6_vip/)
- [proxy_advertisement.advertise_on_public_default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_vip/)
- [proxy_advertisement.advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/)
- [proxy_advertisement.do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/do_not_advertise/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
