---
page_title: "xcsh_subnet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet reference."
---

# xcsh_subnet reference

<a id="canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85e75e31191ce70b44050e647902a020ef041a027abacc2d73c4704209079b73"></a>

## Property reference — Property reference / 63c2bfef66fb / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- Property reference

<a id="canonical-cad6f4320823f0bb9c6da01d962c881ae46c23a62a30455f535558ada3172574"></a>

## Direct properties — Property reference / 63c2bfef66fb / 3

<a id="canonical-10e5297e5ea9b22a86b2e5c1f048349bb7c07febe4735a90aae8e262298486ac"></a>

<a id="canonical-217cd34eeb44f71cd0aff965173bd262051c2d40daf67d010ab9f55d27ea9c9b"></a>

## annotations property — Property reference / 63c2bfef66fb / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-dc679315875cb7cd67797244b7a0f2b81a17bf99c90ed8332706f02822c90331): complete subsection reference.

- [connect_to_slo](data-sources--subnet--reference--group-001.md#canonical-46c21d95b78d70d1def68e910b0233161ae2b53bfcbab43b6caf63533f279f82): complete subsection reference.

<a id="canonical-18f562e67be759d76122c4167f0495a175b4eb22fff0e8c0a59e738ba3ebd74e"></a>

<a id="canonical-180b2b1c3d3d691206151aee4c6e5e04a01f24eba4616164c667fe08660c8ed8"></a>

## description property — Property reference / 63c2bfef66fb / 5

Type: `"string"`. Computed.

Description of the Subnet.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-104a1b48de3b59a37e9a64c7ce3fc64f4078018528a74c83018311bee699ac69"></a>

<a id="canonical-c5c4f81713b954fd8ea071dcdfc356a55894dbd10d0863202e1fbe68bc2ec3ce"></a>

## id property — Property reference / 63c2bfef66fb / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated_nw](data-sources--subnet--reference--group-001.md#canonical-52a72afef5d80609be631d87cdedad1e613c3b2df7e3691ddc2ba10f609ce52c): complete subsection reference.

<a id="canonical-4be13a1927f61648522956d182b350148e94acb0542216e25d76113011748f6f"></a>

<a id="canonical-fb2a113bca1702e5fc2a06af41919b9514a054dbca0a4b39512a6c6eb984e30c"></a>

## labels property — Property reference / 63c2bfef66fb / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-e8687335f83cefb8b891615b8992ebf4c98d8079e7bf81bb3ace32198d9f7e6d"></a>

<a id="canonical-2d2a3874e844de0833cb20855fe067133f97d453597cc998de9086a3109bd42b"></a>

## name property — Property reference / 63c2bfef66fb / 8

Type: `"string"`. Required.

Name of the Subnet.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-13935a67e57a90df9dcc9331087e3d1e84feb79c4eb8a185b2719ca56c52a517"></a>

<a id="canonical-915428279dc177be54a59f4c81c01f477fd72b2c281e7ccea57ada5b5dde54b9"></a>

## namespace property — Property reference / 63c2bfef66fb / 9

Type: `"string"`. Required.

Namespace where the Subnet exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3): complete subsection reference.

<a id="canonical-2e66ed3e64cb54bd0eb8c567eed97cf328661572d8c131828565175c89276cb2"></a>

## All schema paths — Property reference / 63c2bfef66fb / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--subnet--reference--group-001.md#canonical-10e5297e5ea9b22a86b2e5c1f048349bb7c07febe4735a90aae8e262298486ac) |
| `connect_to_layer2` | [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-26f9f34f10981515519c1105127f25bb8d2b58febe2242df1fc3be2dc2a5ee75) |
| `connect_to_layer2.layer2_intf_ref` | [connect_to_layer2.layer2_intf_ref](data-sources--subnet--reference--group-001.md#canonical-df1a3bf76d7ae90e3e2a540ff625c5818f3b79abe9a04a328263a99be54ca6fb) |
| `connect_to_layer2.layer2_intf_ref.name` | [connect_to_layer2.layer2_intf_ref.name](data-sources--subnet--reference--group-001.md#canonical-90fb47b548d27018ecd835b54ae306f69fd84fee8e76613f1fb707a7c9baf34c) |
| `connect_to_layer2.layer2_intf_ref.namespace` | [connect_to_layer2.layer2_intf_ref.namespace](data-sources--subnet--reference--group-001.md#canonical-c914ba0090edffbab721d17332a1f44fa278f8178e760a1d3916cc712f6b1deb) |
| `connect_to_layer2.layer2_intf_ref.tenant` | [connect_to_layer2.layer2_intf_ref.tenant](data-sources--subnet--reference--group-001.md#canonical-b46605545ec57009a6e66cf40888450794c0784da6a2b0ef1b0837c1172ea4c7) |
| `connect_to_slo` | [connect_to_slo](data-sources--subnet--reference--group-001.md#canonical-f122e0bf3dbaf56391e21f105254761e98454bcb1b5f919f464d9d1a0550d7b7) |
| `description` | [description](data-sources--subnet--reference--group-001.md#canonical-18f562e67be759d76122c4167f0495a175b4eb22fff0e8c0a59e738ba3ebd74e) |
| `id` | [id](data-sources--subnet--reference--group-001.md#canonical-104a1b48de3b59a37e9a64c7ce3fc64f4078018528a74c83018311bee699ac69) |
| `isolated_nw` | [isolated_nw](data-sources--subnet--reference--group-001.md#canonical-9eb6fb15abbb9f68a6a087b6571477c65a76b6732bedd791e8bad4fe410685e8) |
| `labels` | [labels](data-sources--subnet--reference--group-001.md#canonical-4be13a1927f61648522956d182b350148e94acb0542216e25d76113011748f6f) |
| `name` | [name](data-sources--subnet--reference--group-001.md#canonical-e8687335f83cefb8b891615b8992ebf4c98d8079e7bf81bb3ace32198d9f7e6d) |
| `namespace` | [namespace](data-sources--subnet--reference--group-001.md#canonical-13935a67e57a90df9dcc9331087e3d1e84feb79c4eb8a185b2719ca56c52a517) |
| `site_subnet_params` | [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-cc8f12fc682cd50d1191b00ddf09a24fe00d90e0fe1f3173748d00adddbefb32) |
| `site_subnet_params.dhcp` | [site_subnet_params.dhcp](data-sources--subnet--reference--group-001.md#canonical-046d5db191553fe217a2702bb8d4030550b2d4797693e558454b6ff09bffd87a) |
| `site_subnet_params.site` | [site_subnet_params.site](data-sources--subnet--reference--group-001.md#canonical-03d2a7c19c6a54a9d1388cf66928f23e6482e91fa1157f856d244a7e7dc75dd2) |
| `site_subnet_params.site.name` | [site_subnet_params.site.name](data-sources--subnet--reference--group-001.md#canonical-11677e8018a4e019c6f3e5ce7e94053c3b82a9b3578aa8707fa938692fbdb8a0) |
| `site_subnet_params.site.namespace` | [site_subnet_params.site.namespace](data-sources--subnet--reference--group-001.md#canonical-1526f6097e2ced7b7c0979099b0ec61f924cb6c6130743337872303b7d146ca2) |
| `site_subnet_params.site.tenant` | [site_subnet_params.site.tenant](data-sources--subnet--reference--group-001.md#canonical-5987869f16aa5f021666e2ab48b85187420e59f25579cb944d7ead8f8e358cfd) |
| `site_subnet_params.static_ip` | [site_subnet_params.static_ip](data-sources--subnet--reference--group-001.md#canonical-ec3519da365c2d46e2668d0bcdb708254da756377adb88111d08921957f0ecbd) |
| `site_subnet_params.subnet_dhcp_server_params` | [site_subnet_params.subnet_dhcp_server_params](data-sources--subnet--reference--group-001.md#canonical-5d93bf23ff8d4a277f58d2030c4dd33800a6d57ebd20ad232be8c3d1d4b27505) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks](data-sources--subnet--reference--group-001.md#canonical-071c528788baeac407c6e99f46b8d096bca481af30feec38d6c9872710bfbf07) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix](data-sources--subnet--reference--group-001.md#canonical-49dacef18027e3a6cc17295db1c9d2c13b98d42b9ac07845c30e4cb1fae4b523) |

<a id="canonical-25523f18a13b24ef7bc0de4cd3739cdade26156a6d30afae4fc3cbae7d317c9c"></a>

## Next pages — Property reference / 63c2bfef66fb / 11

- [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-dc679315875cb7cd67797244b7a0f2b81a17bf99c90ed8332706f02822c90331)
- [connect_to_slo](data-sources--subnet--reference--group-001.md#canonical-46c21d95b78d70d1def68e910b0233161ae2b53bfcbab43b6caf63533f279f82)
- [isolated_nw](data-sources--subnet--reference--group-001.md#canonical-52a72afef5d80609be631d87cdedad1e613c3b2df7e3691ddc2ba10f609ce52c)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-dc679315875cb7cd67797244b7a0f2b81a17bf99c90ed8332706f02822c90331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-108f2089b348479f7a57a788dd54292138d1120521d182ae11bd9d164cc49cc6"></a>

## connect_to_layer2 — connect_to_layer2 / afe6d2000f89 / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- connect_to_layer2

<a id="canonical-26f9f34f10981515519c1105127f25bb8d2b58febe2242df1fc3be2dc2a5ee75"></a>

Type: `"single"`. Computed.

\[OneOf: connect\_to\_layer2, connect\_to\_slo, isolated\_nw\] Configuration parameter for connect
to layer2.

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

OneOf alternatives in this subsection:

- [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-26f9f34f10981515519c1105127f25bb8d2b58febe2242df1fc3be2dc2a5ee75)
- [connect_to_slo](data-sources--subnet--reference--group-001.md#canonical-f122e0bf3dbaf56391e21f105254761e98454bcb1b5f919f464d9d1a0550d7b7)
- [isolated_nw](data-sources--subnet--reference--group-001.md#canonical-9eb6fb15abbb9f68a6a087b6571477c65a76b6732bedd791e8bad4fe410685e8)

Select alternatives according to the provider validators above.

<a id="canonical-cfc1478f25d522de56471f87db974ffd95cef906339436196f1bea92b7055da4"></a>

## Direct properties — connect_to_layer2 / afe6d2000f89 / 3

- [layer2_intf_ref](data-sources--subnet--reference--group-001.md#canonical-73a08ceadfd01ddfa2abba216b9f6cc14bcf1e58e1cc99c22ca06e21bd27535a): complete subsection reference.

<a id="canonical-2ef19a71493de897cae58232860439b3ce3c04b15af250d3811dc44f6bcbd495"></a>

## Next pages — connect_to_layer2 / afe6d2000f89 / 4

- [connect_to_layer2.layer2_intf_ref](data-sources--subnet--reference--group-001.md#canonical-73a08ceadfd01ddfa2abba216b9f6cc14bcf1e58e1cc99c22ca06e21bd27535a)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-73a08ceadfd01ddfa2abba216b9f6cc14bcf1e58e1cc99c22ca06e21bd27535a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c454c019f51746497484593aeb3ab4f5a93b1c5b95119a02140df8f23d0a4f36"></a>

## connect_to_layer2.layer2_intf_ref — connect_to_layer2.layer2_intf_ref / c4bb52fe3323 / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-dc679315875cb7cd67797244b7a0f2b81a17bf99c90ed8332706f02822c90331)
- connect_to_layer2.layer2_intf_ref

<a id="canonical-df1a3bf76d7ae90e3e2a540ff625c5818f3b79abe9a04a328263a99be54ca6fb"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-23692b0481b6e5851462152d03699eb68686a03ebb84ee05c8ce97bf64e94e5b"></a>

## Direct properties — connect_to_layer2.layer2_intf_ref / c4bb52fe3323 / 3

<a id="canonical-90fb47b548d27018ecd835b54ae306f69fd84fee8e76613f1fb707a7c9baf34c"></a>

<a id="canonical-7e8f397a599e4460b635a6923cc67389779bc6ce570758c77118d10f3ab75033"></a>

## name property — connect_to_layer2.layer2_intf_ref / c4bb52fe3323 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-c914ba0090edffbab721d17332a1f44fa278f8178e760a1d3916cc712f6b1deb"></a>

<a id="canonical-932626f54d253e868e948c52a3e872353e64f724ee0a294b5a2a12bf393b9ec3"></a>

## namespace property — connect_to_layer2.layer2_intf_ref / c4bb52fe3323 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-b46605545ec57009a6e66cf40888450794c0784da6a2b0ef1b0837c1172ea4c7"></a>

<a id="canonical-ebf9660cd061590b228b407e2c9765000dc8563435b3b6d48b251deab8c56e9d"></a>

## tenant property — connect_to_layer2.layer2_intf_ref / c4bb52fe3323 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-880fc7871181d37bede38cdb5c277274a7557c0d7ff5d8d9390001aed9171b25"></a>

## Next pages — connect_to_layer2.layer2_intf_ref / c4bb52fe3323 / 7

- [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-dc679315875cb7cd67797244b7a0f2b81a17bf99c90ed8332706f02822c90331)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-46c21d95b78d70d1def68e910b0233161ae2b53bfcbab43b6caf63533f279f82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3acb4a20bddefdfc9ed12ac6c0f4933770325fda26621b95b0f009970df1690f"></a>

## connect_to_slo — connect_to_slo / ff074931058e / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- connect_to_slo

<a id="canonical-f122e0bf3dbaf56391e21f105254761e98454bcb1b5f919f464d9d1a0550d7b7"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for connect to slo.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-ca1989eaf6dd336f83b38cb1d89a2e356cec74b77191ca0d3621175c57a22ebd"></a>

## Direct properties — connect_to_slo / ff074931058e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb8bd1a5ea930ab44c37ee3a98309532e71e78475bdddc3045137015aa2509ae"></a>

## Next pages — connect_to_slo / ff074931058e / 4

- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-52a72afef5d80609be631d87cdedad1e613c3b2df7e3691ddc2ba10f609ce52c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f87d86d72b7f8ef098e0c47de5371b2cb5cca21f1443beb4dc44fb4e65f21d2"></a>

## isolated_nw — isolated_nw / 20502cf1a049 / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- isolated_nw

<a id="canonical-9eb6fb15abbb9f68a6a087b6571477c65a76b6732bedd791e8bad4fe410685e8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for isolated nw.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-abcf6555dfcbe213b97374f75a4b77395562b76ec0da77e70ad1812e21e60ecc"></a>

## Direct properties — isolated_nw / 20502cf1a049 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16b1c841252b112c560f687c114101db49fcd61d10ef0820f8a47db1e1bf90b3"></a>

## Next pages — isolated_nw / 20502cf1a049 / 4

- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f76eff54fd80403bde2bc06796f719a56dca02a2a0849a9d97431faa2fc1be9"></a>

## site_subnet_params — site_subnet_params / 3ba15df5094d / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- site_subnet_params

<a id="canonical-cc8f12fc682cd50d1191b00ddf09a24fe00d90e0fe1f3173748d00adddbefb32"></a>

Type: `"list"`. Computed.

Site Subnet Parameters. Configure subnet parameters per site.

Upstream description:

Configure subnet parameters per site.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-d1d19d4cf95187ecd0a7f5b0d7d10b7c8f5d6209a20f40d3bbaf8fda50536d66"></a>

## Direct properties — site_subnet_params / 3ba15df5094d / 3

- [dhcp](data-sources--subnet--reference--group-001.md#canonical-61852139e81000de48f0cef25ae74d07abec53d8c4f8697d6795eaed16e4536f): complete subsection reference.

- [site](data-sources--subnet--reference--group-001.md#canonical-52c3328fcc73d67c8332514123668299bb7c308545df0eae7b416d5a7db32d63): complete subsection reference.

- [static_ip](data-sources--subnet--reference--group-001.md#canonical-bf23adf1194c23de29dde4fd471009633426f8e7f10ddd746f941cc45a17a183): complete subsection reference.

- [subnet_dhcp_server_params](data-sources--subnet--reference--group-001.md#canonical-af8592bb82ff5974c6113da4a7a2e8be77e152567e4a06d2731ea5d354d3f414): complete subsection reference.

<a id="canonical-33fc179065d69c6268256bcaa0151b59724d158700637d019d12ce9c5e759f23"></a>

## Next pages — site_subnet_params / 3ba15df5094d / 4

- [site_subnet_params.dhcp](data-sources--subnet--reference--group-001.md#canonical-61852139e81000de48f0cef25ae74d07abec53d8c4f8697d6795eaed16e4536f)
- [site_subnet_params.site](data-sources--subnet--reference--group-001.md#canonical-52c3328fcc73d67c8332514123668299bb7c308545df0eae7b416d5a7db32d63)
- [site_subnet_params.static_ip](data-sources--subnet--reference--group-001.md#canonical-bf23adf1194c23de29dde4fd471009633426f8e7f10ddd746f941cc45a17a183)
- [site_subnet_params.subnet_dhcp_server_params](data-sources--subnet--reference--group-001.md#canonical-af8592bb82ff5974c6113da4a7a2e8be77e152567e4a06d2731ea5d354d3f414)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-61852139e81000de48f0cef25ae74d07abec53d8c4f8697d6795eaed16e4536f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f3aea62c35a7d1f8c351a6d5626220a6b17fd58b75142dbc5183599884d9c0b"></a>

## site_subnet_params.dhcp — site_subnet_params.dhcp / 76df40071ecd / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- site_subnet_params.dhcp

<a id="canonical-046d5db191553fe217a2702bb8d4030550b2d4797693e558454b6ff09bffd87a"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-a5f1409097449f27e574896a889ffa2950ba4fa5c58ea758c5155f751e7c264b"></a>

## Direct properties — site_subnet_params.dhcp / 76df40071ecd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15a341d1bd0a10d44ccc91290ed04f8d6442cfe76e626d12fe0656ab664664ed"></a>

## Next pages — site_subnet_params.dhcp / 76df40071ecd / 4

- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-52c3328fcc73d67c8332514123668299bb7c308545df0eae7b416d5a7db32d63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a169382996d993339549a2dcd4b6656e1aadcde5607181568415ffd65daffa7b"></a>

## site_subnet_params.site — site_subnet_params.site / 0f80a489a99b / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- site_subnet_params.site

<a id="canonical-03d2a7c19c6a54a9d1388cf66928f23e6482e91fa1157f856d244a7e7dc75dd2"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-64477d144de347d488fde6df02eda866df5aa614c20a09ecca4ae442244f07e8"></a>

## Direct properties — site_subnet_params.site / 0f80a489a99b / 3

<a id="canonical-11677e8018a4e019c6f3e5ce7e94053c3b82a9b3578aa8707fa938692fbdb8a0"></a>

<a id="canonical-25f4710ce981164e9f3d0c924e67652a9b992023cf842b539b6725dda639ccd8"></a>

## name property — site_subnet_params.site / 0f80a489a99b / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1526f6097e2ced7b7c0979099b0ec61f924cb6c6130743337872303b7d146ca2"></a>

<a id="canonical-f715d8c11b506f6d65c1cd12a53d34403d489aca369f1bef64201f07d51393cd"></a>

## namespace property — site_subnet_params.site / 0f80a489a99b / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-5987869f16aa5f021666e2ab48b85187420e59f25579cb944d7ead8f8e358cfd"></a>

<a id="canonical-842327d724ad1bf33b13432a3bff4e3bc2787b27a07df0a0beadf5a64815698f"></a>

## tenant property — site_subnet_params.site / 0f80a489a99b / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3bca70ffa609be32783722df2cd5a12de3b947bef946eec7d870ff0cb3c94bb1"></a>

## Next pages — site_subnet_params.site / 0f80a489a99b / 7

- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-bf23adf1194c23de29dde4fd471009633426f8e7f10ddd746f941cc45a17a183"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1281af2a88c17386896dd819ff74b25b35ba1ed894c9e655ac9075908813eefe"></a>

## site_subnet_params.static_ip — site_subnet_params.static_ip / c4c3e1075859 / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- site_subnet_params.static_ip

<a id="canonical-ec3519da365c2d46e2668d0bcdb708254da756377adb88111d08921957f0ecbd"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-89cde42b9126a232bb245edbf686e8f8685aca92107e2d4c7e8e06814048cd65"></a>

## Direct properties — site_subnet_params.static_ip / c4c3e1075859 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3b377ca3659fc66f0404c43d11d63aad55768ac0225b72db9120f8ce5c7e89d7"></a>

## Next pages — site_subnet_params.static_ip / c4c3e1075859 / 4

- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-af8592bb82ff5974c6113da4a7a2e8be77e152567e4a06d2731ea5d354d3f414"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f68230495abdbe2257efa84d4803525373866310e84bcf99b8f7094b8795a7ff"></a>

## site_subnet_params.subnet_dhcp_server_params — site_subnet_params.subnet_dhcp_server_params / 7888732c71f8 / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- site_subnet_params.subnet_dhcp_server_params

<a id="canonical-5d93bf23ff8d4a277f58d2030c4dd33800a6d57ebd20ad232be8c3d1d4b27505"></a>

Type: `"single"`. Computed.

Subnet DHCP parameters will be a subset of network\_interface.dhcpserverparameterstype as all
features in network\_interface.dhcpserverparameterstype may not be supported in a subnet.

Upstream description:

Subnet DHCP parameters will be a subset of network\_interface.dhcpserverparameterstype as all
features in network\_interface.dhcpserverparameterstype may not be supported in a subnet.

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

<a id="canonical-b0b58bd0bbdd25b698cb01845f29d9b2b86e40c8967312d1e6f6d73d90750af0"></a>

## Direct properties — site_subnet_params.subnet_dhcp_server_params / 7888732c71f8 / 3

- [dhcp_networks](data-sources--subnet--reference--group-001.md#canonical-8d87cd76815ce19cc1bef99e1fa5529512ab4ec7531a8e73131b94c87b33f325): complete subsection reference.

<a id="canonical-8da48a7d56b65d4971c1b2e450deec2c11ed9c3e7478f62b5e0f52985df1331c"></a>

## Next pages — site_subnet_params.subnet_dhcp_server_params / 7888732c71f8 / 4

- [site_subnet_params.subnet_dhcp_server_params.dhcp_networks](data-sources--subnet--reference--group-001.md#canonical-8d87cd76815ce19cc1bef99e1fa5529512ab4ec7531a8e73131b94c87b33f325)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)

<a id="canonical-8d87cd76815ce19cc1bef99e1fa5529512ab4ec7531a8e73131b94c87b33f325"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34f89920bb8eb7bc5a34d37e7b5b523aac4ec0846c230c4a8586edba104a7264"></a>

## site_subnet_params.subnet_dhcp_server_params.dhcp_networks — site_subnet_params.subnet_dhcp_server_params.dhcp_networks / 8419d58b5d2f / 2

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-c70d6debec98139114312c88618b1d3cd9e0ea3c6f250e87f9252a47de09e9c3)
- [site_subnet_params.subnet_dhcp_server_params](data-sources--subnet--reference--group-001.md#canonical-af8592bb82ff5974c6113da4a7a2e8be77e152567e4a06d2731ea5d354d3f414)
- site_subnet_params.subnet_dhcp_server_params.dhcp_networks

<a id="canonical-071c528788baeac407c6e99f46b8d096bca481af30feec38d6c9872710bfbf07"></a>

Type: `"list"`. Computed.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7a813fa2922e35899c19d03e0040b288b9f9535692b27b3a785532fad6ba3336"></a>

## Direct properties — site_subnet_params.subnet_dhcp_server_params.dhcp_networks / 8419d58b5d2f / 3

<a id="canonical-49dacef18027e3a6cc17295db1c9d2c13b98d42b9ac07845c30e4cb1fae4b523"></a>

<a id="canonical-fb0f6edc6e023534299b46bc8676638ecb03916517eb63393844e87f71f31705"></a>

## network_prefix property — site_subnet_params.subnet_dhcp_server_params.dhcp_networks / 8419d58b5d2f / 4

Type: `"string"`. Computed.

Exclusive with \[\] Network prefix for subnet.

Upstream description:

Exclusive with \[\] Network prefix for subnet.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2e42a402ca2f4f3139ec5be536286755a7805a9d8147eb9b120e480b3188d9d1"></a>

## Next pages — site_subnet_params.subnet_dhcp_server_params.dhcp_networks / 8419d58b5d2f / 5

- [site_subnet_params.subnet_dhcp_server_params](data-sources--subnet--reference--group-001.md#canonical-af8592bb82ff5974c6113da4a7a2e8be77e152567e4a06d2731ea5d354d3f414)
- [xcsh_subnet](../data-sources/subnet.md#canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624)
