---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info"
subcategory: ""
description: "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management."
xcsh_docs: {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap target initiator secret blindfold secret info"], "body_bytes": 5281, "body_sha256": "sha256:d5717982bf6a8a9225d706b20169dfd26885c045afb1d8e681813de74cfbb1e8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret:blindfold_secret_info", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret", "path": "documentation/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_target_initiator_secret/blindfold_secret_info/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3322321230331020-0100100110023131-0111023101130232-3213312223203302-2201122113320010-2230002030222100-2312221020031100-1011330023231132", "registry_path": "docs/guides/data-sources--fleet--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_target_initiator_secret", "blindfold_secret_info"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap target initiator secret blindfold secret info decryption provider"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info--decryption_provider", "description": "Name of the Secret Management Access object that contains information about the backend Secret Management service.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_target_initiator_secret", "blindfold_secret_info", "decryption_provider"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap target initiator secret blindfold secret info location"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info--location", "description": "Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_target_initiator_secret", "blindfold_secret_info", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap target initiator secret blindfold secret info store provider"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info--store_provider", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_target_initiator_secret", "blindfold_secret_info", "store_provider"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_target_initiator_secret/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["fleetCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_target_initiator_secret/)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

## Direct properties

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info--decryption_provider"></a>

### decryption_provider property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info--location"></a>

### location property

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info--store_provider"></a>

### store_provider property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
