---
page_title: "enable_api_discovery"
subcategory: "Load Balancing"
description: "Specifies the settings used for API discovery."
xcsh_docs: {"aliases": ["enable api discovery"], "body_bytes": 4467, "body_sha256": "sha256:ffb88c5b34fd63f16cbd93d2dddeef41591ac025f3ee67ed2d5c89269011059a", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:default_api_auth_discovery", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:discovered_api_settings", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:enable_learn_from_redirect_traffic"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/enable_api_discovery/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery:ConflictingObjectAttributes:custom_api_auth_discovery,default_api_auth_discovery", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery:ConflictingObjectAttributes:custom_api_auth_discovery,default_api_auth_discovery", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:default_api_auth_discovery", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery:ConflictingObjectAttributes:disable_learn_from_redirect_traffic,enable_learn_from_redirect_traffic", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery:ConflictingObjectAttributes:disable_learn_from_redirect_traffic,enable_learn_from_redirect_traffic", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:enable_learn_from_redirect_traffic", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery"], "schema_version": 1, "sections": [{"aliases": ["api crawler"], "anchor": "section", "description": "API Crawler message.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_crawler:ConflictingObjectAttributes:api_crawler_config,disable_api_crawler", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_crawler:ConflictingObjectAttributes:api_crawler_config,disable_api_crawler", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:disable_api_crawler", "type": "conflicts"}], "schema_path": ["enable_api_discovery", "api_crawler"], "syntax": "block", "type": "object"}, {"aliases": ["api discovery from code scan"], "anchor": "section", "description": "Select Code Base and Repositories.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_discovery_from_code_scan:RequiredObjectAttributes:code_base_integrations", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations", "type": "requires"}], "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan"], "syntax": "block", "type": "object"}, {"aliases": ["custom api auth discovery"], "anchor": "section", "description": "API Discovery Advanced settings.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery"], "syntax": "block", "type": "object"}, {"aliases": ["default api auth discovery"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:default_api_auth_discovery", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "default_api_auth_discovery"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable learn from redirect traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "disable_learn_from_redirect_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovered api settings"], "anchor": "section", "description": "Configure Discovered API Settings.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:discovered_api_settings", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-enable_api_discovery--discovered_api_settings--purge_duration_for_inactive_discovered_apis", "enforcement": "provider-schema", "group": "enable_api_discovery.discovered_api_settings:RequiredObjectAttributes:purge_duration_for_inactive_discovered_apis", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:discovered_api_settings", "type": "requires"}], "schema_path": ["enable_api_discovery", "discovered_api_settings"], "syntax": "block", "type": "object"}, {"aliases": ["enable learn from redirect traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:enable_learn_from_redirect_traffic", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "enable_learn_from_redirect_traffic"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_api_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specifies the settings used for API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- enable_api_discovery

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_api_auth_discovery",
    "default_api_auth_discovery"),
  validators.ConflictingObjectAttributes("disable_learn_from_redirect_traffic",
    "enable_learn_from_redirect_traffic")}
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
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_api_discovery {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/): complete subsection reference.

- [api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/): complete subsection reference.

- [custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/): complete subsection reference.

- [default_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/default_api_auth_discovery/): complete subsection reference.

- [disable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/disable_learn_from_redirect_traffic/): complete subsection reference.

- [discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/discovered_api_settings/): complete subsection reference.

- [enable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/enable_learn_from_redirect_traffic/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/)
- [enable_api_discovery.api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/)
- [enable_api_discovery.custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/)
- [enable_api_discovery.default_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/default_api_auth_discovery/)
- [enable_api_discovery.disable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/disable_learn_from_redirect_traffic/)
- [enable_api_discovery.discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/discovered_api_settings/)
- [enable_api_discovery.enable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/enable_learn_from_redirect_traffic/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
