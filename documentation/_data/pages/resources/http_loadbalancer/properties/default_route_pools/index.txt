---
page_title: "default_route_pools"
subcategory: "Load Balancing"
description: "Origin Pools used when no route is specified (default route)"
xcsh_docs: {"aliases": ["default route pools"], "body_bytes": 5059, "body_sha256": "sha256:72fd05d4c094f4eb4f715c62007ab1c5c254b6f7d95a391554245b9b09f11c4c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:cluster", "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:endpoint_subsets", "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:pool"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/default_route_pools/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_route_pools:ConflictingListObjectAttributes:cluster,pool", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_route_pools:ConflictingListObjectAttributes:cluster,pool", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:pool", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_route_pools"], "schema_version": 1, "sections": [{"aliases": ["default route pools cluster"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:cluster", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_route_pools--cluster--name", "enforcement": "provider-schema", "group": "default_route_pools.cluster:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:cluster", "type": "requires"}], "schema_path": ["default_route_pools", "cluster"], "syntax": "block", "type": "object"}, {"aliases": ["default route pools endpoint subsets"], "anchor": "section", "description": "Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8s or Consul cluster, the label of", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:endpoint_subsets", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_route_pools", "endpoint_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["default route pools pool"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_route_pools--pool--name", "enforcement": "provider-schema", "group": "default_route_pools.pool:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools:pool", "type": "requires"}], "schema_path": ["default_route_pools", "pool"], "syntax": "block", "type": "object"}, {"aliases": ["default route pools priority"], "anchor": "schema-default_route_pools--priority", "description": "Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_route_pools", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["default route pools weight"], "anchor": "schema-default_route_pools--weight", "description": "Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_route_pools", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_route_pools", "weight"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_route_pools/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Origin Pools used when no route is specified (default route)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_route_pools

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- default_route_pools

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools used when no route is specified (default route).

Upstream description:

Origin Pools used when no route is specified (default route)

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
default_route_pools {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_route_pools/cluster/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_route_pools/endpoint_subsets/): complete subsection reference.

- [pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_route_pools/pool/): complete subsection reference.

<a id="schema-default_route_pools--priority"></a>

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

<a id="schema-default_route_pools--weight"></a>

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

- [default_route_pools.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_route_pools/cluster/)
- [default_route_pools.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_route_pools/endpoint_subsets/)
- [default_route_pools.pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_route_pools/pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
