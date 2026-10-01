---
page_title: "stateful_service.advertise_options.advertise_custom"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1926, "body_sha256": "sha256:8bb7e562c89132cfa4567bff4788afc8df81d1d74bc284810f74bfebe550db65", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_custom.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- stateful_service.advertise_options.advertise_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where",
    "ports")}
```

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
advertise_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_where](resources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where.md): complete subsection reference.

- [ports](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--properties--stateful_service--advertise_options--advertise_custom--advertise_where.md)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [xcsh_workload](../resources/workload.md)
