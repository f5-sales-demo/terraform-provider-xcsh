---
page_title: "discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["discovery consul access info connection info tls info key url clear secret info"], "body_bytes": 4554, "body_sha256": "sha256:7307e2771a93b14261b61b6c0c8be12e7ee039733bdfb6fb5cd49f9ce47f9af5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:clear_secret_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url", "path": "documentation/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0012201201132201-2113310323100122-1032223311130201-1111200003201301-1212221232011100-0213123113301323-3302213002211023-0012003322201122", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--url", "enforcement": "provider-schema", "group": "discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:clear_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "access_info", "connection_info", "tls_info", "key_url", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:clear_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "access_info", "connection_info", "tls_info", "key_url", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info:key_url:clear_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "access_info", "connection_info", "tls_info", "key_url", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/)
- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/)
- [discovery_consul.access_info.connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/)
- [discovery_consul.access_info.connection_info.tls_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/)
- [discovery_consul.access_info.connection_info.tls_info.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/)
- discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info

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

<a id="schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--url"></a>

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

- [discovery_consul.access_info.connection_info.tls_info.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
