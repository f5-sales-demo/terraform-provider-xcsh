---
page_title: "protected_cookies"
subcategory: "Load Balancing"
description: "Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will"
xcsh_docs: {"aliases": ["protected cookies"], "body_bytes": 8664, "body_sha256": "sha256:4ba9029c87be02cfe6bc71144e08ffd0c5828b7be5851595c6e76139f860ae0a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:add_httponly", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:add_secure", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:disable_tampering_protection", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:enable_tampering_protection", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:ignore_httponly", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:ignore_max_age", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:ignore_secure", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:samesite_lax", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:samesite_none", "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:samesite_strict"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/protected_cookies/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protected_cookies"], "schema_version": 1, "sections": [{"aliases": ["add httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:add_httponly", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "add_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["add secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:add_secure", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "add_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable tampering protection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:disable_tampering_protection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "disable_tampering_protection"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable tampering protection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:enable_tampering_protection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "enable_tampering_protection"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:ignore_httponly", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore max age"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:ignore_max_age", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_max_age"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore samesite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:ignore_samesite", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_samesite"], "syntax": "attribute", "type": "object"}, {"aliases": ["ignore secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:ignore_secure", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["max age value"], "anchor": "schema-protected_cookies--max_age_value", "description": "Exclusive with Add max age attribute.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "max_age_value"], "syntax": "attribute", "type": "number"}, {"aliases": ["name"], "anchor": "schema-protected_cookies--name", "description": "Name of the Cookie.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["samesite lax"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:samesite_lax", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "samesite_lax"], "syntax": "attribute", "type": "object"}, {"aliases": ["samesite none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:samesite_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "samesite_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["samesite strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:protected_cookies:samesite_strict", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "samesite_strict"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/protected_cookies/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protected_cookies

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- protected_cookies

<a id="section"></a>

Type: `"list"`. Computed.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite.

Upstream description:

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

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

## Direct properties

- [add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/add_httponly/): complete subsection reference.

- [add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/add_secure/): complete subsection reference.

- [disable_tampering_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/disable_tampering_protection/): complete subsection reference.

- [enable_tampering_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/enable_tampering_protection/): complete subsection reference.

- [ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/ignore_httponly/): complete subsection reference.

- [ignore_max_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/ignore_max_age/): complete subsection reference.

- [ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/ignore_samesite/): complete subsection reference.

- [ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/ignore_secure/): complete subsection reference.

<a id="schema-protected_cookies--max_age_value"></a>

### max_age_value property

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

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

Type: `"string"`. Computed.

Cookie Name. Name of the Cookie.

Upstream description:

Name of the Cookie.

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

- [samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/samesite_lax/): complete subsection reference.

- [samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/samesite_none/): complete subsection reference.

- [samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/samesite_strict/): complete subsection reference.

## Next pages

- [protected_cookies.add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/add_httponly/)
- [protected_cookies.add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/add_secure/)
- [protected_cookies.disable_tampering_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/disable_tampering_protection/)
- [protected_cookies.enable_tampering_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/enable_tampering_protection/)
- [protected_cookies.ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/ignore_httponly/)
- [protected_cookies.ignore_max_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/ignore_max_age/)
- [protected_cookies.ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/ignore_samesite/)
- [protected_cookies.ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/ignore_secure/)
- [protected_cookies.samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/samesite_lax/)
- [protected_cookies.samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/samesite_none/)
- [protected_cookies.samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/protected_cookies/samesite_strict/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
