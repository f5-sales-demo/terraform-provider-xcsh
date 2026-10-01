---
page_title: "default_pool.origin_servers.k8s_service.snat_pool"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.k8s_service.snat_pool for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1989, "body_sha256": "sha256:2372bc6f70e1ef2fe242288bee47405adec745231a83485671fd419da78e42ce", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:no_snat_pool", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:snat_pool"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--snat_pool.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "k8s_service", "snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.k8s_service.snat_pool for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.k8s_service.snat_pool

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](data-sources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md)
- default_pool.origin_servers.k8s_service.snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

## Direct properties

- [no_snat_pool](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--snat_pool--no_snat_pool.md): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--snat_pool--snat_pool.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--snat_pool--no_snat_pool.md)
- [default_pool.origin_servers.k8s_service.snat_pool.snat_pool](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--snat_pool--snat_pool.md)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
