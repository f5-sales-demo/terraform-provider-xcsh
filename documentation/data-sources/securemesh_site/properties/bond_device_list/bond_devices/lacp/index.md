---
page_title: "bond_device_list.bond_devices.lacp"
subcategory: ""
description: "bond_device_list.bond_devices.lacp for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 2228, "body_sha256": "sha256:de2c75687797c48ac023862b1ece59d792a5f1387b5ff85346ba64ec2d13fd20", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:bond_device_list:bond_devices:lacp", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:bond_device_list:bond_devices", "path": "documentation/data-sources/securemesh_site/properties/bond_device_list/bond_devices/lacp/index.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["bond_device_list", "bond_devices", "lacp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/bond_device_list/bond_devices/lacp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bond_device_list.bond_devices.lacp for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list.bond_devices.lacp

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/bond_device_list/)
- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/bond_device_list/bond_devices/)
- bond_device_list.bond_devices.lacp

<a id="section"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

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

## Direct properties

<a id="schema-bond_device_list--bond_devices--lacp--rate"></a>

### rate property

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

## Next pages

- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/bond_device_list/bond_devices/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
