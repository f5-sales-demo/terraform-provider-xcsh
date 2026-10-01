---
page_title: "webhook.http_config.client_cert_obj"
subcategory: ""
description: "webhook.http_config.client_cert_obj for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1410, "body_sha256": "sha256:3a186a037e2825feff0a169cdd6975f16c78c2e1b7dd085526ea7f176885db30", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "path": "docs/guides/resources--alert_receiver--properties--webhook--http_config--client_cert_obj.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "client_cert_obj"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.client_cert_obj for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
