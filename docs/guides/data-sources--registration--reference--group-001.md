---
page_title: "xcsh_registration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration reference."
---

# xcsh_registration reference

<a id="canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6e84446356b18b6f6b11637b399f6dc31c0789c53639840c8189f82c796965a"></a>

## Property reference — Property reference / 55a850ce7617 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- Property reference

<a id="canonical-3bbf37b4acc02d800f641ddfe6c95b3c88bf0f16a159a17efd40b67be43a7569"></a>

## Direct properties — Property reference / 55a850ce7617 / 3

<a id="canonical-0d8b034f819323c1a25d2e03ce3f4944bb6ce89f9bc82901bf0a92d528703516"></a>

<a id="canonical-2efbc59ca720bd475adda5af704132341f0e3f3cea6f07f9d0aa11ff4af803cd"></a>

## annotations property — Property reference / 55a850ce7617 / 4

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

<a id="canonical-3a690b4db2bde307334f2ea4d77531e7aafde3d4b5160c09bf9b3a84f873eb92"></a>

<a id="canonical-84d0bebdd1b53cdc3ad3a293bc3f07208ec39558780e85350c3f4069df5d2d76"></a>

## description property — Property reference / 55a850ce7617 / 5

Type: `"string"`. Computed.

Description of the Registration.

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

<a id="canonical-c4f42698bcd6aa1e14c13c4ce356efe99c317d1fd999c2cd88860cfae49a7bb8"></a>

<a id="canonical-35f63ff3ce23655a143365185805848ee13f703ee2c9cf6438a3782ae0144c80"></a>

## id property — Property reference / 55a850ce7617 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251): complete subsection reference.

<a id="canonical-ab21e67e1e3f18f5e85f844e8a974b8379e70eb3148f405b04908965f11c887b"></a>

<a id="canonical-25936e3e8f5440cf2ad6f14ee4b6d2d1050a8ccf4ea887f4c385c65efbc17cab"></a>

## labels property — Property reference / 55a850ce7617 / 7

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

<a id="canonical-86e991b70142cefb27991dec48bea1cdf66cd9c0ef695e87a990f5febb0de47f"></a>

<a id="canonical-c5d591580abfb0851cec9649b9f799196405acb403f7a0dd4f108c33669a996e"></a>

## name property — Property reference / 55a850ce7617 / 8

Type: `"string"`. Required.

Name of the Registration.

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

<a id="canonical-0451ba7180ca4ae64678d6ee9540eeab05c9e74a5f8b314887b21761cd8ccba2"></a>

<a id="canonical-322b1d3fbe7b7ffc911a7cf5a58c32d73118113f83646167c1d05c886f135ae1"></a>

## namespace property — Property reference / 55a850ce7617 / 9

Type: `"string"`. Required.

Namespace where the Registration exists.

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

- [passport](data-sources--registration--reference--group-001.md#canonical-2bb99d2d6317422b03cc3b89a4cc4e42278c4b94e5f3a453f7c72c70228ebba5): complete subsection reference.

<a id="canonical-f2e23b317e188987a2066e88ea128e91802fd38660a572937d855dd2b21c49f5"></a>

<a id="canonical-3b9737d8fac3ea5651ea36a60fd4558a04d24024095f0931aed242a26b6483b8"></a>

## token property — Property reference / 55a850ce7617 / 10

Type: `"string"`. Computed.

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

<a id="canonical-067b6c07e09ed7945d62c41d1d08472a7d243c49ff0b262ec6592745417c41a2"></a>

## All schema paths — Property reference / 55a850ce7617 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--registration--reference--group-001.md#canonical-0d8b034f819323c1a25d2e03ce3f4944bb6ce89f9bc82901bf0a92d528703516) |
| `description` | [description](data-sources--registration--reference--group-001.md#canonical-3a690b4db2bde307334f2ea4d77531e7aafde3d4b5160c09bf9b3a84f873eb92) |
| `id` | [id](data-sources--registration--reference--group-001.md#canonical-c4f42698bcd6aa1e14c13c4ce356efe99c317d1fd999c2cd88860cfae49a7bb8) |
| `infra` | [infra](data-sources--registration--reference--group-001.md#canonical-d6aa00385557f0cabda8d4e6a87649ead136e7d9962a60756cf65628cb25feae) |
| `infra.availability_zone` | [infra.availability_zone](data-sources--registration--reference--group-001.md#canonical-d2339daaf652b79f88f403c3963349828bdf2497f89289a9e8a182d3f1e692af) |
| `infra.bond_config` | [infra.bond_config](data-sources--registration--reference--group-001.md#canonical-856b55a0ac73d1f2bde0447fe5350dd861a7561d57d307bd6b37419cf97c7cb7) |
| `infra.bond_config.interfaces` | [infra.bond_config.interfaces](data-sources--registration--reference--group-001.md#canonical-f74856cc1675346a8da2264d108b29d2d4045b00210826ca709e0376655ef612) |
| `infra.bond_config.mode` | [infra.bond_config.mode](data-sources--registration--reference--group-001.md#canonical-e1b5cb8aa00ed55afd4309b418e3fc3986431a37dc8aa2bf5ba43a98567bc70b) |
| `infra.bond_config.name` | [infra.bond_config.name](data-sources--registration--reference--group-001.md#canonical-b96897b0af62a6c65068f4cb680f56e344ddbd76d90cf6816e90da53b34e20b7) |
| `infra.certified_hw` | [infra.certified_hw](data-sources--registration--reference--group-001.md#canonical-1eed2376602876afb0d3169d40ce65dc64756adf97b190abf8ef17b85856a148) |
| `infra.domain` | [infra.domain](data-sources--registration--reference--group-001.md#canonical-965ebc7b99947f28930165ebf1f6040e9d100acb58b04c087dcc5ac8dd50abad) |
| `infra.hostname` | [infra.hostname](data-sources--registration--reference--group-001.md#canonical-252d1d38e7e98d718c6473b027502dd249ab950294eb28600f4289d0f7120198) |
| `infra.hugepages` | [infra.hugepages](data-sources--registration--reference--group-001.md#canonical-cb87cb5a30dec946d6f992db740a233b00389ee73c11121643751f3857190f12) |
| `infra.hugepages.free` | [infra.hugepages.free](data-sources--registration--reference--group-001.md#canonical-51aa84b7c464d23af72dc85b3261e91cd1210126ca1c56a3e43655b08c8768be) |
| `infra.hugepages.page_size` | [infra.hugepages.page_size](data-sources--registration--reference--group-001.md#canonical-0575c14cd7a7e997bc1e06ce88f4c5fdfaf3dc0f045ecf7e5697fd20d187ae8f) |
| `infra.hugepages.total` | [infra.hugepages.total](data-sources--registration--reference--group-001.md#canonical-fcb41ea3a17b3cbbb2a8306b23867841938e1fc63e585ad1e255235c150a2f85) |
| `infra.hw_info` | [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-b2dfbc22aa5aa18c5b4b1ce255ef13745006c9d9558aaec8f98f49549988119f) |
| `infra.hw_info.bios` | [infra.hw_info.bios](data-sources--registration--reference--group-001.md#canonical-25e2cea9aac7727a1acd5601d41c78c369e928e992ae2f624ebf611d68b5d4da) |
| `infra.hw_info.bios.date` | [infra.hw_info.bios.date](data-sources--registration--reference--group-001.md#canonical-5c00f25eda66394f2953add04a91a8ebdc83181f643a37debf00d38248497250) |
| `infra.hw_info.bios.vendor` | [infra.hw_info.bios.vendor](data-sources--registration--reference--group-001.md#canonical-05c0fae09bfe7492087c3f5d82f5d6932676bf6e57b39d44ad5a73f817a7b753) |
| `infra.hw_info.bios.version` | [infra.hw_info.bios.version](data-sources--registration--reference--group-001.md#canonical-33e62923885b0139969ef55b62acf9fea0757a4a4173587a55ae26e8bbc7ad14) |
| `infra.hw_info.board` | [infra.hw_info.board](data-sources--registration--reference--group-001.md#canonical-92842c11c11e9c61ba4dc29145a8c7021a7da13cf2eeb6749a8d7a9c51ac338d) |
| `infra.hw_info.board.asset_tag` | [infra.hw_info.board.asset_tag](data-sources--registration--reference--group-001.md#canonical-8fa043ea8a73781a5904ee7aded79ce5f8607631256a9d937396bf7b56084088) |
| `infra.hw_info.board.name` | [infra.hw_info.board.name](data-sources--registration--reference--group-001.md#canonical-ae78efaf4464e74d85c9030e5579848bbdcc7137f99c30b66ded0ca3f0e9d215) |
| `infra.hw_info.board.serial` | [infra.hw_info.board.serial](data-sources--registration--reference--group-001.md#canonical-b1e605730ef31538839bd2c2bf34c189b3eb3c02da1c1a5c548d187089719afc) |
| `infra.hw_info.board.vendor` | [infra.hw_info.board.vendor](data-sources--registration--reference--group-001.md#canonical-0303bb642b498d19463d93b5d817572a0ad07a01f67b47cd75389865c4542663) |
| `infra.hw_info.board.version` | [infra.hw_info.board.version](data-sources--registration--reference--group-001.md#canonical-971d1e802caef39fc5013172555cf04c522c0a1dcbacc8ab74202e66c5566319) |
| `infra.hw_info.chassis` | [infra.hw_info.chassis](data-sources--registration--reference--group-001.md#canonical-bd2cea604c3adad46692715e3fb84be8f68b0f820cb69a184c20076584646811) |
| `infra.hw_info.chassis.asset_tag` | [infra.hw_info.chassis.asset_tag](data-sources--registration--reference--group-001.md#canonical-3a2d1d7f8c7a8515774b22d63ca89c836ad9b209f342c7dc66df74e50fac7a01) |
| `infra.hw_info.chassis.serial` | [infra.hw_info.chassis.serial](data-sources--registration--reference--group-001.md#canonical-9c61611eed59edfe15d0f3c7adbd782e48e3187441e1e10b700d8a675862c568) |
| `infra.hw_info.chassis.type` | [infra.hw_info.chassis.type](data-sources--registration--reference--group-001.md#canonical-5012e862b0cc32091b68aa1f32b6748c81e9370905832773eeed2865acc32142) |
| `infra.hw_info.chassis.vendor` | [infra.hw_info.chassis.vendor](data-sources--registration--reference--group-001.md#canonical-b9c497d3c92aa1de78dd58411523c9f4a4952098010d4865c4cefb9059546c55) |
| `infra.hw_info.chassis.version` | [infra.hw_info.chassis.version](data-sources--registration--reference--group-001.md#canonical-238babd77520c0e04cf15dd11c6838304e5e0bd3b49dca57307c3e1b70acf1c7) |
| `infra.hw_info.cpu` | [infra.hw_info.cpu](data-sources--registration--reference--group-001.md#canonical-80098b0065f5f696339b7393ec10eb533aa8cc3fa2f24e0b3f7f5bcf2ccf3007) |
| `infra.hw_info.cpu.cache` | [infra.hw_info.cpu.cache](data-sources--registration--reference--group-001.md#canonical-fcc8e6055acbea8ee12534c0c7f9215c12513e7fd2845586645af89725acd296) |
| `infra.hw_info.cpu.cores` | [infra.hw_info.cpu.cores](data-sources--registration--reference--group-001.md#canonical-362e4e9c3e21b4bf7fb5720222d4285bb7f5f9f1e4c35f5509fdf9dfd38a5704) |
| `infra.hw_info.cpu.cpus` | [infra.hw_info.cpu.cpus](data-sources--registration--reference--group-001.md#canonical-ff2e2bcd64e4b2f543abe0959767a99e0be8b207ec4a2f29952e2528177c0008) |
| `infra.hw_info.cpu.model` | [infra.hw_info.cpu.model](data-sources--registration--reference--group-001.md#canonical-49dc542a6c52186f8383ad511a78ef82f41fac88aad856788ca2a2c93a1196d9) |
| `infra.hw_info.cpu.speed` | [infra.hw_info.cpu.speed](data-sources--registration--reference--group-001.md#canonical-321452b7326f07388ffce1ca4b33766ffa50dabb4e002eeb6ceb36d3f070f153) |
| `infra.hw_info.cpu.threads` | [infra.hw_info.cpu.threads](data-sources--registration--reference--group-001.md#canonical-07fd5607758550961175b3fe368c5581c9b98bd6449bc72f42c60b6ae4822b6e) |
| `infra.hw_info.cpu.vendor` | [infra.hw_info.cpu.vendor](data-sources--registration--reference--group-001.md#canonical-35c56c97eca26f451f70001af450b05d03d21f551e232ed2dc063ed44ed902db) |
| `infra.hw_info.gpu` | [infra.hw_info.gpu](data-sources--registration--reference--group-001.md#canonical-c0b726570e47e146e616a0de6d2ee35390042345c603f10042ac3cd1f8e03a9c) |
| `infra.hw_info.gpu.cuda_version` | [infra.hw_info.gpu.cuda_version](data-sources--registration--reference--group-001.md#canonical-9077d869f49f37bb1ba51a668ca293e8256e72718864fefe034e60fc96fad118) |
| `infra.hw_info.gpu.driver_version` | [infra.hw_info.gpu.driver_version](data-sources--registration--reference--group-001.md#canonical-6a7c2324a7590f6d2b10afbac966023c0a609af2763f3979062e8f713c393b45) |
| `infra.hw_info.gpu.gpu_device` | [infra.hw_info.gpu.gpu_device](data-sources--registration--reference--group-001.md#canonical-b5c4c33c56dd0b7e6208af67ddbbc19a1d9505b0ad4e234aa805088e303ad87f) |
| `infra.hw_info.gpu.gpu_device.id` | [infra.hw_info.gpu.gpu_device.id](data-sources--registration--reference--group-001.md#canonical-b6392d2308c986dd21a1483741e9431340a33d7a84cdeab5b5671faa8acc8603) |
| `infra.hw_info.gpu.gpu_device.processes` | [infra.hw_info.gpu.gpu_device.processes](data-sources--registration--reference--group-001.md#canonical-f60635c2b7b143af084618f97048f4b05eeaf1a9f392bf120baf0119a943a5d1) |
| `infra.hw_info.gpu.gpu_device.product_name` | [infra.hw_info.gpu.gpu_device.product_name](data-sources--registration--reference--group-001.md#canonical-d400a7dfd493c4408106da0872785d7e377c168665491e648f23c8c69900a6d0) |
| `infra.hw_info.kernel` | [infra.hw_info.kernel](data-sources--registration--reference--group-001.md#canonical-01c6feed763af81f81ed315517d22fd0dab1025aae71c58c308018fe97b3dff7) |
| `infra.hw_info.kernel.architecture` | [infra.hw_info.kernel.architecture](data-sources--registration--reference--group-001.md#canonical-0d7b937a84946686f0cb6f25123f6f7e61eae8556644da650ffde38142fda775) |
| `infra.hw_info.kernel.release` | [infra.hw_info.kernel.release](data-sources--registration--reference--group-001.md#canonical-f2eb1bf197e129ec6734dea416c4a9a76596b2274f76fb6e23b050c71e51bc19) |
| `infra.hw_info.kernel.version` | [infra.hw_info.kernel.version](data-sources--registration--reference--group-001.md#canonical-8ca4d13fb5611eb64af3548a815a3bfac95efea8258c947a747844e3b66af260) |
| `infra.hw_info.memory` | [infra.hw_info.memory](data-sources--registration--reference--group-001.md#canonical-9d21295a158b1e882cccaade22deb31b90c4d5949b3ede8506022eeff2cd8013) |
| `infra.hw_info.memory.size_mb` | [infra.hw_info.memory.size_mb](data-sources--registration--reference--group-001.md#canonical-23b9ed00f9b0857d8300b0470ba0d9da9a37217902f9508b9055968c5ebf760d) |
| `infra.hw_info.memory.speed` | [infra.hw_info.memory.speed](data-sources--registration--reference--group-001.md#canonical-dc40017ab764925ace9addedb36de16bf17939b0d1c7d61e18cc286870803af4) |
| `infra.hw_info.memory.type` | [infra.hw_info.memory.type](data-sources--registration--reference--group-001.md#canonical-93e7ebdbd605199e6135215a9484edcdac4615368062cc646564c26172042944) |
| `infra.hw_info.network` | [infra.hw_info.network](data-sources--registration--reference--group-001.md#canonical-5a8f3ea4c380f801a5b8b1e3061c2076523239a4345287142b75aac13ec0d546) |
| `infra.hw_info.network.driver` | [infra.hw_info.network.driver](data-sources--registration--reference--group-001.md#canonical-152120246033b39a842e843ae63de2b10edd2a80f773fe1c04fbcc4d85d70e37) |
| `infra.hw_info.network.ip_address` | [infra.hw_info.network.ip_address](data-sources--registration--reference--group-001.md#canonical-47a74c1555448891ed7634a850db7ed4b3113ff86c3d7b7c8cec27000e1a4887) |
| `infra.hw_info.network.link_quality` | [infra.hw_info.network.link_quality](data-sources--registration--reference--group-001.md#canonical-b300aa4c50290b00ebf108cd134d0b621ec0342dc084585d07cd5245506fe0f3) |
| `infra.hw_info.network.link_type` | [infra.hw_info.network.link_type](data-sources--registration--reference--group-001.md#canonical-2f68f06443b41f97fa8162afb54cf0f1d3def77e58482c340ec7fe46809d96b0) |
| `infra.hw_info.network.mac_address` | [infra.hw_info.network.mac_address](data-sources--registration--reference--group-001.md#canonical-151fa0746dda4bed7c9feb36d62b4072ae9efced55295ac057b146fc46b4a769) |
| `infra.hw_info.network.name` | [infra.hw_info.network.name](data-sources--registration--reference--group-001.md#canonical-a9b0908cee823ffd57db6ce4743d9195b192e793f8baaf87cadae13c2d4556a4) |
| `infra.hw_info.network.port` | [infra.hw_info.network.port](data-sources--registration--reference--group-001.md#canonical-8fac2d641d1e5b4ed61629a7c9dbfba98d356067f9097dce0682489dbad45080) |
| `infra.hw_info.network.speed` | [infra.hw_info.network.speed](data-sources--registration--reference--group-001.md#canonical-7d93d488c4f26ad530d2ca586d592fbd582755a9b0aef674cdaa6c41886d9f83) |
| `infra.hw_info.numa_nodes` | [infra.hw_info.numa_nodes](data-sources--registration--reference--group-001.md#canonical-210a77c6eb3796b1d7bd28a29b6535a5acf40653775af4d49c82cf60f2353c7e) |
| `infra.hw_info.os` | [infra.hw_info.os](data-sources--registration--reference--group-001.md#canonical-d2890a2e03daa3522ac7d7502356bf2972c20128c6a2cf5acbd1e2826e26bdad) |
| `infra.hw_info.os.architecture` | [infra.hw_info.os.architecture](data-sources--registration--reference--group-001.md#canonical-7b79cd287fc76c7917e4414b2638b1e5156ba51c8774da3c29750e651c982bce) |
| `infra.hw_info.os.name` | [infra.hw_info.os.name](data-sources--registration--reference--group-001.md#canonical-ad04f6535672b0fb4fd32dde5f1a02c4ffbd8e4e161db6ef14487722a2ad2f0a) |
| `infra.hw_info.os.release` | [infra.hw_info.os.release](data-sources--registration--reference--group-001.md#canonical-e6e6926d320c95adaa2c83f69d8094c3d49b26c3d96edad7ce82a42db8c1e466) |
| `infra.hw_info.os.vendor` | [infra.hw_info.os.vendor](data-sources--registration--reference--group-001.md#canonical-841c64929fd5968c9845050d3833e8e470a82d8c4a45d94316b957ff544ff4ee) |
| `infra.hw_info.os.version` | [infra.hw_info.os.version](data-sources--registration--reference--group-001.md#canonical-fba7d5af0ef9f7d612b897f9db891b8d2ce406ff9eb7c42dde3b94f29a5b201d) |
| `infra.hw_info.product` | [infra.hw_info.product](data-sources--registration--reference--group-001.md#canonical-01f5ce00011b541227ad6cc0b4767f4347cbbfc629e971f13440d96cc9245661) |
| `infra.hw_info.product.name` | [infra.hw_info.product.name](data-sources--registration--reference--group-001.md#canonical-e01342aad605978eca8f3bcf9fcbb7b3b2fc85028ff3e0e97cbee1b58efb77cb) |
| `infra.hw_info.product.serial` | [infra.hw_info.product.serial](data-sources--registration--reference--group-001.md#canonical-cf7a5cd1ff019a238809bb68179098be186b3fd497f955e23521a4b1d7f33f95) |
| `infra.hw_info.product.vendor` | [infra.hw_info.product.vendor](data-sources--registration--reference--group-001.md#canonical-f493c823b6ad2a390201978488cefa2a3e62964853942d59bb4d73781a80f8e8) |
| `infra.hw_info.product.version` | [infra.hw_info.product.version](data-sources--registration--reference--group-001.md#canonical-f2c6b857bb6b1011ebf24f742898556b7d5d6c57b5a18c71a9e49942eb3b32ae) |
| `infra.hw_info.storage` | [infra.hw_info.storage](data-sources--registration--reference--group-001.md#canonical-13af4265aa1b701288082f18907a01c7e7716c4c008b11a03bb9c1f749d622f5) |
| `infra.hw_info.storage.driver` | [infra.hw_info.storage.driver](data-sources--registration--reference--group-001.md#canonical-9bed739cc0bea1358741b592678e776e48026e5ab93f911a00d90f3b3032742a) |
| `infra.hw_info.storage.model` | [infra.hw_info.storage.model](data-sources--registration--reference--group-001.md#canonical-be99188b2e9e2dd736950b24824554eaf134aad77d76f707f4719b49495aee6d) |
| `infra.hw_info.storage.name` | [infra.hw_info.storage.name](data-sources--registration--reference--group-001.md#canonical-75540bfb81eee9729e2caf3200960f40484310302bbcb18b02dfac0390083a52) |
| `infra.hw_info.storage.serial` | [infra.hw_info.storage.serial](data-sources--registration--reference--group-001.md#canonical-11a4c9023dfd4cafadf40ce965e21bac81e4b02da103d0e756400202d0bb3e42) |
| `infra.hw_info.storage.size_gb` | [infra.hw_info.storage.size_gb](data-sources--registration--reference--group-001.md#canonical-3b9ec393e2f7801aabee1ac914bcebce8b79c8ba839db6ea30a1d44b71df8bd2) |
| `infra.hw_info.storage.vendor` | [infra.hw_info.storage.vendor](data-sources--registration--reference--group-001.md#canonical-bea58b6f2b8dbdd28b3512e14ae8035a4d8cf05f31f53c2952ba7a05ac00d920) |
| `infra.hw_info.usb` | [infra.hw_info.usb](data-sources--registration--reference--group-001.md#canonical-61cecb10c88e6fe786fe81d4d87c18458f820d627857688381e8265c5887c3f2) |
| `infra.hw_info.usb.address` | [infra.hw_info.usb.address](data-sources--registration--reference--group-001.md#canonical-5fdd6997b2b0eac64c008e2a6ba461fd6e7fc9219bb308533a74fe1b81da1171) |
| `infra.hw_info.usb.b_device_class` | [infra.hw_info.usb.b_device_class](data-sources--registration--reference--group-001.md#canonical-96aa255e4448dbc5a1bc85c733b1b8c8b82ed118786318d8fd1af1b1935ba943) |
| `infra.hw_info.usb.b_device_protocol` | [infra.hw_info.usb.b_device_protocol](data-sources--registration--reference--group-001.md#canonical-c913a09218aa68ff12363aacb3361d44e822ec1634b384d0751ff0543f3e61fb) |
| `infra.hw_info.usb.b_device_sub_class` | [infra.hw_info.usb.b_device_sub_class](data-sources--registration--reference--group-001.md#canonical-bdc69450e46b26cb29d3753c1a68afac6a0ecd2a9a02e61a9c6701d2cf8d4036) |
| `infra.hw_info.usb.b_max_packet_size` | [infra.hw_info.usb.b_max_packet_size](data-sources--registration--reference--group-001.md#canonical-306311b6133b1b9dc494ca9838ed4a462ecddab4c977d8ac5831acb52ec14164) |
| `infra.hw_info.usb.bcd_device` | [infra.hw_info.usb.bcd_device](data-sources--registration--reference--group-001.md#canonical-789197a203a2f0ece7b6919432fe6b1c5642e973ef45e14f1e6b399c6bb64f2f) |
| `infra.hw_info.usb.bcd_usb` | [infra.hw_info.usb.bcd_usb](data-sources--registration--reference--group-001.md#canonical-ee1bb48a34b54d63c9c840253d7d430dddec7f2e1c53bfc6272b50f67caedb5b) |
| `infra.hw_info.usb.bus` | [infra.hw_info.usb.bus](data-sources--registration--reference--group-001.md#canonical-64dacdef66b6986fa581a67b34726f243c2ededb686210243fb54eaca1f55c0b) |
| `infra.hw_info.usb.description_spec` | [infra.hw_info.usb.description_spec](data-sources--registration--reference--group-001.md#canonical-6193bbf8d42e11147e5eac2b3881cd93545ed86174514d2b73e3d34d4b5faea6) |
| `infra.hw_info.usb.i_manufacturer` | [infra.hw_info.usb.i_manufacturer](data-sources--registration--reference--group-001.md#canonical-0e94491e5288737232042fc2e8c12f03936502bcbca8c571e811357274e5a9f4) |
| `infra.hw_info.usb.i_product` | [infra.hw_info.usb.i_product](data-sources--registration--reference--group-001.md#canonical-d6f1283505b477f56c38d721c835cc75ec7f476c2e61e138c5da466f967b45d9) |
| `infra.hw_info.usb.i_serial` | [infra.hw_info.usb.i_serial](data-sources--registration--reference--group-001.md#canonical-5fb9e28656ab99422836e876608c051e18cbcf7007818bbec5964e5909851331) |
| `infra.hw_info.usb.id_product` | [infra.hw_info.usb.id_product](data-sources--registration--reference--group-001.md#canonical-4fd647f1b068835ddfe591dc0ec2ea3f39bddc1bf2d827acc047ea5e78ca87ac) |
| `infra.hw_info.usb.id_vendor` | [infra.hw_info.usb.id_vendor](data-sources--registration--reference--group-001.md#canonical-f6651f4d154e60c4d50f54a88b17ebb72245c21365c645fd3caa3f017f700ffe) |
| `infra.hw_info.usb.port` | [infra.hw_info.usb.port](data-sources--registration--reference--group-001.md#canonical-b777ede2990f42c748101d076ec99dc7a524ac2156db51296f55def0db2df339) |
| `infra.hw_info.usb.product_name` | [infra.hw_info.usb.product_name](data-sources--registration--reference--group-001.md#canonical-0919dbc918d4ea98352697b8190479027a59fd0f33498f3facaa9fdeacfbad6a) |
| `infra.hw_info.usb.speed` | [infra.hw_info.usb.speed](data-sources--registration--reference--group-001.md#canonical-a145f97d57f574ea9876d8ec6f899ff1fc963087e5ad80a50ecbc5b534aabb44) |
| `infra.hw_info.usb.usb_type` | [infra.hw_info.usb.usb_type](data-sources--registration--reference--group-001.md#canonical-05390abad848080e7ae2c829f69edd6302ce1605ce3c173e0d3b6c599dae03e8) |
| `infra.hw_info.usb.vendor_name` | [infra.hw_info.usb.vendor_name](data-sources--registration--reference--group-001.md#canonical-ff897000262d59ec8ee0848773ac4c220865541ea07cc59d92d772265ae0cdaf) |
| `infra.instance_id` | [infra.instance_id](data-sources--registration--reference--group-001.md#canonical-0d8985889d9450eed805c9f4283de05aeaa8d0a192f8593256464a82d7570f86) |
| `infra.interfaces` | [infra.interfaces](data-sources--registration--reference--group-001.md#canonical-cc176c11d0ed98583899474f02130f872dd7be318f058aa564620735b54ac56f) |
| `infra.internet_proxy` | [infra.internet_proxy](data-sources--registration--reference--group-001.md#canonical-7049e6f2d7d90e37a3d0812d70ad82ef248843da0d2288cd25d70ad191e61de6) |
| `infra.internet_proxy.http_proxy` | [infra.internet_proxy.http_proxy](data-sources--registration--reference--group-001.md#canonical-34c1a7845752198b925e2fb8e2b36c6fcc09c84b7f407b4ea80c3cbddff7c921) |
| `infra.internet_proxy.https_proxy` | [infra.internet_proxy.https_proxy](data-sources--registration--reference--group-001.md#canonical-12599a84ffb4dbc4f2da290685c68a8f926896113b25f4e783cb8e2623e11926) |
| `infra.internet_proxy.no_proxy` | [infra.internet_proxy.no_proxy](data-sources--registration--reference--group-001.md#canonical-087c6453087d26821db98825a356da83631b4e41cdb6c5711bb10f54bfa33df8) |
| `infra.internet_proxy.proxy_cacert_url` | [infra.internet_proxy.proxy_cacert_url](data-sources--registration--reference--group-001.md#canonical-0dfaf03d195bd17bc39c6318a22d1d173d5c39d98e0be4906bb7e21f1cf70b0e) |
| `infra.is_slo_static` | [infra.is_slo_static](data-sources--registration--reference--group-001.md#canonical-0b4494d0ec753a7782d96ae69d7c84dd3b8c335edcbfbccf21c44f45a9d7a361) |
| `infra.machine_id` | [infra.machine_id](data-sources--registration--reference--group-001.md#canonical-3fc8295a6016ea89a3f3723fec1294bcb258c72705c7d4e9485ed44e311c8bc6) |
| `infra.provider_ref` | [infra.provider_ref](data-sources--registration--reference--group-001.md#canonical-d0d44cb39e7530be55ba8662129c27be1ea9256ce2cf4bc11f1e422795207539) |
| `infra.sw_info` | [infra.sw_info](data-sources--registration--reference--group-001.md#canonical-031d3b5037360d922c55b5daf24317d6e7b1822b4871c64df862c84dd254f96d) |
| `infra.sw_info.sw_version` | [infra.sw_info.sw_version](data-sources--registration--reference--group-001.md#canonical-cfce203206f6e802e1ee679524cdac4667b96aa3784998d58ce5a4ae1ca09e94) |
| `infra.timestamp` | [infra.timestamp](data-sources--registration--reference--group-001.md#canonical-04d75bad395ae5e48575a37d5ae78bee65e851e56f5df7675616a1750817a876) |
| `infra.zone` | [infra.zone](data-sources--registration--reference--group-001.md#canonical-066bf3bf63ea67001fb3ad67ccf5d2e3916b1810722893104c50ba79fbec6e0f) |
| `labels` | [labels](data-sources--registration--reference--group-001.md#canonical-ab21e67e1e3f18f5e85f844e8a974b8379e70eb3148f405b04908965f11c887b) |
| `name` | [name](data-sources--registration--reference--group-001.md#canonical-86e991b70142cefb27991dec48bea1cdf66cd9c0ef695e87a990f5febb0de47f) |
| `namespace` | [namespace](data-sources--registration--reference--group-001.md#canonical-0451ba7180ca4ae64678d6ee9540eeab05c9e74a5f8b314887b21761cd8ccba2) |
| `passport` | [passport](data-sources--registration--reference--group-001.md#canonical-a56dab5f7a353a4caef5e774a989f31214a11e309edf4baa760ad1a85d540da4) |
| `passport.cluster_name` | [passport.cluster_name](data-sources--registration--reference--group-001.md#canonical-50de270cda18adce9a38b906bbc50d1ec284863326768adb35bed7b3d09631fb) |
| `passport.cluster_size` | [passport.cluster_size](data-sources--registration--reference--group-001.md#canonical-042836eafe92a9cf2d08a70e2f69399b6b192556fdbe103abb6f18e4c4a9fce4) |
| `passport.cluster_type` | [passport.cluster_type](data-sources--registration--reference--group-001.md#canonical-6c0e49d49f85cc2f03048bbc054b2ebd422d1f1dfeffa1a5f5d7f30422863bb9) |
| `passport.default_os_version` | [passport.default_os_version](data-sources--registration--reference--group-001.md#canonical-89b724215a31292ee4898fd234e4d0fad59ed2f383310b2eb6173f04c5268f12) |
| `passport.default_sw_version` | [passport.default_sw_version](data-sources--registration--reference--group-001.md#canonical-833fe530a0c541863bd4b92f5a9a98ed988dd4b80bd4c25b2d06460cdcf8c197) |
| `passport.latitude` | [passport.latitude](data-sources--registration--reference--group-001.md#canonical-533bdddd825426963cc0b937e55698abc17753010201b326a527b715680cca31) |
| `passport.longitude` | [passport.longitude](data-sources--registration--reference--group-001.md#canonical-d4126dbdd48ba1e553b18aea9d57cf4dae8b5d1b0018a040091541fefeb000e5) |
| `passport.operating_system_version` | [passport.operating_system_version](data-sources--registration--reference--group-001.md#canonical-a7e92f460ecfc8a59c9798c6233c3f332ac815df1319233a3e3f1225bba1bf06) |
| `passport.private_network_name` | [passport.private_network_name](data-sources--registration--reference--group-001.md#canonical-e5299d5a774c8ee43d037817a69f5af86d0fe7f61a644da4ae51cb46799cd8b0) |
| `passport.volterra_software_version` | [passport.volterra_software_version](data-sources--registration--reference--group-001.md#canonical-55cc1eb104b02d9bf76ea5507241bcaee426d82ec9bce7e0396680a40e926a29) |
| `token` | [token](data-sources--registration--reference--group-001.md#canonical-f2e23b317e188987a2066e88ea128e91802fd38660a572937d855dd2b21c49f5) |

<a id="canonical-cbfeb76a0022d38b3bad9b6745847d730bcd1ce4e8d32ea0ac2fcfd7e5188345"></a>

## Next pages — Property reference / 55a850ce7617 / 12

- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [passport](data-sources--registration--reference--group-001.md#canonical-2bb99d2d6317422b03cc3b89a4cc4e42278c4b94e5f3a453f7c72c70228ebba5)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd844fdb5d1424fbb60cd81592db63901e5ddcd0a23a4c8af5f83c0aac770000"></a>

## infra — infra / f880ec804e55 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- infra

<a id="canonical-d6aa00385557f0cabda8d4e6a87649ead136e7d9962a60756cf65628cb25feae"></a>

Type: `"single"`. Computed.

InfraMetadata stores information about instance infrastructure.

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

<a id="canonical-ac281dababa2eb38b667b7bb225ecf71097a30052abe694f76dcb98164014429"></a>

## Direct properties — infra / f880ec804e55 / 3

<a id="canonical-d2339daaf652b79f88f403c3963349828bdf2497f89289a9e8a182d3f1e692af"></a>

<a id="canonical-5cb71277f450209f87a77e1df312b8f9350c784954924dc72c03de9a5f83fdca"></a>

## availability_zone property — infra / f880ec804e55 / 4

Type: `"string"`. Computed.

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

- [bond_config](data-sources--registration--reference--group-001.md#canonical-a4430190a55aa93a53b54484946dcd3bf4d709a3d804f2afd0a63b9363274b04): complete subsection reference.

<a id="canonical-1eed2376602876afb0d3169d40ce65dc64756adf97b190abf8ef17b85856a148"></a>

<a id="canonical-85b8ec419a663cafa5f7d412491b78e6cd8d0ab75a0b162eee992e020abfc527"></a>

## certified_hw property — infra / f880ec804e55 / 5

Type: `"string"`. Computed.

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

<a id="canonical-965ebc7b99947f28930165ebf1f6040e9d100acb58b04c087dcc5ac8dd50abad"></a>

<a id="canonical-70b45fd0bb15f5d0d42856a8ab10813c5750b71572876d11560bb56c97390050"></a>

## domain property — infra / f880ec804e55 / 6

Type: `"string"`. Computed.

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

<a id="canonical-252d1d38e7e98d718c6473b027502dd249ab950294eb28600f4289d0f7120198"></a>

<a id="canonical-42f3887bbfc1de553cbfdee0fe1a8506c4d0cdd24d2158cd7b4ece42f35309e3"></a>

## hostname property — infra / f880ec804e55 / 7

Type: `"string"`. Computed.

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Upstream description:

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

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

- [hugepages](data-sources--registration--reference--group-001.md#canonical-983535f1dda3f78ee763b6d731cf15b162b6da119cef30d71f0e25891bf7aef0): complete subsection reference.

- [hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4): complete subsection reference.

<a id="canonical-0d8985889d9450eed805c9f4283de05aeaa8d0a192f8593256464a82d7570f86"></a>

<a id="canonical-b968292b425771005f10505436bc8ac2f81e183190f3b5d2d825d0313555d337"></a>

## instance_id property — infra / f880ec804e55 / 8

Type: `"string"`. Computed.

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

- [interfaces](data-sources--registration--reference--group-001.md#canonical-66e79b237f37efd5879541c286d9c104de02734a72db584f56cabcfc76bf017d): complete subsection reference.

- [internet_proxy](data-sources--registration--reference--group-001.md#canonical-8a5def9df187ad5173d54e74d358a2239e37981290fd63c52a9a3057a477131b): complete subsection reference.

<a id="canonical-0b4494d0ec753a7782d96ae69d7c84dd3b8c335edcbfbccf21c44f45a9d7a361"></a>

<a id="canonical-d6cb4235bbb21f6de570d1505c025b35b4f5f3fd770c3c2e3f0f36e277460090"></a>

## is_slo_static property — infra / f880ec804e55 / 9

Type: `"bool"`. Computed.

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

<a id="canonical-3fc8295a6016ea89a3f3723fec1294bcb258c72705c7d4e9485ed44e311c8bc6"></a>

<a id="canonical-280d345721269dd51c5764ce80445cd5a6320c8f7bb9b7396472e33b3b8124b0"></a>

## machine_id property — infra / f880ec804e55 / 10

Type: `"string"`. Computed.

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

<a id="canonical-d0d44cb39e7530be55ba8662129c27be1ea9256ce2cf4bc11f1e422795207539"></a>

<a id="canonical-0969fb3e76123907f832582891738da748b5cb82bcf10b9ac123aba59b8df84a"></a>

## provider_ref property — infra / f880ec804e55 / 11

Type: `"string"`. Computed.

\[Enum:
UNKNOWN|AWS|GOOGLE|AZURE|VMWARE|KVM|OTHER|VOLTERRA|IBMCLOUD|UNKNOWN\_K8S|AWS\_K8S|GCP\_K8S|AZURE\_K8S|VMWARE\_K8S|KVM\_K8S|OTHER\_K8S|VOLTERRA\_K8S|IBMCLOUD\_K8S|F5OS|RSERIES|OCI|NUTANIX|OPENSTACK|EQUINIX|OPENSHIFT\_VIRTUALIZATION|KUBERNETES\]
Infrastructure provider enum for registration. It describes where is instance running. Provider was
not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other
provider, which was not identified by system. Possible values are \`UNKNOWN\`, \`AWS\`, \`GOOGLE\`,
\`AZURE\`, \`VMWARE\`, \`KVM\`, \`OTHER\`, \`VOLTERRA\`, \`IBMCLOUD\`, \`UNKNOWN\_K8S\`,
\`AWS\_K8S\`, \`GCP\_K8S\`, \`AZURE\_K8S\`, \`VMWARE\_K8S\`, \`KVM\_K8S\`, \`OTHER\_K8S\`,
\`VOLTERRA\_K8S\`, \`IBMCLOUD\_K8S\`, \`F5OS\`, \`RSERIES\`, \`OCI\`, \`NUTANIX\`, \`OPENSTACK\`,
\`EQUINIX\`, \`OPENSHIFT\_VIRTUALIZATION\`, \`KUBERNETES\`.

- [sw_info](data-sources--registration--reference--group-001.md#canonical-6dd402088ad909db8bc33ca16006763d6574c8e94ebcf113e42c2bc20d5edbf3): complete subsection reference.

<a id="canonical-04d75bad395ae5e48575a37d5ae78bee65e851e56f5df7675616a1750817a876"></a>

<a id="canonical-d24ff08ca9bff01a1e9ded29f115da7bd0f70a2f44853891c4b6454a7537a31a"></a>

## timestamp property — infra / f880ec804e55 / 12

Type: `"string"`. Computed.

It's used to verify machine have acceptable time difference from server.

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

<a id="canonical-066bf3bf63ea67001fb3ad67ccf5d2e3916b1810722893104c50ba79fbec6e0f"></a>

<a id="canonical-223a92ec1fbf4e628564600ff541402044d91d8fcc86b5c024ee0fc76b89d7b1"></a>

## zone property — infra / f880ec804e55 / 13

Type: `"string"`. Computed.

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

<a id="canonical-7ed45a8aaf21d1ccf6b994b4d377bc11dd7c4ff03dc30d9230c93499b232b529"></a>

## Next pages — infra / f880ec804e55 / 14

- [infra.bond_config](data-sources--registration--reference--group-001.md#canonical-a4430190a55aa93a53b54484946dcd3bf4d709a3d804f2afd0a63b9363274b04)
- [infra.hugepages](data-sources--registration--reference--group-001.md#canonical-983535f1dda3f78ee763b6d731cf15b162b6da119cef30d71f0e25891bf7aef0)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [infra.interfaces](data-sources--registration--reference--group-001.md#canonical-66e79b237f37efd5879541c286d9c104de02734a72db584f56cabcfc76bf017d)
- [infra.internet_proxy](data-sources--registration--reference--group-001.md#canonical-8a5def9df187ad5173d54e74d358a2239e37981290fd63c52a9a3057a477131b)
- [infra.sw_info](data-sources--registration--reference--group-001.md#canonical-6dd402088ad909db8bc33ca16006763d6574c8e94ebcf113e42c2bc20d5edbf3)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-a4430190a55aa93a53b54484946dcd3bf4d709a3d804f2afd0a63b9363274b04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98083107f0a472f6c4cd8d0075e2b2448be4b60004f68a76798a541a21c3a8bd"></a>

## infra.bond_config — infra.bond_config / 8ff69ee600cf / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- infra.bond_config

<a id="canonical-856b55a0ac73d1f2bde0447fe5350dd861a7561d57d307bd6b37419cf97c7cb7"></a>

Type: `"single"`. Computed.

Bond device configuration for VPM registration.

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

<a id="canonical-ef1c8abc88711ea070dde514545493fc73f93dddcd5cabe0379893aab0393c8c"></a>

## Direct properties — infra.bond_config / 8ff69ee600cf / 3

<a id="canonical-f74856cc1675346a8da2264d108b29d2d4045b00210826ca709e0376655ef612"></a>

<a id="canonical-270c1dbf4f9bbd80b0676ceefda0859feeee390541da1193f0e2f6764874bcf4"></a>

## interfaces property — infra.bond_config / 8ff69ee600cf / 4

Type: `["list", "string"]`. Computed.

Member Interfaces. Configuration parameter for interfaces

Upstream description:

Configuration parameter for interfaces

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

<a id="canonical-e1b5cb8aa00ed55afd4309b418e3fc3986431a37dc8aa2bf5ba43a98567bc70b"></a>

<a id="canonical-5ae0a4567e44701f23656b0c6786faf807b47af882f3c2c6c6d0e60277cd71fc"></a>

## mode property — infra.bond_config / 8ff69ee600cf / 5

Type: `"string"`. Computed.

\[Enum: BOND\_MODE\_UNSPECIFIED|ACTIVE\_BACKUP|LACP\_802\_3AD\] Bonding mode for bond device
configuration Bond mode is not specified Active-backup bond mode (one interface active, others as
backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are
\`BOND\_MODE\_UNSPECIFIED\`, \`ACTIVE\_BACKUP\`, \`LACP\_802\_3AD\`. Defaults to
\`BOND\_MODE\_UNSPECIFIED\`.

Upstream description:

Bonding mode for bond device configuration

Bond mode is not specified Active-backup bond mode (one interface active, others as backup) IEEE
802.3ad Dynamic link aggregation (LACP)

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

<a id="canonical-b96897b0af62a6c65068f4cb680f56e344ddbd76d90cf6816e90da53b34e20b7"></a>

<a id="canonical-de7a09ef5a0a3a5f97f0c45cad22bd7e4531ba1f0237c7e5305fa250278b9aef"></a>

## name property — infra.bond_config / 8ff69ee600cf / 6

Type: `"string"`. Computed.

Bond Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-b083dfa93a050adde1cb1f9901c99e9882dfcec2d178070759ebb2fb9ca6ad6b"></a>

## Next pages — infra.bond_config / 8ff69ee600cf / 7

- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-983535f1dda3f78ee763b6d731cf15b162b6da119cef30d71f0e25891bf7aef0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9a6bb724dd915d809b6db69bf760a46fa9008852ea2385f6204da0f590f5f60"></a>

## infra.hugepages — infra.hugepages / bad897618372 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- infra.hugepages

<a id="canonical-cb87cb5a30dec946d6f992db740a233b00389ee73c11121643751f3857190f12"></a>

Type: `"list"`. Computed.

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

<a id="canonical-e460504e8b3bdc91de3187565be1e417385ee7027b0c7479aad45b34d88d6361"></a>

## Direct properties — infra.hugepages / bad897618372 / 3

<a id="canonical-51aa84b7c464d23af72dc85b3261e91cd1210126ca1c56a3e43655b08c8768be"></a>

<a id="canonical-9743c21882170beddce8cef5dadf69e07bc6b840f27566391cf870686ef78b27"></a>

## free property — infra.hugepages / bad897618372 / 4

Type: `"number"`. Computed.

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

<a id="canonical-0575c14cd7a7e997bc1e06ce88f4c5fdfaf3dc0f045ecf7e5697fd20d187ae8f"></a>

<a id="canonical-c6726a8f213bb2a640d11e29e3807a1f213d8f7e6c156068a37e49dfb9cbb8e3"></a>

## page_size property — infra.hugepages / bad897618372 / 5

Type: `"number"`. Computed.

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

<a id="canonical-fcb41ea3a17b3cbbb2a8306b23867841938e1fc63e585ad1e255235c150a2f85"></a>

<a id="canonical-e526fe954a813cfd757a623ad8e3dd86e076eb1538b7ffee87e56c78d44cfdec"></a>

## total property — infra.hugepages / bad897618372 / 6

Type: `"number"`. Computed.

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

<a id="canonical-1c52fa0560a32037531d589c98922753b2935ff87caa51c98bcff7a99637324f"></a>

## Next pages — infra.hugepages / bad897618372 / 7

- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd322441352abbc99857df1c74a03b634b1587ed702915d822efc738f36e7daf"></a>

## infra.hw_info — infra.hw_info / 92143c6871ed / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- infra.hw_info

<a id="canonical-b2dfbc22aa5aa18c5b4b1ce255ef13745006c9d9558aaec8f98f49549988119f"></a>

Type: `"single"`. Computed.

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

<a id="canonical-c44a4c3de71c99f598fd8766396edbcbc019faede531d4fa95a38517b31f285c"></a>

## Direct properties — infra.hw_info / 92143c6871ed / 3

- [bios](data-sources--registration--reference--group-001.md#canonical-0049c2de62c5c2869ec7ecbb337428745bb10a1abdd99c7a03d298a98fb7dc26): complete subsection reference.

- [board](data-sources--registration--reference--group-001.md#canonical-303fe47badcc2391b6a5dbcfa0257ac8b41cb1a6c8c5678cdccc9d978783779e): complete subsection reference.

- [chassis](data-sources--registration--reference--group-001.md#canonical-41d8831cde6b958eecff88a99eee41756b33b0ed02ffc794d2649326394085be): complete subsection reference.

- [cpu](data-sources--registration--reference--group-001.md#canonical-b489cef08fc7374e2103578f6a57c27e3623956e4788b2758c9881abdf190e8b): complete subsection reference.

- [gpu](data-sources--registration--reference--group-001.md#canonical-18b728ff653601cfc02eb75dc3762a3e51502ab8882dc313a79c4849a55f8a34): complete subsection reference.

- [kernel](data-sources--registration--reference--group-001.md#canonical-e9411e953ab23e15aadccc3395f3df6adfb9f5a64b4f185ea2607e2adfab317a): complete subsection reference.

- [memory](data-sources--registration--reference--group-001.md#canonical-3b03cfa26c33c61bb43fdf5473a0e964162fd45461daf983d7069d000a398fc5): complete subsection reference.

- [network](data-sources--registration--reference--group-001.md#canonical-820c5be99108b8f517fbf7b22318d3e3729ace16af7100f3038c1fdd4ae5b065): complete subsection reference.

<a id="canonical-210a77c6eb3796b1d7bd28a29b6535a5acf40653775af4d49c82cf60f2353c7e"></a>

<a id="canonical-90e2a4cb91e8db9b1d34ab31cdc79cc8e162c45d14125b2308804722ec42fcee"></a>

## numa_nodes property — infra.hw_info / 92143c6871ed / 4

Type: `"number"`. Computed.

Non-uniform memory access (NUMA) nodes count.

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

- [os](data-sources--registration--reference--group-001.md#canonical-8bcbb6bac391159990da9202be8fed6f04f430ad1886e198be1cc0cb11bac1fd): complete subsection reference.

- [product](data-sources--registration--reference--group-001.md#canonical-5cd2a5f5f07aec658752a5088ccdd898ebe9fd093a1f2413b6a637b7c0f8a46b): complete subsection reference.

- [storage](data-sources--registration--reference--group-001.md#canonical-db1e154b948c3d7a4bdf736a72b0b6c1dac635f96835094b06de205af71cadbd): complete subsection reference.

- [usb](data-sources--registration--reference--group-001.md#canonical-542fdb8614ec2c371d6ee65c34fdcc803998f066bcbaac2bb79cdee1119d16b9): complete subsection reference.

<a id="canonical-a6ab940df1c3c9a90c2ec1c2ce4c01451761daf4d0c800acdf9bb379a9c2e610"></a>

## Next pages — infra.hw_info / 92143c6871ed / 5

- [infra.hw_info.bios](data-sources--registration--reference--group-001.md#canonical-0049c2de62c5c2869ec7ecbb337428745bb10a1abdd99c7a03d298a98fb7dc26)
- [infra.hw_info.board](data-sources--registration--reference--group-001.md#canonical-303fe47badcc2391b6a5dbcfa0257ac8b41cb1a6c8c5678cdccc9d978783779e)
- [infra.hw_info.chassis](data-sources--registration--reference--group-001.md#canonical-41d8831cde6b958eecff88a99eee41756b33b0ed02ffc794d2649326394085be)
- [infra.hw_info.cpu](data-sources--registration--reference--group-001.md#canonical-b489cef08fc7374e2103578f6a57c27e3623956e4788b2758c9881abdf190e8b)
- [infra.hw_info.gpu](data-sources--registration--reference--group-001.md#canonical-18b728ff653601cfc02eb75dc3762a3e51502ab8882dc313a79c4849a55f8a34)
- [infra.hw_info.kernel](data-sources--registration--reference--group-001.md#canonical-e9411e953ab23e15aadccc3395f3df6adfb9f5a64b4f185ea2607e2adfab317a)
- [infra.hw_info.memory](data-sources--registration--reference--group-001.md#canonical-3b03cfa26c33c61bb43fdf5473a0e964162fd45461daf983d7069d000a398fc5)
- [infra.hw_info.network](data-sources--registration--reference--group-001.md#canonical-820c5be99108b8f517fbf7b22318d3e3729ace16af7100f3038c1fdd4ae5b065)
- [infra.hw_info.os](data-sources--registration--reference--group-001.md#canonical-8bcbb6bac391159990da9202be8fed6f04f430ad1886e198be1cc0cb11bac1fd)
- [infra.hw_info.product](data-sources--registration--reference--group-001.md#canonical-5cd2a5f5f07aec658752a5088ccdd898ebe9fd093a1f2413b6a637b7c0f8a46b)
- [infra.hw_info.storage](data-sources--registration--reference--group-001.md#canonical-db1e154b948c3d7a4bdf736a72b0b6c1dac635f96835094b06de205af71cadbd)
- [infra.hw_info.usb](data-sources--registration--reference--group-001.md#canonical-542fdb8614ec2c371d6ee65c34fdcc803998f066bcbaac2bb79cdee1119d16b9)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-0049c2de62c5c2869ec7ecbb337428745bb10a1abdd99c7a03d298a98fb7dc26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c2c74f1dd7c5f25469137d0fa4b9a508f7d82b3ebc23159e8514e74babc7b9f"></a>

## infra.hw_info.bios — infra.hw_info.bios / 7c2615a9aa27 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.bios

<a id="canonical-25e2cea9aac7727a1acd5601d41c78c369e928e992ae2f624ebf611d68b5d4da"></a>

Type: `"single"`. Computed.

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

<a id="canonical-7bb44bbedc5dac4621082e9e2bc94e2acd16827111c6f502f2d9adbe937733b1"></a>

## Direct properties — infra.hw_info.bios / 7c2615a9aa27 / 3

<a id="canonical-5c00f25eda66394f2953add04a91a8ebdc83181f643a37debf00d38248497250"></a>

<a id="canonical-930e5e54d2ae7cfdd35b95a5d2eb0562b46f5cf568354ff5801b33821ba851bc"></a>

## date property — infra.hw_info.bios / 7c2615a9aa27 / 4

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_date.

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

<a id="canonical-05c0fae09bfe7492087c3f5d82f5d6932676bf6e57b39d44ad5a73f817a7b753"></a>

<a id="canonical-6d481843922edf72f5aad5f60cf46a94b160ff39f863807502a8ac9256db8f59"></a>

## vendor property — infra.hw_info.bios / 7c2615a9aa27 / 5

Type: `"string"`. Computed.

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

<a id="canonical-33e62923885b0139969ef55b62acf9fea0757a4a4173587a55ae26e8bbc7ad14"></a>

<a id="canonical-ca7cfd667ed5b7c6b29ea135e8a11acf0e2eba93b706c092a5d1dfc437b4abe4"></a>

## version property — infra.hw_info.bios / 7c2615a9aa27 / 6

Type: `"string"`. Computed.

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

<a id="canonical-ad4c23b1d8121b854489239347bfd5942c34cdf0f9b4d29bccbc82d551a3fb21"></a>

## Next pages — infra.hw_info.bios / 7c2615a9aa27 / 7

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-303fe47badcc2391b6a5dbcfa0257ac8b41cb1a6c8c5678cdccc9d978783779e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-288ee201a615c17503ba3efae62b2d8c583e3b4c5b3ba0e0a20a102c4c62b55c"></a>

## infra.hw_info.board — infra.hw_info.board / 3744564e58c5 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.board

<a id="canonical-92842c11c11e9c61ba4dc29145a8c7021a7da13cf2eeb6749a8d7a9c51ac338d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-190193c543c8e84f6c675f7d970dff31f478b6232b325aad5eef69768e23ea24"></a>

## Direct properties — infra.hw_info.board / 3744564e58c5 / 3

<a id="canonical-8fa043ea8a73781a5904ee7aded79ce5f8607631256a9d937396bf7b56084088"></a>

<a id="canonical-34b197aecab221bacb7ce1af826b49f3b37d1621ecdb69a0fedffd7356e89cbd"></a>

## asset_tag property — infra.hw_info.board / 3744564e58c5 / 4

Type: `"string"`. Computed.

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

<a id="canonical-ae78efaf4464e74d85c9030e5579848bbdcc7137f99c30b66ded0ca3f0e9d215"></a>

<a id="canonical-70f50d7ca751b90754e391d66f3928c9b06a2a7c3b4138be75255b3e67db5b65"></a>

## name property — infra.hw_info.board / 3744564e58c5 / 5

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_name.

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

<a id="canonical-b1e605730ef31538839bd2c2bf34c189b3eb3c02da1c1a5c548d187089719afc"></a>

<a id="canonical-0a33731bb5a497b1ec736b151cb4c58dc29eba6943936bfb8ea341c4a996b407"></a>

## serial property — infra.hw_info.board / 3744564e58c5 / 6

Type: `"string"`. Computed.

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

<a id="canonical-0303bb642b498d19463d93b5d817572a0ad07a01f67b47cd75389865c4542663"></a>

<a id="canonical-0c7dd016490f9512f79690545926968c2a2e89a6e71c213b9a9df84a72f29373"></a>

## vendor property — infra.hw_info.board / 3744564e58c5 / 7

Type: `"string"`. Computed.

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

<a id="canonical-971d1e802caef39fc5013172555cf04c522c0a1dcbacc8ab74202e66c5566319"></a>

<a id="canonical-22d12b0ae4cb46fc3bab4fdb208d3edfc4a8576dab936d420f51fb14bc214877"></a>

## version property — infra.hw_info.board / 3744564e58c5 / 8

Type: `"string"`. Computed.

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

<a id="canonical-635a2a3d1c5d835097ffbd6fe2ab757a3b472b0cb55fa6f2a3e6522bc9aaf3f6"></a>

## Next pages — infra.hw_info.board / 3744564e58c5 / 9

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-41d8831cde6b958eecff88a99eee41756b33b0ed02ffc794d2649326394085be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4d8ea8df4809e62389d4acd599d84cada9da70fd042322c539cfdbf861ebcdf"></a>

## infra.hw_info.chassis — infra.hw_info.chassis / c334d2c70e62 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.chassis

<a id="canonical-bd2cea604c3adad46692715e3fb84be8f68b0f820cb69a184c20076584646811"></a>

Type: `"single"`. Computed.

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

<a id="canonical-6057008a0649700ae93adb81be3b5b774ecadbc3008a8513981438fe279f977a"></a>

## Direct properties — infra.hw_info.chassis / c334d2c70e62 / 3

<a id="canonical-3a2d1d7f8c7a8515774b22d63ca89c836ad9b209f342c7dc66df74e50fac7a01"></a>

<a id="canonical-5cddcd409a1f41e0a6e65a85cfb2c0c98f41a6a3b5acb096bfce5cde946a2cbb"></a>

## asset_tag property — infra.hw_info.chassis / c334d2c70e62 / 4

Type: `"string"`. Computed.

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

<a id="canonical-9c61611eed59edfe15d0f3c7adbd782e48e3187441e1e10b700d8a675862c568"></a>

<a id="canonical-a07b67f7042bcb9160cf28186687b3560ff13be6ac1f634b5d8a9528dfdef78d"></a>

## serial property — infra.hw_info.chassis / c334d2c70e62 / 5

Type: `"string"`. Computed.

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

<a id="canonical-5012e862b0cc32091b68aa1f32b6748c81e9370905832773eeed2865acc32142"></a>

<a id="canonical-7302a9746e33d6a5d6f303c0e6b330a22620007bb72ffea15d1fbe5d22db473e"></a>

## type property — infra.hw_info.chassis / c334d2c70e62 / 6

Type: `"number"`. Computed.

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

<a id="canonical-b9c497d3c92aa1de78dd58411523c9f4a4952098010d4865c4cefb9059546c55"></a>

<a id="canonical-f830d62e946cf2b19cb1f76de54202b6d8df5e92a1ab6dccb42e8a64e7d07296"></a>

## vendor property — infra.hw_info.chassis / c334d2c70e62 / 7

Type: `"string"`. Computed.

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

<a id="canonical-238babd77520c0e04cf15dd11c6838304e5e0bd3b49dca57307c3e1b70acf1c7"></a>

<a id="canonical-0e680ee3bb0aa3e9bebad98f7ac160915a2f33eea5596a9e0adc890c3f0239bf"></a>

## version property — infra.hw_info.chassis / c334d2c70e62 / 8

Type: `"string"`. Computed.

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

<a id="canonical-3264e4997a0bb643de6988722d56a905b91ade4996833e76284ac5a6aaa79093"></a>

## Next pages — infra.hw_info.chassis / c334d2c70e62 / 9

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-b489cef08fc7374e2103578f6a57c27e3623956e4788b2758c9881abdf190e8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de900fbeddc9a97f87e313e4b35d7f8ff74874e55e84ba89b002f78f1ab64cba"></a>

## infra.hw_info.cpu — infra.hw_info.cpu / ffec3b70ea64 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.cpu

<a id="canonical-80098b0065f5f696339b7393ec10eb533aa8cc3fa2f24e0b3f7f5bcf2ccf3007"></a>

Type: `"single"`. Computed.

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

<a id="canonical-4064102dd8048303bebe1ea6217a0fe5cf17d764c78cf0d42c9ac5d54e2e7f70"></a>

## Direct properties — infra.hw_info.cpu / ffec3b70ea64 / 3

<a id="canonical-fcc8e6055acbea8ee12534c0c7f9215c12513e7fd2845586645af89725acd296"></a>

<a id="canonical-eb575b9abc6ae11ebf73ccedd3a7d113bb93f74edeee67cbbe9b03b861f1c8d4"></a>

## cache property — infra.hw_info.cpu / ffec3b70ea64 / 4

Type: `"number"`. Computed.

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

<a id="canonical-362e4e9c3e21b4bf7fb5720222d4285bb7f5f9f1e4c35f5509fdf9dfd38a5704"></a>

<a id="canonical-534309acde3208fa3d9b8c45f56fcfe4077321784f711d48ab3590edea5ef073"></a>

## cores property — infra.hw_info.cpu / ffec3b70ea64 / 5

Type: `"number"`. Computed.

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

<a id="canonical-ff2e2bcd64e4b2f543abe0959767a99e0be8b207ec4a2f29952e2528177c0008"></a>

<a id="canonical-8c1b0be5aea3239f8199349abe6f633a7c1559a11d4fc21a0a9494c675ade498"></a>

## cpus property — infra.hw_info.cpu / ffec3b70ea64 / 6

Type: `"number"`. Computed.

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

<a id="canonical-49dc542a6c52186f8383ad511a78ef82f41fac88aad856788ca2a2c93a1196d9"></a>

<a id="canonical-d61add795ab4baec00672fac9e4429d460c4e03322758763e2a4714ab2d5c818"></a>

## model property — infra.hw_info.cpu / ffec3b70ea64 / 7

Type: `"string"`. Computed.

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

<a id="canonical-321452b7326f07388ffce1ca4b33766ffa50dabb4e002eeb6ceb36d3f070f153"></a>

<a id="canonical-a342aeb37d918b7936d3623cc1cc574a6138086711deb073ff604b4bb2f50649"></a>

## speed property — infra.hw_info.cpu / ffec3b70ea64 / 8

Type: `"number"`. Computed.

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

<a id="canonical-07fd5607758550961175b3fe368c5581c9b98bd6449bc72f42c60b6ae4822b6e"></a>

<a id="canonical-c1f13c49da0514c11b2ba291c4956f2548a6a92df54917000e33329fc9e07d83"></a>

## threads property — infra.hw_info.cpu / ffec3b70ea64 / 9

Type: `"number"`. Computed.

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

<a id="canonical-35c56c97eca26f451f70001af450b05d03d21f551e232ed2dc063ed44ed902db"></a>

<a id="canonical-be00ca6172b210d24c60a06df3b71bb93691340229852d8c7b64416d0842c47f"></a>

## vendor property — infra.hw_info.cpu / ffec3b70ea64 / 10

Type: `"string"`. Computed.

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

<a id="canonical-ef4e5e812d686256491656c61759b2a9a7c3992d7e2d46d593a3b45cc4ee2b91"></a>

## Next pages — infra.hw_info.cpu / ffec3b70ea64 / 11

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-18b728ff653601cfc02eb75dc3762a3e51502ab8882dc313a79c4849a55f8a34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c07a7544e994911ed151f0e6c962ad5cd37cb23e147a9a91bcf00bb2eaad4cd"></a>

## infra.hw_info.gpu — infra.hw_info.gpu / d40cb3e508cd / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.gpu

<a id="canonical-c0b726570e47e146e616a0de6d2ee35390042345c603f10042ac3cd1f8e03a9c"></a>

Type: `"single"`. Computed.

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

<a id="canonical-39926d94cb0b3bf7dd49c793b0805c1816feaa38b094313f9de849503918ca13"></a>

## Direct properties — infra.hw_info.gpu / d40cb3e508cd / 3

<a id="canonical-9077d869f49f37bb1ba51a668ca293e8256e72718864fefe034e60fc96fad118"></a>

<a id="canonical-2f3cb81f234953e3db8de8e6e2b05ed3a0e7aa5ec36fed9c6973319475a91b74"></a>

## cuda_version property — infra.hw_info.gpu / d40cb3e508cd / 4

Type: `"string"`. Computed.

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

<a id="canonical-6a7c2324a7590f6d2b10afbac966023c0a609af2763f3979062e8f713c393b45"></a>

<a id="canonical-107347010e6a2224f91795dc383068e264406504b3d850606fd2255ad99381a9"></a>

## driver_version property — infra.hw_info.gpu / d40cb3e508cd / 5

Type: `"string"`. Computed.

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

- [gpu_device](data-sources--registration--reference--group-001.md#canonical-8819b1b710bb4a42dbcdecd72147fe99f6bac6e0b34aef6ac7bcc09b63f767c6): complete subsection reference.

<a id="canonical-5bd8860736bf594511c2c23f7f735c46531b02155e93fc724b7d43e682081052"></a>

## Next pages — infra.hw_info.gpu / d40cb3e508cd / 6

- [infra.hw_info.gpu.gpu_device](data-sources--registration--reference--group-001.md#canonical-8819b1b710bb4a42dbcdecd72147fe99f6bac6e0b34aef6ac7bcc09b63f767c6)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-8819b1b710bb4a42dbcdecd72147fe99f6bac6e0b34aef6ac7bcc09b63f767c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8c6f0018a9cf399ff05ced017a90b614c7e6b704f74abb3a1e4917250e6fd05"></a>

## infra.hw_info.gpu.gpu_device — infra.hw_info.gpu.gpu_device / 803779798dc9 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [infra.hw_info.gpu](data-sources--registration--reference--group-001.md#canonical-18b728ff653601cfc02eb75dc3762a3e51502ab8882dc313a79c4849a55f8a34)
- infra.hw_info.gpu.gpu_device

<a id="canonical-b5c4c33c56dd0b7e6208af67ddbbc19a1d9505b0ad4e234aa805088e303ad87f"></a>

Type: `"list"`. Computed.

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

<a id="canonical-dc0b92eaf206c7bb64591ccf07fb254869f83cda0fda871a1140a7bdf0ac5d38"></a>

## Direct properties — infra.hw_info.gpu.gpu_device / 803779798dc9 / 3

<a id="canonical-b6392d2308c986dd21a1483741e9431340a33d7a84cdeab5b5671faa8acc8603"></a>

<a id="canonical-8cbd1d9a61459a2a4bf0c233046b1c9a06e3683f44fe279b79b2d888a324ca6b"></a>

## id property — infra.hw_info.gpu.gpu_device / 803779798dc9 / 4

Type: `"string"`. Computed.

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

<a id="canonical-f60635c2b7b143af084618f97048f4b05eeaf1a9f392bf120baf0119a943a5d1"></a>

<a id="canonical-f74344c367aa95ef0014d66582dfb1d758905f570386dc7605512dce495bc345"></a>

## processes property — infra.hw_info.gpu.gpu_device / 803779798dc9 / 5

Type: `"string"`. Computed.

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

<a id="canonical-d400a7dfd493c4408106da0872785d7e377c168665491e648f23c8c69900a6d0"></a>

<a id="canonical-aa5668e9b1d92c443d94f4be1bda906a1ba1c5c742046dc6c8a4f4264ccf5a4f"></a>

## product_name property — infra.hw_info.gpu.gpu_device / 803779798dc9 / 6

Type: `"string"`. Computed.

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

<a id="canonical-119277f64b9dca5d6a601649c0f5d4d38473accd8f353f5c909ca74a1599889a"></a>

## Next pages — infra.hw_info.gpu.gpu_device / 803779798dc9 / 7

- [infra.hw_info.gpu](data-sources--registration--reference--group-001.md#canonical-18b728ff653601cfc02eb75dc3762a3e51502ab8882dc313a79c4849a55f8a34)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-e9411e953ab23e15aadccc3395f3df6adfb9f5a64b4f185ea2607e2adfab317a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a17e2417d567d4dab6bb11562810a8e3f468c1108785481e56b1ccd86d142e16"></a>

## infra.hw_info.kernel — infra.hw_info.kernel / 3cacf974948b / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.kernel

<a id="canonical-01c6feed763af81f81ed315517d22fd0dab1025aae71c58c308018fe97b3dff7"></a>

Type: `"single"`. Computed.

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

<a id="canonical-36776295b8622f8591f36b3be2c141700860c5e807a1473ce94cb9fbae4ceb67"></a>

## Direct properties — infra.hw_info.kernel / 3cacf974948b / 3

<a id="canonical-0d7b937a84946686f0cb6f25123f6f7e61eae8556644da650ffde38142fda775"></a>

<a id="canonical-19be6ae9c8af1c5cbeeefe0befdb545ddd3e0bcbfb33ee3de3bdc62503350fbc"></a>

## architecture property — infra.hw_info.kernel / 3cacf974948b / 4

Type: `"string"`. Computed.

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

<a id="canonical-f2eb1bf197e129ec6734dea416c4a9a76596b2274f76fb6e23b050c71e51bc19"></a>

<a id="canonical-ad9dfb58d1b619668229cf6fa50c9b0deb86d7e6cdc8ee3eec00300b1023d307"></a>

## release property — infra.hw_info.kernel / 3cacf974948b / 5

Type: `"string"`. Computed.

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

<a id="canonical-8ca4d13fb5611eb64af3548a815a3bfac95efea8258c947a747844e3b66af260"></a>

<a id="canonical-6593637eca43aa6b64b9d763b01268afa69262d84b97b9fa5eda2d8ac91574b7"></a>

## version property — infra.hw_info.kernel / 3cacf974948b / 6

Type: `"string"`. Computed.

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

<a id="canonical-6d25782f19e08066cffce7a6bcfa2719f516f20f01640e9bd7ea12622ae1e243"></a>

## Next pages — infra.hw_info.kernel / 3cacf974948b / 7

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-3b03cfa26c33c61bb43fdf5473a0e964162fd45461daf983d7069d000a398fc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8db53502081a0f74b3acf38a5e4c4b4e60e780247d1be9e77969ab23a48a9377"></a>

## infra.hw_info.memory — infra.hw_info.memory / 401fce50cb0b / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.memory

<a id="canonical-9d21295a158b1e882cccaade22deb31b90c4d5949b3ede8506022eeff2cd8013"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1ce3ccb8079793de97158b14dc86a5248223593e904e9adedd94cd6072271150"></a>

## Direct properties — infra.hw_info.memory / 401fce50cb0b / 3

<a id="canonical-23b9ed00f9b0857d8300b0470ba0d9da9a37217902f9508b9055968c5ebf760d"></a>

<a id="canonical-f1ba699b292aa3d483aa1171c38d7e01becba0727622eb4ceb61b5f0b944fa20"></a>

## size_mb property — infra.hw_info.memory / 401fce50cb0b / 4

Type: `"number"`. Computed.

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

<a id="canonical-dc40017ab764925ace9addedb36de16bf17939b0d1c7d61e18cc286870803af4"></a>

<a id="canonical-55e9d79065f906f802978475f55c5f2595c4feb5ae342695cdedb5a54828fdb1"></a>

## speed property — infra.hw_info.memory / 401fce50cb0b / 5

Type: `"number"`. Computed.

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

<a id="canonical-93e7ebdbd605199e6135215a9484edcdac4615368062cc646564c26172042944"></a>

<a id="canonical-7b494df5f15e1acdbb73e3d607e43f742e3fb7f5557b10bc292e7aea11f49ce5"></a>

## type property — infra.hw_info.memory / 401fce50cb0b / 6

Type: `"string"`. Computed.

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

<a id="canonical-9df5fe0ab85e0e03673d19d6f5e1df67876fe1d32ef2e0bd40e98ad6b9e7b624"></a>

## Next pages — infra.hw_info.memory / 401fce50cb0b / 7

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-820c5be99108b8f517fbf7b22318d3e3729ace16af7100f3038c1fdd4ae5b065"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae6a3d5b2d5a6fa09aa64403d90241b7af2f22e9169c12885b6133603787ec46"></a>

## infra.hw_info.network — infra.hw_info.network / e1fc6bb9569a / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.network

<a id="canonical-5a8f3ea4c380f801a5b8b1e3061c2076523239a4345287142b75aac13ec0d546"></a>

Type: `"list"`. Computed.

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

<a id="canonical-13fc17940ec154bf3acfef10f71bb89d77a752d55f6e94a2841973ea41ee7381"></a>

## Direct properties — infra.hw_info.network / e1fc6bb9569a / 3

<a id="canonical-152120246033b39a842e843ae63de2b10edd2a80f773fe1c04fbcc4d85d70e37"></a>

<a id="canonical-f09994bd91925cff0f565a37443a342829636d2cd1c785019f3dc06ba8dc9352"></a>

## driver property — infra.hw_info.network / e1fc6bb9569a / 4

Type: `"string"`. Computed.

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

<a id="canonical-47a74c1555448891ed7634a850db7ed4b3113ff86c3d7b7c8cec27000e1a4887"></a>

<a id="canonical-89ab99231fff5df14142e99cde71d7f73279fc406eaa5226a0589a79c200e1e4"></a>

## ip_address property — infra.hw_info.network / e1fc6bb9569a / 5

Type: `["list", "string"]`. Computed.

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

<a id="canonical-b300aa4c50290b00ebf108cd134d0b621ec0342dc084585d07cd5245506fe0f3"></a>

<a id="canonical-d7f29e73bab211bcded0e07d79d0486eff26db1372cd8a4a295687cfdf5ecc54"></a>

## link_quality property — infra.hw_info.network / e1fc6bb9569a / 6

Type: `"string"`. Computed.

\[Enum: QUALITY\_UNKNOWN|QUALITY\_GOOD|QUALITY\_POOR|QUALITY\_DISABLED\] Link quality determined by
VER using different probes Unknown quality Link quality is good Link quality is poor Quality
disabled. Possible values are \`QUALITY\_UNKNOWN\`, \`QUALITY\_GOOD\`, \`QUALITY\_POOR\`,
\`QUALITY\_DISABLED\`. Defaults to \`QUALITY\_UNKNOWN\`.

Upstream description:

Link quality determined by VER using different probes

Unknown quality Link quality is good Link quality is poor Quality disabled.

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

<a id="canonical-2f68f06443b41f97fa8162afb54cf0f1d3def77e58482c340ec7fe46809d96b0"></a>

<a id="canonical-d191dbaf6f185afafbda3dcaef7cc039feb93802de70f0424887dc4c3b7529bb"></a>

## link_type property — infra.hw_info.network / e1fc6bb9569a / 7

Type: `"string"`. Computed.

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

<a id="canonical-151fa0746dda4bed7c9feb36d62b4072ae9efced55295ac057b146fc46b4a769"></a>

<a id="canonical-6af2842ea7903ff0bbec2a34c9df59c389157fb101e385e35a366d5ecf599476"></a>

## mac_address property — infra.hw_info.network / e1fc6bb9569a / 8

Type: `"string"`. Computed.

MAC Address. MAC address on interface.

Upstream description:

MAC address on interface.

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

<a id="canonical-a9b0908cee823ffd57db6ce4743d9195b192e793f8baaf87cadae13c2d4556a4"></a>

<a id="canonical-af91927d0d5cb5f2294fafd2aed435b865daeb1ed14b462689cad88349055fdc"></a>

## name property — infra.hw_info.network / e1fc6bb9569a / 9

Type: `"string"`. Computed.

Name. Name of device, eg. Eth0.

Upstream description:

Name of device, eg. Eth0.

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

<a id="canonical-8fac2d641d1e5b4ed61629a7c9dbfba98d356067f9097dce0682489dbad45080"></a>

<a id="canonical-79164b2778de109855a3a9dc70340ce90bc767d0159e68428f2cc16048814390"></a>

## port property — infra.hw_info.network / e1fc6bb9569a / 10

Type: `"string"`. Computed.

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

<a id="canonical-7d93d488c4f26ad530d2ca586d592fbd582755a9b0aef674cdaa6c41886d9f83"></a>

<a id="canonical-71b54bd4dc653f6dd580776c2f94ff6c8c3a44ef094fae5f2c1abe79b287af74"></a>

## speed property — infra.hw_info.network / e1fc6bb9569a / 11

Type: `"number"`. Computed.

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

<a id="canonical-ce22c192504ca5dfd7c3549e38c59863c9080b37238777286a729a8f083aae62"></a>

## Next pages — infra.hw_info.network / e1fc6bb9569a / 12

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-8bcbb6bac391159990da9202be8fed6f04f430ad1886e198be1cc0cb11bac1fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1d67204ed4b6869a19a9502439fb7fc86d7743d8fbc41631640b53a983d4b4d"></a>

## infra.hw_info.os — infra.hw_info.os / b489084c71a9 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.os

<a id="canonical-d2890a2e03daa3522ac7d7502356bf2972c20128c6a2cf5acbd1e2826e26bdad"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1d888d83a3f5fde024bb01565ebf45257aa0a1094e595aeb527a5f6c7b35e6de"></a>

## Direct properties — infra.hw_info.os / b489084c71a9 / 3

<a id="canonical-7b79cd287fc76c7917e4414b2638b1e5156ba51c8774da3c29750e651c982bce"></a>

<a id="canonical-ae57304965749f3f10e2efa88f37e3c0857694bcf0d66695136449ddf68a9212"></a>

## architecture property — infra.hw_info.os / b489084c71a9 / 4

Type: `"string"`. Computed.

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

<a id="canonical-ad04f6535672b0fb4fd32dde5f1a02c4ffbd8e4e161db6ef14487722a2ad2f0a"></a>

<a id="canonical-86f5b5373af92a1a42876065c7ad85dbfbc272576308976f30aabbeca0b25531"></a>

## name property — infra.hw_info.os / b489084c71a9 / 5

Type: `"string"`. Computed.

Name. Name of OS.

Upstream description:

Name of OS.

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

<a id="canonical-e6e6926d320c95adaa2c83f69d8094c3d49b26c3d96edad7ce82a42db8c1e466"></a>

<a id="canonical-5761447a6adbd29d8dfbe9b930b17ad54b6c1d10f545606aa85a22388f117104"></a>

## release property — infra.hw_info.os / b489084c71a9 / 6

Type: `"string"`. Computed.

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

<a id="canonical-841c64929fd5968c9845050d3833e8e470a82d8c4a45d94316b957ff544ff4ee"></a>

<a id="canonical-6e9b71f517086dd57f6531023317974939ea366f65a4c4aaf307db9b1bcdb474"></a>

## vendor property — infra.hw_info.os / b489084c71a9 / 7

Type: `"string"`. Computed.

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

<a id="canonical-fba7d5af0ef9f7d612b897f9db891b8d2ce406ff9eb7c42dde3b94f29a5b201d"></a>

<a id="canonical-9ce8cac1e6c88c00c67b238316b258aed3915d2ce21429601ffc68bb38e89bca"></a>

## version property — infra.hw_info.os / b489084c71a9 / 8

Type: `"string"`. Computed.

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

<a id="canonical-e52a676713fb90145338de37d10ec993c442fb641b3169fa04c0fa24f6df5d75"></a>

## Next pages — infra.hw_info.os / b489084c71a9 / 9

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-5cd2a5f5f07aec658752a5088ccdd898ebe9fd093a1f2413b6a637b7c0f8a46b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-631b10e426f143e15d493e61875c68ed767744d9f4678fdc1643b1d4e94bfd39"></a>

## infra.hw_info.product — infra.hw_info.product / a2f73a559ffa / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.product

<a id="canonical-01f5ce00011b541227ad6cc0b4767f4347cbbfc629e971f13440d96cc9245661"></a>

Type: `"single"`. Computed.

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

<a id="canonical-64b774ba0397e725a9316d7ebfbe9ea0ec9c502bec806c4122f009e3ae7f63a2"></a>

## Direct properties — infra.hw_info.product / a2f73a559ffa / 3

<a id="canonical-e01342aad605978eca8f3bcf9fcbb7b3b2fc85028ff3e0e97cbee1b58efb77cb"></a>

<a id="canonical-893fca213ede96a7227576cf36c69f2b02f088a59174954011adea45cbc33ae2"></a>

## name property — infra.hw_info.product / a2f73a559ffa / 4

Type: `"string"`. Computed.

Name. Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product\_name.

Upstream description:

Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product\_name.

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

<a id="canonical-cf7a5cd1ff019a238809bb68179098be186b3fd497f955e23521a4b1d7f33f95"></a>

<a id="canonical-4a43099dbdff5baffea3f1b64e07a9bf85c373630823de1f7801ba1669064159"></a>

## serial property — infra.hw_info.product / a2f73a559ffa / 5

Type: `"string"`. Computed.

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

<a id="canonical-f493c823b6ad2a390201978488cefa2a3e62964853942d59bb4d73781a80f8e8"></a>

<a id="canonical-311ff3aa38f4b57f97c555d9f7c7ae7bcdaa7f35d4736565198f854c80edbd44"></a>

## vendor property — infra.hw_info.product / a2f73a559ffa / 6

Type: `"string"`. Computed.

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

<a id="canonical-f2c6b857bb6b1011ebf24f742898556b7d5d6c57b5a18c71a9e49942eb3b32ae"></a>

<a id="canonical-966b26d64160920a7148244cce19ab260cb5ebec2c538b39066740b3b0164d12"></a>

## version property — infra.hw_info.product / a2f73a559ffa / 7

Type: `"string"`. Computed.

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

<a id="canonical-a4b7afd2311d0d01dc3d6a2a6cc657b1b22ba40e7ae2ac815eb9a40cfa07f91f"></a>

## Next pages — infra.hw_info.product / a2f73a559ffa / 8

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-db1e154b948c3d7a4bdf736a72b0b6c1dac635f96835094b06de205af71cadbd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c48f4e7dfbaa92b44a9e0adde5af1f4e259681e2dbbddb44e614de6369ef0ebe"></a>

## infra.hw_info.storage — infra.hw_info.storage / c3bf08db5597 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.storage

<a id="canonical-13af4265aa1b701288082f18907a01c7e7716c4c008b11a03bb9c1f749d622f5"></a>

Type: `"list"`. Computed.

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

<a id="canonical-8911df33d92aa4fe8a7ae44853f5e40c1be35fa1109088d0ed046fcca163cce5"></a>

## Direct properties — infra.hw_info.storage / c3bf08db5597 / 3

<a id="canonical-9bed739cc0bea1358741b592678e776e48026e5ab93f911a00d90f3b3032742a"></a>

<a id="canonical-3a1155f539ac06e99cfbc290e27dcb28dee7a0449dffe20f3a1b32c08b3d8dad"></a>

## driver property — infra.hw_info.storage / c3bf08db5597 / 4

Type: `"string"`. Computed.

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

<a id="canonical-be99188b2e9e2dd736950b24824554eaf134aad77d76f707f4719b49495aee6d"></a>

<a id="canonical-930a7afa886308d487d52b7ff9398bdbe62fd54188c2c85ce92b35de1058e39b"></a>

## model property — infra.hw_info.storage / c3bf08db5597 / 5

Type: `"string"`. Computed.

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

<a id="canonical-75540bfb81eee9729e2caf3200960f40484310302bbcb18b02dfac0390083a52"></a>

<a id="canonical-0430927fdfee0d7f8bb03cce33ec9dd58e0acdf720d1f6c79c98119a3404cdc0"></a>

## name property — infra.hw_info.storage / c3bf08db5597 / 6

Type: `"string"`. Computed.

Name. Name of device, eg. Nvme0n1.

Upstream description:

Name of device, eg. Nvme0n1.

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

<a id="canonical-11a4c9023dfd4cafadf40ce965e21bac81e4b02da103d0e756400202d0bb3e42"></a>

<a id="canonical-8541fd3fc7a4428468f116c19332d83313765f3318b34f04c96a93aba0542598"></a>

## serial property — infra.hw_info.storage / c3bf08db5597 / 7

Type: `"string"`. Computed.

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

<a id="canonical-3b9ec393e2f7801aabee1ac914bcebce8b79c8ba839db6ea30a1d44b71df8bd2"></a>

<a id="canonical-312560dc39b2062b961ce2221a66cafc7564528eb8e9e01a95a333249dc982a0"></a>

## size_gb property — infra.hw_info.storage / c3bf08db5597 / 8

Type: `"number"`. Computed.

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

<a id="canonical-bea58b6f2b8dbdd28b3512e14ae8035a4d8cf05f31f53c2952ba7a05ac00d920"></a>

<a id="canonical-3bd401235af664530823a08ed85ff6f24667eec4de5426bcd4d9165b412e1e90"></a>

## vendor property — infra.hw_info.storage / c3bf08db5597 / 9

Type: `"string"`. Computed.

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

<a id="canonical-ef8f2d2bce11eafc1d523afc197d47e5f13f3c16dfbffee6394dc950fa6fa670"></a>

## Next pages — infra.hw_info.storage / c3bf08db5597 / 10

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-542fdb8614ec2c371d6ee65c34fdcc803998f066bcbaac2bb79cdee1119d16b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fcc5307756e76c7f2a7beaf8e424f3c293fee62154af4eea9e92d34836908bf2"></a>

## infra.hw_info.usb — infra.hw_info.usb / ab5ddbf2b93f / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- infra.hw_info.usb

<a id="canonical-61cecb10c88e6fe786fe81d4d87c18458f820d627857688381e8265c5887c3f2"></a>

Type: `"list"`. Computed.

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

<a id="canonical-adad24cecef829f92f0ac45e351c97db8a40451f988c599e9ddd131573475eb4"></a>

## Direct properties — infra.hw_info.usb / ab5ddbf2b93f / 3

<a id="canonical-5fdd6997b2b0eac64c008e2a6ba461fd6e7fc9219bb308533a74fe1b81da1171"></a>

<a id="canonical-bdb890e05f580b8776fe44920db18d37e69569e68f70d4158002e6f82b1dad8e"></a>

## address property — infra.hw_info.usb / ab5ddbf2b93f / 4

Type: `"number"`. Computed.

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

<a id="canonical-96aa255e4448dbc5a1bc85c733b1b8c8b82ed118786318d8fd1af1b1935ba943"></a>

<a id="canonical-48a30aaa047bc120dd52c7bc49c3694e286f522b9f418a01a2a0b28a5e4ba765"></a>

## b_device_class property — infra.hw_info.usb / ab5ddbf2b93f / 5

Type: `"string"`. Computed.

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

<a id="canonical-c913a09218aa68ff12363aacb3361d44e822ec1634b384d0751ff0543f3e61fb"></a>

<a id="canonical-0eb4e427bf95fb74a6c7160e5894fbf8d6327f10a718cbb376789d10da1c7a24"></a>

## b_device_protocol property — infra.hw_info.usb / ab5ddbf2b93f / 6

Type: `"string"`. Computed.

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

<a id="canonical-bdc69450e46b26cb29d3753c1a68afac6a0ecd2a9a02e61a9c6701d2cf8d4036"></a>

<a id="canonical-bb4f231d764d5d3011e3bdd8f3407bf7397d09ef62d9a5f04f1ca2a0bae754dc"></a>

## b_device_sub_class property — infra.hw_info.usb / ab5ddbf2b93f / 7

Type: `"string"`. Computed.

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

<a id="canonical-306311b6133b1b9dc494ca9838ed4a462ecddab4c977d8ac5831acb52ec14164"></a>

<a id="canonical-3c2c4218b93a65c235f925cb078d2c148c72c5aa2e5364b7322b5d077bae505b"></a>

## b_max_packet_size property — infra.hw_info.usb / ab5ddbf2b93f / 8

Type: `"number"`. Computed.

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

<a id="canonical-789197a203a2f0ece7b6919432fe6b1c5642e973ef45e14f1e6b399c6bb64f2f"></a>

<a id="canonical-1253413f85719f68a5e438422b8a5b16f83c37491c8b3c50f3651c7cdfa40c2d"></a>

## bcd_device property — infra.hw_info.usb / ab5ddbf2b93f / 9

Type: `"string"`. Computed.

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

<a id="canonical-ee1bb48a34b54d63c9c840253d7d430dddec7f2e1c53bfc6272b50f67caedb5b"></a>

<a id="canonical-5ad0988f4b7557793d9c54296057eab9de3e9bc0b81d998104c0ce5d58b87913"></a>

## bcd_usb property — infra.hw_info.usb / ab5ddbf2b93f / 10

Type: `"string"`. Computed.

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

<a id="canonical-64dacdef66b6986fa581a67b34726f243c2ededb686210243fb54eaca1f55c0b"></a>

<a id="canonical-5139499080c5722687ce1c68d27d7a4879970860f114edaaba989f4d50c6a9cf"></a>

## bus property — infra.hw_info.usb / ab5ddbf2b93f / 11

Type: `"number"`. Computed.

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

<a id="canonical-6193bbf8d42e11147e5eac2b3881cd93545ed86174514d2b73e3d34d4b5faea6"></a>

<a id="canonical-d010899c50908a8b88f79f885e44e0cb100170746242df106c1029c2c96bf6bd"></a>

## description_spec property — infra.hw_info.usb / ab5ddbf2b93f / 12

Type: `"string"`. Computed.

Description. Device description.

<a id="canonical-0e94491e5288737232042fc2e8c12f03936502bcbca8c571e811357274e5a9f4"></a>

<a id="canonical-ee2f5f58caff7d22dc0027b206f3f1337c0902b35b653acdc7628d5811809556"></a>

## i_manufacturer property — infra.hw_info.usb / ab5ddbf2b93f / 13

Type: `"string"`. Computed.

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

<a id="canonical-d6f1283505b477f56c38d721c835cc75ec7f476c2e61e138c5da466f967b45d9"></a>

<a id="canonical-756dbedd1c1ddab804c8027ea23eca92efb221c32046c5f9142dd8909cf3e4c0"></a>

## i_product property — infra.hw_info.usb / ab5ddbf2b93f / 14

Type: `"string"`. Computed.

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

<a id="canonical-5fb9e28656ab99422836e876608c051e18cbcf7007818bbec5964e5909851331"></a>

<a id="canonical-67e8ea54418a6582ce1dfdc872dddd8acffe84b4636b1336aae12d17bd353db6"></a>

## i_serial property — infra.hw_info.usb / ab5ddbf2b93f / 15

Type: `"string"`. Computed.

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

<a id="canonical-4fd647f1b068835ddfe591dc0ec2ea3f39bddc1bf2d827acc047ea5e78ca87ac"></a>

<a id="canonical-13a08a1426547b137648a1ae14423bda973bf9dd160fe60c99b869f54319a608"></a>

## id_product property — infra.hw_info.usb / ab5ddbf2b93f / 16

Type: `"string"`. Computed.

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

<a id="canonical-f6651f4d154e60c4d50f54a88b17ebb72245c21365c645fd3caa3f017f700ffe"></a>

<a id="canonical-ad8df8e56ae352d92f2e3df35c680eaea2b609d5315d009f416806bc5a199544"></a>

## id_vendor property — infra.hw_info.usb / ab5ddbf2b93f / 17

Type: `"string"`. Computed.

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

<a id="canonical-b777ede2990f42c748101d076ec99dc7a524ac2156db51296f55def0db2df339"></a>

<a id="canonical-99d6bc61fb9389566d09eff9119647c654e6c91738cac9dd967c0306e5f89580"></a>

## port property — infra.hw_info.usb / ab5ddbf2b93f / 18

Type: `"number"`. Computed.

Port on which the device was detected in decimal.

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

<a id="canonical-0919dbc918d4ea98352697b8190479027a59fd0f33498f3facaa9fdeacfbad6a"></a>

<a id="canonical-f98f01cbf977da0af958ca97cb14f0ed617f762ee5e3bc2e719f358e695514be"></a>

## product_name property — infra.hw_info.usb / ab5ddbf2b93f / 19

Type: `"string"`. Computed.

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

<a id="canonical-a145f97d57f574ea9876d8ec6f899ff1fc963087e5ad80a50ecbc5b534aabb44"></a>

<a id="canonical-4a2535f42c6efbd1ab7060280c73e66fecbe0f476ac50d66304d1fdafc0b4a38"></a>

## speed property — infra.hw_info.usb / ab5ddbf2b93f / 20

Type: `"string"`. Computed.

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

<a id="canonical-05390abad848080e7ae2c829f69edd6302ce1605ce3c173e0d3b6c599dae03e8"></a>

<a id="canonical-de627d5cf4ad55d0bff43a42019ad2205f4f068fb33cee6b230aeb8f2002ddff"></a>

## usb_type property — infra.hw_info.usb / ab5ddbf2b93f / 21

Type: `"string"`. Computed.

\[Enum: UNKNOWN\_USB|INTERNAL|REGISTERED|CONFIGURABLE\] Type of USB device Unknown USB device type
Internal USB present in Certified HW USB device present during node registration USB device that can
be matched by USB rules. Possible values are \`UNKNOWN\_USB\`, \`INTERNAL\`, \`REGISTERED\`,
\`CONFIGURABLE\`. Defaults to \`UNKNOWN\_USB\`.

Upstream description:

Type of USB device

Unknown USB device type Internal USB present in Certified HW USB device present during node
registration USB device that can be matched by USB rules.

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

<a id="canonical-ff897000262d59ec8ee0848773ac4c220865541ea07cc59d92d772265ae0cdaf"></a>

<a id="canonical-860c923c2f11a5a010436c77e43c6d7ed849f4a8ba6281849e3aace96e4ea6e1"></a>

## vendor_name property — infra.hw_info.usb / ab5ddbf2b93f / 22

Type: `"string"`. Computed.

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

<a id="canonical-0ce505552739b458eaa7dd92fc2e838a2fee7ec8eb1000d69d22d3c97c525e12"></a>

## Next pages — infra.hw_info.usb / ab5ddbf2b93f / 23

- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-03d094c4da89d5f13def9527b667d4dad3891cd1bb70c71bb3ce32634fc0e4d4)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-66e79b237f37efd5879541c286d9c104de02734a72db584f56cabcfc76bf017d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3482e66e14d30a27f2b0b879c2302e9970b1ac25ef801590563e6e9eba0bba53"></a>

## infra.interfaces — infra.interfaces / 055c4a269853 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- infra.interfaces

<a id="canonical-cc176c11d0ed98583899474f02130f872dd7be318f058aa564620735b54ac56f"></a>

Type: `"single"`. Computed.

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

<a id="canonical-5b1776f7445b2926bbdd94ffabfd6294f16d3d8830e67025258fdce4a2b58787"></a>

## Direct properties — infra.interfaces / 055c4a269853 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26214c0e60cb06b44aba789940058b57540eca8eb76bb710b7d2dc45b7e5a8f3"></a>

## Next pages — infra.interfaces / 055c4a269853 / 4

- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-8a5def9df187ad5173d54e74d358a2239e37981290fd63c52a9a3057a477131b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e0ac8bbe54da28675e218e2cc14aaf96099970592f016b9d62f685db445df51"></a>

## infra.internet_proxy — infra.internet_proxy / 927fb9d4f123 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- infra.internet_proxy

<a id="canonical-7049e6f2d7d90e37a3d0812d70ad82ef248843da0d2288cd25d70ad191e61de6"></a>

Type: `"single"`. Computed.

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

<a id="canonical-5a7f5e86ac57061eb6d0d624ee7721599e1bb4940d4e169be625e320a3976556"></a>

## Direct properties — infra.internet_proxy / 927fb9d4f123 / 3

<a id="canonical-34c1a7845752198b925e2fb8e2b36c6fcc09c84b7f407b4ea80c3cbddff7c921"></a>

<a id="canonical-2f0cf7509eb3d67dca166ee7a2ef644112287f77f48cd95cf91cac1ea2029a18"></a>

## http_proxy property — infra.internet_proxy / 927fb9d4f123 / 4

Type: `"string"`. Computed.

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

<a id="canonical-12599a84ffb4dbc4f2da290685c68a8f926896113b25f4e783cb8e2623e11926"></a>

<a id="canonical-571b034610a7748a0d80a1b14bcbaa83062b51a1930e14302f5a4c43f78affe4"></a>

## https_proxy property — infra.internet_proxy / 927fb9d4f123 / 5

Type: `"string"`. Computed.

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

<a id="canonical-087c6453087d26821db98825a356da83631b4e41cdb6c5711bb10f54bfa33df8"></a>

<a id="canonical-82788b452620d73f2504973d72650f10a092c5bd1814e91fd49ba188c8678efd"></a>

## no_proxy property — infra.internet_proxy / 927fb9d4f123 / 6

Type: `"string"`. Computed.

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

<a id="canonical-0dfaf03d195bd17bc39c6318a22d1d173d5c39d98e0be4906bb7e21f1cf70b0e"></a>

<a id="canonical-45950dbea7480c5e27143d320a9f9c7a0bc013d6ef583a6cf6fa17d47569064d"></a>

## proxy_cacert_url property — infra.internet_proxy / 927fb9d4f123 / 7

Type: `"string"`. Computed.

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

<a id="canonical-6196487b91b66c66f3dd96d7c20a523ab2e3d510f90f0f49760781fb959ea07f"></a>

## Next pages — infra.internet_proxy / 927fb9d4f123 / 8

- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-6dd402088ad909db8bc33ca16006763d6574c8e94ebcf113e42c2bc20d5edbf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-094ae6e332738e305c1e0e8640fee438ef6e4370a895e641f6859900cd6f1fd2"></a>

## infra.sw_info — infra.sw_info / dd1858e43640 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- infra.sw_info

<a id="canonical-031d3b5037360d922c55b5daf24317d6e7b1822b4871c64df862c84dd254f96d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-d97da9b9a734b4eb61cdbf181289befc82f777663269cb20d2bdaed7b7e9ded5"></a>

## Direct properties — infra.sw_info / dd1858e43640 / 3

<a id="canonical-cfce203206f6e802e1ee679524cdac4667b96aa3784998d58ce5a4ae1ca09e94"></a>

<a id="canonical-2e5cef309fb76501131543eadb80614024865a280f6f0d0eb7d55ff66f8b99d8"></a>

## sw_version property — infra.sw_info / dd1858e43640 / 4

Type: `"string"`. Computed.

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

<a id="canonical-5b9e9e42790f76ac87d9daf51679ff3959c2caa6556019b5882e2aff03924e9b"></a>

## Next pages — infra.sw_info / dd1858e43640 / 5

- [infra](data-sources--registration--reference--group-001.md#canonical-7aa07ccd59a9d38f3083c54865ad44476af276e2b74a2c9d2a70c80954695251)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-2bb99d2d6317422b03cc3b89a4cc4e42278c4b94e5f3a453f7c72c70228ebba5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6dd7cd6445cd1b5bbd8e3c27c2bf1c39c5ec5c67f9a2ab47858a9822fefb5784"></a>

## passport — passport / 6bb165dc659a / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- passport

<a id="canonical-a56dab5f7a353a4caef5e774a989f31214a11e309edf4baa760ad1a85d540da4"></a>

Type: `"single"`. Computed.

Passport stores information about identification and node configuration provided by CE during
registration. It can be manually updated by user during approval.

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

<a id="canonical-4aabcd3602bfa0a230d0b409049c9bed179379c6451d4af8ca361d7fde9932eb"></a>

## Direct properties — passport / 6bb165dc659a / 3

<a id="canonical-50de270cda18adce9a38b906bbc50d1ec284863326768adb35bed7b3d09631fb"></a>

<a id="canonical-27129a4213dfb3810ba31de4dc30b3236e5482062d350181c6c76690a67fe02d"></a>

## cluster_name property — passport / 6bb165dc659a / 4

Type: `"string"`. Computed.

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

<a id="canonical-042836eafe92a9cf2d08a70e2f69399b6b192556fdbe103abb6f18e4c4a9fce4"></a>

<a id="canonical-8fa623363782c3853c0573664c135e28abdfed1c92d9714fde651249be9138cb"></a>

## cluster_size property — passport / 6bb165dc659a / 5

Type: `"number"`. Computed.

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

<a id="canonical-6c0e49d49f85cc2f03048bbc054b2ebd422d1f1dfeffa1a5f5d7f30422863bb9"></a>

<a id="canonical-b629d9c6405852a8d31ea11b415c311c23b19f417ee9e80dea6df0bdd53331b4"></a>

## cluster_type property — passport / 6bb165dc659a / 6

Type: `"string"`. Computed.

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

- [default_os_version](data-sources--registration--reference--group-001.md#canonical-7174da0146bed552b61b530e823e582d300395b65a5384f5f6365234b97f48f4): complete subsection reference.

- [default_sw_version](data-sources--registration--reference--group-001.md#canonical-c5be527e8c6acdd4559846c6fcb016b06caa3cd29d93508df4ea2c76bd9a677d): complete subsection reference.

<a id="canonical-533bdddd825426963cc0b937e55698abc17753010201b326a527b715680cca31"></a>

<a id="canonical-4744a2d711f7821ab0bac336c1693b5c394e01bcd05b89edf472d9c4bd6b436e"></a>

## latitude property — passport / 6bb165dc659a / 7

Type: `"number"`. Computed.

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

<a id="canonical-d4126dbdd48ba1e553b18aea9d57cf4dae8b5d1b0018a040091541fefeb000e5"></a>

<a id="canonical-b678ef7dc3b1b47e336e594b55d43ee7b87a6db7efb17ea76616098b5290d8ad"></a>

## longitude property — passport / 6bb165dc659a / 8

Type: `"number"`. Computed.

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

<a id="canonical-a7e92f460ecfc8a59c9798c6233c3f332ac815df1319233a3e3f1225bba1bf06"></a>

<a id="canonical-691be627cac943f31622624b8fb21233a987491787897fc939e101caad90d086"></a>

## operating_system_version property — passport / 6bb165dc659a / 9

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Upstream description:

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

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

<a id="canonical-e5299d5a774c8ee43d037817a69f5af86d0fe7f61a644da4ae51cb46799cd8b0"></a>

<a id="canonical-55f51876a64ed39e483cdf8996de7d4677c0af01df0c277974320dfdd9a87d57"></a>

## private_network_name property — passport / 6bb165dc659a / 10

Type: `"string"`. Computed.

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

Upstream description:

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

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

<a id="canonical-55cc1eb104b02d9bf76ea5507241bcaee426d82ec9bce7e0396680a40e926a29"></a>

<a id="canonical-2384a88aed845479359d7601dd57bc29208e21b82c54ca6ceb0e883b8583c22f"></a>

## volterra_software_version property — passport / 6bb165dc659a / 11

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

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

<a id="canonical-5947e215a503eb40e39085529a01d9b60611a23fe2862808b59e390a2e236e5d"></a>

## Next pages — passport / 6bb165dc659a / 12

- [passport.default_os_version](data-sources--registration--reference--group-001.md#canonical-7174da0146bed552b61b530e823e582d300395b65a5384f5f6365234b97f48f4)
- [passport.default_sw_version](data-sources--registration--reference--group-001.md#canonical-c5be527e8c6acdd4559846c6fcb016b06caa3cd29d93508df4ea2c76bd9a677d)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-7174da0146bed552b61b530e823e582d300395b65a5384f5f6365234b97f48f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa87e89b933f85382cf72be5a7a7d666f9535320fec424e55da53f5d24f625e4"></a>

## passport.default_os_version — passport.default_os_version / 38ef96ab49cf / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [passport](data-sources--registration--reference--group-001.md#canonical-2bb99d2d6317422b03cc3b89a4cc4e42278c4b94e5f3a453f7c72c70228ebba5)
- passport.default_os_version

<a id="canonical-89b724215a31292ee4898fd234e4d0fad59ed2f383310b2eb6173f04c5268f12"></a>

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

<a id="canonical-528e9dd4d6e33a3235811022f02998a708cd4010ceabad0ba01cadc646985e67"></a>

## Direct properties — passport.default_os_version / 38ef96ab49cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aaf8663356d1acee212f217885a4b548e83f80bda9964b879690f73ee54db93c"></a>

## Next pages — passport.default_os_version / 38ef96ab49cf / 4

- [passport](data-sources--registration--reference--group-001.md#canonical-2bb99d2d6317422b03cc3b89a4cc4e42278c4b94e5f3a453f7c72c70228ebba5)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-c5be527e8c6acdd4559846c6fcb016b06caa3cd29d93508df4ea2c76bd9a677d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32144598711004b81b24434494042759d45d758f91f907906ca68957c8f5836c"></a>

## passport.default_sw_version — passport.default_sw_version / 8abea8c73932 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [passport](data-sources--registration--reference--group-001.md#canonical-2bb99d2d6317422b03cc3b89a4cc4e42278c4b94e5f3a453f7c72c70228ebba5)
- passport.default_sw_version

<a id="canonical-833fe530a0c541863bd4b92f5a9a98ed988dd4b80bd4c25b2d06460cdcf8c197"></a>

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

<a id="canonical-cb1dd8e7d8e77dbfbbc9b6151405cbdc0938ad5f1dc9dedebbc3ad39b5503c94"></a>

## Direct properties — passport.default_sw_version / 8abea8c73932 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19bc5984d445ed27b59f444866f4aae2f230b8b9535c6260e331840c64a1dc6c"></a>

## Next pages — passport.default_sw_version / 8abea8c73932 / 4

- [passport](data-sources--registration--reference--group-001.md#canonical-2bb99d2d6317422b03cc3b89a4cc4e42278c4b94e5f3a453f7c72c70228ebba5)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
