---
page_title: "dedicated_management_interface"
subcategory: ""
description: "Dedicated Interface Configuration."
xcsh_docs: {"aliases": ["dedicated management interface"], "body_bytes": 4536, "body_sha256": "sha256:bfa2786ae29042ac95ee327c8f4b30aeebecf24e61f5ee7c6df4473317ef56a2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:dedicated_management_interface:cluster"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:dedicated_management_interface", "parent_id": "xcsh-docs:data-sources:network_interface:reference", "path": "documentation/data-sources/network_interface/properties/dedicated_management_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0031110031232313-1231101221020103-3111000020011113-3113300131113000-1123010222111130-0023230223301022-2202230222001100-3100123330301302", "registry_path": "docs/guides/data-sources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dedicated_management_interface"], "schema_version": 1, "sections": [{"aliases": ["dedicated management interface cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_management_interface:cluster", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_management_interface", "cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["dedicated management interface device"], "anchor": "schema-dedicated_management_interface--device", "description": "Name of the device for which interface is configured.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_management_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_management_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["dedicated management interface mtu"], "anchor": "schema-dedicated_management_interface--mtu", "description": "Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between 512 and 9000.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_management_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_management_interface", "mtu"], "syntax": "attribute", "type": "number"}, {"aliases": ["dedicated management interface node"], "anchor": "schema-dedicated_management_interface--node", "description": "Exclusive with Configuration will apply to a device on the given node of the site.", "document_id": "xcsh-docs:data-sources:network_interface:properties:dedicated_management_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dedicated_management_interface", "node"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/dedicated_management_interface/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Dedicated Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dedicated_management_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- dedicated_management_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dedicated management interface.

Upstream description:

Dedicated Interface Configuration.

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

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/cluster/): complete subsection reference.

<a id="schema-dedicated_management_interface--device"></a>

### device property

Type: `"string"`. Computed.

Name of the device for which interface is configured.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node of the site.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [dedicated_management_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/dedicated_management_interface/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
