---
page_title: "custom_proxy"
subcategory: ""
description: "custom_proxy for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 5325, "body_sha256": "sha256:0cc9a5072f0da910d41ce19ecd7b4b3d32f2b6424b2e898b7f01b7722dd9c2a2", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:disable_re_tunnel", "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:enable_re_tunnel", "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:password"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--custom_proxy.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/custom_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_proxy for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_proxy

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- custom_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_proxy, f5\_proxy, private\_adn\] Configuration parameter for custom proxy.

Upstream description:

Custom Enterprise Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("proxy_ip_address",
    "proxy_port"),
  validators.ConflictingObjectAttributes("disable_re_tunnel",
    "enable_re_tunnel")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-use_for_re_tunnel_choice": "[\"disable_re_tunnel\",\"enable_re_tunnel\"]"
}
```

OneOf alternatives in this subsection:

- [custom_proxy](resources--securemesh_site_v2--properties--custom_proxy.md#section)
- [f5_proxy](resources--securemesh_site_v2--properties--f5_proxy.md#section)
- [private_adn](resources--securemesh_site_v2--properties--private_adn.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_re_tunnel](resources--securemesh_site_v2--properties--custom_proxy--disable_re_tunnel.md): complete subsection reference.

- [enable_re_tunnel](resources--securemesh_site_v2--properties--custom_proxy--enable_re_tunnel.md): complete subsection reference.

- [password](resources--securemesh_site_v2--properties--custom_proxy--password.md): complete subsection reference.

<a id="schema-custom_proxy--proxy_ip_address"></a>

### proxy_ip_address property

Type: `"string"`. Optional.

Specify the IPv4 Address of the internal Enterprise Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="schema-custom_proxy--proxy_port"></a>

### proxy_port property

Type: `"number"`. Optional.

Specify the Port of the internal Enterprise Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-custom_proxy--username"></a>

### username property

Type: `"string"`. Optional.

If the internal Enterprise Proxy is using basic authentication, specify the username. This is an
optional field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
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

- [custom_proxy.disable_re_tunnel](resources--securemesh_site_v2--properties--custom_proxy--disable_re_tunnel.md)
- [custom_proxy.enable_re_tunnel](resources--securemesh_site_v2--properties--custom_proxy--enable_re_tunnel.md)
- [custom_proxy.password](resources--securemesh_site_v2--properties--custom_proxy--password.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
