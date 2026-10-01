---
page_title: "webhook.url"
subcategory: ""
description: "webhook.url for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1673, "body_sha256": "sha256:a10b8391734638d1ac2e0bc0d879da0939f58e53d66d9cce5ff3c778c22d6740", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:url", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:url:blindfold_secret_info", "xcsh-docs:resources:alert_receiver:properties:webhook:url:clear_secret_info"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:url", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook", "path": "docs/guides/resources--alert_receiver--properties--webhook--url.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.url for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.url

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- webhook.url

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
url {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--alert_receiver--properties--webhook--url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--properties--webhook--url--clear_secret_info.md): complete subsection reference.

## Next pages

- [webhook.url.blindfold_secret_info](resources--alert_receiver--properties--webhook--url--blindfold_secret_info.md)
- [webhook.url.clear_secret_info](resources--alert_receiver--properties--webhook--url--clear_secret_info.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
