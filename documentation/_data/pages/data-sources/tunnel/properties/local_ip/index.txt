---
page_title: "local_ip"
subcategory: ""
description: "Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS available are - 1. Local Interface - Network Interface from which IP address and network will be selected 2. IP Address - IP address and network can be configured explicitly."
xcsh_docs: {"aliases": ["local ip"], "body_bytes": 1218, "body_sha256": "sha256:c405ab0d17a38da5d8b1c7978d67050ac18bb15665b25d01c8954aed9f8a0486", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:intf", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip", "parent_id": "xcsh-docs:data-sources:tunnel:reference", "path": "documentation/data-sources/tunnel/properties/local_ip/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0121023110232123-0310011211312111-3032100012100203-0013123232002201-2113201321232021-1330120312221311-2203322102310323-1310230011303112", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip"], "schema_version": 1, "sections": [{"aliases": ["local ip intf"], "anchor": "section", "description": "Provides the local interface to pick up source IP and network for transporting encapsulated packet.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:intf", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "intf"], "syntax": "attribute", "type": "object"}, {"aliases": ["local ip ip address"], "anchor": "section", "description": "Provides the configuration to pick up source IP and network for transporting encapsulated packet.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "ip_address"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS available are - 1. Local Interface - Network Interface from which IP address and network will be selected 2. IP Address - IP address and network can be configured explicitly.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["tunnelCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- local_ip

<a id="section"></a>

Type: `"single"`. Computed.

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - 1. Local Interface - Network Interface from which IP address and network will be
selected 2. IP Address - IP address and network can be configured explicitly.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"intf\",\"ip_address\"]"
}
```

## Direct properties

- [intf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/intf/): complete subsection reference.

- [ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/): complete subsection reference.
