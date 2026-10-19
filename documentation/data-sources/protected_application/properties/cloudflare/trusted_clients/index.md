---
page_title: "cloudflare.trusted_clients"
subcategory: ""
description: "Define your allowlists to skip Bot Defense inference processing."
xcsh_docs: {"aliases": ["cloudflare trusted clients"], "body_bytes": 2628, "body_sha256": "sha256:be2c2f9cc7e94c648c3f9d65d0181dfa7e608e774ac580fbc930ffe325fb1f63", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients:http_header", "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients:metadata"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare", "path": "documentation/data-sources/protected_application/properties/cloudflare/trusted_clients/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "trusted_clients"], "schema_version": 1, "sections": [{"aliases": ["cloudflare trusted clients http header"], "anchor": "section", "description": "Request header name and value pairs.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients:http_header", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "trusted_clients", "http_header"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudflare trusted clients ip prefix"], "anchor": "schema-cloudflare--trusted_clients--ip_prefix", "description": "Exclusive with IP prefix string.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "trusted_clients", "ip_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudflare trusted clients metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "trusted_clients", "metadata"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/trusted_clients/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Define your allowlists to skip Bot Defense inference processing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.trusted_clients

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/)
- cloudflare.trusted_clients

<a id="section"></a>

Type: `"list"`. Computed.

Define your allowlists to skip Bot Defense inference processing.

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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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

## Direct properties

- [http_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/trusted_clients/http_header/): complete subsection reference.

<a id="schema-cloudflare--trusted_clients--ip_prefix"></a>

### ip_prefix property

Type: `"string"`. Computed.

Exclusive with \[http\_header\] IP prefix string.

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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/trusted_clients/metadata/): complete subsection reference.
