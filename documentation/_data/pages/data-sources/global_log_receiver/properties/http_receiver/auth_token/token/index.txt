---
page_title: "http_receiver.auth_token.token"
subcategory: ""
description: "http_receiver.auth_token.token for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2163, "body_sha256": "sha256:10e9df90cc1f13ebdc3bc953c33a5b2494d4a635f152c53fa750ab0dd6f5c7ca", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token:token:blindfold_secret_info", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token:token:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token:token", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:auth_token", "path": "documentation/data-sources/global_log_receiver/properties/http_receiver/auth_token/token/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["http_receiver", "auth_token", "token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/http_receiver/auth_token/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_receiver.auth_token.token for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# http_receiver.auth_token.token

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/)
- [http_receiver.auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_token/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_token/token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_token/token/clear_secret_info/): complete subsection reference.

## Next pages

- [http_receiver.auth_token.token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_token/token/blindfold_secret_info/)
- [http_receiver.auth_token.token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_token/token/clear_secret_info/)
- [http_receiver.auth_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/auth_token/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
