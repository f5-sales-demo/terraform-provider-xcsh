---
page_title: "ingress_gw"
subcategory: "Infrastructure"
description: "Single interface AWS ingress site."
xcsh_docs: {"aliases": ["ingress gw"], "body_bytes": 3627, "body_sha256": "sha256:55cb731685365373a072a79b89e0e631a386d9ecf9da3772b62fb48a9aa9d658", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port", "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes", "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:performance_enhancement_mode"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "documentation/resources/aws_vpc_site/properties/ingress_gw/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-003.md", "relationships": [{"anchor": "schema-ingress_gw--aws_certified_hw", "enforcement": "provider-schema", "group": "ingress_gw:RequiredObjectAttributes:aws_certified_hw,az_nodes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw:RequiredObjectAttributes:aws_certified_hw,az_nodes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw"], "schema_version": 1, "sections": [{"aliases": ["allowed vip port"], "anchor": "section", "description": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,disable_allowed_vip_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,disable_allowed_vip_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:disable_allowed_vip_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_http_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_http_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:custom_ports,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:disable_allowed_vip_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_https_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port:ConflictingObjectAttributes:use_http_port,use_https_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:use_https_port", "type": "conflicts"}], "schema_path": ["ingress_gw", "allowed_vip_port"], "syntax": "block", "type": "object"}, {"aliases": ["aws certified hw"], "anchor": "schema-ingress_gw--aws_certified_hw", "description": "Name for AWS certified hardware.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "aws_certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["az nodes"], "anchor": "section", "description": "Only Single AZ or Three AZ(s) nodes are supported currently.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-ingress_gw--az_nodes--aws_az_name", "enforcement": "provider-schema", "group": "ingress_gw.az_nodes:RequiredListObjectAttributes:aws_az_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes", "type": "requires"}], "schema_path": ["ingress_gw", "az_nodes"], "syntax": "block", "type": "object"}, {"aliases": ["performance enhancement mode"], "anchor": "section", "description": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:performance_enhancement_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l7_enhanced", "type": "conflicts"}], "schema_path": ["ingress_gw", "performance_enhancement_mode"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_gw/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Single interface AWS ingress site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- ingress_gw

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS Ingress Gateway. Single interface AWS ingress site.

Upstream description:

Single interface AWS ingress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_certified_hw",
    "az_nodes")}
```

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
ingress_gw {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/): complete subsection reference.

<a id="schema-ingress_gw--aws_certified_hw"></a>

### aws_certified_hw property

Type: `"string"`. Optional.

\[Enum: aws-byol-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The only
possible value is \`aws-byol-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("aws-byol-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-voltmesh"
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
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/az_nodes/): complete subsection reference.

- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/performance_enhancement_mode/): complete subsection reference.

## Next pages

- [ingress_gw.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/)
- [ingress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/az_nodes/)
- [ingress_gw.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/performance_enhancement_mode/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
