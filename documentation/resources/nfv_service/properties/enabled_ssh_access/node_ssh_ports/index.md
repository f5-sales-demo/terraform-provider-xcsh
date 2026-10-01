---
page_title: "enabled_ssh_access.node_ssh_ports"
subcategory: ""
description: "enabled_ssh_access.node_ssh_ports for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 4486, "body_sha256": "sha256:b7538202e30005f32540abe2816ea2e57d5ad707880a9cdcfb5d34e2eb27d6ed", "child_ids": [], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "parent_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access", "path": "documentation/resources/nfv_service/properties/enabled_ssh_access/node_ssh_ports/index.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["enabled_ssh_access", "node_ssh_ports"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/enabled_ssh_access/node_ssh_ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enabled_ssh_access.node_ssh_ports for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enabled_ssh_access.node_ssh_ports

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [enabled_ssh_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/)
- enabled_ssh_access.node_ssh_ports

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Management Node SSH Port. Enter TCP port and node name per node.

Upstream description:

Enter TCP port and node name per node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("node_name",
    "ssh_port")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.repeated.max_items": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2"
  }
}
```

Terraform syntax:

```terraform
node_ssh_ports {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-enabled_ssh_access--node_ssh_ports--node_name"></a>

### node_name property

Type: `"string"`. Optional.

Node name will be used to match a particular node with the desired TCP port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-enabled_ssh_access--node_ssh_ports--ssh_port"></a>

### ssh_port property

Type: `"number"`. Optional.

SSH Port. Enter TCP port per node.

Upstream description:

Enter TCP port per node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1024, 65535),
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
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [enabled_ssh_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
