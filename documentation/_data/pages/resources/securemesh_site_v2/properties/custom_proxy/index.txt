---
page_title: "custom_proxy"
subcategory: ""
description: "Custom Enterprise Proxy."
xcsh_docs: {"aliases": ["custom proxy"], "body_bytes": 5986, "body_sha256": "sha256:c2b86371ebf56a753300a304b531177a10260ce4b4510fe25585436ecf750ffb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:disable_re_tunnel", "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:enable_re_tunnel", "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:password"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/custom_proxy/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_proxy:ConflictingObjectAttributes:disable_re_tunnel,enable_re_tunnel", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:disable_re_tunnel", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_proxy:ConflictingObjectAttributes:disable_re_tunnel,enable_re_tunnel", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:enable_re_tunnel", "type": "conflicts"}, {"anchor": "schema-custom_proxy--proxy_ip_address", "enforcement": "provider-schema", "group": "custom_proxy:RequiredObjectAttributes:proxy_ip_address,proxy_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "type": "requires"}, {"anchor": "schema-custom_proxy--proxy_port", "enforcement": "provider-schema", "group": "custom_proxy:RequiredObjectAttributes:proxy_ip_address,proxy_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_proxy"], "schema_version": 1, "sections": [{"aliases": ["disable re tunnel"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:disable_re_tunnel", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "disable_re_tunnel"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable re tunnel"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:enable_re_tunnel", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "enable_re_tunnel"], "syntax": "attribute", "type": "object"}, {"aliases": ["password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:password", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_proxy.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_proxy.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["custom_proxy", "password"], "syntax": "block", "type": "object"}, {"aliases": ["proxy ip address"], "anchor": "schema-custom_proxy--proxy_ip_address", "description": "Specify the IPv4 Address of the internal Enterprise Proxy.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "proxy_ip_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["proxy port"], "anchor": "schema-custom_proxy--proxy_port", "description": "Specify the Port of the internal Enterprise Proxy.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "proxy_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["authentication", "credential setup", "credentials", "username"], "anchor": "schema-custom_proxy--username", "description": "If the internal Enterprise Proxy is using basic authentication, specify the username. This is an optional field.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "username"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/custom_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Custom Enterprise Proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_proxy

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
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

- [custom_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/custom_proxy/#section)
- [f5_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/f5_proxy/#section)
- [private_adn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/private_adn/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_re_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/custom_proxy/disable_re_tunnel/): complete subsection reference.

- [enable_re_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/custom_proxy/enable_re_tunnel/): complete subsection reference.

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/custom_proxy/password/): complete subsection reference.

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

- [custom_proxy.disable_re_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/custom_proxy/disable_re_tunnel/)
- [custom_proxy.enable_re_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/custom_proxy/enable_re_tunnel/)
- [custom_proxy.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/custom_proxy/password/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
