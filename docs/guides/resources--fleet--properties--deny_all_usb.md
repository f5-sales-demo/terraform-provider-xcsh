---
page_title: "deny_all_usb"
subcategory: ""
description: "deny_all_usb for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 753, "body_sha256": "sha256:30dc157f903f928dcdb1f84219b46d32e4f70528ed6381f56cb81658eb002da9", "canonical_id": "xcsh-docs:resources:fleet:properties:deny_all_usb", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:deny_all_usb", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "docs/guides/resources--fleet--properties--deny_all_usb.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["deny_all_usb"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/deny_all_usb/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "deny_all_usb for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# deny_all_usb

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- deny_all_usb

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for deny all usb.

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
deny_all_usb = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--fleet--reference.md)
- [xcsh_fleet](../resources/fleet.md)
