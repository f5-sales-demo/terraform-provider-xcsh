---
page_title: "webhook.http_config.client_cert_obj"
subcategory: ""
description: "webhook.http_config.client_cert_obj for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1310, "body_sha256": "sha256:36564e5311aa0b75e16d85e4bd33aaaa85805e928bdf7fe67d20970fd3459a17", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "path": "docs/guides/data-sources--alert_receiver--properties--webhook--http_config--client_cert_obj.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "client_cert_obj"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/client_cert_obj/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.client_cert_obj for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.client_cert_obj

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [webhook](data-sources--alert_receiver--properties--webhook.md)
- [webhook.http_config](data-sources--alert_receiver--properties--webhook--http_config.md)
- webhook.http_config.client_cert_obj

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [use_tls_obj](data-sources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md): complete subsection reference.

## Next pages

- [webhook.http_config.client_cert_obj.use_tls_obj](data-sources--alert_receiver--properties--webhook--http_config--client_cert_obj--use_tls_obj.md)
- [webhook.http_config](data-sources--alert_receiver--properties--webhook--http_config.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
