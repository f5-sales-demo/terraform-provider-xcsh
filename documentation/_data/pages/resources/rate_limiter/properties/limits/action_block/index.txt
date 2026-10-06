---
page_title: "limits.action_block"
subcategory: "Security"
description: "Action where a user is blocked from making further requests after exceeding rate limit threshold."
xcsh_docs: {"aliases": ["limits action block"], "body_bytes": 1825, "body_sha256": "sha256:ce3c3a41eef6fa79b8f571d8f54b5fb78ef2f982f503ecd44d4da803b5fe3857", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block", "parent_id": "xcsh-docs:resources:rate_limiter:properties:limits", "path": "documentation/resources/rate_limiter/properties/limits/action_block/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3330231013002132-1002333201033023-3203301313302131-2201003132020220-0001323103320111-3330113203231300-1211302013131322-2012323210131033", "registry_path": "docs/guides/resources--rate_limiter--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["limits", "action_block"], "schema_version": 1, "sections": [{"aliases": ["limits action block hours"], "anchor": "section", "description": "Input Duration Hours.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block", "hours"], "syntax": "block", "type": "object"}, {"aliases": ["limits action block minutes"], "anchor": "section", "description": "Input Duration Minutes.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block", "minutes"], "syntax": "block", "type": "object"}, {"aliases": ["limits action block seconds"], "anchor": "section", "description": "Input Duration Seconds.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block", "seconds"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/properties/limits/action_block/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits.action_block

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/)
- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/)
- limits.action_block

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes"),
  validators.ConflictingObjectAttributes("hours",
    "seconds"),
  validators.ConflictingObjectAttributes("minutes",
    "seconds")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

## Direct properties

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/minutes/): complete subsection reference.

- [seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/seconds/): complete subsection reference.
