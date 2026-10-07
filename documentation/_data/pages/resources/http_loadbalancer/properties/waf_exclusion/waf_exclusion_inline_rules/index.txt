---
page_title: "waf_exclusion.waf_exclusion_inline_rules"
subcategory: "Load Balancing"
description: "A list of WAF exclusion rules that will be applied inline."
xcsh_docs: {"aliases": ["waf exclusion waf exclusion inline rules"], "body_bytes": 1192, "body_sha256": "sha256:6e28d161f3c5a875bc005f7fa2847ae0ade78a69a7bc31e814855c71e87dde6e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion", "path": "documentation/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-028.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules"], "schema_version": 1, "sections": [{"aliases": ["waf exclusion waf exclusion inline rules rules"], "anchor": "section", "description": "An ordered list of WAF Exclusions specific to this Load Balancer.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--exact_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--exact_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_prefix", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_prefix", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:path_prefix,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_regex", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_regex", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:path_prefix,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--suffix_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--suffix_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:waf_skip_processing", "type": "conflicts"}], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "A list of WAF exclusion rules that will be applied inline.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion.waf_exclusion_inline_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/)
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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/): complete subsection reference.
