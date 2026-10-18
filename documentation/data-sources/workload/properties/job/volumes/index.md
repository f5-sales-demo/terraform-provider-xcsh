---
page_title: "job.volumes"
subcategory: "Container"
description: "Volumes for the job."
xcsh_docs: {"aliases": ["job volumes"], "body_bytes": 2985, "body_sha256": "sha256:d116c7135fd7bd3e29f60a5dd32174383ed8700eb294e8133e9f6332cf108fad", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:volumes:empty_dir", "xcsh-docs:data-sources:workload:properties:job:volumes:host_path", "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:volumes", "parent_id": "xcsh-docs:data-sources:workload:properties:job", "path": "documentation/data-sources/workload/properties/job/volumes/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011", "registry_path": "docs/guides/data-sources--workload--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "volumes"], "schema_version": 1, "sections": [{"aliases": ["job volumes empty dir"], "anchor": "section", "description": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:empty_dir", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "volumes", "empty_dir"], "syntax": "attribute", "type": "object"}, {"aliases": ["job volumes host path"], "anchor": "section", "description": "Volume containing a host mapped path into the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:host_path", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "volumes", "host_path"], "syntax": "attribute", "type": "object"}, {"aliases": ["job volumes name"], "anchor": "schema-job--volumes--name", "description": "Name of the volume.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "volumes", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["job volumes persistent volume"], "anchor": "section", "description": "Volume containing the Persistent Storage for the workload.", "document_id": "xcsh-docs:data-sources:workload:properties:job:volumes:persistent_volume", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "volumes", "persistent_volume"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/volumes/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Volumes for the job.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["workloadCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- job.volumes

<a id="section"></a>

Type: `"list"`. Computed.

Volumes. Volumes for the job.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-17T10:10:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

## Direct properties

- [empty_dir](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/empty_dir/): complete subsection reference.

- [host_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/host_path/): complete subsection reference.

<a id="schema-job--volumes--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-17T10:10:42+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/volumes/persistent_volume/): complete subsection reference.
