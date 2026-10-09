---
page_title: "advanced_tcp_profile.disable_tcp_advanced_profile"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["advanced tcp profile disable tcp advanced profile"], "body_bytes": 1129, "body_sha256": "sha256:daff02dbaaed3788d8ef1884ff1895794f3955d7b6bd989103b1cce7790513c3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "parent_id": "xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile", "path": "documentation/resources/application_profiles/properties/advanced_tcp_profile/disable_tcp_advanced_profile/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0201221302013100-2013130322130320-1200303113112222-1312203330331102-2031323221233100-0000021120300111-3303331233222322-3330311103011330", "registry_path": "docs/guides/resources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_tcp_profile", "disable_tcp_advanced_profile"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/advanced_tcp_profile/disable_tcp_advanced_profile/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["application_profilesCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
