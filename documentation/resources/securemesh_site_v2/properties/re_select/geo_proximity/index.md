---
page_title: "re_select.geo_proximity"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["re select geo proximity"], "body_bytes": 1017, "body_sha256": "sha256:04441eb4bdab57cf4e4504cc03c4641d837d9ddf8d1ce25a44e9bcec1624e4b3", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "path": "documentation/resources/securemesh_site_v2/properties/re_select/geo_proximity/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3021321102110221-1110320033102320-2311303220231220-1200120002001130-1321022332032201-0033232102033300-3223013231331303-3001230322223100", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_select", "geo_proximity"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/re_select/geo_proximity/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_select.geo_proximity

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [re_select](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/re_select/)
- re_select.geo_proximity

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for geo proximity.

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

Terraform syntax:

```terraform
geo_proximity = {}
```

This is an empty object or choice marker. It has no direct properties.
