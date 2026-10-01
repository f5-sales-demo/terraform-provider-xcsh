---
page_title: "webhook.http_config.auth_token"
subcategory: ""
description: "webhook.http_config.auth_token for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1333, "body_sha256": "sha256:785aeb3aa8e8c1fb29335c181c27b784c1ed8168838f71be428f6c928d840d27", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token:token"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:auth_token", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "path": "docs/guides/resources--alert_receiver--properties--webhook--http_config--auth_token.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "auth_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/auth_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.auth_token for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.auth_token

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- webhook.http_config.auth_token

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Access Token. Authentication Token for access.

Upstream description:

Authentication Token for access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
auth_token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [token](resources--alert_receiver--properties--webhook--http_config--auth_token--token.md): complete subsection reference.

## Next pages

- [webhook.http_config.auth_token.token](resources--alert_receiver--properties--webhook--http_config--auth_token--token.md)
- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
