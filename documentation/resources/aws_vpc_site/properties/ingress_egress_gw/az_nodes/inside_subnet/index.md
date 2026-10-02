---
page_title: "ingress_egress_gw.az_nodes.inside_subnet"
subcategory: "Infrastructure"
description: "Parameters for AWS subnet."
xcsh_docs: {"aliases": ["ingress egress gw az nodes inside subnet"], "body_bytes": 3383, "body_sha256": "sha256:dc27d222700dcd55abaa0a2f60f66d1c75fcfa270ee188db34d922f0e62d9458", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes", "path": "documentation/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3333313031011133-1311220301031012-1031202100120003-0200332010120102-3113033211331030-3301100301001212-0320011313120011-0021220301020301", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-002.md", "relationships": [{"anchor": "schema-ingress_egress_gw--az_nodes--inside_subnet--existing_subnet_id", "enforcement": "provider-schema", "group": "ingress_egress_gw.az_nodes.inside_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.az_nodes.inside_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet"], "schema_version": 1, "sections": [{"aliases": ["existing subnet id"], "anchor": "schema-ingress_egress_gw--az_nodes--inside_subnet--existing_subnet_id", "description": "Exclusive with Information about existing subnet ID.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet", "existing_subnet_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw--az_nodes--inside_subnet--subnet_param--ipv4", "enforcement": "provider-schema", "group": "ingress_egress_gw.az_nodes.inside_subnet.subnet_param:RequiredObjectAttributes:ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:az_nodes:inside_subnet:subnet_param", "type": "requires"}], "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet", "subnet_param"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for AWS subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes.inside_subnet

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/)
- ingress_egress_gw.az_nodes.inside_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
inside_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw--az_nodes--inside_subnet--existing_subnet_id"></a>

### existing_subnet_id property

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/): complete subsection reference.

## Next pages

- [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/az_nodes/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
