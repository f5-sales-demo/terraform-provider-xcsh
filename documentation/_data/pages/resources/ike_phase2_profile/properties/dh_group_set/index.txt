---
page_title: "dh_group_set"
subcategory: ""
description: "Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile."
xcsh_docs: {"aliases": ["dh group set"], "body_bytes": 2945, "body_sha256": "sha256:107d6842e854adefe726d4f7be1e95962c2d8cea5a9be64b9d8139b6b1f1ce54", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase2_profile:properties:dh_group_set", "parent_id": "xcsh-docs:resources:ike_phase2_profile:reference", "path": "documentation/resources/ike_phase2_profile/properties/dh_group_set/index.md", "product": "distributed-cloud", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0133212113121220-2130201020223302-2301133222102133-0032320323013023-1212213333211310-3200230010202120-2303020033323020-2303101213112202", "registry_path": "docs/guides/resources--ike_phase2_profile--reference--group-001.md", "relationships": [{"anchor": "schema-dh_group_set--dh_groups", "enforcement": "provider-schema", "group": "dh_group_set:RequiredObjectAttributes:dh_groups", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:ike_phase2_profile:properties:dh_group_set", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["dh_group_set"], "schema_version": 1, "sections": [{"aliases": ["dh groups"], "anchor": "schema-dh_group_set--dh_groups", "description": "Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile.", "document_id": "xcsh-docs:resources:ike_phase2_profile:properties:dh_group_set", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dh_group_set", "dh_groups"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase2_profile/properties/dh_group_set/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of this profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dh_group_set

Breadcrumbs:

- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/)
- dh_group_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dh\_group\_set, disable\_pfs; Default: disable\_pfs\] Choose the acceptable Diffie
Hellman(DH) Group or Groups that you are willing to accept as part of this profile.

Upstream description:

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dh_groups")}
```

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

- [dh_group_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/dh_group_set/#section)
- [disable_pfs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/disable_pfs/#section)

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
Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile. Possible values are \`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`,
\`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`, \`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`,
\`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`. Defaults to \`DH\_GROUP\_DEFAULT\`.

Upstream description:

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/)
- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
