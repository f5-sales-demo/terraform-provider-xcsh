---
page_title: "routes.route_destination.destinations.endpoint_subsets"
subcategory: ""
description: "routes.route_destination.destinations.endpoint_subsets for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2298, "body_sha256": "sha256:bc4eb2ec1e87ad80584bf09f1f8d856a44042ca35bba5364a12b9305af653a75", "canonical_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations:endpoint_subsets", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations:endpoint_subsets", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations", "path": "docs/guides/resources--route--properties--routes--route_destination--destinations--endpoint_subsets.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_destination", "destinations", "endpoint_subsets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/destinations/endpoint_subsets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_destination.destinations.endpoint_subsets for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.route_destination.destinations.endpoint_subsets

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.route_destination](resources--route--properties--routes--route_destination.md)
- [routes.route_destination.destinations](resources--route--properties--routes--route_destination--destinations.md)
- routes.route_destination.destinations.endpoint_subsets

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached
to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be
selected by the load balancer Labels field of endpoint object's metadata is used for subset..

Upstream description:

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

Terraform syntax:

```terraform
endpoint_subsets {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.route_destination.destinations](resources--route--properties--routes--route_destination--destinations.md)
- [xcsh_route](../resources/route.md)
