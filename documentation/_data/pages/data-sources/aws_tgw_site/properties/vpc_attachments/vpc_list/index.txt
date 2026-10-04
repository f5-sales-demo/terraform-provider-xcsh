---
page_title: "vpc_attachments.vpc_list"
subcategory: ""
description: "List of VPC attachments to transit gateway."
xcsh_docs: {"aliases": ["vpc attachments vpc list"], "body_bytes": 2894, "body_sha256": "sha256:a2e191e7ef84f2dd03e6c46c97a2fb049a6124985c1c1cee175d785b154ea262", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments", "path": "documentation/data-sources/aws_tgw_site/properties/vpc_attachments/vpc_list/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0100101120202303-3222030000301000-3222211001000330-0130021323301301-0013002010232121-3120121121232301-1022321302331302-3203311222310020", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vpc_attachments", "vpc_list"], "schema_version": 1, "sections": [{"aliases": ["vpc attachments vpc list labels"], "anchor": "section", "description": "Add labels for the VPC attachment. These labels can then be used in policies such as enhanced firewall.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vpc_attachments", "vpc_list", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["vpc attachments vpc list vpc id"], "anchor": "schema-vpc_attachments--vpc_list--vpc_id", "description": "Information about existing VPC.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vpc_attachments", "vpc_list", "vpc_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vpc_attachments/vpc_list/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of VPC attachments to transit gateway.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc_attachments.vpc_list

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vpc_attachments/)
- vpc_attachments.vpc_list

<a id="section"></a>

Type: `"list"`. Computed.

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

## Direct properties

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vpc_attachments/vpc_list/labels/): complete subsection reference.

<a id="schema-vpc_attachments--vpc_list--vpc_id"></a>

### vpc_id property

Type: `"string"`. Computed.

VPC ID. Information about existing VPC.

Upstream description:

Information about existing VPC.

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

- [vpc_attachments.vpc_list.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vpc_attachments/vpc_list/labels/)
- [vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vpc_attachments/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
