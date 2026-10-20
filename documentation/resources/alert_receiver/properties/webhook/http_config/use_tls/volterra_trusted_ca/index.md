---
page_title: "webhook.http_config.use_tls.volterra_trusted_ca"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["webhook http config use tls volterra trusted ca"], "body_bytes": 1349, "body_sha256": "sha256:ae141b844a809d7879929a4a9bf0cbf74465b41b61be4a952bc23562d5ac1fbd", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls:volterra_trusted_ca", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:use_tls", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/use_tls/volterra_trusted_ca/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3231300230002132-1122230313302133-3111223223032033-1022113031333210-0310103333120312-2101221001200012-0331102202123022-3030211012211102", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "use_tls", "volterra_trusted_ca"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/use_tls/volterra_trusted_ca/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.use_tls.volterra_trusted_ca

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- [webhook.http_config.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/use_tls/)
- webhook.http_config.use_tls.volterra_trusted_ca

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.
