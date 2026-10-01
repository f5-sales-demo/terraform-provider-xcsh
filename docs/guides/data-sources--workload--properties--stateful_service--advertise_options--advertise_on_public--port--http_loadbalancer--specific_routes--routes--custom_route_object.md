---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4137, "body_sha256": "sha256:b75e5a685ebcc853c4e66f9a1affa44e0a9d0ffa737b4b74a080044df3fed58c", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object:caching_disable", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object:caching_inherit", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object:route_ref"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes", "routes", "custom_route_object"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/custom_route_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes.md)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object

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

- [caching_disable](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--caching_disable.md): complete subsection reference.

- [caching_inherit](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--caching_inherit.md): complete subsection reference.

- [route_ref](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--route_ref.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--caching_disable.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--caching_inherit.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--route_ref.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes.md)
- [xcsh_workload](../data-sources/workload.md)
