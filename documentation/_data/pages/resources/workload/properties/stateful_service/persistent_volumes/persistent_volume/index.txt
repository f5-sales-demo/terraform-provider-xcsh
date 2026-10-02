---
page_title: "stateful_service.persistent_volumes.persistent_volume"
subcategory: "Container"
description: "Volume containing the Persistent Storage for the workload."
xcsh_docs: {"aliases": ["stateful service persistent volumes persistent volume"], "body_bytes": 2245, "body_sha256": "sha256:85c7fea82cc03db0002a7c8ac01fff03677b46226fca99e2bc32a9c1000a5362", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:mount", "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes", "path": "documentation/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2112312000232122-2312223222320010-2033200313031001-1131110333101220-3321201023230022-2131001030110101-1110201233102113-0013122123223132", "registry_path": "docs/guides/resources--workload--reference--group-029.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume"], "schema_version": 1, "sections": [{"aliases": ["mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:mount", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--persistent_volumes--persistent_volume--mount--mount_path", "enforcement": "provider-schema", "group": "stateful_service.persistent_volumes.persistent_volume.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:mount", "type": "requires"}], "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume", "mount"], "syntax": "block", "type": "object"}, {"aliases": ["storage"], "anchor": "section", "description": "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--persistent_volumes--persistent_volume--storage--class_name", "enforcement": "provider-schema", "group": "stateful_service.persistent_volumes.persistent_volume.storage:ConflictingObjectAttributes:class_name,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.persistent_volumes.persistent_volume.storage:ConflictingObjectAttributes:class_name,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage:default", "type": "conflicts"}, {"anchor": "schema-stateful_service--persistent_volumes--persistent_volume--storage--storage_size", "enforcement": "provider-schema", "group": "stateful_service.persistent_volumes.persistent_volume.storage:RequiredObjectAttributes:storage_size", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "type": "requires"}], "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume", "storage"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Volume containing the Persistent Storage for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.persistent_volumes.persistent_volume

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.persistent_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/)
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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/mount/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/storage/): complete subsection reference.

## Next pages

- [stateful_service.persistent_volumes.persistent_volume.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/mount/)
- [stateful_service.persistent_volumes.persistent_volume.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/persistent_volume/storage/)
- [stateful_service.persistent_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/persistent_volumes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
