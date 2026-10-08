---
page_title: "cloudfront"
subcategory: ""
description: "Bot Defense policy configuration for AWS Cloudfront."
xcsh_docs: {"aliases": ["cloudfront"], "body_bytes": 6746, "body_sha256": "sha256:7da7e15ce3e873253150e94e1b266e5687cb5151d2d8074be55c8368ac458bb2", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:aws_configuration_id_selector", "xcsh-docs:data-sources:protected_application:properties:cloudfront:aws_configuration_tag_selector", "xcsh-docs:data-sources:protected_application:properties:cloudfront:disable_aws_configuration", "xcsh-docs:data-sources:protected_application:properties:cloudfront:disable_js_insert", "xcsh-docs:data-sources:protected_application:properties:cloudfront:disable_mobile_sdk", "xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules", "xcsh-docs:data-sources:protected_application:properties:cloudfront:manual_js_insert", "xcsh-docs:data-sources:protected_application:properties:cloudfront:mobile_sdk_config", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront", "parent_id": "xcsh-docs:data-sources:protected_application:reference", "path": "documentation/data-sources/protected_application/properties/cloudfront/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront"], "schema_version": 1, "sections": [{"aliases": ["cloudfront aws configuration id selector"], "anchor": "section", "description": "List of CloudFront distributions.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:aws_configuration_id_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "aws_configuration_id_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront aws configuration tag selector"], "anchor": "section", "description": "CloudFront distribution tag list.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:aws_configuration_tag_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "aws_configuration_tag_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront continue mitigation action hdr"], "anchor": "schema-cloudfront--continue_mitigation_action_hdr", "description": "A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "continue_mitigation_action_hdr"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront data sample"], "anchor": "schema-cloudfront--data_sample", "description": "Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576 == 1 MiByte)", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "data_sample"], "syntax": "attribute", "type": "number"}, {"aliases": ["cloudfront disable aws configuration"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:disable_aws_configuration", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "disable_aws_configuration"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront disable js insert"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:disable_js_insert", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "disable_js_insert"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront disable mobile sdk"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:disable_mobile_sdk", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "disable_mobile_sdk"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront js insertion rules"], "anchor": "section", "description": "This defines custom JavaScript insertion rules for Bot Defense Policy.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront loglevel"], "anchor": "schema-cloudfront--loglevel", "description": "Select the level of logging desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) - LOG_UNDEFINED: Undefined - LOG_ERROR: Error Log only errors - LOG_WARNING: Warning Log malicious requests - LOG_INFO: Info Log all requests - LOG_DEBUG: Debug Log debugging data.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "loglevel"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront manual js insert"], "anchor": "section", "description": "Insert JavaScript manually.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:manual_js_insert", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "manual_js_insert"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront mobile sdk config"], "anchor": "section", "description": "Mobile SDK configuration.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:mobile_sdk_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "mobile_sdk_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints"], "anchor": "section", "description": "List of protected endpoints (max 128 items)", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront timeout", "duration"], "anchor": "schema-cloudfront--timeout", "description": "The timeout for the inference check, in milliseconds.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["cloudfront trusted clients"], "anchor": "section", "description": "Define your allowlists to skip Bot Defense inference processing.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudfront", "trusted_clients"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Bot Defense policy configuration for AWS Cloudfront.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- cloudfront

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense policy configuration for AWS Cloudfront.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-aws_configuration_type_choice": "[\"aws_configuration_id_selector\",\"aws_configuration_tag_selector\",\"disable_aws_configuration\"]",
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insertion_rules\",\"manual_js_insert\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

## Direct properties

- [aws_configuration_id_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/aws_configuration_id_selector/): complete subsection reference.

- [aws_configuration_tag_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/aws_configuration_tag_selector/): complete subsection reference.

<a id="schema-cloudfront--continue_mitigation_action_hdr"></a>

### continue_mitigation_action_hdr property

Type: `"string"`. Computed.

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="schema-cloudfront--data_sample"></a>

### data_sample property

Type: `"number"`. Computed.

Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576
== 1 MiByte).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1048576,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1048576"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1048576"
  }
}
```

- [disable_aws_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/disable_aws_configuration/): complete subsection reference.

- [disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/disable_js_insert/): complete subsection reference.

- [disable_mobile_sdk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/disable_mobile_sdk/): complete subsection reference.

- [js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/js_insertion_rules/): complete subsection reference.

<a id="schema-cloudfront--loglevel"></a>

### loglevel property

Type: `"string"`. Computed.

\[Enum: LOG\_UNDEFINED|LOG\_ERROR|LOG\_WARNING|LOG\_INFO|LOG\_DEBUG\] Select the level of logging
desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) -
LOG\_UNDEFINED: Undefined - LOG\_ERROR: Error Log only errors - LOG\_WARNING: Warning Log malicious
requests - LOG\_INFO: Info Log all requests - LOG\_DEBUG: Debug Log debugging data. Possible values
are \`LOG\_UNDEFINED\`, \`LOG\_ERROR\`, \`LOG\_WARNING\`, \`LOG\_INFO\`, \`LOG\_DEBUG\`. Defaults to
\`LOG\_UNDEFINED\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "LOG_UNDEFINED",
  "enum": [
    "LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [manual_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/manual_js_insert/): complete subsection reference.

- [mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/): complete subsection reference.

- [protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/): complete subsection reference.

<a id="schema-cloudfront--timeout"></a>

### timeout property

Type: `"number"`. Computed.

The timeout for the inference check, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/trusted_clients/): complete subsection reference.
