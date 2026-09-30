---
page_title: "origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator"
subcategory: ""
description: "origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2691, "body_sha256": "sha256:56af96b589b79fd762380392aace8bf23dec79a07e16621d51de8df9a559292e", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:site", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator:virtual_site"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip", "path": "docs/guides/resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "site_locator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [origin_pools](resources--bigip_http_proxy--properties--origin_pools.md)
- [origin_pools.pools](resources--bigip_http_proxy--properties--origin_pools--pools.md)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site.md): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site.md): complete subsection reference.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
