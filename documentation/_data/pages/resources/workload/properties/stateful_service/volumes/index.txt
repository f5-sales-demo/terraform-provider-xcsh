---
page_title: "stateful_service.volumes"
subcategory: "Container"
description: "Ephemeral volumes for the service."
xcsh_docs: {"aliases": ["stateful service volumes"], "body_bytes": 3336, "body_sha256": "sha256:5e438152047b4001f159cfee6ec7749085ffdc9ca0f3e33e396dbd152f66337e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "xcsh-docs:resources:workload:properties:stateful_service:volumes:host_path"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:volumes", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service", "path": "documentation/resources/workload/properties/stateful_service/volumes/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213", "registry_path": "docs/guides/resources--workload--reference--group-029.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.volumes:ConflictingListObjectAttributes:empty_dir,host_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.volumes:ConflictingListObjectAttributes:empty_dir,host_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:host_path", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "volumes"], "schema_version": 1, "sections": [{"aliases": ["stateful service volumes empty dir"], "anchor": "section", "description": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--volumes--empty_dir--size_limit", "enforcement": "provider-schema", "group": "stateful_service.volumes.empty_dir:RequiredObjectAttributes:size_limit", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:empty_dir", "type": "requires"}], "schema_path": ["stateful_service", "volumes", "empty_dir"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service volumes host path"], "anchor": "section", "description": "Volume containing a host mapped path into the workload.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:host_path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--volumes--host_path--path", "enforcement": "provider-schema", "group": "stateful_service.volumes.host_path:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes:host_path", "type": "requires"}], "schema_path": ["stateful_service", "volumes", "host_path"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service volumes name"], "anchor": "schema-stateful_service--volumes--name", "description": "Name of the volume.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:volumes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "volumes", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/volumes/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Ephemeral volumes for the service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.volumes

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- stateful_service.volumes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ephemeral Volumes. Ephemeral volumes for the service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("empty_dir",
    "host_path")}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [empty_dir](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/volumes/empty_dir/): complete subsection reference.

- [host_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/volumes/host_path/): complete subsection reference.

<a id="schema-stateful_service--volumes--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the volume.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
