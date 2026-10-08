---
page_title: "advanced_tcp_profile"
subcategory: ""
description: "BIG-IP Advanced TCP Profile."
xcsh_docs: {"aliases": ["advanced tcp profile"], "body_bytes": 1676, "body_sha256": "sha256:2971876d2760f27f99d73dfa2b7b1b9fb901bc8e3239b8ce4903db7cf44e3844", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile", "parent_id": "xcsh-docs:resources:application_profiles:reference", "path": "documentation/resources/application_profiles/properties/advanced_tcp_profile/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0102012130211220-0033013023011312-0020312201233333-0110230303020102-0021300320132223-2232021231133230-0130333111122200-0300200330323322", "registry_path": "docs/guides/resources--application_profiles--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "advanced_tcp_profile:ConflictingObjectAttributes:disable_tcp_advanced_profile,enable_tcp_advanced_profile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_tcp_profile:ConflictingObjectAttributes:disable_tcp_advanced_profile,enable_tcp_advanced_profile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_tcp_profile"], "schema_version": 1, "sections": [{"aliases": ["advanced tcp profile disable tcp advanced profile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_tcp_profile", "disable_tcp_advanced_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced tcp profile enable tcp advanced profile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_tcp_profile", "enable_tcp_advanced_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/advanced_tcp_profile/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "BIG-IP Advanced TCP Profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["application_profilesCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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

Additional upstream details:

BIG-IP Advanced TCP Profile.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
