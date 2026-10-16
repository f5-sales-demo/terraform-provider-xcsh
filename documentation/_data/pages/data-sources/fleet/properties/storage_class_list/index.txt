---
page_title: "storage_class_list"
subcategory: ""
description: "Add additional custom storage classes in Kubernetes for this fleet."
xcsh_docs: {"aliases": ["storage class list"], "body_bytes": 852, "body_sha256": "sha256:5ece0f25b8accd98041b3be004e17816a0f9e69a155aac444969816c5a8a6025", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_class_list", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/storage_class_list/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2023110232112222-3321330312213101-2201202331303111-2221322320001110-1303303110032013-3100010023303012-0321310033103202-3032223021313121", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list"], "schema_version": 1, "sections": [{"aliases": ["storage class list storage classes"], "anchor": "section", "description": "List of custom storage classes.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["storage_class_list", "storage_classes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_class_list/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Add additional custom storage classes in Kubernetes for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["fleetCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- storage_class_list

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

- [storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/): complete subsection reference.
