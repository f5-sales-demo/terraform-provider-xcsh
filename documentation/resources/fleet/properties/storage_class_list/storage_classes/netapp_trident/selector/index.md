---
page_title: "storage_class_list.storage_classes.netapp_trident.selector"
subcategory: ""
description: "Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a volume. The volume will have the aspects defined in the chosen virtual pool."
xcsh_docs: {"aliases": ["storage class list storage classes netapp trident selector"], "body_bytes": 1452, "body_sha256": "sha256:d695902a407db594c76c34dd792850d281846a29787adf5baab64a2cc97f32aa", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident:selector", "parent_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "path": "documentation/resources/fleet/properties/storage_class_list/storage_classes/netapp_trident/selector/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1231211312010313-3221012302202022-1123112203201012-0332131023302331-2312131323222331-3131200232222010-0210211001112201-2201230101112003", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "netapp_trident", "selector"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_class_list/storage_classes/netapp_trident/selector/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a volume. The volume will have the aspects defined in the chosen virtual pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["fleetCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.netapp_trident.selector

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/)
- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/)
- [storage_class_list.storage_classes.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/netapp_trident/)
- storage_class_list.storage_classes.netapp_trident.selector

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
selector {}
```

This is an empty object or choice marker. It has no direct properties.
