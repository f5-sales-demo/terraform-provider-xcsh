---
page_title: "single_lb_app.enable_discovery"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3516, "body_sha256": "sha256:611f5c439582dce158b8cc38e5c3a00ff9ad3a8f57fa3a90d704862ecd6a4681", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_discovery_from_code_scan", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:custom_api_auth_discovery", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:default_api_auth_discovery", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:disable_learn_from_redirect_traffic", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:discovered_api_settings", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:enable_learn_from_redirect_traffic"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app", "path": "docs/guides/data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [single_lb_app](data-sources--http_loadbalancer--properties--single_lb_app.md)
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

- [api_crawler](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler.md): complete subsection reference.

- [api_discovery_from_code_scan](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_discovery_from_code_scan.md): complete subsection reference.

- [custom_api_auth_discovery](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--custom_api_auth_discovery.md): complete subsection reference.

- [default_api_auth_discovery](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--default_api_auth_discovery.md): complete subsection reference.

- [disable_learn_from_redirect_traffic](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--disable_learn_from_redirect_traffic.md): complete subsection reference.

- [discovered_api_settings](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--discovered_api_settings.md): complete subsection reference.

- [enable_learn_from_redirect_traffic](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--enable_learn_from_redirect_traffic.md): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler.md)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_discovery_from_code_scan.md)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--custom_api_auth_discovery.md)
- [single_lb_app.enable_discovery.default_api_auth_discovery](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--default_api_auth_discovery.md)
- [single_lb_app.enable_discovery.disable_learn_from_redirect_traffic](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--disable_learn_from_redirect_traffic.md)
- [single_lb_app.enable_discovery.discovered_api_settings](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--discovered_api_settings.md)
- [single_lb_app.enable_discovery.enable_learn_from_redirect_traffic](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--enable_learn_from_redirect_traffic.md)
- [single_lb_app](data-sources--http_loadbalancer--properties--single_lb_app.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
