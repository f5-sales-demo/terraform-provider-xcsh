---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3894, "body_sha256": "sha256:e5b8700f963c5b9cddc7e665b76db6346df3f6943858a55f5d2901ab7cad855d", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object:caching_disable", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object:caching_inherit", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object:route_ref"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:custom_route_object", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes", "routes", "custom_route_object"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/custom_route_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](data-sources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object

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

- [caching_disable](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--caching_disable.md): complete subsection reference.

- [caching_inherit](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--caching_inherit.md): complete subsection reference.

- [route_ref](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--route_ref.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--caching_disable.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--caching_inherit.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--custom_route_object--route_ref.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes.md)
- [xcsh_workload](../data-sources/workload.md)
