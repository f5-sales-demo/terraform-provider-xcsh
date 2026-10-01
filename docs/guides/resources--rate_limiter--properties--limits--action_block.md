---
page_title: "limits.action_block"
subcategory: "Security"
description: "limits.action_block for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1955, "body_sha256": "sha256:2ff31ab112f6fecd24fef88f5aa62c799a0f3ae10c40bb32a921cea39b5221d1", "canonical_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block", "child_ids": ["xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds"], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block", "parent_id": "xcsh-docs:resources:rate_limiter:properties:limits", "path": "docs/guides/resources--rate_limiter--properties--limits--action_block.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["limits", "action_block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/properties/limits/action_block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "limits.action_block for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits.action_block

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md)
- [Property reference](resources--rate_limiter--reference.md)
- [limits](resources--rate_limiter--properties--limits.md)
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

- [hours](resources--rate_limiter--properties--limits--action_block--hours.md): complete subsection reference.

- [minutes](resources--rate_limiter--properties--limits--action_block--minutes.md): complete subsection reference.

- [seconds](resources--rate_limiter--properties--limits--action_block--seconds.md): complete subsection reference.

## Next pages

- [limits.action_block.hours](resources--rate_limiter--properties--limits--action_block--hours.md)
- [limits.action_block.minutes](resources--rate_limiter--properties--limits--action_block--minutes.md)
- [limits.action_block.seconds](resources--rate_limiter--properties--limits--action_block--seconds.md)
- [limits](resources--rate_limiter--properties--limits.md)
- [xcsh_rate_limiter](../resources/rate_limiter.md)
