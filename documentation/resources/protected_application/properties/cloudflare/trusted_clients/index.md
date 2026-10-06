---
page_title: "cloudflare.trusted_clients"
subcategory: ""
description: "Define your allowlists to skip Bot Defense inference processing."
xcsh_docs: {"aliases": ["cloudflare trusted clients"], "body_bytes": 3138, "body_sha256": "sha256:4f1387820a1043b5dbceb5c33e61db21ec2019c5ffc2c3031671b0dd7e81cfcc", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header", "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:metadata"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare", "path": "documentation/resources/protected_application/properties/cloudflare/trusted_clients/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "schema-cloudflare--trusted_clients--ip_prefix", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients:ConflictingListObjectAttributes:http_header,ip_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients:ConflictingListObjectAttributes:http_header,ip_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "trusted_clients"], "schema_version": 1, "sections": [{"aliases": ["cloudflare trusted clients http header"], "anchor": "section", "description": "Request header name and value pairs.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients.http_header:RequiredObjectAttributes:headers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "type": "requires"}], "schema_path": ["cloudflare", "trusted_clients", "http_header"], "syntax": "block", "type": "object"}, {"aliases": ["cloudflare trusted clients ip prefix"], "anchor": "schema-cloudflare--trusted_clients--ip_prefix", "description": "Exclusive with IP prefix string.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "trusted_clients", "ip_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudflare trusted clients metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudflare--trusted_clients--metadata--name", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:metadata", "type": "requires"}], "schema_path": ["cloudflare", "trusted_clients", "metadata"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/trusted_clients/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Define your allowlists to skip Bot Defense inference processing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.trusted_clients

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- cloudflare.trusted_clients

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Define your allowlists to skip Bot Defense inference processing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
trusted_clients {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/http_header/): complete subsection reference.

<a id="schema-cloudflare--trusted_clients--ip_prefix"></a>

### ip_prefix property

Type: `"string"`. Optional.

Exclusive with \[http\_header\] IP prefix string.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/metadata/): complete subsection reference.
