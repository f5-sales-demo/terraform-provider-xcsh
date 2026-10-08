---
page_title: "local_ip"
subcategory: ""
description: "Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS available are - 1. Local Interface - Network Interface from which IP address and network will be selected 2. IP Address - IP address and network can be configured explicitly."
xcsh_docs: {"aliases": ["local ip"], "body_bytes": 1517, "body_sha256": "sha256:66a87a504e9f0079fd3c333bccda98c837e00a8018cd3ba2c74530cb46da61b6", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:intf", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip", "parent_id": "xcsh-docs:resources:tunnel:reference", "path": "documentation/resources/tunnel/properties/local_ip/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_ip:ConflictingObjectAttributes:intf,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:intf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip:ConflictingObjectAttributes:intf,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip"], "schema_version": 1, "sections": [{"aliases": ["local ip intf"], "anchor": "section", "description": "Provides the local interface to pick up source IP and network for transporting encapsulated packet.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:intf", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "intf"], "syntax": "block", "type": "object"}, {"aliases": ["local ip ip address"], "anchor": "section", "description": "Provides the configuration to pick up source IP and network for transporting encapsulated packet.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address:ConflictingObjectAttributes:auto,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:auto", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address:ConflictingObjectAttributes:auto,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "type": "conflicts"}], "schema_path": ["local_ip", "ip_address"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS available are - 1. Local Interface - Network Interface from which IP address and network will be selected 2. IP Address - IP address and network can be configured explicitly.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["tunnelCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- local_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - 1. Local Interface - Network Interface from which IP address and network will be
selected 2. IP Address - IP address and network can be configured explicitly.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("intf",
    "ip_address")}
```

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

Terraform syntax:

```terraform
local_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [intf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/): complete subsection reference.

- [ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/): complete subsection reference.
