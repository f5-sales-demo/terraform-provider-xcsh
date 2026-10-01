---
page_title: "xcsh_registration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration reference."
---

# xcsh_registration reference

<a id="canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f608e6d7f925219b460621d0726763791da36d56226ce03b9aa84cd1040240aa"></a>

## Property reference — Property reference / ea1062c425d0 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- Property reference

<a id="canonical-55aa0bcfe83b13c9b6f3cb53d68b169b789fd675f5078705ef52ba9f30859b82"></a>

## Direct properties — Property reference / ea1062c425d0 / 3

<a id="canonical-014d7a17db1e80cd617a18219268e96bd976caa7e9e711f6c134936606261fab"></a>

<a id="canonical-813cae1d6953ff5f8a51988f261885dc15e2a0e26accaece7a479c8c66a5f1cf"></a>

## annotations property — Property reference / ea1062c425d0 / 4

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

<a id="canonical-700354b6669bf0ab4fea206cbbf13ed046233b6938636bc379021237bed55ff8"></a>

<a id="canonical-d25121964385d2eadd77713999f524583708f734115ed2707117ee8fec325994"></a>

## description property — Property reference / ea1062c425d0 / 5

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

<a id="canonical-27552a8fee52cc83dfe31b210920d537c5ae80da3c0a62c4bfe7cb53acfd5672"></a>

<a id="canonical-3ed44925d1d747f98c65ce845182e5931ea4ccf0e132f1bb50cc8fa94770a58c"></a>

## disable property — Property reference / ea1062c425d0 / 6

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

<a id="canonical-bd6715120686e59d5355bedccf1f2e0a4cf8d718803af761be037eb7d7651c89"></a>

<a id="canonical-3ca87ca5f5417e577ae1788b7287a989fd93a24a1619256be07d7b3d083f2613"></a>

## id property — Property reference / ea1062c425d0 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5): complete subsection reference.

<a id="canonical-d36fb260271ce4e102ac7199a920bbbb37197d4c7b7d60479a0a33185896c80f"></a>

<a id="canonical-580ed61771e5f6d5d03c13205f08f8622c715ff9a76e4c8b5b9f06a82846c090"></a>

## labels property — Property reference / ea1062c425d0 / 8

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

<a id="canonical-b83d746b3ac0656ca8efaf27839679370c983cf7437ae9f677acab04a5d346d2"></a>

<a id="canonical-d2c4dd9af8045c24c627627815dcf1ca4ffd92fc75746e5060f176d54470626b"></a>

## name property — Property reference / ea1062c425d0 / 9

Type: `"string"`. Required.

Name of the Registration. Must be unique within the namespace.

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

<a id="canonical-b59b9d5fddb10a1821517d7e433bf6c19c34d9540d71cb151c35de9216bf6dea"></a>

<a id="canonical-61913d492288cc03a001293c443d238a185f12f4d41f6c9664204387ab2a54a7"></a>

## namespace property — Property reference / ea1062c425d0 / 10

Type: `"string"`. Required.

Namespace where the Registration is created.

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

- [passport](resources--registration--reference--group-001.md#canonical-5847986b1feea2a4410054d5badbb079e13e8d02a85f50e7d1513ded659f8267): complete subsection reference.

- [timeouts](resources--registration--reference--group-001.md#canonical-e736dfc22ce4ac100b194673d184aade88f5363c160371788e6bc9c276d81652): complete subsection reference.

<a id="canonical-afbdaa38ad15c835f94ca34784f82a6deae750944d1b47563d570ee15e474661"></a>

<a id="canonical-9d8d94fd8b777a77f87ac48dbb0a7cb056e972aa13e301344ecca16de69b215b"></a>

## token property — Property reference / ea1062c425d0 / 11

Type: `"string"`. Required.

Token is used for machine and tenant identification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "security",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 20
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

<a id="canonical-39f2a8875148672c087898aff7cbc6a5ce6c49ee83c145c971f9a94f2b2fc983"></a>

## All schema paths — Property reference / ea1062c425d0 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--registration--reference--group-001.md#canonical-014d7a17db1e80cd617a18219268e96bd976caa7e9e711f6c134936606261fab) |
| `description` | [description](resources--registration--reference--group-001.md#canonical-700354b6669bf0ab4fea206cbbf13ed046233b6938636bc379021237bed55ff8) |
| `disable` | [disable](resources--registration--reference--group-001.md#canonical-27552a8fee52cc83dfe31b210920d537c5ae80da3c0a62c4bfe7cb53acfd5672) |
| `id` | [id](resources--registration--reference--group-001.md#canonical-bd6715120686e59d5355bedccf1f2e0a4cf8d718803af761be037eb7d7651c89) |
| `infra` | [infra](resources--registration--reference--group-001.md#canonical-3f898b9932f81db1b2749869bdc1a68644589b3305f200adc9f97fb67665c5a7) |
| `infra.availability_zone` | [infra.availability_zone](resources--registration--reference--group-001.md#canonical-53d2c555c02f42ac2acda7312637b32977842d10237d62e5e3f9cf613652fe95) |
| `infra.bond_config` | [infra.bond_config](resources--registration--reference--group-001.md#canonical-1711756b2112aee9221eb98c4aa6546f27f71fa762b9ce5b86ff3e31e88a6acc) |
| `infra.bond_config.interfaces` | [infra.bond_config.interfaces](resources--registration--reference--group-001.md#canonical-0f77dc1113bc7124edce076e43d2c8e2f7a611ea759d72fbc2504819d633bc33) |
| `infra.bond_config.mode` | [infra.bond_config.mode](resources--registration--reference--group-001.md#canonical-866a902897ab3742c37a3f3df1e464e31739ddca0564a696016e504523a8c0d1) |
| `infra.bond_config.name` | [infra.bond_config.name](resources--registration--reference--group-001.md#canonical-fe47aefc466f30c79c052396a037e9b5a07d485f1ed04884022d1206790cbc5f) |
| `infra.certified_hw` | [infra.certified_hw](resources--registration--reference--group-001.md#canonical-9e670446bd86d16a47f75dc82f02843ce2acb3320f2a54444d70cd3a5e4e24e0) |
| `infra.domain` | [infra.domain](resources--registration--reference--group-001.md#canonical-a719ce3b07261f50801a7a6c5bac104d2c425905a54fee2e9aa76837e3169d05) |
| `infra.hostname` | [infra.hostname](resources--registration--reference--group-001.md#canonical-6faa1f67dd1f39c3391693d99e5765fc4463b1d4107329bde597f9c19f6db088) |
| `infra.hugepages` | [infra.hugepages](resources--registration--reference--group-001.md#canonical-8ec6d9821c5f0649e1cc6dbfa73348d4fec3f14a0361582bf8e841f8f4f780ce) |
| `infra.hugepages.free` | [infra.hugepages.free](resources--registration--reference--group-001.md#canonical-aeaabcea4057b10c331f4b1ebc88196a1c09fa300595b9b9dd87e31915e6f319) |
| `infra.hugepages.page_size` | [infra.hugepages.page_size](resources--registration--reference--group-001.md#canonical-4fa79b2a657ad4681459ce07a5198533745d957aeda28fcd6c157f9801f20cf0) |
| `infra.hugepages.total` | [infra.hugepages.total](resources--registration--reference--group-001.md#canonical-46e837a0b16435caae57d5c295bee56939caab6fa5d8e3a489fb30754d65bf9d) |
| `infra.hw_info` | [infra.hw_info](resources--registration--reference--group-001.md#canonical-969d0c6fd11fc46c4a44cb1c1d77e5783086f992f53010b54268ac585da5c1a8) |
| `infra.hw_info.bios` | [infra.hw_info.bios](resources--registration--reference--group-001.md#canonical-a684563e7210b9f6bc80d2a69e5e1c403534656e1c148cffd15f99e06cac4bee) |
| `infra.hw_info.bios.date` | [infra.hw_info.bios.date](resources--registration--reference--group-001.md#canonical-25379bb9ec7502b7c0c01304a479b5f339b4cdbcf9df5a077bb10dec93650414) |
| `infra.hw_info.bios.vendor` | [infra.hw_info.bios.vendor](resources--registration--reference--group-001.md#canonical-42921eea47ae8aa7133bf84cfec20ea24c8331167a097d7d635c3beee9d0658d) |
| `infra.hw_info.bios.version` | [infra.hw_info.bios.version](resources--registration--reference--group-001.md#canonical-8443d2712f9d5ff3abecf203de2c074ab51b1e5890b15d33b08027f544df12d7) |
| `infra.hw_info.board` | [infra.hw_info.board](resources--registration--reference--group-001.md#canonical-ae3d652e4ed148282c8c8ff6695c1c5600c2a29dfdf86f8fb66930275d1a84ff) |
| `infra.hw_info.board.asset_tag` | [infra.hw_info.board.asset_tag](resources--registration--reference--group-001.md#canonical-094f44cfe2dc821887d9f67ee31cb2b99a286975f8d3a33a7d42542cbc452c3f) |
| `infra.hw_info.board.name` | [infra.hw_info.board.name](resources--registration--reference--group-001.md#canonical-4a031a904f9694b383787a257e912cac6134637b0f72f520d13a1e2663cdd26f) |
| `infra.hw_info.board.serial` | [infra.hw_info.board.serial](resources--registration--reference--group-001.md#canonical-5f090bcd0f07dd0f0fe21ba2da373a99dea69f8d3dd7ef12a4b81127ad818199) |
| `infra.hw_info.board.vendor` | [infra.hw_info.board.vendor](resources--registration--reference--group-001.md#canonical-5e8ddcec633570197e14af5ab203902d7857180d2a73a6d846ee65ead31cf9c9) |
| `infra.hw_info.board.version` | [infra.hw_info.board.version](resources--registration--reference--group-001.md#canonical-ff2d07c13a11d8f7000a6f14dfef1d9dcbbc5b521816c2c0ea75490f0710f2cf) |
| `infra.hw_info.chassis` | [infra.hw_info.chassis](resources--registration--reference--group-001.md#canonical-761617314004971994a7e39b29e4e4aa4b3ffce50960a6abe43fab3e0b8ea642) |
| `infra.hw_info.chassis.asset_tag` | [infra.hw_info.chassis.asset_tag](resources--registration--reference--group-001.md#canonical-73eb21b22eafc524b2fefc9ca25fac58e2c70effbf259e87dc8b42784a09e012) |
| `infra.hw_info.chassis.serial` | [infra.hw_info.chassis.serial](resources--registration--reference--group-001.md#canonical-ee005e8737ddd86b6ece137a9df2a234853d67991b0105b5a6a9e000ad4b23cc) |
| `infra.hw_info.chassis.type` | [infra.hw_info.chassis.type](resources--registration--reference--group-001.md#canonical-843562fc73d14e9f3c7f81594b8cbc491af0ac96b41b0f2a8b062732e0bbe5d6) |
| `infra.hw_info.chassis.vendor` | [infra.hw_info.chassis.vendor](resources--registration--reference--group-001.md#canonical-863826996c1e21f9cdf7ed4ac29cba6e08e63cf705fefc0f57d8f594c4a49ca1) |
| `infra.hw_info.chassis.version` | [infra.hw_info.chassis.version](resources--registration--reference--group-001.md#canonical-f8e03429720962bade703cb7de8c09d597ecdbb2ff60bf58244039f1d999eeba) |
| `infra.hw_info.cpu` | [infra.hw_info.cpu](resources--registration--reference--group-001.md#canonical-2cee95b43f9c73eadf0306cbd5dad4dbd33ba67b168d535c502402f7850219de) |
| `infra.hw_info.cpu.cache` | [infra.hw_info.cpu.cache](resources--registration--reference--group-001.md#canonical-fa856a116614374a85407517f8ac0d09f35396db1e9e833d4eff89824de98bf2) |
| `infra.hw_info.cpu.cores` | [infra.hw_info.cpu.cores](resources--registration--reference--group-001.md#canonical-8df80c2a47b1e9e339e2594f622daa9146071236845d5c73d7308f3fa7dff24f) |
| `infra.hw_info.cpu.cpus` | [infra.hw_info.cpu.cpus](resources--registration--reference--group-001.md#canonical-f5869ada211ee5e929b264d501cbf7a49531a84deda358d5299f3142f5d3bbce) |
| `infra.hw_info.cpu.model` | [infra.hw_info.cpu.model](resources--registration--reference--group-001.md#canonical-26f3e166c14148ba29f1ed2383c57ba5fa079d7fccf6c45242fb222e93ec0453) |
| `infra.hw_info.cpu.speed` | [infra.hw_info.cpu.speed](resources--registration--reference--group-001.md#canonical-2c1b73df433ffb276b324083ab3decb854118df8682fb4eff1beda35514c7578) |
| `infra.hw_info.cpu.threads` | [infra.hw_info.cpu.threads](resources--registration--reference--group-001.md#canonical-079870743c583c26123381207c6704d83b78af674928fa7839ec0a579d3126d9) |
| `infra.hw_info.cpu.vendor` | [infra.hw_info.cpu.vendor](resources--registration--reference--group-001.md#canonical-5f8edf93828bf92377fbf26e0c2926fdda44e06cb066e1deece99a1ac9983096) |
| `infra.hw_info.gpu` | [infra.hw_info.gpu](resources--registration--reference--group-001.md#canonical-f0146d9a4fe41322e1b40be5e49f9b48d0726beefc00ce7f584c03367857394d) |
| `infra.hw_info.gpu.cuda_version` | [infra.hw_info.gpu.cuda_version](resources--registration--reference--group-001.md#canonical-e23a3568afbc554010628d57a220f00c696ba29fed4e50f95db9d53d51cd4c18) |
| `infra.hw_info.gpu.driver_version` | [infra.hw_info.gpu.driver_version](resources--registration--reference--group-001.md#canonical-c496b39f550a4c6cead52ecfd4ab1b95bc538464904d2c46a5a45d591ef39f69) |
| `infra.hw_info.gpu.gpu_device` | [infra.hw_info.gpu.gpu_device](resources--registration--reference--group-001.md#canonical-28e9e0907f217bb3d3734a7a160ee554c3e6150d8040755276baea5f103ef545) |
| `infra.hw_info.gpu.gpu_device.id` | [infra.hw_info.gpu.gpu_device.id](resources--registration--reference--group-001.md#canonical-a6d3b14b2b1ebe153c5398218574b84fa847eb344c6699988ef8f13226f2ac65) |
| `infra.hw_info.gpu.gpu_device.processes` | [infra.hw_info.gpu.gpu_device.processes](resources--registration--reference--group-001.md#canonical-c4342dd3d54d0b8bc2735cf15b222f0ffa2be5d44305cbec6ac212256b183edf) |
| `infra.hw_info.gpu.gpu_device.product_name` | [infra.hw_info.gpu.gpu_device.product_name](resources--registration--reference--group-001.md#canonical-1f12b1e1db6cb56db26e2793cec9022041f51305031d29353358bb3084f241e4) |
| `infra.hw_info.kernel` | [infra.hw_info.kernel](resources--registration--reference--group-001.md#canonical-6ca1a07d68f8bbc4fec9c951ede7cc6fba9516cf081afc5807accb4c2021a18d) |
| `infra.hw_info.kernel.architecture` | [infra.hw_info.kernel.architecture](resources--registration--reference--group-001.md#canonical-2958f4db98c3aeaeba69af34929dccf4ac30e4c6dbe7a6951e449b53bced5fb8) |
| `infra.hw_info.kernel.release` | [infra.hw_info.kernel.release](resources--registration--reference--group-001.md#canonical-9b0e2c85a807d410b0d67d4059ecc9dc651317fe8deaf956199461ddbe2a0d17) |
| `infra.hw_info.kernel.version` | [infra.hw_info.kernel.version](resources--registration--reference--group-001.md#canonical-43e7021e32489a51dfeffc01a56697708814b6a90d9318f8171f8e10a16850ef) |
| `infra.hw_info.memory` | [infra.hw_info.memory](resources--registration--reference--group-001.md#canonical-80d295de4e47f2e88ca4a1c1be6d930b081dd220686e4fab9da48f490837533c) |
| `infra.hw_info.memory.size_mb` | [infra.hw_info.memory.size_mb](resources--registration--reference--group-001.md#canonical-0b63d7e999e402ce8d5bd1b6cb0adf81481d1aeb52cc2a52085b81992bdf3a0c) |
| `infra.hw_info.memory.speed` | [infra.hw_info.memory.speed](resources--registration--reference--group-001.md#canonical-d6cc2c7d59262629a302a1234d049c5b07291a1db4790ca1007320fc03a3da03) |
| `infra.hw_info.memory.type` | [infra.hw_info.memory.type](resources--registration--reference--group-001.md#canonical-9c191138e6abec8896113258f3d6d1f1c583a160311ab209a744b4fbe0bfa509) |
| `infra.hw_info.network` | [infra.hw_info.network](resources--registration--reference--group-001.md#canonical-259012996ccf1e7492bebc548651c3136e5c69c339912ba437631b310eb6acbb) |
| `infra.hw_info.network.driver` | [infra.hw_info.network.driver](resources--registration--reference--group-001.md#canonical-e3f6c88cbe2f9cafca506fa6b90c2fa04c8827277f901d9cf689f4490cf7fda9) |
| `infra.hw_info.network.ip_address` | [infra.hw_info.network.ip_address](resources--registration--reference--group-001.md#canonical-04ee52f2fcbe58c4eea60b4a83ac90fec31b9ae619ad3b02bc94274fd7c96f4c) |
| `infra.hw_info.network.link_quality` | [infra.hw_info.network.link_quality](resources--registration--reference--group-001.md#canonical-eb3b8d4239cc039d0ea18f3247392780cde30cd6b9e012a12f1c6a627ca89467) |
| `infra.hw_info.network.link_type` | [infra.hw_info.network.link_type](resources--registration--reference--group-001.md#canonical-37798952b43c5d312ca761d6cb37714358c16d2fb1ed8cb46fb6c1dfe516ea51) |
| `infra.hw_info.network.mac_address` | [infra.hw_info.network.mac_address](resources--registration--reference--group-001.md#canonical-a3c5d5405c220be760268125f1c48865b6c127f26c27b6a6d7c3a64bea06b8f7) |
| `infra.hw_info.network.name` | [infra.hw_info.network.name](resources--registration--reference--group-001.md#canonical-c0edabbf2b77f7484b70dd56ce3f8ed171b9bac6c3a5cc774f0f347da809a78b) |
| `infra.hw_info.network.port` | [infra.hw_info.network.port](resources--registration--reference--group-001.md#canonical-e9580e7c2d1448b3b086e2e09039ee370ca479aa902edb7dd33a7e8355b64069) |
| `infra.hw_info.network.speed` | [infra.hw_info.network.speed](resources--registration--reference--group-001.md#canonical-a360a3275fb960023cdcdb5ae81735b3c6af2a498d697a6aa1b4ec09b7c060b4) |
| `infra.hw_info.numa_nodes` | [infra.hw_info.numa_nodes](resources--registration--reference--group-001.md#canonical-a84d04d419a2dfa1a6cbac2152d60c222ff52eb17d991422a51ae3cd3088edd6) |
| `infra.hw_info.os` | [infra.hw_info.os](resources--registration--reference--group-001.md#canonical-6d08774c468187ee4765d04f7cd821ac7465f741b2e5434ae4832ea4dd299f72) |
| `infra.hw_info.os.architecture` | [infra.hw_info.os.architecture](resources--registration--reference--group-001.md#canonical-4f79c2e92416df8cc51f11d0e6566217199c88d0955d4becf50ee5dd9cf15fdd) |
| `infra.hw_info.os.name` | [infra.hw_info.os.name](resources--registration--reference--group-001.md#canonical-c2d6f16b01f29d250e0a125424992e357182c189928fbeed4f1b7187fbbeb910) |
| `infra.hw_info.os.release` | [infra.hw_info.os.release](resources--registration--reference--group-001.md#canonical-f394db4b7259f120d148d12054feae00052f04fd3fd5ef2c211ceb1f53c5c00f) |
| `infra.hw_info.os.vendor` | [infra.hw_info.os.vendor](resources--registration--reference--group-001.md#canonical-5d53a3e10c6040100e7982140dfaa6249facbc7c648d77585d7d955a17c6430d) |
| `infra.hw_info.os.version` | [infra.hw_info.os.version](resources--registration--reference--group-001.md#canonical-cff9ea6c47b406b1fd4e1c1d42bf37ef76cce3ffbdd6d5eff191e2f166ebbbf9) |
| `infra.hw_info.product` | [infra.hw_info.product](resources--registration--reference--group-001.md#canonical-2b2accc84cad6e7ec6bf3759572e6667a1b080926da06de43078eaf3a0f99f41) |
| `infra.hw_info.product.name` | [infra.hw_info.product.name](resources--registration--reference--group-001.md#canonical-f885f3b0a8998a4526b2fcab6470e42a83722ace1eafc9daccedd1f68f5e46ed) |
| `infra.hw_info.product.serial` | [infra.hw_info.product.serial](resources--registration--reference--group-001.md#canonical-c8983b978e1606813a17f27b8848e15e8d8e66fe1549f5413efaf6fe53ecee54) |
| `infra.hw_info.product.vendor` | [infra.hw_info.product.vendor](resources--registration--reference--group-001.md#canonical-82385f3f1b79256b8c5f4ac9fc348426e8ce6bb0d137a06ced1513916d3fdd2d) |
| `infra.hw_info.product.version` | [infra.hw_info.product.version](resources--registration--reference--group-001.md#canonical-54f11d6244095f83a9b506db760e36546fe7e6f8ce5df3c0a92fb3326e23e674) |
| `infra.hw_info.storage` | [infra.hw_info.storage](resources--registration--reference--group-001.md#canonical-a40af56489ffe975ec3d30828bff812eeb821258c2aad5f3aea21afdb8c51f69) |
| `infra.hw_info.storage.driver` | [infra.hw_info.storage.driver](resources--registration--reference--group-001.md#canonical-2a98cd82e733f3c86aa35972ebb81c572fcc25c71df22e4fdeb327150ab8964f) |
| `infra.hw_info.storage.model` | [infra.hw_info.storage.model](resources--registration--reference--group-001.md#canonical-df90c7f9ec70796f413b30849fc9e9e0cbc9743e8f7092a5d6115bb2f0760eb5) |
| `infra.hw_info.storage.name` | [infra.hw_info.storage.name](resources--registration--reference--group-001.md#canonical-cf835d6f7a5544afb7b4699b7ff641f24403cc81e1dc1259315dd99c3c0a37a6) |
| `infra.hw_info.storage.serial` | [infra.hw_info.storage.serial](resources--registration--reference--group-001.md#canonical-6a1d1336c2f676f14ca2275682d7ab6dc5564db330b4303453dc103fe50a7835) |
| `infra.hw_info.storage.size_gb` | [infra.hw_info.storage.size_gb](resources--registration--reference--group-001.md#canonical-dbb99b885b642d2bc59bc3dde8839a474ae51a67a4929a2c1571111d015c0845) |
| `infra.hw_info.storage.vendor` | [infra.hw_info.storage.vendor](resources--registration--reference--group-001.md#canonical-91b379cb0b18d996bfbbf7268b2c7326427bc0a43f6830eb1a96032ec36e900e) |
| `infra.hw_info.usb` | [infra.hw_info.usb](resources--registration--reference--group-001.md#canonical-ede6e1b7b67983d962742fa3ee4fd3e664c4e24e61ab66e2a8df392d7ad552b5) |
| `infra.hw_info.usb.address` | [infra.hw_info.usb.address](resources--registration--reference--group-001.md#canonical-d71c4af252c647cdf99de2f2fcdac51c15a66be44ddcfa955042aa6a456fa841) |
| `infra.hw_info.usb.b_device_class` | [infra.hw_info.usb.b_device_class](resources--registration--reference--group-001.md#canonical-2de97a180e7ebe1dba93ffbe294c24788cd8d9a448644db7272eed72bd6bbfe6) |
| `infra.hw_info.usb.b_device_protocol` | [infra.hw_info.usb.b_device_protocol](resources--registration--reference--group-001.md#canonical-55352536d2ec1967ea505975145f92a7936d3f8d185838c9138820c58ed2557c) |
| `infra.hw_info.usb.b_device_sub_class` | [infra.hw_info.usb.b_device_sub_class](resources--registration--reference--group-001.md#canonical-4c5bbbec7cddef1a76dc9f57c10fa0df8313674a08853e4f51593ed315d05732) |
| `infra.hw_info.usb.b_max_packet_size` | [infra.hw_info.usb.b_max_packet_size](resources--registration--reference--group-001.md#canonical-536a16e89eece88227c981448f146d659a50b11ec2209699cb02379443294897) |
| `infra.hw_info.usb.bcd_device` | [infra.hw_info.usb.bcd_device](resources--registration--reference--group-001.md#canonical-7ef3aa7ad159ce26f5c62e1abb34ab8c86cda74144ebe1b8e48989563d28a6b4) |
| `infra.hw_info.usb.bcd_usb` | [infra.hw_info.usb.bcd_usb](resources--registration--reference--group-001.md#canonical-02f0a3a24710cce172dc4719b49024e51f11dc5acf0c6be56bfc2b4bc3679bf7) |
| `infra.hw_info.usb.bus` | [infra.hw_info.usb.bus](resources--registration--reference--group-001.md#canonical-0bb83ab98a35a6d2ee4fc04e04a8be1b9b4aed3f02f9184478943ddddaf81424) |
| `infra.hw_info.usb.description_spec` | [infra.hw_info.usb.description_spec](resources--registration--reference--group-001.md#canonical-d273d5afa9746b3f868ad59ab35a5d47805f1891706c03fd8ff1d8c2f151ff0b) |
| `infra.hw_info.usb.i_manufacturer` | [infra.hw_info.usb.i_manufacturer](resources--registration--reference--group-001.md#canonical-4da22b3052d7324b129720c3292b52c820c95f54713473668c44411863c19342) |
| `infra.hw_info.usb.i_product` | [infra.hw_info.usb.i_product](resources--registration--reference--group-001.md#canonical-a0ce31e9c5acc232f1b470f087a317654dc74f2d923f70961e6f40dfdf222947) |
| `infra.hw_info.usb.i_serial` | [infra.hw_info.usb.i_serial](resources--registration--reference--group-001.md#canonical-58ae3ba085bff88c0b3de42fd7e7c7b86335c9a6176e91e1bd6cd74b17404ad6) |
| `infra.hw_info.usb.id_product` | [infra.hw_info.usb.id_product](resources--registration--reference--group-001.md#canonical-2a05762248b1aa90b11eb444561f3277707061ae364de092ea1444100fc142a1) |
| `infra.hw_info.usb.id_vendor` | [infra.hw_info.usb.id_vendor](resources--registration--reference--group-001.md#canonical-5d20199e109fb0a6fac44bfbefff9fd1cf523e630796db9931507ea730b05417) |
| `infra.hw_info.usb.port` | [infra.hw_info.usb.port](resources--registration--reference--group-001.md#canonical-46d88860b17b25de8ed0e4000b204eded3e48aa3214aeb5536f6b19c48da2576) |
| `infra.hw_info.usb.product_name` | [infra.hw_info.usb.product_name](resources--registration--reference--group-001.md#canonical-8b57dadce7ad01eed939df81b5d12cd95dbb50beefb61a4b4e2ad40ccd5dfd6b) |
| `infra.hw_info.usb.speed` | [infra.hw_info.usb.speed](resources--registration--reference--group-001.md#canonical-4d17fd3861030c9f67233204c24f22b5168b08f56a66e5433bb4dd2464f49466) |
| `infra.hw_info.usb.usb_type` | [infra.hw_info.usb.usb_type](resources--registration--reference--group-001.md#canonical-6ccb44e3e922c0c9df8d0dc1e707b6241ec9e2e17e48738c08977e2e009e70e6) |
| `infra.hw_info.usb.vendor_name` | [infra.hw_info.usb.vendor_name](resources--registration--reference--group-001.md#canonical-2759e2fa04050536f7238e5bcb76fb1d8299d36fcdeb14a23b40b22b75ae0010) |
| `infra.instance_id` | [infra.instance_id](resources--registration--reference--group-001.md#canonical-2fbc5813e1f329641caa0f50a7c7ebc25ce666fe975aa6db103f612bdfcff465) |
| `infra.interfaces` | [infra.interfaces](resources--registration--reference--group-001.md#canonical-bd56d895993bfaf21838b305c335a55132b69a010499b9bf80495ff7976b7048) |
| `infra.internet_proxy` | [infra.internet_proxy](resources--registration--reference--group-001.md#canonical-0286836835e664c81e393c999087fccf39c27a2a48735def7c82fbe51a4cf7e2) |
| `infra.internet_proxy.http_proxy` | [infra.internet_proxy.http_proxy](resources--registration--reference--group-001.md#canonical-fb27176385f54809eb55678eb50ae931ce8930aef9a60534459a9f13ced376f0) |
| `infra.internet_proxy.https_proxy` | [infra.internet_proxy.https_proxy](resources--registration--reference--group-001.md#canonical-a19d1ae17c3187dbf7c123786700cd454da30d6085137aae434cd420c9ae3fad) |
| `infra.internet_proxy.no_proxy` | [infra.internet_proxy.no_proxy](resources--registration--reference--group-001.md#canonical-6b85f9eb9d2e7e9a9f6307cbe4d127b38bd3cf247489e27fae9966b3d42e284f) |
| `infra.internet_proxy.proxy_cacert_url` | [infra.internet_proxy.proxy_cacert_url](resources--registration--reference--group-001.md#canonical-1feb7bda989daee0919016a1548910143b4136b7429afc2dec9bee1351a2fc49) |
| `infra.is_slo_static` | [infra.is_slo_static](resources--registration--reference--group-001.md#canonical-5f496040538208b03cb7022ce8b90830b56ac11e4a9f600a0a6e2cd074cea7a6) |
| `infra.machine_id` | [infra.machine_id](resources--registration--reference--group-001.md#canonical-d0fc0a3361c4daff2da34d8fe7e107c96e54c4010ad0f4857daecbc1f01cadab) |
| `infra.provider_ref` | [infra.provider_ref](resources--registration--reference--group-001.md#canonical-14bad0f0e6b46e419b692f5ca79f925844d96e88ac7b14f63f6fc15244199243) |
| `infra.sw_info` | [infra.sw_info](resources--registration--reference--group-001.md#canonical-17e0ce779c7417a871c90e844aaa36193600486c545ec855189c849e326ec714) |
| `infra.sw_info.sw_version` | [infra.sw_info.sw_version](resources--registration--reference--group-001.md#canonical-c0435a0d1a22bca5ed857369a72e15cf19667a1e40de2d8d809206116e508ebe) |
| `infra.timestamp` | [infra.timestamp](resources--registration--reference--group-001.md#canonical-5357625e40b402e09062b99ace7e1001d0f70d770a7ff95c5db94af41fadcce0) |
| `infra.zone` | [infra.zone](resources--registration--reference--group-001.md#canonical-c253f05172ca04f6a5cb58b233e67fb7fe0999d36323620c591510b43a7e4150) |
| `labels` | [labels](resources--registration--reference--group-001.md#canonical-d36fb260271ce4e102ac7199a920bbbb37197d4c7b7d60479a0a33185896c80f) |
| `name` | [name](resources--registration--reference--group-001.md#canonical-b83d746b3ac0656ca8efaf27839679370c983cf7437ae9f677acab04a5d346d2) |
| `namespace` | [namespace](resources--registration--reference--group-001.md#canonical-b59b9d5fddb10a1821517d7e433bf6c19c34d9540d71cb151c35de9216bf6dea) |
| `passport` | [passport](resources--registration--reference--group-001.md#canonical-f076677f370ad60d9dddd74f1050f6f7d6a31e3fa90b9b90666451f552983320) |
| `passport.cluster_name` | [passport.cluster_name](resources--registration--reference--group-001.md#canonical-b240fe63a307cced4170a9e0347049038f21a24b71a76830f6c1637c2087ffc2) |
| `passport.cluster_size` | [passport.cluster_size](resources--registration--reference--group-001.md#canonical-bbe0a65f851b262179fa1382e6470b4268dc87c0d6d02f20f74a39133ed70bec) |
| `passport.cluster_type` | [passport.cluster_type](resources--registration--reference--group-001.md#canonical-dc85dba647164396c1eaf6b1fa272ced0be422178bfde2f706b613a08345004b) |
| `passport.default_os_version` | [passport.default_os_version](resources--registration--reference--group-001.md#canonical-9c77537c12bb80ff502fc6f2c46c6097ed75f3d0c326beee1dd56b75d1c98006) |
| `passport.default_sw_version` | [passport.default_sw_version](resources--registration--reference--group-001.md#canonical-277aefba9c4ee82470e495a2e0708af8313ba0552d7c7c3a56fd2671fdcf87e8) |
| `passport.latitude` | [passport.latitude](resources--registration--reference--group-001.md#canonical-4a44c90927eba75dd5c7e2bc0639ff8e515d04e37ca073107ef1696884544547) |
| `passport.longitude` | [passport.longitude](resources--registration--reference--group-001.md#canonical-9caba410c68b78db0cf09ff92d67328494b091098bae51ee0e15ab7dad59b0b5) |
| `passport.operating_system_version` | [passport.operating_system_version](resources--registration--reference--group-001.md#canonical-9f34be31e04ae07290faec52abe6bab992699b1b7e30ca5ee39cf94ec6239cba) |
| `passport.private_network_name` | [passport.private_network_name](resources--registration--reference--group-001.md#canonical-84b238010b49ae155fa3d4d0dc1ff592ada41db38d518f1db64c24d7343fc984) |
| `passport.volterra_software_version` | [passport.volterra_software_version](resources--registration--reference--group-001.md#canonical-5151e86391962b756697b0f7a6ec0d79da79bf65951756d68f2f1548f6db70ac) |
| `timeouts` | [timeouts](resources--registration--reference--group-001.md#canonical-3e0eae774c99299344a7ee3643b1023a92a764086da984c5ca32b15449acbc2c) |
| `timeouts.create` | [timeouts.create](resources--registration--reference--group-001.md#canonical-79b98fe3898b519d835f709c9870bdac164da0fa173bfcb9e587ee4d0a2ec223) |
| `timeouts.delete` | [timeouts.delete](resources--registration--reference--group-001.md#canonical-df1328339c334387e8be09b7f1aa90fe6146daa1ed1e3b19a6cb5e496e51a7ff) |
| `timeouts.read` | [timeouts.read](resources--registration--reference--group-001.md#canonical-8ffde01174dbfe4bc41a0394db6fcbd55a8215830904cbe72ff14ac0ede32e91) |
| `timeouts.update` | [timeouts.update](resources--registration--reference--group-001.md#canonical-0c0e77f0f101db141f126cbd074b3a6f1c7d0195567a3d6b617f5c93683435e4) |
| `token` | [token](resources--registration--reference--group-001.md#canonical-afbdaa38ad15c835f94ca34784f82a6deae750944d1b47563d570ee15e474661) |

<a id="canonical-af8b626c1ca96d51621c483b41735e2f5cf57c46bbe070d40c53f0287f2d44b5"></a>

## Next pages — Property reference / ea1062c425d0 / 13

- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [passport](resources--registration--reference--group-001.md#canonical-5847986b1feea2a4410054d5badbb079e13e8d02a85f50e7d1513ded659f8267)
- [timeouts](resources--registration--reference--group-001.md#canonical-e736dfc22ce4ac100b194673d184aade88f5363c160371788e6bc9c276d81652)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd5cee29a2118b454b1da678d14e600ed149abfcd64566798d85e1cebd0f0482"></a>

## infra — infra / eee6906df63a / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- infra

<a id="canonical-3f898b9932f81db1b2749869bdc1a68644589b3305f200adc9f97fb67665c5a7"></a>

Type: `"object"`. single nested block, Optional.

InfraMetadata stores information about instance infrastructure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hostname",
    "interfaces")}
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
infra {
  # Configure direct properties listed below.
}
```

<a id="canonical-20b7313a3c90fc4e5b37e50bf0241b4fa6c8fd838b1004830520447a2d305ab6"></a>

## Direct properties — infra / eee6906df63a / 3

<a id="canonical-53d2c555c02f42ac2acda7312637b32977842d10237d62e5e3f9cf613652fe95"></a>

<a id="canonical-995a7984a2065774131ba183984ee7def2ec74098cc73fac5b459467f28ef360"></a>

## availability_zone property — infra / eee6906df63a / 4

Type: `"string"`. Optional.

Availability Zone is a high-availability offering that protects your applications and data from
datacenter failures.

Upstream description:

An Availability Zone is a high-availability offering that protects your applications and data from
datacenter failures.

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

- [bond_config](resources--registration--reference--group-001.md#canonical-1985ce79edd7168fac907a0482fd85b218cda2c71ee32b9569491dad217e274a): complete subsection reference.

<a id="canonical-9e670446bd86d16a47f75dc82f02843ce2acb3320f2a54444d70cd3a5e4e24e0"></a>

<a id="canonical-989c2425d087f1ec032ad10622226a9cb35673161e6ddd56cb406fa0e8a601f9"></a>

## certified_hw property — infra / eee6906df63a / 5

Type: `"string"`. Optional.

Certified HW name used to map with F5XC certified\_hardware definition.

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

<a id="canonical-a719ce3b07261f50801a7a6c5bac104d2c425905a54fee2e9aa76837e3169d05"></a>

<a id="canonical-0047bb5648c3f902179c3d9371bcf3fc1e10bceb34cba03682106b012e368baf"></a>

## domain property — infra / eee6906df63a / 6

Type: `"string"`. Optional.

Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5
Distributed Cloud.

Upstream description:

Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5
Distributed Cloud.

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

<a id="canonical-6faa1f67dd1f39c3391693d99e5765fc4463b1d4107329bde597f9c19f6db088"></a>

<a id="canonical-fb52b2ae310972aab4bc650bea4d9b9cdb74f98881d3defaafdf49bc4ffe44ef"></a>

## hostname property — infra / eee6906df63a / 7

Type: `"string"`. Optional.

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Upstream description:

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [hugepages](resources--registration--reference--group-001.md#canonical-e140c78431807e189faf0079882140e2953c50e4e423acb2992e525127665b24): complete subsection reference.

- [hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9): complete subsection reference.

<a id="canonical-2fbc5813e1f329641caa0f50a7c7ebc25ce666fe975aa6db103f612bdfcff465"></a>

<a id="canonical-fdad9ce88abcb4d465ef8cddf4efd32f2949368df578b9412d3335ff4d5ccc31"></a>

## instance_id property — infra / eee6906df63a / 8

Type: `"string"`. Optional.

Instance ID (assigned by infrastructure provider).

Upstream description:

Instance ID (assigned by infrastructure provider)

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

- [interfaces](resources--registration--reference--group-001.md#canonical-5d541690a45c7e6d8588920979547f7f93ba22b917ca21d2201a0d90029aca80): complete subsection reference.

- [internet_proxy](resources--registration--reference--group-001.md#canonical-3fb989f79dda1da83e11b9c550bbea6bdc33b362a10f2f8fd41e2c04519b49ab): complete subsection reference.

<a id="canonical-5f496040538208b03cb7022ce8b90830b56ac11e4a9f600a0a6e2cd074cea7a6"></a>

<a id="canonical-4230dac656501c77a9575daaea9e7261bef48650ef3eeefda96c01e60c3f14d9"></a>

## is_slo_static property — infra / eee6906df63a / 9

Type: `"bool"`. Optional.

Is SLO Static. Indicates whether the SLO is static.

Upstream description:

Indicates whether the SLO is static.

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

<a id="canonical-d0fc0a3361c4daff2da34d8fe7e107c96e54c4010ad0f4857daecbc1f01cadab"></a>

<a id="canonical-5041157dae141555d38a1b8da845567ef3f556fe7de37a90daf764c5a57ce4ca"></a>

## machine_id property — infra / eee6906df63a / 10

Type: `"string"`. Optional.

Machine ID - generated by operating system.

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

<a id="canonical-14bad0f0e6b46e419b692f5ca79f925844d96e88ac7b14f63f6fc15244199243"></a>

<a id="canonical-59f5fcaac3689405307ce856deedd57c6cc1277b006a3f60e59186ee659062e4"></a>

## provider_ref property — infra / eee6906df63a / 11

Type: `"string"`. Optional.

\[Enum:
UNKNOWN|AWS|GOOGLE|AZURE|VMWARE|KVM|OTHER|VOLTERRA|IBMCLOUD|UNKNOWN\_K8S|AWS\_K8S|GCP\_K8S|AZURE\_K8S|VMWARE\_K8S|KVM\_K8S|OTHER\_K8S|VOLTERRA\_K8S|IBMCLOUD\_K8S|F5OS|RSERIES|OCI|NUTANIX|OPENSTACK|EQUINIX|OPENSHIFT\_VIRTUALIZATION|KUBERNETES\]
Infrastructure provider enum for registration. It describes where is instance running. Provider was
not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other
provider, which was not identified by system. Possible values are \`UNKNOWN\`, \`AWS\`, \`GOOGLE\`,
\`AZURE\`, \`VMWARE\`, \`KVM\`, \`OTHER\`, \`VOLTERRA\`, \`IBMCLOUD\`, \`UNKNOWN\_K8S\`,
\`AWS\_K8S\`, \`GCP\_K8S\`, \`AZURE\_K8S\`, \`VMWARE\_K8S\`, \`KVM\_K8S\`, \`OTHER\_K8S\`,
\`VOLTERRA\_K8S\`, \`IBMCLOUD\_K8S\`, \`F5OS\`, \`RSERIES\`, \`OCI\`, \`NUTANIX\`, \`OPENSTACK\`,
\`EQUINIX\`, \`OPENSHIFT\_VIRTUALIZATION\`, \`KUBERNETES\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN",
    "AWS",
    "GOOGLE",
    "AZURE",
    "VMWARE",
    "KVM",
    "OTHER",
    "VOLTERRA",
    "IBMCLOUD",
    "UNKNOWN_K8S",
    "AWS_K8S",
    "GCP_K8S",
    "AZURE_K8S",
    "VMWARE_K8S",
    "KVM_K8S",
    "OTHER_K8S",
    "VOLTERRA_K8S",
    "IBMCLOUD_K8S",
    "F5OS",
    "RSERIES",
    "OCI",
    "NUTANIX",
    "OPENSTACK",
    "EQUINIX",
    "OPENSHIFT_VIRTUALIZATION",
    "KUBERNETES"),
}
```

- [sw_info](resources--registration--reference--group-001.md#canonical-36e66a63e4ea15731911a13fb5e36195bfaa706c2dc77f01ed49a134cc25f1a8): complete subsection reference.

<a id="canonical-5357625e40b402e09062b99ace7e1001d0f70d770a7ff95c5db94af41fadcce0"></a>

<a id="canonical-014d32a37e843e66474c6a6ccf8fc259ec1484bf8a38148ee99d27e88ce7efd2"></a>

## timestamp property — infra / eee6906df63a / 12

Type: `"string"`. Optional.

It's used to verify machine have acceptable time difference from server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "temporal",
    "constraintType": "string",
    "deterministic": true,
    "format": "date-time",
    "formatDescription": "ISO 8601 date-time (e.g., 2026-01-19T12:00:00Z)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 20,
    "pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d{3})?Z?$",
    "validation": {
      "standard": "ISO 8601"
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

<a id="canonical-c253f05172ca04f6a5cb58b233e67fb7fe0999d36323620c591510b43a7e4150"></a>

<a id="canonical-633ed85f518cd9a712b7eedfc1c12597a751aed405abd6f11b1a1c9df8a03385"></a>

## zone property — infra / eee6906df63a / 13

Type: `"string"`. Optional.

Instance zone (or region), depends on provider.

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

<a id="canonical-3f50aae54cd76e3154b40929c7d576a8129c7501eb7bbf7c0538fef263819094"></a>

## Next pages — infra / eee6906df63a / 14

- [infra.bond_config](resources--registration--reference--group-001.md#canonical-1985ce79edd7168fac907a0482fd85b218cda2c71ee32b9569491dad217e274a)
- [infra.hugepages](resources--registration--reference--group-001.md#canonical-e140c78431807e189faf0079882140e2953c50e4e423acb2992e525127665b24)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [infra.interfaces](resources--registration--reference--group-001.md#canonical-5d541690a45c7e6d8588920979547f7f93ba22b917ca21d2201a0d90029aca80)
- [infra.internet_proxy](resources--registration--reference--group-001.md#canonical-3fb989f79dda1da83e11b9c550bbea6bdc33b362a10f2f8fd41e2c04519b49ab)
- [infra.sw_info](resources--registration--reference--group-001.md#canonical-36e66a63e4ea15731911a13fb5e36195bfaa706c2dc77f01ed49a134cc25f1a8)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-1985ce79edd7168fac907a0482fd85b218cda2c71ee32b9569491dad217e274a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c75070f40eea980ee95b9aedb557b53b6145e9fb9fb5c69b81de6d7dbe76fea4"></a>

## infra.bond_config — infra.bond_config / 8bc49e26c21c / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- infra.bond_config

<a id="canonical-1711756b2112aee9221eb98c4aa6546f27f71fa762b9ce5b86ff3e31e88a6acc"></a>

Type: `"object"`. single nested block, Optional.

Bond device configuration for VPM registration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces",
    "name")}
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
bond_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-7cdbea4eadcbb9692d1bbadc9f82c706ca8f29307aaa43ce47122f9d5c619f61"></a>

## Direct properties — infra.bond_config / 8bc49e26c21c / 3

<a id="canonical-0f77dc1113bc7124edce076e43d2c8e2f7a611ea759d72fbc2504819d633bc33"></a>

<a id="canonical-7fa784ac8fe8bef1bc8463fadb70a6e2aa73a9f61bae103c13a0128fa202f49b"></a>

## interfaces property — infra.bond_config / 8bc49e26c21c / 4

Type: `["list", "string"]`. Optional.

Member Interfaces. Configuration parameter for interfaces

Upstream description:

Configuration parameter for interfaces

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-866a902897ab3742c37a3f3df1e464e31739ddca0564a696016e504523a8c0d1"></a>

<a id="canonical-fc92ca50b937394048974d856da83bfbfa3497ea89e60d74082305e2bab2ef3d"></a>

## mode property — infra.bond_config / 8bc49e26c21c / 5

Type: `"string"`. Optional.

\[Enum: BOND\_MODE\_UNSPECIFIED|ACTIVE\_BACKUP|LACP\_802\_3AD\] Bonding mode for bond device
configuration Bond mode is not specified Active-backup bond mode (one interface active, others as
backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are
\`BOND\_MODE\_UNSPECIFIED\`, \`ACTIVE\_BACKUP\`, \`LACP\_802\_3AD\`. Defaults to
\`BOND\_MODE\_UNSPECIFIED\`.

Upstream description:

Bonding mode for bond device configuration

Bond mode is not specified Active-backup bond mode (one interface active, others as backup) IEEE
802.3ad Dynamic link aggregation (LACP)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BOND_MODE_UNSPECIFIED",
    "ACTIVE_BACKUP",
    "LACP_802_3AD"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BOND_MODE_UNSPECIFIED",
  "enum": [
    "BOND_MODE_UNSPECIFIED",
    "ACTIVE_BACKUP",
    "LACP_802_3AD"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fe47aefc466f30c79c052396a037e9b5a07d485f1ed04884022d1206790cbc5f"></a>

<a id="canonical-13a2ae4e7c23e4347e2c719fd7e5eca5e8b9c47217ad5cd72f3f40e9cd8a6f82"></a>

## name property — infra.bond_config / 8bc49e26c21c / 6

Type: `"string"`. Optional.

Bond Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
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

<a id="canonical-33a2b8a0e8ad0af05651e28449d4c3be27b749f669417bed1fb411cfffc161ac"></a>

## Next pages — infra.bond_config / 8bc49e26c21c / 7

- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-e140c78431807e189faf0079882140e2953c50e4e423acb2992e525127665b24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bab5797853a3ecbceca60a097162e8e1e46daba2d91979060632c300e01762e0"></a>

## infra.hugepages — infra.hugepages / 187554bbb069 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- infra.hugepages

<a id="canonical-8ec6d9821c5f0649e1cc6dbfa73348d4fec3f14a0361582bf8e841f8f4f780ce"></a>

Type: `"object"`. list nested block, Optional.

Hugepage settings for CE on K8s SMV2 site.

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
hugepages {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ce54c106565b41d5096e7d3a3ea42832723098b951345fa2416a990d10e5a2e"></a>

## Direct properties — infra.hugepages / 187554bbb069 / 3

<a id="canonical-aeaabcea4057b10c331f4b1ebc88196a1c09fa300595b9b9dd87e31915e6f319"></a>

<a id="canonical-9dafb285248aa9f8094bb9ec9add2b4ebcd6397073fe2fda9fe51ab3b23299be"></a>

## free property — infra.hugepages / 187554bbb069 / 4

Type: `"number"`. Optional.

Free Hugepages. Total number of free hugepages present.

Upstream description:

Total number of free hugepages present.

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

<a id="canonical-4fa79b2a657ad4681459ce07a5198533745d957aeda28fcd6c157f9801f20cf0"></a>

<a id="canonical-d2f13f121c439fd8f943b40b0120c99696e20a2468be760d67b5f7c3858365f1"></a>

## page_size property — infra.hugepages / 187554bbb069 / 5

Type: `"number"`. Optional.

Hugepage Size. Size of each hugepage.

Upstream description:

Size of each hugepage.

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

<a id="canonical-46e837a0b16435caae57d5c295bee56939caab6fa5d8e3a489fb30754d65bf9d"></a>

<a id="canonical-69b0fee7492f20d408f6c4acae4520007f1b93596ba09488aa357409351d8933"></a>

## total property — infra.hugepages / 187554bbb069 / 6

Type: `"number"`. Optional.

Total Hugepages. Total number of hugepages present.

Upstream description:

Total number of hugepages present.

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

<a id="canonical-615e840f240a8db06d9f5613b912b21b81e5b99b95bd42c509391c651db6f362"></a>

## Next pages — infra.hugepages / 187554bbb069 / 7

- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1569a8b051b01ff670d33ca8b8092cd53095d1594bbcc6995a89ff5fb38a98b1"></a>

## infra.hw_info — infra.hw_info / 85dd5bbd5f1b / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- infra.hw_info

<a id="canonical-969d0c6fd11fc46c4a44cb1c1d77e5783086f992f53010b54268ac585da5c1a8"></a>

Type: `"object"`. single nested block, Optional.

OsInfo holds information about host OS and HW.

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
hw_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-ab417c8c00148f7cd3fe59312ffae4ee90d1bb08915167d1e4e02dc471cab47d"></a>

## Direct properties — infra.hw_info / 85dd5bbd5f1b / 3

- [bios](resources--registration--reference--group-001.md#canonical-ec9f9abb7f1b748a576de2a844a80c4fbf239e1c19168afce7a3eab823f3b4fb): complete subsection reference.

- [board](resources--registration--reference--group-001.md#canonical-313da29ff10c72526216d797a2fe8b1a268de9921a1a48af46b7aea848073e31): complete subsection reference.

- [chassis](resources--registration--reference--group-001.md#canonical-3bedd7e505a15bd7fad12f01fe6e78b16a666f4d4454b0c056cf18a3ee3508ec): complete subsection reference.

- [cpu](resources--registration--reference--group-001.md#canonical-53e35cd6709cf440ca6e39fd6f13e48a7d867627eac5ac23e0e964bfc148835a): complete subsection reference.

- [gpu](resources--registration--reference--group-001.md#canonical-85fff7bff0fd46e3031bd8c3a7d86ce5a16713901e6fbfaa5fcf2a81cc241174): complete subsection reference.

- [kernel](resources--registration--reference--group-001.md#canonical-d3ef85aa49e3d96fbcb96043cabd4dfad02f2d0e963bc2d229754d01059bfe63): complete subsection reference.

- [memory](resources--registration--reference--group-001.md#canonical-d7f20f3ec103f6aefe8a97108cb53d80439976afa150c733208530a209e07f10): complete subsection reference.

- [network](resources--registration--reference--group-001.md#canonical-ed0c15d762897fd9ba0f1beaea5bda031a285d223270e6c05753458f9af9b041): complete subsection reference.

<a id="canonical-a84d04d419a2dfa1a6cbac2152d60c222ff52eb17d991422a51ae3cd3088edd6"></a>

<a id="canonical-71cc72db71f5268f28d56c0cded9e44872d79e4e02165515f7a6319152fbc925"></a>

## numa_nodes property — infra.hw_info / 85dd5bbd5f1b / 4

Type: `"number"`. Optional.

Non-uniform memory access (NUMA) nodes count.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
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
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0"
  }
}
```

- [os](resources--registration--reference--group-001.md#canonical-6d04f2f6855e80f52a32f04ce3cae37d7d913ca39b6fc1cac07ac7fae794f175): complete subsection reference.

- [product](resources--registration--reference--group-001.md#canonical-d5749798617b5ca83af058243e0d42cfe45df7ed56d10e46615ecf9f2004ba35): complete subsection reference.

- [storage](resources--registration--reference--group-001.md#canonical-f77d981ef10641c80095b7b3fdf86ed91651a28ccf8136da10f8cf0cbe71eb58): complete subsection reference.

- [usb](resources--registration--reference--group-001.md#canonical-9f77fc86b5ad100ec50d3b6f52dbf0d122ee56ea28b91f3c86f25482ce0c2270): complete subsection reference.

<a id="canonical-741165af2c3af2fdf30585322e421835c333e2c9aa15b5c9c0de1e48c27c9678"></a>

## Next pages — infra.hw_info / 85dd5bbd5f1b / 5

- [infra.hw_info.bios](resources--registration--reference--group-001.md#canonical-ec9f9abb7f1b748a576de2a844a80c4fbf239e1c19168afce7a3eab823f3b4fb)
- [infra.hw_info.board](resources--registration--reference--group-001.md#canonical-313da29ff10c72526216d797a2fe8b1a268de9921a1a48af46b7aea848073e31)
- [infra.hw_info.chassis](resources--registration--reference--group-001.md#canonical-3bedd7e505a15bd7fad12f01fe6e78b16a666f4d4454b0c056cf18a3ee3508ec)
- [infra.hw_info.cpu](resources--registration--reference--group-001.md#canonical-53e35cd6709cf440ca6e39fd6f13e48a7d867627eac5ac23e0e964bfc148835a)
- [infra.hw_info.gpu](resources--registration--reference--group-001.md#canonical-85fff7bff0fd46e3031bd8c3a7d86ce5a16713901e6fbfaa5fcf2a81cc241174)
- [infra.hw_info.kernel](resources--registration--reference--group-001.md#canonical-d3ef85aa49e3d96fbcb96043cabd4dfad02f2d0e963bc2d229754d01059bfe63)
- [infra.hw_info.memory](resources--registration--reference--group-001.md#canonical-d7f20f3ec103f6aefe8a97108cb53d80439976afa150c733208530a209e07f10)
- [infra.hw_info.network](resources--registration--reference--group-001.md#canonical-ed0c15d762897fd9ba0f1beaea5bda031a285d223270e6c05753458f9af9b041)
- [infra.hw_info.os](resources--registration--reference--group-001.md#canonical-6d04f2f6855e80f52a32f04ce3cae37d7d913ca39b6fc1cac07ac7fae794f175)
- [infra.hw_info.product](resources--registration--reference--group-001.md#canonical-d5749798617b5ca83af058243e0d42cfe45df7ed56d10e46615ecf9f2004ba35)
- [infra.hw_info.storage](resources--registration--reference--group-001.md#canonical-f77d981ef10641c80095b7b3fdf86ed91651a28ccf8136da10f8cf0cbe71eb58)
- [infra.hw_info.usb](resources--registration--reference--group-001.md#canonical-9f77fc86b5ad100ec50d3b6f52dbf0d122ee56ea28b91f3c86f25482ce0c2270)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-ec9f9abb7f1b748a576de2a844a80c4fbf239e1c19168afce7a3eab823f3b4fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f55c81f80b25b3ba7a73b5a6ea5a36699fd4f14645a047dcdf71de10f7ec422"></a>

## infra.hw_info.bios — infra.hw_info.bios / d718da64f30a / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.bios

<a id="canonical-a684563e7210b9f6bc80d2a69e5e1c403534656e1c148cffd15f99e06cac4bee"></a>

Type: `"object"`. single nested block, Optional.

Bios Data. BIOS information.

Upstream description:

BIOS information.

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
bios {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0da10ff8ce142b77dc8d89721c5f0c7f2795cacb70fbc5430d16d06c1753f1f"></a>

## Direct properties — infra.hw_info.bios / d718da64f30a / 3

<a id="canonical-25379bb9ec7502b7c0c01304a479b5f339b4cdbcf9df5a077bb10dec93650414"></a>

<a id="canonical-d8766c317a95db5bae5a5581afcfe48946f7cf0ce29bf9e44408583145e45fae"></a>

## date property — infra.hw_info.bios / d718da64f30a / 4

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/bios\_date.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(10, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "temporal",
    "constraintType": "string",
    "deterministic": true,
    "format": "date",
    "formatDescription": "ISO 8601 date (e.g., 2026-01-19)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 10,
    "pattern": "^\\d{4}-\\d{2}-\\d{2}$",
    "validation": {
      "standard": "ISO 8601"
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

<a id="canonical-42921eea47ae8aa7133bf84cfec20ea24c8331167a097d7d635c3beee9d0658d"></a>

<a id="canonical-a8205f7a4b9bc5cd06981a1b711215ae644943c9b2415ca9d742aa068edcc1ca"></a>

## vendor property — infra.hw_info.bios / d718da64f30a / 5

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/bios\_vendor.

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

<a id="canonical-8443d2712f9d5ff3abecf203de2c074ab51b1e5890b15d33b08027f544df12d7"></a>

<a id="canonical-4243000283e12009df71f7ee45ea155d0a3c129edfb6cd579e0388c6fe8a8a2a"></a>

## version property — infra.hw_info.bios / d718da64f30a / 6

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/bios\_version.

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

<a id="canonical-f88500e4a2e1d15eee34cc8e28643f17f7fd10f41d9efd60c704621d510f666e"></a>

## Next pages — infra.hw_info.bios / d718da64f30a / 7

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-313da29ff10c72526216d797a2fe8b1a268de9921a1a48af46b7aea848073e31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49593f15997a2388af2d0c381209f3a7a62f66d70a37354dac0fc13722fd6edc"></a>

## infra.hw_info.board — infra.hw_info.board / 56de1bf9f53f / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.board

<a id="canonical-ae3d652e4ed148282c8c8ff6695c1c5600c2a29dfdf86f8fb66930275d1a84ff"></a>

Type: `"object"`. single nested block, Optional.

Board Details. Board information.

Upstream description:

Board information.

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
board {
  # Configure direct properties listed below.
}
```

<a id="canonical-bdb111c88c58d9d49aa679a5ea6ad1c841b254da1030aafe5c9aa7cb68fd54db"></a>

## Direct properties — infra.hw_info.board / 56de1bf9f53f / 3

<a id="canonical-094f44cfe2dc821887d9f67ee31cb2b99a286975f8d3a33a7d42542cbc452c3f"></a>

<a id="canonical-9b8aeaca93a030aaae0d1c6d6d73488f44ef110c00c52dc0a2bfe69924a6569f"></a>

## asset_tag property — infra.hw_info.board / 56de1bf9f53f / 4

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_asset\_tag.

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

<a id="canonical-4a031a904f9694b383787a257e912cac6134637b0f72f520d13a1e2663cdd26f"></a>

<a id="canonical-1a6221794db798bf8bd67a537f4f94fb284ecb290256a854bbdbde0d48d6eeb4"></a>

## name property — infra.hw_info.board / 56de1bf9f53f / 5

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5f090bcd0f07dd0f0fe21ba2da373a99dea69f8d3dd7ef12a4b81127ad818199"></a>

<a id="canonical-c77f50b709b298717c47bbf35357317fdc1e93d939e57bfb78fe13dc39464cbf"></a>

## serial property — infra.hw_info.board / 56de1bf9f53f / 6

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_serial.

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

<a id="canonical-5e8ddcec633570197e14af5ab203902d7857180d2a73a6d846ee65ead31cf9c9"></a>

<a id="canonical-43f59444125ab316e41a9d9b4bf4850e43d185ef35d08a30d0fcb3794907b2d3"></a>

## vendor property — infra.hw_info.board / 56de1bf9f53f / 7

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_vendor.

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

<a id="canonical-ff2d07c13a11d8f7000a6f14dfef1d9dcbbc5b521816c2c0ea75490f0710f2cf"></a>

<a id="canonical-66fd8854bbc719eff8e8ce6948a07f44ede36094aa1a03ba67f5015d39b4481e"></a>

## version property — infra.hw_info.board / 56de1bf9f53f / 8

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_version.

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

<a id="canonical-cf15d96a7fb3c17ae4ecbcd6871dd06fc12f377531c54e32b3aa3813513c8817"></a>

## Next pages — infra.hw_info.board / 56de1bf9f53f / 9

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-3bedd7e505a15bd7fad12f01fe6e78b16a666f4d4454b0c056cf18a3ee3508ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-826ef0e886de4e83b412f276c3dd8cad8eb2e8209f688737aa7db1025426615a"></a>

## infra.hw_info.chassis — infra.hw_info.chassis / 4835347347f6 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.chassis

<a id="canonical-761617314004971994a7e39b29e4e4aa4b3ffce50960a6abe43fab3e0b8ea642"></a>

Type: `"object"`. single nested block, Optional.

Chassis Details. Chassis information.

Upstream description:

Chassis information.

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
chassis {
  # Configure direct properties listed below.
}
```

<a id="canonical-1999a08adc42a107d9698b4e9699827b77625ddb0d9f1fce65a7df835d0285fe"></a>

## Direct properties — infra.hw_info.chassis / 4835347347f6 / 3

<a id="canonical-73eb21b22eafc524b2fefc9ca25fac58e2c70effbf259e87dc8b42784a09e012"></a>

<a id="canonical-5089d8f44e1e33906ebfa845291cad91f0518e2c89eecc1f879f4426b3d97416"></a>

## asset_tag property — infra.hw_info.chassis / 4835347347f6 / 4

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_asset\_tag.

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

<a id="canonical-ee005e8737ddd86b6ece137a9df2a234853d67991b0105b5a6a9e000ad4b23cc"></a>

<a id="canonical-70016df06327a7c0a198a5b96a9438f20cbf2b034976eb3afa1900f6e8d8e132"></a>

## serial property — infra.hw_info.chassis / 4835347347f6 / 5

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_serial.

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

<a id="canonical-843562fc73d14e9f3c7f81594b8cbc491af0ac96b41b0f2a8b062732e0bbe5d6"></a>

<a id="canonical-d9522b5b59c218c8e65aff9a41da81d9dd6a2e7e8f2e2cadb337f906a43eaedb"></a>

## type property — infra.hw_info.chassis / 4835347347f6 / 6

Type: `"number"`. Optional.

Information from /sys/class/dmi/ID/chassis\_type.

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

<a id="canonical-863826996c1e21f9cdf7ed4ac29cba6e08e63cf705fefc0f57d8f594c4a49ca1"></a>

<a id="canonical-cd2f5116d50f201465e5b652a82b9a7a908362da0ed3e246398800a60dad0f78"></a>

## vendor property — infra.hw_info.chassis / 4835347347f6 / 7

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_vendor.

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

<a id="canonical-f8e03429720962bade703cb7de8c09d597ecdbb2ff60bf58244039f1d999eeba"></a>

<a id="canonical-fffc3ccfb5436629bd9eaae5299d7d81ff04818304a8b84c831b7809989c19ec"></a>

## version property — infra.hw_info.chassis / 4835347347f6 / 8

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_version.

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

<a id="canonical-847269466a2c37c08e328481fb89a6f1956469acca2dde452a31a9a1131dcf15"></a>

## Next pages — infra.hw_info.chassis / 4835347347f6 / 9

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-53e35cd6709cf440ca6e39fd6f13e48a7d867627eac5ac23e0e964bfc148835a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a25c79ec459e0835583fa5176e265c647d3ebce8d589d59c6f52a5cab2bd354"></a>

## infra.hw_info.cpu — infra.hw_info.cpu / f8344ba0c71a / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.cpu

<a id="canonical-2cee95b43f9c73eadf0306cbd5dad4dbd33ba67b168d535c502402f7850219de"></a>

Type: `"object"`. single nested block, Optional.

CPU Information. CPU information.

Upstream description:

CPU information.

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
cpu {
  # Configure direct properties listed below.
}
```

<a id="canonical-2e64cb2eccbeda24894a099d4a63942aa056543c15e460a874d04ffa85aa3c7d"></a>

## Direct properties — infra.hw_info.cpu / f8344ba0c71a / 3

<a id="canonical-fa856a116614374a85407517f8ac0d09f35396db1e9e833d4eff89824de98bf2"></a>

<a id="canonical-3dbef16e163432f1cc72d90cefe127a4ee50d099d56bf37c43e74206cdbd0927"></a>

## cache property — infra.hw_info.cpu / f8344ba0c71a / 4

Type: `"number"`. Optional.

Cache. CPU cache size in KB.

Upstream description:

CPU cache size in KB.

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

<a id="canonical-8df80c2a47b1e9e339e2594f622daa9146071236845d5c73d7308f3fa7dff24f"></a>

<a id="canonical-ca8ad8cc263f91da31976c8be7b2fad71b2e76c9dba533452dd65e0d3468e84d"></a>

## cores property — infra.hw_info.cpu / f8344ba0c71a / 5

Type: `"number"`. Optional.

Cores. Number of physical CPU cores.

Upstream description:

Number of physical CPU cores.

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

<a id="canonical-f5869ada211ee5e929b264d501cbf7a49531a84deda358d5299f3142f5d3bbce"></a>

<a id="canonical-c81f6e5d3e146a03843a89f8e2f067e357be768f4d5dbb06f562e8efac46af25"></a>

## cpus property — infra.hw_info.cpu / f8344ba0c71a / 6

Type: `"number"`. Optional.

CPUs. Number of physical CPUs.

Upstream description:

Number of physical CPUs.

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

<a id="canonical-26f3e166c14148ba29f1ed2383c57ba5fa079d7fccf6c45242fb222e93ec0453"></a>

<a id="canonical-b095a2ad371bc0adb17e926f59130c3062f71bdc6791f2a257fec3fe30cc56f0"></a>

## model property — infra.hw_info.cpu / f8344ba0c71a / 7

Type: `"string"`. Optional.

Model. CPU model

Upstream description:

CPU model

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

<a id="canonical-2c1b73df433ffb276b324083ab3decb854118df8682fb4eff1beda35514c7578"></a>

<a id="canonical-f2868c5549086887934d3cf34c961e9699ef0359287fd09cad4f687f128de75d"></a>

## speed property — infra.hw_info.cpu / f8344ba0c71a / 8

Type: `"number"`. Optional.

Speed. CPU clock rate in MHz.

Upstream description:

CPU clock rate in MHz.

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

<a id="canonical-079870743c583c26123381207c6704d83b78af674928fa7839ec0a579d3126d9"></a>

<a id="canonical-9be7968d4bc753293e54927851d0c81ae474d696e8e94716dce02f100391f364"></a>

## threads property — infra.hw_info.cpu / f8344ba0c71a / 9

Type: `"number"`. Optional.

Threads. Number of logical (HT) CPU cores.

Upstream description:

Number of logical (HT) CPU cores.

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

<a id="canonical-5f8edf93828bf92377fbf26e0c2926fdda44e06cb066e1deece99a1ac9983096"></a>

<a id="canonical-34e5a64d815a30382223028301f3d52316070a14a1e16ff67e968c631804fa84"></a>

## vendor property — infra.hw_info.cpu / f8344ba0c71a / 10

Type: `"string"`. Optional.

Vendor. CPU vendor.

Upstream description:

CPU vendor.

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

<a id="canonical-fcbd640055af28ea6a32a0067336362ba76dc27af693064702a42d1397753fc2"></a>

## Next pages — infra.hw_info.cpu / f8344ba0c71a / 11

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-85fff7bff0fd46e3031bd8c3a7d86ce5a16713901e6fbfaa5fcf2a81cc241174"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e152a071976eaf2fcecb605d2a22e7dab5036c91b8ce17ed41c0f4f371235d3e"></a>

## infra.hw_info.gpu — infra.hw_info.gpu / 8b90a62b3999 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.gpu

<a id="canonical-f0146d9a4fe41322e1b40be5e49f9b48d0726beefc00ce7f584c03367857394d"></a>

Type: `"object"`. single nested block, Optional.

GPU. GPU information on server.

Upstream description:

GPU information on server.

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
gpu {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c6ab7e2d59440aa799e37d8082ca4af4749d3ee4dbda1e707b7e07018143190"></a>

## Direct properties — infra.hw_info.gpu / 8b90a62b3999 / 3

<a id="canonical-e23a3568afbc554010628d57a220f00c696ba29fed4e50f95db9d53d51cd4c18"></a>

<a id="canonical-0a5d41573be0099ab5c5155650e4384f931d47920d9b2205f47d4c041addebe7"></a>

## cuda_version property — infra.hw_info.gpu / 8b90a62b3999 / 4

Type: `"string"`. Optional.

Cuda Version. GPU Cuda Version.

Upstream description:

GPU Cuda Version.

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

<a id="canonical-c496b39f550a4c6cead52ecfd4ab1b95bc538464904d2c46a5a45d591ef39f69"></a>

<a id="canonical-2aa1f47e8651d0f9824c9d9f0ddb66b1b60cbde99bc07e834a66e73dc82c7d5f"></a>

## driver_version property — infra.hw_info.gpu / 8b90a62b3999 / 5

Type: `"string"`. Optional.

Driver Version. GPU Driver Version.

Upstream description:

GPU Driver Version.

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

- [gpu_device](resources--registration--reference--group-001.md#canonical-7d7c934191c416204dc605dcc71328deb24f71454c8d13d9e2a2c4bd8d79e48b): complete subsection reference.

<a id="canonical-95eddfb9d9e85e2991c72d0daa153a9007585ae92d0b17224930570fcadde058"></a>

## Next pages — infra.hw_info.gpu / 8b90a62b3999 / 6

- [infra.hw_info.gpu.gpu_device](resources--registration--reference--group-001.md#canonical-7d7c934191c416204dc605dcc71328deb24f71454c8d13d9e2a2c4bd8d79e48b)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-7d7c934191c416204dc605dcc71328deb24f71454c8d13d9e2a2c4bd8d79e48b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ea555fe9dee4acd4d43c3417eeab48c95e78d1ca8b3f2015d9dc98e0e2811b7"></a>

## infra.hw_info.gpu.gpu_device — infra.hw_info.gpu.gpu_device / ab06cc148126 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [infra.hw_info.gpu](resources--registration--reference--group-001.md#canonical-85fff7bff0fd46e3031bd8c3a7d86ce5a16713901e6fbfaa5fcf2a81cc241174)
- infra.hw_info.gpu.gpu_device

<a id="canonical-28e9e0907f217bb3d3734a7a160ee554c3e6150d8040755276baea5f103ef545"></a>

Type: `"object"`. list nested block, Optional.

GPU devices. List of GPU devices in server.

Upstream description:

List of GPU devices in server.

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
gpu_device {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ddd0ca2ab297b5d80b86c06f64418e5c880738aca9e4fe11167e0f4748fdafd"></a>

## Direct properties — infra.hw_info.gpu.gpu_device / ab06cc148126 / 3

<a id="canonical-a6d3b14b2b1ebe153c5398218574b84fa847eb344c6699988ef8f13226f2ac65"></a>

<a id="canonical-1260d2051743515a4e7d10a1cdda9350bbbe56b3af4e0b0a96baa93edcccaee0"></a>

## id property — infra.hw_info.gpu.gpu_device / ab06cc148126 / 4

Type: `"string"`. Optional.

GPU ID. GPU ID

Upstream description:

GPU ID

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

<a id="canonical-c4342dd3d54d0b8bc2735cf15b222f0ffa2be5d44305cbec6ac212256b183edf"></a>

<a id="canonical-f73480aeef28cd7219940cc4bbf4ec4e2164d2b30696474796aa5dcdc3e4cf5a"></a>

## processes property — infra.hw_info.gpu.gpu_device / ab06cc148126 / 5

Type: `"string"`. Optional.

Processes. GPU Processes.

Upstream description:

GPU Processes.

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

<a id="canonical-1f12b1e1db6cb56db26e2793cec9022041f51305031d29353358bb3084f241e4"></a>

<a id="canonical-40a98966b5694071ae515126e31f9762f8e2151214ad52bcff19e2c92d27e51d"></a>

## product_name property — infra.hw_info.gpu.gpu_device / ab06cc148126 / 6

Type: `"string"`. Optional.

Product Name. GPU Product Name.

Upstream description:

GPU Product Name.

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

<a id="canonical-28ed8ec3a53d8fe041dddaddadbdf1f1bdaede8bceba800fb21d52402688ab6a"></a>

## Next pages — infra.hw_info.gpu.gpu_device / ab06cc148126 / 7

- [infra.hw_info.gpu](resources--registration--reference--group-001.md#canonical-85fff7bff0fd46e3031bd8c3a7d86ce5a16713901e6fbfaa5fcf2a81cc241174)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-d3ef85aa49e3d96fbcb96043cabd4dfad02f2d0e963bc2d229754d01059bfe63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c50ef28a1f443942dca2bbc0a6115ae552bda81d6c7a152fbdbaf1a3ee024a34"></a>

## infra.hw_info.kernel — infra.hw_info.kernel / 3842a71e077d / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.kernel

<a id="canonical-6ca1a07d68f8bbc4fec9c951ede7cc6fba9516cf081afc5807accb4c2021a18d"></a>

Type: `"object"`. single nested block, Optional.

Kernel. Kernel information.

Upstream description:

Kernel information.

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
kernel {
  # Configure direct properties listed below.
}
```

<a id="canonical-30d503dd04c6a13ccaad14eb6f83332bcc8a34b11f84a9e40d15570a87285232"></a>

## Direct properties — infra.hw_info.kernel / 3842a71e077d / 3

<a id="canonical-2958f4db98c3aeaeba69af34929dccf4ac30e4c6dbe7a6951e449b53bced5fb8"></a>

<a id="canonical-7df997a7eb7e8645ccfa628ffe6909bc51eb0b52aab692c390c7cb68d09f9d42"></a>

## architecture property — infra.hw_info.kernel / 3842a71e077d / 4

Type: `"string"`. Optional.

Architecture. Kernel architecture.

Upstream description:

Kernel architecture.

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

<a id="canonical-9b0e2c85a807d410b0d67d4059ecc9dc651317fe8deaf956199461ddbe2a0d17"></a>

<a id="canonical-1b7b209c3a9d16ea3d88665277aab7ba94e0df1a94dbab1c9ff8620e926670b0"></a>

## release property — infra.hw_info.kernel / 3842a71e077d / 5

Type: `"string"`. Optional.

Release. Kernel release.

Upstream description:

Kernel release.

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

<a id="canonical-43e7021e32489a51dfeffc01a56697708814b6a90d9318f8171f8e10a16850ef"></a>

<a id="canonical-2d577feefe39b226b3284b489806e5f81b2369d915f3fed07ef561652609f313"></a>

## version property — infra.hw_info.kernel / 3842a71e077d / 6

Type: `"string"`. Optional.

Version. Kernel version.

Upstream description:

Kernel version.

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

<a id="canonical-f8dfba5c42b2a9a5cf9fcb580efad5f7b125840888f56b3e9778d23a011bf6df"></a>

## Next pages — infra.hw_info.kernel / 3842a71e077d / 7

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-d7f20f3ec103f6aefe8a97108cb53d80439976afa150c733208530a209e07f10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b29ee680a65f63982ac06d2b68c7f5045ebf1abebae9cbf98658c10aba537d4"></a>

## infra.hw_info.memory — infra.hw_info.memory / 7f3dea369240 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.memory

<a id="canonical-80d295de4e47f2e88ca4a1c1be6d930b081dd220686e4fab9da48f490837533c"></a>

Type: `"object"`. single nested block, Optional.

Memory Information. Memory information.

Upstream description:

Memory information.

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
memory {
  # Configure direct properties listed below.
}
```

<a id="canonical-b2b1c6d4e20b75ba01d264a56dd8f43f5d4cf6e0f0efe03bc2b3f443f58a9694"></a>

## Direct properties — infra.hw_info.memory / 7f3dea369240 / 3

<a id="canonical-0b63d7e999e402ce8d5bd1b6cb0adf81481d1aeb52cc2a52085b81992bdf3a0c"></a>

<a id="canonical-ef1d74e5787a720a6e9587815240dff5b671822d190cbe587202a745d2579cd1"></a>

## size_mb property — infra.hw_info.memory / 7f3dea369240 / 4

Type: `"number"`. Optional.

RAM. RAM size in MB.

Upstream description:

RAM size in MB.

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

<a id="canonical-d6cc2c7d59262629a302a1234d049c5b07291a1db4790ca1007320fc03a3da03"></a>

<a id="canonical-0d45fbda5881f0fa00ceb942ae1758de70a0b0bf6d506ba52a5867074c9111e8"></a>

## speed property — infra.hw_info.memory / 7f3dea369240 / 5

Type: `"number"`. Optional.

Speed. RAM data rate in MT/s.

Upstream description:

RAM data rate in MT/s.

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

<a id="canonical-9c191138e6abec8896113258f3d6d1f1c583a160311ab209a744b4fbe0bfa509"></a>

<a id="canonical-692929e674da01741c37c75c40661ba7e6bb8ed6675d99c05b2c0a05a1513df6"></a>

## type property — infra.hw_info.memory / 7f3dea369240 / 6

Type: `"string"`. Optional.

Type. Type of memory, eg. DDR4.

Upstream description:

Type of memory, eg. DDR4.

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

<a id="canonical-172af95d339e4e8c7cdeebd002aeff4760decf9c6d20ff186ded3dee247296e3"></a>

## Next pages — infra.hw_info.memory / 7f3dea369240 / 7

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-ed0c15d762897fd9ba0f1beaea5bda031a285d223270e6c05753458f9af9b041"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c35dfa6e75a2e1c8f215a63b8b7b610ff6eb5166141d815ee278f0c46e55097f"></a>

## infra.hw_info.network — infra.hw_info.network / d29d3c4826b3 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.network

<a id="canonical-259012996ccf1e7492bebc548651c3136e5c69c339912ba437631b310eb6acbb"></a>

Type: `"object"`. list nested block, Optional.

Network. List of network devices in server.

Upstream description:

List of network devices in server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "array",
    "maxItems": 32,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
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
network {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5379c6bc6f57387838efeb091bf3bc482e28135bc717dcb7bd0661de19617b3"></a>

## Direct properties — infra.hw_info.network / d29d3c4826b3 / 3

<a id="canonical-e3f6c88cbe2f9cafca506fa6b90c2fa04c8827277f901d9cf689f4490cf7fda9"></a>

<a id="canonical-6213f9c6391b7e4348f6fffb68c69cc7d5ca56ba5345236c71210e96b796dbeb"></a>

## driver property — infra.hw_info.network / d29d3c4826b3 / 4

Type: `"string"`. Optional.

Driver. Driver of device, eg. E1000e.

Upstream description:

Driver of device, eg. E1000e.

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

<a id="canonical-04ee52f2fcbe58c4eea60b4a83ac90fec31b9ae619ad3b02bc94274fd7c96f4c"></a>

<a id="canonical-94c476cc421d712b936617cf4d59f2409d4478469b20cb0a6b4716b5b161116d"></a>

## ip_address property — infra.hw_info.network / d29d3c4826b3 / 5

Type: `["list", "string"]`. Optional.

IP Address. IP address on interface.

Upstream description:

IP address on interface.

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

<a id="canonical-eb3b8d4239cc039d0ea18f3247392780cde30cd6b9e012a12f1c6a627ca89467"></a>

<a id="canonical-8eed494177e25d6a78bef5f13ea7b92e789ea3e5323bde62830a5f89b836f753"></a>

## link_quality property — infra.hw_info.network / d29d3c4826b3 / 6

Type: `"string"`. Optional.

\[Enum: QUALITY\_UNKNOWN|QUALITY\_GOOD|QUALITY\_POOR|QUALITY\_DISABLED\] Link quality determined by
VER using different probes Unknown quality Link quality is good Link quality is poor Quality
disabled. Possible values are \`QUALITY\_UNKNOWN\`, \`QUALITY\_GOOD\`, \`QUALITY\_POOR\`,
\`QUALITY\_DISABLED\`. Defaults to \`QUALITY\_UNKNOWN\`.

Upstream description:

Link quality determined by VER using different probes

Unknown quality Link quality is good Link quality is poor Quality disabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("QUALITY_UNKNOWN",
    "QUALITY_GOOD",
    "QUALITY_POOR",
    "QUALITY_DISABLED"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "QUALITY_UNKNOWN",
  "enum": [
    "QUALITY_UNKNOWN",
    "QUALITY_GOOD",
    "QUALITY_POOR",
    "QUALITY_DISABLED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-37798952b43c5d312ca761d6cb37714358c16d2fb1ed8cb46fb6c1dfe516ea51"></a>

<a id="canonical-ed292b04ea25077d24fb661a1164dc89c648cf099d0df448d69069d6eaf8806b"></a>

## link_type property — infra.hw_info.network / d29d3c4826b3 / 7

Type: `"string"`. Optional.

\[Enum:
LINK\_TYPE\_UNKNOWN|LINK\_TYPE\_ETHERNET|LINK\_TYPE\_WIFI\_802\_11AC|LINK\_TYPE\_WIFI\_802\_11BGN|LINK\_TYPE\_4G|LINK\_TYPE\_WIFI|LINK\_TYPE\_WAN\]
Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of
type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link. Possible values are
\`LINK\_TYPE\_UNKNOWN\`, \`LINK\_TYPE\_ETHERNET\`, \`LINK\_TYPE\_WIFI\_802\_11AC\`,
\`LINK\_TYPE\_WIFI\_802\_11BGN\`, \`LINK\_TYPE\_4G\`, \`LINK\_TYPE\_WIFI\`, \`LINK\_TYPE\_WAN\`.
Defaults to \`LINK\_TYPE\_UNKNOWN\`.

Upstream description:

Link type of interface determined operationally

Link type unknown Link type ethernet Wi-Fi link of type 802.11ac Wi-Fi link of type 802.11bgn Link
type 4G Wi-Fi link Wan link.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("LINK_TYPE_UNKNOWN",
    "LINK_TYPE_ETHERNET",
    "LINK_TYPE_WIFI_802_11AC",
    "LINK_TYPE_WIFI_802_11BGN",
    "LINK_TYPE_4G",
    "LINK_TYPE_WIFI",
    "LINK_TYPE_WAN"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "LINK_TYPE_UNKNOWN",
  "enum": [
    "LINK_TYPE_UNKNOWN",
    "LINK_TYPE_ETHERNET",
    "LINK_TYPE_WIFI_802_11AC",
    "LINK_TYPE_WIFI_802_11BGN",
    "LINK_TYPE_4G",
    "LINK_TYPE_WIFI",
    "LINK_TYPE_WAN"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a3c5d5405c220be760268125f1c48865b6c127f26c27b6a6d7c3a64bea06b8f7"></a>

<a id="canonical-438e7064b9eedea03b64c2994e14d44f8a71af6cb253be74b8f79fe7a7d2ae3f"></a>

## mac_address property — infra.hw_info.network / d29d3c4826b3 / 8

Type: `"string"`. Optional.

MAC Address. MAC address on interface.

Upstream description:

MAC address on interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(17, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`),
    ""),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "formatDescription": "MAC address (e.g., 00:1A:2B:3C:4D:5E)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c0edabbf2b77f7484b70dd56ce3f8ed171b9bac6c3a5cc774f0f347da809a78b"></a>

<a id="canonical-435a75956702276b655510f1acebd7018e9b28f20d6e4a11b04879d5743e540e"></a>

## name property — infra.hw_info.network / d29d3c4826b3 / 9

Type: `"string"`. Optional.

Name. Name of device, eg. Eth0.

Upstream description:

Name of device, eg. Eth0.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e9580e7c2d1448b3b086e2e09039ee370ca479aa902edb7dd33a7e8355b64069"></a>

<a id="canonical-7366ee33c373a4034cbc636fd372a43c25f559c8d6c8858217589eb0dc58ca2e"></a>

## port property — infra.hw_info.network / d29d3c4826b3 / 10

Type: `"string"`. Optional.

Port. Used port, eg. Tp.

Upstream description:

Used port, eg. Tp.

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

<a id="canonical-a360a3275fb960023cdcdb5ae81735b3c6af2a498d697a6aa1b4ec09b7c060b4"></a>

<a id="canonical-254e320595958076cfc31edd96dc8fc3cee5941234534ffbdb106cc65306b682"></a>

## speed property — infra.hw_info.network / d29d3c4826b3 / 11

Type: `"number"`. Optional.

Speed. Device max supported speed in Mbps.

Upstream description:

Device max supported speed in Mbps.

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

<a id="canonical-b3934ed77c4b76cc713ef307a16a491fea5b57786b580951a9fbaecdb312737a"></a>

## Next pages — infra.hw_info.network / d29d3c4826b3 / 12

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-6d04f2f6855e80f52a32f04ce3cae37d7d913ca39b6fc1cac07ac7fae794f175"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aab81f7c2f5314f9bc5bd92fbb0ba51867384d452247785665f737012fab327"></a>

## infra.hw_info.os — infra.hw_info.os / 5a5f4f7dcd46 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.os

<a id="canonical-6d08774c468187ee4765d04f7cd821ac7465f741b2e5434ae4832ea4dd299f72"></a>

Type: `"object"`. single nested block, Optional.

OS. Details of Operating System.

Upstream description:

Details of Operating System.

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
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-dcde7dddf0eded912e1c3f70da6194c7bc9ea1863f209b2e0fa5f80c41d1332a"></a>

## Direct properties — infra.hw_info.os / 5a5f4f7dcd46 / 3

<a id="canonical-4f79c2e92416df8cc51f11d0e6566217199c88d0955d4becf50ee5dd9cf15fdd"></a>

<a id="canonical-2e5621fda9804b35b40968090f7316f25d4da8bc9523c87fdb044934b6a08043"></a>

## architecture property — infra.hw_info.os / 5a5f4f7dcd46 / 4

Type: `"string"`. Optional.

Architecture. Architecture of OS.

Upstream description:

Architecture of OS.

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

<a id="canonical-c2d6f16b01f29d250e0a125424992e357182c189928fbeed4f1b7187fbbeb910"></a>

<a id="canonical-682aa60e97f8de1fdc0fd4fd74b70f4e6e2ac74f9c8306b36764ac688958ce6e"></a>

## name property — infra.hw_info.os / 5a5f4f7dcd46 / 5

Type: `"string"`. Optional.

Name. Name of OS.

Upstream description:

Name of OS.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f394db4b7259f120d148d12054feae00052f04fd3fd5ef2c211ceb1f53c5c00f"></a>

<a id="canonical-1f557a07bf131bb401fc9a6abd921a68c7174546bf8742512cbf529fa844fcf7"></a>

## release property — infra.hw_info.os / 5a5f4f7dcd46 / 6

Type: `"string"`. Optional.

Release. Release of the OS.

Upstream description:

Release of the OS.

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

<a id="canonical-5d53a3e10c6040100e7982140dfaa6249facbc7c648d77585d7d955a17c6430d"></a>

<a id="canonical-1d6893a97a6b2017392578cdf140ae8750d04e975fe6b12f5af8b10b9a511790"></a>

## vendor property — infra.hw_info.os / 5a5f4f7dcd46 / 7

Type: `"string"`. Optional.

Vendor. Vendor of OS.

Upstream description:

Vendor of OS.

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

<a id="canonical-cff9ea6c47b406b1fd4e1c1d42bf37ef76cce3ffbdd6d5eff191e2f166ebbbf9"></a>

<a id="canonical-142570f3bcc20530e65878a9aa570519eff94e8add8d236e2b36aea68fe3521b"></a>

## version property — infra.hw_info.os / 5a5f4f7dcd46 / 8

Type: `"string"`. Optional.

Version. Version of OS.

Upstream description:

Version of OS.

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

<a id="canonical-84551f2847adeb1cc050d5aae6ce306aade50208edd15b90a41ab3ffa28394f8"></a>

## Next pages — infra.hw_info.os / 5a5f4f7dcd46 / 9

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-d5749798617b5ca83af058243e0d42cfe45df7ed56d10e46615ecf9f2004ba35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85a8e0c66430b2fcc85b43b8c1a5c779dfd2b19a44bc45829433f6a4e230d03a"></a>

## infra.hw_info.product — infra.hw_info.product / a67ecf93b8d6 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.product

<a id="canonical-2b2accc84cad6e7ec6bf3759572e6667a1b080926da06de43078eaf3a0f99f41"></a>

Type: `"object"`. single nested block, Optional.

Product Information. Product information.

Upstream description:

Product information.

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
product {
  # Configure direct properties listed below.
}
```

<a id="canonical-b4ed23a1abe2407c64920b8657b384dc4ce23f20d7184872f76e009cd1e3fdda"></a>

## Direct properties — infra.hw_info.product / a67ecf93b8d6 / 3

<a id="canonical-f885f3b0a8998a4526b2fcab6470e42a83722ace1eafc9daccedd1f68f5e46ed"></a>

<a id="canonical-716a7c3bc31cad3185245bc7527a1653987f824893942034cdf8927cafc9a54d"></a>

## name property — infra.hw_info.product / a67ecf93b8d6 / 4

Type: `"string"`. Optional.

Name. Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product\_name.

Upstream description:

Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product\_name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c8983b978e1606813a17f27b8848e15e8d8e66fe1549f5413efaf6fe53ecee54"></a>

<a id="canonical-28d9ea049c73034c587b4ab334a72ec888bb630b9ed4fd9ec9e5173627c07ffb"></a>

## serial property — infra.hw_info.product / a67ecf93b8d6 / 5

Type: `"string"`. Optional.

Serial number, eg. For AWS 00000000-0000-4000-8000-23460f645a1f. Info taken from
/sys/class/dmi/ID/product\_serial.

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

<a id="canonical-82385f3f1b79256b8c5f4ac9fc348426e8ce6bb0d137a06ced1513916d3fdd2d"></a>

<a id="canonical-5e193795cb901dbed144259b0e20afb6191c2d1f4b32b1dcbcebfad2b2718aa3"></a>

## vendor property — infra.hw_info.product / a67ecf93b8d6 / 6

Type: `"string"`. Optional.

Vendor. Vendor name, eg. For AWS Amazon EC2. Info taken from /sys/class/dmi/ID/product\_vendor.

Upstream description:

Vendor name, eg. For AWS Amazon EC2. Info taken from /sys/class/dmi/ID/product\_vendor.

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

<a id="canonical-54f11d6244095f83a9b506db760e36546fe7e6f8ce5df3c0a92fb3326e23e674"></a>

<a id="canonical-d6bfde677b8e4478b96103dfa408aa7cc89f43aee693a318ddf57d50cbe353a9"></a>

## version property — infra.hw_info.product / a67ecf93b8d6 / 7

Type: `"string"`. Optional.

Version name. Info taken from /sys/class/dmi/ID/product\_version.

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

<a id="canonical-aa70768627b4de0a42542969a734a29c2f5b6423818f8109c0d74605f6ad827d"></a>

## Next pages — infra.hw_info.product / a67ecf93b8d6 / 8

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-f77d981ef10641c80095b7b3fdf86ed91651a28ccf8136da10f8cf0cbe71eb58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f1819ec5a2fd3d420dcf10fbc8dcdfabad91a9093523b5817448163bbee1022"></a>

## infra.hw_info.storage — infra.hw_info.storage / e318490885d2 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.storage

<a id="canonical-a40af56489ffe975ec3d30828bff812eeb821258c2aad5f3aea21afdb8c51f69"></a>

Type: `"object"`. list nested block, Optional.

Storage. List of storage devices in server.

Upstream description:

List of storage devices in server.

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
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-dfc9d86041ee8c3ae1e1f70857bc714c46ad82cf475b28fc2873f6417b219b64"></a>

## Direct properties — infra.hw_info.storage / e318490885d2 / 3

<a id="canonical-2a98cd82e733f3c86aa35972ebb81c572fcc25c71df22e4fdeb327150ab8964f"></a>

<a id="canonical-a9ec04365f29c26c512ecad63460ccd4bcd618e3bb1ceee95cda22d5ee6533a3"></a>

## driver property — infra.hw_info.storage / e318490885d2 / 4

Type: `"string"`. Optional.

Driver. Driver of device.

Upstream description:

Driver of device.

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

<a id="canonical-df90c7f9ec70796f413b30849fc9e9e0cbc9743e8f7092a5d6115bb2f0760eb5"></a>

<a id="canonical-2d5d65ce2097d6279dd56ab6fb35bc0e541c13e807db3ea3a19fc2f6ea976a66"></a>

## model property — infra.hw_info.storage / e318490885d2 / 5

Type: `"string"`. Optional.

Model. Model of device.

Upstream description:

Model of device.

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

<a id="canonical-cf835d6f7a5544afb7b4699b7ff641f24403cc81e1dc1259315dd99c3c0a37a6"></a>

<a id="canonical-f843cb57285c29a4f59cca500734a231bd32f1149ac85bca3fc5ff6b7498d80f"></a>

## name property — infra.hw_info.storage / e318490885d2 / 6

Type: `"string"`. Optional.

Name. Name of device, eg. Nvme0n1.

Upstream description:

Name of device, eg. Nvme0n1.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6a1d1336c2f676f14ca2275682d7ab6dc5564db330b4303453dc103fe50a7835"></a>

<a id="canonical-f0e0ab7500f8a7ba9449326d4b3e72f902358ae0fbbba64d023a53c2b809cf81"></a>

## serial property — infra.hw_info.storage / e318490885d2 / 7

Type: `"string"`. Optional.

Serial Number. Serial of device.

Upstream description:

Serial of device.

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

<a id="canonical-dbb99b885b642d2bc59bc3dde8839a474ae51a67a4929a2c1571111d015c0845"></a>

<a id="canonical-94908b951aa3cddb38e119b2d9077d3d82756b9f54d6cc1226b527e5901ad0f3"></a>

## size_gb property — infra.hw_info.storage / e318490885d2 / 8

Type: `"number"`. Optional.

Size(GB). Device size in GB.

Upstream description:

Device size in GB.

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

<a id="canonical-91b379cb0b18d996bfbbf7268b2c7326427bc0a43f6830eb1a96032ec36e900e"></a>

<a id="canonical-acddb1dec175e9938aa3d120f5622921f9f3abfc0998bdd4d8117d39a57ed912"></a>

## vendor property — infra.hw_info.storage / e318490885d2 / 9

Type: `"string"`. Optional.

Vendor. Vendor of device.

Upstream description:

Vendor of device.

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

<a id="canonical-77a4d4bd5136804f1777f3baa2d6b75cfbb68aaabf94595a8d96d4088a4421c0"></a>

## Next pages — infra.hw_info.storage / e318490885d2 / 10

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-9f77fc86b5ad100ec50d3b6f52dbf0d122ee56ea28b91f3c86f25482ce0c2270"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a342e9007553359f8b0c43e611683bc8ca344b1860d674b8b61d8b7101f512f"></a>

## infra.hw_info.usb — infra.hw_info.usb / 0956fd4f4de5 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- infra.hw_info.usb

<a id="canonical-ede6e1b7b67983d962742fa3ee4fd3e664c4e24e61ab66e2a8df392d7ad552b5"></a>

Type: `"object"`. list nested block, Optional.

USB devices. List of USB devices in server.

Upstream description:

List of USB devices in server.

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
usb {
  # Configure direct properties listed below.
}
```

<a id="canonical-71408d60fd98dd4398ed3f14be29e8610ff3a90d16545bd0453a38622e727d47"></a>

## Direct properties — infra.hw_info.usb / 0956fd4f4de5 / 3

<a id="canonical-d71c4af252c647cdf99de2f2fcdac51c15a66be44ddcfa955042aa6a456fa841"></a>

<a id="canonical-3845d02888b2cbad3c423d7ddd79bcd2106f1fb95713cdd61d75ec2968a86ed6"></a>

## address property — infra.hw_info.usb / 0956fd4f4de5 / 4

Type: `"number"`. Optional.

Address of the device on the bus in decimal.

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

<a id="canonical-2de97a180e7ebe1dba93ffbe294c24788cd8d9a448644db7272eed72bd6bbfe6"></a>

<a id="canonical-5c8b8fdbaeeccdf3604aafccf52daaf4c0c93fb3b550c149fa349da9a1797dd0"></a>

## b_device_class property — infra.hw_info.usb / 0956fd4f4de5 / 5

Type: `"string"`. Optional.

Class. The class of this device.

Upstream description:

The class of this device.

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

<a id="canonical-55352536d2ec1967ea505975145f92a7936d3f8d185838c9138820c58ed2557c"></a>

<a id="canonical-14540b97dd6ca003fbfb0a1789a712699e30be38c4728449dbb902076c1738d0"></a>

## b_device_protocol property — infra.hw_info.usb / 0956fd4f4de5 / 6

Type: `"string"`. Optional.

The protocol (within the sub-class) of this device.

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

<a id="canonical-4c5bbbec7cddef1a76dc9f57c10fa0df8313674a08853e4f51593ed315d05732"></a>

<a id="canonical-c17aac7861a80aff55f3b27c671c0532acd2fae5bcfe6b5a5101d58e1c6c0af8"></a>

## b_device_sub_class property — infra.hw_info.usb / 0956fd4f4de5 / 7

Type: `"string"`. Optional.

The sub-class (within the class) of this device.

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

<a id="canonical-536a16e89eece88227c981448f146d659a50b11ec2209699cb02379443294897"></a>

<a id="canonical-b1478e68b08dc80c51215efc6221a1b0ecb7a9ddfb2200b71ed6dffc59454b7e"></a>

## b_max_packet_size property — infra.hw_info.usb / 0956fd4f4de5 / 8

Type: `"number"`. Optional.

Max packet size. Maximum size of the control transfer.

Upstream description:

Maximum size of the control transfer.

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

<a id="canonical-7ef3aa7ad159ce26f5c62e1abb34ab8c86cda74144ebe1b8e48989563d28a6b4"></a>

<a id="canonical-b0a3b24513cf0a86cccb93d23c5219e31addedbaaf75c5f85ee8d7c4b33fc454"></a>

## bcd_device property — infra.hw_info.usb / 0956fd4f4de5 / 9

Type: `"string"`. Optional.

BCD Device. The device version.

Upstream description:

The device version.

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

<a id="canonical-02f0a3a24710cce172dc4719b49024e51f11dc5acf0c6be56bfc2b4bc3679bf7"></a>

<a id="canonical-e60f95d98e2409a39d229c899bb0e781e25698b3107a67995a7684a70f0e200c"></a>

## bcd_usb property — infra.hw_info.usb / 0956fd4f4de5 / 10

Type: `"string"`. Optional.

BCD Spec. USB Specification Release Number.

Upstream description:

USB Specification Release Number.

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

<a id="canonical-0bb83ab98a35a6d2ee4fc04e04a8be1b9b4aed3f02f9184478943ddddaf81424"></a>

<a id="canonical-c45e44f266c69c26c91b5089085b9a9220203faade98ecd5cb674c075487516e"></a>

## bus property — infra.hw_info.usb / 0956fd4f4de5 / 11

Type: `"number"`. Optional.

The bus on which the device was detected in decimal.

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

<a id="canonical-d273d5afa9746b3f868ad59ab35a5d47805f1891706c03fd8ff1d8c2f151ff0b"></a>

<a id="canonical-71a5a4081fb3ba7b0dfe0d4953fa54d6885a204574f7f6c742bb402c48841d4a"></a>

## description_spec property — infra.hw_info.usb / 0956fd4f4de5 / 12

Type: `"string"`. Optional.

Description. Device description.

<a id="canonical-4da22b3052d7324b129720c3292b52c820c95f54713473668c44411863c19342"></a>

<a id="canonical-1dc7cc7811c164f60db11459ac8be553117d35052fd6ab7fc0a731f211c49301"></a>

## i_manufacturer property — infra.hw_info.usb / 0956fd4f4de5 / 13

Type: `"string"`. Optional.

Manufacturer. Manufacturer name.

Upstream description:

Manufacturer name.

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

<a id="canonical-a0ce31e9c5acc232f1b470f087a317654dc74f2d923f70961e6f40dfdf222947"></a>

<a id="canonical-9f733e0af62edee67252072c5f985f1c540d103b25628f56f3137d0a57bad54c"></a>

## i_product property — infra.hw_info.usb / 0956fd4f4de5 / 14

Type: `"string"`. Optional.

Device product. Product name reported by device.

Upstream description:

Product name reported by device.

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

<a id="canonical-58ae3ba085bff88c0b3de42fd7e7c7b86335c9a6176e91e1bd6cd74b17404ad6"></a>

<a id="canonical-ac5d13e915801c1533e8e19a2baea43d2492031748cf45480093e12fdf8c12e0"></a>

## i_serial property — infra.hw_info.usb / 0956fd4f4de5 / 15

Type: `"string"`. Optional.

Index of Serial Number String Descriptor.

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

<a id="canonical-2a05762248b1aa90b11eb444561f3277707061ae364de092ea1444100fc142a1"></a>

<a id="canonical-43e37860a8c39f40d6c8d4b749f4d86ffbee527b683feec5b3912d2c84d88043"></a>

## id_product property — infra.hw_info.usb / 0956fd4f4de5 / 16

Type: `"string"`. Optional.

Product ID (Assigned by Manufacturer) in hex.

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

<a id="canonical-5d20199e109fb0a6fac44bfbefff9fd1cf523e630796db9931507ea730b05417"></a>

<a id="canonical-a5e2026906ebc2badba15a12befb9018ad449273d52c15c66a0ca1c590f09ed1"></a>

## id_vendor property — infra.hw_info.usb / 0956fd4f4de5 / 17

Type: `"string"`. Optional.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

Upstream description:

Vendor ID (Assigned by USB Org) in hex.

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

<a id="canonical-46d88860b17b25de8ed0e4000b204eded3e48aa3214aeb5536f6b19c48da2576"></a>

<a id="canonical-849ead8d840328e88e76bd35fbd4ed607477b43fe21b225191dff5f4103adb5a"></a>

## port property — infra.hw_info.usb / 0956fd4f4de5 / 18

Type: `"number"`. Optional.

Port on which the device was detected in decimal.

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
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8b57dadce7ad01eed939df81b5d12cd95dbb50beefb61a4b4e2ad40ccd5dfd6b"></a>

<a id="canonical-d19c407de402c02a2c4e84410e0d434a1569c440f82c9d81348b8dd538e2ccda"></a>

## product_name property — infra.hw_info.usb / 0956fd4f4de5 / 19

Type: `"string"`. Optional.

Product ID translated to name (if available).

Upstream description:

Product ID translated to name (if available)

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

<a id="canonical-4d17fd3861030c9f67233204c24f22b5168b08f56a66e5433bb4dd2464f49466"></a>

<a id="canonical-ee5dac0fda6c697a4ea2f38e231a541ccc979dfafd68bf2e39836e55a1f8f115"></a>

## speed property — infra.hw_info.usb / 0956fd4f4de5 / 20

Type: `"string"`. Optional.

The negotiated operating speed for the device.

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

<a id="canonical-6ccb44e3e922c0c9df8d0dc1e707b6241ec9e2e17e48738c08977e2e009e70e6"></a>

<a id="canonical-64e6f113950beea82981db38fd552667e3220829e84c62307be74a2e0c24f65f"></a>

## usb_type property — infra.hw_info.usb / 0956fd4f4de5 / 21

Type: `"string"`. Optional.

\[Enum: UNKNOWN\_USB|INTERNAL|REGISTERED|CONFIGURABLE\] Type of USB device Unknown USB device type
Internal USB present in Certified HW USB device present during node registration USB device that can
be matched by USB rules. Possible values are \`UNKNOWN\_USB\`, \`INTERNAL\`, \`REGISTERED\`,
\`CONFIGURABLE\`. Defaults to \`UNKNOWN\_USB\`.

Upstream description:

Type of USB device

Unknown USB device type Internal USB present in Certified HW USB device present during node
registration USB device that can be matched by USB rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN_USB",
    "INTERNAL",
    "REGISTERED",
    "CONFIGURABLE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNKNOWN_USB",
  "enum": [
    "UNKNOWN_USB",
    "INTERNAL",
    "REGISTERED",
    "CONFIGURABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2759e2fa04050536f7238e5bcb76fb1d8299d36fcdeb14a23b40b22b75ae0010"></a>

<a id="canonical-15a72ba65e4fa66fb1e5cd48d72e65b4db9b647faebbab0c00d5210b046f3208"></a>

## vendor_name property — infra.hw_info.usb / 0956fd4f4de5 / 22

Type: `"string"`. Optional.

Vendor ID translated to name (if available).

Upstream description:

Vendor ID translated to name (if available)

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

<a id="canonical-b3fc20329f69451a5e8d3fd8296886731c70bc7dbda753b89dd6c701ce6c802a"></a>

## Next pages — infra.hw_info.usb / 0956fd4f4de5 / 23

- [infra.hw_info](resources--registration--reference--group-001.md#canonical-784fee08325939c192703d25e7c385cbf4462a658c6b409115028b9f9854a9e9)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-5d541690a45c7e6d8588920979547f7f93ba22b917ca21d2201a0d90029aca80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bcf92fb3623cd60a824d04afd563c9622587a1d90afd51a1c470deef2a75643"></a>

## infra.interfaces — infra.interfaces / 337e84baf786 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- infra.interfaces

<a id="canonical-bd56d895993bfaf21838b305c335a55132b69a010499b9bf80495ff7976b7048"></a>

Type: `"object"`. single nested block, Optional.

Machine interfaces present during registration time.

Receipt-pinned upstream constraints:

```json
{
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

Terraform syntax:

```terraform
interfaces {}
```

<a id="canonical-d7e882c5fa23ae77d3404142e225b31a8b7f4fc1705f195163ed9a041a7bfdab"></a>

## Direct properties — infra.interfaces / 337e84baf786 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1cf7646ac76de7deb581f11947f94338d40a55f9503fb0abf482494330d8aa65"></a>

## Next pages — infra.interfaces / 337e84baf786 / 4

- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-3fb989f79dda1da83e11b9c550bbea6bdc33b362a10f2f8fd41e2c04519b49ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d4609e7d44e0122d59f283e6439072ca1505e77564fc5c60f40107c8662d69d"></a>

## infra.internet_proxy — infra.internet_proxy / a1bbbe73fe75 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- infra.internet_proxy

<a id="canonical-0286836835e664c81e393c999087fccf39c27a2a48735def7c82fbe51a4cf7e2"></a>

Type: `"object"`. single nested block, Optional.

Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.

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
internet_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-e76161ed4ed5ef967cfc8b7f0b8b3282518574577f0b01787cb9482ae4be84e3"></a>

## Direct properties — infra.internet_proxy / a1bbbe73fe75 / 3

<a id="canonical-fb27176385f54809eb55678eb50ae931ce8930aef9a60534459a9f13ced376f0"></a>

<a id="canonical-5d3c96fb47251e5c8e75da4a4861bd579b06296ca21a173f9a99a47c39f478f8"></a>

## http_proxy property — infra.internet_proxy / a1bbbe73fe75 / 4

Type: `"string"`. Optional.

It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by
HTTPSProxy or NoProxy.

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

<a id="canonical-a19d1ae17c3187dbf7c123786700cd454da30d6085137aae434cd420c9ae3fad"></a>

<a id="canonical-c86d63c0b9ebdad5ca5649834f99dc67964350371a933cf8693322c9afb247fc"></a>

## https_proxy property — infra.internet_proxy / a1bbbe73fe75 / 5

Type: `"string"`. Optional.

It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.

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

<a id="canonical-6b85f9eb9d2e7e9a9f6307cbe4d127b38bd3cf247489e27fae9966b3d42e284f"></a>

<a id="canonical-b7b6e416b41b6bbc3b3c7f318f5502dc4972779eff62c11219f926a2503eec9f"></a>

## no_proxy property — infra.internet_proxy / a1bbbe73fe75 / 6

Type: `"string"`. Optional.

It specifies a string that contains comma-separated values specifying hosts that should be excluded
from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix
in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (\*). An IP address..

Upstream description:

It specifies a string that contains comma-separated values specifying hosts that should be excluded
from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix
in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (\*). An IP address prefix
and domain name can also include a literal port number (192.0.2.103:80). A domain name matches that
name and all subdomains. A domain name with a leading "." matches subdomains only. For example
"example.com" matches "example.com" and "bar.example.com"; ".y.com" matches "x.y.com" but not
"y.com". A single asterisk (\*) indicates that no proxying should be done.

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

<a id="canonical-1feb7bda989daee0919016a1548910143b4136b7429afc2dec9bee1351a2fc49"></a>

<a id="canonical-27fa423977f9b20195b72bd4eecf20016631bde6cf4298708f602b6d1d2015e9"></a>

## proxy_cacert_url property — infra.internet_proxy / a1bbbe73fe75 / 7

Type: `"string"`. Optional.

Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate
value.

Upstream description:

Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate
value.

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

<a id="canonical-93860182b9b98bbade9cd86171e0249843d2d50ceedb066a0a66852c0c74fdb5"></a>

## Next pages — infra.internet_proxy / a1bbbe73fe75 / 8

- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-36e66a63e4ea15731911a13fb5e36195bfaa706c2dc77f01ed49a134cc25f1a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-970eab8c670b7b2e15b30a457c0e452147206f765487d547074025511defc5bc"></a>

## infra.sw_info — infra.sw_info / ed3d4b5ac74b / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- infra.sw_info

<a id="canonical-17e0ce779c7417a871c90e844aaa36193600486c545ec855189c849e326ec714"></a>

Type: `"object"`. single nested block, Optional.

SWInfo holds information about sw version.

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
sw_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-d10dcb415190a19fbf24d4214969e38561c22f88c75d291035bde8e8e2b28786"></a>

## Direct properties — infra.sw_info / ed3d4b5ac74b / 3

<a id="canonical-c0435a0d1a22bca5ed857369a72e15cf19667a1e40de2d8d809206116e508ebe"></a>

<a id="canonical-4ab10487e0a60654be30fb8d9cf67e13f9e4b069d253d6ac22dd332064c472bb"></a>

## sw_version property — infra.sw_info / ed3d4b5ac74b / 4

Type: `"string"`. Optional.

SW Version. SW Version in the site.

Upstream description:

SW Version in the site.

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

<a id="canonical-6c21fd4d272e71b97d2fed21c389b1692ed57b6cda7ccc1fe9fbffacd746aeb6"></a>

## Next pages — infra.sw_info / ed3d4b5ac74b / 5

- [infra](resources--registration--reference--group-001.md#canonical-541c32674e5c6df017fc170bfab3b9681a83aa69c8a95ef7c20beefffe4f65d5)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-5847986b1feea2a4410054d5badbb079e13e8d02a85f50e7d1513ded659f8267"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f151f948e37c240f31364c9bb76106b9ed6cdeae3047e87ea427dee4ba9ed959"></a>

## passport — passport / cd9dda162340 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- passport

<a id="canonical-f076677f370ad60d9dddd74f1050f6f7d6a31e3fa90b9b90666451f552983320"></a>

Type: `"object"`. single nested block, Optional.

Passport stores information about identification and node configuration provided by CE during
registration. It can be manually updated by user during approval.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_name",
    "cluster_type",
    "latitude",
    "longitude"),
  validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version"),
  validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]",
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
passport {
  # Configure direct properties listed below.
}
```

<a id="canonical-673a15c8c24246232b1c6d32eeea104fd5ed0ca376d171114a03b13d1be00170"></a>

## Direct properties — passport / cd9dda162340 / 3

<a id="canonical-b240fe63a307cced4170a9e0347049038f21a24b71a76830f6c1637c2087ffc2"></a>

<a id="canonical-a8bd6238254a96e20924694de61e8ec2da85bae1a65ad60a83ab8f1dfd0f1701"></a>

## cluster_name property — passport / cd9dda162340 / 4

Type: `"string"`. Optional.

Cluster Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-bbe0a65f851b262179fa1382e6470b4268dc87c0d6d02f20f74a39133ed70bec"></a>

<a id="canonical-41ee2b6afa6c0a9c9faeae51080d511a1d04f291ea06c88d4c7377f7cee575ef"></a>

## cluster_size property — passport / cd9dda162340 / 5

Type: `"number"`. Optional.

Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single
master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time,
cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after..

Upstream description:

Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single
master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time,
cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after
installation. It does not interact with auto-scaling as only pool nodes are scaled.

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
    "ves.io.schema.rules.int32.in": "[0,1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.in": "[0,1,3]"
  }
}
```

<a id="canonical-dc85dba647164396c1eaf6b1fa272ced0be422178bfde2f706b613a08345004b"></a>

<a id="canonical-894cee1cfbc0663b4de0d9407ed38a64ca6f4423cb6fb19e658971f057966fe6"></a>

## cluster_type property — passport / cd9dda162340 / 6

Type: `"string"`. Optional.

Cluster Type. Cluster or grouping configuration

Upstream description:

Cluster or grouping configuration

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

- [default_os_version](resources--registration--reference--group-001.md#canonical-1947a268a1f969d1acf752ec75ed3e9632a15b2537ddfaec838e05323d0e3380): complete subsection reference.

- [default_sw_version](resources--registration--reference--group-001.md#canonical-6ccabbf5195e7909510d50b05a15b62fbbe883466af385475e1f0cdb7704f1ac): complete subsection reference.

<a id="canonical-4a44c90927eba75dd5c7e2bc0639ff8e515d04e37ca073107ef1696884544547"></a>

<a id="canonical-f01a4680b85424cbe9e0db8655e41037f8fc06ece7fcf4b25994b916d5dfd9d4"></a>

## latitude property — passport / cd9dda162340 / 7

Type: `"number"`. Optional.

Latitude. Geographic location of this site.

Upstream description:

Geographic location of this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-9caba410c68b78db0cf09ff92d67328494b091098bae51ee0e15ab7dad59b0b5"></a>

<a id="canonical-d94b06a164138d2208593d4a5745cee83fc690a251b40e2a985245b9c11c4f7a"></a>

## longitude property — passport / cd9dda162340 / 8

Type: `"number"`. Optional.

Longitude. Geographic location of this site.

Upstream description:

Geographic location of this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-9f34be31e04ae07290faec52abe6bab992699b1b7e30ca5ee39cf94ec6239cba"></a>

<a id="canonical-46861d7b75514acfc315a974dbe9405de3f8d7f3d8f74bda8ae550e75ab8ae92"></a>

## operating_system_version property — passport / cd9dda162340 / 9

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Upstream description:

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-84b238010b49ae155fa3d4d0dc1ff592ada41db38d518f1db64c24d7343fc984"></a>

<a id="canonical-b21b7a70fd3b99f71b1bbea185425e788bf36404be2fbf11f1fcaad643820e33"></a>

## private_network_name property — passport / cd9dda162340 / 10

Type: `"string"`. Optional.

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

Upstream description:

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

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

<a id="canonical-5151e86391962b756697b0f7a6ec0d79da79bf65951756d68f2f1548f6db70ac"></a>

<a id="canonical-7b1e314759f864e2073431a2e25b645c7b0c1a4dfa4b9b10c2ce9946bbbecb69"></a>

## volterra_software_version property — passport / cd9dda162340 / 11

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-796a918fd5bbd09216d89eb08dfcf1f5cf4df211e349a5bdc9fc9b67da1089cd"></a>

## Next pages — passport / cd9dda162340 / 12

- [passport.default_os_version](resources--registration--reference--group-001.md#canonical-1947a268a1f969d1acf752ec75ed3e9632a15b2537ddfaec838e05323d0e3380)
- [passport.default_sw_version](resources--registration--reference--group-001.md#canonical-6ccabbf5195e7909510d50b05a15b62fbbe883466af385475e1f0cdb7704f1ac)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-1947a268a1f969d1acf752ec75ed3e9632a15b2537ddfaec838e05323d0e3380"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b35d2aae788c5ffe7cc6834b8e1ffa2b6261d66cc1f6e632b0baff96ed6f3729"></a>

## passport.default_os_version — passport.default_os_version / 7d99328f9250 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [passport](resources--registration--reference--group-001.md#canonical-5847986b1feea2a4410054d5badbb079e13e8d02a85f50e7d1513ded659f8267)
- passport.default_os_version

<a id="canonical-9c77537c12bb80ff502fc6f2c46c6097ed75f3d0c326beee1dd56b75d1c98006"></a>

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
default_os_version = {}
```

<a id="canonical-9e6d871996f11ed4add669269291ba4b09ee0e7b4c66b0b21b05baab69af477c"></a>

## Direct properties — passport.default_os_version / 7d99328f9250 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ab9fbf890c2c2ab5691fc52f75ab0ad5d1953a36e116eb618db3c54be49e157"></a>

## Next pages — passport.default_os_version / 7d99328f9250 / 4

- [passport](resources--registration--reference--group-001.md#canonical-5847986b1feea2a4410054d5badbb079e13e8d02a85f50e7d1513ded659f8267)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-6ccabbf5195e7909510d50b05a15b62fbbe883466af385475e1f0cdb7704f1ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e578224215f0de11818dcb2c1e01fe60478060b4858ced657f6d0605bb3547b"></a>

## passport.default_sw_version — passport.default_sw_version / ea5b885517dd / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [passport](resources--registration--reference--group-001.md#canonical-5847986b1feea2a4410054d5badbb079e13e8d02a85f50e7d1513ded659f8267)
- passport.default_sw_version

<a id="canonical-277aefba9c4ee82470e495a2e0708af8313ba0552d7c7c3a56fd2671fdcf87e8"></a>

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
default_sw_version = {}
```

<a id="canonical-1e57f8560eec808ab7841edc6e72fe4658a381882766745ad9d4a70020b61586"></a>

## Direct properties — passport.default_sw_version / ea5b885517dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-421688bd6a1fd53774d24a46184469b5a4dee847942b83369a50a2de555e2c1d"></a>

## Next pages — passport.default_sw_version / ea5b885517dd / 4

- [passport](resources--registration--reference--group-001.md#canonical-5847986b1feea2a4410054d5badbb079e13e8d02a85f50e7d1513ded659f8267)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-e736dfc22ce4ac100b194673d184aade88f5363c160371788e6bc9c276d81652"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d756a587cf6c8c8384108a6f8dbb8ad038f5ec4d0a3ae279486e60580e87115"></a>

## timeouts — timeouts / 879cb86b42b6 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- timeouts

<a id="canonical-3e0eae774c99299344a7ee3643b1023a92a764086da984c5ca32b15449acbc2c"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f2fd41f098514fda4dcb43ae2790b162b99787097d4c01c8e609917baa415d9e"></a>

## Direct properties — timeouts / 879cb86b42b6 / 3

<a id="canonical-79b98fe3898b519d835f709c9870bdac164da0fa173bfcb9e587ee4d0a2ec223"></a>

<a id="canonical-57a752ee41916838ef761e207041983cc6bd881fc2fbf6b6a5fbd71841e18166"></a>

## create property — timeouts / 879cb86b42b6 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-df1328339c334387e8be09b7f1aa90fe6146daa1ed1e3b19a6cb5e496e51a7ff"></a>

<a id="canonical-4525222fc00b9ab1930b6ab7bff49173de0a1b973b909753a960b8e45e1280e5"></a>

## delete property — timeouts / 879cb86b42b6 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-8ffde01174dbfe4bc41a0394db6fcbd55a8215830904cbe72ff14ac0ede32e91"></a>

<a id="canonical-5cd43e562f7a789e6c36be1bc2ad1ca1d0f56e763acc5fa7f2a1d499e1476027"></a>

## read property — timeouts / 879cb86b42b6 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0c0e77f0f101db141f126cbd074b3a6f1c7d0195567a3d6b617f5c93683435e4"></a>

<a id="canonical-b279e5de4eb423938c90c9c9fc944710a21e92040ed4ac5258c5ffa1925dbc6c"></a>

## update property — timeouts / 879cb86b42b6 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a02f8b5fc69e1095a3cbbc967d0bc580f0b82a01f0763f067c93b9429dd762f4"></a>

## Next pages — timeouts / 879cb86b42b6 / 8

- [Property reference](resources--registration--reference--group-001.md#canonical-0b136b7db3ab5dafffeec47c165a5aa86d57e234b9a5411c9f9a732d5dfa37c3)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
