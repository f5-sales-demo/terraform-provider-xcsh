---
page_title: "advanced_tcp_profile"
subcategory: ""
description: "advanced_tcp_profile for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 2303, "body_sha256": "sha256:30676fe58ff6d2656c237caf97e059b3eea441b7777e1834db7385cef0a93bf1", "child_ids": ["xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile", "parent_id": "xcsh-docs:resources:application_profiles:reference", "path": "documentation/resources/application_profiles/properties/advanced_tcp_profile/index.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["advanced_tcp_profile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/advanced_tcp_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_tcp_profile for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_tcp_profile

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
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

- [disable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/advanced_tcp_profile/disable_tcp_advanced_profile/): complete subsection reference.

- [enable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/advanced_tcp_profile/enable_tcp_advanced_profile/): complete subsection reference.

## Next pages

- [advanced_tcp_profile.disable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/advanced_tcp_profile/disable_tcp_advanced_profile/)
- [advanced_tcp_profile.enable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/advanced_tcp_profile/enable_tcp_advanced_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
