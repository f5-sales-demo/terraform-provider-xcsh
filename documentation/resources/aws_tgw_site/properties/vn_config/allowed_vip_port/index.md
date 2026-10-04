---
page_title: "vn_config.allowed_vip_port"
subcategory: ""
description: "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site."
xcsh_docs: {"aliases": ["vn config allowed vip port"], "body_bytes": 4511, "body_sha256": "sha256:023a90bc57f096c7f45be08a1191ef5663b86498882ec33ce7d162b3a2d43202", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:disable_allowed_vip_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_https_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_https_port"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "path": "documentation/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3003102010221332-0101222211103223-2312030111033332-2212001213031132-2301133033032321-2201310021330003-0201123333111201-2300332211300213", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:custom_ports,disable_allowed_vip_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:custom_ports,disable_allowed_vip_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:use_http_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port:ConflictingObjectAttributes:use_http_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_https_port", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "allowed_vip_port"], "schema_version": 1, "sections": [{"aliases": ["vn config allowed vip port custom ports"], "anchor": "section", "description": "List of Custom port.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-vn_config--allowed_vip_port--custom_ports--port_ranges", "enforcement": "provider-schema", "group": "vn_config.allowed_vip_port.custom_ports:RequiredObjectAttributes:port_ranges", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:custom_ports", "type": "requires"}], "schema_path": ["vn_config", "allowed_vip_port", "custom_ports"], "syntax": "block", "type": "object"}, {"aliases": ["vn config allowed vip port disable allowed vip port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:disable_allowed_vip_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "disable_allowed_vip_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config allowed vip port use http https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_https_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "use_http_https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config allowed vip port use http port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_http_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "use_http_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config allowed vip port use https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port:use_https_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "allowed_vip_port", "use_https_port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.allowed_vip_port

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- vn_config.allowed_vip_port

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
allowed_vip_port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/custom_ports/): complete subsection reference.

- [disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/disable_allowed_vip_port/): complete subsection reference.

- [use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_https_port/): complete subsection reference.

- [use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_port/): complete subsection reference.

- [use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_https_port/): complete subsection reference.

## Next pages

- [vn_config.allowed_vip_port.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/custom_ports/)
- [vn_config.allowed_vip_port.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/disable_allowed_vip_port/)
- [vn_config.allowed_vip_port.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_https_port/)
- [vn_config.allowed_vip_port.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_port/)
- [vn_config.allowed_vip_port.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_https_port/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
