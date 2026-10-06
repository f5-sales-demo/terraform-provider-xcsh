---
page_title: "stateful_service.volumes.empty_dir"
subcategory: "Container"
description: "Volume containing a temporary directory whose lifetime is the same as a replica of a workload."
xcsh_docs: {"aliases": ["stateful service volumes empty dir"], "body_bytes": 2089, "body_sha256": "sha256:505ffd74f04a5435e38690a9c9dd03574047159e1d5df8e4d35fc7c6d16c7c5e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir:mount"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes", "path": "documentation/resources/workload/properties/stateful_service/volumes/empty_dir/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1233111023320131-2313133001020012-3013001223020200-0301013301200313-0310102301020311-1212003232313031-1122230103210021-0203223003200300", "registry_path": "docs/guides/resources--workload--reference--group-028.md", "relationships": [{"anchor": "schema-stateful_service--volumes--empty_dir--size_limit", "enforcement": "provider-schema", "group": "stateful_service.volumes.empty_dir:RequiredObjectAttributes:size_limit", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "volumes", "empty_dir"], "schema_version": 1, "sections": [{"aliases": ["stateful service volumes empty dir mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir:mount", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--volumes--empty_dir--mount--mount_path", "enforcement": "provider-schema", "group": "stateful_service.volumes.empty_dir.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir:mount", "type": "requires"}], "schema_path": ["stateful_service", "volumes", "empty_dir", "mount"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service volumes empty dir size limit"], "anchor": "schema-stateful_service--volumes--empty_dir--size_limit", "description": "Configuration parameter for size limit", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "volumes", "empty_dir", "size_limit"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/volumes/empty_dir/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.volumes.empty_dir

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/volumes/)
- stateful_service.volumes.empty_dir

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("size_limit")}
```

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
empty_dir {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/volumes/empty_dir/mount/): complete subsection reference.

<a id="schema-stateful_service--volumes--empty_dir--size_limit"></a>

### size_limit property

Type: `"number"`. Optional.

Size Limit (in GiB). Configuration parameter for size limit

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```
