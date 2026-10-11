---
page_title: "vpc_attachments.vpc_list"
subcategory: ""
description: "List of VPC attachments to transit gateway."
xcsh_docs: {"aliases": ["vpc attachments vpc list"], "body_bytes": 2422, "body_sha256": "sha256:16e145d2d31a22e9f9bc0312f00f61399573c0300064faa4566e3fca92eeb5bf", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments", "path": "documentation/data-sources/aws_tgw_site/properties/vpc_attachments/vpc_list/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0100101120202303-3222030000301000-3222211001000330-0130021323301301-0013002010232121-3120121121232301-1022321302331302-3203311222310020", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vpc_attachments", "vpc_list"], "schema_version": 1, "sections": [{"aliases": ["vpc attachments vpc list labels"], "anchor": "section", "description": "Add labels for the VPC attachment. These labels can then be used in policies such as enhanced firewall.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vpc_attachments", "vpc_list", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["vpc attachments vpc list vpc id"], "anchor": "schema-vpc_attachments--vpc_list--vpc_id", "description": "Information about existing VPC.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vpc_attachments", "vpc_list", "vpc_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vpc_attachments/vpc_list/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of VPC attachments to transit gateway.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
