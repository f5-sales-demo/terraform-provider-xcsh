---
page_title: "rule_list.rules"
subcategory: "DNS"
description: "Rules to perform load balancing."
xcsh_docs: {"aliases": ["rule list rules"], "body_bytes": 3797, "body_sha256": "sha256:4c097933910cb1779d9947d949306a7a6a9803cb2ff0fa4e5937012f2fa7ec37", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:asn_list", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:pool"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list", "path": "documentation/data-sources/dns_load_balancer/properties/rule_list/rules/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103", "registry_path": "docs/guides/data-sources--dns_load_balancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "sections": [{"aliases": ["rule list rules asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:asn_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "asn_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules asn matcher"], "anchor": "section", "description": "Match any AS number contained in the list of bgp_asn_sets.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "asn_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules geo location label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "geo_location_label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules geo location set"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "geo_location_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "ip_prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules ip prefix set"], "anchor": "section", "description": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules pool"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules:pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules score"], "anchor": "schema-rule_list--rules--score", "description": "When multiple load balancing rules match a query, the one with the highest score is chosen.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "score"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Rules to perform load balancing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/)
- rule_list.rules

<a id="section"></a>

Type: `"list"`. Computed.

Load Balancing Rules. Rules to perform load balancing.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
