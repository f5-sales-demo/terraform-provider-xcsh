---
page_title: "cloudfront.aws_configuration_tag_selector"
subcategory: ""
description: "CloudFront distribution tag list."
xcsh_docs: {"aliases": ["cloudfront aws configuration tag selector"], "body_bytes": 3001, "body_sha256": "sha256:f9512dbb681522f0f83b7d81abad48775b9e570011d6e568a2d3b348f44263d7", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:aws_configuration_tag_selector", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront", "path": "documentation/resources/protected_application/properties/cloudfront/aws_configuration_tag_selector/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2301003033010132-2011002322022103-1032023300113333-3203030110102133-1233003321210103-1010302203112203-1332321303311222-0210232022202322", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "schema-cloudfront--aws_configuration_tag_selector--tags", "enforcement": "provider-schema", "group": "cloudfront.aws_configuration_tag_selector:RequiredObjectAttributes:tags", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:aws_configuration_tag_selector", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "aws_configuration_tag_selector"], "schema_version": 1, "sections": [{"aliases": ["tags"], "anchor": "schema-cloudfront--aws_configuration_tag_selector--tags", "description": "List contains the Cloudfront distribution selection by tags key is a AWS tag name, and the value is regular expression to match.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:aws_configuration_tag_selector", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "aws_configuration_tag_selector", "tags"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/aws_configuration_tag_selector/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "CloudFront distribution tag list.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.aws_configuration_tag_selector

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- cloudfront.aws_configuration_tag_selector

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Distribution Tag List. CloudFront distribution tag list.

Upstream description:

CloudFront distribution tag list.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tags")}
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
aws_configuration_tag_selector {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cloudfront--aws_configuration_tag_selector--tags"></a>

### tags property

Type: `["map", "string"]`. Optional.

List contains the Cloudfront distribution selection by tags key is a AWS tag name, and the value is
regular expression to match.

Upstream description:

List contains the Cloudfront distribution selection by tags key is a AWS tag name, and the value is
regular expression to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.min_pairs": "1",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.regex": "true",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.min_pairs": "1",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.regex": "true",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
