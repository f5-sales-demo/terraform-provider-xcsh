---
page_title: "palo_alto_fw_service.panorama_server"
subcategory: ""
description: "Panorama Server Type."
xcsh_docs: {"aliases": ["palo alto fw service panorama server"], "body_bytes": 4864, "body_sha256": "sha256:66f5b8de1a6f5e80f5fe67c39cf1f600aa0b7916effb30fb6401017802ac4f7a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "path": "documentation/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1011323001103321-2030100221203210-2122232000230032-2223123210312101-0202020002101023-0101121233000113-3313223221110312-1303032221032111", "registry_path": "docs/guides/resources--nfv_service--reference--group-004.md", "relationships": [{"anchor": "schema-palo_alto_fw_service--panorama_server--server", "enforcement": "provider-schema", "group": "palo_alto_fw_service.panorama_server:RequiredObjectAttributes:server", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "panorama_server"], "schema_version": 1, "sections": [{"aliases": ["authorization key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.panorama_server.authorization_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.panorama_server.authorization_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["palo_alto_fw_service", "panorama_server", "authorization_key"], "syntax": "block", "type": "object"}, {"aliases": ["device group name"], "anchor": "schema-palo_alto_fw_service--panorama_server--device_group_name", "description": "Device Group Name.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "panorama_server", "device_group_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["server"], "anchor": "schema-palo_alto_fw_service--panorama_server--server", "description": "Panorama Server Address to which the firewall should connect to.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "panorama_server", "server"], "syntax": "attribute", "type": "string"}, {"aliases": ["template stack name"], "anchor": "schema-palo_alto_fw_service--panorama_server--template_stack_name", "description": "Template Stack Name.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "panorama_server", "template_stack_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Panorama Server Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.panorama_server

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/)
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

- [authorization_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/): complete subsection reference.

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

- [palo_alto_fw_service.panorama_server.authorization_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
