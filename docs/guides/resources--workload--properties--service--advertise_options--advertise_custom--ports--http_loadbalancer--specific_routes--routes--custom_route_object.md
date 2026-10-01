---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4103, "body_sha256": "sha256:56f3e80e6c9f6d1ba11ca72378951293f829e52d11ea91a9b757ad8a9a7cb281", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_disable", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_inherit", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object:route_ref"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "custom_route_object"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/custom_route_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.ports](resources--workload--properties--service--advertise_options--advertise_custom--ports.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit")}
```

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

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

## Direct properties

- [caching_disable](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--caching_disable.md): complete subsection reference.

- [caching_inherit](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--caching_inherit.md): complete subsection reference.

- [route_ref](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--route_ref.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--caching_disable.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--caching_inherit.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--route_ref.md)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md)
- [xcsh_workload](../resources/workload.md)
