---
page_title: "cloudflare.protected_endpoints"
subcategory: ""
description: "List of protected endpoints (max 128 items)"
xcsh_docs: {"aliases": ["cloudflare protected endpoints"], "body_bytes": 6653, "body_sha256": "sha256:5febf072ce73d414083ee2c5296e939e2dc3bace2fb30d6111065e5e2e175a31", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:any_domain", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:domain", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:metadata", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:path", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare", "path": "documentation/resources/protected_application/properties/cloudflare/protected_endpoints/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200", "registry_path": "docs/guides/resources--protected_application--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints:ConflictingListObjectAttributes:mobile_client,web_client", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints:ConflictingListObjectAttributes:mobile_client,web_mobile_client", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints:ConflictingListObjectAttributes:mobile_client,web_client", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints:ConflictingListObjectAttributes:web_client,web_mobile_client", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints:ConflictingListObjectAttributes:mobile_client,web_mobile_client", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints:ConflictingListObjectAttributes:web_client,web_mobile_client", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client", "type": "conflicts"}, {"anchor": "schema-cloudflare--protected_endpoints--http_methods", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints:RequiredListObjectAttributes:http_methods", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "protected_endpoints"], "schema_version": 1, "sections": [{"aliases": ["cloudflare protected endpoints any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:any_domain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudflare protected endpoints domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:domain", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudflare--protected_endpoints--domain--exact_value", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--protected_endpoints--domain--exact_value", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--protected_endpoints--domain--regex_value", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.domain:ConflictingObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--protected_endpoints--domain--regex_value", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--protected_endpoints--domain--suffix_value", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.domain:ConflictingObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:domain", "type": "conflicts"}, {"anchor": "schema-cloudflare--protected_endpoints--domain--suffix_value", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.domain:ConflictingObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:domain", "type": "conflicts"}], "schema_path": ["cloudflare", "protected_endpoints", "domain"], "syntax": "block", "type": "object"}, {"aliases": ["cloudflare protected endpoints http methods"], "anchor": "schema-cloudflare--protected_endpoints--http_methods", "description": "List of HTTP methods.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "http_methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["cloudflare protected endpoints metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudflare--protected_endpoints--metadata--name", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:metadata", "type": "requires"}], "schema_path": ["cloudflare", "protected_endpoints", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["cloudflare protected endpoints mobile client"], "anchor": "section", "description": "Mobile client configuration OPTIONS.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.mobile_client:ConflictingObjectAttributes:block,continue", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.mobile_client:ConflictingObjectAttributes:block,continue", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue", "type": "conflicts"}], "schema_path": ["cloudflare", "protected_endpoints", "mobile_client"], "syntax": "block", "type": "object"}, {"aliases": ["cloudflare protected endpoints path"], "anchor": "section", "description": "URI Path", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:path", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudflare--protected_endpoints--path--path", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.path:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:path", "type": "requires"}], "schema_path": ["cloudflare", "protected_endpoints", "path"], "syntax": "block", "type": "object"}, {"aliases": ["cloudflare protected endpoints query"], "anchor": "schema-cloudflare--protected_endpoints--query", "description": "Enter a regular expression to match your query parameters of interest.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "query"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudflare protected endpoints web client"], "anchor": "section", "description": "Web client configuration OPTIONS.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_client:ConflictingObjectAttributes:block,continue", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_client:ConflictingObjectAttributes:block,redirect", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_client:ConflictingObjectAttributes:block,continue", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_client:ConflictingObjectAttributes:continue,redirect", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_client:ConflictingObjectAttributes:block,redirect", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:redirect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_client:ConflictingObjectAttributes:continue,redirect", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:redirect", "type": "conflicts"}], "schema_path": ["cloudflare", "protected_endpoints", "web_client"], "syntax": "block", "type": "object"}, {"aliases": ["cloudflare protected endpoints web mobile client"], "anchor": "section", "description": "Web and Mobile client configuration OPTIONS.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_mobile,continue_mobile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_mobile", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_web,continue_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_web,redirect_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_mobile,continue_mobile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_web,continue_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:continue_web,redirect_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_web,redirect_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:redirect_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:continue_web,redirect_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:redirect_web", "type": "conflicts"}], "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/protected_endpoints/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of protected endpoints (max 128 items)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- cloudflare.protected_endpoints

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints (max 128 items).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_client"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_mobile_client"),
  validators.ConflictingListObjectAttributes("web_client",
    "web_mobile_client")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_endpoints {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/domain/): complete subsection reference.

<a id="schema-cloudflare--protected_endpoints--http_methods"></a>

### http_methods property

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/metadata/): complete subsection reference.

- [mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/path/): complete subsection reference.

<a id="schema-cloudflare--protected_endpoints--query"></a>

### query property

Type: `"string"`. Optional.

Enter a regular expression to match your query parameters of interest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

- [web_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_client/): complete subsection reference.

- [web_mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/): complete subsection reference.
