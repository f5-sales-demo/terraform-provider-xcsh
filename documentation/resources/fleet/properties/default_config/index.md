---
page_title: "default_config"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default config"], "body_bytes": 1380, "body_sha256": "sha256:e7207ad2415ad0dc0d76b18fadb0089817f19eec459cf7723e34e492cb8eb057", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:default_config", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/default_config/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-3310002022222312-0033100302223102-0112021232310302-1233201223101302-0221021103321211-0101113223023222-3000011131123213-1121030200103313", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_config"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/default_config/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["fleetCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_config

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- default_config

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_config, device\_list, interface\_list; Default: default\_config\] Enable this
option

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

- [default_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/default_config/#section)
- [device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/#section)
- [interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/interface_list/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_config = {}
```

This is an empty object or choice marker. It has no direct properties.
