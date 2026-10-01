---
page_title: "origin_pools.pools.origin_servers"
subcategory: ""
description: "origin_pools.pools.origin_servers for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3264, "body_sha256": "sha256:00160d65610c34c9b1d0d5541e09f8d6d17e1ce2b4a2a6d5a1c5b7151d19e1a1", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:automatic_port", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:lb_port", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools", "path": "docs/guides/data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [origin_pools](data-sources--bigip_http_proxy--properties--origin_pools.md)
- [origin_pools.pools](data-sources--bigip_http_proxy--properties--origin_pools--pools.md)
- origin_pools.pools.origin_servers

<a id="section"></a>

Type: `"single"`. Computed.

List of origin Servers for the BIG-IP HTTP Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]"
}
```

## Direct properties

- [automatic_port](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--automatic_port.md): complete subsection reference.

- [health_checks](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks.md): complete subsection reference.

- [lb_port](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--lb_port.md): complete subsection reference.

- [origin_servers](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers.md): complete subsection reference.

<a id="schema-origin_pools--pools--origin_servers--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [origin_pools.pools.origin_servers.automatic_port](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--automatic_port.md)
- [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks.md)
- [origin_pools.pools.origin_servers.lb_port](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--lb_port.md)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers.md)
- [origin_pools.pools](data-sources--bigip_http_proxy--properties--origin_pools--pools.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
