---
page_title: "webhook.http_config.client_cert_obj"
subcategory: ""
description: "Configuration for client certificate."
xcsh_docs: {"aliases": ["webhook http config client cert obj"], "body_bytes": 1186, "body_sha256": "sha256:dff12faf51411e2d9a4a75a534524b2ffc18980828141bd34b23d832d53e4547", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config", "path": "documentation/data-sources/alert_receiver/properties/webhook/http_config/client_cert_obj/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-2311000221310311-1210201200003322-2321232110300330-3023130122002031-3131003023220200-1203200302202301-1223013222002320-0211023111213222", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "client_cert_obj"], "schema_version": 1, "sections": [{"aliases": ["webhook http config client cert obj use tls obj"], "anchor": "section", "description": "Reference to client certificate object.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj", "use_tls_obj"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/client_cert_obj/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration for client certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
