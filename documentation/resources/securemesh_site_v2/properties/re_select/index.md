---
page_title: "re_select"
subcategory: ""
description: "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s)."
xcsh_docs: {"aliases": ["re select"], "body_bytes": 1822, "body_sha256": "sha256:d10a89c094e051e10f42125ba3660e9f99a65846db54d9e93bf9a5fd8caca904", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/re_select/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-016.md", "relationships": [{"anchor": "schema-re_select--specific_geography", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:geo_proximity,specific_geography", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "type": "conflicts"}, {"anchor": "schema-re_select--specific_geography", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:specific_geography,specific_re", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:geo_proximity,specific_geography", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:geo_proximity,specific_re", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:geo_proximity,specific_re", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_select:ConflictingObjectAttributes:specific_geography,specific_re", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["re_select"], "schema_version": 1, "sections": [{"aliases": ["re select geo proximity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "geo_proximity"], "syntax": "attribute", "type": "object"}, {"aliases": ["re select specific geography"], "anchor": "schema-re_select--specific_geography", "description": "Geographic selection for the site's Regional Edge connections.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_select", "specific_geography"], "syntax": "attribute", "type": "string"}, {"aliases": ["re select specific re"], "anchor": "section", "description": "Select specific REs. This is useful when a site needs to deterministically connect to a set of REs. A site will always be connected to 2 REs.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_select", "specific_re"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/re_select/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
EnumExtractionComplete: false
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
