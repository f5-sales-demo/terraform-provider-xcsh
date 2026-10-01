---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4206, "body_sha256": "sha256:1f4d6c4dc222566fef3111f318f7e66dfaabedc0c1544104595420cff267ca5a", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:simple_route"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](data-sources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--properties--service--advertise_options--advertise_custom--ports.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="section"></a>

Type: `"list"`. Computed.

Routes. Routes for this loadbalancer.

Upstream description:

Routes for this loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

## Direct properties

- [custom_route_object](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object.md): complete subsection reference.

- [direct_response_route](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route.md): complete subsection reference.

- [redirect_route](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route.md): complete subsection reference.

- [simple_route](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--simple_route.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes.md)
- [xcsh_workload](../data-sources/workload.md)
