---
page_title: "origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check"
subcategory: ""
description: "origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1774, "body_sha256": "sha256:4be32861a2291182bb3d1716bd1def9386542b111c387adfe93020537d938b7f", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:icmp_health_check", "child_ids": [], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:icmp_health_check", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check", "path": "docs/guides/resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks--health_check--icmp_health_check.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "health_check", "icmp_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/icmp_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [origin_pools](resources--bigip_http_proxy--properties--origin_pools.md)
- [origin_pools.pools](resources--bigip_http_proxy--properties--origin_pools--pools.md)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md)
- [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks.md)
- [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks--health_check.md)
- origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check

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

- [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks--health_check.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
