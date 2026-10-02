---
page_title: "waf_exclusion.waf_exclusion_inline_rules"
subcategory: "Load Balancing"
description: "A list of WAF exclusion rules that will be applied inline."
xcsh_docs: {"aliases": ["waf exclusion waf exclusion inline rules"], "body_bytes": 1717, "body_sha256": "sha256:8ec5bd35422cd9a76fc4bc67adccc0406b7fab5a20fa8569e846103eda45c807", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion", "path": "documentation/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules"], "schema_version": 1, "sections": [{"aliases": ["rules"], "anchor": "section", "description": "An ordered list of WAF Exclusions specific to this Load Balancer.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--exact_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--exact_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_prefix", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_prefix", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:path_prefix,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_regex", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_regex", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:path_prefix,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--suffix_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--suffix_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:waf_skip_processing", "type": "conflicts"}], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A list of WAF exclusion rules that will be applied inline.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Terraform syntax:

```terraform
waf_exclusion_inline_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/)
- [waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/waf_exclusion/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
