---
page_title: "enable_api_discovery"
subcategory: "Load Balancing"
description: "Specifies the settings used for API discovery."
xcsh_docs: {"aliases": ["enable api discovery"], "body_bytes": 2535, "body_sha256": "sha256:505394271db763036e8c8a92a6677d197324c31116282e61fc37ca88bfa3d028", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_crawler", "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:default_api_auth_discovery", "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:discovered_api_settings", "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:enable_learn_from_redirect_traffic"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/enable_api_discovery/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery api crawler"], "anchor": "section", "description": "API Crawler message.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_crawler", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler"], "syntax": "block", "type": "object"}, {"aliases": ["enable api discovery api discovery from code scan"], "anchor": "section", "description": "Select Code Base and Repositories.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan"], "syntax": "block", "type": "object"}, {"aliases": ["enable api discovery custom api auth discovery"], "anchor": "section", "description": "API Discovery Advanced settings.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:custom_api_auth_discovery", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery"], "syntax": "block", "type": "object"}, {"aliases": ["enable api discovery default api auth discovery"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:default_api_auth_discovery", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "default_api_auth_discovery"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable api discovery disable learn from redirect traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "disable_learn_from_redirect_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable api discovery discovered api settings"], "anchor": "section", "description": "Configure Discovered API Settings.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:discovered_api_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "discovered_api_settings"], "syntax": "block", "type": "object"}, {"aliases": ["enable api discovery enable learn from redirect traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:enable_learn_from_redirect_traffic", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "enable_learn_from_redirect_traffic"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_api_discovery/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Specifies the settings used for API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- enable_api_discovery

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

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

- [api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/api_crawler/): complete subsection reference.

- [api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/): complete subsection reference.

- [custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/custom_api_auth_discovery/): complete subsection reference.

- [default_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/default_api_auth_discovery/): complete subsection reference.

- [disable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/disable_learn_from_redirect_traffic/): complete subsection reference.

- [discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/discovered_api_settings/): complete subsection reference.

- [enable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/enable_api_discovery/enable_learn_from_redirect_traffic/): complete subsection reference.
