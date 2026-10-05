---
page_title: "origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher"
subcategory: "Load Balancing"
description: "Match any AS number contained in the list of bgp_asn_sets."
xcsh_docs: {"aliases": ["origin server subset rule list origin server subset rules asn matcher"], "body_bytes": 2275, "body_sha256": "sha256:2ed4d331ed2a20059fab2ac8adc24096fd6f154d4fff88edd8b958276c5906fe", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher:asn_sets"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "path": "documentation/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2131232131220022-3231312032001311-1001123321021221-1010012233233102-3032020321321013-2003031121101100-2013021031132131-2032111122221213", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-022.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher:RequiredObjectAttributes:asn_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher:asn_sets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "asn_matcher"], "schema_version": 1, "sections": [{"aliases": ["origin server subset rule list origin server subset rules asn matcher asn sets"], "anchor": "section", "description": "A list of references to bgp_asn_set objects.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher:asn_sets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "asn_matcher", "asn_sets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_matcher/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Match any AS number contained in the list of bgp_asn_sets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [origin_server_subset_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/)
- [origin_server_subset_rule_list.origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher

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

- [asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_matcher/asn_sets/): complete subsection reference.

## Next pages

- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/asn_matcher/asn_sets/)
- [origin_server_subset_rule_list.origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
