---
page_title: "advanced_tcp_profile"
subcategory: ""
description: "BIG-IP Advanced TCP Profile."
xcsh_docs: {"aliases": ["advanced tcp profile"], "body_bytes": 1646, "body_sha256": "sha256:801a96e527dcd9af4241e300dd36865cc68b682169062bef2043d52cb580fb0d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile", "parent_id": "xcsh-docs:resources:application_profiles:reference", "path": "documentation/resources/application_profiles/properties/advanced_tcp_profile/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0102012130211220-0033013023011312-0020312201233333-0110230303020102-0021300320132223-2232021231133230-0130333111122200-0300200330323322", "registry_path": "docs/guides/resources--application_profiles--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "advanced_tcp_profile:ConflictingObjectAttributes:disable_tcp_advanced_profile,enable_tcp_advanced_profile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_tcp_profile:ConflictingObjectAttributes:disable_tcp_advanced_profile,enable_tcp_advanced_profile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_tcp_profile"], "schema_version": 1, "sections": [{"aliases": ["advanced tcp profile disable tcp advanced profile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_tcp_profile", "disable_tcp_advanced_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced tcp profile enable tcp advanced profile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_tcp_profile", "enable_tcp_advanced_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/advanced_tcp_profile/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "BIG-IP Advanced TCP Profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
