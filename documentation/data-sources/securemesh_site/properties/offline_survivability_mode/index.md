---
page_title: "offline_survivability_mode"
subcategory: ""
description: "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is "
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "offline survivability mode", "tls certificates"], "body_bytes": 2826, "body_sha256": "sha256:9ea9cbe5f565768492e8255dda90d5a22b4a5402bbcaa584341d4256caaba5ca", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "xcsh-docs:data-sources:securemesh_site:properties:offline_survivability_mode:no_offline_survivability_mode"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:offline_survivability_mode", "parent_id": "xcsh-docs:data-sources:securemesh_site:reference", "path": "documentation/data-sources/securemesh_site/properties/offline_survivability_mode/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2020103213000102-1110301232330123-2322321023321200-2131023113310313-2233322222100111-2032312323132100-0321122012303013-0320311231012313", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["offline_survivability_mode"], "schema_version": 1, "sections": [{"aliases": ["enable offline survivability mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["offline_survivability_mode", "enable_offline_survivability_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["no offline survivability mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:offline_survivability_mode:no_offline_survivability_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["offline_survivability_mode", "no_offline_survivability_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is ", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# offline_survivability_mode

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- offline_survivability_mode

<a id="section"></a>

Type: `"single"`. Computed.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

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

- [enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/offline_survivability_mode/enable_offline_survivability_mode/): complete subsection reference.

- [no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/offline_survivability_mode/no_offline_survivability_mode/): complete subsection reference.

## Next pages

- [offline_survivability_mode.enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/offline_survivability_mode/enable_offline_survivability_mode/)
- [offline_survivability_mode.no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/offline_survivability_mode/no_offline_survivability_mode/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
