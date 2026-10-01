---
page_title: "vpc"
subcategory: "Infrastructure"
description: "vpc for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2044, "body_sha256": "sha256:189acfa1788febb24dc55891c136a24a5f8e1f0be44b771994c1d0019461ccab", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:vpc", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:vpc:new_vpc"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:vpc", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:reference", "path": "docs/guides/data-sources--aws_vpc_site--properties--vpc.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vpc"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/vpc/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vpc for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- vpc

<a id="section"></a>

Type: `"single"`. Computed.

Defines choice about AWS VPC for a view.

Upstream description:

This defines choice about AWS VPC for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"new_vpc\",\"vpc_id\"]"
}
```

## Direct properties

- [new_vpc](data-sources--aws_vpc_site--properties--vpc--new_vpc.md): complete subsection reference.

<a id="schema-vpc--vpc_id"></a>

### vpc_id property

Type: `"string"`. Computed.

Exclusive with \[new\_vpc\] Information about existing VPC ID.

Upstream description:

Exclusive with \[new\_vpc\] Information about existing VPC ID.

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

- [vpc.new_vpc](data-sources--aws_vpc_site--properties--vpc--new_vpc.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
