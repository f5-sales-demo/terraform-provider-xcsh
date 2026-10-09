---
page_title: "service.volumes.host_path"
subcategory: "Container"
description: "Volume containing a host mapped path into the workload."
xcsh_docs: {"aliases": ["service volumes host path"], "body_bytes": 2122, "body_sha256": "sha256:43ed6c4c4c35c4bff408bc8c8dd5cb27232b2b958f9c318192487f7a7771d8e9", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:volumes:host_path:mount"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:volumes:host_path", "parent_id": "xcsh-docs:data-sources:workload:properties:service:volumes", "path": "documentation/data-sources/workload/properties/service/volumes/host_path/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3220202111330130-0310032103002210-0122302221113022-3022122321110002-1211123003210300-3121213331203021-3230312212312232-2213002303130131", "registry_path": "docs/guides/data-sources--workload--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "volumes", "host_path"], "schema_version": 1, "sections": [{"aliases": ["service volumes host path mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:service:volumes:host_path:mount", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "volumes", "host_path", "mount"], "syntax": "attribute", "type": "object"}, {"aliases": ["service volumes host path path"], "anchor": "schema-service--volumes--host_path--path", "description": "Path of the directory on the host.", "document_id": "xcsh-docs:data-sources:workload:properties:service:volumes:host_path", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "volumes", "host_path", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/volumes/host_path/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Volume containing a host mapped path into the workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes.host_path

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/)
- service.volumes.host_path

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

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/host_path/mount/): complete subsection reference.

<a id="schema-service--volumes--host_path--path"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
