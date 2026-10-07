---
page_title: "webhook.http_config.use_tls.use_server_verification.ca_cert_obj"
subcategory: ""
description: "Configuration for CA certificate."
xcsh_docs: {"aliases": ["webhook http config use tls use server verification ca cert obj"], "body_bytes": 1671, "body_sha256": "sha256:5d454c94ed6a80894ad4f7881a10f4e7c0509255adfbea037e6803d319200a35", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj:trusted_ca"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "path": "documentation/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3213001032311223-1032133001311020-1320031222113033-3213112031013112-0210122320133311-3321123033023310-1230310300323333-0033010033103333", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification", "ca_cert_obj"], "schema_version": 1, "sections": [{"aliases": ["webhook http config use tls use server verification ca cert obj trusted ca"], "anchor": "section", "description": "Reference to client certificate object.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj:trusted_ca", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification", "ca_cert_obj", "trusted_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Configuration for CA certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.use_tls.use_server_verification.ca_cert_obj

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/)
- [webhook.http_config.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/use_tls/)
- [webhook.http_config.use_tls.use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/)
- webhook.http_config.use_tls.use_server_verification.ca_cert_obj

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for ca cert obj.

Additional upstream details:

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

## Direct properties

- [trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/trusted_ca/): complete subsection reference.
