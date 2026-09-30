---
page_title: "reauth_disabled"
subcategory: ""
description: "reauth_disabled for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": [], "body_bytes": 1289, "body_sha256": "sha256:a00be139357fd32a00761e23da9e33e2f288a68eba9c1fb7933208c1193ac667", "canonical_id": "xcsh-docs:resources:ike_phase1_profile:properties:reauth_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase1_profile:properties:reauth_disabled", "parent_id": "xcsh-docs:resources:ike_phase1_profile:reference", "path": "docs/guides/resources--ike_phase1_profile--properties--reauth_disabled.md", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["reauth_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/properties/reauth_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "reauth_disabled for xcsh_ike_phase1_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# reauth_disabled

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md)
- [Property reference](resources--ike_phase1_profile--reference.md)
- reauth_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [reauth_disabled](resources--ike_phase1_profile--properties--reauth_disabled.md#section)
- [reauth_timeout_days](resources--ike_phase1_profile--properties--reauth_timeout_days.md#section)
- [reauth_timeout_hours](resources--ike_phase1_profile--properties--reauth_timeout_hours.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
reauth_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--ike_phase1_profile--reference.md)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md)
