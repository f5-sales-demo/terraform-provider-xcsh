---
page_title: "default_pool.origin_servers.k8s_service.vk8s_networks"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.k8s_service.vk8s_networks for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1289, "body_sha256": "sha256:ae2c22889e8bb4bd1a86f18996c08185f8c715893035a8ec82ed94f9eef89d00", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:vk8s_networks", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:vk8s_networks", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--vk8s_networks.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "k8s_service", "vk8s_networks"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/vk8s_networks/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.k8s_service.vk8s_networks for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_pool.origin_servers.k8s_service.vk8s_networks

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md)
- default_pool.origin_servers.k8s_service.vk8s_networks

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vk8s networks.

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
vk8s_networks = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
