---
page_title: "http_receiver.use_tls.mtls_enable.key_url.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["http receiver use tls mtls enable key url clear secret info"], "body_bytes": 3860, "body_sha256": "sha256:a718f49711b219b0d2f6568442545df0564d6da51a4ed3bffb2281c65ba24382", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url", "path": "documentation/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0102312301010323-0101310233032311-3203011200311011-1323233023200231-2201020323132312-3220133231222220-2123300110200312-1010333011223212", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_receiver", "use_tls", "mtls_enable", "key_url", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "use_tls", "mtls_enable", "key_url", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "use_tls", "mtls_enable", "key_url", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.use_tls.mtls_enable.key_url.clear_secret_info

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/)
- [http_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/)
- [http_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/)
- [http_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/)
- http_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url"></a>

### url property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

- [http_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
