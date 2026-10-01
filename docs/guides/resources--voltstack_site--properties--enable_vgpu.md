---
page_title: "enable_vgpu"
subcategory: ""
description: "enable_vgpu for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3845, "body_sha256": "sha256:4320bc72ff15e6ebe21d3c920e497d8d3c451f501a8861cd8004a8120581ca09", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:enable_vgpu", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:enable_vgpu", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "docs/guides/resources--voltstack_site--properties--enable_vgpu.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_vgpu"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/enable_vgpu/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_vgpu for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_vgpu

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
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

Upstream description:

Set feature to be enabled

Operate with a degraded vGPU performance Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation
Enable NVIDIA Virtual Compute Server.

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

Upstream description:

Set License Server Address.

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

Upstream description:

Set License Server port number.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [Property reference](resources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
