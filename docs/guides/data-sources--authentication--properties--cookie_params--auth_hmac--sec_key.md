---
page_title: "cookie_params.auth_hmac.sec_key"
subcategory: ""
description: "cookie_params.auth_hmac.sec_key for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 1708, "body_sha256": "sha256:2934d7fb413172b9ead2fde1ce3e75cc21ddb1f05796b83b67328b709d28dce9", "canonical_id": "xcsh-docs:data-sources:authentication:properties:cookie_params:auth_hmac:sec_key", "child_ids": ["xcsh-docs:data-sources:authentication:properties:cookie_params:auth_hmac:sec_key:blindfold_secret_info", "xcsh-docs:data-sources:authentication:properties:cookie_params:auth_hmac:sec_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:properties:cookie_params:auth_hmac:sec_key", "parent_id": "xcsh-docs:data-sources:authentication:properties:cookie_params:auth_hmac", "path": "docs/guides/data-sources--authentication--properties--cookie_params--auth_hmac--sec_key.md", "provider_name": "authentication", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cookie_params", "auth_hmac", "sec_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/properties/cookie_params/auth_hmac/sec_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cookie_params.auth_hmac.sec_key for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_params.auth_hmac.sec_key

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md)
- [Property reference](data-sources--authentication--reference.md)
- [cookie_params](data-sources--authentication--properties--cookie_params.md)
- [cookie_params.auth_hmac](data-sources--authentication--properties--cookie_params--auth_hmac.md)
- cookie_params.auth_hmac.sec_key

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

- [blindfold_secret_info](data-sources--authentication--properties--cookie_params--auth_hmac--sec_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--authentication--properties--cookie_params--auth_hmac--sec_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [cookie_params.auth_hmac.sec_key.blindfold_secret_info](data-sources--authentication--properties--cookie_params--auth_hmac--sec_key--blindfold_secret_info.md)
- [cookie_params.auth_hmac.sec_key.clear_secret_info](data-sources--authentication--properties--cookie_params--auth_hmac--sec_key--clear_secret_info.md)
- [cookie_params.auth_hmac](data-sources--authentication--properties--cookie_params--auth_hmac.md)
- [xcsh_authentication](../data-sources/authentication.md)
