---
page_title: "protected_cookies"
subcategory: "Load Balancing"
description: "Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will"
xcsh_docs: {"aliases": ["protected cookies"], "body_bytes": 6072, "body_sha256": "sha256:ffe1737d8c39a0764d93c37580097c56ee61c46f3253f78031b3a2fbd51bc4b5", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:add_httponly", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:add_secure", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:disable_tampering_protection", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:enable_tampering_protection", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:ignore_httponly", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:ignore_max_age", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:ignore_samesite", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:ignore_secure", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:samesite_lax", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:samesite_none", "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:samesite_strict"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/protected_cookies/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protected_cookies"], "schema_version": 1, "sections": [{"aliases": ["protected cookies add httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:add_httponly", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "add_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies add secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:add_secure", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "add_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies disable tampering protection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:disable_tampering_protection", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "disable_tampering_protection"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies enable tampering protection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:enable_tampering_protection", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "enable_tampering_protection"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies ignore httponly"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:ignore_httponly", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_httponly"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies ignore max age"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:ignore_max_age", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_max_age"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies ignore samesite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:ignore_samesite", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_samesite"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies ignore secure"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:ignore_secure", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "ignore_secure"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies max age value"], "anchor": "schema-protected_cookies--max_age_value", "description": "Exclusive with Add max age attribute.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "max_age_value"], "syntax": "attribute", "type": "number"}, {"aliases": ["protected cookies name"], "anchor": "schema-protected_cookies--name", "description": "Name of the Cookie.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["protected cookies samesite lax"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:samesite_lax", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "samesite_lax"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies samesite none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:samesite_none", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "samesite_none"], "syntax": "attribute", "type": "object"}, {"aliases": ["protected cookies samesite strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:protected_cookies:samesite_strict", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protected_cookies", "samesite_strict"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/protected_cookies/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF (monitoring or blocking) will", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protected_cookies

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- protected_cookies

<a id="section"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/add_httponly/): complete subsection reference.

- [add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/add_secure/): complete subsection reference.

- [disable_tampering_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/disable_tampering_protection/): complete subsection reference.

- [enable_tampering_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/enable_tampering_protection/): complete subsection reference.

- [ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/ignore_httponly/): complete subsection reference.

- [ignore_max_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/ignore_max_age/): complete subsection reference.

- [ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/ignore_samesite/): complete subsection reference.

- [ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/ignore_secure/): complete subsection reference.

<a id="schema-protected_cookies--max_age_value"></a>

### max_age_value property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/samesite_lax/): complete subsection reference.

- [samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/samesite_none/): complete subsection reference.

- [samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/protected_cookies/samesite_strict/): complete subsection reference.
