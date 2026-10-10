---
page_title: "disable_path_normalize"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable path normalize"], "body_bytes": 1360, "body_sha256": "sha256:7aab18a262216241bf2400de3786f6ac38c13792ba7e20c4645ec32b2b1545c4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:disable_path_normalize", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "documentation/resources/virtual_host/properties/disable_path_normalize/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0002023231210102-0232311230110330-2231021201011211-3000222000033100-3121230313301113-1303121121223033-1322331301232132-0220030301200000", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_path_normalize"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/disable_path_normalize/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_path_normalize

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- disable_path_normalize

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_path\_normalize, enable\_path\_normalize; Default: disable\_path\_normalize\]
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

OneOf alternatives in this subsection:

- [disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/disable_path_normalize/#section)
- [enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/enable_path_normalize/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.
