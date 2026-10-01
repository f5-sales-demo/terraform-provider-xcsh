---
page_title: "reauth_timeout_days"
subcategory: ""
description: "reauth_timeout_days for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": [], "body_bytes": 2098, "body_sha256": "sha256:4309ae07e3550d6fd977bcab6ae4de19b16e54c7a244222659720c8559c9b3a3", "canonical_id": "xcsh-docs:resources:ike_phase1_profile:properties:reauth_timeout_days", "child_ids": [], "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase1_profile:properties:reauth_timeout_days", "parent_id": "xcsh-docs:resources:ike_phase1_profile:reference", "path": "docs/guides/resources--ike_phase1_profile--properties--reauth_timeout_days.md", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["reauth_timeout_days"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/properties/reauth_timeout_days/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "reauth_timeout_days for xcsh_ike_phase1_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# reauth_timeout_days

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md)
- [Property reference](resources--ike_phase1_profile--reference.md)
- reauth_timeout_days

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
reauth_timeout_days {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-reauth_timeout_days--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

## Next pages

- [Property reference](resources--ike_phase1_profile--reference.md)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md)
