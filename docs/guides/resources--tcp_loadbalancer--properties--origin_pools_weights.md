---
page_title: "origin_pools_weights"
subcategory: "Load Balancing"
description: "origin_pools_weights for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4464, "body_sha256": "sha256:e08c876be66300b36bb5f84b91af0ef0b306132949359ab96d0480b3be4a1aaf", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:origin_pools_weights", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:origin_pools_weights:cluster", "xcsh-docs:resources:tcp_loadbalancer:properties:origin_pools_weights:endpoint_subsets", "xcsh-docs:resources:tcp_loadbalancer:properties:origin_pools_weights:pool"], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:origin_pools_weights", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:reference", "path": "docs/guides/resources--tcp_loadbalancer--properties--origin_pools_weights.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools_weights"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/origin_pools_weights/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools_weights for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools_weights

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- origin_pools_weights

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Origin pools and weights used for this load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
origin_pools_weights {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster](resources--tcp_loadbalancer--properties--origin_pools_weights--cluster.md): complete subsection reference.

- [endpoint_subsets](resources--tcp_loadbalancer--properties--origin_pools_weights--endpoint_subsets.md): complete subsection reference.

- [pool](resources--tcp_loadbalancer--properties--origin_pools_weights--pool.md): complete subsection reference.

<a id="schema-origin_pools_weights--priority"></a>

### priority property

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

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

<a id="schema-origin_pools_weights--weight"></a>

### weight property

Type: `"number"`. Optional.

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

- [origin_pools_weights.cluster](resources--tcp_loadbalancer--properties--origin_pools_weights--cluster.md)
- [origin_pools_weights.endpoint_subsets](resources--tcp_loadbalancer--properties--origin_pools_weights--endpoint_subsets.md)
- [origin_pools_weights.pool](resources--tcp_loadbalancer--properties--origin_pools_weights--pool.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
