---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2620, "body_sha256": "sha256:747ee6603b5d8566b7e52ed2bfebb3bc70cb9c8b2083f957eb8bbdea514bed26", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_disable", "child_ids": [], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object:caching_disable", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:custom_route_object", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object--caching_disable.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "custom_route_object", "caching_disable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/custom_route_object/caching_disable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object.md)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
caching_disable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--properties--service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--custom_route_object.md)
- [xcsh_workload](../resources/workload.md)
