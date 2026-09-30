---
page_title: "azure_event_hubs_receiver"
subcategory: ""
description: "azure_event_hubs_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4481, "body_sha256": "sha256:8cd0a02d16fada84cd5b9efb441ded63cb453e82d148e14a69c3da839761de3b", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_event_hubs_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "docs/guides/resources--global_log_receiver--properties--azure_event_hubs_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_event_hubs_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_event_hubs_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_event_hubs_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# azure_event_hubs_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- azure_event_hubs_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Azure Event Hubs Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("instance",
    "namespace")}
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
azure_event_hubs_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [connection_string](resources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string.md): complete subsection reference.

<a id="schema-azure_event_hubs_receiver--instance"></a>

### instance property

Type: `"string"`. Optional.

Event Hubs Instance name into which logs should be stored.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

<a id="schema-azure_event_hubs_receiver--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Event Hubs Namespace is namespace with instance into which logs should be stored.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

## Next pages

- [azure_event_hubs_receiver.connection_string](resources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
