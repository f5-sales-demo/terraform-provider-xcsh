---
page_title: "xcsh_protocol_inspection reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection reference."
---

# xcsh_protocol_inspection reference

<a id="canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60016d7e0d08e604ac5ad4bcfe25954f716674dba7ea1d896308a948b556f1a1"></a>

## Property reference — Property reference / 84df3bdea478 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
- Property reference

<a id="canonical-3880dd701fbe0a45a6732376be5d225a5e97711dd13f0fa19ae47af6da7c98e7"></a>

## Direct properties — Property reference / 84df3bdea478 / 3

<a id="canonical-9339e1abd5f92aa550737ae7aecf9317bbea0e2688d7758cc4bc46261eaccc43"></a>

<a id="canonical-c73587899dc908393666ec516bbbef56619ccf6b71ad389a6c6f11a75b67096a"></a>

## action property — Property reference / 84df3bdea478 / 4

Type: `"string"`. Computed.

\[Enum: ALLOW|DENY|DROP\] Action after inspection - ALLOW: Allow Allow traffic - DENY: Deny Throw
RST error for TCP and ICMP error for UDP - DROP: DROP Silently drop traffic. Possible values are
\`ALLOW\`, \`DENY\`, \`DROP\`. Defaults to \`ALLOW\`. Server applies default when omitted.

Upstream description:

Action after inspection

&#8203;- ALLOW: Allow

Allow traffic &#8203;- DENY: Deny

Throw RST error for TCP and ICMP error for UDP &#8203;- DROP: DROP

Silently drop traffic.

Receipt-pinned upstream constraints:

```json
{
  "default": "ALLOW",
  "enum": [
    "ALLOW",
    "DENY",
    "DROP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-02740a947d83de36ef90b7bf3648c7f485b533d58390ae7d16d841b79741cc4f"></a>

<a id="canonical-8dea9d733913168dba98a4193a0920bc731142cfe5bb81cf6757fcb05fa12ac3"></a>

## annotations property — Property reference / 84df3bdea478 / 5

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

<a id="canonical-916494966b846c458eddcada4e7be8198f57529ec47ae8a331ae98133db28f47"></a>

<a id="canonical-838cd1ec2aeed08d892f9cbdeeb96499373e49a35ab33650396615c6a456d807"></a>

## description property — Property reference / 84df3bdea478 / 6

Type: `"string"`. Computed.

Description of the ProtocolInspection.

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

- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-6b239633331aad186c765c0999395a86d4499fc49ef59e4f24dad20c3bab205f): complete subsection reference.

- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-12bd750e4ba08c97d1fd09b490943d2855cc4b18a9ac2b25bf3c9a2c6d273328): complete subsection reference.

<a id="canonical-cb9c9a6195ad3603368197e558566b232a09bdf0e546cb13e6ecf5dd13878844"></a>

<a id="canonical-f7c162a0e57f4c882b2411ce8e5e9b04356ef491b9613b09af9b74a4afc21dc1"></a>

## id property — Property reference / 84df3bdea478 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-c05c7593991d9e8f66031b02e1830a13e0d140dfbc74aaee5a4a0deefad3358e"></a>

<a id="canonical-8bc2e03e37f8369b4cfd7e46bbb8367d13d08ceb1f612662e40408bcd2fb98c1"></a>

## labels property — Property reference / 84df3bdea478 / 8

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

<a id="canonical-ae73257b2a25c1c46b3d4ec10d21100e0c877935ebcb3bf3987f1dca350294a2"></a>

<a id="canonical-e2a6f6750dcab05005543ed645747ecbc37cdd7fe9ee525fa5f2efa0a2257095"></a>

## name property — Property reference / 84df3bdea478 / 9

Type: `"string"`. Required.

Name of the ProtocolInspection.

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

<a id="canonical-46dadbe77e9b16e2f18de7756ce242337379cc2a89b107777ffa05e7d13538e7"></a>

<a id="canonical-7beead420b1fa8c1c8acfbe79a51fdabcb7dbccd01055b62555af54c1882bdb3"></a>

## namespace property — Property reference / 84df3bdea478 / 10

Type: `"string"`. Required.

Namespace where the ProtocolInspection exists.

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

<a id="canonical-aac6a310cd012be78ead8b530ed4c9043ddb19b34df64314c745159982be2ca7"></a>

## All schema paths — Property reference / 84df3bdea478 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--protocol_inspection--reference--group-001.md#canonical-9339e1abd5f92aa550737ae7aecf9317bbea0e2688d7758cc4bc46261eaccc43) |
| `annotations` | [annotations](data-sources--protocol_inspection--reference--group-001.md#canonical-02740a947d83de36ef90b7bf3648c7f485b533d58390ae7d16d841b79741cc4f) |
| `description` | [description](data-sources--protocol_inspection--reference--group-001.md#canonical-916494966b846c458eddcada4e7be8198f57529ec47ae8a331ae98133db28f47) |
| `enable_disable_compliance_checks` | [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-9776b6c59fdd47da17ba4d5a4354dd7368a84afbd225c6828bfdb1eed11c6c38) |
| `enable_disable_compliance_checks.disable_compliance_checks` | [enable_disable_compliance_checks.disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-2b3b24b15f8f6919341c349220236a9e834b401bc8dc5f6da1bca7761f2a6ac0) |
| `enable_disable_compliance_checks.enable_compliance_checks` | [enable_disable_compliance_checks.enable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-dba8b911d72a03bac1766992bbb3958fa4f31d581083a57a7a8ee3fdaa207fbb) |
| `enable_disable_compliance_checks.enable_compliance_checks.name` | [enable_disable_compliance_checks.enable_compliance_checks.name](data-sources--protocol_inspection--reference--group-001.md#canonical-bf96d2a0bde954c7975899089c6fd4f67150b0ab16a7b7b675058b2d55a5ea2e) |
| `enable_disable_compliance_checks.enable_compliance_checks.namespace` | [enable_disable_compliance_checks.enable_compliance_checks.namespace](data-sources--protocol_inspection--reference--group-001.md#canonical-ac615fc7293000851427e1839071645323f157eefbcd52db982a8fd5587aeb57) |
| `enable_disable_compliance_checks.enable_compliance_checks.tenant` | [enable_disable_compliance_checks.enable_compliance_checks.tenant](data-sources--protocol_inspection--reference--group-001.md#canonical-b02289c16798b107f23c37edd9273452ef3a296815e46f49197d0f5b0a32b6d6) |
| `enable_disable_signatures` | [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-0d77b3bf7086f968b7ec8bef58c9d84adfff8f4ea1bf15fa66ea0f64140ddd41) |
| `enable_disable_signatures.disable_signature` | [enable_disable_signatures.disable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-168c03655b94656071623c591986df8fb3f056611aba69acb095b0a36946bb8d) |
| `enable_disable_signatures.enable_signature` | [enable_disable_signatures.enable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-aef16fd6bdcd9bd5492082fcdf5e69f09ed326092f42be2120c01b9f8001fe52) |
| `id` | [id](data-sources--protocol_inspection--reference--group-001.md#canonical-cb9c9a6195ad3603368197e558566b232a09bdf0e546cb13e6ecf5dd13878844) |
| `labels` | [labels](data-sources--protocol_inspection--reference--group-001.md#canonical-c05c7593991d9e8f66031b02e1830a13e0d140dfbc74aaee5a4a0deefad3358e) |
| `name` | [name](data-sources--protocol_inspection--reference--group-001.md#canonical-ae73257b2a25c1c46b3d4ec10d21100e0c877935ebcb3bf3987f1dca350294a2) |
| `namespace` | [namespace](data-sources--protocol_inspection--reference--group-001.md#canonical-46dadbe77e9b16e2f18de7756ce242337379cc2a89b107777ffa05e7d13538e7) |

<a id="canonical-e80f769c54f23508958f2f8ba2e0b0fba4c11f52f88ba19c543e558b9f7c1162"></a>

## Next pages — Property reference / 84df3bdea478 / 12

- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-6b239633331aad186c765c0999395a86d4499fc49ef59e4f24dad20c3bab205f)
- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-12bd750e4ba08c97d1fd09b490943d2855cc4b18a9ac2b25bf3c9a2c6d273328)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)

<a id="canonical-6b239633331aad186c765c0999395a86d4499fc49ef59e4f24dad20c3bab205f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47647b6e94088e1d240d9814fe2005435775131afec377ee8d35f17bab21f09a"></a>

## enable_disable_compliance_checks — enable_disable_compliance_checks / 04a629745b5b / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb)
- enable_disable_compliance_checks

<a id="canonical-9776b6c59fdd47da17ba4d5a4354dd7368a84afbd225c6828bfdb1eed11c6c38"></a>

Type: `"single"`. Computed.

Enable Disable Compliance Checks Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compliance_check_choice": "[\"disable_compliance_checks\",\"enable_compliance_checks\"]"
}
```

<a id="canonical-d230dd74974bea1891e5ee58aca94c4163760ec39e029624ef39fcce3e851985"></a>

## Direct properties — enable_disable_compliance_checks / 04a629745b5b / 3

- [disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-ab66187914f1fa7c45e1b001ce91a88b3dd71eb8a9308c53a403a03e638fdd0c): complete subsection reference.

- [enable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-851ea412956ebddf36eb47a0999010f5ed8b0cd2af149e5368e2596aee5d0f76): complete subsection reference.

<a id="canonical-c31eaf14c995d72cec7ef86253b4f0580aeb6a1980a2f7973f736cead5225676"></a>

## Next pages — enable_disable_compliance_checks / 04a629745b5b / 4

- [enable_disable_compliance_checks.disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-ab66187914f1fa7c45e1b001ce91a88b3dd71eb8a9308c53a403a03e638fdd0c)
- [enable_disable_compliance_checks.enable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-851ea412956ebddf36eb47a0999010f5ed8b0cd2af149e5368e2596aee5d0f76)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)

<a id="canonical-ab66187914f1fa7c45e1b001ce91a88b3dd71eb8a9308c53a403a03e638fdd0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a021a3ffa3651a05518dc630ec82171c5dcaf768e9b4dee08ce7977645f32a3"></a>

## enable_disable_compliance_checks.disable_compliance_checks — enable_disable_compliance_checks.disable_compliance_checks / 3774bc622bb9 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb)
- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-6b239633331aad186c765c0999395a86d4499fc49ef59e4f24dad20c3bab205f)
- enable_disable_compliance_checks.disable_compliance_checks

<a id="canonical-2b3b24b15f8f6919341c349220236a9e834b401bc8dc5f6da1bca7761f2a6ac0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable compliance checks.

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

<a id="canonical-742b930c00617c561a3ecf0f82341d594ed3eb9860589424f22f1f6dca9daa5a"></a>

## Direct properties — enable_disable_compliance_checks.disable_compliance_checks / 3774bc622bb9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-676d0b4a4f7eb482183454ed6f004fbc498f95ff209d7917b51d6a3e1be9cffc"></a>

## Next pages — enable_disable_compliance_checks.disable_compliance_checks / 3774bc622bb9 / 4

- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-6b239633331aad186c765c0999395a86d4499fc49ef59e4f24dad20c3bab205f)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)

<a id="canonical-851ea412956ebddf36eb47a0999010f5ed8b0cd2af149e5368e2596aee5d0f76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2377b98827b70839d86c5a92fd87396bb71df81682c34cf29b55115dae6fb47"></a>

## enable_disable_compliance_checks.enable_compliance_checks — enable_disable_compliance_checks.enable_compliance_checks / 79ba4881a5cf / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb)
- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-6b239633331aad186c765c0999395a86d4499fc49ef59e4f24dad20c3bab205f)
- enable_disable_compliance_checks.enable_compliance_checks

<a id="canonical-dba8b911d72a03bac1766992bbb3958fa4f31d581083a57a7a8ee3fdaa207fbb"></a>

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

<a id="canonical-0be18b7ba61baa6c3f45aa6eaf6ce1c4ed221a8f3199a280d016ba76e4a34602"></a>

## Direct properties — enable_disable_compliance_checks.enable_compliance_checks / 79ba4881a5cf / 3

<a id="canonical-bf96d2a0bde954c7975899089c6fd4f67150b0ab16a7b7b675058b2d55a5ea2e"></a>

<a id="canonical-063f4e76f496384d0dfc44c0085244451177cfb32767507107a4a52706c6e46d"></a>

## name property — enable_disable_compliance_checks.enable_compliance_checks / 79ba4881a5cf / 4

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

<a id="canonical-ac615fc7293000851427e1839071645323f157eefbcd52db982a8fd5587aeb57"></a>

<a id="canonical-ca1aee134a02eadac5ffe5a03ffaa6abae58c1e82caabea9cbd6118b824119ed"></a>

## namespace property — enable_disable_compliance_checks.enable_compliance_checks / 79ba4881a5cf / 5

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

<a id="canonical-b02289c16798b107f23c37edd9273452ef3a296815e46f49197d0f5b0a32b6d6"></a>

<a id="canonical-eb7b14a0ae5e19cd0cc6dc2a9b00522ba23c603dd91f9a1ee358ad1c6d8a4e9a"></a>

## tenant property — enable_disable_compliance_checks.enable_compliance_checks / 79ba4881a5cf / 6

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

<a id="canonical-b85780581d23ccc94dcd39963a8515068799278b7a0b3db898b467b4b59a56ef"></a>

## Next pages — enable_disable_compliance_checks.enable_compliance_checks / 79ba4881a5cf / 7

- [enable_disable_compliance_checks](data-sources--protocol_inspection--reference--group-001.md#canonical-6b239633331aad186c765c0999395a86d4499fc49ef59e4f24dad20c3bab205f)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)

<a id="canonical-12bd750e4ba08c97d1fd09b490943d2855cc4b18a9ac2b25bf3c9a2c6d273328"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac396637ec66840e4b237cedf71c287f51932899c22bfdebeead6c676f37cc9a"></a>

## enable_disable_signatures — enable_disable_signatures / a023113079dc / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb)
- enable_disable_signatures

<a id="canonical-0d77b3bf7086f968b7ec8bef58c9d84adfff8f4ea1bf15fa66ea0f64140ddd41"></a>

Type: `"single"`. Computed.

Configuration parameter for enable disable signatures.

Upstream description:

Enable Disable Signature Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signature_choice": "[\"disable_signature\",\"enable_signature\"]"
}
```

<a id="canonical-7603cc25a1670d543f0350195c704569137c5f0d4ba61c8d050b77c904a101f2"></a>

## Direct properties — enable_disable_signatures / a023113079dc / 3

- [disable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-71e644301628c27e6998bce41008b697d628379d0d52569a3f21106f074b5560): complete subsection reference.

- [enable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-f4c69e84220f66c8f1db2e10bd2a3257e66dd93b8c7def264401079b00727ff9): complete subsection reference.

<a id="canonical-8564fefb009277fda2816ccad5423ed03e21d90ab6258824f0a22fc506f8ec95"></a>

## Next pages — enable_disable_signatures / a023113079dc / 4

- [enable_disable_signatures.disable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-71e644301628c27e6998bce41008b697d628379d0d52569a3f21106f074b5560)
- [enable_disable_signatures.enable_signature](data-sources--protocol_inspection--reference--group-001.md#canonical-f4c69e84220f66c8f1db2e10bd2a3257e66dd93b8c7def264401079b00727ff9)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)

<a id="canonical-71e644301628c27e6998bce41008b697d628379d0d52569a3f21106f074b5560"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8de80aa72201365b1f4d5fa2ad82096e7393fcf83a545631cc31e9dcff9f507e"></a>

## enable_disable_signatures.disable_signature — enable_disable_signatures.disable_signature / 65624c37eff1 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb)
- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-12bd750e4ba08c97d1fd09b490943d2855cc4b18a9ac2b25bf3c9a2c6d273328)
- enable_disable_signatures.disable_signature

<a id="canonical-168c03655b94656071623c591986df8fb3f056611aba69acb095b0a36946bb8d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable signature.

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

<a id="canonical-5700f4f7b26dbaf490938e682d6ac588b0aa9663e5cd4cd6108a8b8fe052d9aa"></a>

## Direct properties — enable_disable_signatures.disable_signature / 65624c37eff1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40e68a56682267de2476dafd1ac487b05cc4101a9105b7ba1bea70f7187522bc"></a>

## Next pages — enable_disable_signatures.disable_signature / 65624c37eff1 / 4

- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-12bd750e4ba08c97d1fd09b490943d2855cc4b18a9ac2b25bf3c9a2c6d273328)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)

<a id="canonical-f4c69e84220f66c8f1db2e10bd2a3257e66dd93b8c7def264401079b00727ff9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4432b9fb2965cd8ea43577b3299634288edd4e84596fd552a3628bbc3956ac0e"></a>

## enable_disable_signatures.enable_signature — enable_disable_signatures.enable_signature / 98bc8c42c1e9 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
- [Property reference](data-sources--protocol_inspection--reference--group-001.md#canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb)
- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-12bd750e4ba08c97d1fd09b490943d2855cc4b18a9ac2b25bf3c9a2c6d273328)
- enable_disable_signatures.enable_signature

<a id="canonical-aef16fd6bdcd9bd5492082fcdf5e69f09ed326092f42be2120c01b9f8001fe52"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable signature.

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

<a id="canonical-d17f21f482166f31b6f581abc3553ccca568a4438ac2c4c988f2fb3819e33176"></a>

## Direct properties — enable_disable_signatures.enable_signature / 98bc8c42c1e9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66c1ff13aff917e8bedd0c8666600ae8fb9ad14b28a9b32f644e6b2f8427b339"></a>

## Next pages — enable_disable_signatures.enable_signature / 98bc8c42c1e9 / 4

- [enable_disable_signatures](data-sources--protocol_inspection--reference--group-001.md#canonical-12bd750e4ba08c97d1fd09b490943d2855cc4b18a9ac2b25bf3c9a2c6d273328)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
