---
page_title: "re_select"
subcategory: ""
description: "re_select for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1932, "body_sha256": "sha256:f3ca65fe8eb443166bd82f15bbc7ba836f11b9df11768c2b3755b96513f818e1", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:re_select:geo_proximity", "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--re_select.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_select"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/re_select/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_select for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_select

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
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

- [geo_proximity](resources--securemesh_site_v2--properties--re_select--geo_proximity.md): complete subsection reference.

<a id="schema-re_select--specific_geography"></a>

### specific_geography property

Type: `"string"`. Optional.

Geographic selection for the site's Regional Edge connections.

- [specific_re](resources--securemesh_site_v2--properties--re_select--specific_re.md): complete subsection reference.

## Next pages

- [re_select.geo_proximity](resources--securemesh_site_v2--properties--re_select--geo_proximity.md)
- [re_select.specific_re](resources--securemesh_site_v2--properties--re_select--specific_re.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
