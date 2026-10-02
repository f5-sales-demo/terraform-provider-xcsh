---
page_title: "qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["qradar receiver use tls mtls enable key url clear secret info"], "body_bytes": 4277, "body_sha256": "sha256:ddbb2977c97431a35a37d6756b5e1a77356bba6c613f926b056c658bc8923cb1", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable:key_url", "path": "documentation/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2022312021032211-3333213111301231-0130111113321001-2222032122300130-0003121003223301-2000122202202202-3222023132210333-0110023213202232", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-004.md", "relationships": [{"anchor": "schema-qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url", "enforcement": "provider-schema", "group": "qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["qradar_receiver", "use_tls", "mtls_enable", "key_url", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "mtls_enable", "key_url", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "mtls_enable", "key_url", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/clear_secret_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/)
- [qradar_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/)
- [qradar_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/)
- [qradar_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/)
- qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info

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

<a id="schema-qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-qradar_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url"></a>

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

- [qradar_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
