---
page_title: "dedicated_management_interface"
subcategory: ""
description: "Dedicated Interface Configuration."
xcsh_docs: {"aliases": ["dedicated management interface"], "body_bytes": 5398, "body_sha256": "sha256:e8ebd7963e82b6f040451058f33e27a850361942b0be364491961a469c6b28f6", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:dedicated_management_interface:cluster"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface", "parent_id": "xcsh-docs:resources:network_interface:reference", "path": "documentation/resources/network_interface/properties/dedicated_management_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2133212122221100-0000122231110223-2300111023113203-0222121110123023-1311302001131223-3002111002332020-1002010011031222-2022311230202232", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [{"anchor": "schema-dedicated_management_interface--node", "enforcement": "provider-schema", "group": "dedicated_management_interface:ConflictingObjectAttributes:cluster,node", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dedicated_management_interface:ConflictingObjectAttributes:cluster,node", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface:cluster", "type": "conflicts"}, {"anchor": "schema-dedicated_management_interface--device", "enforcement": "provider-schema", "group": "dedicated_management_interface:RequiredObjectAttributes:device", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["dedicated_management_interface"], "schema_version": 1, "sections": [{"aliases": ["cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface:cluster", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_management_interface", "cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["device"], "anchor": "schema-dedicated_management_interface--device", "description": "Name of the device for which interface is configured.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_management_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["mtu"], "anchor": "schema-dedicated_management_interface--mtu", "description": "Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between 512 and 9000.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_management_interface", "mtu"], "syntax": "attribute", "type": "number"}, {"aliases": ["node"], "anchor": "schema-dedicated_management_interface--node", "description": "Exclusive with Configuration will apply to a device on the given node of the site.", "document_id": "xcsh-docs:resources:network_interface:properties:dedicated_management_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_management_interface", "node"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/dedicated_management_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Dedicated Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dedicated_management_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- dedicated_management_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dedicated management interface.

Upstream description:

Dedicated Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node")}
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
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]"
}
```

Terraform syntax:

```terraform
dedicated_management_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_management_interface/cluster/): complete subsection reference.

<a id="schema-dedicated_management_interface--device"></a>

### device property

Type: `"string"`. Optional.

Name of the device for which interface is configured.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-dedicated_management_interface--mtu"></a>

### mtu property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="schema-dedicated_management_interface--node"></a>

### node property

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [dedicated_management_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_management_interface/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
