---
page_title: "webhook"
subcategory: ""
description: "Webhook configuration to send alert notifications."
xcsh_docs: {"aliases": ["webhook"], "body_bytes": 1492, "body_sha256": "sha256:e993f4439020e220fd18883c36623286bb6983f07a8cdc2f373a1ef84c423a3f", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "xcsh-docs:data-sources:alert_receiver:properties:webhook:url"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook", "parent_id": "xcsh-docs:data-sources:alert_receiver:reference", "path": "documentation/data-sources/alert_receiver/properties/webhook/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0300223223302111-0101231033213323-0110013230113210-2010301321102001-2123131011023312-3331032020333110-1222231013311122-3133302001000303", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook"], "schema_version": 1, "sections": [{"aliases": ["webhook http config"], "anchor": "section", "description": "Configuration for HTTP endpoint.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["webhook url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:url", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "url"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Webhook configuration to send alert notifications.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
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

- [http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/): complete subsection reference.

- [url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/url/): complete subsection reference.

## Next pages

- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/)
- [webhook.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/url/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
