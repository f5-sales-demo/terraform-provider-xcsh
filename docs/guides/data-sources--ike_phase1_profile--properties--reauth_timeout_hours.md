---
page_title: "reauth_timeout_hours"
subcategory: ""
description: "reauth_timeout_hours for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": [], "body_bytes": 1694, "body_sha256": "sha256:faa36b359094f61d294140113f8262520f284857fe63382882b398898c786956", "canonical_id": "xcsh-docs:data-sources:ike_phase1_profile:properties:reauth_timeout_hours", "child_ids": [], "collection_id": "xcsh-docs:data-sources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase1_profile:properties:reauth_timeout_hours", "parent_id": "xcsh-docs:data-sources:ike_phase1_profile:reference", "path": "docs/guides/data-sources--ike_phase1_profile--properties--reauth_timeout_hours.md", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["reauth_timeout_hours"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase1_profile/properties/reauth_timeout_hours/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "reauth_timeout_hours for xcsh_ike_phase1_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# reauth_timeout_hours

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md)
- [Property reference](data-sources--ike_phase1_profile--reference.md)
- reauth_timeout_hours

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout hours.

Upstream description:

Input Hours.

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

<a id="schema-reauth_timeout_hours--duration"></a>

### duration property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
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
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

## Next pages

- [Property reference](data-sources--ike_phase1_profile--reference.md)
- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md)
