---
page_title: "webhook.http_config.no_authorization"
subcategory: ""
description: "webhook.http_config.no_authorization for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1137, "body_sha256": "sha256:1b437f2b99dba3cac8514cce619a7a9f180ad1f080ba1ec017db9dff9f40224b", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_authorization", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:no_authorization", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "path": "docs/guides/resources--alert_receiver--properties--webhook--http_config--no_authorization.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "no_authorization"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/no_authorization/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.no_authorization for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.no_authorization

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- webhook.http_config.no_authorization

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no authorization.

Upstream description:

This can be used for messages where no values are needed.

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
no_authorization = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
