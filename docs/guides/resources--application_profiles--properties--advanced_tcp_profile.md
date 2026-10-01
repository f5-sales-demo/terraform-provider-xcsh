---
page_title: "advanced_tcp_profile"
subcategory: ""
description: "advanced_tcp_profile for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1895, "body_sha256": "sha256:807ef036724d80b20a19bb571f8d4d2e7c4ca01b9639a12b5af302f05e58e650", "canonical_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile", "child_ids": ["xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile", "parent_id": "xcsh-docs:resources:application_profiles:reference", "path": "docs/guides/resources--application_profiles--properties--advanced_tcp_profile.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_tcp_profile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/advanced_tcp_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_tcp_profile for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_tcp_profile

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- [Property reference](resources--application_profiles--reference.md)
- advanced_tcp_profile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for advanced tcp profile.

Upstream description:

BIG-IP Advanced TCP Profile.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_tcp_advanced_profile",
    "enable_tcp_advanced_profile")}
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
  "x-ves-oneof-field-tcp_advanced_profile_choice": "[\"disable_tcp_advanced_profile\",\"enable_tcp_advanced_profile\"]"
}
```

Terraform syntax:

```terraform
advanced_tcp_profile {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_tcp_advanced_profile](resources--application_profiles--properties--advanced_tcp_profile--disable_tcp_advanced_profile.md): complete subsection reference.

- [enable_tcp_advanced_profile](resources--application_profiles--properties--advanced_tcp_profile--enable_tcp_advanced_profile.md): complete subsection reference.

## Next pages

- [advanced_tcp_profile.disable_tcp_advanced_profile](resources--application_profiles--properties--advanced_tcp_profile--disable_tcp_advanced_profile.md)
- [advanced_tcp_profile.enable_tcp_advanced_profile](resources--application_profiles--properties--advanced_tcp_profile--enable_tcp_advanced_profile.md)
- [Property reference](resources--application_profiles--reference.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
