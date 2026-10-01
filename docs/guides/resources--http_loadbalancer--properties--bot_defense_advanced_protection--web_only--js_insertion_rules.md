---
page_title: "bot_defense_advanced_protection.web_only.js_insertion_rules"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.web_only.js_insertion_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2220, "body_sha256": "sha256:dccf632cf1c2f8525736bd9ac825076cc5199a74957f10609680bab0de1256ce", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules:exclude_list", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules:rules"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insertion_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.web_only.js_insertion_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.web_only.js_insertion_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense_advanced_protection](resources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md)
- bot_defense_advanced_protection.web_only.js_insertion_rules

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

- [exclude_list](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules--exclude_list.md): complete subsection reference.

- [rules](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules--rules.md): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules--exclude_list.md)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules--rules.md)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
