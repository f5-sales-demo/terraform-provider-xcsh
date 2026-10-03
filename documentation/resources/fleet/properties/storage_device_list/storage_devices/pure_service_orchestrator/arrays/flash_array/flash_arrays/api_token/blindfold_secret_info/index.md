---
page_title: "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info"
subcategory: ""
description: "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management."
xcsh_docs: {"aliases": ["storage device list storage devices pure service orchestrator arrays flash array flash arrays api token blindfold secret info"], "body_bytes": 6660, "body_sha256": "sha256:d61dc887a55628d0b83094561dcb4cdb1c3d028d01da71decdc0f91f86e2a312", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays:api_token:blindfold_secret_info", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays:api_token", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/flash_arrays/api_token/blindfold_secret_info/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3322310000203022-3222000303221130-3312130212322301-2010232221331200-0123032312123002-1201201201101120-1302101123332032-0222332301231223", "registry_path": "docs/guides/resources--fleet--reference--group-004.md", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays:api_token:blindfold_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "flash_arrays", "api_token", "blindfold_secret_info"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices pure service orchestrator arrays flash array flash arrays api token blindfold secret info decryption provider"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--decryption_provider", "description": "Name of the Secret Management Access object that contains information about the backend Secret Management service.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays:api_token:blindfold_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "flash_arrays", "api_token", "blindfold_secret_info", "decryption_provider"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices pure service orchestrator arrays flash array flash arrays api token blindfold secret info location"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--location", "description": "Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays:api_token:blindfold_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "flash_arrays", "api_token", "blindfold_secret_info", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices pure service orchestrator arrays flash array flash arrays api token blindfold secret info store provider"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--store_provider", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays:api_token:blindfold_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "flash_arrays", "api_token", "blindfold_secret_info", "store_provider"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/flash_arrays/api_token/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/flash_arrays/)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/flash_arrays/api_token/)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--decryption_provider"></a>

### decryption_provider property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
  }
}
```

<a id="schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--location"></a>

### location property

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--api_token--blindfold_secret_info--store_provider"></a>

### store_provider property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
  }
}
```

## Next pages

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/flash_arrays/api_token/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
