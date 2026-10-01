---
page_title: "default_pool_list.pools"
subcategory: "Load Balancing"
description: "default_pool_list.pools for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4279, "body_sha256": "sha256:153b3219782987396f5fef0be29b644c1c4906c03b8fa1e0b3634141746acb0d", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list:pools", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list:pools:cluster", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list:pools:endpoint_subsets", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list:pools:pool"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list:pools", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool_list", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool_list--pools.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool_list", "pools"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool_list/pools/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool_list.pools for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool_list.pools

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool_list](data-sources--http_loadbalancer--properties--default_pool_list.md)
- default_pool_list.pools

<a id="section"></a>

Type: `"list"`. Computed.

Origin Pools. List of Origin Pools.

Upstream description:

List of Origin Pools.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [cluster](data-sources--http_loadbalancer--properties--default_pool_list--pools--cluster.md): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--properties--default_pool_list--pools--endpoint_subsets.md): complete subsection reference.

- [pool](data-sources--http_loadbalancer--properties--default_pool_list--pools--pool.md): complete subsection reference.

<a id="schema-default_pool_list--pools--priority"></a>

### priority property

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-default_pool_list--pools--weight"></a>

### weight property

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [default_pool_list.pools.cluster](data-sources--http_loadbalancer--properties--default_pool_list--pools--cluster.md)
- [default_pool_list.pools.endpoint_subsets](data-sources--http_loadbalancer--properties--default_pool_list--pools--endpoint_subsets.md)
- [default_pool_list.pools.pool](data-sources--http_loadbalancer--properties--default_pool_list--pools--pool.md)
- [default_pool_list](data-sources--http_loadbalancer--properties--default_pool_list.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
