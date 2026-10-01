---
page_title: "protected_cookies"
subcategory: "Load Balancing"
description: "protected_cookies for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 8703, "body_sha256": "sha256:6e40931819e678d8e28be8ab1bc92a4ad68221daffe186d08b18f694c08c2ec0", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:add_httponly", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:add_secure", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:disable_tampering_protection", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:enable_tampering_protection", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_httponly", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_max_age", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_secure", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_lax", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_none", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_strict"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--protected_cookies.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protected_cookies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/protected_cookies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protected_cookies for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protected_cookies

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- protected_cookies

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite.

Upstream description:

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("disable_tampering_protection",
    "enable_tampering_protection"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_cookies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [add_httponly](resources--http_loadbalancer--properties--protected_cookies--add_httponly.md): complete subsection reference.

- [add_secure](resources--http_loadbalancer--properties--protected_cookies--add_secure.md): complete subsection reference.

- [disable_tampering_protection](resources--http_loadbalancer--properties--protected_cookies--disable_tampering_protection.md): complete subsection reference.

- [enable_tampering_protection](resources--http_loadbalancer--properties--protected_cookies--enable_tampering_protection.md): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--properties--protected_cookies--ignore_httponly.md): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--properties--protected_cookies--ignore_max_age.md): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--properties--protected_cookies--ignore_samesite.md): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--properties--protected_cookies--ignore_secure.md): complete subsection reference.

<a id="schema-protected_cookies--max_age_value"></a>

### max_age_value property

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="schema-protected_cookies--name"></a>

### name property

Type: `"string"`. Optional.

Cookie Name. Name of the Cookie.

Upstream description:

Name of the Cookie.

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
  "x-f5xc-constraints": {
    "category": "discovery",
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](resources--http_loadbalancer--properties--protected_cookies--samesite_lax.md): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--properties--protected_cookies--samesite_none.md): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--properties--protected_cookies--samesite_strict.md): complete subsection reference.

## Next pages

- [protected_cookies.add_httponly](resources--http_loadbalancer--properties--protected_cookies--add_httponly.md)
- [protected_cookies.add_secure](resources--http_loadbalancer--properties--protected_cookies--add_secure.md)
- [protected_cookies.disable_tampering_protection](resources--http_loadbalancer--properties--protected_cookies--disable_tampering_protection.md)
- [protected_cookies.enable_tampering_protection](resources--http_loadbalancer--properties--protected_cookies--enable_tampering_protection.md)
- [protected_cookies.ignore_httponly](resources--http_loadbalancer--properties--protected_cookies--ignore_httponly.md)
- [protected_cookies.ignore_max_age](resources--http_loadbalancer--properties--protected_cookies--ignore_max_age.md)
- [protected_cookies.ignore_samesite](resources--http_loadbalancer--properties--protected_cookies--ignore_samesite.md)
- [protected_cookies.ignore_secure](resources--http_loadbalancer--properties--protected_cookies--ignore_secure.md)
- [protected_cookies.samesite_lax](resources--http_loadbalancer--properties--protected_cookies--samesite_lax.md)
- [protected_cookies.samesite_none](resources--http_loadbalancer--properties--protected_cookies--samesite_none.md)
- [protected_cookies.samesite_strict](resources--http_loadbalancer--properties--protected_cookies--samesite_strict.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
