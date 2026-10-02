---
page_title: "default_pool_list.pools.endpoint_subsets"
subcategory: "Load Balancing"
description: "Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8s or Consul cluster, the label of"
xcsh_docs: {"aliases": ["backend servers", "default pool list pools endpoint subsets", "origin servers", "upstream servers"], "body_bytes": 2611, "body_sha256": "sha256:71e17d8ed2418374555b16fe11713848bb25937c376d1f8936d78bae2277765f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:endpoint_subsets", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools", "path": "documentation/resources/http_loadbalancer/properties/default_pool_list/pools/endpoint_subsets/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3210332310321130-0113222213001001-2113013221322330-3212101123033102-3001002230002303-2132202333322300-1201201130002212-1120123101133323", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool_list", "pools", "endpoint_subsets"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool_list/pools/endpoint_subsets/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8s or Consul cluster, the label of", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool_list.pools.endpoint_subsets

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool_list/)
- [default_pool_list.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool_list/pools/)
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

- [default_pool_list.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool_list/pools/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
