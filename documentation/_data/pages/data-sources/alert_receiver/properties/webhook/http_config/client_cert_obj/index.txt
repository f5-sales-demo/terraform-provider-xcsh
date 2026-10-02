---
page_title: "webhook.http_config.client_cert_obj"
subcategory: ""
description: "Configuration for client certificate."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "webhook http config client cert obj"], "body_bytes": 1712, "body_sha256": "sha256:7e840c33a60eae1095ddfd2326c48d994d5843f2c220f6ec264b6564b626ad86", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "path": "documentation/data-sources/alert_receiver/properties/webhook/http_config/client_cert_obj/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2311000221310311-1210201200003322-2321232110300330-3023130122002031-3131003023220200-1203200302202301-1223013222002320-0211023111213222", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "client_cert_obj"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "use tls obj"], "anchor": "section", "description": "Reference to client certificate object.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj", "use_tls_obj"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/client_cert_obj/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for client certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.client_cert_obj

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/)
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

- [use_tls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/): complete subsection reference.

## Next pages

- [webhook.http_config.client_cert_obj.use_tls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
