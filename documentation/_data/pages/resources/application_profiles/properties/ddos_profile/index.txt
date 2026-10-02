---
page_title: "ddos_profile"
subcategory: ""
description: "BIG-IP DDoS Protection Rules."
xcsh_docs: {"aliases": ["ddos profile"], "body_bytes": 2159, "body_sha256": "sha256:0d34763ac9b1e1d7ae31d5a46cd1f977eebcae291c31ad1519080a8804495e41", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:application_profiles:properties:ddos_profile:disable_ddos_mitigation", "xcsh-docs:resources:application_profiles:properties:ddos_profile:enable_ddos_mitigation"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:ddos_profile", "parent_id": "xcsh-docs:resources:application_profiles:reference", "path": "documentation/resources/application_profiles/properties/ddos_profile/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2202332031223123-2301302010321012-1011102311002222-1100030332313031-3311332110021011-2231133020100211-2122201323020110-1321212330120300", "registry_path": "docs/guides/resources--application_profiles--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ddos_profile:ConflictingObjectAttributes:disable_ddos_mitigation,enable_ddos_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:ddos_profile:disable_ddos_mitigation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ddos_profile:ConflictingObjectAttributes:disable_ddos_mitigation,enable_ddos_mitigation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:application_profiles:properties:ddos_profile:enable_ddos_mitigation", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_profile"], "schema_version": 1, "sections": [{"aliases": ["disable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:ddos_profile:disable_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "disable_ddos_mitigation"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:application_profiles:properties:ddos_profile:enable_ddos_mitigation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "enable_ddos_mitigation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/ddos_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "BIG-IP DDoS Protection Rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_profile

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- ddos_profile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Upstream description:

BIG-IP DDoS Protection Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_ddos_mitigation",
    "enable_ddos_mitigation")}
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
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

Terraform syntax:

```terraform
ddos_profile {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/ddos_profile/disable_ddos_mitigation/): complete subsection reference.

- [enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/ddos_profile/enable_ddos_mitigation/): complete subsection reference.

## Next pages

- [ddos_profile.disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/ddos_profile/disable_ddos_mitigation/)
- [ddos_profile.enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/ddos_profile/enable_ddos_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
