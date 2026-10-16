---
page_title: "webhook.http_config.use_tls.disable_sni"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["webhook http config use tls disable sni"], "body_bytes": 1317, "body_sha256": "sha256:3ea3aec668354fbffdeaa335aad8bb87a2c441af9009ae644fa411d00be9f8ea", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:disable_sni", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/use_tls/disable_sni/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0320030233030310-1013031020112330-1322120301100231-2033031312001131-2123300130000202-3123120132213321-0122310003223200-3102022222220203", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "disable_sni"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/use_tls/disable_sni/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.use_tls.disable_sni

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- [webhook.http_config.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/)
- webhook.http_config.use_tls.disable_sni

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

Additional upstream details:

This can be used for messages where no values are needed.

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
disable_sni = {}
```

This is an empty object or choice marker. It has no direct properties.
