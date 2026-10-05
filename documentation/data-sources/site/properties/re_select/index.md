---
page_title: "re_select"
subcategory: "Infrastructure"
description: "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s)."
xcsh_docs: {"aliases": ["re select"], "body_bytes": 1482, "body_sha256": "sha256:c92249e8e4ac8ad6dd820006609ae72fa0f0360aa87448aea84fa69dce855c44", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:re_select:geo_proximity", "xcsh-docs:data-sources:site:properties:re_select:specific_re"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:re_select", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/re_select/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2301320020302102-2123113003103212-2301203100031011-2111021021231131-3112322001232300-1301212000120020-2302213323322121-3003032212333033", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["re_select"], "schema_version": 1, "sections": [{"aliases": ["re select geo proximity"], "anchor": "section", "description": "Configuration parameter for geo proximity.", "document_id": "xcsh-docs:data-sources:site:properties:re_select:geo_proximity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "geo_proximity"], "syntax": "attribute", "type": "object"}, {"aliases": ["re select specific geography"], "anchor": "schema-re_select--specific_geography", "description": "Geographic selection for the site's Regional Edge connections.", "document_id": "xcsh-docs:data-sources:site:properties:re_select", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "specific_geography"], "syntax": "attribute", "type": "string"}, {"aliases": ["re select specific re"], "anchor": "section", "description": "Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs.", "document_id": "xcsh-docs:data-sources:site:properties:re_select:specific_re", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_select", "specific_re"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/re_select/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_select

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- re_select

<a id="section"></a>

Type: `"single"`. Computed.

Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).

## Direct properties

- [geo_proximity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/re_select/geo_proximity/): complete subsection reference.

<a id="schema-re_select--specific_geography"></a>

### specific_geography property

Type: `"string"`. Computed.

Geographic selection for the site's Regional Edge connections.

- [specific_re](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/re_select/specific_re/): complete subsection reference.

## Next pages

- [re_select.geo_proximity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/re_select/geo_proximity/)
- [re_select.specific_re](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/re_select/specific_re/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
