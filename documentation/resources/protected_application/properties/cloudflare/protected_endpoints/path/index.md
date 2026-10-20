---
page_title: "cloudflare.protected_endpoints.path"
subcategory: ""
description: "URI Path"
xcsh_docs: {"aliases": ["cloudflare protected endpoints path"], "body_bytes": 3093, "body_sha256": "sha256:eee06f120b2f936d27e80e2ed8756b5f597dade045ab6c60e99aa41e8d8fdb65", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:path", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "path": "documentation/resources/protected_application/properties/cloudflare/protected_endpoints/path/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1001300212102331-3212330102111323-3100201020003211-0232102000201230-2211133203031111-3132133032303233-3312123120323022-0303222121321212", "registry_path": "docs/guides/resources--protected_application--reference--group-001.md", "relationships": [{"anchor": "schema-cloudflare--protected_endpoints--path--path", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.path:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:path", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "path"], "schema_version": 1, "sections": [{"aliases": ["cloudflare protected endpoints path caseinsensitive"], "anchor": "schema-cloudflare--protected_endpoints--path--caseinsensitive", "description": "Should path be searched case insensitive;", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:path", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "path", "caseinsensitive"], "syntax": "attribute", "type": "bool"}, {"aliases": ["cloudflare protected endpoints path path"], "anchor": "schema-cloudflare--protected_endpoints--path--path", "description": "URI Path", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:path", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "path", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/protected_endpoints/path/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "URI Path", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.path

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/)
- cloudflare.protected_endpoints.path

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Path. URI Path

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
path {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cloudflare--protected_endpoints--path--caseinsensitive"></a>

### caseinsensitive property

Type: `"bool"`. Optional.

Should path be searched case insensitive;.

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

<a id="schema-cloudflare--protected_endpoints--path--path"></a>

### path property

Type: `"string"`. Optional.

Path. URI Path

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-20T00:10:01+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  }
}
```
