---
page_title: "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool"
subcategory: ""
description: "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2636, "body_sha256": "sha256:0bdd552fc00f3760790d844582f6f39466882d164ef064ad3ad789f70f08e646", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:no_snat_pool", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:snat_pool"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip", "path": "docs/guides/resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [origin_pools](resources--bigip_http_proxy--properties--origin_pools.md)
- [origin_pools.pools](resources--bigip_http_proxy--properties--origin_pools--pools.md)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--no_snat_pool.md): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--snat_pool.md): complete subsection reference.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--no_snat_pool.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--snat_pool.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
