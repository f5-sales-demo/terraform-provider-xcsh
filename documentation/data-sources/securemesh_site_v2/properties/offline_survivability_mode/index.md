---
page_title: "offline_survivability_mode"
subcategory: ""
description: "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is"
xcsh_docs: {"aliases": ["offline survivability mode"], "body_bytes": 1834, "body_sha256": "sha256:a1e4ea6ea3d7c2a301858449be1669985be7858e278cbb6cd5dc0d3c9daffd40", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:offline_survivability_mode:enable_offline_survivability_mode", "xcsh-docs:data-sources:securemesh_site_v2:properties:offline_survivability_mode:no_offline_survivability_mode"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:offline_survivability_mode", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/offline_survivability_mode/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1021200122303311-3233121032110203-1312312313203131-3300220212031202-1330203323130121-3303330222102333-1033013213023320-1320312232330232", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["offline_survivability_mode"], "schema_version": 1, "sections": [{"aliases": ["offline survivability mode enable offline survivability mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:offline_survivability_mode:enable_offline_survivability_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["offline_survivability_mode", "enable_offline_survivability_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["offline survivability mode no offline survivability mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:offline_survivability_mode:no_offline_survivability_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["offline_survivability_mode", "no_offline_survivability_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# offline_survivability_mode

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- offline_survivability_mode

<a id="section"></a>

Type: `"single"`. Computed.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

## Direct properties

- [enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/offline_survivability_mode/enable_offline_survivability_mode/): complete subsection reference.

- [no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/offline_survivability_mode/no_offline_survivability_mode/): complete subsection reference.
