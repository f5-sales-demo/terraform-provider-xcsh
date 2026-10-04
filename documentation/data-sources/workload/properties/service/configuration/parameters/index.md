---
page_title: "service.configuration.parameters"
subcategory: "Container"
description: "Parameters for the workload."
xcsh_docs: {"aliases": ["service configuration parameters"], "body_bytes": 2515, "body_sha256": "sha256:45013284b848ff9e8704beeefc9c7a1abac8897e60fc750efa36ddc2354e18cc", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:configuration:parameters:env_var", "xcsh-docs:data-sources:workload:properties:service:configuration:parameters:file"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:configuration:parameters", "parent_id": "xcsh-docs:data-sources:workload:properties:service:configuration", "path": "documentation/data-sources/workload/properties/service/configuration/parameters/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0012011320103113-3033302112011313-2110032033231102-1120002333001312-0230133122323022-2322200011110000-1310321323010332-0020233203221233", "registry_path": "docs/guides/data-sources--workload--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "configuration", "parameters"], "schema_version": 1, "sections": [{"aliases": ["service configuration parameters env var"], "anchor": "section", "description": "Environment Variable.", "document_id": "xcsh-docs:data-sources:workload:properties:service:configuration:parameters:env_var", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "configuration", "parameters", "env_var"], "syntax": "attribute", "type": "object"}, {"aliases": ["service configuration parameters file"], "anchor": "section", "description": "Configuration File for the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:service:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "configuration", "parameters", "file"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/configuration/parameters/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Parameters for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.configuration.parameters

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/configuration/)
- service.configuration.parameters

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/configuration/parameters/env_var/): complete subsection reference.

- [file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/configuration/parameters/file/): complete subsection reference.

## Next pages

- [service.configuration.parameters.env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/configuration/parameters/env_var/)
- [service.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/configuration/parameters/file/)
- [service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/configuration/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
