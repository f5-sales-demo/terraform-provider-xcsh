---
page_title: "job.configuration.parameters"
subcategory: "Container"
description: "job.configuration.parameters for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 2360, "body_sha256": "sha256:5064c7a981f7c2a15d262f9e12d816dbb648726cd2fb4ddc3daea870e9c8d1f8", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:configuration:parameters:env_var", "xcsh-docs:data-sources:workload:properties:job:configuration:parameters:file"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:configuration:parameters", "parent_id": "xcsh-docs:data-sources:workload:properties:job:configuration", "path": "documentation/data-sources/workload/properties/job/configuration/parameters/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["job", "configuration", "parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/configuration/parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.configuration.parameters for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# job.configuration.parameters

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/)
- job.configuration.parameters

<a id="section"></a>

Type: `"list"`. Computed.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/env_var/): complete subsection reference.

- [file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/file/): complete subsection reference.

## Next pages

- [job.configuration.parameters.env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/env_var/)
- [job.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/parameters/file/)
- [job.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/configuration/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
