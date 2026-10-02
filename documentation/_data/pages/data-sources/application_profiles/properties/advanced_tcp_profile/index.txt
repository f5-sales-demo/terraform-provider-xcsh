---
page_title: "advanced_tcp_profile"
subcategory: ""
description: "BIG-IP Advanced TCP Profile."
xcsh_docs: {"aliases": ["advanced tcp profile"], "body_bytes": 1993, "body_sha256": "sha256:cb210aa477bcd4c80b224edb18726db401806c95214740a839063dabe4813045", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile", "parent_id": "xcsh-docs:data-sources:application_profiles:reference", "path": "documentation/data-sources/application_profiles/properties/advanced_tcp_profile/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3021122313210131-2313033223133231-0330112103210330-1123322311012132-1112331223212022-3311220221212131-2333232332331221-1001103023222323", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_tcp_profile"], "schema_version": 1, "sections": [{"aliases": ["disable tcp advanced profile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile:disable_tcp_advanced_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_tcp_profile", "disable_tcp_advanced_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable tcp advanced profile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile:enable_tcp_advanced_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_tcp_profile", "enable_tcp_advanced_profile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/advanced_tcp_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "BIG-IP Advanced TCP Profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_tcp_profile

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
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

- [disable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/advanced_tcp_profile/disable_tcp_advanced_profile/): complete subsection reference.

- [enable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/advanced_tcp_profile/enable_tcp_advanced_profile/): complete subsection reference.

## Next pages

- [advanced_tcp_profile.disable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/advanced_tcp_profile/disable_tcp_advanced_profile/)
- [advanced_tcp_profile.enable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/advanced_tcp_profile/enable_tcp_advanced_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
