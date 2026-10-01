---
page_title: "custom_proxy.password"
subcategory: ""
description: "custom_proxy.password for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1544, "body_sha256": "sha256:c45abe811187fbccc3f68b98d093c384c501afb55463214a9a57ffda4b0e163b", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:custom_proxy:password", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:custom_proxy:password:blindfold_secret_info", "xcsh-docs:data-sources:securemesh_site_v2:properties:custom_proxy:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:custom_proxy:password", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:custom_proxy", "path": "docs/guides/data-sources--securemesh_site_v2--properties--custom_proxy--password.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_proxy", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/custom_proxy/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_proxy.password for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_proxy.password

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [custom_proxy](data-sources--securemesh_site_v2--properties--custom_proxy.md)
- custom_proxy.password

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

- [blindfold_secret_info](data-sources--securemesh_site_v2--properties--custom_proxy--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--securemesh_site_v2--properties--custom_proxy--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [custom_proxy.password.blindfold_secret_info](data-sources--securemesh_site_v2--properties--custom_proxy--password--blindfold_secret_info.md)
- [custom_proxy.password.clear_secret_info](data-sources--securemesh_site_v2--properties--custom_proxy--password--clear_secret_info.md)
- [custom_proxy](data-sources--securemesh_site_v2--properties--custom_proxy.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
