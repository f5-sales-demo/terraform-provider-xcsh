---
page_title: "reauth_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["reauth disabled"], "body_bytes": 1355, "body_sha256": "sha256:49ce47797b6b3400efcc6f1c581e7dbe37551446e249fcb3ef0d2e17d110f2d1", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike1:properties:reauth_disabled", "parent_id": "xcsh-docs:data-sources:ike1:reference", "path": "documentation/data-sources/ike1/properties/reauth_disabled/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3033312231301021-1312332110313310-0022003320203320-3311113320322230-2303310002023031-0303231333300103-2003132232103312-3130230321003312", "registry_path": "docs/guides/data-sources--ike1--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["reauth_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike1/properties/reauth_disabled/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["ike1CreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# reauth_disabled

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/)
- reauth_disabled

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

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

- [reauth_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_disabled/#section)
- [reauth_timeout_days](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_days/#section)
- [reauth_timeout_hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/reauth_timeout_hours/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
