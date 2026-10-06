---
page_title: "protected_cookies"
subcategory: "Load Balancing"
description: "Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will"
xcsh_docs: {"aliases": ["protected cookies"], "body_bytes": 7472, "body_sha256": "sha256:0389e4356a3927d2264629c7cdf722369e2123fede0e3dce7f34b19712498c89", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:add_httponly", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:add_secure", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:disable_tampering_protection", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:enable_tampering_protection", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_httponly", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_max_age", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_secure", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_lax", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_none", "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_strict"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/protected_cookies/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [{"anchor": "schema-protected_cookies--max_age_value", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:ignore_max_age,max_age_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:add_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:add_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:disable_tampering_protection,enable_tampering_protection", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:disable_tampering_protection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:disable_tampering_protection,enable_tampering_protection", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:enable_tampering_protection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:ignore_max_age,max_age_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_max_age", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protected_cookies:ConflictingListObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_strict", "type": "conflicts"}, {"anchor": "schema-protected_cookies--name", "enforcement": "provider-schema", "group": "protected_cookies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["protected_cookies"], "schema_version": 1, "sections": [{"aliases": ["protected cookies add httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:add_httponly", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "add_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies add secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:add_secure", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "add_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies disable tampering protection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:disable_tampering_protection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "disable_tampering_protection"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies enable tampering protection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:enable_tampering_protection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "enable_tampering_protection"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies ignore httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_httponly", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies ignore max age"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_max_age", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_max_age"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies ignore samesite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_samesite"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies ignore secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:ignore_secure", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies max age value"], "anchor": "schema-protected_cookies--max_age_value", "description": "Exclusive with Add max age attribute.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "max_age_value"], "syntax": "attribute", "type": "number"}, {"aliases": ["protected cookies name"], "anchor": "schema-protected_cookies--name", "description": "Name of the Cookie.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["protected cookies samesite lax"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_lax", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "samesite_lax"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies samesite none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "samesite_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies samesite strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:protected_cookies:samesite_strict", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "samesite_strict"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/protected_cookies/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protected_cookies

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- protected_cookies

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/add_httponly/): complete subsection reference.

- [add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/add_secure/): complete subsection reference.

- [disable_tampering_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/disable_tampering_protection/): complete subsection reference.

- [enable_tampering_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/enable_tampering_protection/): complete subsection reference.

- [ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/ignore_httponly/): complete subsection reference.

- [ignore_max_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/ignore_max_age/): complete subsection reference.

- [ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/ignore_samesite/): complete subsection reference.

- [ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/ignore_secure/): complete subsection reference.

<a id="schema-protected_cookies--max_age_value"></a>

### max_age_value property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/samesite_lax/): complete subsection reference.

- [samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/samesite_none/): complete subsection reference.

- [samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/protected_cookies/samesite_strict/): complete subsection reference.
