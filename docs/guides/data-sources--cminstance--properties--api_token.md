---
page_title: "api_token"
subcategory: ""
description: "api_token for xcsh_cminstance."
xcsh_docs: {"aliases": [], "body_bytes": 1275, "body_sha256": "sha256:2aaa96a621fb107c3eb5650c8d53aa3fbdf2faa370c85276820865931db4f28a", "canonical_id": "xcsh-docs:data-sources:cminstance:properties:api_token", "child_ids": ["xcsh-docs:data-sources:cminstance:properties:api_token:blindfold_secret_info", "xcsh-docs:data-sources:cminstance:properties:api_token:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cminstance:properties:api_token", "parent_id": "xcsh-docs:data-sources:cminstance:reference", "path": "docs/guides/data-sources--cminstance--properties--api_token.md", "provider_name": "cminstance", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cminstance/properties/api_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_token for xcsh_cminstance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cminstanceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_token

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md)
- [Property reference](data-sources--cminstance--reference.md)
- api_token

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

- [blindfold_secret_info](data-sources--cminstance--properties--api_token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--cminstance--properties--api_token--clear_secret_info.md): complete subsection reference.

## Next pages

- [api_token.blindfold_secret_info](data-sources--cminstance--properties--api_token--blindfold_secret_info.md)
- [api_token.clear_secret_info](data-sources--cminstance--properties--api_token--clear_secret_info.md)
- [Property reference](data-sources--cminstance--reference.md)
- [xcsh_cminstance](../data-sources/cminstance.md)
