---
page_title: "origin_pools.pools.origin_servers.origin_servers"
subcategory: ""
description: "origin_pools.pools.origin_servers.origin_servers for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 4009, "body_sha256": "sha256:fb6f8cd8ef47036aef497df4ee72e6bfe92fade3eb69e1a55b7e445d036007bd", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:public_ip", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:public_name"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers", "path": "docs/guides/resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.origin_servers for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.origin_servers

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [origin_pools](resources--bigip_http_proxy--properties--origin_pools.md)
- [origin_pools.pools](resources--bigip_http_proxy--properties--origin_pools--pools.md)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md)
- origin_pools.pools.origin_servers.origin_servers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Origin Servers. List of origin servers for Proxy.

Upstream description:

List of origin servers for Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("k8s_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_ip"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_name"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
```

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

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [k8s_service](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service.md): complete subsection reference.

- [private_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md): complete subsection reference.

- [public_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--public_ip.md): complete subsection reference.

- [public_name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--public_name.md): complete subsection reference.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md)
- [origin_pools.pools.origin_servers.origin_servers.public_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--public_ip.md)
- [origin_pools.pools.origin_servers.origin_servers.public_name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--public_name.md)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
