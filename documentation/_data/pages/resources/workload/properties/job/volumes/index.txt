---
page_title: "job.volumes"
subcategory: "Container"
description: "Volumes for the job."
xcsh_docs: {"aliases": ["job volumes"], "body_bytes": 3621, "body_sha256": "sha256:aae43732d14186c6bbcc71cb3659751953145a2ae55843d0b7c84e415a6c92a1", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:volumes:empty_dir", "xcsh-docs:resources:workload:properties:job:volumes:host_path", "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:volumes", "parent_id": "xcsh-docs:resources:workload:properties:job", "path": "documentation/resources/workload/properties/job/volumes/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131", "registry_path": "docs/guides/resources--workload--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:empty_dir,host_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:empty_dir", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:empty_dir,persistent_volume", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:empty_dir", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:empty_dir,host_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:host_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:host_path,persistent_volume", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:host_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:empty_dir,persistent_volume", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.volumes:ConflictingListObjectAttributes:host_path,persistent_volume", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "volumes"], "schema_version": 1, "sections": [{"aliases": ["job volumes empty dir"], "anchor": "section", "description": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "document_id": "xcsh-docs:resources:workload:properties:job:volumes:empty_dir", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-job--volumes--empty_dir--size_limit", "enforcement": "provider-schema", "group": "job.volumes.empty_dir:RequiredObjectAttributes:size_limit", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:empty_dir", "type": "requires"}], "schema_path": ["job", "volumes", "empty_dir"], "syntax": "block", "type": "object"}, {"aliases": ["job volumes host path"], "anchor": "section", "description": "Volume containing a host mapped path into the workload.", "document_id": "xcsh-docs:resources:workload:properties:job:volumes:host_path", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-job--volumes--host_path--path", "enforcement": "provider-schema", "group": "job.volumes.host_path:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:volumes:host_path", "type": "requires"}], "schema_path": ["job", "volumes", "host_path"], "syntax": "block", "type": "object"}, {"aliases": ["job volumes name"], "anchor": "schema-job--volumes--name", "description": "Name of the volume.", "document_id": "xcsh-docs:resources:workload:properties:job:volumes", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "volumes", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["job volumes persistent volume"], "anchor": "section", "description": "Volume containing the Persistent Storage for the workload.", "document_id": "xcsh-docs:resources:workload:properties:job:volumes:persistent_volume", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "volumes", "persistent_volume"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/volumes/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Volumes for the job.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["workloadCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.volumes

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- job.volumes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Volumes. Volumes for the job.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("empty_dir",
    "host_path"),
  validators.ConflictingListObjectAttributes("empty_dir",
    "persistent_volume"),
  validators.ConflictingListObjectAttributes("host_path",
    "persistent_volume")}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
volumes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [empty_dir](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/empty_dir/): complete subsection reference.

- [host_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/host_path/): complete subsection reference.

<a id="schema-job--volumes--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the volume.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [persistent_volume](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/volumes/persistent_volume/): complete subsection reference.
