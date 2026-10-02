---
page_title: "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules"
subcategory: "Load Balancing"
description: "This defines custom JavaScript insertion rules for Bot Defense Policy."
xcsh_docs: {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection both web and mobile js insertion rules"], "body_bytes": 2850, "body_sha256": "sha256:581c76369aa03c2e16bb7a02441c503b51836801d70eaecce3ede10008faaa0e", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:rules"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "path": "documentation/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-013.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insertion_rules"], "schema_version": 1, "sections": [{"aliases": ["exclude list"], "anchor": "section", "description": "Optional JavaScript insertions exclude list of domain and path matchers.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insertion_rules", "exclude_list"], "syntax": "block", "type": "object"}, {"aliases": ["rules"], "anchor": "section", "description": "Required list of pages to insert Bot Defense client JavaScript.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:rules:domain", "type": "conflicts"}], "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insertion_rules", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines custom JavaScript insertion rules for Bot Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [bot_defense_advanced_protection.both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/): complete subsection reference.

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/rules/): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/rules/)
- [bot_defense_advanced_protection.both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
