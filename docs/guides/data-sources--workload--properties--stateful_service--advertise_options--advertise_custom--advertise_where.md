---
page_title: "stateful_service.advertise_options.advertise_custom.advertise_where"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.advertise_where for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3153, "body_sha256": "sha256:c64dce3991254e14b7beed481de97c7e9db243852ce636c8e47f69e85038892a", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:site", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:virtual_site", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where:vk8s_service"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "advertise_where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/advertise_where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.advertise_where for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.advertise_where

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- stateful_service.advertise_options.advertise_custom.advertise_where

<a id="section"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [site](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--site.md): complete subsection reference.

- [virtual_site](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--virtual_site.md): complete subsection reference.

- [vk8s_service](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--vk8s_service.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--site.md)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--virtual_site.md)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where--vk8s_service.md)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [xcsh_workload](../data-sources/workload.md)
