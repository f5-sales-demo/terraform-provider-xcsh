---
page_title: "peers.external"
subcategory: ""
description: "peers.external for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 17451, "body_sha256": "sha256:1b47ddee5d720d3fafd2faf760af1ba77db8a3ecbbc5dfdff409dc24192d4032", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:external", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:external:default_gateway", "xcsh-docs:resources:bgp:properties:peers:external:default_gateway_v6", "xcsh-docs:resources:bgp:properties:peers:external:disable_spec", "xcsh-docs:resources:bgp:properties:peers:external:disable_v6", "xcsh-docs:resources:bgp:properties:peers:external:external_connector", "xcsh-docs:resources:bgp:properties:peers:external:family_inet", "xcsh-docs:resources:bgp:properties:peers:external:from_site", "xcsh-docs:resources:bgp:properties:peers:external:from_site_v6", "xcsh-docs:resources:bgp:properties:peers:external:interface", "xcsh-docs:resources:bgp:properties:peers:external:interface_list", "xcsh-docs:resources:bgp:properties:peers:external:no_authentication"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external", "parent_id": "xcsh-docs:resources:bgp:properties:peers", "path": "docs/guides/resources--bgp--properties--peers--external.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "external"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.external for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- peers.external

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

External BGP Peer. External BGP Peer parameters.

Upstream description:

External BGP Peer parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn",
    "port"),
  validators.ConflictingObjectAttributes("address",
    "default_gateway"),
  validators.ConflictingObjectAttributes("address",
    "disable_spec"),
  validators.ConflictingObjectAttributes("address",
    "external_connector"),
  validators.ConflictingObjectAttributes("address",
    "from_site"),
  validators.ConflictingObjectAttributes("address",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("address",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "default_gateway_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway",
    "disable_spec"),
  validators.ConflictingObjectAttributes("default_gateway",
    "external_connector"),
  validators.ConflictingObjectAttributes("default_gateway",
    "from_site"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("disable_spec",
    "external_connector"),
  validators.ConflictingObjectAttributes("disable_spec",
    "from_site"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("disable_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("external_connector",
    "from_site"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("interface",
    "interface_list"),
  validators.ConflictingObjectAttributes("md5_auth_key",
    "no_authentication"),
  validators.ConflictingObjectAttributes("subnet_begin_offset",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("subnet_begin_offset_v6",
    "subnet_end_offset_v6")}
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
  "x-ves-oneof-field-address_choice": "[\"address\",\"default_gateway\",\"disable\",\"external_connector\",\"from_site\",\"subnet_begin_offset\",\"subnet_end_offset\"]",
  "x-ves-oneof-field-address_choice_v6": "[\"address_ipv6\",\"default_gateway_v6\",\"disable_v6\",\"from_site_v6\",\"subnet_begin_offset_v6\",\"subnet_end_offset_v6\"]",
  "x-ves-oneof-field-auth_choice": "[\"md5_auth_key\",\"no_authentication\"]",
  "x-ves-oneof-field-interface_choice": "[\"interface\",\"interface_list\"]"
}
```

Terraform syntax:

```terraform
external {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-peers--external--address"></a>

### address property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Upstream description:

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="schema-peers--external--address_ipv6"></a>

### address_ipv6 property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Upstream description:

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="schema-peers--external--asn"></a>

### asn property

Type: `"number"`. Optional.

ASN. Autonomous System Number for BGP peer.

Upstream description:

Autonomous System Number for BGP peer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [default_gateway](resources--bgp--properties--peers--external--default_gateway.md): complete subsection reference.

- [default_gateway_v6](resources--bgp--properties--peers--external--default_gateway_v6.md): complete subsection reference.

- [disable_spec](resources--bgp--properties--peers--external--disable_spec.md): complete subsection reference.

- [disable_v6](resources--bgp--properties--peers--external--disable_v6.md): complete subsection reference.

- [external_connector](resources--bgp--properties--peers--external--external_connector.md): complete subsection reference.

- [family_inet](resources--bgp--properties--peers--external--family_inet.md): complete subsection reference.

- [from_site](resources--bgp--properties--peers--external--from_site.md): complete subsection reference.

- [from_site_v6](resources--bgp--properties--peers--external--from_site_v6.md): complete subsection reference.

- [interface](resources--bgp--properties--peers--external--interface.md): complete subsection reference.

- [interface_list](resources--bgp--properties--peers--external--interface_list.md): complete subsection reference.

<a id="schema-peers--external--md5_auth_key"></a>

### md5_auth_key property

Type: `"string"`. Optional.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

Upstream description:

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385)

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

- [no_authentication](resources--bgp--properties--peers--external--no_authentication.md): complete subsection reference.

<a id="schema-peers--external--port"></a>

### port property

Type: `"number"`. Optional.

Peer Port. Peer TCP port number.

Upstream description:

Peer TCP port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
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
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-peers--external--subnet_begin_offset"></a>

### subnet_begin_offset property

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-peers--external--subnet_begin_offset_v6"></a>

### subnet_begin_offset_v6 property

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-peers--external--subnet_end_offset"></a>

### subnet_end_offset property

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-peers--external--subnet_end_offset_v6"></a>

### subnet_end_offset_v6 property

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

## Next pages

- [peers.external.default_gateway](resources--bgp--properties--peers--external--default_gateway.md)
- [peers.external.default_gateway_v6](resources--bgp--properties--peers--external--default_gateway_v6.md)
- [peers.external.disable_spec](resources--bgp--properties--peers--external--disable_spec.md)
- [peers.external.disable_v6](resources--bgp--properties--peers--external--disable_v6.md)
- [peers.external.external_connector](resources--bgp--properties--peers--external--external_connector.md)
- [peers.external.family_inet](resources--bgp--properties--peers--external--family_inet.md)
- [peers.external.from_site](resources--bgp--properties--peers--external--from_site.md)
- [peers.external.from_site_v6](resources--bgp--properties--peers--external--from_site_v6.md)
- [peers.external.interface](resources--bgp--properties--peers--external--interface.md)
- [peers.external.interface_list](resources--bgp--properties--peers--external--interface_list.md)
- [peers.external.no_authentication](resources--bgp--properties--peers--external--no_authentication.md)
- [peers](resources--bgp--properties--peers.md)
- [xcsh_bgp](../resources/bgp.md)
