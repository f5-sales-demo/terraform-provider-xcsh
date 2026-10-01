---
page_title: "routes.simple_route.advanced_options.response_cookies_to_add.secret_value"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.response_cookies_to_add.secret_value for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2449, "body_sha256": "sha256:33a0b7e0b7974a461606dff926f1353a2899ce0e1efb84965713444469bd9173", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "path": "docs/guides/data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--secret_value.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.response_cookies_to_add.secret_value for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.response_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](data-sources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md)
- routes.simple_route.advanced_options.response_cookies_to_add.secret_value

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--secret_value--blindfold_secret_info.md)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--secret_value--clear_secret_info.md)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
