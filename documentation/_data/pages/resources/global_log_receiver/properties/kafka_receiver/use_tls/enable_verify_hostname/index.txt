---
page_title: "kafka_receiver.use_tls.enable_verify_hostname"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["kafka receiver use tls enable verify hostname"], "body_bytes": 1506, "body_sha256": "sha256:4610025a18626215f913da0f06dd927660c82d67804bbc66d148097c3a41cead", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_hostname", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls", "path": "documentation/resources/global_log_receiver/properties/kafka_receiver/use_tls/enable_verify_hostname/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0202222011220001-3010030131003222-2221003210311023-1330030111201211-1011121213132032-2202332321012002-3101201213231331-1211200123312310", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kafka_receiver", "use_tls", "enable_verify_hostname"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/kafka_receiver/use_tls/enable_verify_hostname/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_receiver.use_tls.enable_verify_hostname

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/)
- [kafka_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/)
- kafka_receiver.use_tls.enable_verify_hostname

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
enable_verify_hostname = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [kafka_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
