---
page_title: "origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info"
subcategory: "Load Balancing"
description: "origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3977, "body_sha256": "sha256:3485d58094d3682934610836c86dc707a147a01dda279f7efb0f1ed7bf3caf30", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key", "path": "docs/guides/resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "private_key", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [origin_pool](resources--cdn_loadbalancer--properties--origin_pool.md)
- [origin_pool.use_tls](resources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls.md)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates.md)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key.md)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--url"></a>

### url property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--properties--origin_pool--use_tls--use_mtls--tls_certificates--private_key.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
