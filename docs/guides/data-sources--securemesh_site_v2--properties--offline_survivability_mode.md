---
page_title: "offline_survivability_mode"
subcategory: ""
description: "offline_survivability_mode for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2349, "body_sha256": "sha256:3dfdbff204a7e6387a0255e234a89117179f56990c0865d7b60723bfa7069988", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:offline_survivability_mode", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:offline_survivability_mode:enable_offline_survivability_mode", "xcsh-docs:data-sources:securemesh_site_v2:properties:offline_survivability_mode:no_offline_survivability_mode"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:offline_survivability_mode", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "docs/guides/data-sources--securemesh_site_v2--properties--offline_survivability_mode.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["offline_survivability_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "offline_survivability_mode for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# offline_survivability_mode

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
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

- [enable_offline_survivability_mode](data-sources--securemesh_site_v2--properties--offline_survivability_mode--enable_offline_survivability_mode.md): complete subsection reference.

- [no_offline_survivability_mode](data-sources--securemesh_site_v2--properties--offline_survivability_mode--no_offline_survivability_mode.md): complete subsection reference.

## Next pages

- [offline_survivability_mode.enable_offline_survivability_mode](data-sources--securemesh_site_v2--properties--offline_survivability_mode--enable_offline_survivability_mode.md)
- [offline_survivability_mode.no_offline_survivability_mode](data-sources--securemesh_site_v2--properties--offline_survivability_mode--no_offline_survivability_mode.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
