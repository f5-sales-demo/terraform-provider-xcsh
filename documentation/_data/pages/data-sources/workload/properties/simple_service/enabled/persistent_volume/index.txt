---
page_title: "simple_service.enabled.persistent_volume"
subcategory: "Container"
description: "Volume containing the Persistent Storage for the workload."
xcsh_docs: {"aliases": ["simple service enabled persistent volume"], "body_bytes": 1989, "body_sha256": "sha256:5b1329f61ea58a378b03ced2a68015b35be777781dc23b1556781daa71df6a03", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:mount", "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled", "path": "documentation/data-sources/workload/properties/simple_service/enabled/persistent_volume/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331", "registry_path": "docs/guides/data-sources--workload--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "enabled", "persistent_volume"], "schema_version": 1, "sections": [{"aliases": ["mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:mount", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "enabled", "persistent_volume", "mount"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage"], "anchor": "section", "description": "Persistent storage configuration is used to configure Persistent Volume Claim (PVC)", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:enabled:persistent_volume:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "enabled", "persistent_volume", "storage"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/enabled/persistent_volume/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Volume containing the Persistent Storage for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.enabled.persistent_volume

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/)
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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/mount/): complete subsection reference.

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/storage/): complete subsection reference.

## Next pages

- [simple_service.enabled.persistent_volume.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/mount/)
- [simple_service.enabled.persistent_volume.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/persistent_volume/storage/)
- [simple_service.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/enabled/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
