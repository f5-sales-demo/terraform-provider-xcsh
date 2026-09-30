---
page_title: "offline_survivability_mode"
subcategory: ""
description: "offline_survivability_mode for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2289, "body_sha256": "sha256:a9fad0ced7e03c56b5abc334c516ed57bf568ed03ff07708525c1ccbc6ca6f66", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:offline_survivability_mode", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "xcsh-docs:data-sources:aws_tgw_site:properties:offline_survivability_mode:no_offline_survivability_mode"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:offline_survivability_mode", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "path": "docs/guides/data-sources--aws_tgw_site--properties--offline_survivability_mode.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["offline_survivability_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "offline_survivability_mode for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# offline_survivability_mode

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
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

- [enable_offline_survivability_mode](data-sources--aws_tgw_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md): complete subsection reference.

- [no_offline_survivability_mode](data-sources--aws_tgw_site--properties--offline_survivability_mode--no_offline_survivability_mode.md): complete subsection reference.

## Next pages

- [offline_survivability_mode.enable_offline_survivability_mode](data-sources--aws_tgw_site--properties--offline_survivability_mode--enable_offline_survivability_mode.md)
- [offline_survivability_mode.no_offline_survivability_mode](data-sources--aws_tgw_site--properties--offline_survivability_mode--no_offline_survivability_mode.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
