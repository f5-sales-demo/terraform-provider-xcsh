---
page_title: "cloudflare"
subcategory: ""
description: "cloudflare for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 6993, "body_sha256": "sha256:c3c1ea8f97e8d5b1d7fdd11659f0eab2e8d33b6ee2e994aaf8004b39d7021209", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudflare", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:disable_js_insert", "xcsh-docs:resources:protected_application:properties:cloudflare:disable_mobile_sdk", "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules", "xcsh-docs:resources:protected_application:properties:cloudflare:manual_js_insert", "xcsh-docs:resources:protected_application:properties:cloudflare:mobile_sdk_config", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare", "parent_id": "xcsh-docs:resources:protected_application:reference", "path": "docs/guides/resources--protected_application--properties--cloudflare.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudflare

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
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

- [disable_js_insert](resources--protected_application--properties--cloudflare--disable_js_insert.md): complete subsection reference.

- [disable_mobile_sdk](resources--protected_application--properties--cloudflare--disable_mobile_sdk.md): complete subsection reference.

- [js_insertion_rules](resources--protected_application--properties--cloudflare--js_insertion_rules.md): complete subsection reference.

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

- [manual_js_insert](resources--protected_application--properties--cloudflare--manual_js_insert.md): complete subsection reference.

- [mobile_sdk_config](resources--protected_application--properties--cloudflare--mobile_sdk_config.md): complete subsection reference.

- [protected_endpoints](resources--protected_application--properties--cloudflare--protected_endpoints.md): complete subsection reference.

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

- [trusted_clients](resources--protected_application--properties--cloudflare--trusted_clients.md): complete subsection reference.

## Next pages

- [cloudflare.disable_js_insert](resources--protected_application--properties--cloudflare--disable_js_insert.md)
- [cloudflare.disable_mobile_sdk](resources--protected_application--properties--cloudflare--disable_mobile_sdk.md)
- [cloudflare.js_insertion_rules](resources--protected_application--properties--cloudflare--js_insertion_rules.md)
- [cloudflare.manual_js_insert](resources--protected_application--properties--cloudflare--manual_js_insert.md)
- [cloudflare.mobile_sdk_config](resources--protected_application--properties--cloudflare--mobile_sdk_config.md)
- [cloudflare.protected_endpoints](resources--protected_application--properties--cloudflare--protected_endpoints.md)
- [cloudflare.trusted_clients](resources--protected_application--properties--cloudflare--trusted_clients.md)
- [Property reference](resources--protected_application--reference.md)
- [xcsh_protected_application](../resources/protected_application.md)
