---
page_title: "job.volumes.empty_dir"
subcategory: "Container"
description: "job.volumes.empty_dir for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1748, "body_sha256": "sha256:fbc2d5ef131628e073aada31afa126d222577c64318ed7097f5391bbcb2822e2", "canonical_id": "xcsh-docs:data-sources:workload:properties:job:volumes:empty_dir", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:volumes:empty_dir:mount"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:volumes:empty_dir", "parent_id": "xcsh-docs:data-sources:workload:properties:job:volumes", "path": "docs/guides/data-sources--workload--properties--job--volumes--empty_dir.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "volumes", "empty_dir"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/volumes/empty_dir/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.volumes.empty_dir for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes.empty_dir

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [job](data-sources--workload--properties--job.md)
- [job.volumes](data-sources--workload--properties--job--volumes.md)
- job.volumes.empty_dir

<a id="section"></a>

Type: `"single"`. Computed.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

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

## Direct properties

- [mount](data-sources--workload--properties--job--volumes--empty_dir--mount.md): complete subsection reference.

<a id="schema-job--volumes--empty_dir--size_limit"></a>

### size_limit property

Type: `"number"`. Computed.

Size Limit (in GiB). Configuration parameter for size limit

Upstream description:

Configuration parameter for size limit

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [job.volumes.empty_dir.mount](data-sources--workload--properties--job--volumes--empty_dir--mount.md)
- [job.volumes](data-sources--workload--properties--job--volumes.md)
- [xcsh_workload](../data-sources/workload.md)
