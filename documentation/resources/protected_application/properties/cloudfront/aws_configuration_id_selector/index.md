---
page_title: "cloudfront.aws_configuration_id_selector"
subcategory: ""
description: "List of CloudFront distributions."
xcsh_docs: {"aliases": ["cloudfront aws configuration id selector"], "body_bytes": 2991, "body_sha256": "sha256:0ca25ab065141d0958c206b276e0476ae98ec7aed306c13a114b1f235819b733", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:aws_configuration_id_selector", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront", "path": "documentation/resources/protected_application/properties/cloudfront/aws_configuration_id_selector/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2001311023332132-2323212030212103-0022332300333333-3003021123320310-0021023320310331-3000001132213121-1331211303103221-2312121021022113", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "schema-cloudfront--aws_configuration_id_selector--ids", "enforcement": "provider-schema", "group": "cloudfront.aws_configuration_id_selector:RequiredObjectAttributes:ids", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:aws_configuration_id_selector", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "aws_configuration_id_selector"], "schema_version": 1, "sections": [{"aliases": ["cloudfront aws configuration id selector ids"], "anchor": "schema-cloudfront--aws_configuration_id_selector--ids", "description": "Add AWS CloudFront distribution ID, e.g. ABCDEFGHI0JKLM.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:aws_configuration_id_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "aws_configuration_id_selector", "ids"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/aws_configuration_id_selector/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of CloudFront distributions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.aws_configuration_id_selector

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- cloudfront.aws_configuration_id_selector

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aws configuration id selector.

Additional upstream details:

List of CloudFront distributions.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ids")}
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
aws_configuration_id_selector {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cloudfront--aws_configuration_id_selector--ids"></a>

### ids property

Type: `["list", "string"]`. Optional.

Add AWS CloudFront distribution ID, e.g. ABCDEFGHI0JKLM.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "32",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[A-Z0-9]+$",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "32",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[A-Z0-9]+$",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
