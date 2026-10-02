---
page_title: "origin_pools_weights"
subcategory: ""
description: "Origin pools with weights and priorities used for this load balancer."
xcsh_docs: {"aliases": ["backend servers", "origin pools weights", "origin servers", "upstream servers"], "body_bytes": 4988, "body_sha256": "sha256:37afcaa9b6b842dee8f5104ab73f7d7ca16b3980aae4a513e493cd9979fb012d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:cluster", "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:endpoint_subsets", "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:pool"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights", "parent_id": "xcsh-docs:resources:udp_loadbalancer:reference", "path": "documentation/resources/udp_loadbalancer/properties/origin_pools_weights/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303", "registry_path": "docs/guides/resources--udp_loadbalancer--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools_weights:ConflictingListObjectAttributes:cluster,pool", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools_weights:ConflictingListObjectAttributes:cluster,pool", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:pool", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pools_weights"], "schema_version": 1, "sections": [{"aliases": ["cluster"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:cluster", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_pools_weights--cluster--name", "enforcement": "provider-schema", "group": "origin_pools_weights.cluster:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:cluster", "type": "requires"}], "schema_path": ["origin_pools_weights", "cluster"], "syntax": "block", "type": "object"}, {"aliases": ["endpoint subsets"], "anchor": "section", "description": "Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8s or Consul cluster, the label of", "document_id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:endpoint_subsets", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pools_weights", "endpoint_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["pool"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_pools_weights--pool--name", "enforcement": "provider-schema", "group": "origin_pools_weights.pool:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights:pool", "type": "requires"}], "schema_path": ["origin_pools_weights", "pool"], "syntax": "block", "type": "object"}, {"aliases": ["priority"], "anchor": "schema-origin_pools_weights--priority", "description": "Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool as lowest priority origin pool Priority of 1 means highest priority and is considered active. When active origin pool is not available, lower priority origin pools are made active as per the increasing priority.", "document_id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools_weights", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["weight"], "anchor": "schema-origin_pools_weights--weight", "description": "Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.", "document_id": "xcsh-docs:resources:udp_loadbalancer:properties:origin_pools_weights", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools_weights", "weight"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/origin_pools_weights/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Origin pools with weights and priorities used for this load balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools_weights

Breadcrumbs:

- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/)
- origin_pools_weights

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Origin pools with weights and priorities used for this load balancer.

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

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/origin_pools_weights/cluster/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/origin_pools_weights/endpoint_subsets/): complete subsection reference.

- [pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/origin_pools_weights/pool/): complete subsection reference.

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

- [origin_pools_weights.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/origin_pools_weights/cluster/)
- [origin_pools_weights.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/origin_pools_weights/endpoint_subsets/)
- [origin_pools_weights.pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/origin_pools_weights/pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/)
- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/)
