---
page_title: "aws.not_managed.node_list.interface_list.bond_interface.lacp"
subcategory: ""
description: "aws.not_managed.node_list.interface_list.bond_interface.lacp for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2445, "body_sha256": "sha256:995f80aac31ef07b21eaa05e2c62ba9dde22900256b14e34e5b848e03279f16e", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:bond_interface:lacp", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:bond_interface:lacp", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:bond_interface", "path": "docs/guides/data-sources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--bond_interface--lacp.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws", "not_managed", "node_list", "interface_list", "bond_interface", "lacp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/bond_interface/lacp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws.not_managed.node_list.interface_list.bond_interface.lacp for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.not_managed.node_list.interface_list.bond_interface.lacp

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [aws](data-sources--securemesh_site_v2--properties--aws.md)
- [aws.not_managed](data-sources--securemesh_site_v2--properties--aws--not_managed.md)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--properties--aws--not_managed--node_list.md)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list.md)
- [aws.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--bond_interface.md)
- aws.not_managed.node_list.interface_list.bond_interface.lacp

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

<a id="schema-aws--not_managed--node_list--interface_list--bond_interface--lacp--rate"></a>

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

- [aws.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--properties--aws--not_managed--node_list--interface_list--bond_interface.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
