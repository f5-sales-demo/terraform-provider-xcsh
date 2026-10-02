---
page_title: "default_sriov_interface"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default sriov interface"], "body_bytes": 1530, "body_sha256": "sha256:4272b92eb87621016d76ef2db441e605242f18a409e8c40251259697e3f2c2e3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:default_sriov_interface", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/default_sriov_interface/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1033001023033030-3300110122022001-2102101121002010-0102011111321302-3323110310023332-2112010011212302-1110320322001132-3102320123230202", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_sriov_interface"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/default_sriov_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_sriov_interface

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- default_sriov_interface

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_sriov\_interface, sriov\_interfaces; Default: default\_sriov\_interface\]
Configuration parameter for default sriov interface.

Upstream description:

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

- [default_sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/default_sriov_interface/#section)
- [sriov_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/sriov_interfaces/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
