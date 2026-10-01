---
page_title: "discovery_consul.access_info.http_basic_auth_info"
subcategory: ""
description: "discovery_consul.access_info.http_basic_auth_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2384, "body_sha256": "sha256:f8babc49844f028593a737837de70ed2747c83da30d28d8b5a05c2abacae2278", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info", "path": "docs/guides/resources--discovery--properties--discovery_consul--access_info--http_basic_auth_info.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.access_info.http_basic_auth_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.http_basic_auth_info

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- [discovery_consul](resources--discovery--properties--discovery_consul.md)
- [discovery_consul.access_info](resources--discovery--properties--discovery_consul--access_info.md)
- discovery_consul.access_info.http_basic_auth_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters to access Hashicorp Consul.

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
http_basic_auth_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [passwd_url](resources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url.md): complete subsection reference.

<a id="schema-discovery_consul--access_info--http_basic_auth_info--user_name"></a>

### user_name property

Type: `"string"`. Optional.

User Name. Username in consul.

Upstream description:

Username in consul.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

## Next pages

- [discovery_consul.access_info.http_basic_auth_info.passwd_url](resources--discovery--properties--discovery_consul--access_info--http_basic_auth_info--passwd_url.md)
- [discovery_consul.access_info](resources--discovery--properties--discovery_consul--access_info.md)
- [xcsh_discovery](../resources/discovery.md)
