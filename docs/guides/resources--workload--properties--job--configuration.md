---
page_title: "job.configuration"
subcategory: "Container"
description: "job.configuration for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 945, "body_sha256": "sha256:906528c918fe6e17d2e3539e30c6aeb58d0098e7090ec5c06d8161757316ff47", "canonical_id": "xcsh-docs:resources:workload:properties:job:configuration", "child_ids": ["xcsh-docs:resources:workload:properties:job:configuration:parameters"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:configuration", "parent_id": "xcsh-docs:resources:workload:properties:job", "path": "docs/guides/resources--workload--properties--job--configuration.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "configuration"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/configuration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.configuration for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# job.configuration

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [job](resources--workload--properties--job.md)
- job.configuration

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameters of the workload.

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
configuration {
  # Configure direct properties listed below.
}
```

## Direct properties

- [parameters](resources--workload--properties--job--configuration--parameters.md): complete subsection reference.

## Next pages

- [job.configuration.parameters](resources--workload--properties--job--configuration--parameters.md)
- [job](resources--workload--properties--job.md)
- [xcsh_workload](../resources/workload.md)
