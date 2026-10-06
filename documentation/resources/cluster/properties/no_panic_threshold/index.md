---
page_title: "no_panic_threshold"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["no panic threshold"], "body_bytes": 1316, "body_sha256": "sha256:38b1f08511cba1194ba8c0d3be6884cb3ab04f73f97a32125825b442189aca20", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:no_panic_threshold", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/no_panic_threshold/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3003220312211122-0130132300200133-0003203213303111-3231033101300201-1323232130312102-0321223312002221-2012322311301030-2120333223120322", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["no_panic_threshold"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/no_panic_threshold/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
