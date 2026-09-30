---
page_title: "service.deploy_options"
subcategory: "Container"
description: "service.deploy_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3953, "body_sha256": "sha256:7441e9242f790b357269b836f71fa7934d36f7edb356737272990f6d0fbb62ae", "canonical_id": "xcsh-docs:resources:workload:properties:service:deploy_options", "child_ids": ["xcsh-docs:resources:workload:properties:service:deploy_options:all_res", "xcsh-docs:resources:workload:properties:service:deploy_options:default_virtual_sites", "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_ce_sites", "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_ce_virtual_sites", "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_sites", "xcsh-docs:resources:workload:properties:service:deploy_options:deploy_re_virtual_sites"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:deploy_options", "parent_id": "xcsh-docs:resources:workload:properties:service", "path": "docs/guides/resources--workload--properties--service--deploy_options.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "deploy_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/deploy_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.deploy_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.deploy_options

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- service.deploy_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_res",
    "default_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_re_sites",
    "deploy_re_virtual_sites")}
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
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

Terraform syntax:

```terraform
deploy_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_res](resources--workload--properties--service--deploy_options--all_res.md): complete subsection reference.

- [default_virtual_sites](resources--workload--properties--service--deploy_options--default_virtual_sites.md): complete subsection reference.

- [deploy_ce_sites](resources--workload--properties--service--deploy_options--deploy_ce_sites.md): complete subsection reference.

- [deploy_ce_virtual_sites](resources--workload--properties--service--deploy_options--deploy_ce_virtual_sites.md): complete subsection reference.

- [deploy_re_sites](resources--workload--properties--service--deploy_options--deploy_re_sites.md): complete subsection reference.

- [deploy_re_virtual_sites](resources--workload--properties--service--deploy_options--deploy_re_virtual_sites.md): complete subsection reference.

## Next pages

- [service.deploy_options.all_res](resources--workload--properties--service--deploy_options--all_res.md)
- [service.deploy_options.default_virtual_sites](resources--workload--properties--service--deploy_options--default_virtual_sites.md)
- [service.deploy_options.deploy_ce_sites](resources--workload--properties--service--deploy_options--deploy_ce_sites.md)
- [service.deploy_options.deploy_ce_virtual_sites](resources--workload--properties--service--deploy_options--deploy_ce_virtual_sites.md)
- [service.deploy_options.deploy_re_sites](resources--workload--properties--service--deploy_options--deploy_re_sites.md)
- [service.deploy_options.deploy_re_virtual_sites](resources--workload--properties--service--deploy_options--deploy_re_virtual_sites.md)
- [service](resources--workload--properties--service.md)
- [xcsh_workload](../resources/workload.md)
