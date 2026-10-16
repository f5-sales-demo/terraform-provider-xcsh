---
page_title: "no_panic_threshold"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["no panic threshold"], "body_bytes": 1316, "body_sha256": "sha256:38b1f08511cba1194ba8c0d3be6884cb3ab04f73f97a32125825b442189aca20", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:no_panic_threshold", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/no_panic_threshold/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3003220312211122-0130132300200133-0003203213303111-3231033101300201-1323232130312102-0321223312002221-2012322311301030-2120333223120322", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["no_panic_threshold"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/no_panic_threshold/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["clusterCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_panic_threshold

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- no_panic_threshold

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_panic\_threshold, panic\_threshold; Default: no\_panic\_threshold\] Configuration
parameter for no panic threshold.

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

- [no_panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/no_panic_threshold/#section)
- [panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-panic_threshold)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_panic_threshold = {}
```

This is an empty object or choice marker. It has no direct properties.
