---
page_title: "ingress_gw.allowed_vip_port"
subcategory: "Infrastructure"
description: "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site."
xcsh_docs: {"aliases": ["ingress gw allowed vip port"], "body_bytes": 4532, "body_sha256": "sha256:d586fd4c064b4d3895c6f904c573c41a90f0a85b0994eca69bc0759a87d0b663", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw", "path": "documentation/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,disable_allowed_vip_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,disable_allowed_vip_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw", "allowed_vip_port"], "schema_version": 1, "sections": [{"aliases": ["custom ports"], "anchor": "section", "description": "List of Custom port.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_gw--allowed_vip_port--custom_ports--port_ranges", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port.custom_ports:RequiredObjectAttributes:port_ranges", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "requires"}], "schema_path": ["ingress_gw", "allowed_vip_port", "custom_ports"], "syntax": "block", "type": "object"}, {"aliases": ["disable allowed vip port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "allowed_vip_port", "disable_allowed_vip_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use http https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "allowed_vip_port", "use_http_https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use http port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "allowed_vip_port", "use_http_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["use https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "allowed_vip_port", "use_https_port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.allowed_vip_port

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/)
- ingress_gw.allowed_vip_port

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

- [custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/custom_ports/): complete subsection reference.

- [disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/disable_allowed_vip_port/): complete subsection reference.

- [use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_https_port/): complete subsection reference.

- [use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_port/): complete subsection reference.

- [use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_https_port/): complete subsection reference.

## Next pages

- [ingress_gw.allowed_vip_port.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/custom_ports/)
- [ingress_gw.allowed_vip_port.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/disable_allowed_vip_port/)
- [ingress_gw.allowed_vip_port.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_https_port/)
- [ingress_gw.allowed_vip_port.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_http_port/)
- [ingress_gw.allowed_vip_port.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/use_https_port/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
