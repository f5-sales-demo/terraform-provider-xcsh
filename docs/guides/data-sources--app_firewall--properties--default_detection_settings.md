---
page_title: "default_detection_settings"
subcategory: "Security"
description: "default_detection_settings for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1365, "body_sha256": "sha256:a859f4bbf8221ea00201df8ce5d3b2a913e7fb591398a779079a60ee69583378", "canonical_id": "xcsh-docs:data-sources:app_firewall:properties:default_detection_settings", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:default_detection_settings", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "docs/guides/data-sources--app_firewall--properties--default_detection_settings.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_detection_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/default_detection_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_detection_settings for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_detection_settings

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md)
- [Property reference](data-sources--app_firewall--reference.md)
- default_detection_settings

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_detection\_settings, detection\_settings; Default: default\_detection\_settings\]
Configuration parameter for default detection settings. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [default_detection_settings](data-sources--app_firewall--properties--default_detection_settings.md#section)
- [detection_settings](data-sources--app_firewall--properties--detection_settings.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--app_firewall--reference.md)
- [xcsh_app_firewall](../data-sources/app_firewall.md)
