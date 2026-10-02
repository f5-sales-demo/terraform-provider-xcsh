---
page_title: "more_option.response_headers_to_add.secret_value.blindfold_secret_info"
subcategory: "Load Balancing"
description: "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management."
xcsh_docs: {"aliases": ["more option response headers to add secret value blindfold secret info"], "body_bytes": 4876, "body_sha256": "sha256:36257a231a81c214a9c949db6ff0e7b647146c36b04c0dc78fe7bb0b2ecddc59", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_headers_to_add:secret_value:blindfold_secret_info", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_headers_to_add:secret_value", "path": "documentation/data-sources/http_loadbalancer/properties/more_option/response_headers_to_add/secret_value/blindfold_secret_info/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0210300123133110-2331100013321200-1031202323100310-1123013031321012-1021100031121211-2313132101122002-3110101102302301-2323102022030001", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-021.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["more_option", "response_headers_to_add", "secret_value", "blindfold_secret_info"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "decryption provider", "origin servers", "upstream servers"], "anchor": "schema-more_option--response_headers_to_add--secret_value--blindfold_secret_info--decryption_provider", "description": "Name of the Secret Management Access object that contains information about the backend Secret Management service.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_headers_to_add:secret_value:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_headers_to_add", "secret_value", "blindfold_secret_info", "decryption_provider"], "syntax": "attribute", "type": "string"}, {"aliases": ["location"], "anchor": "schema-more_option--response_headers_to_add--secret_value--blindfold_secret_info--location", "description": "Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_headers_to_add:secret_value:blindfold_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_headers_to_add", "secret_value", "blindfold_secret_info", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["store provider"], "anchor": "schema-more_option--response_headers_to_add--secret_value--blindfold_secret_info--store_provider", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_headers_to_add:secret_value:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_headers_to_add", "secret_value", "blindfold_secret_info", "store_provider"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/more_option/response_headers_to_add/secret_value/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option.response_headers_to_add.secret_value.blindfold_secret_info

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/)
- [more_option.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_headers_to_add/)
- [more_option.response_headers_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_headers_to_add/secret_value/)
- more_option.response_headers_to_add.secret_value.blindfold_secret_info

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

<a id="schema-more_option--response_headers_to_add--secret_value--blindfold_secret_info--decryption_provider"></a>

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

<a id="schema-more_option--response_headers_to_add--secret_value--blindfold_secret_info--location"></a>

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

<a id="schema-more_option--response_headers_to_add--secret_value--blindfold_secret_info--store_provider"></a>

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

- [more_option.response_headers_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_headers_to_add/secret_value/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
