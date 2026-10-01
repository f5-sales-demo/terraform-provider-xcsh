---
page_title: "origin_servers.origin_servers.k8s_service.snat_pool.snat_pool"
subcategory: ""
description: "origin_servers.origin_servers.k8s_service.snat_pool.snat_pool for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2439, "body_sha256": "sha256:f401f399ee9783069f2629c798b199e3881cc5ad44b91c60788f03607466f914", "canonical_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool:snat_pool", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool:snat_pool", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool", "path": "docs/guides/data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "origin_servers", "k8s_service", "snat_pool", "snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.origin_servers.k8s_service.snat_pool.snat_pool for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- [origin_servers](data-sources--dns_proxy--properties--origin_servers.md)
- [origin_servers.origin_servers](data-sources--dns_proxy--properties--origin_servers--origin_servers.md)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool.md)
- origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="schema-origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
