---
page_title: "code_base_integration.bitbucket_server"
subcategory: ""
description: "Configuration parameter for bitbucket server."
xcsh_docs: {"aliases": ["code base integration bitbucket server"], "body_bytes": 3381, "body_sha256": "sha256:26d214a3bbb117e83591b0246c3712a9e87cace82984699581b629e3c5e30dab", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server", "parent_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration", "path": "documentation/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2233300220000300-2201300211120013-0002332300233220-3012101120033023-0213023212111120-0302220111313222-3302330000303311-2020232020102220", "registry_path": "docs/guides/data-sources--code_base_integration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "bitbucket_server"], "schema_version": 1, "sections": [{"aliases": ["code base integration bitbucket server passwd"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server:passwd", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "bitbucket_server", "passwd"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration bitbucket server url"], "anchor": "schema-code_base_integration--bitbucket_server--url", "description": "URL or URI reference", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "bitbucket_server", "url"], "syntax": "attribute", "type": "string"}, {"aliases": ["code base integration bitbucket server username"], "anchor": "schema-code_base_integration--bitbucket_server--username", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "bitbucket_server", "username"], "syntax": "attribute", "type": "string"}, {"aliases": ["code base integration bitbucket server verify ssl"], "anchor": "schema-code_base_integration--bitbucket_server--verify_ssl", "description": "Configuration parameter for verify ssl", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:bitbucket_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["code_base_integration", "bitbucket_server", "verify_ssl"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration parameter for bitbucket server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.bitbucket_server

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/)
- code_base_integration.bitbucket_server

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for bitbucket server.

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

- [passwd](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/bitbucket_server/passwd/): complete subsection reference.

<a id="schema-code_base_integration--bitbucket_server--url"></a>

### url property

Type: `"string"`. Computed.

BitBucket Server URL. URL or URI reference

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

Type: `"string"`. Computed.

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

Type: `"bool"`. Computed.

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
