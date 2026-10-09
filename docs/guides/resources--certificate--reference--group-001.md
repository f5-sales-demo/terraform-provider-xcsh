---
page_title: "xcsh_certificate reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate reference."
---

# xcsh_certificate reference

<a id="canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- Property reference

<a id="canonical-1032333300302201-3220332313112200-3003102333023113-0030202022003231-2102001101333333-1332121201202201-3033300012210132-1311321222133201"></a>

### Direct properties for `xcsh_certificate`

<a id="canonical-0320220311103002-1100000320010211-0023313322031110-0211201123101220-0102012203201200-2032211233103030-1330132322012212-3322131010100202"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

- [blindfold](resources--certificate--reference--group-001.md#canonical-1133330130112321-1330211011221231-3112320213031213-1000313300133122-1013102303022130-2300232200213211-0121031230120100-0201032002231233): complete subsection reference.

- [certificate_chain](resources--certificate--reference--group-001.md#canonical-1030320000233232-2232302213110323-3310002333011103-3020003203301303-3332202000210013-1213222102122030-2002010110130123-2123301311203302): complete subsection reference.

<a id="canonical-0120232130213103-0001203133221331-2100132201133130-3333330022201311-3300233010333110-1031203021302101-2331022012031000-3303231032102113"></a>

<a id="canonical-1320122003230231-2020100331123012-1003002022320313-3221013022023311-3203010210121023-2300113220110111-2111132022232301-2131210031111312"></a>

#### `certificate_url` property

Type: `"string"`. Optional, Computed.

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-3131031131220113-3330221330321201-1023222220022220-0312222333120011-2110302312132303-2122013133211112-3110020100032323-3230021122333211): complete subsection reference.

<a id="canonical-2100113220220032-0013033300222122-1101030123020132-2012311302022211-0311012202232331-3022232331120112-0210101101313312-0100320213313020"></a>

<a id="canonical-3210120012232312-2121330203222203-3112221010030003-2102130223200030-3223201113303133-2112203002310012-3333313113302203-1100010132330211"></a>

#### `description` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3302011201332201-1011130333020332-2000332013012220-3200032331010300-1330013110302301-1001133331223120-2023103331221011-0023313132321032"></a>

<a id="canonical-0303112333001220-3301203022322322-3103030222102130-1300222301231111-2210022323303013-0221103023111221-3032302111301000-2122000333020022"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

- [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-3232212033302331-1222133011022110-3212321322101213-0313001110013132-0330213212211122-0302330121213011-1313201320102231-0233302020223133): complete subsection reference.

<a id="canonical-0323000032312003-0203202331330222-3020123330112102-2113333223131110-1032223011103231-3220223110332311-2301003103222200-0122211113222131"></a>

<a id="canonical-3133302111000022-0232011011000000-1223321113333130-0122332210030301-1202202321013031-0010102021011133-3323331210123220-0100220310221301"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1031133332223333-1310130312111213-0332203121120233-2332033020233003-3100123323330031-2001322013023210-1023320203311003-2210322121100310"></a>

<a id="canonical-3221012323011333-2233222123232133-2231301323220010-0331121231212032-1000010120223211-0203033210021022-0113200010221331-1212003330313021"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

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

<a id="canonical-0133301113113232-1303121313102300-1010102233023201-0020120302311111-2332322122021030-2100103001331231-3331010222033103-0011100132212202"></a>

<a id="canonical-0223332303002030-2231323322022331-1211303322023132-0202222111000123-3212130310012211-2002102320121231-1133020212223200-3023233311003300"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Certificate. Must be unique within the namespace.

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0002231220003012-0111032313012131-0322312332223333-0022203321203111-3031323303133223-1132231231211330-0301020233320132-0332303332001313"></a>

<a id="canonical-1200222320122220-3103312202211220-0103322303110220-0200221230301300-2112033301300302-0301010200320203-2113321302123211-0032131230021022"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Certificate is created.

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [private_key](resources--certificate--reference--group-001.md#canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322): complete subsection reference.

- [timeouts](resources--certificate--reference--group-001.md#canonical-0022303111210311-0212221001121021-2111110133110332-1103310031101303-1020011013221332-0130113102021020-1110332113310011-1103211100200222): complete subsection reference.

- [use_system_defaults](resources--certificate--reference--group-001.md#canonical-1011132030030013-2320300233313232-2321133121131212-2123120301103000-1003011033212130-0213210031100131-0330120313310313-1100200323301322): complete subsection reference.

<a id="canonical-1201220031312011-2002202102010023-1213020031013303-1322200220320103-2323120203031110-2211130121101202-0020210320003122-2321332311022110"></a>

### All schema paths for `xcsh_certificate`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--certificate--reference--group-001.md#canonical-0320220311103002-1100000320010211-0023313322031110-0211201123101220-0102012203201200-2032211233103030-1330132322012212-3322131010100202) |
| `blindfold` | [blindfold](resources--certificate--reference--group-001.md#canonical-3031322122003022-1212000100100000-3210321203322000-1011231222333213-0222010032133000-2023222002222220-2110302030131000-2000232302110233) |
| `blindfold.algorithm` | [blindfold.algorithm](resources--certificate--reference--group-001.md#canonical-0200012230203110-1232233122302013-1301121332313321-0220221333211323-1021031011322030-1011110301131002-2210320220233023-0000321332201113) |
| `blindfold.certificate_file` | [blindfold.certificate_file](resources--certificate--reference--group-001.md#canonical-3010333233200030-3300002033021100-3222021001032200-0212121031301030-0330012321122321-1122333231100033-0031020303333212-1033012130221101) |
| `blindfold.certificate_pem` | [blindfold.certificate_pem](resources--certificate--reference--group-001.md#canonical-0001211331000003-3122013102130113-1312030022001322-2321312200111122-2002312211312321-0222113120101120-2000303010312111-2223113000200212) |
| `blindfold.chain_identity` | [blindfold.chain_identity](resources--certificate--reference--group-001.md#canonical-0111120021211033-0131200021232200-2321322131032333-0101021110000221-0230121313023322-3031232132220312-2010220002300301-1103030332023221) |
| `blindfold.context_digest` | [blindfold.context_digest](resources--certificate--reference--group-001.md#canonical-1033322111213313-1033020013102001-0212121222103011-0210113221020112-1012323310201311-0132020122000132-3203010330313101-0211221231001231) |
| `blindfold.encrypted_location` | [blindfold.encrypted_location](resources--certificate--reference--group-001.md#canonical-3223123323002302-0302303202213311-3031120210122110-2121020110133000-2232321233133000-2330130210231012-3102321102020121-3213112123030011) |
| `blindfold.expires_at` | [blindfold.expires_at](resources--certificate--reference--group-001.md#canonical-3102223123200023-1212132233013032-1120133320132000-0111121100102301-2121310102123020-0203011200002321-3211223032002200-1032212133232101) |
| `blindfold.fingerprint` | [blindfold.fingerprint](resources--certificate--reference--group-001.md#canonical-3213103120312322-3032301220312112-1102120023003130-2302203133110012-0312111231213213-3301221312210230-0123032123121313-1231313131000021) |
| `blindfold.id` | [blindfold.id](resources--certificate--reference--group-001.md#canonical-1130132321321030-2202101100023331-2032000321321101-0232300321300300-3330032003330102-3121302003020013-3311000013312232-1010021002103310) |
| `blindfold.material_version` | [blindfold.material_version](resources--certificate--reference--group-001.md#canonical-2133323012322123-3313213322003321-3330002202120100-0213121013222301-3010322022203002-1232211333011021-3010233023032013-0231233133213011) |
| `blindfold.passphrase_env` | [blindfold.passphrase_env](resources--certificate--reference--group-001.md#canonical-3021022312232221-1331013132033311-0121120332110321-3210202203312010-0201333033111032-1202210100010320-3021022330011321-3312120203011010) |
| `blindfold.passphrase_wo` | [blindfold.passphrase_wo](resources--certificate--reference--group-001.md#canonical-1112313322103110-0133121212330012-2332323120100112-0000122030001113-0121032312200100-0313303323202022-2133112311220113-3301031220011011) |
| `blindfold.pkcs12_file` | [blindfold.pkcs12_file](resources--certificate--reference--group-001.md#canonical-1020203220320202-1202210223201332-3100002232103120-1323303023123323-1102212203111022-3313301001131022-1301223020310120-2012011201121312) |
| `blindfold.pkcs12_wo` | [blindfold.pkcs12_wo](resources--certificate--reference--group-001.md#canonical-2212022003303103-1013213103011120-2303211300301203-0310012100002001-3021332210211333-1112002100330223-1113033012211203-1323002003011311) |
| `blindfold.policy` | [blindfold.policy](resources--certificate--reference--group-001.md#canonical-2132123101023203-0111210023022303-3111303330231223-1220003012100112-3220310313320201-2320313030201112-1200100330230100-1022013231312113) |
| `blindfold.prepared_identity` | [blindfold.prepared_identity](resources--certificate--reference--group-001.md#canonical-3132001201103012-0303123331212303-3201233021332000-2023003033323223-2001311000000111-3230313023220302-2100312111013331-0132212203102002) |
| `blindfold.private_key_file` | [blindfold.private_key_file](resources--certificate--reference--group-001.md#canonical-1010030330201301-1133203213233220-3121120031011222-1012022002112231-0311013010333101-1020023311000111-3132330230333011-0002201111210010) |
| `blindfold.private_key_wo` | [blindfold.private_key_wo](resources--certificate--reference--group-001.md#canonical-1330213223011221-2222212021231110-2131113313213101-1023131303333300-3130233221200033-1100121221011311-1010300303332102-1000332002302010) |
| `blindfold.spki_identity` | [blindfold.spki_identity](resources--certificate--reference--group-001.md#canonical-3211101310302022-2333020112000122-0332112231301111-1113331221013220-3322011321021312-0313000330330111-1333333011211232-3130221200331013) |
| `certificate_chain` | [certificate_chain](resources--certificate--reference--group-001.md#canonical-1330301200113122-3202132011031313-2233013030020112-1010133233032132-1131231131000232-0131233310311330-1213323111312102-1011302321023122) |
| `certificate_chain.name` | [certificate_chain.name](resources--certificate--reference--group-001.md#canonical-3100102310320022-2001321103213123-0332112213002323-2221231200312220-0202301212320121-1230121322331213-1232132021203012-0121013303100313) |
| `certificate_chain.namespace` | [certificate_chain.namespace](resources--certificate--reference--group-001.md#canonical-3303011203122322-1221332120333011-0312101111221311-2002220331020300-0201203012303303-2112130020110000-1111013200311313-1111031212033212) |
| `certificate_chain.tenant` | [certificate_chain.tenant](resources--certificate--reference--group-001.md#canonical-3121330331022102-2013220212222130-3332322302230102-1103233332120203-1331031100021020-3330200011030302-0310300020331320-0203230220201332) |
| `certificate_url` | [certificate_url](resources--certificate--reference--group-001.md#canonical-0120232130213103-0001203133221331-2100132201133130-3333330022201311-3300233010333110-1031203021302101-2331022012031000-3303231032102113) |
| `custom_hash_algorithms` | [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-0332010030333003-3020030330020200-0323221203003331-1013103210232131-0222132321131330-1310101323302011-1300122110213332-2223123113131111) |
| `custom_hash_algorithms.hash_algorithms` | [custom_hash_algorithms.hash_algorithms](resources--certificate--reference--group-001.md#canonical-3110300333231130-3310010200131301-3023003021121333-3333230133221303-3131031032220301-3010012221131003-2013120100300213-0321023300310313) |
| `description` | [description](resources--certificate--reference--group-001.md#canonical-2100113220220032-0013033300222122-1101030123020132-2012311302022211-0311012202232331-3022232331120112-0210101101313312-0100320213313020) |
| `disable` | [disable](resources--certificate--reference--group-001.md#canonical-3302011201332201-1011130333020332-2000332013012220-3200032331010300-1330013110302301-1001133331223120-2023103331221011-0023313132321032) |
| `disable_ocsp_stapling` | [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-3112210023202301-0330301120233313-3300023132021120-3202331002133131-1010200312002010-3302112331202331-1120331210032120-1030211100333130) |
| `id` | [ID](resources--certificate--reference--group-001.md#canonical-0323000032312003-0203202331330222-3020123330112102-2113333223131110-1032223011103231-3220223110332311-2301003103222200-0122211113222131) |
| `labels` | [labels](resources--certificate--reference--group-001.md#canonical-1031133332223333-1310130312111213-0332203121120233-2332033020233003-3100123323330031-2001322013023210-1023320203311003-2210322121100310) |
| `name` | [name](resources--certificate--reference--group-001.md#canonical-0133301113113232-1303121313102300-1010102233023201-0020120302311111-2332322122021030-2100103001331231-3331010222033103-0011100132212202) |
| `namespace` | [namespace](resources--certificate--reference--group-001.md#canonical-0002231220003012-0111032313012131-0322312332223333-0022203321203111-3031323303133223-1132231231211330-0301020233320132-0332303332001313) |
| `private_key` | [private_key](resources--certificate--reference--group-001.md#canonical-0100130131110220-0011113301322303-1310022010313231-1132001300000030-2032203200130133-0102112301032301-1130003031303201-2001103211102333) |
| `private_key.blindfold_secret_info` | [private_key.blindfold_secret_info](resources--certificate--reference--group-001.md#canonical-2231111122010112-0121323323011231-2320123113013123-2320233213210120-2030020301002303-1112232200021020-2221102232222130-2120021011011330) |
| `private_key.blindfold_secret_info.decryption_provider` | [private_key.blindfold_secret_info.decryption_provider](resources--certificate--reference--group-001.md#canonical-2230002020113101-0010223013223001-0023222311313032-1001021123200303-0132202013121202-1031103013232231-2233211123202223-0231331231301021) |
| `private_key.blindfold_secret_info.location` | [private_key.blindfold_secret_info.location](resources--certificate--reference--group-001.md#canonical-1031022210300020-0331210101210101-0113301301311002-1320102133002130-3230212211213312-3221000011220122-3023312133112102-2313233121203232) |
| `private_key.blindfold_secret_info.store_provider` | [private_key.blindfold_secret_info.store_provider](resources--certificate--reference--group-001.md#canonical-1022312313300303-3330120320200312-1223210113301132-0010202320302300-1232311333132031-0001111002311220-3022333202321222-0020300021110202) |
| `private_key.clear_secret_info` | [private_key.clear_secret_info](resources--certificate--reference--group-001.md#canonical-0232100033021131-0200222301033031-1100332131030221-1133110310203222-1122211202113033-0303220221133302-0200212021332222-0210312222120202) |
| `private_key.clear_secret_info.provider_ref` | [private_key.clear_secret_info.provider_ref](resources--certificate--reference--group-001.md#canonical-1322302211223031-0333203031303322-1313121200221133-1121100002222013-3333002122121013-3231101302103311-0220323110133302-1311013332020313) |
| `private_key.clear_secret_info.url` | [private_key.clear_secret_info.url](resources--certificate--reference--group-001.md#canonical-3032023123012320-0111332113111321-2002131103110320-1213210132200100-0321200110221232-0212000300110103-0233111221211333-0002022300022330) |
| `timeouts` | [timeouts](resources--certificate--reference--group-001.md#canonical-0000303133321230-1012002213111110-0031031201201332-2222102332200222-3132331013210231-2301313033030213-0210202323230022-0130230330123212) |
| `timeouts.create` | [timeouts.create](resources--certificate--reference--group-001.md#canonical-2103123102303330-0320110110203203-0010002101020110-0122223020320211-0100313010001123-3122202321233220-0031021030311012-0022033101332000) |
| `timeouts.delete` | [timeouts.delete](resources--certificate--reference--group-001.md#canonical-1002220213101223-3320020313122203-2030320231313200-0303133310200030-0232223001021120-3010020033032301-3023333120220321-0310112011103333) |
| `timeouts.read` | [timeouts.read](resources--certificate--reference--group-001.md#canonical-1133311000020021-0311113100113302-1232321120310122-2330113330313303-2210311313202012-3232311300021212-3111222301201022-3323030232232123) |
| `timeouts.update` | [timeouts.update](resources--certificate--reference--group-001.md#canonical-1120011233003201-2000212210010112-0033011100022110-3231030033323133-3230321202231000-3233032332223033-2102210201101110-0231222103012330) |
| `use_system_defaults` | [use_system_defaults](resources--certificate--reference--group-001.md#canonical-3101230030230231-1320130120122112-2020012023230002-2003032110301202-1121311100130003-0001321111002231-2003020200001100-2312032130010111) |

<a id="canonical-1133330130112321-1330211011221231-3112320213031213-1000313300133122-1013102303022130-2300232200213211-0121031230120100-0201032002231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blindfold` properties

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- blindfold

<a id="canonical-3031322122003022-1212000100100000-3210321203322000-1011231222333213-0222010032133000-2023222002222220-2110302030131000-2000232302110233"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-3233131132211221-1213131331020133-3120301011221003-1033332203033101-3322200020300332-3222021330231021-3213122221120112-3132023013030002"></a>

### Direct properties for `blindfold`

<a id="canonical-0200012230203110-1232233122302013-1301121332313321-0220221333211323-1021031011322030-1011110301131002-2210320220233023-0000321332201113"></a>

#### `blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-3010333233200030-3300002033021100-3222021001032200-0212121031301030-0330012321122321-1122333231100033-0031020303333212-1033012130221101"></a>

<a id="canonical-1111101302223323-3031210130232122-3131033010003032-1213030113131001-2321333331223102-3000020112100131-2222130131110221-0100222221012303"></a>

#### `blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-0001211331000003-3122013102130113-1312030022001322-2321312200111122-2002312211312321-0222113120101120-2000303010312111-2223113000200212"></a>

<a id="canonical-3303133221100103-2230120130022213-1131322331300302-0200221130223121-0322023002230222-2310123100111010-1233102213011331-2120033033233121"></a>

#### `blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-0111120021211033-0131200021232200-2321322131032333-0101021110000221-0230121313023322-3031232132220312-2010220002300301-1103030332023221"></a>

<a id="canonical-2030020230122113-1201013120010123-1211302013332001-0011003003103013-1123020111203302-0303323113131033-3100132030013001-1022331112103320"></a>

#### `blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-1033322111213313-1033020013102001-0212121222103011-0210113221020112-1012323310201311-0132020122000132-3203010330313101-0211221231001231"></a>

<a id="canonical-1221300331003302-2031223012213300-1032323001311133-3100211233032110-3122333001233330-3121310332122123-2002210221133132-3101000030233330"></a>

#### `blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-3223123323002302-0302303202213311-3031120210122110-2121020110133000-2232321233133000-2330130210231012-3102321102020121-3213112123030011"></a>

<a id="canonical-2023000323000210-1231330212113210-3012312111010221-3311132323330310-2332333133311123-1020300213313223-0031333323233022-1232020310301113"></a>

#### `blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-3102223123200023-1212132233013032-1120133320132000-0111121100102301-2121310102123020-0203011200002321-3211223032002200-1032212133232101"></a>

<a id="canonical-2110212033203031-0322233302231222-3331310220003212-3321212000023000-1020023131023201-0201231012122231-3211011323110200-0203120111011033"></a>

#### `blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-3213103120312322-3032301220312112-1102120023003130-2302203133110012-0312111231213213-3301221312210230-0123032123121313-1231313131000021"></a>

<a id="canonical-1110002012011331-3220030231133001-1213210220303331-3113110221320011-2330132122022233-2032303121223023-0333022101133321-1012331223033011"></a>

#### `blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-1130132321321030-2202101100023331-2032000321321101-0232300321300300-3330032003330102-3121302003020013-3311000013312232-1010021002103310"></a>

<a id="canonical-2030021022221132-1200231112201032-0001220331010001-1133300120213130-0330302331300033-3011311132033023-3233312121013222-1213300021120212"></a>

#### `blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-2133323012322123-3313213322003321-3330002202120100-0213121013222301-3010322022203002-1232211333011021-3010233023032013-0231233133213011"></a>

<a id="canonical-0133331302220321-3202313220332330-0001012200222211-2302123232203130-1030112031103231-0000021020332113-3333320120302330-1022231030302101"></a>

#### `blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-3021022312232221-1331013132033311-0121120332110321-3210202203312010-0201333033111032-1202210100010320-3021022330011321-3312120203011010"></a>

<a id="canonical-0013200113300321-1122323231232201-2023301131031113-1330111213133211-3331102211112233-1120301200101210-3020022231122313-1300223110232121"></a>

#### `blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-1112313322103110-0133121212330012-2332323120100112-0000122030001113-0121032312200100-0313303323202022-2133112311220113-3301031220011011"></a>

<a id="canonical-0310213111312311-1200313223032221-1333201022003110-2303112102020130-0210213130120011-1120121133023313-2330022132122002-1010311210230132"></a>

#### `blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1020203220320202-1202210223201332-3100002232103120-1323303023123323-1102212203111022-3313301001131022-1301223020310120-2012011201121312"></a>

<a id="canonical-3320300122313103-0331101021223332-2330133230231001-0332233101130033-1100332002231331-2203122321230222-3133230133330331-1210330230010203"></a>

#### `blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-2212022003303103-1013213103011120-2303211300301203-0310012100002001-3021332210211333-1112002100330223-1113033012211203-1323002003011311"></a>

<a id="canonical-0211302001111110-1300023121033311-2202221013331320-0130232012112211-1233213220110013-0020033111212132-0203203222312110-0222132112330002"></a>

#### `blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2132123101023203-0111210023022303-3111303330231223-1220003012100112-3220310313320201-2320313030201112-1200100330230100-1022013231312113"></a>

<a id="canonical-2322113030002103-0330131032002303-2033030012112303-2123001132121013-0133211023011012-0232332211222222-1202022021130201-0103303021230213"></a>

#### `blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-3132001201103012-0303123331212303-3201233021332000-2023003033323223-2001311000000111-3230313023220302-2100312111013331-0132212203102002"></a>

<a id="canonical-0321022213132102-0210230322211133-3120202301311110-3111303003130223-0003022211221023-3331332232310033-0122221200313300-3111111301311322"></a>

#### `blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-1010030330201301-1133203213233220-3121120031011222-1012022002112231-0311013010333101-1020023311000111-3132330230333011-0002201111210010"></a>

<a id="canonical-2321111000122101-1000332201102022-0301330032111132-0211010323303322-1010320230322122-2131210011031103-1233033130303302-0222100110211321"></a>

#### `blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-1330213223011221-2222212021231110-2131113313213101-1023131303333300-3130233221200033-1100121221011311-1010300303332102-1000332002302010"></a>

<a id="canonical-0320311123333133-3320323323101121-0013100303023312-3122213003011222-0310022211301203-2113230003130212-2010002001310212-3033133130120311"></a>

#### `blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3211101310302022-2333020112000122-0332112231301111-1113331221013220-3322011321021312-0313000330330111-1333333011211232-3130221200331013"></a>

<a id="canonical-0310323021013131-1310131111202311-0213330231230210-2103323313112201-3010011333211121-2203122003000222-1230131023333031-3101120101333221"></a>

#### `blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-1030320000233232-2232302213110323-3310002333011103-3020003203301303-3332202000210013-1213222102122030-2002010110130123-2123301311203302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `certificate_chain` properties

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- certificate_chain

<a id="canonical-1330301200113122-3202132011031313-2233013030020112-1010133233032132-1131231131000232-0131233310311330-1213323111312102-1011302321023122"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
certificate_chain {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010012321132012-1201100030230020-1201320202121111-3111111021122300-0022221110312023-2321031330200203-1132122123201213-0331000000023300"></a>

### Direct properties for `certificate_chain`

<a id="canonical-3100102310320022-2001321103213123-0332112213002323-2221231200312220-0202301212320121-1230121322331213-1232132021203012-0121013303100313"></a>

#### `certificate_chain.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3303011203122322-1221332120333011-0312101111221311-2002220331020300-0201203012303303-2112130020110000-1111013200311313-1111031212033212"></a>

<a id="canonical-2330022201323033-1011132333211021-1221021321101302-0313001221211311-3133133331130210-0322231231231111-1030101102233112-2003002021302332"></a>

#### `certificate_chain.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3121330331022102-2013220212222130-3332322302230102-1103233332120203-1331031100021020-3330200011030302-0310300020331320-0203230220201332"></a>

<a id="canonical-0321132231303210-1322303130323020-2112033201320133-1002022330012101-1232012112300320-0302323100233231-1222303322002130-3300303221203102"></a>

#### `certificate_chain.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3131031131220113-3330221330321201-1023222220022220-0312222333120011-2110302312132303-2122013133211112-3110020100032323-3230021122333211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- custom_hash_algorithms

<a id="canonical-0332010030333003-3020030330020200-0323221203003331-1013103210232131-0222132321131330-1310101323302011-1300122110213332-2223123113131111"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_hash\_algorithms, disable\_ocsp\_stapling, use\_system\_defaults; Default:
use\_system\_defaults\] Specifies the hash algorithms to be used.

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

- [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-0332010030333003-3020030330020200-0323221203003331-1013103210232131-0222132321131330-1310101323302011-1300122110213332-2223123113131111)
- [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-3112210023202301-0330301120233313-3300023132021120-3202331002133131-1010200312002010-3302112331202331-1120331210032120-1030211100333130)
- [use_system_defaults](resources--certificate--reference--group-001.md#canonical-3101230030230231-1320130120122112-2020012023230002-2003032110301202-1121311100130003-0001321111002231-2003020200001100-2312032130010111)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130122011202323-1101200223301122-1302232231202330-3311302201102113-2012101121333332-2332110213001201-3032003023002023-3100322222102230"></a>

### Direct properties for `custom_hash_algorithms`

<a id="canonical-3110300333231130-3310010200131301-3023003021121333-3333230133221303-3131031032220301-3010012221131003-2013120100300213-0321023300310313"></a>

#### `custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3232212033302331-1222133011022110-3212321322101213-0313001110013132-0330213212211122-0302330121213011-1313201320102231-0233302020223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- disable_ocsp_stapling

<a id="canonical-3112210023202301-0330301120233313-3300023132021120-3202331002133131-1010200312002010-3302112331202331-1120331210032120-1030211100333130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

Additional upstream details:

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
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_key` properties

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- private_key

<a id="canonical-0100130131110220-0011113301322303-1310022010313231-1132001300000030-2032203200130133-0102112301032301-1130003031303201-2001103211102333"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323221230032120-1132011311110302-0131310032220310-2321102000211103-2221333120222211-2120211333031313-1233322323303012-3310233203031023"></a>

### Direct properties for `private_key`

- [blindfold_secret_info](resources--certificate--reference--group-001.md#canonical-1212130133113120-3213313012221220-1110001201213313-2131112331202231-2020230313231212-0330310130022011-1120030320133221-3001003323101233): complete subsection reference.

- [clear_secret_info](resources--certificate--reference--group-001.md#canonical-3002130021002013-0003231011201212-2110203232203110-2331222133010022-1020311000320201-2113111203202220-1313300012200102-2201233122303102): complete subsection reference.

<a id="canonical-1212130133113120-3213313012221220-1110001201213313-2131112331202231-2020230313231212-0330310130022011-1120030320133221-3001003323101233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [private_key](resources--certificate--reference--group-001.md#canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322)
- private_key.blindfold_secret_info

<a id="canonical-2231111122010112-0121323323011231-2320123113013123-2320233213210120-2030020301002303-1112232200021020-2221102232222130-2120021011011330"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-1210330131331010-1021111132121230-1312013010021002-2233111231132021-3233013012311302-1110001203311010-2213103110012212-3022311232312032"></a>

### Direct properties for `private_key.blindfold_secret_info`

<a id="canonical-2230002020113101-0010223013223001-0023222311313032-1001021123200303-0132202013121202-1031103013232231-2233211123202223-0231331231301021"></a>

#### `private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1031022210300020-0331210101210101-0113301301311002-1320102133002130-3230212211213312-3221000011220122-3023312133112102-2313233121203232"></a>

<a id="canonical-1203233002013110-1210101013321003-1200232122111201-1001133030130122-0013302000201221-2122131023020231-2331223131022100-0110310032310021"></a>

#### `private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1022312313300303-3330120320200312-1223210113301132-0010202320302300-1232311333132031-0001111002311220-3022333202321222-0020300021110202"></a>

<a id="canonical-2120312133210223-2132210312210310-1020330010221213-2320011100102000-2132120021111130-2102003222012103-0310213000331300-3210102233330320"></a>

#### `private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3002130021002013-0003231011201212-2110203232203110-2331222133010022-1020311000320201-2113111203202220-1313300012200102-2201233122303102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [private_key](resources--certificate--reference--group-001.md#canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322)
- private_key.clear_secret_info

<a id="canonical-0232100033021131-0200222301033031-1100332131030221-1133110310203222-1122211202113033-0303220221133302-0200212021332222-0210312222120202"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-1333321220201021-0131022001131101-0232203112211033-2111223320123220-0330132031033301-2223121232210023-3031202033001313-3000300012333230"></a>

### Direct properties for `private_key.clear_secret_info`

<a id="canonical-1322302211223031-0333203031303322-1313121200221133-1121100002222013-3333002122121013-3231101302103311-0220323110133302-1311013332020313"></a>

#### `private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3032023123012320-0111332113111321-2002131103110320-1213210132200100-0321200110221232-0212000300110103-0233111221211333-0002022300022330"></a>

<a id="canonical-1211023231131213-1102101103303311-3020231130330030-0102023200003120-2103333122123102-3013123012230013-1211012323010332-1300320113111221"></a>

#### `private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0022303111210311-0212221001121021-2111110133110332-1103310031101303-1020011013221332-0130113102021020-1110332113310011-1103211100200222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- timeouts

<a id="canonical-0000303133321230-1012002213111110-0031031201201332-2222102332200222-3132331013210231-2301313033030213-0210202323230022-0130230330123212"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233000100202132-1102032113000320-2031311002130031-3320320230002003-1211001233121001-0333230121101203-1000230121202110-0101211330001003"></a>

### Direct properties for `timeouts`

<a id="canonical-2103123102303330-0320110110203203-0010002101020110-0122223020320211-0100313010001123-3122202321233220-0031021030311012-0022033101332000"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1002220213101223-3320020313122203-2030320231313200-0303133310200030-0232223001021120-3010020033032301-3023333120220321-0310112011103333"></a>

<a id="canonical-2310131222131102-0220332022111332-3123220303212123-2130131310011222-0020303000113110-2310330023100203-3132331321200010-2111110130311313"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1133311000020021-0311113100113302-1232321120310122-2330113330313303-2210311313202012-3232311300021212-3111222301201022-3323030232232123"></a>

<a id="canonical-2133012230311210-0203133312033231-3213312231011023-3221213230222132-1000313333223130-2012001002330111-3102321223300310-3130113123111232"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1120011233003201-2000212210010112-0033011100022110-3231030033323133-3230321202231000-3233032332223033-2102210201101110-0231222103012330"></a>

<a id="canonical-3031111003133313-0330102333012101-3323232223210323-2213301320200302-1210120211102203-1020203222020023-3001302232332313-2212030133012312"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1011132030030013-2320300233313232-2321133121131212-2123120301103000-1003011033212130-0213210031100131-0330120313310313-1100200323301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_system_defaults` properties

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- use_system_defaults

<a id="canonical-3101230030230231-1320130120122112-2020012023230002-2003032110301202-1121311100130003-0001321111002231-2003020200001100-2312032130010111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

Additional upstream details:

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
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.
