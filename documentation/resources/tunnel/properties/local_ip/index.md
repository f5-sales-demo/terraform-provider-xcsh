---
page_title: "local_ip"
subcategory: ""
description: "Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS available are - 1. Local Interface - Network Interface from which IP address and network will be selected 2. IP Address - IP address and network can be configured explicitly."
xcsh_docs: {"aliases": ["local ip"], "body_bytes": 2257, "body_sha256": "sha256:89c679a59809da8721a3e7f82e2e58b06aaf24ee073d4d45beea53a06bf4c55d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:intf", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip", "parent_id": "xcsh-docs:resources:tunnel:reference", "path": "documentation/resources/tunnel/properties/local_ip/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3310002120310311-3233112233301322-3202003303031322-3200110222231231-1110211002321202-2023230133212221-3020020030123203-3312013200013013", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_ip:ConflictingObjectAttributes:intf,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:intf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip:ConflictingObjectAttributes:intf,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip"], "schema_version": 1, "sections": [{"aliases": ["intf"], "anchor": "section", "description": "Provides the local interface to pick up source IP and network for transporting encapsulated packet.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:intf", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "intf"], "syntax": "block", "type": "object"}, {"aliases": ["ip address"], "anchor": "section", "description": "Provides the configuration to pick up source IP and network for transporting encapsulated packet.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address:ConflictingObjectAttributes:auto,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:auto", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address:ConflictingObjectAttributes:auto,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "type": "conflicts"}], "schema_path": ["local_ip", "ip_address"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS available are - 1. Local Interface - Network Interface from which IP address and network will be selected 2. IP Address - IP address and network can be configured explicitly.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - &#8203;1. Local Interface - Network Interface from which IP address and network will
be selected &#8203;2. IP Address - IP address and network can be configured explicitly.

Provider validators and defaults (from schema source):

```go
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

## Next pages

- [local_ip.intf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/intf/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
