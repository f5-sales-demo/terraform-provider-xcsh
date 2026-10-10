---
page_title: "rule_list.rules.ip_prefix_set"
subcategory: "DNS"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["rule list rules ip prefix set"], "body_bytes": 1876, "body_sha256": "sha256:d9d0435068fec360ac210c3fa8d995fea910ec727cfbcfbeccc700c62d99b282", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set:prefix_sets"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "path": "documentation/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1021103121300020-3113310130101200-2121330213301130-1123000102230021-2310311211332002-0032132003201311-1312223122232100-2322120232033331", "registry_path": "docs/guides/resources--dns_load_balancer--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.ip_prefix_set:RequiredObjectAttributes:prefix_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set:prefix_sets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["rule list rules ip prefix set invert matcher"], "anchor": "schema-rule_list--rules--ip_prefix_set--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "ip_prefix_set", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["rule list rules ip prefix set prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set:prefix_sets", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "ip_prefix_set", "prefix_sets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.ip_prefix_set

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/)
- rule_list.rules.ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--ip_prefix_set--invert_matcher"></a>

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/): complete subsection reference.
