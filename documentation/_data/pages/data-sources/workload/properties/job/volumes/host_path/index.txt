---
page_title: "job.volumes.host_path"
subcategory: "Container"
description: "Volume containing a host mapped path into the workload."
xcsh_docs: {"aliases": ["job volumes host path"], "body_bytes": 2090, "body_sha256": "sha256:f87dadb7a5a09333e4a59ab28f0bab1d71adc6d0a4e528fe322fb08a0366981b", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:volumes:host_path:mount"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:volumes:host_path", "parent_id": "xcsh-docs:data-sources:workload:properties:job:volumes", "path": "documentation/data-sources/workload/properties/job/volumes/host_path/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-1101123330022122-0312023230320121-2333332030310233-1331312111301103-0120032000312031-2200113123130032-2333123300102003-0212211023011122", "registry_path": "docs/guides/data-sources--workload--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "volumes", "host_path"], "schema_version": 1, "sections": [{"aliases": ["job volumes host path mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:host_path:mount", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "volumes", "host_path", "mount"], "syntax": "attribute", "type": "object"}, {"aliases": ["job volumes host path path"], "anchor": "schema-job--volumes--host_path--path", "description": "Path of the directory on the host.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:host_path", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "volumes", "host_path", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/volumes/host_path/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Volume containing a host mapped path into the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["workloadCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes.host_path

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/)
- job.volumes.host_path

<a id="section"></a>

Type: `"single"`. Computed.

Volume containing a host mapped path into the workload.

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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/host_path/mount/): complete subsection reference.

<a id="schema-job--volumes--host_path--path"></a>

### path property

Type: `"string"`. Computed.

Path. Path of the directory on the host.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "minLength": 1,
    "pattern": "[^\\\\0]+"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  }
}
```
