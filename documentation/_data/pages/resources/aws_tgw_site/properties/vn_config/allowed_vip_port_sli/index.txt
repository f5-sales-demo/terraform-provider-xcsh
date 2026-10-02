---
page_title: "vn_config.allowed_vip_port_sli"
subcategory: ""
description: "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site."
xcsh_docs: {"aliases": ["vn config allowed vip port sli"], "body_bytes": 4583, "body_sha256": "sha256:d8167c528ff568e88cad52a7d13ad167f5555a2432c4b34ffed830859a5c1202", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:disable_allowed_vip_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_https_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_https_port"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "path": "documentation/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1301303321121123-2321223232201013-3102330121020211-2120201131031303-3021010211110131-2331233133123310-2023212112103320-1313123230030221", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:custom_ports,disable_allowed_vip_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:custom_ports,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:custom_ports,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:custom_ports,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:custom_ports,disable_allowed_vip_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:disable_allowed_vip_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:custom_ports,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:use_http_https_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:use_http_https_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:custom_ports,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:use_http_https_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:use_http_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:custom_ports,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:disable_allowed_vip_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:use_http_https_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli:ConflictingObjectAttributes:use_http_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_https_port", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "allowed_vip_port_sli"], "schema_version": 1, "sections": [{"aliases": ["custom ports"], "anchor": "section", "description": "List of Custom port.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-vn_config--allowed_vip_port_sli--custom_ports--port_ranges", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port_sli.custom_ports:RequiredObjectAttributes:port_ranges", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:custom_ports", "type": "requires"}], "schema_path": ["vn_config", "allowed_vip_port_sli", "custom_ports"], "syntax": "block", "type": "object"}, {"aliases": ["disable allowed vip port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:disable_allowed_vip_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port_sli", "disable_allowed_vip_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use http https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_https_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port_sli", "use_http_https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use http port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_http_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port_sli", "use_http_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli:use_https_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port_sli", "use_https_port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.allowed_vip_port_sli

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- vn_config.allowed_vip_port_sli

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port_sli {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/): complete subsection reference.

- [disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/disable_allowed_vip_port/): complete subsection reference.

- [use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_https_port/): complete subsection reference.

- [use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_port/): complete subsection reference.

- [use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_https_port/): complete subsection reference.

## Next pages

- [vn_config.allowed_vip_port_sli.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/)
- [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/disable_allowed_vip_port/)
- [vn_config.allowed_vip_port_sli.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_https_port/)
- [vn_config.allowed_vip_port_sli.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_port/)
- [vn_config.allowed_vip_port_sli.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_https_port/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
