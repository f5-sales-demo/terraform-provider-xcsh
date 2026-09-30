---
page_title: "stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2460, "body_sha256": "sha256:6151e855d5107ec94b4f8704251cbf9d4bc8bf1d935ba7ed504b841edda4ecce", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service:site", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service:virtual_site"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--vk8s_service.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "advertise_where", "vk8s_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/advertise_where/vk8s_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where.md)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="section"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

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

## Direct properties

- [site](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--vk8s_service--site.md): complete subsection reference.

- [virtual_site](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--vk8s_service--virtual_site.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--vk8s_service--site.md)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--vk8s_service--virtual_site.md)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where.md)
- [xcsh_workload](../data-sources/workload.md)
