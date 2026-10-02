---
page_title: "default_subset"
subcategory: ""
description: "List of key-value pairs that define default subset. This subset can be referred in fallback_policy which gets used when route specifies no metadata or no subset matching the metadata exists."
xcsh_docs: {"aliases": ["default subset"], "body_bytes": 1469, "body_sha256": "sha256:14f528835d90abd68b35c9ba5805776fb525fdfc1abc0d0db1e5ef38555428e8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:default_subset", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "documentation/data-sources/cluster/properties/default_subset/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3211100033212002-3323310310011303-3231020001323310-3133321001312311-3031123102313023-3321012012100132-3231002312200331-1303113213312131", "registry_path": "docs/guides/data-sources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_subset"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/default_subset/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of key-value pairs that define default subset. This subset can be referred in fallback_policy which gets used when route specifies no metadata or no subset matching the metadata exists.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_subset

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- default_subset

<a id="section"></a>

Type: `"single"`. Computed.

List of key-value pairs that define default subset. This subset can be referred in fallback\_policy
which gets used when route specifies no metadata or no subset matching the metadata exists.

Upstream description:

List of key-value pairs that define default subset. This subset can be referred in fallback\_policy
which gets used when route specifies no metadata or no subset matching the metadata exists.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  }
}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
