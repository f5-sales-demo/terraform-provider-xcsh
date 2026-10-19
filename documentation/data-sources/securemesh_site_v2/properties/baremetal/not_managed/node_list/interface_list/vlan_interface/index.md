---
page_title: "baremetal.not_managed.node_list.interface_list.vlan_interface"
subcategory: ""
description: "Configuration parameter for vlan interface."
xcsh_docs: {"aliases": ["baremetal not managed node list interface list vlan interface"], "body_bytes": 3380, "body_sha256": "sha256:ee629c8755851fbd70f48bda12106c666c5f6a7bdc41cd32f89b29f85de8c70c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:vlan_interface", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list", "path": "documentation/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/vlan_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0100332010022300-0003110001013331-1302332000010220-0320133313322201-2313103122022232-1233102002201012-2121231022011032-2032111233031310", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "vlan_interface"], "schema_version": 1, "sections": [{"aliases": ["baremetal not managed node list interface list vlan interface device"], "anchor": "schema-baremetal--not_managed--node_list--interface_list--vlan_interface--device", "description": "Select a parent interface from the dropdown.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:vlan_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "vlan_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["baremetal not managed node list interface list vlan interface vlan id"], "anchor": "schema-baremetal--not_managed--node_list--interface_list--vlan_interface--vlan_id", "description": "Configure the VLAN tag for this interface.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:vlan_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "vlan_interface", "vlan_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/vlan_interface/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Configuration parameter for vlan interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# baremetal.not_managed.node_list.interface_list.vlan_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [baremetal](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/)
- [baremetal.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/)
- [baremetal.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/)
- [baremetal.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/)
- baremetal.not_managed.node_list.interface_list.vlan_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for vlan interface.

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

<a id="schema-baremetal--not_managed--node_list--interface_list--vlan_interface--device"></a>

### device property

Type: `"string"`. Computed.

Select a parent interface from the dropdown.

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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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

<a id="schema-baremetal--not_managed--node_list--interface_list--vlan_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Computed.

Configure the VLAN tag for this interface.

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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```
