---
page_title: "xcsh_nginx_service_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery reference."
---

# xcsh_nginx_service_discovery reference

<a id="canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2652646df7c62eddc07e999d4b5655a217df07b00e08c172cd535ab1539d2ab"></a>

## Property reference — Property reference / 2a4a88477386 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- Property reference

<a id="canonical-9bcc7ec5ec45029f7c273a1760ba9b49596dbe399135eac385615d77ffe43903"></a>

## Direct properties — Property reference / 2a4a88477386 / 3

<a id="canonical-49deaea7dca3c5fbb6cc96bace279a85490c4a08ee54a4c4664038da2fe5ec44"></a>

<a id="canonical-3ff386646dcb8425e1f4653c3721f045e31c19dd0429e99027f128d4d52ab427"></a>

## annotations property — Property reference / 2a4a88477386 / 4

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

<a id="canonical-c829eb8451f454d7e409b099f307c3a87866e8d598ae77a99e8110cb779ff836"></a>

<a id="canonical-704df5c2e7aea134a56029246f1dea908325af25049541343ef5e476eceae7bd"></a>

## description property — Property reference / 2a4a88477386 / 5

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

<a id="canonical-af0942b6373b44b83f99bf7c2d465a29a13c6a17861c362b12da73cf30027730"></a>

<a id="canonical-9dc4ff4a7d7fc48240c25669c1897a616de0d120bbc383eac624d160612d5eef"></a>

## disable property — Property reference / 2a4a88477386 / 6

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

- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-2054a037c9d64d979b7e5977309895a409f9af4abcca00bd24c7e70e5e31ebf6): complete subsection reference.

<a id="canonical-f2b6b90a9fdaf0829d60cbf623a74b95ec1215963f68caa0991c866a0cf0673a"></a>

<a id="canonical-777bbd7bfcf61469ef7638e3dc9afa92903ead5379eacce146697277d6de4cee"></a>

## id property — Property reference / 2a4a88477386 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-4b2ad1edb35313cbe746277829a4c4bd6d41a359ab49d671736f6034e85067a1"></a>

<a id="canonical-1cc53fb8ba86ce386425af3233afeee187c8d431fba9a77b57bc6ae057144095"></a>

## labels property — Property reference / 2a4a88477386 / 8

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

<a id="canonical-590b08169160e866a44d942f650f632961bd5bbcaa0bd3544033c34783bc318b"></a>

<a id="canonical-6ecf415c0625e8b53cb239e4cad101f5e9110f7770e6ce54cd582e51309b91e9"></a>

## name property — Property reference / 2a4a88477386 / 9

Type: `"string"`. Required.

Name of the Nginx Service Discovery. Must be unique within the namespace.

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

<a id="canonical-8244e39d021882f536a87a9425c86547a26fd31a6d805757c665c0a42fa81e31"></a>

<a id="canonical-dea01eaecb5f97c6a9a65f570b3131d82acc9ddf7d6a11135d2fbcdf8b008bbb"></a>

## namespace property — Property reference / 2a4a88477386 / 10

Type: `"string"`. Required.

Namespace where the Nginx Service Discovery is created.

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

- [server_block_filters](resources--nginx_service_discovery--reference--group-001.md#canonical-aa29d75af6e1655c8211234f59f20760fb3899fff39a441e11340c60ac7f4f9b): complete subsection reference.

- [timeouts](resources--nginx_service_discovery--reference--group-001.md#canonical-fe720eb086b57545b2fb0f63f5461fc0b63c51e4b8e942105bd7f5b7cfd791f1): complete subsection reference.

<a id="canonical-e5d77c435819a5b580547577ad408509d5ccf2d770e2ecef6534ca599e178528"></a>

## All schema paths — Property reference / 2a4a88477386 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--nginx_service_discovery--reference--group-001.md#canonical-49deaea7dca3c5fbb6cc96bace279a85490c4a08ee54a4c4664038da2fe5ec44) |
| `description` | [description](resources--nginx_service_discovery--reference--group-001.md#canonical-c829eb8451f454d7e409b099f307c3a87866e8d598ae77a99e8110cb779ff836) |
| `disable` | [disable](resources--nginx_service_discovery--reference--group-001.md#canonical-af0942b6373b44b83f99bf7c2d465a29a13c6a17861c362b12da73cf30027730) |
| `discovery_target` | [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-8af76100b856e03baa6cedc650239dded2f83e5a172d99c153af3d3b409de9a7) |
| `discovery_target.config_sync_group` | [discovery_target.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-005dab6cfd7088d4f47befb529abe950df908ef39ad0a85375a0e231b7a72ece) |
| `discovery_target.config_sync_group.config_sync_group` | [discovery_target.config_sync_group.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-16f422cfa2a46f5f8ce21fafe59b8ab271532b0658e5d68917f2e2ce25d6d9fd) |
| `discovery_target.config_sync_group.config_sync_group.kind` | [discovery_target.config_sync_group.config_sync_group.kind](resources--nginx_service_discovery--reference--group-001.md#canonical-27130147c44350b40fa7d045d4ab27a9caa45ad36932bff6389b91b7a6e1c67c) |
| `discovery_target.config_sync_group.config_sync_group.name` | [discovery_target.config_sync_group.config_sync_group.name](resources--nginx_service_discovery--reference--group-001.md#canonical-2798f643301c60f000f879546ed91c2ccba821ffe57d820ea0a08dea4d2a9879) |
| `discovery_target.config_sync_group.config_sync_group.namespace` | [discovery_target.config_sync_group.config_sync_group.namespace](resources--nginx_service_discovery--reference--group-001.md#canonical-7d70a9fdd08c376fbbbdcb46c3bed10f765774aadcd7fdb023ce6af0fbe351dc) |
| `discovery_target.config_sync_group.config_sync_group.tenant` | [discovery_target.config_sync_group.config_sync_group.tenant](resources--nginx_service_discovery--reference--group-001.md#canonical-20598e24525bc4a6f7e0c992c6018ed29a508d61226abdd9612fd5deed4dcc61) |
| `discovery_target.config_sync_group.config_sync_group.uid` | [discovery_target.config_sync_group.config_sync_group.uid](resources--nginx_service_discovery--reference--group-001.md#canonical-6252daaa85b3d69b6e075a443e96556e9b07fe0bbdc82d6f588673a0db94c0fa) |
| `discovery_target.nginx_instance` | [discovery_target.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-78505e910ac7bdd8d5fb2e27ed54d006019cd2ec718a1fe780188857d47c9b0b) |
| `discovery_target.nginx_instance.nginx_instance` | [discovery_target.nginx_instance.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-e8099d466372c0d9dbba4102956518ba5f1a1a8f9fefbc589fd98db5d2de6e9e) |
| `discovery_target.nginx_instance.nginx_instance.kind` | [discovery_target.nginx_instance.nginx_instance.kind](resources--nginx_service_discovery--reference--group-001.md#canonical-b5be6b837223b0f83c865035e126b672e7ecb3f9e4e3dbdfd87bbb34d292a5a5) |
| `discovery_target.nginx_instance.nginx_instance.name` | [discovery_target.nginx_instance.nginx_instance.name](resources--nginx_service_discovery--reference--group-001.md#canonical-9d5bffaff73aee7f1d2ec2cd9a1c2026e7b6e085b0d6dd77b769e7e798d6b1d2) |
| `discovery_target.nginx_instance.nginx_instance.namespace` | [discovery_target.nginx_instance.nginx_instance.namespace](resources--nginx_service_discovery--reference--group-001.md#canonical-dc741cef4a95896994f1ba5f1f68142ab5c3ae15173bf65d921792d985a26b78) |
| `discovery_target.nginx_instance.nginx_instance.tenant` | [discovery_target.nginx_instance.nginx_instance.tenant](resources--nginx_service_discovery--reference--group-001.md#canonical-70191eb5c643d9c3dc666b3d9df5c1e8c4564574c587d551e89ccf81cd2895c1) |
| `discovery_target.nginx_instance.nginx_instance.uid` | [discovery_target.nginx_instance.nginx_instance.uid](resources--nginx_service_discovery--reference--group-001.md#canonical-4e70255fbd81eba06a2e117d8f177ab16cb200f5e714fc9edce08eaaf7ab68a7) |
| `id` | [id](resources--nginx_service_discovery--reference--group-001.md#canonical-f2b6b90a9fdaf0829d60cbf623a74b95ec1215963f68caa0991c866a0cf0673a) |
| `labels` | [labels](resources--nginx_service_discovery--reference--group-001.md#canonical-4b2ad1edb35313cbe746277829a4c4bd6d41a359ab49d671736f6034e85067a1) |
| `name` | [name](resources--nginx_service_discovery--reference--group-001.md#canonical-590b08169160e866a44d942f650f632961bd5bbcaa0bd3544033c34783bc318b) |
| `namespace` | [namespace](resources--nginx_service_discovery--reference--group-001.md#canonical-8244e39d021882f536a87a9425c86547a26fd31a6d805757c665c0a42fa81e31) |
| `server_block_filters` | [server_block_filters](resources--nginx_service_discovery--reference--group-001.md#canonical-c86a6a9bf1ea1c8cd3622faaff016bf6befc32c7d0e812592a034a244d6d4c8a) |
| `server_block_filters.name_regex` | [server_block_filters.name_regex](resources--nginx_service_discovery--reference--group-001.md#canonical-971b7ba000cd9913a4b38ab088e2f0dfb60e876735cda14887a63080c0fc810d) |
| `server_block_filters.port_ranges` | [server_block_filters.port_ranges](resources--nginx_service_discovery--reference--group-001.md#canonical-7b634dae1709d55836179ea8a146c8ddac9f3cb05ceae9ef50b7d65a7299bf61) |
| `timeouts` | [timeouts](resources--nginx_service_discovery--reference--group-001.md#canonical-d0fca5828481edc50c29f4d915ab395980f080db92b39ff51f8b959d7ced822d) |
| `timeouts.create` | [timeouts.create](resources--nginx_service_discovery--reference--group-001.md#canonical-bc3b9923a5115c97da21b102dde71575da63a437f9e725dea1eae052e2514e08) |
| `timeouts.delete` | [timeouts.delete](resources--nginx_service_discovery--reference--group-001.md#canonical-40a61f66ac4f17b16204ab3867932b2cfc1d93dbe13f8a1a5067f30dee515ee7) |
| `timeouts.read` | [timeouts.read](resources--nginx_service_discovery--reference--group-001.md#canonical-78157e1c1a76a1d8e033d3091f312e756f8fe9e4f00e6302901af0d62e1dc120) |
| `timeouts.update` | [timeouts.update](resources--nginx_service_discovery--reference--group-001.md#canonical-eb609d4d71d73f24017ef9ce1881bdf681475ccbcb61a407791b99572d32fcc1) |

<a id="canonical-1371b6c7f1fe4227c79185b390932f1f1755278f39d1089e12cf3b9e7b3822e1"></a>

## Next pages — Property reference / 2a4a88477386 / 12

- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-2054a037c9d64d979b7e5977309895a409f9af4abcca00bd24c7e70e5e31ebf6)
- [server_block_filters](resources--nginx_service_discovery--reference--group-001.md#canonical-aa29d75af6e1655c8211234f59f20760fb3899fff39a441e11340c60ac7f4f9b)
- [timeouts](resources--nginx_service_discovery--reference--group-001.md#canonical-fe720eb086b57545b2fb0f63f5461fc0b63c51e4b8e942105bd7f5b7cfd791f1)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)

<a id="canonical-2054a037c9d64d979b7e5977309895a409f9af4abcca00bd24c7e70e5e31ebf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78cc0181d0ef2e7d57f4b60aeb5188c359cd8bcdefa7efb8ee8da0792805b336"></a>

## discovery_target — discovery_target / 4b1d7280c07c / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- discovery_target

<a id="canonical-8af76100b856e03baa6cedc650239dded2f83e5a172d99c153af3d3b409de9a7"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for discovery target.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("config_sync_group",
    "nginx_instance")}
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
  "x-ves-oneof-field-target": "[\"config_sync_group\",\"nginx_instance\"]"
}
```

Terraform syntax:

```terraform
discovery_target {
  # Configure direct properties listed below.
}
```

<a id="canonical-059dcdb4f9cb76271479a5ea9295d6bd03504f516755acc3996b7a2dd76c14e5"></a>

## Direct properties — discovery_target / 4b1d7280c07c / 3

- [config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-8598c4fb654452cc64a846b13a58e10f57ba05580c1e9d7cdb693b0e6987ca1a): complete subsection reference.

- [nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-c24f0ca3ef08690c9d09aca923162a6dd0553a3d68f23de93a2a6dbe2cc8809f): complete subsection reference.

<a id="canonical-484f0212a0401645d9de1c7186a7bcf5db93effbc26e2ecafd9d0eaab7dad7cb"></a>

## Next pages — discovery_target / 4b1d7280c07c / 4

- [discovery_target.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-8598c4fb654452cc64a846b13a58e10f57ba05580c1e9d7cdb693b0e6987ca1a)
- [discovery_target.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-c24f0ca3ef08690c9d09aca923162a6dd0553a3d68f23de93a2a6dbe2cc8809f)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)

<a id="canonical-8598c4fb654452cc64a846b13a58e10f57ba05580c1e9d7cdb693b0e6987ca1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79cabb87a1dc94c145be29b7e14c1a106cdba3cda23c872d4ed763a64e37ec67"></a>

## discovery_target.config_sync_group — discovery_target.config_sync_group / 7100f4010c93 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-2054a037c9d64d979b7e5977309895a409f9af4abcca00bd24c7e70e5e31ebf6)
- discovery_target.config_sync_group

<a id="canonical-005dab6cfd7088d4f47befb529abe950df908ef39ad0a85375a0e231b7a72ece"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for config sync group.

Upstream description:

Select new ConfigSyncGroup.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("config_sync_group")}
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
config_sync_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-b21facb32dd46cde213feeb64ff4c6545014d7032a44e5ba529fd4144e3eaa98"></a>

## Direct properties — discovery_target.config_sync_group / 7100f4010c93 / 3

- [config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-1eb09d1ad3f93d00963cbe48fb0d3681090123af24c22b9cb648c00dac95eccc): complete subsection reference.

<a id="canonical-be7ad85298017eb0d1f4fb5b35b5d1ed1754345230412b161a3ab4cf572a68e0"></a>

## Next pages — discovery_target.config_sync_group / 7100f4010c93 / 4

- [discovery_target.config_sync_group.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-1eb09d1ad3f93d00963cbe48fb0d3681090123af24c22b9cb648c00dac95eccc)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-2054a037c9d64d979b7e5977309895a409f9af4abcca00bd24c7e70e5e31ebf6)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)

<a id="canonical-1eb09d1ad3f93d00963cbe48fb0d3681090123af24c22b9cb648c00dac95eccc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8384f36501737c82a54edd48b4c4ae90ee63c3dc34ac683775c14fec49ce3533"></a>

## discovery_target.config_sync_group.config_sync_group — discovery_target.config_sync_group.config_sync_group / 1052f40f50f9 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-2054a037c9d64d979b7e5977309895a409f9af4abcca00bd24c7e70e5e31ebf6)
- [discovery_target.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-8598c4fb654452cc64a846b13a58e10f57ba05580c1e9d7cdb693b0e6987ca1a)
- discovery_target.config_sync_group.config_sync_group

<a id="canonical-16f422cfa2a46f5f8ce21fafe59b8ab271532b0658e5d68917f2e2ce25d6d9fd"></a>

Type: `"object"`. list nested block, Optional.

Reference. Select new ConfigSyncGroup.

Upstream description:

Select new ConfigSyncGroup.

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
config_sync_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-ecf4766eb6c0b712684e785e31b05a7f12ce623d143891279ba7800d9a1d26c3"></a>

## Direct properties — discovery_target.config_sync_group.config_sync_group / 1052f40f50f9 / 3

<a id="canonical-27130147c44350b40fa7d045d4ab27a9caa45ad36932bff6389b91b7a6e1c67c"></a>

<a id="canonical-949d8edf28185539824b5edf7dc3c068acf46dc34a8e80cffee16cd37793e08b"></a>

## kind property — discovery_target.config_sync_group.config_sync_group / 1052f40f50f9 / 4

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

<a id="canonical-2798f643301c60f000f879546ed91c2ccba821ffe57d820ea0a08dea4d2a9879"></a>

<a id="canonical-539b860ac3a048726a37a2f1212f46f4379aa7aa0cf28aab5f4dc980816ec484"></a>

## name property — discovery_target.config_sync_group.config_sync_group / 1052f40f50f9 / 5

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

<a id="canonical-7d70a9fdd08c376fbbbdcb46c3bed10f765774aadcd7fdb023ce6af0fbe351dc"></a>

<a id="canonical-8c842d6b1fb87a471fe30c9ded1956ff2342acb61769be4bb6f78e6c1cfb13eb"></a>

## namespace property — discovery_target.config_sync_group.config_sync_group / 1052f40f50f9 / 6

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

<a id="canonical-20598e24525bc4a6f7e0c992c6018ed29a508d61226abdd9612fd5deed4dcc61"></a>

<a id="canonical-244aa40d6984ee498c5eddcdcc8e2413d0a4e796914c2c50d6609005860ffac1"></a>

## tenant property — discovery_target.config_sync_group.config_sync_group / 1052f40f50f9 / 7

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

<a id="canonical-6252daaa85b3d69b6e075a443e96556e9b07fe0bbdc82d6f588673a0db94c0fa"></a>

<a id="canonical-e4a04ac12a58987885aa4abb6f7d3be068279e855f12737da82ef92aae4d7872"></a>

## uid property — discovery_target.config_sync_group.config_sync_group / 1052f40f50f9 / 8

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

<a id="canonical-aaeee883bf5bc19af73fc22d122ec31132b8b12553158665fe09e5511f03e76d"></a>

## Next pages — discovery_target.config_sync_group.config_sync_group / 1052f40f50f9 / 9

- [discovery_target.config_sync_group](resources--nginx_service_discovery--reference--group-001.md#canonical-8598c4fb654452cc64a846b13a58e10f57ba05580c1e9d7cdb693b0e6987ca1a)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)

<a id="canonical-c24f0ca3ef08690c9d09aca923162a6dd0553a3d68f23de93a2a6dbe2cc8809f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43f2af28d4b8a443881763e8e9fc087ba4cdb67245e753c242105af74359a52c"></a>

## discovery_target.nginx_instance — discovery_target.nginx_instance / 7d7d5107b209 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-2054a037c9d64d979b7e5977309895a409f9af4abcca00bd24c7e70e5e31ebf6)
- discovery_target.nginx_instance

<a id="canonical-78505e910ac7bdd8d5fb2e27ed54d006019cd2ec718a1fe780188857d47c9b0b"></a>

Type: `"object"`. single nested block, Optional.

NGINXInstance Reference. Select new NGINX Instance.

Upstream description:

Select new NGINX Instance.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("nginx_instance")}
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
nginx_instance {
  # Configure direct properties listed below.
}
```

<a id="canonical-529631d2ab2d025bb03823818467c39b1a0176ece30bb98c7b7c70ea76776978"></a>

## Direct properties — discovery_target.nginx_instance / 7d7d5107b209 / 3

- [nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-3c4e30b9a162f7376a1a38c72fa468261c65d89fbff418ff24fe93461f885a64): complete subsection reference.

<a id="canonical-c2a89cac414c6ccd8aadba1e02cd712ac9de800c4703698ab3f2c5876d105bb6"></a>

## Next pages — discovery_target.nginx_instance / 7d7d5107b209 / 4

- [discovery_target.nginx_instance.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-3c4e30b9a162f7376a1a38c72fa468261c65d89fbff418ff24fe93461f885a64)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-2054a037c9d64d979b7e5977309895a409f9af4abcca00bd24c7e70e5e31ebf6)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)

<a id="canonical-3c4e30b9a162f7376a1a38c72fa468261c65d89fbff418ff24fe93461f885a64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93c19ed17e4a75114a16c975f5dee520356fcb6c7d334157cdb7090b4d817e88"></a>

## discovery_target.nginx_instance.nginx_instance — discovery_target.nginx_instance.nginx_instance / 21d1e0765ee7 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- [discovery_target](resources--nginx_service_discovery--reference--group-001.md#canonical-2054a037c9d64d979b7e5977309895a409f9af4abcca00bd24c7e70e5e31ebf6)
- [discovery_target.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-c24f0ca3ef08690c9d09aca923162a6dd0553a3d68f23de93a2a6dbe2cc8809f)
- discovery_target.nginx_instance.nginx_instance

<a id="canonical-e8099d466372c0d9dbba4102956518ba5f1a1a8f9fefbc589fd98db5d2de6e9e"></a>

Type: `"object"`. list nested block, Optional.

Reference. Select new NGINX Instance.

Upstream description:

Select new NGINX Instance.

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
nginx_instance {
  # Configure direct properties listed below.
}
```

<a id="canonical-c413a8b49c97dca110eff6f89fee50f59474f1f66d6dfc9107f2721e0a94f4ed"></a>

## Direct properties — discovery_target.nginx_instance.nginx_instance / 21d1e0765ee7 / 3

<a id="canonical-b5be6b837223b0f83c865035e126b672e7ecb3f9e4e3dbdfd87bbb34d292a5a5"></a>

<a id="canonical-97449b781832e62316f90aedf5ee2ca65b1a112a11b154388aae54c6c6aa4486"></a>

## kind property — discovery_target.nginx_instance.nginx_instance / 21d1e0765ee7 / 4

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

<a id="canonical-9d5bffaff73aee7f1d2ec2cd9a1c2026e7b6e085b0d6dd77b769e7e798d6b1d2"></a>

<a id="canonical-1be2d3afa86537efb55644033a7eebef07bb250094be41723d62371b3d6c1d08"></a>

## name property — discovery_target.nginx_instance.nginx_instance / 21d1e0765ee7 / 5

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

<a id="canonical-dc741cef4a95896994f1ba5f1f68142ab5c3ae15173bf65d921792d985a26b78"></a>

<a id="canonical-1b3be46213309ecb20f98172cd3a786ae468635b3a76ff5d2bd3e17dea949872"></a>

## namespace property — discovery_target.nginx_instance.nginx_instance / 21d1e0765ee7 / 6

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

<a id="canonical-70191eb5c643d9c3dc666b3d9df5c1e8c4564574c587d551e89ccf81cd2895c1"></a>

<a id="canonical-9e0f4a6b92e610d4938f8026bdecd1b8ed56b837d8085a7d7fa4f39c9fcbfb0e"></a>

## tenant property — discovery_target.nginx_instance.nginx_instance / 21d1e0765ee7 / 7

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

<a id="canonical-4e70255fbd81eba06a2e117d8f177ab16cb200f5e714fc9edce08eaaf7ab68a7"></a>

<a id="canonical-4753c2d4bfbef9be9a957ff96ae711a44212f233762f9b59cd8cbf6b3f1fba4a"></a>

## uid property — discovery_target.nginx_instance.nginx_instance / 21d1e0765ee7 / 8

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

<a id="canonical-aa358ce2c06b5fbb78572d47a115d80141403b2655334fac2469f09a470a75a6"></a>

## Next pages — discovery_target.nginx_instance.nginx_instance / 21d1e0765ee7 / 9

- [discovery_target.nginx_instance](resources--nginx_service_discovery--reference--group-001.md#canonical-c24f0ca3ef08690c9d09aca923162a6dd0553a3d68f23de93a2a6dbe2cc8809f)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)

<a id="canonical-aa29d75af6e1655c8211234f59f20760fb3899fff39a441e11340c60ac7f4f9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b439d3883494353c0819d03056526e40cdbd38aa3ef3af4145b3fb2492d296ab"></a>

## server_block_filters — server_block_filters / 2d950ae7e90c / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- server_block_filters

<a id="canonical-c86a6a9bf1ea1c8cd3622faaff016bf6befc32c7d0e812592a034a244d6d4c8a"></a>

Type: `"object"`. list nested block, Optional.

Filters discovered server blocks based on server name, domain and ports. Atleast, one field should
be populated for each filter. X-textBlockContent: If no filters are specified, all server blocks
will be discovered by default.

Upstream description:

Filters discovered server blocks based on server name, domain and ports. Atleast, one field should
be populated for each filter.

X-textBlockContent: If no filters are specified, all server blocks will be discovered by default.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_block_filters {
  # Configure direct properties listed below.
}
```

<a id="canonical-add48c2d910cff360a8ac749aa8b8ec7252f2120521ad7fd366d68f417718635"></a>

## Direct properties — server_block_filters / 2d950ae7e90c / 3

<a id="canonical-971b7ba000cd9913a4b38ab088e2f0dfb60e876735cda14887a63080c0fc810d"></a>

<a id="canonical-94cb6474890d53a6c5ef0a8a3c57e7a0c76c8969bc1807b26657a93ae8e195d2"></a>

## name_regex property — server_block_filters / 2d950ae7e90c / 4

Type: `"string"`. Optional.

Regular expression to match the server name or domain that must be discovered.

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-7b634dae1709d55836179ea8a146c8ddac9f3cb05ceae9ef50b7d65a7299bf61"></a>

<a id="canonical-de38321ab5e66c3c01ea64b5f5498ce8e37d121219c9acbf6526729150fad170"></a>

## port_ranges property — server_block_filters / 2d950ae7e90c / 5

Type: `"string"`. Optional.

String containing a comma separated list of individual service ports or port ranges. Each port range
consists of a single port or two ports separated by '-'. For example, 8000-8191.

Upstream description:

A string containing a comma separated list of individual service ports or port ranges. Each port
range consists of a single port or two ports separated by "-". For example, 8000-8191. Maximum
number of ports allowed is 1024.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-1c309f65f94961a05e376dedfcf0a8c11b60bf4c6622ae115947084f1209c920"></a>

## Next pages — server_block_filters / 2d950ae7e90c / 6

- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)

<a id="canonical-fe720eb086b57545b2fb0f63f5461fc0b63c51e4b8e942105bd7f5b7cfd791f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72cb6bf91a6332f36318d7c9ee4431c9ab495705c605f9037a35bc2057df2fe7"></a>

## timeouts — timeouts / f8c8a4c7846c / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- timeouts

<a id="canonical-d0fca5828481edc50c29f4d915ab395980f080db92b39ff51f8b959d7ced822d"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-a46610675b5c66a397c967f6ad4bd61eb372c4c289da693e0e27dc3a09faa371"></a>

## Direct properties — timeouts / f8c8a4c7846c / 3

<a id="canonical-bc3b9923a5115c97da21b102dde71575da63a437f9e725dea1eae052e2514e08"></a>

<a id="canonical-e690a70cc2d054438d63ce0693c153f5c6c500d1a720d55e4a030ea59fff14da"></a>

## create property — timeouts / f8c8a4c7846c / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-40a61f66ac4f17b16204ab3867932b2cfc1d93dbe13f8a1a5067f30dee515ee7"></a>

<a id="canonical-c6f09ac25b12ef4898a8398a3f1620211752bc6253aaf4844ec0d2df4e94eee0"></a>

## delete property — timeouts / f8c8a4c7846c / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-78157e1c1a76a1d8e033d3091f312e756f8fe9e4f00e6302901af0d62e1dc120"></a>

<a id="canonical-f2857171e9d49db65a880635ef4230e1cbfbecbdbe36c349b3d63323d6dca67a"></a>

## read property — timeouts / f8c8a4c7846c / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-eb609d4d71d73f24017ef9ce1881bdf681475ccbcb61a407791b99572d32fcc1"></a>

<a id="canonical-f14ffdf43fe23ce32eda3dfb8c1ce023ec96b938a297d210eb2c52204e7db746"></a>

## update property — timeouts / f8c8a4c7846c / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3cff868f603e4bc2f5dbfe0bbc04c970032ccba644942cd0a27e3f2f6269c165"></a>

## Next pages — timeouts / f8c8a4c7846c / 8

- [Property reference](resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b)
