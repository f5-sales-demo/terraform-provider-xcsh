---
page_title: "job.volumes.persistent_volume"
subcategory: "Container"
description: "job.volumes.persistent_volume for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1337, "body_sha256": "sha256:1614a3cf33f0d339b654af87fe3312bc82967896218dbebd0f033658664ed5d5", "canonical_id": "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume:mount", "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume:storage"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume", "parent_id": "xcsh-docs:data-sources:workload:properties:job:volumes", "path": "docs/guides/data-sources--workload--properties--job--volumes--persistent_volume.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "volumes", "persistent_volume"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/volumes/persistent_volume/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.volumes.persistent_volume for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes.persistent_volume

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [job](data-sources--workload--properties--job.md)
- [job.volumes](data-sources--workload--properties--job--volumes.md)
- job.volumes.persistent_volume

<a id="section"></a>

Type: `"single"`. Computed.

Volume containing the Persistent Storage for the workload.

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

- [mount](data-sources--workload--properties--job--volumes--persistent_volume--mount.md): complete subsection reference.

- [storage](data-sources--workload--properties--job--volumes--persistent_volume--storage.md): complete subsection reference.

## Next pages

- [job.volumes.persistent_volume.mount](data-sources--workload--properties--job--volumes--persistent_volume--mount.md)
- [job.volumes.persistent_volume.storage](data-sources--workload--properties--job--volumes--persistent_volume--storage.md)
- [job.volumes](data-sources--workload--properties--job--volumes.md)
- [xcsh_workload](../data-sources/workload.md)
