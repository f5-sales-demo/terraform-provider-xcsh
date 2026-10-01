---
page_title: "origin_servers"
subcategory: ""
description: "origin_servers for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1133, "body_sha256": "sha256:74755cf742870406e61c7ae8e1bbe9ebb773f61a6fada3817ff711848fac9dd3", "canonical_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:health_checks", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers", "parent_id": "xcsh-docs:data-sources:dns_proxy:reference", "path": "docs/guides/data-sources--dns_proxy--properties--origin_servers.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- origin_servers

<a id="section"></a>

Type: `"single"`. Computed.

List of origin Servers for the DNS proxy.

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

- [health_checks](data-sources--dns_proxy--properties--origin_servers--health_checks.md): complete subsection reference.

- [origin_servers](data-sources--dns_proxy--properties--origin_servers--origin_servers.md): complete subsection reference.

## Next pages

- [origin_servers.health_checks](data-sources--dns_proxy--properties--origin_servers--health_checks.md)
- [origin_servers.origin_servers](data-sources--dns_proxy--properties--origin_servers--origin_servers.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
