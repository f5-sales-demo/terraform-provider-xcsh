---
page_title: "vpc_attachments.vpc_list"
subcategory: ""
description: "List of VPC attachments to transit gateway."
xcsh_docs: {"aliases": ["vpc attachments vpc list"], "body_bytes": 3127, "body_sha256": "sha256:f14832b2cf68ef441643e3b9b181a069c901ddfe14964d7c33485deeb0e79dcd", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments", "path": "documentation/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2222113021122020-0012033011131223-1212101133221212-1133313010011213-2032103312331322-1002312232020110-2300122030323031-0113000122111302", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vpc_attachments", "vpc_list"], "schema_version": 1, "sections": [{"aliases": ["vpc attachments vpc list labels"], "anchor": "section", "description": "Add labels for the VPC attachment. These labels can then be used in policies such as enhanced firewall.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vpc_attachments", "vpc_list", "labels"], "syntax": "block", "type": "object"}, {"aliases": ["vpc attachments vpc list vpc id"], "anchor": "schema-vpc_attachments--vpc_list--vpc_id", "description": "Information about existing VPC.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vpc_attachments", "vpc_list", "vpc_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of VPC attachments to transit gateway.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc_attachments.vpc_list

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/)
- vpc_attachments.vpc_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of VPC attachments to transit gateway.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
vpc_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/labels/): complete subsection reference.

<a id="schema-vpc_attachments--vpc_list--vpc_id"></a>

### vpc_id property

Type: `"string"`. Optional.

VPC ID. Information about existing VPC.

Upstream description:

Information about existing VPC.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

## Next pages

- [vpc_attachments.vpc_list.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/labels/)
- [vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
