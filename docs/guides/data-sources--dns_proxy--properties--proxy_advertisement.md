---
page_title: "proxy_advertisement"
subcategory: ""
description: "proxy_advertisement for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3391, "body_sha256": "sha256:f9545a19f62808745316d99cc0d659460a71b83b1049260f23bbeba779285573", "canonical_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_custom", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_dualstack_on_public", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_dualstack_vip", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_ipv6_vip", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_vip", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement:do_not_advertise"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement", "parent_id": "xcsh-docs:data-sources:dns_proxy:reference", "path": "docs/guides/data-sources--dns_proxy--properties--proxy_advertisement.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/proxy_advertisement/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Property reference](data-sources--dns_proxy--reference.md)
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

- [advertise_custom](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom.md): complete subsection reference.

- [advertise_dualstack_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public.md): complete subsection reference.

- [advertise_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public.md): complete subsection reference.

- [advertise_on_public_default_dualstack_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_dualstack_vip.md): complete subsection reference.

- [advertise_on_public_default_ipv6_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_ipv6_vip.md): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_vip.md): complete subsection reference.

- [advertise_v6_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public.md): complete subsection reference.

- [do_not_advertise](data-sources--dns_proxy--properties--proxy_advertisement--do_not_advertise.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom.md)
- [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public.md)
- [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public.md)
- [proxy_advertisement.advertise_on_public_default_dualstack_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_dualstack_vip.md)
- [proxy_advertisement.advertise_on_public_default_ipv6_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_ipv6_vip.md)
- [proxy_advertisement.advertise_on_public_default_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_vip.md)
- [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public.md)
- [proxy_advertisement.do_not_advertise](data-sources--dns_proxy--properties--proxy_advertisement--do_not_advertise.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
