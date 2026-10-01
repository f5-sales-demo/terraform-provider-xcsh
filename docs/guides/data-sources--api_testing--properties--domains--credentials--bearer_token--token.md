---
page_title: "domains.credentials.bearer_token.token"
subcategory: ""
description: "domains.credentials.bearer_token.token for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 1846, "body_sha256": "sha256:c52a4610113395dd5c686e1fc937568b7b88239844ce6082878c5a77d14e4032", "canonical_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token:token", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token:token:blindfold_secret_info", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token:token:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token:token", "parent_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token", "path": "docs/guides/data-sources--api_testing--properties--domains--credentials--bearer_token--token.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "credentials", "bearer_token", "token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/credentials/bearer_token/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials.bearer_token.token for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.bearer_token.token

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md)
- [Property reference](data-sources--api_testing--reference.md)
- [domains](data-sources--api_testing--properties--domains.md)
- [domains.credentials](data-sources--api_testing--properties--domains--credentials.md)
- [domains.credentials.bearer_token](data-sources--api_testing--properties--domains--credentials--bearer_token.md)
- domains.credentials.bearer_token.token

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

- [blindfold_secret_info](data-sources--api_testing--properties--domains--credentials--bearer_token--token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--properties--domains--credentials--bearer_token--token--clear_secret_info.md): complete subsection reference.

## Next pages

- [domains.credentials.bearer_token.token.blindfold_secret_info](data-sources--api_testing--properties--domains--credentials--bearer_token--token--blindfold_secret_info.md)
- [domains.credentials.bearer_token.token.clear_secret_info](data-sources--api_testing--properties--domains--credentials--bearer_token--token--clear_secret_info.md)
- [domains.credentials.bearer_token](data-sources--api_testing--properties--domains--credentials--bearer_token.md)
- [xcsh_api_testing](../data-sources/api_testing.md)
