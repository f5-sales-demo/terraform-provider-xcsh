---
page_title: "routes.route_destination.destinations"
subcategory: ""
description: "When requests have to distributed among multiple upstream clusters, multiple destinations are configured, each having its own cluster and weight. Traffic is distributed among clusters based on the weight configured. Example: destinations: - cluster: - kind: F5 xc.vega.cfg.adc.cluster.object uid: cluster-1 weight: 20 -"
xcsh_docs: {"aliases": ["routes route destination destinations"], "body_bytes": 5172, "body_sha256": "sha256:3e0234e7a5bcba40d2523dbc0d9c4274408c7e42b3e3b1be012283d9d1655b1b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:route_destination:destinations:cluster", "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations:endpoint_subsets"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "documentation/data-sources/route/properties/routes/route_destination/destinations/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0030003212020121-1223121010102011-2322312132003222-2210220313332211-1221133210032101-1103010232121021-2321313022100233-2111112303331310", "registry_path": "docs/guides/data-sources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "destinations"], "schema_version": 1, "sections": [{"aliases": ["routes route destination destinations cluster"], "anchor": "section", "description": "Indicates the upstream cluster to which the request should be sent. If the cluster does not exist ServiceUnavailable response will be sent.", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations:cluster", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "route_destination", "destinations", "cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination destinations endpoint subsets"], "anchor": "section", "description": "Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer Labels field of endpoint object's metadata is used for subset matching. For endpoints which", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations:endpoint_subsets", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "destinations", "endpoint_subsets"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination destinations priority"], "anchor": "schema-routes--route_destination--destinations--priority", "description": "Priority of this cluster, valid only with multiple destinations are configured. Value of 0 will make the cluster as lowest priority upstream cluster Priority of 1 means highest priority and is considered active. When active cluster is not available, lower priority clusters are made active as per the increasing", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "destinations", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["routes route destination destinations weight"], "anchor": "schema-routes--route_destination--destinations--weight", "description": "When requests have to distributed among multiple upstream clusters, multiple destinations are configured, each having its own cluster and weight. Traffic is distributed among clusters based on the weight configured. Example: destinations: - cluster: - kind: F5 xc.vega.cfg.adc.cluster.object uid: cluster-1 weight: 20 -", "document_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "destinations", "weight"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/destinations/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "When requests have to distributed among multiple upstream clusters, multiple destinations are configured, each having its own cluster and weight. Traffic is distributed among clusters based on the weight configured. Example: destinations: - cluster: - kind: F5 xc.vega.cfg.adc.cluster.object uid: cluster-1 weight: 20 -", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["routeCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.destinations

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- routes.route_destination.destinations

<a id="section"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/destinations/cluster/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/destinations/endpoint_subsets/): complete subsection reference.

<a id="schema-routes--route_destination--destinations--priority"></a>

### priority property

Type: `"number"`. Computed.

Priority of this cluster, valid only with multiple destinations are configured. Value of 0 will make
the cluster as lowest priority upstream cluster Priority of 1 means highest priority and is
considered active. When active cluster is not available, lower priority clusters are made active as
per the increasing priority.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
