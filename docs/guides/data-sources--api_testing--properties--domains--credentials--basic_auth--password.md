---
page_title: "domains.credentials.basic_auth.password"
subcategory: ""
description: "domains.credentials.basic_auth.password for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 1846, "body_sha256": "sha256:85107b477c932c92a7942fe5bbe8755a807f686c9bf06f90dd5dfe600b0c7c9d", "canonical_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password:blindfold_secret_info", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password", "parent_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth", "path": "docs/guides/data-sources--api_testing--properties--domains--credentials--basic_auth--password.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "credentials", "basic_auth", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/credentials/basic_auth/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials.basic_auth.password for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.basic_auth.password

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md)
- [Property reference](data-sources--api_testing--reference.md)
- [domains](data-sources--api_testing--properties--domains.md)
- [domains.credentials](data-sources--api_testing--properties--domains--credentials.md)
- [domains.credentials.basic_auth](data-sources--api_testing--properties--domains--credentials--basic_auth.md)
- domains.credentials.basic_auth.password

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

- [blindfold_secret_info](data-sources--api_testing--properties--domains--credentials--basic_auth--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--properties--domains--credentials--basic_auth--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [domains.credentials.basic_auth.password.blindfold_secret_info](data-sources--api_testing--properties--domains--credentials--basic_auth--password--blindfold_secret_info.md)
- [domains.credentials.basic_auth.password.clear_secret_info](data-sources--api_testing--properties--domains--credentials--basic_auth--password--clear_secret_info.md)
- [domains.credentials.basic_auth](data-sources--api_testing--properties--domains--credentials--basic_auth.md)
- [xcsh_api_testing](../data-sources/api_testing.md)
