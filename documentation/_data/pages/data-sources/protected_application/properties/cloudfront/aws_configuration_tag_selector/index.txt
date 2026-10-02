---
page_title: "cloudfront.aws_configuration_tag_selector"
subcategory: ""
description: "CloudFront distribution tag list."
xcsh_docs: {"aliases": ["cloudfront aws configuration tag selector"], "body_bytes": 2734, "body_sha256": "sha256:7d7bff1cd259a5ffb37b73207357be4c556a60082616fe3cad7da1565d94132d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:aws_configuration_tag_selector", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront", "path": "documentation/data-sources/protected_application/properties/cloudfront/aws_configuration_tag_selector/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2210221302211311-0330010103131230-3211123033330312-3223100021013003-1132131312103322-1110022122230303-0130233332121223-1231100333311331", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "aws_configuration_tag_selector"], "schema_version": 1, "sections": [{"aliases": ["tags"], "anchor": "schema-cloudfront--aws_configuration_tag_selector--tags", "description": "List contains the Cloudfront distribution selection by tags key is a AWS tag name, and the value is regular expression to match.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:aws_configuration_tag_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "aws_configuration_tag_selector", "tags"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/aws_configuration_tag_selector/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "CloudFront distribution tag list.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.aws_configuration_tag_selector

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- cloudfront.aws_configuration_tag_selector

<a id="section"></a>

Type: `"single"`. Computed.

Distribution Tag List. CloudFront distribution tag list.

Upstream description:

CloudFront distribution tag list.

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

## Direct properties

<a id="schema-cloudfront--aws_configuration_tag_selector--tags"></a>

### tags property

Type: `["map", "string"]`. Computed.

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

- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
