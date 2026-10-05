---
page_title: "re_select"
subcategory: ""
description: "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s)."
xcsh_docs: {"aliases": ["re select"], "body_bytes": 2340, "body_sha256": "sha256:841e3766785c60eccf342cd49a43727f14d879f34185ae065321f61bf5581d65", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/re_select/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [{"anchor": "schema-re_select--specific_geography", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:geo_proximity,specific_geography", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "type": "conflicts"}, {"anchor": "schema-re_select--specific_geography", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:specific_geography,specific_re", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:geo_proximity,specific_geography", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:geo_proximity,specific_re", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:geo_proximity,specific_re", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:specific_geography,specific_re", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["re_select"], "schema_version": 1, "sections": [{"aliases": ["re select geo proximity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "geo_proximity"], "syntax": "attribute", "type": "object"}, {"aliases": ["re select specific geography"], "anchor": "schema-re_select--specific_geography", "description": "Geographic selection for the site's Regional Edge connections.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "specific_geography"], "syntax": "attribute", "type": "string"}, {"aliases": ["re select specific re"], "anchor": "section", "description": "Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_select", "specific_re"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/re_select/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_select

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- re_select

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("geo_proximity",
    "specific_geography"),
  validators.ConflictingObjectAttributes("geo_proximity",
    "specific_re"),
  validators.ConflictingObjectAttributes("specific_geography",
    "specific_re")}
```

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

Terraform syntax:

```terraform
re_select {
  # Configure direct properties listed below.
}
```

## Direct properties

- [geo_proximity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/re_select/geo_proximity/): complete subsection reference.

<a id="schema-re_select--specific_geography"></a>

### specific_geography property

Type: `"string"`. Optional.

Geographic selection for the site's Regional Edge connections.

- [specific_re](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/re_select/specific_re/): complete subsection reference.

## Next pages

- [re_select.geo_proximity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/re_select/geo_proximity/)
- [re_select.specific_re](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/re_select/specific_re/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
