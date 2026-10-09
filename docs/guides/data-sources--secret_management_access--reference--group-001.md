---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- Property reference

<a id="canonical-3210232011222331-3231121320001131-1331102131312302-2002321021120122-3011102202101012-0310331313333112-2331230012000330-3103023301000233"></a>

### Direct properties for `xcsh_secret_management_access`

- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133): complete subsection reference.

<a id="canonical-0020022133311021-2022112212231122-3002022022213000-2321300310310202-1303120203321303-1111311002032120-3232000222223133-3213020221113313"></a>

<a id="canonical-0101013131023012-1210103121021002-2211321010120112-0230033120303303-0322303122313133-2123010231123110-3232021021231211-0131321201301331"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

<a id="canonical-3122210100112130-1313303203321031-2030323222331321-0330210013020311-1031201211022211-0002022031330033-3031210210201132-3022033203223031"></a>

<a id="canonical-1100230113200012-3201210012331203-2111031102013322-1222233020103302-1101010122223001-2211020310023301-0121222322112100-1122213230003123"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the SecretManagementAccess.

Additional upstream details:

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

<a id="canonical-3300331133320113-3021133221213331-3322100223310211-0010012221331102-2232301032010230-3332203302330332-3233103133102310-2230101011212021"></a>

<a id="canonical-0331201101330010-1333023202023120-2120020111310333-3223032131113100-1112122323020012-2100021030033020-0310132010300323-1301103303320120"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1223222333112320-1303011301113202-0103111032130113-3023313003111001-2000231233330002-2211033023210132-3312110223232331-3222000022020112"></a>

<a id="canonical-3132021103322203-0313313332322330-0123110301111013-1122130322032112-1110120213212302-2020003012223133-2023103100223230-2132122200332300"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-0221012311032323-1211103132120012-0010211303010031-1103303112001001-2211232232020232-3023112313212212-0133010211003102-2322322313200312"></a>

<a id="canonical-2223101232302001-3010312002211323-2321330021000003-1020320323012011-3010222013330033-3121121112133321-1302201222021321-1001100220003302"></a>

#### `name` property

Type: `"string"`. Required.

Name of the SecretManagementAccess.

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

<a id="canonical-1211001101231101-2231021223312012-1231030031230121-1313210332331310-3113333033022033-2212221201122022-3230113201003223-3232311111031233"></a>

<a id="canonical-0122100132131010-0303331301103111-3221333323212122-0233113231112321-0200000033212223-3023313112000122-2230131120020130-0221012121102223"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the SecretManagementAccess exists.

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

<a id="canonical-1100332300303212-0302032103100023-2012103030312130-3011032122333203-3100102012022131-0101201213232200-1130000200313130-3111201031320322"></a>

<a id="canonical-2012121111212300-1032102130220211-1032210310100212-3300002331020120-1002010320221330-3310033130011110-1233223313211132-1010112203011132"></a>

#### `provider_name` property

Type: `"string"`. Computed.

Name given to this secret management backend. site.provider needs to be unique, and will be
referenced for using this object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101): complete subsection reference.

<a id="canonical-1302122102332100-1211222111002301-3332210310230100-1300310101110122-1033323022300312-1033202021301013-1013223121213230-1010131222101300"></a>

### All schema paths for `xcsh_secret_management_access`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `access_info` | [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1321022221213221-0020110111230331-3112013223203201-2022311211233302-2001110021330213-2003213320321322-2020012200211213-1312130320321012) |
| `access_info.rest_auth_info` | [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1320233201201101-1201330311212301-0222020003313211-2131223110112020-1312031130313222-1101110100012033-2122012312300030-3123123321000300) |
| `access_info.rest_auth_info.basic_auth` | [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2320002323003120-1231310201333211-0112330332133133-1030220201331310-0332320031312313-2220011333102303-0320110202320330-3230123122210202) |
| `access_info.rest_auth_info.basic_auth.password` | [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-2322323102101021-2030302123320230-0021232020212113-3001113300332000-0102320223201220-1321231331103111-0300230233313001-0002211210022320) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-2131220230312112-1003322331022133-0320120200303022-0033311223000011-0103223201312033-3233122010202100-3110310012103300-1023212011000021) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-001.md#canonical-3100331020123101-3220010230212110-1200302130233032-0032223323101331-1012002311110022-2321013100102001-1330302012213330-1101122023303313) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-001.md#canonical-2022311031003332-1232210332023001-2212321022110122-0111312121120133-0133313210213212-3233000200002232-3112331220110300-1201023130330332) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-001.md#canonical-2023022213122100-3300000332201231-0010322102300002-2212323213130331-3213103133200332-0133123212123131-2021010321311102-1303200011110322) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-1023032203302020-0003323100002000-3033120200333123-0031021010133130-1020123212223301-3210223310031102-2203032300212313-0211020123130123) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-001.md#canonical-0001101123000030-1003112211133130-0132330022013332-0322101012122110-0230230001001011-2002330302003120-1102002213213313-0032331101200101) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.url` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.url](data-sources--secret_management_access--reference--group-001.md#canonical-2030122203203233-3130301111300133-0101310302010032-2221033203333113-2322312021013221-2132213112021211-3310321010033232-3310121302213120) |
| `access_info.rest_auth_info.basic_auth.username` | [access_info.rest_auth_info.basic_auth.username](data-sources--secret_management_access--reference--group-001.md#canonical-0233323213032201-1112203333110232-1333303130223032-3322300201133202-0200132013313100-2303233000222311-3121313330300213-1111232331033302) |
| `access_info.rest_auth_info.headers_auth` | [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2011112102113323-2031321013332332-0123111121020322-3131300221121021-1332231130310330-0100012011231013-2103321200201202-3103200010232102) |
| `access_info.rest_auth_info.headers_auth.headers` | [access_info.rest_auth_info.headers_auth.headers](data-sources--secret_management_access--reference--group-001.md#canonical-2023110200112120-1322333132030021-2302011231030012-0332010123220132-2202011213133132-0131122232320213-1231002221212303-2100312100011323) |
| `access_info.rest_auth_info.query_params_auth` | [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-0201033223131030-1130030212123032-1030221223120301-3110230101031233-1233223132303003-3311101000130133-2123300300120213-3331231023101002) |
| `access_info.rest_auth_info.query_params_auth.query_params` | [access_info.rest_auth_info.query_params_auth.query_params](data-sources--secret_management_access--reference--group-001.md#canonical-0011000122222030-3211231011101312-0020101233023130-3321232012021000-1233133100200231-1212212320303321-0023120202231221-0130232221222310) |
| `access_info.scheme` | [access_info.scheme](data-sources--secret_management_access--reference--group-001.md#canonical-3032302011210022-1121311233220103-1230322100222110-1120213120331101-1322101210231233-2201300303331003-0332130121010211-0012001232210210) |
| `access_info.server_endpoint` | [access_info.server_endpoint](data-sources--secret_management_access--reference--group-001.md#canonical-0002020123320112-1213103113213313-1120323212103202-1022311130212120-1232331032030321-2311312023033223-1001213132130302-0220003301203313) |
| `access_info.tls_config` | [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-2203201032131033-1021100022222100-2002003021221310-1210100033233311-1202033222032221-3313130122231312-3101012302133313-1220130232231033) |
| `access_info.tls_config.cert_params` | [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-2213122301012012-0000300100312320-2022211030311300-0123202030212023-3120031112120123-1102123201103330-3330203213321312-1233032133113222) |
| `access_info.tls_config.cert_params.certificates` | [access_info.tls_config.cert_params.certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2110212023202120-1133030112012123-2301222133112222-1312120020030213-3213211012212222-3230130233110221-0112022223301312-3001302300321021) |
| `access_info.tls_config.cert_params.certificates.kind` | [access_info.tls_config.cert_params.certificates.kind](data-sources--secret_management_access--reference--group-001.md#canonical-0113131011332303-1313131022031221-1201111202212301-1133301123101012-1321302321103321-0030312222012310-2313200301302330-3330312232122010) |
| `access_info.tls_config.cert_params.certificates.name` | [access_info.tls_config.cert_params.certificates.name](data-sources--secret_management_access--reference--group-001.md#canonical-1022031202133120-1322302013003031-2003001012010332-3112310011200111-0330030332231212-0210232222211113-2302111033210300-0010013333222210) |
| `access_info.tls_config.cert_params.certificates.namespace` | [access_info.tls_config.cert_params.certificates.namespace](data-sources--secret_management_access--reference--group-001.md#canonical-2302130233313022-2133013233300310-1131013231002002-2202111332020021-0020213010233313-1003110100302032-0000001223323003-2102023033110121) |
| `access_info.tls_config.cert_params.certificates.tenant` | [access_info.tls_config.cert_params.certificates.tenant](data-sources--secret_management_access--reference--group-001.md#canonical-2023133000010201-2111303032303220-0233002310101210-0121321000030200-1332301113111030-1222223133320102-0021002303120220-1333312201123320) |
| `access_info.tls_config.cert_params.certificates.uid` | [access_info.tls_config.cert_params.certificates.uid](data-sources--secret_management_access--reference--group-001.md#canonical-3301303130001033-0013331323000000-0333133132200133-3133122330212202-2002230031112100-3211101320001030-0132002103033100-1223120111130022) |
| `access_info.tls_config.cert_params.cipher_suites` | [access_info.tls_config.cert_params.cipher_suites](data-sources--secret_management_access--reference--group-001.md#canonical-1020030112031312-3000310203203331-2123232212001332-3123012331231322-0321300301231102-3232200303203103-1130020032322222-0232311032330011) |
| `access_info.tls_config.cert_params.maximum_protocol_version` | [access_info.tls_config.cert_params.maximum_protocol_version](data-sources--secret_management_access--reference--group-001.md#canonical-0300100113211332-0022233111021023-1320201111130102-1313100212303133-0112213123022120-2231213013300023-0221331032210132-0201102210123210) |
| `access_info.tls_config.cert_params.minimum_protocol_version` | [access_info.tls_config.cert_params.minimum_protocol_version](data-sources--secret_management_access--reference--group-001.md#canonical-0320013333312130-1002022111022312-0203133100100120-3211103002210202-1300023303030323-3112201111203130-3211011303300301-2311230301002202) |
| `access_info.tls_config.cert_params.skip_server_verification` | [access_info.tls_config.cert_params.skip_server_verification](data-sources--secret_management_access--reference--group-001.md#canonical-2132300231222233-1020221121030211-0221333203332203-0133131031330000-0021220122300002-2211020212333302-2201203313113320-2111020120112112) |
| `access_info.tls_config.cert_params.tls_validation_params` | [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-1313132013331301-1022311210223030-0000030232233231-2311012200110231-3112030323013302-2202023231323321-1021232311213211-3031303232220032) |
| `access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification` | [access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification](data-sources--secret_management_access--reference--group-001.md#canonical-1032310210233100-0330132203201222-1211102130203331-3302101230221012-3033202002131331-0233213211010101-2200002221001103-2131312311230032) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-3020202221322302-1010133021111101-3103213113030031-0132000033312122-3103320022330000-2321332000112212-3322003101333230-2213311112130033) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-0011333100021110-2032212333221023-1311001220020113-3233123213300202-2212012032332321-0213100100230230-2121221023120323-0130031120033110) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](data-sources--secret_management_access--reference--group-001.md#canonical-0312032301000210-3103022130322030-1131020112113000-2201023023031101-1312220100231132-2001301220220212-0123213233311301-2022021022103200) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](data-sources--secret_management_access--reference--group-001.md#canonical-1310220101131202-0100001130230321-3033330331112123-0110010232313002-2111130221301232-2223130212222132-1303311320021110-3001100201023331) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--secret_management_access--reference--group-001.md#canonical-2133113022223010-0120311021130301-3023000100231000-2302013312333203-3011023310223103-0213331332331122-0001300123103020-1130202310230030) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--secret_management_access--reference--group-001.md#canonical-3323100102202321-2332031020202201-1033230211333132-0202223111010212-0000033133323230-1331320332102002-0133022011303312-3130132101310112) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](data-sources--secret_management_access--reference--group-001.md#canonical-1303323033233032-0130311221000233-3331200300231221-3312303233112202-3032331121101312-0120320300021300-2320021211322220-2232100130333213) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url](data-sources--secret_management_access--reference--group-001.md#canonical-2113222321200120-1201111132130130-2203211320022103-2003211303320011-3003303023331320-2212032223221323-2100200010032030-0312121303231231) |
| `access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names` | [access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names](data-sources--secret_management_access--reference--group-001.md#canonical-0002302330330222-1130332300200213-3233100311021111-1310303212011023-3322031000132031-1213030103100122-0230103331330022-3330020110013000) |
| `access_info.tls_config.cert_params.volterra_trusted_ca` | [access_info.tls_config.cert_params.volterra_trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-0123112310323301-3202211211010010-2130211130210323-2300111201232131-1013230022010210-3130320332100012-1310133122322001-2332300223113233) |
| `access_info.tls_config.common_params` | [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-1103203122021022-3002220212010200-3131203100120021-0131222302231112-1332321003021333-2300130033220033-2000101002203311-3323102012002322) |
| `access_info.tls_config.common_params.cipher_suites` | [access_info.tls_config.common_params.cipher_suites](data-sources--secret_management_access--reference--group-001.md#canonical-3230322012213021-1023003033321213-0210012111120310-1122130203002221-3122000322021111-1131121112333101-3312310032013221-2130302230121112) |
| `access_info.tls_config.common_params.maximum_protocol_version` | [access_info.tls_config.common_params.maximum_protocol_version](data-sources--secret_management_access--reference--group-001.md#canonical-0323301101033133-2110003220020032-0323210222333220-2301200200102011-2100300232200330-3131330202323221-0010102311122101-3000111323233203) |
| `access_info.tls_config.common_params.minimum_protocol_version` | [access_info.tls_config.common_params.minimum_protocol_version](data-sources--secret_management_access--reference--group-001.md#canonical-1032323031121103-0130103301311111-1000020000100323-0111131013121201-1322113103323302-0010033132113202-0132223023130001-2001200233100231) |
| `access_info.tls_config.common_params.tls_certificates` | [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-3323302002301203-2303102011222212-1122111100230133-2202321212220011-1310301221313220-0000331231221313-0100112012032101-1222001020103231) |
| `access_info.tls_config.common_params.tls_certificates.certificate_url` | [access_info.tls_config.common_params.tls_certificates.certificate_url](data-sources--secret_management_access--reference--group-001.md#canonical-2013130102301020-0213010113000321-2230300211012101-2110210223313303-2023000313322312-0311331133320330-0023230222323330-1020220021110121) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](data-sources--secret_management_access--reference--group-001.md#canonical-0013322302301232-2233303112233202-0302030212220222-0220011320321230-0303110213013032-3232321311310311-0202230323333233-2100223133020003) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--secret_management_access--reference--group-001.md#canonical-2300022023333230-0230100001123230-1230013320210000-3002012002020020-1332231300100203-1100131011113001-3033021133132310-1023003132101212) |
| `access_info.tls_config.common_params.tls_certificates.description_spec` | [access_info.tls_config.common_params.tls_certificates.description_spec](data-sources--secret_management_access--reference--group-001.md#canonical-2022132113301312-1111232023232303-1221332103020123-2010332212030332-0201331120111013-3202300203130012-0221301200222100-3010023322312113) |
| `access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling` | [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](data-sources--secret_management_access--reference--group-001.md#canonical-2301023220100010-0320220122321110-2120203020003311-0333333011232010-1032313032002300-0333213112212312-2201323020322030-3013003101101301) |
| `access_info.tls_config.common_params.tls_certificates.private_key` | [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-0021302312122233-2112222102301000-1331223110333001-1132032013010223-1322101312101032-3010113211321221-0221313303202311-1312201132122122) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-2212010130332122-1303103322011222-3030132023131312-1203000212201000-3030300203130223-0021300131222110-1333023233012111-2212112212011303) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-001.md#canonical-1113110033102300-1331103000031232-0103332101200132-2021131213222110-2303321311002320-0200300312023001-0100103323332212-2312211312032103) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-001.md#canonical-1230022323200132-1022013011100221-0213131332003222-1222210111231231-3232323120212201-2121023300310233-3112012123213320-0022210330010003) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-001.md#canonical-0201211012022310-2123003310231001-2133203310301121-2302002012021202-0330133012333210-1331311221300031-0131313301221212-3110001301110100) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3102220123101120-0321332330321020-1003131113103011-1000002330103033-3112023132313032-1133302312112100-1321212232121322-1123221133302021) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-001.md#canonical-0120320202103310-0323220321223103-3201213102300112-3232123322132223-2011311220211113-0101000303130102-3032332110022232-0302211120323122) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--secret_management_access--reference--group-001.md#canonical-3321212121332100-2030312333120000-1202220312223003-1103203331001310-2233302333221111-2120033211230212-3321032233012231-2012003121120111) |
| `access_info.tls_config.common_params.tls_certificates.use_system_defaults` | [access_info.tls_config.common_params.tls_certificates.use_system_defaults](data-sources--secret_management_access--reference--group-001.md#canonical-0221030011001233-1220131011110221-2112300021012233-0301022111112111-3210301220202233-1213011122231203-0232001111032120-3121113230331330) |
| `access_info.tls_config.common_params.validation_params` | [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-2310202013033320-0302233011130222-0311210021331122-1031030033221210-1102101101331122-1132133133020033-2133011021210100-1130132330231221) |
| `access_info.tls_config.common_params.validation_params.skip_hostname_verification` | [access_info.tls_config.common_params.validation_params.skip_hostname_verification](data-sources--secret_management_access--reference--group-001.md#canonical-0022113231002332-1332022301001112-2033100011111012-2123301312132021-1013302032023130-1320101133032103-2313100130222202-3203100232111303) |
| `access_info.tls_config.common_params.validation_params.trusted_ca` | [access_info.tls_config.common_params.validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-0112021103031313-1303200332123020-3132123331232303-1031310120202220-3313010312100222-2212332231321021-0101012211000311-3310213102222103) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-1103231332101211-1333330113300120-1030012120031321-3131332330102033-2303022000311030-0022303111121023-3322203123211030-2200132332232030) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--secret_management_access--reference--group-001.md#canonical-1333313301023221-1223133121110202-0031220012233033-0333213030300200-0032002132113212-0100311003130312-3000122030212002-3222313231320210) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--secret_management_access--reference--group-001.md#canonical-0100311303002312-1131303113001211-3333032310033311-1131213110213100-2110222330212132-2002100133210220-2003021311020020-1333320203010101) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--secret_management_access--reference--group-001.md#canonical-0322313213003323-1202133300331332-2120113001121300-3213122231210231-3033202323111132-3233003100310321-0330301112021331-0312322310113332) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--secret_management_access--reference--group-001.md#canonical-3311313120212301-2011301223022332-1212202102131313-2332100103020212-1012320301300212-1131000210302212-3322022031002122-1012322202222212) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--secret_management_access--reference--group-001.md#canonical-0331333322023123-2021011310013103-3320302021130112-3131133200221212-0011102201210012-1312030022232232-2322000133232323-0022123102121210) |
| `access_info.tls_config.common_params.validation_params.trusted_ca_url` | [access_info.tls_config.common_params.validation_params.trusted_ca_url](data-sources--secret_management_access--reference--group-001.md#canonical-2133201221000213-0231311300330220-0112013100123012-1302032322100113-1320212032001130-3133002220211113-1021022203031020-0002231110012310) |
| `access_info.tls_config.common_params.validation_params.verify_subject_alt_names` | [access_info.tls_config.common_params.validation_params.verify_subject_alt_names](data-sources--secret_management_access--reference--group-001.md#canonical-1032230033231110-3201011101300223-3030011213110213-1220231101101113-0003320101122332-3031031132332220-3303300312101222-1232121231302003) |
| `access_info.tls_config.default_session_key_caching` | [access_info.tls_config.default_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-2131321032020123-0232231210002310-2031113022030210-3121310200233322-3333210211133131-0000310331313310-2330301011100132-3323211313212231) |
| `access_info.tls_config.disable_session_key_caching` | [access_info.tls_config.disable_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-2032010020020012-1233223300322003-1213013022322221-2202210101010333-3111213332330112-0220031223131002-2112031122020333-1212301310002223) |
| `access_info.tls_config.disable_sni` | [access_info.tls_config.disable_sni](data-sources--secret_management_access--reference--group-001.md#canonical-0000330111213131-3031223010221210-2130332300203033-1201323121001001-0120113220113311-2320210310313030-3310311321132301-0131032013233331) |
| `access_info.tls_config.max_session_keys` | [access_info.tls_config.max_session_keys](data-sources--secret_management_access--reference--group-001.md#canonical-2032013333213230-3300320223333220-2033233103202101-3220311331110103-3121132030122232-2210011212311302-1021003203020221-1323111130232100) |
| `access_info.tls_config.sni` | [access_info.tls_config.sni](data-sources--secret_management_access--reference--group-001.md#canonical-1111312120133212-1332221003133322-3302033021111130-0211222212122331-2122221013010223-3123110000122031-1322131030202210-3103220221030012) |
| `access_info.tls_config.use_host_header_as_sni` | [access_info.tls_config.use_host_header_as_sni](data-sources--secret_management_access--reference--group-001.md#canonical-0213020113222102-3012113030302333-2013103220233122-0031232211300312-0112113112010332-0330333203322113-0002300002102002-1103033312322010) |
| `access_info.vault_auth_info` | [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3121013112330211-1031313011202211-1103021331302210-0001023222103233-1232110133333331-2330331113021031-0132321111213333-3230100003301132) |
| `access_info.vault_auth_info.app_role_auth` | [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2003300311230032-3101330022111120-0301320313012000-3332130101131323-2013122012110023-2021212303131033-3101112300131203-0330303120003020) |
| `access_info.vault_auth_info.app_role_auth.role_id` | [access_info.vault_auth_info.app_role_auth.role_id](data-sources--secret_management_access--reference--group-001.md#canonical-3003333122011202-0301020320203102-1303032111133312-0322311230213230-2323121323310331-1030300103102231-1212321312021333-0022030321023212) |
| `access_info.vault_auth_info.app_role_auth.secret_id` | [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-2301133211213110-3003223123030322-1223012110113120-3210202011132312-1301110333011231-0033210033323113-1211310200120311-2110123233032210) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-1123031301231032-3020202231313010-2332313230102321-2121322011001102-2301300313333300-2220111003331022-1133011301302331-2312210000121001) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-001.md#canonical-3332221120110303-3131221100012212-0203120011321300-1003320322300312-2203113102020113-3021022100231131-0031322110102001-2123230021000022) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-001.md#canonical-2012012203201112-2213202132201311-3333332010202020-1330310201300220-1001332123001010-2322133103313313-0220302223232023-1330130021211123) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-001.md#canonical-1212122113102213-2312200130202113-0230330303320033-1022231322322123-0211110333130303-1321113133213112-1113111100012203-2121001021323130) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3100233101303200-3021132120010212-1222112320313312-0321113032123231-0010020012201210-1300211221213200-3021001030132220-3210301222211221) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-001.md#canonical-0111110323030220-1221133112102302-3020312111000201-0311122000330210-0330112102332332-3120022101211203-0303313312033320-0003121122320122) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url](data-sources--secret_management_access--reference--group-001.md#canonical-0301122312022223-0320030333010321-3201331123120332-0111023123113100-0130203323013232-1132031333302111-0231322123002210-3113321120000333) |
| `access_info.vault_auth_info.token` | [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-001.md#canonical-0022122320212220-3202033221010132-0222312101211131-2321320311103112-2202101021123031-2230111222301311-1222303121333330-2300210021233310) |
| `access_info.vault_auth_info.token.blindfold_secret_info` | [access_info.vault_auth_info.token.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-2000101120331303-3230012210023113-2302112301002300-1312033301222103-2033032211313030-1102303102002311-0233021232002302-3132030301020322) |
| `access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-001.md#canonical-2123103202031232-3023230213333130-2111033331311223-0021031013201033-0012313001301213-0303213111300011-0233012332003320-0332021212303101) |
| `access_info.vault_auth_info.token.blindfold_secret_info.location` | [access_info.vault_auth_info.token.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-001.md#canonical-2330203333102013-1201020331222013-3103132332000221-2222301231231220-1100023022033331-1031203020102231-3320313131030333-0112210330131331) |
| `access_info.vault_auth_info.token.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-001.md#canonical-1110201202123120-1123130112012000-1123231320031230-2232133100301222-3211303011333001-2210111230022023-1201303322113212-3213103321300030) |
| `access_info.vault_auth_info.token.clear_secret_info` | [access_info.vault_auth_info.token.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322100230302323-2211102110110000-0101322312020333-1300211312023200-3321301023133030-1330203320200110-0001032011213110-3213020033123313) |
| `access_info.vault_auth_info.token.clear_secret_info.provider_ref` | [access_info.vault_auth_info.token.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-001.md#canonical-0010303232000132-0233132320002133-3220031313120133-1311103230331022-2010130101232303-0003031133222233-0300313333132302-1020113100331111) |
| `access_info.vault_auth_info.token.clear_secret_info.url` | [access_info.vault_auth_info.token.clear_secret_info.url](data-sources--secret_management_access--reference--group-001.md#canonical-2123203212021322-2130123100211003-1133310310010113-0222011121122013-2113321003002302-3310303312032130-2010110013012303-1010030222321013) |
| `annotations` | [annotations](data-sources--secret_management_access--reference--group-001.md#canonical-0020022133311021-2022112212231122-3002022022213000-2321300310310202-1303120203321303-1111311002032120-3232000222223133-3213020221113313) |
| `description` | [description](data-sources--secret_management_access--reference--group-001.md#canonical-3122210100112130-1313303203321031-2030323222331321-0330210013020311-1031201211022211-0002022031330033-3031210210201132-3022033203223031) |
| `id` | [ID](data-sources--secret_management_access--reference--group-001.md#canonical-3300331133320113-3021133221213331-3322100223310211-0010012221331102-2232301032010230-3332203302330332-3233103133102310-2230101011212021) |
| `labels` | [labels](data-sources--secret_management_access--reference--group-001.md#canonical-1223222333112320-1303011301113202-0103111032130113-3023313003111001-2000231233330002-2211033023210132-3312110223232331-3222000022020112) |
| `name` | [name](data-sources--secret_management_access--reference--group-001.md#canonical-0221012311032323-1211103132120012-0010211303010031-1103303112001001-2211232232020232-3023112313212212-0133010211003102-2322322313200312) |
| `namespace` | [namespace](data-sources--secret_management_access--reference--group-001.md#canonical-1211001101231101-2231021223312012-1231030031230121-1313210332331310-3113333033022033-2212221201122022-3230113201003223-3232311111031233) |
| `provider_name` | [provider_name](data-sources--secret_management_access--reference--group-001.md#canonical-1100332300303212-0302032103100023-2012103030312130-3011032122333203-3100102012022131-0101201213232200-1130000200313130-3111201031320322) |
| `where` | [where](data-sources--secret_management_access--reference--group-001.md#canonical-3002002312300203-3333333000211233-1230031002223332-2233310003333231-0321230033230330-2132001131033032-1210230322002222-0123221101110303) |
| `where.site` | [where.site](data-sources--secret_management_access--reference--group-001.md#canonical-3332301333021022-3213033023310121-1000023201332233-3131300113321222-2103321032312310-0030301012000231-3233102010210212-1120213112220330) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--secret_management_access--reference--group-001.md#canonical-2012121010311023-0320322323221133-2231213331200223-0303303332211213-0300022003020012-2113213130312000-0131000211021201-0220110313303302) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--secret_management_access--reference--group-001.md#canonical-1111122123111022-3033221332301213-0100023031220221-1221230002220021-0302100212310033-2001123033002000-2331010113112222-2120221221310123) |
| `where.site.network_type` | [where.site.network_type](data-sources--secret_management_access--reference--group-001.md#canonical-2210021123301000-3323031211313332-1301202310021222-3311220210320320-1222331201230200-3112100013321230-0021221113101111-2120133122103022) |
| `where.site.ref` | [where.site.ref](data-sources--secret_management_access--reference--group-001.md#canonical-2320023130132003-3222300321212333-0111020000010222-2033323012310311-2201013331210103-3033122023121112-1030120122210330-0031121111132021) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--secret_management_access--reference--group-001.md#canonical-3223132211021303-3233223303220121-1101102203131010-3000103220102132-3133111113201222-0333301320020230-1100211301310112-0231011302000102) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--secret_management_access--reference--group-001.md#canonical-0030100011312200-3032013002233113-3313003313012123-1000310112313100-3223131023210120-2202311112123023-1030330031222233-3311212202002223) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--secret_management_access--reference--group-001.md#canonical-0122223323201200-0210313020303120-0111132033302032-0300113022210113-3033033323121122-0033112310031321-2213033101303330-1311301302301023) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--secret_management_access--reference--group-001.md#canonical-1133311323211200-3113102312303032-2210222003122001-3130233331203020-2103000302023203-2001012123133332-2202013331202303-3210301321330302) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--secret_management_access--reference--group-001.md#canonical-3211200102122130-3011210131320233-3031132330313121-3032221130032300-1110221212131010-3110133100313223-1323102122201312-3332231020120212) |
| `where.virtual_network` | [where.virtual_network](data-sources--secret_management_access--reference--group-001.md#canonical-0330112231003221-3223230112200210-2101230313300030-0003230131212010-2032120121313002-1012323320223003-1310000313322201-1010102300232330) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--secret_management_access--reference--group-001.md#canonical-1130233211123201-0120130313000322-0111210222311320-2022133233103301-3212131022333332-3131312323313102-2002132023230033-3100210201020200) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--secret_management_access--reference--group-001.md#canonical-1213000012233011-3321212101033221-0310320123202300-0102231032000321-2233030002033002-3000201102111311-3123210201300121-0130233331122330) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--secret_management_access--reference--group-001.md#canonical-2310132331232320-3333330011323000-3123213032332310-3121321221202020-0120003012300221-2202133102201030-2012032323000331-0320233321222013) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--secret_management_access--reference--group-001.md#canonical-1030131032213201-1113312220113313-3220321110301301-2001003223320010-2013333020133113-2103121320000313-0103010110323003-3033230301331000) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--secret_management_access--reference--group-001.md#canonical-0010230300222313-0310010122111122-1003033123032201-2211303032310212-2102212132022321-0233321120212131-2213012203021101-3112021032331111) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--secret_management_access--reference--group-001.md#canonical-1301030132220110-3202201203323322-2132223330211023-2030020002121122-3013220213231130-2222003131220322-2022102101023310-1203001102213032) |
| `where.virtual_site` | [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-2001033311032022-3000322323110031-3221011302233023-1300031320330201-0112132010323310-2221301210010212-2301013033120303-2032321220111202) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-1132313222220030-3302030131103023-0000322313002232-0221120003000322-2011313122022312-0333303011102330-0132121301002002-0223321212101203) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-1120002302323123-0012311313203013-0200032113300311-2030232003001321-3312010021210023-0201002020323002-0113303130112323-0301221133302001) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--secret_management_access--reference--group-002.md#canonical-1333032023122311-3110002120230000-0200021213113010-3121030313210303-2002002201001113-2301010232013211-2101322221203110-1323303303200021) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--secret_management_access--reference--group-002.md#canonical-1232031000220131-0320210023132201-0233211020113323-3032000231123321-0020213122020321-2221002211321031-0031213203220013-3000022201231213) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--secret_management_access--reference--group-002.md#canonical-2202231031301311-3112211023032113-2302032230320220-0232123232133030-3313130222220100-3313301121023100-2113122112232221-0211030101330320) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--secret_management_access--reference--group-002.md#canonical-3232232030013323-0203121131202312-2322212303200312-3131010003100030-3213023132001000-1003301212231222-1211231122330020-2012012113233330) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--secret_management_access--reference--group-002.md#canonical-1021230013230332-2210123030202320-0233211200001322-3301113201322201-0130231201123001-1211120132102222-2231003301111031-2313310131000020) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--secret_management_access--reference--group-002.md#canonical-0220112000000022-1113232310030133-0333012222320322-3020331312311010-0123103320331032-0223231323123300-3322232031311202-1030330001001103) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--secret_management_access--reference--group-002.md#canonical-3311022013110010-2133332013330030-2020231331130110-2233232200330123-3131303133133111-0033013113332232-0302321021000013-1023330330013103) |

<a id="canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- access_info

<a id="canonical-1321022221213221-0020110111230331-3112013223203201-2022311211233302-2001110021330213-2003213320321322-2020012200211213-1312130320321012"></a>

Type: `"single"`. Computed.

HostAccessInfoType contains the information about how to connect to the remote host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"rest_auth_info\",\"vault_auth_info\"]"
}
```

<a id="canonical-1232333011113313-0311231032020023-2032212231000230-3311201103120113-2120103332010222-3213300003312313-3303002120130200-2120023333332100"></a>

### Direct properties for `access_info`

- [rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131): complete subsection reference.

<a id="canonical-3032302011210022-1121311233220103-1230322100222110-1120213120331101-1322101210231233-2201300303331003-0332130121010211-0012001232210210"></a>

<a id="canonical-2101222120130202-1222121213013232-1123100312010200-1330021222213030-2322303320103303-0020323000312122-2220321133313132-3320223303102023"></a>

#### `access_info.scheme` property

Type: `"string"`. Computed.

\[Enum: HTTP|HTTPS\] SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.
Possible values are \`HTTP\`, \`HTTPS\`. Defaults to \`HTTP\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "HTTP",
  "enum": [
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0002020123320112-1213103113213313-1120323212103202-1022311130212120-1232331032030321-2311312023033223-1001213132130302-0220003301203313"></a>

<a id="canonical-1112302011022012-0200101113202210-3233121101001213-0011023202220302-1111302232131320-3031310110303112-2010133132102031-2220033010120111"></a>

#### `access_info.server_endpoint` property

Type: `"string"`. Computed.

Endpoint to connect to, in host:port format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223): complete subsection reference.

- [vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022): complete subsection reference.

<a id="canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.rest_auth_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- access_info.rest_auth_info

<a id="canonical-1320233201201101-1201330311212301-0222020003313211-2131223110112020-1312031130313222-1101110100012033-2122012312300030-3123123321000300"></a>

Type: `"single"`. Computed.

Authentication parameters for REST based hosts.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"basic_auth\",\"headers_auth\",\"query_params_auth\"]"
}
```

<a id="canonical-3312102303012222-0231230333202123-0032012302322313-1120320101022130-0221112122331132-3013120330230200-2203110103111133-2022200212122211"></a>

### Direct properties for `access_info.rest_auth_info`

- [basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301): complete subsection reference.

- [headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2220011133213131-3101131111102021-2320030231133223-0001320232332110-3230133310022213-3033122313002031-1022302012302010-2123102011221030): complete subsection reference.

- [query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-3123211020031221-1033310323003001-3333320020332010-2311112310201202-2122332031233311-1221220333200110-1203202212131100-1123210312001010): complete subsection reference.

<a id="canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.rest_auth_info.basic_auth` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- access_info.rest_auth_info.basic_auth

<a id="canonical-2320002323003120-1231310201333211-0112330332133133-1030220201331310-0332320031312313-2220011333102303-0320110202320330-3230123122210202"></a>

Type: `"single"`. Computed.

AuthnTypeBasicAuth is used for using basic\_auth mode of HTTP authentication.

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

<a id="canonical-1121221020332322-1322202232320331-2233133020012023-1023320022221012-1332322021131210-2111103020321300-0120300020211113-2212321022021202"></a>

### Direct properties for `access_info.rest_auth_info.basic_auth`

- [password](data-sources--secret_management_access--reference--group-001.md#canonical-0232033033232332-3300113133233020-2200212102320012-0212102103200202-2100020112220023-1123033031113112-3222110212233103-0101132102132212): complete subsection reference.

<a id="canonical-0233323213032201-1112203333110232-1333303130223032-3322300201133202-0200132013313100-2303233000222311-3121313330300213-1111232331033302"></a>

<a id="canonical-0030233000120332-0100323102333123-3223233033022101-3013012331223323-1000212033010133-0023121300303300-2111303323330110-3303223221312011"></a>

#### `access_info.rest_auth_info.basic_auth.username` property

Type: `"string"`. Computed.

The username to encode in Basic Auth scheme.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0232033033232332-3300113133233020-2200212102320012-0212102103200202-2100020112220023-1123033031113112-3222110212233103-0101132102132212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.rest_auth_info.basic_auth.password` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301)
- access_info.rest_auth_info.basic_auth.password

<a id="canonical-2322323102101021-2030302123320230-0021232020212113-3001113300332000-0102320223201220-1321231331103111-0300230233313001-0002211210022320"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3233322331202230-1120022330232212-2200022021132112-2113001211131231-1232000301130010-1311020001213020-1233113220022031-1332330101122132"></a>

### Direct properties for `access_info.rest_auth_info.basic_auth.password`

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3330003132112301-3201302110220203-2321021101021220-3110132003321121-0303310323130333-0211101321212113-0012102100223031-2232231033213330): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3010333002113202-3302323331302311-1031121123121330-2320131130223022-3210312232203203-2110300300102023-0010023313013122-3333303020301003): complete subsection reference.

<a id="canonical-3330003132112301-3201302110220203-2321021101021220-3110132003321121-0303310323130333-0211101321212113-0012102100223031-2232231033213330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301)
- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-0232033033232332-3300113133233020-2200212102320012-0212102103200202-2100020112220023-1123033031113112-3222110212233103-0101132102132212)
- access_info.rest_auth_info.basic_auth.password.blindfold_secret_info

<a id="canonical-2131220230312112-1003322331022133-0320120200303022-0033311223000011-0103223201312033-3233122010202100-3110310012103300-1023212011000021"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0123300221332331-3213213333203231-1300221103112320-0000000023230202-3212001111102331-1203002023021122-0333320031331331-3112001133100123"></a>

### Direct properties for `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info`

<a id="canonical-3100331020123101-3220010230212110-1200302130233032-0032223323101331-1012002311110022-2321013100102001-1330302012213330-1101122023303313"></a>

#### `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-2022311031003332-1232210332023001-2212321022110122-0111312121120133-0133313210213212-3233000200002232-3112331220110300-1201023130330332"></a>

<a id="canonical-0211223201001103-2012120002313002-1203312002120101-2200123310121113-3333021011222101-1232213312031200-0213023330112220-3331002003033223"></a>

#### `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-2023022213122100-3300000332201231-0010322102300002-2212323213130331-3213103133200332-0133123212123131-2021010321311102-1303200011110322"></a>

<a id="canonical-3122112233112203-0303223300121023-2110331223321300-1001201332133230-2032022202331231-0330031301023012-1021303210123023-3320200202020003"></a>

#### `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3010333002113202-3302323331302311-1031121123121330-2320131130223022-3210312232203203-2110300300102023-0010023313013122-3333303020301003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.rest_auth_info.basic_auth.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301)
- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-0232033033232332-3300113133233020-2200212102320012-0212102103200202-2100020112220023-1123033031113112-3222110212233103-0101132102132212)
- access_info.rest_auth_info.basic_auth.password.clear_secret_info

<a id="canonical-1023032203302020-0003323100002000-3033120200333123-0031021010133130-1020123212223301-3210223310031102-2203032300212313-0211020123130123"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1300303023321103-2133121121101113-2133132302210213-2030211231333201-0012022203031112-2023231122321223-0000012232300312-2002233200100322"></a>

### Direct properties for `access_info.rest_auth_info.basic_auth.password.clear_secret_info`

<a id="canonical-0001101123000030-1003112211133130-0132330022013332-0322101012122110-0230230001001011-2002330302003120-1102002213213313-0032331101200101"></a>

#### `access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2030122203203233-3130301111300133-0101310302010032-2221033203333113-2322312021013221-2132213112021211-3310321010033232-3310121302213120"></a>

<a id="canonical-3102100013101022-3210121011332003-3112120312110323-2102320303113301-1212100130022103-0331212321310101-3011012302233301-1301310032330321"></a>

#### `access_info.rest_auth_info.basic_auth.password.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-2220011133213131-3101131111102021-2320030231133223-0001320232332110-3230133310022213-3033122313002031-1022302012302010-2123102011221030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.rest_auth_info.headers_auth` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- access_info.rest_auth_info.headers_auth

<a id="canonical-2011112102113323-2031321013332332-0123111121020322-3131300221121021-1332231130310330-0100012011231013-2103321200201202-3103200010232102"></a>

Type: `"single"`. Computed.

AuthnTypeHeaders is used for setting headers for authentication.

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

<a id="canonical-2021331033031011-0313323023100121-3111331223210032-1232332121020322-1023000302311133-3022102312112122-1033033121320313-0113233132022221"></a>

### Direct properties for `access_info.rest_auth_info.headers_auth`

- [headers](data-sources--secret_management_access--reference--group-001.md#canonical-1003003331312200-3323113111311113-0213222002100131-1121230213000312-0320022003312030-2011321112231130-0022312302021103-3230123311311122): complete subsection reference.

<a id="canonical-1003003331312200-3323113111311113-0213222002100131-1121230213000312-0320022003312030-2011321112231130-0022312302021103-3230123311311122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.rest_auth_info.headers_auth.headers` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2220011133213131-3101131111102021-2320030231133223-0001320232332110-3230133310022213-3033122313002031-1022302012302010-2123102011221030)
- access_info.rest_auth_info.headers_auth.headers

<a id="canonical-2023110200112120-1322333132030021-2302011231030012-0332010123220132-2202011213133132-0131122232320213-1231002221212303-2100312100011323"></a>

Type: `"single"`. Computed.

The set of authentication headers to pass in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123211020031221-1033310323003001-3333320020332010-2311112310201202-2122332031233311-1221220333200110-1203202212131100-1123210312001010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.rest_auth_info.query_params_auth` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- access_info.rest_auth_info.query_params_auth

<a id="canonical-0201033223131030-1130030212123032-1030221223120301-3110230101031233-1233223132303003-3311101000130133-2123300300120213-3331231023101002"></a>

Type: `"single"`. Computed.

AuthnTypeQueryParams is used for setting query\_params for authentication.

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

<a id="canonical-1301000223002130-1212232322102102-0332122002330211-1320033320222113-3112033322321001-1011033302201330-0100131333331321-3111121011120302"></a>

### Direct properties for `access_info.rest_auth_info.query_params_auth`

- [query_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131023122000323-2202200220101021-1332203131103313-2230102121223331-0010322101130331-0122002331202321-2113020132300322-2332112121123032): complete subsection reference.

<a id="canonical-3131023122000323-2202200220101021-1332203131103313-2230102121223331-0010322101130331-0122002331202321-2113020132300322-2332112121123032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.rest_auth_info.query_params_auth.query_params` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-3123211020031221-1033310323003001-3333320020332010-2311112310201202-2122332031233311-1221220333200110-1203202212131100-1123210312001010)
- access_info.rest_auth_info.query_params_auth.query_params

<a id="canonical-0011000122222030-3211231011101312-0020101233023130-3321232012021000-1233133100200231-1212212320303321-0023120202231221-0130232221222310"></a>

Type: `"single"`. Computed.

The set of authentication parameters to be passed as query parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- access_info.tls_config

<a id="canonical-2203201032131033-1021100022222100-2002003021221310-1210100033233311-1202033222032221-3313130122231312-3101012302133313-1220130232231033"></a>

Type: `"single"`. Computed.

TLS configuration for upstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]",
  "x-ves-oneof-field-tls_params_choice": "[\"cert_params\",\"common_params\"]"
}
```

<a id="canonical-3221330301333100-3332210211021103-1021201102001110-2113012013303100-1323023122121103-2211203312300312-0330211202132020-0130031210322211"></a>

### Direct properties for `access_info.tls_config`

- [cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223): complete subsection reference.

- [common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102): complete subsection reference.

- [default_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-1210103112123023-1311003031333303-0033320301002230-1333002123130023-2302111201030130-2311312223001120-3320000021000013-2132010033213110): complete subsection reference.

- [disable_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-0011323301322002-2201102122202131-2130020111322130-1323121110112002-3121221011021313-3000012122012233-3313000221030220-2013020223002321): complete subsection reference.

- [disable_sni](data-sources--secret_management_access--reference--group-001.md#canonical-3313002020213130-1211220100111301-2010321010111112-1221223002230323-2012201001201323-0033002301121033-0211321133002012-3230002322001001): complete subsection reference.

<a id="canonical-2032013333213230-3300320223333220-2033233103202101-3220311331110103-3121132030122232-2210011212311302-1021003203020221-1323111130232100"></a>

<a id="canonical-1121233223030221-1030230200012031-1212323021021121-1323121231203022-2313210023011130-1112333003210203-3110030310212033-3322333012323332"></a>

#### `access_info.tls_config.max_session_keys` property

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

<a id="canonical-1111312120133212-1332221003133322-3302033021111130-0211222212122331-2122221013010223-3123110000122031-1322131030202210-3103220221030012"></a>

<a id="canonical-1330300001111332-0021023331213311-0000110011100133-2021322102221013-2330210113331230-0221322313031223-3033320222110231-0000213202003030"></a>

#### `access_info.tls_config.sni` property

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_host_header_as_sni](data-sources--secret_management_access--reference--group-001.md#canonical-1323322302212031-2102331223003212-2232132102320011-3313132032020021-3013331110313121-0310233210023203-0320100300121321-1033322112222012): complete subsection reference.

<a id="canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.cert_params` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.cert_params

<a id="canonical-2213122301012012-0000300100312320-2022211030311300-0123202030212023-3120031112120123-1102123201103330-3330203213321312-1233032133113222"></a>

Type: `"single"`. Computed.

Certificate Parameters for authentication, TLS ciphers, and trust store.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"tls_validation_params\",\"volterra_trusted_ca\"]"
}
```

<a id="canonical-1111331130001300-0101030321323331-1200123321322002-0123331130330132-3000101331001211-0221010203013301-2010303222303002-0033311220220322"></a>

### Direct properties for `access_info.tls_config.cert_params`

- [certificates](data-sources--secret_management_access--reference--group-001.md#canonical-3223110302010102-0112021220103031-1023330021121210-2202021320210313-0100213121031322-2210312201321303-3230321311203132-1103022001111310): complete subsection reference.

<a id="canonical-1020030112031312-3000310203203331-2123232212001332-3123012331231322-0321300301231102-3232200303203103-1130020032322222-0232311032330011"></a>

<a id="canonical-2013011223332333-2102233202211232-0233200322302101-3111322121103323-0003021120222222-2233213202220311-3203311303032311-0032321002020030"></a>

#### `access_info.tls_config.cert_params.cipher_suites` property

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0300100113211332-0022233111021023-1320201111130102-1313100212303133-0112213123022120-2231213013300023-0221331032210132-0201102210123210"></a>

<a id="canonical-0123200211200220-0223322012223302-0302322113321202-0231312300323021-2121332013132213-2201303021001020-2222031030233113-3121302123210320"></a>

#### `access_info.tls_config.cert_params.maximum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0320013333312130-1002022111022312-0203133100100120-3211103002210202-1300023303030323-3112201111203130-3211011303300301-2311230301002202"></a>

<a id="canonical-1232002031003030-2221320202032000-1223121031011003-3202103123123113-3201023310320311-3332201111332021-3121200123200213-2032102232211013"></a>

#### `access_info.tls_config.cert_params.minimum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [skip_server_verification](data-sources--secret_management_access--reference--group-001.md#canonical-2133021130203313-1233113323230223-0313000111132020-2220211301331123-0030220000122311-3113101332103210-2202210131013211-2222322321013300): complete subsection reference.

- [tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0212121101200211-3311122021003102-3323022130010010-1220003310202223-2133133021200212-3203110132212111-0110011212300112-1130202332133031): complete subsection reference.

- [volterra_trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-1010303013132223-3201132001022120-0210133300210210-3123033023101113-1110313213030032-1202123311312003-2321210220210001-3231102311132323): complete subsection reference.

<a id="canonical-3223110302010102-0112021220103031-1023330021121210-2202021320210313-0100213121031322-2210312201321303-3230321311203132-1103022001111310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.cert_params.certificates` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- access_info.tls_config.cert_params.certificates

<a id="canonical-2110212023202120-1133030112012123-2301222133112222-1312120020030213-3213211012212222-3230130233110221-0112022223301312-3001302300321021"></a>

Type: `"list"`. Computed.

Client TLS Certificate required for mTLS authentication.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1301312112222313-2001023203223132-3003310000202210-0002201323132012-2330232312230003-3322330222211101-3122013213330100-2322201210110210"></a>

### Direct properties for `access_info.tls_config.cert_params.certificates`

<a id="canonical-0113131011332303-1313131022031221-1201111202212301-1133301123101012-1321302321103321-0030312222012310-2313200301302330-3330312232122010"></a>

#### `access_info.tls_config.cert_params.certificates.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-1022031202133120-1322302013003031-2003001012010332-3112310011200111-0330030332231212-0210232222211113-2302111033210300-0010013333222210"></a>

<a id="canonical-1320211101213203-3321322031221030-3030201001331232-0122231311012310-0032332011133213-0303320130033323-0231101311310010-3102233233102310"></a>

#### `access_info.tls_config.cert_params.certificates.name` property

Type: `"string"`. Computed.

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

<a id="canonical-2302130233313022-2133013233300310-1131013231002002-2202111332020021-0020213010233313-1003110100302032-0000001223323003-2102023033110121"></a>

<a id="canonical-0020003203222102-0121302002111311-2123121221013203-0113013330313333-0300030110312201-0313232211330010-2312000223230102-2310132030012200"></a>

#### `access_info.tls_config.cert_params.certificates.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-2023133000010201-2111303032303220-0233002310101210-0121321000030200-1332301113111030-1222223133320102-0021002303120220-1333312201123320"></a>

<a id="canonical-3201300113312033-2300032022123202-2101213000022311-0102020112131033-0112001310010203-2123231023111330-2001120321021211-3133113223300102"></a>

#### `access_info.tls_config.cert_params.certificates.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3301303130001033-0013331323000000-0333133132200133-3133122330212202-2002230031112100-3211101320001030-0132002103033100-1223120111130022"></a>

<a id="canonical-1200033113312111-1030011001001232-0022130011111100-1012021301133010-0101200013230330-0102133122331332-3330110003120021-0210100310311203"></a>

#### `access_info.tls_config.cert_params.certificates.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-2133021130203313-1233113323230223-0313000111132020-2220211301331123-0030220000122311-3113101332103210-2202210131013211-2222322321013300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.cert_params.skip_server_verification` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- access_info.tls_config.cert_params.skip_server_verification

<a id="canonical-2132300231222233-1020221121030211-0221333203332203-0133131031330000-0021220122300002-2211020212333302-2201203313113320-2111020120112112"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212121101200211-3311122021003102-3323022130010010-1220003310202223-2133133021200212-3203110132212111-0110011212300112-1130202332133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.cert_params.tls_validation_params` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- access_info.tls_config.cert_params.tls_validation_params

<a id="canonical-1313132013331301-1022311210223030-0000030232233231-2311012200110231-3112030323013302-2202023231323321-1021232311213211-3031303232220032"></a>

Type: `"single"`. Computed.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-1110303120121220-0233033133303012-3202003313133232-1002331202011202-3123323200101011-0032002131202310-0210131321301011-3103103132210322"></a>

### Direct properties for `access_info.tls_config.cert_params.tls_validation_params`

<a id="canonical-1032310210233100-0330132203201222-1211102130203331-3302101230221012-3033202002131331-0233213211010101-2200002221001103-2131312311230032"></a>

#### `access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification` property

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-0021003012231102-1113101021031010-1123303330032223-1230302120220031-3233030320222203-0030320213202202-3132332220213000-3023131312113012): complete subsection reference.

<a id="canonical-2113222321200120-1201111132130130-2203211320022103-2003211303320011-3003303023331320-2212032223221323-2100200010032030-0312121303231231"></a>

<a id="canonical-1102320333231003-2212333031022232-1323333001331333-2020322112300202-2210310022222303-1221003201221303-2101211211001333-3220121133120012"></a>

#### `access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-0002302330330222-1130332300200213-3233100311021111-1310303212011023-3322031000132031-1213030103100122-0230103331330022-3330020110013000"></a>

<a id="canonical-3132200112003000-0311003020200003-2101212220332032-0110323313331312-2312031020211302-2110300011013310-0001130132012033-0120303330103030"></a>

#### `access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-0021003012231102-1113101021031010-1123303330032223-1230302120220031-3233030320222203-0030320213202202-3132332220213000-3023131312113012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.cert_params.tls_validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0212121101200211-3311122021003102-3323022130010010-1220003310202223-2133133021200212-3203110132212111-0110011212300112-1130202332133031)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca

<a id="canonical-3020202221322302-1010133021111101-3103213113030031-0132000033312122-3103320022330000-2321332000112212-3322003101333230-2213311112130033"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

<a id="canonical-0212202123212022-3022200222113302-1031103101111223-2332103201320101-2133020312021320-2102220300023332-2231002011330103-3322122002123302"></a>

### Direct properties for `access_info.tls_config.cert_params.tls_validation_params.trusted_ca`

- [trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-0110131112111211-1221112123033022-3301310001110213-0311032032222233-0220321220111223-2202200221320332-1013200313202213-3232021212132103): complete subsection reference.

<a id="canonical-0110131112111211-1221112123033022-3301310001110213-0311032032222233-0220321220111223-2202200221320332-1013200313202213-3232021212132103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0212121101200211-3311122021003102-3323022130010010-1220003310202223-2133133021200212-3203110132212111-0110011212300112-1130202332133031)
- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-0021003012231102-1113101021031010-1123303330032223-1230302120220031-3233030320222203-0030320213202202-3132332220213000-3023131312113012)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list

<a id="canonical-0011333100021110-2032212333221023-1311001220020113-3233123213300202-2212012032332321-0213100100230230-2121221023120323-0130031120033110"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2032233300213133-3103032312323023-1031333101110321-1233323212000130-0321011212313231-2321032332030112-0033010123211203-3002331212000321"></a>

### Direct properties for `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-0312032301000210-3103022130322030-1131020112113000-2201023023031101-1312220100231132-2001301220220212-0123213233311301-2022021022103200"></a>

#### `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-1310220101131202-0100001130230321-3033330331112123-0110010232313002-2111130221301232-2223130212222132-1303311320021110-3001100201023331"></a>

<a id="canonical-1210302233100110-1003000200310211-2232213001103322-0010133130223031-2030300022000001-3132331111110002-0111232311121010-2312110123023103"></a>

#### `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` property

Type: `"string"`. Computed.

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

<a id="canonical-2133113022223010-0120311021130301-3023000100231000-2302013312333203-3011023310223103-0213331332331122-0001300123103020-1130202310230030"></a>

<a id="canonical-0320133313213213-0321202113303200-2333231323130210-3030322033021201-3113220331011201-3233212113313203-0233231221233122-2022301310231201"></a>

#### `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-3323100102202321-2332031020202201-1033230211333132-0202223111010212-0000033133323230-1331320332102002-0133022011303312-3130132101310112"></a>

<a id="canonical-1232230333303322-3130222130111121-3002003030202221-3302030211210000-1033031333010033-0023100120111212-2102202330001313-3130311001313020"></a>

#### `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-1303323033233032-0130311221000233-3331200300231221-3312303233112202-3032331121101312-0120320300021300-2320021211322220-2232100130333213"></a>

<a id="canonical-3332201222212102-0132100232021302-1113122322013320-0212212333222013-3101123322330301-1333020003303110-1231211332132033-3021012111012311"></a>

#### `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-1010303013132223-3201132001022120-0210133300210210-3123033023101113-1110313213030032-1202123311312003-2321210220210001-3231102311132323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.cert_params.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- access_info.tls_config.cert_params.volterra_trusted_ca

<a id="canonical-0123112310323301-3202211211010010-2130211130210323-2300111201232131-1013230022010210-3130320332100012-1310133122322001-2332300223113233"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.common_params

<a id="canonical-1103203122021022-3002220212010200-3131203100120021-0131222302231112-1332321003021333-2300130033220033-2000101002203311-3323102012002322"></a>

Type: `"single"`. Computed.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

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

<a id="canonical-3002033232310310-0030331312333130-2302130233000003-0130212123112100-1231331123131213-2112310323110211-0312120223321303-1232222320320002"></a>

### Direct properties for `access_info.tls_config.common_params`

<a id="canonical-3230322012213021-1023003033321213-0210012111120310-1122130203002221-3122000322021111-1131121112333101-3312310032013221-2130302230121112"></a>

#### `access_info.tls_config.common_params.cipher_suites` property

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0323301101033133-2110003220020032-0323210222333220-2301200200102011-2100300232200330-3131330202323221-0010102311122101-3000111323233203"></a>

<a id="canonical-0322303323311201-2232130321112133-1212313020231231-3222313233323211-0000010023000123-2132201013121133-0313201213031133-2123210332012133"></a>

#### `access_info.tls_config.common_params.maximum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1032323031121103-0130103301311111-1000020000100323-0111131013121201-1322113103323302-0010033132113202-0132223023130001-2001200233100231"></a>

<a id="canonical-2131133211023032-0111210030022100-0010133333212110-0201003323123213-1333221032031011-3220230231022302-1011230121322131-3021222230030301"></a>

#### `access_info.tls_config.common_params.minimum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312): complete subsection reference.

- [validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0210313100110112-3122102002210000-1211212311023133-3110030202023333-3302121303223220-0320032333320132-0313201002332222-2100311203122102): complete subsection reference.

<a id="canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.tls_certificates` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- access_info.tls_config.common_params.tls_certificates

<a id="canonical-3323302002301203-2303102011222212-1122111100230133-2202321212220011-1310301221313220-0000331231221313-0100112012032101-1222001020103231"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

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

<a id="canonical-0020213100223313-1322213331131320-1111133322133213-2222301101300011-0210232200122301-3210030332132202-3321101301023033-0221230113330233"></a>

### Direct properties for `access_info.tls_config.common_params.tls_certificates`

<a id="canonical-2013130102301020-0213010113000321-2230300211012101-2110210223313303-2023000313322312-0311331133320330-0023230222323330-1020220021110121"></a>

#### `access_info.tls_config.common_params.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](data-sources--secret_management_access--reference--group-001.md#canonical-0110032232203310-1222333123120133-3013132321021220-0200302121131311-1303213330133100-0011300221203021-0032303132320222-1311033000012033): complete subsection reference.

<a id="canonical-2022132113301312-1111232023232303-1221332103020123-2010332212030332-0201331120111013-3202300203130012-0221301200222100-3010023322312113"></a>

<a id="canonical-1323002123203001-1301130300002211-2200223230310112-1013032001313322-3030001230233100-0332213030233220-0330021332113223-2222101131213113"></a>

#### `access_info.tls_config.common_params.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--secret_management_access--reference--group-001.md#canonical-0003022330330313-2030232201210223-0202003033120100-2230302110300101-0331323332021332-3110011033310020-3231011111312021-0122331102022131): complete subsection reference.

- [private_key](data-sources--secret_management_access--reference--group-001.md#canonical-1301230322222320-1302321123212130-1023332303302213-3033032012112002-2202110332212332-0011110311012320-2332331130220033-2002133021021122): complete subsection reference.

- [use_system_defaults](data-sources--secret_management_access--reference--group-001.md#canonical-3130030022223102-0323033203032120-3020100320133323-0121230311203203-1333310130210032-3233123112131222-2210223233132233-2020231011123222): complete subsection reference.

<a id="canonical-0110032232203310-1222333123120133-3013132321021220-0200302121131311-1303213330133100-0011300221203021-0032303132320222-1311033000012033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-0013322302301232-2233303112233202-0302030212220222-0220011320321230-0303110213013032-3232321311310311-0202230323333233-2100223133020003"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-0323312000220300-0323020210332011-1313133030302102-3102011023021320-1130121032303301-3201211123322202-2223133110122001-1211203021222031"></a>

### Direct properties for `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms`

<a id="canonical-2300022023333230-0230100001123230-1230013320210000-3002012002020020-1332231300100203-1100131011113001-3033021133132310-1023003132101212"></a>

#### `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0003022330330313-2030232201210223-0202003033120100-2230302110300101-0331323332021332-3110011033310020-3231011111312021-0122331102022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-2301023220100010-0320220122321110-2120203020003311-0333333011232010-1032313032002300-0333213112212312-2201323020322030-3013003101101301"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301230322222320-1302321123212130-1023332303302213-3033032012112002-2202110332212332-0011110311012320-2332331130220033-2002133021021122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- access_info.tls_config.common_params.tls_certificates.private_key

<a id="canonical-0021302312122233-2112222102301000-1331223110333001-1132032013010223-1322101312101032-3010113211321221-0221313303202311-1312201132122122"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1312011330322123-0030310220012101-1312213213113132-0203223220230321-0013023003203112-1231222003031001-2133111200301102-1201123320000220"></a>

### Direct properties for `access_info.tls_config.common_params.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-0320200332323203-2122000132300020-3111121303202213-0123001310310221-1212133202322322-2032033221312033-3233021130220010-0121103233220222): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-0333021122331232-3201202312000300-1101032010131000-0131020333212013-2210100202010223-3013112100112333-3311320212230322-1312102133302320): complete subsection reference.

<a id="canonical-0320200332323203-2122000132300020-3111121303202213-0123001310310221-1212133202322322-2032033221312033-3233021130220010-0121103233220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-1301230322222320-1302321123212130-1023332303302213-3033032012112002-2202110332212332-0011110311012320-2332331130220033-2002133021021122)
- access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2212010130332122-1303103322011222-3030132023131312-1203000212201000-3030300203130223-0021300131222110-1333023233012111-2212112212011303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2202031130002330-1311022313303330-2210020101322011-2012232100332330-0320100101011020-2201031133202212-2302130333210021-3110202330130030"></a>

### Direct properties for `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-1113110033102300-1331103000031232-0103332101200132-2021131213222110-2303321311002320-0200300312023001-0100103323332212-2312211312032103"></a>

#### `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1230022323200132-1022013011100221-0213131332003222-1222210111231231-3232323120212201-2121023300310233-3112012123213320-0022210330010003"></a>

<a id="canonical-1130021202333002-1303301332212002-3102101133022203-1102101133231331-2303110302100232-3303021210031010-3133310303213213-0012101011220220"></a>

#### `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-0201211012022310-2123003310231001-2133203310301121-2302002012021202-0330133012333210-1331311221300031-0131313301221212-3110001301110100"></a>

<a id="canonical-2232133031202300-2100131213220202-1231333111023332-3131000022200110-3112333301323020-0131202310310202-0131221031330233-3123132120031311"></a>

#### `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0333021122331232-3201202312000300-1101032010131000-0131020333212013-2210100202010223-3013112100112333-3311320212230322-1312102133302320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-1301230322222320-1302321123212130-1023332303302213-3033032012112002-2202110332212332-0011110311012320-2332331130220033-2002133021021122)
- access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-3102220123101120-0321332330321020-1003131113103011-1000002330103033-3112023132313032-1133302312112100-1321212232121322-1123221133302021"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3213210122002103-2311003331231023-0102323331132103-3102211221101213-1032203321200310-0303231213010333-1212113322202123-0200011331233023"></a>

### Direct properties for `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0120320202103310-0323220321223103-3201213102300112-3232123322132223-2011311220211113-0101000303130102-3032332110022232-0302211120323122"></a>

#### `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3321212121332100-2030312333120000-1202220312223003-1103203331001310-2233302333221111-2120033211230212-3321032233012231-2012003121120111"></a>

<a id="canonical-1100110202113001-1321011122313313-1232110321022302-1201210111033210-0000212223002213-1222220113313111-2313210321110222-1021202131032102"></a>

#### `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-3130030022223102-0323033203032120-3020100320133323-0121230311203203-1333310130210032-3233123112131222-2210223233132233-2020231011123222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- access_info.tls_config.common_params.tls_certificates.use_system_defaults

<a id="canonical-0221030011001233-1220131011110221-2112300021012233-0301022111112111-3210301220202233-1213011122231203-0232001111032120-3121113230331330"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210313100110112-3122102002210000-1211212311023133-3110030202023333-3302121303223220-0320032333320132-0313201002332222-2100311203122102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.validation_params` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- access_info.tls_config.common_params.validation_params

<a id="canonical-2310202013033320-0302233011130222-0311210021331122-1031030033221210-1102101101331122-1132133133020033-2133011021210100-1130132330231221"></a>

Type: `"single"`. Computed.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-2321131010021020-3033202033102033-3212132102310333-0220120103313302-2110032303000021-1321231000203133-3233303101323022-2211220333202113"></a>

### Direct properties for `access_info.tls_config.common_params.validation_params`

<a id="canonical-0022113231002332-1332022301001112-2033100011111012-2123301312132021-1013302032023130-1320101133032103-2313100130222202-3203100232111303"></a>

#### `access_info.tls_config.common_params.validation_params.skip_hostname_verification` property

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-3323022232231121-0202001001312303-2231100110230321-2130122233101103-3231100313122112-1003222223102130-1121012032301232-0213203023200030): complete subsection reference.

<a id="canonical-2133201221000213-0231311300330220-0112013100123012-1302032322100113-1320212032001130-3133002220211113-1021022203031020-0002231110012310"></a>

<a id="canonical-0203321003211313-3033213230313010-2230323123002021-3202221233123332-0321212013102011-1211331022032010-2330020002121223-1021230313032211"></a>

#### `access_info.tls_config.common_params.validation_params.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-1032230033231110-3201011101300223-3030011213110213-1220231101101113-0003320101122332-3031031132332220-3303300312101222-1232121231302003"></a>

<a id="canonical-1102110031201001-1031113313313313-3100001233011210-1221331011120020-2010331121311003-3011231011232000-3221201312121103-2003301000102230"></a>

#### `access_info.tls_config.common_params.validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-3323022232231121-0202001001312303-2231100110230321-2130122233101103-3231100313122112-1003222223102130-1121012032301232-0213203023200030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0210313100110112-3122102002210000-1211212311023133-3110030202023333-3302121303223220-0320032333320132-0313201002332222-2100311203122102)
- access_info.tls_config.common_params.validation_params.trusted_ca

<a id="canonical-0112021103031313-1303200332123020-3132123331232303-1031310120202220-3313010312100222-2212332231321021-0101012211000311-3310213102222103"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

<a id="canonical-0000033121231032-3112123010300002-1310300222020321-2130210110000302-3110133302331001-3010002211020312-2030311221133113-2103302333103101"></a>

### Direct properties for `access_info.tls_config.common_params.validation_params.trusted_ca`

- [trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-3213223132332031-0010001010211223-2310021322133310-0221230213332012-3023122303213312-2101331020231003-1311031113300131-2323011120232000): complete subsection reference.

<a id="canonical-3213223132332031-0010001010211223-2310021322133310-0221230213332012-3023122303213312-2101331020231003-1311031113300131-2323011120232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0210313100110112-3122102002210000-1211212311023133-3110030202023333-3302121303223220-0320032333320132-0313201002332222-2100311203122102)
- [access_info.tls_config.common_params.validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-3323022232231121-0202001001312303-2231100110230321-2130122233101103-3231100313122112-1003222223102130-1121012032301232-0213203023200030)
- access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-1103231332101211-1333330113300120-1030012120031321-3131332330102033-2303022000311030-0022303111121023-3322203123211030-2200132332232030"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1032312020022022-3301023002010110-1133033230310032-3130202020303011-1201123120101232-0323121000200322-2031332112131331-0132210320333220"></a>

### Direct properties for `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-1333313301023221-1223133121110202-0031220012233033-0333213030300200-0032002132113212-0100311003130312-3000122030212002-3222313231320210"></a>

#### `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-0100311303002312-1131303113001211-3333032310033311-1131213110213100-2110222330212132-2002100133210220-2003021311020020-1333320203010101"></a>

<a id="canonical-2303121210222331-1200000223112331-2330301211213111-0221313311003120-2002221321131000-1122303111332113-2231000323002103-2001021111202310"></a>

#### `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name` property

Type: `"string"`. Computed.

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

<a id="canonical-0322313213003323-1202133300331332-2120113001121300-3213122231210231-3033202323111132-3233003100310321-0330301112021331-0312322310113332"></a>

<a id="canonical-1023202023022113-0300222133321231-0300331113222231-2212002010232032-1203202321021233-2022332021333220-3303231303222330-2331331132201233"></a>

#### `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-3311313120212301-2011301223022332-1212202102131313-2332100103020212-1012320301300212-1131000210302212-3322022031002122-1012322202222212"></a>

<a id="canonical-0020102302111222-0302323103132121-3031300033200321-3100311132202333-0220132000211002-1232003220223123-1221112133010322-3113101031131233"></a>

#### `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-0331333322023123-2021011310013103-3320302021130112-3131133200221212-0011102201210012-1312030022232232-2322000133232323-0022123102121210"></a>

<a id="canonical-2102122302130332-2231111332111121-3132202022100132-2231313120030021-2322023321022101-0020112002003002-2201130132030100-2210201301322223"></a>

#### `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-1210103112123023-1311003031333303-0033320301002230-1333002123130023-2302111201030130-2311312223001120-3320000021000013-2132010033213110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.default_session_key_caching

<a id="canonical-2131321032020123-0232231210002310-2031113022030210-3121310200233322-3333210211133131-0000310331313310-2330301011100132-3323211313212231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011323301322002-2201102122202131-2130020111322130-1323121110112002-3121221011021313-3000012122012233-3313000221030220-2013020223002321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.disable_session_key_caching

<a id="canonical-2032010020020012-1233223300322003-1213013022322221-2202210101010333-3111213332330112-0220031223131002-2112031122020333-1212301310002223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313002020213130-1211220100111301-2010321010111112-1221223002230323-2012201001201323-0033002301121033-0211321133002012-3230002322001001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.disable_sni` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.disable_sni

<a id="canonical-0000330111213131-3031223010221210-2130332300203033-1201323121001001-0120113220113311-2320210310313030-3310311321132301-0131032013233331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323322302212031-2102331223003212-2232132102320011-3313132032020021-3013331110313121-0310233210023203-0320100300121321-1033322112222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.tls_config.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.use_host_header_as_sni

<a id="canonical-0213020113222102-3012113030302333-2013103220233122-0031232211300312-0112113112010332-0330333203322113-0002300002102002-1103033312322010"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.vault_auth_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- access_info.vault_auth_info

<a id="canonical-3121013112330211-1031313011202211-1103021331302210-0001023222103233-1232110133333331-2330331113021031-0132321111213333-3230100003301132"></a>

Type: `"single"`. Computed.

Authentication parameters for Hashicorp Vault hosts.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_params": "[\"app_role_auth\",\"token\"]"
}
```

<a id="canonical-1203233223020320-3120110202010201-0010000132333111-3220333200211100-3031031300233311-2110222212022103-0223111022113021-1203003031221021"></a>

### Direct properties for `access_info.vault_auth_info`

- [app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022): complete subsection reference.

- [token](data-sources--secret_management_access--reference--group-001.md#canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203): complete subsection reference.

<a id="canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.vault_auth_info.app_role_auth` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- access_info.vault_auth_info.app_role_auth

<a id="canonical-2003300311230032-3101330022111120-0301320313012000-3332130101131323-2013122012110023-2021212303131033-3101112300131203-0330303120003020"></a>

Type: `"single"`. Computed.

AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.

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

<a id="canonical-2220022102021121-1200031222213311-2122122001000121-1031332100323030-1232321301110331-1032332203112000-3203210300120223-3232101132333322"></a>

### Direct properties for `access_info.vault_auth_info.app_role_auth`

<a id="canonical-3003333122011202-0301020320203102-1303032111133312-0322311230213230-2323121323310331-1030300103102231-1212321312021333-0022030321023212"></a>

#### `access_info.vault_auth_info.app_role_auth.role_id` property

Type: `"string"`. Computed.

Role ID. Role-ID to be used for authentication.

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

- [secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020): complete subsection reference.

<a id="canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.vault_auth_info.app_role_auth.secret_id` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022)
- access_info.vault_auth_info.app_role_auth.secret_id

<a id="canonical-2301133211213110-3003223123030322-1223012110113120-3210202011132312-1301110333011231-0033210033323113-1211310200120311-2110123233032210"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1222300321231010-0110301302231121-2012002000220002-2112311133313000-2231011000100102-1122001002230311-2320203021012030-2230130131002113"></a>

### Direct properties for `access_info.vault_auth_info.app_role_auth.secret_id`

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-1313201332032131-2130101023010103-3323213031000023-0211023332102230-1100103002222230-2202201123302213-0303312223222203-1020132213130330): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3210131331231023-2130202300320212-3120220201021330-1100313132122113-3000202123200023-3001310010123021-0110223232203320-0003110101320221): complete subsection reference.

<a id="canonical-1313201332032131-2130101023010103-3323213031000023-0211023332102230-1100103002222230-2202201123302213-0303312223222203-1020132213130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022)
- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020)
- access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info

<a id="canonical-1123031301231032-3020202231313010-2332313230102321-2121322011001102-2301300313333300-2220111003331022-1133011301302331-2312210000121001"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1221320333303111-3312322133100232-0031223312003011-2221303203200110-2133313322222302-3102223101302023-1310311112211122-0230132031010012"></a>

### Direct properties for `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info`

<a id="canonical-3332221120110303-3131221100012212-0203120011321300-1003320322300312-2203113102020113-3021022100231131-0031322110102001-2123230021000022"></a>

#### `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-2012012203201112-2213202132201311-3333332010202020-1330310201300220-1001332123001010-2322133103313313-0220302223232023-1330130021211123"></a>

<a id="canonical-1002232002201200-2220012212103111-1003031231310200-1012300020111203-1002300022133121-2312221131202332-3020303300133220-2032033113202330"></a>

#### `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-1212122113102213-2312200130202113-0230330303320033-1022231322322123-0211110333130303-1321113133213112-1113111100012203-2121001021323130"></a>

<a id="canonical-3002320321112130-0101332223120013-0311020211000301-3113321102220023-0001331012121233-3222322032322222-3033232230113130-0120110023323201"></a>

#### `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3210131331231023-2130202300320212-3120220201021330-1100313132122113-3000202123200023-3001310010123021-0110223232203320-0003110101320221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022)
- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020)
- access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info

<a id="canonical-3100233101303200-3021132120010212-1222112320313312-0321113032123231-0010020012201210-1300211221213200-3021001030132220-3210301222211221"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2003210301303323-3013233232112322-0232210012033213-0221033232010301-0220333110032121-2211031230102122-3232233123130130-1200322232223331"></a>

### Direct properties for `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info`

<a id="canonical-0111110323030220-1221133112102302-3020312111000201-0311122000330210-0330112102332332-3120022101211203-0303313312033320-0003121122320122"></a>

#### `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0301122312022223-0320030333010321-3201331123120332-0111023123113100-0130203323013232-1132031333302111-0231322123002210-3113321120000333"></a>

<a id="canonical-0202133003031130-1103202200012101-3003033203223031-3301030132000000-1031321113030131-1130033111020112-3112100130023103-0121302113010232"></a>

#### `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.vault_auth_info.token` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- access_info.vault_auth_info.token

<a id="canonical-0022122320212220-3202033221010132-0222312101211131-2321320311103112-2202101021123031-2230111222301311-1222303121333330-2300210021233310"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1320302113303233-3300010302002100-2110000000023020-2332211221302101-0132030333012323-3312220302030011-2211103323332233-0100211020332322"></a>

### Direct properties for `access_info.vault_auth_info.token`

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3322012010221021-3220213301022210-3120103012201323-0210322010330132-3001131332202320-3212201213112200-1131211210133311-3303103222123010): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-1032012332233213-1220021321001313-2132023132202000-1000030120323133-3112100120311330-1100023213013300-1000230130203301-1020101322330113): complete subsection reference.

<a id="canonical-3322012010221021-3220213301022210-3120103012201323-0210322010330132-3001131332202320-3212201213112200-1131211210133311-3303103222123010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.vault_auth_info.token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-001.md#canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203)
- access_info.vault_auth_info.token.blindfold_secret_info

<a id="canonical-2000101120331303-3230012210023113-2302112301002300-1312033301222103-2033032211313030-1102303102002311-0233021232002302-3132030301020322"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3012103111030000-0322203133300023-3221112221133100-3000322220313003-1030200220030313-0231312202033333-0103110231103023-2233201220132132"></a>

### Direct properties for `access_info.vault_auth_info.token.blindfold_secret_info`

<a id="canonical-2123103202031232-3023230213333130-2111033331311223-0021031013201033-0012313001301213-0303213111300011-0233012332003320-0332021212303101"></a>

#### `access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-2330203333102013-1201020331222013-3103132332000221-2222301231231220-1100023022033331-1031203020102231-3320313131030333-0112210330131331"></a>

<a id="canonical-1100220200030023-3333012332112021-1120122302330210-0102000322332103-0322021301000031-0222233233113210-3012232322203130-2031132112332211"></a>

#### `access_info.vault_auth_info.token.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-1110201202123120-1123130112012000-1123231320031230-2232133100301222-3211303011333001-2210111230022023-1201303322113212-3213103321300030"></a>

<a id="canonical-0310003222331202-0220003122202110-3202301132300001-2110031321313222-2001202222313232-1101132211123230-1012112012300012-1103333133030122"></a>

#### `access_info.vault_auth_info.token.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1032012332233213-1220021321001313-2132023132202000-1000030120323133-3112100120311330-1100023213013300-1000230130203301-1020101322330113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `access_info.vault_auth_info.token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-001.md#canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203)
- access_info.vault_auth_info.token.clear_secret_info

<a id="canonical-1322100230302323-2211102110110000-0101322312020333-1300211312023200-3321301023133030-1330203320200110-0001032011213110-3213020033123313"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2112122023211311-1321102010333133-0001120110133113-2330013030233110-0111132132000103-2210011111131022-0103202231123301-1001133031011233"></a>

### Direct properties for `access_info.vault_auth_info.token.clear_secret_info`

<a id="canonical-0010303232000132-0233132320002133-3220031313120133-1311103230331022-2010130101232303-0003031133222233-0300313333132302-1020113100331111"></a>

#### `access_info.vault_auth_info.token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2123203212021322-2130123100211003-1133310310010113-0222011121122013-2113321003002302-3310303312032130-2010110013012303-1010030222321013"></a>

<a id="canonical-1133300133020102-0213021003203023-0310230123221331-0330132031230211-0311322031113011-0121011222212331-1032133033200232-2033303232022113"></a>

#### `access_info.vault_auth_info.token.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- where

<a id="canonical-3002002312300203-3333333000211233-1230031002223332-2233310003333231-0321230033230330-2132001131033032-1210230322002222-0123221101110303"></a>

Type: `"single"`. Computed.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

<a id="canonical-3122321222122310-2120120322130020-3013021302203333-3030223110001002-2303323211020311-0001121231012303-0010312213321012-0001231112130021"></a>

### Direct properties for `where`

- [site](data-sources--secret_management_access--reference--group-001.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321): complete subsection reference.

- [virtual_network](data-sources--secret_management_access--reference--group-001.md#canonical-1010100312222110-1101123221213010-1320012022131302-2121232311200201-0133002000213230-3132212101210021-0321212330221100-1322311110200013): complete subsection reference.

- [virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212): complete subsection reference.

<a id="canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- where.site

<a id="canonical-3332301333021022-3213033023310121-1000023201332233-3131300113321222-2103321032312310-0030301012000231-3233102010210212-1120213112220330"></a>

Type: `"single"`. Computed.

This specifies a direct reference to a site configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-1022203332102202-1303321303031203-1013302032002312-0332122121012133-1210111310122120-0003210103001322-2300022033001001-2311102200113230"></a>

### Direct properties for `where.site`

- [disable_internet_vip](data-sources--secret_management_access--reference--group-001.md#canonical-0020112212231223-0132311323130321-1120222122103021-0120213121301200-3113000133332023-2332011232112303-3330133000113322-3001220111210031): complete subsection reference.

- [enable_internet_vip](data-sources--secret_management_access--reference--group-001.md#canonical-2111130133133231-1331101233332023-0113032000223033-1012103203012032-3011031332333030-1132330302303130-0330112330331310-3021223003231021): complete subsection reference.

<a id="canonical-2210021123301000-3323031211313332-1301202310021222-3311220210320320-1222331201230200-3112100013321230-0021221113101111-2120133122103022"></a>

<a id="canonical-1322330321000133-2021010333013111-1302312011121221-1321101211032013-0032221102212103-2231002101001221-1320130303023203-3111313212132301"></a>

#### `where.site.network_type` property

Type: `"string"`. Computed.

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

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

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

- [ref](data-sources--secret_management_access--reference--group-001.md#canonical-3322223332330102-2330232001102203-2213010010102323-3200103311331011-0100130131321202-0222213330321201-0100121301230212-1331121303010101): complete subsection reference.

<a id="canonical-0020112212231223-0132311323130321-1120222122103021-0120213121301200-3113000133332023-2332011232112303-3330133000113322-3001220111210031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.site](data-sources--secret_management_access--reference--group-001.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- where.site.disable_internet_vip

<a id="canonical-2012121010311023-0320322323221133-2231213331200223-0303303332211213-0300022003020012-2113213130312000-0131000211021201-0220110313303302"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111130133133231-1331101233332023-0113032000223033-1012103203012032-3011031332333030-1132330302303130-0330112330331310-3021223003231021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.site](data-sources--secret_management_access--reference--group-001.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- where.site.enable_internet_vip

<a id="canonical-1111122123111022-3033221332301213-0100023031220221-1221230002220021-0302100212310033-2001123033002000-2331010113112222-2120221221310123"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322223332330102-2330232001102203-2213010010102323-3200103311331011-0100130131321202-0222213330321201-0100121301230212-1331121303010101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.ref` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.site](data-sources--secret_management_access--reference--group-001.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- where.site.ref

<a id="canonical-2320023130132003-3222300321212333-0111020000010222-2033323012310311-2201013331210103-3033122023121112-1030120122210330-0031121111132021"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3113313213100323-0233131322133001-2221321211003313-2201001332003123-0112100202112020-2220022320330013-1212102122032010-1133131211131103"></a>

### Direct properties for `where.site.ref`

<a id="canonical-3223132211021303-3233223303220121-1101102203131010-3000103220102132-3133111113201222-0333301320020230-1100211301310112-0231011302000102"></a>

#### `where.site.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-0030100011312200-3032013002233113-3313003313012123-1000310112313100-3223131023210120-2202311112123023-1030330031222233-3311212202002223"></a>

<a id="canonical-1323113113332311-1133303213320101-2032202111002211-2212023023010303-2232123013223232-2311021022310032-1321303310210033-2102121011010322"></a>

#### `where.site.ref.name` property

Type: `"string"`. Computed.

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

<a id="canonical-0122223323201200-0210313020303120-0111132033302032-0300113022210113-3033033323121122-0033112310031321-2213033101303330-1311301302301023"></a>

<a id="canonical-2320111310030010-3110220230332201-3031222021331220-3311202012030312-2301032323301331-3323210121001232-0312111301320301-2123302133031101"></a>

#### `where.site.ref.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-1133311323211200-3113102312303032-2210222003122001-3130233331203020-2103000302023203-2001012123133332-2202013331202303-3210301321330302"></a>

<a id="canonical-2303202103302231-0121112022020232-0232323001100111-1232202212310112-3031223131211101-0211213210211330-3122201312023330-1210213332000121"></a>

#### `where.site.ref.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3211200102122130-3011210131320233-3031132330313121-3032221130032300-1110221212131010-3110133100313223-1323102122201312-3332231020120212"></a>

<a id="canonical-3310112032113320-2221112302223011-2211323231330020-2203222323102301-0131211322113033-2223110132321022-0303001103032133-1201230201333302"></a>

#### `where.site.ref.uid` property

Type: `"string"`. Computed.

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

<a id="canonical-1010100312222110-1101123221213010-1320012022131302-2121232311200201-0133002000213230-3132212101210021-0321212330221100-1322311110200013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- where.virtual_network

<a id="canonical-0330112231003221-3223230112200210-2101230313300030-0003230131212010-2032120121313002-1012323320223003-1310000313322201-1010102300232330"></a>

Type: `"single"`. Computed.

This specifies a direct reference to a network configuration object.

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

<a id="canonical-0033210230130022-2121223303312102-1012001023332201-0331302120000311-0120130320133101-0021023222330123-2313301311131233-0223203220122213"></a>

### Direct properties for `where.virtual_network`

- [ref](data-sources--secret_management_access--reference--group-001.md#canonical-2321122003110310-1200023331323221-3113213010023032-0123112030221013-1132302221123030-2113312011101323-3130330322130320-1231001033133300): complete subsection reference.

<a id="canonical-2321122003110310-1200023331323221-3113213010023032-0123112030221013-1132302221123030-2113312011101323-3130330322130320-1231001033133300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network.ref` properties

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-001.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.virtual_network](data-sources--secret_management_access--reference--group-001.md#canonical-1010100312222110-1101123221213010-1320012022131302-2121232311200201-0133002000213230-3132212101210021-0321212330221100-1322311110200013)
- where.virtual_network.ref

<a id="canonical-1130233211123201-0120130313000322-0111210222311320-2022133233103301-3212131022333332-3131312323313102-2002132023230033-3100210201020200"></a>

Type: `"list"`. Computed.

Reference. A virtual network direct reference.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2210132132020100-0102103001001223-1101222222113322-0303202011200222-1003100131333113-1211333120332223-1123020011011321-2233203323130301"></a>

### Direct properties for `where.virtual_network.ref`

<a id="canonical-1213000012233011-3321212101033221-0310320123202300-0102231032000321-2233030002033002-3000201102111311-3123210201300121-0130233331122330"></a>

#### `where.virtual_network.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-2310132331232320-3333330011323000-3123213032332310-3121321221202020-0120003012300221-2202133102201030-2012032323000331-0320233321222013"></a>

<a id="canonical-3301223213131003-3213101123001132-2320031023203000-3310001023222032-1320301100310033-1001102223312332-2003000112032303-1313303030003321"></a>

#### `where.virtual_network.ref.name` property

Type: `"string"`. Computed.

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

<a id="canonical-1030131032213201-1113312220113313-3220321110301301-2001003223320010-2013333020133113-2103121320000313-0103010110323003-3033230301331000"></a>

<a id="canonical-2010313333023310-0313220031211202-2211303310031203-2013010022101321-3110000223022113-1331212020310101-1110131121201012-1033001002110010"></a>

#### `where.virtual_network.ref.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-0010230300222313-0310010122111122-1003033123032201-2211303032310212-2102212132022321-0233321120212131-2213012203021101-3112021032331111"></a>

<a id="canonical-0233101033103200-2222033102201111-1032213322203033-3002221332123322-3121030213011032-3101133302212110-2030011220010010-0332001123203202"></a>

#### `where.virtual_network.ref.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-1301030132220110-3202201203323322-2132223330211023-2030020002121122-3013220213231130-2222003131220322-2022102101023310-1203001102213032"></a>

<a id="canonical-1233233010300302-1232303311200223-3232332310220301-3031021102133302-2002110213212022-2103202200020020-0111330033122303-1303312323211123"></a>

#### `where.virtual_network.ref.uid` property

Type: `"string"`. Computed.

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
