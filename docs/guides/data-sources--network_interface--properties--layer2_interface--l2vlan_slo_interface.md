---
page_title: "layer2_interface.l2vlan_slo_interface"
subcategory: ""
description: "layer2_interface.l2vlan_slo_interface for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1886, "body_sha256": "sha256:ea58f6d511b676cc54c3253d0e9b7c0828b90a3de093b1b421c6100117fb4bd8", "canonical_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "parent_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "path": "docs/guides/data-sources--network_interface--properties--layer2_interface--l2vlan_slo_interface.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["layer2_interface", "l2vlan_slo_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/layer2_interface/l2vlan_slo_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "layer2_interface.l2vlan_slo_interface for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# layer2_interface.l2vlan_slo_interface

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- [layer2_interface](data-sources--network_interface--properties--layer2_interface.md)
- layer2_interface.l2vlan_slo_interface

<a id="section"></a>

Type: `"single"`. Computed.

Layer2 Site Local Outside VLAN Interface Configuration.

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

<a id="schema-layer2_interface--l2vlan_slo_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Computed.

VLAN ID. VLAN ID

Upstream description:

VLAN ID

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

## Next pages

- [layer2_interface](data-sources--network_interface--properties--layer2_interface.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
