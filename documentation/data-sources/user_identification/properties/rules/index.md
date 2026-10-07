---
page_title: "rules"
subcategory: ""
description: "An ordered list of rules that are evaluated sequentially against the input fields extracted from an API request in order to determine a user identifier. Evaluation of the rules is terminated once a user identifier has been extracted."
xcsh_docs: {"aliases": ["rules"], "body_bytes": 9944, "body_sha256": "sha256:07d2d3b6cbea93fea3232dde760bb9b79c6c473929bbd78e9eb45e3ae6a0003e", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:user_identification:properties:rules:client_asn", "xcsh-docs:data-sources:user_identification:properties:rules:client_city", "xcsh-docs:data-sources:user_identification:properties:rules:client_country", "xcsh-docs:data-sources:user_identification:properties:rules:client_ip", "xcsh-docs:data-sources:user_identification:properties:rules:client_region", "xcsh-docs:data-sources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "xcsh-docs:data-sources:user_identification:properties:rules:ip_and_tls_fingerprint", "xcsh-docs:data-sources:user_identification:properties:rules:ja4_tls_fingerprint", "xcsh-docs:data-sources:user_identification:properties:rules:none", "xcsh-docs:data-sources:user_identification:properties:rules:tls_fingerprint"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:user_identification:properties:rules", "parent_id": "xcsh-docs:data-sources:user_identification:reference", "path": "documentation/data-sources/user_identification/properties/rules/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0023320101101103-1310232200212121-0320003011113033-0312100132213231-1113210010123023-0103110220112213-2032223203203200-1013303102232233", "registry_path": "docs/guides/data-sources--user_identification--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules"], "schema_version": 1, "sections": [{"aliases": ["rules client asn"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:client_asn", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_asn"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules client city"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:client_city", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_city"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules client country"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:client_country", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_country"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules client ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:client_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules client region"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:client_region", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "client_region"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules cookie name"], "anchor": "schema-rules--cookie_name", "description": "Exclusive with Use the HTTP cookie value for the given name as user identifier.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "cookie_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules http header name"], "anchor": "schema-rules--http_header_name", "description": "Exclusive with Use the HTTP header value for the given name as user identifier.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "http_header_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules ip and http header name"], "anchor": "schema-rules--ip_and_http_header_name", "description": "Exclusive with Name of HTTP header from which the value should be extracted.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ip_and_http_header_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules ip and ja4 tls fingerprint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ip_and_ja4_tls_fingerprint"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules ip and tls fingerprint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:ip_and_tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ip_and_tls_fingerprint"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules ja4 tls fingerprint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:ja4_tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "ja4_tls_fingerprint"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules jwt claim name"], "anchor": "schema-rules--jwt_claim_name", "description": "Exclusive with Use the JWT claim value as user identifier.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "jwt_claim_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:none", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "none"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules query param key"], "anchor": "schema-rules--query_param_key", "description": "Exclusive with Use the query parameter value for the given key as user identifier.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "query_param_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules tls fingerprint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:user_identification:properties:rules:tls_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "tls_fingerprint"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/user_identification/properties/rules/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "An ordered list of rules that are evaluated sequentially against the input fields extracted from an API request in order to determine a user identifier. Evaluation of the rules is terminated once a user identifier has been extracted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["user_identificationCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/)
- rules

<a id="section"></a>

Type: `"list"`. Computed.

An ordered list of rules that are evaluated sequentially against the input fields extracted from an
API request in order to determine a user identifier. Evaluation of the rules is terminated once a
user identifier has been extracted.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [client_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/client_asn/): complete subsection reference.

- [client_city](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/client_city/): complete subsection reference.

- [client_country](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/client_country/): complete subsection reference.

- [client_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/client_ip/): complete subsection reference.

- [client_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/client_region/): complete subsection reference.

<a id="schema-rules--cookie_name"></a>

### cookie_name property

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user identifier.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="schema-rules--http_header_name"></a>

### http_header_name property

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user identifier.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="schema-rules--ip_and_http_header_name"></a>

### ip_and_http_header_name property

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [ip_and_ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/ip_and_ja4_tls_fingerprint/): complete subsection reference.

- [ip_and_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/ip_and_tls_fingerprint/): complete subsection reference.

- [ja4_tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/ja4_tls_fingerprint/): complete subsection reference.

<a id="schema-rules--jwt_claim_name"></a>

### jwt_claim_name property

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/none/): complete subsection reference.

<a id="schema-rules--query_param_key"></a>

### query_param_key property

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user identifier.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tls_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/tls_fingerprint/): complete subsection reference.
