---
page_title: "http_receiver.auth_token.token"
subcategory: ""
description: "http_receiver.auth_token.token for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1764, "body_sha256": "sha256:c4377bdc6d8113b751eac0ec836fdabf2d9fa72a8d6f34e62cdbea8df72099c8", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token:token", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token:token:blindfold_secret_info", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token:token:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token:token", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token", "path": "docs/guides/data-sources--global_log_receiver--properties--http_receiver--auth_token--token.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_receiver", "auth_token", "token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/http_receiver/auth_token/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_receiver.auth_token.token for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.auth_token.token

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [http_receiver](data-sources--global_log_receiver--properties--http_receiver.md)
- [http_receiver.auth_token](data-sources--global_log_receiver--properties--http_receiver--auth_token.md)
- http_receiver.auth_token.token

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

- [blindfold_secret_info](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--clear_secret_info.md): complete subsection reference.

## Next pages

- [http_receiver.auth_token.token.blindfold_secret_info](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--blindfold_secret_info.md)
- [http_receiver.auth_token.token.clear_secret_info](data-sources--global_log_receiver--properties--http_receiver--auth_token--token--clear_secret_info.md)
- [http_receiver.auth_token](data-sources--global_log_receiver--properties--http_receiver--auth_token.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
