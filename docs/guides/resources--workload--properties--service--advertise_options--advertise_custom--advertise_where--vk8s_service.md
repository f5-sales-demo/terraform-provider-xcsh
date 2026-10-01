---
page_title: "service.advertise_options.advertise_custom.advertise_where.vk8s_service"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.advertise_where.vk8s_service for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2648, "body_sha256": "sha256:a638cc5669f2038628a6e8e98c03d29f138d7cbd10f162cdf78547733676d568", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:site", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service:virtual_site"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--vk8s_service.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where", "vk8s_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/vk8s_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.advertise_where.vk8s_service for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.advertise_where.vk8s_service

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where.md)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--vk8s_service--site.md): complete subsection reference.

- [virtual_site](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--vk8s_service--virtual_site.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--vk8s_service--site.md)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--vk8s_service--virtual_site.md)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where.md)
- [xcsh_workload](../resources/workload.md)
