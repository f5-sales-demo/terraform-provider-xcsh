---
page_title: "discovery_k8s.access_info.connection_info"
subcategory: ""
description: "Configuration details to access discovery service REST API."
xcsh_docs: {"aliases": ["discovery k8s access info connection info"], "body_bytes": 3087, "body_sha256": "sha256:7bbb808f097ec5ddeff865622240ec86641d9abde3caec04fb1a96ed4e923ece", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info", "path": "documentation/resources/discovery/properties/discovery_k8s/access_info/connection_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3102311222101313-3002123010110023-0001213313233230-1111231130222303-3323223313013213-1003311301200300-2031022333113221-0021113332233333", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "schema-discovery_k8s--access_info--connection_info--api_server", "enforcement": "provider-schema", "group": "discovery_k8s.access_info.connection_info:RequiredObjectAttributes:api_server", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "access_info", "connection_info"], "schema_version": 1, "sections": [{"aliases": ["api server"], "anchor": "schema-discovery_k8s--access_info--connection_info--api_server", "description": "API server must be a fully qualified domain string and port specified as host:port pair.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "connection_info", "api_server"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls info"], "anchor": "section", "description": "TLS config for client of discovery service.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "access_info", "connection_info", "tls_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/access_info/connection_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration details to access discovery service REST API.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.access_info.connection_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- [discovery_k8s.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/)
- discovery_k8s.access_info.connection_info

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

<a id="schema-discovery_k8s--access_info--connection_info--api_server"></a>

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

- [tls_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/): complete subsection reference.

## Next pages

- [discovery_k8s.access_info.connection_info.tls_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/)
- [discovery_k8s.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
