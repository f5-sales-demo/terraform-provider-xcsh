---
page_title: "custom_storage_config.storage_device_list"
subcategory: ""
description: "Add additional custom storage classes in Kubernetes for this fleet."
xcsh_docs: {"aliases": ["custom storage config storage device list"], "body_bytes": 1590, "body_sha256": "sha256:e245176fa0f44f1f38ecc778dd22c5d6962e4e30b4ec0c22674e841ea90d3bed", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/storage_device_list/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list"], "schema_version": 1, "sections": [{"aliases": ["storage devices"], "anchor": "section", "description": "List of custom storage devices.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_device_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Add additional custom storage classes in Kubernetes for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_device_list

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- custom_storage_config.storage_device_list

<a id="section"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this fleet.

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

- [storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
