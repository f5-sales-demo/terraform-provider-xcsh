---
page_title: "stateful_service.volumes.empty_dir"
subcategory: "Container"
description: "Volume containing a temporary directory whose lifetime is the same as a replica of a workload."
xcsh_docs: {"aliases": ["stateful service volumes empty dir"], "body_bytes": 2059, "body_sha256": "sha256:0a1f4686f001a90341bb4fddc12a16fa63ae27b7f41e9485abf5833387a1b856", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir:mount"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes", "path": "documentation/resources/workload/properties/stateful_service/volumes/empty_dir/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1233111023320131-2313133001020012-3013001223020200-0301013301200313-0310102301020311-1212003232313031-1122230103210021-0203223003200300", "registry_path": "docs/guides/resources--workload--reference--group-029.md", "relationships": [{"anchor": "schema-stateful_service--volumes--empty_dir--size_limit", "enforcement": "provider-schema", "group": "stateful_service.volumes.empty_dir:RequiredObjectAttributes:size_limit", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "volumes", "empty_dir"], "schema_version": 1, "sections": [{"aliases": ["stateful service volumes empty dir mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir:mount", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--volumes--empty_dir--mount--mount_path", "enforcement": "provider-schema", "group": "stateful_service.volumes.empty_dir.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir:mount", "type": "requires"}], "schema_path": ["stateful_service", "volumes", "empty_dir", "mount"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service volumes empty dir size limit"], "anchor": "schema-stateful_service--volumes--empty_dir--size_limit", "description": "Configuration parameter for size limit", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "volumes", "empty_dir", "size_limit"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/volumes/empty_dir/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
