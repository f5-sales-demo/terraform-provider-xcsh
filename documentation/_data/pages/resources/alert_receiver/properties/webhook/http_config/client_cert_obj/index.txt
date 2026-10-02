---
page_title: "webhook.http_config.client_cert_obj"
subcategory: ""
description: "Configuration for client certificate."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "webhook http config client cert obj"], "body_bytes": 1812, "body_sha256": "sha256:52c53ec4ba5e81155ea60e4e8128cb873fae50c79ce4085aa4ff918bade970e9", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0001133132213100-1310233321222231-3033101332123331-2331133220012333-1233030013101022-0021213321003121-2121331002230132-3011020232321100", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "client_cert_obj"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "use tls obj"], "anchor": "section", "description": "Reference to client certificate object.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:client_cert_obj:use_tls_obj", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["webhook", "http_config", "client_cert_obj", "use_tls_obj"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration for client certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.client_cert_obj

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- webhook.http_config.client_cert_obj

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
client_cert_obj {
  # Configure direct properties listed below.
}
```

## Direct properties

- [use_tls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/): complete subsection reference.

## Next pages

- [webhook.http_config.client_cert_obj.use_tls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/client_cert_obj/use_tls_obj/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
