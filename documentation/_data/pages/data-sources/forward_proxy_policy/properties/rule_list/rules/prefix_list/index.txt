---
page_title: "rule_list.rules.prefix_list"
subcategory: "Security"
description: "List of IPv4 prefixes that represent an endpoint."
xcsh_docs: {"aliases": ["rule list rules prefix list"], "body_bytes": 2334, "body_sha256": "sha256:3d7c1eda56bb9b4f39de1cb9d6f438407359efc43fe30b057a9eb19c2ba3b65e", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:prefix_list", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules", "path": "documentation/data-sources/forward_proxy_policy/properties/rule_list/rules/prefix_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3212130112120111-0303230130202111-3022100232223233-3002201320311201-2201332023302332-2211231010010220-3022300313111010-2321121102331231", "registry_path": "docs/guides/data-sources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "prefix_list"], "schema_version": 1, "sections": [{"aliases": ["prefixes"], "anchor": "schema-rule_list--rules--prefix_list--prefixes", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "prefix_list", "prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/rule_list/rules/prefix_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of IPv4 prefixes that represent an endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.prefix_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/)
- rule_list.rules.prefix_list

<a id="section"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="schema-rule_list--rules--prefix_list--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
