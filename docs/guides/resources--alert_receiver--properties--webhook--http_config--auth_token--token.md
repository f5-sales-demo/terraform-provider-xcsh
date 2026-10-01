---
page_title: "webhook.http_config.auth_token.token"
subcategory: ""
description: "webhook.http_config.auth_token.token for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2130, "body_sha256": "sha256:a4c36809435e59bbb0dd8a092e14a692908aa7c762b0d32b7271ce876f3fb55a", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:blindfold_secret_info", "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token:clear_secret_info"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "path": "docs/guides/resources--alert_receiver--properties--webhook--http_config--auth_token--token.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "auth_token", "token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/auth_token/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.auth_token.token for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.auth_token.token

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- [webhook.http_config.auth_token](resources--alert_receiver--properties--webhook--http_config--auth_token.md)
- webhook.http_config.auth_token.token

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--alert_receiver--properties--webhook--http_config--auth_token--token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--properties--webhook--http_config--auth_token--token--clear_secret_info.md): complete subsection reference.

## Next pages

- [webhook.http_config.auth_token.token.blindfold_secret_info](resources--alert_receiver--properties--webhook--http_config--auth_token--token--blindfold_secret_info.md)
- [webhook.http_config.auth_token.token.clear_secret_info](resources--alert_receiver--properties--webhook--http_config--auth_token--token--clear_secret_info.md)
- [webhook.http_config.auth_token](resources--alert_receiver--properties--webhook--http_config--auth_token.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
