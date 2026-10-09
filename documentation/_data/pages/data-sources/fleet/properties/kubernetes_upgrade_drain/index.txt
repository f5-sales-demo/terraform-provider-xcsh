---
page_title: "kubernetes_upgrade_drain"
subcategory: ""
description: "Specify how worker nodes within a site will be upgraded."
xcsh_docs: {"aliases": ["kubernetes upgrade drain"], "body_bytes": 1181, "body_sha256": "sha256:52784fcf8e919680553b248cf8afc418a75b91cee753ca864a2a2344f61a3ccc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain:enable_upgrade_drain"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/kubernetes_upgrade_drain/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2102203213212232-2121010023212310-1000010030032213-0112022300111120-1031323232020023-0321132223313110-0333211122013003-3212333331332202", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kubernetes_upgrade_drain"], "schema_version": 1, "sections": [{"aliases": ["kubernetes upgrade drain disable upgrade drain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "disable_upgrade_drain"], "syntax": "attribute", "type": "object"}, {"aliases": ["kubernetes upgrade drain enable upgrade drain"], "anchor": "section", "description": "Specify batch upgrade settings for worker nodes within a site.", "document_id": "xcsh-docs:data-sources:fleet:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/kubernetes_upgrade_drain/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Specify how worker nodes within a site will be upgraded.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["fleetCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- kubernetes_upgrade_drain

<a id="section"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

## Direct properties

- [disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/kubernetes_upgrade_drain/disable_upgrade_drain/): complete subsection reference.

- [enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/kubernetes_upgrade_drain/enable_upgrade_drain/): complete subsection reference.
