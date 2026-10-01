---
page_title: "simple_service.enabled.persistent_volume"
subcategory: "Container"
description: "simple_service.enabled.persistent_volume for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1491, "body_sha256": "sha256:e6e3c8706ac53f67db9183455748fbe325113d3be95e2647e9ff03e4520eed1a", "canonical_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:mount", "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled", "path": "docs/guides/data-sources--workload--properties--simple_service--enabled--persistent_volume.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["simple_service", "enabled", "persistent_volume"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/enabled/persistent_volume/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "simple_service.enabled.persistent_volume for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.enabled.persistent_volume

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [simple_service](data-sources--workload--properties--simple_service.md)
- [simple_service.enabled](data-sources--workload--properties--simple_service--enabled.md)
- simple_service.enabled.persistent_volume

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

- [mount](data-sources--workload--properties--simple_service--enabled--persistent_volume--mount.md): complete subsection reference.

- [storage](data-sources--workload--properties--simple_service--enabled--persistent_volume--storage.md): complete subsection reference.

## Next pages

- [simple_service.enabled.persistent_volume.mount](data-sources--workload--properties--simple_service--enabled--persistent_volume--mount.md)
- [simple_service.enabled.persistent_volume.storage](data-sources--workload--properties--simple_service--enabled--persistent_volume--storage.md)
- [simple_service.enabled](data-sources--workload--properties--simple_service--enabled.md)
- [xcsh_workload](../data-sources/workload.md)
