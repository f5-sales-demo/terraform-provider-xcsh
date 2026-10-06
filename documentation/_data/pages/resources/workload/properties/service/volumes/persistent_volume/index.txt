---
page_title: "service.volumes.persistent_volume"
subcategory: "Container"
description: "Volume containing the Persistent Storage for the workload."
xcsh_docs: {"aliases": ["service volumes persistent volume"], "body_bytes": 1404, "body_sha256": "sha256:54055c4e8837a53998222bf8284db317f71b76a0f69f307e5031789879b86055", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:mount", "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume", "parent_id": "xcsh-docs:resources:workload:properties:service:volumes", "path": "documentation/resources/workload/properties/service/volumes/persistent_volume/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3111313320323303-3333313230133022-1132322302230100-3232032012300012-1300203002120002-1333332300333233-1332212331012323-1223110121113331", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "volumes", "persistent_volume"], "schema_version": 1, "sections": [{"aliases": ["service volumes persistent volume mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:mount", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--volumes--persistent_volume--mount--mount_path", "enforcement": "provider-schema", "group": "service.volumes.persistent_volume.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:mount", "type": "requires"}], "schema_path": ["service", "volumes", "persistent_volume", "mount"], "syntax": "block", "type": "object"}, {"aliases": ["service volumes persistent volume storage"], "anchor": "section", "description": "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--volumes--persistent_volume--storage--class_name", "enforcement": "provider-schema", "group": "service.volumes.persistent_volume.storage:ConflictingObjectAttributes:class_name,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.volumes.persistent_volume.storage:ConflictingObjectAttributes:class_name,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage:default", "type": "conflicts"}, {"anchor": "schema-service--volumes--persistent_volume--storage--storage_size", "enforcement": "provider-schema", "group": "service.volumes.persistent_volume.storage:RequiredObjectAttributes:storage_size", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:volumes:persistent_volume:storage", "type": "requires"}], "schema_path": ["service", "volumes", "persistent_volume", "storage"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/volumes/persistent_volume/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Volume containing the Persistent Storage for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes.persistent_volume

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/)
- service.volumes.persistent_volume

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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/persistent_volume/mount/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/persistent_volume/storage/): complete subsection reference.
