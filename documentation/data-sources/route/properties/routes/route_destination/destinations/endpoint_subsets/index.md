---
page_title: "routes.route_destination.destinations.endpoint_subsets"
subcategory: ""
description: "Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer Labels field of endpoint object's metadata is used for subset matching. For endpoints which"
xcsh_docs: {"aliases": ["routes route destination destinations endpoint subsets"], "body_bytes": 2329, "body_sha256": "sha256:15edb047a8da89c3ebd3183f56311f749675e4e5b8f63b9afa3a373d39b23b79", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations:endpoint_subsets", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination:destinations", "path": "documentation/data-sources/route/properties/routes/route_destination/destinations/endpoint_subsets/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1312330300302211-0331122033121010-0110312211112133-3130111000103233-1303311322212122-3331120300002312-0102002021301102-0321122312100103", "registry_path": "docs/guides/data-sources--route--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "destinations", "endpoint_subsets"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/destinations/endpoint_subsets/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer Labels field of endpoint object's metadata is used for subset matching. For endpoints which", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["routeCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.destinations.endpoint_subsets

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- [routes.route_destination.destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/destinations/)
- routes.route_destination.destinations.endpoint_subsets

<a id="section"></a>

Type: `"single"`. Computed.

Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached
to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be
selected by the load balancer

Labels field of endpoint object's metadata is used for subset matching. For endpoints which are
discovered in K8s or Consul cluster, the label of the service is merged with endpoint's labels. In
case of Consul, the label is derived from the "Tag" field. For labels that are common between
configured endpoint and discovered service, labels from discovered service takes precedence.

List of key-value pairs that will be used as matching metadata. Only those endpoints of upstream
cluster which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
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

This is an empty object or choice marker. It has no direct properties.
