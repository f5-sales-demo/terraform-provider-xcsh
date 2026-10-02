---
page_title: "single_lb_app.enable_discovery"
subcategory: "Load Balancing"
description: "Specifies the settings used for API discovery."
xcsh_docs: {"aliases": ["single lb app enable discovery"], "body_bytes": 4459, "body_sha256": "sha256:f6156f7183d67571b91598eda25bfff20bd02093f99e83a275a2f6aedd11272e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:default_api_auth_discovery", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:disable_learn_from_redirect_traffic", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:discovered_api_settings", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:enable_learn_from_redirect_traffic"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app", "path": "documentation/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery"], "schema_version": 1, "sections": [{"aliases": ["api crawler"], "anchor": "section", "description": "API Crawler message.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler"], "syntax": "attribute", "type": "object"}, {"aliases": ["api discovery from code scan"], "anchor": "section", "description": "Select Code Base and Repositories.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_discovery_from_code_scan"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom api auth discovery"], "anchor": "section", "description": "API Discovery Advanced settings.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "custom_api_auth_discovery"], "syntax": "attribute", "type": "object"}, {"aliases": ["default api auth discovery"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:default_api_auth_discovery", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "default_api_auth_discovery"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable learn from redirect traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:disable_learn_from_redirect_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "disable_learn_from_redirect_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovered api settings"], "anchor": "section", "description": "Configure Discovered API Settings.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:discovered_api_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "discovered_api_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable learn from redirect traffic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:enable_learn_from_redirect_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "enable_learn_from_redirect_traffic"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies the settings used for API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/)
- single_lb_app.enable_discovery

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/): complete subsection reference.

- [api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/): complete subsection reference.

- [custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/custom_api_auth_discovery/): complete subsection reference.

- [default_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/default_api_auth_discovery/): complete subsection reference.

- [disable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/disable_learn_from_redirect_traffic/): complete subsection reference.

- [discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/discovered_api_settings/): complete subsection reference.

- [enable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/enable_learn_from_redirect_traffic/): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_discovery_from_code_scan/)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/custom_api_auth_discovery/)
- [single_lb_app.enable_discovery.default_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/default_api_auth_discovery/)
- [single_lb_app.enable_discovery.disable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/disable_learn_from_redirect_traffic/)
- [single_lb_app.enable_discovery.discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/discovered_api_settings/)
- [single_lb_app.enable_discovery.enable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/enable_learn_from_redirect_traffic/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
