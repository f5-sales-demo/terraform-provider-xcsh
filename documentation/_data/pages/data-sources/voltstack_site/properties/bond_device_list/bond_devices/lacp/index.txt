---
page_title: "bond_device_list.bond_devices.lacp"
subcategory: ""
description: "LACP parameters for the bond device."
xcsh_docs: {"aliases": ["bond device list bond devices lacp"], "body_bytes": 2220, "body_sha256": "sha256:29947caa5d27c5de528a7a1f870a4e482bc0eb5a352c7f63602760b8dbde2194", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:bond_device_list:bond_devices:lacp", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:bond_device_list:bond_devices", "path": "documentation/data-sources/voltstack_site/properties/bond_device_list/bond_devices/lacp/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2302233310312113-3030002231200101-1123222213221021-0303112333101231-0020232032320011-1223130103022020-2101031132023313-2300332130112233", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bond_device_list", "bond_devices", "lacp"], "schema_version": 1, "sections": [{"aliases": ["bond device list bond devices lacp rate"], "anchor": "schema-bond_device_list--bond_devices--lacp--rate", "description": "Interval in seconds to transmit LACP packets.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:bond_device_list:bond_devices:lacp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "lacp", "rate"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/bond_device_list/bond_devices/lacp/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "LACP parameters for the bond device.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list.bond_devices.lacp

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/bond_device_list/)
- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/bond_device_list/bond_devices/)
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

- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/bond_device_list/bond_devices/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
