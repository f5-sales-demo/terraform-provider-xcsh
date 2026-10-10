---
page_title: "business_logic_markup_setting.enable"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["business logic markup setting enable"], "body_bytes": 1012, "body_sha256": "sha256:62e2b1fc12beea6b46fe95087ac04fc9bb3b2ea08d10e34b22a556b713b8c3ec", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:enable", "parent_id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting", "path": "documentation/resources/app_type/properties/business_logic_markup_setting/enable/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3212031320200323-1032122133100010-2132303001322023-0010131003113211-2213333112013221-1030000033220230-1310021320210103-1030132323201200", "registry_path": "docs/guides/resources--app_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["business_logic_markup_setting", "enable"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_type/properties/business_logic_markup_setting/enable/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_typeCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# business_logic_markup_setting.enable

Breadcrumbs:

- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/)
- [business_logic_markup_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/business_logic_markup_setting/)
- business_logic_markup_setting.enable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
enable = {}
```

This is an empty object or choice marker. It has no direct properties.
