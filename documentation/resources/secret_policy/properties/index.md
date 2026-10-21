---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_secret_policy."
xcsh_docs: {"aliases": ["secret policy"], "body_bytes": 11748, "body_sha256": "sha256:9e6ebd12ec1b44ebb5d10fe13aff1b4afda2b8cb91c6b909fdf02d598fc5c4e0", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:secret_policy:properties:rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:secret_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_policy:reference", "parent_id": "xcsh-docs:resources:secret_policy:fundamentals", "path": "documentation/resources/secret_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "secret_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1213313330302120-1100122222322313-2323333222220332-2211122300300103-3232201121133200-0113220013210033-2322130013023011-1132310023231200", "registry_path": "docs/guides/resources--secret_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["allow f5xc"], "anchor": "schema-allow_f5xc", "description": "If allow_f5xc is set to true, it allows relevant F5XC infrastructure services to decrypt the secret encrypted using this policy.", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allow_f5xc"], "syntax": "attribute", "type": "bool"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["decrypt cache timeout", "duration"], "anchor": "schema-decrypt_cache_timeout", "description": "Decrypt_cache_timeout contains the amount of time a decrypted secret is cached in wingman. Value for this parameter is a string ending in the suffix \"s\" (indicating seconds), suffix \"m\" (indicating minutes) or suffix \"h\" (indicating hours)", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["decrypt_cache_timeout"], "syntax": "attribute", "type": "string"}, {"aliases": ["deletion time"], "anchor": "schema-deletion_time", "description": "deletion time", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["deletion_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "id", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["marked for delete"], "anchor": "schema-marked_for_delete", "description": "marked for delete", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["marked_for_delete"], "syntax": "attribute", "type": "bool"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["rule combining algorithm"], "anchor": "schema-rule_combining_algorithm", "description": "rule combining algorithm", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["FIRST_MATCH"], "version": 1}], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_combining_algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules"], "anchor": "section", "description": "rules", "document_id": "xcsh-docs:resources:secret_policy:properties:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["uid"], "anchor": "schema-uid", "description": "uid", "document_id": "xcsh-docs:resources:secret_policy:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_policy/properties/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Property reference for xcsh_secret_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["secret_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_secret_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/)
- Property reference

## Direct properties

<a id="schema-allow_f5xc"></a>

### allow_f5xc property

Type: `"bool"`. Optional, Computed.

Additional upstream details:

If allow\_f5xc is set to true, it allows relevant F5XC infrastructure services to decrypt the secret
encrypted using this policy.

Provider validators and defaults (from schema source):

```go
Default: booldefault.StaticBool(false)
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

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Additional upstream details:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-decrypt_cache_timeout"></a>

### decrypt_cache_timeout property

Type: `"string"`. Optional, Computed.

Additional upstream details:

Decrypt\_cache\_timeout contains the amount of time a decrypted secret is cached in wingman. Value
for this parameter is a string ending in the suffix "s" (indicating seconds), suffix "m" (indicating
minutes) or suffix "h" (indicating hours)

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("30s")
EnumExtractionComplete: true
Validators: []validator.String{stringvalidator.RegexMatches(secretPolicyDuration,
    "Use a positive duration ending in s, m, or h.")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-21T11:00:09+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-deletion_time"></a>

### deletion_time property

Type: `"string"`. Computed.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Additional upstream details:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-21T11:00:09+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-marked_for_delete"></a>

### marked_for_delete property

Type: `"bool"`. Computed.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{validators.NameValidator()}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-21T11:00:09+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{validators.NamespaceValidator()}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-21T11:00:09+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-rule_combining_algorithm"></a>

### rule_combining_algorithm property

Type: `"string"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("FIRST_MATCH")
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["FIRST_MATCH"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{stringvalidator.OneOf("FIRST_MATCH")}
```

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/rules/): complete subsection reference.

<a id="schema-uid"></a>

### uid property

Type: `"string"`. Computed.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_f5xc` | [allow_f5xc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-allow_f5xc) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-annotations) |
| `decrypt_cache_timeout` | [decrypt_cache_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-decrypt_cache_timeout) |
| `deletion_time` | [deletion_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-deletion_time) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-labels) |
| `marked_for_delete` | [marked_for_delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-marked_for_delete) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-namespace) |
| `rule_combining_algorithm` | [rule_combining_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-rule_combining_algorithm) |
| `rules` | [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/rules/#section) |
| `rules.action` | [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/rules/#schema-rules--action) |
| `rules.client_selector` | [rules.client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/rules/#schema-rules--client_selector) |
| `rules.description` | [rules.description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/rules/#schema-rules--description) |
| `rules.name` | [rules.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/rules/#schema-rules--name) |
| `uid` | [uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_policy/properties/#schema-uid) |
