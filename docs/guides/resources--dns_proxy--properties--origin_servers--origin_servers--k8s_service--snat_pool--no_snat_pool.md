---
page_title: "origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool"
subcategory: ""
description: "origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1423, "body_sha256": "sha256:40c498c72da2a30867668718d314714df487c977917055d7d17128d00d7b917e", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool:no_snat_pool", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool:no_snat_pool", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool", "path": "docs/guides/resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--no_snat_pool.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "origin_servers", "k8s_service", "snat_pool", "no_snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/no_snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [origin_servers](resources--dns_proxy--properties--origin_servers.md)
- [origin_servers.origin_servers](resources--dns_proxy--properties--origin_servers--origin_servers.md)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool.md)
- origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
