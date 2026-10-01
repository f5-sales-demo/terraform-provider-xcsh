---
page_title: "http_proxy.more_option.response_cookies_to_add.secret_value"
subcategory: ""
description: "http_proxy.more_option.response_cookies_to_add.secret_value for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2004, "body_sha256": "sha256:cada0d3bd2b1c98fb4a503a608d617fe4424c208c31ca5cb1a3c12dac85fbb10", "canonical_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value", "child_ids": ["xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_cookies_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:proxy:properties:http_proxy:more_option:response_cookies_to_add", "path": "docs/guides/data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_proxy", "more_option", "response_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/http_proxy/more_option/response_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_proxy.more_option.response_cookies_to_add.secret_value for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.response_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [http_proxy](data-sources--proxy--properties--http_proxy.md)
- [http_proxy.more_option](data-sources--proxy--properties--http_proxy--more_option.md)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md)
- http_proxy.more_option.response_cookies_to_add.secret_value

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

- [blindfold_secret_info](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md)
- [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md)
- [xcsh_proxy](../data-sources/proxy.md)
