---
page_title: "dh_group_set"
subcategory: ""
description: "Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile."
xcsh_docs: {"aliases": ["dh group set"], "body_bytes": 2385, "body_sha256": "sha256:3b5d74e3999c3b2ffbc56e96b64e370dfa2beb53cb21e4f0950ad6ae01a8d395", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike2:properties:dh_group_set", "parent_id": "xcsh-docs:resources:ike2:reference", "path": "documentation/resources/ike2/properties/dh_group_set/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1132030321233131-0112333220310223-0120130001103312-1033130032311312-1222102313212001-3232201300232232-0030223131101212-3011200112322331", "registry_path": "docs/guides/resources--ike2--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dh_group_set"], "schema_version": 1, "sections": [{"aliases": ["dh groups"], "anchor": "schema-dh_group_set--dh_groups", "description": "Group or collection configuration", "document_id": "xcsh-docs:resources:ike2:properties:dh_group_set", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dh_group_set", "dh_groups"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/properties/dh_group_set/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile.

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

Upstream description:

Group or collection configuration

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/properties/)
- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike2/)
