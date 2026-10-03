---
page_title: "blocked_services.blocked_service"
subcategory: ""
description: "Blocking or denial configuration"
xcsh_docs: {"aliases": ["blocked services blocked service"], "body_bytes": 8291, "body_sha256": "sha256:8783bd12687b04963544d503cdf64ac9aaae0f981ee8fac6e7b9a01a6fd0aae7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:blocked_services:blocked_service:dns", "xcsh-docs:data-sources:securemesh_site:properties:blocked_services:blocked_service:ssh", "xcsh-docs:data-sources:securemesh_site:properties:blocked_services:blocked_service:web_user_interface"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:blocked_services:blocked_service", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:blocked_services", "path": "documentation/data-sources/securemesh_site/properties/blocked_services/blocked_service/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["blocked_services", "blocked_service"], "schema_version": 1, "sections": [{"aliases": ["blocked services blocked service dns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:blocked_services:blocked_service:dns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_services", "blocked_service", "dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["blocked services blocked service network type"], "anchor": "schema-blocked_services--blocked_service--network_type", "description": "Different types of virtual networks understood by the system Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL provides connectivity to public (outside) network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type is local to every site. Two virtual", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:blocked_services:blocked_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_services", "blocked_service", "network_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["blocked services blocked service ssh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:blocked_services:blocked_service:ssh", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_services", "blocked_service", "ssh"], "syntax": "attribute", "type": "object"}, {"aliases": ["blocked services blocked service web user interface"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:blocked_services:blocked_service:web_user_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_services", "blocked_service", "web_user_interface"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/blocked_services/blocked_service/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Blocking or denial configuration", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services.blocked_service

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/blocked_services/)
- blocked_services.blocked_service

<a id="section"></a>

Type: `"list"`. Computed.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/blocked_services/blocked_service/dns/): complete subsection reference.

<a id="schema-blocked_services--blocked_service--network_type"></a>

### network_type property

Type: `"string"`. Computed.

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

- [ssh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/blocked_services/blocked_service/ssh/): complete subsection reference.

- [web_user_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/blocked_services/blocked_service/web_user_interface/): complete subsection reference.

## Next pages

- [blocked_services.blocked_service.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/blocked_services/blocked_service/dns/)
- [blocked_services.blocked_service.ssh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/blocked_services/blocked_service/ssh/)
- [blocked_services.blocked_service.web_user_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/blocked_services/blocked_service/web_user_interface/)
- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/blocked_services/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
