---
page_title: "policy_based_challenge.rule_list.rules.spec.ip_prefix_list"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec.ip_prefix_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3002, "body_sha256": "sha256:ac5d078a8f68cb8af21c0edade8a8ba3b3c0d39c9d85d44db8783f2277fdc663", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "docs/guides/data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_prefix_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_prefix_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_prefix_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec.ip_prefix_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# policy_based_challenge.rule_list.rules.spec.ip_prefix_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [policy_based_challenge](data-sources--http_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="section"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

## Direct properties

<a id="schema-policy_based_challenge--rule_list--rules--spec--ip_prefix_list--invert_match"></a>

### invert_match property

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="schema-policy_based_challenge--rule_list--rules--spec--ip_prefix_list--ip_prefixes"></a>

### ip_prefixes property

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
