---
page_title: "site_subnet_params.subnet_dhcp_server_params"
subcategory: ""
description: "site_subnet_params.subnet_dhcp_server_params for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 1383, "body_sha256": "sha256:f6dc289cf87b7f3afc079ceaae087e34873e8343dd5d299aeca4e4e5e1c956ec", "canonical_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params", "child_ids": ["xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params:dhcp_networks"], "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params:subnet_dhcp_server_params", "parent_id": "xcsh-docs:data-sources:subnet:properties:site_subnet_params", "path": "docs/guides/data-sources--subnet--properties--site_subnet_params--subnet_dhcp_server_params.md", "provider_name": "subnet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_subnet_params", "subnet_dhcp_server_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/properties/site_subnet_params/subnet_dhcp_server_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_subnet_params.subnet_dhcp_server_params for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# site_subnet_params.subnet_dhcp_server_params

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md)
- [Property reference](data-sources--subnet--reference.md)
- [site_subnet_params](data-sources--subnet--properties--site_subnet_params.md)
- site_subnet_params.subnet_dhcp_server_params

<a id="section"></a>

Type: `"single"`. Computed.

Subnet DHCP parameters will be a subset of network\_interface.dhcpserverparameterstype as all
features in network\_interface.dhcpserverparameterstype may not be supported in a subnet.

Upstream description:

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

- [dhcp_networks](data-sources--subnet--properties--site_subnet_params--subnet_dhcp_server_params--dhcp_networks.md): complete subsection reference.

## Next pages

- [site_subnet_params.subnet_dhcp_server_params.dhcp_networks](data-sources--subnet--properties--site_subnet_params--subnet_dhcp_server_params--dhcp_networks.md)
- [site_subnet_params](data-sources--subnet--properties--site_subnet_params.md)
- [xcsh_subnet](../data-sources/subnet.md)
