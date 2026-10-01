---
page_title: "baremetal.not_managed.node_list.interface_list.bond_interface.lacp"
subcategory: ""
description: "baremetal.not_managed.node_list.interface_list.bond_interface.lacp for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2897, "body_sha256": "sha256:03e3bb68cfe53e85adcfcff5270767a1dd75b705d799b0a33c3f99f890081cc5", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:bond_interface:lacp", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:bond_interface:lacp", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:bond_interface", "path": "docs/guides/resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--bond_interface--lacp.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "bond_interface", "lacp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/bond_interface/lacp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "baremetal.not_managed.node_list.interface_list.bond_interface.lacp for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# baremetal.not_managed.node_list.interface_list.bond_interface.lacp

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [baremetal](resources--securemesh_site_v2--properties--baremetal.md)
- [baremetal.not_managed](resources--securemesh_site_v2--properties--baremetal--not_managed.md)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list.md)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list.md)
- [baremetal.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--bond_interface.md)
- baremetal.not_managed.node_list.interface_list.bond_interface.lacp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
```

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
lacp {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-baremetal--not_managed--node_list--interface_list--bond_interface--lacp--rate"></a>

### rate property

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

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

- [baremetal.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--bond_interface.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
