---
page_title: "layer2_interface.l2vlan_interface"
subcategory: ""
description: "layer2_interface.l2vlan_interface for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 3153, "body_sha256": "sha256:023efe7ea1f02a712ed70d76663d5a79c7068a43a53f49429443baaf4476f613", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_interface", "parent_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "path": "documentation/data-sources/network_interface/properties/layer2_interface/l2vlan_interface/index.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["layer2_interface", "l2vlan_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/layer2_interface/l2vlan_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "layer2_interface.l2vlan_interface for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# layer2_interface.l2vlan_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/)
- layer2_interface.l2vlan_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for l2vlan interface.

Upstream description:

Layer2 VLAN Interface Configuration.

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

<a id="schema-layer2_interface--l2vlan_interface--device"></a>

### device property

Type: `"string"`. Computed.

Ethernet Device. Physical ethernet interface.

Upstream description:

Physical ethernet interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-layer2_interface--l2vlan_interface--vlan_id"></a>

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

- [layer2_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
