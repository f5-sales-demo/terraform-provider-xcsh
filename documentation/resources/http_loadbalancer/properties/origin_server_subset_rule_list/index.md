---
page_title: "origin_server_subset_rule_list"
subcategory: "Load Balancing"
description: "List of Origin Pools."
xcsh_docs: {"aliases": ["backend servers", "origin server subset rule list", "origin servers", "upstream servers"], "body_bytes": 1591, "body_sha256": "sha256:508f1dcc3e08a47926a7c5b02eba96e2f6e7bba3aa916e1ecee4e5f0f7cdbae4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/origin_server_subset_rule_list/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_server_subset_rule_list"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "origin server subset rules", "origin servers", "upstream servers"], "anchor": "section", "description": "Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN, Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin Server Subset is a sequential engine where rules are evaluated one after the other. It's important to define the correct", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:client_selector,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:ConflictingListObjectAttributes:client_selector,none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:none", "type": "conflicts"}, {"anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--origin_server_subsets_action", "enforcement": "provider-schema", "group": "origin_server_subset_rule_list.origin_server_subset_rules:RequiredListObjectAttributes:origin_server_subsets_action", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "type": "requires"}], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/origin_server_subset_rule_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of Origin Pools.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- origin_server_subset_rule_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Origin Server Subset Rule List Type. List of Origin Pools.

Upstream description:

List of Origin Pools.

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
origin_server_subset_rule_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/): complete subsection reference.

## Next pages

- [origin_server_subset_rule_list.origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
