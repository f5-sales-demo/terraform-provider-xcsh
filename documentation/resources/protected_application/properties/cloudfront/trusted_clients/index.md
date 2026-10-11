---
page_title: "cloudfront.trusted_clients"
subcategory: ""
description: "Define your allowlists to skip Bot Defense inference processing."
xcsh_docs: {"aliases": ["cloudfront trusted clients"], "body_bytes": 3138, "body_sha256": "sha256:158580b4e7bb9c87e23ab992c21460b2805666996c5018af9abdb74b7b2fb135", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients:http_header", "xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients:metadata"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront", "path": "documentation/resources/protected_application/properties/cloudfront/trusted_clients/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1331313303213213-1122122223112222-3001222033333102-3201121101201113-1303223000202010-0120230110212212-1113101221303311-3220103301020330", "registry_path": "docs/guides/resources--protected_application--reference--group-004.md", "relationships": [{"anchor": "schema-cloudfront--trusted_clients--ip_prefix", "enforcement": "provider-schema", "group": "cloudfront.trusted_clients:ConflictingListObjectAttributes:http_header,ip_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.trusted_clients:ConflictingListObjectAttributes:http_header,ip_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients:http_header", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "trusted_clients"], "schema_version": 1, "sections": [{"aliases": ["cloudfront trusted clients http header"], "anchor": "section", "description": "Request header name and value pairs.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients:http_header", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.trusted_clients.http_header:RequiredObjectAttributes:headers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients:http_header:headers", "type": "requires"}], "schema_path": ["cloudfront", "trusted_clients", "http_header"], "syntax": "block", "type": "object"}, {"aliases": ["cloudfront trusted clients ip prefix"], "anchor": "schema-cloudfront--trusted_clients--ip_prefix", "description": "Exclusive with IP prefix string.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "trusted_clients", "ip_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront trusted clients metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudfront--trusted_clients--metadata--name", "enforcement": "provider-schema", "group": "cloudfront.trusted_clients.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:trusted_clients:metadata", "type": "requires"}], "schema_path": ["cloudfront", "trusted_clients", "metadata"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/trusted_clients/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Define your allowlists to skip Bot Defense inference processing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.trusted_clients

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- cloudfront.trusted_clients

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/trusted_clients/http_header/): complete subsection reference.

<a id="schema-cloudfront--trusted_clients--ip_prefix"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/trusted_clients/metadata/): complete subsection reference.
