---
page_title: "job.volumes.empty_dir"
subcategory: "Container"
description: "Volume containing a temporary directory whose lifetime is the same as a replica of a workload."
xcsh_docs: {"aliases": ["job volumes empty dir"], "body_bytes": 1703, "body_sha256": "sha256:47685be6ca38dce6d9b3767348ff10ee80d2e9555da58f52370522a40f2cba31", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:volumes:empty_dir:mount"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:volumes:empty_dir", "parent_id": "xcsh-docs:data-sources:workload:properties:job:volumes", "path": "documentation/data-sources/workload/properties/job/volumes/empty_dir/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-2130120111010100-0110033232023301-3123203333112202-2023301113120022-2031322212021222-1301300101103333-2131320002113230-3100111121000003", "registry_path": "docs/guides/data-sources--workload--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "volumes", "empty_dir"], "schema_version": 1, "sections": [{"aliases": ["job volumes empty dir mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:empty_dir:mount", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "volumes", "empty_dir", "mount"], "syntax": "attribute", "type": "object"}, {"aliases": ["job volumes empty dir size limit"], "anchor": "schema-job--volumes--empty_dir--size_limit", "description": "Configuration parameter for size limit", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:empty_dir", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "volumes", "empty_dir", "size_limit"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/volumes/empty_dir/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["workloadCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes.empty_dir

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/)
- job.volumes.empty_dir

<a id="section"></a>

Type: `"single"`. Computed.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

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

## Direct properties

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/empty_dir/mount/): complete subsection reference.

<a id="schema-job--volumes--empty_dir--size_limit"></a>

### size_limit property

Type: `"number"`. Computed.

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
