---
page_title: "gre.gre_parameters.tunnel_eps"
subcategory: ""
description: "gre.gre_parameters.tunnel_eps for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 5497, "body_sha256": "sha256:1aa8afa3d41b93cc68907e6464263328b9a5570b4bbd8479bc8b1e0dffb1a7e7", "canonical_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "child_ids": [], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "parent_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "path": "docs/guides/resources--external_connector--properties--gre--gre_parameters--tunnel_eps.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gre", "gre_parameters", "tunnel_eps"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/gre_parameters/tunnel_eps/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gre.gre_parameters.tunnel_eps for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# gre.gre_parameters.tunnel_eps

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
- [gre](resources--external_connector--properties--gre.md)
- [gre.gre_parameters](resources--external_connector--properties--gre--gre_parameters.md)
- gre.gre_parameters.tunnel_eps

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Configure tunnel parameters, source, destination, IP addresses.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

CE site is composed of multiple nodes. Choose a node that will be part of this external connection.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [gre.gre_parameters](resources--external_connector--properties--gre--gre_parameters.md)
- [xcsh_external_connector](../resources/external_connector.md)
