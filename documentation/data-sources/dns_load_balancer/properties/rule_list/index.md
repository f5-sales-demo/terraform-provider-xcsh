---
page_title: "rule_list"
subcategory: "DNS"
description: "List of the Load Balancing Rules."
xcsh_docs: {"aliases": ["rule list"], "body_bytes": 845, "body_sha256": "sha256:ec3ee19726917a3da722e3b17256f83b9d9f61344d7ed4ece9db17dbce7ea0c8", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "path": "documentation/data-sources/dns_load_balancer/properties/rule_list/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223", "registry_path": "docs/guides/data-sources--dns_load_balancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "sections": [{"aliases": ["rule list rules"], "anchor": "section", "description": "Rules to perform load balancing.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/rule_list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of the Load Balancing Rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- rule_list

<a id="section"></a>

Type: `"single"`. Computed.

Load Balancing Rule List. List of the Load Balancing Rules.

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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/): complete subsection reference.
