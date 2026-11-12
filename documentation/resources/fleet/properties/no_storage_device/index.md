---
page_title: "no_storage_device"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["no storage device"], "body_bytes": 1312, "body_sha256": "sha256:60a6209265933511f7a30cb65c099a1a744f2547d4b57661f44cc01137a687c8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:no_storage_device", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/no_storage_device/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-1323333231331210-1331121020203033-1001232023202221-0323221132032011-2122012311233313-0333321001230223-1230201132012120-3320022013311213", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["no_storage_device"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/no_storage_device/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["fleetCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_storage_device

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- no_storage_device

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_storage\_device, storage\_device\_list; Default: no\_storage\_device\] Configuration
parameter for no storage device.

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

- [no_storage_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/no_storage_device/#section)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_storage_device = {}
```

This is an empty object or choice marker. It has no direct properties.
