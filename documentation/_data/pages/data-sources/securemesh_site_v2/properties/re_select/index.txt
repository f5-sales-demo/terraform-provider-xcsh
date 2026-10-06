---
page_title: "re_select"
subcategory: ""
description: "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s)."
xcsh_docs: {"aliases": ["re select"], "body_bytes": 1344, "body_sha256": "sha256:1dedc37d160739753afe515bf9789d97613eb3a6b835aeaf1d22b25a53feb3a3", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:geo_proximity", "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:specific_re"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/re_select/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2121113003333022-2122100030330221-0011311122211220-2010201213300022-0132320300020113-3222232131112223-3232100101011113-0213130312020332", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_select"], "schema_version": 1, "sections": [{"aliases": ["re select geo proximity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:geo_proximity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "geo_proximity"], "syntax": "attribute", "type": "object"}, {"aliases": ["re select specific geography"], "anchor": "schema-re_select--specific_geography", "description": "Geographic selection for the site's Regional Edge connections.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "specific_geography"], "syntax": "attribute", "type": "string"}, {"aliases": ["re select specific re"], "anchor": "section", "description": "Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:specific_re", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_select", "specific_re"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/re_select/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_select

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- re_select

<a id="section"></a>

Type: `"single"`. Computed.

Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-re_selection_choice": "[\"geo_proximity\", \"specific_geography\", \"specific_re\"]"
}
```

## Direct properties

- [geo_proximity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/re_select/geo_proximity/): complete subsection reference.

<a id="schema-re_select--specific_geography"></a>

### specific_geography property

Type: `"string"`. Computed.

Geographic selection for the site's Regional Edge connections.

- [specific_re](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/re_select/specific_re/): complete subsection reference.
