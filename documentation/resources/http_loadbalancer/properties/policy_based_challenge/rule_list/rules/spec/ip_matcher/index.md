---
page_title: "policy_based_challenge.rule_list.rules.spec.ip_matcher"
subcategory: "Load Balancing"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["policy based challenge rule list rules spec ip matcher"], "body_bytes": 2373, "body_sha256": "sha256:c4415de56481beba7a9efdfec110aabf07938c8d0e9443fdeecb9f90de4997b4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher:prefix_sets"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "path": "documentation/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec.ip_matcher:RequiredObjectAttributes:prefix_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher:prefix_sets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_matcher"], "schema_version": 1, "sections": [{"aliases": ["policy based challenge rule list rules spec ip matcher invert matcher"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--ip_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["policy based challenge rule list rules spec ip matcher prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher:prefix_sets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "ip_matcher", "prefix_sets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_matcher/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.ip_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-policy_based_challenge--rule_list--rules--spec--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/ip_matcher/prefix_sets/): complete subsection reference.
