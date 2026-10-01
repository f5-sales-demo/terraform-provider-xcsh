---
page_title: "arg_matchers"
subcategory: ""
description: "arg_matchers for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 4575, "body_sha256": "sha256:df801408bb0eae9a6eab1856a7e42682eb976765e9e59391d8b5d78b3a3bd7cc", "canonical_id": "xcsh-docs:data-sources:service_policy_rule:properties:arg_matchers", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:arg_matchers:check_not_present", "xcsh-docs:data-sources:service_policy_rule:properties:arg_matchers:check_present", "xcsh-docs:data-sources:service_policy_rule:properties:arg_matchers:item"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:arg_matchers", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "docs/guides/data-sources--service_policy_rule--properties--arg_matchers.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["arg_matchers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/arg_matchers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "arg_matchers for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# arg_matchers

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- arg_matchers

<a id="section"></a>

Type: `"list"`. Computed.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

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
    }
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

- [check_not_present](data-sources--service_policy_rule--properties--arg_matchers--check_not_present.md): complete subsection reference.

- [check_present](data-sources--service_policy_rule--properties--arg_matchers--check_present.md): complete subsection reference.

<a id="schema-arg_matchers--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--service_policy_rule--properties--arg_matchers--item.md): complete subsection reference.

<a id="schema-arg_matchers--name"></a>

### name property

Type: `"string"`. Computed.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

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
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [arg_matchers.check_not_present](data-sources--service_policy_rule--properties--arg_matchers--check_not_present.md)
- [arg_matchers.check_present](data-sources--service_policy_rule--properties--arg_matchers--check_present.md)
- [arg_matchers.item](data-sources--service_policy_rule--properties--arg_matchers--item.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
