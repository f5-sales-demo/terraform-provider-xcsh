---
page_title: "job.configuration.parameters"
subcategory: "Container"
description: "job.configuration.parameters for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1961, "body_sha256": "sha256:7cd3cc84e1be7249f563d571454b7a4509de31348c0ccc8a10957aaa1dbf6062", "canonical_id": "xcsh-docs:data-sources:workload:properties:job:configuration:parameters", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:configuration:parameters:env_var", "xcsh-docs:data-sources:workload:properties:job:configuration:parameters:file"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:configuration:parameters", "parent_id": "xcsh-docs:data-sources:workload:properties:job:configuration", "path": "docs/guides/data-sources--workload--properties--job--configuration--parameters.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "configuration", "parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/configuration/parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.configuration.parameters for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.configuration.parameters

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [job](data-sources--workload--properties--job.md)
- [job.configuration](data-sources--workload--properties--job--configuration.md)
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

- [env_var](data-sources--workload--properties--job--configuration--parameters--env_var.md): complete subsection reference.

- [file](data-sources--workload--properties--job--configuration--parameters--file.md): complete subsection reference.

## Next pages

- [job.configuration.parameters.env_var](data-sources--workload--properties--job--configuration--parameters--env_var.md)
- [job.configuration.parameters.file](data-sources--workload--properties--job--configuration--parameters--file.md)
- [job.configuration](data-sources--workload--properties--job--configuration.md)
- [xcsh_workload](../data-sources/workload.md)
