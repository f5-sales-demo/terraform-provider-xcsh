---
page_title: "azure_receiver.compression"
subcategory: ""
description: "Compression Type."
xcsh_docs: {"aliases": ["azure receiver compression"], "body_bytes": 2911, "body_sha256": "sha256:dfcec2e140d363afb52d5638ecf844449990c334732f5dc242fb852b4970a211", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_default", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_gzip", "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_none"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver", "path": "documentation/resources/global_log_receiver/properties/azure_receiver/compression/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2212231022022020-2232331232100113-0300233311201203-3300213121033120-1133003222102120-1310021311213323-1001131113123113-1002030321000210", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_none", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_receiver", "compression"], "schema_version": 1, "sections": [{"aliases": ["compression default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_default", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "compression", "compression_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["compression gzip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_gzip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "compression", "compression_gzip"], "syntax": "attribute", "type": "object"}, {"aliases": ["compression none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:compression:compression_none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "compression", "compression_none"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/compression/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Compression Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/)
- azure_receiver.compression

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

- [compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/compression_default/): complete subsection reference.

- [compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/compression_gzip/): complete subsection reference.

- [compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/compression_none/): complete subsection reference.

## Next pages

- [azure_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/compression_default/)
- [azure_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/compression_gzip/)
- [azure_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/compression/compression_none/)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
