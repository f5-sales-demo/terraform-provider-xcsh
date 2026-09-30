---
page_title: "bot_defense_advanced_protection.web_only"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.web_only for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3438, "body_sha256": "sha256:856eecf7ea2b63030120cdc1fbb34f972cb14e594b56b9f204c04edfa16af22c", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:disable_js_insert", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insertion_rules", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:web"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense_advanced_protection", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.web_only for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense_advanced_protection.web_only

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense_advanced_protection](resources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- bot_defense_advanced_protection.web_only

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Web. Web only configuration.

Upstream description:

Web only configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

Terraform syntax:

```terraform
web_only {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_js_insert](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--disable_js_insert.md): complete subsection reference.

- [js_insert_all_pages](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages.md): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except.md): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules.md): complete subsection reference.

- [web](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--web.md): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.web_only.disable_js_insert](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--disable_js_insert.md)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages.md)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except.md)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insertion_rules.md)
- [bot_defense_advanced_protection.web_only.web](resources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--web.md)
- [bot_defense_advanced_protection](resources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
