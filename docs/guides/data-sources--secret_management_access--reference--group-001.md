---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210232011222331-3231121320001131-1331102131312302-2002321021120122-3011102202101012-0310331313333112-2331230012000330-3103023301000233"></a>

## Property reference — Property reference / 003011133013 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- Property reference

<a id="canonical-0101013131023012-1210103121021002-2211321010120112-0230033120303303-0322303122313133-2123010231123110-3232021021231211-0131321201301331"></a>

## Direct properties — Property reference / 003011133013 / 3

- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133): complete subsection reference.

<a id="canonical-0020022133311021-2022112212231122-3002022022213000-2321300310310202-1303120203321303-1111311002032120-3232000222223133-3213020221113313"></a>

<a id="canonical-1100230113200012-3201210012331203-2111031102013322-1222233020103302-1101010122223001-2211020310023301-0121222322112100-1122213230003123"></a>

## annotations property — Property reference / 003011133013 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
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

<a id="canonical-3122210100112130-1313303203321031-2030323222331321-0330210013020311-1031201211022211-0002022031330033-3031210210201132-3022033203223031"></a>

<a id="canonical-0331201101330010-1333023202023120-2120020111310333-3223032131113100-1112122323020012-2100021030033020-0310132010300323-1301103303320120"></a>

## description property — Property reference / 003011133013 / 5

Type: `"string"`. Computed.

Description of the SecretManagementAccess.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3132021103322203-0313313332322330-0123110301111013-1122130322032112-1110120213212302-2020003012223133-2023103100223230-2132122200332300"></a>

## ID property — Property reference / 003011133013 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1223222333112320-1303011301113202-0103111032130113-3023313003111001-2000231233330002-2211033023210132-3312110223232331-3222000022020112"></a>

<a id="canonical-2223101232302001-3010312002211323-2321330021000003-1020320323012011-3010222013330033-3121121112133321-1302201222021321-1001100220003302"></a>

## labels property — Property reference / 003011133013 / 7

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

<a id="canonical-0221012311032323-1211103132120012-0010211303010031-1103303112001001-2211232232020232-3023112313212212-0133010211003102-2322322313200312"></a>

<a id="canonical-0122100132131010-0303331301103111-3221333323212122-0233113231112321-0200000033212223-3023313112000122-2230131120020130-0221012121102223"></a>

## name property — Property reference / 003011133013 / 8

Type: `"string"`. Required.

Name of the SecretManagementAccess.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2012121111212300-1032102130220211-1032210310100212-3300002331020120-1002010320221330-3310033130011110-1233223313211132-1010112203011132"></a>

## namespace property — Property reference / 003011133013 / 9

Type: `"string"`. Required.

Namespace where the SecretManagementAccess exists.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1302122102332100-1211222111002301-3332210310230100-1300310101110122-1033323022300312-1033202021301013-1013223121213230-1010131222101300"></a>

## provider_name property — Property reference / 003011133013 / 10

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101): complete subsection reference.

<a id="canonical-2113302301211330-0020312311323211-1213322122223211-2103011002203200-3223200330220122-0332111112332032-2321111110031231-0012301022001123"></a>

## All schema paths — Property reference / 003011133013 / 11

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
| `access_info.vault_auth_info.app_role_auth` | [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-002.md#canonical-2003300311230032-3101330022111120-0301320313012000-3332130101131323-2013122012110023-2021212303131033-3101112300131203-0330303120003020) |
| `access_info.vault_auth_info.app_role_auth.role_id` | [access_info.vault_auth_info.app_role_auth.role_id](data-sources--secret_management_access--reference--group-002.md#canonical-3003333122011202-0301020320203102-1303032111133312-0322311230213230-2323121323310331-1030300103102231-1212321312021333-0022030321023212) |
| `access_info.vault_auth_info.app_role_auth.secret_id` | [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-002.md#canonical-2301133211213110-3003223123030322-1223012110113120-3210202011132312-1301110333011231-0033210033323113-1211310200120311-2110123233032210) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-1123031301231032-3020202231313010-2332313230102321-2121322011001102-2301300313333300-2220111003331022-1133011301302331-2312210000121001) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-002.md#canonical-3332221120110303-3131221100012212-0203120011321300-1003320322300312-2203113102020113-3021022100231131-0031322110102001-2123230021000022) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-002.md#canonical-2012012203201112-2213202132201311-3333332010202020-1330310201300220-1001332123001010-2322133103313313-0220302223232023-1330130021211123) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-002.md#canonical-1212122113102213-2312200130202113-0230330303320033-1022231322322123-0211110333130303-1321113133213112-1113111100012203-2121001021323130) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-3100233101303200-3021132120010212-1222112320313312-0321113032123231-0010020012201210-1300211221213200-3021001030132220-3210301222211221) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-002.md#canonical-0111110323030220-1221133112102302-3020312111000201-0311122000330210-0330112102332332-3120022101211203-0303313312033320-0003121122320122) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url](data-sources--secret_management_access--reference--group-002.md#canonical-0301122312022223-0320030333010321-3201331123120332-0111023123113100-0130203323013232-1132031333302111-0231322123002210-3113321120000333) |
| `access_info.vault_auth_info.token` | [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-0022122320212220-3202033221010132-0222312101211131-2321320311103112-2202101021123031-2230111222301311-1222303121333330-2300210021233310) |
| `access_info.vault_auth_info.token.blindfold_secret_info` | [access_info.vault_auth_info.token.blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-2000101120331303-3230012210023113-2302112301002300-1312033301222103-2033032211313030-1102303102002311-0233021232002302-3132030301020322) |
| `access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider](data-sources--secret_management_access--reference--group-002.md#canonical-2123103202031232-3023230213333130-2111033331311223-0021031013201033-0012313001301213-0303213111300011-0233012332003320-0332021212303101) |
| `access_info.vault_auth_info.token.blindfold_secret_info.location` | [access_info.vault_auth_info.token.blindfold_secret_info.location](data-sources--secret_management_access--reference--group-002.md#canonical-2330203333102013-1201020331222013-3103132332000221-2222301231231220-1100023022033331-1031203020102231-3320313131030333-0112210330131331) |
| `access_info.vault_auth_info.token.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.store_provider](data-sources--secret_management_access--reference--group-002.md#canonical-1110201202123120-1123130112012000-1123231320031230-2232133100301222-3211303011333001-2210111230022023-1201303322113212-3213103321300030) |
| `access_info.vault_auth_info.token.clear_secret_info` | [access_info.vault_auth_info.token.clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-1322100230302323-2211102110110000-0101322312020333-1300211312023200-3321301023133030-1330203320200110-0001032011213110-3213020033123313) |
| `access_info.vault_auth_info.token.clear_secret_info.provider_ref` | [access_info.vault_auth_info.token.clear_secret_info.provider_ref](data-sources--secret_management_access--reference--group-002.md#canonical-0010303232000132-0233132320002133-3220031313120133-1311103230331022-2010130101232303-0003031133222233-0300313333132302-1020113100331111) |
| `access_info.vault_auth_info.token.clear_secret_info.url` | [access_info.vault_auth_info.token.clear_secret_info.url](data-sources--secret_management_access--reference--group-002.md#canonical-2123203212021322-2130123100211003-1133310310010113-0222011121122013-2113321003002302-3310303312032130-2010110013012303-1010030222321013) |
| `annotations` | [annotations](data-sources--secret_management_access--reference--group-001.md#canonical-0020022133311021-2022112212231122-3002022022213000-2321300310310202-1303120203321303-1111311002032120-3232000222223133-3213020221113313) |
| `description` | [description](data-sources--secret_management_access--reference--group-001.md#canonical-3122210100112130-1313303203321031-2030323222331321-0330210013020311-1031201211022211-0002022031330033-3031210210201132-3022033203223031) |
| `id` | [ID](data-sources--secret_management_access--reference--group-001.md#canonical-3300331133320113-3021133221213331-3322100223310211-0010012221331102-2232301032010230-3332203302330332-3233103133102310-2230101011212021) |
| `labels` | [labels](data-sources--secret_management_access--reference--group-001.md#canonical-1223222333112320-1303011301113202-0103111032130113-3023313003111001-2000231233330002-2211033023210132-3312110223232331-3222000022020112) |
| `name` | [name](data-sources--secret_management_access--reference--group-001.md#canonical-0221012311032323-1211103132120012-0010211303010031-1103303112001001-2211232232020232-3023112313212212-0133010211003102-2322322313200312) |
| `namespace` | [namespace](data-sources--secret_management_access--reference--group-001.md#canonical-1211001101231101-2231021223312012-1231030031230121-1313210332331310-3113333033022033-2212221201122022-3230113201003223-3232311111031233) |
| `provider_name` | [provider_name](data-sources--secret_management_access--reference--group-001.md#canonical-1100332300303212-0302032103100023-2012103030312130-3011032122333203-3100102012022131-0101201213232200-1130000200313130-3111201031320322) |
| `where` | [where](data-sources--secret_management_access--reference--group-002.md#canonical-3002002312300203-3333333000211233-1230031002223332-2233310003333231-0321230033230330-2132001131033032-1210230322002222-0123221101110303) |
| `where.site` | [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-3332301333021022-3213033023310121-1000023201332233-3131300113321222-2103321032312310-0030301012000231-3233102010210212-1120213112220330) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-2012121010311023-0320322323221133-2231213331200223-0303303332211213-0300022003020012-2113213130312000-0131000211021201-0220110313303302) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-1111122123111022-3033221332301213-0100023031220221-1221230002220021-0302100212310033-2001123033002000-2331010113112222-2120221221310123) |
| `where.site.network_type` | [where.site.network_type](data-sources--secret_management_access--reference--group-002.md#canonical-2210021123301000-3323031211313332-1301202310021222-3311220210320320-1222331201230200-3112100013321230-0021221113101111-2120133122103022) |
| `where.site.ref` | [where.site.ref](data-sources--secret_management_access--reference--group-002.md#canonical-2320023130132003-3222300321212333-0111020000010222-2033323012310311-2201013331210103-3033122023121112-1030120122210330-0031121111132021) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--secret_management_access--reference--group-002.md#canonical-3223132211021303-3233223303220121-1101102203131010-3000103220102132-3133111113201222-0333301320020230-1100211301310112-0231011302000102) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--secret_management_access--reference--group-002.md#canonical-0030100011312200-3032013002233113-3313003313012123-1000310112313100-3223131023210120-2202311112123023-1030330031222233-3311212202002223) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--secret_management_access--reference--group-002.md#canonical-0122223323201200-0210313020303120-0111132033302032-0300113022210113-3033033323121122-0033112310031321-2213033101303330-1311301302301023) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--secret_management_access--reference--group-002.md#canonical-1133311323211200-3113102312303032-2210222003122001-3130233331203020-2103000302023203-2001012123133332-2202013331202303-3210301321330302) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--secret_management_access--reference--group-002.md#canonical-3211200102122130-3011210131320233-3031132330313121-3032221130032300-1110221212131010-3110133100313223-1323102122201312-3332231020120212) |
| `where.virtual_network` | [where.virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-0330112231003221-3223230112200210-2101230313300030-0003230131212010-2032120121313002-1012323320223003-1310000313322201-1010102300232330) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--secret_management_access--reference--group-002.md#canonical-1130233211123201-0120130313000322-0111210222311320-2022133233103301-3212131022333332-3131312323313102-2002132023230033-3100210201020200) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--secret_management_access--reference--group-002.md#canonical-1213000012233011-3321212101033221-0310320123202300-0102231032000321-2233030002033002-3000201102111311-3123210201300121-0130233331122330) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--secret_management_access--reference--group-002.md#canonical-2310132331232320-3333330011323000-3123213032332310-3121321221202020-0120003012300221-2202133102201030-2012032323000331-0320233321222013) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--secret_management_access--reference--group-002.md#canonical-1030131032213201-1113312220113313-3220321110301301-2001003223320010-2013333020133113-2103121320000313-0103010110323003-3033230301331000) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--secret_management_access--reference--group-002.md#canonical-0010230300222313-0310010122111122-1003033123032201-2211303032310212-2102212132022321-0233321120212131-2213012203021101-3112021032331111) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--secret_management_access--reference--group-002.md#canonical-1301030132220110-3202201203323322-2132223330211023-2030020002121122-3013220213231130-2222003131220322-2022102101023310-1203001102213032) |
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

<a id="canonical-3022001023332200-1132333222012120-3121301332302011-0103223031201010-0210022221031300-2220020320030021-0031311200001203-0003201221022221"></a>

## Next pages — Property reference / 003011133013 / 12

- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232333011113313-0311231032020023-2032212231000230-3311201103120113-2120103332010222-3213300003312313-3303002120130200-2120023333332100"></a>

## access_info — access_info / 111213221223 / 2

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

<a id="canonical-2101222120130202-1222121213013232-1123100312010200-1330021222213030-2322303320103303-0020323000312122-2220321133313132-3320223303102023"></a>

## Direct properties — access_info / 111213221223 / 3

- [rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131): complete subsection reference.

<a id="canonical-3032302011210022-1121311233220103-1230322100222110-1120213120331101-1322101210231233-2201300303331003-0332130121010211-0012001232210210"></a>

<a id="canonical-1112302011022012-0200101113202210-3233121101001213-0011023202220302-1111302232131320-3031310110303112-2010133132102031-2220033010120111"></a>

## scheme property — access_info / 111213221223 / 4

Type: `"string"`. Computed.

\[Enum: HTTP|HTTPS\] SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.
Possible values are \`HTTP\`, \`HTTPS\`. Defaults to \`HTTP\`.

Upstream description:

SchemeType is used to indicate URL scheme

HTTP:// scheme HTTPS:// scheme.

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

<a id="canonical-1300110030113011-0011301212012223-3211010323102110-3300020110211301-1113033312230012-3310011130110322-2313323030200300-2021031202102203"></a>

## server_endpoint property — access_info / 111213221223 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1212332201002123-1113211301122330-1103001020033012-0302332323322021-2330023210131120-0211320102211311-1032231010133300-2112323002121331"></a>

## Next pages — access_info / 111213221223 / 6

- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312102303012222-0231230333202123-0032012302322313-1120320101022130-0221112122331132-3013120330230200-2203110103111133-2022200212122211"></a>

## access_info.rest_auth_info — rest_auth_info / 302223020230 / 2

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

<a id="canonical-2123300331122300-0011130223002112-2313113032321223-3323223320011110-0122232113101103-1020220330002131-1323001133312123-1220111002021223"></a>

## Direct properties — rest_auth_info / 302223020230 / 3

- [basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301): complete subsection reference.

- [headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2220011133213131-3101131111102021-2320030231133223-0001320232332110-3230133310022213-3033122313002031-1022302012302010-2123102011221030): complete subsection reference.

- [query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-3123211020031221-1033310323003001-3333320020332010-2311112310201202-2122332031233311-1221220333200110-1203202212131100-1123210312001010): complete subsection reference.

<a id="canonical-3022230001322200-1030212330310002-3121033033122100-0101131232012021-1101123030200311-3002012213100110-3223010321211312-3010123210320111"></a>

## Next pages — rest_auth_info / 302223020230 / 4

- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301)
- [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2220011133213131-3101131111102021-2320030231133223-0001320232332110-3230133310022213-3033122313002031-1022302012302010-2123102011221030)
- [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-3123211020031221-1033310323003001-3333320020332010-2311112310201202-2122332031233311-1221220333200110-1203202212131100-1123210312001010)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121221020332322-1322202232320331-2233133020012023-1023320022221012-1332322021131210-2111103020321300-0120300020211113-2212321022021202"></a>

## access_info.rest_auth_info.basic_auth — basic_auth / 032111211311 / 2

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

<a id="canonical-0030233000120332-0100323102333123-3223233033022101-3013012331223323-1000212033010133-0023121300303300-2111303323330110-3303223221312011"></a>

## Direct properties — basic_auth / 032111211311 / 3

- [password](data-sources--secret_management_access--reference--group-001.md#canonical-0232033033232332-3300113133233020-2200212102320012-0212102103200202-2100020112220023-1123033031113112-3222110212233103-0101132102132212): complete subsection reference.

<a id="canonical-0233323213032201-1112203333110232-1333303130223032-3322300201133202-0200132013313100-2303233000222311-3121313330300213-1111232331033302"></a>

<a id="canonical-1123331201221131-1020221032021010-2213130003110122-0012331121312123-1033021033010230-1332002023111123-2132001003001330-0331221013120022"></a>

## username property — basic_auth / 032111211311 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3101233302112103-1201311113220100-0232011101011000-3212132110322120-3301112232210131-1020122002300301-2202221313100210-0330123332302211"></a>

## Next pages — basic_auth / 032111211311 / 5

- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-0232033033232332-3300113133233020-2200212102320012-0212102103200202-2100020112220023-1123033031113112-3222110212233103-0101132102132212)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0232033033232332-3300113133233020-2200212102320012-0212102103200202-2100020112220023-1123033031113112-3222110212233103-0101132102132212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233322331202230-1120022330232212-2200022021132112-2113001211131231-1232000301130010-1311020001213020-1233113220022031-1332330101122132"></a>

## access_info.rest_auth_info.basic_auth.password — password / 123322301121 / 2

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

<a id="canonical-3110330211111101-3131202203101331-3001223122121220-2213311023011310-0121230103130123-1020210121000223-2131221302233200-0321322213313231"></a>

## Direct properties — password / 123322301121 / 3

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3330003132112301-3201302110220203-2321021101021220-3110132003321121-0303310323130333-0211101321212113-0012102100223031-2232231033213330): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3010333002113202-3302323331302311-1031121123121330-2320131130223022-3210312232203203-2110300300102023-0010023313013122-3333303020301003): complete subsection reference.

<a id="canonical-3300023011001211-0102001112233031-1013313313320321-1302032012230301-3310112321231300-1222111302101013-3330313330122132-3311132032313001"></a>

## Next pages — password / 123322301121 / 4

- [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3330003132112301-3201302110220203-2321021101021220-3110132003321121-0303310323130333-0211101321212113-0012102100223031-2232231033213330)
- [access_info.rest_auth_info.basic_auth.password.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-3010333002113202-3302323331302311-1031121123121330-2320131130223022-3210312232203203-2110300300102023-0010023313013122-3333303020301003)
- [access_info.rest_auth_info.basic_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2332210232020211-2222222113310221-0101203303121203-2133122120101003-2122202120203322-3300313221010130-2220113232032322-3003233230300301)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3330003132112301-3201302110220203-2321021101021220-3110132003321121-0303310323130333-0211101321212113-0012102100223031-2232231033213330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123300221332331-3213213333203231-1300221103112320-0000000023230202-3212001111102331-1203002023021122-0333320031331331-3112001133100123"></a>

## access_info.rest_auth_info.basic_auth.password.blindfold_secret_info — blindfold_secret_info / 320132320030 / 2

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

<a id="canonical-0211223201001103-2012120002313002-1203312002120101-2200123310121113-3333021011222101-1232213312031200-0213023330112220-3331002003033223"></a>

## Direct properties — blindfold_secret_info / 320132320030 / 3

<a id="canonical-3100331020123101-3220010230212110-1200302130233032-0032223323101331-1012002311110022-2321013100102001-1330302012213330-1101122023303313"></a>

<a id="canonical-3122112233112203-0303223300121023-2110331223321300-1001201332133230-2032022202331231-0330031301023012-1021303210123023-3320200202020003"></a>

## decryption_provider property — blindfold_secret_info / 320132320030 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3200113303021112-1001003301211200-2021311112032123-2020112003232000-0130223201111223-3322133313312232-0203130333033130-0120132331121112"></a>

## location property — blindfold_secret_info / 320132320030 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0311220013232201-2310132020030112-3130203103120203-3212021320102310-2323333211323100-2032222022223313-1331312211031300-2110020301331121"></a>

## store_provider property — blindfold_secret_info / 320132320030 / 6

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0303112322011131-1310220322012301-3000313123102320-0232001103320113-1211123232132321-3022312221320301-3230111122120321-1303320131233101"></a>

## Next pages — blindfold_secret_info / 320132320030 / 7

- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-0232033033232332-3300113133233020-2200212102320012-0212102103200202-2100020112220023-1123033031113112-3222110212233103-0101132102132212)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3010333002113202-3302323331302311-1031121123121330-2320131130223022-3210312232203203-2110300300102023-0010023313013122-3333303020301003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300303023321103-2133121121101113-2133132302210213-2030211231333201-0012022203031112-2023231122321223-0000012232300312-2002233200100322"></a>

## access_info.rest_auth_info.basic_auth.password.clear_secret_info — clear_secret_info / 121132003332 / 2

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

<a id="canonical-3102100013101022-3210121011332003-3112120312110323-2102320303113301-1212100130022103-0331212321310101-3011012302233301-1301310032330321"></a>

## Direct properties — clear_secret_info / 121132003332 / 3

<a id="canonical-0001101123000030-1003112211133130-0132330022013332-0322101012122110-0230230001001011-2002330302003120-1102002213213313-0032331101200101"></a>

<a id="canonical-2212011301331032-2012302330003223-3231013300013131-3220321100123023-2322122123213323-3221121013223031-1330203022202032-1120322331301233"></a>

## provider_ref property — clear_secret_info / 121132003332 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2030122203203233-3130301111300133-0101310302010032-2221033203333113-2322312021013221-2132213112021211-3310321010033232-3310121302213120"></a>

<a id="canonical-1112223321111203-3113121213010012-0010202131102120-0122232332323322-2023233031310031-2103320232223123-1331221011001212-2013122121333232"></a>

## URL property — clear_secret_info / 121132003332 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0222132210220320-0131022302002032-0330123223213220-3032000012213012-2103132233203232-3120000331310100-3300130202003232-2021030001220313"></a>

## Next pages — clear_secret_info / 121132003332 / 6

- [access_info.rest_auth_info.basic_auth.password](data-sources--secret_management_access--reference--group-001.md#canonical-0232033033232332-3300113133233020-2200212102320012-0212102103200202-2100020112220023-1123033031113112-3222110212233103-0101132102132212)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2220011133213131-3101131111102021-2320030231133223-0001320232332110-3230133310022213-3033122313002031-1022302012302010-2123102011221030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021331033031011-0313323023100121-3111331223210032-1232332121020322-1023000302311133-3022102312112122-1033033121320313-0113233132022221"></a>

## access_info.rest_auth_info.headers_auth — headers_auth / 032302133121 / 2

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

<a id="canonical-2212012332202033-2021023202110221-1202311020232031-0032112332332203-2102020020331032-1100032111100201-0231313012212110-2101312102230110"></a>

## Direct properties — headers_auth / 032302133121 / 3

- [headers](data-sources--secret_management_access--reference--group-001.md#canonical-1003003331312200-3323113111311113-0213222002100131-1121230213000312-0320022003312030-2011321112231130-0022312302021103-3230123311311122): complete subsection reference.

<a id="canonical-3010223001332202-1120032003322223-1223000030230012-3212333330210213-0130103312022323-3001033211112010-2202022211210103-2021230002133200"></a>

## Next pages — headers_auth / 032302133121 / 4

- [access_info.rest_auth_info.headers_auth.headers](data-sources--secret_management_access--reference--group-001.md#canonical-1003003331312200-3323113111311113-0213222002100131-1121230213000312-0320022003312030-2011321112231130-0022312302021103-3230123311311122)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1003003331312200-3323113111311113-0213222002100131-1121230213000312-0320022003312030-2011321112231130-0022312302021103-3230123311311122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311223020002220-1000212333113223-3221022023232310-1322100321203232-1202202122223203-0220100030312021-0212220103022301-3231222112133333"></a>

## access_info.rest_auth_info.headers_auth.headers — headers / 021100221110 / 2

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
    "category": "general",
    "constraintType": "object",
    "maxProperties": 16,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0023100331232122-2120301331011200-3230132110021000-2303211121031210-0213032031303020-3332313023012023-1100112101011123-2010222330311121"></a>

## Direct properties — headers / 021100221110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211213012022233-2101303132103033-3023230331333211-3232103010132321-0130020321002212-2333033033210213-3213110032232010-2330030013230133"></a>

## Next pages — headers / 021100221110 / 4

- [access_info.rest_auth_info.headers_auth](data-sources--secret_management_access--reference--group-001.md#canonical-2220011133213131-3101131111102021-2320030231133223-0001320232332110-3230133310022213-3033122313002031-1022302012302010-2123102011221030)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3123211020031221-1033310323003001-3333320020332010-2311112310201202-2122332031233311-1221220333200110-1203202212131100-1123210312001010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301000223002130-1212232322102102-0332122002330211-1320033320222113-3112033322321001-1011033302201330-0100131333331321-3111121011120302"></a>

## access_info.rest_auth_info.query_params_auth — query_params_auth / 220220021002 / 2

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

<a id="canonical-0332022102331211-0233211133110212-3010023231331013-1003020010203231-0310331221201231-3330333211333213-1212322301132131-1120123220330011"></a>

## Direct properties — query_params_auth / 220220021002 / 3

- [query_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131023122000323-2202200220101021-1332203131103313-2230102121223331-0010322101130331-0122002331202321-2113020132300322-2332112121123032): complete subsection reference.

<a id="canonical-2011030032000333-3332010331210033-1330302020301022-3303110120100220-0131123130023211-3120013100122221-1120320210110001-1132122212022203"></a>

## Next pages — query_params_auth / 220220021002 / 4

- [access_info.rest_auth_info.query_params_auth.query_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131023122000323-2202200220101021-1332203131103313-2230102121223331-0010322101130331-0122002331202321-2113020132300322-2332112121123032)
- [access_info.rest_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-1020312222021101-2300221223300003-1222330010013302-1002123323210312-2133212322121110-2131301202121000-3202203112232112-1201200230102131)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3131023122000323-2202200220101021-1332203131103313-2230102121223331-0010322101130331-0122002331202321-2113020132300322-2332112121123032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110323111232003-1023021101313100-0010102213222003-0322100010013320-3100303231311113-2332011301130211-3013121303313130-2111003022033210"></a>

## access_info.rest_auth_info.query_params_auth.query_params — query_params / 202332313110 / 2

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
    "category": "general",
    "constraintType": "object",
    "maxProperties": 16,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1231103022123202-1133030221221103-1310101120302201-1231220232122231-3111113221011231-3113221131230312-0310002322211212-1120113322112223"></a>

## Direct properties — query_params / 202332313110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300200122220013-2312122210021111-1112201102133000-0300032202311002-2033231303301033-3130110200203220-1233300202133100-1121331210033121"></a>

## Next pages — query_params / 202332313110 / 4

- [access_info.rest_auth_info.query_params_auth](data-sources--secret_management_access--reference--group-001.md#canonical-3123211020031221-1033310323003001-3333320020332010-2311112310201202-2122332031233311-1221220333200110-1203202212131100-1123210312001010)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221330301333100-3332210211021103-1021201102001110-2113012013303100-1323023122121103-2211203312300312-0330211202132020-0130031210322211"></a>

## access_info.tls_config — tls_config / 031200321020 / 2

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

<a id="canonical-1121233223030221-1030230200012031-1212323021021121-1323121231203022-2313210023011130-1112333003210203-3110030310212033-3322333012323332"></a>

## Direct properties — tls_config / 031200321020 / 3

- [cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223): complete subsection reference.

- [common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102): complete subsection reference.

- [default_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-1210103112123023-1311003031333303-0033320301002230-1333002123130023-2302111201030130-2311312223001120-3320000021000013-2132010033213110): complete subsection reference.

- [disable_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-0011323301322002-2201102122202131-2130020111322130-1323121110112002-3121221011021313-3000012122012233-3313000221030220-2013020223002321): complete subsection reference.

- [disable_sni](data-sources--secret_management_access--reference--group-001.md#canonical-3313002020213130-1211220100111301-2010321010111112-1221223002230323-2012201001201323-0033002301121033-0211321133002012-3230002322001001): complete subsection reference.

<a id="canonical-2032013333213230-3300320223333220-2033233103202101-3220311331110103-3121132030122232-2210011212311302-1021003203020221-1323111130232100"></a>

<a id="canonical-1330300001111332-0021023331213311-0000110011100133-2021322102221013-2330210113331230-0221322313031223-3033320222110231-0000213202003030"></a>

## max_session_keys property — tls_config / 031200321020 / 4

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2232033131321322-1111203100302312-2323303302030132-3331001210132102-2300323200031022-2332202202320221-3311110032302121-2122233222011210"></a>

## sni property — tls_config / 031200321020 / 5

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2030132132312203-2011320022213232-2311122131212022-0331300100203023-2301133222001330-0002011333122013-0132211102322311-0320311000201211"></a>

## Next pages — tls_config / 031200321020 / 6

- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [access_info.tls_config.default_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-1210103112123023-1311003031333303-0033320301002230-1333002123130023-2302111201030130-2311312223001120-3320000021000013-2132010033213110)
- [access_info.tls_config.disable_session_key_caching](data-sources--secret_management_access--reference--group-001.md#canonical-0011323301322002-2201102122202131-2130020111322130-1323121110112002-3121221011021313-3000012122012233-3313000221030220-2013020223002321)
- [access_info.tls_config.disable_sni](data-sources--secret_management_access--reference--group-001.md#canonical-3313002020213130-1211220100111301-2010321010111112-1221223002230323-2012201001201323-0033002301121033-0211321133002012-3230002322001001)
- [access_info.tls_config.use_host_header_as_sni](data-sources--secret_management_access--reference--group-001.md#canonical-1323322302212031-2102331223003212-2232132102320011-3313132032020021-3013331110313121-0310233210023203-0320100300121321-1033322112222012)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111331130001300-0101030321323331-1200123321322002-0123331130330132-3000101331001211-0221010203013301-2010303222303002-0033311220220322"></a>

## access_info.tls_config.cert_params — cert_params / 222103310111 / 2

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

<a id="canonical-2013011223332333-2102233202211232-0233200322302101-3111322121103323-0003021120222222-2233213202220311-3203311303032311-0032321002020030"></a>

## Direct properties — cert_params / 222103310111 / 3

- [certificates](data-sources--secret_management_access--reference--group-001.md#canonical-3223110302010102-0112021220103031-1023330021121210-2202021320210313-0100213121031322-2210312201321303-3230321311203132-1103022001111310): complete subsection reference.

<a id="canonical-1020030112031312-3000310203203331-2123232212001332-3123012331231322-0321300301231102-3232200303203103-1130020032322222-0232311032330011"></a>

<a id="canonical-0123200211200220-0223322012223302-0302322113321202-0231312300323021-2121332013132213-2201303021001020-2222031030233113-3121302123210320"></a>

## cipher_suites property — cert_params / 222103310111 / 4

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1232002031003030-2221320202032000-1223121031011003-3202103123123113-3201023310320311-3332201111332021-3121200123200213-2032102232211013"></a>

## maximum_protocol_version property — cert_params / 222103310111 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-3110123213311003-1231302232002330-1223312220222232-3021100222123110-3033312120300203-1103123002210313-1312012021231033-3333130021101301"></a>

## minimum_protocol_version property — cert_params / 222103310111 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-1212313121200210-3103013202213223-3220222113210022-1011030322230031-3003211103203131-3330302021012213-0121310022122221-3001220120000103"></a>

## Next pages — cert_params / 222103310111 / 7

- [access_info.tls_config.cert_params.certificates](data-sources--secret_management_access--reference--group-001.md#canonical-3223110302010102-0112021220103031-1023330021121210-2202021320210313-0100213121031322-2210312201321303-3230321311203132-1103022001111310)
- [access_info.tls_config.cert_params.skip_server_verification](data-sources--secret_management_access--reference--group-001.md#canonical-2133021130203313-1233113323230223-0313000111132020-2220211301331123-0030220000122311-3113101332103210-2202210131013211-2222322321013300)
- [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0212121101200211-3311122021003102-3323022130010010-1220003310202223-2133133021200212-3203110132212111-0110011212300112-1130202332133031)
- [access_info.tls_config.cert_params.volterra_trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-1010303013132223-3201132001022120-0210133300210210-3123033023101113-1110313213030032-1202123311312003-2321210220210001-3231102311132323)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3223110302010102-0112021220103031-1023330021121210-2202021320210313-0100213121031322-2210312201321303-3230321311203132-1103022001111310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301312112222313-2001023203223132-3003310000202210-0002201323132012-2330232312230003-3322330222211101-3122013213330100-2322201210110210"></a>

## access_info.tls_config.cert_params.certificates — certificates / 012011001302 / 2

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

<a id="canonical-1320211101213203-3321322031221030-3030201001331232-0122231311012310-0032332011133213-0303320130033323-0231101311310010-3102233233102310"></a>

## Direct properties — certificates / 012011001302 / 3

<a id="canonical-0113131011332303-1313131022031221-1201111202212301-1133301123101012-1321302321103321-0030312222012310-2313200301302330-3330312232122010"></a>

<a id="canonical-0020003203222102-0121302002111311-2123121221013203-0113013330313333-0300030110312201-0313232211330010-2312000223230102-2310132030012200"></a>

## kind property — certificates / 012011001302 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3201300113312033-2300032022123202-2101213000022311-0102020112131033-0112001310010203-2123231023111330-2001120321021211-3133113223300102"></a>

## name property — certificates / 012011001302 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1200033113312111-1030011001001232-0022130011111100-1012021301133010-0101200013230330-0102133122331332-3330110003120021-0210100310311203"></a>

## namespace property — certificates / 012011001302 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1222030000312223-0112331323032023-3311103121203230-1302320030223133-1213013123331131-3321002330323323-1001300131310023-2132110102120330"></a>

## tenant property — certificates / 012011001302 / 7

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1300011100003101-1112030103132123-2000333233220210-2031120002202033-3002131130213010-1200101100130221-2012212230203023-0032233233003113"></a>

## uid property — certificates / 012011001302 / 8

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0311203331333032-3112303002202010-1003311131202020-0010303023033210-3311020332123230-1312203212010001-1310330331322233-1012030030302310"></a>

## Next pages — certificates / 012011001302 / 9

- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2133021130203313-1233113323230223-0313000111132020-2220211301331123-0030220000122311-3113101332103210-2202210131013211-2222322321013300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031030212000122-2201131311213111-0010203000231302-1201132301202033-0011100220000323-0013010320001110-1301112303110213-0333210110232030"></a>

## access_info.tls_config.cert_params.skip_server_verification — skip_server_verification / 213121000112 / 2

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

<a id="canonical-1100333200110302-2220300321110212-0221303223202103-3222012131302113-3020332201312122-3110012322133032-1220111010023232-1220321230213003"></a>

## Direct properties — skip_server_verification / 213121000112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021120202210332-0033130031020011-0001310210100110-0212221113321003-1011233121300020-0022111311203110-2020131112212332-3121203332201023"></a>

## Next pages — skip_server_verification / 213121000112 / 4

- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0212121101200211-3311122021003102-3323022130010010-1220003310202223-2133133021200212-3203110132212111-0110011212300112-1130202332133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110303120121220-0233033133303012-3202003313133232-1002331202011202-3123323200101011-0032002131202310-0210131321301011-3103103132210322"></a>

## access_info.tls_config.cert_params.tls_validation_params — tls_validation_params / 110221121312 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- access_info.tls_config.cert_params.tls_validation_params

<a id="canonical-1313132013331301-1022311210223030-0000030232233231-2311012200110231-3112030323013302-2202023231323321-1021232311213211-3031303232220032"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

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

<a id="canonical-1102320333231003-2212333031022232-1323333001331333-2020322112300202-2210310022222303-1221003201221303-2101211211001333-3220121133120012"></a>

## Direct properties — tls_validation_params / 110221121312 / 3

<a id="canonical-1032310210233100-0330132203201222-1211102130203331-3302101230221012-3033202002131331-0233213211010101-2200002221001103-2131312311230032"></a>

<a id="canonical-3132200112003000-0311003020200003-2101212220332032-0110323313331312-2312031020211302-2110300011013310-0001130132012033-0120303330103030"></a>

## skip_hostname_verification property — tls_validation_params / 110221121312 / 4

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

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

<a id="canonical-2010332110013222-3300202300200303-1223102230111210-1323232110013230-3003302033011112-0323232101100200-1110123320010120-2011032321103112"></a>

## trusted_ca_url property — tls_validation_params / 110221121312 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3113332102131130-1320101022120302-3212020312233221-2023333032110301-1122320101300223-2130021200232211-3230103321031211-0211333331231223"></a>

## verify_subject_alt_names property — tls_validation_params / 110221121312 / 6

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

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

<a id="canonical-3322131101011131-0212230012110233-1310212303333313-0220133111012100-2310132222321313-3013222201330012-0010033210231102-1132020102022333"></a>

## Next pages — tls_validation_params / 110221121312 / 7

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-0021003012231102-1113101021031010-1123303330032223-1230302120220031-3233030320222203-0030320213202202-3132332220213000-3023131312113012)
- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0021003012231102-1113101021031010-1123303330032223-1230302120220031-3233030320222203-0030320213202202-3132332220213000-3023131312113012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212202123212022-3022200222113302-1031103101111223-2332103201320101-2133020312021320-2102220300023332-2231002011330103-3322122002123302"></a>

## access_info.tls_config.cert_params.tls_validation_params.trusted_ca — trusted_ca / 021300303023 / 2

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

Upstream description:

Reference to Root CA Certificate.

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

<a id="canonical-1100032201001310-0311201101003202-1213222333021203-2312010100202102-2032032102313213-3231122021013223-3010213110111232-0223111203312230"></a>

## Direct properties — trusted_ca / 021300303023 / 3

- [trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-0110131112111211-1221112123033022-3301310001110213-0311032032222233-0220321220111223-2202200221320332-1013200313202213-3232021212132103): complete subsection reference.

<a id="canonical-0232023311022110-3111213131233330-1002013111033013-3101310330103030-2103030301332130-1021310330022303-2101002302231122-2000310130230112"></a>

## Next pages — trusted_ca / 021300303023 / 4

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-0110131112111211-1221112123033022-3301310001110213-0311032032222233-0220321220111223-2202200221320332-1013200313202213-3232021212132103)
- [access_info.tls_config.cert_params.tls_validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0212121101200211-3311122021003102-3323022130010010-1220003310202223-2133133021200212-3203110132212111-0110011212300112-1130202332133031)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0110131112111211-1221112123033022-3301310001110213-0311032032222233-0220321220111223-2202200221320332-1013200313202213-3232021212132103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032233300213133-3103032312323023-1031333101110321-1233323212000130-0321011212313231-2321032332030112-0033010123211203-3002331212000321"></a>

## access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 331203002010 / 2

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

Upstream description:

Reference to Root CA Certificate.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1210302233100110-1003000200310211-2232213001103322-0010133130223031-2030300022000001-3132331111110002-0111232311121010-2312110123023103"></a>

## Direct properties — trusted_ca_list / 331203002010 / 3

<a id="canonical-0312032301000210-3103022130322030-1131020112113000-2201023023031101-1312220100231132-2001301220220212-0123213233311301-2022021022103200"></a>

<a id="canonical-0320133313213213-0321202113303200-2333231323130210-3030322033021201-3113220331011201-3233212113313203-0233231221233122-2022301310231201"></a>

## kind property — trusted_ca_list / 331203002010 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1232230333303322-3130222130111121-3002003030202221-3302030211210000-1033031333010033-0023100120111212-2102202330001313-3130311001313020"></a>

## name property — trusted_ca_list / 331203002010 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3332201222212102-0132100232021302-1113122322013320-0212212333222013-3101123322330301-1333020003303110-1231211332132033-3021012111012311"></a>

## namespace property — trusted_ca_list / 331203002010 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3200230030333130-2000323021313032-3111131322230023-3103103321201320-3111112123313121-1013313231131133-0331133000230311-1200231312230331"></a>

## tenant property — trusted_ca_list / 331203002010 / 7

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0233100020020203-3012211201331012-2020102321233012-0221112222122113-2100313213020301-0313203223013123-3313233300133333-3330131332222313"></a>

## uid property — trusted_ca_list / 331203002010 / 8

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2022020132330030-2130120311221221-1001212110221123-2312312100110210-3221203100010032-3323311220101321-0220012021313333-1232200231030001"></a>

## Next pages — trusted_ca_list / 331203002010 / 9

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-0021003012231102-1113101021031010-1123303330032223-1230302120220031-3233030320222203-0030320213202202-3132332220213000-3023131312113012)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1010303013132223-3201132001022120-0210133300210210-3123033023101113-1110313213030032-1202123311312003-2321210220210001-3231102311132323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101003313112013-2311300210310230-0331332233121301-0131320120313032-1022320011031203-2302331323211323-0011132300031000-1201103221220332"></a>

## access_info.tls_config.cert_params.volterra_trusted_ca — volterra_trusted_ca / 133331013132 / 2

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

<a id="canonical-1102230112202312-3211302231322221-2120121301212123-1011030203121032-2012323133110312-2232132012001330-2201023131010320-1010022132110023"></a>

## Direct properties — volterra_trusted_ca / 133331013132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133200330121322-1101301231103120-3130102211331311-0013201333220233-3033302212210113-3112122033303002-2302210232120102-2021313003231221"></a>

## Next pages — volterra_trusted_ca / 133331013132 / 4

- [access_info.tls_config.cert_params](data-sources--secret_management_access--reference--group-001.md#canonical-3131122331330221-0201323301032030-0003111313000332-1023303231323133-1202132333213220-2210102231303032-2212213102233201-3302103322130223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002033232310310-0030331312333130-2302130233000003-0130212123112100-1231331123131213-2112310323110211-0312120223321303-1232222320320002"></a>

## access_info.tls_config.common_params — common_params / 113231121011 / 2

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

Upstream description:

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

<a id="canonical-0322303323311201-2232130321112133-1212313020231231-3222313233323211-0000010023000123-2132201013121133-0313201213031133-2123210332012133"></a>

## Direct properties — common_params / 113231121011 / 3

<a id="canonical-3230322012213021-1023003033321213-0210012111120310-1122130203002221-3122000322021111-1131121112333101-3312310032013221-2130302230121112"></a>

<a id="canonical-2131133211023032-0111210030022100-0010133333212110-0201003323123213-1333221032031011-3220230231022302-1011230121322131-3021222230030301"></a>

## cipher_suites property — common_params / 113231121011 / 4

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3023122231033022-2103021110000202-3131210113030010-3231322310113233-1322300330330232-0033123121020220-3131031322033120-1000222000121021"></a>

## maximum_protocol_version property — common_params / 113231121011 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-1001103021332113-3222121202323000-3330023221321002-1110102103011111-2312131331132221-3032200210221123-3122222103333030-2020210310012320"></a>

## minimum_protocol_version property — common_params / 113231121011 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-0012302023220032-0103303021210320-0222111223032002-2302120023232123-1230331302321023-1021010322102302-0332021001323221-3231320131000233"></a>

## Next pages — common_params / 113231121011 / 7

- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0210313100110112-3122102002210000-1211212311023133-3110030202023333-3302121303223220-0320032333320132-0313201002332222-2100311203122102)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020213100223313-1322213331131320-1111133322133213-2222301101300011-0210232200122301-3210030332132202-3321101301023033-0221230113330233"></a>

## access_info.tls_config.common_params.tls_certificates — tls_certificates / 031331002221 / 2

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

Upstream description:

Set of TLS certificates.

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

<a id="canonical-1323002123203001-1301130300002211-2200223230310112-1013032001313322-3030001230233100-0332213030233220-0330021332113223-2222101131213113"></a>

## Direct properties — tls_certificates / 031331002221 / 3

<a id="canonical-2013130102301020-0213010113000321-2230300211012101-2110210223313303-2023000313322312-0311331133320330-0023230222323330-1020220021110121"></a>

<a id="canonical-0302110333111031-3000210222103303-1313323301032132-1031102010112322-2332301100320233-0011331113002130-3233112201332011-1023303132320002"></a>

## certificate_url property — tls_certificates / 031331002221 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3122120101013232-2023020301203013-3331321103133033-1321101030033130-0202112111001013-0201121021222103-3003300112302320-2102231310202002"></a>

## description_spec property — tls_certificates / 031331002221 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--secret_management_access--reference--group-001.md#canonical-0003022330330313-2030232201210223-0202003033120100-2230302110300101-0331323332021332-3110011033310020-3231011111312021-0122331102022131): complete subsection reference.

- [private_key](data-sources--secret_management_access--reference--group-001.md#canonical-1301230322222320-1302321123212130-1023332303302213-3033032012112002-2202110332212332-0011110311012320-2332331130220033-2002133021021122): complete subsection reference.

- [use_system_defaults](data-sources--secret_management_access--reference--group-001.md#canonical-3130030022223102-0323033203032120-3020100320133323-0121230311203203-1333310130210032-3233123112131222-2210223233132233-2020231011123222): complete subsection reference.

<a id="canonical-2320231213102123-1322230120001122-2211003030203021-1330231231211332-3303223223233001-2221320033123003-2103101300003011-1012322323022101"></a>

## Next pages — tls_certificates / 031331002221 / 6

- [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](data-sources--secret_management_access--reference--group-001.md#canonical-0110032232203310-1222333123120133-3013132321021220-0200302121131311-1303213330133100-0011300221203021-0032303132320222-1311033000012033)
- [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](data-sources--secret_management_access--reference--group-001.md#canonical-0003022330330313-2030232201210223-0202003033120100-2230302110300101-0331323332021332-3110011033310020-3231011111312021-0122331102022131)
- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-1301230322222320-1302321123212130-1023332303302213-3033032012112002-2202110332212332-0011110311012320-2332331130220033-2002133021021122)
- [access_info.tls_config.common_params.tls_certificates.use_system_defaults](data-sources--secret_management_access--reference--group-001.md#canonical-3130030022223102-0323033203032120-3020100320133323-0121230311203203-1333310130210032-3233123112131222-2210223233132233-2020231011123222)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0110032232203310-1222333123120133-3013132321021220-0200302121131311-1303213330133100-0011300221203021-0032303132320222-1311033000012033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323312000220300-0323020210332011-1313133030302102-3102011023021320-1130121032303301-3201211123322202-2223133110122001-1211203021222031"></a>

## access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 103023200303 / 2

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

<a id="canonical-2331211310212333-0001222313103231-2130221132130133-1002111330011210-1102310111221113-0201313023200302-1310311211231211-2000232032310310"></a>

## Direct properties — custom_hash_algorithms / 103023200303 / 3

<a id="canonical-2300022023333230-0230100001123230-1230013320210000-3002012002020020-1332231300100203-1100131011113001-3033021133132310-1023003132101212"></a>

<a id="canonical-0323230303131300-2132021002300331-3103133300303203-3013123230031210-0201231212321313-1221301021233000-0022110323010001-2102212200032102"></a>

## hash_algorithms property — custom_hash_algorithms / 103023200303 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2111301220112203-0021332131132111-3100122311233002-1321300023032123-3210232113020001-2202130310023232-1301212131323330-1232032202020322"></a>

## Next pages — custom_hash_algorithms / 103023200303 / 5

- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0003022330330313-2030232201210223-0202003033120100-2230302110300101-0331323332021332-3110011033310020-3231011111312021-0122331102022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003321211323002-0032131000230210-1330231302301323-1212020213202321-3313000210203203-0222222311320012-1032132310012320-1330203300300100"></a>

## access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 202001003001 / 2

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

<a id="canonical-1031011002310303-0313323101123100-1032203322023013-3222330311331001-0021011331023211-0110213100012101-0022333200130202-3023232132232133"></a>

## Direct properties — disable_ocsp_stapling / 202001003001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002031111320232-0130123212222131-0321222003211202-3010101131222033-2123203211103322-1113021122010122-3110221031020332-2212100210110121"></a>

## Next pages — disable_ocsp_stapling / 202001003001 / 4

- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1301230322222320-1302321123212130-1023332303302213-3033032012112002-2202110332212332-0011110311012320-2332331130220033-2002133021021122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312011330322123-0030310220012101-1312213213113132-0203223220230321-0013023003203112-1231222003031001-2133111200301102-1201123320000220"></a>

## access_info.tls_config.common_params.tls_certificates.private_key — private_key / 131011200301 / 2

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

<a id="canonical-3010210123110233-0131103022203102-3301312300322121-2330323312102120-0002103130320301-1322130112212002-3030200113121033-0131113322213220"></a>

## Direct properties — private_key / 131011200301 / 3

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-0320200332323203-2122000132300020-3111121303202213-0123001310310221-1212133202322322-2032033221312033-3233021130220010-0121103233220222): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-0333021122331232-3201202312000300-1101032010131000-0131020333212013-2210100202010223-3013112100112333-3311320212230322-1312102133302320): complete subsection reference.

<a id="canonical-2113331102202030-0332210213212203-3302020001201222-1321302023031222-1010012133212203-2322222300223213-0212112121310001-2003010323312031"></a>

## Next pages — private_key / 131011200301 / 4

- [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-0320200332323203-2122000132300020-3111121303202213-0123001310310221-1212133202322322-2032033221312033-3233021130220010-0121103233220222)
- [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](data-sources--secret_management_access--reference--group-001.md#canonical-0333021122331232-3201202312000300-1101032010131000-0131020333212013-2210100202010223-3013112100112333-3311320212230322-1312102133302320)
- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0320200332323203-2122000132300020-3111121303202213-0123001310310221-1212133202322322-2032033221312033-3233021130220010-0121103233220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202031130002330-1311022313303330-2210020101322011-2012232100332330-0320100101011020-2201031133202212-2302130333210021-3110202330130030"></a>

## access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 032322221033 / 2

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

<a id="canonical-1130021202333002-1303301332212002-3102101133022203-1102101133231331-2303110302100232-3303021210031010-3133310303213213-0012101011220220"></a>

## Direct properties — blindfold_secret_info / 032322221033 / 3

<a id="canonical-1113110033102300-1331103000031232-0103332101200132-2021131213222110-2303321311002320-0200300312023001-0100103323332212-2312211312032103"></a>

<a id="canonical-2232133031202300-2100131213220202-1231333111023332-3131000022200110-3112333301323020-0131202310310202-0131221031330233-3123132120031311"></a>

## decryption_provider property — blindfold_secret_info / 032322221033 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3132320112030022-1332232300111202-0133033330232020-2100132000012201-2020103033332312-3213033320103011-0310103001120003-3031320331320130"></a>

## location property — blindfold_secret_info / 032322221033 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3220123130211233-2303132213322210-3121133121200111-2130321220103030-2000032300011022-2233003223102313-1123201332202123-2312222121201311"></a>

## store_provider property — blindfold_secret_info / 032322221033 / 6

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3300330112030001-3321321301123133-1012312232001021-1303333311123130-2312203312223133-0330013303231131-3110120033220001-3123311203021031"></a>

## Next pages — blindfold_secret_info / 032322221033 / 7

- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-1301230322222320-1302321123212130-1023332303302213-3033032012112002-2202110332212332-0011110311012320-2332331130220033-2002133021021122)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0333021122331232-3201202312000300-1101032010131000-0131020333212013-2210100202010223-3013112100112333-3311320212230322-1312102133302320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213210122002103-2311003331231023-0102323331132103-3102211221101213-1032203321200310-0303231213010333-1212113322202123-0200011331233023"></a>

## access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info — clear_secret_info / 332033200011 / 2

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

<a id="canonical-1100110202113001-1321011122313313-1232110321022302-1201210111033210-0000212223002213-1222220113313111-2313210321110222-1021202131032102"></a>

## Direct properties — clear_secret_info / 332033200011 / 3

<a id="canonical-0120320202103310-0323220321223103-3201213102300112-3232123322132223-2011311220211113-0101000303130102-3032332110022232-0302211120323122"></a>

<a id="canonical-1013022031310202-2311110123300303-1103031022313001-2310233201301313-2323001001013121-2120312000210013-3212020202211002-0030300202231032"></a>

## provider_ref property — clear_secret_info / 332033200011 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3321212121332100-2030312333120000-1202220312223003-1103203331001310-2233302333221111-2120033211230212-3321032233012231-2012003121120111"></a>

<a id="canonical-0322221003033000-3032102233332312-2211100102302303-1102003200201201-0332210230001032-0113112000132122-1310133321002022-3200322220210221"></a>

## URL property — clear_secret_info / 332033200011 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2202013013001111-3112330032323321-3220122213311103-1310102131021010-2130130323303320-3211013020022110-3121311131000022-2120312123033131"></a>

## Next pages — clear_secret_info / 332033200011 / 6

- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--reference--group-001.md#canonical-1301230322222320-1302321123212130-1023332303302213-3033032012112002-2202110332212332-0011110311012320-2332331130220033-2002133021021122)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3130030022223102-0323033203032120-3020100320133323-0121230311203203-1333310130210032-3233123112131222-2210223233132233-2020231011123222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312201020320221-1211332113100301-1130122030320332-2322201023300102-1022022201011002-2030213200213213-3211020331103023-0032200312032113"></a>

## access_info.tls_config.common_params.tls_certificates.use_system_defaults — use_system_defaults / 013313001310 / 2

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

<a id="canonical-2302112123232133-3112232031232002-2231233313221232-1113121030111331-1322332001022232-1133023100312331-3313230033210301-2033210033133132"></a>

## Direct properties — use_system_defaults / 013313001310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211232330310311-0312332203010020-0233211203121132-1030111023030002-3111001300100120-1300031310212100-2320120300311203-3220110323231321"></a>

## Next pages — use_system_defaults / 013313001310 / 4

- [access_info.tls_config.common_params.tls_certificates](data-sources--secret_management_access--reference--group-001.md#canonical-2033001231003230-0030211003323131-3230010320121333-3133230111211123-3102103313211111-1221132313301310-0023222023313312-0303223210322312)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0210313100110112-3122102002210000-1211212311023133-3110030202023333-3302121303223220-0320032333320132-0313201002332222-2100311203122102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321131010021020-3033202033102033-3212132102310333-0220120103313302-2110032303000021-1321231000203133-3233303101323022-2211220333202113"></a>

## access_info.tls_config.common_params.validation_params — validation_params / 030203303332 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- access_info.tls_config.common_params.validation_params

<a id="canonical-2310202013033320-0302233011130222-0311210021331122-1031030033221210-1102101101331122-1132133133020033-2133011021210100-1130132330231221"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

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

<a id="canonical-0203321003211313-3033213230313010-2230323123002021-3202221233123332-0321212013102011-1211331022032010-2330020002121223-1021230313032211"></a>

## Direct properties — validation_params / 030203303332 / 3

<a id="canonical-0022113231002332-1332022301001112-2033100011111012-2123301312132021-1013302032023130-1320101133032103-2313100130222202-3203100232111303"></a>

<a id="canonical-1102110031201001-1031113313313313-3100001233011210-1221331011120020-2010331121311003-3011231011232000-3221201312121103-2003301000102230"></a>

## skip_hostname_verification property — validation_params / 030203303332 / 4

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

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

<a id="canonical-3130201022000320-2333300003130013-1032111032312322-1121213232013100-3221302130333303-1002130031133120-0303232130330022-0323333201322222"></a>

## trusted_ca_url property — validation_params / 030203303332 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3023021010301200-0112233122211130-1100321030201011-2102020222022013-0322322203320312-1232312003132011-3303313311301020-2320201210232132"></a>

## verify_subject_alt_names property — validation_params / 030203303332 / 6

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

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

<a id="canonical-1230030212321330-0311113212212331-1110203131130130-3221311301310123-3310101020010031-0302032320021010-2031131101210233-3123231312021001"></a>

## Next pages — validation_params / 030203303332 / 7

- [access_info.tls_config.common_params.validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-3323022232231121-0202001001312303-2231100110230321-2130122233101103-3231100313122112-1003222223102130-1121012032301232-0213203023200030)
- [access_info.tls_config.common_params](data-sources--secret_management_access--reference--group-001.md#canonical-0000013110300012-1102302103331223-2031123231303130-1202332002112232-0310323100020032-2223022103203121-0011310330031231-1102120023032102)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3323022232231121-0202001001312303-2231100110230321-2130122233101103-3231100313122112-1003222223102130-1121012032301232-0213203023200030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000033121231032-3112123010300002-1310300222020321-2130210110000302-3110133302331001-3010002211020312-2030311221133113-2103302333103101"></a>

## access_info.tls_config.common_params.validation_params.trusted_ca — trusted_ca / 033030301333 / 2

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

Upstream description:

Reference to Root CA Certificate.

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

<a id="canonical-0303132032231223-2103300003210310-1310012330203113-0323113332233233-1002133233233301-1221011220321111-2013321023232301-3321023112021302"></a>

## Direct properties — trusted_ca / 033030301333 / 3

- [trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-3213223132332031-0010001010211223-2310021322133310-0221230213332012-3023122303213312-2101331020231003-1311031113300131-2323011120232000): complete subsection reference.

<a id="canonical-2103330301311122-1203302032211233-3232030001120022-2022132330221212-1022222220311331-2032031121122020-3313232121221131-2321211231020202"></a>

## Next pages — trusted_ca / 033030301333 / 4

- [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--secret_management_access--reference--group-001.md#canonical-3213223132332031-0010001010211223-2310021322133310-0221230213332012-3023122303213312-2101331020231003-1311031113300131-2323011120232000)
- [access_info.tls_config.common_params.validation_params](data-sources--secret_management_access--reference--group-001.md#canonical-0210313100110112-3122102002210000-1211212311023133-3110030202023333-3302121303223220-0320032333320132-0313201002332222-2100311203122102)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3213223132332031-0010001010211223-2310021322133310-0221230213332012-3023122303213312-2101331020231003-1311031113300131-2323011120232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032312020022022-3301023002010110-1133033230310032-3130202020303011-1201123120101232-0323121000200322-2031332112131331-0132210320333220"></a>

## access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 132011111011 / 2

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

Upstream description:

Reference to Root CA Certificate.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2303121210222331-1200000223112331-2330301211213111-0221313311003120-2002221321131000-1122303111332113-2231000323002103-2001021111202310"></a>

## Direct properties — trusted_ca_list / 132011111011 / 3

<a id="canonical-1333313301023221-1223133121110202-0031220012233033-0333213030300200-0032002132113212-0100311003130312-3000122030212002-3222313231320210"></a>

<a id="canonical-1023202023022113-0300222133321231-0300331113222231-2212002010232032-1203202321021233-2022332021333220-3303231303222330-2331331132201233"></a>

## kind property — trusted_ca_list / 132011111011 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0020102302111222-0302323103132121-3031300033200321-3100311132202333-0220132000211002-1232003220223123-1221112133010322-3113101031131233"></a>

## name property — trusted_ca_list / 132011111011 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2102122302130332-2231111332111121-3132202022100132-2231313120030021-2322023321022101-0020112002003002-2201130132030100-2210201301322223"></a>

## namespace property — trusted_ca_list / 132011111011 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2011002331321231-0030130331303123-1230023331121300-1330313113303310-2323222033123302-1300121110010333-1022010130311021-1301121222220311"></a>

## tenant property — trusted_ca_list / 132011111011 / 7

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3201111231012012-3032323033211323-2003200230130313-1000223033211101-1303033301003013-1121201031012021-2101131011220020-0102323021132131"></a>

## uid property — trusted_ca_list / 132011111011 / 8

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2313220130032003-2232223100302332-3213322320210333-1321101022020200-0122231032031213-1030222302111010-2003202332012022-3302303231021101"></a>

## Next pages — trusted_ca_list / 132011111011 / 9

- [access_info.tls_config.common_params.validation_params.trusted_ca](data-sources--secret_management_access--reference--group-001.md#canonical-3323022232231121-0202001001312303-2231100110230321-2130122233101103-3231100313122112-1003222223102130-1121012032301232-0213203023200030)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1210103112123023-1311003031333303-0033320301002230-1333002123130023-2302111201030130-2311312223001120-3320000021000013-2132010033213110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212030130120332-3002212323123302-1300231123031133-2312000201323102-2101003003130332-3232123323003231-3033103221211331-3212023031330220"></a>

## access_info.tls_config.default_session_key_caching — default_session_key_caching / 211122101202 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.default_session_key_caching

<a id="canonical-2131321032020123-0232231210002310-2031113022030210-3121310200233322-3333210211133131-0000310331313310-2330301011100132-3323211313212231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching.

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

<a id="canonical-1002213032103121-2202232202222023-0211023222201031-0302222100331030-1203130200202213-2100000223033213-3223100203222011-3110023203221321"></a>

## Direct properties — default_session_key_caching / 211122101202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310221332201102-0002200230033301-3012223103332231-3331122121022302-2222112032301111-2201333320203122-3210311303301220-0221122211121201"></a>

## Next pages — default_session_key_caching / 211122101202 / 4

- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0011323301322002-2201102122202131-2130020111322130-1323121110112002-3121221011021313-3000012122012233-3313000221030220-2013020223002321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332031013023131-0122002330033020-0231300310130012-1203030322301112-0023202122121122-2200012320332013-2321311130201013-2220202201003303"></a>

## access_info.tls_config.disable_session_key_caching — disable_session_key_caching / 031333310122 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.disable_session_key_caching

<a id="canonical-2032010020020012-1233223300322003-1213013022322221-2202210101010333-3111213332330112-0220031223131002-2112031122020333-1212301310002223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

<a id="canonical-3331003031033202-2032222203213102-3130212011221202-0323132032130033-3330331200313232-2130222132212313-1301123232311312-1101223033202123"></a>

## Direct properties — disable_session_key_caching / 031333310122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001232210323312-1013121133103002-2021020322213110-2131102203220320-2031313321003211-2302313120132312-3030123000031122-1220330010210112"></a>

## Next pages — disable_session_key_caching / 031333310122 / 4

- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3313002020213130-1211220100111301-2010321010111112-1221223002230323-2012201001201323-0033002301121033-0211321133002012-3230002322001001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323103022110100-3003201200233111-1210013122332333-2223331233321303-2332121030111322-1000020000010303-1001000032033020-2033030332012013"></a>

## access_info.tls_config.disable_sni — disable_sni / 200310023123 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.disable_sni

<a id="canonical-0000330111213131-3031223010221210-2130332300203033-1201323121001001-0120113220113311-2320210310313030-3310311321132301-0131032013233331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

<a id="canonical-0113132113020110-2232321033021330-2002311012032210-2201011030311332-1221202222103000-1122033100120200-1010231220113100-3011320210332122"></a>

## Direct properties — disable_sni / 200310023123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132302321303013-0111133213122221-0301011213321301-2230003033231220-1100121123123000-2313120023313103-1201102002003230-2210001230221300"></a>

## Next pages — disable_sni / 200310023123 / 4

- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1323322302212031-2102331223003212-2232132102320011-3313132032020021-3013331110313121-0310233210023203-0320100300121321-1033322112222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003313130332231-2222132330300311-3332102020212222-2002333200032133-3121221031331221-1213311010020201-3101130113113333-3301110201203002"></a>

## access_info.tls_config.use_host_header_as_sni — use_host_header_as_sni / 032131203330 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- access_info.tls_config.use_host_header_as_sni

<a id="canonical-0213020113222102-3012113030302333-2013103220233122-0031232211300312-0112113112010332-0330333203322113-0002300002102002-1103033312322010"></a>

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

<a id="canonical-0000301222103103-2331221312100201-1022130133102133-2211011233110132-1010222231322232-0021033302021331-3030330320311100-0210103101330201"></a>

## Direct properties — use_host_header_as_sni / 032131203330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213222131021103-1200200300332203-3131203333330102-1312221331321003-3103310221100133-2322111302332312-2222230121010320-2030133210112201"></a>

## Next pages — use_host_header_as_sni / 032131203330 / 4

- [access_info.tls_config](data-sources--secret_management_access--reference--group-001.md#canonical-3112300310130311-1322310032001022-1121202023203333-0013122330212132-1302320121130013-1001122210211023-2203321122112132-1130133120030223)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203233223020320-3120110202010201-0010000132333111-3220333200211100-3031031300233311-2110222212022103-0223111022113021-1203003031221021"></a>

## access_info.vault_auth_info — vault_auth_info / 032232323310 / 2

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

<a id="canonical-1133221031032022-0210202022333033-0022001231012223-3101223322020321-2113201012300021-0321133112302223-2110130022100102-1223032001321231"></a>

## Direct properties — vault_auth_info / 032232323310 / 3

- [app_role_auth](data-sources--secret_management_access--reference--group-002.md#canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022): complete subsection reference.

- [token](data-sources--secret_management_access--reference--group-002.md#canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203): complete subsection reference.
