---
page_title: "xcsh_app_type reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type reference."
---

# xcsh_app_type reference

<a id="canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c964687c8674dda7c98c2c65aca6db2b101cfbd92a093668f4e752afb9eff17c"></a>

## Property reference — Property reference / 14bba3307417 / 2

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
- Property reference

<a id="canonical-04cc5af73a4ac20265f18ec074a1694a26bbc430e9a960cf95528db0efad6f51"></a>

## Direct properties — Property reference / 14bba3307417 / 3

<a id="canonical-84cf1472c495eabe27aa955f0728a4c4d5ca0b26a7dfafc322f4ed0d7a4cbb3f"></a>

<a id="canonical-fdad887056fa5b344f212f497a82c01b65223630bd65f7dd4e3e0651e221f743"></a>

## annotations property — Property reference / 14bba3307417 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-ebe940e35a9b4f909a32e5ae6bf9c1c235b368a6eb8ff1b1139d98c7c9234f6b): complete subsection reference.

<a id="canonical-cbc9f9849bbacae75aa209a8d1feada05af4d520dd2308f66906f2de1ee95b90"></a>

<a id="canonical-5968657290f33beaae40857394bddb8e2926691fffc9dc4ddee52fd30fd25dcb"></a>

## description property — Property reference / 14bba3307417 / 5

Type: `"string"`. Optional.

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

<a id="canonical-a644d0d7cbef91f5f3f182a346fadbe558ef9228f554f562ecbf11e2e9a45e6c"></a>

<a id="canonical-aa5d42bc5fc98924b46adb41adcafc85877fa1cc7ac61b07559b0aa42bc3d0d2"></a>

## disable property — Property reference / 14bba3307417 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [features](resources--app_type--reference--group-001.md#canonical-75641fd00001519d99f2c3756755f540dbcda18a23274d78fa89e4356aca8e68): complete subsection reference.

<a id="canonical-ef3a9b8aa25d116a958d32a51d70fe858cf45e40fa4aaa7bdcfeb9d064577d9e"></a>

<a id="canonical-20560eb37ea92c53f0223539e2eca2a9bccbdbd31c3fda86849a51671bd2e137"></a>

## id property — Property reference / 14bba3307417 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1e31b3260034dc456e70b8f239afaa19bcfd6d509d356a82a27adf94c14fa332"></a>

<a id="canonical-f3d53e76a7f944b42632fc2288d8cba479d1e12016a414eb8cc7ccaccdd4e608"></a>

## labels property — Property reference / 14bba3307417 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-65e7add93e7bb879455701d398312de28fc904302c2600c248f5f13e1180e8e4"></a>

<a id="canonical-64d516d0d1703aca438679a219e8b762fc8d7a0fcc7d3b6f8519b60a9efb0dfb"></a>

## name property — Property reference / 14bba3307417 / 9

Type: `"string"`. Required.

Name of the App Type. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

<a id="canonical-5b8f36065830ab135f82f873bf1d4c2b303ca271555c35268d1650cb8112594c"></a>

<a id="canonical-9da806c16225f3b51d04ba47bfc087ee88e5befa64da3e9f42a2d575778ed04c"></a>

## namespace property — Property reference / 14bba3307417 / 10

Type: `"string"`. Required.

Namespace where the App Type is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [timeouts](resources--app_type--reference--group-001.md#canonical-f628b03181ef3059af8d0ec014e64e5d8a030eedcf7fc01fbb97dff98510810a): complete subsection reference.

<a id="canonical-a786f02cb9f447eab059e3ba03eb96bbc82d88ecaa1c65d09bb3cbdd661b2025"></a>

## All schema paths — Property reference / 14bba3307417 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--app_type--reference--group-001.md#canonical-84cf1472c495eabe27aa955f0728a4c4d5ca0b26a7dfafc322f4ed0d7a4cbb3f) |
| `business_logic_markup_setting` | [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-f29528f6c9bbaa45a37846a139a0ea38cdbfdfaa9ce19c40eed1cab57f77b64a) |
| `business_logic_markup_setting.disable_spec` | [business_logic_markup_setting.disable_spec](resources--app_type--reference--group-001.md#canonical-f2eea2e68dc5138e950ce12db3b0fecdcf371fe67d82b444d6db79fcb9919a9f) |
| `business_logic_markup_setting.discovered_api_settings` | [business_logic_markup_setting.discovered_api_settings](resources--app_type--reference--group-001.md#canonical-c53fb657eac9ef46fdaa443ea14b07fa14ed12bfb33b4ed395872b20c304a3d9) |
| `business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis` | [business_logic_markup_setting.discovered_api_settings.purge_duration_for_inactive_discovered_apis](resources--app_type--reference--group-001.md#canonical-399444798904510fab9757c6be8a2d0c8c5fa2790e96a93f49dba56d086bd8f8) |
| `business_logic_markup_setting.enable` | [business_logic_markup_setting.enable](resources--app_type--reference--group-001.md#canonical-eab262d9cf4f1af09facf6d81ea2fcf82ac118206d6ba1629a7f9220ada7834c) |
| `description` | [description](resources--app_type--reference--group-001.md#canonical-cbc9f9849bbacae75aa209a8d1feada05af4d520dd2308f66906f2de1ee95b90) |
| `disable` | [disable](resources--app_type--reference--group-001.md#canonical-a644d0d7cbef91f5f3f182a346fadbe558ef9228f554f562ecbf11e2e9a45e6c) |
| `features` | [features](resources--app_type--reference--group-001.md#canonical-0d80f43b5b725d5280614e6d4f61665776266571a07f68b8c7043afc6dd63c6c) |
| `features.type` | [features.type](resources--app_type--reference--group-001.md#canonical-d7d268cffa24a74e16496f60cfad967946c46ad72a56558bae1477b02ec11b32) |
| `id` | [id](resources--app_type--reference--group-001.md#canonical-ef3a9b8aa25d116a958d32a51d70fe858cf45e40fa4aaa7bdcfeb9d064577d9e) |
| `labels` | [labels](resources--app_type--reference--group-001.md#canonical-1e31b3260034dc456e70b8f239afaa19bcfd6d509d356a82a27adf94c14fa332) |
| `name` | [name](resources--app_type--reference--group-001.md#canonical-65e7add93e7bb879455701d398312de28fc904302c2600c248f5f13e1180e8e4) |
| `namespace` | [namespace](resources--app_type--reference--group-001.md#canonical-5b8f36065830ab135f82f873bf1d4c2b303ca271555c35268d1650cb8112594c) |
| `timeouts` | [timeouts](resources--app_type--reference--group-001.md#canonical-e1767fc8e9d50caf44b850a76a72bbfee2ff1bee55b56526ff103a9396c2b072) |
| `timeouts.create` | [timeouts.create](resources--app_type--reference--group-001.md#canonical-1e2f74cde9c0716390894d2ed2208aef8c862c2046a7525088a0f13067e82c2e) |
| `timeouts.delete` | [timeouts.delete](resources--app_type--reference--group-001.md#canonical-edfdea0f9387284d6a4d8670210a069fd06a2ddbb15e761b3e19ad744d80213f) |
| `timeouts.read` | [timeouts.read](resources--app_type--reference--group-001.md#canonical-f0cc800fbf8709025fcacb8084a04f44062afc162fe546c99ef4916c014b70b3) |
| `timeouts.update` | [timeouts.update](resources--app_type--reference--group-001.md#canonical-a10b8cf53d020c94e9704db1971e3260ee3e80aa1e7d7f21470ed4fc7a0360a2) |

<a id="canonical-bc2b09a5ce7b25e4521db5569eea2e224530367e617c59602fabbfc337ab15fd"></a>

## Next pages — Property reference / 14bba3307417 / 12

- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-ebe940e35a9b4f909a32e5ae6bf9c1c235b368a6eb8ff1b1139d98c7c9234f6b)
- [features](resources--app_type--reference--group-001.md#canonical-75641fd00001519d99f2c3756755f540dbcda18a23274d78fa89e4356aca8e68)
- [timeouts](resources--app_type--reference--group-001.md#canonical-f628b03181ef3059af8d0ec014e64e5d8a030eedcf7fc01fbb97dff98510810a)
- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)

<a id="canonical-ebe940e35a9b4f909a32e5ae6bf9c1c235b368a6eb8ff1b1139d98c7c9234f6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee4ab257eeaedb101c041baa9c946b6f539fc660b5a5213cbd705a7f3dc2c4e8"></a>

## business_logic_markup_setting — business_logic_markup_setting / dc6191c5d0ec / 2

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
- [Property reference](resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- business_logic_markup_setting

<a id="canonical-f29528f6c9bbaa45a37846a139a0ea38cdbfdfaa9ce19c40eed1cab57f77b64a"></a>

Type: `"object"`. single nested block, Optional.

Settings specifying how API Discovery will be performed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
business_logic_markup_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-2fd8c876a91e97a34b342fbad972ed61a90243122b4447f9b22b5bf249312bc8"></a>

## Direct properties — business_logic_markup_setting / dc6191c5d0ec / 3

- [disable_spec](resources--app_type--reference--group-001.md#canonical-d48d7a007b08f8ddaf0d25127828f6fa286b936bb813dc6d7dea6d90dfe3678f): complete subsection reference.

- [discovered_api_settings](resources--app_type--reference--group-001.md#canonical-132daf54188c2c81c41e3d493c38b460155bc11faacbe130b5c617069cef3990): complete subsection reference.

- [enable](resources--app_type--reference--group-001.md#canonical-e637883b4e69f4049ecc1e8b047435e5a7fd61e94c00fa2c742789134c7bb860): complete subsection reference.

<a id="canonical-880d8befc620d08bcf66074278442956d67e4b0e0ab99de8647a9b38ee1d4f3a"></a>

## Next pages — business_logic_markup_setting / dc6191c5d0ec / 4

- [business_logic_markup_setting.disable_spec](resources--app_type--reference--group-001.md#canonical-d48d7a007b08f8ddaf0d25127828f6fa286b936bb813dc6d7dea6d90dfe3678f)
- [business_logic_markup_setting.discovered_api_settings](resources--app_type--reference--group-001.md#canonical-132daf54188c2c81c41e3d493c38b460155bc11faacbe130b5c617069cef3990)
- [business_logic_markup_setting.enable](resources--app_type--reference--group-001.md#canonical-e637883b4e69f4049ecc1e8b047435e5a7fd61e94c00fa2c742789134c7bb860)
- [Property reference](resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)

<a id="canonical-d48d7a007b08f8ddaf0d25127828f6fa286b936bb813dc6d7dea6d90dfe3678f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10cca99c3420f9a72bb2bf8312ac19f715d7c4923f5b0accce4566d13fe69b39"></a>

## business_logic_markup_setting.disable_spec — business_logic_markup_setting.disable_spec / fd5359ac6490 / 2

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
- [Property reference](resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-ebe940e35a9b4f909a32e5ae6bf9c1c235b368a6eb8ff1b1139d98c7c9234f6b)
- business_logic_markup_setting.disable_spec

<a id="canonical-f2eea2e68dc5138e950ce12db3b0fecdcf371fe67d82b444d6db79fcb9919a9f"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-30c6ed549d51e6c876230bd55b6cf64be382b8ffbcb0c963028b89f1aa6b5db3"></a>

## Direct properties — business_logic_markup_setting.disable_spec / fd5359ac6490 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71e8aff1587af840e4e955427410a289cc8ef7ab68994c9a0dc829cf94aff8a8"></a>

## Next pages — business_logic_markup_setting.disable_spec / fd5359ac6490 / 4

- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-ebe940e35a9b4f909a32e5ae6bf9c1c235b368a6eb8ff1b1139d98c7c9234f6b)
- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)

<a id="canonical-132daf54188c2c81c41e3d493c38b460155bc11faacbe130b5c617069cef3990"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a55826f11e790fde214044f7c77abf88f04190dca8f14387c9e3c7ded8575f87"></a>

## business_logic_markup_setting.discovered_api_settings — business_logic_markup_setting.discovered_api_settings / f5b3d1e17d19 / 2

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
- [Property reference](resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-ebe940e35a9b4f909a32e5ae6bf9c1c235b368a6eb8ff1b1139d98c7c9234f6b)
- business_logic_markup_setting.discovered_api_settings

<a id="canonical-c53fb657eac9ef46fdaa443ea14b07fa14ed12bfb33b4ed395872b20c304a3d9"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("purge_duration_for_inactive_discovered_apis")}
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
discovered_api_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-1c801bb82aad3b684de01a56238fb027aa7034f23964a87cfeef6d00d0234bf1"></a>

## Direct properties — business_logic_markup_setting.discovered_api_settings / f5b3d1e17d19 / 3

<a id="canonical-399444798904510fab9757c6be8a2d0c8c5fa2790e96a93f49dba56d086bd8f8"></a>

<a id="canonical-58405aa2190ab3c3ec102b25ee5848c007a4f3b9ac0961cf4e72b0ea2b333769"></a>

## purge_duration_for_inactive_discovered_apis property — business_logic_markup_setting.discovered_api_settings / f5b3d1e17d19 / 4

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-ddaee6d9da043cd218ca7685e627bd0e4b742025aa8b146403d95f8a63c45de8"></a>

## Next pages — business_logic_markup_setting.discovered_api_settings / f5b3d1e17d19 / 5

- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-ebe940e35a9b4f909a32e5ae6bf9c1c235b368a6eb8ff1b1139d98c7c9234f6b)
- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)

<a id="canonical-e637883b4e69f4049ecc1e8b047435e5a7fd61e94c00fa2c742789134c7bb860"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ff3b1633214a22c3c763e272eb5837c2d5cafebbf4a2f73660a0b19e52cee51"></a>

## business_logic_markup_setting.enable — business_logic_markup_setting.enable / 17937d14a093 / 2

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
- [Property reference](resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-ebe940e35a9b4f909a32e5ae6bf9c1c235b368a6eb8ff1b1139d98c7c9234f6b)
- business_logic_markup_setting.enable

<a id="canonical-eab262d9cf4f1af09facf6d81ea2fcf82ac118206d6ba1629a7f9220ada7834c"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable = {}
```

<a id="canonical-130f19c2bd9537dac414d3a2f1a2c4879b85da4e9dc766312f730b87e47161c1"></a>

## Direct properties — business_logic_markup_setting.enable / 17937d14a093 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f44363529eb137cf2a1c0d66b25769795f1536b4e6571517d7eccbd28f2bb0b"></a>

## Next pages — business_logic_markup_setting.enable / 17937d14a093 / 4

- [business_logic_markup_setting](resources--app_type--reference--group-001.md#canonical-ebe940e35a9b4f909a32e5ae6bf9c1c235b368a6eb8ff1b1139d98c7c9234f6b)
- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)

<a id="canonical-75641fd00001519d99f2c3756755f540dbcda18a23274d78fa89e4356aca8e68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f18c1266e633a17d65c7fd6e7dc9f6b85f68638a31d13a08ab74970c06c8617"></a>

## features — features / 15ee5da06c25 / 2

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
- [Property reference](resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- features

<a id="canonical-0d80f43b5b725d5280614e6d4f61665776266571a07f68b8c7043afc6dd63c6c"></a>

Type: `"object"`. list nested block, Optional.

List of various advanced security features enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
features {
  # Configure direct properties listed below.
}
```

<a id="canonical-40e4095fe57d798e4247872f7a6de5bd27084e1bc4b454051df41ed027018976"></a>

## Direct properties — features / 15ee5da06c25 / 3

<a id="canonical-d7d268cffa24a74e16496f60cfad967946c46ad72a56558bae1477b02ec11b32"></a>

<a id="canonical-4bfca72457aa28d18a2c6c39aebe8f730c4a88ba92684f1dc71194c958c2731e"></a>

## type property — features / 15ee5da06c25 / 4

Type: `"string"`. Optional.

\[Enum:
BUSINESS\_LOGIC\_MARKUP|TIMESERIES\_ANOMALY\_DETECTION|PER\_REQ\_ANOMALY\_DETECTION|USER\_BEHAVIOR\_ANALYSIS\]
Enumeration for advanced security features supported API Discovery enables generation of model for
various API interactions between services of App type. Enable analysis of timeseries for various
metric collected like requests, errors, latency etc. Enable anomaly detection per API request, i.e.
Possible values are \`BUSINESS\_LOGIC\_MARKUP\`, \`TIMESERIES\_ANOMALY\_DETECTION\`,
\`PER\_REQ\_ANOMALY\_DETECTION\`, \`USER\_BEHAVIOR\_ANALYSIS\`. Defaults to
\`BUSINESS\_LOGIC\_MARKUP\`.

Upstream description:

Enumeration for advanced security features supported

API Discovery enables generation of model for various API interactions between services of App type.
Enable analysis of timeseries for various metric collected like requests, errors, latency etc.
Enable anomaly detection per API request, i.e. The probability density function (PDF) charts
generation for API endpoints Enable user behavior analysis.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BUSINESS_LOGIC_MARKUP",
    "TIMESERIES_ANOMALY_DETECTION",
    "PER_REQ_ANOMALY_DETECTION",
    "USER_BEHAVIOR_ANALYSIS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BUSINESS_LOGIC_MARKUP",
  "enum": [
    "BUSINESS_LOGIC_MARKUP",
    "TIMESERIES_ANOMALY_DETECTION",
    "PER_REQ_ANOMALY_DETECTION",
    "USER_BEHAVIOR_ANALYSIS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-498c5037e4435235c4f42d5056d36778d5eb0368be820c90808ec665975d44c2"></a>

## Next pages — features / 15ee5da06c25 / 5

- [Property reference](resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)

<a id="canonical-f628b03181ef3059af8d0ec014e64e5d8a030eedcf7fc01fbb97dff98510810a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c698d0e9c17b876de1ffcf4a228f59818dc896856eb6727aa52666d5f40831e"></a>

## timeouts — timeouts / 3026feaa76e8 / 2

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
- [Property reference](resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- timeouts

<a id="canonical-e1767fc8e9d50caf44b850a76a72bbfee2ff1bee55b56526ff103a9396c2b072"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-ff4a9fb34c26db956000eb3f9e1525dd1a3939748c30b2867fa4a2e2bd737b81"></a>

## Direct properties — timeouts / 3026feaa76e8 / 3

<a id="canonical-1e2f74cde9c0716390894d2ed2208aef8c862c2046a7525088a0f13067e82c2e"></a>

<a id="canonical-77fe4951c24cfe102e6ae9cdf9e66f28b10be8aec9d8e1baa064b019ac418bd4"></a>

## create property — timeouts / 3026feaa76e8 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-edfdea0f9387284d6a4d8670210a069fd06a2ddbb15e761b3e19ad744d80213f"></a>

<a id="canonical-cbe82377a1c601f8db7b7e17e89eb0d8f7fe3e11024b019b5673ed10350e6fc8"></a>

## delete property — timeouts / 3026feaa76e8 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-f0cc800fbf8709025fcacb8084a04f44062afc162fe546c99ef4916c014b70b3"></a>

<a id="canonical-26ab1a103c6b7052d415f453b9c1c65f6043d2b909f70d405f8e5e8a76d0f49d"></a>

## read property — timeouts / 3026feaa76e8 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-a10b8cf53d020c94e9704db1971e3260ee3e80aa1e7d7f21470ed4fc7a0360a2"></a>

<a id="canonical-ec515976576e83e3f41d9080690e0e3f30f4f5c1d0ee50f45d80ba4e8e0ef8e6"></a>

## update property — timeouts / 3026feaa76e8 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a0eaeae31f2642f74a5999528c661025e6e9361670d40e5e25553680dc8b744b"></a>

## Next pages — timeouts / 3026feaa76e8 / 8

- [Property reference](resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
