---
page_title: "discovery_k8s.access_info.connection_info"
subcategory: ""
description: "discovery_k8s.access_info.connection_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2586, "body_sha256": "sha256:b08b3e3e563d78ef9590181f6d1339e983ddc53a5288ddf957d0f74601649645", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info", "path": "docs/guides/resources--discovery--properties--discovery_k8s--access_info--connection_info.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s", "access_info", "connection_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/access_info/connection_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.access_info.connection_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# discovery_k8s.access_info.connection_info

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- [discovery_k8s](resources--discovery--properties--discovery_k8s.md)
- [discovery_k8s.access_info](resources--discovery--properties--discovery_k8s--access_info.md)
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

- [tls_info](resources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info.md): complete subsection reference.

## Next pages

- [discovery_k8s.access_info.connection_info.tls_info](resources--discovery--properties--discovery_k8s--access_info--connection_info--tls_info.md)
- [discovery_k8s.access_info](resources--discovery--properties--discovery_k8s--access_info.md)
- [xcsh_discovery](../resources/discovery.md)
