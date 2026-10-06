---
page_title: "f5_big_ip_aws_service"
subcategory: ""
description: "Virtual BIG-IP specification for AWS."
xcsh_docs: {"aliases": ["f5 big ip aws service"], "body_bytes": 6544, "body_sha256": "sha256:ef5e2dcd693453b07d1c4156af24c975f06c1a920f9a267f419e38e2b2396118", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:admin_password", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "parent_id": "xcsh-docs:resources:nfv_service:reference", "path": "documentation/resources/nfv_service/properties/f5_big_ip_aws_service/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [{"anchor": "schema-f5_big_ip_aws_service--admin_username", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service:RequiredObjectAttributes:admin_username,nodes,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "type": "requires"}, {"anchor": "schema-f5_big_ip_aws_service--ssh_key", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service:RequiredObjectAttributes:admin_username,nodes,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service:RequiredObjectAttributes:admin_username,nodes,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service"], "schema_version": 1, "sections": [{"aliases": ["f5 big ip aws service admin password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:admin_password", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:admin_password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:admin_password:clear_secret_info", "type": "conflicts"}], "schema_path": ["f5_big_ip_aws_service", "admin_password"], "syntax": "block", "type": "object"}, {"aliases": ["f5 big ip aws service admin username"], "anchor": "schema-f5_big_ip_aws_service--admin_username", "description": "Admin Username for BIG-IP.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "admin_username"], "syntax": "attribute", "type": "string"}, {"aliases": ["f5 big ip aws service aws tgw site params"], "anchor": "section", "description": "BIG-IP AWS TGW site specification.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["f5_big_ip_aws_service", "aws_tgw_site_params"], "syntax": "block", "type": "object"}, {"aliases": ["f5 big ip aws service endpoint service"], "anchor": "section", "description": "Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies the destination with a new destination address.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-f5_big_ip_aws_service--endpoint_service--configured_vip", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:automatic_vip,configured_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:advertise_on_slo_ip,advertise_on_slo_ip_external", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:advertise_on_slo_ip,disable_advertise_on_slo_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:advertise_on_slo_ip,advertise_on_slo_ip_external", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip_external", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:advertise_on_slo_ip_external,disable_advertise_on_slo_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:advertise_on_slo_ip_external", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:automatic_vip,configured_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:automatic_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_tcp_ports,default_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_tcp_ports,http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_tcp_ports,https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_tcp_ports,no_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_udp_ports,no_udp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_tcp_ports,default_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:default_tcp_ports,http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:default_tcp_ports,https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:default_tcp_ports,no_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:default_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:advertise_on_slo_ip,disable_advertise_on_slo_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:disable_advertise_on_slo_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:advertise_on_slo_ip_external,disable_advertise_on_slo_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:disable_advertise_on_slo_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_tcp_ports,http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:default_tcp_ports,http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:http_port,https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:http_port,no_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_tcp_ports,https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:default_tcp_ports,https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:http_port,https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:https_port,no_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_tcp_ports,no_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:default_tcp_ports,no_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:http_port,no_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:https_port,no_tcp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_tcp_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.endpoint_service:ConflictingObjectAttributes:custom_udp_ports,no_udp_ports", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:no_udp_ports", "type": "conflicts"}], "schema_path": ["f5_big_ip_aws_service", "endpoint_service"], "syntax": "block", "type": "object"}, {"aliases": ["f5 big ip aws service market place image"], "anchor": "section", "description": "BIG-IP AWS Pay as You Go Image Selection.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.market_place_image:ConflictingObjectAttributes:awafpay_g200_mbps,awafpay_g3_gbps", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g200_mbps", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.market_place_image:ConflictingObjectAttributes:awafpay_g200_mbps,awafpay_g3_gbps", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:market_place_image:awafpay_g3_gbps", "type": "conflicts"}], "schema_path": ["f5_big_ip_aws_service", "market_place_image"], "syntax": "block", "type": "object"}, {"aliases": ["f5 big ip aws service nodes"], "anchor": "section", "description": "Specify how and where the service nodes are spawned.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-f5_big_ip_aws_service--nodes--tunnel_prefix", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:ConflictingListObjectAttributes:automatic_prefix,tunnel_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:ConflictingListObjectAttributes:automatic_prefix,tunnel_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:automatic_prefix", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:ConflictingListObjectAttributes:mgmt_subnet,reserved_mgmt_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:ConflictingListObjectAttributes:mgmt_subnet,reserved_mgmt_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes:reserved_mgmt_subnet", "type": "conflicts"}, {"anchor": "schema-f5_big_ip_aws_service--nodes--aws_az_name", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:RequiredListObjectAttributes:aws_az_name,node_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "type": "requires"}, {"anchor": "schema-f5_big_ip_aws_service--nodes--node_name", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.nodes:RequiredListObjectAttributes:aws_az_name,node_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:nodes", "type": "requires"}], "schema_path": ["f5_big_ip_aws_service", "nodes"], "syntax": "block", "type": "object"}, {"aliases": ["f5 big ip aws service ssh key"], "anchor": "schema-f5_big_ip_aws_service--ssh_key", "description": "Public SSH key for accessing the Big IP nodes.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "ssh_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["f5 big ip aws service tags"], "anchor": "schema-f5_big_ip_aws_service--tags", "description": "AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify, organize, search for, and filter resources in AWS console.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "tags"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Virtual BIG-IP specification for AWS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- f5_big_ip_aws_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: f5\_big\_ip\_aws\_service, palo\_alto\_fw\_service\] Virtual BIG-IP AWS. Virtual BIG-IP
specification for AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("admin_username",
    "nodes",
    "ssh_key")}
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
  "x-ves-oneof-field-image_choice": "[\"market_place_image\"]",
  "x-ves-oneof-field-site_type_choice": "[\"aws_tgw_site_params\"]"
}
```

OneOf alternatives in this subsection:

- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/#section)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
f5_big_ip_aws_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/admin_password/): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--admin_username"></a>

### admin_username property

Type: `"string"`. Optional.

Admin Username. Admin Username for BIG-IP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [aws_tgw_site_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/): complete subsection reference.

- [endpoint_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/): complete subsection reference.

- [market_place_image](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/market_place_image/): complete subsection reference.

- [nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/nodes/): complete subsection reference.

<a id="schema-f5_big_ip_aws_service--ssh_key"></a>

### ssh_key property

Type: `"string"`. Optional.

Public SSH key for accessing the Big IP nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-f5_big_ip_aws_service--tags"></a>

### tags property

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":40},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":127,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"127\",\"ves.io.schema.rules.map.max_pairs\":\"40\",\"ves.io.schema.rules.map.values.string.max_len\":\"255\"},\"values\":{\"maxLength\":255,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 40
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 127,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "127",
      "ves.io.schema.rules.map.max_pairs": "40",
      "ves.io.schema.rules.map.values.string.max_len": "255"
    },
    "values": {
      "maxLength": 255,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```
