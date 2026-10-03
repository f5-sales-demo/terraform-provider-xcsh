---
page_title: "https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["https management advertise on slo sli tls certificates private key clear secret info"], "body_bytes": 4060, "body_sha256": "sha256:0a6c4e57d7b5a7d128e9321d14580a000880fc460e628bc24acfbb24e19a2877", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_certificates:private_key:clear_secret_info", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_certificates:private_key", "path": "documentation/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_certificates/private_key/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1231030232233310-1212233131111122-2222133203002230-3300200312020232-3023200320003111-3021232002312233-3210333323101223-0210313223112022", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_slo_sli", "tls_certificates", "private_key", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["https management advertise on slo sli tls certificates private key clear secret info provider ref"], "anchor": "schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_certificates:private_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_sli", "tls_certificates", "private_key", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["https management advertise on slo sli tls certificates private key clear secret info url"], "anchor": "schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_certificates:private_key:clear_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_sli", "tls_certificates", "private_key", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_certificates/private_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/)
- [https_management.advertise_on_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/)
- [https_management.advertise_on_slo_sli.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_certificates/)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_certificates/private_key/)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info

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

<a id="schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-https_management--advertise_on_slo_sli--tls_certificates--private_key--clear_secret_info--url"></a>

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

- [https_management.advertise_on_slo_sli.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_certificates/private_key/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
