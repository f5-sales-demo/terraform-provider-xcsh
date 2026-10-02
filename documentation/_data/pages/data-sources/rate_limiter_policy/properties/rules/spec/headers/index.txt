---
page_title: "rules.spec.headers"
subcategory: "Security"
description: "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header"
xcsh_docs: {"aliases": ["rules spec headers"], "body_bytes": 5408, "body_sha256": "sha256:2805b8f42354ef1f5aa5ac0ec1cdd05251cb77a1d9ed1b570775f51ae3fe6782", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers:check_not_present", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers:check_present", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers:item"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec", "path": "documentation/data-sources/rate_limiter_policy/properties/rules/spec/headers/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201", "registry_path": "docs/guides/data-sources--rate_limiter_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "spec", "headers"], "schema_version": 1, "sections": [{"aliases": ["check not present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers:check_not_present", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "headers", "check_not_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["check present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers:check_present", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "headers", "check_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["invert matcher"], "anchor": "schema-rules--spec--headers--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "headers", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["item", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers:item", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "headers", "item"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-rules--spec--headers--name", "description": "A case-insensitive HTTP header name.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "headers", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/headers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.headers

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/)
- rules.spec.headers

<a id="section"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

## Direct properties

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/check_present/): complete subsection reference.

<a id="schema-rules--spec--headers--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/item/): complete subsection reference.

<a id="schema-rules--spec--headers--name"></a>

### name property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "maxLength": 63,
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [rules.spec.headers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/check_not_present/)
- [rules.spec.headers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/check_present/)
- [rules.spec.headers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/item/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
