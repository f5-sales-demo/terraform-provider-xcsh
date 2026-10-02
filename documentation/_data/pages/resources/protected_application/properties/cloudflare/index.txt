---
page_title: "cloudflare"
subcategory: ""
description: "Bot Defense policy configuration for Cloudflare."
xcsh_docs: {"aliases": ["cloudflare"], "body_bytes": 8000, "body_sha256": "sha256:feb9b3fee0b84b88ff5dcd50bb9bfe5221b1ad518137edc9c3eb975e5587dea5", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:disable_js_insert", "xcsh-docs:resources:protected_application:properties:cloudflare:disable_mobile_sdk", "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules", "xcsh-docs:resources:protected_application:properties:cloudflare:manual_js_insert", "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare", "parent_id": "xcsh-docs:resources:protected_application:reference", "path": "documentation/resources/protected_application/properties/cloudflare/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111", "registry_path": "docs/guides/resources--protected_application--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare:ConflictingObjectAttributes:disable_js_insert,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare:ConflictingObjectAttributes:disable_js_insert,manual_js_insert", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare:ConflictingObjectAttributes:disable_mobile_sdk,mobile_sdk_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:disable_mobile_sdk", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare:ConflictingObjectAttributes:disable_js_insert,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare:ConflictingObjectAttributes:js_insertion_rules,manual_js_insert", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare:ConflictingObjectAttributes:disable_js_insert,manual_js_insert", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:manual_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare:ConflictingObjectAttributes:js_insertion_rules,manual_js_insert", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:manual_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare:ConflictingObjectAttributes:disable_mobile_sdk,mobile_sdk_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare:RequiredObjectAttributes:protected_endpoints", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare"], "schema_version": 1, "sections": [{"aliases": ["continue mitigation action hdr"], "anchor": "schema-cloudflare--continue_mitigation_action_hdr", "description": "A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "continue_mitigation_action_hdr"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable js insert"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:disable_js_insert", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "disable_js_insert"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable mobile sdk"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:disable_mobile_sdk", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "disable_mobile_sdk"], "syntax": "attribute", "type": "object"}, {"aliases": ["js insertion rules"], "anchor": "section", "description": "This defines custom JavaScript insertion rules for Bot Defense Policy.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:rules", "type": "requires"}], "schema_path": ["cloudflare", "js_insertion_rules"], "syntax": "block", "type": "object"}, {"aliases": ["loglevel"], "anchor": "schema-cloudflare--loglevel", "description": "Select the level of logging desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) - LOG_UNDEFINED: Undefined - LOG_ERROR: Error Log only errors - LOG_WARNING: Warning Log malicious requests - LOG_INFO: Info Log all requests - LOG_DEBUG: Debug Log debugging data.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "loglevel"], "syntax": "attribute", "type": "string"}, {"aliases": ["manual js insert"], "anchor": "section", "description": "Insert JavaScript manually.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:manual_js_insert", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "manual_js_insert"], "syntax": "block", "type": "object"}, {"aliases": ["mobile sdk config"], "anchor": "section", "description": "Mobile SDK configuration.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "mobile_sdk_config"], "syntax": "block", "type": "object"}, {"aliases": ["protected endpoints"], "anchor": "section", "description": "List of protected endpoints (max 128 items)", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeout"], "anchor": "schema-cloudflare--timeout", "description": "The timeout for the inference check, in milliseconds.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["trusted clients"], "anchor": "section", "description": "Define your allowlists to skip Bot Defense inference processing.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudflare", "trusted_clients"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Bot Defense policy configuration for Cloudflare.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- cloudflare

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense policy configuration for Cloudflare.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_endpoints"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "manual_js_insert"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insertion_rules",
    "manual_js_insert")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insertion_rules\",\"manual_js_insert\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
cloudflare {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cloudflare--continue_mitigation_action_hdr"></a>

### continue_mitigation_action_hdr property

Type: `"string"`. Optional.

Case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Upstream description:

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/disable_js_insert/): complete subsection reference.

- [disable_mobile_sdk](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/disable_mobile_sdk/): complete subsection reference.

- [js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/): complete subsection reference.

<a id="schema-cloudflare--loglevel"></a>

### loglevel property

Type: `"string"`. Optional.

\[Enum: LOG\_UNDEFINED|LOG\_ERROR|LOG\_WARNING|LOG\_INFO|LOG\_DEBUG\] Select the level of logging
desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) -
LOG\_UNDEFINED: Undefined - LOG\_ERROR: Error Log only errors - LOG\_WARNING: Warning Log malicious
requests - LOG\_INFO: Info Log all requests - LOG\_DEBUG: Debug Log debugging data. Possible values
are \`LOG\_UNDEFINED\`, \`LOG\_ERROR\`, \`LOG\_WARNING\`, \`LOG\_INFO\`, \`LOG\_DEBUG\`. Defaults to
\`LOG\_UNDEFINED\`.

Upstream description:

Select the level of logging desired. Levels are cumulative (e.g. Debug includes Error, Warning, and
Informational)

&#8203;- LOG\_UNDEFINED: Undefined

&#8203;- LOG\_ERROR: Error

Log only errors &#8203;- LOG\_WARNING: Warning

Log malicious requests &#8203;- LOG\_INFO: Info

Log all requests &#8203;- LOG\_DEBUG: Debug

Log debugging data.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"),
}
```

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

- [manual_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/manual_js_insert/): complete subsection reference.

- [mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/mobile_sdk_config/): complete subsection reference.

- [protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/): complete subsection reference.

<a id="schema-cloudflare--timeout"></a>

### timeout property

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/): complete subsection reference.

## Next pages

- [cloudflare.disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/disable_js_insert/)
- [cloudflare.disable_mobile_sdk](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/disable_mobile_sdk/)
- [cloudflare.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/)
- [cloudflare.manual_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/manual_js_insert/)
- [cloudflare.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/mobile_sdk_config/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/)
- [cloudflare.trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
