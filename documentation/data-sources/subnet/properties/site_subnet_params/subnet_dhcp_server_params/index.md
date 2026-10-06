---
page_title: "site_subnet_params.subnet_dhcp_server_params"
subcategory: ""
description: "Subnet DHCP parameters will be a subset of network_interface.dhcpserverparameterstype as all features in network_interface.dhcpserverparameterstype may not be supported in a subnet."
xcsh_docs: {"aliases": ["site subnet params subnet dhcp server params"], "body_bytes": 1177, "body_sha256": "sha256:8c1248c491c90c2d0f5ae82cf21aeee84c9a5ba2c66eef51baa8995597af7cb3", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params", "parent_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params", "path": "documentation/data-sources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2233201121022323-2002333311211310-3012010103312210-2213220232202332-1313320111021112-1332102200123102-1303013222113103-1110310333100110", "registry_path": "docs/guides/data-sources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_subnet_params", "subnet_dhcp_server_params"], "schema_version": 1, "sections": [{"aliases": ["site subnet params subnet dhcp server params dhcp networks"], "anchor": "section", "description": "List of networks from which DHCP server can allocate IP addresses.", "document_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["site_subnet_params", "subnet_dhcp_server_params", "dhcp_networks"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Subnet DHCP parameters will be a subset of network_interface.dhcpserverparameterstype as all features in network_interface.dhcpserverparameterstype may not be supported in a subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params.subnet_dhcp_server_params

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/)
- [site_subnet_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/)
- site_subnet_params.subnet_dhcp_server_params

<a id="section"></a>

Type: `"single"`. Computed.

Subnet DHCP parameters will be a subset of network\_interface.dhcpserverparameterstype as all
features in network\_interface.dhcpserverparameterstype may not be supported in a subnet.

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

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/dhcp_networks/): complete subsection reference.
