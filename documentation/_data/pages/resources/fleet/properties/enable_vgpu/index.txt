---
page_title: "enable_vgpu"
subcategory: ""
description: "Licensing configuration for NVIDIA vGPU."
xcsh_docs: {"aliases": ["enable vgpu"], "body_bytes": 3491, "body_sha256": "sha256:9b21fffec4375d96efd790fedd0404948592adf835bfa1257bb13f83be89a922", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:enable_vgpu", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/enable_vgpu/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0023220101031001-1331310302022102-0213011031321233-3112103310211311-1212211111023103-3221232112323313-3332022200013311-1333313210022010", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [{"anchor": "schema-enable_vgpu--server_port", "enforcement": "provider-schema", "group": "enable_vgpu:RequiredObjectAttributes:server_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:enable_vgpu", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_vgpu"], "schema_version": 1, "sections": [{"aliases": ["enable vgpu feature type"], "anchor": "schema-enable_vgpu--feature_type", "description": "Set feature to be enabled Operate with a degraded vGPU performance Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation Enable NVIDIA Virtual Compute Server.", "document_id": "xcsh-docs:resources:fleet:properties:enable_vgpu", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_vgpu", "feature_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable vgpu server address"], "anchor": "schema-enable_vgpu--server_address", "description": "Set License Server Address.", "document_id": "xcsh-docs:resources:fleet:properties:enable_vgpu", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_vgpu", "server_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable vgpu server port"], "anchor": "schema-enable_vgpu--server_port", "description": "Set License Server port number.", "document_id": "xcsh-docs:resources:fleet:properties:enable_vgpu", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_vgpu", "server_port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/enable_vgpu/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Licensing configuration for NVIDIA vGPU.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_vgpu

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- enable_vgpu

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Licensing configuration for NVIDIA vGPU.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server_port")}
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

Terraform syntax:

```terraform
enable_vgpu {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-enable_vgpu--feature_type"></a>

### feature_type property

Type: `"string"`. Optional.

\[Enum: UNLICENSED|VGPU|VWS|VCS\] Set feature to be enabled Operate with a degraded vGPU performance
Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation Enable NVIDIA Virtual Compute Server.
Possible values are \`UNLICENSED\`, \`VGPU\`, \`VWS\`, \`VCS\`. Defaults to \`UNLICENSED\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNLICENSED",
    "VGPU",
    "VWS",
    "VCS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNLICENSED",
  "enum": [
    "UNLICENSED",
    "VGPU",
    "VWS",
    "VCS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-enable_vgpu--server_address"></a>

### server_address property

Type: `"string"`. Optional.

License Server Address. Set License Server Address.

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
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

<a id="schema-enable_vgpu--server_port"></a>

### server_port property

Type: `"number"`. Optional.

License Server Port Number. Set License Server port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```
