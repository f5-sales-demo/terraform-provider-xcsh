---
page_title: "service.advertise_options"
subcategory: "Container"
description: "service.advertise_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2681, "body_sha256": "sha256:da87854586c8d4b9b86d0b02f741fa85e61b3bb9f2227e98875f11b908b3894e", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public", "xcsh-docs:resources:workload:properties:service:advertise_options:do_not_advertise"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options", "parent_id": "xcsh-docs:resources:workload:properties:service", "path": "docs/guides/resources--workload--properties--service--advertise_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- service.advertise_options

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

- [advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md): complete subsection reference.

- [advertise_in_cluster](resources--workload--properties--service--advertise_options--advertise_in_cluster.md): complete subsection reference.

- [advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md): complete subsection reference.

- [do_not_advertise](resources--workload--properties--service--advertise_options--do_not_advertise.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom](resources--workload--properties--service--advertise_options--advertise_custom.md)
- [service.advertise_options.advertise_in_cluster](resources--workload--properties--service--advertise_options--advertise_in_cluster.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.do_not_advertise](resources--workload--properties--service--advertise_options--do_not_advertise.md)
- [service](resources--workload--properties--service.md)
- [xcsh_workload](../resources/workload.md)
