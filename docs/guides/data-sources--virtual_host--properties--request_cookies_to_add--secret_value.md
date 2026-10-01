---
page_title: "request_cookies_to_add.secret_value"
subcategory: ""
description: "request_cookies_to_add.secret_value for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1630, "body_sha256": "sha256:84a4843b11a9e2182a6b087e240d78dda8de6244e20b5b0b0d6a5f723b343cca", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:request_cookies_to_add:secret_value", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:request_cookies_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:virtual_host:properties:request_cookies_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:request_cookies_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:request_cookies_to_add", "path": "docs/guides/data-sources--virtual_host--properties--request_cookies_to_add--secret_value.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["request_cookies_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/request_cookies_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_cookies_to_add.secret_value for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_cookies_to_add.secret_value

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- [request_cookies_to_add](data-sources--virtual_host--properties--request_cookies_to_add.md)
- request_cookies_to_add.secret_value

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

- [blindfold_secret_info](data-sources--virtual_host--properties--request_cookies_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--properties--request_cookies_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--virtual_host--properties--request_cookies_to_add--secret_value--blindfold_secret_info.md)
- [request_cookies_to_add.secret_value.clear_secret_info](data-sources--virtual_host--properties--request_cookies_to_add--secret_value--clear_secret_info.md)
- [request_cookies_to_add](data-sources--virtual_host--properties--request_cookies_to_add.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
