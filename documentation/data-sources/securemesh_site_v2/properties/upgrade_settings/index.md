---
page_title: "upgrade_settings"
subcategory: ""
description: "Specify how a site will be upgraded."
xcsh_docs: {"aliases": ["upgrade settings"], "body_bytes": 962, "body_sha256": "sha256:e0d69566f23d90259c737c8ae5b4eca6194592a121db312f578c83cb604c6ea7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/upgrade_settings/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["upgrade_settings"], "schema_version": 1, "sections": [{"aliases": ["upgrade settings kubernetes upgrade drain"], "anchor": "section", "description": "Specify how worker nodes within a site will be upgraded.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/upgrade_settings/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Specify how a site will be upgraded.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upgrade_settings

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- upgrade_settings

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for upgrade settings.

Additional upstream details:

Specify how a site will be upgraded.

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

- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/): complete subsection reference.
