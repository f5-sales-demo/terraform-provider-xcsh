---
page_title: "webhook.http_config.use_tls.use_server_verification"
subcategory: ""
description: "Upstream TLS Validation Context."
xcsh_docs: {"aliases": ["webhook http config use tls use server verification"], "body_bytes": 1443, "body_sha256": "sha256:064399f19b3999729f5a44275777686b2b7f864323dc135bce2a20a0a2213598", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls", "path": "documentation/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0001031210332313-3310013011030213-1220232223110312-0030002100112211-2132023301203221-2030213200231232-2232011311010032-3131031103222113", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification"], "schema_version": 1, "sections": [{"aliases": ["webhook http config use tls use server verification ca cert obj"], "anchor": "section", "description": "Configuration for CA certificate.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification", "ca_cert_obj"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Upstream TLS Validation Context.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.use_tls.use_server_verification

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/)
- [webhook.http_config.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/use_tls/)
- webhook.http_config.use_tls.use_server_verification

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for use server verification.

Additional upstream details:

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

- [ca_cert_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/): complete subsection reference.
