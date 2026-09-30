---
page_title: "stateful_service.persistent_volumes.persistent_volume"
subcategory: "Container"
description: "stateful_service.persistent_volumes.persistent_volume for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1648, "body_sha256": "sha256:b6573fd8a38c8685c00efdd31196fe11f455eb51754b41c7e3ca438a565ab17f", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:mount", "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes", "path": "docs/guides/resources--workload--properties--stateful_service--persistent_volumes--persistent_volume.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.persistent_volumes.persistent_volume for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# stateful_service.persistent_volumes.persistent_volume

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.persistent_volumes](resources--workload--properties--stateful_service--persistent_volumes.md)
- stateful_service.persistent_volumes.persistent_volume

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

- [mount](resources--workload--properties--stateful_service--persistent_volumes--persistent_volume--mount.md): complete subsection reference.

- [storage](resources--workload--properties--stateful_service--persistent_volumes--persistent_volume--storage.md): complete subsection reference.

## Next pages

- [stateful_service.persistent_volumes.persistent_volume.mount](resources--workload--properties--stateful_service--persistent_volumes--persistent_volume--mount.md)
- [stateful_service.persistent_volumes.persistent_volume.storage](resources--workload--properties--stateful_service--persistent_volumes--persistent_volume--storage.md)
- [stateful_service.persistent_volumes](resources--workload--properties--stateful_service--persistent_volumes.md)
- [xcsh_workload](../resources/workload.md)
