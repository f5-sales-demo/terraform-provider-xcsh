---
page_title: "disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disabled"], "body_bytes": 1185, "body_sha256": "sha256:343547617856a3ddcf6b455af1fb6ed28c77c9282aca9a848e9e0676fb27c54c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:disabled", "parent_id": "xcsh-docs:resources:cloud_link:reference", "path": "documentation/resources/cloud_link/properties/disabled/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1022131120022203-1330132311200122-3203320013133123-1312200301110003-2220313021320223-0321311322113312-1231023302122033-0222221310010130", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/disabled/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disabled

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disabled, enabled\] Enable this option

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

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/disabled/#section)
- [enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/enabled/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
