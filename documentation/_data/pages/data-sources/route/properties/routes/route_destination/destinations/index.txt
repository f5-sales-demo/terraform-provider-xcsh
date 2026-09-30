---
page_title: "routes.route_destination.destinations"
subcategory: ""
description: "routes.route_destination.destinations for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 6512, "body_sha256": "sha256:510599219ce696fd30a25a4c354d55d47825aeb9b6654f1b9ce36a71687ca3e8", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:route_destination:destinations:cluster", "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations:endpoint_subsets"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "documentation/data-sources/route/properties/routes/route_destination/destinations/index.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["routes", "route_destination", "destinations"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/destinations/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.destinations for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/destinations/cluster/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/destinations/endpoint_subsets/): complete subsection reference.

<a id="schema-routes--route_destination--destinations--priority"></a>

### priority property

Type: `"number"`. Computed.

Priority of this cluster, valid only with multiple destinations are configured. Value of 0 will make
the cluster as lowest priority upstream cluster Priority of 1 means highest priority and is
considered active. When active cluster is not available, lower priority clusters are made active as
per..

Upstream description:

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

Type: `"number"`. Computed.

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

- [routes.route_destination.destinations.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/destinations/cluster/)
- [routes.route_destination.destinations.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/destinations/endpoint_subsets/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
