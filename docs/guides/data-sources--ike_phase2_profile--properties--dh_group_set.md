---
page_title: "dh_group_set"
subcategory: ""
description: "dh_group_set for xcsh_ike_phase2_profile."
xcsh_docs: {"aliases": [], "body_bytes": 2384, "body_sha256": "sha256:da7248e7be9cf91304fb5786b6fdcf79ea9cb29b76076d98d3d778fe4e02b22c", "canonical_id": "xcsh-docs:data-sources:ike_phase2_profile:properties:dh_group_set", "child_ids": [], "collection_id": "xcsh-docs:data-sources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase2_profile:properties:dh_group_set", "parent_id": "xcsh-docs:data-sources:ike_phase2_profile:reference", "path": "docs/guides/data-sources--ike_phase2_profile--properties--dh_group_set.md", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dh_group_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase2_profile/properties/dh_group_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dh_group_set for xcsh_ike_phase2_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dh_group_set

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md)
- [Property reference](data-sources--ike_phase2_profile--reference.md)
- dh_group_set

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: dh\_group\_set, disable\_pfs; Default: disable\_pfs\] Choose the acceptable Diffie
Hellman(DH) Group or Groups that you are willing to accept as part of this profile.

Upstream description:

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile.

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

- [dh_group_set](data-sources--ike_phase2_profile--properties--dh_group_set.md#section)
- [disable_pfs](data-sources--ike_phase2_profile--properties--disable_pfs.md#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-dh_group_set--dh_groups"></a>

### dh_groups property

Type: `["list", "string"]`. Computed.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile. Possible values are \`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`,
\`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`, \`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`,
\`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`. Defaults to \`DH\_GROUP\_DEFAULT\`.

Upstream description:

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [Property reference](data-sources--ike_phase2_profile--reference.md)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md)
