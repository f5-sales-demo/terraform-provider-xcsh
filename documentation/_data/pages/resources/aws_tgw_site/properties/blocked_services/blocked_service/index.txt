---
page_title: "blocked_services.blocked_service"
subcategory: ""
description: "Blocking or denial configuration"
xcsh_docs: {"aliases": ["blocked services blocked service"], "body_bytes": 9135, "body_sha256": "sha256:3e1e785a5c47567187a958863c6bdc9b9feed4b84c00dee038b835f597876035", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:dns", "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:ssh", "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:web_user_interface"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services", "path": "documentation/resources/aws_tgw_site/properties/blocked_services/blocked_service/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,ssh", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,ssh", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:ssh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:ssh,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:ssh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:web_user_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:ssh,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:web_user_interface", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["blocked_services", "blocked_service"], "schema_version": 1, "sections": [{"aliases": ["blocked services blocked service dns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:dns", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_services", "blocked_service", "dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["blocked services blocked service network type"], "anchor": "schema-blocked_services--blocked_service--network_type", "description": "Different types of virtual networks understood by the system Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL provides connectivity to public (outside) network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type is local to every site. Two virtual", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["VIRTUAL_NETWORK_GLOBAL", "VIRTUAL_NETWORK_IP_AUTO", "VIRTUAL_NETWORK_IP_FABRIC", "VIRTUAL_NETWORK_MANAGEMENT", "VIRTUAL_NETWORK_PER_SITE", "VIRTUAL_NETWORK_PUBLIC", "VIRTUAL_NETWORK_SEGMENT", "VIRTUAL_NETWORK_SITE_LOCAL", "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE", "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE", "VIRTUAL_NETWORK_SITE_SERVICE", "VIRTUAL_NETWORK_SRV6_NETWORK", "VIRTUAL_NETWORK_VER_INTERNAL", "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_services", "blocked_service", "network_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["blocked services blocked service ssh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:ssh", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_services", "blocked_service", "ssh"], "syntax": "attribute", "type": "object"}, {"aliases": ["blocked services blocked service web user interface"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services:blocked_service:web_user_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["blocked_services", "blocked_service", "web_user_interface"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/blocked_services/blocked_service/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Blocking or denial configuration", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services.blocked_service

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/)
- blocked_services.blocked_service

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
blocked_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/blocked_service/dns/): complete subsection reference.

<a id="schema-blocked_services--blocked_service--network_type"></a>

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

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [ssh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/blocked_service/ssh/): complete subsection reference.

- [web_user_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/blocked_service/web_user_interface/): complete subsection reference.
