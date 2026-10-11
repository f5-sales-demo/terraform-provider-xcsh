---
page_title: "dh_group_set"
subcategory: ""
description: "Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile."
xcsh_docs: {"aliases": ["dh group set"], "body_bytes": 1980, "body_sha256": "sha256:d0042383454c89e5f051d9b7dc2b1fcb186d6795e14c97315ffaf6169eded7d7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike2:properties:dh_group_set", "parent_id": "xcsh-docs:resources:ike2:reference", "path": "documentation/resources/ike2/properties/dh_group_set/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1132030321233131-0112333220310223-0120130001103312-1033130032311312-1222102313212001-3232201300232232-0030223131101212-3011200112322331", "registry_path": "docs/guides/resources--ike2--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dh_group_set"], "schema_version": 1, "sections": [{"aliases": ["dh group set dh groups"], "anchor": "schema-dh_group_set--dh_groups", "description": "Group or collection configuration", "document_id": "xcsh-docs:resources:ike2:properties:dh_group_set", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dh_group_set", "dh_groups"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/properties/dh_group_set/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["ike2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dh_group_set

Breadcrumbs:

- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/)
- dh_group_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dh\_group\_set, disable\_pfs; Default: disable\_pfs\] Choose the acceptable Diffie
Hellman(DH) Group or Groups that you are willing to accept as part of this profile.

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

- [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/dh_group_set/#section)
- [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/disable_pfs/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dh_group_set {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-dh_group_set--dh_groups"></a>

### dh_groups property

Type: `["list", "string"]`. Optional.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Diffie Hellman Groups. Group or collection configuration. Possible values are
\`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`, \`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`,
\`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`, \`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`.
Defaults to \`DH\_GROUP\_DEFAULT\`.

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
