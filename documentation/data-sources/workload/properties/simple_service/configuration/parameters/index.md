---
page_title: "simple_service.configuration.parameters"
subcategory: "Container"
description: "Parameters for the workload."
xcsh_docs: {"aliases": ["simple service configuration parameters"], "body_bytes": 2613, "body_sha256": "sha256:c806ed47608bf67ca730208838fd29537a78bc66c10c82e4153ca663b57594b7", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:env_var", "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:file"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration", "path": "documentation/data-sources/workload/properties/simple_service/configuration/parameters/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1312302100222221-3310233130230312-3011122312201321-3020123032330102-1223232031312102-2111321113311321-1112011300203202-2330023033123211", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "configuration", "parameters"], "schema_version": 1, "sections": [{"aliases": ["env var"], "anchor": "section", "description": "Environment Variable.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:env_var", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "env_var"], "syntax": "attribute", "type": "object"}, {"aliases": ["file"], "anchor": "section", "description": "Configuration File for the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "file"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/configuration/parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.configuration.parameters

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/)
- simple_service.configuration.parameters

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

- [env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/env_var/): complete subsection reference.

- [file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/file/): complete subsection reference.

## Next pages

- [simple_service.configuration.parameters.env_var](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/env_var/)
- [simple_service.configuration.parameters.file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/file/)
- [simple_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
