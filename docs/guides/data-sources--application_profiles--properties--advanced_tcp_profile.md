---
page_title: "advanced_tcp_profile"
subcategory: ""
description: "advanced_tcp_profile for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1585, "body_sha256": "sha256:18ecc41d4a058f4862e46a7fad6201f26237c10e66b3457c1e26f51a1c89bdb1", "canonical_id": "xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile", "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile"], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile", "parent_id": "xcsh-docs:data-sources:application_profiles:reference", "path": "docs/guides/data-sources--application_profiles--properties--advanced_tcp_profile.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_tcp_profile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/advanced_tcp_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_tcp_profile for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_tcp_profile

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- [Property reference](data-sources--application_profiles--reference.md)
- advanced_tcp_profile

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for advanced tcp profile.

Upstream description:

BIG-IP Advanced TCP Profile.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tcp_advanced_profile_choice": "[\"disable_tcp_advanced_profile\",\"enable_tcp_advanced_profile\"]"
}
```

## Direct properties

- [disable_tcp_advanced_profile](data-sources--application_profiles--properties--advanced_tcp_profile--disable_tcp_advanced_profile.md): complete subsection reference.

- [enable_tcp_advanced_profile](data-sources--application_profiles--properties--advanced_tcp_profile--enable_tcp_advanced_profile.md): complete subsection reference.

## Next pages

- [advanced_tcp_profile.disable_tcp_advanced_profile](data-sources--application_profiles--properties--advanced_tcp_profile--disable_tcp_advanced_profile.md)
- [advanced_tcp_profile.enable_tcp_advanced_profile](data-sources--application_profiles--properties--advanced_tcp_profile--enable_tcp_advanced_profile.md)
- [Property reference](data-sources--application_profiles--reference.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
