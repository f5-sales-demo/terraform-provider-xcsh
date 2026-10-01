---
page_title: "dscp"
subcategory: ""
description: "dscp for xcsh_forwarding_class."
xcsh_docs: {"aliases": [], "body_bytes": 4222, "body_sha256": "sha256:ae671a07cf3352b8126e258f68880b48b96b8759b7b1e13c26a3e940cf20c52f", "canonical_id": "xcsh-docs:resources:forwarding_class:properties:dscp", "child_ids": [], "collection_id": "xcsh-docs:resources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:resources:forwarding_class:properties:dscp", "parent_id": "xcsh-docs:resources:forwarding_class:reference", "path": "docs/guides/resources--forwarding_class--properties--dscp.md", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dscp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forwarding_class/properties/dscp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dscp for xcsh_forwarding_class.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dscp

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md)
- [Property reference](resources--forwarding_class--reference.md)
- dscp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [dscp](resources--forwarding_class--properties--dscp.md#section)
- [no_marking](resources--forwarding_class--properties--no_marking.md#section)
- [tos_value](resources--forwarding_class--reference.md#schema-tos_value)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dscp {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-dscp--drop_precedence"></a>

### drop_precedence property

Type: `"string"`. Optional.

\[Enum: DSCP\_AF\_LOW|DSCP\_AF\_MEDIUM|DSCP\_AF\_HIGH|DSCP\_AF\_POLICER\] DSCP Assured forwarding
drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop
precedence value is taken from output of policer. Possible values are \`DSCP\_AF\_LOW\`,
\`DSCP\_AF\_MEDIUM\`, \`DSCP\_AF\_HIGH\`, \`DSCP\_AF\_POLICER\`.

Upstream description:

DSCP Assured forwarding drop precedence

DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop precedence
value is taken from output of policer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DSCP_AF_LOW",
    "DSCP_AF_MEDIUM",
    "DSCP_AF_HIGH",
    "DSCP_AF_POLICER"),
}
```

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

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"),
}
```

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

- [Property reference](resources--forwarding_class--reference.md)
- [xcsh_forwarding_class](../resources/forwarding_class.md)
