---
page_title: "job.deploy_options"
subcategory: "Container"
description: "job.deploy_options for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4801, "body_sha256": "sha256:78352ac849f670e99a103ecd99e9b986f6ee8ac91a831416ea0b80a482e66198", "child_ids": ["xcsh-docs:resources:workload:properties:job:deploy_options:all_res", "xcsh-docs:resources:workload:properties:job:deploy_options:default_virtual_sites", "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_sites", "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_ce_virtual_sites", "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_sites", "xcsh-docs:resources:workload:properties:job:deploy_options:deploy_re_virtual_sites"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:deploy_options", "parent_id": "xcsh-docs:resources:workload:properties:job", "path": "documentation/resources/workload/properties/job/deploy_options/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["job", "deploy_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/deploy_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.deploy_options for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.deploy_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- job.deploy_options

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

- [all_res](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/all_res/): complete subsection reference.

- [default_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/default_virtual_sites/): complete subsection reference.

- [deploy_ce_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_ce_sites/): complete subsection reference.

- [deploy_ce_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/): complete subsection reference.

- [deploy_re_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_re_sites/): complete subsection reference.

- [deploy_re_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_re_virtual_sites/): complete subsection reference.

## Next pages

- [job.deploy_options.all_res](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/all_res/)
- [job.deploy_options.default_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/default_virtual_sites/)
- [job.deploy_options.deploy_ce_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_ce_sites/)
- [job.deploy_options.deploy_ce_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_ce_virtual_sites/)
- [job.deploy_options.deploy_re_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_re_sites/)
- [job.deploy_options.deploy_re_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/deploy_options/deploy_re_virtual_sites/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
