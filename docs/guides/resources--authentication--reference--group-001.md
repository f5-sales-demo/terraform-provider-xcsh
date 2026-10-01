---
page_title: "xcsh_authentication reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication reference."
---

# xcsh_authentication reference

<a id="canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133000131000111-2132122000333231-0232300333121012-1000130302222323-0221031213013003-0120313232211330-1300323332120122-3112030000333012"></a>

## Property reference — Property reference / 322303232323 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- Property reference

<a id="canonical-0233213111232311-1001100022320123-0010301002221113-1321231303003203-3332320133311132-1232333300123303-2013020110211121-1001200201121010"></a>

## Direct properties — Property reference / 322303232323 / 3

<a id="canonical-3213330020013132-2303230133302033-1210211111213322-2112031233300131-0222012300033311-0330301132020323-0122322003010223-1000310301121002"></a>

<a id="canonical-2301103233330233-0012332032202312-2230021021112210-3331032032010000-1231223122021000-3033000320030203-2120331111011030-0112202010121303"></a>

## annotations property — Property reference / 322303232323 / 4

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

- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110): complete subsection reference.

<a id="canonical-1332030110022312-1010221013220200-1010012303311302-3003121130201020-1121023203230302-0322010133032123-2112112201231123-1132221331001100"></a>

<a id="canonical-0203223100320032-2302320030033102-2123000013232211-1323103030013300-1111213110200000-2032232320003222-1012011312013201-1100331030033222"></a>

## description property — Property reference / 322303232323 / 5

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

<a id="canonical-3020023133011330-3120202321210310-2210222302223112-1212131111313123-2310201322001230-2202110010313020-2200210122002011-1121312323000002"></a>

<a id="canonical-2231100210033322-3030100211133332-3322011123022103-1203132200133111-1220310013031330-2000011032233022-2011000030100011-1323331002101323"></a>

## disable property — Property reference / 322303232323 / 6

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

<a id="canonical-2020332131012101-2010133300321131-2321303100032130-2302231130303103-1331222323012132-2131113023220100-3220222212330012-3313202112210301"></a>

<a id="canonical-1303320232002232-2223112322131011-3323311121101133-2001321310312033-1210003122313001-2333200131013223-3023212013333121-3212011220121102"></a>

## ID property — Property reference / 322303232323 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0222300012132200-1023033320233301-1231110333333001-3323332312012303-2022211312100322-3301033321121222-0121231311000300-0000121231213120"></a>

<a id="canonical-1003021233220022-2301312023203300-2231022003213120-0002032120103031-3130310321102021-3320123102201223-2000210212323300-0000331120231000"></a>

## labels property — Property reference / 322303232323 / 8

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

<a id="canonical-1023002012211102-1102131033202302-2220112312000220-2302122112331230-3321030202303230-0100310012211030-0202311222021201-0322223121210331"></a>

<a id="canonical-3310232312111333-2233022102130223-3202332202330020-3310001211021123-2011321120332200-2323221220302203-0231203122313221-3233330311131301"></a>

## name property — Property reference / 322303232323 / 9

Type: `"string"`. Required.

Name of the Authentication. Must be unique within the namespace.

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

<a id="canonical-0112311213123212-2311301332313120-2121300133013332-3123033320313102-0102210012000011-2101013230020212-2310003023232230-2131310112303322"></a>

<a id="canonical-2323020333231323-1111012030303320-0033221103112131-1020202001223312-2210120200020100-3302210130213223-1122330233303103-2122300333333033"></a>

## namespace property — Property reference / 322303232323 / 10

Type: `"string"`. Required.

Namespace where the Authentication is created.

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

- [oidc_auth](resources--authentication--reference--group-001.md#canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222): complete subsection reference.

- [timeouts](resources--authentication--reference--group-001.md#canonical-3303312123102202-1020231312022121-0002100023333100-3230010222333203-3101213130130031-3312230101321112-2323131100121313-1211003123300030): complete subsection reference.

<a id="canonical-3223021331012213-1231132130121122-0222010012322310-1021030023010302-1210022320113311-0023103002232322-2303111233210333-1031220033222121"></a>

## All schema paths — Property reference / 322303232323 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--authentication--reference--group-001.md#canonical-3213330020013132-2303230133302033-1210211111213322-2112031233300131-0222012300033311-0330301132020323-0122322003010223-1000310301121002) |
| `cookie_params` | [cookie_params](resources--authentication--reference--group-001.md#canonical-1101213232331223-0100112132333322-1022303222110222-1022211310230230-2000222311010232-2132122020312120-3111300020203023-2113300031101100) |
| `cookie_params.auth_hmac` | [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-3120122211310231-0302232102020030-2221012223012100-0321123300001231-2132230332010232-0023111200220131-1222131013232200-3320000323021311) |
| `cookie_params.auth_hmac.prim_key` | [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-0232333313131021-1132010222322010-0332021101220010-3203012212131122-3002110020110330-0322201132013300-3133120322030333-1210130012230110) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-3010211001113213-3222120201200323-0331032311210131-1223132330302202-1320310301003230-1132020223300000-1211313321233202-3210332231313120) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](resources--authentication--reference--group-001.md#canonical-3330032102320101-1222011020311202-2330231201202222-0102132310032320-2313112011013321-3220112222332013-1021110133120001-0113231112023033) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](resources--authentication--reference--group-001.md#canonical-1223132331203130-2113001101133302-2323031312303321-2013302302102121-0331002133333310-1300130133003333-1010102201121123-1102230331112101) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](resources--authentication--reference--group-001.md#canonical-3313013131022332-1312021103320033-0010022021233131-2232130330220003-1122120221300323-3011031000221103-0010333033330221-0002112213010100) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info` | [cookie_params.auth_hmac.prim_key.clear_secret_info](resources--authentication--reference--group-001.md#canonical-2200202303020111-3032330013203121-2203030202022122-2123130330031223-1333212122020210-0302133211122032-2322230300201023-1210121210121320) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](resources--authentication--reference--group-001.md#canonical-1231313112201122-0331001202322011-0021300003233003-0120332201200233-3230223312333101-3301001213311020-3033303032203011-0301313301300010) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [cookie_params.auth_hmac.prim_key.clear_secret_info.url](resources--authentication--reference--group-001.md#canonical-3313131101113231-2313223211110200-0203121333023130-0210333221101332-1101001222032111-0223013112120130-1312031010100230-1022221102102331) |
| `cookie_params.auth_hmac.prim_key_expiry` | [cookie_params.auth_hmac.prim_key_expiry](resources--authentication--reference--group-001.md#canonical-2203202122211202-0021301113102211-0333310110033002-2211211121230200-0231302330333100-3113211330121301-0201231202123312-2101201310100213) |
| `cookie_params.auth_hmac.sec_key` | [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-3003002300003200-3131132232001220-0300033211203000-1210123221332121-1121022112201213-0101213013300211-2211231121232011-3200013230311211) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-0311230322203213-0323212233112301-1132302300013101-3311101303213213-1202002131202032-3020313131033122-0131212332021031-1103201313111223) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](resources--authentication--reference--group-001.md#canonical-3330012013012330-3210330123320312-3003113322120121-3322010233020021-3111013321010000-2111033131121321-1033101001220123-0231031332333323) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](resources--authentication--reference--group-001.md#canonical-2220031132310110-2232001010032011-3223020330023221-3023301010221232-3130133320231131-2131213303303003-3023022010322001-3122302003033033) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](resources--authentication--reference--group-001.md#canonical-1232020101033202-2332130301221221-2113023021011213-2211223323200231-2232200011200103-1031003030011022-0301313033110001-0102311032031201) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info` | [cookie_params.auth_hmac.sec_key.clear_secret_info](resources--authentication--reference--group-001.md#canonical-3133320002221001-1301022000013022-2112022320203033-3331123032211030-1211332231333030-1303102231233320-3032110010222322-2313220130332230) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](resources--authentication--reference--group-001.md#canonical-2320222312102233-0112231312121223-2132213100213111-0003301300212321-2221303310132203-1230233032010133-2121312300100310-0113320133111123) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [cookie_params.auth_hmac.sec_key.clear_secret_info.url](resources--authentication--reference--group-001.md#canonical-2130331000331000-2011312232011132-3110132222002311-1130100312232003-0223120331322220-1003031113002320-1311233030001222-0200001211132112) |
| `cookie_params.auth_hmac.sec_key_expiry` | [cookie_params.auth_hmac.sec_key_expiry](resources--authentication--reference--group-001.md#canonical-1001102121012200-1102230003000031-2202210301012322-2301020212203212-1232122210210030-0200211200211120-2230300120301321-3112022303330132) |
| `cookie_params.cookie_expiry` | [cookie_params.cookie_expiry](resources--authentication--reference--group-001.md#canonical-1313321212312323-0101303332213330-3032333010310313-0032231303101320-3121021333032123-1131112012200130-2022301321330223-2320220212222332) |
| `cookie_params.cookie_refresh_interval` | [cookie_params.cookie_refresh_interval](resources--authentication--reference--group-001.md#canonical-1201331231123200-1033303211131131-2301312331323020-3310103122223231-0011110203002010-1002322101323321-1302230012120022-3031010223133213) |
| `cookie_params.kms_key_hmac` | [cookie_params.kms_key_hmac](resources--authentication--reference--group-001.md#canonical-3000331311232232-3120331032102320-3010203102030100-2313113231211322-0101223303221221-2021131203333231-3022103102230303-1321122102231132) |
| `cookie_params.session_expiry` | [cookie_params.session_expiry](resources--authentication--reference--group-001.md#canonical-3310120231022031-1022330010202302-1323303300030000-2112112112003311-0333331133200322-1310203210031232-2120132020321130-2112230102103013) |
| `description` | [description](resources--authentication--reference--group-001.md#canonical-1332030110022312-1010221013220200-1010012303311302-3003121130201020-1121023203230302-0322010133032123-2112112201231123-1132221331001100) |
| `disable` | [disable](resources--authentication--reference--group-001.md#canonical-3020023133011330-3120202321210310-2210222302223112-1212131111313123-2310201322001230-2202110010313020-2200210122002011-1121312323000002) |
| `id` | [ID](resources--authentication--reference--group-001.md#canonical-2020332131012101-2010133300321131-2321303100032130-2302231130303103-1331222323012132-2131113023220100-3220222212330012-3313202112210301) |
| `labels` | [labels](resources--authentication--reference--group-001.md#canonical-0222300012132200-1023033320233301-1231110333333001-3323332312012303-2022211312100322-3301033321121222-0121231311000300-0000121231213120) |
| `name` | [name](resources--authentication--reference--group-001.md#canonical-1023002012211102-1102131033202302-2220112312000220-2302122112331230-3321030202303230-0100310012211030-0202311222021201-0322223121210331) |
| `namespace` | [namespace](resources--authentication--reference--group-001.md#canonical-0112311213123212-2311301332313120-2121300133013332-3123033320313102-0102210012000011-2101013230020212-2310003023232230-2131310112303322) |
| `oidc_auth` | [oidc_auth](resources--authentication--reference--group-001.md#canonical-1020321322010212-2322011202212131-2312221012232220-2200100320320002-0302212110023330-3320233123301010-0310331131003021-2302032202201302) |
| `oidc_auth.client_secret` | [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-3301101301321023-2011110013333210-3213130320102120-0200323012033111-2001113221301231-2001321030003300-0103101003110222-1101320020002221) |
| `oidc_auth.client_secret.blindfold_secret_info` | [oidc_auth.client_secret.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-3112223030122003-3311022113123021-1012321033100201-1330210011222211-0322321113033222-0211301103101131-2001302130302310-2330123201113011) |
| `oidc_auth.client_secret.blindfold_secret_info.decryption_provider` | [oidc_auth.client_secret.blindfold_secret_info.decryption_provider](resources--authentication--reference--group-001.md#canonical-3032000211303020-3111222201301213-3110210311231003-0123202122030323-2011231233132313-1020123011023321-3332331332333113-3230312211301222) |
| `oidc_auth.client_secret.blindfold_secret_info.location` | [oidc_auth.client_secret.blindfold_secret_info.location](resources--authentication--reference--group-001.md#canonical-3103332001310332-2311032121130212-2021323020323211-2333322213012121-0021130101123113-1232313033221333-3201311202131033-2230222001203010) |
| `oidc_auth.client_secret.blindfold_secret_info.store_provider` | [oidc_auth.client_secret.blindfold_secret_info.store_provider](resources--authentication--reference--group-001.md#canonical-1333321210333211-1030113202321301-2233201122232031-0133331231101110-0131213131123312-2233312013301103-2012313321003222-1112333121211200) |
| `oidc_auth.client_secret.clear_secret_info` | [oidc_auth.client_secret.clear_secret_info](resources--authentication--reference--group-001.md#canonical-3203222021300223-0320333011131231-3210022223302113-2221022133313231-0303311210030313-3231323321003202-3200232121122330-1230203213310333) |
| `oidc_auth.client_secret.clear_secret_info.provider_ref` | [oidc_auth.client_secret.clear_secret_info.provider_ref](resources--authentication--reference--group-001.md#canonical-0220021221313313-0302300310000211-2312121023323002-1202023011203110-1212023121321100-3321101110312011-2301213022033113-3311110010232331) |
| `oidc_auth.client_secret.clear_secret_info.url` | [oidc_auth.client_secret.clear_secret_info.url](resources--authentication--reference--group-001.md#canonical-2022031232123012-1323012232322221-0123100230332100-0133203321120013-1322333132112112-3031132112110023-2312021003020010-1202222230110303) |
| `oidc_auth.oidc_auth_params` | [oidc_auth.oidc_auth_params](resources--authentication--reference--group-001.md#canonical-3303333101012201-3033132212301013-0123001103012333-3311330000131230-3121022231302303-0313302331131233-2013302031003000-1123313220310323) |
| `oidc_auth.oidc_auth_params.auth_endpoint_url` | [oidc_auth.oidc_auth_params.auth_endpoint_url](resources--authentication--reference--group-001.md#canonical-0212111010130212-3030131210001013-2312121202032330-1012121011211102-0020201212313312-1330333313013132-3102200200020013-1333210321132033) |
| `oidc_auth.oidc_auth_params.end_session_endpoint_url` | [oidc_auth.oidc_auth_params.end_session_endpoint_url](resources--authentication--reference--group-001.md#canonical-0322001030333122-1311201122313033-3213002003131330-3120013022202021-1101310231023332-2213000113322210-1323332233002130-2112211132202133) |
| `oidc_auth.oidc_auth_params.token_endpoint_url` | [oidc_auth.oidc_auth_params.token_endpoint_url](resources--authentication--reference--group-001.md#canonical-1101001113231211-2100212210022013-0220102132331232-0233100202212323-2232212233203031-3203330232020201-2212213210121020-0330332010102222) |
| `oidc_auth.oidc_client_id` | [oidc_auth.oidc_client_id](resources--authentication--reference--group-001.md#canonical-0031301232313200-0332332303330011-2113303112223233-0001203302101223-0123113231312001-2113130033230301-1122003032230033-2002132023332230) |
| `oidc_auth.oidc_well_known_config_url` | [oidc_auth.oidc_well_known_config_url](resources--authentication--reference--group-001.md#canonical-1302020102023020-3232002103113320-0001022022010332-3000111312013213-1031202230221101-1003302323001011-0321333002103303-3100231112321120) |
| `timeouts` | [timeouts](resources--authentication--reference--group-001.md#canonical-1312231330210232-0212102012112211-1223332203033223-3223332113131001-2123311112211002-0013321121321310-1020230313031121-1313333103211212) |
| `timeouts.create` | [timeouts.create](resources--authentication--reference--group-001.md#canonical-3321132333212111-0133322211021130-0112221211201122-2113033201213112-0112031230221000-2100030130322010-3310230312322012-2303133323222213) |
| `timeouts.delete` | [timeouts.delete](resources--authentication--reference--group-001.md#canonical-0012121131300302-3202002100121211-2020200213003013-2232323000011000-1013200333302210-0231131301011121-2013031333310123-2200032211220232) |
| `timeouts.read` | [timeouts.read](resources--authentication--reference--group-001.md#canonical-1233023122310303-0012130113300330-1233321100131222-0321030003111211-3032331203012133-3202011012011022-0020113222030320-2021302133033133) |
| `timeouts.update` | [timeouts.update](resources--authentication--reference--group-001.md#canonical-0320131111121222-3013313230101121-1122012213022203-0312310022332122-2312311022320220-3031233102221113-0110333332223131-0130000120322232) |

<a id="canonical-1123311321021201-1001130003121313-2232031330300103-3221323211101222-0230322303023121-0231212202312111-3001021300000201-3010223130232300"></a>

## Next pages — Property reference / 322303232323 / 12

- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222)
- [timeouts](resources--authentication--reference--group-001.md#canonical-3303312123102202-1020231312022121-0002100023333100-3230010222333203-3101213130130031-3312230101321112-2323131100121313-1211003123300030)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101123023010201-1302031102032211-2220031120203003-3211032031223223-1103131031330233-0031310210331312-1003210013132310-0202303212202002"></a>

## cookie_params — cookie_params / 033301332013 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- cookie_params

<a id="canonical-1101213232331223-0100112132333322-1022303222110222-1022211310230230-2000222311010232-2132122020312120-3111300020203023-2113300031101100"></a>

Type: `"object"`. single nested block, Optional.

Specifies different cookie related config parameters for authentication.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auth_hmac",
    "kms_key_hmac")}
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
  "x-ves-oneof-field-secret_choice": "[\"auth_hmac\",\"kms_key_hmac\"]"
}
```

Terraform syntax:

```terraform
cookie_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303221132033221-1303200211101200-2233232323303331-2023321121113331-0311331031131012-0133033003021203-3123202320300200-3230001310012302"></a>

## Direct properties — cookie_params / 033301332013 / 3

- [auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301): complete subsection reference.

<a id="canonical-1313321212312323-0101303332213330-3032333010310313-0032231303101320-3121021333032123-1131112012200130-2022301321330223-2320220212222332"></a>

<a id="canonical-3233011333231220-1231213311112202-3023023021313113-2113131022312113-2300123321321312-3132112333032022-0113222311310111-3210130321233302"></a>

## cookie_expiry property — cookie_params / 033301332013 / 4

Type: `"number"`. Optional.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client-side after which client will not
be setting the cookie as part of the request.

Upstream description:

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client-side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-1201331231123200-1033303211131131-2301312331323020-3310103122223231-0011110203002010-1002322101323321-1302230012120022-3031010223133213"></a>

<a id="canonical-1323203133102133-0230023021022202-1200023100130132-1201022231211022-1210110210333112-2301021313102203-3002330023301013-3102332332322011"></a>

## cookie_refresh_interval property — cookie_params / 033301332013 / 5

Type: `"number"`. Optional.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session..

Upstream description:

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

- [kms_key_hmac](resources--authentication--reference--group-001.md#canonical-2003212203022110-1010121230313123-3220100012131210-2330120021133202-3003222133010123-1212113021300220-2111221011111000-1132303310113010): complete subsection reference.

<a id="canonical-3310120231022031-1022330010202302-1323303300030000-2112112112003311-0333331133200322-1310203210031232-2120132020321130-2112230102103013"></a>

<a id="canonical-2201123111032223-0321020233321212-1011200202102033-2211322003021123-3030012321210321-1110312322300213-1203200021222122-0201133122021200"></a>

## session_expiry property — cookie_params / 033301332013 / 6

Type: `"number"`. Optional.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Upstream description:

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1296000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1296000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "1296000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1296000"
  }
}
```

<a id="canonical-2113110131220110-2102203320001010-1233332033123031-0323231202220222-2100110022032320-1103211133100032-2003232230131121-3032332213321013"></a>

## Next pages — cookie_params / 033301332013 / 7

- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301)
- [cookie_params.kms_key_hmac](resources--authentication--reference--group-001.md#canonical-2003212203022110-1010121230313123-3220100012131210-2330120021133202-3003222133010123-1212113021300220-2111221011111000-1132303310113010)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001033322231012-3311332200022300-3321131323233032-3101110112100331-1201223211322103-2003332112330031-3002323130110011-3103220323212223"></a>

## cookie_params.auth_hmac — auth_hmac / 220223102203 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- cookie_params.auth_hmac

<a id="canonical-3120122211310231-0302232102020030-2221012223012100-0321123300001231-2132230332010232-0023111200220131-1222131013232200-3320000323021311"></a>

Type: `"object"`. single nested block, Optional.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Upstream description:

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prim_key_expiry",
    "sec_key_expiry")}
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
auth_hmac {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112312300333231-2003033030020220-3100300333000133-0302333031102112-3322233121213130-1101033311222100-0010120130132220-1223203310121013"></a>

## Direct properties — auth_hmac / 220223102203 / 3

- [prim_key](resources--authentication--reference--group-001.md#canonical-3032121123310031-0232123133330200-3020103210122032-1020011331320323-3300310120322032-2321011320320102-0302123300213023-1221300110301223): complete subsection reference.

<a id="canonical-2203202122211202-0021301113102211-0333310110033002-2211211121230200-0231302330333100-3113211330121301-0201231202123312-2101201310100213"></a>

<a id="canonical-3122310011233031-2203020012122331-0010312000010021-2021333331012021-2123000033231323-3211322001331222-2011302010021131-0220021132000010"></a>

## prim_key_expiry property — auth_hmac / 220223102203 / 4

Type: `"string"`. Optional.

HMAC Primary Key Expiry. Primary HMAC Key Expiry time.

Upstream description:

Primary HMAC Key Expiry time.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [sec_key](resources--authentication--reference--group-001.md#canonical-0330232223232000-1131330103022012-0211033123123212-2310203020323113-2013220222333112-2110003310103313-3110100023010300-1032320232220033): complete subsection reference.

<a id="canonical-1001102121012200-1102230003000031-2202210301012322-2301020212203212-1232122210210030-0200211200211120-2230300120301321-3112022303330132"></a>

<a id="canonical-1313123122333322-3323233020003310-0111230101201301-0101213213233031-0223020221203113-1320231032233102-2322313111023302-0133312300011020"></a>

## sec_key_expiry property — auth_hmac / 220223102203 / 5

Type: `"string"`. Optional.

HMAC Secondary Key Expiry. Secondary HMAC Key Expiry time.

Upstream description:

Secondary HMAC Key Expiry time.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

<a id="canonical-0313113002032002-0203110303303100-3101032331122121-0332301310000300-1012033003310130-1231320320312132-0101300331312312-3111002123321313"></a>

## Next pages — auth_hmac / 220223102203 / 6

- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-3032121123310031-0232123133330200-3020103210122032-1020011331320323-3300310120322032-2321011320320102-0302123300213023-1221300110301223)
- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-0330232223232000-1131330103022012-0211033123123212-2310203020323113-2013220222333112-2110003310103313-3110100023010300-1032320232220033)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-3032121123310031-0232123133330200-3020103210122032-1020011331320323-3300310120322032-2321011320320102-0302123300213023-1221300110301223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012312321003211-2133322223110200-0321023132211022-1313010300100110-2000122221221131-1011023313102002-3002323001101001-3232110121100002"></a>

## cookie_params.auth_hmac.prim_key — prim_key / 312213013102 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301)
- cookie_params.auth_hmac.prim_key

<a id="canonical-0232333313131021-1132010222322010-0332021101220010-3203012212131122-3002110020110330-0322201132013300-3133120322030333-1210130012230110"></a>

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
prim_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122000222312320-2112102013312102-1223020121031120-1211013332311120-2333111213101220-0323122110002102-0302333120210312-2111302210100301"></a>

## Direct properties — prim_key / 312213013102 / 3

- [blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-2023331000303023-1213133321020300-0201323202322310-0200030213131311-1201321201102222-2310200133300000-1001013332312003-3212000330130200): complete subsection reference.

- [clear_secret_info](resources--authentication--reference--group-001.md#canonical-0110322113231023-1313013002322120-0200300300220201-1020001113013322-1011332330012122-0223120112212030-2330101233220320-3231211021331210): complete subsection reference.

<a id="canonical-1113222331130131-1120301102223133-3013110310303202-1320132101323003-3222331313013223-0202323032000301-1333220123001213-1223232222301020"></a>

## Next pages — prim_key / 312213013102 / 4

- [cookie_params.auth_hmac.prim_key.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-2023331000303023-1213133321020300-0201323202322310-0200030213131311-1201321201102222-2310200133300000-1001013332312003-3212000330130200)
- [cookie_params.auth_hmac.prim_key.clear_secret_info](resources--authentication--reference--group-001.md#canonical-0110322113231023-1313013002322120-0200300300220201-1020001113013322-1011332330012122-0223120112212030-2330101233220320-3231211021331210)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-2023331000303023-1213133321020300-0201323202322310-0200030213131311-1201321201102222-2310200133300000-1001013332312003-3212000330130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323032103322311-1320210332201131-0000023000300223-2221021033230002-3221302011322302-1231131212201321-1121032313002222-2013332031300223"></a>

## cookie_params.auth_hmac.prim_key.blindfold_secret_info — blindfold_secret_info / 300313101133 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301)
- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-3032121123310031-0232123133330200-3020103210122032-1020011331320323-3300310120322032-2321011320320102-0302123300213023-1221300110301223)
- cookie_params.auth_hmac.prim_key.blindfold_secret_info

<a id="canonical-3010211001113213-3222120201200323-0331032311210131-1223132330302202-1320310301003230-1132020223300000-1211313321233202-3210332231313120"></a>

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

<a id="canonical-3321123332132131-0033021330233323-2120132203231233-2200130220201310-0113000230000233-0233010001321122-1101331312113200-2201213123001020"></a>

## Direct properties — blindfold_secret_info / 300313101133 / 3

<a id="canonical-3330032102320101-1222011020311202-2330231201202222-0102132310032320-2313112011013321-3220112222332013-1021110133120001-0113231112023033"></a>

<a id="canonical-2102312212131202-1023321333110000-1101230021231301-2020121022112023-2310230211103013-3012000321232330-2203322131001321-1120021212001303"></a>

## decryption_provider property — blindfold_secret_info / 300313101133 / 4

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

<a id="canonical-1223132331203130-2113001101133302-2323031312303321-2013302302102121-0331002133333310-1300130133003333-1010102201121123-1102230331112101"></a>

<a id="canonical-2313320120002120-3033303320002300-0122011003101132-1020022203310122-3111323122130121-0221033311020232-3010322330313303-3012310310222131"></a>

## location property — blindfold_secret_info / 300313101133 / 5

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

<a id="canonical-3313013131022332-1312021103320033-0010022021233131-2232130330220003-1122120221300323-3011031000221103-0010333033330221-0002112213010100"></a>

<a id="canonical-0203103200331213-3202113202030320-0322213330023113-0121323013133202-0112212221130001-2120210011112131-3210301330130110-3221321223022002"></a>

## store_provider property — blindfold_secret_info / 300313101133 / 6

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

<a id="canonical-1333112000102120-1310033332320033-1111113331121333-3202231013033310-3323002022321311-0031030230030230-3033113111032001-2200213311321213"></a>

## Next pages — blindfold_secret_info / 300313101133 / 7

- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-3032121123310031-0232123133330200-3020103210122032-1020011331320323-3300310120322032-2321011320320102-0302123300213023-1221300110301223)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-0110322113231023-1313013002322120-0200300300220201-1020001113013322-1011332330012122-0223120112212030-2330101233220320-3231211021331210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113130021302213-3020032213100010-0212020203320001-3332322322001221-3120123121110030-1331100100003101-2130231120320232-0121023020103333"></a>

## cookie_params.auth_hmac.prim_key.clear_secret_info — clear_secret_info / 302332020333 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301)
- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-3032121123310031-0232123133330200-3020103210122032-1020011331320323-3300310120322032-2321011320320102-0302123300213023-1221300110301223)
- cookie_params.auth_hmac.prim_key.clear_secret_info

<a id="canonical-2200202303020111-3032330013203121-2203030202022122-2123130330031223-1333212122020210-0302133211122032-2322230300201023-1210121210121320"></a>

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

<a id="canonical-2030000232002202-3231332201310301-3122012202023200-3033222223122210-2330310013313223-3333103131223330-0310022203030103-3201333333031100"></a>

## Direct properties — clear_secret_info / 302332020333 / 3

<a id="canonical-1231313112201122-0331001202322011-0021300003233003-0120332201200233-3230223312333101-3301001213311020-3033303032203011-0301313301300010"></a>

<a id="canonical-0212303002320030-0100211130022323-2322033103200010-2330001020032222-1222101100222101-0113311213213202-0312202230122012-0132012213200323"></a>

## provider_ref property — clear_secret_info / 302332020333 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3313131101113231-2313223211110200-0203121333023130-0210333221101332-1101001222032111-0223013112120130-1312031010100230-1022221102102331"></a>

<a id="canonical-0110202121303010-1131201321231223-1303222003202222-1231210000232122-2021011313112102-3021300301112232-2313120130233131-0132200103303213"></a>

## URL property — clear_secret_info / 302332020333 / 5

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

<a id="canonical-3322333030332210-0122230021312333-2113303012030301-2203332320102003-2310303201201130-1103003122011120-1202121320131200-0303120111202310"></a>

## Next pages — clear_secret_info / 302332020333 / 6

- [cookie_params.auth_hmac.prim_key](resources--authentication--reference--group-001.md#canonical-3032121123310031-0232123133330200-3020103210122032-1020011331320323-3300310120322032-2321011320320102-0302123300213023-1221300110301223)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-0330232223232000-1131330103022012-0211033123123212-2310203020323113-2013220222333112-2110003310103313-3110100023010300-1032320232220033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000031320102130-0333220300330033-3103100211120201-1111002231131133-2012022002310010-0001232031023333-2301301021300310-2012312221120323"></a>

## cookie_params.auth_hmac.sec_key — sec_key / 301230203201 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301)
- cookie_params.auth_hmac.sec_key

<a id="canonical-3003002300003200-3131132232001220-0300033211203000-1210123221332121-1121022112201213-0101213013300211-2211231121232011-3200013230311211"></a>

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
sec_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333003332032310-2220100110231313-1121333220013321-0202033012013001-0113203111230121-3133011102220211-0303112300203023-3133300322231202"></a>

## Direct properties — sec_key / 301230203201 / 3

- [blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-2223131323101331-0323001131102102-0003203303301121-0313011012110331-0130210312132223-1312211233201210-2002332002122231-0322122303032021): complete subsection reference.

- [clear_secret_info](resources--authentication--reference--group-001.md#canonical-1230230223233213-1001301023002000-2332011130021320-3311233201232230-3310332210203120-3102021322203121-2313021001201123-2222113310101133): complete subsection reference.

<a id="canonical-1022312033203313-3332032323301323-3301232120131123-2000023023132313-0030001013120133-1331220230220111-2100200221211012-2220331222332123"></a>

## Next pages — sec_key / 301230203201 / 4

- [cookie_params.auth_hmac.sec_key.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-2223131323101331-0323001131102102-0003203303301121-0313011012110331-0130210312132223-1312211233201210-2002332002122231-0322122303032021)
- [cookie_params.auth_hmac.sec_key.clear_secret_info](resources--authentication--reference--group-001.md#canonical-1230230223233213-1001301023002000-2332011130021320-3311233201232230-3310332210203120-3102021322203121-2313021001201123-2222113310101133)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-2223131323101331-0323001131102102-0003203303301121-0313011012110331-0130210312132223-1312211233201210-2002332002122231-0322122303032021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133002002211102-0300031001203102-3220111303030220-1002102210211020-1301022012312312-1302100001222310-3000033321231011-2031221323122220"></a>

## cookie_params.auth_hmac.sec_key.blindfold_secret_info — blindfold_secret_info / 013101130013 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301)
- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-0330232223232000-1131330103022012-0211033123123212-2310203020323113-2013220222333112-2110003310103313-3110100023010300-1032320232220033)
- cookie_params.auth_hmac.sec_key.blindfold_secret_info

<a id="canonical-0311230322203213-0323212233112301-1132302300013101-3311101303213213-1202002131202032-3020313131033122-0131212332021031-1103201313111223"></a>

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

<a id="canonical-0320111311313112-1023212113012213-2203323130013232-0302103130101222-0331223133220230-0222310023133303-1202020013220301-2212031300102020"></a>

## Direct properties — blindfold_secret_info / 013101130013 / 3

<a id="canonical-3330012013012330-3210330123320312-3003113322120121-3322010233020021-3111013321010000-2111033131121321-1033101001220123-0231031332333323"></a>

<a id="canonical-0232110221032302-2122001213220131-2222032013011133-2102323230311211-3113302133232210-0311200202201131-3032130113321332-2031130310111220"></a>

## decryption_provider property — blindfold_secret_info / 013101130013 / 4

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

<a id="canonical-2220031132310110-2232001010032011-3223020330023221-3023301010221232-3130133320231131-2131213303303003-3023022010322001-3122302003033033"></a>

<a id="canonical-2020023110202032-3320000321120033-1211012202213033-3133302233133001-1032031322120101-3012300010002233-0320211302331223-0120133132213102"></a>

## location property — blindfold_secret_info / 013101130013 / 5

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

<a id="canonical-1232020101033202-2332130301221221-2113023021011213-2211223323200231-2232200011200103-1031003030011022-0301313033110001-0102311032031201"></a>

<a id="canonical-3312323321110231-3131200230030322-3211122232201212-1213032221023020-1112210131102120-3023311023202220-3130031221131232-3323020133223112"></a>

## store_provider property — blindfold_secret_info / 013101130013 / 6

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

<a id="canonical-2123333220320230-3011102100232213-3302000322332300-2323110201321123-0233312133311232-3223222110022332-3002001210323130-1332333020022010"></a>

## Next pages — blindfold_secret_info / 013101130013 / 7

- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-0330232223232000-1131330103022012-0211033123123212-2310203020323113-2013220222333112-2110003310103313-3110100023010300-1032320232220033)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-1230230223233213-1001301023002000-2332011130021320-3311233201232230-3310332210203120-3102021322203121-2313021001201123-2222113310101133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200000301312322-0013120012002311-3310331332301120-3120110002233003-3021010330323303-1200320302333221-2110013330003123-3103121102213302"></a>

## cookie_params.auth_hmac.sec_key.clear_secret_info — clear_secret_info / 210033020002 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- [cookie_params.auth_hmac](resources--authentication--reference--group-001.md#canonical-1030303313120200-2123312122320320-3323022100021303-2022132030133313-3130201121102120-2132213101202231-1230102221321112-2231231033210301)
- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-0330232223232000-1131330103022012-0211033123123212-2310203020323113-2013220222333112-2110003310103313-3110100023010300-1032320232220033)
- cookie_params.auth_hmac.sec_key.clear_secret_info

<a id="canonical-3133320002221001-1301022000013022-2112022320203033-3331123032211030-1211332231333030-1303102231233320-3032110010222322-2313220130332230"></a>

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

<a id="canonical-0013201020212101-0220012132330213-1210213002113203-0231110231202200-3322022303123110-1131122000013001-1103323301001311-2312233320230332"></a>

## Direct properties — clear_secret_info / 210033020002 / 3

<a id="canonical-2320222312102233-0112231312121223-2132213100213111-0003301300212321-2221303310132203-1230233032010133-2121312300100310-0113320133111123"></a>

<a id="canonical-2200330120112333-1130121000333212-3233310132123100-2223213301320023-0210112102102200-3003201231302100-0033310122013230-1220031233203010"></a>

## provider_ref property — clear_secret_info / 210033020002 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2130331000331000-2011312232011132-3110132222002311-1130100312232003-0223120331322220-1003031113002320-1311233030001222-0200001211132112"></a>

<a id="canonical-2133020313202222-1333303333212033-0113031113231103-0132030312323300-1321312122322331-1110131103230103-3123213133223012-3320332300302010"></a>

## URL property — clear_secret_info / 210033020002 / 5

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

<a id="canonical-1232102211113101-3222100313212311-0031233201021212-0323310311010030-2013100101312002-0100213220222231-1110322333113030-3112011010023011"></a>

## Next pages — clear_secret_info / 210033020002 / 6

- [cookie_params.auth_hmac.sec_key](resources--authentication--reference--group-001.md#canonical-0330232223232000-1131330103022012-0211033123123212-2310203020323113-2013220222333112-2110003310103313-3110100023010300-1032320232220033)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-2003212203022110-1010121230313123-3220100012131210-2330120021133202-3003222133010123-1212113021300220-2111221011111000-1132303310113010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103200001203201-2013022132011132-2323030031113212-3222132211133232-1300313310122231-0122001022110300-0031301013222130-1101021013333000"></a>

## cookie_params.kms_key_hmac — kms_key_hmac / 013002311233 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- cookie_params.kms_key_hmac

<a id="canonical-3000331311232232-3120331032102320-3010203102030100-2313113231211322-0101223303221221-2021131203333231-3022103102230303-1321122102231132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for kms key hmac.

Upstream description:

Reference to KMS Key Object.

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
kms_key_hmac = {}
```

<a id="canonical-3010310130010331-0223121201002223-3011321222112012-3121231012300113-0131322031333112-3333321011100022-0003302132223211-3230100003201301"></a>

## Direct properties — kms_key_hmac / 013002311233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330203102211322-1310022210201102-1112233033103200-1110313333311010-1000212110132033-2203101111221021-1221300213233202-0223133223313111"></a>

## Next pages — kms_key_hmac / 013002311233 / 4

- [cookie_params](resources--authentication--reference--group-001.md#canonical-0320202312203102-2010321021022301-0100331102230320-0330333111033032-1022101322132232-0213230210110001-0311123301230122-1030121000202110)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231020302303220-1312210330213020-0002202233311332-3020003332302112-3031312020232300-2202001223210021-0122122000221031-1312332213022211"></a>

## oidc_auth — oidc_auth / 202311000313 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- oidc_auth

<a id="canonical-1020321322010212-2322011202212131-2312221012232220-2200100320320002-0302212110023330-3320233123301010-0310331131003021-2302032202201302"></a>

Type: `"object"`. single nested block, Optional.

OIDCAuthType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("oidc_client_id"),
  validators.ConflictingObjectAttributes("oidc_auth_params",
    "oidc_well_known_config_url")}
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
  "x-ves-oneof-field-auth_params_choice": "[\"oidc_auth_params\",\"oidc_well_known_config_url\"]"
}
```

Terraform syntax:

```terraform
oidc_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102031333110231-3231322200311333-3202311320211310-1103321033103212-2113313223021010-1100302111333323-2020331113032222-0232333230231112"></a>

## Direct properties — oidc_auth / 202311000313 / 3

- [client_secret](resources--authentication--reference--group-001.md#canonical-1320230130103302-2003312321123101-3113322013121030-2001000113312133-1110211313130221-2210321120320211-2002022121312232-2333123113001213): complete subsection reference.

- [oidc_auth_params](resources--authentication--reference--group-001.md#canonical-2121211332103330-0132113033112313-0200321101212122-0022032210032230-2112313233102322-2101230323030132-2030321103332312-1323132000201232): complete subsection reference.

<a id="canonical-0031301232313200-0332332303330011-2113303112223233-0001203302101223-0123113231312001-2113130033230301-1122003032230033-2002132023332230"></a>

<a id="canonical-2010301013001133-0322111232131321-2032012010220031-2201030102332110-1222212312023113-0233011302300110-0003112020113003-0030200130322300"></a>

## oidc_client_id property — oidc_auth / 202311000313 / 4

Type: `"string"`. Optional.

Client ID used while sending the Authorization Request to OIDC server.

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

<a id="canonical-1302020102023020-3232002103113320-0001022022010332-3000111312013213-1031202230221101-1003302323001011-0321333002103303-3100231112321120"></a>

<a id="canonical-1030030133123033-3233200011220112-2121321112303022-2032011301133123-2312021303231001-3200221231301033-0021200202231210-1032221303310200"></a>

## oidc_well_known_config_url property — oidc_auth / 202311000313 / 5

Type: `"string"`. Optional.

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

Upstream description:

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0220213213003021-1232101110002122-3111011033222233-3131320100003303-2331033303330331-3232101130332221-1232230111213131-1320331220023033"></a>

## Next pages — oidc_auth / 202311000313 / 6

- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-1320230130103302-2003312321123101-3113322013121030-2001000113312133-1110211313130221-2210321120320211-2002022121312232-2333123113001213)
- [oidc_auth.oidc_auth_params](resources--authentication--reference--group-001.md#canonical-2121211332103330-0132113033112313-0200321101212122-0022032210032230-2112313233102322-2101230323030132-2030321103332312-1323132000201232)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-1320230130103302-2003312321123101-3113322013121030-2001000113312133-1110211313130221-2210321120320211-2002022121312232-2333123113001213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110103010021221-3312133131310223-0313022310213111-2111203113020103-2212332332100000-3333210213303122-3021020220320110-0201200233222133"></a>

## oidc_auth.client_secret — client_secret / 231202330320 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222)
- oidc_auth.client_secret

<a id="canonical-3301101301321023-2011110013333210-3213130320102120-0200323012033111-2001113221301231-2001321030003300-0103101003110222-1101320020002221"></a>

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
client_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010110211012201-0310021302122231-0112103030020220-3310311321330231-2310311331301032-3203311003321213-0213231213003001-0121130121120121"></a>

## Direct properties — client_secret / 231202330320 / 3

- [blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-3213223111302231-3212123332110113-3203323230132110-0020103210133203-1120131202230000-2111010203313023-2102112032111031-1112023331010301): complete subsection reference.

- [clear_secret_info](resources--authentication--reference--group-001.md#canonical-0311323102113013-3003023003021301-3212010223330222-2113102332202020-3221322333321120-0300020322031321-3203112311333130-0322011033311321): complete subsection reference.

<a id="canonical-0303032313002122-3200130033122222-1230230002211122-3201220213223201-1001222032000233-3231312133213323-0121220023001021-3002210020030011"></a>

## Next pages — client_secret / 231202330320 / 4

- [oidc_auth.client_secret.blindfold_secret_info](resources--authentication--reference--group-001.md#canonical-3213223111302231-3212123332110113-3203323230132110-0020103210133203-1120131202230000-2111010203313023-2102112032111031-1112023331010301)
- [oidc_auth.client_secret.clear_secret_info](resources--authentication--reference--group-001.md#canonical-0311323102113013-3003023003021301-3212010223330222-2113102332202020-3221322333321120-0300020322031321-3203112311333130-0322011033311321)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-3213223111302231-3212123332110113-3203323230132110-0020103210133203-1120131202230000-2111010203313023-2102112032111031-1112023331010301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133322113112020-2022313311201111-3300012220213320-0122130323232021-3000003120100011-0023012133130122-1233031312113011-1300230212301321"></a>

## oidc_auth.client_secret.blindfold_secret_info — blindfold_secret_info / 320021210223 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222)
- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-1320230130103302-2003312321123101-3113322013121030-2001000113312133-1110211313130221-2210321120320211-2002022121312232-2333123113001213)
- oidc_auth.client_secret.blindfold_secret_info

<a id="canonical-3112223030122003-3311022113123021-1012321033100201-1330210011222211-0322321113033222-0211301103101131-2001302130302310-2330123201113011"></a>

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

<a id="canonical-1333022230212121-2233100020223022-0011131201310101-3213110212202012-2121322212002313-0112303300302332-3103033131023102-2203002221313123"></a>

## Direct properties — blindfold_secret_info / 320021210223 / 3

<a id="canonical-3032000211303020-3111222201301213-3110210311231003-0123202122030323-2011231233132313-1020123011023321-3332331332333113-3230312211301222"></a>

<a id="canonical-0203320032220232-2013331213011333-0122331031101102-1303222022031330-2313232210002320-1221123202320001-3203032122303120-1030120211230333"></a>

## decryption_provider property — blindfold_secret_info / 320021210223 / 4

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

<a id="canonical-3103332001310332-2311032121130212-2021323020323211-2333322213012121-0021130101123113-1232313033221333-3201311202131033-2230222001203010"></a>

<a id="canonical-3132223330010002-2231121212002232-3302011233102121-0100010212002330-1012211321221232-0332331230122022-0023220330111202-1220333311300220"></a>

## location property — blindfold_secret_info / 320021210223 / 5

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

<a id="canonical-1333321210333211-1030113202321301-2233201122232031-0133331231101110-0131213131123312-2233312013301103-2012313321003222-1112333121211200"></a>

<a id="canonical-2110123101030121-1213102222211021-2312112131130132-2203130023030313-0202330213201220-3301212131301330-3112233102012103-3203210303111200"></a>

## store_provider property — blindfold_secret_info / 320021210223 / 6

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

<a id="canonical-1022300013220300-0032323333131102-0300220300220201-1333320132331231-2200133331121221-2300102230113202-2210201313011022-3002010021013120"></a>

## Next pages — blindfold_secret_info / 320021210223 / 7

- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-1320230130103302-2003312321123101-3113322013121030-2001000113312133-1110211313130221-2210321120320211-2002022121312232-2333123113001213)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-0311323102113013-3003023003021301-3212010223330222-2113102332202020-3221322333321120-0300020322031321-3203112311333130-0322011033311321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023201330303030-3100133131022310-2100012200001120-0131102223101132-2332121321311333-3210001010323310-0322111032312203-3022223100311312"></a>

## oidc_auth.client_secret.clear_secret_info — clear_secret_info / 121312312022 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222)
- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-1320230130103302-2003312321123101-3113322013121030-2001000113312133-1110211313130221-2210321120320211-2002022121312232-2333123113001213)
- oidc_auth.client_secret.clear_secret_info

<a id="canonical-3203222021300223-0320333011131231-3210022223302113-2221022133313231-0303311210030313-3231323321003202-3200232121122330-1230203213310333"></a>

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

<a id="canonical-0111332021010310-2332002111111211-0023033011212000-1220232030221021-3203010112213131-2232300010112231-3203111123332330-0110122102131330"></a>

## Direct properties — clear_secret_info / 121312312022 / 3

<a id="canonical-0220021221313313-0302300310000211-2312121023323002-1202023011203110-1212023121321100-3321101110312011-2301213022033113-3311110010232331"></a>

<a id="canonical-1111002131132020-2231330003023020-0223022212132120-3021210332312200-1133123131030201-0130233100013303-1132302111301222-1333211210213100"></a>

## provider_ref property — clear_secret_info / 121312312022 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2022031232123012-1323012232322221-0123100230332100-0133203321120013-1322333132112112-3031132112110023-2312021003020010-1202222230110303"></a>

<a id="canonical-1010131210110011-0222031211120313-2301031111113310-0220131112200033-3021001321231123-2223030232012111-3013323300021310-3132333120222313"></a>

## URL property — clear_secret_info / 121312312022 / 5

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

<a id="canonical-0021230223322103-2002033212130022-3331300232010300-3010322012203032-0002011203031131-3330131121131110-2302132102313030-2221320313123213"></a>

## Next pages — clear_secret_info / 121312312022 / 6

- [oidc_auth.client_secret](resources--authentication--reference--group-001.md#canonical-1320230130103302-2003312321123101-3113322013121030-2001000113312133-1110211313130221-2210321120320211-2002022121312232-2333123113001213)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-2121211332103330-0132113033112313-0200321101212122-0022032210032230-2112313233102322-2101230323030132-2030321103332312-1323132000201232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122021103223233-1000111211210301-1330100013333202-0001100301032023-3331231223012031-0301211202321232-2013101121202313-0323300333030031"></a>

## oidc_auth.oidc_auth_params — oidc_auth_params / 122331111010 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [oidc_auth](resources--authentication--reference--group-001.md#canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222)
- oidc_auth.oidc_auth_params

<a id="canonical-3303333101012201-3033132212301013-0123001103012333-3311330000131230-3121022231302303-0313302331131233-2013302031003000-1123313220310323"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for oidc auth params.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("auth_endpoint_url",
    "end_session_endpoint_url",
    "token_endpoint_url")}
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
oidc_auth_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132332212012201-3032330011301310-2230312102233003-2103323331331230-2201000103202111-3012301121103102-3201101133230010-0113011131121133"></a>

## Direct properties — oidc_auth_params / 122331111010 / 3

<a id="canonical-0212111010130212-3030131210001013-2312121202032330-1012121011211102-0020201212313312-1330333313013132-3102200200020013-1333210321132033"></a>

<a id="canonical-1130132132203220-0313013013110002-2233211032023220-3023110321031002-3233221021113303-0331303101123122-1311133212232103-1202031221102000"></a>

## auth_endpoint_url property — oidc_auth_params / 122331111010 / 4

Type: `"string"`. Optional.

URL of the authorization server's authorization endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0322001030333122-1311201122313033-3213002003131330-3120013022202021-1101310231023332-2213000113322210-1323332233002130-2112211132202133"></a>

<a id="canonical-2021110322223023-0221120202112010-1031220012010121-2231110332223011-0200320332201310-3212012200022322-3021330030211032-2011002330310302"></a>

## end_session_endpoint_url property — oidc_auth_params / 122331111010 / 5

Type: `"string"`. Optional.

URL of the authorization server's Logout endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1101001113231211-2100212210022013-0220102132331232-0233100202212323-2232212233203031-3203330232020201-2212213210121020-0330332010102222"></a>

<a id="canonical-0202330023223303-3203031312223001-3001221302232333-2120122130220313-0303212111032022-3132123133032333-2033203221233133-1100133000031100"></a>

## token_endpoint_url property — oidc_auth_params / 122331111010 / 6

Type: `"string"`. Optional.

URL of the authorization server's Token endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3322330001112130-3221323302013300-2333232311311202-1131232030310102-1022213100103022-0033031113132210-3202302320220022-2000232112022122"></a>

## Next pages — oidc_auth_params / 122331111010 / 7

- [oidc_auth](resources--authentication--reference--group-001.md#canonical-0232111213313223-0233011011200031-1232001100210323-2032032203100301-0100031032120210-2330333301330320-3200031233123113-2032233203332222)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)

<a id="canonical-3303312123102202-1020231312022121-0002100023333100-3230010222333203-3101213130130031-3312230101321112-2323131100121313-1211003123300030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110102122320130-2233232030202021-0123220113113032-3133011330210323-0203032031102221-0001012301101033-2323001222332230-3023223011031012"></a>

## timeouts — timeouts / 210021110333 / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- timeouts

<a id="canonical-1312231330210232-0212102012112211-1223332203033223-3223332113131001-2123311112211002-0013321121321310-1020230313031121-1313333103211212"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033003333033123-1131211123000103-2331223131221321-3102233332310122-3101132110210133-3333313311122030-1131211321121103-2113030101110223"></a>

## Direct properties — timeouts / 210021110333 / 3

<a id="canonical-3321132333212111-0133322211021130-0112221211201122-2113033201213112-0112031230221000-2100030130322010-3310230312322012-2303133323222213"></a>

<a id="canonical-3300212121333031-1333102310123021-1102033123111310-0212123330021230-0220222102233102-3310303320103331-1030000003013020-3030201100103112"></a>

## create property — timeouts / 210021110333 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0012121131300302-3202002100121211-2020200213003013-2232323000011000-1013200333302210-0231131301011121-2013031333310123-2200032211220232"></a>

<a id="canonical-3102332033333113-1132332212333322-0312213032300130-2303312010112233-0223301212203000-0122020310110131-2032202232032220-1100010023203320"></a>

## delete property — timeouts / 210021110333 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1233023122310303-0012130113300330-1233321100131222-0321030003111211-3032331203012133-3202011012011022-0020113222030320-2021302133033133"></a>

<a id="canonical-0221202002130300-2023202110311110-1021310311203203-3010111003231322-1121011312130121-0001222110000320-0113303122131131-2223302033203202"></a>

## read property — timeouts / 210021110333 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0320131111121222-3013313230101121-1122012213022203-0312310022332122-2312311022320220-3031233102221113-0110333332223131-0130000120322232"></a>

<a id="canonical-0331122202033031-1030221213113311-3013012312233031-0330201212132021-0301201201032323-1231001112000131-3013301131101120-1133230111110130"></a>

## update property — timeouts / 210021110333 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3301123310212133-0233232021323002-3330321320230231-3021320302222032-0312130010213022-3033200321323313-3300210313110003-2233113020220131"></a>

## Next pages — timeouts / 210021110333 / 8

- [Property reference](resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
