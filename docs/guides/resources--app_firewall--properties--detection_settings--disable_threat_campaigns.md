---
page_title: "detection_settings.disable_threat_campaigns"
subcategory: "Security"
description: "detection_settings.disable_threat_campaigns for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1049, "body_sha256": "sha256:f0b9bc07e1babc5c4cbf3e5ae388ab0a7f8c2120c652b878064024f661d9c32b", "canonical_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:disable_threat_campaigns", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings:disable_threat_campaigns", "parent_id": "xcsh-docs:resources:app_firewall:properties:detection_settings", "path": "docs/guides/resources--app_firewall--properties--detection_settings--disable_threat_campaigns.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["detection_settings", "disable_threat_campaigns"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/disable_threat_campaigns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "detection_settings.disable_threat_campaigns for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.disable_threat_campaigns

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- [detection_settings](resources--app_firewall--properties--detection_settings.md)
- detection_settings.disable_threat_campaigns

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
disable_threat_campaigns = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [detection_settings](resources--app_firewall--properties--detection_settings.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
