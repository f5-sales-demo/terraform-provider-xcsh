---
page_title: "webhook"
subcategory: ""
description: "webhook for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1084, "body_sha256": "sha256:fda7c5b9183ace76bc81410c404e5e05d19f64c1e1bc572db917c94c8e515e25", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "xcsh-docs:data-sources:alert_receiver:properties:webhook:url"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook", "parent_id": "xcsh-docs:data-sources:alert_receiver:reference", "path": "docs/guides/data-sources--alert_receiver--properties--webhook.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- webhook

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [http_config](data-sources--alert_receiver--properties--webhook--http_config.md): complete subsection reference.

- [url](data-sources--alert_receiver--properties--webhook--url.md): complete subsection reference.

## Next pages

- [webhook.http_config](data-sources--alert_receiver--properties--webhook--http_config.md)
- [webhook.url](data-sources--alert_receiver--properties--webhook--url.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
