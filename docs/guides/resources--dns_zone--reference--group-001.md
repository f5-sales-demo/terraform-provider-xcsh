---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d37706e9fbc0191b4436f7a90b7506b053082f5840867c601272d70d124eed9"></a>

## Property reference — Property reference / 6f4349d8c84a / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- Property reference

<a id="canonical-3cf8b17db909cedd5a1ce52a08026a3aa2228b4b3e3a018a18f1c4a079ecdeb2"></a>

## Direct properties — Property reference / 6f4349d8c84a / 3

<a id="canonical-caf722e3d735c8c76ff229cd97f5abda98877641a4cb15fcd240c88a9105b081"></a>

<a id="canonical-b97515310961b887ae815d1e9e8ba869228612a30b48bd441d1b3b3ea611fe2d"></a>

## annotations property — Property reference / 6f4349d8c84a / 4

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

<a id="canonical-5c790971a2427e4acd9e51b47e0aec80e0205b7454e3695ec39476f18a54a722"></a>

<a id="canonical-1923382d7b2d27e0605d6b7c6732786027f54626b55b69cdf400ad64db0e785c"></a>

## description property — Property reference / 6f4349d8c84a / 5

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

<a id="canonical-226aab1cd2d8fe6842b04a8cee399c8751ce037af3f4401cd0c6062efd796087"></a>

<a id="canonical-5133786eb7bd69b6f8922283afafc3a1e72729e65145f5f34651139b0a7f93e6"></a>

## disable property — Property reference / 6f4349d8c84a / 6

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

<a id="canonical-09dae83edf4c8feaff00bbd6cc6cfa26f57523d2ed962e56383c0fa1da8e3d52"></a>

<a id="canonical-c92b918d859df8c53d2e044c622809f6f1a730a2dacd6dfbb7290287d9d360ad"></a>

## id property — Property reference / 6f4349d8c84a / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-342735d119486d6927df32db3bfba41d2cc2c312f02046ebc75ed08cd395b1e4"></a>

<a id="canonical-21292dd50b04b56e3ac8501434284a991c2080ac68285a8dcc0d47595d880187"></a>

## labels property — Property reference / 6f4349d8c84a / 8

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

<a id="canonical-1bc0f5b0bf2cb599b20f3e7f05ec226301a9b9475a15ff7de49ce6191e5cd158"></a>

<a id="canonical-514df2f6672be633db0a2c6fcd2a24a42283e91b0057bc202118a418a9a09bf5"></a>

## name property — Property reference / 6f4349d8c84a / 9

Type: `"string"`. Required.

Domain name for the DNS Zone (e.g., example.com). Must be a valid DNS domain name.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.DomainValidator(),
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

<a id="canonical-4e68b09f514f857c663996861b49bc075212b3f03b998b1e338ab4bba1ed247f"></a>

<a id="canonical-41905b3b2ec9422e3d7bcffc91494fb2d9fd403a8b73e15229f904b813d43630"></a>

## namespace property — Property reference / 6f4349d8c84a / 10

Type: `"string"`. Optional, Computed.

Namespace for the DNS Zone. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6): complete subsection reference.

- [secondary](resources--dns_zone--reference--group-003.md#canonical-915ef6a26cadd3cdb1d03e8f1f87b29e6e5292057e8da86defb920ce53ea35c9): complete subsection reference.

- [timeouts](resources--dns_zone--reference--group-003.md#canonical-592d8ee330c4fcf5e49e421377380abc8cd9a954a4313bed2da72b7dc44025bf): complete subsection reference.

<a id="canonical-816be94d918c7fc6386ff8c49be33f9b02ee736a4fa524a8935261d755692b95"></a>

## All schema paths — Property reference / 6f4349d8c84a / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_zone--reference--group-001.md#canonical-caf722e3d735c8c76ff229cd97f5abda98877641a4cb15fcd240c88a9105b081) |
| `description` | [description](resources--dns_zone--reference--group-001.md#canonical-5c790971a2427e4acd9e51b47e0aec80e0205b7454e3695ec39476f18a54a722) |
| `disable` | [disable](resources--dns_zone--reference--group-001.md#canonical-226aab1cd2d8fe6842b04a8cee399c8751ce037af3f4401cd0c6062efd796087) |
| `id` | [id](resources--dns_zone--reference--group-001.md#canonical-09dae83edf4c8feaff00bbd6cc6cfa26f57523d2ed962e56383c0fa1da8e3d52) |
| `labels` | [labels](resources--dns_zone--reference--group-001.md#canonical-342735d119486d6927df32db3bfba41d2cc2c312f02046ebc75ed08cd395b1e4) |
| `name` | [name](resources--dns_zone--reference--group-001.md#canonical-1bc0f5b0bf2cb599b20f3e7f05ec226301a9b9475a15ff7de49ce6191e5cd158) |
| `namespace` | [namespace](resources--dns_zone--reference--group-001.md#canonical-4e68b09f514f857c663996861b49bc075212b3f03b998b1e338ab4bba1ed247f) |
| `primary` | [primary](resources--dns_zone--reference--group-001.md#canonical-f0df71152037b454f5fd2a1339729cbe2aa0d1f15cbeb8ada01eab617ed14435) |
| `primary.allow_http_lb_managed_records` | [primary.allow_http_lb_managed_records](resources--dns_zone--reference--group-001.md#canonical-58855fea13abbe96c61a63d99aaea640a2d4225ad44a8572c16f61d357895921) |
| `primary.default_rr_set_group` | [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-264380bf2df9264ada7036448f23d9e49047139c6c0a0c209a7e48a3b073375e) |
| `primary.default_rr_set_group.a_record` | [primary.default_rr_set_group.a_record](resources--dns_zone--reference--group-001.md#canonical-bc515a86c7bd5786c45d78ea53a7da119d458d08ced64c98b27442cbdc59cf3a) |
| `primary.default_rr_set_group.a_record.name` | [primary.default_rr_set_group.a_record.name](resources--dns_zone--reference--group-001.md#canonical-9815ed0fac4144419a01fa239953df034ca7ad531b505bf565adfd53da56dca9) |
| `primary.default_rr_set_group.a_record.values` | [primary.default_rr_set_group.a_record.values](resources--dns_zone--reference--group-001.md#canonical-fc68bcb484297dea3f074b7d477a43b3c89cd74184b566d1ca83d6a892575b9f) |
| `primary.default_rr_set_group.aaaa_record` | [primary.default_rr_set_group.aaaa_record](resources--dns_zone--reference--group-001.md#canonical-2e858ac2aebe340e0d8720d6171b4361ee24d6d7753ee1e1ae5c6571903a7d7f) |
| `primary.default_rr_set_group.aaaa_record.name` | [primary.default_rr_set_group.aaaa_record.name](resources--dns_zone--reference--group-001.md#canonical-68c8ff285be349e835e304b76da836ec6453f38fd346fcb2dcba6103b309bb94) |
| `primary.default_rr_set_group.aaaa_record.values` | [primary.default_rr_set_group.aaaa_record.values](resources--dns_zone--reference--group-001.md#canonical-e140258a76b9b94181d55bce9c51d1d9a2cff7c2d06bafa130938c1ab0879666) |
| `primary.default_rr_set_group.afsdb_record` | [primary.default_rr_set_group.afsdb_record](resources--dns_zone--reference--group-001.md#canonical-fe5af3bb5c6d98d77036196a82f14ecefdf16fe4139b7af831f555889fde3974) |
| `primary.default_rr_set_group.afsdb_record.name` | [primary.default_rr_set_group.afsdb_record.name](resources--dns_zone--reference--group-001.md#canonical-6467bf4bbb9917a5148418f7c9b74e30c9dddff165f088310862574b20d97bc2) |
| `primary.default_rr_set_group.afsdb_record.values` | [primary.default_rr_set_group.afsdb_record.values](resources--dns_zone--reference--group-001.md#canonical-54de531e70e2280e4555e80b4221ca8003b2374b5db05984b98d08c83bce34be) |
| `primary.default_rr_set_group.afsdb_record.values.hostname` | [primary.default_rr_set_group.afsdb_record.values.hostname](resources--dns_zone--reference--group-001.md#canonical-658aa1e4f89509ff7d11b4ea3102457c6a3ff1d0b05270e7f04add6fb9b46d89) |
| `primary.default_rr_set_group.afsdb_record.values.subtype` | [primary.default_rr_set_group.afsdb_record.values.subtype](resources--dns_zone--reference--group-001.md#canonical-de8cffd9c56157fba0bd11d23964094d2ea820e8631912560201ebbfbe7f1f5c) |
| `primary.default_rr_set_group.alias_record` | [primary.default_rr_set_group.alias_record](resources--dns_zone--reference--group-001.md#canonical-5a3c83fdec994eadf604c445e075882be2a69cc8f77a070e537294662b2386e8) |
| `primary.default_rr_set_group.alias_record.value` | [primary.default_rr_set_group.alias_record.value](resources--dns_zone--reference--group-001.md#canonical-9637fdf46e48d862d25345b956b40615802b7e0a6f153ac8e5d4b20de8503d8f) |
| `primary.default_rr_set_group.caa_record` | [primary.default_rr_set_group.caa_record](resources--dns_zone--reference--group-001.md#canonical-b8d44d2e87764a9b2c90680730ec575e0407d7ee62253655216c862d1b057a5a) |
| `primary.default_rr_set_group.caa_record.name` | [primary.default_rr_set_group.caa_record.name](resources--dns_zone--reference--group-001.md#canonical-b77b32cbfa0dfedcbdb719240443c6a62a726d0dd4fbf178f9a8376bd5596df8) |
| `primary.default_rr_set_group.caa_record.values` | [primary.default_rr_set_group.caa_record.values](resources--dns_zone--reference--group-001.md#canonical-64fb576a89a17ed8e2428602af5391e8cff4cc935fc264dfba9538c4d416092f) |
| `primary.default_rr_set_group.caa_record.values.flags` | [primary.default_rr_set_group.caa_record.values.flags](resources--dns_zone--reference--group-001.md#canonical-5fad0d88b8648d1237ada6db08df2da9ea5aaf3d02c3dc79fe0a06f458a42d67) |
| `primary.default_rr_set_group.caa_record.values.tag` | [primary.default_rr_set_group.caa_record.values.tag](resources--dns_zone--reference--group-001.md#canonical-9eb39ef685866bf0f1f7000b856b5c1a98fdb496fd17a305659298c5521d670f) |
| `primary.default_rr_set_group.caa_record.values.value` | [primary.default_rr_set_group.caa_record.values.value](resources--dns_zone--reference--group-001.md#canonical-d5c1c9e6498512ee919e1691ac59f08a888121cd6d3bc792ef77ed09a539b9b1) |
| `primary.default_rr_set_group.cds_record` | [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-0796127575044d96a7fcf8811ee2b4c1f8bcd15271301e18df7974373863b0b7) |
| `primary.default_rr_set_group.cds_record.name` | [primary.default_rr_set_group.cds_record.name](resources--dns_zone--reference--group-001.md#canonical-92a08edfc0e5c442570b332bab4da398c888db6d05f7713aac3899f7db907e21) |
| `primary.default_rr_set_group.cds_record.values` | [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-c83ad09a767e5950e115f6867f002b6416bb279e3ef56c5b081f7f39deabae59) |
| `primary.default_rr_set_group.cds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.cds_record.values.ds_key_algorithm](resources--dns_zone--reference--group-001.md#canonical-f3fd150f9f78da28539e8404efe49e16d46d5de179f65862cc6f933cd8734d5a) |
| `primary.default_rr_set_group.cds_record.values.key_tag` | [primary.default_rr_set_group.cds_record.values.key_tag](resources--dns_zone--reference--group-001.md#canonical-51c0e71b3fc9f06f65b1306d4f5300b139da322f75cafa849dd08720698e9f38) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest](resources--dns_zone--reference--group-001.md#canonical-39c73459a239f389b4d6a9d7e7c8fc0736a471202bdba84b36d82a9cd91cd676) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest.digest](resources--dns_zone--reference--group-001.md#canonical-b934967095930068e105f4ab0cb3b0bf07d92b8465100dfad35956ca1e44c83e) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest](resources--dns_zone--reference--group-001.md#canonical-ffd71854cdd8cdc872c50924f84a8f61a6ebb8584d792ac16052fa7e8bba8845) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest.digest](resources--dns_zone--reference--group-001.md#canonical-39146718da7d7d11d747999f92a09bc4e8b5c7f1582cc2db2851238e6647f8ea) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest](resources--dns_zone--reference--group-001.md#canonical-c78be4f701a4bf07e83705a01b3d467dc5fd8d6c3e4b7dfc7d31943ebcd26ca4) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest.digest](resources--dns_zone--reference--group-001.md#canonical-34e231455619605086e5093411b76c2b08d43da3a12c2504635807a4acef7a76) |
| `primary.default_rr_set_group.cert_record` | [primary.default_rr_set_group.cert_record](resources--dns_zone--reference--group-001.md#canonical-35b54fc48af4a141a054910a2e7fbadf2d60aa44dc7f9d501f0107f41557e0e3) |
| `primary.default_rr_set_group.cert_record.name` | [primary.default_rr_set_group.cert_record.name](resources--dns_zone--reference--group-001.md#canonical-ac59a80ab2be105145a9fed01f1451008dcbb5c7b5487df5e76efbf08363e288) |
| `primary.default_rr_set_group.cert_record.values` | [primary.default_rr_set_group.cert_record.values](resources--dns_zone--reference--group-001.md#canonical-60d828a22f1921316eeab51f0d1fa3c99e5ca29144ef0557aa559b8a910c93cd) |
| `primary.default_rr_set_group.cert_record.values.algorithm` | [primary.default_rr_set_group.cert_record.values.algorithm](resources--dns_zone--reference--group-001.md#canonical-9acbe02274496d3cac6ed9d945a8fe145e17b094cacd66d875328dfa2c5a82f1) |
| `primary.default_rr_set_group.cert_record.values.cert_key_tag` | [primary.default_rr_set_group.cert_record.values.cert_key_tag](resources--dns_zone--reference--group-001.md#canonical-047427b75fe967c06976aea655ecb001c6e05ae7b60ed06921f4c3ad4023eaab) |
| `primary.default_rr_set_group.cert_record.values.cert_type` | [primary.default_rr_set_group.cert_record.values.cert_type](resources--dns_zone--reference--group-001.md#canonical-ca91483b5c9aa839b4ff31fb39951fd87223c33f9318f956cc87218ad72a247e) |
| `primary.default_rr_set_group.cert_record.values.certificate` | [primary.default_rr_set_group.cert_record.values.certificate](resources--dns_zone--reference--group-001.md#canonical-5c937c0b8d067565d6fb3e46e29a20d7995ada62bcba17b1e096f20e26e35c31) |
| `primary.default_rr_set_group.cname_record` | [primary.default_rr_set_group.cname_record](resources--dns_zone--reference--group-001.md#canonical-e71d65a7afabb6f07076babaedb18c3f43be759348f4857518a48e5be3df5b60) |
| `primary.default_rr_set_group.cname_record.name` | [primary.default_rr_set_group.cname_record.name](resources--dns_zone--reference--group-001.md#canonical-1f2b78cfa39de47786e27406dea1bf2a80d9aa441867106e165a6c9ba29f4888) |
| `primary.default_rr_set_group.cname_record.value` | [primary.default_rr_set_group.cname_record.value](resources--dns_zone--reference--group-001.md#canonical-0d53f753398168cc5dff897ed0cd633037f9ab57235a35922b46f7c4c964349d) |
| `primary.default_rr_set_group.description_spec` | [primary.default_rr_set_group.description_spec](resources--dns_zone--reference--group-001.md#canonical-b3bf216484665be82c1bbeeb64e18359bdb240a7d228c49a4bc20b6989abf80d) |
| `primary.default_rr_set_group.ds_record` | [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-002.md#canonical-9a8239d3f6bd83c459c22ef486f4c378221782004dfb8e3a5d39360c90d26599) |
| `primary.default_rr_set_group.ds_record.name` | [primary.default_rr_set_group.ds_record.name](resources--dns_zone--reference--group-002.md#canonical-030fe7896e1262c76289a9a3f9091d258e31213fbe2bc6c41082f253915b7de1) |
| `primary.default_rr_set_group.ds_record.values` | [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-e3c78ae8cda23c10de36087f9932beab06ecfb1e8428a8efc086365130c3d376) |
| `primary.default_rr_set_group.ds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.ds_record.values.ds_key_algorithm](resources--dns_zone--reference--group-002.md#canonical-dc814eb94f7e4540d97aa0f674efe083d85be8ab38b631f425c83ff8769aa0b1) |
| `primary.default_rr_set_group.ds_record.values.key_tag` | [primary.default_rr_set_group.ds_record.values.key_tag](resources--dns_zone--reference--group-002.md#canonical-b0d259435bbce603cb39c15c71202071116b260097462ea7460239baa77947f9) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest](resources--dns_zone--reference--group-002.md#canonical-660c74bcb25854ab8d14ffb718bbddd445fa7ab9a7df83405671a32301c46bc6) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest.digest](resources--dns_zone--reference--group-002.md#canonical-5fd0911a73ebed263dbf35bf53e7ee69434fe7f6f94ae2500451bd8f41e8c79e) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest](resources--dns_zone--reference--group-002.md#canonical-b97fdf25b052ddf367d6b53680ee1f754afc5353965f8c0153ec098fbda20e18) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest.digest](resources--dns_zone--reference--group-002.md#canonical-cd0e5a6bf2aad3fdd226e1283345c5b36db455557fffbbbde1aa6aa85d85a065) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest](resources--dns_zone--reference--group-002.md#canonical-8c6af6f4266a18cc3cc3cbd7461359a94e19fc65ba1f80337cb31d8e9a4a193f) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest.digest](resources--dns_zone--reference--group-002.md#canonical-51cbd828a95301b45ae40bdff967a9b714ee07b6b28d1a09bb390bef2b57ef2c) |
| `primary.default_rr_set_group.eui48_record` | [primary.default_rr_set_group.eui48_record](resources--dns_zone--reference--group-002.md#canonical-6d9f2388659ebdaf2fcfe6b0f2f271cbc4d5c63056a473a10138589f2d8f28e2) |
| `primary.default_rr_set_group.eui48_record.name` | [primary.default_rr_set_group.eui48_record.name](resources--dns_zone--reference--group-002.md#canonical-4643cc862286990aa94e26b0dc51730fc91a3e1f2fba83d2f444c5c112766b56) |
| `primary.default_rr_set_group.eui48_record.value` | [primary.default_rr_set_group.eui48_record.value](resources--dns_zone--reference--group-002.md#canonical-e15a020edb3937514f16667c7688377b386cdf56083660d8f0f3316bd52ba8b7) |
| `primary.default_rr_set_group.eui64_record` | [primary.default_rr_set_group.eui64_record](resources--dns_zone--reference--group-002.md#canonical-ebe180c5d3e1c34ef5f2f1e977ad310c579a32590a1ac31657d9e3ab7aad149e) |
| `primary.default_rr_set_group.eui64_record.name` | [primary.default_rr_set_group.eui64_record.name](resources--dns_zone--reference--group-002.md#canonical-546dbc00927c0775ab0758ce4678f60b57eacd7dfe3be49a724ffa24f519b500) |
| `primary.default_rr_set_group.eui64_record.value` | [primary.default_rr_set_group.eui64_record.value](resources--dns_zone--reference--group-002.md#canonical-cc2ae5c35a1070f40a050b53c9a858221c3fa14ebdaf5580ed47cfc867f5bde7) |
| `primary.default_rr_set_group.lb_record` | [primary.default_rr_set_group.lb_record](resources--dns_zone--reference--group-002.md#canonical-1b095e419ce2862479d03cc3908cd350afa5d9c3ce508a054705bae9513efbb4) |
| `primary.default_rr_set_group.lb_record.name` | [primary.default_rr_set_group.lb_record.name](resources--dns_zone--reference--group-002.md#canonical-6218b9dc934bebc42c63c59e66ac55428a771e4443c2adbc1312743a610593c1) |
| `primary.default_rr_set_group.lb_record.value` | [primary.default_rr_set_group.lb_record.value](resources--dns_zone--reference--group-002.md#canonical-fb4077bba3708dedad1cfaaa331b2bbe4fc84f66f8665441370ae38095453c8e) |
| `primary.default_rr_set_group.lb_record.value.name` | [primary.default_rr_set_group.lb_record.value.name](resources--dns_zone--reference--group-002.md#canonical-24aab09cac04bc1c11f181fd830609a1bcc4788e48f7b677bc5448f4bc934372) |
| `primary.default_rr_set_group.lb_record.value.namespace` | [primary.default_rr_set_group.lb_record.value.namespace](resources--dns_zone--reference--group-002.md#canonical-9d8bafa6c2451cf317d8ae5b0aed0555ea76ed23f60f75c5a55f3c53742499c7) |
| `primary.default_rr_set_group.lb_record.value.tenant` | [primary.default_rr_set_group.lb_record.value.tenant](resources--dns_zone--reference--group-002.md#canonical-b5ba9b0e6a3fe065e76a365155db93486c227ea8239cc7ae88a35059ddea1adb) |
| `primary.default_rr_set_group.loc_record` | [primary.default_rr_set_group.loc_record](resources--dns_zone--reference--group-002.md#canonical-9e557309f9eac6a4bc2b884e0a577755d66b9ec3f95e1616bda7f60f2e20c5c4) |
| `primary.default_rr_set_group.loc_record.name` | [primary.default_rr_set_group.loc_record.name](resources--dns_zone--reference--group-002.md#canonical-066fefe0825e0c60ff626d11568add5af771d513cef10beccdfd45d09af062e2) |
| `primary.default_rr_set_group.loc_record.values` | [primary.default_rr_set_group.loc_record.values](resources--dns_zone--reference--group-002.md#canonical-00cf7713e51c975ba65921feb73b1efd972ffcc99a03db844fdfb3302d0c4472) |
| `primary.default_rr_set_group.loc_record.values.altitude` | [primary.default_rr_set_group.loc_record.values.altitude](resources--dns_zone--reference--group-002.md#canonical-9d8964e9a5d11bf5c75f9739eb9db4816b06e1939542ccd485978cbe0c03e86d) |
| `primary.default_rr_set_group.loc_record.values.horizontal_precision` | [primary.default_rr_set_group.loc_record.values.horizontal_precision](resources--dns_zone--reference--group-002.md#canonical-2bd350ed9c5acc6a60094649cb92facf875aa91ed44165fe0a0e358c077a8579) |
| `primary.default_rr_set_group.loc_record.values.latitude_degree` | [primary.default_rr_set_group.loc_record.values.latitude_degree](resources--dns_zone--reference--group-002.md#canonical-d1fec37de3bb14c81be96593c7de296d82476ee78f059744f77607267ce6abc5) |
| `primary.default_rr_set_group.loc_record.values.latitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.latitude_hemisphere](resources--dns_zone--reference--group-002.md#canonical-c4eb1990e977902e83d715742d4680814c0d2948318d2b38e92e7bc14292cb47) |
| `primary.default_rr_set_group.loc_record.values.latitude_minute` | [primary.default_rr_set_group.loc_record.values.latitude_minute](resources--dns_zone--reference--group-002.md#canonical-c3804645da21022ed8c33cd070f2434b516b93c5f874f17457c2fa055c1c3d5c) |
| `primary.default_rr_set_group.loc_record.values.latitude_second` | [primary.default_rr_set_group.loc_record.values.latitude_second](resources--dns_zone--reference--group-002.md#canonical-5ed597640e5f0e50ac085c907670d70a0b7462c1c5b5c57174915c8456410adb) |
| `primary.default_rr_set_group.loc_record.values.location_diameter` | [primary.default_rr_set_group.loc_record.values.location_diameter](resources--dns_zone--reference--group-002.md#canonical-81c800cb145f1528cad20e877ad40147a3efd7ea7e68438f4aca96eab4a6f6ed) |
| `primary.default_rr_set_group.loc_record.values.longitude_degree` | [primary.default_rr_set_group.loc_record.values.longitude_degree](resources--dns_zone--reference--group-002.md#canonical-1b382697513f05b9d2fd51a7d563df7c255f443961c550a9978656ceadd365e3) |
| `primary.default_rr_set_group.loc_record.values.longitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.longitude_hemisphere](resources--dns_zone--reference--group-002.md#canonical-4f8378def2c49a8086e385a4e1a2d13107ed84e36e2fd23db11657738700a90e) |
| `primary.default_rr_set_group.loc_record.values.longitude_minute` | [primary.default_rr_set_group.loc_record.values.longitude_minute](resources--dns_zone--reference--group-002.md#canonical-d1d1e8f16eef15134ba677473cc59b624faab18b29c13977e09dce3c21ecbead) |
| `primary.default_rr_set_group.loc_record.values.longitude_second` | [primary.default_rr_set_group.loc_record.values.longitude_second](resources--dns_zone--reference--group-002.md#canonical-8bb3542dfa5035da4ff7fb753e2ddb24af7c34758f46aad7d41e30a23a6216dc) |
| `primary.default_rr_set_group.loc_record.values.vertical_precision` | [primary.default_rr_set_group.loc_record.values.vertical_precision](resources--dns_zone--reference--group-002.md#canonical-b3925685de3c303112d42201206a99c349a5ea127ac676b7a8a73434681d5e70) |
| `primary.default_rr_set_group.mx_record` | [primary.default_rr_set_group.mx_record](resources--dns_zone--reference--group-002.md#canonical-cd2866444f0039a87f9dfb81b18fcac562f91cd82c82cc457509c35d17996c09) |
| `primary.default_rr_set_group.mx_record.name` | [primary.default_rr_set_group.mx_record.name](resources--dns_zone--reference--group-002.md#canonical-790a25b4b60b9c945cd20ffbe5f39f564dc45cd3537cba2e3289e0e22cf06107) |
| `primary.default_rr_set_group.mx_record.values` | [primary.default_rr_set_group.mx_record.values](resources--dns_zone--reference--group-002.md#canonical-033347cf71da161052c41945518d91e0ee663af7e0873affe6c518b5944bef2c) |
| `primary.default_rr_set_group.mx_record.values.domain` | [primary.default_rr_set_group.mx_record.values.domain](resources--dns_zone--reference--group-002.md#canonical-e6afeab19592737972142b63a3a2bd58af8c0921a902a1d10761271288007cb1) |
| `primary.default_rr_set_group.mx_record.values.priority` | [primary.default_rr_set_group.mx_record.values.priority](resources--dns_zone--reference--group-002.md#canonical-31254d157ac754a16d019b2e7f1120b1a262fdc2f39a2da6ccec5a6e8ec51870) |
| `primary.default_rr_set_group.naptr_record` | [primary.default_rr_set_group.naptr_record](resources--dns_zone--reference--group-002.md#canonical-f284891cca16b8890fcd1a2102116fda0d3c6ba9f73a04eed2e4986551723574) |
| `primary.default_rr_set_group.naptr_record.name` | [primary.default_rr_set_group.naptr_record.name](resources--dns_zone--reference--group-002.md#canonical-f382a8f5facda2b934541d6e75bd0a8280922745fadc550bdf19ab589833efaa) |
| `primary.default_rr_set_group.naptr_record.values` | [primary.default_rr_set_group.naptr_record.values](resources--dns_zone--reference--group-002.md#canonical-067e57912a5a5bc322daaf5f5520e66975cff88dd64671618c8399f6c84007b5) |
| `primary.default_rr_set_group.naptr_record.values.flags` | [primary.default_rr_set_group.naptr_record.values.flags](resources--dns_zone--reference--group-002.md#canonical-4f26f91ae0b817661cf20551820ff6c0f434050dece4fbac1c94629bf30b64db) |
| `primary.default_rr_set_group.naptr_record.values.order` | [primary.default_rr_set_group.naptr_record.values.order](resources--dns_zone--reference--group-002.md#canonical-03eb2b6a420a1dc4e1b25b286b85ae38f5f208230066ff0635b8e47ddbdc7fae) |
| `primary.default_rr_set_group.naptr_record.values.preference` | [primary.default_rr_set_group.naptr_record.values.preference](resources--dns_zone--reference--group-002.md#canonical-23bd86196bcddf126bbb4782f1a574d0f6bcfcb50523ef9105eefc5ee9a50db3) |
| `primary.default_rr_set_group.naptr_record.values.regexp` | [primary.default_rr_set_group.naptr_record.values.regexp](resources--dns_zone--reference--group-002.md#canonical-5f44a64cea55cbbbc3dab71cc05755781225a286db4b4f49bdd2490a507961d8) |
| `primary.default_rr_set_group.naptr_record.values.replacement` | [primary.default_rr_set_group.naptr_record.values.replacement](resources--dns_zone--reference--group-002.md#canonical-730a8606cfd6b899a434e53828353c715462c3753499595e112bddf269cd429d) |
| `primary.default_rr_set_group.naptr_record.values.service` | [primary.default_rr_set_group.naptr_record.values.service](resources--dns_zone--reference--group-002.md#canonical-10e80ffc47cbf4dea28ab4b6515b6367f426a526683ea820de32b46018716c69) |
| `primary.default_rr_set_group.ns_record` | [primary.default_rr_set_group.ns_record](resources--dns_zone--reference--group-002.md#canonical-dfed0301fbe283864bf9bd2b9fbeb55a94f74a45eb7554e38a91103e063c7d4d) |
| `primary.default_rr_set_group.ns_record.name` | [primary.default_rr_set_group.ns_record.name](resources--dns_zone--reference--group-002.md#canonical-8f0d5f6e90e707e21635ecce39b204e2ff11ad6bf8b4cdf8b0249d15959962d4) |
| `primary.default_rr_set_group.ns_record.values` | [primary.default_rr_set_group.ns_record.values](resources--dns_zone--reference--group-002.md#canonical-d7ff7ae2fa88e9c462bc0fa0ea4fa5d2aa7f3c6f388e4c9ded55cc42b66ca912) |
| `primary.default_rr_set_group.ptr_record` | [primary.default_rr_set_group.ptr_record](resources--dns_zone--reference--group-002.md#canonical-807291065fb9cfa38b6c64db67fee14fc9efde9e28e255ffb56390a12222486d) |
| `primary.default_rr_set_group.ptr_record.name` | [primary.default_rr_set_group.ptr_record.name](resources--dns_zone--reference--group-002.md#canonical-52403f8cc3e81e7dae112077eb51259578ee73b95e9acd42dcb700941e3f3cc0) |
| `primary.default_rr_set_group.ptr_record.values` | [primary.default_rr_set_group.ptr_record.values](resources--dns_zone--reference--group-002.md#canonical-26ecdee7f42f5bdcc7493a1c8ca348e4e9343f20ec41278d1460d3daca6ca1ec) |
| `primary.default_rr_set_group.srv_record` | [primary.default_rr_set_group.srv_record](resources--dns_zone--reference--group-002.md#canonical-ee78b04d1479134cb6e3b100325c9637a0cee5180efb3f38a0b43a06c255cf04) |
| `primary.default_rr_set_group.srv_record.name` | [primary.default_rr_set_group.srv_record.name](resources--dns_zone--reference--group-002.md#canonical-892b2e142ed598d71de4cabef9c1bc4aa2987574c21aca20c4652d0265254250) |
| `primary.default_rr_set_group.srv_record.values` | [primary.default_rr_set_group.srv_record.values](resources--dns_zone--reference--group-002.md#canonical-6f51465f25fa658a28feb5082a46701b6be7e3cc66862ffd2a90e0f31fa35128) |
| `primary.default_rr_set_group.srv_record.values.port` | [primary.default_rr_set_group.srv_record.values.port](resources--dns_zone--reference--group-002.md#canonical-c8926f76973e7c43fe3242c26af502f5e73f43e5c84663c8c2c4066f0bc99937) |
| `primary.default_rr_set_group.srv_record.values.priority` | [primary.default_rr_set_group.srv_record.values.priority](resources--dns_zone--reference--group-002.md#canonical-266a72825273cb0a974f4d62ca9c3e55cd68dae4b23bd68e0faf8312c93fa042) |
| `primary.default_rr_set_group.srv_record.values.target` | [primary.default_rr_set_group.srv_record.values.target](resources--dns_zone--reference--group-002.md#canonical-ebe90cf90d41a1802d631a3bd1ec8968f92cdf8c80bc3c4a46c6e9b92b7624ac) |
| `primary.default_rr_set_group.srv_record.values.weight` | [primary.default_rr_set_group.srv_record.values.weight](resources--dns_zone--reference--group-002.md#canonical-2be083b3b7e09102505865e3aa594af7285fdadd7c581c9752ae30dfb3791c52) |
| `primary.default_rr_set_group.sshfp_record` | [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0fd30bf7c1875249db91f9c537b8225215894cf72c7cd23b94b26cf0f8dc7271) |
| `primary.default_rr_set_group.sshfp_record.name` | [primary.default_rr_set_group.sshfp_record.name](resources--dns_zone--reference--group-002.md#canonical-76ec0625dadda3c1b142ed0604077f5953ee57cb55917294178508097b910665) |
| `primary.default_rr_set_group.sshfp_record.values` | [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-3b834ce5df8abc2cffb5fb163a98eb0558f04369d3132dd2f4ac9f222ae7b0ff) |
| `primary.default_rr_set_group.sshfp_record.values.algorithm` | [primary.default_rr_set_group.sshfp_record.values.algorithm](resources--dns_zone--reference--group-002.md#canonical-10db005a50e8c44fc7290f9d2f9684a9de7bb779c702ab31490bb808784570e3) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](resources--dns_zone--reference--group-002.md#canonical-f0b12d0e07d64327d5d8522a9a6c2396a9c87f31e247be32f21a726ab895dfc4) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint](resources--dns_zone--reference--group-002.md#canonical-ecf32208677aa77b02118ecdef17b087a18313a019ea116a77a5d6e4819a5e5a) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](resources--dns_zone--reference--group-002.md#canonical-1bcaf9f2f5ba78cd87b099f5c0680865084e889559b268b6a13a071bf3e13dd2) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint](resources--dns_zone--reference--group-002.md#canonical-7286c2536ebe17bb4c1bf9e05fe034acfb305b536c2b548e8001a8abf25b40a0) |
| `primary.default_rr_set_group.tlsa_record` | [primary.default_rr_set_group.tlsa_record](resources--dns_zone--reference--group-002.md#canonical-7852f0cbf82c87e7752c2701c7bafc40b7acdf7562f056eac5af057e4a7ac413) |
| `primary.default_rr_set_group.tlsa_record.name` | [primary.default_rr_set_group.tlsa_record.name](resources--dns_zone--reference--group-002.md#canonical-6bde7fa9b9b7f10792fe47688d3e14ad174747f466c91137bf3f334d72f57b27) |
| `primary.default_rr_set_group.tlsa_record.values` | [primary.default_rr_set_group.tlsa_record.values](resources--dns_zone--reference--group-002.md#canonical-8faf6eb248629989492dcb2ed2ea43e55f91dcc0a2b420a190096a3797d74fef) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_association_data` | [primary.default_rr_set_group.tlsa_record.values.certificate_association_data](resources--dns_zone--reference--group-002.md#canonical-a2d596fe8bb4dbf1f31b02b1083e7d66a4cbe490b59971c8ebec0989774398af) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_usage` | [primary.default_rr_set_group.tlsa_record.values.certificate_usage](resources--dns_zone--reference--group-002.md#canonical-1728f6bff7d5d58ecc2b8670bfcb9b0a56ed09dd6eab5ab4aa947db80e606d94) |
| `primary.default_rr_set_group.tlsa_record.values.matching_type` | [primary.default_rr_set_group.tlsa_record.values.matching_type](resources--dns_zone--reference--group-002.md#canonical-4b892046d6b20111b985e40f055edf027b29245f6ed5eea383caa2f0d638accb) |
| `primary.default_rr_set_group.tlsa_record.values.selector` | [primary.default_rr_set_group.tlsa_record.values.selector](resources--dns_zone--reference--group-002.md#canonical-63aa0f3a70d812b3106bb95d9a8fcfd330d198468302dd8fba83f07f70217d75) |
| `primary.default_rr_set_group.ttl` | [primary.default_rr_set_group.ttl](resources--dns_zone--reference--group-001.md#canonical-0df548e4b0e455d90e09f0e2801d21ecb478de9d25b5ecd9f14817491cd681db) |
| `primary.default_rr_set_group.txt_record` | [primary.default_rr_set_group.txt_record](resources--dns_zone--reference--group-002.md#canonical-51eb4b36eacea5e3ba339251f95e878b86ed4f201376c314f06494824a96b4cd) |
| `primary.default_rr_set_group.txt_record.name` | [primary.default_rr_set_group.txt_record.name](resources--dns_zone--reference--group-002.md#canonical-4968725855434b4643f7b2034a2ccdc7e31b857736974d8e805a1665f0b99801) |
| `primary.default_rr_set_group.txt_record.values` | [primary.default_rr_set_group.txt_record.values](resources--dns_zone--reference--group-002.md#canonical-e3e41e3d2e77427ff8f6a06044316afd824e365605175f9c4406490c46d5b60b) |
| `primary.default_soa_parameters` | [primary.default_soa_parameters](resources--dns_zone--reference--group-002.md#canonical-e095b6512a3a9b738e0020e504d7db4fcef2c0cd5f503558aa285f3fd21dd86a) |
| `primary.dnssec_mode` | [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-8420098636cb935a0db948198999de0e4ce6f8e65a9b01c432e93c44b17b0f10) |
| `primary.dnssec_mode.disable_spec` | [primary.dnssec_mode.disable_spec](resources--dns_zone--reference--group-002.md#canonical-12161a39eb4a6edf5d9c31b9b0a896fb95df47794f43279921d14fc27cc5a329) |
| `primary.dnssec_mode.enable` | [primary.dnssec_mode.enable](resources--dns_zone--reference--group-002.md#canonical-a315e64ebb2dddba4926df681d252f8fe0f797565c3281d1b52ca5c320a8d4ab) |
| `primary.rr_set_group` | [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-ba051d9ab361e297286390078ab285b50cfc51e3aed2e56572cfc4dfa9833b5b) |
| `primary.rr_set_group.metadata` | [primary.rr_set_group.metadata](resources--dns_zone--reference--group-002.md#canonical-4213c2db1f14c049218d6f76e32da858c4416482a5f6ccf3d0849845318a1570) |
| `primary.rr_set_group.metadata.description_spec` | [primary.rr_set_group.metadata.description_spec](resources--dns_zone--reference--group-002.md#canonical-7a10cd34072cadd9797218e748cc495e7fb91ec9c1b8f0a0c9f76ce220a317e3) |
| `primary.rr_set_group.metadata.name` | [primary.rr_set_group.metadata.name](resources--dns_zone--reference--group-002.md#canonical-8790e588869b7922f402c751d2bd03243fd6e79ec9dd7bc4142bfbc2b09c8a06) |
| `primary.rr_set_group.rr_set` | [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-79799cf7b30e119a5eac90eac36a5a208ee438e2a64272523bfd74e1854c9651) |
| `primary.rr_set_group.rr_set.a_record` | [primary.rr_set_group.rr_set.a_record](resources--dns_zone--reference--group-002.md#canonical-8dd140157ae9c87cfbc41526dfd3c7c7241cf49da3b2080ca71e87a77149390d) |
| `primary.rr_set_group.rr_set.a_record.name` | [primary.rr_set_group.rr_set.a_record.name](resources--dns_zone--reference--group-002.md#canonical-282ce1c269b09e98675174a6a6eb7532098f46073faadfe7f1a355d5c813c44d) |
| `primary.rr_set_group.rr_set.a_record.values` | [primary.rr_set_group.rr_set.a_record.values](resources--dns_zone--reference--group-002.md#canonical-30282b52b230c3a3ccbf3bf92b3bdd8fbb302049f4828470ca53b1fb3fade7da) |
| `primary.rr_set_group.rr_set.aaaa_record` | [primary.rr_set_group.rr_set.aaaa_record](resources--dns_zone--reference--group-002.md#canonical-7b1128bd0d6eeaa2259c980bc6ac16205e543406b20055f7986986483f340f03) |
| `primary.rr_set_group.rr_set.aaaa_record.name` | [primary.rr_set_group.rr_set.aaaa_record.name](resources--dns_zone--reference--group-002.md#canonical-ebd80a66c368b51a8308bc3f2517f58006c3f136aa521b7d2382975c1e66e36b) |
| `primary.rr_set_group.rr_set.aaaa_record.values` | [primary.rr_set_group.rr_set.aaaa_record.values](resources--dns_zone--reference--group-002.md#canonical-c8a71a044e9497b75ba1411c1a5dd599ad5ae07cd3dce415ae54e1714abde1e0) |
| `primary.rr_set_group.rr_set.afsdb_record` | [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--reference--group-002.md#canonical-66d3aaebb6048a48acab610fb8028d67037db69cb4cf1443d5c192aba7cd57ca) |
| `primary.rr_set_group.rr_set.afsdb_record.name` | [primary.rr_set_group.rr_set.afsdb_record.name](resources--dns_zone--reference--group-002.md#canonical-67ca590d9fe208c0154ace0469fb06192b0cde6a337baf030ca0c8d356acba0d) |
| `primary.rr_set_group.rr_set.afsdb_record.values` | [primary.rr_set_group.rr_set.afsdb_record.values](resources--dns_zone--reference--group-002.md#canonical-b282af3970514edb73bc073ae1a9c82d03fbdfd9e1aad7092e5c089c2e7b6363) |
| `primary.rr_set_group.rr_set.afsdb_record.values.hostname` | [primary.rr_set_group.rr_set.afsdb_record.values.hostname](resources--dns_zone--reference--group-002.md#canonical-920a10fc824691abc18d9d82e59591fcd5d9ab0a64cdb272d52ff1deb7a723fd) |
| `primary.rr_set_group.rr_set.afsdb_record.values.subtype` | [primary.rr_set_group.rr_set.afsdb_record.values.subtype](resources--dns_zone--reference--group-002.md#canonical-8055fd07af621914a80b330494014cebc42d79cbdb404d0861b0843bcec5a839) |
| `primary.rr_set_group.rr_set.alias_record` | [primary.rr_set_group.rr_set.alias_record](resources--dns_zone--reference--group-002.md#canonical-cb37358245a37f0df01bf2235362769b3e5106bb39fa7bc443c7e279fab5f890) |
| `primary.rr_set_group.rr_set.alias_record.value` | [primary.rr_set_group.rr_set.alias_record.value](resources--dns_zone--reference--group-002.md#canonical-e34f3ac0a97f1936774fc5c34034c8691a66add26920cd07c9145b8fed955666) |
| `primary.rr_set_group.rr_set.caa_record` | [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--reference--group-002.md#canonical-e32abe2d4e40424bd27ae885e935e982be381024e47243cd2cb22f4c1eba16d3) |
| `primary.rr_set_group.rr_set.caa_record.name` | [primary.rr_set_group.rr_set.caa_record.name](resources--dns_zone--reference--group-002.md#canonical-22bb4b2d10fc4a7bb51e28a46ca20aee1cce67765dab55d84f6d80eec48f6d63) |
| `primary.rr_set_group.rr_set.caa_record.values` | [primary.rr_set_group.rr_set.caa_record.values](resources--dns_zone--reference--group-002.md#canonical-97234d1a0474aa1584b1f4ca274d45d94201e2ee95e3a77416092e3752cae84c) |
| `primary.rr_set_group.rr_set.caa_record.values.flags` | [primary.rr_set_group.rr_set.caa_record.values.flags](resources--dns_zone--reference--group-002.md#canonical-311e1baeed18139c9290b06bd70eef3f6c34409adb7e2fd72628a58a93420435) |
| `primary.rr_set_group.rr_set.caa_record.values.tag` | [primary.rr_set_group.rr_set.caa_record.values.tag](resources--dns_zone--reference--group-002.md#canonical-3cb28684816b038f222aaee02ff5c4fa5de30d157ee44cbd234f11f8ed558b05) |
| `primary.rr_set_group.rr_set.caa_record.values.value` | [primary.rr_set_group.rr_set.caa_record.values.value](resources--dns_zone--reference--group-002.md#canonical-00748efaf6475feaf4091b6ea7bcbd67c29b5af03464b53b71b2619c25cdc4de) |
| `primary.rr_set_group.rr_set.cds_record` | [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-c6a39ab792433adef8073b7657a226afef3e28b5f3c9ce3adc071918ec516c1b) |
| `primary.rr_set_group.rr_set.cds_record.name` | [primary.rr_set_group.rr_set.cds_record.name](resources--dns_zone--reference--group-002.md#canonical-622273816904063720be4a445807bd6e2d9bb772682524be554e16290b4ce85e) |
| `primary.rr_set_group.rr_set.cds_record.values` | [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-b9d03af879fbf303ba3c9f6b7ce82300e99118f460c9c5d724debecf8b2e6b22) |
| `primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm](resources--dns_zone--reference--group-002.md#canonical-e9dcbc8572f5e29292dee0e84557230c9eb84220e6769e0511d8907785d1d17d) |
| `primary.rr_set_group.rr_set.cds_record.values.key_tag` | [primary.rr_set_group.rr_set.cds_record.values.key_tag](resources--dns_zone--reference--group-002.md#canonical-ed19a1f02148abe7ec3420024abfb01e5eacf1fb644b4fc3215d8afc52573950) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](resources--dns_zone--reference--group-002.md#canonical-cf7936a2b8e89093a68cc243f70840ff38113d84fe24a5601ca4edb96c5f210d) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest](resources--dns_zone--reference--group-002.md#canonical-9e2065437348b11c448b7d54908708e3a167c5f528123c8aab79cde5f6d4cea7) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](resources--dns_zone--reference--group-002.md#canonical-04bdbaef8648dc1ef6eca5d4cdd5249991d7b8e29d3a7bcf5dd68126a9a82042) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest](resources--dns_zone--reference--group-002.md#canonical-29b2299625740f56b9b27980dad2bdb4248b88f2adb4846eec8375dec02618a7) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](resources--dns_zone--reference--group-003.md#canonical-d8f281ec335da889a1ba25fb8a5c29073c8afeda2f6b315664b68e8e11d185a1) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest](resources--dns_zone--reference--group-003.md#canonical-5c5eb52018291efb8da653ca186de1798913273a3d50e12af0a2a9b536d8c5b5) |
| `primary.rr_set_group.rr_set.cert_record` | [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--reference--group-003.md#canonical-c6c789327ff609e0c4d0f7246b08495c0a44715d2cdfe6693e61181f6d9fabbc) |
| `primary.rr_set_group.rr_set.cert_record.name` | [primary.rr_set_group.rr_set.cert_record.name](resources--dns_zone--reference--group-003.md#canonical-883b8f4d056f3ba5391dc53a07a971eb151a6af5553b9fbe2af8dcbd2cfc7c48) |
| `primary.rr_set_group.rr_set.cert_record.values` | [primary.rr_set_group.rr_set.cert_record.values](resources--dns_zone--reference--group-003.md#canonical-b90e051c6f5245cd6b931c4799adac5d7b6ccdc0cbba85220f4e66e643c9a028) |
| `primary.rr_set_group.rr_set.cert_record.values.algorithm` | [primary.rr_set_group.rr_set.cert_record.values.algorithm](resources--dns_zone--reference--group-003.md#canonical-6519a134216376074eea5e54df670d57e4aed2ac75e75ba6b1994e4656a07d0b) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_key_tag` | [primary.rr_set_group.rr_set.cert_record.values.cert_key_tag](resources--dns_zone--reference--group-003.md#canonical-7313060d7ba1accc9e97a6e90ade63803a87d21d58d889389ad56d9966ecacba) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_type` | [primary.rr_set_group.rr_set.cert_record.values.cert_type](resources--dns_zone--reference--group-003.md#canonical-9dcf7e2cd9163ae1d6ba7e638993b0aadf203c81cf43f7feecc02080a49563cc) |
| `primary.rr_set_group.rr_set.cert_record.values.certificate` | [primary.rr_set_group.rr_set.cert_record.values.certificate](resources--dns_zone--reference--group-003.md#canonical-4c30b9997468b3c96afca5388c5bc71e89e301ebed5db1c7297874023c83f3b9) |
| `primary.rr_set_group.rr_set.cname_record` | [primary.rr_set_group.rr_set.cname_record](resources--dns_zone--reference--group-003.md#canonical-34d5905aa5802da9f46e479b550c7491281f0405fb25fae463fa6f458d9884db) |
| `primary.rr_set_group.rr_set.cname_record.name` | [primary.rr_set_group.rr_set.cname_record.name](resources--dns_zone--reference--group-003.md#canonical-685796b1c132cc30372659ebf71a3aa8ef0af128f0675536969172b1ca5f9f5f) |
| `primary.rr_set_group.rr_set.cname_record.value` | [primary.rr_set_group.rr_set.cname_record.value](resources--dns_zone--reference--group-003.md#canonical-984dd91ec5eba33dc59729ba2e7cf1b6084260aab91f13f3eddaefd6c2598991) |
| `primary.rr_set_group.rr_set.description_spec` | [primary.rr_set_group.rr_set.description_spec](resources--dns_zone--reference--group-002.md#canonical-9ff8e9a41b9eca99db5ee61a44fa5bd8d0e3d7183c700a0e974fdd8c7ab4dba5) |
| `primary.rr_set_group.rr_set.ds_record` | [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-76e22f22caa83f6f4831e7a0452d5f93fc822ec65b28dd68a95f137eb2650252) |
| `primary.rr_set_group.rr_set.ds_record.name` | [primary.rr_set_group.rr_set.ds_record.name](resources--dns_zone--reference--group-003.md#canonical-0ee745cb5fe1cf9e99c4061c4bdd7ef53c8436108b0edcf5d103ab3f68f9424d) |
| `primary.rr_set_group.rr_set.ds_record.values` | [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-908e4fcb9d02f4efd483ec03224e499c35f27be5c3c5a8022415fec8defa57a1) |
| `primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm](resources--dns_zone--reference--group-003.md#canonical-bba743dd2af769047a7ab5600837bd376326100d105f97c633cd015d7cb21d30) |
| `primary.rr_set_group.rr_set.ds_record.values.key_tag` | [primary.rr_set_group.rr_set.ds_record.values.key_tag](resources--dns_zone--reference--group-003.md#canonical-4a1f693cd57aab471d3cae5363ebcd752fe3982993bbcbe45df6e30d73b8f52d) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](resources--dns_zone--reference--group-003.md#canonical-057dde4bdc8f160e145a875295645849a5f7a61c8236827c96b69e882be9bee0) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest](resources--dns_zone--reference--group-003.md#canonical-ae773b8f3fd8a2b65550424391a9838cfc6d0b1a9caa4bf60a4c754e22b4794f) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](resources--dns_zone--reference--group-003.md#canonical-ebc154f09d5c3662308baae866ca53275f66009db0414bbd28684b2ea7a262d8) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest](resources--dns_zone--reference--group-003.md#canonical-52ba0723e72fde65372e463a99933bb3c85201def6d246fded085fec6731a6c0) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](resources--dns_zone--reference--group-003.md#canonical-88e4f14212003d0ca2e0023f047c43081cdc4f24f50ed0d126428644a87988b0) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest](resources--dns_zone--reference--group-003.md#canonical-78de4e9dad0a8cf23ead45a1ae6c6e466937a222244fec1babfa6c315d09a80c) |
| `primary.rr_set_group.rr_set.eui48_record` | [primary.rr_set_group.rr_set.eui48_record](resources--dns_zone--reference--group-003.md#canonical-216f2a93b61d77861a018d38382993d0c72bba0cfdae7a7fb91d0bf9f2f1466e) |
| `primary.rr_set_group.rr_set.eui48_record.name` | [primary.rr_set_group.rr_set.eui48_record.name](resources--dns_zone--reference--group-003.md#canonical-826be9317cc42134d49761bcbb9da7be7f3687e66838719d1f4134827d25f2c8) |
| `primary.rr_set_group.rr_set.eui48_record.value` | [primary.rr_set_group.rr_set.eui48_record.value](resources--dns_zone--reference--group-003.md#canonical-022deb453359246e49bf1093f8fe79c766548ddd753e11d794810a5f999e8273) |
| `primary.rr_set_group.rr_set.eui64_record` | [primary.rr_set_group.rr_set.eui64_record](resources--dns_zone--reference--group-003.md#canonical-8795d399362d50041a23c675da96a20f0c13bab3158eda96a55b1d98648cf36f) |
| `primary.rr_set_group.rr_set.eui64_record.name` | [primary.rr_set_group.rr_set.eui64_record.name](resources--dns_zone--reference--group-003.md#canonical-c31ee61a959964b682f4300d2b6ec8294a8a6bc2d164db34ec2b6b62a9193bce) |
| `primary.rr_set_group.rr_set.eui64_record.value` | [primary.rr_set_group.rr_set.eui64_record.value](resources--dns_zone--reference--group-003.md#canonical-526fb9d2d6a2f79daeeeb0f21113b86ddc1b5595d0ab579ad7ec4f8169e67a26) |
| `primary.rr_set_group.rr_set.lb_record` | [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--reference--group-003.md#canonical-0963412d2570059ec0fd1bb7f81afd24b014a937af631f868bf8c11c278a1caf) |
| `primary.rr_set_group.rr_set.lb_record.name` | [primary.rr_set_group.rr_set.lb_record.name](resources--dns_zone--reference--group-003.md#canonical-092301468a776edec18b35e09de595cc76754d3e183ad7ea64d85436f95e90a0) |
| `primary.rr_set_group.rr_set.lb_record.value` | [primary.rr_set_group.rr_set.lb_record.value](resources--dns_zone--reference--group-003.md#canonical-8c09cb5e801be9d767f469e89945bcb793c3dd09d8126a958ab21c4b42fd7c08) |
| `primary.rr_set_group.rr_set.lb_record.value.name` | [primary.rr_set_group.rr_set.lb_record.value.name](resources--dns_zone--reference--group-003.md#canonical-9c17df4933a8e2f72a38ca9d5ff468100e685fb9cba7861f9c64fb2b13b18cf1) |
| `primary.rr_set_group.rr_set.lb_record.value.namespace` | [primary.rr_set_group.rr_set.lb_record.value.namespace](resources--dns_zone--reference--group-003.md#canonical-96912b2a8e8449d3f6d2d5c2dca72ce8b6c171c6cbbe60bd2c9de1bd6c46c9b9) |
| `primary.rr_set_group.rr_set.lb_record.value.tenant` | [primary.rr_set_group.rr_set.lb_record.value.tenant](resources--dns_zone--reference--group-003.md#canonical-7a766e0d21f09e5d37e5d86f5e7a627a80e821c75cd12d6b4f3d821242128862) |
| `primary.rr_set_group.rr_set.loc_record` | [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--reference--group-003.md#canonical-636c023c97807452576c77476421f5eb6bb35edf416e0167139921b53a2eb4ae) |
| `primary.rr_set_group.rr_set.loc_record.name` | [primary.rr_set_group.rr_set.loc_record.name](resources--dns_zone--reference--group-003.md#canonical-a117fbd3d9fa53fdadc0489b3e9b7fa1970fc1f9900fde483e46847dd2efae5d) |
| `primary.rr_set_group.rr_set.loc_record.values` | [primary.rr_set_group.rr_set.loc_record.values](resources--dns_zone--reference--group-003.md#canonical-f7d6025c5d0a342c8dc31082d1464284f59eb8c7bac6c824c8b66447a777e894) |
| `primary.rr_set_group.rr_set.loc_record.values.altitude` | [primary.rr_set_group.rr_set.loc_record.values.altitude](resources--dns_zone--reference--group-003.md#canonical-075b763ef87bfa9e7d04f05a0f075e873294e905dd2cfc04c0ee53c23732345c) |
| `primary.rr_set_group.rr_set.loc_record.values.horizontal_precision` | [primary.rr_set_group.rr_set.loc_record.values.horizontal_precision](resources--dns_zone--reference--group-003.md#canonical-3202ef51df86585adebd7356bbddc8a8f922daa59c3f9b810657a63137218a41) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.latitude_degree](resources--dns_zone--reference--group-003.md#canonical-cad3a6f735448a8648b20c0dba50d2b49245d81f1e5fb65df425cd075b80d5e5) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere](resources--dns_zone--reference--group-003.md#canonical-f472b64ba9f0475581d835682ebb6785745d0cc57999f63e8cc4f20bdb7c0953) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.latitude_minute](resources--dns_zone--reference--group-003.md#canonical-809a050700878ffe17f48214928ab27c9e51225f7434920edfde73b7bc261227) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_second` | [primary.rr_set_group.rr_set.loc_record.values.latitude_second](resources--dns_zone--reference--group-003.md#canonical-396f2a2b678dd4b0405a7d44f2e8c2221573dbb4024c1eab56b642d76b6f6bba) |
| `primary.rr_set_group.rr_set.loc_record.values.location_diameter` | [primary.rr_set_group.rr_set.loc_record.values.location_diameter](resources--dns_zone--reference--group-003.md#canonical-0071547aaf7adfc1b2628d11f6276a45161ba9af834cc15d0ca9aaf09e5734b1) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.longitude_degree](resources--dns_zone--reference--group-003.md#canonical-562f2dae7d0c91e049ca0f91942daec69bd656023e692f1dd9e1b175c8babeda) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere](resources--dns_zone--reference--group-003.md#canonical-f59cd940a3507533c5774de792684ab235e4dce3f30d5cb169d9a6a4e7491ec6) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.longitude_minute](resources--dns_zone--reference--group-003.md#canonical-ffd5f4d66b6a3eb3519dc0538369951c19867fd6321857f3f20020ed0431e474) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_second` | [primary.rr_set_group.rr_set.loc_record.values.longitude_second](resources--dns_zone--reference--group-003.md#canonical-03c35979f23f11c073add48109d06c6ac2a540d9d2414eb1f14c4adb22880e5d) |
| `primary.rr_set_group.rr_set.loc_record.values.vertical_precision` | [primary.rr_set_group.rr_set.loc_record.values.vertical_precision](resources--dns_zone--reference--group-003.md#canonical-cbf9fa4da7821153e718b3ff369c7183a0ca0f973869f840c0c3c42a0c961e01) |
| `primary.rr_set_group.rr_set.mx_record` | [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--reference--group-003.md#canonical-9e3370feec0c55291c08db96df2194baabb12c768b33cb3cddc10437e97ab47d) |
| `primary.rr_set_group.rr_set.mx_record.name` | [primary.rr_set_group.rr_set.mx_record.name](resources--dns_zone--reference--group-003.md#canonical-1236085d5ecf20eae4f5fa14fee88c0c0fd3c8ec84fbb68b0e31ad25560c4874) |
| `primary.rr_set_group.rr_set.mx_record.values` | [primary.rr_set_group.rr_set.mx_record.values](resources--dns_zone--reference--group-003.md#canonical-a2fe245d5e4931c056f19fb8e6ac8a6d8b72d4bde3cd29240fa02a603778ca85) |
| `primary.rr_set_group.rr_set.mx_record.values.domain` | [primary.rr_set_group.rr_set.mx_record.values.domain](resources--dns_zone--reference--group-003.md#canonical-5f4ba535d1edbeddd1c6356deb9bb9677d512444efd8afbc8cf8fce15227439d) |
| `primary.rr_set_group.rr_set.mx_record.values.priority` | [primary.rr_set_group.rr_set.mx_record.values.priority](resources--dns_zone--reference--group-003.md#canonical-d5dca6c95f9fa58ea364282e59c407fcc07f6bd37f78cdb05232b6fb0a5afff1) |
| `primary.rr_set_group.rr_set.naptr_record` | [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--reference--group-003.md#canonical-174a598251a89294d1034269e76dd69eaacf4124db2746df363fd297c318c3c5) |
| `primary.rr_set_group.rr_set.naptr_record.name` | [primary.rr_set_group.rr_set.naptr_record.name](resources--dns_zone--reference--group-003.md#canonical-601aeb4171436d4d9fc174284c44912a8eaef4383f8537141b4fed46d80a9306) |
| `primary.rr_set_group.rr_set.naptr_record.values` | [primary.rr_set_group.rr_set.naptr_record.values](resources--dns_zone--reference--group-003.md#canonical-2a82f0fb3a36ddd0612945fa5c2d664cd3bd1dd977332cbfba1de92d0e5574cf) |
| `primary.rr_set_group.rr_set.naptr_record.values.flags` | [primary.rr_set_group.rr_set.naptr_record.values.flags](resources--dns_zone--reference--group-003.md#canonical-80f9ed0b5ccdadc391bf09dd8c08d861d599974f723a3ab1c754ce5b16c30f2b) |
| `primary.rr_set_group.rr_set.naptr_record.values.order` | [primary.rr_set_group.rr_set.naptr_record.values.order](resources--dns_zone--reference--group-003.md#canonical-5d1cd4ad75ca4ccd4d3cef0b07f217d67027a1d0a410619adec6381b2dd8194d) |
| `primary.rr_set_group.rr_set.naptr_record.values.preference` | [primary.rr_set_group.rr_set.naptr_record.values.preference](resources--dns_zone--reference--group-003.md#canonical-8566f9301985ace9f00cb0ea1eba778ab4c6d99463e38d0b973d457663e4e166) |
| `primary.rr_set_group.rr_set.naptr_record.values.regexp` | [primary.rr_set_group.rr_set.naptr_record.values.regexp](resources--dns_zone--reference--group-003.md#canonical-4a68c86f0c4c7dd168a996f36ca9a7db961cc9a1fe6454253c9bb062f2958b30) |
| `primary.rr_set_group.rr_set.naptr_record.values.replacement` | [primary.rr_set_group.rr_set.naptr_record.values.replacement](resources--dns_zone--reference--group-003.md#canonical-62ee7daadb719740fd2b76bfabaac8e34803844ed2a11ce3668f37a165ae45e5) |
| `primary.rr_set_group.rr_set.naptr_record.values.service` | [primary.rr_set_group.rr_set.naptr_record.values.service](resources--dns_zone--reference--group-003.md#canonical-4c145849b97de94db32f9e2261d596f8b8892da67f55d59b9d7718bb5bf759c0) |
| `primary.rr_set_group.rr_set.ns_record` | [primary.rr_set_group.rr_set.ns_record](resources--dns_zone--reference--group-003.md#canonical-6d0da6f39f1c1a95debdc1c1c6d52a6893916066fb023ed7d8cbd02e89967db9) |
| `primary.rr_set_group.rr_set.ns_record.name` | [primary.rr_set_group.rr_set.ns_record.name](resources--dns_zone--reference--group-003.md#canonical-f82119045a704373f1efeea7b1f0404a1992d72a4e6b1739783610c596b52c97) |
| `primary.rr_set_group.rr_set.ns_record.values` | [primary.rr_set_group.rr_set.ns_record.values](resources--dns_zone--reference--group-003.md#canonical-871f908069ca417673b4062df2115e047d13e5eab2ff7a4fcf015af266861b32) |
| `primary.rr_set_group.rr_set.ptr_record` | [primary.rr_set_group.rr_set.ptr_record](resources--dns_zone--reference--group-003.md#canonical-c7a4bd4a0895c873dd89a3c37091fac0ea10b7944148db1708a9c7c05ace08e0) |
| `primary.rr_set_group.rr_set.ptr_record.name` | [primary.rr_set_group.rr_set.ptr_record.name](resources--dns_zone--reference--group-003.md#canonical-d0af55dd139061e43de94a5f083b92b88a8e046ab40fab0127d6c5456219898e) |
| `primary.rr_set_group.rr_set.ptr_record.values` | [primary.rr_set_group.rr_set.ptr_record.values](resources--dns_zone--reference--group-003.md#canonical-8bdb1ba55e569f47921b0f9936034d730c78e1a5ed1f9f826339692c2cfc69a7) |
| `primary.rr_set_group.rr_set.srv_record` | [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--reference--group-003.md#canonical-817b0e7c207fc263379fa1d01565484986050a2c402bfa62d6d323d349a6f417) |
| `primary.rr_set_group.rr_set.srv_record.name` | [primary.rr_set_group.rr_set.srv_record.name](resources--dns_zone--reference--group-003.md#canonical-ca746817db665410290c46f927cdd64edaecc6cc3875b097d1ee8d6dbcd3a7e1) |
| `primary.rr_set_group.rr_set.srv_record.values` | [primary.rr_set_group.rr_set.srv_record.values](resources--dns_zone--reference--group-003.md#canonical-3c5642a92bf376371be812a3d267061b806a34915c17603f8b93dc55555dce55) |
| `primary.rr_set_group.rr_set.srv_record.values.port` | [primary.rr_set_group.rr_set.srv_record.values.port](resources--dns_zone--reference--group-003.md#canonical-38819f0f92a3313ebd7c19221a6e3553e9db171c7fab6a29543aaaf1c4acee21) |
| `primary.rr_set_group.rr_set.srv_record.values.priority` | [primary.rr_set_group.rr_set.srv_record.values.priority](resources--dns_zone--reference--group-003.md#canonical-3c3eea868da3b9132eaa62ea62bbbe0117b650b1f9f70fc25edd49feb3abbb04) |
| `primary.rr_set_group.rr_set.srv_record.values.target` | [primary.rr_set_group.rr_set.srv_record.values.target](resources--dns_zone--reference--group-003.md#canonical-2cbc29682d5eeb9817e01110f128d129a79dbfefe2478c6334f3fddd9a612a8b) |
| `primary.rr_set_group.rr_set.srv_record.values.weight` | [primary.rr_set_group.rr_set.srv_record.values.weight](resources--dns_zone--reference--group-003.md#canonical-558d091d205638efe260332d23671ac0225c5c09e5e4f54bf1801cf076d59018) |
| `primary.rr_set_group.rr_set.sshfp_record` | [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-1967a16573c65b03213f601611ad04a692421a6a6f1b81d789ffbc9d27f766d9) |
| `primary.rr_set_group.rr_set.sshfp_record.name` | [primary.rr_set_group.rr_set.sshfp_record.name](resources--dns_zone--reference--group-003.md#canonical-67396cbc8ac3069dcc1d732631262d787d510e87d6af64fa6116f1c14729c7cf) |
| `primary.rr_set_group.rr_set.sshfp_record.values` | [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-f0428d7ad2c072f342aae4d74424ea1bc946e3741750d5234c23038ab5ffbeed) |
| `primary.rr_set_group.rr_set.sshfp_record.values.algorithm` | [primary.rr_set_group.rr_set.sshfp_record.values.algorithm](resources--dns_zone--reference--group-003.md#canonical-82a6a44f04985998dbeb4deffd9dfd784b47f9cef84175038a357b4d59fa2f35) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](resources--dns_zone--reference--group-003.md#canonical-5c84b2dbc3bc128a5308ba63d84f119482580c7f0156118f3e7472a4f99d34f9) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint](resources--dns_zone--reference--group-003.md#canonical-c69527b20ecdfd86d6680c0fff40b7cb3febe39304400f7b519eff403ad4e88a) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](resources--dns_zone--reference--group-003.md#canonical-90f9b4eca7ffcbce57b501b825f9aeda043f270956e838b6073eac5b46c0c466) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint](resources--dns_zone--reference--group-003.md#canonical-6390f91ea90efb936cbc34fbcc74c133445a50f607a35e79160cf274508dd2a1) |
| `primary.rr_set_group.rr_set.tlsa_record` | [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--reference--group-003.md#canonical-f4b10a1483b2c1444b579ca7ac76fe78c1a92d9dee3b1d3d389c1db2bafa6157) |
| `primary.rr_set_group.rr_set.tlsa_record.name` | [primary.rr_set_group.rr_set.tlsa_record.name](resources--dns_zone--reference--group-003.md#canonical-75d06f0a4375ff2e6c8e6e5606244fd6b8465707069521930a839d682c20b6c5) |
| `primary.rr_set_group.rr_set.tlsa_record.values` | [primary.rr_set_group.rr_set.tlsa_record.values](resources--dns_zone--reference--group-003.md#canonical-a5bcf0deed728e6c55cd10d5f902eb9866a51aeff88192086e85b4993f0df530) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data](resources--dns_zone--reference--group-003.md#canonical-85e5d88290496c08fc66922f9e50e30de2766d9a104059f7d94236338887526c) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage](resources--dns_zone--reference--group-003.md#canonical-45f20f0c9bd93c8a9743b2fb729702905cda4f7453bf7f13fa8f8e93bc97cea6) |
| `primary.rr_set_group.rr_set.tlsa_record.values.matching_type` | [primary.rr_set_group.rr_set.tlsa_record.values.matching_type](resources--dns_zone--reference--group-003.md#canonical-184f7d6bf598fe1eb725e4c6d8e97384a44bcf706d9cd305af747d68bf2ecfe1) |
| `primary.rr_set_group.rr_set.tlsa_record.values.selector` | [primary.rr_set_group.rr_set.tlsa_record.values.selector](resources--dns_zone--reference--group-003.md#canonical-6d005b3242721466cb53f22bcc059c0120183f799391761935f3aedbe97f08fc) |
| `primary.rr_set_group.rr_set.ttl` | [primary.rr_set_group.rr_set.ttl](resources--dns_zone--reference--group-002.md#canonical-a6e4f2f33b06dd83e43b40c32106049d1d4ee1159ad909cb7307028520494070) |
| `primary.rr_set_group.rr_set.txt_record` | [primary.rr_set_group.rr_set.txt_record](resources--dns_zone--reference--group-003.md#canonical-50f54e8c2155c9e2d42d4031039d03e156041ad2780e0cdec0efb2220ff78c7e) |
| `primary.rr_set_group.rr_set.txt_record.name` | [primary.rr_set_group.rr_set.txt_record.name](resources--dns_zone--reference--group-003.md#canonical-8c86994d06088715740813022cae3507b28558c19790e9f9be3b9a60abfe2124) |
| `primary.rr_set_group.rr_set.txt_record.values` | [primary.rr_set_group.rr_set.txt_record.values](resources--dns_zone--reference--group-003.md#canonical-6789051a40fa4494af829f64d671b6a4ddc6a9468581b91a23293dfc8e4722a4) |
| `primary.soa_parameters` | [primary.soa_parameters](resources--dns_zone--reference--group-003.md#canonical-99a3950ee58f7323132356a10be24888d9022566671717ccd57139a105d9eaa6) |
| `primary.soa_parameters.expire` | [primary.soa_parameters.expire](resources--dns_zone--reference--group-003.md#canonical-ec7b000fd4458b9312b9d2a9276a4972791b8aa11a22db7dfda35b52d2eb7235) |
| `primary.soa_parameters.negative_ttl` | [primary.soa_parameters.negative_ttl](resources--dns_zone--reference--group-003.md#canonical-b94507c79fdfc2f52cf7e0c1c6a31faa126f8cb37632c5bfd6e210e17cd44309) |
| `primary.soa_parameters.refresh` | [primary.soa_parameters.refresh](resources--dns_zone--reference--group-003.md#canonical-b6ea1b18927f2c4348fcd616ee9e4d5afa016462c5d82304c14bf363a0fd2a36) |
| `primary.soa_parameters.retry` | [primary.soa_parameters.retry](resources--dns_zone--reference--group-003.md#canonical-1eaf7d30c2e59d3b7afe14b6c2771c7bb412c9cc69779a0b2a1de34e214727e5) |
| `primary.soa_parameters.ttl` | [primary.soa_parameters.ttl](resources--dns_zone--reference--group-003.md#canonical-cd67c0aa5a10d3e7828d7681a8e21e86f36fbefb3684a63ff66fa6bb9777c880) |
| `secondary` | [secondary](resources--dns_zone--reference--group-003.md#canonical-2b11d4bd34678030dd7d0e737680d35191aa7af0d4c737ab4b790cdbc6aeb819) |
| `secondary.primary_servers` | [secondary.primary_servers](resources--dns_zone--reference--group-003.md#canonical-cae88882b862e89c7e51461eed76abc7cb2cd77780e4ee1b2a2bbbe12b3b00ca) |
| `secondary.tsig_key_algorithm` | [secondary.tsig_key_algorithm](resources--dns_zone--reference--group-003.md#canonical-9da4604c27a10e80196fcf20dac84bedfae8c71e6e470be38aa0f4ef965a5ecd) |
| `secondary.tsig_key_name` | [secondary.tsig_key_name](resources--dns_zone--reference--group-003.md#canonical-765f3c5b1dde3d9a808c80a986177b9032dc3779cd52eda6f054eb82c8689478) |
| `secondary.tsig_key_value` | [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-5ff70a5f1b2717ab7dca34ac281aaf5361a73042d3103d60c24a78c3c61f6579) |
| `secondary.tsig_key_value.blindfold_secret_info` | [secondary.tsig_key_value.blindfold_secret_info](resources--dns_zone--reference--group-003.md#canonical-f119a7ff1489aa5b282fc812809501b01cfacca27a44c9cf97537a6c7e23661b) |
| `secondary.tsig_key_value.blindfold_secret_info.decryption_provider` | [secondary.tsig_key_value.blindfold_secret_info.decryption_provider](resources--dns_zone--reference--group-003.md#canonical-8f9fa3b635860b2d8146b2e17cab38fb31654c13779a97d610ec5c6ede9eb46f) |
| `secondary.tsig_key_value.blindfold_secret_info.location` | [secondary.tsig_key_value.blindfold_secret_info.location](resources--dns_zone--reference--group-003.md#canonical-cbf02dfc34bd9372c79d9d21d31441817d87117a1df1f7aa8d22c24a73fcaefd) |
| `secondary.tsig_key_value.blindfold_secret_info.store_provider` | [secondary.tsig_key_value.blindfold_secret_info.store_provider](resources--dns_zone--reference--group-003.md#canonical-2d9fc3d2c35218a5221e72b30b00d0e091afede412db1d46070eb0ef1641cd21) |
| `secondary.tsig_key_value.clear_secret_info` | [secondary.tsig_key_value.clear_secret_info](resources--dns_zone--reference--group-003.md#canonical-8b79518335a7b1d15a02cace936e0a124b205b59e56b424deeb213b607720d2c) |
| `secondary.tsig_key_value.clear_secret_info.provider_ref` | [secondary.tsig_key_value.clear_secret_info.provider_ref](resources--dns_zone--reference--group-003.md#canonical-377a887f8fa8ff470cf33d563a948a3112f49014e732aeacc900ce431351f2c7) |
| `secondary.tsig_key_value.clear_secret_info.url` | [secondary.tsig_key_value.clear_secret_info.url](resources--dns_zone--reference--group-003.md#canonical-aa2eae8a5166e7cad4ab8797e66fc837cd46187724615c1e838485e3971be751) |
| `timeouts` | [timeouts](resources--dns_zone--reference--group-003.md#canonical-30e39fcb34f8437dcda977f5d0332283e66222720e8328865080d09fea604511) |
| `timeouts.create` | [timeouts.create](resources--dns_zone--reference--group-003.md#canonical-3076d76a9317e888315a3dd82513307bb4f512be658355d81908792a2deea226) |
| `timeouts.delete` | [timeouts.delete](resources--dns_zone--reference--group-003.md#canonical-f944a9b66f978d13077ce0332d13de6a271ab5753b9eca7c5fa39514a327b469) |
| `timeouts.read` | [timeouts.read](resources--dns_zone--reference--group-003.md#canonical-393204c14936e2ae4c4d32a9932b5476d51310e569daab5ed92c4c36cb8865c8) |
| `timeouts.update` | [timeouts.update](resources--dns_zone--reference--group-003.md#canonical-fb12c93332b9fccf35b695cd6df6ddb9e37e8cc9fd514e7eee5b3674e8f973b9) |

<a id="canonical-3b05cb3c2859e8f6fca7b94a48f57a80132a446b26c4cffb23e9c84194f6f704"></a>

## Next pages — Property reference / 6f4349d8c84a / 12

- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-915ef6a26cadd3cdb1d03e8f1f87b29e6e5292057e8da86defb920ce53ea35c9)
- [timeouts](resources--dns_zone--reference--group-003.md#canonical-592d8ee330c4fcf5e49e421377380abc8cd9a954a4313bed2da72b7dc44025bf)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6546cf9d13e1a6e309dedb90de1f85e9cc301841128aef61901a98206ea58913"></a>

## primary — primary / 4d761ac56d62 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- primary

<a id="canonical-f0df71152037b454f5fd2a1339729cbe2aa0d1f15cbeb8ada01eab617ed14435"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: primary, secondary\] PrimaryDNSCreateSpecType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_soa_parameters",
    "soa_parameters")}
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
  "x-ves-oneof-field-soa_record_parameters_choice": "[\"default_soa_parameters\",\"soa_parameters\"]"
}
```

OneOf alternatives in this subsection:

- [primary](resources--dns_zone--reference--group-001.md#canonical-f0df71152037b454f5fd2a1339729cbe2aa0d1f15cbeb8ada01eab617ed14435)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-2b11d4bd34678030dd7d0e737680d35191aa7af0d4c737ab4b790cdbc6aeb819)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
primary {
  # Configure direct properties listed below.
}
```

<a id="canonical-abd3ca133d4d85df15de1f2ffc96ef885d3f2ea1b10b6532e9995ca58131d091"></a>

## Direct properties — primary / 4d761ac56d62 / 3

<a id="canonical-58855fea13abbe96c61a63d99aaea640a2d4225ad44a8572c16f61d357895921"></a>

<a id="canonical-fd3632b1f58354932480a668ffc06c53edca1b306e8f6456d9356e90e659f872"></a>

## allow_http_lb_managed_records property — primary / 4d761ac56d62 / 4

Type: `"bool"`. Optional.

Option to allow user-created HTTP, TCP, and CDN load balancer related resource records to be
automatically managed in a protected RRset.

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

- [default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d): complete subsection reference.

- [default_soa_parameters](resources--dns_zone--reference--group-002.md#canonical-011cd0aebba1ce421435bec972cff46926c8a27bfd644b58dd735e03f67530f1): complete subsection reference.

- [dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-a0163bba2e298b8beeb8b130e64bd81043610e52aeb892e347bb2aab203f53d2): complete subsection reference.

- [rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94): complete subsection reference.

- [soa_parameters](resources--dns_zone--reference--group-003.md#canonical-4888a3369cbc116c1b65198b284eb32389412885bed470f1c01e31f21955301c): complete subsection reference.

<a id="canonical-150dc9a69d4f5331eb56301e1200af7fcec29445d7a58c4f9f30fe869b665b34"></a>

## Next pages — primary / 4d761ac56d62 / 5

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_soa_parameters](resources--dns_zone--reference--group-002.md#canonical-011cd0aebba1ce421435bec972cff46926c8a27bfd644b58dd735e03f67530f1)
- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-a0163bba2e298b8beeb8b130e64bd81043610e52aeb892e347bb2aab203f53d2)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.soa_parameters](resources--dns_zone--reference--group-003.md#canonical-4888a3369cbc116c1b65198b284eb32389412885bed470f1c01e31f21955301c)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86c8a7ddaec0e1676c099e89681fdaf25adb7d3c154e42a8b1f07abdc0504641"></a>

## primary.default_rr_set_group — primary.default_rr_set_group / af969c559c6e / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- primary.default_rr_set_group

<a id="canonical-264380bf2df9264ada7036448f23d9e49047139c6c0a0c209a7e48a3b073375e"></a>

Type: `"object"`. list nested block, Optional.

Add and manage DNS resource record sets part of Default set group.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ttl"),
  validators.ConflictingListObjectAttributes("a_record",
    "aaaa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("tlsa_record",
    "txt_record")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
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
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

Terraform syntax:

```terraform
default_rr_set_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-9c36b43ebf16a696f45ae5aeafbe0ecef25909d7e533cab5ed4ee9b6b2e94d00"></a>

## Direct properties — primary.default_rr_set_group / af969c559c6e / 3

- [a_record](resources--dns_zone--reference--group-001.md#canonical-cb1254a2f7ef35dc2f0f3fd921e365589b34dcdc3628e3ffbd891a0f1d8f4adb): complete subsection reference.

- [aaaa_record](resources--dns_zone--reference--group-001.md#canonical-4df113dc32862c88be04572548acd25169eacb006983d134eac3fb446ec62b3c): complete subsection reference.

- [afsdb_record](resources--dns_zone--reference--group-001.md#canonical-a69a70ed761e7a1d9741be052f5614d3e18c16578a4b33d18e979035d2d4e0ad): complete subsection reference.

- [alias_record](resources--dns_zone--reference--group-001.md#canonical-745055fcacba5a73de051b599651b7825ae2e8322365b3b4dad8c5903c892981): complete subsection reference.

- [caa_record](resources--dns_zone--reference--group-001.md#canonical-32f71c9ef6262dec136a95c07d164118c7675c46b0b59e2cde8f00731f1fd37d): complete subsection reference.

- [cds_record](resources--dns_zone--reference--group-001.md#canonical-794f910af9f92715f1b01c03d4bf830e8f08186fae3d28c3d8e7b0bc155cda37): complete subsection reference.

- [cert_record](resources--dns_zone--reference--group-001.md#canonical-c3e649004dc62fa19d277fa4563448bf796df6c0bce76da41fe6872c5748c611): complete subsection reference.

- [cname_record](resources--dns_zone--reference--group-001.md#canonical-8ef00259f52c959e585f74aed99b7ff0449ac0da39979cbc38adf23ab232805c): complete subsection reference.

<a id="canonical-b3bf216484665be82c1bbeeb64e18359bdb240a7d228c49a4bc20b6989abf80d"></a>

<a id="canonical-cb35ca219dd7ec4e3c70016a6b30a7ae0e399faf165addd8c2a1de98eeda89ed"></a>

## description_spec property — primary.default_rr_set_group / af969c559c6e / 4

Type: `"string"`. Optional.

Comment. Human-readable description text

- [ds_record](resources--dns_zone--reference--group-001.md#canonical-a95a01f6aa08e943ea1267e9b696669ccf3c5fcf7acfd21343f4a70bf68b5673): complete subsection reference.

- [eui48_record](resources--dns_zone--reference--group-002.md#canonical-d0061623363fa2c7be32ad499afa216264a6374fea4ded410c8eb10ecae277a5): complete subsection reference.

- [eui64_record](resources--dns_zone--reference--group-002.md#canonical-34b84b028eff581da400c6b35883d2c8d1506ce9b2ab4580ae71ad7b2311a5f5): complete subsection reference.

- [lb_record](resources--dns_zone--reference--group-002.md#canonical-802c5376839db15e30f65fef43748962271c0826de4f85bc9c570e4b2968f34c): complete subsection reference.

- [loc_record](resources--dns_zone--reference--group-002.md#canonical-01f7a3b77dbf0a6d8939bf5a77c55698576dea8907d7f0a4c95d5604d913baff): complete subsection reference.

- [mx_record](resources--dns_zone--reference--group-002.md#canonical-42513c672c75487deb183579303e80a4c071e22205160f1118a0cea4e78708f6): complete subsection reference.

- [naptr_record](resources--dns_zone--reference--group-002.md#canonical-29861e75cf9362a30496f5183138a125bae01dd5822188db7f03bc58fccf8ff2): complete subsection reference.

- [ns_record](resources--dns_zone--reference--group-002.md#canonical-781801294746217c2a3523dd3f86ab70b44e13a8ea217f252c073dad72ab2732): complete subsection reference.

- [ptr_record](resources--dns_zone--reference--group-002.md#canonical-cc42794f9714583aafa8ea4d6820d1615e7962dce5b458be9befc0a7a0b921ac): complete subsection reference.

- [srv_record](resources--dns_zone--reference--group-002.md#canonical-1706c61fb364648d60d5a0aecef6e3fb6011b6a966e1b65b079e1ef8f7de562f): complete subsection reference.

- [sshfp_record](resources--dns_zone--reference--group-002.md#canonical-027f21fbf6635b09d7990db940a88858d0819a3562710a5cfec958515218b6e9): complete subsection reference.

- [tlsa_record](resources--dns_zone--reference--group-002.md#canonical-549151928d884e1b5ee19bbe822cdb0e8a9cebaf669d1ae353e4aa003c5bb1da): complete subsection reference.

<a id="canonical-0df548e4b0e455d90e09f0e2801d21ecb478de9d25b5ecd9f14817491cd681db"></a>

<a id="canonical-351730e3695af4896ddeaaa009cbd655768d181996763bad7a1ed7be3169b008"></a>

## ttl property — primary.default_rr_set_group / af969c559c6e / 5

Type: `"number"`. Optional.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(60, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

- [txt_record](resources--dns_zone--reference--group-002.md#canonical-fdb7631ec4d613093c3880c355db38c0eec42adcd38295c86269c3657197c858): complete subsection reference.

<a id="canonical-f895ed1bcc8279d12b26a6aafe962845c4add69849ae0ef8aa00f20f25b5347b"></a>

## Next pages — primary.default_rr_set_group / af969c559c6e / 6

- [primary.default_rr_set_group.a_record](resources--dns_zone--reference--group-001.md#canonical-cb1254a2f7ef35dc2f0f3fd921e365589b34dcdc3628e3ffbd891a0f1d8f4adb)
- [primary.default_rr_set_group.aaaa_record](resources--dns_zone--reference--group-001.md#canonical-4df113dc32862c88be04572548acd25169eacb006983d134eac3fb446ec62b3c)
- [primary.default_rr_set_group.afsdb_record](resources--dns_zone--reference--group-001.md#canonical-a69a70ed761e7a1d9741be052f5614d3e18c16578a4b33d18e979035d2d4e0ad)
- [primary.default_rr_set_group.alias_record](resources--dns_zone--reference--group-001.md#canonical-745055fcacba5a73de051b599651b7825ae2e8322365b3b4dad8c5903c892981)
- [primary.default_rr_set_group.caa_record](resources--dns_zone--reference--group-001.md#canonical-32f71c9ef6262dec136a95c07d164118c7675c46b0b59e2cde8f00731f1fd37d)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-794f910af9f92715f1b01c03d4bf830e8f08186fae3d28c3d8e7b0bc155cda37)
- [primary.default_rr_set_group.cert_record](resources--dns_zone--reference--group-001.md#canonical-c3e649004dc62fa19d277fa4563448bf796df6c0bce76da41fe6872c5748c611)
- [primary.default_rr_set_group.cname_record](resources--dns_zone--reference--group-001.md#canonical-8ef00259f52c959e585f74aed99b7ff0449ac0da39979cbc38adf23ab232805c)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-001.md#canonical-a95a01f6aa08e943ea1267e9b696669ccf3c5fcf7acfd21343f4a70bf68b5673)
- [primary.default_rr_set_group.eui48_record](resources--dns_zone--reference--group-002.md#canonical-d0061623363fa2c7be32ad499afa216264a6374fea4ded410c8eb10ecae277a5)
- [primary.default_rr_set_group.eui64_record](resources--dns_zone--reference--group-002.md#canonical-34b84b028eff581da400c6b35883d2c8d1506ce9b2ab4580ae71ad7b2311a5f5)
- [primary.default_rr_set_group.lb_record](resources--dns_zone--reference--group-002.md#canonical-802c5376839db15e30f65fef43748962271c0826de4f85bc9c570e4b2968f34c)
- [primary.default_rr_set_group.loc_record](resources--dns_zone--reference--group-002.md#canonical-01f7a3b77dbf0a6d8939bf5a77c55698576dea8907d7f0a4c95d5604d913baff)
- [primary.default_rr_set_group.mx_record](resources--dns_zone--reference--group-002.md#canonical-42513c672c75487deb183579303e80a4c071e22205160f1118a0cea4e78708f6)
- [primary.default_rr_set_group.naptr_record](resources--dns_zone--reference--group-002.md#canonical-29861e75cf9362a30496f5183138a125bae01dd5822188db7f03bc58fccf8ff2)
- [primary.default_rr_set_group.ns_record](resources--dns_zone--reference--group-002.md#canonical-781801294746217c2a3523dd3f86ab70b44e13a8ea217f252c073dad72ab2732)
- [primary.default_rr_set_group.ptr_record](resources--dns_zone--reference--group-002.md#canonical-cc42794f9714583aafa8ea4d6820d1615e7962dce5b458be9befc0a7a0b921ac)
- [primary.default_rr_set_group.srv_record](resources--dns_zone--reference--group-002.md#canonical-1706c61fb364648d60d5a0aecef6e3fb6011b6a966e1b65b079e1ef8f7de562f)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-027f21fbf6635b09d7990db940a88858d0819a3562710a5cfec958515218b6e9)
- [primary.default_rr_set_group.tlsa_record](resources--dns_zone--reference--group-002.md#canonical-549151928d884e1b5ee19bbe822cdb0e8a9cebaf669d1ae353e4aa003c5bb1da)
- [primary.default_rr_set_group.txt_record](resources--dns_zone--reference--group-002.md#canonical-fdb7631ec4d613093c3880c355db38c0eec42adcd38295c86269c3657197c858)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-cb1254a2f7ef35dc2f0f3fd921e365589b34dcdc3628e3ffbd891a0f1d8f4adb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-810713290c4ec7cb91f32db0bf7e0d54ea863cbec1939f4554133bd16373a006"></a>

## primary.default_rr_set_group.a_record — primary.default_rr_set_group.a_record / 336a68ee2d04 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.a_record

<a id="canonical-bc515a86c7bd5786c45d78ea53a7da119d458d08ced64c98b27442cbdc59cf3a"></a>

Type: `"object"`. single nested block, Optional.

DNSAResourceRecord. A Records

Upstream description:

A Records

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
a_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-dc53ef359972315e7d3117bb542d205645a81b895ce44b26a00dad90e16fb16d"></a>

## Direct properties — primary.default_rr_set_group.a_record / 336a68ee2d04 / 3

<a id="canonical-9815ed0fac4144419a01fa239953df034ca7ad531b505bf565adfd53da56dca9"></a>

<a id="canonical-2c8705e8ba41c533ff77d7eb0eb8779ccab15f3d1b2fc2020b86852e336fc88c"></a>

## name property — primary.default_rr_set_group.a_record / 336a68ee2d04 / 4

Type: `"string"`. Optional.

Record name, please provide only the specific subdomain or record name without the base domain.

Upstream description:

A Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-fc68bcb484297dea3f074b7d477a43b3c89cd74184b566d1ca83d6a892575b9f"></a>

<a id="canonical-bec65cd509887c347fcf4f737b53f8aa19244ff66121e1ab13106c76bdd1b90c"></a>

## values property — primary.default_rr_set_group.a_record / 336a68ee2d04 / 5

Type: `["list", "string"]`. Optional.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

Upstream description:

A valid IPv4 address, for example: 192.0.2.242.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b2787366ab402156c26dbbbac4012ad2544ae80eaed771689e7024d92473bf47"></a>

## Next pages — primary.default_rr_set_group.a_record / 336a68ee2d04 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-4df113dc32862c88be04572548acd25169eacb006983d134eac3fb446ec62b3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8501c2f8dfa393c6c6f045ab77bcd32b61f4db5b856c1c3c237a6e4dd7f8de53"></a>

## primary.default_rr_set_group.aaaa_record — primary.default_rr_set_group.aaaa_record / 97c95b75806a / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.aaaa_record

<a id="canonical-2e858ac2aebe340e0d8720d6171b4361ee24d6d7753ee1e1ae5c6571903a7d7f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aaaa record.

Upstream description:

RecordSet for AAAA Records.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
aaaa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-aecae983399bb5c1a279e3e884783e38d62770830be249d2842b0e3395e5853e"></a>

## Direct properties — primary.default_rr_set_group.aaaa_record / 97c95b75806a / 3

<a id="canonical-68c8ff285be349e835e304b76da836ec6453f38fd346fcb2dcba6103b309bb94"></a>

<a id="canonical-82d47d617b8b047e11dffcd8f7f01340ce256f4036617d1efc6f9e0a440169b7"></a>

## name property — primary.default_rr_set_group.aaaa_record / 97c95b75806a / 4

Type: `"string"`. Optional.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-e140258a76b9b94181d55bce9c51d1d9a2cff7c2d06bafa130938c1ab0879666"></a>

<a id="canonical-55363e28b0b9939b751ac879d692d8b50aa68ce20a750b1185a906335220af37"></a>

## values property — primary.default_rr_set_group.aaaa_record / 97c95b75806a / 5

Type: `["list", "string"]`. Optional.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Upstream description:

A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1d76bfa392a559e3eca060c5f8b0a5a366a5958690005048b069d309eab3b0e4"></a>

## Next pages — primary.default_rr_set_group.aaaa_record / 97c95b75806a / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-a69a70ed761e7a1d9741be052f5614d3e18c16578a4b33d18e979035d2d4e0ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-497649a0872aad339afc9893db52d9d80a4aff75f5958008277ecd4d7a372ed9"></a>

## primary.default_rr_set_group.afsdb_record — primary.default_rr_set_group.afsdb_record / 9021b5d68f30 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.afsdb_record

<a id="canonical-fe5af3bb5c6d98d77036196a82f14ecefdf16fe4139b7af831f555889fde3974"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for afsdb record.

Upstream description:

DNS AFSDB Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
afsdb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-d050cb20355ee6d096a8e398e7d62b841f9b50c8fbaf9e596f60ddeca597ee6e"></a>

## Direct properties — primary.default_rr_set_group.afsdb_record / 9021b5d68f30 / 3

<a id="canonical-6467bf4bbb9917a5148418f7c9b74e30c9dddff165f088310862574b20d97bc2"></a>

<a id="canonical-cd4f4998d65b2252c91d0669ca1a7d7fd67d772fea3f41a9fcc0f0eba6c41c81"></a>

## name property — primary.default_rr_set_group.afsdb_record / 9021b5d68f30 / 4

Type: `"string"`. Optional.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-001.md#canonical-fd73030e74e851f4c1da892a81edd5ac30b7b28f043559212201713706531114): complete subsection reference.

<a id="canonical-ad716b3214939aa9e4cfc404760856f3abd747b14b6750f05c212166d37e3ec1"></a>

## Next pages — primary.default_rr_set_group.afsdb_record / 9021b5d68f30 / 5

- [primary.default_rr_set_group.afsdb_record.values](resources--dns_zone--reference--group-001.md#canonical-fd73030e74e851f4c1da892a81edd5ac30b7b28f043559212201713706531114)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-fd73030e74e851f4c1da892a81edd5ac30b7b28f043559212201713706531114"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c00a29d1abd10a52053758a919d8e7169ca4e80c16e68fe5196c35220ef4a12c"></a>

## primary.default_rr_set_group.afsdb_record.values — primary.default_rr_set_group.afsdb_record.values / f6320bf71514 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.afsdb_record](resources--dns_zone--reference--group-001.md#canonical-a69a70ed761e7a1d9741be052f5614d3e18c16578a4b33d18e979035d2d4e0ad)
- primary.default_rr_set_group.afsdb_record.values

<a id="canonical-54de531e70e2280e4555e80b4221ca8003b2374b5db05984b98d08c83bce34be"></a>

Type: `"object"`. list nested block, Optional.

AFSDB Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("hostname")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-241bc3e02ed1d7398b9b43d9b150e5b5f53902e487c0917fff7141b07f124a53"></a>

## Direct properties — primary.default_rr_set_group.afsdb_record.values / f6320bf71514 / 3

<a id="canonical-658aa1e4f89509ff7d11b4ea3102457c6a3ff1d0b05270e7f04add6fb9b46d89"></a>

<a id="canonical-0820da0fe242bf3b95b971dce7696fe768ede114c082c074d076d2557da96751"></a>

## hostname property — primary.default_rr_set_group.afsdb_record.values / f6320bf71514 / 4

Type: `"string"`. Optional.

Server name of the AFS cell database server or the DCE name server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-de8cffd9c56157fba0bd11d23964094d2ea820e8631912560201ebbfbe7f1f5c"></a>

<a id="canonical-3d5d0ec486a515d0214bed4e60a9c6626d1ec58b5d02e675885b19ee06149082"></a>

## subtype property — primary.default_rr_set_group.afsdb_record.values / f6320bf71514 / 5

Type: `"string"`. Optional.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Upstream description:

AFS Volume Location Server or DCE Authentication Server.

&#8203;- NONE: NONE

&#8203;- AFSVolumeLocationServer: AFS Volume Location Server

&#8203;- DCEAuthenticationServer: DCE Authentication Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b13ae5168d9ba0d5433b6f247b0f69c03efafd805fb0d521844c332ab9e82675"></a>

## Next pages — primary.default_rr_set_group.afsdb_record.values / f6320bf71514 / 6

- [primary.default_rr_set_group.afsdb_record](resources--dns_zone--reference--group-001.md#canonical-a69a70ed761e7a1d9741be052f5614d3e18c16578a4b33d18e979035d2d4e0ad)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-745055fcacba5a73de051b599651b7825ae2e8322365b3b4dad8c5903c892981"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-034996a0ef1cbb4bea036a689481e7d587956104db5e66d666a861f781b3d970"></a>

## primary.default_rr_set_group.alias_record — primary.default_rr_set_group.alias_record / cfb5d39271ac / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.alias_record

<a id="canonical-5a3c83fdec994eadf604c445e075882be2a69cc8f77a070e537294662b2386e8"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for alias record.

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
alias_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-c8ea6383aa1cc4e468f9ff4ee78dd0e28c24a168728def29614191aa4cdc1780"></a>

## Direct properties — primary.default_rr_set_group.alias_record / cfb5d39271ac / 3

<a id="canonical-9637fdf46e48d862d25345b956b40615802b7e0a6f153ac8e5d4b20de8503d8f"></a>

<a id="canonical-1e7d274b74fdc6305723e1cfc60c094bccd5c23884b2aaf3aae004458018a81e"></a>

## value property — primary.default_rr_set_group.alias_record / cfb5d39271ac / 4

Type: `"string"`. Optional.

Domain. A valid domain name, for example: example.com.

Upstream description:

A valid domain name, for example: example.com.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-73a1c2e20095cd6ae62eac326f1d70e8e46b7895b0bef623d3107182f0263f21"></a>

## Next pages — primary.default_rr_set_group.alias_record / cfb5d39271ac / 5

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-32f71c9ef6262dec136a95c07d164118c7675c46b0b59e2cde8f00731f1fd37d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63cb86ceef279b7b6a3aeb0ce5ed5accb4f410e6c86cd562803d60e2303e1f95"></a>

## primary.default_rr_set_group.caa_record — primary.default_rr_set_group.caa_record / 24e74aecaeb6 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.caa_record

<a id="canonical-b8d44d2e87764a9b2c90680730ec575e0407d7ee62253655216c862d1b057a5a"></a>

Type: `"object"`. single nested block, Optional.

DNSCAAResourceRecord.

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
caa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-e7a9674172ea4913dd3ae2604f8c34914d543e99c77faa99e1e8c2542d2b3c00"></a>

## Direct properties — primary.default_rr_set_group.caa_record / 24e74aecaeb6 / 3

<a id="canonical-b77b32cbfa0dfedcbdb719240443c6a62a726d0dd4fbf178f9a8376bd5596df8"></a>

<a id="canonical-e74568995a3b31b1b2c38877099585778b1b4880984edf26e3edb0b928395985"></a>

## name property — primary.default_rr_set_group.caa_record / 24e74aecaeb6 / 4

Type: `"string"`. Optional.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-001.md#canonical-3bfdc57817178309dc09e7cb0932c0c2b43347deab97f3314b728276801d33fa): complete subsection reference.

<a id="canonical-832fb67007c5b9a66536d917f2c7b352f4fe9dee60558e7cd270ffd09e4b93e3"></a>

## Next pages — primary.default_rr_set_group.caa_record / 24e74aecaeb6 / 5

- [primary.default_rr_set_group.caa_record.values](resources--dns_zone--reference--group-001.md#canonical-3bfdc57817178309dc09e7cb0932c0c2b43347deab97f3314b728276801d33fa)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-3bfdc57817178309dc09e7cb0932c0c2b43347deab97f3314b728276801d33fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5c38d90626bc20d07e9f2c9ab3880a1fcc0090848eedb73c1d93aca09a0061a"></a>

## primary.default_rr_set_group.caa_record.values — primary.default_rr_set_group.caa_record.values / 0209f5d87d1a / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.caa_record](resources--dns_zone--reference--group-001.md#canonical-32f71c9ef6262dec136a95c07d164118c7675c46b0b59e2cde8f00731f1fd37d)
- primary.default_rr_set_group.caa_record.values

<a id="canonical-64fb576a89a17ed8e2428602af5391e8cff4cc935fc264dfba9538c4d416092f"></a>

Type: `"object"`. list nested block, Optional.

CAA Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-b922aa8517b18f974ebed6373d9dc746ad0199b6a6ff943cd09e8b35c425070d"></a>

## Direct properties — primary.default_rr_set_group.caa_record.values / 0209f5d87d1a / 3

<a id="canonical-5fad0d88b8648d1237ada6db08df2da9ea5aaf3d02c3dc79fe0a06f458a42d67"></a>

<a id="canonical-949b62dbc86fdb1082d2ee5b4d0192cf9e872a2e070bb54a182b33b72f076593"></a>

## flags property — primary.default_rr_set_group.caa_record.values / 0209f5d87d1a / 4

Type: `"number"`. Optional.

Flag should be an integer between 0 and 255.

Upstream description:

This flag should be an integer between 0 and 255.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-9eb39ef685866bf0f1f7000b856b5c1a98fdb496fd17a305659298c5521d670f"></a>

<a id="canonical-cac7fdcf3a15d18cd1f83d1b89c9100d0ab03b4f5e556dd20b10ca678774f5a5"></a>

## tag property — primary.default_rr_set_group.caa_record.values / 0209f5d87d1a / 5

Type: `"string"`. Optional.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Upstream description:

Tag for categorization and filtering

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("issue",
    "issuewild",
    "iodef"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="canonical-d5c1c9e6498512ee919e1691ac59f08a888121cd6d3bc792ef77ed09a539b9b1"></a>

<a id="canonical-47f6a0d9e2b2a626ed88e0ca00ace24933f623002a7c54d23998a9dba7481af2"></a>

## value property — primary.default_rr_set_group.caa_record.values / 0209f5d87d1a / 6

Type: `"string"`. Optional.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-d82e1d59eff93fe31a33f89e4bc291d94860fe8b59437fa0083be4574b198fb3"></a>

## Next pages — primary.default_rr_set_group.caa_record.values / 0209f5d87d1a / 7

- [primary.default_rr_set_group.caa_record](resources--dns_zone--reference--group-001.md#canonical-32f71c9ef6262dec136a95c07d164118c7675c46b0b59e2cde8f00731f1fd37d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-794f910af9f92715f1b01c03d4bf830e8f08186fae3d28c3d8e7b0bc155cda37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-275736e74f37ca594bc39b56d0d3f3ff40ab5d53adc3e009683d90582ea6e9b6"></a>

## primary.default_rr_set_group.cds_record — primary.default_rr_set_group.cds_record / 8edf911c8b4d / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.cds_record

<a id="canonical-0796127575044d96a7fcf8811ee2b4c1f8bcd15271301e18df7974373863b0b7"></a>

Type: `"object"`. single nested block, Optional.

DNS CDS Record. DNS CDS Record.

Upstream description:

DNS CDS Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
cds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b7ad37970ded4560246beb21cce928ba825c91191abc00a4d56932517dabc7a"></a>

## Direct properties — primary.default_rr_set_group.cds_record / 8edf911c8b4d / 3

<a id="canonical-92a08edfc0e5c442570b332bab4da398c888db6d05f7713aac3899f7db907e21"></a>

<a id="canonical-742052206b55077bcce7c3df5010fc3abda78cd17088c9bb2d612a3f7efdb91e"></a>

## name property — primary.default_rr_set_group.cds_record / 8edf911c8b4d / 4

Type: `"string"`. Optional.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-001.md#canonical-a142e4b9d79d18af2d2f44aca68f97521046c0d578ef61f2d2f12c23f5b5c452): complete subsection reference.

<a id="canonical-2bef6397fc19536ae9b9577f4f94077ef1f3a1f178a7814ac588a2a02096bf42"></a>

## Next pages — primary.default_rr_set_group.cds_record / 8edf911c8b4d / 5

- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-a142e4b9d79d18af2d2f44aca68f97521046c0d578ef61f2d2f12c23f5b5c452)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-a142e4b9d79d18af2d2f44aca68f97521046c0d578ef61f2d2f12c23f5b5c452"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-882549a2f6732ff531bb673cc2dc153fa088d89ac3a83413694137f96759eff2"></a>

## primary.default_rr_set_group.cds_record.values — primary.default_rr_set_group.cds_record.values / 58436460285c / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-794f910af9f92715f1b01c03d4bf830e8f08186fae3d28c3d8e7b0bc155cda37)
- primary.default_rr_set_group.cds_record.values

<a id="canonical-c83ad09a767e5950e115f6867f002b6416bb279e3ef56c5b081f7f39deabae59"></a>

Type: `"object"`. list nested block, Optional.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key_tag"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha256_digest"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha384_digest"),
  validators.ConflictingListObjectAttributes("sha256_digest",
    "sha384_digest")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-d25a71511f212bc820a9c54c0b9ad9b5b67e789e577579ae8f22f8fd48c984d3"></a>

## Direct properties — primary.default_rr_set_group.cds_record.values / 58436460285c / 3

<a id="canonical-f3fd150f9f78da28539e8404efe49e16d46d5de179f65862cc6f933cd8734d5a"></a>

<a id="canonical-d9c331b312dfa481af0d4ad057303fdaef5170e54649ad78d3fa4e282f6e9d3e"></a>

## ds_key_algorithm property — primary.default_rr_set_group.cds_record.values / 58436460285c / 4

Type: `"string"`. Optional.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-51c0e71b3fc9f06f65b1306d4f5300b139da322f75cafa849dd08720698e9f38"></a>

<a id="canonical-35a39c7f729f7af4dd05480509862a6c94056080f9f762ec63b0b45598ad96b7"></a>

## key_tag property — primary.default_rr_set_group.cds_record.values / 58436460285c / 5

Type: `"number"`. Optional.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](resources--dns_zone--reference--group-001.md#canonical-1547acb3b331b4e81b7b0f30af9eae5a8a502c06b7aa2227956bf88fcd21f8b9): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-001.md#canonical-ca7a3bf29db338ce49850f066809c0d4dcd5327d54809683968c80a04f907919): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-001.md#canonical-a20c4875d92ad2ae11a0aaadcb9673680530ce10f643a27c77e6e3729c5fcbb6): complete subsection reference.

<a id="canonical-38c20eea1b327a5c61fd69dbeb3dbb8e003c7b566367047478795679d1f0b090"></a>

## Next pages — primary.default_rr_set_group.cds_record.values / 58436460285c / 6

- [primary.default_rr_set_group.cds_record.values.sha1_digest](resources--dns_zone--reference--group-001.md#canonical-1547acb3b331b4e81b7b0f30af9eae5a8a502c06b7aa2227956bf88fcd21f8b9)
- [primary.default_rr_set_group.cds_record.values.sha256_digest](resources--dns_zone--reference--group-001.md#canonical-ca7a3bf29db338ce49850f066809c0d4dcd5327d54809683968c80a04f907919)
- [primary.default_rr_set_group.cds_record.values.sha384_digest](resources--dns_zone--reference--group-001.md#canonical-a20c4875d92ad2ae11a0aaadcb9673680530ce10f643a27c77e6e3729c5fcbb6)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-794f910af9f92715f1b01c03d4bf830e8f08186fae3d28c3d8e7b0bc155cda37)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-1547acb3b331b4e81b7b0f30af9eae5a8a502c06b7aa2227956bf88fcd21f8b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a15eb069abda75cbcebe17f6c17963f4ba2ae4a47660f1c8ca985661c97faba2"></a>

## primary.default_rr_set_group.cds_record.values.sha1_digest — primary.default_rr_set_group.cds_record.values.sha1_digest / 02d90d089aa0 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-794f910af9f92715f1b01c03d4bf830e8f08186fae3d28c3d8e7b0bc155cda37)
- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-a142e4b9d79d18af2d2f44aca68f97521046c0d578ef61f2d2f12c23f5b5c452)
- primary.default_rr_set_group.cds_record.values.sha1_digest

<a id="canonical-39c73459a239f389b4d6a9d7e7c8fc0736a471202bdba84b36d82a9cd91cd676"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha1_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-b6b82c86df5169a932cf33489f9997781ae8f6f84ef0b3468ff809deb9424164"></a>

## Direct properties — primary.default_rr_set_group.cds_record.values.sha1_digest / 02d90d089aa0 / 3

<a id="canonical-b934967095930068e105f4ab0cb3b0bf07d92b8465100dfad35956ca1e44c83e"></a>

<a id="canonical-ed2ee003575584025610817dbdd6e7c4751edf577b503f0cae808c8e4b2e12f9"></a>

## digest property — primary.default_rr_set_group.cds_record.values.sha1_digest / 02d90d089aa0 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-dd6ac59c36f4f11ad2a16d4cf2602199bd8e38b7fe8997b71e80a74db16d9b60"></a>

## Next pages — primary.default_rr_set_group.cds_record.values.sha1_digest / 02d90d089aa0 / 5

- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-a142e4b9d79d18af2d2f44aca68f97521046c0d578ef61f2d2f12c23f5b5c452)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-ca7a3bf29db338ce49850f066809c0d4dcd5327d54809683968c80a04f907919"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5276abd326252aeb830cfde9b4efbd18d4194d589da095daa5ca6718384883f"></a>

## primary.default_rr_set_group.cds_record.values.sha256_digest — primary.default_rr_set_group.cds_record.values.sha256_digest / 146347ea6e5a / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-794f910af9f92715f1b01c03d4bf830e8f08186fae3d28c3d8e7b0bc155cda37)
- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-a142e4b9d79d18af2d2f44aca68f97521046c0d578ef61f2d2f12c23f5b5c452)
- primary.default_rr_set_group.cds_record.values.sha256_digest

<a id="canonical-ffd71854cdd8cdc872c50924f84a8f61a6ebb8584d792ac16052fa7e8bba8845"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha256_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-80a2d5c3383f060670342863d0986f44e2d08becfb6b52ba956ad7fac23d9245"></a>

## Direct properties — primary.default_rr_set_group.cds_record.values.sha256_digest / 146347ea6e5a / 3

<a id="canonical-39146718da7d7d11d747999f92a09bc4e8b5c7f1582cc2db2851238e6647f8ea"></a>

<a id="canonical-71cb6d43a6de2e12027fea30adb84003dfb7cb64e08a56ebf44790517340d6fe"></a>

## digest property — primary.default_rr_set_group.cds_record.values.sha256_digest / 146347ea6e5a / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-c0134e24123ef2823e59af7728662fac678cb3e48a89d58cba6a1ff34d6b3e50"></a>

## Next pages — primary.default_rr_set_group.cds_record.values.sha256_digest / 146347ea6e5a / 5

- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-a142e4b9d79d18af2d2f44aca68f97521046c0d578ef61f2d2f12c23f5b5c452)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-a20c4875d92ad2ae11a0aaadcb9673680530ce10f643a27c77e6e3729c5fcbb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6eb5de40daeb82e99757c047617483ef25e0a2f7be1a9c8aec0500c4078b0b61"></a>

## primary.default_rr_set_group.cds_record.values.sha384_digest — primary.default_rr_set_group.cds_record.values.sha384_digest / 7f4b1f360e9b / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-794f910af9f92715f1b01c03d4bf830e8f08186fae3d28c3d8e7b0bc155cda37)
- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-a142e4b9d79d18af2d2f44aca68f97521046c0d578ef61f2d2f12c23f5b5c452)
- primary.default_rr_set_group.cds_record.values.sha384_digest

<a id="canonical-c78be4f701a4bf07e83705a01b3d467dc5fd8d6c3e4b7dfc7d31943ebcd26ca4"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha384 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha384_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-5142d63a66a688aa088d7c5179d22abc117453cd5c26d093f80b577725b08d8e"></a>

## Direct properties — primary.default_rr_set_group.cds_record.values.sha384_digest / 7f4b1f360e9b / 3

<a id="canonical-34e231455619605086e5093411b76c2b08d43da3a12c2504635807a4acef7a76"></a>

<a id="canonical-9062aa941adde4967bb3ef22b868d46b041c57a4e0748db22cd62012fcd36947"></a>

## digest property — primary.default_rr_set_group.cds_record.values.sha384_digest / 7f4b1f360e9b / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(96, 96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-89c9c5adbd89beee501218c709802c87b41093b633feae88483189244451505d"></a>

## Next pages — primary.default_rr_set_group.cds_record.values.sha384_digest / 7f4b1f360e9b / 5

- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-a142e4b9d79d18af2d2f44aca68f97521046c0d578ef61f2d2f12c23f5b5c452)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-c3e649004dc62fa19d277fa4563448bf796df6c0bce76da41fe6872c5748c611"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dac2cd235e2549db70816bba45e0182d29a58ced6eba94aee14a019a4bf598ad"></a>

## primary.default_rr_set_group.cert_record — primary.default_rr_set_group.cert_record / 46d1e2a4adc9 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.cert_record

<a id="canonical-35b54fc48af4a141a054910a2e7fbadf2d60aa44dc7f9d501f0107f41557e0e3"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cert record.

Upstream description:

DNS CERT Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
cert_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-896466d59ecf8be21a978ab01051673634bca43311b9862b947933b65f727cda"></a>

## Direct properties — primary.default_rr_set_group.cert_record / 46d1e2a4adc9 / 3

<a id="canonical-ac59a80ab2be105145a9fed01f1451008dcbb5c7b5487df5e76efbf08363e288"></a>

<a id="canonical-a170198842e1ddd29f01ff225ba22b5c51005248d460450e55f9aa7e729b175b"></a>

## name property — primary.default_rr_set_group.cert_record / 46d1e2a4adc9 / 4

Type: `"string"`. Optional.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-001.md#canonical-a4c61301e879017bb039b72d3c278834d18420b0ecd34d92113667fdb70f70dd): complete subsection reference.

<a id="canonical-afc2a19f9a576e238cb762bf9594d568dc52eb45b000b36bbf903ae9c07f38c6"></a>

## Next pages — primary.default_rr_set_group.cert_record / 46d1e2a4adc9 / 5

- [primary.default_rr_set_group.cert_record.values](resources--dns_zone--reference--group-001.md#canonical-a4c61301e879017bb039b72d3c278834d18420b0ecd34d92113667fdb70f70dd)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-a4c61301e879017bb039b72d3c278834d18420b0ecd34d92113667fdb70f70dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b703048ac935373a576ef5b58b55a5516ace9cd1ccb570fb0e8257ce3557223e"></a>

## primary.default_rr_set_group.cert_record.values — primary.default_rr_set_group.cert_record.values / 0c32ba3e6c1f / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.cert_record](resources--dns_zone--reference--group-001.md#canonical-c3e649004dc62fa19d277fa4563448bf796df6c0bce76da41fe6872c5748c611)
- primary.default_rr_set_group.cert_record.values

<a id="canonical-60d828a22f1921316eeab51f0d1fa3c99e5ca29144ef0557aa559b8a910c93cd"></a>

Type: `"object"`. list nested block, Optional.

CERT Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("cert_key_tag",
    "certificate")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-eae263efe7c116dd8cb1f1b8f81e3038211a832d42267aa391b342164ca4109f"></a>

## Direct properties — primary.default_rr_set_group.cert_record.values / 0c32ba3e6c1f / 3

<a id="canonical-9acbe02274496d3cac6ed9d945a8fe145e17b094cacd66d875328dfa2c5a82f1"></a>

<a id="canonical-8add8b3c64b7571e1bb84e66a278e1799405f8778977bfbff2aeaba41416324b"></a>

## algorithm property — primary.default_rr_set_group.cert_record.values / 0c32ba3e6c1f / 4

Type: `"string"`. Optional.

\[Enum: RESERVEDALGORITHM|RSAMD5|DH|DSASHA1|ECC|RSASHA1ALGORITHM|INDIRECT|PRIVATEDNS|PRIVATEOID\]
CERT algorithm value must be compatible with the specified algorithm. - RESERVEDALGORITHM:
RESERVEDALGORITHM - RSAMD5: RSAMD5 - DH: DH - DSASHA1: DSASHA1 - ECC: ECC - RSASHA1ALGORITHM:
RSA-SHA1 - INDIRECT: INDIRECT - PRIVATEDNS: PRIVATEDNS - PRIVATEOID: PRIVATEOID. Possible values are
\`RESERVEDALGORITHM\`, \`RSAMD5\`, \`DH\`, \`DSASHA1\`, \`ECC\`, \`RSASHA1ALGORITHM\`, \`INDIRECT\`,
\`PRIVATEDNS\`, \`PRIVATEOID\`. Defaults to \`RESERVEDALGORITHM\`.

Upstream description:

CERT algorithm value must be compatible with the specified algorithm.

&#8203;- RESERVEDALGORITHM: RESERVEDALGORITHM

&#8203;- RSAMD5: RSAMD5

&#8203;- DH: DH

&#8203;- DSASHA1: DSASHA1

&#8203;- ECC: ECC

&#8203;- RSASHA1ALGORITHM: RSA-SHA1

&#8203;- INDIRECT: INDIRECT

&#8203;- PRIVATEDNS: PRIVATEDNS

&#8203;- PRIVATEOID: PRIVATEOID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "RESERVEDALGORITHM",
  "enum": [
    "RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-047427b75fe967c06976aea655ecb001c6e05ae7b60ed06921f4c3ad4023eaab"></a>

<a id="canonical-7b76db56df143904f8e6ee80abac4b5a5878bcfbf8fbf018e563a40a03a84300"></a>

## cert_key_tag property — primary.default_rr_set_group.cert_record.values / 0c32ba3e6c1f / 5

Type: `"number"`. Optional.

Key Tag. Tag for categorization and filtering

Upstream description:

Tag for categorization and filtering

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-ca91483b5c9aa839b4ff31fb39951fd87223c33f9318f956cc87218ad72a247e"></a>

<a id="canonical-e0071914f87f7e52bb2666a12051650ed3255e84412c0edfa41f161fe3c8c016"></a>

## cert_type property — primary.default_rr_set_group.cert_record.values / 0c32ba3e6c1f / 6

Type: `"string"`. Optional.

\[Enum: INVALIDCERTTYPE|PKIX|SPKI|PGP|IPKIX|ISPKI|IPGP|ACPKIX|IACPKIX|URI\_|OID\] CERT type value
must be compatible with the specified types. - INVALIDCERTTYPE: INVALIDCERTTYPE - PKIX: PKIX - SPKI:
SPKI - PGP: PGP - IPKIX: IPKIX - ISPKI: ISPKI - IPGP: IPGP - ACPKIX: ACPKIX - IACPKIX: IACPKIX -
URI\_: URI - OID: OID. Possible values are \`INVALIDCERTTYPE\`, \`PKIX\`, \`SPKI\`, \`PGP\`,
\`IPKIX\`, \`ISPKI\`, \`IPGP\`, \`ACPKIX\`, \`IACPKIX\`, \`URI\_\`, \`OID\`. Defaults to
\`INVALIDCERTTYPE\`.

Upstream description:

CERT type value must be compatible with the specified types.

&#8203;- INVALIDCERTTYPE: INVALIDCERTTYPE

&#8203;- PKIX: PKIX

&#8203;- SPKI: SPKI

&#8203;- PGP: PGP

&#8203;- IPKIX: IPKIX

&#8203;- ISPKI: ISPKI

&#8203;- IPGP: IPGP

&#8203;- ACPKIX: ACPKIX

&#8203;- IACPKIX: IACPKIX

&#8203;- URI\_: URI

&#8203;- OID: OID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALIDCERTTYPE",
  "enum": [
    "INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5c937c0b8d067565d6fb3e46e29a20d7995ada62bcba17b1e096f20e26e35c31"></a>

<a id="canonical-e09a7177b39fbe911571c0e0df507448c9cc311dd92cec4263b93b26088ac1cc"></a>

## certificate property — primary.default_rr_set_group.cert_record.values / 0c32ba3e6c1f / 7

Type: `"string"`. Optional.

Certificate. Certificate in base 64 format.

Upstream description:

Certificate in base 64 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-891ea645b865d0498401182a53eecef44b8aed7a4be6d953dfe4290a820df256"></a>

## Next pages — primary.default_rr_set_group.cert_record.values / 0c32ba3e6c1f / 8

- [primary.default_rr_set_group.cert_record](resources--dns_zone--reference--group-001.md#canonical-c3e649004dc62fa19d277fa4563448bf796df6c0bce76da41fe6872c5748c611)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-8ef00259f52c959e585f74aed99b7ff0449ac0da39979cbc38adf23ab232805c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95ba5559397743f5e0d7cd6c749fefd894df1f016073152a5a9828111d7e69e1"></a>

## primary.default_rr_set_group.cname_record — primary.default_rr_set_group.cname_record / 74c00877848d / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.cname_record

<a id="canonical-e71d65a7afabb6f07076babaedb18c3f43be759348f4857518a48e5be3df5b60"></a>

Type: `"object"`. single nested block, Optional.

DNSCNAMEResourceRecord.

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
cname_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-6f1d9e3fe953dd266fadb77358331bc8666e0b86e8d21460ef08ada30076b611"></a>

## Direct properties — primary.default_rr_set_group.cname_record / 74c00877848d / 3

<a id="canonical-1f2b78cfa39de47786e27406dea1bf2a80d9aa441867106e165a6c9ba29f4888"></a>

<a id="canonical-8f0d2d7d46656cb30217b0d999b66bc444e18b28d313f8d201d0b0b991d907c5"></a>

## name property — primary.default_rr_set_group.cname_record / 74c00877848d / 4

Type: `"string"`. Optional.

CName Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-0d53f753398168cc5dff897ed0cd633037f9ab57235a35922b46f7c4c964349d"></a>

<a id="canonical-18f627e183406076a5a89eb5b4447ceb3c3b28ca9d825dbef103a9cde718add5"></a>

## value property — primary.default_rr_set_group.cname_record / 74c00877848d / 5

Type: `"string"`. Optional.

Domain. Configuration parameter for value

Upstream description:

Configuration parameter for value

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-d2f584a975018c23681cb67cde63ab3093fb6d2cb3ea6fd5ce0bc42143cf8557"></a>

## Next pages — primary.default_rr_set_group.cname_record / 74c00877848d / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-a95a01f6aa08e943ea1267e9b696669ccf3c5fcf7acfd21343f4a70bf68b5673"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
