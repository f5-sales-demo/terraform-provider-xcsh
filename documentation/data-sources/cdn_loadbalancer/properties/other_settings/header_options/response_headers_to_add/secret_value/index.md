---
page_title: "other_settings.header_options.response_headers_to_add.secret_value"
subcategory: "Load Balancing"
description: "other_settings.header_options.response_headers_to_add.secret_value for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2795, "body_sha256": "sha256:2f80b0a89e79415aec2628d9f2c548759efc41e49974e563738dcc830cc2657e", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add", "path": "documentation/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/secret_value/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["other_settings", "header_options", "response_headers_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "other_settings.header_options.response_headers_to_add.secret_value for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings.header_options.response_headers_to_add.secret_value

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [other_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/)
- [other_settings.header_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/header_options/)
- [other_settings.header_options.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/)
- other_settings.header_options.response_headers_to_add.secret_value

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/secret_value/blindfold_secret_info/)
- [other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/secret_value/clear_secret_info/)
- [other_settings.header_options.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
