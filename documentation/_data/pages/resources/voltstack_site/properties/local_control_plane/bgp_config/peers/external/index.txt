---
page_title: "local_control_plane.bgp_config.peers.external"
subcategory: ""
description: "local_control_plane.bgp_config.peers.external for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 20810, "body_sha256": "sha256:ed1d624cd2a0a498f9374e59d2205fe03738b6093892d03432a453ac2e8c54c8", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:default_gateway", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:default_gateway_v6", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:disable_spec", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:disable_v6", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:external_connector", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:family_inet", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:from_site", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:from_site_v6", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface_list", "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:no_authentication"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers", "path": "documentation/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_control_plane.bgp_config.peers.external for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_control_plane.bgp_config.peers.external

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- local_control_plane.bgp_config.peers.external

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

<a id="schema-local_control_plane--bgp_config--peers--external--address"></a>

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

<a id="schema-local_control_plane--bgp_config--peers--external--address_ipv6"></a>

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

<a id="schema-local_control_plane--bgp_config--peers--external--asn"></a>

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

- [default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/default_gateway/): complete subsection reference.

- [default_gateway_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/default_gateway_v6/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/disable_spec/): complete subsection reference.

- [disable_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/disable_v6/): complete subsection reference.

- [external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/external_connector/): complete subsection reference.

- [family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/): complete subsection reference.

- [from_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/from_site/): complete subsection reference.

- [from_site_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/from_site_v6/): complete subsection reference.

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/interface/): complete subsection reference.

- [interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/interface_list/): complete subsection reference.

<a id="schema-local_control_plane--bgp_config--peers--external--md5_auth_key"></a>

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

- [no_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/no_authentication/): complete subsection reference.

<a id="schema-local_control_plane--bgp_config--peers--external--port"></a>

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

<a id="schema-local_control_plane--bgp_config--peers--external--subnet_begin_offset"></a>

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

<a id="schema-local_control_plane--bgp_config--peers--external--subnet_begin_offset_v6"></a>

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

<a id="schema-local_control_plane--bgp_config--peers--external--subnet_end_offset"></a>

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

<a id="schema-local_control_plane--bgp_config--peers--external--subnet_end_offset_v6"></a>

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

- [local_control_plane.bgp_config.peers.external.default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/default_gateway/)
- [local_control_plane.bgp_config.peers.external.default_gateway_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/default_gateway_v6/)
- [local_control_plane.bgp_config.peers.external.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/disable_spec/)
- [local_control_plane.bgp_config.peers.external.disable_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/disable_v6/)
- [local_control_plane.bgp_config.peers.external.external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/external_connector/)
- [local_control_plane.bgp_config.peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/family_inet/)
- [local_control_plane.bgp_config.peers.external.from_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/from_site/)
- [local_control_plane.bgp_config.peers.external.from_site_v6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/from_site_v6/)
- [local_control_plane.bgp_config.peers.external.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/interface/)
- [local_control_plane.bgp_config.peers.external.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/interface_list/)
- [local_control_plane.bgp_config.peers.external.no_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/no_authentication/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
