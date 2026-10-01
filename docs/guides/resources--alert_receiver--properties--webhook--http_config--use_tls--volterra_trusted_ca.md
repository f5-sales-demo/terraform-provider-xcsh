---
page_title: "webhook.http_config.use_tls.volterra_trusted_ca"
subcategory: ""
description: "webhook.http_config.use_tls.volterra_trusted_ca for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1287, "body_sha256": "sha256:cebb8d45b01bf1e67359f0f49f7fb9763b6d7661a654ab9387974a323fa826d2", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:volterra_trusted_ca", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:volterra_trusted_ca", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "path": "docs/guides/resources--alert_receiver--properties--webhook--http_config--use_tls--volterra_trusted_ca.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "volterra_trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/use_tls/volterra_trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.use_tls.volterra_trusted_ca for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.use_tls.volterra_trusted_ca

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- [webhook.http_config.use_tls](resources--alert_receiver--properties--webhook--http_config--use_tls.md)
- webhook.http_config.use_tls.volterra_trusted_ca

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [webhook.http_config.use_tls](resources--alert_receiver--properties--webhook--http_config--use_tls.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
