---
page_title: "disable_https_management"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable https management"], "body_bytes": 1388, "body_sha256": "sha256:81eaa7735aeb13270de8da8e4d6cba005d73d02c02b14b8e982256604aa08690", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:disable_https_management", "parent_id": "xcsh-docs:resources:nfv_service:reference", "path": "documentation/resources/nfv_service/properties/disable_https_management/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1213021122302111-2303132302331132-2033120022132312-0232210233303323-2001032101312213-3131012003332023-0101021121323031-2100213002233100", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_https_management"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/disable_https_management/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_https_management

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- disable_https_management

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_https\_management, https\_management; Default: disable\_https\_management\]
Configuration parameter for disable https management.

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

- [disable_https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/disable_https_management/#section)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_https_management = {}
```

This is an empty object or choice marker. It has no direct properties.
