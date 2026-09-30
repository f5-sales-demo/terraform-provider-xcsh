---
page_title: "rule_list.rules"
subcategory: "DNS"
description: "rule_list.rules for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 5163, "body_sha256": "sha256:bfb492a0752318b8305146c3d4acc8cf0ef97f2368fc0aeece8cce9a48ba6e4e", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:asn_list", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:pool"], "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list", "path": "documentation/data-sources/dns_load_balancer/properties/rule_list/rules/index.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list.rules

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/)
- rule_list.rules

<a id="section"></a>

Type: `"list"`. Computed.

Load Balancing Rules. Rules to perform load balancing.

Upstream description:

Rules to perform load balancing.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/): complete subsection reference.

- [geo_location_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_label_selector/): complete subsection reference.

- [geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_set/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/): complete subsection reference.

- [pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/pool/): complete subsection reference.

<a id="schema-rule_list--rules--score"></a>

### score property

Type: `"number"`. Computed.

When multiple load balancing rules match a query, the one with the highest score is chosen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32767,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32767"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32767"
  }
}
```

## Next pages

- [rule_list.rules.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_list/)
- [rule_list.rules.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/)
- [rule_list.rules.geo_location_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_label_selector/)
- [rule_list.rules.geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_set/)
- [rule_list.rules.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/)
- [rule_list.rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/)
- [rule_list.rules.pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/pool/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
