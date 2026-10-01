---
page_title: "xcsh_nat_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy reference."
---

# xcsh_nat_policy reference

<a id="canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b529dc5a3d5312a2d2895261c444f08236b0e06f66ae97d0889f94e714a2dfd4"></a>

## Property reference — Property reference / 41f75688a175 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- Property reference

<a id="canonical-71718953cd77b429469bd0a44f68feb958f7af96236d733b0fe1df399ad90349"></a>

## Direct properties — Property reference / 41f75688a175 / 3

<a id="canonical-ba917a15bc6bc65d2156755e6bfa1e523f1b3841db3eb0ed51a24056b3f5f0dc"></a>

<a id="canonical-c04191b2b9dd5130a75e68ba49d31daca89c2bbfed5c72189691de9c1b1b71f1"></a>

## annotations property — Property reference / 41f75688a175 / 4

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

<a id="canonical-c125130e8f53f01833b09c5af235ead038d12e431c11fdaabbe99c93f13ab992"></a>

<a id="canonical-8a0481c7ac3370453783c1739506b62b131ffa8965ac77f3581c1fdddce796a8"></a>

## description property — Property reference / 41f75688a175 / 5

Type: `"string"`. Computed.

Description of the NATPolicy.

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

<a id="canonical-60cdbcd4bbb01190ed71bd13a76d0a3b5977c5df350819ae23c0307d8a9298d9"></a>

<a id="canonical-c34fa2a63f2599ca533dfd5312da03f4eeff5106968d17d44ded5f9b9a900c0b"></a>

## id property — Property reference / 41f75688a175 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-b58d280ffc3ac4a33e34f3225adfc94908e0ab3f654b2326edd1a47f4e4c0366"></a>

<a id="canonical-bba95b99304b552446f0afddcad1c5fd5e4deff9450bacf5bd67fdcef67d7062"></a>

## labels property — Property reference / 41f75688a175 / 7

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

<a id="canonical-cff81ddb586ea04db713840b4c0b37d1141e970a026045c27bc01e01d1da3729"></a>

<a id="canonical-e21beca1a5c6122223500be09072355e4c8e39cf61fff9c61d46a5ac000c4ef5"></a>

## name property — Property reference / 41f75688a175 / 8

Type: `"string"`. Required.

Name of the NATPolicy.

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

<a id="canonical-7692a894ea525de07744b7cd70270643d4b8d19de9ab587c89e293a9a70c4141"></a>

<a id="canonical-30714c606614fff61803ef18e2edeb54cc4cb67ce7f5ffb67881b8259d58ddab"></a>

## namespace property — Property reference / 41f75688a175 / 9

Type: `"string"`. Required.

Namespace where the NATPolicy exists.

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

- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec): complete subsection reference.

- [site](data-sources--nat_policy--reference--group-001.md#canonical-45c6954ad1fe44c2f95f33065ed54ee51b8ee9ad280c257c2c757d100b9b89d4): complete subsection reference.

<a id="canonical-c034c34f5116b8465bc966066ef94076b09c9f1cc754a293cdd538ec68ff553d"></a>

## All schema paths — Property reference / 41f75688a175 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nat_policy--reference--group-001.md#canonical-ba917a15bc6bc65d2156755e6bfa1e523f1b3841db3eb0ed51a24056b3f5f0dc) |
| `description` | [description](data-sources--nat_policy--reference--group-001.md#canonical-c125130e8f53f01833b09c5af235ead038d12e431c11fdaabbe99c93f13ab992) |
| `id` | [id](data-sources--nat_policy--reference--group-001.md#canonical-60cdbcd4bbb01190ed71bd13a76d0a3b5977c5df350819ae23c0307d8a9298d9) |
| `labels` | [labels](data-sources--nat_policy--reference--group-001.md#canonical-b58d280ffc3ac4a33e34f3225adfc94908e0ab3f654b2326edd1a47f4e4c0366) |
| `name` | [name](data-sources--nat_policy--reference--group-001.md#canonical-cff81ddb586ea04db713840b4c0b37d1141e970a026045c27bc01e01d1da3729) |
| `namespace` | [namespace](data-sources--nat_policy--reference--group-001.md#canonical-7692a894ea525de07744b7cd70270643d4b8d19de9ab587c89e293a9a70c4141) |
| `rules` | [rules](data-sources--nat_policy--reference--group-001.md#canonical-dafa45a5831de76a4c8bbf730e92b17f46da2c46c77196428139641e99573e19) |
| `rules.action` | [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-dc742ea365aa3bce75c14b1ac664c62ad7b656b95d238191184f555f029b3a28) |
| `rules.action.dynamic` | [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-675d01771e7c17a802460e0ff0b9def1fe25fb88bbcb177e51293268d6bfd990) |
| `rules.action.dynamic.elastic_ips` | [rules.action.dynamic.elastic_ips](data-sources--nat_policy--reference--group-001.md#canonical-7398ad1a9e52c319e3a9481328b36d01fc3752e8aa45eba8daf7c6084eb5381e) |
| `rules.action.dynamic.elastic_ips.refs` | [rules.action.dynamic.elastic_ips.refs](data-sources--nat_policy--reference--group-001.md#canonical-394a254ab7285565ae72834bbc3ec617b3306da2b53c8581adbfb6e64bb9f45c) |
| `rules.action.dynamic.elastic_ips.refs.kind` | [rules.action.dynamic.elastic_ips.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-a1a05c65379e48a18a94695f9a1d062d35c16c009a4c0d694377814c3ac8e37b) |
| `rules.action.dynamic.elastic_ips.refs.name` | [rules.action.dynamic.elastic_ips.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-1829f5e778332fd4dde1901cd15f05cb5972b5f1117d342d8da523884ad1da9b) |
| `rules.action.dynamic.elastic_ips.refs.namespace` | [rules.action.dynamic.elastic_ips.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-0661baa2a9f92f355be04a320f21b2e09395cfb9dd2b8f764b33836ee3b72b8e) |
| `rules.action.dynamic.elastic_ips.refs.tenant` | [rules.action.dynamic.elastic_ips.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-88b21c6928c2f6230455c3cd6285f3d7ac37c6a9103c630de8a4877e256cfc6b) |
| `rules.action.dynamic.elastic_ips.refs.uid` | [rules.action.dynamic.elastic_ips.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-84b7c44944c0dfaa129fc39bbbdd47c28c4fdb3d0de08c5d27f0ae30f2622f65) |
| `rules.action.dynamic.pools` | [rules.action.dynamic.pools](data-sources--nat_policy--reference--group-001.md#canonical-cd9cd2ab6d14ca207d19eda7dad06f6b18c773558b5b3f6ececc55b211696bdf) |
| `rules.action.dynamic.pools.prefixes` | [rules.action.dynamic.pools.prefixes](data-sources--nat_policy--reference--group-001.md#canonical-496bd8c0fbeeb71745eea5263e98a726b4b47558c2b007190c8dfaf1428a4583) |
| `rules.action.virtual_cidr` | [rules.action.virtual_cidr](data-sources--nat_policy--reference--group-001.md#canonical-12b540322e5ffaea969e20064893bdabedc1c50493d8ee3167041b884fbd403d) |
| `rules.cloud_connect` | [rules.cloud_connect](data-sources--nat_policy--reference--group-001.md#canonical-fada7afcfe5def21a83f8c87d7a01cbec473371ee02b242877d11c95b1aefad2) |
| `rules.cloud_connect.refs` | [rules.cloud_connect.refs](data-sources--nat_policy--reference--group-001.md#canonical-7e7c5dc7b8a3ebee4ed91e597a6bbecb5ec08275225236e2c95400f17beab4ee) |
| `rules.cloud_connect.refs.kind` | [rules.cloud_connect.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-1e276bb67b73dc166b43896f44e03176d5611dd2af519a75dfbad33018659e99) |
| `rules.cloud_connect.refs.name` | [rules.cloud_connect.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-54ff417b159579259f1c075f3ab3e59b07d39ba6169add957d33072b665c8b2a) |
| `rules.cloud_connect.refs.namespace` | [rules.cloud_connect.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-9ff338ef03c189d2faa3387f07ecd6e7154ed98ba6415629e922d4da80660023) |
| `rules.cloud_connect.refs.tenant` | [rules.cloud_connect.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-14ddbd5da067cdb8e7386d94686437d38afa3a7c1d9620a24fb5b48b91c24880) |
| `rules.cloud_connect.refs.uid` | [rules.cloud_connect.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-bd40ba6cc19feabdf903cfd9ca45dbd2f4c7a38c039491b8f9a4732f92025895) |
| `rules.criteria` | [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-da9bc5b563e3e5129d132eb92d1ee9e5f5a06bc9dede2bc9aaa04f0d633015d1) |
| `rules.criteria.any` | [rules.criteria.any](data-sources--nat_policy--reference--group-001.md#canonical-25e5c031ba32963f2357b46d5d3428c6c7ed6235bce99fc13631bfe7e8e72ff5) |
| `rules.criteria.destination_cidr` | [rules.criteria.destination_cidr](data-sources--nat_policy--reference--group-001.md#canonical-8d8b876e08f2928589ff1f5346bf56c95d61be2d5e7ca535158a1324d3199006) |
| `rules.criteria.icmp` | [rules.criteria.icmp](data-sources--nat_policy--reference--group-001.md#canonical-2b764068cc67fbb794bcd21c01825ed511e212242be311de43f094dccd0f2974) |
| `rules.criteria.site_local_inside_network` | [rules.criteria.site_local_inside_network](data-sources--nat_policy--reference--group-001.md#canonical-9c7cc8cd6e33db8a6e22f764b91edbfcfe4687ad7fed3db1a43b38654ba43a7a) |
| `rules.criteria.site_local_network` | [rules.criteria.site_local_network](data-sources--nat_policy--reference--group-001.md#canonical-fd66aae39133f525153c86f973307f2516dbbb306aa008eeb3b0f5203c92b947) |
| `rules.criteria.source_cidr` | [rules.criteria.source_cidr](data-sources--nat_policy--reference--group-001.md#canonical-b53da93c3f5e140327a58bbedbcb35bd61e9f4021d125adef1f2fa70d65253da) |
| `rules.criteria.tcp` | [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-80763633846414c2a9586f0f55145acbc35fb69fcdc34f2b1277d2e7fcacbaa0) |
| `rules.criteria.tcp.destination_port` | [rules.criteria.tcp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-306ab0529782257fd9fe782d74c90162557c94e17ab8b4c2f7f8cb00b75f33c5) |
| `rules.criteria.tcp.destination_port.no_port_match` | [rules.criteria.tcp.destination_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-0ae6b2edb9cdb4ed42adbdceb1f03d02bcbf5a022cb15a741afc69c8dab67ff2) |
| `rules.criteria.tcp.destination_port.port` | [rules.criteria.tcp.destination_port.port](data-sources--nat_policy--reference--group-001.md#canonical-f7e006498edaa4de595ac7f6a61216d974b310dd0c58386dd7e6ed7a698bd3ad) |
| `rules.criteria.tcp.destination_port.port_ranges` | [rules.criteria.tcp.destination_port.port_ranges](data-sources--nat_policy--reference--group-001.md#canonical-ad924a32864e3b514fecb9685487317c5afb6e95af9c0dd4c53652d17d1839f1) |
| `rules.criteria.tcp.source_port` | [rules.criteria.tcp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-42a584df71a5169a3c312f4dd3dd3193d88ae7e5ae12cf4e7f3a293c1bcc2142) |
| `rules.criteria.tcp.source_port.no_port_match` | [rules.criteria.tcp.source_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-726ac133e91c11d0ad0400a15f0ca5698a3c8e434679445739968c229db62910) |
| `rules.criteria.tcp.source_port.port` | [rules.criteria.tcp.source_port.port](data-sources--nat_policy--reference--group-001.md#canonical-b0db40ad3baa9e073e6be0fdd593d4c0b23f15b528eb83d7ca460f6f99d9dd27) |
| `rules.criteria.tcp.source_port.port_ranges` | [rules.criteria.tcp.source_port.port_ranges](data-sources--nat_policy--reference--group-001.md#canonical-02de2f063134b9e090ff0449d8fc22d5bdbffb452b550eba808454aac1bb6d75) |
| `rules.criteria.udp` | [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-af356fd749638b9fc5edc0550eca4559fd3ff702c227d487594f476f4721432f) |
| `rules.criteria.udp.destination_port` | [rules.criteria.udp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-bdba660aee91f1fd8ffc30c5a14ff3cd6211d5fa0b28cabb6ff429d699d12043) |
| `rules.criteria.udp.destination_port.no_port_match` | [rules.criteria.udp.destination_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-e95b4398177c481e13ea35f23ecebc2f9d391e89cbb2e6f3361183144dc107e7) |
| `rules.criteria.udp.destination_port.port` | [rules.criteria.udp.destination_port.port](data-sources--nat_policy--reference--group-001.md#canonical-a56919d49549e7314ebf33c8027fe7300096bc851104f638eb28ea797e00baf0) |
| `rules.criteria.udp.destination_port.port_ranges` | [rules.criteria.udp.destination_port.port_ranges](data-sources--nat_policy--reference--group-001.md#canonical-c5904fc8bdf8f5af20bfaf78509726a215643a9a284f726e9d0f9befa61b9556) |
| `rules.criteria.udp.source_port` | [rules.criteria.udp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-9b1ff8ddcac2b463c02da40d4789634c381b0e9ae88369a248c2ba9c8dac57b4) |
| `rules.criteria.udp.source_port.no_port_match` | [rules.criteria.udp.source_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-33811b1c5a1084af454940e05d2f70bfc8b1fffc399ad4e447be3ae342b46c56) |
| `rules.criteria.udp.source_port.port` | [rules.criteria.udp.source_port.port](data-sources--nat_policy--reference--group-001.md#canonical-fabc27657bf7405e78daed60e351336b87f12be16b0ad3cfcf3df86630d5fd97) |
| `rules.criteria.udp.source_port.port_ranges` | [rules.criteria.udp.source_port.port_ranges](data-sources--nat_policy--reference--group-001.md#canonical-c3eaf1a283859f06d560cd74599c08fe41824198e3d279959a001ee6e469df4f) |
| `rules.disable_spec` | [rules.disable_spec](data-sources--nat_policy--reference--group-001.md#canonical-88c9e9fbde7468b5fc33c3f9313d129c854e568548a3407bc0b8c8d66d986e9a) |
| `rules.enable` | [rules.enable](data-sources--nat_policy--reference--group-001.md#canonical-1a3408b30acaf4e386b89b43606584666fa292cd42261ab4cddb2b5d5d3d6a5e) |
| `rules.name` | [rules.name](data-sources--nat_policy--reference--group-001.md#canonical-575ab2e88313af4c5ca1a3a63b13cdf43738dec3f586df1616fcb38f2f8e737b) |
| `rules.node_interface` | [rules.node_interface](data-sources--nat_policy--reference--group-001.md#canonical-d5873746b2f5e4d0b1a3b689f123d0d82a5ac496279b298b2bd1310ad2b25b76) |
| `rules.node_interface.list` | [rules.node_interface.list](data-sources--nat_policy--reference--group-001.md#canonical-cdedbf90e8f4582f89c2fbde6b458947a1d239972e7e31777d74de3db05dd557) |
| `rules.node_interface.list.interface` | [rules.node_interface.list.interface](data-sources--nat_policy--reference--group-001.md#canonical-1bf0558418267940feb9960756fcdaf1775347de6f208b507c3b80be22de23aa) |
| `rules.node_interface.list.interface.kind` | [rules.node_interface.list.interface.kind](data-sources--nat_policy--reference--group-001.md#canonical-c036f81a9463902e126525d539e2937d9cf185acde07c3525e00c18d83b0f9b5) |
| `rules.node_interface.list.interface.name` | [rules.node_interface.list.interface.name](data-sources--nat_policy--reference--group-001.md#canonical-7d5bf562ed1263bd1d9f2380eb8647bf7fb5d5b78807831d150964ed92cb4cf3) |
| `rules.node_interface.list.interface.namespace` | [rules.node_interface.list.interface.namespace](data-sources--nat_policy--reference--group-001.md#canonical-dc6aad8f081d6cbf6c4dd5f631b93cd9c491747ffcb1b3918f5e1ac3b772bab1) |
| `rules.node_interface.list.interface.tenant` | [rules.node_interface.list.interface.tenant](data-sources--nat_policy--reference--group-001.md#canonical-686f589c846e31e0170be955f508d3c4ed7be265ace5e11362e8441a2ab336b3) |
| `rules.node_interface.list.interface.uid` | [rules.node_interface.list.interface.uid](data-sources--nat_policy--reference--group-001.md#canonical-9d0609ea0f1d60f919317ed20250667da31abce1364b3e00d4a9b6fcc603a4f2) |
| `rules.node_interface.list.node` | [rules.node_interface.list.node](data-sources--nat_policy--reference--group-001.md#canonical-d30febadbd5dfdff686a294fec43c10245ded2efebac297faab894b89084f0ca) |
| `rules.segment` | [rules.segment](data-sources--nat_policy--reference--group-001.md#canonical-391d1b8612c574c88f42cc55473b80a985f67cbd4c9194cab7b55d0655131507) |
| `rules.segment.refs` | [rules.segment.refs](data-sources--nat_policy--reference--group-001.md#canonical-b36d5d6d5572b52a5470cc80fd83a0aa4e6f6b4a0e48c5f66a94d7bc00fb7d9b) |
| `rules.segment.refs.kind` | [rules.segment.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-e92cf2cd7da2ca7f36cd1d640769187dec22392e6bab8cc1a6444162ff2f8e4c) |
| `rules.segment.refs.name` | [rules.segment.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-812dd0ac8e84b14a922b376f80ed54fa0e7dc90fa976af60eccb561e755cc645) |
| `rules.segment.refs.namespace` | [rules.segment.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-47caa2540e372de16221730512f7c80ca90b62f411659aba5165f872e49eb92e) |
| `rules.segment.refs.tenant` | [rules.segment.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-dba3dd7a2ca885d6f0e9cc985f9c9bc91f0a530b6b4da0d7faf354bf216667c6) |
| `rules.segment.refs.uid` | [rules.segment.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-942da4a7e7fe27a087187994da85ad9a61405f580e8928ceeb4d667e9314059b) |
| `rules.virtual_network` | [rules.virtual_network](data-sources--nat_policy--reference--group-001.md#canonical-cbddabff3be16256f397b8ce641364f4e761b08a6569483538c831104e69f58a) |
| `rules.virtual_network.refs` | [rules.virtual_network.refs](data-sources--nat_policy--reference--group-001.md#canonical-dc139ddfa35fac76edfabe18064b992a352e417347492380ac96457f79f7c814) |
| `rules.virtual_network.refs.kind` | [rules.virtual_network.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-dfa84393b6aba3d05dda73bd5ada6de1a61340ff957211478c5f5c5c2c22b629) |
| `rules.virtual_network.refs.name` | [rules.virtual_network.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-a6b57d8ed5e53112bee299a8a222bde4091cef6ff6088e39a9cec404c4c8a408) |
| `rules.virtual_network.refs.namespace` | [rules.virtual_network.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-d9e9e186c070c5c324255920345899595d10661c86c59a0e4433a40f18669029) |
| `rules.virtual_network.refs.tenant` | [rules.virtual_network.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-07a380f925223dff4a10bd416f0979c8dd627b68dcad6c0b4f30c6167a057021) |
| `rules.virtual_network.refs.uid` | [rules.virtual_network.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-75829fa69cdad677a11fdbc95c692574c97ddd2313f69b0f56687e9ee1509381) |
| `site` | [site](data-sources--nat_policy--reference--group-001.md#canonical-77ffd0e49759954f42b9c58321477d4bf44abda43c11fd8f72a18b9eb782d9d1) |
| `site.refs` | [site.refs](data-sources--nat_policy--reference--group-001.md#canonical-d7c609af0cc1b8007802c33c5f50ff03ec9a44338726c55c9c41fdf878264bf3) |
| `site.refs.kind` | [site.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-b1ba199a1a0964c210a184f37c95b55a5bf3e156a097e142ea59745bde3b3080) |
| `site.refs.name` | [site.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-eff62a7af5df1091c52d18c1f3f85d698de144c691508e3e5bc5c8a246142de4) |
| `site.refs.namespace` | [site.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-5161eea35ad523c5f813bed60b5893588aa2728229583e6cd83234c7640b44dd) |
| `site.refs.tenant` | [site.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-4d2418860e418bd1010d0af20af40eb383e69bf12822a08f9cd8b43669d4295a) |
| `site.refs.uid` | [site.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-114d73cd64913e78479af865b2faeac3bec2b8eb155f9c42538abed10db15287) |

<a id="canonical-24fe284c337a2daeb690352de90a2b5fa6b10a2b5821a6cdff3ec8e189ea604d"></a>

## Next pages — Property reference / 41f75688a175 / 11

- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [site](data-sources--nat_policy--reference--group-001.md#canonical-45c6954ad1fe44c2f95f33065ed54ee51b8ee9ad280c257c2c757d100b9b89d4)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31d4b5c1cffcfa18f37a2100d39a25879d5e5938b6d54f195c9d6d3c03f9e5b7"></a>

## rules — rules / 61beb1842d93 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- rules

<a id="canonical-dafa45a5831de76a4c8bbf730e92b17f46da2c46c77196428139641e99573e19"></a>

Type: `"list"`. Computed.

List of rules to apply under the NAT Policy. Rule that matches first would be applied.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

<a id="canonical-6db1dacf675cbc7fe07a71a6912af54e6f7d15b03d76a6f654f8332503a2378c"></a>

## Direct properties — rules / 61beb1842d93 / 3

- [action](data-sources--nat_policy--reference--group-001.md#canonical-57395017dadb349df1fa0f634e692c7ccbab391bd784103196451d95e5e26ce8): complete subsection reference.

- [cloud_connect](data-sources--nat_policy--reference--group-001.md#canonical-82fa5def3be374335167fbaa29134b4b1425da8c402e7f210b2e6617ec7abc0c): complete subsection reference.

- [criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb): complete subsection reference.

- [disable_spec](data-sources--nat_policy--reference--group-001.md#canonical-a4e3b0ae0e9db9059d9e05237babd022505e67d8a50020ea1ac6e1d061e60e76): complete subsection reference.

- [enable](data-sources--nat_policy--reference--group-001.md#canonical-7e494423c38e6c628469aebfa17b674ef405e452a32abbffbbaa5f0b494e5af3): complete subsection reference.

<a id="canonical-575ab2e88313af4c5ca1a3a63b13cdf43738dec3f586df1616fcb38f2f8e737b"></a>

<a id="canonical-5bd14028480f51363fce6748a01c21f94b3e9a2108cfa5f28cdcdb5f101107ea"></a>

## name property — rules / 61beb1842d93 / 4

Type: `"string"`. Computed.

Name. Name of the Rule.

Upstream description:

Name of the Rule.

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

- [node_interface](data-sources--nat_policy--reference--group-001.md#canonical-d27f3d9d29707b5ddf7466df4f1f8d4aa1f3a0f4ee55814bc229ea0b0ebde515): complete subsection reference.

- [segment](data-sources--nat_policy--reference--group-001.md#canonical-dabe55b282aee5012d817c9b22a7d59b8bde82652e4797c4c03a6c31eadd2f75): complete subsection reference.

- [virtual_network](data-sources--nat_policy--reference--group-001.md#canonical-6a7c35e7a50dfb77cc131a8a4eba991f9f38f74e9f6981b6d6faf384ce6e15e9): complete subsection reference.

<a id="canonical-b4503d512cf38446b8dd067813dee36e3ff6c12c6a13c5f0e4e39649caf617ca"></a>

## Next pages — rules / 61beb1842d93 / 5

- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-57395017dadb349df1fa0f634e692c7ccbab391bd784103196451d95e5e26ce8)
- [rules.cloud_connect](data-sources--nat_policy--reference--group-001.md#canonical-82fa5def3be374335167fbaa29134b4b1425da8c402e7f210b2e6617ec7abc0c)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [rules.disable_spec](data-sources--nat_policy--reference--group-001.md#canonical-a4e3b0ae0e9db9059d9e05237babd022505e67d8a50020ea1ac6e1d061e60e76)
- [rules.enable](data-sources--nat_policy--reference--group-001.md#canonical-7e494423c38e6c628469aebfa17b674ef405e452a32abbffbbaa5f0b494e5af3)
- [rules.node_interface](data-sources--nat_policy--reference--group-001.md#canonical-d27f3d9d29707b5ddf7466df4f1f8d4aa1f3a0f4ee55814bc229ea0b0ebde515)
- [rules.segment](data-sources--nat_policy--reference--group-001.md#canonical-dabe55b282aee5012d817c9b22a7d59b8bde82652e4797c4c03a6c31eadd2f75)
- [rules.virtual_network](data-sources--nat_policy--reference--group-001.md#canonical-6a7c35e7a50dfb77cc131a8a4eba991f9f38f74e9f6981b6d6faf384ce6e15e9)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-57395017dadb349df1fa0f634e692c7ccbab391bd784103196451d95e5e26ce8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-803199a37261e935fff7afb9d797153095e7d8bfb3b8da944d19adb7c8d0718f"></a>

## rules.action — rules.action / a9b1b9bdbb3d / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- rules.action

<a id="canonical-dc742ea365aa3bce75c14b1ac664c62ad7b656b95d238191184f555f029b3a28"></a>

Type: `"single"`. Computed.

Action to apply on the packet if the NAT rule is applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_nat_choice": "[\"dynamic\",\"virtual_cidr\"]"
}
```

<a id="canonical-ceaed26f2fdbc4adca9939c8cbe193db33e930cedebf604e7ba346a4f063059a"></a>

## Direct properties — rules.action / a9b1b9bdbb3d / 3

- [dynamic](data-sources--nat_policy--reference--group-001.md#canonical-95dc6c9003cb28a6a7a8b38f783b413a25dc1854c03a93be73714467bf93c962): complete subsection reference.

<a id="canonical-12b540322e5ffaea969e20064893bdabedc1c50493d8ee3167041b884fbd403d"></a>

<a id="canonical-fbbc24407739ac29d72685cc2015080704bb1c55a27a99ea02e3106f6bc17515"></a>

## virtual_cidr property — rules.action / a9b1b9bdbb3d / 4

Type: `"string"`. Computed.

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR..

Upstream description:

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR
192.0.2.0/24, the virtual CIDR has 100.100.100.0/24.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-1dd1220f6965da8bbadf8c3d2ccdc5954f6867d174ef3129c80282187f89ed6c"></a>

## Next pages — rules.action / a9b1b9bdbb3d / 5

- [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-95dc6c9003cb28a6a7a8b38f783b413a25dc1854c03a93be73714467bf93c962)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-95dc6c9003cb28a6a7a8b38f783b413a25dc1854c03a93be73714467bf93c962"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b18b0313f785abb647aa119f643f4ba6a8b3515e741f0c2cd0be3ac7017d577a"></a>

## rules.action.dynamic — rules.action.dynamic / e78b01e75a75 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-57395017dadb349df1fa0f634e692c7ccbab391bd784103196451d95e5e26ce8)
- rules.action.dynamic

<a id="canonical-675d01771e7c17a802460e0ff0b9def1fe25fb88bbcb177e51293268d6bfd990"></a>

Type: `"single"`. Computed.

Dynamic Pool. Dynamic Pool Configuration.

Upstream description:

Dynamic Pool Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-pool_choice": "[\"elastic_ips\",\"pools\"]"
}
```

<a id="canonical-a71fb61e65adb5dfcbe16acfe9f389e1969bdc64b2ec0a73eb1be0036571879a"></a>

## Direct properties — rules.action.dynamic / e78b01e75a75 / 3

- [elastic_ips](data-sources--nat_policy--reference--group-001.md#canonical-2c0e8cdf6c2b7eb980190de0d8341301aa2c0fcabd5c89913bd053fa87c05db8): complete subsection reference.

- [pools](data-sources--nat_policy--reference--group-001.md#canonical-b11bbbdedf378496f1896a91054f7534d34fd9c1c5d0be29fe9b7ee21be51581): complete subsection reference.

<a id="canonical-8ee9db2d99c3f120f181d0153dc1ec727d87759eb764d05dfbdbc4e4ef597967"></a>

## Next pages — rules.action.dynamic / e78b01e75a75 / 4

- [rules.action.dynamic.elastic_ips](data-sources--nat_policy--reference--group-001.md#canonical-2c0e8cdf6c2b7eb980190de0d8341301aa2c0fcabd5c89913bd053fa87c05db8)
- [rules.action.dynamic.pools](data-sources--nat_policy--reference--group-001.md#canonical-b11bbbdedf378496f1896a91054f7534d34fd9c1c5d0be29fe9b7ee21be51581)
- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-57395017dadb349df1fa0f634e692c7ccbab391bd784103196451d95e5e26ce8)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-2c0e8cdf6c2b7eb980190de0d8341301aa2c0fcabd5c89913bd053fa87c05db8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebbc155c909635a28b2d2de3653ee3ba399fdc2aaf412303c0d3777f7ced35d0"></a>

## rules.action.dynamic.elastic_ips — rules.action.dynamic.elastic_ips / f9e6d2a06dd1 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-57395017dadb349df1fa0f634e692c7ccbab391bd784103196451d95e5e26ce8)
- [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-95dc6c9003cb28a6a7a8b38f783b413a25dc1854c03a93be73714467bf93c962)
- rules.action.dynamic.elastic_ips

<a id="canonical-7398ad1a9e52c319e3a9481328b36d01fc3752e8aa45eba8daf7c6084eb5381e"></a>

Type: `"single"`. Computed.

List of references to Cloud Elastic IP Object.

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

<a id="canonical-623b5297fc8db4b5a007b60abc85a1608b3c4df331ec9f7bff6f05cdc3e187d0"></a>

## Direct properties — rules.action.dynamic.elastic_ips / f9e6d2a06dd1 / 3

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-f7292b05debf0f99eaf01fa3c3d6c0e0953692eed2f2901d4d356ced1e6fe2ea): complete subsection reference.

<a id="canonical-51a14f9a8004d53ac72a184a01a1fb38507bfecffbe7945640a160118b2fbc44"></a>

## Next pages — rules.action.dynamic.elastic_ips / f9e6d2a06dd1 / 4

- [rules.action.dynamic.elastic_ips.refs](data-sources--nat_policy--reference--group-001.md#canonical-f7292b05debf0f99eaf01fa3c3d6c0e0953692eed2f2901d4d356ced1e6fe2ea)
- [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-95dc6c9003cb28a6a7a8b38f783b413a25dc1854c03a93be73714467bf93c962)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-f7292b05debf0f99eaf01fa3c3d6c0e0953692eed2f2901d4d356ced1e6fe2ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ed11b3cbbc756bebdff7841c6d40864bd351cc1d587b28a4ab5ef9828a2e238"></a>

## rules.action.dynamic.elastic_ips.refs — rules.action.dynamic.elastic_ips.refs / aba4dc907dc4 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-57395017dadb349df1fa0f634e692c7ccbab391bd784103196451d95e5e26ce8)
- [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-95dc6c9003cb28a6a7a8b38f783b413a25dc1854c03a93be73714467bf93c962)
- [rules.action.dynamic.elastic_ips](data-sources--nat_policy--reference--group-001.md#canonical-2c0e8cdf6c2b7eb980190de0d8341301aa2c0fcabd5c89913bd053fa87c05db8)
- rules.action.dynamic.elastic_ips.refs

<a id="canonical-394a254ab7285565ae72834bbc3ec617b3306da2b53c8581adbfb6e64bb9f45c"></a>

Type: `"list"`. Computed.

Reference to one or more cloud elastic IP objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-5b8f3867a5ec0dc7749551f7de6c09879a68d2e6591680c312d6deedf8d482a0"></a>

## Direct properties — rules.action.dynamic.elastic_ips.refs / aba4dc907dc4 / 3

<a id="canonical-a1a05c65379e48a18a94695f9a1d062d35c16c009a4c0d694377814c3ac8e37b"></a>

<a id="canonical-8c54a0927f88987f2cddbbed89273371c810e31b9181618a8a5e34942f4403e9"></a>

## kind property — rules.action.dynamic.elastic_ips.refs / aba4dc907dc4 / 4

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

<a id="canonical-1829f5e778332fd4dde1901cd15f05cb5972b5f1117d342d8da523884ad1da9b"></a>

<a id="canonical-d07e45a7ca34c272ea2a86f2c6519432b33508850884c0695c87548f88162a7b"></a>

## name property — rules.action.dynamic.elastic_ips.refs / aba4dc907dc4 / 5

Type: `"string"`. Computed.

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

<a id="canonical-0661baa2a9f92f355be04a320f21b2e09395cfb9dd2b8f764b33836ee3b72b8e"></a>

<a id="canonical-968438bc3edf93853e193b037d187347a1d6e01466b2f5af3920a56d8fae6119"></a>

## namespace property — rules.action.dynamic.elastic_ips.refs / aba4dc907dc4 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-88b21c6928c2f6230455c3cd6285f3d7ac37c6a9103c630de8a4877e256cfc6b"></a>

<a id="canonical-11fea6e1b6b0222bb7a39a384aea204692faa9ee7d077748efdb99d0f7ea4e10"></a>

## tenant property — rules.action.dynamic.elastic_ips.refs / aba4dc907dc4 / 7

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

<a id="canonical-84b7c44944c0dfaa129fc39bbbdd47c28c4fdb3d0de08c5d27f0ae30f2622f65"></a>

<a id="canonical-1137c525a2b679dfa77c8aacb45da1f83ae596bc4b7718d5007d4c12edd2b094"></a>

## uid property — rules.action.dynamic.elastic_ips.refs / aba4dc907dc4 / 8

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

<a id="canonical-9ca1881bdf7ceac9b39216422582b5bbbbbcfb66fd4bdc1ce027e469c9bc6d1a"></a>

## Next pages — rules.action.dynamic.elastic_ips.refs / aba4dc907dc4 / 9

- [rules.action.dynamic.elastic_ips](data-sources--nat_policy--reference--group-001.md#canonical-2c0e8cdf6c2b7eb980190de0d8341301aa2c0fcabd5c89913bd053fa87c05db8)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-b11bbbdedf378496f1896a91054f7534d34fd9c1c5d0be29fe9b7ee21be51581"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6b5505076517e02c0ffa0b6943b7e3ad87e26473caea3be9524dc844c6366b0"></a>

## rules.action.dynamic.pools — rules.action.dynamic.pools / 3bd4d3ff0ab6 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-57395017dadb349df1fa0f634e692c7ccbab391bd784103196451d95e5e26ce8)
- [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-95dc6c9003cb28a6a7a8b38f783b413a25dc1854c03a93be73714467bf93c962)
- rules.action.dynamic.pools

<a id="canonical-cd9cd2ab6d14ca207d19eda7dad06f6b18c773558b5b3f6ececc55b211696bdf"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-25ab8b4103a789a92edbaaf19f3dcf91ceb547fd008741eca1b61ea294974ebb"></a>

## Direct properties — rules.action.dynamic.pools / 3bd4d3ff0ab6 / 3

<a id="canonical-496bd8c0fbeeb71745eea5263e98a726b4b47558c2b007190c8dfaf1428a4583"></a>

<a id="canonical-fadf4f335de3f298ee0d51855290a498531d1a3aa99ef7a3fb876beedef6f5e7"></a>

## prefixes property — rules.action.dynamic.pools / 3bd4d3ff0ab6 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-eb95f79df512846fb53bac22a1841415089948fcaa73ead4d5ffdf42e2fb0f99"></a>

## Next pages — rules.action.dynamic.pools / 3bd4d3ff0ab6 / 5

- [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-95dc6c9003cb28a6a7a8b38f783b413a25dc1854c03a93be73714467bf93c962)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-82fa5def3be374335167fbaa29134b4b1425da8c402e7f210b2e6617ec7abc0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1696e0d445220bcee400cba1bf79643ae2d2be63b0abb96a279c07ca56b304d"></a>

## rules.cloud_connect — rules.cloud_connect / b39e964e57b1 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- rules.cloud_connect

<a id="canonical-fada7afcfe5def21a83f8c87d7a01cbec473371ee02b242877d11c95b1aefad2"></a>

Type: `"single"`. Computed.

Configuration parameter for cloud connect.

Upstream description:

Reference to Cloud connect Object.

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

<a id="canonical-732b8e22d7014303c9130678ff4151878cff9e088b6fe94a01da2f9daf3e9b98"></a>

## Direct properties — rules.cloud_connect / b39e964e57b1 / 3

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-a80bc8563b5a9bedf58a9959750691dfa642f96ab4b449eaa8c0925ba961ae72): complete subsection reference.

<a id="canonical-88594834db61e65b595f79d1da2c53f428553e050ffea363fed56780017fe0a2"></a>

## Next pages — rules.cloud_connect / b39e964e57b1 / 4

- [rules.cloud_connect.refs](data-sources--nat_policy--reference--group-001.md#canonical-a80bc8563b5a9bedf58a9959750691dfa642f96ab4b449eaa8c0925ba961ae72)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-a80bc8563b5a9bedf58a9959750691dfa642f96ab4b449eaa8c0925ba961ae72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a3881e954dd87a9c0a9500ff725a791fb4787369b8fcec29cb4aa15a118d774"></a>

## rules.cloud_connect.refs — rules.cloud_connect.refs / e95eb60581ca / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.cloud_connect](data-sources--nat_policy--reference--group-001.md#canonical-82fa5def3be374335167fbaa29134b4b1425da8c402e7f210b2e6617ec7abc0c)
- rules.cloud_connect.refs

<a id="canonical-7e7c5dc7b8a3ebee4ed91e597a6bbecb5ec08275225236e2c95400f17beab4ee"></a>

Type: `"list"`. Computed.

Cloud Connect. Reference to Cloud Connect Object.

Upstream description:

Reference to Cloud Connect Object.

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

<a id="canonical-1e62b5812e60beb3439677f0438e3632bdc19b5a2e5564e059e73e20afe5f71e"></a>

## Direct properties — rules.cloud_connect.refs / e95eb60581ca / 3

<a id="canonical-1e276bb67b73dc166b43896f44e03176d5611dd2af519a75dfbad33018659e99"></a>

<a id="canonical-8ff4e0b5f34156fa95ac994331614962dbf074af62f4eebb56101ea2fdff9708"></a>

## kind property — rules.cloud_connect.refs / e95eb60581ca / 4

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

<a id="canonical-54ff417b159579259f1c075f3ab3e59b07d39ba6169add957d33072b665c8b2a"></a>

<a id="canonical-b5e71406c79ee42fbf48585e8c0af33465b1306ac1dab26c45f62acab054bb2f"></a>

## name property — rules.cloud_connect.refs / e95eb60581ca / 5

Type: `"string"`. Computed.

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

<a id="canonical-9ff338ef03c189d2faa3387f07ecd6e7154ed98ba6415629e922d4da80660023"></a>

<a id="canonical-9cdbb171ac049be213ee06bfecc7307363b82023e8b14b1769286048b762a230"></a>

## namespace property — rules.cloud_connect.refs / e95eb60581ca / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-14ddbd5da067cdb8e7386d94686437d38afa3a7c1d9620a24fb5b48b91c24880"></a>

<a id="canonical-a339a0ef401d4c1603b1787c221e507e2e061371155f1b1469b199df8bbdd0bb"></a>

## tenant property — rules.cloud_connect.refs / e95eb60581ca / 7

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

<a id="canonical-bd40ba6cc19feabdf903cfd9ca45dbd2f4c7a38c039491b8f9a4732f92025895"></a>

<a id="canonical-687c0f53e09093982c5c4beacd54d0556a81307ace3209f3680679d9540ce806"></a>

## uid property — rules.cloud_connect.refs / e95eb60581ca / 8

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

<a id="canonical-8dd070193e835328d3c638f01b0e780a8d1410447a8f50be849190ed8709f109"></a>

## Next pages — rules.cloud_connect.refs / e95eb60581ca / 9

- [rules.cloud_connect](data-sources--nat_policy--reference--group-001.md#canonical-82fa5def3be374335167fbaa29134b4b1425da8c402e7f210b2e6617ec7abc0c)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04bf58c7de9819f9ea5ecf0381371fd76c63aacb28e73089bf2c7be244672c4c"></a>

## rules.criteria — rules.criteria / fe715b451713 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- rules.criteria

<a id="canonical-da9bc5b563e3e5129d132eb92d1ee9e5f5a06bc9dede2bc9aaa04f0d633015d1"></a>

Type: `"single"`. Computed.

Match criteria of the packet to apply the NAT Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-protocol_choice": "[\"any\",\"icmp\",\"tcp\",\"udp\"]"
}
```

<a id="canonical-ac39989eaa13597c68a7ed7644944f37292f50b13105141182b4abaa76a2eb68"></a>

## Direct properties — rules.criteria / fe715b451713 / 3

- [any](data-sources--nat_policy--reference--group-001.md#canonical-d5cebc900fd576f74ff3cd0f7e448aeeec855e7bd2049e070095c6f290c88f6b): complete subsection reference.

<a id="canonical-8d8b876e08f2928589ff1f5346bf56c95d61be2d5e7ca535158a1324d3199006"></a>

<a id="canonical-f5398620c05bb144778aa20858c8f53c11705664034da0de130163800c3e05b3"></a>

## destination_cidr property — rules.criteria / fe715b451713 / 4

Type: `["list", "string"]`. Computed.

Destination IP. Destination IP of the packet to match.

Upstream description:

Destination IP of the packet to match.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [icmp](data-sources--nat_policy--reference--group-001.md#canonical-a51c63334256a0c38a8e44725756fa45e33f97b64d262fd9067b655ced1ec0f7): complete subsection reference.

- [site_local_inside_network](data-sources--nat_policy--reference--group-001.md#canonical-af6cf36af097c2abc0be80f7d72e2a932f4bb897b31cf27c9f6e27b7d3e5ca43): complete subsection reference.

- [site_local_network](data-sources--nat_policy--reference--group-001.md#canonical-17dd46d5dde0ab0c7faf3e5e5eb883a31e82db8dc5b612ee8c9f23e98f9cb107): complete subsection reference.

<a id="canonical-b53da93c3f5e140327a58bbedbcb35bd61e9f4021d125adef1f2fa70d65253da"></a>

<a id="canonical-929805c1210f2f2ed654fea079011bb31ad4c02b98c20701a20a4588e77b4c11"></a>

## source_cidr property — rules.criteria / fe715b451713 / 5

Type: `["list", "string"]`. Computed.

Source IP. Source IP of the packet to match.

Upstream description:

Source IP of the packet to match.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [tcp](data-sources--nat_policy--reference--group-001.md#canonical-323262920a8f8a12b1ce22ae6f9acd70db7913fd17d861916089c468de577b27): complete subsection reference.

- [udp](data-sources--nat_policy--reference--group-001.md#canonical-cc94026d041bb1a9a760ce463f9899d1055eaf3a9087642efaa0b3052ed199bd): complete subsection reference.

<a id="canonical-be25e35b420e5b893cecc1d565c618a231d00a2bf22aa43a5c26c7fffb13bef4"></a>

## Next pages — rules.criteria / fe715b451713 / 6

- [rules.criteria.any](data-sources--nat_policy--reference--group-001.md#canonical-d5cebc900fd576f74ff3cd0f7e448aeeec855e7bd2049e070095c6f290c88f6b)
- [rules.criteria.icmp](data-sources--nat_policy--reference--group-001.md#canonical-a51c63334256a0c38a8e44725756fa45e33f97b64d262fd9067b655ced1ec0f7)
- [rules.criteria.site_local_inside_network](data-sources--nat_policy--reference--group-001.md#canonical-af6cf36af097c2abc0be80f7d72e2a932f4bb897b31cf27c9f6e27b7d3e5ca43)
- [rules.criteria.site_local_network](data-sources--nat_policy--reference--group-001.md#canonical-17dd46d5dde0ab0c7faf3e5e5eb883a31e82db8dc5b612ee8c9f23e98f9cb107)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-323262920a8f8a12b1ce22ae6f9acd70db7913fd17d861916089c468de577b27)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-cc94026d041bb1a9a760ce463f9899d1055eaf3a9087642efaa0b3052ed199bd)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-d5cebc900fd576f74ff3cd0f7e448aeeec855e7bd2049e070095c6f290c88f6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bc65495af80aee02322aa0c400f49ba476ad5891da7fa53b82fcf5154e3fd63"></a>

## rules.criteria.any — rules.criteria.any / eafbebe44bc9 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- rules.criteria.any

<a id="canonical-25e5c031ba32963f2357b46d5d3428c6c7ed6235bce99fc13631bfe7e8e72ff5"></a>

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

<a id="canonical-4fbc5f1faed176ec9ce90fbd89ba33d8c8c509a79b8acad4736e696fc41b4a9b"></a>

## Direct properties — rules.criteria.any / eafbebe44bc9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fad3e2d8ae6565eab1011c6c1514361c426c7d5975f665a97fda591e3a20107d"></a>

## Next pages — rules.criteria.any / eafbebe44bc9 / 4

- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-a51c63334256a0c38a8e44725756fa45e33f97b64d262fd9067b655ced1ec0f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd0bca44355a63ab30533e4197774c6258afcea2dac47efe422cf35a25320015"></a>

## rules.criteria.icmp — rules.criteria.icmp / bfa0f5f1518e / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- rules.criteria.icmp

<a id="canonical-2b764068cc67fbb794bcd21c01825ed511e212242be311de43f094dccd0f2974"></a>

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

<a id="canonical-3f095ab5facc0d2693e6a8b6e4041b603076b2338100446d66da185554901708"></a>

## Direct properties — rules.criteria.icmp / bfa0f5f1518e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4e429d19d4fa353c4905098a918e8768723db7684ae1c3796495a1d727d6ff22"></a>

## Next pages — rules.criteria.icmp / bfa0f5f1518e / 4

- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-af6cf36af097c2abc0be80f7d72e2a932f4bb897b31cf27c9f6e27b7d3e5ca43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7302506b5097ab8d077508ed59eac080c0dcfd8cd69bc2e9f80774fbf722701a"></a>

## rules.criteria.site_local_inside_network — rules.criteria.site_local_inside_network / d3c6ed178ed2 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- rules.criteria.site_local_inside_network

<a id="canonical-9c7cc8cd6e33db8a6e22f764b91edbfcfe4687ad7fed3db1a43b38654ba43a7a"></a>

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

<a id="canonical-c98ff4fd9a26c76dcfeb7183b42fc07cbfaf6996a95b5ebb2e31c0b54875ab25"></a>

## Direct properties — rules.criteria.site_local_inside_network / d3c6ed178ed2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d224ee0c8a016d086c4eb09cb1882b652e1c921e987985608887fbc58c078c8"></a>

## Next pages — rules.criteria.site_local_inside_network / d3c6ed178ed2 / 4

- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-17dd46d5dde0ab0c7faf3e5e5eb883a31e82db8dc5b612ee8c9f23e98f9cb107"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f41f48a1aa5b915f0be994b2051d4d501486e97d25f6cb8fb73988c2b39dec34"></a>

## rules.criteria.site_local_network — rules.criteria.site_local_network / f21612fc3f7c / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- rules.criteria.site_local_network

<a id="canonical-fd66aae39133f525153c86f973307f2516dbbb306aa008eeb3b0f5203c92b947"></a>

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

<a id="canonical-d303b3ca69292dabdf99be825e02a8c83baa2bc4bc3c5f0abf1c1b426523e059"></a>

## Direct properties — rules.criteria.site_local_network / f21612fc3f7c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e81b2d55c196bd96e4aa62f1e8b3081462cb1867fbe9adc88d40fe1fc28a9fe3"></a>

## Next pages — rules.criteria.site_local_network / f21612fc3f7c / 4

- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-323262920a8f8a12b1ce22ae6f9acd70db7913fd17d861916089c468de577b27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da40987d440343a4f4a4ae43c02e408b469d54808f18616fb533b1beae93ae88"></a>

## rules.criteria.tcp — rules.criteria.tcp / 287c668f16f5 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- rules.criteria.tcp

<a id="canonical-80763633846414c2a9586f0f55145acbc35fb69fcdc34f2b1277d2e7fcacbaa0"></a>

Type: `"single"`. Computed.

Action to apply on the packet if the NAT rule is applied.

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

<a id="canonical-d9ecbd778f6b6f8da5e2c211583a76b333ab19d1787a1162e3cc4c3909c2e864"></a>

## Direct properties — rules.criteria.tcp / 287c668f16f5 / 3

- [destination_port](data-sources--nat_policy--reference--group-001.md#canonical-453a8146986fa706616a25c3b67cc29e66c6b5348070240f2f411d9e479c92f6): complete subsection reference.

- [source_port](data-sources--nat_policy--reference--group-001.md#canonical-1829f31851bda2f6c53dd4f7f2ee64b42d010606823ecc244bfddd4c04598e51): complete subsection reference.

<a id="canonical-f1a3336f1d5b610aceaef2eb055f4821888678dddc6b0b960400652185b02205"></a>

## Next pages — rules.criteria.tcp / 287c668f16f5 / 4

- [rules.criteria.tcp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-453a8146986fa706616a25c3b67cc29e66c6b5348070240f2f411d9e479c92f6)
- [rules.criteria.tcp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-1829f31851bda2f6c53dd4f7f2ee64b42d010606823ecc244bfddd4c04598e51)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-453a8146986fa706616a25c3b67cc29e66c6b5348070240f2f411d9e479c92f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06632b0a8a1f7966abbfd124af45650d3ee7cd3c619220bb61600a1f56850f6a"></a>

## rules.criteria.tcp.destination_port — rules.criteria.tcp.destination_port / d96c469e490a / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-323262920a8f8a12b1ce22ae6f9acd70db7913fd17d861916089c468de577b27)
- rules.criteria.tcp.destination_port

<a id="canonical-306ab0529782257fd9fe782d74c90162557c94e17ab8b4c2f7f8cb00b75f33c5"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-e0432e13612398a35dfbbdee161e7339d5dc33adab472cdbaab2c6bcf71c5805"></a>

## Direct properties — rules.criteria.tcp.destination_port / d96c469e490a / 3

- [no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-9b56fa2772143d67d704037b17dccfd04b4cc76359f102a270b37b740002548a): complete subsection reference.

<a id="canonical-f7e006498edaa4de595ac7f6a61216d974b310dd0c58386dd7e6ed7a698bd3ad"></a>

<a id="canonical-a47d9098e2a34f267d6c3c2a1fe794f0c6f9055be46d0512636a26a841ab2113"></a>

## port property — rules.criteria.tcp.destination_port / d96c469e490a / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-ad924a32864e3b514fecb9685487317c5afb6e95af9c0dd4c53652d17d1839f1"></a>

<a id="canonical-b31bf39e57f44080c967594f8e9b25cf38ae8cfe0b7901bdfa10b80861999388"></a>

## port_ranges property — rules.criteria.tcp.destination_port / d96c469e490a / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-c20113547901f87882718adaac9e0415db6691467d9397c29113d1dba48bb0f0"></a>

## Next pages — rules.criteria.tcp.destination_port / d96c469e490a / 6

- [rules.criteria.tcp.destination_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-9b56fa2772143d67d704037b17dccfd04b4cc76359f102a270b37b740002548a)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-323262920a8f8a12b1ce22ae6f9acd70db7913fd17d861916089c468de577b27)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-9b56fa2772143d67d704037b17dccfd04b4cc76359f102a270b37b740002548a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea2f5c88d0f0150e13ad87371cb0d7ec0c522027d41595ecb23cba7fe78ab660"></a>

## rules.criteria.tcp.destination_port.no_port_match — rules.criteria.tcp.destination_port.no_port_match / 8f11c422b67e / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-323262920a8f8a12b1ce22ae6f9acd70db7913fd17d861916089c468de577b27)
- [rules.criteria.tcp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-453a8146986fa706616a25c3b67cc29e66c6b5348070240f2f411d9e479c92f6)
- rules.criteria.tcp.destination_port.no_port_match

<a id="canonical-0ae6b2edb9cdb4ed42adbdceb1f03d02bcbf5a022cb15a741afc69c8dab67ff2"></a>

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

<a id="canonical-b400a38fb9fb8a510776c312835d2597462bde43bf86b4126eeb60fa5f330b9a"></a>

## Direct properties — rules.criteria.tcp.destination_port.no_port_match / 8f11c422b67e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-33ec4ab10f951e030d3aa4311b4da6cd1e785620b55484f059e953c2bfd5ed81"></a>

## Next pages — rules.criteria.tcp.destination_port.no_port_match / 8f11c422b67e / 4

- [rules.criteria.tcp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-453a8146986fa706616a25c3b67cc29e66c6b5348070240f2f411d9e479c92f6)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-1829f31851bda2f6c53dd4f7f2ee64b42d010606823ecc244bfddd4c04598e51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa50b2feb1d9b2f0585e218e9ba7d28a65a2a521e82b4143faecbbc6414b15d1"></a>

## rules.criteria.tcp.source_port — rules.criteria.tcp.source_port / 898fccc6e907 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-323262920a8f8a12b1ce22ae6f9acd70db7913fd17d861916089c468de577b27)
- rules.criteria.tcp.source_port

<a id="canonical-42a584df71a5169a3c312f4dd3dd3193d88ae7e5ae12cf4e7f3a293c1bcc2142"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-34e6911f63f33d73d15620b80751a2d7ad4269fdcc77fa4650a31dc5a4a67fc5"></a>

## Direct properties — rules.criteria.tcp.source_port / 898fccc6e907 / 3

- [no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-7f14ff9f05b343d8dc634178c8e36b527a9edd80fe8e0808120cf7f4d4472047): complete subsection reference.

<a id="canonical-b0db40ad3baa9e073e6be0fdd593d4c0b23f15b528eb83d7ca460f6f99d9dd27"></a>

<a id="canonical-356dcec8bbeb3ed75ba4ce860bf114591c05406b156bc5fe8201574fb9b2fb36"></a>

## port property — rules.criteria.tcp.source_port / 898fccc6e907 / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-02de2f063134b9e090ff0449d8fc22d5bdbffb452b550eba808454aac1bb6d75"></a>

<a id="canonical-c8cb097d955c766c3ca4ff5e51232b3ab6cb28f177f8b361d4b0b9d1ffbc16e9"></a>

## port_ranges property — rules.criteria.tcp.source_port / 898fccc6e907 / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-b1db3b4693b83edda3b5624f3bcdacf64449cacaf6a9d681f23b2acd8d3b02d0"></a>

## Next pages — rules.criteria.tcp.source_port / 898fccc6e907 / 6

- [rules.criteria.tcp.source_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-7f14ff9f05b343d8dc634178c8e36b527a9edd80fe8e0808120cf7f4d4472047)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-323262920a8f8a12b1ce22ae6f9acd70db7913fd17d861916089c468de577b27)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-7f14ff9f05b343d8dc634178c8e36b527a9edd80fe8e0808120cf7f4d4472047"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c83d73a8ac259050994c17fb45aaa11047c31c77e8c1dafa2de3c30eefe2de29"></a>

## rules.criteria.tcp.source_port.no_port_match — rules.criteria.tcp.source_port.no_port_match / befdb6681f17 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-323262920a8f8a12b1ce22ae6f9acd70db7913fd17d861916089c468de577b27)
- [rules.criteria.tcp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-1829f31851bda2f6c53dd4f7f2ee64b42d010606823ecc244bfddd4c04598e51)
- rules.criteria.tcp.source_port.no_port_match

<a id="canonical-726ac133e91c11d0ad0400a15f0ca5698a3c8e434679445739968c229db62910"></a>

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

<a id="canonical-b39f6a52c2c8f1d9b1d7ae004a0d7ae58e9f5b67e3650bfc3d5402b2a985b77f"></a>

## Direct properties — rules.criteria.tcp.source_port.no_port_match / befdb6681f17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b3b34a7a063139a52f1cef584aa21261db0cd2732337d27ff2e2273bf288853"></a>

## Next pages — rules.criteria.tcp.source_port.no_port_match / befdb6681f17 / 4

- [rules.criteria.tcp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-1829f31851bda2f6c53dd4f7f2ee64b42d010606823ecc244bfddd4c04598e51)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-cc94026d041bb1a9a760ce463f9899d1055eaf3a9087642efaa0b3052ed199bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63d7825e105e9127c7d7380c04ae471ec9ad5d2cbd937dde7f6a05f9b2f6f7e6"></a>

## rules.criteria.udp — rules.criteria.udp / b16710295708 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- rules.criteria.udp

<a id="canonical-af356fd749638b9fc5edc0550eca4559fd3ff702c227d487594f476f4721432f"></a>

Type: `"single"`. Computed.

Action to apply on the packet if the NAT rule is applied.

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

<a id="canonical-dbe56c356432527f40f9c346364036150c08567f57e7a94fa7ded2ff8e3807f0"></a>

## Direct properties — rules.criteria.udp / b16710295708 / 3

- [destination_port](data-sources--nat_policy--reference--group-001.md#canonical-ba4f61ba72796b94546c83492839d93cd5a20a8f11d217850e2ef711a233e0c8): complete subsection reference.

- [source_port](data-sources--nat_policy--reference--group-001.md#canonical-daafb6da1fa3d87c06f13aec79418965232740b1ed5485aeafe82ae7545b989d): complete subsection reference.

<a id="canonical-9b74ae4e1aba691cc89d3f038afd0771e52b96064263bdcba42dd1f43d5aac9e"></a>

## Next pages — rules.criteria.udp / b16710295708 / 4

- [rules.criteria.udp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-ba4f61ba72796b94546c83492839d93cd5a20a8f11d217850e2ef711a233e0c8)
- [rules.criteria.udp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-daafb6da1fa3d87c06f13aec79418965232740b1ed5485aeafe82ae7545b989d)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-ba4f61ba72796b94546c83492839d93cd5a20a8f11d217850e2ef711a233e0c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-257fba07f46b0629bcd40ff253b15faa11c21092b0d725f3436a6369645d4f89"></a>

## rules.criteria.udp.destination_port — rules.criteria.udp.destination_port / a2500e0cabc8 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-cc94026d041bb1a9a760ce463f9899d1055eaf3a9087642efaa0b3052ed199bd)
- rules.criteria.udp.destination_port

<a id="canonical-bdba660aee91f1fd8ffc30c5a14ff3cd6211d5fa0b28cabb6ff429d699d12043"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-7d93b985b1061018effe67c8191adca45a9187535218d0cb26616507380dc2bc"></a>

## Direct properties — rules.criteria.udp.destination_port / a2500e0cabc8 / 3

- [no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-18c09d67326c597d0af204683477f5588d1696ebe4ea3ead5b7df5a1d4564c6c): complete subsection reference.

<a id="canonical-a56919d49549e7314ebf33c8027fe7300096bc851104f638eb28ea797e00baf0"></a>

<a id="canonical-7945fccef5edca41dcb9ec6b9232db7f67f65326128f4688d2208e79c8a1129a"></a>

## port property — rules.criteria.udp.destination_port / a2500e0cabc8 / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-c5904fc8bdf8f5af20bfaf78509726a215643a9a284f726e9d0f9befa61b9556"></a>

<a id="canonical-6d890c37c779e6245ec917cdf204870fc4c7e0d05797374ba4adb0533bd1ce06"></a>

## port_ranges property — rules.criteria.udp.destination_port / a2500e0cabc8 / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-c27ec8d9aaa26bd059fa7ffc02dd9f2d0978331d36fe87ac278bdf02c6189c09"></a>

## Next pages — rules.criteria.udp.destination_port / a2500e0cabc8 / 6

- [rules.criteria.udp.destination_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-18c09d67326c597d0af204683477f5588d1696ebe4ea3ead5b7df5a1d4564c6c)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-cc94026d041bb1a9a760ce463f9899d1055eaf3a9087642efaa0b3052ed199bd)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-18c09d67326c597d0af204683477f5588d1696ebe4ea3ead5b7df5a1d4564c6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c58bfc95417a8d0d629df1c743dc428251349b26229ee1757380bcb88b9ba3d5"></a>

## rules.criteria.udp.destination_port.no_port_match — rules.criteria.udp.destination_port.no_port_match / 06703d5d5783 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-cc94026d041bb1a9a760ce463f9899d1055eaf3a9087642efaa0b3052ed199bd)
- [rules.criteria.udp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-ba4f61ba72796b94546c83492839d93cd5a20a8f11d217850e2ef711a233e0c8)
- rules.criteria.udp.destination_port.no_port_match

<a id="canonical-e95b4398177c481e13ea35f23ecebc2f9d391e89cbb2e6f3361183144dc107e7"></a>

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

<a id="canonical-fac36d7fca4349dfa7883a876d4de34112fe55adf7054f36152b933d59d252b5"></a>

## Direct properties — rules.criteria.udp.destination_port.no_port_match / 06703d5d5783 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0a985c76793591b93acd2b54d9b23282ca762a8004196f25dc5c6800b34c3e33"></a>

## Next pages — rules.criteria.udp.destination_port.no_port_match / 06703d5d5783 / 4

- [rules.criteria.udp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-ba4f61ba72796b94546c83492839d93cd5a20a8f11d217850e2ef711a233e0c8)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-daafb6da1fa3d87c06f13aec79418965232740b1ed5485aeafe82ae7545b989d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25b58bce1115f2dc64cc35c82ead462a6cb1f32670b05a1a408ffb291f8d7c5e"></a>

## rules.criteria.udp.source_port — rules.criteria.udp.source_port / 970d88486d2d / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-cc94026d041bb1a9a760ce463f9899d1055eaf3a9087642efaa0b3052ed199bd)
- rules.criteria.udp.source_port

<a id="canonical-9b1ff8ddcac2b463c02da40d4789634c381b0e9ae88369a248c2ba9c8dac57b4"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-62a7e26b7fe10bda92962502857f9251f5f2ce0ab2fabf4d8f626bc0fc581b0d"></a>

## Direct properties — rules.criteria.udp.source_port / 970d88486d2d / 3

- [no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-49005c3546d06250dcfc5b5d4ad579e2ec5e8e23f53b852035c4a57325f7cfa9): complete subsection reference.

<a id="canonical-fabc27657bf7405e78daed60e351336b87f12be16b0ad3cfcf3df86630d5fd97"></a>

<a id="canonical-5a0a805ded6e990d1f38486ebaff7f883bd0692f65607cd9d71210d558c2690a"></a>

## port property — rules.criteria.udp.source_port / 970d88486d2d / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-c3eaf1a283859f06d560cd74599c08fe41824198e3d279959a001ee6e469df4f"></a>

<a id="canonical-548da09da6b01a811b82b521ba6a92d28aa14222bfc524da15114b8fd4714cf6"></a>

## port_ranges property — rules.criteria.udp.source_port / 970d88486d2d / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-b0b076ef50b184058aabb6c14d90ea18a6096d3913f0f0b7cdc6d53c76d7a352"></a>

## Next pages — rules.criteria.udp.source_port / 970d88486d2d / 6

- [rules.criteria.udp.source_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-49005c3546d06250dcfc5b5d4ad579e2ec5e8e23f53b852035c4a57325f7cfa9)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-cc94026d041bb1a9a760ce463f9899d1055eaf3a9087642efaa0b3052ed199bd)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-49005c3546d06250dcfc5b5d4ad579e2ec5e8e23f53b852035c4a57325f7cfa9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9471aca9220006069e9ce9551770577ea7ba3919761cc340c7f0b18ec662373"></a>

## rules.criteria.udp.source_port.no_port_match — rules.criteria.udp.source_port.no_port_match / 55acb965d314 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-ee379833005155e78b79cfc9b9c938cd59ca5be07fef9ce2fcbf2359473d79cb)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-cc94026d041bb1a9a760ce463f9899d1055eaf3a9087642efaa0b3052ed199bd)
- [rules.criteria.udp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-daafb6da1fa3d87c06f13aec79418965232740b1ed5485aeafe82ae7545b989d)
- rules.criteria.udp.source_port.no_port_match

<a id="canonical-33811b1c5a1084af454940e05d2f70bfc8b1fffc399ad4e447be3ae342b46c56"></a>

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

<a id="canonical-aed20a1a6ab2ef33ec1e5a0742af0292addfcd298ae7fc5542bed37acf7cba32"></a>

## Direct properties — rules.criteria.udp.source_port.no_port_match / 55acb965d314 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b5e30db833805e06a9f3e1cd5d24d629f8c4b4dbdcec1841c092cb87475f9d43"></a>

## Next pages — rules.criteria.udp.source_port.no_port_match / 55acb965d314 / 4

- [rules.criteria.udp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-daafb6da1fa3d87c06f13aec79418965232740b1ed5485aeafe82ae7545b989d)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-a4e3b0ae0e9db9059d9e05237babd022505e67d8a50020ea1ac6e1d061e60e76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ae91d26136d927fc6871b73b4b332acc7bef3dfd4a2cfaa3791772691468ad4"></a>

## rules.disable_spec — rules.disable_spec / e4c7b46b09a1 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- rules.disable_spec

<a id="canonical-88c9e9fbde7468b5fc33c3f9313d129c854e568548a3407bc0b8c8d66d986e9a"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-fbd635b9334897c66a3d7a4a97063651c7be25abed1c715bfde670e800e05356"></a>

## Direct properties — rules.disable_spec / e4c7b46b09a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5398aa873097ea514b8de07e4dfd850d621c383043494e42fbfac95dfd95049a"></a>

## Next pages — rules.disable_spec / e4c7b46b09a1 / 4

- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-7e494423c38e6c628469aebfa17b674ef405e452a32abbffbbaa5f0b494e5af3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1721dc9d915416b2ff7f5c425c031bb2e077e942f0f079aad96edde0a6ccf781"></a>

## rules.enable — rules.enable / bb5fb9fe48bc / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- rules.enable

<a id="canonical-1a3408b30acaf4e386b89b43606584666fa292cd42261ab4cddb2b5d5d3d6a5e"></a>

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

<a id="canonical-3adb42c7c95b1c09f6f0c7d2393fb1dc5718dc296fc2ac3adaea9a5ad44c7f97"></a>

## Direct properties — rules.enable / bb5fb9fe48bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cc082964568259c676e83c196ffd84400e3766a2eac6e8289fe9484914548c99"></a>

## Next pages — rules.enable / bb5fb9fe48bc / 4

- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-d27f3d9d29707b5ddf7466df4f1f8d4aa1f3a0f4ee55814bc229ea0b0ebde515"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24e5e6960e75969e866f279661300bbcf1d24ef1d5ea675b23b88062fe145f67"></a>

## rules.node_interface — rules.node_interface / 4b3303eeccdd / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- rules.node_interface

<a id="canonical-d5873746b2f5e4d0b1a3b689f123d0d82a5ac496279b298b2bd1310ad2b25b76"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-f8ec88ad0c8a2d24de9708db0506e7dae5dac8a8dde88acd5b13c6d654ef1b04"></a>

## Direct properties — rules.node_interface / 4b3303eeccdd / 3

- [list](data-sources--nat_policy--reference--group-001.md#canonical-012a383096fbc63cf768a14e0ae3e7e48c2873db85e173adac0978e816734542): complete subsection reference.

<a id="canonical-6c37c5cb69536fdb11ae0ad1af621cf103aa498f4d8aaa6d5d5d89e0dfdd5f9b"></a>

## Next pages — rules.node_interface / 4b3303eeccdd / 4

- [rules.node_interface.list](data-sources--nat_policy--reference--group-001.md#canonical-012a383096fbc63cf768a14e0ae3e7e48c2873db85e173adac0978e816734542)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-012a383096fbc63cf768a14e0ae3e7e48c2873db85e173adac0978e816734542"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15475af43648cc366d06c7c59d93e5cc45a44eadfa34edf1b819a5b750a63eaa"></a>

## rules.node_interface.list — rules.node_interface.list / 185ea8e673be / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.node_interface](data-sources--nat_policy--reference--group-001.md#canonical-d27f3d9d29707b5ddf7466df4f1f8d4aa1f3a0f4ee55814bc229ea0b0ebde515)
- rules.node_interface.list

<a id="canonical-cdedbf90e8f4582f89c2fbde6b458947a1d239972e7e31777d74de3db05dd557"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-3686cd7b436c247536b6b1824e911656830a004c5740c89d2024b7bd498da7c2"></a>

## Direct properties — rules.node_interface.list / 185ea8e673be / 3

- [interface](data-sources--nat_policy--reference--group-001.md#canonical-00206525ab03d4a4fe437e22789f17c6e2f7319eeb09681795e506863afbc341): complete subsection reference.

<a id="canonical-d30febadbd5dfdff686a294fec43c10245ded2efebac297faab894b89084f0ca"></a>

<a id="canonical-e55d1528582f82a4d640fc71b84ba574b0b85cb61ef42e15e3495aaee9c3149c"></a>

## node property — rules.node_interface.list / 185ea8e673be / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-363d007b98fde6f15e5a53106e5b2ca0a424cf6278471d495ae0f20032c88e02"></a>

## Next pages — rules.node_interface.list / 185ea8e673be / 5

- [rules.node_interface.list.interface](data-sources--nat_policy--reference--group-001.md#canonical-00206525ab03d4a4fe437e22789f17c6e2f7319eeb09681795e506863afbc341)
- [rules.node_interface](data-sources--nat_policy--reference--group-001.md#canonical-d27f3d9d29707b5ddf7466df4f1f8d4aa1f3a0f4ee55814bc229ea0b0ebde515)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-00206525ab03d4a4fe437e22789f17c6e2f7319eeb09681795e506863afbc341"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dde98366ada5c2aa05101297668af004a71831c665add0c0d2ea1fa7cd54f8f"></a>

## rules.node_interface.list.interface — rules.node_interface.list.interface / 6a7f2bd63fbb / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.node_interface](data-sources--nat_policy--reference--group-001.md#canonical-d27f3d9d29707b5ddf7466df4f1f8d4aa1f3a0f4ee55814bc229ea0b0ebde515)
- [rules.node_interface.list](data-sources--nat_policy--reference--group-001.md#canonical-012a383096fbc63cf768a14e0ae3e7e48c2873db85e173adac0978e816734542)
- rules.node_interface.list.interface

<a id="canonical-1bf0558418267940feb9960756fcdaf1775347de6f208b507c3b80be22de23aa"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-f9a99b9f44f3ca1fb9d9835e3f2df2408cd21ef8b767141e74ce478f259450bf"></a>

## Direct properties — rules.node_interface.list.interface / 6a7f2bd63fbb / 3

<a id="canonical-c036f81a9463902e126525d539e2937d9cf185acde07c3525e00c18d83b0f9b5"></a>

<a id="canonical-bf8d872a3f4274180a8aa3d59d7bf572545043b0b532cd6b8d128617ee43176d"></a>

## kind property — rules.node_interface.list.interface / 6a7f2bd63fbb / 4

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

<a id="canonical-7d5bf562ed1263bd1d9f2380eb8647bf7fb5d5b78807831d150964ed92cb4cf3"></a>

<a id="canonical-f52a9988592493e4eac719715f453368f07004732d4e1cce88d0c0772fa76b05"></a>

## name property — rules.node_interface.list.interface / 6a7f2bd63fbb / 5

Type: `"string"`. Computed.

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

<a id="canonical-dc6aad8f081d6cbf6c4dd5f631b93cd9c491747ffcb1b3918f5e1ac3b772bab1"></a>

<a id="canonical-89cedbb3f23cac540e2c420478a206f7a7f225cefee4b4b70f5ece004270303c"></a>

## namespace property — rules.node_interface.list.interface / 6a7f2bd63fbb / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-686f589c846e31e0170be955f508d3c4ed7be265ace5e11362e8441a2ab336b3"></a>

<a id="canonical-8787fc94154c28776e89d0c7c32141c6e1a9e491fbe961421c4efbbdc4dfe849"></a>

## tenant property — rules.node_interface.list.interface / 6a7f2bd63fbb / 7

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

<a id="canonical-9d0609ea0f1d60f919317ed20250667da31abce1364b3e00d4a9b6fcc603a4f2"></a>

<a id="canonical-269f7d8242ca990f26a19d2132ae5f4f592134c743dc58c3b49de84d69f26db2"></a>

## uid property — rules.node_interface.list.interface / 6a7f2bd63fbb / 8

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

<a id="canonical-bf9a6053b3858c95971903d177d98a92ecc23f6d058c12e23ae78e422e755844"></a>

## Next pages — rules.node_interface.list.interface / 6a7f2bd63fbb / 9

- [rules.node_interface.list](data-sources--nat_policy--reference--group-001.md#canonical-012a383096fbc63cf768a14e0ae3e7e48c2873db85e173adac0978e816734542)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-dabe55b282aee5012d817c9b22a7d59b8bde82652e4797c4c03a6c31eadd2f75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5855c1dc967246e3ff470e1ca46af9b85092215a7fc8ad198b75d06b971c3c71"></a>

## rules.segment — rules.segment / 00b8936a2767 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- rules.segment

<a id="canonical-391d1b8612c574c88f42cc55473b80a985f67cbd4c9194cab7b55d0655131507"></a>

Type: `"single"`. Computed.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

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

<a id="canonical-7fef935ff4bcb32cc910b70dff6d145788f3812505a2fd673df71617298908f6"></a>

## Direct properties — rules.segment / 00b8936a2767 / 3

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-d7121c454ccdfa7d80a715b0f0c34ea549663c8ea253a50d506c8a42e5469667): complete subsection reference.

<a id="canonical-39b433dcacae40b34dc21b996d89a52dfc776134f0c7633aee769324a59674c8"></a>

## Next pages — rules.segment / 00b8936a2767 / 4

- [rules.segment.refs](data-sources--nat_policy--reference--group-001.md#canonical-d7121c454ccdfa7d80a715b0f0c34ea549663c8ea253a50d506c8a42e5469667)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-d7121c454ccdfa7d80a715b0f0c34ea549663c8ea253a50d506c8a42e5469667"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87f5857161254049d51481b7e9100d363e9a24b7901331a9a4f006113e96213f"></a>

## rules.segment.refs — rules.segment.refs / 05f4a9275b92 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.segment](data-sources--nat_policy--reference--group-001.md#canonical-dabe55b282aee5012d817c9b22a7d59b8bde82652e4797c4c03a6c31eadd2f75)
- rules.segment.refs

<a id="canonical-b36d5d6d5572b52a5470cc80fd83a0aa4e6f6b4a0e48c5f66a94d7bc00fb7d9b"></a>

Type: `"list"`. Computed.

Segment. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

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

<a id="canonical-fd276acced88d4da53e98c548b042272045092bbf7658875a0831adf0f93c241"></a>

## Direct properties — rules.segment.refs / 05f4a9275b92 / 3

<a id="canonical-e92cf2cd7da2ca7f36cd1d640769187dec22392e6bab8cc1a6444162ff2f8e4c"></a>

<a id="canonical-2cf1a71d0ff409c331cc1a07338208b8831cf775ef361fed062b62b9ce205e45"></a>

## kind property — rules.segment.refs / 05f4a9275b92 / 4

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

<a id="canonical-812dd0ac8e84b14a922b376f80ed54fa0e7dc90fa976af60eccb561e755cc645"></a>

<a id="canonical-b90f51b3f0f1f2942c5a27fead79a4eaecc980825810ccd28b45163c0f0bed9f"></a>

## name property — rules.segment.refs / 05f4a9275b92 / 5

Type: `"string"`. Computed.

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

<a id="canonical-47caa2540e372de16221730512f7c80ca90b62f411659aba5165f872e49eb92e"></a>

<a id="canonical-c8af2951226029f8cb9427c8b14e32a89d354e4beeeaee4e6230167a9b872678"></a>

## namespace property — rules.segment.refs / 05f4a9275b92 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-dba3dd7a2ca885d6f0e9cc985f9c9bc91f0a530b6b4da0d7faf354bf216667c6"></a>

<a id="canonical-a7658652317a252095c7a125a935b1dce16923cac5812ce2a50a2ec803556785"></a>

## tenant property — rules.segment.refs / 05f4a9275b92 / 7

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

<a id="canonical-942da4a7e7fe27a087187994da85ad9a61405f580e8928ceeb4d667e9314059b"></a>

<a id="canonical-23ff691e922be041af367764159ef3cfb022452fcfac7c0a43d5dfa028f47d32"></a>

## uid property — rules.segment.refs / 05f4a9275b92 / 8

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

<a id="canonical-5dedb0c97316e08ebcddeffc4c44ff94feaa8cf4e15018dbfcc4968bdbe8866b"></a>

## Next pages — rules.segment.refs / 05f4a9275b92 / 9

- [rules.segment](data-sources--nat_policy--reference--group-001.md#canonical-dabe55b282aee5012d817c9b22a7d59b8bde82652e4797c4c03a6c31eadd2f75)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-6a7c35e7a50dfb77cc131a8a4eba991f9f38f74e9f6981b6d6faf384ce6e15e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59a7d144031c8c1e3006d5c3c45dedf9ac2da349c0b5304b3ee8c9a3ce7e543b"></a>

## rules.virtual_network — rules.virtual_network / 3e072e185b43 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- rules.virtual_network

<a id="canonical-cbddabff3be16256f397b8ce641364f4e761b08a6569483538c831104e69f58a"></a>

Type: `"single"`. Computed.

Carries the reference to virtual network.

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

<a id="canonical-e9413a68949a5a760e892d7e52ba68fc534616bf265c8473991522764bc43935"></a>

## Direct properties — rules.virtual_network / 3e072e185b43 / 3

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-e5ca6a6e6f97bf64abb4814426942888e1ec233b78664ba186921d96dcb26c1e): complete subsection reference.

<a id="canonical-b31cff9cdac47037799948a79a09ce9287183ead9d8bd8b3c7041a4b0a33af8e"></a>

## Next pages — rules.virtual_network / 3e072e185b43 / 4

- [rules.virtual_network.refs](data-sources--nat_policy--reference--group-001.md#canonical-e5ca6a6e6f97bf64abb4814426942888e1ec233b78664ba186921d96dcb26c1e)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-e5ca6a6e6f97bf64abb4814426942888e1ec233b78664ba186921d96dcb26c1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb38f1c192918cc99c1f79f19d48135cc39e59bba4806239f39c5b0e2e840a04"></a>

## rules.virtual_network.refs — rules.virtual_network.refs / 22c092f2672c / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-25524751b3c336960de7b8cb83dfff6dceadaa4f59e848a6afbe182680d568ec)
- [rules.virtual_network](data-sources--nat_policy--reference--group-001.md#canonical-6a7c35e7a50dfb77cc131a8a4eba991f9f38f74e9f6981b6d6faf384ce6e15e9)
- rules.virtual_network.refs

<a id="canonical-dc139ddfa35fac76edfabe18064b992a352e417347492380ac96457f79f7c814"></a>

Type: `"list"`. Computed.

Virtual Network Reference. Reference to virtual network.

Upstream description:

Reference to virtual network.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-cc4f0cd26d26a3267a33e96f4c47455a42030d69832600383e231709a320e0e0"></a>

## Direct properties — rules.virtual_network.refs / 22c092f2672c / 3

<a id="canonical-dfa84393b6aba3d05dda73bd5ada6de1a61340ff957211478c5f5c5c2c22b629"></a>

<a id="canonical-e4d105bdf098316711ce082453c3ebe1f4366a98d521f1a7eaf0c0311aa19381"></a>

## kind property — rules.virtual_network.refs / 22c092f2672c / 4

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

<a id="canonical-a6b57d8ed5e53112bee299a8a222bde4091cef6ff6088e39a9cec404c4c8a408"></a>

<a id="canonical-1992722b530a49174362d9359ff09b6b83c4e77104cdee451e8fb99f3405c9f2"></a>

## name property — rules.virtual_network.refs / 22c092f2672c / 5

Type: `"string"`. Computed.

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

<a id="canonical-d9e9e186c070c5c324255920345899595d10661c86c59a0e4433a40f18669029"></a>

<a id="canonical-30a0930728456dbf717bf8371d74af41def772b6572fad6acaefedb477966034"></a>

## namespace property — rules.virtual_network.refs / 22c092f2672c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-07a380f925223dff4a10bd416f0979c8dd627b68dcad6c0b4f30c6167a057021"></a>

<a id="canonical-09342297c29224cb610f76aa13375463bfd63fb0844423b30b4e8833a25676d1"></a>

## tenant property — rules.virtual_network.refs / 22c092f2672c / 7

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

<a id="canonical-75829fa69cdad677a11fdbc95c692574c97ddd2313f69b0f56687e9ee1509381"></a>

<a id="canonical-2443b896b54b98d1b754bf7c8140c137dfb87302febc21816fd8b760de85f23b"></a>

## uid property — rules.virtual_network.refs / 22c092f2672c / 8

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

<a id="canonical-c0a6a9c4eba55a2a004e4d5d8ea2623c1477577f337a401ffa20d5c60fa3342f"></a>

## Next pages — rules.virtual_network.refs / 22c092f2672c / 9

- [rules.virtual_network](data-sources--nat_policy--reference--group-001.md#canonical-6a7c35e7a50dfb77cc131a8a4eba991f9f38f74e9f6981b6d6faf384ce6e15e9)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-45c6954ad1fe44c2f95f33065ed54ee51b8ee9ad280c257c2c757d100b9b89d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e83eb4df5cbe7b4a0e1330d60b3ecbfebdf4a4004745e48252d229c24495b8bc"></a>

## site — site / b0d23b48edd9 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- site

<a id="canonical-77ffd0e49759954f42b9c58321477d4bf44abda43c11fd8f72a18b9eb782d9d1"></a>

Type: `"single"`. Computed.

Site Reference Type. Reference to Site Object.

Upstream description:

Reference to Site Object.

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

<a id="canonical-c50b9011ddc0294c465a71df53a2fd99d57c8898b225d07925eae145a81e0f11"></a>

## Direct properties — site / b0d23b48edd9 / 3

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-8e9bec9d0b9f0c65aece518036ba7f23ac94aae0a93c177fc926b326f6ce1207): complete subsection reference.

<a id="canonical-be5ee4c57d27658236a4389b625b400d7e0f1cc00fb5971e33f4397e0e41ec47"></a>

## Next pages — site / b0d23b48edd9 / 4

- [site.refs](data-sources--nat_policy--reference--group-001.md#canonical-8e9bec9d0b9f0c65aece518036ba7f23ac94aae0a93c177fc926b326f6ce1207)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-8e9bec9d0b9f0c65aece518036ba7f23ac94aae0a93c177fc926b326f6ce1207"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da3ed72657391477b852e2e7069f02be02e7e9f906560fd7d05ffdad26081335"></a>

## site.refs — site.refs / 1920c30f1cc2 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-3700bcc776d70423b66533c4058868d9b850878907e066987ec2e221b01c7899)
- [site](data-sources--nat_policy--reference--group-001.md#canonical-45c6954ad1fe44c2f95f33065ed54ee51b8ee9ad280c257c2c757d100b9b89d4)
- site.refs

<a id="canonical-d7c609af0cc1b8007802c33c5f50ff03ec9a44338726c55c9c41fdf878264bf3"></a>

Type: `"list"`. Computed.

Site. Reference to Site Object.

Upstream description:

Reference to Site Object.

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

<a id="canonical-2fc18c9466c8f1d5d52047896e0635f69de15332cd23e27ce5c3ac73cf87979e"></a>

## Direct properties — site.refs / 1920c30f1cc2 / 3

<a id="canonical-b1ba199a1a0964c210a184f37c95b55a5bf3e156a097e142ea59745bde3b3080"></a>

<a id="canonical-a81fa8b8e0daed121297e9191147b3c68c638f82cbab56c84a0d32b3010d02dd"></a>

## kind property — site.refs / 1920c30f1cc2 / 4

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

<a id="canonical-eff62a7af5df1091c52d18c1f3f85d698de144c691508e3e5bc5c8a246142de4"></a>

<a id="canonical-77cdafd653116a518394c3d17c7740aefc562be0410993ce70aa029fa185e537"></a>

## name property — site.refs / 1920c30f1cc2 / 5

Type: `"string"`. Computed.

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

<a id="canonical-5161eea35ad523c5f813bed60b5893588aa2728229583e6cd83234c7640b44dd"></a>

<a id="canonical-bae9e93f9d6f0c868127e7d086a54b9af94326ec9b5b95ca2e37c331ea19ef42"></a>

## namespace property — site.refs / 1920c30f1cc2 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-4d2418860e418bd1010d0af20af40eb383e69bf12822a08f9cd8b43669d4295a"></a>

<a id="canonical-6b0159e6e5345f80244f49d6ff0df44f0d204087f5de0b00489a1dd1976d3604"></a>

## tenant property — site.refs / 1920c30f1cc2 / 7

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

<a id="canonical-114d73cd64913e78479af865b2faeac3bec2b8eb155f9c42538abed10db15287"></a>

<a id="canonical-aa9495fc50b1f04a46df3de25a069c95d13c6377895c39d95743976d137cfcba"></a>

## uid property — site.refs / 1920c30f1cc2 / 8

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

<a id="canonical-3ff4976abbd6937502773e0194ceeccde13c78e03da77a915b86f2c056715faa"></a>

## Next pages — site.refs / 1920c30f1cc2 / 9

- [site](data-sources--nat_policy--reference--group-001.md#canonical-45c6954ad1fe44c2f95f33065ed54ee51b8ee9ad280c257c2c757d100b9b89d4)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
