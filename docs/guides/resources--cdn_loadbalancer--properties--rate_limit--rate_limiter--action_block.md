---
page_title: "rate_limit.rate_limiter.action_block"
subcategory: "Load Balancing"
description: "rate_limit.rate_limiter.action_block for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2341, "body_sha256": "sha256:35b06620f1c10053ae976e3d649dff463fd52e38dbc707523ef50d135de7c577", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "path": "docs/guides/resources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "action_block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.rate_limiter.action_block for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.action_block

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [rate_limit](resources--cdn_loadbalancer--properties--rate_limit.md)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter.md)
- rate_limit.rate_limiter.action_block

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

- [hours](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block--hours.md): complete subsection reference.

- [minutes](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block--minutes.md): complete subsection reference.

- [seconds](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block--seconds.md): complete subsection reference.

## Next pages

- [rate_limit.rate_limiter.action_block.hours](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block--hours.md)
- [rate_limit.rate_limiter.action_block.minutes](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block--minutes.md)
- [rate_limit.rate_limiter.action_block.seconds](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block--seconds.md)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
