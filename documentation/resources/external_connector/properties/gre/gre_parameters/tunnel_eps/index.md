---
page_title: "gre.gre_parameters.tunnel_eps"
subcategory: ""
description: "Configure tunnel parameters, source, destination, IP addresses."
xcsh_docs: {"aliases": ["gre gre parameters tunnel eps"], "body_bytes": 5537, "body_sha256": "sha256:c7c98920a4e66eafc761b438c1c41f2f9957474aa93894af889d0087c32e5b57", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "parent_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "path": "documentation/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3010112031211331-0301021120032210-0212123101022030-3302111100331202-2220301013332022-1320010300030001-0003013201312303-1013110321001212", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [{"anchor": "schema-gre--gre_parameters--tunnel_eps--interface", "enforcement": "provider-schema", "group": "gre.gre_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}, {"anchor": "schema-gre--gre_parameters--tunnel_eps--local_tunnel_ip", "enforcement": "provider-schema", "group": "gre.gre_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}, {"anchor": "schema-gre--gre_parameters--tunnel_eps--node", "enforcement": "provider-schema", "group": "gre.gre_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}, {"anchor": "schema-gre--gre_parameters--tunnel_eps--remote_tunnel_ip", "enforcement": "provider-schema", "group": "gre.gre_parameters.tunnel_eps:RequiredListObjectAttributes:interface,local_tunnel_ip,node,remote_tunnel_ip", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gre", "gre_parameters", "tunnel_eps"], "schema_version": 1, "sections": [{"aliases": ["gre gre parameters tunnel eps interface"], "anchor": "schema-gre--gre_parameters--tunnel_eps--interface", "description": "For the chosen node, specify the interface that will be the tunnel source.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "tunnel_eps", "interface"], "syntax": "attribute", "type": "string"}, {"aliases": ["gre gre parameters tunnel eps local tunnel ip"], "anchor": "schema-gre--gre_parameters--tunnel_eps--local_tunnel_ip", "description": "For a particular tunnel on a node, specify the local tunnel IP Address i.e. The IP address of the tunnel on the CE node itself and a subnet prefix length.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "tunnel_eps", "local_tunnel_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["gre gre parameters tunnel eps node"], "anchor": "schema-gre--gre_parameters--tunnel_eps--node", "description": "A CE site is composed of multiple nodes. Choose a node that will be part of this external connection.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "tunnel_eps", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["gre gre parameters tunnel eps remote tunnel ip"], "anchor": "schema-gre--gre_parameters--tunnel_eps--remote_tunnel_ip", "description": "For a particular tunnel on a node, specify the remote tunnel IP Address i.e. The IP address of the tunnel on the remote gateway and a subnet prefix length.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gre", "gre_parameters", "tunnel_eps", "remote_tunnel_ip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Configure tunnel parameters, source, destination, IP addresses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["external_connectorCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre.gre_parameters.tunnel_eps

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/)
- [gre.gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/)
- gre.gre_parameters.tunnel_eps

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Configure tunnel parameters, source, destination, IP addresses.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("interface",
    "local_tunnel_ip",
    "node",
    "remote_tunnel_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tunnel_eps {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-gre--gre_parameters--tunnel_eps--interface"></a>

### interface property

Type: `"string"`. Optional.

For the chosen node, specify the interface that will be the tunnel source.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
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

<a id="schema-gre--gre_parameters--tunnel_eps--local_tunnel_ip"></a>

### local_tunnel_ip property

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the local tunnel IP Address i.e. The IP address of the
tunnel on the CE node itself and a subnet prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="schema-gre--gre_parameters--tunnel_eps--node"></a>

### node property

Type: `"string"`. Optional.

A CE site is composed of multiple nodes. Choose a node that will be part of this external
connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
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

<a id="schema-gre--gre_parameters--tunnel_eps--remote_tunnel_ip"></a>

### remote_tunnel_ip property

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the remote tunnel IP Address i.e. The IP address of the
tunnel on the remote gateway and a subnet prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```
