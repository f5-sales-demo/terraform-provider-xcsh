---
page_title: "disable_url_categorization"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable url categorization"], "body_bytes": 1430, "body_sha256": "sha256:f31d5ff5cd438b516c720c7287a53d1fcfec692914f3c4f796641c14553db0fe", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:disable_url_categorization", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/disable_url_categorization/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0302003210031020-0002223310123233-0230133013212201-1212020012311333-0112301130131323-0032203021212031-1033011301012302-3123112120202231", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_url_categorization"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/disable_url_categorization/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_url_categorization

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- disable_url_categorization

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_url\_categorization, enable\_url\_categorization; Default:
disable\_url\_categorization\] Enable this option

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

OneOf alternatives in this subsection:

- [disable_url_categorization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/disable_url_categorization/#section)
- [enable_url_categorization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/enable_url_categorization/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_url_categorization = {}
```

This is an empty object or choice marker. It has no direct properties.
