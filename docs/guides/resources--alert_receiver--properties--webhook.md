---
page_title: "webhook"
subcategory: ""
description: "webhook for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1077, "body_sha256": "sha256:c798e39eb0c00f3db0bd45a523ac44ca0f1b13999ef5abeb24b88080ba88a3fb", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "xcsh-docs:resources:alert_receiver:properties:webhook:url"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook", "parent_id": "xcsh-docs:resources:alert_receiver:reference", "path": "docs/guides/resources--alert_receiver--properties--webhook.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# webhook

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- webhook

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Webhook configuration to send alert notifications.

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
webhook {
  # Configure direct properties listed below.
}
```

## Direct properties

- [http_config](resources--alert_receiver--properties--webhook--http_config.md): complete subsection reference.

- [url](resources--alert_receiver--properties--webhook--url.md): complete subsection reference.

## Next pages

- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- [webhook.url](resources--alert_receiver--properties--webhook--url.md)
- [Property reference](resources--alert_receiver--reference.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
