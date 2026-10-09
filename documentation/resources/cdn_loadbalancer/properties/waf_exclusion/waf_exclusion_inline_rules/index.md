---
page_title: "waf_exclusion.waf_exclusion_inline_rules"
subcategory: "Load Balancing"
description: "A list of WAF exclusion rules that will be applied inline."
xcsh_docs: {"aliases": ["waf exclusion waf exclusion inline rules"], "body_bytes": 1187, "body_sha256": "sha256:68aa12538c6fb936ea1756e01d3d64dc59fafde27ba6c01207e8d5a3072d52e6", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion", "path": "documentation/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules"], "schema_version": 1, "sections": [{"aliases": ["waf exclusion waf exclusion inline rules rules"], "anchor": "section", "description": "An ordered list of WAF Exclusions specific to this Load Balancer.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "A list of WAF exclusion rules that will be applied inline.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion.waf_exclusion_inline_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/)
- waf_exclusion.waf_exclusion_inline_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
waf_exclusion_inline_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/): complete subsection reference.
