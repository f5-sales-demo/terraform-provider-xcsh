---
page_title: "ingress_egress_gw"
subcategory: "Infrastructure"
description: "ingress_egress_gw for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 13067, "body_sha256": "sha256:b6bcc023040dbcff55a1c8a87815ecdeeec81e11c15ff85e4e19ce165cb4b384", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:active_forward_proxy_policies", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:active_network_policies", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:allowed_vip_port_sli", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:forward_proxy_allow_all", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:global_network_list", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:no_dc_cluster_group", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:no_forward_proxy", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:no_global_network", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:no_inside_static_routes", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:no_network_policy", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:no_outside_static_routes", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:outside_static_routes", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:sm_connection_public_ip", "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:sm_connection_pvt_ip"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "documentation/resources/aws_vpc_site/properties/ingress_egress_gw/index.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["ingress_egress_gw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- ingress_egress_gw

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface AWS ingress/egress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_certified_hw",
    "az_nodes"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("inside_static_routes",
    "no_inside_static_routes"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/#section)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/#section)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ingress_egress_gw {
  # Configure direct properties listed below.
}
```

## Direct properties

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/active_network_policies/): complete subsection reference.

- [allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/): complete subsection reference.

- [allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/): complete subsection reference.

<a id="schema-ingress_egress_gw--aws_certified_hw"></a>

### aws_certified_hw property

Type: `"string"`. Optional.

\[Enum: aws-byol-multi-nic-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The
only possible value is \`aws-byol-multi-nic-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("aws-byol-multi-nic-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-multi-nic-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/): complete subsection reference.

- [dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/): complete subsection reference.

- [dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/forward_proxy_allow_all/): complete subsection reference.

- [global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/): complete subsection reference.

- [inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_dc_cluster_group/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_forward_proxy/): complete subsection reference.

- [no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_global_network/): complete subsection reference.

- [no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_inside_static_routes/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_network_policy/): complete subsection reference.

- [no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_outside_static_routes/): complete subsection reference.

- [outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/): complete subsection reference.

- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/sm_connection_pvt_ip/): complete subsection reference.

## Next pages

- [ingress_egress_gw.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/)
- [ingress_egress_gw.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/active_forward_proxy_policies/)
- [ingress_egress_gw.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/active_network_policies/)
- [ingress_egress_gw.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port/)
- [ingress_egress_gw.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/allowed_vip_port_sli/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/)
- [ingress_egress_gw.dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/)
- [ingress_egress_gw.dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/)
- [ingress_egress_gw.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/forward_proxy_allow_all/)
- [ingress_egress_gw.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/global_network_list/)
- [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/)
- [ingress_egress_gw.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_dc_cluster_group/)
- [ingress_egress_gw.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_forward_proxy/)
- [ingress_egress_gw.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_global_network/)
- [ingress_egress_gw.no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_inside_static_routes/)
- [ingress_egress_gw.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_network_policy/)
- [ingress_egress_gw.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/no_outside_static_routes/)
- [ingress_egress_gw.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/)
- [ingress_egress_gw.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/)
- [ingress_egress_gw.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/sm_connection_public_ip/)
- [ingress_egress_gw.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/sm_connection_pvt_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
