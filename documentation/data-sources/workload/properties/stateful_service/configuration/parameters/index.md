---
page_title: "stateful_service.configuration.parameters"
subcategory: "Container"
description: "Parameters for the workload."
xcsh_docs: {"aliases": ["stateful service configuration parameters"], "body_bytes": 2641, "body_sha256": "sha256:3b3ed283cb194ec0bc6c54cc0006e5bcd6680e777ae9dd988d01aa6424bea54e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:env_var", "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:file"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration", "path": "documentation/data-sources/workload/properties/stateful_service/configuration/parameters/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3122320321220010-2102312233010312-0232123013221122-2001223312112201-2022020012221002-0030032122333331-3230312102013102-2113113003030230", "registry_path": "docs/guides/data-sources--workload--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "configuration", "parameters"], "schema_version": 1, "sections": [{"aliases": ["stateful service configuration parameters env var"], "anchor": "section", "description": "Environment Variable.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:env_var", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "configuration", "parameters", "env_var"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service configuration parameters file"], "anchor": "section", "description": "Configuration File for the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "configuration", "parameters", "file"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/configuration/parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.configuration.parameters

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/)
- stateful_service.configuration.parameters

<a id="section"></a>

Type: `"list"`. Computed.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/parameters/env_var/): complete subsection reference.

- [file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/parameters/file/): complete subsection reference.

## Next pages

- [stateful_service.configuration.parameters.env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/parameters/env_var/)
- [stateful_service.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/parameters/file/)
- [stateful_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/configuration/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
