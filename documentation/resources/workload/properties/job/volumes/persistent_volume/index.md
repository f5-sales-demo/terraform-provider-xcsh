---
page_title: "job.volumes.persistent_volume"
subcategory: "Container"
description: "Volume containing the Persistent Storage for the workload."
xcsh_docs: {"aliases": ["job volumes persistent volume"], "body_bytes": 1931, "body_sha256": "sha256:ab3fa75ccdd636a6b3b82885dd4056f6bd5a030480f24557de21d6a7c19e0b43", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:mount", "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:storage"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume", "parent_id": "xcsh-docs:resources:workload:properties:job:volumes", "path": "documentation/resources/workload/properties/job/volumes/persistent_volume/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311", "registry_path": "docs/guides/resources--workload--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "volumes", "persistent_volume"], "schema_version": 1, "sections": [{"aliases": ["job volumes persistent volume mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:mount", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-job--volumes--persistent_volume--mount--mount_path", "enforcement": "provider-schema", "group": "job.volumes.persistent_volume.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:mount", "type": "requires"}], "schema_path": ["job", "volumes", "persistent_volume", "mount"], "syntax": "block", "type": "object"}, {"aliases": ["job volumes persistent volume storage"], "anchor": "section", "description": "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)", "document_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:storage", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-job--volumes--persistent_volume--storage--class_name", "enforcement": "provider-schema", "group": "job.volumes.persistent_volume.storage:ConflictingObjectAttributes:class_name,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes.persistent_volume.storage:ConflictingObjectAttributes:class_name,default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:storage:default", "type": "conflicts"}, {"anchor": "schema-job--volumes--persistent_volume--storage--storage_size", "enforcement": "provider-schema", "group": "job.volumes.persistent_volume.storage:RequiredObjectAttributes:storage_size", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume:storage", "type": "requires"}], "schema_path": ["job", "volumes", "persistent_volume", "storage"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/volumes/persistent_volume/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Volume containing the Persistent Storage for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes.persistent_volume

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/)
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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/persistent_volume/mount/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/persistent_volume/storage/): complete subsection reference.

## Next pages

- [job.volumes.persistent_volume.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/persistent_volume/mount/)
- [job.volumes.persistent_volume.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/persistent_volume/storage/)
- [job.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
