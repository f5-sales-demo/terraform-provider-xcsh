---
page_title: "deny_all_usb"
subcategory: ""
description: "deny_all_usb for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 852, "body_sha256": "sha256:da1dcac60597ad4ad8bc1bb12501129e9b04bc041fa013a6e639874545b728fa", "canonical_id": "xcsh-docs:resources:fleet:properties:deny_all_usb", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:deny_all_usb", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "docs/guides/resources--fleet--properties--deny_all_usb.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["deny_all_usb"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/deny_all_usb/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "deny_all_usb for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
