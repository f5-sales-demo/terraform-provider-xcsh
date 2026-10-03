---
page_title: "kafka_receiver.compression"
subcategory: ""
description: "Compression Type."
xcsh_docs: {"aliases": ["kafka receiver compression"], "body_bytes": 2911, "body_sha256": "sha256:6fb5ae588cf62361ab8d950c2956938576bb6f87d7e3184a3bbfae033fd47bea", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_default", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_gzip", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_none"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "path": "documentation/resources/global_log_receiver/properties/kafka_receiver/compression/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_none", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["kafka_receiver", "compression"], "schema_version": 1, "sections": [{"aliases": ["kafka receiver compression compression default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_default", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "compression", "compression_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["kafka receiver compression compression gzip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_gzip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "compression", "compression_gzip"], "syntax": "attribute", "type": "object"}, {"aliases": ["kafka receiver compression compression none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "compression", "compression_none"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/kafka_receiver/compression/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Compression Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/)
- kafka_receiver.compression

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

## Direct properties

- [compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/compression_default/): complete subsection reference.

- [compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/compression_gzip/): complete subsection reference.

- [compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/compression_none/): complete subsection reference.

## Next pages

- [kafka_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/compression_default/)
- [kafka_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/compression_gzip/)
- [kafka_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/compression_none/)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
