---
page_title: "disabled"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disabled"], "body_bytes": 1193, "body_sha256": "sha256:ec45714ed21309608d7c916cc24e1c181e60b819c781ec59b00158929265e3de", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_k8s:properties:disabled", "parent_id": "xcsh-docs:resources:virtual_k8s:reference", "path": "documentation/resources/virtual_k8s/properties/disabled/index.md", "product": "distributed-cloud", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1000110211130130-2231200300130313-0012222033232021-3302311001032023-2030022233131330-2020303113200313-0130032100202113-2320003221333033", "registry_path": "docs/guides/resources--virtual_k8s--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_k8s/properties/disabled/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disabled

Breadcrumbs:

- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/properties/)
- disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disabled, isolated\] Enable this option

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

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/properties/disabled/#section)
- [isolated](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/properties/isolated/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
