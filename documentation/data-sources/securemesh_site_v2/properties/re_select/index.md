---
page_title: "re_select"
subcategory: ""
description: "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s)."
xcsh_docs: {"aliases": ["re select"], "body_bytes": 1904, "body_sha256": "sha256:b423b48fd9bc634de96a62791129a9d6ad02b160365f3725c82da58e7ff34def", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:geo_proximity", "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:specific_re"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/re_select/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2121113003333022-2122100030330221-0011311122211220-2010201213300022-0132320300020113-3222232131112223-3232100101011113-0213130312020332", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_select"], "schema_version": 1, "sections": [{"aliases": ["geo proximity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:geo_proximity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "geo_proximity"], "syntax": "attribute", "type": "object"}, {"aliases": ["specific geography"], "anchor": "schema-re_select--specific_geography", "description": "Geographic selection for the site's Regional Edge connections.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "specific_geography"], "syntax": "attribute", "type": "string"}, {"aliases": ["specific re"], "anchor": "section", "description": "Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:specific_re", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_select", "specific_re"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/re_select/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [re_select.geo_proximity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/re_select/geo_proximity/)
- [re_select.specific_re](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/re_select/specific_re/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
