---
page_title: "forward_proxy_pbr.forward_proxy_pbr_rules.http_list"
subcategory: ""
description: "URLListType."
xcsh_docs: {"aliases": ["forward proxy pbr forward proxy pbr rules http list"], "body_bytes": 1278, "body_sha256": "sha256:f4f61d9eb381d2df175bf3d46df462ed788931c4a9fe89687d6b1ff9a554fef5", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list:http_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "parent_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "path": "documentation/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1323233330201000-1123120330313023-2232010013232112-0133020112302123-0333100330230222-0111103031022313-2310300210200300-0003022222231013", "registry_path": "docs/guides/data-sources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "http_list"], "schema_version": 1, "sections": [{"aliases": ["forward proxy pbr forward proxy pbr rules http list http list"], "anchor": "section", "description": "URLs for HTTP connections.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list:http_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "http_list", "http_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "URLListType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# forward_proxy_pbr.forward_proxy_pbr_rules.http_list

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/)
- [forward_proxy_pbr.forward_proxy_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list

<a id="section"></a>

Type: `"single"`. Computed.

URLListType.

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

- [http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/http_list/): complete subsection reference.
