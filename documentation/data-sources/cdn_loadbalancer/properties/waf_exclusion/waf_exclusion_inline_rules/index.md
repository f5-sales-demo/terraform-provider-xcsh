---
page_title: "waf_exclusion.waf_exclusion_inline_rules"
subcategory: "Load Balancing"
description: "A list of WAF exclusion rules that will be applied inline."
xcsh_docs: {"aliases": ["waf exclusion waf exclusion inline rules"], "body_bytes": 1603, "body_sha256": "sha256:30cbae0ea7b99ec0fb38bc8cd028efc80dad5b0a07cf2e746c80afa294800f13", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion", "path": "documentation/data-sources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules"], "schema_version": 1, "sections": [{"aliases": ["waf exclusion waf exclusion inline rules rules"], "anchor": "section", "description": "An ordered list of WAF Exclusions specific to this Load Balancer.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "A list of WAF exclusion rules that will be applied inline.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion.waf_exclusion_inline_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/waf_exclusion/)
- waf_exclusion.waf_exclusion_inline_rules

<a id="section"></a>

Type: `"single"`. Computed.

List of WAF exclusion rules that will be applied inline.

Upstream description:

A list of WAF exclusion rules that will be applied inline.

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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/)
- [waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/waf_exclusion/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
