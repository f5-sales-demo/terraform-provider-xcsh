---
page_title: "ipsec.ipsec_tunnel_parameters.tunnel_eps"
subcategory: ""
description: "Configure tunnel parameters, local and remote IP addresses."
xcsh_docs: {"aliases": ["ipsec ipsec tunnel parameters tunnel eps"], "body_bytes": 5279, "body_sha256": "sha256:7d1d5a65c3cfcef5bacab993a45e4c72dfbcdc0b4c490ae5ad4286c4de6714b7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "parent_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters", "path": "documentation/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1203320103000013-1201323323000200-3123321232320022-3310203201110313-3330131333201213-3210222223331222-2023033222220310-3303012310100120", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ipsec_tunnel_parameters", "tunnel_eps"], "schema_version": 1, "sections": [{"aliases": ["ipsec ipsec tunnel parameters tunnel eps interface"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--interface", "description": "For the chosen node, specify the interface that will be the tunnel source.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "tunnel_eps", "interface"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ipsec tunnel parameters tunnel eps local tunnel ip"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--local_tunnel_ip", "description": "For a particular tunnel on a node, specify the local tunnel IP Address i.e. The IP address of the tunnel on the CE node itself and a subnet prefix length.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "tunnel_eps", "local_tunnel_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ipsec tunnel parameters tunnel eps node"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--node", "description": "A CE site is composed of multiple nodes. Choose a node that will be part of this external connection.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "tunnel_eps", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipsec ipsec tunnel parameters tunnel eps remote tunnel ip"], "anchor": "schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--remote_tunnel_ip", "description": "For a particular tunnel on a node, specify the remote tunnel IP Address i.e. The IP address of the tunnel on the remote gateway and a subnet prefix length.", "document_id": "xcsh-docs:data-sources:external_connector:properties:ipsec:ipsec_tunnel_parameters:tunnel_eps", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipsec", "ipsec_tunnel_parameters", "tunnel_eps", "remote_tunnel_ip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/tunnel_eps/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configure tunnel parameters, local and remote IP addresses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["external_connectorCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ipsec_tunnel_parameters.tunnel_eps

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/)
- [ipsec.ipsec_tunnel_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/ipsec_tunnel_parameters/)
- ipsec.ipsec_tunnel_parameters.tunnel_eps

<a id="section"></a>

Type: `"list"`. Computed.

Configure tunnel parameters, local and remote IP addresses.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

## Direct properties

<a id="schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--interface"></a>

### interface property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--local_tunnel_ip"></a>

### local_tunnel_ip property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--node"></a>

### node property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-ipsec--ipsec_tunnel_parameters--tunnel_eps--remote_tunnel_ip"></a>

### remote_tunnel_ip property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
