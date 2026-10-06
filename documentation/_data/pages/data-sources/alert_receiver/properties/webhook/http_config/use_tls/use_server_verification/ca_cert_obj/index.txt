---
page_title: "webhook.http_config.use_tls.use_server_verification.ca_cert_obj"
subcategory: ""
description: "Configuration for CA certificate."
xcsh_docs: {"aliases": ["webhook http config use tls use server verification ca cert obj"], "body_bytes": 1671, "body_sha256": "sha256:5d454c94ed6a80894ad4f7881a10f4e7c0509255adfbea037e6803d319200a35", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj:trusted_ca"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification", "path": "documentation/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3213001032311223-1032133001311020-1320031222113033-3213112031013112-0210122320133311-3321123033023310-1230310300323333-0033010033103333", "registry_path": "docs/guides/data-sources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification", "ca_cert_obj"], "schema_version": 1, "sections": [{"aliases": ["webhook http config use tls use server verification ca cert obj trusted ca"], "anchor": "section", "description": "Reference to client certificate object.", "document_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:http_config:use_tls:use_server_verification:ca_cert_obj:trusted_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["webhook", "http_config", "use_tls", "use_server_verification", "ca_cert_obj", "trusted_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/http_config/use_tls/use_server_verification/ca_cert_obj/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration for CA certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
