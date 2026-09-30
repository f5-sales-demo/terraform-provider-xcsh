---
page_title: "webhook.http_config.use_tls.use_server_verification.ca_cert_obj"
subcategory: ""
description: "webhook.http_config.use_tls.use_server_verification.ca_cert_obj for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1744, "body_sha256": "sha256:7f3086a320943a2c7ac2a663246a908e1659064ae0b69b6b2bbcf9eb8ec2fafc", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj:trusted_ca"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "path": "docs/guides/resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification", "ca_cert_obj"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.use_tls.use_server_verification.ca_cert_obj for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# webhook.http_config.use_tls.use_server_verification.ca_cert_obj

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [webhook](resources--alert_receiver--properties--webhook.md)
- [webhook.http_config](resources--alert_receiver--properties--webhook--http_config.md)
- [webhook.http_config.use_tls](resources--alert_receiver--properties--webhook--http_config--use_tls.md)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification.md)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ca cert obj.

Upstream description:

Configuration for CA certificate.

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
ca_cert_obj {
  # Configure direct properties listed below.
}
```

## Direct properties

- [trusted_ca](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca.md): complete subsection reference.

## Next pages

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj.trusted_ca](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj--trusted_ca.md)
- [webhook.http_config.use_tls.use_server_verification](resources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
