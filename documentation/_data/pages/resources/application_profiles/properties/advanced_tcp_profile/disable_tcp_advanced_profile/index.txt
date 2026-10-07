---
page_title: "advanced_tcp_profile.disable_tcp_advanced_profile"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["advanced tcp profile disable tcp advanced profile"], "body_bytes": 1129, "body_sha256": "sha256:daff02dbaaed3788d8ef1884ff1895794f3955d7b6bd989103b1cce7790513c3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "parent_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile", "path": "documentation/resources/application_profiles/properties/advanced_tcp_profile/disable_tcp_advanced_profile/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0201221302013100-2013130322130320-1200303113112222-1312203330331102-2031323221233100-0000021120300111-3303331233222322-3330311103011330", "registry_path": "docs/guides/resources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_tcp_profile", "disable_tcp_advanced_profile"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/advanced_tcp_profile/disable_tcp_advanced_profile/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["application_profilesCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_tcp_profile.disable_tcp_advanced_profile

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [advanced_tcp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/advanced_tcp_profile/)
- advanced_tcp_profile.disable_tcp_advanced_profile

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tcp advanced profile.

Additional upstream details:

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

Terraform syntax:

```terraform
disable_tcp_advanced_profile = {}
```

This is an empty object or choice marker. It has no direct properties.
