---
page_title: "site_subnet_params.subnet_dhcp_server_params"
subcategory: ""
description: "Subnet DHCP parameters will be a subset of network_interface.dhcpserverparameterstype as all features in network_interface.dhcpserverparameterstype may not be supported in a subnet."
xcsh_docs: {"aliases": ["site subnet params subnet dhcp server params"], "body_bytes": 1299, "body_sha256": "sha256:d70f9c380f72daa7c159e79b46a2dff6abeaf6774c2833a613c997ea393bf2e2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params", "parent_id": "xcsh-docs:resources:subnet:properties:site_subnet_params", "path": "documentation/resources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3222100203231312-0002211232302021-3221300102032233-3133331312322010-0103123030010313-2202020001001100-2330211233023033-3003221313122203", "registry_path": "docs/guides/resources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_subnet_params", "subnet_dhcp_server_params"], "schema_version": 1, "sections": [{"aliases": ["site subnet params subnet dhcp server params dhcp networks"], "anchor": "section", "description": "List of networks from which DHCP server can allocate IP addresses.", "document_id": "xcsh-docs:resources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["site_subnet_params", "subnet_dhcp_server_params", "dhcp_networks"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Subnet DHCP parameters will be a subset of network_interface.dhcpserverparameterstype as all features in network_interface.dhcpserverparameterstype may not be supported in a subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["subnetCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params.subnet_dhcp_server_params

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- [site_subnet_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/)
- site_subnet_params.subnet_dhcp_server_params

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
subnet_dhcp_server_params {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/dhcp_networks/): complete subsection reference.
