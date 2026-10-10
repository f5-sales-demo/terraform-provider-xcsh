---
page_title: "discovery_k8s.access_info.connection_info.tls_info"
subcategory: ""
description: "TLS config for client of discovery service."
xcsh_docs: {"aliases": ["discovery k8s access info connection info tls info"], "body_bytes": 4783, "body_sha256": "sha256:7748329f49fd70f57412f67ac92479d2408400e1ec7bce50be7585131fbfa8f8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info:key_url"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info", "path": "documentation/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3213112033331231-0002002323300313-3100230200032300-3302120331312303-3131200202121132-3321213301222231-3121330213302123-1131012112231132", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "access_info", "connection_info", "tls_info"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s access info connection info tls info certificate"], "anchor": "schema-discovery_k8s--access_info--connection_info--tls_info--certificate", "description": "Client certificate is PEM-encoded certificate or certificate-chain.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "connection_info", "tls_info", "certificate"], "syntax": "attribute", "type": "string"}, {"aliases": ["discovery k8s access info connection info tls info key url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info:key_url", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "access_info", "connection_info", "tls_info", "key_url"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery k8s access info connection info tls info server name"], "anchor": "schema-discovery_k8s--access_info--connection_info--tls_info--server_name", "description": "ServerName is passed to the server for SNI and is used in the client to check server certificates against. If ServerName is empty, the hostname used to contact the server is used.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "connection_info", "tls_info", "server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["discovery k8s access info connection info tls info trusted ca url"], "anchor": "schema-discovery_k8s--access_info--connection_info--tls_info--trusted_ca_url", "description": "The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "connection_info", "tls_info", "trusted_ca_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "TLS config for client of discovery service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["discoveryCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.access_info.connection_info.tls_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
- [discovery_k8s.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/)
- [discovery_k8s.access_info.connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/)
- discovery_k8s.access_info.connection_info.tls_info

<a id="section"></a>

Type: `"single"`. Computed.

TLS config for client of discovery service.

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

<a id="schema-discovery_k8s--access_info--connection_info--tls_info--certificate"></a>

### certificate property

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/): complete subsection reference.

<a id="schema-discovery_k8s--access_info--connection_info--tls_info--server_name"></a>

### server_name property

Type: `"string"`. Computed.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-discovery_k8s--access_info--connection_info--tls_info--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```
