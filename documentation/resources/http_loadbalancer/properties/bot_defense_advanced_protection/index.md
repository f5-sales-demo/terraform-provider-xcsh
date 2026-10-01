---
page_title: "bot_defense_advanced_protection"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2711, "body_sha256": "sha256:dc1bf56bd46ab07b11bbf1bdd155ec76ceb4f2d7d486fc9493df7d5c765435c3", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:mobile_only", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/bot_defense_advanced_protection/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["bot_defense_advanced_protection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- bot_defense_advanced_protection

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Advanced Protection - replaces BotDefenseAdvancedType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("both_web_and_mobile",
    "mobile_only"),
  validators.ConflictingObjectAttributes("both_web_and_mobile",
    "web_only"),
  validators.ConflictingObjectAttributes("mobile_only",
    "web_only")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_type_choice": "[\"both_web_and_mobile\",\"mobile_only\",\"web_only\"]"
}
```

Terraform syntax:

```terraform
bot_defense_advanced_protection {
  # Configure direct properties listed below.
}
```

## Direct properties

- [both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/): complete subsection reference.

- [mobile_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/mobile_only/): complete subsection reference.

- [web_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/)
- [bot_defense_advanced_protection.mobile_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/mobile_only/)
- [bot_defense_advanced_protection.web_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
