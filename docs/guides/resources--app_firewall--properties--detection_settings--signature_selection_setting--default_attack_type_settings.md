---
page_title: "detection_settings.signature_selection_setting.default_attack_type_settings"
subcategory: "Security"
description: "detection_settings.signature_selection_setting.default_attack_type_settings for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1353, "body_sha256": "sha256:c2095ee1d19b0d06c8f768073b9063a1266da390998155fb5de07a1702d73159", "canonical_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_attack_type_settings", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_attack_type_settings", "parent_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting", "path": "docs/guides/resources--app_firewall--properties--detection_settings--signature_selection_setting--default_attack_type_settings.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting", "default_attack_type_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/signature_selection_setting/default_attack_type_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "detection_settings.signature_selection_setting.default_attack_type_settings for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting.default_attack_type_settings

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- [detection_settings](resources--app_firewall--properties--detection_settings.md)
- [detection_settings.signature_selection_setting](resources--app_firewall--properties--detection_settings--signature_selection_setting.md)
- detection_settings.signature_selection_setting.default_attack_type_settings

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default attack type settings.

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

Terraform syntax:

```terraform
default_attack_type_settings = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [detection_settings.signature_selection_setting](resources--app_firewall--properties--detection_settings--signature_selection_setting.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
