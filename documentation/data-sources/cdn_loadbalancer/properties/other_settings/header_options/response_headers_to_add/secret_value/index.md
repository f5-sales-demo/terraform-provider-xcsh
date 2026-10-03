---
page_title: "other_settings.header_options.response_headers_to_add.secret_value"
subcategory: "Load Balancing"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["other settings header options response headers to add secret value"], "body_bytes": 2795, "body_sha256": "sha256:2f80b0a89e79415aec2628d9f2c548759efc41e49974e563738dcc830cc2657e", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add", "path": "documentation/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/secret_value/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3113030023011322-2102333220022101-1021022200110333-1213020300210232-2231022221320323-2302200032120332-3100231223231112-2221010210022312", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["other_settings", "header_options", "response_headers_to_add", "secret_value"], "schema_version": 1, "sections": [{"aliases": ["other settings header options response headers to add secret value blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "header_options", "response_headers_to_add", "secret_value", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["other settings header options response headers to add secret value clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add:secret_value:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "header_options", "response_headers_to_add", "secret_value", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/header_options/response_headers_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
