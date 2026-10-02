---
page_title: "offline_survivability_mode"
subcategory: ""
description: "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is"
xcsh_docs: {"aliases": ["offline survivability mode"], "body_bytes": 2816, "body_sha256": "sha256:828937ceb3bd7c36eb41939f4b7394c5e3d4b0681cf12f33de67c0f0455abc9b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "xcsh-docs:data-sources:voltstack_site:properties:offline_survivability_mode:no_offline_survivability_mode"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:offline_survivability_mode", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "documentation/data-sources/voltstack_site/properties/offline_survivability_mode/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3012201203013310-2202301221032113-3211212303232023-3001230312211120-0223122330211122-3321300212212230-1123030231222122-0310112232022022", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["offline_survivability_mode"], "schema_version": 1, "sections": [{"aliases": ["enable offline survivability mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["offline_survivability_mode", "enable_offline_survivability_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["no offline survivability mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:offline_survivability_mode:no_offline_survivability_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["offline_survivability_mode", "no_offline_survivability_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# offline_survivability_mode

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
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

- [enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/offline_survivability_mode/enable_offline_survivability_mode/): complete subsection reference.

- [no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/offline_survivability_mode/no_offline_survivability_mode/): complete subsection reference.

## Next pages

- [offline_survivability_mode.enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/offline_survivability_mode/enable_offline_survivability_mode/)
- [offline_survivability_mode.no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/offline_survivability_mode/no_offline_survivability_mode/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
