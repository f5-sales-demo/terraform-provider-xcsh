---
page_title: "default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info"
subcategory: "Load Balancing"
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["cert", "certificate", "default pool use tls use mtls tls certificates private key clear secret info", "existing certificates", "tls certificates"], "body_bytes": 4164, "body_sha256": "sha256:6346d6d473272d9ea4b37ab91317e14746237f30899bff84c4030c533167658f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:private_key", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3030023122112131-0011033301011123-0221123320010321-2302303333231133-2123332312320212-3220112303022313-3233103010030233-3330310123303003", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "use_tls", "use_mtls", "tls_certificates", "private_key", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-default_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "use_mtls", "tls_certificates", "private_key", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-default_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "use_mtls", "tls_certificates", "private_key", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/)
- [default_pool.use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/)
- [default_pool.use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/tls_certificates/)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/tls_certificates/private_key/)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

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

<a id="schema-default_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-default_pool--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--url"></a>

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

- [default_pool.use_tls.use_mtls.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/tls_certificates/private_key/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
