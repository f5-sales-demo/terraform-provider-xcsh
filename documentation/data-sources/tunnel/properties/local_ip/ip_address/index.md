---
page_title: "local_ip.ip_address"
subcategory: ""
description: "Provides the configuration to pick up source IP and network for transporting encapsulated packet."
xcsh_docs: {"aliases": ["local ip ip address"], "body_bytes": 2085, "body_sha256": "sha256:a7a0907c128a15da2e7897204cbb167af18350c40d7fe13f27148cdb8389fb53", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:auto", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip", "path": "documentation/data-sources/tunnel/properties/local_ip/ip_address/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0220323001102331-0213311030330111-2111232213210111-1033120330031000-0211330211223312-1331320130322020-0330300311013221-0200103200320110", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_ip", "ip_address"], "schema_version": 1, "sections": [{"aliases": ["local ip ip address auto"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:auto", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_ip", "ip_address", "auto"], "syntax": "attribute", "type": "object"}, {"aliases": ["local ip ip address ip address"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "ip_address", "ip_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["local ip ip address virtual network type"], "anchor": "section", "description": "Different types of virtual networks understood by the system.", "document_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["local_ip", "ip_address", "virtual_network_type"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/ip_address/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Provides the configuration to pick up source IP and network for transporting encapsulated packet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/)
- local_ip.ip_address

<a id="section"></a>

Type: `"single"`. Computed.

Provides the configuration to pick up source IP and network for transporting encapsulated packet.

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

## Direct properties

- [auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/auto/): complete subsection reference.

- [ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/ip_address/): complete subsection reference.

- [virtual_network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/): complete subsection reference.

## Next pages

- [local_ip.ip_address.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/auto/)
- [local_ip.ip_address.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/ip_address/)
- [local_ip.ip_address.virtual_network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
