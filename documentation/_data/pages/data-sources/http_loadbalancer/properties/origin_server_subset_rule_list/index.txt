---
page_title: "origin_server_subset_rule_list"
subcategory: "Load Balancing"
description: "List of Origin Pools."
xcsh_docs: {"aliases": ["backend servers", "origin server subset rule list", "origin servers", "upstream servers"], "body_bytes": 949, "body_sha256": "sha256:050043d12458393de8d9ac8fb4cfdacc4179f01481cbecdbfaedc24dde0888f8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-021.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_server_subset_rule_list"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "origin server subset rule list origin server subset rules", "origin servers", "upstream servers"], "anchor": "section", "description": "Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN, Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin Server Subset is a sequential engine where rules are evaluated one after the other. It's important to define the correct", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "List of Origin Pools.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- origin_server_subset_rule_list

<a id="section"></a>

Type: `"single"`. Computed.

Origin Server Subset Rule List Type. List of Origin Pools.

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

- [origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/): complete subsection reference.
