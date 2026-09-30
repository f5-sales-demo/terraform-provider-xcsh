---
page_title: "other_settings.header_options.request_headers_to_add.secret_value"
subcategory: "Load Balancing"
description: "other_settings.header_options.request_headers_to_add.secret_value for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2142, "body_sha256": "sha256:23971b873be4a172b23444336a2ac5fbfa36985cf87879ea7a4fdf83cd6d8b86", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:request_headers_to_add:secret_value", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:request_headers_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:request_headers_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:request_headers_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:request_headers_to_add", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--other_settings--header_options--request_headers_to_add--secret_value.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["other_settings", "header_options", "request_headers_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/header_options/request_headers_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "other_settings.header_options.request_headers_to_add.secret_value for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# other_settings.header_options.request_headers_to_add.secret_value

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [other_settings](data-sources--cdn_loadbalancer--properties--other_settings.md)
- [other_settings.header_options](data-sources--cdn_loadbalancer--properties--other_settings--header_options.md)
- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--properties--other_settings--header_options--request_headers_to_add.md)
- other_settings.header_options.request_headers_to_add.secret_value

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

- [blindfold_secret_info](data-sources--cdn_loadbalancer--properties--other_settings--header_options--request_headers_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--properties--other_settings--header_options--request_headers_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--cdn_loadbalancer--properties--other_settings--header_options--request_headers_to_add--secret_value--blindfold_secret_info.md)
- [other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info](data-sources--cdn_loadbalancer--properties--other_settings--header_options--request_headers_to_add--secret_value--clear_secret_info.md)
- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--properties--other_settings--header_options--request_headers_to_add.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
