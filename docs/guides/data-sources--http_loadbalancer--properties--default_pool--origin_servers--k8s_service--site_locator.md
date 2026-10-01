---
page_title: "default_pool.origin_servers.k8s_service.site_locator"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.k8s_service.site_locator for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2046, "body_sha256": "sha256:e01eb1c91491b4953acea0bd968c8518063bd05604cea355c2b91660a484fd8a", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:site_locator", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:site_locator:site", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:site_locator:virtual_site"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:site_locator", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--site_locator.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "k8s_service", "site_locator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.k8s_service.site_locator for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.k8s_service.site_locator

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](data-sources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md)
- default_pool.origin_servers.k8s_service.site_locator

<a id="section"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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

## Direct properties

- [site](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--site_locator--site.md): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--site_locator--virtual_site.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.k8s_service.site_locator.site](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--site_locator--site.md)
- [default_pool.origin_servers.k8s_service.site_locator.virtual_site](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--site_locator--virtual_site.md)
- [default_pool.origin_servers.k8s_service](data-sources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
