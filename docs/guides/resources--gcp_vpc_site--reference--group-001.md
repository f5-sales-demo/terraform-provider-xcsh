---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6cbba261f283f71a42ca946c853fcaf28ea262e48b099a4b3193b7962d51ae4"></a>

## Property reference — Property reference / ecb0e823acfb / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- Property reference

<a id="canonical-892525950e79f8f0be91e23630822ad6ef35cd694b921cca99ed67248ac35ebb"></a>

## Direct properties — Property reference / ecb0e823acfb / 3

<a id="canonical-1d2078f7cc7121d012ed8b5c512865da90c01697c8d557a689a1a6b0f981bc57"></a>

<a id="canonical-b0321827df9f6156f5542819f5addab3c4979b35242b345cac4d9904ca74522c"></a>

## address property — Property reference / ecb0e823acfb / 4

Type: `"string"`. Optional, Computed.

Site's geographical address that can be used to determine its latitude and longitude.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-18add28bee441d3d531062ac60749c6c76c95afcc4bfbf9d55cd6c1e538098b3): complete subsection reference.

<a id="canonical-4790436fa0ef414fe593dd98a170f0d5e2f7156e22404d49374dd2bf6103951d"></a>

<a id="canonical-a50490a4240fbe46f3de75720b8f23de531994b09cdb0bbf7019c43e44bdf7fe"></a>

## annotations property — Property reference / ecb0e823acfb / 5

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

- [block_all_services](resources--gcp_vpc_site--reference--group-001.md#canonical-530fa4fdc2c05b7ed2ac80e6e596f07e8e6046f437b406923387bbfbd1fb533d): complete subsection reference.

- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-909c249948254fbec30f63c678012c95380a2edad9b7709aa1548da60ead104e): complete subsection reference.

- [cloud_credentials](resources--gcp_vpc_site--reference--group-001.md#canonical-7605207075b1d259a748aabdbd6f06777b66701612652233add52b297a2687fd): complete subsection reference.

- [coordinates](resources--gcp_vpc_site--reference--group-001.md#canonical-1ee923a929358c616aab84a0ce7c58d4c6808078b75cfa0eaa85cc4699011ac8): complete subsection reference.

- [custom_dns](resources--gcp_vpc_site--reference--group-001.md#canonical-e8c0b808c7aa5333108e5f700c5fdbdbfd782aad00c4616db4c163488eefee4c): complete subsection reference.

- [default_blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-f437cf619b611a651738572b03d9cd7173644469efc9d4bcb509457d475b6c9f): complete subsection reference.

<a id="canonical-4553b17ca8ca673f3332f6befdd168d37a3855c6afa3bb451a7de1d585141d43"></a>

<a id="canonical-6367778d06bc6a4e0573af7f4a1b5c5caed302f337eaf38f2879df080daec1b3"></a>

## description property — Property reference / ecb0e823acfb / 6

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

<a id="canonical-005719b2630e69616394d1363d11d8ca30ff66414a58dc5d32768c6f5b1e417a"></a>

<a id="canonical-08fe399cc16e89be918621419e22213c1f66d97a92510e6df27cd90ba74895a1"></a>

## disable property — Property reference / ecb0e823acfb / 7

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

- [disable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-5d57af9ed321100aa53f9baf8cbe7bec1da9c0cf28e205bab76d17351fe60eb3): complete subsection reference.

<a id="canonical-195f4105e0ee183fe696d1826c908c77d8e9c9b4fb2d19ab7707519ce4de821a"></a>

<a id="canonical-8127bc7d361d16291fe1b104e2e301a1889fbed6e803d8e2f57d20da4e46c6ae"></a>

## disk_size property — Property reference / ecb0e823acfb / 8

Type: `"number"`. Optional, Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(64000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-51a5b7b4447038370ee22e5d23fefdc78599620aa70e55d74e5d89319ed1f8f2): complete subsection reference.

<a id="canonical-8a194ea0eda0e3a18074704cfc39f9eff4cb6f0957ae0e345046d98f11b8586f"></a>

<a id="canonical-9f27304b4342138350a33e293fff00f7dfcaf14db6c24e5b6638cb11d73048e0"></a>

## gcp_labels property — Property reference / ecb0e823acfb / 9

Type: `["map", "string"]`. Optional.

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

Upstream description:

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-4b13fabdd9c862dcf9aff8f02209f41f1925138dea3b45302b5e1f594f86b44e"></a>

<a id="canonical-e09a8475db598594208c8e0a34347a06dd5699a6a36a8e5294b919737a0700ce"></a>

## gcp_region property — Property reference / ecb0e823acfb / 10

Type: `"string"`. Required.

GCP Region. Name for GCP Region.

Upstream description:

Name for GCP Region.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-30e5d7907b1957a054f8a82ddb5d16f89714c5124a6221611def9cca73aff887"></a>

<a id="canonical-c5c62606b51781c0a58d543c052033ea6ce77077cd1fc830fb347a03e5aa2fb5"></a>

## id property — Property reference / ecb0e823acfb / 11

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110): complete subsection reference.

- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc): complete subsection reference.

<a id="canonical-5a934812b9ee77174ec0cdbd462bb759464afea4c2d773e32d1f14bdb12c87c8"></a>

<a id="canonical-04f46e3488731648dc885a2d64eeacfffcdf79a170cfb735b37164198e3c6c45"></a>

## instance_type property — Property reference / ecb0e823acfb / 12

Type: `"string"`. Required.

Select Instance size based on performance needed.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-91992f0d045bc68652deeaf62a5b3849c65f5447dcdd89b44c6c2b97307f6efb): complete subsection reference.

<a id="canonical-91485afe3b26851d2a6d3295435f77fe34fb3bbda400660749f104709326a074"></a>

<a id="canonical-3b85bedc86d83bb23e06a7aae536f5580f6b6cbe2d7af69040df0f4b1f3c3bd0"></a>

## labels property — Property reference / ecb0e823acfb / 13

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

- [log_receiver](resources--gcp_vpc_site--reference--group-003.md#canonical-4b1d87b55107bf727296e0f74b17d3ab715fa9ab4ce733b2a51208d76bc40e16): complete subsection reference.

- [logs_streaming_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-c144b8deb7de45a002c3f118644029a062245be86a46e5968311c10907617f46): complete subsection reference.

<a id="canonical-3aae36927c98842ce6a50195b151aeae1cf823867dc7bac52cb9c0eafb0a7c3f"></a>

<a id="canonical-576956b4582fb37ec82b5cd75f890bb2a7073952695e600e26ce62e27d56cf80"></a>

## name property — Property reference / ecb0e823acfb / 14

Type: `"string"`. Required.

Name of the GCP VPC Site. Must be unique within the namespace.

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

<a id="canonical-82a06d4ea4f1fe4df494c165cd1cb8ffb2659b0da939e44cf715c9a17a46ec5a"></a>

<a id="canonical-19e8c1e1ec791b30fbf1ee82b728417dceacecbf191f807e2c65a598bdf4731d"></a>

## namespace property — Property reference / ecb0e823acfb / 15

Type: `"string"`. Required.

Namespace where the GCP VPC Site is created.

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

- [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-321583ca89a2994ed6ab1089491a0a823abc73ba1e92d28770051589feecf20d): complete subsection reference.

- [os](resources--gcp_vpc_site--reference--group-003.md#canonical-eaae4c5a3696af2cdaa426ef1745cc75f8cc67b832eed73e1b84ffbca525c28b): complete subsection reference.

- [private_connect_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-60d557d0083e9ae6c90a25feb164999e2e99fcf4ad1174a90fa437468a9395e1): complete subsection reference.

- [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-2b156b46727ac0a204195e7e776cb3951197b186926382e9a9299718d824f79f): complete subsection reference.

<a id="canonical-a9fe566bd0f7c324583d1724835d557ccbfaf68ce093dd4399d2fad10e6dfdc6"></a>

<a id="canonical-470711e6916f043200336ca17cab00ac2b29dfb9747b0f45d29bb38f0bc65054"></a>

## ssh_key property — Property reference / ecb0e823acfb / 16

Type: `"string"`. Required.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [sw](resources--gcp_vpc_site--reference--group-003.md#canonical-267f967c717532fd9b7a7d530db900ad6f3ed5285284fb0cc81876b19418dec2): complete subsection reference.

- [timeouts](resources--gcp_vpc_site--reference--group-004.md#canonical-2298b57a32aa2303d946d8a0a735921119bf066b4e9cbcf62828b9ca2db0aad9): complete subsection reference.

- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657): complete subsection reference.

- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-a8bce3447b430d8e2a86211cdc8cabaa307dce6f5e9a69b93a189d355e794ad0): complete subsection reference.

<a id="canonical-f275748c202ae5d495a88683a5d938a643a2dbeaf343afd527281ff2052e28b5"></a>

## All schema paths — Property reference / ecb0e823acfb / 17

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--gcp_vpc_site--reference--group-001.md#canonical-1d2078f7cc7121d012ed8b5c512865da90c01697c8d557a689a1a6b0f981bc57) |
| `admin_password` | [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-c3a0b8541b1aa8db5bb99430dc51ea6903dc63b3eea554655814046a5a2d60dc) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-7e2f54e2ac4eb35999bfef936b7875d1f563c28540651dc6a52ca78f733aa704) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](resources--gcp_vpc_site--reference--group-001.md#canonical-762eaa2b46723d0c87de1f8dc31efd0cac50297ee10f5c1582069270557a7950) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](resources--gcp_vpc_site--reference--group-001.md#canonical-36f5497bb8cf9ba7188412cac2a1f0cc9ea5f702915192756e936c7a232c1abf) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](resources--gcp_vpc_site--reference--group-001.md#canonical-ef35964fe7bd2eb8a378894a9f210d79fdf6762f44560957ec03c318106cd386) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-0b1bf8d3e212b5ce76fcda970c8d85116a5ea3729d8beceaf547df8c4b1b9c52) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](resources--gcp_vpc_site--reference--group-001.md#canonical-d7e2b6825e8126064a998032cc29c6dd91b739d493f89dfcebd569bee153cb9f) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](resources--gcp_vpc_site--reference--group-001.md#canonical-3b801374994c8596a10e0399020b4381e90d6b38d3da065c72d4a7ca246c1202) |
| `annotations` | [annotations](resources--gcp_vpc_site--reference--group-001.md#canonical-4790436fa0ef414fe593dd98a170f0d5e2f7156e22404d49374dd2bf6103951d) |
| `block_all_services` | [block_all_services](resources--gcp_vpc_site--reference--group-001.md#canonical-048fca1479a480db37fe1cc5f8db572fff45c1e0152aed86724d150cab9fe316) |
| `blocked_services` | [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-40ed0b75b1e1bbb6826fe1612f67629b453b86e353fdb8d3c3ace972f7b99a7e) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-ce5d377d88dc557f4bec763db0193206571bc27c1ee52fa37019a3aadbf20a44) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--gcp_vpc_site--reference--group-001.md#canonical-56b33d280d462be0217555498c29f2402b1ef74df98a3a94c23feb2d5935a5fc) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--gcp_vpc_site--reference--group-001.md#canonical-3d4b763eb7a3f722ecdf5e7506faee64b9b5c066b1e31d0eea39a882ad80d448) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--gcp_vpc_site--reference--group-001.md#canonical-d37f2e179d76d69b044a65c18e8ea484b2316e60869d774588f2579067f3e1c2) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--gcp_vpc_site--reference--group-001.md#canonical-26a76fd918e6c2dc193e36927db59e1a8a0638dc85fef8bffe0928da2ceda8f8) |
| `cloud_credentials` | [cloud_credentials](resources--gcp_vpc_site--reference--group-001.md#canonical-a6b3e47ed3196bff633e8bee11ac1478a41dd25b83acc9291fccf74921e259ca) |
| `cloud_credentials.name` | [cloud_credentials.name](resources--gcp_vpc_site--reference--group-001.md#canonical-24cbd6943b0107a6703dca7e690e79b90fbc1bbf41819bff7b15fd2d1593f3d5) |
| `cloud_credentials.namespace` | [cloud_credentials.namespace](resources--gcp_vpc_site--reference--group-001.md#canonical-69c3fc7a78772052c34e5415e55add94d3471f682a6bc3fd4fbcad86e3ae7445) |
| `cloud_credentials.tenant` | [cloud_credentials.tenant](resources--gcp_vpc_site--reference--group-001.md#canonical-d199fe0d1a8ddb8d651d8a0f4fb65f9939523c3a636722f76c7f1c4dd32a0f5b) |
| `coordinates` | [coordinates](resources--gcp_vpc_site--reference--group-001.md#canonical-45eaca03320e66bdc6ad8eefdd9730b334211eb74b2aba12ece8005ea3242199) |
| `coordinates.latitude` | [coordinates.latitude](resources--gcp_vpc_site--reference--group-001.md#canonical-04839decdd9441917a0cac6115838829f6dfdd9a74692c4240a71490e9c80190) |
| `coordinates.longitude` | [coordinates.longitude](resources--gcp_vpc_site--reference--group-001.md#canonical-dbe18856ca8f74ee685853014deeaa4eb1f26acae4402b766fcde0d50eaeeeb2) |
| `custom_dns` | [custom_dns](resources--gcp_vpc_site--reference--group-001.md#canonical-c232eb908f219d5d9f81c0b765f3509f04213b581145dca6c1897f789a3c366f) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](resources--gcp_vpc_site--reference--group-001.md#canonical-ae00ebe6bb587207a157c8672b17a7e3b0cef1dadd3c99965b61b39cf9e7d1df) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](resources--gcp_vpc_site--reference--group-001.md#canonical-ae304ffc069929f5298383fe6840a878cff521b819bc52105fe7fbd421e4904c) |
| `default_blocked_services` | [default_blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-e9e2a082dee8738d6dc8de57f82a85fee2c74d73677af996801695fc7d9bfadb) |
| `description` | [description](resources--gcp_vpc_site--reference--group-001.md#canonical-4553b17ca8ca673f3332f6befdd168d37a3855c6afa3bb451a7de1d585141d43) |
| `disable` | [disable](resources--gcp_vpc_site--reference--group-001.md#canonical-005719b2630e69616394d1363d11d8ca30ff66414a58dc5d32768c6f5b1e417a) |
| `disable_encryption` | [disable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-6580d72f7bfe1f713ab6946a5892531818101ede667a92673b4c220773607b25) |
| `disk_size` | [disk_size](resources--gcp_vpc_site--reference--group-001.md#canonical-195f4105e0ee183fe696d1826c908c77d8e9c9b4fb2d19ab7707519ce4de821a) |
| `enable_encryption` | [enable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-614da5ada84d27e8085ce867a8308c6719470ae34ab110f6c2c37a5fb12652b9) |
| `enable_encryption.kms_key_resource_id` | [enable_encryption.kms_key_resource_id](resources--gcp_vpc_site--reference--group-001.md#canonical-7d7d9d265e0f61dc2cab4dd3996bba72c07bd4bada29d0858f046314d6232894) |
| `enable_encryption.kms_key_ring_id` | [enable_encryption.kms_key_ring_id](resources--gcp_vpc_site--reference--group-001.md#canonical-bfa21b384bdf37cc0da2cc65d55433da149e0ef1f3882d6e7780beaf8b962700) |
| `gcp_labels` | [gcp_labels](resources--gcp_vpc_site--reference--group-001.md#canonical-8a194ea0eda0e3a18074704cfc39f9eff4cb6f0957ae0e345046d98f11b8586f) |
| `gcp_region` | [gcp_region](resources--gcp_vpc_site--reference--group-001.md#canonical-4b13fabdd9c862dcf9aff8f02209f41f1925138dea3b45302b5e1f594f86b44e) |
| `id` | [id](resources--gcp_vpc_site--reference--group-001.md#canonical-30e5d7907b1957a054f8a82ddb5d16f89714c5124a6221611def9cca73aff887) |
| `ingress_egress_gw` | [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-c6ba3e3f7c8cfd37d26ea3e57fe7ab1b7a2f174266e35c4e1914080cc714210a) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-bfb1a0f2c82cda5cda05bf459016cdb72bf99d65da842121f625ba46aadcb916) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-17555357c5c1378d1e8c88a4f2d69b8ccf60233c48e2be8a5dd42299f4aed466) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--gcp_vpc_site--reference--group-002.md#canonical-2138d0f069cb0c87aec03b446f29aa3c29207ced44e73c50552bc58c2910a105) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-466f0d98a341c4e6a362913dd457834276b642e683b62963e599d9a346d708f7) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-5321823f1f15e918b22d1e870db957c2ffe1ccd58cbc3da3a796fbe332a60643) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-3e18e060fb0bd220f2044233bf27775082af2e8dae587c3492352d512ac2a96c) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-e01e25d7ac6bc1e84c85ce135320e9e06317e85a31f93f3ea2e241e8b2f8336e) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](resources--gcp_vpc_site--reference--group-002.md#canonical-3d5cdef532b3b658a72d04af05f7f5be6ed78cdceda3d96eb9427db96a49abb2) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-ee4e6c0e29432561b192b6100a18a01262d0c35f7374636731890505ec6c0430) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-4948b76283667f872336c949f8a73f1848be9e03db0910afa220b7218e6437bf) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-a880a857c5d584ab1c88d4d44fe6d6dc875d71745de890c8feab5add45023087) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-e58cebf376467eded302bf842364ed535f299c8f138ff6e30fcdb1b2698d95d5) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](resources--gcp_vpc_site--reference--group-002.md#canonical-1529af2b1c1953451c85bb5ef26235293fb12ba3b36d96da203eef02ae038ad9) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-1a559c7cba3eae0d49bd1a966d7ff53ec651e1503f98b9b6e90ee7e1491c55d7) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-22c290070148b0a53f52309bc3f7c8097f38a991669279bcaf5494559a92fcd3) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-501f103b08f1f067b8ec1f5805a743df1528958fb42a771e09de7dde8b5c9ff3) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](resources--gcp_vpc_site--reference--group-002.md#canonical-0046e8910fb8a8ecbf7dd3534ffead266b6610709d36ab5c76d74e3875a4f311) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-07557df8f45696b2a22e98b7b2fb12bc07f0cc940737353c6aaf721a0f9f21ca) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-684b0be60f4ba4f4d48263d96c042b4d56bee5c8b0c8489ed227a128eb66d25d) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-12cb66a01c69faea311cd5b34a6a4e1d7598b4a791f914c0a99f414a5f4e09ff) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](resources--gcp_vpc_site--reference--group-002.md#canonical-1d1d9e0d49f03b6179d04ee6b098e10c629c139f3983b0659e29a5f297ff918f) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-4409cf0daeff9527c2f810447d8f30f5f4379d960411db8c085105f96e387462) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-9c1a17a37b5159c4e6af233f3ea2d8be09f66bfc1a5498ff1d53db78b1aaa3ec) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-002.md#canonical-69db5b5870f8eed91524c2f8c7dd28d548b50e4b244e7571215cdb013945db75) |
| `ingress_egress_gw.gcp_certified_hw` | [ingress_egress_gw.gcp_certified_hw](resources--gcp_vpc_site--reference--group-001.md#canonical-de3cb092e4504a727f8d792728f1606e25da90f624db6c0a7edcd8f09ddb2980) |
| `ingress_egress_gw.gcp_zone_names` | [ingress_egress_gw.gcp_zone_names](resources--gcp_vpc_site--reference--group-001.md#canonical-1fcaafaaa27777ee094588a8a662656202b3ad1f29684c517837970e46454fb1) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-f6232a6befb2fa4743cb124e92f64ffcc0eb1b13d53e893227fe7489860a3033) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-a25718402d768acebe3c128d555b3ce87d2cccc30a13c385545b08c567615079) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-6faccf5f50b0fbfd41bc71df4318824e9a713db744c6741fd8a04b9156b4fc82) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-30a7611e7a9939dce982e704f025ad3d6996d4856a24847a645cf9b4e0d2bd11) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--gcp_vpc_site--reference--group-002.md#canonical-f562899604a617e600624a445bcb3afae218efadab8c5520b4e1cb4ec06919a1) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-e2c5f80ea18cc8ad7b1af5f17561930e5ee4dd8e8a87588f1a147fd108cc1cfd) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-13eec7d79efa48b34a29f3764caa4d82cfaca68331892b654b5336b7b4e0b757) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-9551ba1c94e9b38f2b07473534a8df451870eb4e7f6e2bad50b772304dde5733) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-bbcfaddae81b03d5ddc64a3099f6654bfd1f344b96a4d804599dc35aa2d6cd3d) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--gcp_vpc_site--reference--group-002.md#canonical-3214937b6989592445b33b0b90afe5700ecdd3ae8e7059b9949f18bcb0d8bcf4) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-ac388b74f7e2f5d60b6cceea6a9d589a5bb7a20b9d142bfabb2b79a4f8f5bd15) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-d9f871f76d9fab394140da3ab5975d0cea0be27244e3d81f8c2a3fd6fe958d86) |
| `ingress_egress_gw.inside_network` | [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-688741c6be5b66e14589972f597ef9095ba7fc1d056fa3da4f68a1effec97cf5) |
| `ingress_egress_gw.inside_network.existing_network` | [ingress_egress_gw.inside_network.existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-23da4928402c09228ed3df4e2bbeba49ae51992a5745b903d61376abccfdf295) |
| `ingress_egress_gw.inside_network.existing_network.name` | [ingress_egress_gw.inside_network.existing_network.name](resources--gcp_vpc_site--reference--group-002.md#canonical-5226a3218c5b4c1de6651ad472a4300af22c4e9dc1088510f6f7bb0f99343fd8) |
| `ingress_egress_gw.inside_network.new_network` | [ingress_egress_gw.inside_network.new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-86bfa1987f5c1fa62b8ec6f7546ab2f1936e7f8998a6d5be7aed077a098a6689) |
| `ingress_egress_gw.inside_network.new_network.name` | [ingress_egress_gw.inside_network.new_network.name](resources--gcp_vpc_site--reference--group-002.md#canonical-239b52592316f0c10e05080431c80e5a4e89c33257e2f3fde102c600a3d632a4) |
| `ingress_egress_gw.inside_network.new_network_autogenerate` | [ingress_egress_gw.inside_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-dc161626c8b1a2eeea7c3c86d4b74ea19072cddb7bb53283a727e4909a8b36f9) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-1f5a42ecc9836166f7d05a8587619b67138976bbfe4c6567b7a0f67385202d0f) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-f1d24af0e71b3982c6a0eafc9b59fdd30603c18c919b3a50377571b9c8201c52) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-e059814a50ae3fee71312bc400f89142391155727633f789fae00a93c308e027) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](resources--gcp_vpc_site--reference--group-002.md#canonical-01a47ae11c2f491aa43d586952cadd7b5b69212fc57b8d2079a6d206882b25bf) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-002.md#canonical-e8974fee05391eb3de2a4421d40c8c4756341a254b72b5729523a8b10c2ada9f) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-fe771a6ac685c7225f4e7de136e7381b1da2c54118d8147ab7485f0bd5df7d8d) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-002.md#canonical-154670badfd020ea043c1418db56d8e4af30331d3df5d1dad3c5a6d88b6afb19) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--gcp_vpc_site--reference--group-002.md#canonical-c7933a8dd8bbdf32e96f77a83ec0f7c08afed97868c62975add2e59cbbf99ad3) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--gcp_vpc_site--reference--group-002.md#canonical-3dd81986b40c36c9bf5cb4e39e19e4618a479689a0903dcf286aaaa76451b73a) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--gcp_vpc_site--reference--group-002.md#canonical-6afc29e9c08a3f62baf2a17be7bc1a7f2d011638ea825fc96694850e5ea4c656) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--gcp_vpc_site--reference--group-002.md#canonical-a930da51c4b37d2133c21fa7eb182bfd1f3661218f5614132d030ea1f4508455) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--gcp_vpc_site--reference--group-002.md#canonical-6c826cbfd0eedc39ce1207c4ca1de921f576f70b6e4f7c6a46aa49172d09c6d8) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-976f72460a7570d06775a6e472f2dced5e20792e802aa73a2268ca4be3f981e9) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-c81ec834fce277b239413aa90089352e682f42f489bc2e74fe0f4471f399175e) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-6b07e0650516894cc0385b8191e8d9617cbcd7f62185e355007f199f6e5c3c6e) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--gcp_vpc_site--reference--group-002.md#canonical-15bfbcfadc77b064ef8cd1e8893e284d0bddcc112325e4175269fadcce4a9b46) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-c16376571cd79ec1ed524108c28d07af69b5f63e6f578da4527be0f1bbea6f37) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--gcp_vpc_site--reference--group-002.md#canonical-7489972db52e02ef63ce31ebb066f18a468bf198589578debe6dafffb68e3479) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-488b42146d17984186295ce78d62460fd298a0da407db930394c6a207b152287) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--gcp_vpc_site--reference--group-002.md#canonical-1633e3856623517327f69a21c753363aecd13454137d217194ff313742435d39) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-4e69b3ee4801194b6a31086e6fd82f1f8db01b2e446081e8fb36195c88e9c6c9) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--gcp_vpc_site--reference--group-002.md#canonical-26d7b0d8aa65542646e79e959d6f9e497a94680d9b3dca65c055580869bdb731) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--gcp_vpc_site--reference--group-002.md#canonical-cac937c00045c8f407cd86fd838dcf3337477fe364052ccc3baa212d477a4788) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-f5a3834b9c9286f60ea28cfea73da85b9bb4c1e4ac8829b2dce0e9f1fdb51038) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-2551ac5faccad616bc98aa48db994d68f935788c978bf767f4921033eb9b868e) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--gcp_vpc_site--reference--group-002.md#canonical-ef1c0aaa625a8e96944187b927c6a1ef99d38fa1675cbf4c88aa485eaae7e1a7) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--gcp_vpc_site--reference--group-002.md#canonical-2c14dc8eb59474be45b08f4485491b0bb2a7ca4d58c2241f647264e0d52bcf87) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-085035890a888e3762ef6c21539aa4bea318c5d0d7e025c825486c4b52a3299c) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--gcp_vpc_site--reference--group-002.md#canonical-c44ea89329a0f77dc3c87009a0474722b52d04491f2e2c9aca777a20336e12e6) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--gcp_vpc_site--reference--group-002.md#canonical-505929013fc2037925c7c76e5c0f12788af48dc034856846d6880ac9b45fce60) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-86dbae8132f87431ef54afd137385dd386c5cc0f0769e9ff1ae5b98b9ce220b4) |
| `ingress_egress_gw.inside_subnet` | [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-420aacc24ff8ab865eefbc5bc007b11a3c034e8fc757507f75cba043cdfef8d9) |
| `ingress_egress_gw.inside_subnet.existing_subnet` | [ingress_egress_gw.inside_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-b73998122c553819b61386fc7eabb4a5adc8ae7d4eb5794a192f502fc5319506) |
| `ingress_egress_gw.inside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.existing_subnet.subnet_name](resources--gcp_vpc_site--reference--group-002.md#canonical-11237b3012e6657b15c1889e9bc896f85a98eb5cf585866f70ca5027f3aabc9b) |
| `ingress_egress_gw.inside_subnet.new_subnet` | [ingress_egress_gw.inside_subnet.new_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-6c1ccaf5f55d40aea33b0eb4a895f64e2d12649a07752fb2405d4a9fdb386154) |
| `ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-ff733535f87f2758f3e9b630d1c12d40eaef8f66bddd943f19d5f94346907893) |
| `ingress_egress_gw.inside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.new_subnet.subnet_name](resources--gcp_vpc_site--reference--group-002.md#canonical-fe5422ddc18180b47f3b141198fc8fe06f170ee9786976f59fcc9f4de5f94cba) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](resources--gcp_vpc_site--reference--group-002.md#canonical-54cc3c33b03c3cb3a886df18b96c8c80ff6a9107745a8fb7ab5a112f6867ef2c) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](resources--gcp_vpc_site--reference--group-002.md#canonical-d20d7af51f7feb3f13e5c8b9f2c44f98d42561acbfda71e323032b8f06f0180b) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](resources--gcp_vpc_site--reference--group-002.md#canonical-c63bd61d1bc31d7dff51c34e0f1254e3a3f51a6a10207372c6ed162ea80b3d6f) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-01912b556a4c47b59abb1efea77b039733661938973940dc313f390c520e83ae) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](resources--gcp_vpc_site--reference--group-002.md#canonical-3846f730f8f54159e62ed7ce86d63a3104a78f3191c44abc18e1e4897965a180) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-970607dd56e9eb8b1822d5c1967c3f4022244aa9c572ff34ce09e06a9365432c) |
| `ingress_egress_gw.node_number` | [ingress_egress_gw.node_number](resources--gcp_vpc_site--reference--group-002.md#canonical-9a4d3e975620c56fbfdf085d117ee6ac3ac0f551a9eae2e1c9b6e1d7eba8d9a6) |
| `ingress_egress_gw.outside_network` | [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-e3d9fa2dcaa5ba7f8b5427ae8e29854aa4ac2c600b9212abd686b5fd5f232469) |
| `ingress_egress_gw.outside_network.existing_network` | [ingress_egress_gw.outside_network.existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-21de07053b9ffb592d5074f7ac1aea74522765b8848723a38141afd5de67a6a2) |
| `ingress_egress_gw.outside_network.existing_network.name` | [ingress_egress_gw.outside_network.existing_network.name](resources--gcp_vpc_site--reference--group-002.md#canonical-22e6afbbc6bcd329d9bdead2f84010e02dcb60de5473667c516a5aa4e8d39f60) |
| `ingress_egress_gw.outside_network.new_network` | [ingress_egress_gw.outside_network.new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-a16eb8cef119f498fa43d09bb0dd4e8bcd06c353b5fe91aab8580abc3c38c358) |
| `ingress_egress_gw.outside_network.new_network.name` | [ingress_egress_gw.outside_network.new_network.name](resources--gcp_vpc_site--reference--group-002.md#canonical-cd60a18b997ca50b4d503e1ec53e27809c02b2f4634bfe72462f90a163dbecc8) |
| `ingress_egress_gw.outside_network.new_network_autogenerate` | [ingress_egress_gw.outside_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-7456193481ed477d7e767a9ab6a4d8e376cc46a97a6b4503109682ed8e3c7d31) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-5aeb2c18af6df8c821efc8419b22d8f17025852aab1a8a6a917ebb9288f88401) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-b92a885e9021fda1f44f828361d0f61d0c347e5b87f6a5502e614341ebafda4c) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-003.md#canonical-eab63ac4254ea7eb934be257b7d5ac26b247582615580be17033ef45cb52f4b4) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](resources--gcp_vpc_site--reference--group-003.md#canonical-ab75c7c1a7d6608f7fb9bb100d07e6d2f8d65b4fc039ec45eca79b1597bd5d69) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-003.md#canonical-2abd43c8e800b5027639a6bb346cd530122e2839283cd87a221491bf7c8c9da7) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-779c38e03356d852b932c31b1a19ae21008fe9bc77e1f3b38ebfae7b5979dc48) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-003.md#canonical-322abc8ae06ec90139186b47c29fe533a3bffbe99669659d64500d3c8e6dc127) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--gcp_vpc_site--reference--group-003.md#canonical-a6aeddd84d734ea71fbf47899f3b8c3df6a939ac6a636196f3b29b22a7995c34) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--gcp_vpc_site--reference--group-003.md#canonical-b72a67218def7c45a5d200a69196fcb77b7f803999aa8ff09cdb4bd863bea6ee) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--gcp_vpc_site--reference--group-003.md#canonical-204efbed3d6e01a4248b5ce93f1d2d0419bb861adc079dbccd5f5ae3bbf6bf88) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--gcp_vpc_site--reference--group-003.md#canonical-dc4ece3dbc08c41efdb9e501bf3a38264f562507d21a16d73bd6a130216200ec) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--gcp_vpc_site--reference--group-003.md#canonical-7dd6ef8dd787053e0ab8d1ce0a07330dbcf190d1558fd159e0ae07b3355e0052) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-83642e9293fff071211cceb256c429ec50c4abd6eb3010034b2807db0395c81f) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-003.md#canonical-9ec8a134aa837255aaf4a145150149804403c7448cd05dfda8810501599aa288) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-c479744450b567fd5c91903d2f517d0621525a0e085684d99d35270d4a85dc5d) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--gcp_vpc_site--reference--group-003.md#canonical-658f5c09f1e8d67ed4ebaef71a5c16fa439584479e620b6a9afd334f7d944826) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-5862a73c34b43d4eb455da8f2abe438510d3e2e101e0ff135041773a999128fa) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--gcp_vpc_site--reference--group-003.md#canonical-44d229f8691027ca87401be1cd519430a93612afbb928bc6242cb6c20c41b037) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-72d955f9e802b86ff2ae156ed8cb0a4394400917090e75ccbdb50f71f2eafa7f) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--gcp_vpc_site--reference--group-003.md#canonical-82b4bbf9d3958076a4358924fbacad3623599ff13259abe4204700096e6dfb3a) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-c7770aed39817279a03b90f7c1a9596890520bf6d5c1b45b92257c3908380326) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--gcp_vpc_site--reference--group-003.md#canonical-524882a7e347d31b0124fc614f7d14de39bf0b11d504a7697e4dbe268908668a) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--gcp_vpc_site--reference--group-003.md#canonical-649da02f03f021229999887b010dfc174914cf560f61c66b20d8870fa0599a44) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-003.md#canonical-918b6371b0b9b19b3ffd0bb5dadded0b8616c58583c4668fd2e7b6a988530248) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-93fce2ebee00c7160a7196cacacf97c1e9e42b8edd5146ba93e4db393d84f13b) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--gcp_vpc_site--reference--group-003.md#canonical-4c9a6e415e686ced8015543f571b459eb7a880a1b4acc2c1344ad80d7e55be16) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--gcp_vpc_site--reference--group-003.md#canonical-9fb5831e0f3fd694089576e93d1251f900518a71932cce1a9137a68fe67052d5) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-25d295d0eed640b6faa4d4cbbea4cc2fecef6dccded0bd34f3f62ca35135a2fc) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--gcp_vpc_site--reference--group-003.md#canonical-44f36f8939e4408cd178af5fc72369115574a1913072b23040b241aa8c18e920) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--gcp_vpc_site--reference--group-003.md#canonical-264cad398fdc63af1c4e0faeb749eca461e885df71b779f548aebbd05856698a) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-51c19c7a33174067acaef98b42f4eb5d04882a4824e200a94ea3d8ba0098cd1a) |
| `ingress_egress_gw.outside_subnet` | [ingress_egress_gw.outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-9a578c42bc083926bf10f0050692291d5c62860fec9846fbf277927c21b19e59) |
| `ingress_egress_gw.outside_subnet.existing_subnet` | [ingress_egress_gw.outside_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-59bdf79f3cfc8595b169625a3b943b553db84927a66c9f83f75b0e2fc824ac1c) |
| `ingress_egress_gw.outside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.existing_subnet.subnet_name](resources--gcp_vpc_site--reference--group-003.md#canonical-2200db1a4075d94eaf41befbe970b4d417d81c333b26a2538c9c53d1e789ae04) |
| `ingress_egress_gw.outside_subnet.new_subnet` | [ingress_egress_gw.outside_subnet.new_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-6da8d9343fa11eb8f4cd94d58e72097d3fdffe322d5352d71a8b50a97d2b5f3b) |
| `ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-50d00fb1f42454eacd39bbc2e70e6138f2a5eb9b96ebdf36bf2f216daa6f5251) |
| `ingress_egress_gw.outside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.new_subnet.subnet_name](resources--gcp_vpc_site--reference--group-003.md#canonical-dc2b5bfeec8505491435c68a14f4b739a137b3e91c5b857801c0165440349e94) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-49631701caf3f2a0c8e89a3939443712e3409f13a73a281c090000807839aa4a) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-631d4d4fd8afd586b15f074d4aff7ba0342f8377c142c398382ffcbdcb8b4499) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-bc1dcd4e3e5f9762c85874eec773dadfd1b3431fae1cfd29147d34f0dbc0706d) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-b17d5838a74705ff59944c2619fade8824e9fe64811f6e18d73dcb4a37e70804) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-c07a44cd28d6892dec19751848e1bf6d1064f8265c42b850679614a621f70f6b) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-701601133f3569641999edcc0b2b71bfa05ce66f45e50538fb5134ecfdaac584) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--gcp_vpc_site--reference--group-003.md#canonical-29c6b8e11ee5b7f724e32b04e51c8ad49ad6a1bdfadb89db1341494e157ef81d) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-eca9ec7d3159b6ffbd992edde8871a72d19854b02fbfa7d36032163a616f2431) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-64b31924136a5c0b320dfabd7580984c726f0f425c4fb0a115dce902f00a889f) |
| `ingress_gw` | [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-999174b385cdf8f7e4313fe3c8dc8ca1154d6456d47917b6df7b8c1ff499375a) |
| `ingress_gw.gcp_certified_hw` | [ingress_gw.gcp_certified_hw](resources--gcp_vpc_site--reference--group-003.md#canonical-ea631fe01a07f252748b3866ee9bbe405730eec17b005b0a240ba1003acdf922) |
| `ingress_gw.gcp_zone_names` | [ingress_gw.gcp_zone_names](resources--gcp_vpc_site--reference--group-003.md#canonical-6d431fc7b5a8a19add0de709ae2280ef78e7875f8beb409f24b35659537e8b9e) |
| `ingress_gw.local_network` | [ingress_gw.local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-565824f57008a1595cbf974e29bb99faba4c94308f7062b13fa4f9580267ca63) |
| `ingress_gw.local_network.existing_network` | [ingress_gw.local_network.existing_network](resources--gcp_vpc_site--reference--group-003.md#canonical-56a212bfa3e4738c17a733f08b29b275cd216829ec5f3a578be7ce340cf879eb) |
| `ingress_gw.local_network.existing_network.name` | [ingress_gw.local_network.existing_network.name](resources--gcp_vpc_site--reference--group-003.md#canonical-ee22643c3de58de0d3b9d4c108876329c9fe161e3fad85dd084bd2ff5cd3afc4) |
| `ingress_gw.local_network.new_network` | [ingress_gw.local_network.new_network](resources--gcp_vpc_site--reference--group-003.md#canonical-8ac46e259b17cc2c4f02bc880a9b57915c5736e8b84f77c923f7e5b4f89ed424) |
| `ingress_gw.local_network.new_network.name` | [ingress_gw.local_network.new_network.name](resources--gcp_vpc_site--reference--group-003.md#canonical-73dd238691969eb15736ea18e18acd3e9c778d2763b386b24f8c01cc2e5ad9b7) |
| `ingress_gw.local_network.new_network_autogenerate` | [ingress_gw.local_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-003.md#canonical-dc9190c876743380015c38c9e72c06d7c8982193cd1d7c31acf29adf368a2b5f) |
| `ingress_gw.local_subnet` | [ingress_gw.local_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-f58fc041d1f2e4976f7ec01d550eca7fb957bdb25c2e9c9cf2aca087106ec6bd) |
| `ingress_gw.local_subnet.existing_subnet` | [ingress_gw.local_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-4e0ccf60e1809038f4a0e2811e88dcd6ecb3a94d1e8d9ed517d633120fbb479f) |
| `ingress_gw.local_subnet.existing_subnet.subnet_name` | [ingress_gw.local_subnet.existing_subnet.subnet_name](resources--gcp_vpc_site--reference--group-003.md#canonical-124e27b442c38d03c220c4f4c4b873ca20419c31d841c48dee8a7df21c5062b0) |
| `ingress_gw.local_subnet.new_subnet` | [ingress_gw.local_subnet.new_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-837e5a47a9a4677c1f0334683acc38f5f649d5fceec231743dc881ce6611d27b) |
| `ingress_gw.local_subnet.new_subnet.primary_ipv4` | [ingress_gw.local_subnet.new_subnet.primary_ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-ca61ee10fae123ba71c93bbd02d3fb9b82bea456e5ca33e80c30af6104d63bc2) |
| `ingress_gw.local_subnet.new_subnet.subnet_name` | [ingress_gw.local_subnet.new_subnet.subnet_name](resources--gcp_vpc_site--reference--group-003.md#canonical-2939541af982605d8207f82f19a48e7a40a61b105ca9dc1f9b5e855967209fd8) |
| `ingress_gw.node_number` | [ingress_gw.node_number](resources--gcp_vpc_site--reference--group-003.md#canonical-abf2353de80c961bf3868e19b484fae2acdc5f99f22c55876ba2e42564e6772f) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-9df052fe637a6042cce70ccea5ce671ab706ca318d367398ea06bbca947323bd) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-34dd71794ad6565e1e02f055db37d6d1f90964cf2f5f53d859cadbcc71691aaa) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-e062985941c06d8a089bd40059cdd03d8192640f484b813bf0ed18bde53af79e) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-62aa8d0f40e4e868580d32b4a343a97ecb48cb5caabdbffbaaf533f104936d42) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-93722bfea2fc8324e764935d6b69cbdb989605236b55d4995abf4de1fd880ece) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-eec2dd07ca556a8626e84ded35aecdf7d34599b1deb2554f6bb276a2a5675149) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--gcp_vpc_site--reference--group-003.md#canonical-62a31d96180c6c711fa1e85bd5ee7d4c904b7cfc2af7e8e7d96c3b6532270f88) |
| `instance_type` | [instance_type](resources--gcp_vpc_site--reference--group-001.md#canonical-5a934812b9ee77174ec0cdbd462bb759464afea4c2d773e32d1f14bdb12c87c8) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-a1c706be5022d2e31e11e4567d36e5b95beada9218fb2e60e10dc03c54150083) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-38422e31f184402be45a057d8b662544bead9089f37ea79391f371dd32c2dcda) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-541dbdb32270271694c2e6a6d17a195867b124117f93e49b1f9f923de0ceeceb) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-b3dead804495e953126f6405332e8d5e29d2f1f30f661d28f47a1c3ba6f82dc4) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--gcp_vpc_site--reference--group-003.md#canonical-296a7e4cb6785e58cd2982fdee98574a226ca7a06395ebe06c4373d4e44b7edf) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--gcp_vpc_site--reference--group-003.md#canonical-71b23bd989f8931227a0110a2c702a358faca08a6c5a8ce18ea89be67af60920) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--gcp_vpc_site--reference--group-003.md#canonical-2a68c3ec4e2c2d7878c17826f3e7fdb6a3c016e174c61eea8c8bc41a3af0a176) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-c782b3f3c8d9c555bf52dae72211443f85ba398a5f0de682df207602638a4d27) |
| `labels` | [labels](resources--gcp_vpc_site--reference--group-001.md#canonical-91485afe3b26851d2a6d3295435f77fe34fb3bbda400660749f104709326a074) |
| `log_receiver` | [log_receiver](resources--gcp_vpc_site--reference--group-003.md#canonical-a26d857017679cfda085e0f0a49903fcc032db4921c1de00e087766ff383fd49) |
| `log_receiver.name` | [log_receiver.name](resources--gcp_vpc_site--reference--group-003.md#canonical-f2f5d2a070cfa515203de8cb8ee80cf3abca174166b5a2e8ef9173d2b7f1fe5f) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--gcp_vpc_site--reference--group-003.md#canonical-7e41f6b87b4b787e7dc7bd9ca47622c05c4bd48bf3d60078bf5387aead3c5c8c) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--gcp_vpc_site--reference--group-003.md#canonical-c48869e4cb461325b044f277716159ea150be8d6a41e82aeead6f4f56e481987) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-f8b5d095a44382632a1481ab7b7f448fa5e9a46803bfd87fe1e8db304f2c49c1) |
| `name` | [name](resources--gcp_vpc_site--reference--group-001.md#canonical-3aae36927c98842ce6a50195b151aeae1cf823867dc7bac52cb9c0eafb0a7c3f) |
| `namespace` | [namespace](resources--gcp_vpc_site--reference--group-001.md#canonical-82a06d4ea4f1fe4df494c165cd1cb8ffb2659b0da939e44cf715c9a17a46ec5a) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-22243c16b6b4cce94c82f3fbe7d4e9c68740dc568b60d633ec7c7c8e7befe4d4) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-229b89e58d72d2433bb0c5ac8709d6c31d165994c49a623f7f2b6bf0fee71835) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-499d05184531855d870c44e4d7435c92ad1107f7d67d61eb82147726134b1198) |
| `os` | [os](resources--gcp_vpc_site--reference--group-003.md#canonical-88ee877961aa11a6145caff7451f332227f90693cced62a81bcf0b7223977d08) |
| `os.default_os_version` | [os.default_os_version](resources--gcp_vpc_site--reference--group-003.md#canonical-cdaf9a488d747065c0e2be29bf8d3ef53a4f115bc2b19f7d5f74a8441e686501) |
| `os.operating_system_version` | [os.operating_system_version](resources--gcp_vpc_site--reference--group-003.md#canonical-95ae728237cbc1ff0eb7a22667f056dfcdd8ad1483c9a8148491db89ab0644c1) |
| `private_connect_disabled` | [private_connect_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-afc6f45f786e41688d334f4267b4908fd8fc55ce3574f716dd799807571516cb) |
| `private_connectivity` | [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-06c998e2599f7224ad0891dff32b8d5e831d315c925369a3bc90d2a059c30193) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](resources--gcp_vpc_site--reference--group-003.md#canonical-2400ba220f3ac6c28824d7698e46f773874d7129ed4379c3a45c2a254a2b3fce) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](resources--gcp_vpc_site--reference--group-003.md#canonical-fbb9c6a13582b76ad43013fdce2dd23c210dbe977bb08b3e52b95ba6b2ec5ba2) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](resources--gcp_vpc_site--reference--group-003.md#canonical-4a37d85be0d4adb91882efa12d87df43af2d88b446e06c20fa5755c739d2f2ec) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](resources--gcp_vpc_site--reference--group-003.md#canonical-24e7992c2003a746297915b8543f782d15e80f095b0a6aed4f21161079305b1f) |
| `private_connectivity.inside` | [private_connectivity.inside](resources--gcp_vpc_site--reference--group-003.md#canonical-cb234dcf5f479a2e02f80201badd9438dc9b773dc90dc0f45500a7aa9bd25d0a) |
| `private_connectivity.outside` | [private_connectivity.outside](resources--gcp_vpc_site--reference--group-003.md#canonical-b172837f4ecae9888073aff680024aca5ea0c9414ce319ac916e9684264b5db2) |
| `ssh_key` | [ssh_key](resources--gcp_vpc_site--reference--group-001.md#canonical-a9fe566bd0f7c324583d1724835d557ccbfaf68ce093dd4399d2fad10e6dfdc6) |
| `sw` | [sw](resources--gcp_vpc_site--reference--group-003.md#canonical-ac299786eb3acecb781355b7cb3843424b0b2f0aed576816a040d543e47dbb69) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--gcp_vpc_site--reference--group-004.md#canonical-4acb0e49762c1b77f784b71c216bf0747e5d5bf71c9fb49e802f14a4b79a5ae2) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--gcp_vpc_site--reference--group-003.md#canonical-db56b7c2c6159bb17fc629e032371a6860314a8be27262c65fa4a01eed0ab25a) |
| `timeouts` | [timeouts](resources--gcp_vpc_site--reference--group-004.md#canonical-8eae28ef2cad76a6257738eb65a3ced02f012a7344774ac5d6165414e4c0739e) |
| `timeouts.create` | [timeouts.create](resources--gcp_vpc_site--reference--group-004.md#canonical-38844ebfad91166c0a71af7cac77e0d5189b23093dfca2397684076ffb99decf) |
| `timeouts.delete` | [timeouts.delete](resources--gcp_vpc_site--reference--group-004.md#canonical-c85e6158c1df5fd539adb6f62f1e6ca9279420478642c2580b68abf4beb6380b) |
| `timeouts.read` | [timeouts.read](resources--gcp_vpc_site--reference--group-004.md#canonical-36f4b47c9f3598d4fc1327d5795881fa5eec2867966e59f5c023049b21fa9cf4) |
| `timeouts.update` | [timeouts.update](resources--gcp_vpc_site--reference--group-004.md#canonical-1f9c8e19946da3aa6cd54f2ebe2adbd18bf95e97a4f86716025f54e4be5f2066) |
| `voltstack_cluster` | [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-760f6ba6855f451d0477a5feddf0087ff8d9892bce89b49960bd8724d3ce6eba) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-3ba32b782e67df3af8ceede722bd2d1b7bf8edd4a8ee221fd71b218666222fff) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-1a67cca6960532ec5d4ab6d48bfe4e6f84b1f4b8b73099ef459c693a20ebf3b7) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--gcp_vpc_site--reference--group-004.md#canonical-cebb3702efb5f1df66128a68cd662390c387ebd37f7ddd6b6ce461294785bed0) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-5922b67a7d2422b90b1fc35f191fc6777793aa1beb4d934f3cdc26d89c6c5383) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-2a953ea73d73fef189a305d13d645d0c16984f98cbb1100d09e13232b4dba11b) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-35170f83f20c897ab9ec9641b9d6859a2bfc13681ffbe5368a23bf03a607adbd) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-671bf46083feb34481edd240361521c0a8f083d9b3122bac6b213c70a877e3e1) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](resources--gcp_vpc_site--reference--group-004.md#canonical-174e8ac715b4ffdc7cab62329df4226baae2ecb73c1781519868edef9cd49c2d) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-c90eb76da83b43d7df895a68588d5b16e330344cda15fc95b1233bba7d67687f) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-7e2525226c04fae220224868ec0f36dd59c8ca49494ba2e2b5fbda193cec0505) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-8a739ec992939e6c8cd2604af0fbc3ad2158dff1d2015788af793cce8418fdc9) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](resources--gcp_vpc_site--reference--group-004.md#canonical-99640778439a9a333ff988ab515ec375f576d96d03c220596a26ef11e1882c16) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](resources--gcp_vpc_site--reference--group-004.md#canonical-61b2be690f775792f4571c7d5a35917cbedd15cbc5eee850cace31f4b694192d) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-05a8c9ef0c2c963dd798de691c74a7d513a59f60ba740175fa8bc7cf71ba2af2) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-7783eb4daf25d95e089ab6e238146410e4b314c1aec8cd4fac079a29a068022d) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-45928c72e712d901c584aaec9317c761f2f2a2c40956d4abb305e748d2163a4c) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](resources--gcp_vpc_site--reference--group-004.md#canonical-7f91495e7cbcc5464b4388edc0361872933a1fb5d354d14b1d6dd0114bea51fc) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-5178170977e9133e4edcb137f2f1d560c56a8a395b7dafc0efa032f3e92c8140) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-7325fb7aa15c1f2beb8fcafc01348ccc9c6bc37e4082e2268a4ccd242ed661a3) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](resources--gcp_vpc_site--reference--group-004.md#canonical-e4934706ddb00479eb420cd5e501e6e5077f0ef36c2fc301c5a2e752f3d01d17) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-004.md#canonical-b270bff65d2b871fd504af6f6ad2e59a8f63ba12959bd14b00a0643baa5f8ef2) |
| `voltstack_cluster.gcp_certified_hw` | [voltstack_cluster.gcp_certified_hw](resources--gcp_vpc_site--reference--group-004.md#canonical-6c38174d6557a61d5381f449c0bf0c73daf0c4b2908b862cd3b087c068c3da20) |
| `voltstack_cluster.gcp_zone_names` | [voltstack_cluster.gcp_zone_names](resources--gcp_vpc_site--reference--group-004.md#canonical-8e2e4d67b9eff80eeb9025f1b49b8b9f0b000de3d5839f950e0e8b5a5ac88e41) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](resources--gcp_vpc_site--reference--group-004.md#canonical-250995e2a50a5aedccc225c32904a6ff53936fb9bbe7d81d29f7437d0b45d614) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-004.md#canonical-5ddafdefdd62543430af79cd3df7b29849e7c7dcd3bea0d56632a643ea288353) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-291b72062cfee93fe587afb128866ef7fa87c0487de63dc04006707b4a34112b) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-c408f83aa1811ce927b3ac0faa3fb06509244675145ca969cd7dd3d9a853b9b6) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--gcp_vpc_site--reference--group-004.md#canonical-02d612711a27b6e76e946f7c59fbb048f2837082575b8bdc7efecaa6165a696b) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-c145ee1ea8f3c3c83c70c04b893bd7950f0fdb1c03d39d29abbe949f5de999a3) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-5632a2fde54683908df229083bb7f9a3ec627c2c966cf90dfa1a12571a90d859) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-004.md#canonical-d004df18e13d13da8115ee8d893608080f099aa9d2fe372edfb60c4ce2fb4f64) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-004.md#canonical-73486476f74bc188c2d336f365128aa66e03b1fae6c9407f2437fd720014a38d) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--gcp_vpc_site--reference--group-004.md#canonical-6d2e19994c60263a0d89d4a3cecc655565d87997364eded12809b53774bad99c) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-77ae2eb54889df98777dce829cf26ca19d8ca402c045a418b8589b7a11a4cea2) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-6bb17bcaade8d00203487d451d355300735a8c329ccf5e82ce0d84240ade731c) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-2bf47800db46259ad52ac94e9319c52820f937344c07c766342ce6d457f23a96) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](resources--gcp_vpc_site--reference--group-004.md#canonical-dfc9c74019b1f6802985371eea12d70001e7b599de5cd561b74ae24d5b9f16db) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-29915eee296f345040d6fd82d764be1d4e31904ff3ef15e2d92442f48ee94118) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-52d9563136f8de026a87d7b4fb883381605bdf4d15300f2fb947e19c8dbe321b) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](resources--gcp_vpc_site--reference--group-004.md#canonical-e8878600451a685834f2706b24e0c8f7e5b5e1f3a59028efc4e8836a39516456) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](resources--gcp_vpc_site--reference--group-004.md#canonical-981570e351655185177fac094af43bc06253137065e723867cd157d1f5425f0a) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](resources--gcp_vpc_site--reference--group-004.md#canonical-0dc62f3fc25df0223740f2aedf805c5672a0e8cc65b659b0ee2836c87e6a51b7) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-5a986bc0ab426a8dc93c4c643f6dc94236c5e62ae14730948ef59a749068d752) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](resources--gcp_vpc_site--reference--group-004.md#canonical-b899c4074c916ef91648ffe0310cb2432c6a4da1a9010971b0dfa32f91adb25c) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-1375c16ad0b028bc75ae2ce39634a682cf1f7309a0861bb22261785b6e3ffdc6) |
| `voltstack_cluster.node_number` | [voltstack_cluster.node_number](resources--gcp_vpc_site--reference--group-004.md#canonical-0c4eac7ef1c52788f352810d377ccbc461e7c242b0894f80d41868c08042e311) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--reference--group-004.md#canonical-d5bcf5569367f54a95054ce45e587f464533e941209cd060f1542c6a096056f1) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-004.md#canonical-c27e825aae2c35d9fafeaca8adbf4f0a53e6f055124b080ca5bf3f049cd2dccb) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-59d2cb72733d728f978a2e335317dfa7d36dc31008ad9e3adf2179b8c9f7b3b5) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](resources--gcp_vpc_site--reference--group-004.md#canonical-79083fc6040b8ffa85dbefb5da7a2a41e66219c47976fded4ca59a2f7a154179) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-004.md#canonical-d155129d55816cf875cc180bc9831941887f82187031a433eef90c3aca9d7b54) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-004.md#canonical-aeb6f82ed0d3c87b4376b0f877aaa23fa3123205eb298dd9c5408a2987088adf) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-004.md#canonical-3244e851722af24c4619a34896da1794ed4c7ec1dc55a6995e051777d0be8115) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--gcp_vpc_site--reference--group-004.md#canonical-ea4418e1dd97338ab16dadcd61bd067a92aff19d09933eff2c820c7502a13044) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--gcp_vpc_site--reference--group-004.md#canonical-c49afbadc73d70db63c8eb7165c1853fbbc1234eedb4f23807a9031e6f1b16b6) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--gcp_vpc_site--reference--group-004.md#canonical-14d3055b1ec7516694e8d7b6289ba11d7c1d6c4be52346f659c58294ff91dd5c) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--gcp_vpc_site--reference--group-004.md#canonical-09967afd5dde12f56f7e1ef7dccf9e240d632980f0bccf8a758b6b84ebf76672) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--gcp_vpc_site--reference--group-004.md#canonical-c8ad32bdb63c2b4edad6e99fc1f0f2c307e12544e8f7dcebdde651a8b55a2a4f) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-004.md#canonical-59d654a088d6925f97de3e48fab7f5db8a9b1cf42fd156c2724419bb4c84a6f3) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-004.md#canonical-2a24b26f4b58f020024562bd07e823c492b4e0b297f7a9650af842df3b85bf7e) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-cf60c6ee02bc0def55a94f3199b7f8c22279ebd4d479f22a019fd1847a7f9a14) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--gcp_vpc_site--reference--group-004.md#canonical-a0aa7d924e0e1cb413cafdb01f6016aa2b7dcfd861c91cdeae2fcb00ca422386) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-297dc78fc661acb899c262a2e5714d58fe6702ff4486c59962b2851a2d46f5dc) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--gcp_vpc_site--reference--group-004.md#canonical-484dcf9533a112387fa504ae7498e2b5b25aab68a775518bf4e84965d96bb25e) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-117e9fc0a8dcd4b43032a76674e51f12d14415a63655e224f1c1fe3386c559bb) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--gcp_vpc_site--reference--group-004.md#canonical-acb0a7e9d7f2cc6e48e3ab0ec331a2263b87cb782f5cf5dff0ab6796c1cb8ab4) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-1a1f1c3234e7ecffbb08cf9dc95947219dca7764d8894c84fa496383d966f60c) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--gcp_vpc_site--reference--group-004.md#canonical-4fe406a9d724af87e6290b954ebb2f5bcdb9d13afcf1bac6c82256fcd9c7778b) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--gcp_vpc_site--reference--group-004.md#canonical-335941e8353a5d3e85755115d465cd04ecfdab59b44cc5eac5a753f20c92ac94) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-004.md#canonical-3740612018e269ed2f526e7dc4db014f976a0d2bc73138f8dab123daee31ea21) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-6d909efbb462fb3a5c3146ff06577fccd99a10e19468a4b5c76bdcb83863ae12) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--gcp_vpc_site--reference--group-004.md#canonical-9e458e64ee558b09911d62280858a0120aae428aee35a82257c72b5438963cb6) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--gcp_vpc_site--reference--group-004.md#canonical-994b89208b6d1db05bf0a2762aadc4f5dd6485b3e00507bb78e66946a8ebb7e3) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-004.md#canonical-5b1fb5985180339ad98209e369bb3878a057b46bcf668dd2b225b73c8c8d6be8) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--gcp_vpc_site--reference--group-004.md#canonical-d897c305925daa3610f9fb162a7f94dc412d1817aa48b564b39ceaf63f0f3248) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--gcp_vpc_site--reference--group-004.md#canonical-d5ade6dca7076a1afd37f7e9f6f0e7b49e26d39ebf016260c279a4b237a84d23) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](resources--gcp_vpc_site--reference--group-004.md#canonical-3c49603c4c2c5a9034e1097061e22e7b4e590a0589d129996331e9bf58864bb1) |
| `voltstack_cluster.site_local_network` | [voltstack_cluster.site_local_network](resources--gcp_vpc_site--reference--group-004.md#canonical-f658f139f77e0428765d3100c68fde04cffb01647313dac0f40e07a1241e3886) |
| `voltstack_cluster.site_local_network.existing_network` | [voltstack_cluster.site_local_network.existing_network](resources--gcp_vpc_site--reference--group-004.md#canonical-7305db8ef2592e9ff57af0018db8418a678b6b908cd9191c20d663903d4a3202) |
| `voltstack_cluster.site_local_network.existing_network.name` | [voltstack_cluster.site_local_network.existing_network.name](resources--gcp_vpc_site--reference--group-004.md#canonical-c4c9d13b1bd0ad5af58d1a5bf6e0a8835dde3bc01e1ef67b0234981ad7235716) |
| `voltstack_cluster.site_local_network.new_network` | [voltstack_cluster.site_local_network.new_network](resources--gcp_vpc_site--reference--group-004.md#canonical-23585c10aea578ab42f24fa3bb7624c1a9e61287606e1a56e1c8c1fd34f12e15) |
| `voltstack_cluster.site_local_network.new_network.name` | [voltstack_cluster.site_local_network.new_network.name](resources--gcp_vpc_site--reference--group-004.md#canonical-3dcaa675337d0ae63c322f91e21f872abc21b1d20eef0e78666bcb7ccc76d95b) |
| `voltstack_cluster.site_local_network.new_network_autogenerate` | [voltstack_cluster.site_local_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-004.md#canonical-03c01835d37a0e51f6c25972048c9fd54478b62f81073f74f1c15ebd3d237bb0) |
| `voltstack_cluster.site_local_subnet` | [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-56bf4aadfb600193fd1c6b1684f3e5587e30da2af67d6cd36085ff749220967d) |
| `voltstack_cluster.site_local_subnet.existing_subnet` | [voltstack_cluster.site_local_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-049f89ec6068a6ca074106fd52159456f1732fe10b8d877293a7144745f1f76c) |
| `voltstack_cluster.site_local_subnet.existing_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.existing_subnet.subnet_name](resources--gcp_vpc_site--reference--group-004.md#canonical-949c0e46fa330d562bb170c3dbbbcd842edb8a6f4ee75fed653b30d9433d9c1b) |
| `voltstack_cluster.site_local_subnet.new_subnet` | [voltstack_cluster.site_local_subnet.new_subnet](resources--gcp_vpc_site--reference--group-004.md#canonical-bf8d4e45e95e8717eb0d02218c7161a2e671c5e86155d18a96ca0d3cacdd9080) |
| `voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4` | [voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4](resources--gcp_vpc_site--reference--group-004.md#canonical-19dd63bb4c1963a3e54a4cf48ffbb4d5c88d71961ab209b076060cd8e609ed42) |
| `voltstack_cluster.site_local_subnet.new_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.new_subnet.subnet_name](resources--gcp_vpc_site--reference--group-004.md#canonical-4afd6009eb78ac4ca741fe9a8e6540833ba741e7144417827ab98326b12b06c6) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](resources--gcp_vpc_site--reference--group-004.md#canonical-7c2d345d7ba13b8cfbc44093a4f28e5f94367fc74436044ae23d5132e3d4e52b) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-004.md#canonical-01b61b5f410628eff55ddb1d6934093db1b36bc1d23d4b9e3d157e8a65175b42) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](resources--gcp_vpc_site--reference--group-004.md#canonical-8658a534e6cd77f480f217f65542f6488987b44f41724f719b6941b6079b6c39) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](resources--gcp_vpc_site--reference--group-004.md#canonical-6fd7b37cd5297ba229262cf9a4c1a96d0c2fd2bfdd617882da061f2d58604583) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](resources--gcp_vpc_site--reference--group-005.md#canonical-47b3bab9615e8f8ec0219bed7f94047b50ff486639d6da99db405638adf669c3) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](resources--gcp_vpc_site--reference--group-005.md#canonical-8d677aa2dbbc1bb6ebce673026aec17c437bc869e25f8aa2cae326c8e7a573e4) |
| `waf_signatures` | [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-20cdebd9914ff01019312f079981c5ff72c43afeb869ae02a434b180495dfa66) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--gcp_vpc_site--reference--group-005.md#canonical-794d027c6997fac963ad2bf51273c0df9cf5650cdd3ce9388ffaf76f034f8be5) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--gcp_vpc_site--reference--group-005.md#canonical-f1e533a933e3037a95c549327d025ca0fbedfe75c25f86c7bacc27d06c104713) |

<a id="canonical-ba52aaed57d8ec7e04996d2864b4d3498b8fe393f5837ba8862d617bc37e8c99"></a>

## Next pages — Property reference / ecb0e823acfb / 18

- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-18add28bee441d3d531062ac60749c6c76c95afcc4bfbf9d55cd6c1e538098b3)
- [block_all_services](resources--gcp_vpc_site--reference--group-001.md#canonical-530fa4fdc2c05b7ed2ac80e6e596f07e8e6046f437b406923387bbfbd1fb533d)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-909c249948254fbec30f63c678012c95380a2edad9b7709aa1548da60ead104e)
- [cloud_credentials](resources--gcp_vpc_site--reference--group-001.md#canonical-7605207075b1d259a748aabdbd6f06777b66701612652233add52b297a2687fd)
- [coordinates](resources--gcp_vpc_site--reference--group-001.md#canonical-1ee923a929358c616aab84a0ce7c58d4c6808078b75cfa0eaa85cc4699011ac8)
- [custom_dns](resources--gcp_vpc_site--reference--group-001.md#canonical-e8c0b808c7aa5333108e5f700c5fdbdbfd782aad00c4616db4c163488eefee4c)
- [default_blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-f437cf619b611a651738572b03d9cd7173644469efc9d4bcb509457d475b6c9f)
- [disable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-5d57af9ed321100aa53f9baf8cbe7bec1da9c0cf28e205bab76d17351fe60eb3)
- [enable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-51a5b7b4447038370ee22e5d23fefdc78599620aa70e55d74e5d89319ed1f8f2)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-91992f0d045bc68652deeaf62a5b3849c65f5447dcdd89b44c6c2b97307f6efb)
- [log_receiver](resources--gcp_vpc_site--reference--group-003.md#canonical-4b1d87b55107bf727296e0f74b17d3ab715fa9ab4ce733b2a51208d76bc40e16)
- [logs_streaming_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-c144b8deb7de45a002c3f118644029a062245be86a46e5968311c10907617f46)
- [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-321583ca89a2994ed6ab1089491a0a823abc73ba1e92d28770051589feecf20d)
- [os](resources--gcp_vpc_site--reference--group-003.md#canonical-eaae4c5a3696af2cdaa426ef1745cc75f8cc67b832eed73e1b84ffbca525c28b)
- [private_connect_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-60d557d0083e9ae6c90a25feb164999e2e99fcf4ad1174a90fa437468a9395e1)
- [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-2b156b46727ac0a204195e7e776cb3951197b186926382e9a9299718d824f79f)
- [sw](resources--gcp_vpc_site--reference--group-003.md#canonical-267f967c717532fd9b7a7d530db900ad6f3ed5285284fb0cc81876b19418dec2)
- [timeouts](resources--gcp_vpc_site--reference--group-004.md#canonical-2298b57a32aa2303d946d8a0a735921119bf066b4e9cbcf62828b9ca2db0aad9)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-0ca74f7b2129ec6b55d795ed58941e7bb46495b5356da884837bb06cdb4c4657)
- [waf_signatures](resources--gcp_vpc_site--reference--group-005.md#canonical-a8bce3447b430d8e2a86211cdc8cabaa307dce6f5e9a69b93a189d355e794ad0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-18add28bee441d3d531062ac60749c6c76c95afcc4bfbf9d55cd6c1e538098b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e40c766eded359662fb6ec6ad7b82dca6894d6b487eedee87cd6a4289b22034"></a>

## admin_password — admin_password / 301a40269c4d / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- admin_password

<a id="canonical-c3a0b8541b1aa8db5bb99430dc51ea6903dc63b3eea554655814046a5a2d60dc"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-ba3bf2b948e795c34ec5ea12a9e39edcd45716b19f735c4fbfe78feb7e097e60"></a>

## Direct properties — admin_password / 301a40269c4d / 3

- [blindfold_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-95aa93108919da9b3935dfba86e632cbbd4190249b628510370b617ef8ddd390): complete subsection reference.

- [clear_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-63825864fe745e3669971eba568aef98b41fcedafbe0a53c11ba1a4428c18110): complete subsection reference.

<a id="canonical-b4e2c1ef58a47475c3d57c71de6096aa59d95e80b0f1a791fee6292cde9e2e91"></a>

## Next pages — admin_password / 301a40269c4d / 4

- [admin_password.blindfold_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-95aa93108919da9b3935dfba86e632cbbd4190249b628510370b617ef8ddd390)
- [admin_password.clear_secret_info](resources--gcp_vpc_site--reference--group-001.md#canonical-63825864fe745e3669971eba568aef98b41fcedafbe0a53c11ba1a4428c18110)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-95aa93108919da9b3935dfba86e632cbbd4190249b628510370b617ef8ddd390"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f81c59153e24282964b4ef2c0616b99ce269b0ab294d3a60235f83cd040306c8"></a>

## admin_password.blindfold_secret_info — admin_password.blindfold_secret_info / 71e4d7fd80e4 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-18add28bee441d3d531062ac60749c6c76c95afcc4bfbf9d55cd6c1e538098b3)
- admin_password.blindfold_secret_info

<a id="canonical-7e2f54e2ac4eb35999bfef936b7875d1f563c28540651dc6a52ca78f733aa704"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-9dcf34688af13fc9d7e80df9b75b042cbf18c3e5fa42298c0647c35cc698515a"></a>

## Direct properties — admin_password.blindfold_secret_info / 71e4d7fd80e4 / 3

<a id="canonical-762eaa2b46723d0c87de1f8dc31efd0cac50297ee10f5c1582069270557a7950"></a>

<a id="canonical-d42637ff80d4a28b1a14c3b07117ef7da1c4529fd300d1914ce6dd169a1ab6d3"></a>

## decryption_provider property — admin_password.blindfold_secret_info / 71e4d7fd80e4 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-36f5497bb8cf9ba7188412cac2a1f0cc9ea5f702915192756e936c7a232c1abf"></a>

<a id="canonical-7d6c08b0d34ed661c4e9f02f30dee59cd82b7b5a1b132d93b4395401a9fd3397"></a>

## location property — admin_password.blindfold_secret_info / 71e4d7fd80e4 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-ef35964fe7bd2eb8a378894a9f210d79fdf6762f44560957ec03c318106cd386"></a>

<a id="canonical-5fd0290cc83a1cfff60a3d74ddd1d866494e75dce7ea4248230482caeb61b372"></a>

## store_provider property — admin_password.blindfold_secret_info / 71e4d7fd80e4 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-9a9eb905e5587845f285656c0b855dbb17cd66521094fd5fd4c4a5f22dac270f"></a>

## Next pages — admin_password.blindfold_secret_info / 71e4d7fd80e4 / 7

- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-18add28bee441d3d531062ac60749c6c76c95afcc4bfbf9d55cd6c1e538098b3)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-63825864fe745e3669971eba568aef98b41fcedafbe0a53c11ba1a4428c18110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce83079afafc90e074fcd5f2c61d99743af31487e206dd4823e078e3f4abf5d4"></a>

## admin_password.clear_secret_info — admin_password.clear_secret_info / 0d442f82ec76 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-18add28bee441d3d531062ac60749c6c76c95afcc4bfbf9d55cd6c1e538098b3)
- admin_password.clear_secret_info

<a id="canonical-0b1bf8d3e212b5ce76fcda970c8d85116a5ea3729d8beceaf547df8c4b1b9c52"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-fc2b7576cf4428387067d17879e4dd19468b8352430e4e605a5d9e952ceb06ed"></a>

## Direct properties — admin_password.clear_secret_info / 0d442f82ec76 / 3

<a id="canonical-d7e2b6825e8126064a998032cc29c6dd91b739d493f89dfcebd569bee153cb9f"></a>

<a id="canonical-6e6c33f45aedb5fa096518065635001995a67c045258aa2212755610b3ff7a86"></a>

## provider_ref property — admin_password.clear_secret_info / 0d442f82ec76 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3b801374994c8596a10e0399020b4381e90d6b38d3da065c72d4a7ca246c1202"></a>

<a id="canonical-d1c2aedb7e974865c6739bd66c6bbc6a4ca6ba6322a0153f3d7b2b49c44a56e6"></a>

## url property — admin_password.clear_secret_info / 0d442f82ec76 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-7bf949e207e1c54905c05f8e526c6412cd4ea78ab898ddd0815bd6cec3a7d6ef"></a>

## Next pages — admin_password.clear_secret_info / 0d442f82ec76 / 6

- [admin_password](resources--gcp_vpc_site--reference--group-001.md#canonical-18add28bee441d3d531062ac60749c6c76c95afcc4bfbf9d55cd6c1e538098b3)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-530fa4fdc2c05b7ed2ac80e6e596f07e8e6046f437b406923387bbfbd1fb533d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69819df916aa1f9763e5160c2cb9c5f5bbd4ec8a1f5a56bb15bedc8156e80837"></a>

## block_all_services — block_all_services / 15c70dfb19e3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- block_all_services

<a id="canonical-048fca1479a480db37fe1cc5f8db572fff45c1e0152aed86724d150cab9fe316"></a>

Type: `["object", {}]`. Optional.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

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

OneOf alternatives in this subsection:

- [block_all_services](resources--gcp_vpc_site--reference--group-001.md#canonical-048fca1479a480db37fe1cc5f8db572fff45c1e0152aed86724d150cab9fe316)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-40ed0b75b1e1bbb6826fe1612f67629b453b86e353fdb8d3c3ace972f7b99a7e)
- [default_blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-e9e2a082dee8738d6dc8de57f82a85fee2c74d73677af996801695fc7d9bfadb)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
block_all_services = {}
```

<a id="canonical-ca6c4b6935c49912fc089a3393f878887a1c0a886f6892c0c04db5f707e4bb67"></a>

## Direct properties — block_all_services / 15c70dfb19e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9dc271f9b88aabec8ff871b399cae3bb22f3caee2454ee53aace5ae21e7111d3"></a>

## Next pages — block_all_services / 15c70dfb19e3 / 4

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-909c249948254fbec30f63c678012c95380a2edad9b7709aa1548da60ead104e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dd6fd7eeb7f28df47ad9e2e86bd1bcb291b197d009ec093e96e2fa0f1d9c332"></a>

## blocked_services — blocked_services / 8fbb3c36c122 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- blocked_services

<a id="canonical-40ed0b75b1e1bbb6826fe1612f67629b453b86e353fdb8d3c3ace972f7b99a7e"></a>

Type: `"object"`. single nested block, Optional.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-d30f05d39176ea27bb61eefde46f5f628908c8b117319ec9eaa752227fa3b7de"></a>

## Direct properties — blocked_services / 8fbb3c36c122 / 3

- [blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-ca96ae77473a24932e3c816326cf0411534245ef25f16524259a02dc683f09a0): complete subsection reference.

<a id="canonical-a27532942e61e41c347b3f4ac74656e1fb5fcdc020bb3239fe31dc93f5126e4a"></a>

## Next pages — blocked_services / 8fbb3c36c122 / 4

- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-ca96ae77473a24932e3c816326cf0411534245ef25f16524259a02dc683f09a0)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-ca96ae77473a24932e3c816326cf0411534245ef25f16524259a02dc683f09a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea0b573ffdfaa0c2f5450e7cbb06489794c3e357f7db0b18124024671724c056"></a>

## blocked_services.blocked_service — blocked_services.blocked_service / f834de0bb646 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-909c249948254fbec30f63c678012c95380a2edad9b7709aa1548da60ead104e)
- blocked_services.blocked_service

<a id="canonical-ce5d377d88dc557f4bec763db0193206571bc27c1ee52fa37019a3aadbf20a44"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

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
blocked_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-f1db8b50f9dfee670eaa92f723ee6767dc65a28af9b03de7a4deed2d89896e1a"></a>

## Direct properties — blocked_services.blocked_service / f834de0bb646 / 3

- [dns](resources--gcp_vpc_site--reference--group-001.md#canonical-678233f817d796110ad424b13e247b430ac25a2aec4eee5e6da2a2b72ee5c647): complete subsection reference.

<a id="canonical-3d4b763eb7a3f722ecdf5e7506faee64b9b5c066b1e31d0eea39a882ad80d448"></a>

<a id="canonical-5641fcebe955370e84d09b4e43d0a53db401861042cd784c629622d47638dc7e"></a>

## network_type property — blocked_services.blocked_service / f834de0bb646 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ssh](resources--gcp_vpc_site--reference--group-001.md#canonical-0be7a17bd07dc7449c7b47c6384c98def256c53aef15c411337dddb5cf5a7ae4): complete subsection reference.

- [web_user_interface](resources--gcp_vpc_site--reference--group-001.md#canonical-c4e60c6769ee547abf3bf589f10ff0ba673a3909e4c62927da3eef6388d98aac): complete subsection reference.

<a id="canonical-b759312a36215cbe48bbbeed608cfaf6db3201e316b7f0864ea4e15c986af2ca"></a>

## Next pages — blocked_services.blocked_service / f834de0bb646 / 5

- [blocked_services.blocked_service.dns](resources--gcp_vpc_site--reference--group-001.md#canonical-678233f817d796110ad424b13e247b430ac25a2aec4eee5e6da2a2b72ee5c647)
- [blocked_services.blocked_service.ssh](resources--gcp_vpc_site--reference--group-001.md#canonical-0be7a17bd07dc7449c7b47c6384c98def256c53aef15c411337dddb5cf5a7ae4)
- [blocked_services.blocked_service.web_user_interface](resources--gcp_vpc_site--reference--group-001.md#canonical-c4e60c6769ee547abf3bf589f10ff0ba673a3909e4c62927da3eef6388d98aac)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-909c249948254fbec30f63c678012c95380a2edad9b7709aa1548da60ead104e)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-678233f817d796110ad424b13e247b430ac25a2aec4eee5e6da2a2b72ee5c647"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7f84b757804f635c0531dc87fe6922b730ee3c7042f1272adaef7f9cbb1d784"></a>

## blocked_services.blocked_service.dns — blocked_services.blocked_service.dns / fde2b2142cc0 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-909c249948254fbec30f63c678012c95380a2edad9b7709aa1548da60ead104e)
- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-ca96ae77473a24932e3c816326cf0411534245ef25f16524259a02dc683f09a0)
- blocked_services.blocked_service.dns

<a id="canonical-56b33d280d462be0217555498c29f2402b1ef74df98a3a94c23feb2d5935a5fc"></a>

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
dns = {}
```

<a id="canonical-3fa72449d32a95dd803af354ebb6bc1920d9f9e93a4f96395b04d0cc83074094"></a>

## Direct properties — blocked_services.blocked_service.dns / fde2b2142cc0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-caef5bccf8c0c388565d9b8d3efed099ba8f784609e20ee50f55cb2a047cb7fc"></a>

## Next pages — blocked_services.blocked_service.dns / fde2b2142cc0 / 4

- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-ca96ae77473a24932e3c816326cf0411534245ef25f16524259a02dc683f09a0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-0be7a17bd07dc7449c7b47c6384c98def256c53aef15c411337dddb5cf5a7ae4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dca41ed3f6db57a6e3c149258d953d389f1c7f8b09d12600a68df15071e9047"></a>

## blocked_services.blocked_service.ssh — blocked_services.blocked_service.ssh / 36ee10c34431 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-909c249948254fbec30f63c678012c95380a2edad9b7709aa1548da60ead104e)
- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-ca96ae77473a24932e3c816326cf0411534245ef25f16524259a02dc683f09a0)
- blocked_services.blocked_service.ssh

<a id="canonical-d37f2e179d76d69b044a65c18e8ea484b2316e60869d774588f2579067f3e1c2"></a>

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
ssh = {}
```

<a id="canonical-0a40216b36c0291493b92123ab11ab769dbf762bea2152380df572c71d1bac49"></a>

## Direct properties — blocked_services.blocked_service.ssh / 36ee10c34431 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b419071602dab0426dd1aea057d614b9af1af6ded3c9c4896a49e99f6bad438"></a>

## Next pages — blocked_services.blocked_service.ssh / 36ee10c34431 / 4

- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-ca96ae77473a24932e3c816326cf0411534245ef25f16524259a02dc683f09a0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-c4e60c6769ee547abf3bf589f10ff0ba673a3909e4c62927da3eef6388d98aac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a56461f5b55a5e5eaab2a08f8906049fc6485171f67b0f1574a344fed250f241"></a>

## blocked_services.blocked_service.web_user_interface — blocked_services.blocked_service.web_user_interface / 12d43727ec3a / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [blocked_services](resources--gcp_vpc_site--reference--group-001.md#canonical-909c249948254fbec30f63c678012c95380a2edad9b7709aa1548da60ead104e)
- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-ca96ae77473a24932e3c816326cf0411534245ef25f16524259a02dc683f09a0)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-26a76fd918e6c2dc193e36927db59e1a8a0638dc85fef8bffe0928da2ceda8f8"></a>

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
web_user_interface = {}
```

<a id="canonical-aebfe39c2b002c976881be7e19a0e72885334e70cb56d34e2057674074e4ceb9"></a>

## Direct properties — blocked_services.blocked_service.web_user_interface / 12d43727ec3a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5b61f8cd6568f82e44ebc965621b740fb284c28576c676a885687f964f79fbb"></a>

## Next pages — blocked_services.blocked_service.web_user_interface / 12d43727ec3a / 4

- [blocked_services.blocked_service](resources--gcp_vpc_site--reference--group-001.md#canonical-ca96ae77473a24932e3c816326cf0411534245ef25f16524259a02dc683f09a0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-7605207075b1d259a748aabdbd6f06777b66701612652233add52b297a2687fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-453d8eff5655410e1b716536c966104d31fd294e6956d5880a2a44bbae623ed8"></a>

## cloud_credentials — cloud_credentials / 7cdf121dd6d9 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- cloud_credentials

<a id="canonical-a6b3e47ed3196bff633e8bee11ac1478a41dd25b83acc9291fccf74921e259ca"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
cloud_credentials {
  # Configure direct properties listed below.
}
```

<a id="canonical-18e317cbed5b18a703e754033fadca718f48cf0c1b638f856ba07a079423848a"></a>

## Direct properties — cloud_credentials / 7cdf121dd6d9 / 3

<a id="canonical-24cbd6943b0107a6703dca7e690e79b90fbc1bbf41819bff7b15fd2d1593f3d5"></a>

<a id="canonical-04f1c8ee3af1db6d46b55a66bb07578b15281c6ec8776cfed0aeac0635219658"></a>

## name property — cloud_credentials / 7cdf121dd6d9 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-69c3fc7a78772052c34e5415e55add94d3471f682a6bc3fd4fbcad86e3ae7445"></a>

<a id="canonical-7b328973d9efbe4269fe3ef1355204f36720243560892953a83973cf6b6ddc0d"></a>

## namespace property — cloud_credentials / 7cdf121dd6d9 / 5

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
}
```

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

<a id="canonical-d199fe0d1a8ddb8d651d8a0f4fb65f9939523c3a636722f76c7f1c4dd32a0f5b"></a>

<a id="canonical-b31be7ea6d4694a9df3efef0421723436eebd4b7ff086990b865571c169d6a94"></a>

## tenant property — cloud_credentials / 7cdf121dd6d9 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0c313d33a9a44b22f6dd1274d39153e3f24f74f457fb863415fef2ff49bc1511"></a>

## Next pages — cloud_credentials / 7cdf121dd6d9 / 7

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-1ee923a929358c616aab84a0ce7c58d4c6808078b75cfa0eaa85cc4699011ac8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f3485f9eb36d3953737750aa5888eae2dd32e7dfe6b23c8ec5e9714b8e768ac"></a>

## coordinates — coordinates / d5c469716205 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- coordinates

<a id="canonical-45eaca03320e66bdc6ad8eefdd9730b334211eb74b2aba12ece8005ea3242199"></a>

Type: `"object"`. single nested block, Optional.

Coordinates of the site which provides the site physical location.

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
coordinates {
  # Configure direct properties listed below.
}
```

<a id="canonical-36ee947833bcefc435cda19d5fb3152ae979b1972054709044a42778479d8b36"></a>

## Direct properties — coordinates / d5c469716205 / 3

<a id="canonical-04839decdd9441917a0cac6115838829f6dfdd9a74692c4240a71490e9c80190"></a>

<a id="canonical-e67b75b00f75227ada80d6a2a33d0dbb9f568a39e1f53f5a5d92e420106e7a55"></a>

## latitude property — coordinates / d5c469716205 / 4

Type: `"number"`. Optional.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-dbe18856ca8f74ee685853014deeaa4eb1f26acae4402b766fcde0d50eaeeeb2"></a>

<a id="canonical-a191f72850c88d8f668aa575ac1105a05cb65b2d1e74dea6312c931c521b2c24"></a>

## longitude property — coordinates / d5c469716205 / 5

Type: `"number"`. Optional.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-e5822e02d771fcdaef029a20eff6f2d3b104d01da973f950ac92f4bb1fca92e0"></a>

## Next pages — coordinates / d5c469716205 / 6

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-e8c0b808c7aa5333108e5f700c5fdbdbfd782aad00c4616db4c163488eefee4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3de586dca4510b7e270277b4964b2765fc47e443687022e185cba29aeebe4067"></a>

## custom_dns — custom_dns / c922c4abd597 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- custom_dns

<a id="canonical-c232eb908f219d5d9f81c0b765f3509f04213b581145dca6c1897f789a3c366f"></a>

Type: `"object"`. single nested block, Optional.

Custom DNS is the configured for specify CE site.

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
custom_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-8f75dec7de3cfd37ef91e28fa0e71984beb95c29a630e3dbf35704830afef0dc"></a>

## Direct properties — custom_dns / c922c4abd597 / 3

<a id="canonical-ae00ebe6bb587207a157c8672b17a7e3b0cef1dadd3c99965b61b39cf9e7d1df"></a>

<a id="canonical-e4dd1b68c4dedac695ed1c253a9ba95f9170e89089af2ab9cefcadb4017e4a8b"></a>

## inside_nameserver property — custom_dns / c922c4abd597 / 4

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in inside network.

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

<a id="canonical-ae304ffc069929f5298383fe6840a878cff521b819bc52105fe7fbd421e4904c"></a>

<a id="canonical-28001ef6607978bee46aad4bceece9866451606e3015a1c466e33bca53a5a07e"></a>

## outside_nameserver property — custom_dns / c922c4abd597 / 5

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in outside network.

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

<a id="canonical-8e2cdcebfe25a2138212b9c67e922c4859e44d5840ff42395537bb280e600bf4"></a>

## Next pages — custom_dns / c922c4abd597 / 6

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-f437cf619b611a651738572b03d9cd7173644469efc9d4bcb509457d475b6c9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7aafb2e0c85332c21d6f8ae44a4422d455584cad6ea9a5fab16be6ab7cdaef79"></a>

## default_blocked_services — default_blocked_services / ce7e63703207 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- default_blocked_services

<a id="canonical-e9e2a082dee8738d6dc8de57f82a85fee2c74d73677af996801695fc7d9bfadb"></a>

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
default_blocked_services = {}
```

<a id="canonical-69a1995d5ebab42b662e23ac782f62c84e94e245b95db671c0130898168e75e4"></a>

## Direct properties — default_blocked_services / ce7e63703207 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-55d9649ad66925ea00a10de70e8128396cd055f5d954c7284c3b579ccee95ca7"></a>

## Next pages — default_blocked_services / ce7e63703207 / 4

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5d57af9ed321100aa53f9baf8cbe7bec1da9c0cf28e205bab76d17351fe60eb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e446546ff65efdc7cd91bab8e0ddb83b5592298f6d266c4c021a512b4f3da8e"></a>

## disable_encryption — disable_encryption / 3ba9a0a16d1f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- disable_encryption

<a id="canonical-6580d72f7bfe1f713ab6946a5892531818101ede667a92673b4c220773607b25"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_encryption, enable\_encryption; Default: disable\_encryption\] Configuration
parameter for disable encryption.

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

OneOf alternatives in this subsection:

- [disable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-6580d72f7bfe1f713ab6946a5892531818101ede667a92673b4c220773607b25)
- [enable_encryption](resources--gcp_vpc_site--reference--group-001.md#canonical-614da5ada84d27e8085ce867a8308c6719470ae34ab110f6c2c37a5fb12652b9)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_encryption = {}
```

<a id="canonical-705bbe56ad0ba7d3d08c02fa0f87107463151fe3bb389c403a9ce67f3c448391"></a>

## Direct properties — disable_encryption / 3ba9a0a16d1f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-424dcdc7cfb239c6788ec669287462aac057301ddc3d4e3265e7bd384cf66dd9"></a>

## Next pages — disable_encryption / 3ba9a0a16d1f / 4

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-51a5b7b4447038370ee22e5d23fefdc78599620aa70e55d74e5d89319ed1f8f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74d8107dad61c950e083e07a9c988ce2855a56014e922ace521b88aa84b787f7"></a>

## enable_encryption — enable_encryption / a0f63b9f158f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- enable_encryption

<a id="canonical-614da5ada84d27e8085ce867a8308c6719470ae34ab110f6c2c37a5fb12652b9"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("kms_key_resource_id",
    "kms_key_ring_id")}
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
enable_encryption {
  # Configure direct properties listed below.
}
```

<a id="canonical-ac2c93be7364bb8a4e6c0d2736a27dbc3171384d17474f8b3eee65576be73445"></a>

## Direct properties — enable_encryption / a0f63b9f158f / 3

<a id="canonical-7d7d9d265e0f61dc2cab4dd3996bba72c07bd4bada29d0858f046314d6232894"></a>

<a id="canonical-eef42c2c937467b09c9a56336e522f7fd4c5a0ca20cfe7a5cefa401e91290abb"></a>

## kms_key_resource_id property — enable_encryption / a0f63b9f158f / 4

Type: `"string"`. Optional.

GCP KMS Key to be used to encrypt the disk attached to the VM.

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

<a id="canonical-bfa21b384bdf37cc0da2cc65d55433da149e0ef1f3882d6e7780beaf8b962700"></a>

<a id="canonical-f52995cb2503147db20e9b226cccd0cd1adad46aa3e50a3ce25f6b2e7813bacc"></a>

## kms_key_ring_id property — enable_encryption / a0f63b9f158f / 5

Type: `"string"`. Optional.

Key ring in which the CMK to be used to encrypt is present.

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

<a id="canonical-2c450465eb2ca3e6e643288daa164a532f5f1d34370cf0bf8c18622b8acdbd79"></a>

## Next pages — enable_encryption / a0f63b9f158f / 6

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bbd43d46310b02d22eb8a81388d4d1507934674feaa56025ed40a7ac231c600"></a>

## ingress_egress_gw — ingress_egress_gw / 4e4946f67bcb / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- ingress_egress_gw

<a id="canonical-c6ba3e3f7c8cfd37d26ea3e57fe7ab1b7a2f174266e35c4e1914080cc714210a"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface GCP ingress/egress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("gcp_certified_hw",
    "gcp_zone_names"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("inside_static_routes",
    "no_inside_static_routes"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-c6ba3e3f7c8cfd37d26ea3e57fe7ab1b7a2f174266e35c4e1914080cc714210a)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-999174b385cdf8f7e4313fe3c8dc8ca1154d6456d47917b6df7b8c1ff499375a)
- [voltstack_cluster](resources--gcp_vpc_site--reference--group-004.md#canonical-760f6ba6855f451d0477a5feddf0087ff8d9892bce89b49960bd8724d3ce6eba)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ingress_egress_gw {
  # Configure direct properties listed below.
}
```

<a id="canonical-8ad36cb23ccabdb5ef88c8012e38fe99422c6ffb345ff7a062a1b5fe9fbe5828"></a>

## Direct properties — ingress_egress_gw / 4e4946f67bcb / 3

- [active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-7105b339fc23b53004862f942e57d3faabc1403632864c87f7a2d774b0def427): complete subsection reference.

- [active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-b5b4c33b332c2f3597d8f4e3bccfd55bad4b343082aa5b6cda3e4b4c53735cfb): complete subsection reference.

- [active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-261ff2a541c862c8da4a249ad3655fa60954efc402186110687a8752816c111e): complete subsection reference.

- [dc_cluster_group_inside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-250c9d377018d89fb39f874d1fdf70099ca624d2e166fa0b513f2cd7d81f3bb0): complete subsection reference.

- [dc_cluster_group_outside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-9379418bfa11e879e5b850a97ce6b3ba8abd41fdd9ad9ad2531c1ba333f0d44a): complete subsection reference.

- [forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-002.md#canonical-85c804a1ee2f8c0833998b30c7199fdd3a3abd09c91542ce4a85b65b7d6c8a3f): complete subsection reference.

<a id="canonical-de3cb092e4504a727f8d792728f1606e25da90f624db6c0a7edcd8f09ddb2980"></a>

<a id="canonical-a855b772a1f957fa7620ae9d3232d7200f99c4714b53548bcf1643f8aaf8e2ac"></a>

## gcp_certified_hw property — ingress_egress_gw / 4e4946f67bcb / 4

Type: `"string"`. Optional.

\[Enum: gcp-byol-multi-nic-voltmesh\] GCP Certified Hardware. Name for GCP certified hardware. The
only possible value is \`gcp-byol-multi-nic-voltmesh\`.

Upstream description:

Name for GCP certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("gcp-byol-multi-nic-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-multi-nic-voltmesh"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1fcaafaaa27777ee094588a8a662656202b3ad1f29684c517837970e46454fb1"></a>
