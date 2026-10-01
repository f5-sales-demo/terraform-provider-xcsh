---
page_title: "routes.custom_route_object"
subcategory: "Load Balancing"
description: "routes.custom_route_object for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1839, "body_sha256": "sha256:a5cbabaf10fa4b7b2ef7a783d497b500d9d380bef7d403d946b1223d3bdbb15a", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object:caching_disable", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object:route_ref"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes", "path": "docs/guides/data-sources--http_loadbalancer--properties--routes--custom_route_object.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "custom_route_object"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/custom_route_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.custom_route_object for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom_route_object

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- routes.custom_route_object

<a id="section"></a>

Type: `"single"`. Computed.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

## Direct properties

- [caching_disable](data-sources--http_loadbalancer--properties--routes--custom_route_object--caching_disable.md): complete subsection reference.

- [caching_inherit](data-sources--http_loadbalancer--properties--routes--custom_route_object--caching_inherit.md): complete subsection reference.

- [route_ref](data-sources--http_loadbalancer--properties--routes--custom_route_object--route_ref.md): complete subsection reference.

## Next pages

- [routes.custom_route_object.caching_disable](data-sources--http_loadbalancer--properties--routes--custom_route_object--caching_disable.md)
- [routes.custom_route_object.caching_inherit](data-sources--http_loadbalancer--properties--routes--custom_route_object--caching_inherit.md)
- [routes.custom_route_object.route_ref](data-sources--http_loadbalancer--properties--routes--custom_route_object--route_ref.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
