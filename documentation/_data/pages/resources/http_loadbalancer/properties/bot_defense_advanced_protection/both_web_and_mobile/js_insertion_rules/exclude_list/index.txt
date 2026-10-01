---
page_title: "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4772, "body_sha256": "sha256:3abfc73cbf166d60348546739e83d1082a65cecd35461f6445931dbf287850f4", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list:domain", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list:metadata", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list:path"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:js_insertion_rules", "path": "documentation/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "js_insertion_rules", "exclude_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [bot_defense_advanced_protection.both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/path/): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/any_domain/)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/domain/)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/metadata/)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/exclude_list/path/)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/js_insertion_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
