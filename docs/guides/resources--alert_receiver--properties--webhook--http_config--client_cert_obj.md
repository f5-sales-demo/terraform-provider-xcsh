---
page_title: "webhook.http_config.client_cert_obj"
subcategory: ""
description: "webhook.http_config.client_cert_obj for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1311, "body_sha256": "sha256:89833a9c3d5cf0bea4bd2730b25dfd521c7f1a5548756f3cce81d7c91c772acc", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "path": "docs/guides/resources--alert_receiver--properties--webhook--http_config--client_cert_obj.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "client_cert_obj"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.client_cert_obj for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# webhook.http_config.client_cert_obj

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- webhook.http_config.client_cert_obj

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Client Certificate Object. Configuration for client certificate.

Upstream description:

Configuration for client certificate.

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
client_cert_obj {
  # Configure direct properties listed below.
}
```

## Direct properties

- [use_tls_obj](resources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md): complete subsection reference.

## Next pages

- [webhook.http_config.client_cert_obj.use_tls_obj](resources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md)
- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
