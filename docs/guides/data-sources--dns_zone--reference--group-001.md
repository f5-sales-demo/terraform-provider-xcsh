---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a061fac3782c91c71a49e4c86b7f681c8eb0a93601e7a086e6ee3826d7261b80"></a>

## Property reference — Property reference / d4f8b7660ed9 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- Property reference

<a id="canonical-9241c57b6c6d871a3fbb04012b6094ec0c29f32061fd18a420c41f1cfc9fae6f"></a>

## Direct properties — Property reference / d4f8b7660ed9 / 3

<a id="canonical-e9dd2bcbc8c79d2015c244c0a296c3d339e1460343bb688fce2bb16139b9f9e9"></a>

<a id="canonical-6f8ff9bd6114747d6318c21b5fdc4ea62ab3e1f41256455b6da2a9f92e9c530f"></a>

## annotations property — Property reference / d4f8b7660ed9 / 4

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

<a id="canonical-fddd85c262f3bfc757008225c04e43fcf09e9d2f9c60988442bc08f0fa03fcbe"></a>

<a id="canonical-73ed1086b70b9cea179b20467292fe31e3cf6b0a244a1534b817647e0b65bb6e"></a>

## description property — Property reference / d4f8b7660ed9 / 5

Type: `"string"`. Computed.

Description of the DNSZone.

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

<a id="canonical-4b4498f5cb72d11c7af007bc8bb356d0866c25b48d3614b6ed110e601229c28d"></a>

<a id="canonical-9ff3efbe08bee83dc240ff42d0bc51f6e39723f5dfae901b1c1a24a6d212e02b"></a>

## id property — Property reference / d4f8b7660ed9 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-14408e111c2748ac8723ea02c044a95a358e523ac0e4892ff370e9e569c1df7f"></a>

<a id="canonical-caa92427631bff9fb327e7ecde0df8a8b1306dfa7e392828be943c8843b077d2"></a>

## labels property — Property reference / d4f8b7660ed9 / 7

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

<a id="canonical-e1a7635cb461e07bfda0fbf9e86d7d78dc0cbe195397bb6082207def33bebe2c"></a>

<a id="canonical-ee9ca8691a49a92ac792ab7db449c8edd72ee3e8aaeba5eca2d3a954f9d80c68"></a>

## name property — Property reference / d4f8b7660ed9 / 8

Type: `"string"`. Required.

Name of the DNSZone.

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

<a id="canonical-a84e1e3bb3e1a3dcf4b8cf21370681be271d4c524ac83c7b2d479f7405c0c4d7"></a>

<a id="canonical-9c1e09221227083b9af414546faf1885101d789dfb80ed27cb03c5011d27b9f1"></a>

## namespace property — Property reference / d4f8b7660ed9 / 9

Type: `"string"`. Optional, Computed.

Namespace where the DNSZone exists.

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

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd): complete subsection reference.

- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-1416f5ce025b5fd7645bf56ccfba5f70b02bac3f9d43100eb835a8e6d3d840cb): complete subsection reference.

<a id="canonical-c9d823e8634951bd28b10745f1df691080ddfdc9ecdcc8af42b22d608cd1247e"></a>

## All schema paths — Property reference / d4f8b7660ed9 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_zone--reference--group-001.md#canonical-e9dd2bcbc8c79d2015c244c0a296c3d339e1460343bb688fce2bb16139b9f9e9) |
| `description` | [description](data-sources--dns_zone--reference--group-001.md#canonical-fddd85c262f3bfc757008225c04e43fcf09e9d2f9c60988442bc08f0fa03fcbe) |
| `id` | [id](data-sources--dns_zone--reference--group-001.md#canonical-4b4498f5cb72d11c7af007bc8bb356d0866c25b48d3614b6ed110e601229c28d) |
| `labels` | [labels](data-sources--dns_zone--reference--group-001.md#canonical-14408e111c2748ac8723ea02c044a95a358e523ac0e4892ff370e9e569c1df7f) |
| `name` | [name](data-sources--dns_zone--reference--group-001.md#canonical-e1a7635cb461e07bfda0fbf9e86d7d78dc0cbe195397bb6082207def33bebe2c) |
| `namespace` | [namespace](data-sources--dns_zone--reference--group-001.md#canonical-a84e1e3bb3e1a3dcf4b8cf21370681be271d4c524ac83c7b2d479f7405c0c4d7) |
| `primary` | [primary](data-sources--dns_zone--reference--group-001.md#canonical-991194abf4edfd16575548ebc084e3a041bedb94121f2deceb51290542638a3d) |
| `primary.allow_http_lb_managed_records` | [primary.allow_http_lb_managed_records](data-sources--dns_zone--reference--group-001.md#canonical-05875f7e000466835a89dcf6a148db5b48990086e692d272f7939180bf7164f6) |
| `primary.default_rr_set_group` | [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-5748f65740c50bd17bc2bcdf7d61709eee6aa315620872aebee05c3708b35884) |
| `primary.default_rr_set_group.a_record` | [primary.default_rr_set_group.a_record](data-sources--dns_zone--reference--group-001.md#canonical-a2300cf8a33cb09199c979b85d5b0adf5a4eb35c0305e7da0e733c6f2399ea49) |
| `primary.default_rr_set_group.a_record.name` | [primary.default_rr_set_group.a_record.name](data-sources--dns_zone--reference--group-001.md#canonical-a7bc9e04df108de7b707747192087c75a25dfef9458d6c74f88258dfc3fce045) |
| `primary.default_rr_set_group.a_record.values` | [primary.default_rr_set_group.a_record.values](data-sources--dns_zone--reference--group-001.md#canonical-64d6448fdef132fa77bd9d55aa0a5cd75d20ad700e761982c55767ee12c2c6e9) |
| `primary.default_rr_set_group.aaaa_record` | [primary.default_rr_set_group.aaaa_record](data-sources--dns_zone--reference--group-001.md#canonical-373d92e2365372276e53adbddf0365f445a7016cb8cc4ba8a8a3feba7872a5dd) |
| `primary.default_rr_set_group.aaaa_record.name` | [primary.default_rr_set_group.aaaa_record.name](data-sources--dns_zone--reference--group-001.md#canonical-d23c1177f69df50d69b8eb6d6d32e794fb46c7588e912d2338b77ef2ca615e98) |
| `primary.default_rr_set_group.aaaa_record.values` | [primary.default_rr_set_group.aaaa_record.values](data-sources--dns_zone--reference--group-001.md#canonical-b30b80e9d49030569604229f828dd63891c9b3a0b6cc028ee6ee3e29652c56bb) |
| `primary.default_rr_set_group.afsdb_record` | [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-de718cdbd5f577ea03a3a7ad69609fda1fdc6a17862764b256027bc5772b61d5) |
| `primary.default_rr_set_group.afsdb_record.name` | [primary.default_rr_set_group.afsdb_record.name](data-sources--dns_zone--reference--group-001.md#canonical-14c3d71b9ed9b47ff7cc3b5ce45359ff6a03f441bd63dc79e43bf2422b8cbb44) |
| `primary.default_rr_set_group.afsdb_record.values` | [primary.default_rr_set_group.afsdb_record.values](data-sources--dns_zone--reference--group-001.md#canonical-dc0bf29d2c0cb263cccf2fbe8485aef77576e249cb3208c5dae82fb5e721e534) |
| `primary.default_rr_set_group.afsdb_record.values.hostname` | [primary.default_rr_set_group.afsdb_record.values.hostname](data-sources--dns_zone--reference--group-001.md#canonical-423f6991bd603d66561b5226dd7ade0e9d0495d0c7fd06afbe4427467c95c513) |
| `primary.default_rr_set_group.afsdb_record.values.subtype` | [primary.default_rr_set_group.afsdb_record.values.subtype](data-sources--dns_zone--reference--group-001.md#canonical-79d6fab20e287aa16ae993a2fd348ca9ac7f657690f3dfa58c4619998a42cd68) |
| `primary.default_rr_set_group.alias_record` | [primary.default_rr_set_group.alias_record](data-sources--dns_zone--reference--group-001.md#canonical-dbcdbdad5885c364c0821e484bda01b18a52005c0478de41ede87bd48ad713f5) |
| `primary.default_rr_set_group.alias_record.value` | [primary.default_rr_set_group.alias_record.value](data-sources--dns_zone--reference--group-001.md#canonical-03ff40e52ec1b822f4df04e1813f2d8ec6cf412ad06bb333e2cda2a26ca1954f) |
| `primary.default_rr_set_group.caa_record` | [primary.default_rr_set_group.caa_record](data-sources--dns_zone--reference--group-001.md#canonical-95b2a44978900fc848cfd9da410e1c1033f0820aed35c6e846b39c3ab23d5b19) |
| `primary.default_rr_set_group.caa_record.name` | [primary.default_rr_set_group.caa_record.name](data-sources--dns_zone--reference--group-001.md#canonical-c00fe2f75bd0bcb6de41bfd4bc698f8593a9af39e52ba41d1a3d4e7d8646095c) |
| `primary.default_rr_set_group.caa_record.values` | [primary.default_rr_set_group.caa_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1eb96f1ac8cca66f9b5517d9b213b52c0c01a4c858053edbc4e35e6b859f0643) |
| `primary.default_rr_set_group.caa_record.values.flags` | [primary.default_rr_set_group.caa_record.values.flags](data-sources--dns_zone--reference--group-001.md#canonical-bbd76a8107630d7afdf91dcac45c4726e344a037fe940dc4657ffade607d5755) |
| `primary.default_rr_set_group.caa_record.values.tag` | [primary.default_rr_set_group.caa_record.values.tag](data-sources--dns_zone--reference--group-001.md#canonical-f49b5ef1937c2bd435a0ac6cc6f650f25811285367ecef5b5d72a957dc2f7bdc) |
| `primary.default_rr_set_group.caa_record.values.value` | [primary.default_rr_set_group.caa_record.values.value](data-sources--dns_zone--reference--group-001.md#canonical-f0a421166aa6c2881fc5ed0f3a46845747030d99221444526642aa3b7148f51e) |
| `primary.default_rr_set_group.cds_record` | [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-a472e0a669fefebecdab8c5da4362e3e73a7910089f88605db2e7a152b10e7ed) |
| `primary.default_rr_set_group.cds_record.name` | [primary.default_rr_set_group.cds_record.name](data-sources--dns_zone--reference--group-001.md#canonical-61bf656d76ac7455c0beaa1c23f0a76095783ee0bd5303eb8d14d524a3354a01) |
| `primary.default_rr_set_group.cds_record.values` | [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-d3d49d288965a11d71cd95f35101d33caaa799741e8c8f8f32bf3da7d7e52537) |
| `primary.default_rr_set_group.cds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.cds_record.values.ds_key_algorithm](data-sources--dns_zone--reference--group-001.md#canonical-ef9dbe8465bb62417924ccb69c13575b6c68f9518ed1f8a1d9fbb68e108a8a75) |
| `primary.default_rr_set_group.cds_record.values.key_tag` | [primary.default_rr_set_group.cds_record.values.key_tag](data-sources--dns_zone--reference--group-001.md#canonical-470f40f92f3d77628c6949e68e1995cea824138172bcb7890bb233e618573248) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest](data-sources--dns_zone--reference--group-001.md#canonical-932ce44e6439ffca23e08d6da83bcd62c25726cd1a0a76fb12479bf4e4b8d86c) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest.digest](data-sources--dns_zone--reference--group-001.md#canonical-5875a5a41cbdfa4b62f740d1e0ef396aae8af46c461da69596453fea81568956) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest](data-sources--dns_zone--reference--group-001.md#canonical-4ee507782954bdb72e61fcf62c04b6c981180dd977bc323a56a911eeee41d216) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest.digest](data-sources--dns_zone--reference--group-001.md#canonical-3c525a1316bc6cb938031144311f589dd2733055002c2b374e1ea655e13af88a) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest](data-sources--dns_zone--reference--group-001.md#canonical-5519518722cabe2021c0f6af3fbf04850ba200b22a872d9dc177b9efc43f612e) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest.digest](data-sources--dns_zone--reference--group-001.md#canonical-32e43cc11a18543059efa3337113743891a1c2dfdea74e5085fec72fa2358c93) |
| `primary.default_rr_set_group.cert_record` | [primary.default_rr_set_group.cert_record](data-sources--dns_zone--reference--group-001.md#canonical-bcb066de69b36d0f9f55f58fdd2ec7682d9314cfa5d7d47dfea798bdc35285f1) |
| `primary.default_rr_set_group.cert_record.name` | [primary.default_rr_set_group.cert_record.name](data-sources--dns_zone--reference--group-001.md#canonical-f6a3fa4de87213668250e22f61b2987680473eb4f5f0db3515745d27f69215ca) |
| `primary.default_rr_set_group.cert_record.values` | [primary.default_rr_set_group.cert_record.values](data-sources--dns_zone--reference--group-001.md#canonical-b06c6a188f9b294b16adc441378975b81b28554f356aff00553754aee09ba73f) |
| `primary.default_rr_set_group.cert_record.values.algorithm` | [primary.default_rr_set_group.cert_record.values.algorithm](data-sources--dns_zone--reference--group-001.md#canonical-944dbdeadf1752dc7ceda15a10380b9f9a97564d6d62001ecdb31e44bbdc0996) |
| `primary.default_rr_set_group.cert_record.values.cert_key_tag` | [primary.default_rr_set_group.cert_record.values.cert_key_tag](data-sources--dns_zone--reference--group-001.md#canonical-a0c2c76d7f2c8b2f1f20ae4ccf4ef116402a5231119900a0444101dc501c01cd) |
| `primary.default_rr_set_group.cert_record.values.cert_type` | [primary.default_rr_set_group.cert_record.values.cert_type](data-sources--dns_zone--reference--group-001.md#canonical-42811c8d9e9af9117543fb87bf81f96eebcc3c8fb356753a9af29aeb6aa26d45) |
| `primary.default_rr_set_group.cert_record.values.certificate` | [primary.default_rr_set_group.cert_record.values.certificate](data-sources--dns_zone--reference--group-001.md#canonical-dae5ffff77d4be74941bb46e78137f1ae39452fd51058f89145b56b8f5d12975) |
| `primary.default_rr_set_group.cname_record` | [primary.default_rr_set_group.cname_record](data-sources--dns_zone--reference--group-001.md#canonical-99c55d06815c7205b12930a31025f9cd8ca0445f61e6b82d81e531eb8d772219) |
| `primary.default_rr_set_group.cname_record.name` | [primary.default_rr_set_group.cname_record.name](data-sources--dns_zone--reference--group-001.md#canonical-c85c38aafe6406022f0e1ec9cd70304e372a6ba55f51e9a21bcc5937ddd63ac4) |
| `primary.default_rr_set_group.cname_record.value` | [primary.default_rr_set_group.cname_record.value](data-sources--dns_zone--reference--group-001.md#canonical-748c70ce8f42c7624739b3872b69bfd873bb0672551f68b6c9cba530e60d4d90) |
| `primary.default_rr_set_group.description_spec` | [primary.default_rr_set_group.description_spec](data-sources--dns_zone--reference--group-001.md#canonical-17d879209731d27bde4e0df767c1806e179bc7b05be3585ec090404c5cbd6900) |
| `primary.default_rr_set_group.ds_record` | [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-658dc75819079df5add24a0c1c69b129caffda803dc0d041a833ae783815a7dd) |
| `primary.default_rr_set_group.ds_record.name` | [primary.default_rr_set_group.ds_record.name](data-sources--dns_zone--reference--group-001.md#canonical-a45d0972cb1f990aff0214b71e231c3ee7f323ac0edca35bbf6c9f6c28facb79) |
| `primary.default_rr_set_group.ds_record.values` | [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-80485afca70502b050a6561537dacab323366da90417a5269d140391d3c66994) |
| `primary.default_rr_set_group.ds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.ds_record.values.ds_key_algorithm](data-sources--dns_zone--reference--group-001.md#canonical-1a9bf10795c490f393b3b3f0ffee8f3e99efe16e35decd808e98937a269a8da7) |
| `primary.default_rr_set_group.ds_record.values.key_tag` | [primary.default_rr_set_group.ds_record.values.key_tag](data-sources--dns_zone--reference--group-001.md#canonical-b28f44e5ab4287ae4e21a6faf57c28d9360f7538ab54c02273c07af72c77ea97) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest](data-sources--dns_zone--reference--group-001.md#canonical-6b58cd1b299321dd72348ac32e4a41d836c4cbecc578256c855b9a7d8db10203) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest.digest](data-sources--dns_zone--reference--group-001.md#canonical-8b653e94d37b371abbb6712328b5c412d166b37bec2801e6ec9750cb1f8df487) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-654ea7690ee98cc491c81e4eed22052df1995918522198f3e834d26bca9e704d) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-4dafb0b395ca3ef206aed65a460c68cf58e6b70139e720f6a46bf8808010165c) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-63aab37fa7dee6df1e2e69f89a90556dddad581d5b06734ed4961b2fc4c7e1e1) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-9081e150c9ba17b51434b56c2cd5c5b94ad95f1140d7c132059cc2d318debdf6) |
| `primary.default_rr_set_group.eui48_record` | [primary.default_rr_set_group.eui48_record](data-sources--dns_zone--reference--group-002.md#canonical-510e55af08a649aeb199a629b0ba9db14c75629ccaf79c99fc188a72437297e8) |
| `primary.default_rr_set_group.eui48_record.name` | [primary.default_rr_set_group.eui48_record.name](data-sources--dns_zone--reference--group-002.md#canonical-c74b081322eca48f60f206ae6402f0eded2519426e05ddfb8d80494ca758a82c) |
| `primary.default_rr_set_group.eui48_record.value` | [primary.default_rr_set_group.eui48_record.value](data-sources--dns_zone--reference--group-002.md#canonical-0197220591dbfd0a64529cf640880c999cc2267af43393f69041063dd1fbf081) |
| `primary.default_rr_set_group.eui64_record` | [primary.default_rr_set_group.eui64_record](data-sources--dns_zone--reference--group-002.md#canonical-c0f1cdfefed198669ca4168470c520c647eab55ebd61a1e1b66fffa24659c4d8) |
| `primary.default_rr_set_group.eui64_record.name` | [primary.default_rr_set_group.eui64_record.name](data-sources--dns_zone--reference--group-002.md#canonical-54d1d38d99859d69e84f6d4e8bd3f94de95a06a1a6cbc2f97eb5ac5836d1ca02) |
| `primary.default_rr_set_group.eui64_record.value` | [primary.default_rr_set_group.eui64_record.value](data-sources--dns_zone--reference--group-002.md#canonical-aa93a64babaa6873c6e09c529d1eb3c5dd2c183858bceae849ce1672eb64d971) |
| `primary.default_rr_set_group.lb_record` | [primary.default_rr_set_group.lb_record](data-sources--dns_zone--reference--group-002.md#canonical-f91e65320b4b4381fa10b3eaac701378c1bd08564853254068796e84e3380a22) |
| `primary.default_rr_set_group.lb_record.name` | [primary.default_rr_set_group.lb_record.name](data-sources--dns_zone--reference--group-002.md#canonical-27aabec96ea0ba69970a3773cd48c9f72d7bf3daa0910b8d83eafccb8400addf) |
| `primary.default_rr_set_group.lb_record.value` | [primary.default_rr_set_group.lb_record.value](data-sources--dns_zone--reference--group-002.md#canonical-de8c7c93762bffe2dce11df4f3731fda9783de0b393675ae3a93c24719fb898d) |
| `primary.default_rr_set_group.lb_record.value.name` | [primary.default_rr_set_group.lb_record.value.name](data-sources--dns_zone--reference--group-002.md#canonical-5dab3b4e57ea026ddf2d276d60c2b7f8f52ba2ff88b69203052c0106e9e86674) |
| `primary.default_rr_set_group.lb_record.value.namespace` | [primary.default_rr_set_group.lb_record.value.namespace](data-sources--dns_zone--reference--group-002.md#canonical-662a50796055482ab34ff0ac4737b0121aaad7b17f95c9591d7b9ac1116be626) |
| `primary.default_rr_set_group.lb_record.value.tenant` | [primary.default_rr_set_group.lb_record.value.tenant](data-sources--dns_zone--reference--group-002.md#canonical-ec4bfe0f9400cc1230e0aa89c61dba5bb00ae132566a8be7126065a8394d866f) |
| `primary.default_rr_set_group.loc_record` | [primary.default_rr_set_group.loc_record](data-sources--dns_zone--reference--group-002.md#canonical-f26a5d6ad3d522d0e84fb4219b1155bbdb5f12fe2c0a78473cb480b5b9c05ab6) |
| `primary.default_rr_set_group.loc_record.name` | [primary.default_rr_set_group.loc_record.name](data-sources--dns_zone--reference--group-002.md#canonical-53b0121cfd2f282b16a030ab404c66732aa1571b88fe7f73189d294e37d5e26f) |
| `primary.default_rr_set_group.loc_record.values` | [primary.default_rr_set_group.loc_record.values](data-sources--dns_zone--reference--group-002.md#canonical-d7824b5a29aa37b126335cf43c7df72fa3c93c819e6fa1bd9348470647e0a998) |
| `primary.default_rr_set_group.loc_record.values.altitude` | [primary.default_rr_set_group.loc_record.values.altitude](data-sources--dns_zone--reference--group-002.md#canonical-f58209df899332759fb4a1dc3d6275fc8a8c4d95713b2e097726cb10591ceb64) |
| `primary.default_rr_set_group.loc_record.values.horizontal_precision` | [primary.default_rr_set_group.loc_record.values.horizontal_precision](data-sources--dns_zone--reference--group-002.md#canonical-b0621942b4e793169cb300fe2d2448a2da7387adc01c76812a050bbdc1535e27) |
| `primary.default_rr_set_group.loc_record.values.latitude_degree` | [primary.default_rr_set_group.loc_record.values.latitude_degree](data-sources--dns_zone--reference--group-002.md#canonical-d4cd70d24dcbd3393b5377839ddd2b08291c6c15d5fdb478040464e62c314674) |
| `primary.default_rr_set_group.loc_record.values.latitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.latitude_hemisphere](data-sources--dns_zone--reference--group-002.md#canonical-0db02b40642b518b105d14fd51cea1aaee59c3c13870a9365387102812fd0414) |
| `primary.default_rr_set_group.loc_record.values.latitude_minute` | [primary.default_rr_set_group.loc_record.values.latitude_minute](data-sources--dns_zone--reference--group-002.md#canonical-b251c2ddd9f00f7143fa75d1dcb7a3f1156b1d4b610508cf4bdd0d65ddbf2d47) |
| `primary.default_rr_set_group.loc_record.values.latitude_second` | [primary.default_rr_set_group.loc_record.values.latitude_second](data-sources--dns_zone--reference--group-002.md#canonical-752a3ccfd6d90fb1f41d136ce2d55a9722f295ddf8d9bd48afd829e9d0ce6131) |
| `primary.default_rr_set_group.loc_record.values.location_diameter` | [primary.default_rr_set_group.loc_record.values.location_diameter](data-sources--dns_zone--reference--group-002.md#canonical-7e21432814efe3da25a539df9e2e4e3ee0739cacd77e1306219fe8eb63ad78bc) |
| `primary.default_rr_set_group.loc_record.values.longitude_degree` | [primary.default_rr_set_group.loc_record.values.longitude_degree](data-sources--dns_zone--reference--group-002.md#canonical-d8b9b3fd307d57e7d597b8e7f492b3505d5725eb256340e6cedb96c602c7337e) |
| `primary.default_rr_set_group.loc_record.values.longitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.longitude_hemisphere](data-sources--dns_zone--reference--group-002.md#canonical-907acf9a160ca7d155adf5fe57fe334c05b702876451bb03f9980a83e847a97b) |
| `primary.default_rr_set_group.loc_record.values.longitude_minute` | [primary.default_rr_set_group.loc_record.values.longitude_minute](data-sources--dns_zone--reference--group-002.md#canonical-8b6dfe91e887e332abc022a0651aeabfe7b28fcb11707da5dedba2ea5fc33f90) |
| `primary.default_rr_set_group.loc_record.values.longitude_second` | [primary.default_rr_set_group.loc_record.values.longitude_second](data-sources--dns_zone--reference--group-002.md#canonical-1cf24ef1b2ca3382c7b196eb860abb57c3d954555f19b05b41a8a9e0ef597e9b) |
| `primary.default_rr_set_group.loc_record.values.vertical_precision` | [primary.default_rr_set_group.loc_record.values.vertical_precision](data-sources--dns_zone--reference--group-002.md#canonical-953ed9af85851c83a90ec699477753a201b05463d7448473675a8d8cfd1900f4) |
| `primary.default_rr_set_group.mx_record` | [primary.default_rr_set_group.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-4ed508b2d3ed35877f102ee7827c5cc62ca029d4439cee685ce931edf38d5c59) |
| `primary.default_rr_set_group.mx_record.name` | [primary.default_rr_set_group.mx_record.name](data-sources--dns_zone--reference--group-002.md#canonical-3f5b1ca0de9b6a9c0aefdd04b306d8ba37f9dbd5a2ea6e0106abc63c88108e69) |
| `primary.default_rr_set_group.mx_record.values` | [primary.default_rr_set_group.mx_record.values](data-sources--dns_zone--reference--group-002.md#canonical-efbbb9c8523e5b4579896e983fe2c45938c932167dc063660a13874ac189c72f) |
| `primary.default_rr_set_group.mx_record.values.domain` | [primary.default_rr_set_group.mx_record.values.domain](data-sources--dns_zone--reference--group-002.md#canonical-e6bd310996f3c449f55db5b8e45987ed5bd2861671fdf08a97d43bc8636950bb) |
| `primary.default_rr_set_group.mx_record.values.priority` | [primary.default_rr_set_group.mx_record.values.priority](data-sources--dns_zone--reference--group-002.md#canonical-fa98b2319210b46b5847070ac0d532ece87809e6e4fe7fa5d65fb845feaef3a2) |
| `primary.default_rr_set_group.naptr_record` | [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-5d0ea0ba3dd17bad18e52ff0808d262d324c9ea8ee2eb94fbc2d9967b20fd28f) |
| `primary.default_rr_set_group.naptr_record.name` | [primary.default_rr_set_group.naptr_record.name](data-sources--dns_zone--reference--group-002.md#canonical-7420ce7f114f5f72b31bdd14c39ff37f47231e517ada5ed3c7d1cec2c77ae6c2) |
| `primary.default_rr_set_group.naptr_record.values` | [primary.default_rr_set_group.naptr_record.values](data-sources--dns_zone--reference--group-002.md#canonical-b880a18f4f3ce509128b3264cc366709307f04b6a3413f1fb019a3a6502f6d48) |
| `primary.default_rr_set_group.naptr_record.values.flags` | [primary.default_rr_set_group.naptr_record.values.flags](data-sources--dns_zone--reference--group-002.md#canonical-e8e0a0096f3802ee09f0dd51439085b415ebec22bc4688651d50b3f5b2d11fee) |
| `primary.default_rr_set_group.naptr_record.values.order` | [primary.default_rr_set_group.naptr_record.values.order](data-sources--dns_zone--reference--group-002.md#canonical-30359ee2cb5ae61f1caaccfe320338c4045df18a866131b9bb42943fa22e6251) |
| `primary.default_rr_set_group.naptr_record.values.preference` | [primary.default_rr_set_group.naptr_record.values.preference](data-sources--dns_zone--reference--group-002.md#canonical-85263eacad1fec5dde46c2f4cb03693c7f33258da46e3a9f15881e54b2c95917) |
| `primary.default_rr_set_group.naptr_record.values.regexp` | [primary.default_rr_set_group.naptr_record.values.regexp](data-sources--dns_zone--reference--group-002.md#canonical-fc6632232979040382fe4a6b75a397f5fbef5fcf9a34691e6b48f965ce570f22) |
| `primary.default_rr_set_group.naptr_record.values.replacement` | [primary.default_rr_set_group.naptr_record.values.replacement](data-sources--dns_zone--reference--group-002.md#canonical-cf5c3e94a4c2d10003deb2662612846a405caf97a41e076c7605d7a33d4ce29a) |
| `primary.default_rr_set_group.naptr_record.values.service` | [primary.default_rr_set_group.naptr_record.values.service](data-sources--dns_zone--reference--group-002.md#canonical-53074d828f3d8df7f15274ba604d75591586b42f64143590fd1263d4955d3186) |
| `primary.default_rr_set_group.ns_record` | [primary.default_rr_set_group.ns_record](data-sources--dns_zone--reference--group-002.md#canonical-d8175dc81584695aa7f4041c6e47c1ec0c8c1fa080d40238f50857367a1441db) |
| `primary.default_rr_set_group.ns_record.name` | [primary.default_rr_set_group.ns_record.name](data-sources--dns_zone--reference--group-002.md#canonical-a7e062b7494f3e8f1537b896f6584ce3c5918638ac8d2815a529c895c4e24625) |
| `primary.default_rr_set_group.ns_record.values` | [primary.default_rr_set_group.ns_record.values](data-sources--dns_zone--reference--group-002.md#canonical-30623e12f3e7eafe3b405bc7ecd967c7d1feffad4f49293ade8ce40ee836538b) |
| `primary.default_rr_set_group.ptr_record` | [primary.default_rr_set_group.ptr_record](data-sources--dns_zone--reference--group-002.md#canonical-f315e781da4e7643f5968da45534700de6e050fdfb1881c81853089e46ebcbd3) |
| `primary.default_rr_set_group.ptr_record.name` | [primary.default_rr_set_group.ptr_record.name](data-sources--dns_zone--reference--group-002.md#canonical-4b0ad3235582088478973c890faa8570458eae895aa6b9f424bc0f37072bfce4) |
| `primary.default_rr_set_group.ptr_record.values` | [primary.default_rr_set_group.ptr_record.values](data-sources--dns_zone--reference--group-002.md#canonical-b7465227216c220b57117a69b1de54aee08b57f9fe26936a395178622058f3d4) |
| `primary.default_rr_set_group.srv_record` | [primary.default_rr_set_group.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-814b76554d181ac9783b10f0691a158f24b5dfc56b7d7510da06f9706f83e201) |
| `primary.default_rr_set_group.srv_record.name` | [primary.default_rr_set_group.srv_record.name](data-sources--dns_zone--reference--group-002.md#canonical-3b710e4d3a4cc37c3eb4bfac827dc2d7d59d0fefdf48c537d8e51bfb6bb20870) |
| `primary.default_rr_set_group.srv_record.values` | [primary.default_rr_set_group.srv_record.values](data-sources--dns_zone--reference--group-002.md#canonical-c5865dff1e67772d6b443955e7a63965b8c2d3564f98600cf545520ad76b160e) |
| `primary.default_rr_set_group.srv_record.values.port` | [primary.default_rr_set_group.srv_record.values.port](data-sources--dns_zone--reference--group-002.md#canonical-01fa21e199199a8a4852f46e75e76a150b28ab891c01541de93797e825e244f2) |
| `primary.default_rr_set_group.srv_record.values.priority` | [primary.default_rr_set_group.srv_record.values.priority](data-sources--dns_zone--reference--group-002.md#canonical-2367d760698b9569bd7730fcee549527c4b524c203f355f87c184788b8996685) |
| `primary.default_rr_set_group.srv_record.values.target` | [primary.default_rr_set_group.srv_record.values.target](data-sources--dns_zone--reference--group-002.md#canonical-1c9b0e84099a49fbf1524ea59daa5608b39a23ba3a214568b5f97d267388a223) |
| `primary.default_rr_set_group.srv_record.values.weight` | [primary.default_rr_set_group.srv_record.values.weight](data-sources--dns_zone--reference--group-002.md#canonical-6d6f490ad4e1333c4f9e59ac5260d93d9b231af9a22e07ea8b311d3bf2afe6f7) |
| `primary.default_rr_set_group.sshfp_record` | [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-871bdf8acf9f220f6b0b5f0412202c9c90df26d383ecc774ae9bcec5e63c8b26) |
| `primary.default_rr_set_group.sshfp_record.name` | [primary.default_rr_set_group.sshfp_record.name](data-sources--dns_zone--reference--group-002.md#canonical-4b69eafc81e713604cdae06f212dc8909ae82c40d2d3d23ff81debabcc27b62a) |
| `primary.default_rr_set_group.sshfp_record.values` | [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0237ff5a4e21c521d28af88d42e828238c7e32eac8d099cf8d5274649c1c5aad) |
| `primary.default_rr_set_group.sshfp_record.values.algorithm` | [primary.default_rr_set_group.sshfp_record.values.algorithm](data-sources--dns_zone--reference--group-002.md#canonical-3330a0c5c8b7ac4c3c5166b13c280ceda3fab596e1b38de14b89252841fc35cc) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-b12a92e147a6cd9037f0fdede420e6a2b98d797489f70f76046c4380a6c93942) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-a21af051dd1a6e18a631c9bdbdd28a3d01b2e4616969574b29870720c3d5ebaa) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-6193f8e7abb75aa521e3804bd429931cb2d9152431fc4f0f2c2ac5a6a4173379) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-57233d7f1efaa1662e990fefd4533a968b9f42a4b0e6cf17aa384cf20c4b4a6d) |
| `primary.default_rr_set_group.tlsa_record` | [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-0e894d29ce9c45ca3f67578ce312abf73327e74eea7572290520ce6331260d72) |
| `primary.default_rr_set_group.tlsa_record.name` | [primary.default_rr_set_group.tlsa_record.name](data-sources--dns_zone--reference--group-002.md#canonical-6c9cff6a3cb901e4e7a653e4e5ae8f88511183e52f32d034269b08eba4955294) |
| `primary.default_rr_set_group.tlsa_record.values` | [primary.default_rr_set_group.tlsa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-c204f36c2b25620d8db07176a6e442f69cbc114effd68f75a8e81c63cd392648) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_association_data` | [primary.default_rr_set_group.tlsa_record.values.certificate_association_data](data-sources--dns_zone--reference--group-002.md#canonical-114b90e6d79a2ca194cc9475f3180fb535a2dd921e6a92a6251d52f15543a560) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_usage` | [primary.default_rr_set_group.tlsa_record.values.certificate_usage](data-sources--dns_zone--reference--group-002.md#canonical-1b92c812e7267aa382f4e4615b073a6ffb4225fadc997d50cff0ec646d1889ff) |
| `primary.default_rr_set_group.tlsa_record.values.matching_type` | [primary.default_rr_set_group.tlsa_record.values.matching_type](data-sources--dns_zone--reference--group-002.md#canonical-5b30af807acae286e906fa64efe9b74d9731d5666003ac2ea9bd6bba5d22f038) |
| `primary.default_rr_set_group.tlsa_record.values.selector` | [primary.default_rr_set_group.tlsa_record.values.selector](data-sources--dns_zone--reference--group-002.md#canonical-66d7defcd94756c2fa38b7347f004977252e132e95288142f8ddb8f3727f3962) |
| `primary.default_rr_set_group.ttl` | [primary.default_rr_set_group.ttl](data-sources--dns_zone--reference--group-001.md#canonical-58bfbacda295bf891d5f55afdfc00d3da413e1ff06726ba3eb8abc97b00515c6) |
| `primary.default_rr_set_group.txt_record` | [primary.default_rr_set_group.txt_record](data-sources--dns_zone--reference--group-002.md#canonical-179f09d346e4e66438ee5eb5977c299da0e9087c81faf4eff3ff5faf195aec94) |
| `primary.default_rr_set_group.txt_record.name` | [primary.default_rr_set_group.txt_record.name](data-sources--dns_zone--reference--group-002.md#canonical-dee86190a71fbed997427f9bebf3ba94707e62ea06fd08cee554e1bebf9da25b) |
| `primary.default_rr_set_group.txt_record.values` | [primary.default_rr_set_group.txt_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2ec5ce4e5b151ef5534facca66026e7f270eacdee5adf8e2555749bfcbc2417b) |
| `primary.default_soa_parameters` | [primary.default_soa_parameters](data-sources--dns_zone--reference--group-002.md#canonical-604b9eb5eb88a922c4efad10870c5b604ef6fd5e091b8412d7fabd1911f1ff5a) |
| `primary.dnssec_mode` | [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-43ec0c7c7a20efc75929dcdeb54460a21c7c3a901f5d7583999d864e89c44656) |
| `primary.dnssec_mode.disable_spec` | [primary.dnssec_mode.disable_spec](data-sources--dns_zone--reference--group-002.md#canonical-1f44a53646b2983e16addf0adabffbbe33a32a436cad308ad69231a2ff507ee5) |
| `primary.dnssec_mode.enable` | [primary.dnssec_mode.enable](data-sources--dns_zone--reference--group-002.md#canonical-498a22ed638be5660333bfbf9088b47bebcbc8c0ed2897b744446e350fe7dae9) |
| `primary.rr_set_group` | [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-ad8b39448944154c872d1aeff99dc227fee3014b9995e147b635e9205d3427cd) |
| `primary.rr_set_group.metadata` | [primary.rr_set_group.metadata](data-sources--dns_zone--reference--group-002.md#canonical-7c0f3b168b0350ee37250b8d9c425c3c7a4a0d1e15974291fd38004143feb106) |
| `primary.rr_set_group.metadata.description_spec` | [primary.rr_set_group.metadata.description_spec](data-sources--dns_zone--reference--group-002.md#canonical-7af1b8fdb8902aa2e9c75206bcceae664d17a43594b802920e017bd93322f6b5) |
| `primary.rr_set_group.metadata.name` | [primary.rr_set_group.metadata.name](data-sources--dns_zone--reference--group-002.md#canonical-a22dd40a53caea61c516f0de6a7774bd076d78ce717a7592c2bd80bb5ccd09e9) |
| `primary.rr_set_group.rr_set` | [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-f7986ebe5ed53b2253a23fdcf4a471734fa13c9149000f0f64dc99f00df3d425) |
| `primary.rr_set_group.rr_set.a_record` | [primary.rr_set_group.rr_set.a_record](data-sources--dns_zone--reference--group-002.md#canonical-c59b8540e2fc762909995f8b8e6ee395215b462bf481bfd626ae8281815f610d) |
| `primary.rr_set_group.rr_set.a_record.name` | [primary.rr_set_group.rr_set.a_record.name](data-sources--dns_zone--reference--group-002.md#canonical-e74addcf33f7a642f0d8f59c215765864338b525abc3fed865eb2d8789eacf2c) |
| `primary.rr_set_group.rr_set.a_record.values` | [primary.rr_set_group.rr_set.a_record.values](data-sources--dns_zone--reference--group-002.md#canonical-16ed710bcb7475cdd2761e8daf5c2826dcffedcac2a00a7d44db31454b24c5f0) |
| `primary.rr_set_group.rr_set.aaaa_record` | [primary.rr_set_group.rr_set.aaaa_record](data-sources--dns_zone--reference--group-002.md#canonical-8bf46d66a2b3313ca866de0f05be88da6a316c785db81ef4b55557dc1e624639) |
| `primary.rr_set_group.rr_set.aaaa_record.name` | [primary.rr_set_group.rr_set.aaaa_record.name](data-sources--dns_zone--reference--group-002.md#canonical-1e700ec01892ce573ec33cd19df4e054138f54dcddb9743ad33c5ab1e3350f34) |
| `primary.rr_set_group.rr_set.aaaa_record.values` | [primary.rr_set_group.rr_set.aaaa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-6190d72e5925ef8e108017d687a069cf393ff5fe82cf9852d3307b13ce161e89) |
| `primary.rr_set_group.rr_set.afsdb_record` | [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-cfcaa49dcf9066a96ba20bd9808f1c1bb54dbf262798671a2c87814e2d792e7c) |
| `primary.rr_set_group.rr_set.afsdb_record.name` | [primary.rr_set_group.rr_set.afsdb_record.name](data-sources--dns_zone--reference--group-002.md#canonical-0ec6f6220fb57416d5f0bc66a011c58a3da429e50fc14824bacaa6be705e296a) |
| `primary.rr_set_group.rr_set.afsdb_record.values` | [primary.rr_set_group.rr_set.afsdb_record.values](data-sources--dns_zone--reference--group-002.md#canonical-4b060c621406e0c089c13fc11b54200e78848c9f6f974b04aff5c45c4ef36ea2) |
| `primary.rr_set_group.rr_set.afsdb_record.values.hostname` | [primary.rr_set_group.rr_set.afsdb_record.values.hostname](data-sources--dns_zone--reference--group-002.md#canonical-54b52e1de5eebc9831c7c4c8d23358c174fe92afa1601a654ddab7328083b74a) |
| `primary.rr_set_group.rr_set.afsdb_record.values.subtype` | [primary.rr_set_group.rr_set.afsdb_record.values.subtype](data-sources--dns_zone--reference--group-002.md#canonical-f4ca63f5500cb1fcde93a0b41995b85c88603159641f8507ad8e4845fbab7b6c) |
| `primary.rr_set_group.rr_set.alias_record` | [primary.rr_set_group.rr_set.alias_record](data-sources--dns_zone--reference--group-002.md#canonical-fc47f5d66110adfac70f5119899f9d7178afed36f72c82123fa9edcb77280f5f) |
| `primary.rr_set_group.rr_set.alias_record.value` | [primary.rr_set_group.rr_set.alias_record.value](data-sources--dns_zone--reference--group-002.md#canonical-19486adac784aa364f7eac6e2d8b975cefdbdf4afc2d8ccc3e75ec6af1469cfc) |
| `primary.rr_set_group.rr_set.caa_record` | [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--reference--group-002.md#canonical-fc5918f983f498ae444bb488cc4a135f8e8ff33c46c4ad23d2204b3399d43df5) |
| `primary.rr_set_group.rr_set.caa_record.name` | [primary.rr_set_group.rr_set.caa_record.name](data-sources--dns_zone--reference--group-002.md#canonical-e32c4a259cc8d60d06f4bc7ff4cea62e974c47dcb50798481e2d2d0f4fc387f4) |
| `primary.rr_set_group.rr_set.caa_record.values` | [primary.rr_set_group.rr_set.caa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-af0bb68d3b5ed4effec1cb08b11954d1229ecd8d0a3ab9d3a8a4df6aaa9cb44f) |
| `primary.rr_set_group.rr_set.caa_record.values.flags` | [primary.rr_set_group.rr_set.caa_record.values.flags](data-sources--dns_zone--reference--group-002.md#canonical-a5e5f9944ddee3ded3f47a1633918bd93125d1cee1057642c91a66b66e4e84fa) |
| `primary.rr_set_group.rr_set.caa_record.values.tag` | [primary.rr_set_group.rr_set.caa_record.values.tag](data-sources--dns_zone--reference--group-002.md#canonical-5193bad77dae77bb7a831114801917b79ae9bc658b04bec4be1c9354e93a0ca8) |
| `primary.rr_set_group.rr_set.caa_record.values.value` | [primary.rr_set_group.rr_set.caa_record.values.value](data-sources--dns_zone--reference--group-002.md#canonical-2cde9979fbc4c4ee594962c2e1421295bf7e81e5d82091f1b57616ffe8473cff) |
| `primary.rr_set_group.rr_set.cds_record` | [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-d64aa7eac3a3d580782774a205c26e5d97134d25e480706c04d2fed4c6cc66e0) |
| `primary.rr_set_group.rr_set.cds_record.name` | [primary.rr_set_group.rr_set.cds_record.name](data-sources--dns_zone--reference--group-002.md#canonical-f25c900cce87444866f0249de20c7efbdbeb89832e4d9142a385075b4d030a70) |
| `primary.rr_set_group.rr_set.cds_record.values` | [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3b734a95822c51ae245f24a56f83a3fc17571058a321d9687a3023f6de417077) |
| `primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm](data-sources--dns_zone--reference--group-002.md#canonical-7521f90f38d90c4f2033716abf56d5f57533acf551dca88b95a88f538c295e74) |
| `primary.rr_set_group.rr_set.cds_record.values.key_tag` | [primary.rr_set_group.rr_set.cds_record.values.key_tag](data-sources--dns_zone--reference--group-002.md#canonical-d74ed67bed30db1f7081a20ffceccdfcc559a56383caa189ba5ba753719ebdce) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-b92a26c82a6684dbaa7704a01309f75e2440658e5208bceb46322002690bc4dc) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-9fe80f480c43a16430fbbbb33f5337496456b3533213fa5d2798e168e1bacf1a) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-808a206664a36f7c1475f3a670d7bc94d19638d1b9073418f6ba46493ea1001f) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-a23c208940549a0d7d3e46aff7340025c0c2f5d295b3c6feac221cad0b093add) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-863f8fe478da1ff43dcaabb543f724ae4e46d5926d249b11ff85e2552e5c4978) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-a53332073d83865be6d782544db7e1939d6aadd48adfd44920b8ce286146e596) |
| `primary.rr_set_group.rr_set.cert_record` | [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--reference--group-002.md#canonical-7feeb9be87c907fd69d8d2fd398f225bffd3a4c367d015f325d8cd71cb4c6ecf) |
| `primary.rr_set_group.rr_set.cert_record.name` | [primary.rr_set_group.rr_set.cert_record.name](data-sources--dns_zone--reference--group-002.md#canonical-062a2d0ee952165251f63149889fe14a7c4b40bd000d0f599062916f33ae4bb4) |
| `primary.rr_set_group.rr_set.cert_record.values` | [primary.rr_set_group.rr_set.cert_record.values](data-sources--dns_zone--reference--group-002.md#canonical-9c6f945acfbe2f50551c6371776d9397cbad8f5980fb6c33c02a4eb1904a265a) |
| `primary.rr_set_group.rr_set.cert_record.values.algorithm` | [primary.rr_set_group.rr_set.cert_record.values.algorithm](data-sources--dns_zone--reference--group-002.md#canonical-ff3b6466a275be6a487788a0fba01b48fff064df8d181c788326529efa033359) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_key_tag` | [primary.rr_set_group.rr_set.cert_record.values.cert_key_tag](data-sources--dns_zone--reference--group-002.md#canonical-d0dd8ec9d4cda3dd5dda55225f37dab6efc0d64017f8fb6943ba40b4815521c9) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_type` | [primary.rr_set_group.rr_set.cert_record.values.cert_type](data-sources--dns_zone--reference--group-002.md#canonical-417e838d67c19b4a418b0ec96e15a84faf59c776f08e2d0499d2966e8f04b183) |
| `primary.rr_set_group.rr_set.cert_record.values.certificate` | [primary.rr_set_group.rr_set.cert_record.values.certificate](data-sources--dns_zone--reference--group-002.md#canonical-4acd0d999fe0fd86e8fa7f162231c09fd675a2d9913940e7aa2eae6ec323425f) |
| `primary.rr_set_group.rr_set.cname_record` | [primary.rr_set_group.rr_set.cname_record](data-sources--dns_zone--reference--group-002.md#canonical-4c1a7ccf02cf6149d5735bb50814f4093e891ef1df0b8b1daafa0fd212c960c2) |
| `primary.rr_set_group.rr_set.cname_record.name` | [primary.rr_set_group.rr_set.cname_record.name](data-sources--dns_zone--reference--group-002.md#canonical-83432a0fdad6f981172b4f83e093664e00dc0d1f0181e39c718d5a4813ef88f9) |
| `primary.rr_set_group.rr_set.cname_record.value` | [primary.rr_set_group.rr_set.cname_record.value](data-sources--dns_zone--reference--group-002.md#canonical-fc737e60e617bff0cb26bd8de49b5d2eae9c676bf131b6a58e5995f2c95bb9ac) |
| `primary.rr_set_group.rr_set.description_spec` | [primary.rr_set_group.rr_set.description_spec](data-sources--dns_zone--reference--group-002.md#canonical-80e2a92e70d533ededb9b3bcc1ad906dbdfa6c9063e15863197f31d654aafe4f) |
| `primary.rr_set_group.rr_set.ds_record` | [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-fb6edbaed969e5cf8438eb882bad11d7780a30e30965a9ce145c0e34dde02839) |
| `primary.rr_set_group.rr_set.ds_record.name` | [primary.rr_set_group.rr_set.ds_record.name](data-sources--dns_zone--reference--group-002.md#canonical-9810127fbcdbcb69c7b72d8046f2ffd196f04499ecc46a35a5dc1c5ee109af80) |
| `primary.rr_set_group.rr_set.ds_record.values` | [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-82926de713e09cf0d3b41ebf9519e0b8ed216aff7841301b2c22e5fb4c7381cf) |
| `primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm](data-sources--dns_zone--reference--group-002.md#canonical-805b5db828215dd92054223e66c7d06777b9fc6e7e5fb1985a81f5979a85ae27) |
| `primary.rr_set_group.rr_set.ds_record.values.key_tag` | [primary.rr_set_group.rr_set.ds_record.values.key_tag](data-sources--dns_zone--reference--group-002.md#canonical-acaec0306ebb65b08ebe55d00a49855c58a97850bc0fe1bd7ba2c95531fcdf49) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](data-sources--dns_zone--reference--group-003.md#canonical-7504b1c498fa12303cceeac603e555e8ef7b5c3b3d42a6e3044366b022565d45) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest](data-sources--dns_zone--reference--group-003.md#canonical-252b5ef1bf8dd1a84d5310cf5ef457eccc7e1177906ff5805a3520c74e959bd1) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](data-sources--dns_zone--reference--group-003.md#canonical-d8ff2e1b49fcaafcf25674b095dcc6de333733ee24134086260157f195d8926d) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest](data-sources--dns_zone--reference--group-003.md#canonical-008bf175a56b85d9b5811551e0bdb861aa8b46f0df6566476753f75dce30fa4c) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](data-sources--dns_zone--reference--group-003.md#canonical-c3b79f4b9ccb28a571cd170743737b4855fb7db64f25a19fd13cdad1bf2049c3) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest](data-sources--dns_zone--reference--group-003.md#canonical-51a78a6b6df8e956ed5c8cf930e532c05fc017c3de3ad9f901178232b3b919ad) |
| `primary.rr_set_group.rr_set.eui48_record` | [primary.rr_set_group.rr_set.eui48_record](data-sources--dns_zone--reference--group-003.md#canonical-810031a11394fca5e4f73d201eef2df69d1438e9d1e5440c666590ae840a695c) |
| `primary.rr_set_group.rr_set.eui48_record.name` | [primary.rr_set_group.rr_set.eui48_record.name](data-sources--dns_zone--reference--group-003.md#canonical-dc9ae63b5aee65c40119880b21730639ac546cc854cc00710030806cc5518992) |
| `primary.rr_set_group.rr_set.eui48_record.value` | [primary.rr_set_group.rr_set.eui48_record.value](data-sources--dns_zone--reference--group-003.md#canonical-59a2a901a2075caaeaeeefe5883585a095875bb17c859c7df2d30a095b7f9eff) |
| `primary.rr_set_group.rr_set.eui64_record` | [primary.rr_set_group.rr_set.eui64_record](data-sources--dns_zone--reference--group-003.md#canonical-6686d6e014d02ba9405db6396066587a5b4165c26c4ec0edda637861dfacac5b) |
| `primary.rr_set_group.rr_set.eui64_record.name` | [primary.rr_set_group.rr_set.eui64_record.name](data-sources--dns_zone--reference--group-003.md#canonical-eb34300a71ecdda472ea92475fd24d1e682d984728c7d08e3c4d430065a51798) |
| `primary.rr_set_group.rr_set.eui64_record.value` | [primary.rr_set_group.rr_set.eui64_record.value](data-sources--dns_zone--reference--group-003.md#canonical-bd811cdda1c9e4e57406c720f97ace2bdfe9006b450ad361d6963c97754d0cd6) |
| `primary.rr_set_group.rr_set.lb_record` | [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--reference--group-003.md#canonical-4408556c0c0318b5870bac11cf58181cf0b13b448ea5649eefd212cb54472ebe) |
| `primary.rr_set_group.rr_set.lb_record.name` | [primary.rr_set_group.rr_set.lb_record.name](data-sources--dns_zone--reference--group-003.md#canonical-00e48968331000a826f9874fae909196e7bfbcaa8a18076e25d1a2b62745affb) |
| `primary.rr_set_group.rr_set.lb_record.value` | [primary.rr_set_group.rr_set.lb_record.value](data-sources--dns_zone--reference--group-003.md#canonical-d9e2b52b7e6d30aab7103a2c86b6dff244ac5947b9a6f0263d3d9dd10bbc4355) |
| `primary.rr_set_group.rr_set.lb_record.value.name` | [primary.rr_set_group.rr_set.lb_record.value.name](data-sources--dns_zone--reference--group-003.md#canonical-008a89b7d5e756f29d25a8d1c0b3461156b42cd4732fbe6e0796e2598dad7f4a) |
| `primary.rr_set_group.rr_set.lb_record.value.namespace` | [primary.rr_set_group.rr_set.lb_record.value.namespace](data-sources--dns_zone--reference--group-003.md#canonical-f7e80e40d79db1a58e1ab44fad3e1e3f742d03c81033766124fdcdb3a92a7b1f) |
| `primary.rr_set_group.rr_set.lb_record.value.tenant` | [primary.rr_set_group.rr_set.lb_record.value.tenant](data-sources--dns_zone--reference--group-003.md#canonical-bccc7aac05ad64bc0aa8bd088eb7e2196e49c84b628d380023aa5bf1fe8a1391) |
| `primary.rr_set_group.rr_set.loc_record` | [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--reference--group-003.md#canonical-b4bcc079b229881885ecb1aefc7aa1add949d4213e77520181f4f9a04b9545a2) |
| `primary.rr_set_group.rr_set.loc_record.name` | [primary.rr_set_group.rr_set.loc_record.name](data-sources--dns_zone--reference--group-003.md#canonical-de598fe5afb8fb6540f6787451a2c9e105397f6c4cf1a8154d207ff4bf6810e6) |
| `primary.rr_set_group.rr_set.loc_record.values` | [primary.rr_set_group.rr_set.loc_record.values](data-sources--dns_zone--reference--group-003.md#canonical-6559c3f701fe4a0746cf60de728e90a1452bb8249376e95e56e8e5fa91fdd896) |
| `primary.rr_set_group.rr_set.loc_record.values.altitude` | [primary.rr_set_group.rr_set.loc_record.values.altitude](data-sources--dns_zone--reference--group-003.md#canonical-aa79600c34f1bf1e396d885c156dc549ba5e329ee9463efadaae58896866d7fc) |
| `primary.rr_set_group.rr_set.loc_record.values.horizontal_precision` | [primary.rr_set_group.rr_set.loc_record.values.horizontal_precision](data-sources--dns_zone--reference--group-003.md#canonical-caf16c03914926ee8f90d5e3c2e98628a730d594b51270bcc6d7036ec967d493) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.latitude_degree](data-sources--dns_zone--reference--group-003.md#canonical-d4a9b8b6cbc88a12f397731ebf112863e9bb844d2745472299ed4a27701eb31b) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere](data-sources--dns_zone--reference--group-003.md#canonical-15f32dd5be966c25d1cba0206fbe2d8524b91742736491eeebd8525ee1c3b258) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.latitude_minute](data-sources--dns_zone--reference--group-003.md#canonical-e06d796dfe61401114e1ed2001a8c3ba346bd22cb58d13a1b13cf4dab97dd8ee) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_second` | [primary.rr_set_group.rr_set.loc_record.values.latitude_second](data-sources--dns_zone--reference--group-003.md#canonical-a44bf5120df0917dc83ddfe24f56f0a89496fe6d0e381f49d30a856a22e6a208) |
| `primary.rr_set_group.rr_set.loc_record.values.location_diameter` | [primary.rr_set_group.rr_set.loc_record.values.location_diameter](data-sources--dns_zone--reference--group-003.md#canonical-6633655c7690ad2944868f0e28c98ff6aafbcb100ec65d92459aaa45169a1235) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.longitude_degree](data-sources--dns_zone--reference--group-003.md#canonical-950601e71463c53a3da7c8f3d27bfc99d420f1d8b43b1343b61dbeea04ec1e6b) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere](data-sources--dns_zone--reference--group-003.md#canonical-d586e3896825a8588f12a7000a9098d3d10be67775f46f26c00f25959b11d039) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.longitude_minute](data-sources--dns_zone--reference--group-003.md#canonical-4c97d03fb13e2b255346f64e4e7664ac7eaea7ae36115076e5f8698ac82c1c03) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_second` | [primary.rr_set_group.rr_set.loc_record.values.longitude_second](data-sources--dns_zone--reference--group-003.md#canonical-bdaf60beaa069e9b1249e8d4f9b19eb04e4a810c251a97d2459be83f4da7e452) |
| `primary.rr_set_group.rr_set.loc_record.values.vertical_precision` | [primary.rr_set_group.rr_set.loc_record.values.vertical_precision](data-sources--dns_zone--reference--group-003.md#canonical-a0f7d5c839c3150c425cab761b3ef5cf9171e606c6295853aca8318a3e936da5) |
| `primary.rr_set_group.rr_set.mx_record` | [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--reference--group-003.md#canonical-4ec6c672accef470e07480c5ebb5d615e5007733d71ce6ccb767ff8e0c9bd484) |
| `primary.rr_set_group.rr_set.mx_record.name` | [primary.rr_set_group.rr_set.mx_record.name](data-sources--dns_zone--reference--group-003.md#canonical-f53de048122a83156e354674ad2ae36f5f52a4575edc6d0cb72df4ba7fd7da3c) |
| `primary.rr_set_group.rr_set.mx_record.values` | [primary.rr_set_group.rr_set.mx_record.values](data-sources--dns_zone--reference--group-003.md#canonical-845f3db4794d898ba30aff13df27018137858d23bea2f1bbf41c3721b9a3a679) |
| `primary.rr_set_group.rr_set.mx_record.values.domain` | [primary.rr_set_group.rr_set.mx_record.values.domain](data-sources--dns_zone--reference--group-003.md#canonical-a65bdd45520d2d51f3c4425e5e83e4af73f09e9692d99ec9c84fe74ce5170c42) |
| `primary.rr_set_group.rr_set.mx_record.values.priority` | [primary.rr_set_group.rr_set.mx_record.values.priority](data-sources--dns_zone--reference--group-003.md#canonical-52783135a9dd7048e118602f26abc9f723f1de1765a90bb0e2b77a47f7c30d08) |
| `primary.rr_set_group.rr_set.naptr_record` | [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-1fb8e9110182e03b5c8b3763fdfefcfe067428bd38eddbfe00987bfd451a922f) |
| `primary.rr_set_group.rr_set.naptr_record.name` | [primary.rr_set_group.rr_set.naptr_record.name](data-sources--dns_zone--reference--group-003.md#canonical-66176eba368fdc47fa655911a0968015ce58f83f30e9408de2b841a856a9fce7) |
| `primary.rr_set_group.rr_set.naptr_record.values` | [primary.rr_set_group.rr_set.naptr_record.values](data-sources--dns_zone--reference--group-003.md#canonical-a0b5123207883a92c073ab214e9e068d0abe695b09ba8123c886b4651420cb8f) |
| `primary.rr_set_group.rr_set.naptr_record.values.flags` | [primary.rr_set_group.rr_set.naptr_record.values.flags](data-sources--dns_zone--reference--group-003.md#canonical-4d19819777dbafd38e4ea6c699ff6546cd732359a2b263fd23c570adb6ec75c7) |
| `primary.rr_set_group.rr_set.naptr_record.values.order` | [primary.rr_set_group.rr_set.naptr_record.values.order](data-sources--dns_zone--reference--group-003.md#canonical-b58adc7bd9fcac51ab30164a0a6365e133493d0ea8d8dab8361fbf95806b1449) |
| `primary.rr_set_group.rr_set.naptr_record.values.preference` | [primary.rr_set_group.rr_set.naptr_record.values.preference](data-sources--dns_zone--reference--group-003.md#canonical-0ea2ce6e24510c3f1570dd54c719c62425bf49475d3c4f59ab35f0cc8756e68e) |
| `primary.rr_set_group.rr_set.naptr_record.values.regexp` | [primary.rr_set_group.rr_set.naptr_record.values.regexp](data-sources--dns_zone--reference--group-003.md#canonical-11b2cfee7e832939dc5aa8881241847c197cbdf9d6141507bb27128b03fb47a2) |
| `primary.rr_set_group.rr_set.naptr_record.values.replacement` | [primary.rr_set_group.rr_set.naptr_record.values.replacement](data-sources--dns_zone--reference--group-003.md#canonical-a082d684611d2966b115a86dab3f65e0b5c891b4cc636474190595402a520115) |
| `primary.rr_set_group.rr_set.naptr_record.values.service` | [primary.rr_set_group.rr_set.naptr_record.values.service](data-sources--dns_zone--reference--group-003.md#canonical-2ae9174fbc1f230fba6216d9b7d7e127a442e0ddf100979d410e08a13738e0cf) |
| `primary.rr_set_group.rr_set.ns_record` | [primary.rr_set_group.rr_set.ns_record](data-sources--dns_zone--reference--group-003.md#canonical-3cda141896db18e8ec0c07290bb8033ee8a1250b5e6dd409d3b39a56faaf8753) |
| `primary.rr_set_group.rr_set.ns_record.name` | [primary.rr_set_group.rr_set.ns_record.name](data-sources--dns_zone--reference--group-003.md#canonical-2096f321da66047fe88b7d4167da9c51bf668062d05906ba319457e8cdde2171) |
| `primary.rr_set_group.rr_set.ns_record.values` | [primary.rr_set_group.rr_set.ns_record.values](data-sources--dns_zone--reference--group-003.md#canonical-99fc030a555b589c79ab6e502d1865c30ad1456ff856faed6c1532ccdc4bc471) |
| `primary.rr_set_group.rr_set.ptr_record` | [primary.rr_set_group.rr_set.ptr_record](data-sources--dns_zone--reference--group-003.md#canonical-ccc9c1934d838037e7e1e58ec57e5fd0a548bdd400d2361eaf6a0d6106504316) |
| `primary.rr_set_group.rr_set.ptr_record.name` | [primary.rr_set_group.rr_set.ptr_record.name](data-sources--dns_zone--reference--group-003.md#canonical-b85b91bda539f487865a7bb2a0daff9f8c686d7bc06b941bb1906da52e71e754) |
| `primary.rr_set_group.rr_set.ptr_record.values` | [primary.rr_set_group.rr_set.ptr_record.values](data-sources--dns_zone--reference--group-003.md#canonical-ca0299aa6b15e3ec4138389ec7a4f31a0c13d959c245232d8eaba8921a51ea72) |
| `primary.rr_set_group.rr_set.srv_record` | [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--reference--group-003.md#canonical-40a683168442e781ca09825c53ec073e529ad9fc1ab64bb88871eeaed9a97fb0) |
| `primary.rr_set_group.rr_set.srv_record.name` | [primary.rr_set_group.rr_set.srv_record.name](data-sources--dns_zone--reference--group-003.md#canonical-b2ea5c34eac052b424f33a9679b967e16b7a4e9567d02e31abc90d1722997343) |
| `primary.rr_set_group.rr_set.srv_record.values` | [primary.rr_set_group.rr_set.srv_record.values](data-sources--dns_zone--reference--group-003.md#canonical-a6bd35d735280c3a34b61e3450255b35b37036d0abfa24f115017f0669dee80e) |
| `primary.rr_set_group.rr_set.srv_record.values.port` | [primary.rr_set_group.rr_set.srv_record.values.port](data-sources--dns_zone--reference--group-003.md#canonical-c386881c49f2e04db91497ce02d0adae929d7f541e18aecf6b34b426398a4bec) |
| `primary.rr_set_group.rr_set.srv_record.values.priority` | [primary.rr_set_group.rr_set.srv_record.values.priority](data-sources--dns_zone--reference--group-003.md#canonical-e4a39153296d7a1bd1e616c5dee544f1ae992106f8197278e17bc8d1434c6049) |
| `primary.rr_set_group.rr_set.srv_record.values.target` | [primary.rr_set_group.rr_set.srv_record.values.target](data-sources--dns_zone--reference--group-003.md#canonical-aac320bc3d66ecd97bbf8758914a399ea742b7105da0a8bd907b53512420c031) |
| `primary.rr_set_group.rr_set.srv_record.values.weight` | [primary.rr_set_group.rr_set.srv_record.values.weight](data-sources--dns_zone--reference--group-003.md#canonical-37a06d4fdd276352840f3780b72320f6f1a6d078304e307a6ecc86d2cc24fda4) |
| `primary.rr_set_group.rr_set.sshfp_record` | [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-991b84ba14f554879e51f8f8d034a0790e4b9a26e957baa6d0d8002c18e9779a) |
| `primary.rr_set_group.rr_set.sshfp_record.name` | [primary.rr_set_group.rr_set.sshfp_record.name](data-sources--dns_zone--reference--group-003.md#canonical-6d3405e1f9ef813dfbea16372ab5f98947574df26e381ece6f26d093d7515849) |
| `primary.rr_set_group.rr_set.sshfp_record.values` | [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-bdc3cd1f0d8fb860080a18600c3d7135fca67ef979903e95c5b30e633ee89079) |
| `primary.rr_set_group.rr_set.sshfp_record.values.algorithm` | [primary.rr_set_group.rr_set.sshfp_record.values.algorithm](data-sources--dns_zone--reference--group-003.md#canonical-702cd975f7b85a731a98fb9a2648971942e3e22fb47bcdf86d3ec2296706eab0) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-741c22aedb434e08f27bb4dec4dad0883a95f4623c7246060b781e8129745fa9) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-2e01006ab08d3cafd9fd7f0de99b1cd5a4ec57545c4c9ed512dbf72d53b9d03c) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-89b6d74582130a3687d7e0babb4edff1ce48cbf839fd07cbab0d77c9ead56516) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-e18c17e61d7ee3a785d620ea8948ec65613eb02cb2edaca787317544096a5217) |
| `primary.rr_set_group.rr_set.tlsa_record` | [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-51ae6a52fa37f05c53e9569f4fef39edfe86bb9079b50a6d9c6a35872a194ac1) |
| `primary.rr_set_group.rr_set.tlsa_record.name` | [primary.rr_set_group.rr_set.tlsa_record.name](data-sources--dns_zone--reference--group-003.md#canonical-c715be307c0a5c7204a80f8a3d53eeeac710528e34dbd35feefad5c40541cbf1) |
| `primary.rr_set_group.rr_set.tlsa_record.values` | [primary.rr_set_group.rr_set.tlsa_record.values](data-sources--dns_zone--reference--group-003.md#canonical-b6f8693e9a9d359f5ec6cf202b02097d59134e2ec8a3f723f7dfbd92bce497bf) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data](data-sources--dns_zone--reference--group-003.md#canonical-8af5551b904dff0385968098a0abd1622f7ed6659071d202087ce46158436711) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage](data-sources--dns_zone--reference--group-003.md#canonical-4e3d9ca13702c59461ec445aa3d21cdcb25dcf15e11f285a4b8c01b022732c1a) |
| `primary.rr_set_group.rr_set.tlsa_record.values.matching_type` | [primary.rr_set_group.rr_set.tlsa_record.values.matching_type](data-sources--dns_zone--reference--group-003.md#canonical-adc08054a239dd89c74e0b5c5a44dd6cc26e3d904bb21bc4d9f1540ea96bc27a) |
| `primary.rr_set_group.rr_set.tlsa_record.values.selector` | [primary.rr_set_group.rr_set.tlsa_record.values.selector](data-sources--dns_zone--reference--group-003.md#canonical-3143a02aa4d153101881cecd37c3bcb9882c00cbac8ae74db235fdec2ab4c4dd) |
| `primary.rr_set_group.rr_set.ttl` | [primary.rr_set_group.rr_set.ttl](data-sources--dns_zone--reference--group-002.md#canonical-17eceb452bfeaca5fac0370c3b5c5754376f183ece5ec61a3b50b99016a4f79d) |
| `primary.rr_set_group.rr_set.txt_record` | [primary.rr_set_group.rr_set.txt_record](data-sources--dns_zone--reference--group-003.md#canonical-d81f9f35debece6115d30fca7ff3918acaa4efaf215e4358195592b8baaa6304) |
| `primary.rr_set_group.rr_set.txt_record.name` | [primary.rr_set_group.rr_set.txt_record.name](data-sources--dns_zone--reference--group-003.md#canonical-6c6c2cbf3843d3bf3e3a7227b711b5eadc390b920136f98cd51f0024ae2eae2e) |
| `primary.rr_set_group.rr_set.txt_record.values` | [primary.rr_set_group.rr_set.txt_record.values](data-sources--dns_zone--reference--group-003.md#canonical-b72114627c9def5e0051fc5d06b94503f3200ce2dc6930fd188fdb9ac1ead36c) |
| `primary.soa_parameters` | [primary.soa_parameters](data-sources--dns_zone--reference--group-003.md#canonical-fe80e3d93ec4f76e2a897015f27fef777246fbaa833e06034c70fe3b435e7d18) |
| `primary.soa_parameters.expire` | [primary.soa_parameters.expire](data-sources--dns_zone--reference--group-003.md#canonical-613c0cd5abcd80bd093c0b89db6960865bb89714cb9735c9f5d9f52402b55bb2) |
| `primary.soa_parameters.negative_ttl` | [primary.soa_parameters.negative_ttl](data-sources--dns_zone--reference--group-003.md#canonical-d0af71b88d27ec4abd84e8e2822c845a947cd66de1dc5c8b18a5f410e85c7c9a) |
| `primary.soa_parameters.refresh` | [primary.soa_parameters.refresh](data-sources--dns_zone--reference--group-003.md#canonical-d848b58cee658e8309c96991a11a953a280ceaa3b4d37c8862fbd5a71d7dd72c) |
| `primary.soa_parameters.retry` | [primary.soa_parameters.retry](data-sources--dns_zone--reference--group-003.md#canonical-1e5c9aa98707310ae96be7050f5c07a873e5ba4232e28acf9af45879bb53ed15) |
| `primary.soa_parameters.ttl` | [primary.soa_parameters.ttl](data-sources--dns_zone--reference--group-003.md#canonical-e76276a227d1dc59e27e967b12de72c80bb59a9579e0af2044e73ab51a114991) |
| `secondary` | [secondary](data-sources--dns_zone--reference--group-003.md#canonical-f7050a5d3f5c4338703440ab6df7f1041a87d40bcae52230201e1755f8c6f13b) |
| `secondary.primary_servers` | [secondary.primary_servers](data-sources--dns_zone--reference--group-003.md#canonical-adb4c13cb647a05efcc58c669e59d1dbf1cf12d76838b2c9797f67a814c6d80a) |
| `secondary.tsig_key_algorithm` | [secondary.tsig_key_algorithm](data-sources--dns_zone--reference--group-003.md#canonical-fcaf440114d560fdc74e768093e7a4e486ec2d00798f0be5a02e14248aab575e) |
| `secondary.tsig_key_name` | [secondary.tsig_key_name](data-sources--dns_zone--reference--group-003.md#canonical-8e26eaa9fdb935cd2a7650cc159f419518ba1e13e8efa1f46b3ea63e50286979) |
| `secondary.tsig_key_value` | [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-d1bfbfb13633c10816803bd5ccc8f9de6d3de5a269ec5f6fde448d683a3f8dd9) |
| `secondary.tsig_key_value.blindfold_secret_info` | [secondary.tsig_key_value.blindfold_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-fa2710e9450051dae9937dfa4495112c09e86e12f1aa21f53cde58d9f4e3bc37) |
| `secondary.tsig_key_value.blindfold_secret_info.decryption_provider` | [secondary.tsig_key_value.blindfold_secret_info.decryption_provider](data-sources--dns_zone--reference--group-003.md#canonical-a420e74677c48916ab415030785316c7195fa4b7146ae216b19d6dc2da002aa0) |
| `secondary.tsig_key_value.blindfold_secret_info.location` | [secondary.tsig_key_value.blindfold_secret_info.location](data-sources--dns_zone--reference--group-003.md#canonical-70dcd13d65ea27b23ba5fff19893a998b8cc4b9af289ad0c31de510950fd5129) |
| `secondary.tsig_key_value.blindfold_secret_info.store_provider` | [secondary.tsig_key_value.blindfold_secret_info.store_provider](data-sources--dns_zone--reference--group-003.md#canonical-7ee6128e259a9c3e8ad1a1eca5bcc5e3281c02e05fa705839a4c6ecb3a34ad86) |
| `secondary.tsig_key_value.clear_secret_info` | [secondary.tsig_key_value.clear_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-029c1f7b6041911bb160cbeaa3f1c9be7dca710a7c9edee284e6d8f414ac642a) |
| `secondary.tsig_key_value.clear_secret_info.provider_ref` | [secondary.tsig_key_value.clear_secret_info.provider_ref](data-sources--dns_zone--reference--group-003.md#canonical-c6270c5988bc8fb5c0828e1b800f32cbe2f72d1fad7f4f1ff2e65c30f030c355) |
| `secondary.tsig_key_value.clear_secret_info.url` | [secondary.tsig_key_value.clear_secret_info.url](data-sources--dns_zone--reference--group-003.md#canonical-7c6e8732918197fdc42895dafd2e5214ac78dd901130c435e7f5e515d34d7cd7) |

<a id="canonical-23f59fa4002e4041a0a061bd8a3810183b693a18266273df1fc8afc45f6dfb9a"></a>

## Next pages — Property reference / d4f8b7660ed9 / 11

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-1416f5ce025b5fd7645bf56ccfba5f70b02bac3f9d43100eb835a8e6d3d840cb)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31a28577a34910f9691a3759db0b974e79f0e4f7f21275e46984a9cf1d51c08b"></a>

## primary — primary / fc5c38fad7aa / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- primary

<a id="canonical-991194abf4edfd16575548ebc084e3a041bedb94121f2deceb51290542638a3d"></a>

Type: `"single"`. Computed.

\[OneOf: primary, secondary\] PrimaryDNSCreateSpecType.

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

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-991194abf4edfd16575548ebc084e3a041bedb94121f2deceb51290542638a3d)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-f7050a5d3f5c4338703440ab6df7f1041a87d40bcae52230201e1755f8c6f13b)

Select alternatives according to the provider validators above.

<a id="canonical-deb49d1526559364abd56dc6caa5656dcea8e45a9d6a10c1b39e5aaf28e77f63"></a>

## Direct properties — primary / fc5c38fad7aa / 3

<a id="canonical-05875f7e000466835a89dcf6a148db5b48990086e692d272f7939180bf7164f6"></a>

<a id="canonical-73a53cb24d62ff4151a7e44e78448a707702c22554b8e2059ba3614c98e4f6d7"></a>

## allow_http_lb_managed_records property — primary / fc5c38fad7aa / 4

Type: `"bool"`. Computed.

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

- [default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515): complete subsection reference.

- [default_soa_parameters](data-sources--dns_zone--reference--group-002.md#canonical-4739966520d2c572b26ed69d5e165bddefce0cd441f1bdc01a53b0559daf1428): complete subsection reference.

- [dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-317b4b9b10202b68289285d0affdfd75878c388455a4d278e28ad28df983a152): complete subsection reference.

- [rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2): complete subsection reference.

- [soa_parameters](data-sources--dns_zone--reference--group-003.md#canonical-d5c60d18869554ef8153a2525e58f7e65053f5b72480b618ef3dce844ed7aaa9): complete subsection reference.

<a id="canonical-03d9c8956498c96582e73a075c9e123e6ad48b719e545d9ea9a42ffbc0a46506"></a>

## Next pages — primary / fc5c38fad7aa / 5

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_soa_parameters](data-sources--dns_zone--reference--group-002.md#canonical-4739966520d2c572b26ed69d5e165bddefce0cd441f1bdc01a53b0559daf1428)
- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-317b4b9b10202b68289285d0affdfd75878c388455a4d278e28ad28df983a152)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.soa_parameters](data-sources--dns_zone--reference--group-003.md#canonical-d5c60d18869554ef8153a2525e58f7e65053f5b72480b618ef3dce844ed7aaa9)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b864147d7051e1cdd9eea2a31ae46b0dc45bf98be4482714e108c64b9b045db0"></a>

## primary.default_rr_set_group — primary.default_rr_set_group / 8fd02475d7b5 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- primary.default_rr_set_group

<a id="canonical-5748f65740c50bd17bc2bcdf7d61709eee6aa315620872aebee05c3708b35884"></a>

Type: `"list"`. Computed.

Add and manage DNS resource record sets part of Default set group.

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

<a id="canonical-35131a7d171dd787137f89f70a006d6f2afd05424b15a3d9f2e2d70d34da1e2f"></a>

## Direct properties — primary.default_rr_set_group / 8fd02475d7b5 / 3

- [a_record](data-sources--dns_zone--reference--group-001.md#canonical-b680c56066e52a5a73bf168b97413b18ec5712e4795d7cf03549cb56fa5e25c4): complete subsection reference.

- [aaaa_record](data-sources--dns_zone--reference--group-001.md#canonical-9a52d24db50b46650c45deb6829cb13b43d2a452d9bbabac0e861f698ad2c366): complete subsection reference.

- [afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-a6e4c73b3d0e4c6a422ef8ebc107f8fcb4831accac3f298f5fcdddf612a1ed2b): complete subsection reference.

- [alias_record](data-sources--dns_zone--reference--group-001.md#canonical-a1787db40d61a4ede0df65818aed8862c9eb29c3b064204cb07a16743389e38c): complete subsection reference.

- [caa_record](data-sources--dns_zone--reference--group-001.md#canonical-ad735989ad640d243f25f88ca63afcf0aff1a45cf0fd235536cbe0e0a15d4386): complete subsection reference.

- [cds_record](data-sources--dns_zone--reference--group-001.md#canonical-143fca8e1b3c7331722fcf357cb9aa76dfc70cbd1d3cb2b12cb2198d69d63e28): complete subsection reference.

- [cert_record](data-sources--dns_zone--reference--group-001.md#canonical-935f7e4f217a261d45e534d82676c2a545a619156396b4c6e75881f61a0a9ae9): complete subsection reference.

- [cname_record](data-sources--dns_zone--reference--group-001.md#canonical-64ab7b8bf27978045203dbaa7158071b3ec5c7f6164268a2cda56ee8c8f8d6a0): complete subsection reference.

<a id="canonical-17d879209731d27bde4e0df767c1806e179bc7b05be3585ec090404c5cbd6900"></a>

<a id="canonical-b70d14aeda312098892e94a70fda3173c2bb5d55002bbcd62a670995a80ed641"></a>

## description_spec property — primary.default_rr_set_group / 8fd02475d7b5 / 4

Type: `"string"`. Computed.

Comment. Human-readable description text

- [ds_record](data-sources--dns_zone--reference--group-001.md#canonical-3ef88f0643cae6c90abb5fde17697e969b3f9536dffdfc6919f32a93e600fe90): complete subsection reference.

- [eui48_record](data-sources--dns_zone--reference--group-002.md#canonical-ddc3e22cfd3dd01d590d10bc6f4a41ce27ef3658058ab351d7ea921a4d1b9e33): complete subsection reference.

- [eui64_record](data-sources--dns_zone--reference--group-002.md#canonical-4137c5085f685ae669234993993124f44df8c961c4f967580267db0ada9aff97): complete subsection reference.

- [lb_record](data-sources--dns_zone--reference--group-002.md#canonical-7952fecbcf1aa914505a44317d504dfc227553210d64c2d4f31b9a5262998d8b): complete subsection reference.

- [loc_record](data-sources--dns_zone--reference--group-002.md#canonical-f46870a2a1d15407f87baf33c3d6a6fd6909e4329a2462f53cc6f5f502719113): complete subsection reference.

- [mx_record](data-sources--dns_zone--reference--group-002.md#canonical-64661d589d12a3da3ca8de95e02806b23f24e6d6a88e106698338d1eba3a2566): complete subsection reference.

- [naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-80a671af6a9b0043a165a908dfc638927e3a2e31434b16a7a3e00c67065f3853): complete subsection reference.

- [ns_record](data-sources--dns_zone--reference--group-002.md#canonical-b7b5698521f36fec613ee3585b4ac7a5b57373ed28897f7ccc3859a18a84ba13): complete subsection reference.

- [ptr_record](data-sources--dns_zone--reference--group-002.md#canonical-0eb30da1bc8d2b099a6bcb576ef2dcb4c120ae431000dcb33809e1abfe6f4ed7): complete subsection reference.

- [srv_record](data-sources--dns_zone--reference--group-002.md#canonical-ca0b3e919fd28a9269cfb08289e37290cac4f875d0c586ae19393c9685948dce): complete subsection reference.

- [sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-9d9c34211610604207a24eec3afe20b57ba6a3ce26657a692a52f41a5a67fc49): complete subsection reference.

- [tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-91640050b110b8aa2107e7cbd04738b4c1a6f9219edda114d055c15eca75582c): complete subsection reference.

<a id="canonical-58bfbacda295bf891d5f55afdfc00d3da413e1ff06726ba3eb8abc97b00515c6"></a>

<a id="canonical-c6ff3db9e8c57a451ce6441739ae10b2bd0b344cef56cfad3b9e6e57da2ff415"></a>

## ttl property — primary.default_rr_set_group / 8fd02475d7b5 / 5

Type: `"number"`. Computed.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

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

- [txt_record](data-sources--dns_zone--reference--group-002.md#canonical-8b0f9fb60b79b92c52f09acdd60e5210cd48a11f397ffba637e2bb8f2e0cd337): complete subsection reference.

<a id="canonical-a08caa7600ae7995898fa1b1e277f9a492b8bcb01c70b3bd05000af9a39e703d"></a>

## Next pages — primary.default_rr_set_group / 8fd02475d7b5 / 6

- [primary.default_rr_set_group.a_record](data-sources--dns_zone--reference--group-001.md#canonical-b680c56066e52a5a73bf168b97413b18ec5712e4795d7cf03549cb56fa5e25c4)
- [primary.default_rr_set_group.aaaa_record](data-sources--dns_zone--reference--group-001.md#canonical-9a52d24db50b46650c45deb6829cb13b43d2a452d9bbabac0e861f698ad2c366)
- [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-a6e4c73b3d0e4c6a422ef8ebc107f8fcb4831accac3f298f5fcdddf612a1ed2b)
- [primary.default_rr_set_group.alias_record](data-sources--dns_zone--reference--group-001.md#canonical-a1787db40d61a4ede0df65818aed8862c9eb29c3b064204cb07a16743389e38c)
- [primary.default_rr_set_group.caa_record](data-sources--dns_zone--reference--group-001.md#canonical-ad735989ad640d243f25f88ca63afcf0aff1a45cf0fd235536cbe0e0a15d4386)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-143fca8e1b3c7331722fcf357cb9aa76dfc70cbd1d3cb2b12cb2198d69d63e28)
- [primary.default_rr_set_group.cert_record](data-sources--dns_zone--reference--group-001.md#canonical-935f7e4f217a261d45e534d82676c2a545a619156396b4c6e75881f61a0a9ae9)
- [primary.default_rr_set_group.cname_record](data-sources--dns_zone--reference--group-001.md#canonical-64ab7b8bf27978045203dbaa7158071b3ec5c7f6164268a2cda56ee8c8f8d6a0)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-3ef88f0643cae6c90abb5fde17697e969b3f9536dffdfc6919f32a93e600fe90)
- [primary.default_rr_set_group.eui48_record](data-sources--dns_zone--reference--group-002.md#canonical-ddc3e22cfd3dd01d590d10bc6f4a41ce27ef3658058ab351d7ea921a4d1b9e33)
- [primary.default_rr_set_group.eui64_record](data-sources--dns_zone--reference--group-002.md#canonical-4137c5085f685ae669234993993124f44df8c961c4f967580267db0ada9aff97)
- [primary.default_rr_set_group.lb_record](data-sources--dns_zone--reference--group-002.md#canonical-7952fecbcf1aa914505a44317d504dfc227553210d64c2d4f31b9a5262998d8b)
- [primary.default_rr_set_group.loc_record](data-sources--dns_zone--reference--group-002.md#canonical-f46870a2a1d15407f87baf33c3d6a6fd6909e4329a2462f53cc6f5f502719113)
- [primary.default_rr_set_group.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-64661d589d12a3da3ca8de95e02806b23f24e6d6a88e106698338d1eba3a2566)
- [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-80a671af6a9b0043a165a908dfc638927e3a2e31434b16a7a3e00c67065f3853)
- [primary.default_rr_set_group.ns_record](data-sources--dns_zone--reference--group-002.md#canonical-b7b5698521f36fec613ee3585b4ac7a5b57373ed28897f7ccc3859a18a84ba13)
- [primary.default_rr_set_group.ptr_record](data-sources--dns_zone--reference--group-002.md#canonical-0eb30da1bc8d2b099a6bcb576ef2dcb4c120ae431000dcb33809e1abfe6f4ed7)
- [primary.default_rr_set_group.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-ca0b3e919fd28a9269cfb08289e37290cac4f875d0c586ae19393c9685948dce)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-9d9c34211610604207a24eec3afe20b57ba6a3ce26657a692a52f41a5a67fc49)
- [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-91640050b110b8aa2107e7cbd04738b4c1a6f9219edda114d055c15eca75582c)
- [primary.default_rr_set_group.txt_record](data-sources--dns_zone--reference--group-002.md#canonical-8b0f9fb60b79b92c52f09acdd60e5210cd48a11f397ffba637e2bb8f2e0cd337)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-b680c56066e52a5a73bf168b97413b18ec5712e4795d7cf03549cb56fa5e25c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea06839d7d9106b508430eca1cd7953d7734c66ebec7383d1f7a4fcb47668ab0"></a>

## primary.default_rr_set_group.a_record — primary.default_rr_set_group.a_record / 5229c5ec12e2 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.a_record

<a id="canonical-a2300cf8a33cb09199c979b85d5b0adf5a4eb35c0305e7da0e733c6f2399ea49"></a>

Type: `"single"`. Computed.

DNSAResourceRecord. A Records

Upstream description:

A Records

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

<a id="canonical-b582e8176b9a7537426c2cbac2fb1ad9ffa44dcc2096c8c2a8aeef111aea7054"></a>

## Direct properties — primary.default_rr_set_group.a_record / 5229c5ec12e2 / 3

<a id="canonical-a7bc9e04df108de7b707747192087c75a25dfef9458d6c74f88258dfc3fce045"></a>

<a id="canonical-a4f2605c5be4843673eb23eff4de39aef03c1bc683bbc46f313e7ef2dc0863a4"></a>

## name property — primary.default_rr_set_group.a_record / 5229c5ec12e2 / 4

Type: `"string"`. Computed.

Record name, please provide only the specific subdomain or record name without the base domain.

Upstream description:

A Record name, please provide only the specific subdomain or record name without the base domain.

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

<a id="canonical-64d6448fdef132fa77bd9d55aa0a5cd75d20ad700e761982c55767ee12c2c6e9"></a>

<a id="canonical-1c28300da13d89c3a4aedc29b884f3c989224a83800351712149825658e1aee8"></a>

## values property — primary.default_rr_set_group.a_record / 5229c5ec12e2 / 5

Type: `["list", "string"]`. Computed.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

Upstream description:

A valid IPv4 address, for example: 192.0.2.242.

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

<a id="canonical-827cd768446cab109de120a52a2e88db9c0ba05cda24f4981f5759f3864b7a93"></a>

## Next pages — primary.default_rr_set_group.a_record / 5229c5ec12e2 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-9a52d24db50b46650c45deb6829cb13b43d2a452d9bbabac0e861f698ad2c366"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be2c4027e86745b9e3a8b6826d24e98ae3c45979005bdaadbef1a6fc3d1e4501"></a>

## primary.default_rr_set_group.aaaa_record — primary.default_rr_set_group.aaaa_record / b6bcd190ac1a / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.aaaa_record

<a id="canonical-373d92e2365372276e53adbddf0365f445a7016cb8cc4ba8a8a3feba7872a5dd"></a>

Type: `"single"`. Computed.

Configuration parameter for aaaa record.

Upstream description:

RecordSet for AAAA Records.

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

<a id="canonical-e069df8dea1c6a34cffa6cc7bca25ac1563b9ca41f54eb5e2304392673ba9fbe"></a>

## Direct properties — primary.default_rr_set_group.aaaa_record / b6bcd190ac1a / 3

<a id="canonical-d23c1177f69df50d69b8eb6d6d32e794fb46c7588e912d2338b77ef2ca615e98"></a>

<a id="canonical-1815086632118c5f51e08a930803d3133be339bd5ab128a606f21ffbbc4a065a"></a>

## name property — primary.default_rr_set_group.aaaa_record / b6bcd190ac1a / 4

Type: `"string"`. Computed.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

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

<a id="canonical-b30b80e9d49030569604229f828dd63891c9b3a0b6cc028ee6ee3e29652c56bb"></a>

<a id="canonical-f95924eb4ea57d68a6123e0007c5479ee2a2426d93ed6f6603f4d6fd55ff8cda"></a>

## values property — primary.default_rr_set_group.aaaa_record / b6bcd190ac1a / 5

Type: `["list", "string"]`. Computed.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Upstream description:

A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

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

<a id="canonical-ae48b945c2cea1bea8ace268016225daf8a5dff442643ea1a79df863140112e4"></a>

## Next pages — primary.default_rr_set_group.aaaa_record / b6bcd190ac1a / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-a6e4c73b3d0e4c6a422ef8ebc107f8fcb4831accac3f298f5fcdddf612a1ed2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5176bebc17dc3409f9b06d2f75284e4b5fe11e4018dbc2a9029ee47d121d7592"></a>

## primary.default_rr_set_group.afsdb_record — primary.default_rr_set_group.afsdb_record / ba9b131bc7d2 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.afsdb_record

<a id="canonical-de718cdbd5f577ea03a3a7ad69609fda1fdc6a17862764b256027bc5772b61d5"></a>

Type: `"single"`. Computed.

Configuration parameter for afsdb record.

Upstream description:

DNS AFSDB Record.

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

<a id="canonical-5678a198c2a9f61870f1941a86072e01b87f9542c1592b350c722053dbc710d6"></a>

## Direct properties — primary.default_rr_set_group.afsdb_record / ba9b131bc7d2 / 3

<a id="canonical-14c3d71b9ed9b47ff7cc3b5ce45359ff6a03f441bd63dc79e43bf2422b8cbb44"></a>

<a id="canonical-436137fe085a2612979ec9608d06b3fd406325fdd2cdc0f999c6d96151c15d2a"></a>

## name property — primary.default_rr_set_group.afsdb_record / ba9b131bc7d2 / 4

Type: `"string"`. Computed.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

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

- [values](data-sources--dns_zone--reference--group-001.md#canonical-fbbfde0a40030c6d93a78c42436db74dc2d8d6701af023a240d5ee45877573f4): complete subsection reference.

<a id="canonical-146e6a1c37b95407880d6d621f01b4a01a6aa8fa38787abcd18b42e5636f8438"></a>

## Next pages — primary.default_rr_set_group.afsdb_record / ba9b131bc7d2 / 5

- [primary.default_rr_set_group.afsdb_record.values](data-sources--dns_zone--reference--group-001.md#canonical-fbbfde0a40030c6d93a78c42436db74dc2d8d6701af023a240d5ee45877573f4)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-fbbfde0a40030c6d93a78c42436db74dc2d8d6701af023a240d5ee45877573f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e81e1c1c3bff22d8e23fd10b6434344e3642635c64418145da193a6ae1838308"></a>

## primary.default_rr_set_group.afsdb_record.values — primary.default_rr_set_group.afsdb_record.values / f314eea5fe22 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-a6e4c73b3d0e4c6a422ef8ebc107f8fcb4831accac3f298f5fcdddf612a1ed2b)
- primary.default_rr_set_group.afsdb_record.values

<a id="canonical-dc0bf29d2c0cb263cccf2fbe8485aef77576e249cb3208c5dae82fb5e721e534"></a>

Type: `"list"`. Computed.

AFSDB Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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

<a id="canonical-0a6e7867840115d533ef85779ea72e19c276be95fca1c3787923a3d09b513e9c"></a>

## Direct properties — primary.default_rr_set_group.afsdb_record.values / f314eea5fe22 / 3

<a id="canonical-423f6991bd603d66561b5226dd7ade0e9d0495d0c7fd06afbe4427467c95c513"></a>

<a id="canonical-175762cae27b88ffbdd923d1b690cd7868826f217633f339d0afb2267e98e8fe"></a>

## hostname property — primary.default_rr_set_group.afsdb_record.values / f314eea5fe22 / 4

Type: `"string"`. Computed.

Server name of the AFS cell database server or the DCE name server.

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

<a id="canonical-79d6fab20e287aa16ae993a2fd348ca9ac7f657690f3dfa58c4619998a42cd68"></a>

<a id="canonical-f82d75920f823a98f2dec1794f1af5e2f6afa58ca04d371dd128cc342fb8f233"></a>

## subtype property — primary.default_rr_set_group.afsdb_record.values / f314eea5fe22 / 5

Type: `"string"`. Computed.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Upstream description:

AFS Volume Location Server or DCE Authentication Server.

&#8203;- NONE: NONE

&#8203;- AFSVolumeLocationServer: AFS Volume Location Server

&#8203;- DCEAuthenticationServer: DCE Authentication Server.

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

<a id="canonical-369fef520fb1ba05bea192d1d302c13ee17c51e24ac9017b3a2cc173579488ac"></a>

## Next pages — primary.default_rr_set_group.afsdb_record.values / f314eea5fe22 / 6

- [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-a6e4c73b3d0e4c6a422ef8ebc107f8fcb4831accac3f298f5fcdddf612a1ed2b)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-a1787db40d61a4ede0df65818aed8862c9eb29c3b064204cb07a16743389e38c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f8c92a9960252a4f92d9816c24f5bc95bcb0761cbdaf9698f84ad729c001229"></a>

## primary.default_rr_set_group.alias_record — primary.default_rr_set_group.alias_record / 4480cfa9825d / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.alias_record

<a id="canonical-dbcdbdad5885c364c0821e484bda01b18a52005c0478de41ede87bd48ad713f5"></a>

Type: `"single"`. Computed.

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

<a id="canonical-dc82e3b0c2efecb18c9cc8ce88f0f1b83edff34033cde3cf6c23054dec0f6cfe"></a>

## Direct properties — primary.default_rr_set_group.alias_record / 4480cfa9825d / 3

<a id="canonical-03ff40e52ec1b822f4df04e1813f2d8ec6cf412ad06bb333e2cda2a26ca1954f"></a>

<a id="canonical-3cc57eccd2c57c1eee9443593ccfd826a3e61ddb2ba616fec26110cc6aaa6015"></a>

## value property — primary.default_rr_set_group.alias_record / 4480cfa9825d / 4

Type: `"string"`. Computed.

Domain. A valid domain name, for example: example.com.

Upstream description:

A valid domain name, for example: example.com.

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

<a id="canonical-97e26884f737e63e76b8e67e9669afb2af7fce044f9adca58aaf1ef4bd52ab65"></a>

## Next pages — primary.default_rr_set_group.alias_record / 4480cfa9825d / 5

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-ad735989ad640d243f25f88ca63afcf0aff1a45cf0fd235536cbe0e0a15d4386"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63977727d6468c7ba279519e1df4becdf4f8dc8bbf61bea8b88295a1d6144a16"></a>

## primary.default_rr_set_group.caa_record — primary.default_rr_set_group.caa_record / 37292a492d8a / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.caa_record

<a id="canonical-95b2a44978900fc848cfd9da410e1c1033f0820aed35c6e846b39c3ab23d5b19"></a>

Type: `"single"`. Computed.

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

<a id="canonical-b22c61cfe67ad6e118db98a80e66b4212e88ca8f0d9c03ab6122e10280f2dc32"></a>

## Direct properties — primary.default_rr_set_group.caa_record / 37292a492d8a / 3

<a id="canonical-c00fe2f75bd0bcb6de41bfd4bc698f8593a9af39e52ba41d1a3d4e7d8646095c"></a>

<a id="canonical-e58098f73ce2fedbd4f5a898593b9cbfdd8144f0b81de754a81e3fbe7e99e1dc"></a>

## name property — primary.default_rr_set_group.caa_record / 37292a492d8a / 4

Type: `"string"`. Computed.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

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

- [values](data-sources--dns_zone--reference--group-001.md#canonical-431248beff88b3f7e2d3ec8b17ed52403bfe3496b26f52f814f9e812fa487063): complete subsection reference.

<a id="canonical-1a50303a8ff30200599cefa96ad7b488a8fbadeb4ee38d04f76000bd5fbca62f"></a>

## Next pages — primary.default_rr_set_group.caa_record / 37292a492d8a / 5

- [primary.default_rr_set_group.caa_record.values](data-sources--dns_zone--reference--group-001.md#canonical-431248beff88b3f7e2d3ec8b17ed52403bfe3496b26f52f814f9e812fa487063)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-431248beff88b3f7e2d3ec8b17ed52403bfe3496b26f52f814f9e812fa487063"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6649655bb9f54da050747e8c7bd397915852948e6371c7685ee0050d26c82650"></a>

## primary.default_rr_set_group.caa_record.values — primary.default_rr_set_group.caa_record.values / d210c9c599e3 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.caa_record](data-sources--dns_zone--reference--group-001.md#canonical-ad735989ad640d243f25f88ca63afcf0aff1a45cf0fd235536cbe0e0a15d4386)
- primary.default_rr_set_group.caa_record.values

<a id="canonical-1eb96f1ac8cca66f9b5517d9b213b52c0c01a4c858053edbc4e35e6b859f0643"></a>

Type: `"list"`. Computed.

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

<a id="canonical-f4d44f41e8cf1bac55f384f18d73d04cdf8ebd17576762021e3646b00361c979"></a>

## Direct properties — primary.default_rr_set_group.caa_record.values / d210c9c599e3 / 3

<a id="canonical-bbd76a8107630d7afdf91dcac45c4726e344a037fe940dc4657ffade607d5755"></a>

<a id="canonical-2dbc5610415516b5ec201114202e771ebbd8815dabef9d13fa10b021a0a2ea9c"></a>

## flags property — primary.default_rr_set_group.caa_record.values / d210c9c599e3 / 4

Type: `"number"`. Computed.

Flag should be an integer between 0 and 255.

Upstream description:

This flag should be an integer between 0 and 255.

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

<a id="canonical-f49b5ef1937c2bd435a0ac6cc6f650f25811285367ecef5b5d72a957dc2f7bdc"></a>

<a id="canonical-c0074a8a5a184621c973563dd1408c0face7209e61e38d7b93ef4bdd2b78684a"></a>

## tag property — primary.default_rr_set_group.caa_record.values / d210c9c599e3 / 5

Type: `"string"`. Computed.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Upstream description:

Tag for categorization and filtering

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

<a id="canonical-f0a421166aa6c2881fc5ed0f3a46845747030d99221444526642aa3b7148f51e"></a>

<a id="canonical-a4bb89ffd3867df3a62ba3a60bdb9498c30bca7746d77f885ff0e6a722739a9f"></a>

## value property — primary.default_rr_set_group.caa_record.values / d210c9c599e3 / 6

Type: `"string"`. Computed.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

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

<a id="canonical-b6878151f09e78c0fc5a80a293bcd37499722325461b4233cf2f92c4bec9e288"></a>

## Next pages — primary.default_rr_set_group.caa_record.values / d210c9c599e3 / 7

- [primary.default_rr_set_group.caa_record](data-sources--dns_zone--reference--group-001.md#canonical-ad735989ad640d243f25f88ca63afcf0aff1a45cf0fd235536cbe0e0a15d4386)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-143fca8e1b3c7331722fcf357cb9aa76dfc70cbd1d3cb2b12cb2198d69d63e28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4acbcdf96b327a2d990c219103edb117e6e68037b8444c0f214f2640e8dfbf1"></a>

## primary.default_rr_set_group.cds_record — primary.default_rr_set_group.cds_record / 4636200a2c41 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.cds_record

<a id="canonical-a472e0a669fefebecdab8c5da4362e3e73a7910089f88605db2e7a152b10e7ed"></a>

Type: `"single"`. Computed.

DNS CDS Record. DNS CDS Record.

Upstream description:

DNS CDS Record.

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

<a id="canonical-fd7bcaef0229bb6f202ccb187e36856c2b985f9dd91ff927168d62e6ecc24e58"></a>

## Direct properties — primary.default_rr_set_group.cds_record / 4636200a2c41 / 3

<a id="canonical-61bf656d76ac7455c0beaa1c23f0a76095783ee0bd5303eb8d14d524a3354a01"></a>

<a id="canonical-2fbece8fbc0e4659cd7ca460b36848e6ebf6ff83ff2f03e01e32578d31d2c817"></a>

## name property — primary.default_rr_set_group.cds_record / 4636200a2c41 / 4

Type: `"string"`. Computed.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

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

- [values](data-sources--dns_zone--reference--group-001.md#canonical-7832dca787a00ca1a7866ec5415861e0a1105896ccce9e91810929ca2d80cde7): complete subsection reference.

<a id="canonical-464833ea6bcf838c0d854e1a71ce50dc235c5f49689e8f4ce3ef57731a78cc4d"></a>

## Next pages — primary.default_rr_set_group.cds_record / 4636200a2c41 / 5

- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-7832dca787a00ca1a7866ec5415861e0a1105896ccce9e91810929ca2d80cde7)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-7832dca787a00ca1a7866ec5415861e0a1105896ccce9e91810929ca2d80cde7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0882d961c4a336c2b2c34bf73bda474532cdff64869e45bbea541124557f660d"></a>

## primary.default_rr_set_group.cds_record.values — primary.default_rr_set_group.cds_record.values / 08b9a9a0a3a0 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-143fca8e1b3c7331722fcf357cb9aa76dfc70cbd1d3cb2b12cb2198d69d63e28)
- primary.default_rr_set_group.cds_record.values

<a id="canonical-d3d49d288965a11d71cd95f35101d33caaa799741e8c8f8f32bf3da7d7e52537"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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

<a id="canonical-cf6ea97aa1950ff81e39c772e37a0220262e21fa104b7ef18230c1b5d4be67f8"></a>

## Direct properties — primary.default_rr_set_group.cds_record.values / 08b9a9a0a3a0 / 3

<a id="canonical-ef9dbe8465bb62417924ccb69c13575b6c68f9518ed1f8a1d9fbb68e108a8a75"></a>

<a id="canonical-77550ac67143509ba9fc52bd848aa4dd4a8aefe18bf9e74e7bc3dd62de6e82f1"></a>

## ds_key_algorithm property — primary.default_rr_set_group.cds_record.values / 08b9a9a0a3a0 / 4

Type: `"string"`. Computed.

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

<a id="canonical-470f40f92f3d77628c6949e68e1995cea824138172bcb7890bb233e618573248"></a>

<a id="canonical-2f342d60f8e6e330c90253eb94e2348e98d7e31e83b2b694f1e3ea2aaf433ad2"></a>

## key_tag property — primary.default_rr_set_group.cds_record.values / 08b9a9a0a3a0 / 5

Type: `"number"`. Computed.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

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

- [sha1_digest](data-sources--dns_zone--reference--group-001.md#canonical-9f2b219b28ecbc7afbb472f2e6176ac652a10f3afc3015b9978fe18fcb478415): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-001.md#canonical-ad71d4d11c68665ce50ec1831cad00fdeb9427387e1b73e7236a8bb0f1cf90f5): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-001.md#canonical-91e3ebbf703dc4596bd680b04970927dc721a92c4094ab54811f3978d9eaecbd): complete subsection reference.

<a id="canonical-4c22e4f3a9005f69a045dee53ee67f6ea75ac91f9e7d1c9f8fb1127c37e58ff6"></a>

## Next pages — primary.default_rr_set_group.cds_record.values / 08b9a9a0a3a0 / 6

- [primary.default_rr_set_group.cds_record.values.sha1_digest](data-sources--dns_zone--reference--group-001.md#canonical-9f2b219b28ecbc7afbb472f2e6176ac652a10f3afc3015b9978fe18fcb478415)
- [primary.default_rr_set_group.cds_record.values.sha256_digest](data-sources--dns_zone--reference--group-001.md#canonical-ad71d4d11c68665ce50ec1831cad00fdeb9427387e1b73e7236a8bb0f1cf90f5)
- [primary.default_rr_set_group.cds_record.values.sha384_digest](data-sources--dns_zone--reference--group-001.md#canonical-91e3ebbf703dc4596bd680b04970927dc721a92c4094ab54811f3978d9eaecbd)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-143fca8e1b3c7331722fcf357cb9aa76dfc70cbd1d3cb2b12cb2198d69d63e28)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-9f2b219b28ecbc7afbb472f2e6176ac652a10f3afc3015b9978fe18fcb478415"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5142073149d3e48fcf850595eca28246781c34d572f7dca26fc0ffaf4c7c16a7"></a>

## primary.default_rr_set_group.cds_record.values.sha1_digest — primary.default_rr_set_group.cds_record.values.sha1_digest / 0aef08e6db25 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-143fca8e1b3c7331722fcf357cb9aa76dfc70cbd1d3cb2b12cb2198d69d63e28)
- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-7832dca787a00ca1a7866ec5415861e0a1105896ccce9e91810929ca2d80cde7)
- primary.default_rr_set_group.cds_record.values.sha1_digest

<a id="canonical-932ce44e6439ffca23e08d6da83bcd62c25726cd1a0a76fb12479bf4e4b8d86c"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="canonical-1cea1ab301c3c41b7d15a04c5d5e4cc66ce52960033d5bedf33425bea9f47e31"></a>

## Direct properties — primary.default_rr_set_group.cds_record.values.sha1_digest / 0aef08e6db25 / 3

<a id="canonical-5875a5a41cbdfa4b62f740d1e0ef396aae8af46c461da69596453fea81568956"></a>

<a id="canonical-1934dbeee9e13cdd6376394851146f49cf593bbd27e869f8b86a1350950ed9c4"></a>

## digest property — primary.default_rr_set_group.cds_record.values.sha1_digest / 0aef08e6db25 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

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

<a id="canonical-0ad1a032f681eecc9c6e723ef0163375385716b198659a96ec8bd134b5bfd194"></a>

## Next pages — primary.default_rr_set_group.cds_record.values.sha1_digest / 0aef08e6db25 / 5

- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-7832dca787a00ca1a7866ec5415861e0a1105896ccce9e91810929ca2d80cde7)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-ad71d4d11c68665ce50ec1831cad00fdeb9427387e1b73e7236a8bb0f1cf90f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-438e619dc2006d5b589c29dd8359373b88fb26c568c7676aa241b653e8a44f6f"></a>

## primary.default_rr_set_group.cds_record.values.sha256_digest — primary.default_rr_set_group.cds_record.values.sha256_digest / aeb1c720df5f / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-143fca8e1b3c7331722fcf357cb9aa76dfc70cbd1d3cb2b12cb2198d69d63e28)
- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-7832dca787a00ca1a7866ec5415861e0a1105896ccce9e91810929ca2d80cde7)
- primary.default_rr_set_group.cds_record.values.sha256_digest

<a id="canonical-4ee507782954bdb72e61fcf62c04b6c981180dd977bc323a56a911eeee41d216"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="canonical-f6a0c115488dc979bfe1cb1ae9167360227688fa697db2abc98d23b9d7d30b90"></a>

## Direct properties — primary.default_rr_set_group.cds_record.values.sha256_digest / aeb1c720df5f / 3

<a id="canonical-3c525a1316bc6cb938031144311f589dd2733055002c2b374e1ea655e13af88a"></a>

<a id="canonical-5575be91cdbecf62683d7ac36d38df395021333443d30b5d8a8d66fe897c7077"></a>

## digest property — primary.default_rr_set_group.cds_record.values.sha256_digest / aeb1c720df5f / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

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

<a id="canonical-042347d85f7b7e43c2044c0431aa217693c32353fb9bbe236c23a9ae9c85bad4"></a>

## Next pages — primary.default_rr_set_group.cds_record.values.sha256_digest / aeb1c720df5f / 5

- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-7832dca787a00ca1a7866ec5415861e0a1105896ccce9e91810929ca2d80cde7)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-91e3ebbf703dc4596bd680b04970927dc721a92c4094ab54811f3978d9eaecbd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f10292c56c85cb02ffa49a9e9b79278def974c6560473fea55f636ed3e7e9d0"></a>

## primary.default_rr_set_group.cds_record.values.sha384_digest — primary.default_rr_set_group.cds_record.values.sha384_digest / 516cd9837077 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-143fca8e1b3c7331722fcf357cb9aa76dfc70cbd1d3cb2b12cb2198d69d63e28)
- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-7832dca787a00ca1a7866ec5415861e0a1105896ccce9e91810929ca2d80cde7)
- primary.default_rr_set_group.cds_record.values.sha384_digest

<a id="canonical-5519518722cabe2021c0f6af3fbf04850ba200b22a872d9dc177b9efc43f612e"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

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

<a id="canonical-6b5f2456b60bf245b37533b9fa751b2ad546f829d19d723a41603c8e00ab2748"></a>

## Direct properties — primary.default_rr_set_group.cds_record.values.sha384_digest / 516cd9837077 / 3

<a id="canonical-32e43cc11a18543059efa3337113743891a1c2dfdea74e5085fec72fa2358c93"></a>

<a id="canonical-8b47dac15efcace1e8efa3ea31bef4a8fdf6d50f2f517b250c247ac1f67e489c"></a>

## digest property — primary.default_rr_set_group.cds_record.values.sha384_digest / 516cd9837077 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

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

<a id="canonical-f1f7855bac062851eb7bf29e7411ac7690f42abe184101d47a019f61b923f4bb"></a>

## Next pages — primary.default_rr_set_group.cds_record.values.sha384_digest / 516cd9837077 / 5

- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-7832dca787a00ca1a7866ec5415861e0a1105896ccce9e91810929ca2d80cde7)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-935f7e4f217a261d45e534d82676c2a545a619156396b4c6e75881f61a0a9ae9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8a2311f914cc7dc0f54cc9f5df576a253de198cab1b08fdcf4afa05b520b153"></a>

## primary.default_rr_set_group.cert_record — primary.default_rr_set_group.cert_record / 989a53c2abb8 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.cert_record

<a id="canonical-bcb066de69b36d0f9f55f58fdd2ec7682d9314cfa5d7d47dfea798bdc35285f1"></a>

Type: `"single"`. Computed.

Configuration parameter for cert record.

Upstream description:

DNS CERT Record.

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

<a id="canonical-88b55c1abe0ea5bc4da3fc412f3adc679c014ed5a80ff46c4353c3f1e5a7714e"></a>

## Direct properties — primary.default_rr_set_group.cert_record / 989a53c2abb8 / 3

<a id="canonical-f6a3fa4de87213668250e22f61b2987680473eb4f5f0db3515745d27f69215ca"></a>

<a id="canonical-e7d4f916197ad73b794b863b0626d68b2abb4305059b80feea0d544ddeb63ebc"></a>

## name property — primary.default_rr_set_group.cert_record / 989a53c2abb8 / 4

Type: `"string"`. Computed.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

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

- [values](data-sources--dns_zone--reference--group-001.md#canonical-58358b3caa291d2f5438912df9d63aaf78584c9fee435ebb53496334987b2db8): complete subsection reference.

<a id="canonical-4bc5791c1e5150f11ba0d4fee0b2fb1b3cfc03e8d862db4a1bd8b842682cd6b5"></a>

## Next pages — primary.default_rr_set_group.cert_record / 989a53c2abb8 / 5

- [primary.default_rr_set_group.cert_record.values](data-sources--dns_zone--reference--group-001.md#canonical-58358b3caa291d2f5438912df9d63aaf78584c9fee435ebb53496334987b2db8)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-58358b3caa291d2f5438912df9d63aaf78584c9fee435ebb53496334987b2db8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f6b5b2719fdd578058f3637aa8f5e5e30b7f76de0e29f7c042b9e2860ae5f88"></a>

## primary.default_rr_set_group.cert_record.values — primary.default_rr_set_group.cert_record.values / 3de32c47b893 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.cert_record](data-sources--dns_zone--reference--group-001.md#canonical-935f7e4f217a261d45e534d82676c2a545a619156396b4c6e75881f61a0a9ae9)
- primary.default_rr_set_group.cert_record.values

<a id="canonical-b06c6a188f9b294b16adc441378975b81b28554f356aff00553754aee09ba73f"></a>

Type: `"list"`. Computed.

CERT Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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

<a id="canonical-96acef64b0728bea05b960c3ed94e14cb20d2569f3f4b96e069c077aedeab360"></a>

## Direct properties — primary.default_rr_set_group.cert_record.values / 3de32c47b893 / 3

<a id="canonical-944dbdeadf1752dc7ceda15a10380b9f9a97564d6d62001ecdb31e44bbdc0996"></a>

<a id="canonical-fb7b55b17288deac640e7abe9dcacb8892bbe924cc66c72135304f7a12b01503"></a>

## algorithm property — primary.default_rr_set_group.cert_record.values / 3de32c47b893 / 4

Type: `"string"`. Computed.

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

<a id="canonical-a0c2c76d7f2c8b2f1f20ae4ccf4ef116402a5231119900a0444101dc501c01cd"></a>

<a id="canonical-3deddcdc22ba378ab9c5f5ddc186e44c5089e31893715b8d797cb4c5df667b71"></a>

## cert_key_tag property — primary.default_rr_set_group.cert_record.values / 3de32c47b893 / 5

Type: `"number"`. Computed.

Key Tag. Tag for categorization and filtering

Upstream description:

Tag for categorization and filtering

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

<a id="canonical-42811c8d9e9af9117543fb87bf81f96eebcc3c8fb356753a9af29aeb6aa26d45"></a>

<a id="canonical-cfdc05279d14213fc040e5f5845c3c32a80f20abcf539df7c3683d044b6c9cb1"></a>

## cert_type property — primary.default_rr_set_group.cert_record.values / 3de32c47b893 / 6

Type: `"string"`. Computed.

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

<a id="canonical-dae5ffff77d4be74941bb46e78137f1ae39452fd51058f89145b56b8f5d12975"></a>

<a id="canonical-79a7e17b4160ba11b24a997b3ca7ecd4a4ccdfdc9e45a16fd79b715c399cfcc2"></a>

## certificate property — primary.default_rr_set_group.cert_record.values / 3de32c47b893 / 7

Type: `"string"`. Computed.

Certificate. Certificate in base 64 format.

Upstream description:

Certificate in base 64 format.

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

<a id="canonical-feb0646c3c2bcf2cbdefb132ed3501f5bdcf947a0f400988657cbe54ae13379b"></a>

## Next pages — primary.default_rr_set_group.cert_record.values / 3de32c47b893 / 8

- [primary.default_rr_set_group.cert_record](data-sources--dns_zone--reference--group-001.md#canonical-935f7e4f217a261d45e534d82676c2a545a619156396b4c6e75881f61a0a9ae9)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-64ab7b8bf27978045203dbaa7158071b3ec5c7f6164268a2cda56ee8c8f8d6a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef7e59a09d04ea53329cd5a0f7cf9e09bcbfa879bfa75f0ec52836f11e401acf"></a>

## primary.default_rr_set_group.cname_record — primary.default_rr_set_group.cname_record / b4d9e1a207ab / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.cname_record

<a id="canonical-99c55d06815c7205b12930a31025f9cd8ca0445f61e6b82d81e531eb8d772219"></a>

Type: `"single"`. Computed.

DNSCNAMEResourceRecord.

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

<a id="canonical-b9d256ad96c280091356238d7b511d2856993e19c8c095ace181560ad34fc38b"></a>

## Direct properties — primary.default_rr_set_group.cname_record / b4d9e1a207ab / 3

<a id="canonical-c85c38aafe6406022f0e1ec9cd70304e372a6ba55f51e9a21bcc5937ddd63ac4"></a>

<a id="canonical-6594d2d3e90b0fd5d96d0a20a94fbb7ef105ff1a0f7c61f43b0a6a7064fbc076"></a>

## name property — primary.default_rr_set_group.cname_record / b4d9e1a207ab / 4

Type: `"string"`. Computed.

CName Record name, please provide only the specific subdomain or record name without the base
domain.

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

<a id="canonical-748c70ce8f42c7624739b3872b69bfd873bb0672551f68b6c9cba530e60d4d90"></a>

<a id="canonical-83ab9be17a45ae2ee8009716cd2231e42c52fcd0a57de9b68f21e8f1bd0d72c8"></a>

## value property — primary.default_rr_set_group.cname_record / b4d9e1a207ab / 5

Type: `"string"`. Computed.

Domain. Configuration parameter for value

Upstream description:

Configuration parameter for value

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

<a id="canonical-679f302079bdc25886ce05bfab968cdc1705060ccb2e2cb29289277d631166a5"></a>

## Next pages — primary.default_rr_set_group.cname_record / b4d9e1a207ab / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-3ef88f0643cae6c90abb5fde17697e969b3f9536dffdfc6919f32a93e600fe90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df68db21c1c6a20cba56db70312f7c11932fbdd05d0dd3c351e6959aafd44dd9"></a>

## primary.default_rr_set_group.ds_record — primary.default_rr_set_group.ds_record / e7bba35f48af / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.ds_record

<a id="canonical-658dc75819079df5add24a0c1c69b129caffda803dc0d041a833ae783815a7dd"></a>

Type: `"single"`. Computed.

DNS DS Record. DNS DS Record.

Upstream description:

DNS DS Record.

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

<a id="canonical-58a6b5c27d3132bfc6a3cdb49648a929c34f811ada6d80417efdbadec878bfc8"></a>

## Direct properties — primary.default_rr_set_group.ds_record / e7bba35f48af / 3

<a id="canonical-a45d0972cb1f990aff0214b71e231c3ee7f323ac0edca35bbf6c9f6c28facb79"></a>

<a id="canonical-b3094a8c5f2a393aed7d36e10b331f6d4b1b852f9ef068cbe2e6352a33ada96d"></a>

## name property — primary.default_rr_set_group.ds_record / e7bba35f48af / 4

Type: `"string"`. Computed.

DS Record name, please provide only the specific subdomain or record name without the base domain.

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

- [values](data-sources--dns_zone--reference--group-001.md#canonical-5ffdc52aab4896b597f8f36c508520f2da9d29ab4e4d129880080ae08be4b810): complete subsection reference.

<a id="canonical-a9804986ebb065c6e9c330f47fa3a691ca48848a3a0771170b322339b26831d2"></a>

## Next pages — primary.default_rr_set_group.ds_record / e7bba35f48af / 5

- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-5ffdc52aab4896b597f8f36c508520f2da9d29ab4e4d129880080ae08be4b810)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-5ffdc52aab4896b597f8f36c508520f2da9d29ab4e4d129880080ae08be4b810"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd9272249db737b60d84318bcb7149ebabd9f858a8124a160e83256872b09040"></a>

## primary.default_rr_set_group.ds_record.values — primary.default_rr_set_group.ds_record.values / 2f2fffe1e179 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-3ef88f0643cae6c90abb5fde17697e969b3f9536dffdfc6919f32a93e600fe90)
- primary.default_rr_set_group.ds_record.values

<a id="canonical-80485afca70502b050a6561537dacab323366da90417a5269d140391d3c66994"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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

<a id="canonical-bb738b59d3e3cf3e35258d2e96f86318677da21564c36ac5e24c8a9460189022"></a>

## Direct properties — primary.default_rr_set_group.ds_record.values / 2f2fffe1e179 / 3

<a id="canonical-1a9bf10795c490f393b3b3f0ffee8f3e99efe16e35decd808e98937a269a8da7"></a>

<a id="canonical-72acf1de3ea5761b20fb2459734d02f6bd60cd5b449e2018a1ef93833287feed"></a>

## ds_key_algorithm property — primary.default_rr_set_group.ds_record.values / 2f2fffe1e179 / 4

Type: `"string"`. Computed.

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

<a id="canonical-b28f44e5ab4287ae4e21a6faf57c28d9360f7538ab54c02273c07af72c77ea97"></a>

<a id="canonical-1a9b6c685609f033248d91abb0430924edb3d2bba1cd899a2d9c800906b20f27"></a>

## key_tag property — primary.default_rr_set_group.ds_record.values / 2f2fffe1e179 / 5

Type: `"number"`. Computed.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

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

- [sha1_digest](data-sources--dns_zone--reference--group-001.md#canonical-bc9655fb1fe00f7ce65518db68f05baf70e2b4aae5d607244e9c432c6a3229b7): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-001.md#canonical-4808330e5c84e58dc4f70e68f84d9c303de77cfc74b6d1c778a54fd4a226408f): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-8abbf950198894d9b203f18a6e8391e221e5971bf4d7c89b8a0285ee150bf7b7): complete subsection reference.

<a id="canonical-d741661f308e12d90ef4e65ab19384ea937c55d09a4b79cf74ab2a92c14bcb69"></a>

## Next pages — primary.default_rr_set_group.ds_record.values / 2f2fffe1e179 / 6

- [primary.default_rr_set_group.ds_record.values.sha1_digest](data-sources--dns_zone--reference--group-001.md#canonical-bc9655fb1fe00f7ce65518db68f05baf70e2b4aae5d607244e9c432c6a3229b7)
- [primary.default_rr_set_group.ds_record.values.sha256_digest](data-sources--dns_zone--reference--group-001.md#canonical-4808330e5c84e58dc4f70e68f84d9c303de77cfc74b6d1c778a54fd4a226408f)
- [primary.default_rr_set_group.ds_record.values.sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-8abbf950198894d9b203f18a6e8391e221e5971bf4d7c89b8a0285ee150bf7b7)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-3ef88f0643cae6c90abb5fde17697e969b3f9536dffdfc6919f32a93e600fe90)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-bc9655fb1fe00f7ce65518db68f05baf70e2b4aae5d607244e9c432c6a3229b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a590b30201ee4ba90a5dad4b35545c3c1e3125ff4f1791a34b819ee492f0ff27"></a>

## primary.default_rr_set_group.ds_record.values.sha1_digest — primary.default_rr_set_group.ds_record.values.sha1_digest / cf2f450a353e / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-3ef88f0643cae6c90abb5fde17697e969b3f9536dffdfc6919f32a93e600fe90)
- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-5ffdc52aab4896b597f8f36c508520f2da9d29ab4e4d129880080ae08be4b810)
- primary.default_rr_set_group.ds_record.values.sha1_digest

<a id="canonical-6b58cd1b299321dd72348ac32e4a41d836c4cbecc578256c855b9a7d8db10203"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="canonical-29ae9fb7e06d30969b9797efcdcbf3dff0220af0844075004c026a182bb944f4"></a>

## Direct properties — primary.default_rr_set_group.ds_record.values.sha1_digest / cf2f450a353e / 3

<a id="canonical-8b653e94d37b371abbb6712328b5c412d166b37bec2801e6ec9750cb1f8df487"></a>

<a id="canonical-edeb8c632364fbf5b7ba25a2ae0f732d324c1d6369ee3a91ea94e6dc4f880d88"></a>

## digest property — primary.default_rr_set_group.ds_record.values.sha1_digest / cf2f450a353e / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

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

<a id="canonical-6567bd3d858421137197b1eb7f38fa2de5f685a8a729b71af1f9bbdbb9ae9ac4"></a>

## Next pages — primary.default_rr_set_group.ds_record.values.sha1_digest / cf2f450a353e / 5

- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-5ffdc52aab4896b597f8f36c508520f2da9d29ab4e4d129880080ae08be4b810)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-4808330e5c84e58dc4f70e68f84d9c303de77cfc74b6d1c778a54fd4a226408f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
