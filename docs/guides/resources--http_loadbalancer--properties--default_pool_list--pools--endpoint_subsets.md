---
page_title: "default_pool_list.pools.endpoint_subsets"
subcategory: "Load Balancing"
description: "default_pool_list.pools.endpoint_subsets for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2206, "body_sha256": "sha256:c9d3766b46f88ac069864c58b2d04d6a8e4bcbd9970d064b27aba13f677555be", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:endpoint_subsets", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:endpoint_subsets", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool_list--pools--endpoint_subsets.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool_list", "pools", "endpoint_subsets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool_list/pools/endpoint_subsets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool_list.pools.endpoint_subsets for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_pool_list.pools.endpoint_subsets

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool_list](resources--http_loadbalancer--properties--default_pool_list.md)
- [default_pool_list.pools](resources--http_loadbalancer--properties--default_pool_list--pools.md)
- default_pool_list.pools.endpoint_subsets

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_pool_list.pools](resources--http_loadbalancer--properties--default_pool_list--pools.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
