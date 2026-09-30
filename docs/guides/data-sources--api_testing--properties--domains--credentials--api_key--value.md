---
page_title: "domains.credentials.api_key.value"
subcategory: ""
description: "domains.credentials.api_key.value for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 1687, "body_sha256": "sha256:6649fb950c3c275e9b6b071f58dab8ba9984424d9cb548482ebc4546b1d69c8e", "canonical_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key:value", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key:value:blindfold_secret_info", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key:value:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key:value", "parent_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:api_key", "path": "docs/guides/data-sources--api_testing--properties--domains--credentials--api_key--value.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["domains", "credentials", "api_key", "value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/credentials/api_key/value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "domains.credentials.api_key.value for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# domains.credentials.api_key.value

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md)
- [Property reference](data-sources--api_testing--reference.md)
- [domains](data-sources--api_testing--properties--domains.md)
- [domains.credentials](data-sources--api_testing--properties--domains--credentials.md)
- [domains.credentials.api_key](data-sources--api_testing--properties--domains--credentials--api_key.md)
- domains.credentials.api_key.value

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

- [blindfold_secret_info](data-sources--api_testing--properties--domains--credentials--api_key--value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--properties--domains--credentials--api_key--value--clear_secret_info.md): complete subsection reference.

## Next pages

- [domains.credentials.api_key.value.blindfold_secret_info](data-sources--api_testing--properties--domains--credentials--api_key--value--blindfold_secret_info.md)
- [domains.credentials.api_key.value.clear_secret_info](data-sources--api_testing--properties--domains--credentials--api_key--value--clear_secret_info.md)
- [domains.credentials.api_key](data-sources--api_testing--properties--domains--credentials--api_key.md)
- [xcsh_api_testing](../data-sources/api_testing.md)
