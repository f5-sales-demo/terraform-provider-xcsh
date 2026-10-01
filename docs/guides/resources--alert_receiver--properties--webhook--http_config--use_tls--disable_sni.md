---
page_title: "webhook.http_config.use_tls.disable_sni"
subcategory: ""
description: "webhook.http_config.use_tls.disable_sni for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1255, "body_sha256": "sha256:4ef1f82a7a3149afb82a912f0e7d126beefd1194327c5e7e6f864b1ecae26166", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:disable_sni", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:disable_sni", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "path": "docs/guides/resources--alert_receiver--properties--webhook--http_config--use_tls--disable_sni.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "disable_sni"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/use_tls/disable_sni/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.use_tls.disable_sni for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.use_tls.disable_sni

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- [webhook.http_config.use_tls](resources--alert_receiver--properties--webhook--http_config--use_tls.md)
- webhook.http_config.use_tls.disable_sni

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [webhook.http_config.use_tls](resources--alert_receiver--properties--webhook--http_config--use_tls.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
