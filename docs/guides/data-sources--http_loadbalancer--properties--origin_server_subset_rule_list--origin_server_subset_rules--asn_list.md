---
page_title: "origin_server_subset_rule_list.origin_server_subset_rules.asn_list"
subcategory: "Load Balancing"
description: "origin_server_subset_rule_list.origin_server_subset_rules.asn_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3235, "body_sha256": "sha256:1db0760b3591f76ec23b8d9a2ad53bde200ccf5d9cf1e03fe22c60b3088f8515", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_list", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules--asn_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "asn_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_server_subset_rule_list.origin_server_subset_rules.asn_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list.origin_server_subset_rules.asn_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list.md)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules.md)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_list

<a id="section"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="schema-origin_server_subset_rule_list--origin_server_subset_rules--asn_list--as_numbers"></a>

### as_numbers property

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--properties--origin_server_subset_rule_list--origin_server_subset_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
