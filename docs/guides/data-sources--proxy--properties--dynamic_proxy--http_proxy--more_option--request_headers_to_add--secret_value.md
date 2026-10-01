---
page_title: "dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value"
subcategory: ""
description: "dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2293, "body_sha256": "sha256:63145c0a61a8d6993ea754299e1db06a8f010ec21c74a64836008563571dbe33", "canonical_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add:secret_value", "child_ids": ["xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add", "path": "docs/guides/data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "request_headers_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_headers_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.http_proxy](data-sources--proxy--properties--dynamic_proxy--http_proxy.md)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add.md)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value

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

- [blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add.md)
- [xcsh_proxy](../data-sources/proxy.md)
