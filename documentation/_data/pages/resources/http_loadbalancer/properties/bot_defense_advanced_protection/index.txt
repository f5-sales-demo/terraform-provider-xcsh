---
page_title: "bot_defense_advanced_protection"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2612, "body_sha256": "sha256:6397f8e1df8025427f1a821eca0e6bd67e12960e28120be44ce8f8e3be21228d", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:mobile_only", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/bot_defense_advanced_protection/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["bot_defense_advanced_protection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
