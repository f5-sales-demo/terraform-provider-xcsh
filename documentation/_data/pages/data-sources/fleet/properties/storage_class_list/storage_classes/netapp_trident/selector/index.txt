---
page_title: "storage_class_list.storage_classes.netapp_trident.selector"
subcategory: ""
description: "Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a volume. The volume will have the aspects defined in the chosen virtual pool."
xcsh_docs: {"aliases": ["storage class list storage classes netapp trident selector"], "body_bytes": 1397, "body_sha256": "sha256:0e4fe6cdc4b72caee48d2118bdb2842507fc4a36af19d1eebbb1202dbf4431cf", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident:selector", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "path": "documentation/data-sources/fleet/properties/storage_class_list/storage_classes/netapp_trident/selector/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2220110333231123-0331223110210132-1332010323112022-2230113220023112-0222100120211312-0303011321223312-1232321003030230-2010103200321031", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "netapp_trident", "selector"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_class_list/storage_classes/netapp_trident/selector/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a volume. The volume will have the aspects defined in the chosen virtual pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["fleetCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.netapp_trident.selector

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/)
- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/)
- [storage_class_list.storage_classes.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/netapp_trident/)
- storage_class_list.storage_classes.netapp_trident.selector

<a id="section"></a>

Type: `"single"`. Computed.

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

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

This is an empty object or choice marker. It has no direct properties.
