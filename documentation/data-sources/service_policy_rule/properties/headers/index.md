---
page_title: "headers"
subcategory: ""
description: "headers for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 5012, "body_sha256": "sha256:2d8b67bf881d7c705e81a85230a8ce43eaea31e04793758a9e9fdbbc768f8e5e", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:headers:check_not_present", "xcsh-docs:data-sources:service_policy_rule:properties:headers:check_present", "xcsh-docs:data-sources:service_policy_rule:properties:headers:item"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:headers", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "documentation/data-sources/service_policy_rule/properties/headers/index.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["headers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "headers for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# headers

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- headers

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

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/headers/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/headers/check_present/): complete subsection reference.

<a id="schema-headers--invert_matcher"></a>

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

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/headers/item/): complete subsection reference.

<a id="schema-headers--name"></a>

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

- [headers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/headers/check_not_present/)
- [headers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/headers/check_present/)
- [headers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/headers/item/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
