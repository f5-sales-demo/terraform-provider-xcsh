---
page_title: "code_base_integration.bitbucket_server"
subcategory: ""
description: "Configuration parameter for bitbucket server."
xcsh_docs: {"aliases": ["code base integration bitbucket server"], "body_bytes": 3953, "body_sha256": "sha256:7de69a9e7d5c9e5f3d4cc269938341b58cec224e667d7b393cf9a21891130eb5", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration", "path": "documentation/resources/code_base_integration/properties/code_base_integration/bitbucket_server/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1120030323021110-3332221313312113-1230210301103113-2030303033132103-2032321200333101-0031322133320210-3300211013200001-0300233133212230", "registry_path": "docs/guides/resources--code_base_integration--reference--group-001.md", "relationships": [{"anchor": "schema-code_base_integration--bitbucket_server--url", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket_server:RequiredObjectAttributes:url,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server", "type": "requires"}, {"anchor": "schema-code_base_integration--bitbucket_server--username", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket_server:RequiredObjectAttributes:url,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "bitbucket_server"], "schema_version": 1, "sections": [{"aliases": ["code base integration bitbucket server passwd"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket_server.passwd:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "code_base_integration.bitbucket_server.passwd:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd:clear_secret_info", "type": "conflicts"}], "schema_path": ["code_base_integration", "bitbucket_server", "passwd"], "syntax": "block", "type": "object"}, {"aliases": ["code base integration bitbucket server url"], "anchor": "schema-code_base_integration--bitbucket_server--url", "description": "URL or URI reference", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "bitbucket_server", "url"], "syntax": "attribute", "type": "string"}, {"aliases": ["code base integration bitbucket server username"], "anchor": "schema-code_base_integration--bitbucket_server--username", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "bitbucket_server", "username"], "syntax": "attribute", "type": "string"}, {"aliases": ["code base integration bitbucket server verify ssl"], "anchor": "schema-code_base_integration--bitbucket_server--verify_ssl", "description": "Configuration parameter for verify ssl", "document_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:bitbucket_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "bitbucket_server", "verify_ssl"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/bitbucket_server/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration parameter for bitbucket server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.bitbucket_server

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/)
- code_base_integration.bitbucket_server

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bitbucket server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url",
    "username")}
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
bitbucket_server {
  # Configure direct properties listed below.
}
```

## Direct properties

- [passwd](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/): complete subsection reference.

<a id="schema-code_base_integration--bitbucket_server--url"></a>

### url property

Type: `"string"`. Optional.

BitBucket Server URL. URL or URI reference

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
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

<a id="schema-code_base_integration--bitbucket_server--username"></a>

### username property

Type: `"string"`. Optional.

BitBucket Server Username. Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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

<a id="schema-code_base_integration--bitbucket_server--verify_ssl"></a>

### verify_ssl property

Type: `"bool"`. Optional.

Verify SSL. Configuration parameter for verify ssl

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
