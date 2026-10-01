---
page_title: "policy_based_challenge.rule_list.rules.spec.query_params"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec.query_params for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5032, "body_sha256": "sha256:a60adf592ca58ccdae88f65935051d104e2251fd8bead9a4d2ffbc60e02d2d46", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_not_present", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:check_present", "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:item"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec.query_params for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.query_params

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](data-sources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="section"></a>

Type: `"list"`. Computed.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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

- [check_not_present](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params--check_not_present.md): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params--check_present.md): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--query_params--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params--item.md): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--query_params--key"></a>

### key property

Type: `"string"`. Computed.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [policy_based_challenge.rule_list.rules.spec.query_params.check_not_present](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params--check_not_present.md)
- [policy_based_challenge.rule_list.rules.spec.query_params.check_present](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params--check_present.md)
- [policy_based_challenge.rule_list.rules.spec.query_params.item](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params--item.md)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
