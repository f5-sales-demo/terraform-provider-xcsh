---
page_title: "default_pool_list.pools"
subcategory: "Load Balancing"
description: "List of Origin Pools."
xcsh_docs: {"aliases": ["default pool list pools"], "body_bytes": 4072, "body_sha256": "sha256:fb300e157be60256815abc28ba0aa51b95947ce269cb232709eb01acd272fc6a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:cluster", "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:endpoint_subsets", "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:pool"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list", "path": "documentation/resources/http_loadbalancer/properties/default_pool_list/pools/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-017.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool_list.pools:ConflictingListObjectAttributes:cluster,pool", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool_list.pools:ConflictingListObjectAttributes:cluster,pool", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:pool", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool_list", "pools"], "schema_version": 1, "sections": [{"aliases": ["default pool list pools cluster"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:cluster", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_pool_list--pools--cluster--name", "enforcement": "provider-schema", "group": "default_pool_list.pools.cluster:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:cluster", "type": "requires"}], "schema_path": ["default_pool_list", "pools", "cluster"], "syntax": "block", "type": "object"}, {"aliases": ["default pool list pools endpoint subsets"], "anchor": "section", "description": "Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8s or Consul cluster, the label of", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:endpoint_subsets", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool_list", "pools", "endpoint_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["default pool list pools pool"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_pool_list--pools--pool--name", "enforcement": "provider-schema", "group": "default_pool_list.pools.pool:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools:pool", "type": "requires"}], "schema_path": ["default_pool_list", "pools", "pool"], "syntax": "block", "type": "object"}, {"aliases": ["default pool list pools priority"], "anchor": "schema-default_pool_list--pools--priority", "description": "Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool_list", "pools", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["default pool list pools weight"], "anchor": "schema-default_pool_list--pools--weight", "description": "Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool_list:pools", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool_list", "pools", "weight"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool_list/pools/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of Origin Pools.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool_list.pools

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool_list/)
- default_pool_list.pools

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools. List of Origin Pools.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool_list/pools/cluster/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool_list/pools/endpoint_subsets/): complete subsection reference.

- [pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool_list/pools/pool/): complete subsection reference.

<a id="schema-default_pool_list--pools--priority"></a>

### priority property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
