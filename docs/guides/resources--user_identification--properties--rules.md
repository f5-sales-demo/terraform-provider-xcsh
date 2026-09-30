---
page_title: "rules"
subcategory: ""
description: "rules for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 22392, "body_sha256": "sha256:df36f2be295b7b28766ea07dc5d011f6a66c79fd6bfbd3fe457528865aeb5a5f", "canonical_id": "xcsh-docs:resources:user_identification:properties:rules", "child_ids": ["xcsh-docs:resources:user_identification:properties:rules:client_asn", "xcsh-docs:resources:user_identification:properties:rules:client_city", "xcsh-docs:resources:user_identification:properties:rules:client_country", "xcsh-docs:resources:user_identification:properties:rules:client_ip", "xcsh-docs:resources:user_identification:properties:rules:client_region", "xcsh-docs:resources:user_identification:properties:rules:ip_and_ja4_tls_fingerprint", "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "xcsh-docs:resources:user_identification:properties:rules:none", "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint"], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules", "parent_id": "xcsh-docs:resources:user_identification:reference", "path": "docs/guides/resources--user_identification--properties--rules.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md)
- [Property reference](resources--user_identification--reference.md)
- rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules that are evaluated sequentially against the input fields extracted from an API
request in order to determine a user identifier. Evaluation of the rules is terminated once a user
identifier has been extracted.

Upstream description:

An ordered list of rules that are evaluated sequentially against the input fields extracted from an
API request in order to determine a user identifier. Evaluation of the rules is terminated once a
user identifier has been extracted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("client_asn",
    "client_city"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_country"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_asn",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "none"),
  validators.ConflictingListObjectAttributes("client_asn",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_asn",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_country"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_city",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "none"),
  validators.ConflictingListObjectAttributes("client_city",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_city",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_country",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_country",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "none"),
  validators.ConflictingListObjectAttributes("client_country",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_country",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_ip",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "none"),
  validators.ConflictingListObjectAttributes("client_ip",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_ip",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "none"),
  validators.ConflictingListObjectAttributes("client_region",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_region",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "none"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "none"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "none"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("none",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("none",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("query_param_key",
    "tls_fingerprint")}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_asn](resources--user_identification--properties--rules--client_asn.md): complete subsection reference.

- [client_city](resources--user_identification--properties--rules--client_city.md): complete subsection reference.

- [client_country](resources--user_identification--properties--rules--client_country.md): complete subsection reference.

- [client_ip](resources--user_identification--properties--rules--client_ip.md): complete subsection reference.

- [client_region](resources--user_identification--properties--rules--client_region.md): complete subsection reference.

<a id="schema-rules--cookie_name"></a>

### cookie_name property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user identifier.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user identifier.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [ip_and_ja4_tls_fingerprint](resources--user_identification--properties--rules--ip_and_ja4_tls_fingerprint.md): complete subsection reference.

- [ip_and_tls_fingerprint](resources--user_identification--properties--rules--ip_and_tls_fingerprint.md): complete subsection reference.

- [ja4_tls_fingerprint](resources--user_identification--properties--rules--ja4_tls_fingerprint.md): complete subsection reference.

<a id="schema-rules--jwt_claim_name"></a>

### jwt_claim_name property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [none](resources--user_identification--properties--rules--none.md): complete subsection reference.

<a id="schema-rules--query_param_key"></a>

### query_param_key property

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user identifier.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [tls_fingerprint](resources--user_identification--properties--rules--tls_fingerprint.md): complete subsection reference.

## Next pages

- [rules.client_asn](resources--user_identification--properties--rules--client_asn.md)
- [rules.client_city](resources--user_identification--properties--rules--client_city.md)
- [rules.client_country](resources--user_identification--properties--rules--client_country.md)
- [rules.client_ip](resources--user_identification--properties--rules--client_ip.md)
- [rules.client_region](resources--user_identification--properties--rules--client_region.md)
- [rules.ip_and_ja4_tls_fingerprint](resources--user_identification--properties--rules--ip_and_ja4_tls_fingerprint.md)
- [rules.ip_and_tls_fingerprint](resources--user_identification--properties--rules--ip_and_tls_fingerprint.md)
- [rules.ja4_tls_fingerprint](resources--user_identification--properties--rules--ja4_tls_fingerprint.md)
- [rules.none](resources--user_identification--properties--rules--none.md)
- [rules.tls_fingerprint](resources--user_identification--properties--rules--tls_fingerprint.md)
- [Property reference](resources--user_identification--reference.md)
- [xcsh_user_identification](../resources/user_identification.md)
