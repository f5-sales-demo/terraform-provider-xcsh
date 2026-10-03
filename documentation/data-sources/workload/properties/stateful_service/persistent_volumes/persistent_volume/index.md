---
page_title: "stateful_service.persistent_volumes.persistent_volume"
subcategory: "Container"
description: "Volume containing the Persistent Storage for the workload."
xcsh_docs: {"aliases": ["stateful service persistent volumes persistent volume"], "body_bytes": 2149, "body_sha256": "sha256:e08a39137a8111b928bfc39d879c485e5ddbae9f2f4a3dac767e2c4eb37c9b40", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:persistent_volumes:persistent_volume:mount", "xcsh-docs:data-sources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:persistent_volumes:persistent_volume", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:persistent_volumes", "path": "documentation/data-sources/workload/properties/stateful_service/persistent_volumes/persistent_volume/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0313133013332231-3300200130030103-2002302022200030-1003021003231002-1121111231212110-2232311322111203-3202032103100110-0130333321220032", "registry_path": "docs/guides/data-sources--workload--reference--group-028.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume"], "schema_version": 1, "sections": [{"aliases": ["stateful service persistent volumes persistent volume mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:persistent_volumes:persistent_volume:mount", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume", "mount"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service persistent volumes persistent volume storage"], "anchor": "section", "description": "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:persistent_volumes:persistent_volume:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "persistent_volumes", "persistent_volume", "storage"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/persistent_volumes/persistent_volume/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Volume containing the Persistent Storage for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.persistent_volumes.persistent_volume

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.persistent_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/persistent_volumes/)
- stateful_service.persistent_volumes.persistent_volume

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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/persistent_volumes/persistent_volume/mount/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/persistent_volumes/persistent_volume/storage/): complete subsection reference.

## Next pages

- [stateful_service.persistent_volumes.persistent_volume.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/persistent_volumes/persistent_volume/mount/)
- [stateful_service.persistent_volumes.persistent_volume.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/persistent_volumes/persistent_volume/storage/)
- [stateful_service.persistent_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/persistent_volumes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
