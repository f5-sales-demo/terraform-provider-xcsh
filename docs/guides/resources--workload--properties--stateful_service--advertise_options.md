---
page_title: "stateful_service.advertise_options"
subcategory: "Container"
description: "stateful_service.advertise_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2942, "body_sha256": "sha256:e423ff771eaadf5a781b4ae653d41755784fc98f6d75228d235870a1998d6ff7", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:do_not_advertise"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- stateful_service.advertise_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_in_cluster"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
advertise_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_custom](resources--workload--properties--stateful_service--advertise_options--advertise_custom.md): complete subsection reference.

- [advertise_in_cluster](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster.md): complete subsection reference.

- [advertise_on_public](resources--workload--properties--stateful_service--advertise_options--advertise_on_public.md): complete subsection reference.

- [do_not_advertise](resources--workload--properties--stateful_service--advertise_options--do_not_advertise.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom](resources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster.md)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.do_not_advertise](resources--workload--properties--stateful_service--advertise_options--do_not_advertise.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [xcsh_workload](../resources/workload.md)
