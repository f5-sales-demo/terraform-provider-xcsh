---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220210102311023-1013020223213130-3122131320221112-2132100213002111-3030220201330310-0302022220301330-2233102202231001-0122003212122113"></a>

## Property reference — Property reference / 031200003023 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- Property reference

<a id="canonical-0113121010313001-1010222021031220-1221022301103013-0113012202300001-3203100010333333-0032303013222202-3323122130002200-3222233122322023"></a>

## Direct properties — Property reference / 031200003023 / 3

- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100): complete subsection reference.

<a id="canonical-0223313201221312-1013011120202120-3110123320033331-2230130221120023-1330232301213030-1123311022023200-0312311130212312-0321132103211323"></a>

<a id="canonical-1030321020003003-2300232310003121-2130120330121311-0230221103023222-1122202220331320-2311220213032301-0113132333313133-1131003012302112"></a>

## annotations property — Property reference / 031200003023 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-2331100222000021-0111310120210223-3030031213202202-2333100211101220-3111001130310222-3301003032303331-3232313210333130-3100223102031020"></a>

<a id="canonical-1332112302303033-0003220231000030-3223100112230231-3132012230130212-0331330322010202-0013030323233013-0331323322321011-3332111323300333"></a>

## description property — Property reference / 031200003023 / 5

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

<a id="canonical-1323230111130030-3211120002023133-3132101123020030-1200303132133212-2203032122202200-2312002221300030-2231203322103033-2203031013102113"></a>

<a id="canonical-0100212003201102-2333013231220221-2312133112201003-2320233031200200-3301212021110321-3010223023011310-2021032321032211-2103202032221312"></a>

## disable property — Property reference / 031200003023 / 6

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

<a id="canonical-1031201332123121-3211102231022333-0003001202000333-0212310310030010-0100310310323122-2222222112202232-0112011021132121-2330332212021223"></a>

<a id="canonical-1102320200203210-3322121331210113-0220133312202203-3113333122211001-0312211212230112-0220101333321231-1023222133022111-2112001111102101"></a>

## ID property — Property reference / 031200003023 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0220012302320322-0302113133120331-3310230003222102-0230022111121223-0113031031330021-0301331031131332-1100330113221200-2120333200020331"></a>

<a id="canonical-0022120211121100-0032130003332331-1300122333022130-2102120012231222-3202312130212103-0331223221011321-0313123003031222-3003011101100223"></a>

## labels property — Property reference / 031200003023 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
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

<a id="canonical-2222323103103221-0200122230200112-2003303013013012-0011330201103130-1013032103311312-3211022131133021-2302003321203302-2323012311201010"></a>

<a id="canonical-0211130230223020-0333203331200220-3023231332332121-0200112320120110-2131032100302133-0203102332112323-0023012030331320-3313213122133321"></a>

## name property — Property reference / 031200003023 / 9

Type: `"string"`. Required.

Name of the Secret Management Access. Must be unique within the namespace.

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

<a id="canonical-3133031310002231-3111313003203222-3201213203033020-2220001112101120-2010312233332002-1013301332112310-0132102221302133-0112312010000020"></a>

<a id="canonical-3111121332001022-3330132313000110-0012031300033333-1322223321030232-0100110120211323-2312103210011113-0302210230321323-0102323312133003"></a>

## namespace property — Property reference / 031200003023 / 10

Type: `"string"`. Required.

Namespace where the Secret Management Access is created.

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

<a id="canonical-1011003103120023-0202232230032130-2323020033032330-0032222230220011-0003200131331201-0213230300330132-0130102233131032-1031111332130303"></a>

<a id="canonical-3111010233332311-3230013003230332-3032131131230022-1023010020110303-0322113320230312-2020112112322132-0112212211022312-2321013312100323"></a>

## provider_name property — Property reference / 031200003023 / 11

Type: `"string"`. Required.

Name given to this secret management backend. site.provider needs to be unique, and will be
referenced for using this object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [timeouts](resources--secret_management_access--reference--group-002.md#canonical-3323013113011220-0112002023230221-0022123301231112-0311033122021321-0012233310113333-1120230302310030-3302113003112103-0010202130123300): complete subsection reference.

- [where](resources--secret_management_access--reference--group-002.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211): complete subsection reference.

<a id="canonical-0321310101032333-1131231230321013-1311213210133311-2302313113011022-0201201101010131-0023102323201132-2133202103232032-1333322101201202"></a>

## All schema paths — Property reference / 031200003023 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `access_info` | [access_info](resources--secret_management_access--reference--group-001.md#canonical-1222032012011203-3310232220222312-1031202321200133-0022312310213330-0012300111113302-0120301322020011-1110013030330123-0013313102002203) |
| `access_info.rest_auth_info` | [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-1031033332300221-3011202212322003-1132000321310013-2020130113100003-3213331100301023-1102113301213002-2030123210323312-2111013321013330) |
| `access_info.rest_auth_info.basic_auth` | [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-2312112200321323-1122001233003312-2201230022122021-0332132020212221-2201102100030123-3102221033311302-1030311120003011-2002033003032222) |
| `access_info.rest_auth_info.basic_auth.password` | [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-0002311202202012-3130010332330123-1313023111233013-3233111031133232-2012332233101133-3130312001201111-0300203111202301-2220122232333021) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-2033200233022313-2302112120213212-3302333011220221-0222012023002203-2301323311103320-0222230310122023-0101330222011022-1211132333301011) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.decryption_provider](resources--secret_management_access--reference--group-001.md#canonical-1333020331201002-2330300021032321-2021030023320010-3202201201131011-3000131102202022-3333303300320020-0000220332023021-2102312330003133) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.location](resources--secret_management_access--reference--group-001.md#canonical-3231311122012231-0211231222131003-1311012023323333-3311213333002222-1030203022301101-2320030110323322-1311302321012213-1002031023023122) |
| `access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider` | [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info.store_provider](resources--secret_management_access--reference--group-001.md#canonical-0002131330231002-1022323030011200-1311121003002321-1213321133213212-3200033022000010-0030100101222033-1013031023211121-3200230303220211) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-1103300310200022-3101020110020223-1120220122302230-0100132021223003-0221101030213301-1021213303202011-3322233221001121-0303212003213013) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.provider_ref](resources--secret_management_access--reference--group-001.md#canonical-1310001110201212-3331132322211011-3111131211312133-3001312312011313-0201103001101233-2123320130230213-3310023021231011-0210321010310023) |
| `access_info.rest_auth_info.basic_auth.password.clear_secret_info.url` | [access_info.rest_auth_info.basic_auth.password.clear_secret_info.url](resources--secret_management_access--reference--group-001.md#canonical-0101013221311001-1220212010122303-3212212333303331-0323300011301132-3320021202002231-2233121010102122-0330210110023101-0201002030231102) |
| `access_info.rest_auth_info.basic_auth.username` | [access_info.rest_auth_info.basic_auth.username](resources--secret_management_access--reference--group-001.md#canonical-2233333030313032-0121032210000010-3012010133010211-0223312113032231-3331131231202030-2101303212320221-3231231210333010-1213310131332003) |
| `access_info.rest_auth_info.headers_auth` | [access_info.rest_auth_info.headers_auth](resources--secret_management_access--reference--group-001.md#canonical-2220120110122103-3200302102111032-0112003222322013-0210101221001332-2210220131203331-2310221331223230-2300303020332311-3320121113021013) |
| `access_info.rest_auth_info.headers_auth.headers` | [access_info.rest_auth_info.headers_auth.headers](resources--secret_management_access--reference--group-001.md#canonical-1000022023333130-1313222011110232-2012311301232223-3301113030302303-1221133112111223-0011101030011331-0303213212200221-0222030131022320) |
| `access_info.rest_auth_info.query_params_auth` | [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-3222020203131020-2220120113220003-1111010301112010-2233213302112202-2013332030113221-0131300221323100-2233102303233231-2030201332120102) |
| `access_info.rest_auth_info.query_params_auth.query_params` | [access_info.rest_auth_info.query_params_auth.query_params](resources--secret_management_access--reference--group-001.md#canonical-0330312022210021-0031110333310032-0033321211330120-2330220222133313-2213132111200310-0221231302301233-2321020301013301-3230113322123302) |
| `access_info.scheme` | [access_info.scheme](resources--secret_management_access--reference--group-001.md#canonical-0112112121002033-0021232203211000-1103303203313211-1211003012013213-2200113321213230-0331030332313310-3323312001101023-3003031203033102) |
| `access_info.server_endpoint` | [access_info.server_endpoint](resources--secret_management_access--reference--group-001.md#canonical-1330310033321121-0133103101003021-1231202220233301-0313100232020220-3023023230321221-2300122113211121-2001131030011123-0001113213122212) |
| `access_info.tls_config` | [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-1002202302301223-0220232301022112-3031231201312303-2200101231130303-3300121202320210-3012312112002220-2111002002302311-2011331120032110) |
| `access_info.tls_config.cert_params` | [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-0313130300021200-1332201110230311-3302213323223330-2030021203103023-3002223212022302-2021220321320033-3013111210213010-0322130023320303) |
| `access_info.tls_config.cert_params.certificates` | [access_info.tls_config.cert_params.certificates](resources--secret_management_access--reference--group-001.md#canonical-3132302102333311-1131103200100330-2320102112323002-2322301231322032-0223010221112333-0133002132221031-3111220303323202-2233333231130101) |
| `access_info.tls_config.cert_params.certificates.kind` | [access_info.tls_config.cert_params.certificates.kind](resources--secret_management_access--reference--group-001.md#canonical-0110332320230120-2010112300312211-1032023331310213-0112312132200321-1003130130022221-0111013032010101-3110000032100322-0321301023111313) |
| `access_info.tls_config.cert_params.certificates.name` | [access_info.tls_config.cert_params.certificates.name](resources--secret_management_access--reference--group-001.md#canonical-1121102321112131-2031223321320130-3211121033212232-3322120010310010-3333010222011302-2111322220102101-0231211230211303-1233321322321031) |
| `access_info.tls_config.cert_params.certificates.namespace` | [access_info.tls_config.cert_params.certificates.namespace](resources--secret_management_access--reference--group-001.md#canonical-0112313200133110-2002331002030222-3233330112031322-3030322103302132-3012020323223302-3130210011303110-0221313321100200-0032232030300011) |
| `access_info.tls_config.cert_params.certificates.tenant` | [access_info.tls_config.cert_params.certificates.tenant](resources--secret_management_access--reference--group-001.md#canonical-2313021232133013-1320103023123122-1222012030120032-1100210033000111-1101223133021231-1210113311302210-1020021112332021-3103132312333210) |
| `access_info.tls_config.cert_params.certificates.uid` | [access_info.tls_config.cert_params.certificates.uid](resources--secret_management_access--reference--group-001.md#canonical-0031222021001232-3220312300133022-0013000312212021-2131002110131232-3003121212012032-2020331310233023-1111112123231112-3111102002013312) |
| `access_info.tls_config.cert_params.cipher_suites` | [access_info.tls_config.cert_params.cipher_suites](resources--secret_management_access--reference--group-001.md#canonical-3103022213331301-1233101222200022-1330003330200330-3032101130023002-2323201233002221-1120230222113332-2100100101010333-2202311102002030) |
| `access_info.tls_config.cert_params.maximum_protocol_version` | [access_info.tls_config.cert_params.maximum_protocol_version](resources--secret_management_access--reference--group-001.md#canonical-1211221210221213-3131230201122321-0230221033231020-3321300320200210-1133010110031030-2113323003030302-1323031103330000-2010002031331133) |
| `access_info.tls_config.cert_params.minimum_protocol_version` | [access_info.tls_config.cert_params.minimum_protocol_version](resources--secret_management_access--reference--group-001.md#canonical-1100312201321103-3212322330012303-3222313120010231-0332221100330223-3310031300223211-3202102301320202-2333213313032213-1120000300311031) |
| `access_info.tls_config.cert_params.skip_server_verification` | [access_info.tls_config.cert_params.skip_server_verification](resources--secret_management_access--reference--group-001.md#canonical-0222232033020302-0021300031013113-3202313210310213-3310120210333010-0311032013233233-3011100231013002-3331010120033123-3000230003231103) |
| `access_info.tls_config.cert_params.tls_validation_params` | [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-1100220133013321-0010330320201003-3332010020121220-2201133132023212-3300300321213212-1012323330123112-3211322300212103-3103333112022213) |
| `access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification` | [access_info.tls_config.cert_params.tls_validation_params.skip_hostname_verification](resources--secret_management_access--reference--group-001.md#canonical-2013320200302230-0323322100100031-3211230032101113-2112313031300003-3211021012211202-1001032020201103-2311222311121222-1231123000133111) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-1010303223020121-2012020203121300-2203103023020202-2321331103231023-1313320332211030-1112210133231323-0020320322200103-2111322021131000) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-1211003001223000-2332222332101101-2312000331200022-0300132110013233-3032313113113321-1203022122131220-3030321300322131-0231220201130322) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](resources--secret_management_access--reference--group-001.md#canonical-0311330001013003-3312201121223002-0031303211020322-2320223310020022-1013311323320200-0030132020223000-3030121000330233-3332012330032130) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](resources--secret_management_access--reference--group-001.md#canonical-2103233113010130-1200000132002220-2203111101302113-2312110020211032-1323012221030130-0002130123021311-0301302321301232-3212010113131301) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](resources--secret_management_access--reference--group-001.md#canonical-3110333302311213-0123032131330112-3232212011111012-0120212320202033-0002013133031132-0303312203013133-1330211312001300-3011011013122320) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](resources--secret_management_access--reference--group-001.md#canonical-2231122322333222-0223301000201013-3103012021011013-0333222022100123-3003133123222232-0010312302220210-1220303031130220-1013021210231321) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](resources--secret_management_access--reference--group-001.md#canonical-1202002103200000-0103001130130322-2323333033302132-1120303210322303-2113010212233331-0133111300322100-1033323131032211-3013030033111322) |
| `access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url` | [access_info.tls_config.cert_params.tls_validation_params.trusted_ca_url](resources--secret_management_access--reference--group-001.md#canonical-1201210002301130-1030010022212013-1012013323013021-2332333302102112-1221332130303331-1313123211010123-1301132202131011-0312321033110300) |
| `access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names` | [access_info.tls_config.cert_params.tls_validation_params.verify_subject_alt_names](resources--secret_management_access--reference--group-001.md#canonical-3301221121101013-2311333121322323-1100132013012100-3331132011103232-2100230002320111-2232132220113210-1313132303133302-3211302020221321) |
| `access_info.tls_config.cert_params.volterra_trusted_ca` | [access_info.tls_config.cert_params.volterra_trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-0011223123132122-3120310233031211-0312111210103103-1023001011020323-0132211110311302-0312000122330332-2323232231000002-1310000030223123) |
| `access_info.tls_config.common_params` | [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1021313131330113-2131223011003120-3311122112133333-1230022211301213-2211233110202331-0011202223120231-1330331000013300-0311020221031031) |
| `access_info.tls_config.common_params.cipher_suites` | [access_info.tls_config.common_params.cipher_suites](resources--secret_management_access--reference--group-001.md#canonical-0022003202011223-1310320331322202-2231303220310003-2131230303022320-2322310320122231-3102321131223002-2230101030020312-2331313000020101) |
| `access_info.tls_config.common_params.maximum_protocol_version` | [access_info.tls_config.common_params.maximum_protocol_version](resources--secret_management_access--reference--group-001.md#canonical-1110111000102103-3301322023023003-2231303021122101-3320321122232332-2203311203110221-1112311332032310-3222321230203320-3131103021323030) |
| `access_info.tls_config.common_params.minimum_protocol_version` | [access_info.tls_config.common_params.minimum_protocol_version](resources--secret_management_access--reference--group-001.md#canonical-0220110013103230-3122132112231303-0123301021333323-2133321231201331-2020032030023122-0311310221322223-1113301221112130-0002202312032120) |
| `access_info.tls_config.common_params.tls_certificates` | [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-0000322003121202-2123302100002330-1201101301123212-1002103003213232-2321122223001200-1100311300323032-3203203300132001-2030223220312033) |
| `access_info.tls_config.common_params.tls_certificates.certificate_url` | [access_info.tls_config.common_params.tls_certificates.certificate_url](resources--secret_management_access--reference--group-001.md#canonical-3013012321023020-3301020232303303-0302002220002003-0130202322021121-3221332310121033-1103302330023200-3220311121021210-3021101213112122) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](resources--secret_management_access--reference--group-001.md#canonical-1013103223230130-2003133323232103-3103322030200311-1233000001230231-1100313001001103-2102222220221332-3230131010200022-2332321310331212) |
| `access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--secret_management_access--reference--group-001.md#canonical-1002320223323132-2222310220001303-0301222311330203-1310203111000022-1321232312132203-2003231202012133-0033122032021323-0131300002233000) |
| `access_info.tls_config.common_params.tls_certificates.description_spec` | [access_info.tls_config.common_params.tls_certificates.description_spec](resources--secret_management_access--reference--group-001.md#canonical-2002030301030311-0210033133010012-2121030232300100-3302313201113110-0313132013022212-2213312122330202-1102011322130033-1230022230311113) |
| `access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling` | [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](resources--secret_management_access--reference--group-001.md#canonical-0122113310230102-1102001031033302-3131021010002132-3333202120231110-2212220100200210-2203110022033001-2010112130333101-0211320103200302) |
| `access_info.tls_config.common_params.tls_certificates.private_key` | [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-3303311022312211-3101300310203033-2221102101100303-2002222211211333-3210123121002223-2303030102032111-0023120010021132-2222202103221222) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-0132132023022330-2231300031210332-2331221230300113-1110020113332200-1323020222233130-3020322023231023-3022323000003203-1031222322330211) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--secret_management_access--reference--group-001.md#canonical-0212213122231001-1203101303123220-2131110120223200-1310021023112330-3303321322021313-3302101222010230-3011321312002232-1011322212033230) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.location](resources--secret_management_access--reference--group-001.md#canonical-1312322220103301-2102211130110221-2033013311210232-3033020112020230-1121231030222011-3110022112022032-2221221210201020-1002232121200002) |
| `access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--secret_management_access--reference--group-001.md#canonical-3203030113231233-3010221212133223-2331220002111122-3030013013020332-2031303100320132-1321103030030102-3033103302102001-3101112211332221) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-0131233110021013-2221130303202121-2220132100023131-3301000033103221-1110213223132232-0321013003101013-1203320301220220-0033013200223000) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](resources--secret_management_access--reference--group-001.md#canonical-1130011333101112-0003230322322312-0230101130203310-3033223201200023-0001232211203023-3012121111213302-2003021233231023-0000313001131310) |
| `access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url` | [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info.url](resources--secret_management_access--reference--group-001.md#canonical-1223332023303132-3310211122201321-1201112231023133-3222002110021121-0111113133133021-0033021210003212-1310103011330133-0013311100231302) |
| `access_info.tls_config.common_params.tls_certificates.use_system_defaults` | [access_info.tls_config.common_params.tls_certificates.use_system_defaults](resources--secret_management_access--reference--group-001.md#canonical-0232131210233233-2121102011223210-1311213333112330-2312220313212013-0101032232002332-3201232031321301-1012100131302223-0202222321123103) |
| `access_info.tls_config.common_params.validation_params` | [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-3021301002311333-2131313331300332-1212003101212312-0023022102101003-2003320312201231-2330000123321213-3013333011132312-1000030013222320) |
| `access_info.tls_config.common_params.validation_params.skip_hostname_verification` | [access_info.tls_config.common_params.validation_params.skip_hostname_verification](resources--secret_management_access--reference--group-001.md#canonical-1210022112111110-3133113023200112-1011032022100233-0121303023320132-3120200213122032-0111300322032032-3133111220301110-3030223023313322) |
| `access_info.tls_config.common_params.validation_params.trusted_ca` | [access_info.tls_config.common_params.validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-0232320133311111-3231031311213223-0123120113120032-3301210312320011-3110320321131122-1101200102113322-1223330213323303-0331323200023312) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-2131112333022231-1100222202103011-2301201332100303-1110101300112033-1322000302102221-2011210210130020-0332023223210103-0020211113303300) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.kind](resources--secret_management_access--reference--group-001.md#canonical-1003230333331313-3033311313212013-1221321010323213-3122031232321022-2311233300112303-0221202323332012-1010001331311111-1303012020232111) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.name](resources--secret_management_access--reference--group-001.md#canonical-3200301221233030-2030312121322100-1100111032323212-1111230220132103-2020310030203023-3211010122221102-1301102012100311-3311102212232103) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](resources--secret_management_access--reference--group-001.md#canonical-0022202311220012-1021110112130031-0200122212132312-2103302313323102-0001122123330203-2000112230102330-1032210210321220-0100003301202131) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](resources--secret_management_access--reference--group-001.md#canonical-1123230003310003-3201201330111312-2323012010303032-0031231132032231-1221002232010023-2230230001211230-3020310100331112-0000331000000130) |
| `access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list.uid](resources--secret_management_access--reference--group-001.md#canonical-3212230213112132-3013130123310213-1103212301302120-2103122013123003-0110122231012103-2012023212131123-1132322330213031-2302112033213011) |
| `access_info.tls_config.common_params.validation_params.trusted_ca_url` | [access_info.tls_config.common_params.validation_params.trusted_ca_url](resources--secret_management_access--reference--group-001.md#canonical-0223012111020333-1110011022133303-3320233310013300-1123021310230101-1231222021220322-2000110203233321-2031003331321031-3011232313030200) |
| `access_info.tls_config.common_params.validation_params.verify_subject_alt_names` | [access_info.tls_config.common_params.validation_params.verify_subject_alt_names](resources--secret_management_access--reference--group-001.md#canonical-3102312103310212-1002212211020123-3020003102021223-3321023212111302-1333231312201133-0223103130120131-3230221203211020-3103312030001011) |
| `access_info.tls_config.default_session_key_caching` | [access_info.tls_config.default_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-0110010211332033-0022012330313010-0103102130223003-2021213000333002-2102000110112212-2321221233320220-1232110001201230-0312030010130002) |
| `access_info.tls_config.disable_session_key_caching` | [access_info.tls_config.disable_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-1023101211232012-2133221233202220-0333130310310212-3123312103203210-0022000021301203-3011212223213130-1112133101021020-3133112030202123) |
| `access_info.tls_config.disable_sni` | [access_info.tls_config.disable_sni](resources--secret_management_access--reference--group-002.md#canonical-3220101121003121-1210031032210023-3311330011203210-2121100131331213-0310131111032311-0112201222202302-3000201210113112-2320200021033133) |
| `access_info.tls_config.max_session_keys` | [access_info.tls_config.max_session_keys](resources--secret_management_access--reference--group-001.md#canonical-2002310103121301-2132013300230000-0012121102211320-3332332203133023-0123111020021210-0300333323111020-1302203301022003-1122031230330231) |
| `access_info.tls_config.sni` | [access_info.tls_config.sni](resources--secret_management_access--reference--group-001.md#canonical-1133102133000030-2202002223230120-0130212320130123-2013112133120132-3333133113333231-2230201130011102-3221211202320232-2201222022123112) |
| `access_info.tls_config.use_host_header_as_sni` | [access_info.tls_config.use_host_header_as_sni](resources--secret_management_access--reference--group-002.md#canonical-1202101121230132-1102220121002330-3322012213033323-0223032333303303-3312021331322101-3223100320312001-0210301110233331-1010220313001022) |
| `access_info.vault_auth_info` | [access_info.vault_auth_info](resources--secret_management_access--reference--group-002.md#canonical-0012231130130023-2021330132312002-0123100211220102-2320313103311032-0100023300213330-0120211003322132-3020302232210121-1312332033131020) |
| `access_info.vault_auth_info.app_role_auth` | [access_info.vault_auth_info.app_role_auth](resources--secret_management_access--reference--group-002.md#canonical-3323033222310213-2212132222302303-2300300310130212-2321223033301121-1130013311123233-3120313320301211-3310131013103131-2320332111212330) |
| `access_info.vault_auth_info.app_role_auth.role_id` | [access_info.vault_auth_info.app_role_auth.role_id](resources--secret_management_access--reference--group-002.md#canonical-2011231223302313-1102233013010010-3032323221321310-1120031120323020-2013222033120113-3321213010022020-0023130201032233-0001302010320220) |
| `access_info.vault_auth_info.app_role_auth.secret_id` | [access_info.vault_auth_info.app_role_auth.secret_id](resources--secret_management_access--reference--group-002.md#canonical-3221300303331033-1100310300330002-1113112012011212-1132223211313131-1011130012032020-2032102313221320-3310132302012010-0002202032001103) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](resources--secret_management_access--reference--group-002.md#canonical-2222303332023000-0213030232330101-0203332202130311-2111330322321120-3210121302221232-0310200300022013-2221303133203111-2300110333000200) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.decryption_provider](resources--secret_management_access--reference--group-002.md#canonical-1103201211133222-2022321310211322-3023012212132021-0301300023333012-2113121300101231-1122322223310333-3101201220320331-2322103130012020) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.location](resources--secret_management_access--reference--group-002.md#canonical-0100223111200303-2312232133123121-0111033002001000-2013022303120232-0210333301101333-1010311102320311-1211103200221302-0130223233122233) |
| `access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info.store_provider](resources--secret_management_access--reference--group-002.md#canonical-3000103000033310-2312202102203222-1301223100122301-0321203033231020-3123230123020220-1332100302002323-0301101023320300-3220322213122201) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](resources--secret_management_access--reference--group-002.md#canonical-3301212330230011-3030030302020223-1001303301323200-1333201010003033-3212212211012201-1032003223311232-2322221113232302-3321301300223310) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.provider_ref](resources--secret_management_access--reference--group-002.md#canonical-2011101230113322-2222331202313103-2310031023002111-0311023122211200-3322232132100100-3221122303032011-1320301023121231-2030022213020122) |
| `access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url` | [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info.url](resources--secret_management_access--reference--group-002.md#canonical-1220001102011322-1211001303120331-3320023330231122-2300100112111122-2210232001202111-0000310002110233-0131203233122232-3300012023202010) |
| `access_info.vault_auth_info.token` | [access_info.vault_auth_info.token](resources--secret_management_access--reference--group-002.md#canonical-3120322002233003-0212331011010023-3303022100321011-2111102210111013-3322222232232303-1222031012300300-3101013203101020-1213231133102022) |
| `access_info.vault_auth_info.token.blindfold_secret_info` | [access_info.vault_auth_info.token.blindfold_secret_info](resources--secret_management_access--reference--group-002.md#canonical-0022033233301231-1031010202320133-3130133030223221-1220033313033022-2032101001320030-2120123110112333-0002211210210030-1211133131210322) |
| `access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.decryption_provider](resources--secret_management_access--reference--group-002.md#canonical-2200001103122210-3213302133000221-0113330130000201-3031332200102311-2003103122030321-0131132221211202-2111201311133001-0301102112201120) |
| `access_info.vault_auth_info.token.blindfold_secret_info.location` | [access_info.vault_auth_info.token.blindfold_secret_info.location](resources--secret_management_access--reference--group-002.md#canonical-2002220102103120-1333102312110202-2310201013323310-3232130030200303-1030202201333313-1013032000012130-0111230013310221-3111120010021010) |
| `access_info.vault_auth_info.token.blindfold_secret_info.store_provider` | [access_info.vault_auth_info.token.blindfold_secret_info.store_provider](resources--secret_management_access--reference--group-002.md#canonical-2032031202220113-3211110232010110-1320003013022111-1000211220103300-2100212212131312-1011232133113311-1130313230103212-3100323133023103) |
| `access_info.vault_auth_info.token.clear_secret_info` | [access_info.vault_auth_info.token.clear_secret_info](resources--secret_management_access--reference--group-002.md#canonical-3222320230201310-0110300202031310-2311020131213312-3132333232131120-3031302320320301-2200002032213103-2211013003113132-0330103022101013) |
| `access_info.vault_auth_info.token.clear_secret_info.provider_ref` | [access_info.vault_auth_info.token.clear_secret_info.provider_ref](resources--secret_management_access--reference--group-002.md#canonical-1031133233201010-2101320332330202-1300223322200023-2201031301230220-1203220101201031-0112200003032213-0311032300131231-0131332103220003) |
| `access_info.vault_auth_info.token.clear_secret_info.url` | [access_info.vault_auth_info.token.clear_secret_info.url](resources--secret_management_access--reference--group-002.md#canonical-3102301032011000-2131002231311112-2030102120310020-2021323223021020-1321212212320110-3013331003323022-1003001302121022-1313031301311022) |
| `annotations` | [annotations](resources--secret_management_access--reference--group-001.md#canonical-0223313201221312-1013011120202120-3110123320033331-2230130221120023-1330232301213030-1123311022023200-0312311130212312-0321132103211323) |
| `description` | [description](resources--secret_management_access--reference--group-001.md#canonical-2331100222000021-0111310120210223-3030031213202202-2333100211101220-3111001130310222-3301003032303331-3232313210333130-3100223102031020) |
| `disable` | [disable](resources--secret_management_access--reference--group-001.md#canonical-1323230111130030-3211120002023133-3132101123020030-1200303132133212-2203032122202200-2312002221300030-2231203322103033-2203031013102113) |
| `id` | [ID](resources--secret_management_access--reference--group-001.md#canonical-1031201332123121-3211102231022333-0003001202000333-0212310310030010-0100310310323122-2222222112202232-0112011021132121-2330332212021223) |
| `labels` | [labels](resources--secret_management_access--reference--group-001.md#canonical-0220012302320322-0302113133120331-3310230003222102-0230022111121223-0113031031330021-0301331031131332-1100330113221200-2120333200020331) |
| `name` | [name](resources--secret_management_access--reference--group-001.md#canonical-2222323103103221-0200122230200112-2003303013013012-0011330201103130-1013032103311312-3211022131133021-2302003321203302-2323012311201010) |
| `namespace` | [namespace](resources--secret_management_access--reference--group-001.md#canonical-3133031310002231-3111313003203222-3201213203033020-2220001112101120-2010312233332002-1013301332112310-0132102221302133-0112312010000020) |
| `provider_name` | [provider_name](resources--secret_management_access--reference--group-001.md#canonical-1011003103120023-0202232230032130-2323020033032330-0032222230220011-0003200131331201-0213230300330132-0130102233131032-1031111332130303) |
| `timeouts` | [timeouts](resources--secret_management_access--reference--group-002.md#canonical-3233122122221001-2312202221302201-1010010213030313-2131211032312303-1133123022100000-2321313110121320-2333222311323322-2221021302320032) |
| `timeouts.create` | [timeouts.create](resources--secret_management_access--reference--group-002.md#canonical-2113212120312200-0313220020120020-0211221031011231-2003212031232210-2321220033102100-0013032002102223-1030211232000020-1130202110221032) |
| `timeouts.delete` | [timeouts.delete](resources--secret_management_access--reference--group-002.md#canonical-1123131123122002-1233123312231033-0211001110122312-0230031210303331-1100220231131110-3020300213232002-0231000032030212-3222303133131111) |
| `timeouts.read` | [timeouts.read](resources--secret_management_access--reference--group-002.md#canonical-1002013113003230-3002002230012200-1100303321312203-2302010232222000-2222010232022130-3120031220332021-3122333303312323-2103020310212231) |
| `timeouts.update` | [timeouts.update](resources--secret_management_access--reference--group-002.md#canonical-1032311123221330-3330122021302303-1212131230301300-0103022301213131-0202003121332100-0130022313122122-2201303220210131-0201231313331000) |
| `where` | [where](resources--secret_management_access--reference--group-002.md#canonical-2222020230013022-1132023222021231-1021320003032021-2201100230001031-2030331013100203-1020203210021121-2311002100100203-1322032002013221) |
| `where.site` | [where.site](resources--secret_management_access--reference--group-002.md#canonical-2221013003012203-1332011201133123-2303232030033320-1131301220033011-0311113211013200-0101311333101212-0311333312103213-0300130020311300) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-1211223111003203-0122202211311203-1201012323200232-2302220322103103-2331123323320300-2312230301322003-0112300111210001-1210201223000020) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-2100311023212210-1102232230223221-2210002032313010-2131131212330012-2102232111112312-3333211233323222-3310003313132003-1111323132210122) |
| `where.site.network_type` | [where.site.network_type](resources--secret_management_access--reference--group-002.md#canonical-2133013213302313-2320121101202232-0212230133232032-0202200031031311-0203100331121333-2123301320132122-2112301323011112-1110321103233110) |
| `where.site.ref` | [where.site.ref](resources--secret_management_access--reference--group-002.md#canonical-1320013132032111-1220203212322220-1331112031220110-0300320231021031-1203011001211222-3313212110133300-1122231110332323-1332203203322223) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--secret_management_access--reference--group-002.md#canonical-2211020011200301-2211232311132221-2120030010323300-0310123121222120-1332321320213232-3320003302001100-2303230013030122-2221310230222012) |
| `where.site.ref.name` | [where.site.ref.name](resources--secret_management_access--reference--group-002.md#canonical-2322211332302330-1030102303321110-0203030002311110-2003022221303103-0312012330312113-0212021223031212-1200000232022120-2322012030330010) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--secret_management_access--reference--group-002.md#canonical-1011123023031303-2030122112200030-2322320010131213-1110213011332132-3010120022132030-1103013031003220-3113101030123130-1013222200310101) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--secret_management_access--reference--group-002.md#canonical-0011113011123031-2221332322331311-0222003120302103-2310121311323002-1301212321011001-3310132210000013-2100221310030300-3131220121310220) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--secret_management_access--reference--group-002.md#canonical-3230110020233231-1330121120010313-1133311312030003-0321210212131031-1102113121112223-0102310031032313-2132312020022222-1201320030312312) |
| `where.virtual_network` | [where.virtual_network](resources--secret_management_access--reference--group-002.md#canonical-1330333222230132-3013231110212212-1121013333200210-3000203023313023-0232130030113230-1323233331330302-2033333132120333-3231100131122213) |
| `where.virtual_network.ref` | [where.virtual_network.ref](resources--secret_management_access--reference--group-002.md#canonical-2201320010010221-3021211320203123-1003031200323322-1031003100323200-2012222223133132-0200331123122332-3212020010232103-3302320211103232) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](resources--secret_management_access--reference--group-002.md#canonical-0313333132220130-2300033102211000-0132001102133210-2221232022100221-0233133312031133-1322301321323333-0011131122322210-2102230313230213) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](resources--secret_management_access--reference--group-002.md#canonical-3220102222301230-1032331331213212-3233213200031012-2303212023112123-0031230222130311-1220132312033200-3000111123223033-2100300102221120) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](resources--secret_management_access--reference--group-002.md#canonical-2032310033103211-1330233131321310-0320200000300312-1012100323311012-1303121231202102-0203310022331331-0113120201312130-1000132233303321) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](resources--secret_management_access--reference--group-002.md#canonical-2213332131002021-0122113310210030-3313020021222321-3131110313010030-2133231330310222-3031300233331302-3220013222133231-3303120200223222) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](resources--secret_management_access--reference--group-002.md#canonical-1112003301003013-3131233133321003-1330233213200302-3022112133102331-1321012113102211-1211213301330000-2110011102313302-3120032232033032) |
| `where.virtual_site` | [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-3130113011132233-2331020331120121-3331111101001111-1303101020231310-2222101222113212-2222231102212313-0133310332103031-0310123013100121) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-0030100332100002-3302121213031121-2230211333321223-1112222211123122-0302330233200103-1302211012221101-3311002202132013-0131033021103322) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-1313221330113203-2000003300122020-2220130013220132-3001310201212211-0113312031311213-0211323101220312-1132212122003130-0123002303203303) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--secret_management_access--reference--group-002.md#canonical-1021001110312133-0122321112320212-1120030101122101-1021232112012122-0232230311110002-1033311020311232-0123210112201301-3100122232231233) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--secret_management_access--reference--group-002.md#canonical-0223222002022120-3223101332310200-3331011310112033-0331012211101220-2131312031012110-1032232001323322-1013022133310331-0322013223213203) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--secret_management_access--reference--group-002.md#canonical-0202320202032112-0023010100011211-3112303122313333-2122310221100011-3110313300000010-1321311020211002-3123300003123012-2102233113230020) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--secret_management_access--reference--group-002.md#canonical-3023013233232231-1211202113013011-1002323332231322-1102313332103131-0110012213311120-3230301131010332-1133032230111121-1002032020122210) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--secret_management_access--reference--group-002.md#canonical-3032002121222000-3022302132233212-2022033313201212-1020030001331302-0213131100201333-2130031210222331-1132322213203131-3130001230032203) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--secret_management_access--reference--group-002.md#canonical-0203301021331121-2022332330221013-2200221003223022-0020130021222232-2222023232001333-2333031110030212-0012111200323223-2122012310011111) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--secret_management_access--reference--group-002.md#canonical-1232031202132310-3310132131022220-3110112032033303-1033031103222322-0023001311032032-2203321320332033-3300133310303213-2331333311302303) |

<a id="canonical-2130323230102221-0132010111311002-0002332131032332-0120312113312201-3300233230321122-1012121112002100-3111221201331010-3313113211211100"></a>

## Next pages — Property reference / 031200003023 / 13

- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [timeouts](resources--secret_management_access--reference--group-002.md#canonical-3323013113011220-0112002023230221-0022123301231112-0311033122021321-0012233310113333-1120230302310030-3302113003112103-0010202130123300)
- [where](resources--secret_management_access--reference--group-002.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212101331230111-1303123010210201-0231011313131011-0101231013003130-2311001021130323-2131211311312303-2001220123021333-3133223021300000"></a>

## access_info — access_info / 020212331110 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- access_info

<a id="canonical-1222032012011203-3310232220222312-1031202321200133-0022312310213330-0012300111113302-0120301322020011-1110013030330123-0013313102002203"></a>

Type: `"object"`. single nested block, Optional.

HostAccessInfoType contains the information about how to connect to the remote host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server_endpoint"),
  validators.ConflictingObjectAttributes("rest_auth_info",
    "vault_auth_info")}
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
  "x-ves-oneof-field-auth_params": "[\"rest_auth_info\",\"vault_auth_info\"]"
}
```

Terraform syntax:

```terraform
access_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333213001323333-3330010113103232-0311132303200310-3101213121100220-2123112123210120-0231233333101332-3312300212202001-0110301313222131"></a>

## Direct properties — access_info / 020212331110 / 3

- [rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310): complete subsection reference.

<a id="canonical-0112112121002033-0021232203211000-1103303203313211-1211003012013213-2200113321213230-0331030332313310-3323312001101023-3003031203033102"></a>

<a id="canonical-0220221112330130-2213110132021020-2001030221012223-0103111122331312-2021132022000212-3130231220231323-3113021111213320-1303203000031123"></a>

## scheme property — access_info / 020212331110 / 4

Type: `"string"`. Optional.

\[Enum: HTTP|HTTPS\] SchemeType is used to indicate URL scheme HTTP:// scheme HTTPS:// scheme.
Possible values are \`HTTP\`, \`HTTPS\`. Defaults to \`HTTP\`.

Upstream description:

SchemeType is used to indicate URL scheme

HTTP:// scheme HTTPS:// scheme.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HTTP",
    "HTTPS"),
}
```

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

<a id="canonical-1330310033321121-0133103101003021-1231202220233301-0313100232020220-3023023230321221-2300122113211121-2001131030011123-0001113213122212"></a>

<a id="canonical-1311320122030302-3032213210300320-0133022132302021-2131312021131111-0020312120233332-3111203111000202-2201300032232200-0002222001032112"></a>

## server_endpoint property — access_info / 020212331110 / 5

Type: `"string"`. Optional.

Endpoint to connect to, in host:port format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121): complete subsection reference.

- [vault_auth_info](resources--secret_management_access--reference--group-002.md#canonical-2300210120332323-0231220231322210-3032002112232110-2333033211233200-1011201213321300-0120301210302002-0332332302303131-3032233113000312): complete subsection reference.

<a id="canonical-2132101321201002-0032033311301032-3210011111203311-1212112030023301-0210320031110110-1230312202221203-0223211323203011-1132301302133230"></a>

## Next pages — access_info / 020212331110 / 6

- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.vault_auth_info](resources--secret_management_access--reference--group-002.md#canonical-2300210120332323-0231220231322210-3032002112232110-2333033211233200-1011201213321300-0120301210302002-0332332302303131-3032233113000312)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001000110320101-3311220011011222-0332133213031320-1120012303101110-3120113112030033-0230201132010020-2102010332002130-1010332023233121"></a>

## access_info.rest_auth_info — rest_auth_info / 022011322111 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- access_info.rest_auth_info

<a id="canonical-1031033332300221-3011202212322003-1132000321310013-2020130113100003-3213331100301023-1102113301213002-2030123210323312-2111013321013330"></a>

Type: `"object"`. single nested block, Optional.

Authentication parameters for REST based hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("basic_auth",
    "headers_auth"),
  validators.ConflictingObjectAttributes("basic_auth",
    "query_params_auth"),
  validators.ConflictingObjectAttributes("headers_auth",
    "query_params_auth")}
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
  "x-ves-oneof-field-auth_params": "[\"basic_auth\",\"headers_auth\",\"query_params_auth\"]"
}
```

Terraform syntax:

```terraform
rest_auth_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100112002222032-1321330222313210-1030331303202200-2233021033032213-3210012332201133-1220222111002102-0111112222012022-3002230212112030"></a>

## Direct properties — rest_auth_info / 022011322111 / 3

- [basic_auth](resources--secret_management_access--reference--group-001.md#canonical-0001321320301202-3033201003201202-1122100122332311-3102310022031223-2101011312301203-2132121222012011-2302303000311101-2212332023131210): complete subsection reference.

- [headers_auth](resources--secret_management_access--reference--group-001.md#canonical-0232120201223023-3331321121223301-2021133113113131-1112013011310213-1331121210022012-2210021030021232-3110112203003031-3103213201023312): complete subsection reference.

- [query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-1032023300203021-1102132320110003-2303022031202333-1031100030012103-2230031103203233-0222012031222211-1112333310230121-1030332213021130): complete subsection reference.

<a id="canonical-2321213031320221-1332212320201022-1100033303120213-0023323200111230-0031301232330220-0123012122123110-2110230331232213-1021032011011231"></a>

## Next pages — rest_auth_info / 022011322111 / 4

- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-0001321320301202-3033201003201202-1122100122332311-3102310022031223-2101011312301203-2132121222012011-2302303000311101-2212332023131210)
- [access_info.rest_auth_info.headers_auth](resources--secret_management_access--reference--group-001.md#canonical-0232120201223023-3331321121223301-2021133113113131-1112013011310213-1331121210022012-2210021030021232-3110112203003031-3103213201023312)
- [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-1032023300203021-1102132320110003-2303022031202333-1031100030012103-2230031103203233-0222012031222211-1112333310230121-1030332213021130)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0001321320301202-3033201003201202-1122100122332311-3102310022031223-2101011312301203-2132121222012011-2302303000311101-2212332023131210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310032321100131-2323013220101111-0310303203101312-2132012120100023-2020200302300331-1111330130122132-2302101200233333-2303211230231230"></a>

## access_info.rest_auth_info.basic_auth — basic_auth / 103021233032 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- access_info.rest_auth_info.basic_auth

<a id="canonical-2312112200321323-1122001233003312-2201230022122021-0332132020212221-2201102100030123-3102221033311302-1030311120003011-2002033003032222"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
basic_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231011323123102-0012300121132322-1100012311211103-0132221330000200-3020210032330200-0220132110312231-3013311221002202-2023233111102023"></a>

## Direct properties — basic_auth / 103021233032 / 3

- [password](resources--secret_management_access--reference--group-001.md#canonical-0012333222302123-0233222222311132-0111101201331212-0120100200313103-1301033001003000-2130213023301310-2000322120231021-1233022223101201): complete subsection reference.

<a id="canonical-2233333030313032-0121032210000010-3012010133010211-0223312113032231-3331131231202030-2101303212320221-3231231210333010-1213310131332003"></a>

<a id="canonical-0223020213321232-1121212233311002-2121113100121323-0221221330333121-2013020213033030-1121032021221300-1201320303003323-0331300321211120"></a>

## username property — basic_auth / 103021233032 / 4

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3203033200320010-0120200301102321-3303203312133022-2121300000002130-1112130202212111-2013323311331102-2102332232033122-0230103101111110"></a>

## Next pages — basic_auth / 103021233032 / 5

- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-0012333222302123-0233222222311132-0111101201331212-0120100200313103-1301033001003000-2130213023301310-2000322120231021-1233022223101201)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0012333222302123-0233222222311132-0111101201331212-0120100200313103-1301033001003000-2130213023301310-2000322120231021-1233022223101201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333031133223013-2202030323023211-0000013102003312-0201213313021231-0132101101200122-2032133313003210-1322220101332312-1310312231313102"></a>

## access_info.rest_auth_info.basic_auth.password — password / 331320133133 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-0001321320301202-3033201003201202-1122100122332311-3102310022031223-2101011312301203-2132121222012011-2302303000311101-2212332023131210)
- access_info.rest_auth_info.basic_auth.password

<a id="canonical-0002311202202012-3130010332330123-1313023111233013-3233111031133232-2012332233101133-3130312001201111-0300203111202301-2220122232333021"></a>

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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232222213022210-2032121231021232-3212030130211012-3311223222021302-0220112031301023-2131120301302213-1321121003021122-1323312013321111"></a>

## Direct properties — password / 331320133133 / 3

- [blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-3011021101113222-3312312130121303-1133301213100323-2103020132331121-2230232020211023-1120201011222232-3323322211330022-0013302203332233): complete subsection reference.

- [clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-0120133001312011-0211303012020131-0230100102301300-2131203232221300-0202303221121231-1133201102323311-3211100222313112-3222201012122211): complete subsection reference.

<a id="canonical-0322313101200322-2203021103133031-3001112110320301-2322012303010312-1210112021102030-0211121032203311-1120131303313102-2303312102120231"></a>

## Next pages — password / 331320133133 / 4

- [access_info.rest_auth_info.basic_auth.password.blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-3011021101113222-3312312130121303-1133301213100323-2103020132331121-2230232020211023-1120201011222232-3323322211330022-0013302203332233)
- [access_info.rest_auth_info.basic_auth.password.clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-0120133001312011-0211303012020131-0230100102301300-2131203232221300-0202303221121231-1133201102323311-3211100222313112-3222201012122211)
- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-0001321320301202-3033201003201202-1122100122332311-3102310022031223-2101011312301203-2132121222012011-2302303000311101-2212332023131210)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-3011021101113222-3312312130121303-1133301213100323-2103020132331121-2230232020211023-1120201011222232-3323322211330022-0013302203332233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020013013220033-1032301320110133-3303130130323021-2132311001201203-0113013210023010-1111132103002200-2003122021222121-1031330133012022"></a>

## access_info.rest_auth_info.basic_auth.password.blindfold_secret_info — blindfold_secret_info / 010333133010 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-0001321320301202-3033201003201202-1122100122332311-3102310022031223-2101011312301203-2132121222012011-2302303000311101-2212332023131210)
- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-0012333222302123-0233222222311132-0111101201331212-0120100200313103-1301033001003000-2130213023301310-2000322120231021-1233022223101201)
- access_info.rest_auth_info.basic_auth.password.blindfold_secret_info

<a id="canonical-2033200233022313-2302112120213212-3302333011220221-0222012023002203-2301323311103320-0222230310122023-0101330222011022-1211132333301011"></a>

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

<a id="canonical-1111020231302331-0202220233201020-2120332121113023-2233002333120231-2030323331331031-1012211201101213-1122010232013323-2023122300001122"></a>

## Direct properties — blindfold_secret_info / 010333133010 / 3

<a id="canonical-1333020331201002-2330300021032321-2021030023320010-3202201201131011-3000131102202022-3333303300320020-0000220332023021-2102312330003133"></a>

<a id="canonical-2312203232200100-3022002312210332-2033320002112132-2211213312333112-1201001200132203-2231111111110303-2033200133202213-3011033121321230"></a>

## decryption_provider property — blindfold_secret_info / 010333133010 / 4

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

<a id="canonical-3231311122012231-0211231222131003-1311012023323333-3311213333002222-1030203022301101-2320030110323322-1311302321012213-1002031023023122"></a>

<a id="canonical-1311003202323231-1001022211233210-0102201010303212-3021121023131211-3130032130113130-3012110331300211-3222133132010311-2101203023120002"></a>

## location property — blindfold_secret_info / 010333133010 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-0002131330231002-1022323030011200-1311121003002321-1213321133213212-3200033022000010-0030100101222033-1013031023211121-3200230303220211"></a>

<a id="canonical-0303330102200202-0121221220201021-2203102121311011-3213100232231100-2032222312002131-2001112232333200-2200302232032303-3223230101010121"></a>

## store_provider property — blindfold_secret_info / 010333133010 / 6

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

<a id="canonical-2103130232211300-3313002223202303-2303310000310200-3310012131203032-3110003112112201-1320110032001021-1303033202120320-2223323120120123"></a>

## Next pages — blindfold_secret_info / 010333133010 / 7

- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-0012333222302123-0233222222311132-0111101201331212-0120100200313103-1301033001003000-2130213023301310-2000322120231021-1233022223101201)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0120133001312011-0211303012020131-0230100102301300-2131203232221300-0202303221121231-1133201102323311-3211100222313112-3222201012122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032313031333330-1320211210131031-3130203213230210-2330322202221022-1220013211110310-0121032010120311-1302333321321322-3200210222023311"></a>

## access_info.rest_auth_info.basic_auth.password.clear_secret_info — clear_secret_info / 302201320230 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- [access_info.rest_auth_info.basic_auth](resources--secret_management_access--reference--group-001.md#canonical-0001321320301202-3033201003201202-1122100122332311-3102310022031223-2101011312301203-2132121222012011-2302303000311101-2212332023131210)
- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-0012333222302123-0233222222311132-0111101201331212-0120100200313103-1301033001003000-2130213023301310-2000322120231021-1233022223101201)
- access_info.rest_auth_info.basic_auth.password.clear_secret_info

<a id="canonical-1103300310200022-3101020110020223-1120220122302230-0100132021223003-0221101030213301-1021213303202011-3322233221001121-0303212003213013"></a>

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

<a id="canonical-2322221222320313-0033012232113010-3103230321302211-0320202112212333-0123300303303312-1213013032333033-3011020220113220-1301011201001312"></a>

## Direct properties — clear_secret_info / 302201320230 / 3

<a id="canonical-1310001110201212-3331132322211011-3111131211312133-3001312312011313-0201103001101233-2123320130230213-3310023021231011-0210321010310023"></a>

<a id="canonical-3120033301221311-0111122300000122-0022011211133121-0103021113331001-3030112100213031-0320130211230322-0320102222011131-0013033332030123"></a>

## provider_ref property — clear_secret_info / 302201320230 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0101013221311001-1220212010122303-3212212333303331-0323300011301132-3320021202002231-2233121010102122-0330210110023101-0201002030231102"></a>

<a id="canonical-2100000111202123-0020321020301231-0032120220202102-0220311310020302-2223122331110333-2212213330022333-3133120220302132-3013112303123320"></a>

## URL property — clear_secret_info / 302201320230 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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

<a id="canonical-2010111333312223-1000220231110310-1330013302303100-3012331120301021-3323100120110133-0120201121203223-2233213110023000-3111202020130210"></a>

## Next pages — clear_secret_info / 302201320230 / 6

- [access_info.rest_auth_info.basic_auth.password](resources--secret_management_access--reference--group-001.md#canonical-0012333222302123-0233222222311132-0111101201331212-0120100200313103-1301033001003000-2130213023301310-2000322120231021-1233022223101201)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0232120201223023-3331321121223301-2021133113113131-1112013011310213-1331121210022012-2210021030021232-3110112203003031-3103213201023312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203310322100222-2330031301122321-2300120301311112-3110223002013033-2323203213231032-0021201001331200-2302023021332032-0133310010123022"></a>

## access_info.rest_auth_info.headers_auth — headers_auth / 233332030132 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- access_info.rest_auth_info.headers_auth

<a id="canonical-2220120110122103-3200302102111032-0112003222322013-0210101221001332-2210220131203331-2310221331223230-2300303020332311-3320121113021013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
headers_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031112013011332-2302002010202020-1133310310021120-2130013131211222-1200310231110021-1320332001000300-0113111320110311-1112323120301221"></a>

## Direct properties — headers_auth / 233332030132 / 3

- [headers](resources--secret_management_access--reference--group-001.md#canonical-3212121233222013-2011230321321201-3032220201230031-0022100302330010-3012203130103033-1333302302323022-2212321011111022-1121223120111032): complete subsection reference.

<a id="canonical-2332103211210110-1001330321302230-1101310101100221-0201233313110122-1130111230032102-2031133113120222-2323311201310302-3323032230320020"></a>

## Next pages — headers_auth / 233332030132 / 4

- [access_info.rest_auth_info.headers_auth.headers](resources--secret_management_access--reference--group-001.md#canonical-3212121233222013-2011230321321201-3032220201230031-0022100302330010-3012203130103033-1333302302323022-2212321011111022-1121223120111032)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-3212121233222013-2011230321321201-3032220201230031-0022100302330010-3012203130103033-1333302302323022-2212321011111022-1121223120111032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132320000131130-1220010332002232-0332000300100013-2020302132121020-1010130320232301-3300333231010103-1200200313133021-2210112002233133"></a>

## access_info.rest_auth_info.headers_auth.headers — headers / 301002200331 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- [access_info.rest_auth_info.headers_auth](resources--secret_management_access--reference--group-001.md#canonical-0232120201223023-3331321121223301-2021133113113131-1112013011310213-1331121210022012-2210021030021232-3110112203003031-3103213201023312)
- access_info.rest_auth_info.headers_auth.headers

<a id="canonical-1000022023333130-1313222011110232-2012311301232223-3301113030302303-1221133112111223-0011101030011331-0303213212200221-0222030131022320"></a>

Type: `"object"`. single nested block, Optional.

The set of authentication headers to pass in HTTP request.

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

Terraform syntax:

```terraform
headers {}
```

<a id="canonical-3113300303032102-1020002310222002-0310230113322131-1332233300331020-2300310211001130-1222300102102112-0122032331321113-0333303013003132"></a>

## Direct properties — headers / 301002200331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101013222313312-3131001112303222-0323233023113203-0231210302111303-1333012010030300-0202323222201113-0210322213123223-3021320030111120"></a>

## Next pages — headers / 301002200331 / 4

- [access_info.rest_auth_info.headers_auth](resources--secret_management_access--reference--group-001.md#canonical-0232120201223023-3331321121223301-2021133113113131-1112013011310213-1331121210022012-2210021030021232-3110112203003031-3103213201023312)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-1032023300203021-1102132320110003-2303022031202333-1031100030012103-2230031103203233-0222012031222211-1112333310230121-1030332213021130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321010223303301-0221122101321120-1322313011303200-3323230211330201-0031330110021010-1000103220011211-2300210202101010-0001232211021011"></a>

## access_info.rest_auth_info.query_params_auth — query_params_auth / 222331223222 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- access_info.rest_auth_info.query_params_auth

<a id="canonical-3222020203131020-2220120113220003-1111010301112010-2233213302112202-2013332030113221-0131300221323100-2233102303233231-2030201332120102"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
query_params_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030103023003003-0222023121131031-3232312123323102-1200112021301121-3310312222003003-2233310232122202-3121131111001132-1223100312312333"></a>

## Direct properties — query_params_auth / 222331223222 / 3

- [query_params](resources--secret_management_access--reference--group-001.md#canonical-0100012211300003-2112213212002200-3011230321002302-3102330211003103-3020130020003113-0002322113122222-2012010123233122-3122021032200000): complete subsection reference.

<a id="canonical-1130111033230120-1210121320013220-0123312322101012-2222321322233102-0031211030323110-0110120131302213-3010333133121230-2303301112113020"></a>

## Next pages — query_params_auth / 222331223222 / 4

- [access_info.rest_auth_info.query_params_auth.query_params](resources--secret_management_access--reference--group-001.md#canonical-0100012211300003-2112213212002200-3011230321002302-3102330211003103-3020130020003113-0002322113122222-2012010123233122-3122021032200000)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0100012211300003-2112213212002200-3011230321002302-3102330211003103-3020130020003113-0002322113122222-2012010123233122-3122021032200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302212130120103-1101332003100312-1212212313231330-3333001101033032-0032121100103101-1330112023230001-0221020231031021-2203130310212211"></a>

## access_info.rest_auth_info.query_params_auth.query_params — query_params / 013110133210 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.rest_auth_info](resources--secret_management_access--reference--group-001.md#canonical-3321002110312321-0002233221222232-3131302003120121-0000203100321120-2023210130021313-1320010001231210-3123312000311000-0100310020022310)
- [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-1032023300203021-1102132320110003-2303022031202333-1031100030012103-2230031103203233-0222012031222211-1112333310230121-1030332213021130)
- access_info.rest_auth_info.query_params_auth.query_params

<a id="canonical-0330312022210021-0031110333310032-0033321211330120-2330220222133313-2213132111200310-0221231302301233-2321020301013301-3230113322123302"></a>

Type: `"object"`. single nested block, Optional.

The set of authentication parameters to be passed as query parameters.

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

Terraform syntax:

```terraform
query_params {}
```

<a id="canonical-2221211020330012-2003301110023023-3203211232311220-2113233000132022-0112101220231301-2011213332113321-0001130301121010-0021201110121311"></a>

## Direct properties — query_params / 013110133210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313222200323130-2212033310001131-0113210033111303-1110201022101220-0331222110121312-0003322121000013-1213112103100321-0211033131003303"></a>

## Next pages — query_params / 013110133210 / 4

- [access_info.rest_auth_info.query_params_auth](resources--secret_management_access--reference--group-001.md#canonical-1032023300203021-1102132320110003-2303022031202333-1031100030012103-2230031103203233-0222012031222211-1112333310230121-1030332213021130)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201110002320230-2030123210213130-2112301023323032-0132030211131132-2230313112031012-2313333230300213-1130001130000002-1120202020211133"></a>

## access_info.tls_config — tls_config / 020012110220 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- access_info.tls_config

<a id="canonical-1002202302301223-0220232301022112-3031231201312303-2200101231130303-3300121202320210-3012312112002220-2111002002302311-2011331120032110"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for upstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cert_params",
    "common_params"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]",
  "x-ves-oneof-field-tls_params_choice": "[\"cert_params\",\"common_params\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303012120113100-1220313001122332-0100010220331322-0122110320231103-3133132131121003-0100212130201002-3123332310303102-3032321333003122"></a>

## Direct properties — tls_config / 020012110220 / 3

- [cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220): complete subsection reference.

- [common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120): complete subsection reference.

- [default_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-1010002310111333-3331111312010000-2130000001223131-0121101203110301-1123333102202232-1230231311001011-2331030210020320-2212211003301212): complete subsection reference.

- [disable_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-2203002320210003-3112320211113020-1231121133010311-1233120131013313-3001311220012032-3222010112020103-2122300120312001-3113010110113312): complete subsection reference.

- [disable_sni](resources--secret_management_access--reference--group-001.md#canonical-0332120001232113-3013023320312311-2102123303203030-3020002001020103-3033121211320030-0303331223210211-1231303302332132-2220202112120212): complete subsection reference.

<a id="canonical-2002310103121301-2132013300230000-0012121102211320-3332332203133023-0123111020021210-0300333323111020-1302203301022003-1122031230330231"></a>

<a id="canonical-1310133131031212-3120333103132302-0032013302211012-1210303330111022-1030121223213101-0102311003101210-2232200223302210-1232200302020132"></a>

## max_session_keys property — tls_config / 020012110220 / 4

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1133102133000030-2202002223230120-0130212320130123-2013112133120132-3333133113333231-2230201130011102-3221211202320232-2201222022123112"></a>

<a id="canonical-0203013021111230-3200012203113002-3333330011003001-2213130302212321-1233101012012300-0320011013022301-0212313232031330-2011323102130203"></a>

## sni property — tls_config / 020012110220 / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [use_host_header_as_sni](resources--secret_management_access--reference--group-002.md#canonical-0131132312113121-0002111223100322-1223133230330302-0112001130111232-1022232200200232-2011331323211022-3003013122220221-0332220231132233): complete subsection reference.

<a id="canonical-1212021221333302-0332320033011202-2031012022323332-0232011223323322-2211133210111010-0131111320230221-2213021013310031-3302011113032001"></a>

## Next pages — tls_config / 020012110220 / 6

- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [access_info.tls_config.default_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-1010002310111333-3331111312010000-2130000001223131-0121101203110301-1123333102202232-1230231311001011-2331030210020320-2212211003301212)
- [access_info.tls_config.disable_session_key_caching](resources--secret_management_access--reference--group-001.md#canonical-2203002320210003-3112320211113020-1231121133010311-1233120131013313-3001311220012032-3222010112020103-2122300120312001-3113010110113312)
- [access_info.tls_config.disable_sni](resources--secret_management_access--reference--group-001.md#canonical-0332120001232113-3013023320312311-2102123303203030-3020002001020103-3033121211320030-0303331223210211-1231303302332132-2220202112120212)
- [access_info.tls_config.use_host_header_as_sni](resources--secret_management_access--reference--group-002.md#canonical-0131132312113121-0002111223100322-1223133230330302-0112001130111232-1022232200200232-2011331323211022-3003013122220221-0332220231132233)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323022030230131-1212031320100032-3130310121030331-3111331203012230-2023203330013313-3312023233211022-2102203021102132-2210131300022232"></a>

## access_info.tls_config.cert_params — cert_params / 321313000312 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- access_info.tls_config.cert_params

<a id="canonical-0313130300021200-1332201110230311-3302213323223330-2030021203103023-3002223212022302-2021220321320033-3013111210213010-0322130023320303"></a>

Type: `"object"`. single nested block, Optional.

Certificate Parameters for authentication, TLS ciphers, and trust store.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "tls_validation_params"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("tls_validation_params",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"tls_validation_params\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302130323020132-0132201013030123-1113120003133101-3120323330130212-1312311332103232-0201103332010103-0031331032111011-2111130031312230"></a>

## Direct properties — cert_params / 321313000312 / 3

- [certificates](resources--secret_management_access--reference--group-001.md#canonical-3312332300331200-3331110220030031-3103112121301103-3231130233013113-2300320033021120-2331321033322311-1310230202031310-1211311013311133): complete subsection reference.

<a id="canonical-3103022213331301-1233101222200022-1330003330200330-3032101130023002-2323201233002221-1120230222113332-2100100101010333-2202311102002030"></a>

<a id="canonical-3321202120100210-2133203331201210-1302301221000211-0021313112322221-0313103220131330-1210223203332222-2223220300313332-0012311211210100"></a>

## cipher_suites property — cert_params / 321313000312 / 4

Type: `["list", "string"]`. Optional.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1211221210221213-3131230201122321-0230221033231020-3321300320200210-1133010110031030-2113323003030302-1323031103330000-2010002031331133"></a>

<a id="canonical-2201121322330013-1210130001331111-2231200133013310-3011121100203032-2223003103223220-0122232010301222-1213132301231230-2213102111310210"></a>

## maximum_protocol_version property — cert_params / 321313000312 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-1100312201321103-3212322330012303-3222313120010231-0332221100330223-3310031300223211-3202102301320202-2333213313032213-1120000300311031"></a>

<a id="canonical-3133232122210032-3003011223103302-1333000123030223-1003213301300210-3233202311323101-1100322202322021-2201103023323102-1101201210020130"></a>

## minimum_protocol_version property — cert_params / 321313000312 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

- [skip_server_verification](resources--secret_management_access--reference--group-001.md#canonical-1313321312000033-2021120221110111-2033102221321322-3102122033202332-0020102011311023-2121323111211113-2130230310300021-2303130330222000): complete subsection reference.

- [tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-0121031320112211-1123210202002222-2222232302211030-2030211033120231-2211100103310122-1110011313303011-1233131133202213-2132123300322200): complete subsection reference.

- [volterra_trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-1001231221331313-1123332311100232-3211331021222023-3301333132002033-3030012122213310-0002033321033003-1032303110210330-1311012313121221): complete subsection reference.

<a id="canonical-1211223110102320-0231130230323011-0331101230320010-0010012011101330-1221211331112222-1203212321331302-0203031230003210-1000123113333012"></a>

## Next pages — cert_params / 321313000312 / 7

- [access_info.tls_config.cert_params.certificates](resources--secret_management_access--reference--group-001.md#canonical-3312332300331200-3331110220030031-3103112121301103-3231130233013113-2300320033021120-2331321033322311-1310230202031310-1211311013311133)
- [access_info.tls_config.cert_params.skip_server_verification](resources--secret_management_access--reference--group-001.md#canonical-1313321312000033-2021120221110111-2033102221321322-3102122033202332-0020102011311023-2121323111211113-2130230310300021-2303130330222000)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-0121031320112211-1123210202002222-2222232302211030-2030211033120231-2211100103310122-1110011313303011-1233131133202213-2132123300322200)
- [access_info.tls_config.cert_params.volterra_trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-1001231221331313-1123332311100232-3211331021222023-3301333132002033-3030012122213310-0002033321033003-1032303110210330-1311012313121221)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-3312332300331200-3331110220030031-3103112121301103-3231130233013113-2300320033021120-2331321033322311-1310230202031310-1211311013311133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012331222113300-0120010322201112-1313010031000312-2000300003132032-0202111212003011-2020103003223020-2130123320332200-1300110213212103"></a>

## access_info.tls_config.cert_params.certificates — certificates / 301023011321 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- access_info.tls_config.cert_params.certificates

<a id="canonical-3132302102333311-1131103200100330-2320102112323002-2322301231322032-0223010221112333-0133002132221031-3111220303323202-2233333231130101"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323103133131322-1301120322102001-1012131031232030-1110032312020110-0212232320213032-0220122221230200-2130213311132311-1222321101111133"></a>

## Direct properties — certificates / 301023011321 / 3

<a id="canonical-0110332320230120-2010112300312211-1032023331310213-0112312132200321-1003130130022221-0111013032010101-3110000032100322-0321301023111313"></a>

<a id="canonical-1230200323031122-1130100030121332-0030100230021320-1323113322031023-3200212012121000-2233123201331323-0301320130030131-1110211110003112"></a>

## kind property — certificates / 301023011321 / 4

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

<a id="canonical-1121102321112131-2031223321320130-3211121033212232-3322120010310010-3333010222011302-2111322220102101-0231211230211303-1233321322321031"></a>

<a id="canonical-3212232121032101-3331112122210102-3313013110031222-0331122030103032-0032022321320010-0331003220321313-1100321010123320-0330100211223302"></a>

## name property — certificates / 301023011321 / 5

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

<a id="canonical-0112313200133110-2002331002030222-3233330112031322-3030322103302132-3012020323223302-3130210011303110-0221313321100200-0032232030300011"></a>

<a id="canonical-3323012033112130-3003002030103122-0213032330220321-3131210221030311-1312000221232323-1130130212033313-3200330312310231-0132222300112121"></a>

## namespace property — certificates / 301023011321 / 6

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

<a id="canonical-2313021232133013-1320103023123122-1222012030120032-1100210033000111-1101223133021231-1210113311302210-1020021112332021-3103132312333210"></a>

<a id="canonical-3202333232031231-3320003113012013-0313033121023021-2223201023233023-2213331033032003-3123311310103110-3312031232301112-3203103012013013"></a>

## tenant property — certificates / 301023011321 / 7

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

<a id="canonical-0031222021001232-3220312300133022-0013000312212021-2131002110131232-3003121212012032-2020331310233023-1111112123231112-3111102002013312"></a>

<a id="canonical-2230020000332122-0002201333311000-3100032321023012-3331220121333210-2033110023230110-0032003322120220-0113012313301001-1133021102231301"></a>

## uid property — certificates / 301023011321 / 8

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

<a id="canonical-3001021233232030-1101302202223110-2010002030212301-1001310120233203-2200032121303203-0130113012332220-3201003000110032-0132102012131112"></a>

## Next pages — certificates / 301023011321 / 9

- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-1313321312000033-2021120221110111-2033102221321322-3102122033202332-0020102011311023-2121323111211113-2130230310300021-2303130330222000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121002303012120-1230300302100331-2332003213012321-3300011102132111-3020210100312322-0031321330120233-0100221031331120-0313332102202102"></a>

## access_info.tls_config.cert_params.skip_server_verification — skip_server_verification / 133322203200 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- access_info.tls_config.cert_params.skip_server_verification

<a id="canonical-0222232033020302-0021300031013113-3202313210310213-3310120210333010-0311032013233233-3011100231013002-3331010120033123-3000230003231103"></a>

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
skip_server_verification = {}
```

<a id="canonical-2123010011311023-2202201101213131-2020033111311231-1033230022200312-2310000200233022-2000301212320000-2131131210323221-1021303021102231"></a>

## Direct properties — skip_server_verification / 133322203200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203320310030200-1333011113110021-0003220123021200-1221213333013301-2331233112230322-3021222321133200-1220120222100031-3333223030211210"></a>

## Next pages — skip_server_verification / 133322203200 / 4

- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0121031320112211-1123210202002222-2222232302211030-2030211033120231-2211100103310122-1110011313303011-1233131133202213-2132123300322200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001132103201230-1103211312221022-0102130212232300-2032122211132201-0110302200031301-0213232203133333-0313102233033003-0013121202022303"></a>

## access_info.tls_config.cert_params.tls_validation_params — tls_validation_params / 230132230122 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- access_info.tls_config.cert_params.tls_validation_params

<a id="canonical-1100220133013321-0010330320201003-3332010020121220-2201133132023212-3300300321213212-1012323330123112-3211322300212103-3103333112022213"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
tls_validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113323023323311-3220310332310223-3000203113230203-0203333130211112-2222032332012232-1003331123121032-2030010321131320-3011200231133120"></a>

## Direct properties — tls_validation_params / 230132230122 / 3

<a id="canonical-2013320200302230-0323322100100031-3211230032101113-2112313031300003-3211021012211202-1001032020201103-2311222311121222-1231123000133111"></a>

<a id="canonical-2010113023310322-1233001103133332-3302231232002211-0213213201123121-2010130000100201-2133212010332331-0211310001110310-1031312311210021"></a>

## skip_hostname_verification property — tls_validation_params / 230132230122 / 4

Type: `"bool"`. Optional.

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

- [trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-2031030213320222-0001001333133110-0130332232123201-0022203002221000-2200312220313122-0103013012013030-1313302100301021-0233110310001321): complete subsection reference.

<a id="canonical-1201210002301130-1030010022212013-1012013323013021-2332333302102112-1221332130303331-1313123211010123-1301132202131011-0312321033110300"></a>

<a id="canonical-3121303121002320-2010130131002232-2110130122001131-0133212110213133-0032323133222303-0111203023003100-2022033321111222-2330311231133032"></a>

## trusted_ca_url property — tls_validation_params / 230132230122 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-3301221121101013-2311333121322323-1100132013012100-3331132011103232-2100230002320111-2232132220113210-1313132303133302-3211302020221321"></a>

<a id="canonical-1011210030303111-0222221022033310-0212212220101113-2331313121320200-3220230321001331-1010210322201030-1110100333133031-0300113032320332"></a>

## verify_subject_alt_names property — tls_validation_params / 230132230122 / 6

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1300112010233203-3131122003330012-0100132111010113-0213022000203302-2120210103020230-3023122320132002-2021023100303011-0220001201110000"></a>

## Next pages — tls_validation_params / 230132230122 / 7

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-2031030213320222-0001001333133110-0130332232123201-0022203002221000-2200312220313122-0103013012013030-1313302100301021-0233110310001321)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-2031030213320222-0001001333133110-0130332232123201-0022203002221000-2200312220313122-0103013012013030-1313302100301021-0233110310001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322231322121220-1001133010222222-3120012311331220-1321301300030133-1301321233203202-0122211022302312-1002003012133122-1311010321122132"></a>

## access_info.tls_config.cert_params.tls_validation_params.trusted_ca — trusted_ca / 133102300111 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-0121031320112211-1123210202002222-2222232302211030-2030211033120231-2211100103310122-1110011313303011-1233131133202213-2132123300322200)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca

<a id="canonical-1010303223020121-2012020203121300-2203103023020202-2321331103231023-1313320332211030-1112210133231323-0020320322200103-2111322021131000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130321302202000-0012122332320323-1230011003123022-1100233132201231-3130021033020232-0210300300303230-2312311312013102-3313131020020323"></a>

## Direct properties — trusted_ca / 133102300111 / 3

- [trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-0000201332213311-3121221102222010-0103203222233220-0100211322230202-1323101011130130-3332133011313213-2311332022302303-1030013212013210): complete subsection reference.

<a id="canonical-1231200132200311-3232330211001223-1321032131003100-0311230302102022-2113103011200103-0313230100110201-3332000122303001-0323223323232000"></a>

## Next pages — trusted_ca / 133102300111 / 4

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-0000201332213311-3121221102222010-0103203222233220-0100211322230202-1323101011130130-3332133011313213-2311332022302303-1030013212013210)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-0121031320112211-1123210202002222-2222232302211030-2030211033120231-2211100103310122-1110011313303011-1233131133202213-2132123300322200)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0000201332213311-3121221102222010-0103203222233220-0100211322230202-1323101011130130-3332133011313213-2311332022302303-1030013212013210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100121111230120-0113312331113102-3033010313001313-0322133203303331-1112301131030000-3330012003331302-2232120223003233-0222231331302122"></a>

## access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 110103301210 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--reference--group-001.md#canonical-0121031320112211-1123210202002222-2222232302211030-2030211033120231-2211100103310122-1110011313303011-1233131133202213-2132123300322200)
- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-2031030213320222-0001001333133110-0130332232123201-0022203002221000-2200312220313122-0103013012013030-1313302100301021-0233110310001321)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list

<a id="canonical-1211003001223000-2332222332101101-2312000331200022-0300132110013233-3032313113113321-1203022122131220-3030321300322131-0231220201130322"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011033201133003-2302323103100213-2200233110000121-2223131121023211-2301002133320133-1103003303311012-0110313211012313-0321103322022122"></a>

## Direct properties — trusted_ca_list / 110103301210 / 3

<a id="canonical-0311330001013003-3312201121223002-0031303211020322-2320223310020022-1013311323320200-0030132020223000-3030121000330233-3332012330032130"></a>

<a id="canonical-3112201020020201-3221202002210211-2323220111212201-0121310330201300-3303232131120200-2202200321030233-0110330202202023-0101313001233331"></a>

## kind property — trusted_ca_list / 110103301210 / 4

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

<a id="canonical-2103233113010130-1200000132002220-2203111101302113-2312110020211032-1323012221030130-0002130123021311-0301302321301232-3212010113131301"></a>

<a id="canonical-0020333110103233-3202123110113232-2202021231300300-0202101201300201-3002001231230321-0322002021320222-3021023220123213-1333303221212003"></a>

## name property — trusted_ca_list / 110103301210 / 5

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

<a id="canonical-3110333302311213-0123032131330112-3232212011111012-0120212320202033-0002013133031132-0303312203013133-1330211312001300-3011011013122320"></a>

<a id="canonical-0103301021122211-0203031130023112-0113133131230020-3012333021232323-0010002111013020-0111332103323131-3120221130001220-1223123210101023"></a>

## namespace property — trusted_ca_list / 110103301210 / 6

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

<a id="canonical-2231122322333222-0223301000201013-3103012021011013-0333222022100123-3003133123222232-0010312302220210-1220303031130220-1013021210231321"></a>

<a id="canonical-2022100012212232-0333202200232320-0300100212232003-2013202120023101-2233212211001002-3001101322023202-0312222210221331-1320313011011322"></a>

## tenant property — trusted_ca_list / 110103301210 / 7

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

<a id="canonical-1202002103200000-0103001130130322-2323333033302132-1120303210322303-2113010212233331-0133111300322100-1033323131032211-3013030033111322"></a>

<a id="canonical-2200131123322011-3211330023310001-0232002302322002-2333331002320120-3110002322230231-3131201312110321-3030101202320222-0111320111300033"></a>

## uid property — trusted_ca_list / 110103301210 / 8

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

<a id="canonical-0112023112121001-3333130331202200-3011330210203021-1203322303120112-2221230300013310-1312120011210122-0200120233313013-0120133211323022"></a>

## Next pages — trusted_ca_list / 110103301210 / 9

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-2031030213320222-0001001333133110-0130332232123201-0022203002221000-2200312220313122-0103013012013030-1313302100301021-0233110310001321)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-1001231221331313-1123332311100232-3211331021222023-3301333132002033-3030012122213310-0002033321033003-1032303110210330-1311012313121221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003103001013323-3302113223330130-3022323103222332-0333132131021211-1010333110123203-2331113333123312-1000312133213020-2221003021211320"></a>

## access_info.tls_config.cert_params.volterra_trusted_ca — volterra_trusted_ca / 222200333201 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- access_info.tls_config.cert_params.volterra_trusted_ca

<a id="canonical-0011223123132122-3120310233031211-0312111210103103-1023001011020323-0132211110311302-0312000122330332-2323232231000002-1310000030223123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
volterra_trusted_ca = {}
```

<a id="canonical-1310103022212233-3333223221021312-1212303301223000-0212312322103222-1102021210112032-0210132033122103-1232311133101310-2103213322322200"></a>

## Direct properties — volterra_trusted_ca / 222200333201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033220301312213-0221332123102220-2103021000310332-1212113002011032-3022300221211331-1312323032111212-2002231213210300-0032302213003220"></a>

## Next pages — volterra_trusted_ca / 222200333201 / 4

- [access_info.tls_config.cert_params](resources--secret_management_access--reference--group-001.md#canonical-3001022000002111-2201112233212000-1212201103100233-0102310102221331-3301113302111301-1331121113113033-3311230113102221-0211120333012220)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231210020033023-0310002232002311-1032201320023001-3202023102203032-2110001222223200-1222231021300113-2101320002230200-3020111112112100"></a>

## access_info.tls_config.common_params — common_params / 201111010102 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- access_info.tls_config.common_params

<a id="canonical-1021313131330113-2131223011003120-3311122112133333-1230022211301213-2211233110202331-0011202223120231-1330331000013300-0311020221031031"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
common_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300130221112123-2003203201111100-2211103233221330-1020313323200022-1021311223000312-1310100320310310-2033330202221332-0030223312012322"></a>

## Direct properties — common_params / 201111010102 / 3

<a id="canonical-0022003202011223-1310320331322202-2231303220310003-2131230303022320-2322310320122231-3102321131223002-2230101030020312-2331313000020101"></a>

<a id="canonical-3133232122322302-1210311202220112-1311010201213001-1113003323033002-3202231112222001-2230022302231132-2220102111220332-3022020023031101"></a>

## cipher_suites property — common_params / 201111010102 / 4

Type: `["list", "string"]`. Optional.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1110111000102103-3301322023023003-2231303021122101-3320321122232332-2203311203110221-1112311332032310-3222321230203320-3131103021323030"></a>

<a id="canonical-0311033000230301-0203003203221033-1033203123212323-2210212322201330-2122331020121313-3201022223213111-3202201131130111-0213103333231230"></a>

## maximum_protocol_version property — common_params / 201111010102 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-0220110013103230-3122132112231303-0123301021333323-2133321231201331-2020032030023122-0311310221322223-1113301221112130-0002202312032120"></a>

<a id="canonical-0121122110120020-3123123121002200-2210003211301111-2032201131031131-1100211203221310-2320302202331130-2322200210012123-2201012001032121"></a>

## minimum_protocol_version property — common_params / 201111010102 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

- [tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310): complete subsection reference.

- [validation_params](resources--secret_management_access--reference--group-001.md#canonical-2100022310310102-3100211033132010-1321133210103322-2003231000020013-3103032023120112-1312223311122122-3312332033122001-2113122001130333): complete subsection reference.

<a id="canonical-2310300111030000-3021022010213223-2201002102311101-3120020003011301-2012213101332133-2231123202000123-1102200223301032-1231120123333122"></a>

## Next pages — common_params / 201111010102 / 7

- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-2100022310310102-3100211033132010-1321133210103322-2003231000020013-3103032023120112-1312223311122122-3312332033122001-2113122001130333)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122311333101220-3333222312003232-2332013210202303-1211201023111112-3022031332321310-1311333130001013-2211212323100000-1113101103010232"></a>

## access_info.tls_config.common_params.tls_certificates — tls_certificates / 100012330120 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- access_info.tls_config.common_params.tls_certificates

<a id="canonical-0000322003121202-2123302100002330-1201101301123212-1002103003213232-2321122223001200-1100311300323032-3203203300132001-2030223220312033"></a>

Type: `"object"`. list nested block, Optional.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030133332311310-1201010202102102-0102121231021101-2331132130223103-3311100133102021-0332031223300301-3312110120102300-3200121120332232"></a>

## Direct properties — tls_certificates / 100012330120 / 3

<a id="canonical-3013012321023020-3301020232303303-0302002220002003-0130202322021121-3221332310121033-1103302330023200-3220311121021210-3021101213112122"></a>

<a id="canonical-2120123330302202-3300132303210232-0303003322103221-0303111211330231-0031003311033310-1333011012113310-0111022201222222-2112032201312232"></a>

## certificate_url property — tls_certificates / 100012330120 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](resources--secret_management_access--reference--group-001.md#canonical-3301221000322222-2131032022032001-1130313102121210-0212320000200311-3212130010211212-0201222321110103-3032233202010110-0202312311302310): complete subsection reference.

<a id="canonical-2002030301030311-0210033133010012-2121030232300100-3302313201113110-0313132013022212-2213312122330202-1102011322130033-1230022230311113"></a>

<a id="canonical-2200310001111133-2113201122012320-2131103323100000-2311132212133320-3113302003021313-2332010102011223-2003200223222231-2222021200221032"></a>

## description_spec property — tls_certificates / 100012330120 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--secret_management_access--reference--group-001.md#canonical-2211332301212202-1301301300121313-2020332002112321-1322131313332303-1330313211101031-2003001133030200-3020023130320212-2021131203311132): complete subsection reference.

- [private_key](resources--secret_management_access--reference--group-001.md#canonical-0110131301331210-2120332120000100-1103102001211330-2021221032121020-3023120212030003-3220133113122221-0223131211030221-2032123211103013): complete subsection reference.

- [use_system_defaults](resources--secret_management_access--reference--group-001.md#canonical-2332223321023312-3102202320201102-0332313223330202-2100123323203012-3203223010331211-3010330231101330-0131122113322112-1131300013123231): complete subsection reference.

<a id="canonical-1213120010020123-0321220020202123-1220230023231300-0301313112212112-3212321222122033-2312023102130032-3210032302113120-0100231103202321"></a>

## Next pages — tls_certificates / 100012330120 / 6

- [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](resources--secret_management_access--reference--group-001.md#canonical-3301221000322222-2131032022032001-1130313102121210-0212320000200311-3212130010211212-0201222321110103-3032233202010110-0202312311302310)
- [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](resources--secret_management_access--reference--group-001.md#canonical-2211332301212202-1301301300121313-2020332002112321-1322131313332303-1330313211101031-2003001133030200-3020023130320212-2021131203311132)
- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-0110131301331210-2120332120000100-1103102001211330-2021221032121020-3023120212030003-3220133113122221-0223131211030221-2032123211103013)
- [access_info.tls_config.common_params.tls_certificates.use_system_defaults](resources--secret_management_access--reference--group-001.md#canonical-2332223321023312-3102202320201102-0332313223330202-2100123323203012-3203223010331211-3010330231101330-0131122113322112-1131300013123231)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-3301221000322222-2131032022032001-1130313102121210-0212320000200311-3212130010211212-0201222321110103-3032233202010110-0202312311302310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021223013230333-2211032003231133-3212320013312010-3133032302131001-3221333321101110-0230012010312311-0102223230310233-1230101102320033"></a>

## access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 130210012122 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-1013103223230130-2003133323232103-3103322030200311-1233000001230231-1100313001001103-2102222220221332-3230131010200022-2332321310331212"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112310320013232-0301002202330233-0010310223020113-2213022321310020-3102203300123002-3122333020001223-0212030310012133-0220222022122113"></a>

## Direct properties — custom_hash_algorithms / 130210012122 / 3

<a id="canonical-1002320223323132-2222310220001303-0301222311330203-1310203111000022-1321232312132203-2003231202012133-0033122032021323-0131300002233000"></a>

<a id="canonical-3203130220030100-2223322212203131-0023033321002322-3212311113001231-2213312013003303-2122201100322033-3221122211203003-0230022110220300"></a>

## hash_algorithms property — custom_hash_algorithms / 130210012122 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-2021222223000311-2122200301023311-0002013133022100-3303130001131021-2110100121133103-0303200333112332-0030202200331323-0020222221330203"></a>

## Next pages — custom_hash_algorithms / 130210012122 / 5

- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-2211332301212202-1301301300121313-2020332002112321-1322131313332303-1330313211101031-2003001133030200-3020023130320212-2021131203311132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120220221231322-2100022133332213-3000221103123120-3222131031103120-1010001312202110-3333031002101200-3101303322320121-0000232322120223"></a>

## access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 211030122001 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-0122113310230102-1102001031033302-3131021010002132-3333202120231110-2212220100200210-2203110022033001-2010112130333101-0211320103200302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

<a id="canonical-1013303322121113-0023210112323011-1233101333012111-2030210201310201-1211132133302311-0321010231103223-0013330032132303-0203113201300130"></a>

## Direct properties — disable_ocsp_stapling / 211030122001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130311030233101-3303020310310231-1021011202201321-1111011030223310-0030022322201233-0212113333333330-0320333003000001-1102232303012023"></a>

## Next pages — disable_ocsp_stapling / 211030122001 / 4

- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0110131301331210-2120332120000100-1103102001211330-2021221032121020-3023120212030003-3220133113122221-0223131211030221-2032123211103013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202132310123121-1110123230232300-1301331210010000-1021301330131302-3101301122132312-3131321303201003-1103213222300133-2303202223212322"></a>

## access_info.tls_config.common_params.tls_certificates.private_key — private_key / 232011232232 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- access_info.tls_config.common_params.tls_certificates.private_key

<a id="canonical-3303311022312211-3101300310203033-2221102101100303-2002222211211333-3210123121002223-2303030102032111-0023120010021132-2222202103221222"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122211001133321-3123110031221313-1002001123013032-0221332213132033-0222301221110120-2212233321112221-0011231313033001-2220231310301333"></a>

## Direct properties — private_key / 232011232232 / 3

- [blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-2201223011031202-1010320301323300-1000331023033023-0010002100233022-0330303001112313-0300233031021302-3230032311211320-1310110303211210): complete subsection reference.

- [clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-0210121300232321-1232330133221120-3112033031212012-0330022221100232-3222002131110210-0120012231011033-1111231111033012-3131112332221331): complete subsection reference.

<a id="canonical-2130330133021032-3321003002020130-0323112320332230-1003220233331032-1102332130100001-3210323130001102-1322211113300113-0111330202220131"></a>

## Next pages — private_key / 232011232232 / 4

- [access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info](resources--secret_management_access--reference--group-001.md#canonical-2201223011031202-1010320301323300-1000331023033023-0010002100233022-0330303001112313-0300233031021302-3230032311211320-1310110303211210)
- [access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info](resources--secret_management_access--reference--group-001.md#canonical-0210121300232321-1232330133221120-3112033031212012-0330022221100232-3222002131110210-0120012231011033-1111231111033012-3131112332221331)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-2201223011031202-1010320301323300-1000331023033023-0010002100233022-0330303001112313-0300233031021302-3230032311211320-1310110303211210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122000202223003-0023121121323202-2332001310223313-2022010032310110-0020103032003002-1113212212231110-0332103020233002-1121100300022331"></a>

## access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 110232031032 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-0110131301331210-2120332120000100-1103102001211330-2021221032121020-3023120212030003-3220133113122221-0223131211030221-2032123211103013)
- access_info.tls_config.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0132132023022330-2231300031210332-2331221230300113-1110020113332200-1323020222233130-3020322023231023-3022323000003203-1031222322330211"></a>

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

<a id="canonical-1020000223001322-0123202130333131-0133111012232212-0213123332220103-3213131322100311-0131232010103303-3333303010210000-2301033311031111"></a>

## Direct properties — blindfold_secret_info / 110232031032 / 3

<a id="canonical-0212213122231001-1203101303123220-2131110120223200-1310021023112330-3303321322021313-3302101222010230-3011321312002232-1011322212033230"></a>

<a id="canonical-2101110200020033-3323331221102101-3233220221130210-3102310113111020-0332110000212331-1202031003033233-0130101332122330-3103210110111212"></a>

## decryption_provider property — blindfold_secret_info / 110232031032 / 4

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

<a id="canonical-1312322220103301-2102211130110221-2033013311210232-3033020112020230-1121231030222011-3110022112022032-2221221210201020-1002232121200002"></a>

<a id="canonical-1333121233311131-1211200302331300-2301302201300122-1231002301322231-0211300131011123-3331203013133231-1100030300102330-3323230130213212"></a>

## location property — blindfold_secret_info / 110232031032 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-3203030113231233-3010221212133223-2331220002111122-3030013013020332-2031303100320132-1321103030030102-3033103302102001-3101112211332221"></a>

<a id="canonical-2002022312311211-3030103321201133-0101232330212003-2311210102333122-0221033332132010-1031112202101102-1132120332022311-2121002220133331"></a>

## store_provider property — blindfold_secret_info / 110232031032 / 6

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

<a id="canonical-2032131233113033-0021111310303012-1322013022223031-1230032221232322-0232220121223301-3002013213223000-3031303111130130-3232130010230132"></a>

## Next pages — blindfold_secret_info / 110232031032 / 7

- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-0110131301331210-2120332120000100-1103102001211330-2021221032121020-3023120212030003-3220133113122221-0223131211030221-2032123211103013)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0210121300232321-1232330133221120-3112033031212012-0330022221100232-3222002131110210-0120012231011033-1111231111033012-3131112332221331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111203023020201-3120011133110132-1002030130221212-0232022103012222-1110311031113000-2303312202231121-1101233222003110-1213222201312231"></a>

## access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info — clear_secret_info / 202233031100 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-0110131301331210-2120332120000100-1103102001211330-2021221032121020-3023120212030003-3220133113122221-0223131211030221-2032123211103013)
- access_info.tls_config.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0131233110021013-2221130303202121-2220132100023131-3301000033103221-1110213223132232-0321013003101013-1203320301220220-0033013200223000"></a>

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

<a id="canonical-0300220212203011-1200232022230122-0112123012102031-2100203002111130-3233230211310331-2311232210203310-1120232321030121-0303330012312122"></a>

## Direct properties — clear_secret_info / 202233031100 / 3

<a id="canonical-1130011333101112-0003230322322312-0230101130203310-3033223201200023-0001232211203023-3012121111213302-2003021233231023-0000313001131310"></a>

<a id="canonical-1210220001232222-1300202303002301-3333103330332021-0003322123012032-3201031121120003-1133332112313311-2013311230122203-1221210121203211"></a>

## provider_ref property — clear_secret_info / 202233031100 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1223332023303132-3310211122201321-1201112231023133-3222002110021121-0111113133133021-0033021210003212-1310103011330133-0013311100231302"></a>

<a id="canonical-1111313301112321-1202211132313031-2130133330303032-2300020012302201-3110300220132322-2123013110331200-1300320211232113-1103131132210003"></a>

## URL property — clear_secret_info / 202233031100 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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

<a id="canonical-3023212132231230-0101102210333210-2201013103100223-0210002223101131-3312030213310203-0202001013322302-0320223022232001-1003213033213132"></a>

## Next pages — clear_secret_info / 202233031100 / 6

- [access_info.tls_config.common_params.tls_certificates.private_key](resources--secret_management_access--reference--group-001.md#canonical-0110131301331210-2120332120000100-1103102001211330-2021221032121020-3023120212030003-3220133113122221-0223131211030221-2032123211103013)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-2332223321023312-3102202320201102-0332313223330202-2100123323203012-3203223010331211-3010330231101330-0131122113322112-1131300013123231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232210320012032-3101023130223003-0200110031331112-2230011302310201-3111033203010202-1302310312200331-3102213103123030-2302033120330120"></a>

## access_info.tls_config.common_params.tls_certificates.use_system_defaults — use_system_defaults / 133133322203 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- access_info.tls_config.common_params.tls_certificates.use_system_defaults

<a id="canonical-0232131210233233-2121102011223210-1311213333112330-2312220313212013-0101032232002332-3201232031321301-1012100131302223-0202222321123103"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

<a id="canonical-3322313132131122-0011211113101300-2223333222202023-1202221020310013-2033013010120112-2120111301202123-3032020330032032-3301332201020213"></a>

## Direct properties — use_system_defaults / 133133322203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321132211111120-3301323222032312-1010013303031012-0313031312311210-1102202133212000-2113011311331130-2212020101010213-3011230111031010"></a>

## Next pages — use_system_defaults / 133133322203 / 4

- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--reference--group-001.md#canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-2100022310310102-3100211033132010-1321133210103322-2003231000020013-3103032023120112-1312223311122122-3312332033122001-2113122001130333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023231113200202-2113122202302300-2023212320110201-2202221202212320-3332221222201020-3022020332033113-1300310011220313-2333303300123121"></a>

## access_info.tls_config.common_params.validation_params — validation_params / 123320010203 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- access_info.tls_config.common_params.validation_params

<a id="canonical-3021301002311333-2131313331300332-1212003101212312-0023022102101003-2003320312201231-2330000123321213-3013333011132312-1000030013222320"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131230311332203-0000102302302212-0311003233111233-0132321312321322-0013131312232033-0100320330322203-0212220312210031-1003101320103312"></a>

## Direct properties — validation_params / 123320010203 / 3

<a id="canonical-1210022112111110-3133113023200112-1011032022100233-0121303023320132-3120200213122032-0111300322032032-3133111220301110-3030223023313322"></a>

<a id="canonical-2011233001320332-1102120312333131-0122132202020111-1230301110301233-0010121023102032-0123312031002003-2200201320231022-0113030031112033"></a>

## skip_hostname_verification property — validation_params / 123320010203 / 4

Type: `"bool"`. Optional.

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

- [trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-2012030223111220-2011103221001102-2202201031113033-0122332101313003-3001030230231002-3023220101210130-1302111130311130-1122123311010131): complete subsection reference.

<a id="canonical-0223012111020333-1110011022133303-3320233310013300-1123021310230101-1231222021220322-2000110203233321-2031003331321031-3011232313030200"></a>

<a id="canonical-0302202202023023-0032232133021031-3123032303333303-0302320103110330-1112031323013110-0010012110330330-1310030013202121-0313100331021211"></a>

## trusted_ca_url property — validation_params / 123320010203 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-3102312103310212-1002212211020123-3020003102021223-3321023212111302-1333231312201133-0223103130120131-3230221203211020-3103312030001011"></a>

<a id="canonical-2010212132020203-2321332100310312-2130031101333232-2231031132113202-2233121102330031-2221020211232020-1223323020000011-3003131113023123"></a>

## verify_subject_alt_names property — validation_params / 123320010203 / 6

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1001023331303230-1032313303233023-1220010021321323-0211221021021211-1323321301320002-0013221001303333-0220010133310100-1210031211332001"></a>

## Next pages — validation_params / 123320010203 / 7

- [access_info.tls_config.common_params.validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-2012030223111220-2011103221001102-2202201031113033-0122332101313003-3001030230231002-3023220101210130-1302111130311130-1122123311010131)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-2012030223111220-2011103221001102-2202201031113033-0122332101313003-3001030230231002-3023220101210130-1302111130311130-1122123311010131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111310010231012-0110313301223112-1331230313132200-0003213301013233-3302130203200132-3302133312302010-0330303300121221-3010101011132021"></a>

## access_info.tls_config.common_params.validation_params.trusted_ca — trusted_ca / 232003103331 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-2100022310310102-3100211033132010-1321133210103322-2003231000020013-3103032023120112-1312223311122122-3312332033122001-2113122001130333)
- access_info.tls_config.common_params.validation_params.trusted_ca

<a id="canonical-0232320133311111-3231031311213223-0123120113120032-3301210312320011-3110320321131122-1101200102113322-1223330213323303-0331323200023312"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031110220132110-2222323311303023-1232130230130303-1221003311003022-3120332020332223-3310313331202032-2113301122100233-3233130333301003"></a>

## Direct properties — trusted_ca / 232003103331 / 3

- [trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-2202031322113221-3211313331130022-3130001311110031-2301322321312010-1121321113302001-3203102121220333-1102031122233132-2002112033013013): complete subsection reference.

<a id="canonical-0313213002100201-2230101221202333-3300322030301312-3031233220222121-2003110321111030-0320231313221300-0020233002100213-3220021201001032"></a>

## Next pages — trusted_ca / 232003103331 / 4

- [access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list](resources--secret_management_access--reference--group-001.md#canonical-2202031322113221-3211313331130022-3130001311110031-2301322321312010-1121321113302001-3203102121220333-1102031122233132-2002112033013013)
- [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-2100022310310102-3100211033132010-1321133210103322-2003231000020013-3103032023120112-1312223311122122-3312332033122001-2113122001130333)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-2202031322113221-3211313331130022-3130001311110031-2301322321312010-1121321113302001-3203102121220333-1102031122233132-2002112033013013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313220223221301-3010121010310110-3120121232120021-2121031031312321-0131000311121300-2210312000322202-1321110020222213-3212231112213122"></a>

## access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 300331203012 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [access_info.tls_config.common_params](resources--secret_management_access--reference--group-001.md#canonical-1232221100110203-0220321230020230-1210031211101202-2031231201030330-2300232010223312-0303013330022332-3201103220201031-2300322122300120)
- [access_info.tls_config.common_params.validation_params](resources--secret_management_access--reference--group-001.md#canonical-2100022310310102-3100211033132010-1321133210103322-2003231000020013-3103032023120112-1312223311122122-3312332033122001-2113122001130333)
- [access_info.tls_config.common_params.validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-2012030223111220-2011103221001102-2202201031113033-0122332101313003-3001030230231002-3023220101210130-1302111130311130-1122123311010131)
- access_info.tls_config.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-2131112333022231-1100222202103011-2301201332100303-1110101300112033-1322000302102221-2011210210130020-0332023223210103-0020211113303300"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311122210320220-2213101332100213-2111010320011320-2031001122130310-2021002322030222-1103222002132210-1023333130102113-0222101031233212"></a>

## Direct properties — trusted_ca_list / 300331203012 / 3

<a id="canonical-1003230333331313-3033311313212013-1221321010323213-3122031232321022-2311233300112303-0221202323332012-1010001331311111-1303012020232111"></a>

<a id="canonical-2302211321321323-0100331311013320-2103030100302223-2232210333031003-0211300021021231-3213222223330102-1002000032300132-1323221332301023"></a>

## kind property — trusted_ca_list / 300331203012 / 4

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

<a id="canonical-3200301221233030-2030312121322100-1100111032323212-1111230220132103-2020310030203023-3211010122221102-1301102012100311-3311102212232103"></a>

<a id="canonical-1121032232013002-2312321013023311-1201130132113211-2223332120010101-0012200021012001-0121223202020322-3313022011223331-0120220132301102"></a>

## name property — trusted_ca_list / 300331203012 / 5

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

<a id="canonical-0022202311220012-1021110112130031-0200122212132312-2103302313323102-0001122123330203-2000112230102330-1032210210321220-0100003301202131"></a>

<a id="canonical-0213310311231010-2032323112321330-3320012220231233-1212303003301232-3321230230210010-3123301011022321-0031210332223201-1103330103001122"></a>

## namespace property — trusted_ca_list / 300331203012 / 6

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

<a id="canonical-1123230003310003-3201201330111312-2323012010303032-0031231132032231-1221002232010023-2230230001211230-3020310100331112-0000331000000130"></a>

<a id="canonical-0000211332212012-2311031333330121-3320321311011100-3322103312222103-1211231021233302-0013202113331113-2231231123223231-0032130110233033"></a>

## tenant property — trusted_ca_list / 300331203012 / 7

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

<a id="canonical-3212230213112132-3013130123310213-1103212301302120-2103122013123003-0110122231012103-2012023212131123-1132322330213031-2302112033213011"></a>

<a id="canonical-1111033320313210-3031322103001001-2321011000030101-0101110311210300-2000332311211112-2130223310011211-2020132120101012-0101333300001023"></a>

## uid property — trusted_ca_list / 300331203012 / 8

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

<a id="canonical-1031020001331130-0022322231132013-1112000211301102-0313213011000100-1301213300320100-0220211112321111-1012313122322212-0322202211300120"></a>

## Next pages — trusted_ca_list / 300331203012 / 9

- [access_info.tls_config.common_params.validation_params.trusted_ca](resources--secret_management_access--reference--group-001.md#canonical-2012030223111220-2011103221001102-2202201031113033-0122332101313003-3001030230231002-3023220101210130-1302111130311130-1122123311010131)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-1010002310111333-3331111312010000-2130000001223131-0121101203110301-1123333102202232-1230231311001011-2331030210020320-2212211003301212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120213133001013-3022230103233301-3001232322221233-3213210111232310-2310112330033211-2113110213012320-2033222003020303-0121321110311210"></a>

## access_info.tls_config.default_session_key_caching — default_session_key_caching / 103301233330 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- access_info.tls_config.default_session_key_caching

<a id="canonical-0110010211332033-0022012330313010-0103102130223003-2021213000333002-2102000110112212-2321221233320220-1232110001201230-0312030010130002"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_session_key_caching = {}
```

<a id="canonical-0201122332230100-0113113302123231-3310212023001021-3301122103102012-0020323120020221-3120201210303233-2223021311030212-0301031011330200"></a>

## Direct properties — default_session_key_caching / 103301233330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132133301310312-1201210113023323-2303102203301102-1212333102303323-3020222202121110-2210320230101033-3203220020312102-1200002221031101"></a>

## Next pages — default_session_key_caching / 103301233330 / 4

- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-2203002320210003-3112320211113020-1231121133010311-1233120131013313-3001311220012032-3222010112020103-2122300120312001-3113010110113312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233322331233110-3222221030123132-0322202112330022-2113001203231233-1310231303210010-2120112221331021-0312220321113102-0100100011131010"></a>

## access_info.tls_config.disable_session_key_caching — disable_session_key_caching / 101110023112 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [access_info](resources--secret_management_access--reference--group-001.md#canonical-2302030002201031-3031010213110200-2013100112131001-2333012320201231-1232103301212112-3331320111021012-0031323320220221-1211023031200100)
- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- access_info.tls_config.disable_session_key_caching

<a id="canonical-1023101211232012-2133221233202220-0333130310310212-3123312103203210-0022000021301203-3011212223213130-1112133101021020-3133112030202123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_session_key_caching = {}
```

<a id="canonical-2200323302111030-2213000211222311-1001212010200300-3013030100100222-3220300102103211-0131101210030100-1321131003211013-0331023001002302"></a>

## Direct properties — disable_session_key_caching / 101110023112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131223103123002-1233010211133002-1011300230231233-1331132100131233-1311303030212201-0320303021333113-3221331102230103-1323331121330211"></a>

## Next pages — disable_session_key_caching / 101110023112 / 4

- [access_info.tls_config](resources--secret_management_access--reference--group-001.md#canonical-0303313332033223-1101113133210311-0132323103031210-1231020232003323-1331213201300200-1321212120203132-2300130113302022-0011232001322121)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)

<a id="canonical-0332120001232113-3013023320312311-2102123303203030-3020002001020103-3033121211320030-0303331223210211-1231303302332132-2220202112120212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
