---
page_title: "request_logs"
subcategory: ""
description: "Configuration for request logs with sampling choice. Allows selection between sampled (default) or unsampled (full) request logs."
xcsh_docs: {"aliases": ["request logs"], "body_bytes": 2156, "body_sha256": "sha256:1fce68bfb1c5d02dd6669a2b2fc8662a4fe7efeb932b55373d5ea33adad40b03", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:request_logs:sampled", "xcsh-docs:resources:global_log_receiver:properties:request_logs:unsampled"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:request_logs", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/request_logs/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1232021303231213-0303210110003213-3111010032312021-1233200302331232-1321133012013331-0030313121010031-3321002030223202-3213303122202211", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "request_logs:ConflictingObjectAttributes:sampled,unsampled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:sampled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_logs:ConflictingObjectAttributes:sampled,unsampled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:unsampled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["request_logs"], "schema_version": 1, "sections": [{"aliases": ["sampled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:sampled", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_logs", "sampled"], "syntax": "attribute", "type": "object"}, {"aliases": ["unsampled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:request_logs:unsampled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_logs", "unsampled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/request_logs/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for request logs with sampling choice. Allows selection between sampled (default) or unsampled (full) request logs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_logs

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- request_logs

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Upstream description:

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("sampled",
    "unsampled")}
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
  "x-ves-oneof-field-sampling_choice": "[\"sampled\",\"unsampled\"]"
}
```

Terraform syntax:

```terraform
request_logs {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/sampled/): complete subsection reference.

- [unsampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/unsampled/): complete subsection reference.

## Next pages

- [request_logs.sampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/sampled/)
- [request_logs.unsampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/unsampled/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
