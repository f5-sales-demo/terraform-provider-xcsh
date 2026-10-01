---
page_title: "default_pool.origin_servers.k8s_service.outside_network"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.k8s_service.outside_network for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1396, "body_sha256": "sha256:817c30d90476a99e94e837585170d373402a904cfa0e3f5fc49f7a73c03ee56c", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:outside_network", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:outside_network", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--outside_network.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "k8s_service", "outside_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/outside_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.k8s_service.outside_network for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.k8s_service.outside_network

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md)
- default_pool.origin_servers.k8s_service.outside_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
