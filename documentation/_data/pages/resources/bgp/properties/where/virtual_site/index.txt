---
page_title: "where.virtual_site"
subcategory: ""
description: "A reference to virtual_site object."
xcsh_docs: {"aliases": ["where virtual site"], "body_bytes": 8595, "body_sha256": "sha256:2e0768c1601b40452c2957ac7ebd060586da2a0c11e50d9d4e6315eef4bb367d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:where:virtual_site:disable_internet_vip", "xcsh-docs:resources:bgp:properties:where:virtual_site:enable_internet_vip", "xcsh-docs:resources:bgp:properties:where:virtual_site:ref"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:where:virtual_site", "parent_id": "xcsh-docs:resources:bgp:properties:where", "path": "documentation/resources/bgp/properties/where/virtual_site/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0231022200202313-2033330230223132-0302221102233021-2022123222113212-2201231331231000-0221210311023333-0121201210301213-2231101310213200", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:disable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:enable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:RequiredObjectAttributes:ref", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:ref", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "virtual_site"], "schema_version": 1, "sections": [{"aliases": ["where virtual site disable internet vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:disable_internet_vip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["where", "virtual_site", "disable_internet_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["where virtual site enable internet vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:enable_internet_vip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["where", "virtual_site", "enable_internet_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["where virtual site network type"], "anchor": "schema-where--virtual_site--network_type", "description": "Different types of virtual networks understood by the system Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL provides connectivity to public (outside) network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type is local to every site. Two virtual", "document_id": "xcsh-docs:resources:bgp:properties:where:virtual_site", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["where", "virtual_site", "network_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["where virtual site ref"], "anchor": "section", "description": "A virtual_site direct reference.", "document_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["where", "virtual_site", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/where/virtual_site/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A reference to virtual_site object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgpCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_site

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/)
- where.virtual_site

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/disable_internet_vip/): complete subsection reference.

- [enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/enable_internet_vip/): complete subsection reference.

<a id="schema-where--virtual_site--network_type"></a>

### network_type property

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/ref/): complete subsection reference.

## Next pages

- [where.virtual_site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/disable_internet_vip/)
- [where.virtual_site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/enable_internet_vip/)
- [where.virtual_site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/ref/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
