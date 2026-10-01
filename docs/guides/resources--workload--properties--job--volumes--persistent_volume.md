---
page_title: "job.volumes.persistent_volume"
subcategory: "Container"
description: "job.volumes.persistent_volume for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1433, "body_sha256": "sha256:559404f2c1ecd51cf8c26fee76b99441b08f5e88a7296b54e57effb06eb36273", "canonical_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume", "child_ids": ["xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:mount", "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:storage"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume", "parent_id": "xcsh-docs:resources:workload:properties:job:volumes", "path": "docs/guides/resources--workload--properties--job--volumes--persistent_volume.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["job", "volumes", "persistent_volume"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/volumes/persistent_volume/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "job.volumes.persistent_volume for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes.persistent_volume

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [job](resources--workload--properties--job.md)
- [job.volumes](resources--workload--properties--job--volumes.md)
- job.volumes.persistent_volume

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
persistent_volume {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mount](resources--workload--properties--job--volumes--persistent_volume--mount.md): complete subsection reference.

- [storage](resources--workload--properties--job--volumes--persistent_volume--storage.md): complete subsection reference.

## Next pages

- [job.volumes.persistent_volume.mount](resources--workload--properties--job--volumes--persistent_volume--mount.md)
- [job.volumes.persistent_volume.storage](resources--workload--properties--job--volumes--persistent_volume--storage.md)
- [job.volumes](resources--workload--properties--job--volumes.md)
- [xcsh_workload](../resources/workload.md)
