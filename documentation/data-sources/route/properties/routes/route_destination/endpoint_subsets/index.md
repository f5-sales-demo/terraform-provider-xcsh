---
page_title: "routes.route_destination.endpoint_subsets"
subcategory: ""
description: "Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer Labels field of endpoint object's metadata is used for subset matching. For endpoint's which"
xcsh_docs: {"aliases": ["routes route destination endpoint subsets"], "body_bytes": 2722, "body_sha256": "sha256:f5077c3ce8cd9a8e1e527ab6646aca664958479938e47b962e44b275adf64600", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:route_destination:endpoint_subsets", "parent_id": "xcsh-docs:data-sources:route:properties:routes:route_destination", "path": "documentation/data-sources/route/properties/routes/route_destination/endpoint_subsets/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2013022110111230-1123202020123012-1102201000122011-1313332231301300-0203301231132222-3220230221020003-1003012020313203-3102133133012031", "registry_path": "docs/guides/data-sources--route--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "endpoint_subsets"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/route_destination/endpoint_subsets/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer Labels field of endpoint object's metadata is used for subset matching. For endpoint's which", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.endpoint_subsets

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- routes.route_destination.endpoint_subsets

<a id="section"></a>

Type: `"single"`. Computed.

Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached
to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be
selected by the load balancer Labels field of endpoint object's metadata is used for subset..

Upstream description:

Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached
to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be
selected by the load balancer

Labels field of endpoint object's metadata is used for subset matching. For endpoint's which are
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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/route_destination/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
