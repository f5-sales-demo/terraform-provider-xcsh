---
page_title: "xcsh_cloud_elastic_ip reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip reference."
---

# xcsh_cloud_elastic_ip reference

<a id="canonical-c0f91f91cfb563bfe37a3cba69683a2e56a66d3b7c53c84011ee5b0bf7d437a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2fc429775e590a4cd3e7df4daea561f614b6cf60b95db0809b34f91fadf59bc"></a>

## Property reference — Property reference / 1449045ba552 / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)
- Property reference

<a id="canonical-766270201a68c873dd7ed8ae32479d1572ae60853e681dd82d377fa00730dc3b"></a>

## Direct properties — Property reference / 1449045ba552 / 3

<a id="canonical-8dad9198ceb1b839699d23ac17f9e2846597f22934a791f1456694ffcc4e71d9"></a>

<a id="canonical-8cb401d6e083fd372193d47f90b70a4ebd750b1b79895905c5cc1e480aa0adf9"></a>

## annotations property — Property reference / 1449045ba552 / 4

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

<a id="canonical-9c0ca862ac520f887f87559504aac2f9f2a4e0b862606d5ddad7098dca9e019c"></a>

<a id="canonical-43b07a84fc5da1c5ba196a06b86f55cbd50e342d794f1ea80f7d56b688c59957"></a>

## description property — Property reference / 1449045ba552 / 5

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

<a id="canonical-b98053d343e44b57ce5fdee8ca36ad5f2738999cc0781c4798e6713bc6a85bdc"></a>

<a id="canonical-e7d576a6eb4bb7ea86bcdd7078c37ec60c5d23c6265086e5990c0d9d2fad653b"></a>

## disable property — Property reference / 1449045ba552 / 6

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

<a id="canonical-77281b5d1ef0735a7e40cefaeea7940155578dfa5b7aa52bf3d6e2084cf72f70"></a>

<a id="canonical-08acd712b86af8e62f57fc86647d01d71b9e8aae0c271cd2d1226827856577ae"></a>

## id property — Property reference / 1449045ba552 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-cfcfb1efbe8dc1a30667e6e5ab891b01cb9aae1b4ce443916eba1cd3b455bccd"></a>

<a id="canonical-af0456d8ad286dba04b1eeeaa242b99a347e4a9d0a1ed4d5b635102670573ec3"></a>

## item_count property — Property reference / 1449045ba552 / 8

Type: `"number"`. Required.

Number of Elastic Ips / Public Ips associated with this object per Node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8),
}
```

<a id="canonical-93dad7a0903cef61c599e731203dcb41ffea73807f9770b79dd424219fe11e26"></a>

<a id="canonical-6583f5d254b8402afc0bc860260af3b22d8641412ea11fddc51ca56281c6d588"></a>

## labels property — Property reference / 1449045ba552 / 9

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

<a id="canonical-2a3fc12720f7c7ed0ac343ccbc94ae80fe59def74e504f54edaec0cc8a0d5e9d"></a>

<a id="canonical-0ebbbed864a7f287b69b58ad15b71283b2cd9a306e8942fecf86bcfd477c4792"></a>

## name property — Property reference / 1449045ba552 / 10

Type: `"string"`. Required.

Name of the Cloud Elastic IP. Must be unique within the namespace.

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

<a id="canonical-ba6d249763efdb1edae72ac3b7abc89d868ab8d288fa77266500c256440c7aec"></a>

<a id="canonical-9743d7d4d285314d4f81f0954a5ec74416ff9ff4f046ac53da9588b2881b7ad2"></a>

## namespace property — Property reference / 1449045ba552 / 11

Type: `"string"`. Required.

Namespace where the Cloud Elastic IP is created.

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

- [site_ref](resources--cloud_elastic_ip--reference--group-001.md#canonical-926665e656a8bcb5559ec1428ab91d5ef989bbd8885dbe5660828c4ffef32658): complete subsection reference.

- [timeouts](resources--cloud_elastic_ip--reference--group-001.md#canonical-3f376f40f07b520ce97c3025cdf8a5f122242ac2535b5269b40eba1b30c9f2d3): complete subsection reference.

<a id="canonical-0f5c02923ded2de789ef7be99efe370252812f3ac7bf645ff21783c701431d4a"></a>

## All schema paths — Property reference / 1449045ba552 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_elastic_ip--reference--group-001.md#canonical-8dad9198ceb1b839699d23ac17f9e2846597f22934a791f1456694ffcc4e71d9) |
| `description` | [description](resources--cloud_elastic_ip--reference--group-001.md#canonical-9c0ca862ac520f887f87559504aac2f9f2a4e0b862606d5ddad7098dca9e019c) |
| `disable` | [disable](resources--cloud_elastic_ip--reference--group-001.md#canonical-b98053d343e44b57ce5fdee8ca36ad5f2738999cc0781c4798e6713bc6a85bdc) |
| `id` | [id](resources--cloud_elastic_ip--reference--group-001.md#canonical-77281b5d1ef0735a7e40cefaeea7940155578dfa5b7aa52bf3d6e2084cf72f70) |
| `item_count` | [item_count](resources--cloud_elastic_ip--reference--group-001.md#canonical-cfcfb1efbe8dc1a30667e6e5ab891b01cb9aae1b4ce443916eba1cd3b455bccd) |
| `labels` | [labels](resources--cloud_elastic_ip--reference--group-001.md#canonical-93dad7a0903cef61c599e731203dcb41ffea73807f9770b79dd424219fe11e26) |
| `name` | [name](resources--cloud_elastic_ip--reference--group-001.md#canonical-2a3fc12720f7c7ed0ac343ccbc94ae80fe59def74e504f54edaec0cc8a0d5e9d) |
| `namespace` | [namespace](resources--cloud_elastic_ip--reference--group-001.md#canonical-ba6d249763efdb1edae72ac3b7abc89d868ab8d288fa77266500c256440c7aec) |
| `site_ref` | [site_ref](resources--cloud_elastic_ip--reference--group-001.md#canonical-483c50227c5598fb89be78e67d9a6b2208c7bbaea0dd0319af3713d1c15e86a4) |
| `site_ref.kind` | [site_ref.kind](resources--cloud_elastic_ip--reference--group-001.md#canonical-7e4c0bae685d7c05a19d614e7725233fea3bd4c7af90b4d2de2d93a7ed407ddf) |
| `site_ref.name` | [site_ref.name](resources--cloud_elastic_ip--reference--group-001.md#canonical-636e6251ae5b1d04971310d9f993f84bcc2b8396c4dbdeefacc9b38e81888d90) |
| `site_ref.namespace` | [site_ref.namespace](resources--cloud_elastic_ip--reference--group-001.md#canonical-9e0dcd514e962afcdd004872578ad9d42e43539f7442051f52e92840e7838101) |
| `site_ref.tenant` | [site_ref.tenant](resources--cloud_elastic_ip--reference--group-001.md#canonical-e3a73a201079020e29fab640a0c15a053db49eeaea9f0a7ba22fca185f7b9ff2) |
| `site_ref.uid` | [site_ref.uid](resources--cloud_elastic_ip--reference--group-001.md#canonical-a4c9142b5ca6f4a5bb4a3d5e5e5d98b10fb63545ab7aa7b336b7e38714750840) |
| `timeouts` | [timeouts](resources--cloud_elastic_ip--reference--group-001.md#canonical-b18666758ea07ecc09543f90028bff6df9e5bd994d00e7fc6484615adc931714) |
| `timeouts.create` | [timeouts.create](resources--cloud_elastic_ip--reference--group-001.md#canonical-b0414658d78815e467e00fc5b08930869c1d48483d8fbcf48203324edd784b49) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_elastic_ip--reference--group-001.md#canonical-521ef3b88f7a70d9220c699ef33af0632af03484ade0616477065b5f2ec22b68) |
| `timeouts.read` | [timeouts.read](resources--cloud_elastic_ip--reference--group-001.md#canonical-65b3015bab33e09d8db33551d92d55d69d49c41bc5f7e69932ad98a161767c0b) |
| `timeouts.update` | [timeouts.update](resources--cloud_elastic_ip--reference--group-001.md#canonical-a49d6872e57b2bb567839593798f012bdbd9e0181fc8e24d183138eb97418a50) |

<a id="canonical-959af1da239e2fd37428ad1c3966048fad63a59eabc29714939b7f23a96827a6"></a>

## Next pages — Property reference / 1449045ba552 / 13

- [site_ref](resources--cloud_elastic_ip--reference--group-001.md#canonical-926665e656a8bcb5559ec1428ab91d5ef989bbd8885dbe5660828c4ffef32658)
- [timeouts](resources--cloud_elastic_ip--reference--group-001.md#canonical-3f376f40f07b520ce97c3025cdf8a5f122242ac2535b5269b40eba1b30c9f2d3)
- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)

<a id="canonical-926665e656a8bcb5559ec1428ab91d5ef989bbd8885dbe5660828c4ffef32658"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b658770a14840fe3e0e9393666982476baba0f31cfcb9aff98aaacbc1c905cf"></a>

## site_ref — site_ref / e9b8a1771038 / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)
- [Property reference](resources--cloud_elastic_ip--reference--group-001.md#canonical-c0f91f91cfb563bfe37a3cba69683a2e56a66d3b7c53c84011ee5b0bf7d437a4)
- site_ref

<a id="canonical-483c50227c5598fb89be78e67d9a6b2208c7bbaea0dd0319af3713d1c15e86a4"></a>

Type: `"object"`. list nested block, Optional.

Site to which this cloud elastic IP object is attached.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
site_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-d35d5e663edf3f8e86711b6a008f17a4cb2eae8db8295e2afaf0a3c851004cc2"></a>

## Direct properties — site_ref / e9b8a1771038 / 3

<a id="canonical-7e4c0bae685d7c05a19d614e7725233fea3bd4c7af90b4d2de2d93a7ed407ddf"></a>

<a id="canonical-f308096341bc40c194ff29579ab73719b34e7b887395cb0945fcec63b9ddf091"></a>

## kind property — site_ref / e9b8a1771038 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-636e6251ae5b1d04971310d9f993f84bcc2b8396c4dbdeefacc9b38e81888d90"></a>

<a id="canonical-7bab3fd01a6d53110d0e7dc62c3cd4a40fddd1ba4058cb94a2e06a4f6f9a29f5"></a>

## name property — site_ref / e9b8a1771038 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-9e0dcd514e962afcdd004872578ad9d42e43539f7442051f52e92840e7838101"></a>

<a id="canonical-431e9a5c8e8c53b6a5d77c2bd129d5d05c4fc2635c91f5c4aa2d0092ba7672b2"></a>

## namespace property — site_ref / e9b8a1771038 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-e3a73a201079020e29fab640a0c15a053db49eeaea9f0a7ba22fca185f7b9ff2"></a>

<a id="canonical-75b8649704ae068a3e094e02a1de342f8dbe1c86118e9edaf58d623d3489cd9f"></a>

## tenant property — site_ref / e9b8a1771038 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-a4c9142b5ca6f4a5bb4a3d5e5e5d98b10fb63545ab7aa7b336b7e38714750840"></a>

<a id="canonical-04f7cecd55ae95c4316a9411e88186d76578ed1ce20ef7e7ce757fb490924f5f"></a>

## uid property — site_ref / e9b8a1771038 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-24a3a4f8d6d5bd8d577023bc7a5fcfc1ebedfaafb957b9c4f5584e8aa7775838"></a>

## Next pages — site_ref / e9b8a1771038 / 9

- [Property reference](resources--cloud_elastic_ip--reference--group-001.md#canonical-c0f91f91cfb563bfe37a3cba69683a2e56a66d3b7c53c84011ee5b0bf7d437a4)
- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)

<a id="canonical-3f376f40f07b520ce97c3025cdf8a5f122242ac2535b5269b40eba1b30c9f2d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cb65758b9c6ff648856005ef1016b0c17d76d015c61b9784fab0f51121b4bd6"></a>

## timeouts — timeouts / 6d651b505d8d / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)
- [Property reference](resources--cloud_elastic_ip--reference--group-001.md#canonical-c0f91f91cfb563bfe37a3cba69683a2e56a66d3b7c53c84011ee5b0bf7d437a4)
- timeouts

<a id="canonical-b18666758ea07ecc09543f90028bff6df9e5bd994d00e7fc6484615adc931714"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-4f03c7e60a00c78670662acffb7ae42e97c82963fdcc2961486b97df2902b44b"></a>

## Direct properties — timeouts / 6d651b505d8d / 3

<a id="canonical-b0414658d78815e467e00fc5b08930869c1d48483d8fbcf48203324edd784b49"></a>

<a id="canonical-86b0c3a9ec343acc40f1b099634a8cec207a8083673ad33d9afbca540b84947e"></a>

## create property — timeouts / 6d651b505d8d / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-521ef3b88f7a70d9220c699ef33af0632af03484ade0616477065b5f2ec22b68"></a>

<a id="canonical-c18951c9e69394b8d2ba1dda54e0a972cdfdd03059a9c7fd81c534cffc11edca"></a>

## delete property — timeouts / 6d651b505d8d / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-65b3015bab33e09d8db33551d92d55d69d49c41bc5f7e69932ad98a161767c0b"></a>

<a id="canonical-9c723755d8b26391a31bbb2140a21a4882cadec0c78d90965ebd49575e2815dc"></a>

## read property — timeouts / 6d651b505d8d / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-a49d6872e57b2bb567839593798f012bdbd9e0181fc8e24d183138eb97418a50"></a>

<a id="canonical-435c0fdba861ff65e24a426c6b92579aa87acdbe4475bb7e3ba6d6e6a919a062"></a>

## update property — timeouts / 6d651b505d8d / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-b70bfb616e1d37d24864f6ba8aac83761ae951261797b5c6261ba4ce76e38bdd"></a>

## Next pages — timeouts / 6d651b505d8d / 8

- [Property reference](resources--cloud_elastic_ip--reference--group-001.md#canonical-c0f91f91cfb563bfe37a3cba69683a2e56a66d3b7c53c84011ee5b0bf7d437a4)
- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)
