---
page_title: "webhook.http_config.use_tls.use_server_verification"
subcategory: ""
description: "webhook.http_config.use_tls.use_server_verification for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1500, "body_sha256": "sha256:7d16ff8156cb680ec7728dc5161d9e52c4bd4d79d3a9ac57ce207a2b6f831748", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls", "path": "docs/guides/data-sources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.http_config.use_tls.use_server_verification for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.use_tls.use_server_verification

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [webhook](data-sources--alert_receiver--properties--webhook.md)
- [webhook.http_config](data-sources--alert_receiver--properties--webhook--http_config.md)
- [webhook.http_config.use_tls](data-sources--alert_receiver--properties--webhook--http_config--use_tls.md)
- webhook.http_config.use_tls.use_server_verification

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

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

- [ca_cert_obj](data-sources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj.md): complete subsection reference.

## Next pages

- [webhook.http_config.use_tls.use_server_verification.ca_cert_obj](data-sources--alert_receiver--properties--webhook--http_config--use_tls--use_server_verification--ca_cert_obj.md)
- [webhook.http_config.use_tls](data-sources--alert_receiver--properties--webhook--http_config--use_tls.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
