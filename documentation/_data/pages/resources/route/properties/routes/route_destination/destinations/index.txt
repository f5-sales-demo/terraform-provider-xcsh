---
page_title: "routes.route_destination.destinations"
subcategory: ""
description: "When requests have to distributed among multiple upstream clusters, multiple destinations are configured, each having its own cluster and weight. Traffic is distributed among clusters based on the weight configured. Example: destinations: - cluster: - kind: F5 xc.vega.cfg.adc.cluster.object uid: cluster-1 weight: 20 -"
xcsh_docs: {"aliases": ["routes route destination destinations"], "body_bytes": 7119, "body_sha256": "sha256:8746f34f0076e636ab60eca4ef984fa84bccc897b14bad9c4ae4f2933475b300", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:destinations:cluster", "xcsh-docs:resources:route:properties:routes:route_destination:destinations:endpoint_subsets"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/destinations/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2033310032010003-2013013212023023-3331003021220120-3211312230110102-0332232123000011-2322023233233003-2010313220111220-3203133211033012", "registry_path": "docs/guides/resources--route--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.destinations:RequiredListObjectAttributes:cluster", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations:cluster", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "destinations"], "schema_version": 1, "sections": [{"aliases": ["cluster"], "anchor": "section", "description": "Indicates the upstream cluster to which the request should be sent. If the cluster does not exist ServiceUnavailable response will be sent.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations:cluster", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "route_destination", "destinations", "cluster"], "syntax": "block", "type": "object"}, {"aliases": ["endpoint subsets"], "anchor": "section", "description": "Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer Labels field of endpoint object's metadata is used for subset matching. For endpoints which", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations:endpoint_subsets", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "destinations", "endpoint_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["priority"], "anchor": "schema-routes--route_destination--destinations--priority", "description": "Priority of this cluster, valid only with multiple destinations are configured. Value of 0 will make the cluster as lowest priority upstream cluster Priority of 1 means highest priority and is considered active. When active cluster is not available, lower priority clusters are made active as per the increasing", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "destinations", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["weight"], "anchor": "schema-routes--route_destination--destinations--weight", "description": "When requests have to distributed among multiple upstream clusters, multiple destinations are configured, each having its own cluster and weight. Traffic is distributed among clusters based on the weight configured. Example: destinations: - cluster: - kind: F5 xc.vega.cfg.adc.cluster.object uid: cluster-1 weight: 20 -", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "destinations", "weight"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/destinations/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "When requests have to distributed among multiple upstream clusters, multiple destinations are configured, each having its own cluster and weight. Traffic is distributed among clusters based on the weight configured. Example: destinations: - cluster: - kind: F5 xc.vega.cfg.adc.cluster.object uid: cluster-1 weight: 20 -", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.destinations

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- routes.route_destination.destinations

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

When requests have to distributed among multiple upstream clusters, multiple destinations are
configured, each having its own cluster and weight. Traffic is distributed among clusters based on
the weight configured.

Upstream description:

When requests have to distributed among multiple upstream clusters, multiple destinations are
configured, each having its own cluster and weight. Traffic is distributed among clusters based on
the weight configured.

Example: destinations: &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-1 weight: 20 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-2 weight: 30 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-3 weight: 50

This indicates that out of every 100 requests, 50 goes to cluster-3, 30 to cluster-2 and 20 to
cluster-1

When single destination is configured, weight is ignored. All the requests are sent to the cluster
specified in the destination.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("cluster")}
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
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
destinations {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/cluster/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/endpoint_subsets/): complete subsection reference.

<a id="schema-routes--route_destination--destinations--priority"></a>

### priority property

Type: `"number"`. Optional.

Priority of this cluster, valid only with multiple destinations are configured. Value of 0 will make
the cluster as lowest priority upstream cluster Priority of 1 means highest priority and is
considered active. When active cluster is not available, lower priority clusters are made active as
per..

Upstream description:

Priority of this cluster, valid only with multiple destinations are configured. Value of 0 will make
the cluster as lowest priority upstream cluster Priority of 1 means highest priority and is
considered active. When active cluster is not available, lower priority clusters are made active as
per the increasing priority.

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

<a id="schema-routes--route_destination--destinations--weight"></a>

### weight property

Type: `"number"`. Optional.

When requests have to distributed among multiple upstream clusters, multiple destinations are
configured, each having its own cluster and weight. Traffic is distributed among clusters based on
the weight configured.

Upstream description:

When requests have to distributed among multiple upstream clusters, multiple destinations are
configured, each having its own cluster and weight. Traffic is distributed among clusters based on
the weight configured.

Example: destinations: &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-1 weight: 20 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-2 weight: 30 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-3 weight: 10

This indicates that out of every 60 requests, 10 goes to cluster-3, 30 to cluster-2 and 20 to
cluster-1

When single destination is configured, weight is ignored. All the requests are sent to the cluster
specified in the destination.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [routes.route_destination.destinations.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/cluster/)
- [routes.route_destination.destinations.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/endpoint_subsets/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
