---
page_title: "syslog.tls_server.mtls_disabled"
subcategory: "Monitoring"
description: "Selects disabled mTLS for the syslog TLS server."
xcsh_docs: {"aliases": ["disable syslog mtls", "mtls disabled", "syslog tls server mtls disabled"], "body_bytes": 1377, "body_sha256": "sha256:dfd15615c45c92565e2614d7ee2a38e76a45f2a31a4eed4ae234c4b2fee9dae0", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule", "reviewed-summary"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_disabled", "parent_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "path": "documentation/resources/log_receiver/properties/syslog/tls_server/mtls_disabled/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2131320132203000-1012003011033012-0331200120113131-2313003030011020-2302030303123211-3020202232022032-1111333322210032-2113333333213310", "registry_path": "docs/guides/resources--log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["syslog", "tls_server", "mtls_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/tls_server/mtls_disabled/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Selects disabled mTLS for the syslog TLS server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog.tls_server.mtls_disabled

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/)
- [syslog](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/)
- [syslog.tls_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/)
- syslog.tls_server.mtls_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
mtls_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [syslog.tls_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/)
- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
