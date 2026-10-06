---
page_title: "custom_proxy"
subcategory: ""
description: "Custom Enterprise Proxy."
xcsh_docs: {"aliases": ["custom proxy"], "body_bytes": 5359, "body_sha256": "sha256:c31e332a5421b8c4a20d104cef27fca9acfe8d78b339a43218fa01fd3ca3cf1f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:disable_re_tunnel", "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:enable_re_tunnel", "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:password"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/custom_proxy/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_proxy:ConflictingObjectAttributes:disable_re_tunnel,enable_re_tunnel", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:disable_re_tunnel", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_proxy:ConflictingObjectAttributes:disable_re_tunnel,enable_re_tunnel", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:enable_re_tunnel", "type": "conflicts"}, {"anchor": "schema-custom_proxy--proxy_ip_address", "enforcement": "provider-schema", "group": "custom_proxy:RequiredObjectAttributes:proxy_ip_address,proxy_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "type": "requires"}, {"anchor": "schema-custom_proxy--proxy_port", "enforcement": "provider-schema", "group": "custom_proxy:RequiredObjectAttributes:proxy_ip_address,proxy_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_proxy"], "schema_version": 1, "sections": [{"aliases": ["custom proxy disable re tunnel"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:disable_re_tunnel", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "disable_re_tunnel"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom proxy enable re tunnel"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:enable_re_tunnel", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "enable_re_tunnel"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom proxy password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:password", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_proxy.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_proxy.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["custom_proxy", "password"], "syntax": "block", "type": "object"}, {"aliases": ["custom proxy proxy ip address"], "anchor": "schema-custom_proxy--proxy_ip_address", "description": "Specify the IPv4 Address of the internal Enterprise Proxy.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "proxy_ip_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom proxy proxy port"], "anchor": "schema-custom_proxy--proxy_port", "description": "Specify the Port of the internal Enterprise Proxy.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "proxy_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["custom proxy username"], "anchor": "schema-custom_proxy--username", "description": "If the internal Enterprise Proxy is using basic authentication, specify the username. This is an optional field.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:custom_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_proxy", "username"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/custom_proxy/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Custom Enterprise Proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Additional upstream details:

Custom Enterprise Proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
