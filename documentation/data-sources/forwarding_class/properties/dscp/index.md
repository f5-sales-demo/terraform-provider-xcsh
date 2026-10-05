---
page_title: "dscp"
subcategory: ""
description: "DSCP marking setting as per RFC 2475."
xcsh_docs: {"aliases": ["dscp"], "body_bytes": 3988, "body_sha256": "sha256:821dee8a6677a75262b19ea3f8fce070a8d3ffee09f7d3d336ef43dccc80d502", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forwarding_class:properties:dscp", "parent_id": "xcsh-docs:data-sources:forwarding_class:reference", "path": "documentation/data-sources/forwarding_class/properties/dscp/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2311202013321033-0100101033323120-3111331212200200-3133301132032033-3122320200333113-2201133101021101-2300230120331132-1111122123032022", "registry_path": "docs/guides/data-sources--forwarding_class--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dscp"], "schema_version": 1, "sections": [{"aliases": ["dscp drop precedence"], "anchor": "schema-dscp--drop_precedence", "description": "DSCP Assured forwarding drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop precedence value is taken from output of policer.", "document_id": "xcsh-docs:data-sources:forwarding_class:properties:dscp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dscp", "drop_precedence"], "syntax": "attribute", "type": "string"}, {"aliases": ["dscp dscp class"], "anchor": "schema-dscp--dscp_class", "description": "DSCP Precedence Level Values Best Effort service will GET any available bandwidth DSCP Class 1 service DSCP Class 2 service DSCP Class 3 service DSCP Class 4 service Express Forwarding is used for low latency traffic Control is used for routing traffic, not recommended Link Layer traffic like LACP or keepalive, not", "document_id": "xcsh-docs:data-sources:forwarding_class:properties:dscp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dscp", "dscp_class"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forwarding_class/properties/dscp/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "DSCP marking setting as per RFC 2475.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dscp

Breadcrumbs:

- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/)
- dscp

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: dscp, no\_marking, tos\_value; Default: no\_marking\] DSCP Marking setting. DSCP marking
setting as per RFC 2475.

Upstream description:

DSCP marking setting as per RFC 2475.

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

- [dscp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/dscp/#section)
- [no_marking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/no_marking/#section)
- [tos_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-tos_value)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-dscp--drop_precedence"></a>

### drop_precedence property

Type: `"string"`. Computed.

\[Enum: DSCP\_AF\_LOW|DSCP\_AF\_MEDIUM|DSCP\_AF\_HIGH|DSCP\_AF\_POLICER\] DSCP Assured forwarding
drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop
precedence value is taken from output of policer. Possible values are \`DSCP\_AF\_LOW\`,
\`DSCP\_AF\_MEDIUM\`, \`DSCP\_AF\_HIGH\`, \`DSCP\_AF\_POLICER\`.

Upstream description:

DSCP Assured forwarding drop precedence

DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop precedence
value is taken from output of policer.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "DSCP_AF_LOW",
    "DSCP_AF_MEDIUM",
    "DSCP_AF_HIGH",
    "DSCP_AF_POLICER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-dscp--dscp_class"></a>

### dscp_class property

Type: `"string"`. Computed.

\[Enum:
DSCP\_BEST\_EFFORT|DSCP\_CLASS1|DSCP\_CLASS2|DSCP\_CLASS3|DSCP\_CLASS4|DSCP\_EXPRESS\_FORWARDING|DSCP\_CONTROL\_L3|DSCP\_CONTROL\_L2\]
DSCP Precedence Level Values Best Effort service will GET any available bandwidth DSCP Class 1
service DSCP Class 2 service DSCP Class 3 service DSCP Class 4 service Express Forwarding is used
for low latency traffic Control is used for routing traffic, not recommended Link Layer traffic
like.. Possible values are \`DSCP\_BEST\_EFFORT\`, \`DSCP\_CLASS1\`, \`DSCP\_CLASS2\`,
\`DSCP\_CLASS3\`, \`DSCP\_CLASS4\`, \`DSCP\_EXPRESS\_FORWARDING\`, \`DSCP\_CONTROL\_L3\`,
\`DSCP\_CONTROL\_L2\`. Defaults to \`DSCP\_BEST\_EFFORT\`.

Upstream description:

DSCP Precedence Level Values

Best Effort service will GET any available bandwidth DSCP Class 1 service DSCP Class 2 service DSCP
Class 3 service DSCP Class 4 service Express Forwarding is used for low latency traffic Control is
used for routing traffic, not recommended Link Layer traffic like LACP or keepalive, not
recommended.

Receipt-pinned upstream constraints:

```json
{
  "default": "DSCP_BEST_EFFORT",
  "enum": [
    "DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/)
- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/)
