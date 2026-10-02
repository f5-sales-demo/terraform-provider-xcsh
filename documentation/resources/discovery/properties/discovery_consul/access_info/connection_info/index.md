---
page_title: "discovery_consul.access_info.connection_info"
subcategory: ""
description: "Configuration details to access discovery service REST API."
xcsh_docs: {"aliases": ["discovery consul access info connection info"], "body_bytes": 3123, "body_sha256": "sha256:4cfb7f82f815e9bf1f6feab66199536b6694a15fce73cc14da8a4d7280700523", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info", "path": "documentation/resources/discovery/properties/discovery_consul/access_info/connection_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1021002230130101-2331301020313031-0203232222310221-2120111231021320-3112330001001112-0221300302102100-3333022202330012-2022101023021301", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "schema-discovery_consul--access_info--connection_info--api_server", "enforcement": "provider-schema", "group": "discovery_consul.access_info.connection_info:RequiredObjectAttributes:api_server", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "access_info", "connection_info"], "schema_version": 1, "sections": [{"aliases": ["api server"], "anchor": "schema-discovery_consul--access_info--connection_info--api_server", "description": "API server must be a fully qualified domain string and port specified as host:port pair.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "access_info", "connection_info", "api_server"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls info"], "anchor": "section", "description": "TLS config for client of discovery service.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:connection_info:tls_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_consul", "access_info", "connection_info", "tls_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/connection_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration details to access discovery service REST API.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.connection_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/)
- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/)
- discovery_consul.access_info.connection_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration details to access discovery service REST API.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server")}
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
connection_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-discovery_consul--access_info--connection_info--api_server"></a>

### api_server property

Type: `"string"`. Optional.

API server must be a fully qualified domain string and port specified as host:port pair.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

- [tls_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/): complete subsection reference.

## Next pages

- [discovery_consul.access_info.connection_info.tls_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/)
- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
