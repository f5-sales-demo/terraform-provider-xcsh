---
page_title: "nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map"
subcategory: ""
description: "Specify static IPv4 addresses per node."
xcsh_docs: {"aliases": ["nutanix not managed node list interface list dhcp server interface ip map"], "body_bytes": 3900, "body_sha256": "sha256:4c9e6c4ea6763785621feb5195b4258ada0b83170e58a617d1864ffe8a83cd07", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:dhcp_server", "path": "documentation/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/dhcp_server/interface_ip_map/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1300230220001111-2120230133130003-0311321100332200-0220200211203310-2321030013223320-1123232033133010-0133312221201200-1233100210101201", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "dhcp_server", "interface_ip_map"], "schema_version": 1, "sections": [{"aliases": ["nutanix not managed node list interface list dhcp server interface ip map interface ip map"], "anchor": "schema-nutanix--not_managed--node_list--interface_list--dhcp_server--interface_ip_map--interface_ip_map", "description": "Specify static IPv4 addresses per site:node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "dhcp_server", "interface_ip_map", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/dhcp_server/interface_ip_map/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Specify static IPv4 addresses per node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [nutanix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/)
- [nutanix.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/)
- [nutanix.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/)
- [nutanix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/dhcp_server/)
- nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

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
interface_ip_map {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-nutanix--not_managed--node_list--interface_list--dhcp_server--interface_ip_map--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```
