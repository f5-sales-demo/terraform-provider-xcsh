---
page_title: "policy_based_challenge.rule_list.rules.spec.asn_list"
subcategory: "Load Balancing"
description: "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer."
xcsh_docs: {"aliases": ["policy based challenge rule list rules spec asn list"], "body_bytes": 3077, "body_sha256": "sha256:86677b1ec1369ede4305bdf6ebbb811371caba6af19366e859fdacc597f25b63", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "documentation/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/asn_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0003233223300201-0312002012103020-1303002203022023-2111233231333112-1000323100130012-1320213001201033-1202333023000302-1013322233111312", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "asn_list"], "schema_version": 1, "sections": [{"aliases": ["policy based challenge rule list rules spec asn list as numbers"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--asn_list--as_numbers", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "asn_list", "as_numbers"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/asn_list/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.asn_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
asn_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-policy_based_challenge--rule_list--rules--spec--asn_list--as_numbers"></a>

### as_numbers property

Type: `["list", "number"]`. Optional.

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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
