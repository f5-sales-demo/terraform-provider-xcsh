---
page_title: "local_ip.ip_address"
subcategory: ""
description: "Provides the configuration to pick up source IP and network for transporting encapsulated packet."
xcsh_docs: {"aliases": ["local ip ip address"], "body_bytes": 2335, "body_sha256": "sha256:1e64b69ffef9e3f908584beb00b8c47a15688e300726677ceb361f8f8575a04d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:ip_address:auto", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip", "path": "documentation/resources/tunnel/properties/local_ip/ip_address/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0101220131022022-1000210322332102-0021030332331230-1223222200000203-3112120121021221-1333110220021312-3320303000313110-3032130312212310", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address:ConflictingObjectAttributes:auto,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:auto", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address:ConflictingObjectAttributes:auto,ip_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "ip_address"], "schema_version": 1, "sections": [{"aliases": ["auto"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:auto", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_ip", "ip_address", "auto"], "syntax": "attribute", "type": "object"}, {"aliases": ["ip address"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.ip_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.ip_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.ip_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.ip_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.ip_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.ip_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:ipv6", "type": "conflicts"}], "schema_path": ["local_ip", "ip_address", "ip_address"], "syntax": "block", "type": "object"}, {"aliases": ["virtual network type"], "anchor": "section", "description": "Different types of virtual networks understood by the system.", "document_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:public,site_local", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:public,site_local_inside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:public,site_local", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:site_local,site_local_inside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:public,site_local_inside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_ip.ip_address.virtual_network_type:ConflictingObjectAttributes:site_local,site_local_inside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside", "type": "conflicts"}], "schema_path": ["local_ip", "ip_address", "virtual_network_type"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Provides the configuration to pick up source IP and network for transporting encapsulated packet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- local_ip.ip_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Provides the configuration to pick up source IP and network for transporting encapsulated packet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto",
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
  "x-ves-oneof-field-type": "[\"auto\",\"ip_address\"]"
}
```

Terraform syntax:

```terraform
ip_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/auto/): complete subsection reference.

- [ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/): complete subsection reference.

- [virtual_network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/): complete subsection reference.

## Next pages

- [local_ip.ip_address.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/auto/)
- [local_ip.ip_address.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/ip_address/)
- [local_ip.ip_address.virtual_network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/local_ip/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
