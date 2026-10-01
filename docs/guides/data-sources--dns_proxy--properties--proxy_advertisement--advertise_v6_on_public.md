---
page_title: "proxy_advertisement.advertise_v6_on_public"
subcategory: ""
description: "proxy_advertisement.advertise_v6_on_public for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1454, "body_sha256": "sha256:6cae50758f442dbc0e50147e5b4b913e6ef7dcf50459789507d55e1543f7cce0", "canonical_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public:public_ip"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement", "path": "docs/guides/data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement", "advertise_v6_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement.advertise_v6_on_public for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_v6_on_public

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- [proxy_advertisement](data-sources--dns_proxy--properties--proxy_advertisement.md)
- proxy_advertisement.advertise_v6_on_public

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

- [public_ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_v6_on_public.public_ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md)
- [proxy_advertisement](data-sources--dns_proxy--properties--proxy_advertisement.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
