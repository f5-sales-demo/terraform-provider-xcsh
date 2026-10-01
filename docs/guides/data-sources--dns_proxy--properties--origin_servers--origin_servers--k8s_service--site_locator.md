---
page_title: "origin_servers.origin_servers.k8s_service.site_locator"
subcategory: ""
description: "origin_servers.origin_servers.k8s_service.site_locator for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1974, "body_sha256": "sha256:faf3e32e1674eaddfce187e3de6493fa66fe788948fb434467eeea2cc9d44186", "canonical_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator:site", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator:virtual_site"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "path": "docs/guides/data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "origin_servers", "k8s_service", "site_locator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.origin_servers.k8s_service.site_locator for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.k8s_service.site_locator

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- [origin_servers](data-sources--dns_proxy--properties--origin_servers.md)
- [origin_servers.origin_servers](data-sources--dns_proxy--properties--origin_servers--origin_servers.md)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- origin_servers.origin_servers.k8s_service.site_locator

<a id="section"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

## Direct properties

- [site](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md): complete subsection reference.

- [virtual_site](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.k8s_service.site_locator.site](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md)
- [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
