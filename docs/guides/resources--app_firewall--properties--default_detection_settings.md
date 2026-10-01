---
page_title: "default_detection_settings"
subcategory: "Security"
description: "default_detection_settings for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1426, "body_sha256": "sha256:332ad654d0c979574299afaeab465ce9d6c7149e5a805fecbbdf8707655443b1", "canonical_id": "xcsh-docs:resources:app_firewall:properties:default_detection_settings", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:default_detection_settings", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "docs/guides/resources--app_firewall--properties--default_detection_settings.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_detection_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/default_detection_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_detection_settings for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_detection_settings

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- default_detection_settings

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [default_detection_settings](resources--app_firewall--properties--default_detection_settings.md#section)
- [detection_settings](resources--app_firewall--properties--detection_settings.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_detection_settings = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--app_firewall--reference.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
