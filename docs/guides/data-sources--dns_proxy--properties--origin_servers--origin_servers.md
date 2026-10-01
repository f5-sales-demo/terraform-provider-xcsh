---
page_title: "origin_servers.origin_servers"
subcategory: ""
description: "origin_servers.origin_servers for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3027, "body_sha256": "sha256:d1f4effce57165ab8fad85433af01195572ba4f8962860cf421236a6f5f0dc89", "canonical_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_name", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:site_preferences"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers", "path": "docs/guides/data-sources--dns_proxy--properties--origin_servers--origin_servers.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "origin_servers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.origin_servers for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- [origin_servers](data-sources--dns_proxy--properties--origin_servers.md)
- origin_servers.origin_servers

<a id="section"></a>

Type: `"list"`. Computed.

List Of Origin Servers. List of origin servers for Proxy.

Upstream description:

List of origin servers for Proxy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [k8s_service](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md): complete subsection reference.

- [no_preference](data-sources--dns_proxy--properties--origin_servers--origin_servers--no_preference.md): complete subsection reference.

- [public_ip](data-sources--dns_proxy--properties--origin_servers--origin_servers--public_ip.md): complete subsection reference.

- [public_name](data-sources--dns_proxy--properties--origin_servers--origin_servers--public_name.md): complete subsection reference.

- [site_preferences](data-sources--dns_proxy--properties--origin_servers--origin_servers--site_preferences.md): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- [origin_servers.origin_servers.no_preference](data-sources--dns_proxy--properties--origin_servers--origin_servers--no_preference.md)
- [origin_servers.origin_servers.public_ip](data-sources--dns_proxy--properties--origin_servers--origin_servers--public_ip.md)
- [origin_servers.origin_servers.public_name](data-sources--dns_proxy--properties--origin_servers--origin_servers--public_name.md)
- [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--properties--origin_servers--origin_servers--site_preferences.md)
- [origin_servers](data-sources--dns_proxy--properties--origin_servers.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
