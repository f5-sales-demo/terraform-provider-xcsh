---
page_title: "rule_list.rules.asn_matcher"
subcategory: "DNS"
description: "rule_list.rules.asn_matcher for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 1855, "body_sha256": "sha256:53ec39d3323e5f11dba99eb4fc2e669437289e571e96dc08e8c4370be22cca1c", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "path": "documentation/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/index.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["rule_list", "rules", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.asn_matcher for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.asn_matcher

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/)
- rule_list.rules.asn_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
```

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
asn_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

- [asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/): complete subsection reference.

## Next pages

- [rule_list.rules.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
