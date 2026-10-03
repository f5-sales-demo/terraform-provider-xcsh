---
page_title: "http_receiver.compression.compression_gzip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["http receiver compression compression gzip"], "body_bytes": 1504, "body_sha256": "sha256:fe1ff11b19d1d0fa353685b0d02c51bc28237751394d3c503b20591fe3a07aa3", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression:compression_gzip", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:http_receiver:compression", "path": "documentation/resources/global_log_receiver/properties/http_receiver/compression/compression_gzip/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0213103022202120-0223301323013023-2102120011131232-1000330133303213-1123032122212312-1222122300200201-0013211023200232-3030311212132111", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_receiver", "compression", "compression_gzip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/http_receiver/compression/compression_gzip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.compression.compression_gzip

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/)
- [http_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/compression/)
- http_receiver.compression.compression_gzip

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
compression_gzip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/compression/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
