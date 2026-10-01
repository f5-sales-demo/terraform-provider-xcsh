---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4346, "body_sha256": "sha256:d483d98ed4d31f8644f842bb749e3ac418352166f5c3b710e625953a62e7a3db", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_disable", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_inherit", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object:route_ref"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "custom_route_object"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/custom_route_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](resources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

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

- [caching_disable](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--caching_disable.md): complete subsection reference.

- [caching_inherit](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--caching_inherit.md): complete subsection reference.

- [route_ref](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--route_ref.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--caching_disable.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--caching_inherit.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--route_ref.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md)
- [xcsh_workload](../resources/workload.md)
