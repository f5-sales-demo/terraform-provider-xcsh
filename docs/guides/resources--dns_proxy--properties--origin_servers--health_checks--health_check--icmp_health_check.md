---
page_title: "origin_servers.health_checks.health_check.icmp_health_check"
subcategory: ""
description: "origin_servers.health_checks.health_check.icmp_health_check for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1350, "body_sha256": "sha256:e98bd31fa05c90c0b8d40e923e4f0eae6137f54499d5c763b78267f7e6fac347", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:icmp_health_check", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check:icmp_health_check", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check", "path": "docs/guides/resources--dns_proxy--properties--origin_servers--health_checks--health_check--icmp_health_check.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "health_checks", "health_check", "icmp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/health_checks/health_check/icmp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.health_checks.health_check.icmp_health_check for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.health_checks.health_check.icmp_health_check

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [origin_servers](resources--dns_proxy--properties--origin_servers.md)
- [origin_servers.health_checks](resources--dns_proxy--properties--origin_servers--health_checks.md)
- [origin_servers.health_checks.health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check.md)
- origin_servers.health_checks.health_check.icmp_health_check

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for icmp health check.

Upstream description:

This can be used for messages where no values are needed.

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
icmp_health_check = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_servers.health_checks.health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
