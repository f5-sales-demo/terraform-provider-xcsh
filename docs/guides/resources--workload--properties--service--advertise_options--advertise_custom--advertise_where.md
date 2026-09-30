---
page_title: "service.advertise_options.advertise_custom.advertise_where"
subcategory: "Container"
description: "service.advertise_options.advertise_custom.advertise_where for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3292, "body_sha256": "sha256:7cd651a74b51bde143e577bbdb5e333661e293d1852fbdfe3438bdf89cb650c9", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:site", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:virtual_site", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_custom--advertise_where.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_custom.advertise_where for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_custom.advertise_where

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md)
- service.advertise_options.advertise_custom.advertise_where

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service")}
```

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

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--site.md): complete subsection reference.

- [virtual_site](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--virtual_site.md): complete subsection reference.

- [vk8s_service](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--vk8s_service.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.advertise_where.site](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--site.md)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--virtual_site.md)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--properties--service--advertise_options--advertise_custom--advertise_where--vk8s_service.md)
- [service.advertise_options.advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md)
- [xcsh_workload](../resources/workload.md)
