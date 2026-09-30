---
page_title: "webhook.http_config.auth_token.token"
subcategory: ""
description: "webhook.http_config.auth_token.token for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1762, "body_sha256": "sha256:28ee185c8c2086a4e1fd1e1556411fd20746175346185bc15684c8d77b6d7c2b", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token:token", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token:token:blindfold_secret_info", "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token:token:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token:token", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:auth_token", "path": "docs/guides/data-sources--alert_receiver--properties--webhook--http_config--auth_token--token.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "auth_token", "token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/auth_token/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.auth_token.token for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# webhook.http_config.auth_token.token

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [webhook](data-sources--alert_receiver--properties--webhook.md)
- [webhook.http_config](data-sources--alert_receiver--properties--webhook--http_config.md)
- [webhook.http_config.auth_token](data-sources--alert_receiver--properties--webhook--http_config--auth_token.md)
- webhook.http_config.auth_token.token

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

- [blindfold_secret_info](data-sources--alert_receiver--properties--webhook--http_config--auth_token--token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--properties--webhook--http_config--auth_token--token--clear_secret_info.md): complete subsection reference.

## Next pages

- [webhook.http_config.auth_token.token.blindfold_secret_info](data-sources--alert_receiver--properties--webhook--http_config--auth_token--token--blindfold_secret_info.md)
- [webhook.http_config.auth_token.token.clear_secret_info](data-sources--alert_receiver--properties--webhook--http_config--auth_token--token--clear_secret_info.md)
- [webhook.http_config.auth_token](data-sources--alert_receiver--properties--webhook--http_config--auth_token.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
