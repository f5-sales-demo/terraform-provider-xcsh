---
page_title: "https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info"
subcategory: ""
description: "https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 4709, "body_sha256": "sha256:1c4bf6cb5219217d30a42a4976882ef4de9219dbf42cebcc8236bf2fbb59d08b", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_certificates:private_key:blindfold_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_certificates:private_key:blindfold_secret_info", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_certificates:private_key", "path": "docs/guides/data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_slo_sli", "tls_certificates", "private_key", "blindfold_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_certificates/private_key/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [https_management](data-sources--nfv_service--properties--https_management.md)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli.md)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates.md)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key.md)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info

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

<a id="schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info--decryption_provider"></a>

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

<a id="schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info--location"></a>

### location property

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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

<a id="schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--blindfold_secret_info--store_provider"></a>

### store_provider property

Type: `"string"`. Computed.

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

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_certificates--private_key.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
