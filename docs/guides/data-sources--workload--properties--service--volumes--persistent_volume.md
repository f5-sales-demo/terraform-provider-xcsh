---
page_title: "service.volumes.persistent_volume"
subcategory: "Container"
description: "service.volumes.persistent_volume for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1393, "body_sha256": "sha256:5a571673e08405600082c846f9d22337524d22c648e34c393c65578b6fa0358c", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:volumes:persistent_volume", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:volumes:persistent_volume:mount", "xcsh-docs:data-sources:workload:properties:service:volumes:persistent_volume:storage"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:volumes:persistent_volume", "parent_id": "xcsh-docs:data-sources:workload:properties:service:volumes", "path": "docs/guides/data-sources--workload--properties--service--volumes--persistent_volume.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "volumes", "persistent_volume"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/volumes/persistent_volume/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.volumes.persistent_volume for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes.persistent_volume

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.volumes](data-sources--workload--properties--service--volumes.md)
- service.volumes.persistent_volume

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

- [mount](data-sources--workload--properties--service--volumes--persistent_volume--mount.md): complete subsection reference.

- [storage](data-sources--workload--properties--service--volumes--persistent_volume--storage.md): complete subsection reference.

## Next pages

- [service.volumes.persistent_volume.mount](data-sources--workload--properties--service--volumes--persistent_volume--mount.md)
- [service.volumes.persistent_volume.storage](data-sources--workload--properties--service--volumes--persistent_volume--storage.md)
- [service.volumes](data-sources--workload--properties--service--volumes.md)
- [xcsh_workload](../data-sources/workload.md)
