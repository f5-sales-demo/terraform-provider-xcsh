---
page_title: "palo_alto_fw_service.panorama_server"
subcategory: ""
description: "palo_alto_fw_service.panorama_server for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 4509, "body_sha256": "sha256:0dc949b536b5fccdee7dbca3517abf3391729ef0b35a7baed528f8a7e2cb37ed", "canonical_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "path": "docs/guides/resources--nfv_service--properties--palo_alto_fw_service--panorama_server.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["palo_alto_fw_service", "panorama_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.panorama_server for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.panorama_server

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md)
- palo_alto_fw_service.panorama_server

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for panorama server.

Upstream description:

Panorama Server Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server")}
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
panorama_server {
  # Configure direct properties listed below.
}
```

## Direct properties

- [authorization_key](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key.md): complete subsection reference.

<a id="schema-palo_alto_fw_service--panorama_server--device_group_name"></a>

### device_group_name property

Type: `"string"`. Optional.

Device Group Name. Device Group Name.

Upstream description:

Device Group Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-palo_alto_fw_service--panorama_server--server"></a>

### server property

Type: `"string"`. Optional.

Panorama Server Address to which the firewall should connect to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-palo_alto_fw_service--panorama_server--template_stack_name"></a>

### template_stack_name property

Type: `"string"`. Optional.

Template stack name. Template Stack Name.

Upstream description:

Template Stack Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

## Next pages

- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--properties--palo_alto_fw_service--panorama_server--authorization_key.md)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
