---
page_title: "aws_provider"
subcategory: ""
description: "Create AWS Provider Type."
xcsh_docs: {"aliases": ["aws provider"], "body_bytes": 2064, "body_sha256": "sha256:fd5986fae5cbf951bbd866162ada77478ce6104c5a188e1c890943b5305c1db6", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role", "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider", "parent_id": "xcsh-docs:data-sources:cloud_user_account:reference", "path": "documentation/data-sources/cloud_user_account/properties/aws_provider/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0102211100331230-1321103331002200-1232102100033030-2321023132121332-3120112320232012-1031323313003332-2202122131032201-3331030302331011", "registry_path": "docs/guides/data-sources--cloud_user_account--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider"], "schema_version": 1, "sections": [{"aliases": ["aws provider aws account number"], "anchor": "schema-aws_provider--aws_account_number", "description": "12 Digit Account Number.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_account_number"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws provider aws assume role"], "anchor": "section", "description": "AWS Assume Role to Handle Delegated Access.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_assume_role"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws provider aws secret key"], "anchor": "section", "description": "AWS Programmatic Access Credentials type.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_secret_key"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/properties/aws_provider/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Create AWS Provider Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/)
- aws_provider

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for aws provider.

Additional upstream details:

Create AWS Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-aws_authentication_type": "[\"aws_assume_role\",\"aws_secret_key\"]"
}
```

## Direct properties

<a id="schema-aws_provider--aws_account_number"></a>

### aws_account_number property

Type: `"string"`. Computed.

Account Number. 12 Digit Account Number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  }
}
```

- [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_assume_role/): complete subsection reference.

- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/): complete subsection reference.
