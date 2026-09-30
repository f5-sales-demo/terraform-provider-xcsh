---
page_title: "reauth_timeout_hours"
subcategory: ""
description: "reauth_timeout_hours for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": [], "body_bytes": 1595, "body_sha256": "sha256:fc3456e2144f55ca3b0c060fe2f9e17ac233eb219ff6d45bd886394b033e0351", "canonical_id": "xcsh-docs:data-sources:ike_phase1_profile:properties:reauth_timeout_hours", "child_ids": [], "collection_id": "xcsh-docs:data-sources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase1_profile:properties:reauth_timeout_hours", "parent_id": "xcsh-docs:data-sources:ike_phase1_profile:reference", "path": "docs/guides/data-sources--ike_phase1_profile--properties--reauth_timeout_hours.md", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["reauth_timeout_hours"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase1_profile/properties/reauth_timeout_hours/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "reauth_timeout_hours for xcsh_ike_phase1_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
