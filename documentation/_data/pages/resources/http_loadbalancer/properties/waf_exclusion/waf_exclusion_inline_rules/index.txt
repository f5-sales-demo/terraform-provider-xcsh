---
page_title: "waf_exclusion.waf_exclusion_inline_rules"
subcategory: "Load Balancing"
description: "A list of WAF exclusion rules that will be applied inline."
xcsh_docs: {"aliases": ["waf exclusion waf exclusion inline rules"], "body_bytes": 1726, "body_sha256": "sha256:f6409d7094d983c1ce252fe6e365017376d81b357cb1a0afb6cff4d9b7400356", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion", "path": "documentation/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules"], "schema_version": 1, "sections": [{"aliases": ["waf exclusion waf exclusion inline rules rules"], "anchor": "section", "description": "An ordered list of WAF Exclusions specific to this Load Balancer.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--exact_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--exact_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_prefix", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_prefix", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:path_prefix,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_regex", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--path_regex", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:path_prefix,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--suffix_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "schema-waf_exclusion--waf_exclusion_inline_rules--rules--suffix_value", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,exact_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_domain,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:any_path,path_regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:any_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_exclusion.waf_exclusion_inline_rules.rules:ConflictingListObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:waf_exclusion:waf_exclusion_inline_rules:rules:waf_skip_processing", "type": "conflicts"}], "schema_path": ["waf_exclusion", "waf_exclusion_inline_rules", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "A list of WAF exclusion rules that will be applied inline.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/): complete subsection reference.

## Next pages

- [waf_exclusion.waf_exclusion_inline_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/waf_exclusion_inline_rules/rules/)
- [waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/waf_exclusion/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
