---
page_title: "xcsh_authentication reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication reference."
---

# xcsh_authentication reference

<a id="canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- Property reference

<a id="canonical-0002201131230311-0122011031021103-0302113100031200-2022020200133003-1311102032301003-2203103303122023-0112312202122232-2202201122022313"></a>

### Direct properties for `xcsh_authentication`

<a id="canonical-3212033033323310-1213001120121100-1120210133100213-1223023000301033-3132333132313202-2232210323203212-1223021100032310-0221120201132003"></a>

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

- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333): complete subsection reference.

<a id="canonical-0103302321303232-0230012232113302-2221213100331210-2013321013320303-0310121000033232-2331121303130333-2103130123110331-0302121230221200"></a>

<a id="canonical-1232012113030023-3203130000103223-2133211220002322-0122310132003123-3212120230311111-0320112003310213-0303203120323100-1120102030110003"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Authentication.

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

<a id="canonical-2330222120101221-1331300020021212-0113111212203231-0102333221001202-1311212200032300-2310133113231120-0221122203323111-3303111323123013"></a>

<a id="canonical-1111032022210112-1111112111230221-2122200131320032-1120231020300133-0333002122012223-2212012211100121-1332021100320013-2011002311312100"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0031210211202311-2232003213111032-0012112221000122-3032233121333222-0132022221330301-1130331030320322-1310230230023013-0112120130301020"></a>

<a id="canonical-1003323003200233-0223322303323101-3230100333203112-2202131223200312-3210032030011303-0301221020232313-0311230311231102-0321323211311203"></a>

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

<a id="canonical-3123302030023223-0102323200301330-1301132331122312-1322203220031221-1230322322103020-3232213301103231-0003221200301030-1133231122122312"></a>

<a id="canonical-0331031130030330-2333021033303303-1323022310222213-1302131100220310-1313232102132300-3301230032000320-0232333202313230-2213100330111201"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Authentication.

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

<a id="canonical-3202232231030100-0223122221311330-0013222120213031-2102231012002023-3030331111332011-2212213120203012-3030130221311330-1211331213030233"></a>

<a id="canonical-0100202222220322-1313020202223232-1112010020010203-0201102211320003-2231310320311230-1221102222211030-0222202231202202-0310333332001010"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Authentication exists.

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

- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0031110023022323-1212131130130231-2113110302012310-3311202301313322-2312132012100201-2323221103000200-2310003300131031-0221311313011123): complete subsection reference.

<a id="canonical-0232230310233133-1021232032311223-2110210002301202-1301302233212131-2103023221321321-3230301200333130-0231123001333212-0323003113022331"></a>

### All schema paths for `xcsh_authentication`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--authentication--reference--group-001.md#canonical-3212033033323310-1213001120121100-1120210133100213-1223023000301033-3132333132313202-2232210323203212-1223021100032310-0221120201132003) |
| `cookie_params` | [cookie_params](data-sources--authentication--reference--group-001.md#canonical-2332231020032132-3303323121112333-0213232102203202-2010131132320301-0130211113200112-2001103022330131-0121132132111130-3313132232023313) |
| `cookie_params.auth_hmac` | [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-0001010132201221-3100102012120311-3123300202313010-2321013102003312-0213203313103022-0002321230232031-2333131311323231-2120200113211331) |
| `cookie_params.auth_hmac.prim_key` | [cookie_params.auth_hmac.prim_key](data-sources--authentication--reference--group-001.md#canonical-2213003302031211-2021301122000320-2331200120111103-3033322132032111-3330012101210311-2032101020302131-2310312113033131-0300012201030321) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-0310302222223131-1013103002300232-0323030221020103-2022101101300002-3013302203111320-1221230000302112-1122203012100202-2210120023302313) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](data-sources--authentication--reference--group-001.md#canonical-0331033031311122-0113112013302221-2010022133221220-1222122122210310-0120101233301102-2211012203310031-0323233022302021-0101200333323132) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](data-sources--authentication--reference--group-001.md#canonical-2201231121220220-1202110310031231-2223203200310233-3233033001110330-0231113021211130-0133123113003001-3011021032001032-0021111333313322) |
| `cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](data-sources--authentication--reference--group-001.md#canonical-0032023000010120-2313132011300113-0001132232330113-2122023232010000-3220211202033011-1311113132113001-3230123332331303-2120232331210312) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info` | [cookie_params.auth_hmac.prim_key.clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-2022330103303030-1031021320123112-1223330333122201-0013012331123311-3120311221113312-2130103033102332-0330220333020303-2210200323232132) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](data-sources--authentication--reference--group-001.md#canonical-1200022213123203-1133122012013123-3123333220022130-0112001110211233-1301302121331122-2302213022313311-2100003311021231-3212110032313101) |
| `cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [cookie_params.auth_hmac.prim_key.clear_secret_info.url](data-sources--authentication--reference--group-001.md#canonical-2103330123020200-3221311210222102-2323101011133220-1011230302021233-2213312031103021-2220210210331311-2112102033323322-1233111103302000) |
| `cookie_params.auth_hmac.prim_key_expiry` | [cookie_params.auth_hmac.prim_key_expiry](data-sources--authentication--reference--group-001.md#canonical-0313233132223010-0111220212113023-1200133220213133-2310232221103020-0011022003332330-3010211132101220-0223012231320232-3303002203000033) |
| `cookie_params.auth_hmac.sec_key` | [cookie_params.auth_hmac.sec_key](data-sources--authentication--reference--group-001.md#canonical-2003221101333220-2112210130321210-0113310201013000-3220203220023010-2312123003221110-2221001221220202-3113232110030030-0131103233212210) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-1301231113122102-3001211113103003-1031120303010213-1311230130320200-3111331013010021-1021111213011312-2231100323112011-1330231111232103) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](data-sources--authentication--reference--group-001.md#canonical-1122221221121022-2221303201021122-1001131112332223-3011030201113013-2130310230121221-2213132212223212-3112033230220012-1331233020022222) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](data-sources--authentication--reference--group-001.md#canonical-0010230023310001-0320120330323220-2002113232100033-0113232111113030-2123220333132233-0220212222310323-1230110101122100-0110230331011033) |
| `cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](data-sources--authentication--reference--group-001.md#canonical-2033302113022102-2220031310021003-2002201320303332-0221312110301022-3011003013231103-3200311000222122-3012331200223111-2013030013112002) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info` | [cookie_params.auth_hmac.sec_key.clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-1020022000302212-2322110231212323-1331310300031321-0122000021231022-1113203001020212-0132030020220200-2001102030011201-2302003020001032) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](data-sources--authentication--reference--group-001.md#canonical-0231102332032301-3332230132210221-2011310323232122-3111210110233310-1303101110013021-2030200202032112-3012111131210231-0223233223100031) |
| `cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [cookie_params.auth_hmac.sec_key.clear_secret_info.url](data-sources--authentication--reference--group-001.md#canonical-3111131320000211-3000321211203220-3012102113203300-3231002231323033-1013020332100300-0333103310311123-0032013311302023-1310303202021213) |
| `cookie_params.auth_hmac.sec_key_expiry` | [cookie_params.auth_hmac.sec_key_expiry](data-sources--authentication--reference--group-001.md#canonical-0011120103222003-2033000010222023-0031000200013332-3300001221210233-0301003013323310-1122333310130300-1100312100022031-1322320023133211) |
| `cookie_params.cookie_expiry` | [cookie_params.cookie_expiry](data-sources--authentication--reference--group-001.md#canonical-2332331112120000-0123003003221231-2100102122302210-2321132002311003-3303120030232031-1011123200222221-3332121113300221-2303203230022331) |
| `cookie_params.cookie_refresh_interval` | [cookie_params.cookie_refresh_interval](data-sources--authentication--reference--group-001.md#canonical-0112313131223320-3222321321111201-1001211112021301-2020311102121120-2212030211222111-2100330111123030-1000221131300230-3023222312212011) |
| `cookie_params.kms_key_hmac` | [cookie_params.kms_key_hmac](data-sources--authentication--reference--group-001.md#canonical-3312033312031220-1230023033231222-0022330022001300-2220020011301010-0330000100001102-3320030212233322-1102010013100102-0331300333030010) |
| `cookie_params.session_expiry` | [cookie_params.session_expiry](data-sources--authentication--reference--group-001.md#canonical-2132213210132221-0211200322203120-1332210032222231-0200122211330203-0311233032012103-3333230012112230-2120030322322202-3231110231312200) |
| `description` | [description](data-sources--authentication--reference--group-001.md#canonical-0103302321303232-0230012232113302-2221213100331210-2013321013320303-0310121000033232-2331121303130333-2103130123110331-0302121230221200) |
| `id` | [ID](data-sources--authentication--reference--group-001.md#canonical-2330222120101221-1331300020021212-0113111212203231-0102333221001202-1311212200032300-2310133113231120-0221122203323111-3303111323123013) |
| `labels` | [labels](data-sources--authentication--reference--group-001.md#canonical-0031210211202311-2232003213111032-0012112221000122-3032233121333222-0132022221330301-1130331030320322-1310230230023013-0112120130301020) |
| `name` | [name](data-sources--authentication--reference--group-001.md#canonical-3123302030023223-0102323200301330-1301132331122312-1322203220031221-1230322322103020-3232213301103231-0003221200301030-1133231122122312) |
| `namespace` | [namespace](data-sources--authentication--reference--group-001.md#canonical-3202232231030100-0223122221311330-0013222120213031-2102231012002023-3030331111332011-2212213120203012-3030130221311330-1211331213030233) |
| `oidc_auth` | [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-1012130211010111-0210113010133331-2101323133121102-0332200220132311-1202231222102032-1023102220201000-0010111023300100-2303111022233233) |
| `oidc_auth.client_secret` | [oidc_auth.client_secret](data-sources--authentication--reference--group-001.md#canonical-2213132310203303-3002231030012031-3322033121003012-0132310131101232-0011333220311131-3201033002121103-3110332302213120-0122013322020302) |
| `oidc_auth.client_secret.blindfold_secret_info` | [oidc_auth.client_secret.blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-3310011002111320-2223001112003020-2033232313013203-3111330030320022-0101320033322012-3030210231003033-1303211130030013-3231221001030233) |
| `oidc_auth.client_secret.blindfold_secret_info.decryption_provider` | [oidc_auth.client_secret.blindfold_secret_info.decryption_provider](data-sources--authentication--reference--group-001.md#canonical-2223020330032013-2010332103331232-0003132003011012-0300121122132103-0321330133332222-3210320130123130-3232320010313111-2321200233112130) |
| `oidc_auth.client_secret.blindfold_secret_info.location` | [oidc_auth.client_secret.blindfold_secret_info.location](data-sources--authentication--reference--group-001.md#canonical-2331102102230202-2311032321100312-2221231002102202-2203202322133200-0333211210023212-0012210223112201-3222013330213200-2122223312322033) |
| `oidc_auth.client_secret.blindfold_secret_info.store_provider` | [oidc_auth.client_secret.blindfold_secret_info.store_provider](data-sources--authentication--reference--group-001.md#canonical-1310122233120022-3013032011110312-2023123032311323-1233331102123333-2020312011203312-3233100303000330-3331020113123231-0112120123121201) |
| `oidc_auth.client_secret.clear_secret_info` | [oidc_auth.client_secret.clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-3222210121110130-2022100220232332-2321311023020031-3300213300023101-1233133133032020-0102122120031323-0132000031021003-0210333023002200) |
| `oidc_auth.client_secret.clear_secret_info.provider_ref` | [oidc_auth.client_secret.clear_secret_info.provider_ref](data-sources--authentication--reference--group-001.md#canonical-3321202103220300-3223202122120122-1133121101333130-1210311122332300-2331221113332210-1213123210313123-0100210201321221-0113212313201223) |
| `oidc_auth.client_secret.clear_secret_info.url` | [oidc_auth.client_secret.clear_secret_info.url](data-sources--authentication--reference--group-001.md#canonical-0110131311120310-0300302323203101-1023232230132012-3220202120313213-2000033133332231-1330013203001013-1031332111002030-1022331210210033) |
| `oidc_auth.oidc_auth_params` | [oidc_auth.oidc_auth_params](data-sources--authentication--reference--group-001.md#canonical-0203120310310321-3101301132110030-1110013133021100-1102332023223213-1002132201322120-2320333033012322-1121111202220322-2102003121321023) |
| `oidc_auth.oidc_auth_params.auth_endpoint_url` | [oidc_auth.oidc_auth_params.auth_endpoint_url](data-sources--authentication--reference--group-001.md#canonical-3323223302030210-1021021200231031-3201111103302012-0211231210202313-1222023312100121-0012200031300200-1310120213121113-1001222112003313) |
| `oidc_auth.oidc_auth_params.end_session_endpoint_url` | [oidc_auth.oidc_auth_params.end_session_endpoint_url](data-sources--authentication--reference--group-001.md#canonical-1100110223132031-3103001111303033-3321231300030032-1302323111213103-0130100321333313-0332320122332233-3233001233111022-1021322031221231) |
| `oidc_auth.oidc_auth_params.token_endpoint_url` | [oidc_auth.oidc_auth_params.token_endpoint_url](data-sources--authentication--reference--group-001.md#canonical-3133321330103120-3232121211021300-3022302330311303-3010232121210201-0123000222223132-2232111301001033-2112333230112120-0323332321003330) |
| `oidc_auth.oidc_client_id` | [oidc_auth.oidc_client_id](data-sources--authentication--reference--group-001.md#canonical-2311302300023112-0200213232203332-2113003133110223-3322100012001131-2002333132001203-3122310002023023-3113230031233013-3220031303112111) |
| `oidc_auth.oidc_well_known_config_url` | [oidc_auth.oidc_well_known_config_url](data-sources--authentication--reference--group-001.md#canonical-1301103201200330-1111013013130231-1201333222013010-2213230131323221-0200132020032010-2311122320333021-2030020131022121-1320302101211103) |

<a id="canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_params` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- cookie_params

<a id="canonical-2332231020032132-3303323121112333-0213232102203202-2010131132320301-0130211113200112-2001103022330131-0121132132111130-3313132232023313"></a>

Type: `"single"`. Computed.

Specifies different cookie related config parameters for authentication.

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

<a id="canonical-3101022021110111-3020002310233103-1222123012133221-0131331300003233-2332303121021103-0222130222321230-0032030330210301-2321132012212001"></a>

### Direct properties for `cookie_params`

- [auth_hmac](data-sources--authentication--reference--group-001.md#canonical-2303222302130330-3012122012220000-3131320001202302-2323313002202101-3000120022013222-3010101312301201-2131303332100031-3011220001023323): complete subsection reference.

<a id="canonical-2332331112120000-0123003003221231-2100102122302210-2321132002311003-3303120030232031-1011123200222221-3332121113300221-2303203230022331"></a>

<a id="canonical-3330013130033112-2233021303202323-1201123331131110-3231002331130230-1233200023101212-3100302033003020-1321110301101313-1321311313202302"></a>

#### `cookie_params.cookie_expiry` property

Type: `"number"`. Computed.

Specifies in seconds max duration of the allocated cookie. This maps to “Max-Age” attribute in the
session cookie. This will act as an expiry duration on the client-side after which client will not
be setting the cookie as part of the request. Default cookie expiry is 3600 seconds.

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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0112313131223320-3222321321111201-1001211112021301-2020311102121120-2212030211222111-2100330111123030-1000221131300230-3023222312212011"></a>

<a id="canonical-1202222331000022-1331010010300221-1101120213023023-2110233100313301-0312231021132311-2020202311031112-1111123001101210-3210300211020303"></a>

#### `cookie_params.cookie_refresh_interval` property

Type: `"number"`. Computed.

Specifies in seconds refresh interval for session cookie. This is used to keep the active user
active and reduce RE-login. When an incoming cookie's session expiry is still valid, and time to
expire falls behind this interval, RE-issue a cookie with new expiry and with the same original
session expiry. Default refresh interval is 3000 seconds.

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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

- [kms_key_hmac](data-sources--authentication--reference--group-001.md#canonical-1122032132211133-3221102303211012-3310003020230231-2020330020233321-2301211333312333-2220222021012130-3120332310133212-1100101331103330): complete subsection reference.

<a id="canonical-2132213210132221-0211200322203120-1332210032222231-0200122211330203-0311233032012103-3333230012112230-2120030322322202-3231110231312200"></a>

<a id="canonical-2033202111101133-1303233201322332-0200103233022201-0110000333330020-2133012303002212-3122013013103310-1321022033202313-0210011133311123"></a>

#### `cookie_params.session_expiry` property

Type: `"number"`. Computed.

Specifies in seconds max lifetime of an authenticated session after which the user will be forced to
login again. Default session expiry is 86400 seconds(24 hours).

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
    "ves.io.schema.rules.uint32.lte": "1296000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1296000"
  }
}
```

<a id="canonical-2303222302130330-3012122012220000-3131320001202302-2323313002202101-3000120022013222-3010101312301201-2131303332100031-3011220001023323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_params.auth_hmac` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333)
- cookie_params.auth_hmac

<a id="canonical-0001010132201221-3100102012120311-3123300202313010-2321013102003312-0213203313103022-0002321230232031-2333131311323231-2120200113211331"></a>

Type: `"single"`. Computed.

HMAC primary and secondary keys to be used for hashing the Cookie. Each key also have an associated
expiry timestamp, beyond which key is invalid.

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

<a id="canonical-1211102020202111-3333102002113102-3131132113303301-0113311210211222-0120122231003103-1033220031103021-2301213120010301-2312101321301330"></a>

### Direct properties for `cookie_params.auth_hmac`

- [prim_key](data-sources--authentication--reference--group-001.md#canonical-3321032301130201-0113110300133310-0201011031110333-1220021311320202-2203301321130321-0222123101311323-1013312121311031-1302112212300011): complete subsection reference.

<a id="canonical-0313233132223010-0111220212113023-1200133220213133-2310232221103020-0011022003332330-3010211132101220-0223012231320232-3303002203000033"></a>

<a id="canonical-2133102101303030-2013323231002020-0213011312321023-1321313200201103-0130122313212023-3200313033213120-0231302211311031-2332321030121000"></a>

#### `cookie_params.auth_hmac.prim_key_expiry` property

Type: `"string"`. Computed.

HMAC Primary Key Expiry. Primary HMAC Key Expiry time.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [sec_key](data-sources--authentication--reference--group-001.md#canonical-0233322220022012-2130111323222131-2221113230111320-1310210030132213-2233232313033201-0211130320131123-3123320311322201-1001100311222201): complete subsection reference.

<a id="canonical-0011120103222003-2033000010222023-0031000200013332-3300001221210233-0301003013323310-1122333310130300-1100312100022031-1322320023133211"></a>

<a id="canonical-2012221023213300-3300332120311131-2102020021110200-2002103212002130-2203220002113032-1210220033133111-0010231211122101-0013023133322030"></a>

#### `cookie_params.auth_hmac.sec_key_expiry` property

Type: `"string"`. Computed.

HMAC Secondary Key Expiry. Secondary HMAC Key Expiry time.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3321032301130201-0113110300133310-0201011031110333-1220021311320202-2203301321130321-0222123101311323-1013312121311031-1302112212300011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_params.auth_hmac.prim_key` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-2303222302130330-3012122012220000-3131320001202302-2323313002202101-3000120022013222-3010101312301201-2131303332100031-3011220001023323)
- cookie_params.auth_hmac.prim_key

<a id="canonical-2213003302031211-2021301122000320-2331200120111103-3033322132032111-3330012101210311-2032101020302131-2310312113033131-0300012201030321"></a>

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

<a id="canonical-3200032002322232-3211202120001002-2003230232000201-2002200113203032-0222132002201111-0033121022313100-2201030303132010-1030321230310100"></a>

### Direct properties for `cookie_params.auth_hmac.prim_key`

- [blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-1333102221213011-3332010220211130-3333222211011222-2210123203013312-1333131331120203-2112320301111223-0311113330221112-0211031002012100): complete subsection reference.

- [clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-0102110320223032-2203201201021032-1202311003111222-2122212030011010-3210202133223300-2320030020231322-2001122013330101-2211121021210303): complete subsection reference.

<a id="canonical-1333102221213011-3332010220211130-3333222211011222-2210123203013312-1333131331120203-2112320301111223-0311113330221112-0211031002012100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_params.auth_hmac.prim_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-2303222302130330-3012122012220000-3131320001202302-2323313002202101-3000120022013222-3010101312301201-2131303332100031-3011220001023323)
- [cookie_params.auth_hmac.prim_key](data-sources--authentication--reference--group-001.md#canonical-3321032301130201-0113110300133310-0201011031110333-1220021311320202-2203301321130321-0222123101311323-1013312121311031-1302112212300011)
- cookie_params.auth_hmac.prim_key.blindfold_secret_info

<a id="canonical-0310302222223131-1013103002300232-0323030221020103-2022101101300002-3013302203111320-1221230000302112-1122203012100202-2210120023302313"></a>

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

<a id="canonical-2301200102022101-3131233030331213-2103223200021123-1201120202121322-2032323312320303-0233110230303213-2203102002221212-2213022221210012"></a>

### Direct properties for `cookie_params.auth_hmac.prim_key.blindfold_secret_info`

<a id="canonical-0331033031311122-0113112013302221-2010022133221220-1222122122210310-0120101233301102-2211012203310031-0323233022302021-0101200333323132"></a>

#### `cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2201231121220220-1202110310031231-2223203200310233-3233033001110330-0231113021211130-0133123113003001-3011021032001032-0021111333313322"></a>

<a id="canonical-2023130230000232-0231200020202121-2133202100010331-3233002300201330-3000023101013010-3102202311101303-3003002003013122-2100231103330310"></a>

#### `cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` property

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

<a id="canonical-0032023000010120-2313132011300113-0001132232330113-2122023232010000-3220211202033011-1311113132113001-3230123332331303-2120232331210312"></a>

<a id="canonical-2001221010020312-2310331332231000-1011210111011010-0131001033233302-3103302012213310-3331220123120313-0122132300003113-1300112003202231"></a>

#### `cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-0102110320223032-2203201201021032-1202311003111222-2122212030011010-3210202133223300-2320030020231322-2001122013330101-2211121021210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_params.auth_hmac.prim_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-2303222302130330-3012122012220000-3131320001202302-2323313002202101-3000120022013222-3010101312301201-2131303332100031-3011220001023323)
- [cookie_params.auth_hmac.prim_key](data-sources--authentication--reference--group-001.md#canonical-3321032301130201-0113110300133310-0201011031110333-1220021311320202-2203301321130321-0222123101311323-1013312121311031-1302112212300011)
- cookie_params.auth_hmac.prim_key.clear_secret_info

<a id="canonical-2022330103303030-1031021320123112-1223330333122201-0013012331123311-3120311221113312-2130103033102332-0330220333020303-2210200323232132"></a>

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

<a id="canonical-2331111310013221-0130032113100232-0303310203113100-2131313313210221-2230001200231233-0031111330213130-0222131203021020-2121032312110121"></a>

### Direct properties for `cookie_params.auth_hmac.prim_key.clear_secret_info`

<a id="canonical-1200022213123203-1133122012013123-3123333220022130-0112001110211233-1301302121331122-2302213022313311-2100003311021231-3212110032313101"></a>

#### `cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2103330123020200-3221311210222102-2323101011133220-1011230302021233-2213312031103021-2220210210331311-2112102033323322-1233111103302000"></a>

<a id="canonical-1122013330011020-0031013030313320-0210311102003211-0223103030320233-0323310320320311-2311011001330203-3133213012000031-1200211331110031"></a>

#### `cookie_params.auth_hmac.prim_key.clear_secret_info.url` property

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

<a id="canonical-0233322220022012-2130111323222131-2221113230111320-1310210030132213-2233232313033201-0211130320131123-3123320311322201-1001100311222201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_params.auth_hmac.sec_key` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-2303222302130330-3012122012220000-3131320001202302-2323313002202101-3000120022013222-3010101312301201-2131303332100031-3011220001023323)
- cookie_params.auth_hmac.sec_key

<a id="canonical-2003221101333220-2112210130321210-0113310201013000-3220203220023010-2312123003221110-2221001221220202-3113232110030030-0131103233212210"></a>

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

<a id="canonical-0130032201020002-0210201111221311-1033000132013022-1111223233032012-1133013232003312-1303132032303102-0103023112020331-3320131310231313"></a>

### Direct properties for `cookie_params.auth_hmac.sec_key`

- [blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-2100013302203110-0013232130231223-1210201331333101-2120233002013021-3013031120230320-1203102201333203-1020000200032222-3212112301023022): complete subsection reference.

- [clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-3021203302103031-1010001303211213-1113212300120201-3210123022220322-0023101031032121-1012200322302030-3133031113201003-2330330312232003): complete subsection reference.

<a id="canonical-2100013302203110-0013232130231223-1210201331333101-2120233002013021-3013031120230320-1203102201333203-1020000200032222-3212112301023022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_params.auth_hmac.sec_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-2303222302130330-3012122012220000-3131320001202302-2323313002202101-3000120022013222-3010101312301201-2131303332100031-3011220001023323)
- [cookie_params.auth_hmac.sec_key](data-sources--authentication--reference--group-001.md#canonical-0233322220022012-2130111323222131-2221113230111320-1310210030132213-2233232313033201-0211130320131123-3123320311322201-1001100311222201)
- cookie_params.auth_hmac.sec_key.blindfold_secret_info

<a id="canonical-1301231113122102-3001211113103003-1031120303010213-1311230130320200-3111331013010021-1021111213011312-2231100323112011-1330231111232103"></a>

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

<a id="canonical-1032210211002330-2210022222323001-1231001301223101-1101022011101313-1111230001221203-1233023003022122-0223113132212023-2221121231312113"></a>

### Direct properties for `cookie_params.auth_hmac.sec_key.blindfold_secret_info`

<a id="canonical-1122221221121022-2221303201021122-1001131112332223-3011030201113013-2130310230121221-2213132212223212-3112033230220012-1331233020022222"></a>

#### `cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0010230023310001-0320120330323220-2002113232100033-0113232111113030-2123220333132233-0220212222310323-1230110101122100-0110230331011033"></a>

<a id="canonical-2201302323120312-0223103301031012-1122030112001020-1221301033230122-1312100332122210-2121122211300320-1022230023331101-2212330100121133"></a>

#### `cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` property

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

<a id="canonical-2033302113022102-2220031310021003-2002201320303332-0221312110301022-3011003013231103-3200311000222122-3012331200223111-2013030013112002"></a>

<a id="canonical-1321101102003021-0301112130100200-2123011320220231-0323002220313211-2333211111322330-0303201201321230-0112231322200322-3312003223002311"></a>

#### `cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-3021203302103031-1010001303211213-1113212300120201-3210123022220322-0023101031032121-1012200322302030-3133031113201003-2330330312232003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_params.auth_hmac.sec_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333)
- [cookie_params.auth_hmac](data-sources--authentication--reference--group-001.md#canonical-2303222302130330-3012122012220000-3131320001202302-2323313002202101-3000120022013222-3010101312301201-2131303332100031-3011220001023323)
- [cookie_params.auth_hmac.sec_key](data-sources--authentication--reference--group-001.md#canonical-0233322220022012-2130111323222131-2221113230111320-1310210030132213-2233232313033201-0211130320131123-3123320311322201-1001100311222201)
- cookie_params.auth_hmac.sec_key.clear_secret_info

<a id="canonical-1020022000302212-2322110231212323-1331310300031321-0122000021231022-1113203001020212-0132030020220200-2001102030011201-2302003020001032"></a>

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

<a id="canonical-0000030330120311-0132112033201200-2103320223121223-0011331111003013-0233333130031303-1230002101331200-1002222023222212-3223322101203100"></a>

### Direct properties for `cookie_params.auth_hmac.sec_key.clear_secret_info`

<a id="canonical-0231102332032301-3332230132210221-2011310323232122-3111210110233310-1303101110013021-2030200202032112-3012111131210231-0223233223100031"></a>

#### `cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3111131320000211-3000321211203220-3012102113203300-3231002231323033-1013020332100300-0333103310311123-0032013311302023-1310303202021213"></a>

<a id="canonical-3132323032330312-0113122132010301-3120210112130001-0000103013031331-0213131212003020-0131211231213201-0022332333133223-1131012130312111"></a>

#### `cookie_params.auth_hmac.sec_key.clear_secret_info.url` property

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

<a id="canonical-1122032132211133-3221102303211012-3310003020230231-2020330020233321-2301211333312333-2220222021012130-3120332310133212-1100101331103330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_params.kms_key_hmac` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [cookie_params](data-sources--authentication--reference--group-001.md#canonical-1300021200300133-3100001310211100-2030223301303321-2302101322131100-3223222311000313-2113212332213331-0112020123202230-1120121332232333)
- cookie_params.kms_key_hmac

<a id="canonical-3312033312031220-1230023033231222-0022330022001300-2220020011301010-0330000100001102-3320030212233322-1102010013100102-0331300333030010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for kms key hmac.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031110023022323-1212131130130231-2113110302012310-3311202301313322-2312132012100201-2323221103000200-2310003300131031-0221311313011123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oidc_auth` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- oidc_auth

<a id="canonical-1012130211010111-0210113010133331-2101323133121102-0332200220132311-1202231222102032-1023102220201000-0010111023300100-2303111022233233"></a>

Type: `"single"`. Computed.

OIDCAuthType.

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

<a id="canonical-2223102112010232-3100110200000212-3031302313312301-0103003221312310-0301130023321322-1300103132123203-2210002132223003-3110101103333011"></a>

### Direct properties for `oidc_auth`

- [client_secret](data-sources--authentication--reference--group-001.md#canonical-3230011212002033-3231130123201030-0221202213110032-1110301332101132-3330101221132121-1033323000113201-1033122020221130-2030100331302021): complete subsection reference.

- [oidc_auth_params](data-sources--authentication--reference--group-001.md#canonical-1221330010102002-2102120122013232-1202123012303113-0312030123202321-0102003103331033-0213013330332320-1012102300010132-2232213123132012): complete subsection reference.

<a id="canonical-2311302300023112-0200213232203332-2113003133110223-3322100012001131-2002333132001203-3122310002023023-3113230031233013-3220031303112111"></a>

<a id="canonical-1220223120230310-3102323202202203-3221112003122012-2131202300003232-1002120222311332-2030210012013121-2022010211103132-0002020223023313"></a>

#### `oidc_auth.oidc_client_id` property

Type: `"string"`. Computed.

Client ID used while sending the Authorization Request to OIDC server.

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

<a id="canonical-1301103201200330-1111013013130231-1201333222013010-2213230131323221-0200132020032010-2311122320333021-2030020131022121-1320302101211103"></a>

<a id="canonical-0331211223023333-1303011203113332-0000122021133110-1210301110110023-2113031002232012-3132201203302331-3302110301301101-0232302232202202"></a>

#### `oidc_auth.oidc_well_known_config_url` property

Type: `"string"`. Computed.

Exclusive with \[oidc\_auth\_params\] An OIDC well-known configuration URL that will be used to
fetch authentication related endpoints.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3230011212002033-3231130123201030-0221202213110032-1110301332101132-3330101221132121-1033323000113201-1033122020221130-2030100331302021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oidc_auth.client_secret` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0031110023022323-1212131130130231-2113110302012310-3311202301313322-2312132012100201-2323221103000200-2310003300131031-0221311313011123)
- oidc_auth.client_secret

<a id="canonical-2213132310203303-3002231030012031-3322033121003012-0132310131101232-0011333220311131-3201033002121103-3110332302213120-0122013322020302"></a>

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

<a id="canonical-1201211220322101-2330000233303003-1233310201313210-2311301023101332-0030330313220032-0113210030311011-2322112100213323-0110211230003320"></a>

### Direct properties for `oidc_auth.client_secret`

- [blindfold_secret_info](data-sources--authentication--reference--group-001.md#canonical-3001030121221333-3120302031320123-1103310332130203-0231203102212223-3221331002020220-2230310030121203-1001223231100233-3202101332012320): complete subsection reference.

- [clear_secret_info](data-sources--authentication--reference--group-001.md#canonical-1201103321223230-0311223321303011-0101311130031021-1113102222112123-3131213202203133-1331202301310303-2101221330330031-2211123101300333): complete subsection reference.

<a id="canonical-3001030121221333-3120302031320123-1103310332130203-0231203102212223-3221331002020220-2230310030121203-1001223231100233-3202101332012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oidc_auth.client_secret.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0031110023022323-1212131130130231-2113110302012310-3311202301313322-2312132012100201-2323221103000200-2310003300131031-0221311313011123)
- [oidc_auth.client_secret](data-sources--authentication--reference--group-001.md#canonical-3230011212002033-3231130123201030-0221202213110032-1110301332101132-3330101221132121-1033323000113201-1033122020221130-2030100331302021)
- oidc_auth.client_secret.blindfold_secret_info

<a id="canonical-3310011002111320-2223001112003020-2033232313013203-3111330030320022-0101320033322012-3030210231003033-1303211130030013-3231221001030233"></a>

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

<a id="canonical-0201013131022023-3000301010130000-2021001002010311-1223133303122200-0202221312102030-3220303003132101-2102300200200233-3222332222232110"></a>

### Direct properties for `oidc_auth.client_secret.blindfold_secret_info`

<a id="canonical-2223020330032013-2010332103331232-0003132003011012-0300121122132103-0321330133332222-3210320130123130-3232320010313111-2321200233112130"></a>

#### `oidc_auth.client_secret.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2331102102230202-2311032321100312-2221231002102202-2203202322133200-0333211210023212-0012210223112201-3222013330213200-2122223312322033"></a>

<a id="canonical-0020300300113332-2312122000220021-1323212020013023-1101131111320113-3321123221210112-3031021321121110-1031221121112012-3211003223313310"></a>

#### `oidc_auth.client_secret.blindfold_secret_info.location` property

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

<a id="canonical-1310122233120022-3013032011110312-2023123032311323-1233331102123333-2020312011203312-3233100303000330-3331020113123231-0112120123121201"></a>

<a id="canonical-1013313301100101-3032130033333033-0103331201311333-0112232301011131-3121100021103331-3202303303122200-0031121033032333-3010033120211130"></a>

#### `oidc_auth.client_secret.blindfold_secret_info.store_provider` property

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

<a id="canonical-1201103321223230-0311223321303011-0101311130031021-1113102222112123-3131213202203133-1331202301310303-2101221330330031-2211123101300333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oidc_auth.client_secret.clear_secret_info` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0031110023022323-1212131130130231-2113110302012310-3311202301313322-2312132012100201-2323221103000200-2310003300131031-0221311313011123)
- [oidc_auth.client_secret](data-sources--authentication--reference--group-001.md#canonical-3230011212002033-3231130123201030-0221202213110032-1110301332101132-3330101221132121-1033323000113201-1033122020221130-2030100331302021)
- oidc_auth.client_secret.clear_secret_info

<a id="canonical-3222210121110130-2022100220232332-2321311023020031-3300213300023101-1233133133032020-0102122120031323-0132000031021003-0210333023002200"></a>

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

<a id="canonical-3101020302030130-2212303232202001-1012312112113001-2102320031310123-2200031032300322-3212000020122020-1112122030211223-3332220333320002"></a>

### Direct properties for `oidc_auth.client_secret.clear_secret_info`

<a id="canonical-3321202103220300-3223202122120122-1133121101333130-1210311122332300-2331221113332210-1213123210313123-0100210201321221-0113212313201223"></a>

#### `oidc_auth.client_secret.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0110131311120310-0300302323203101-1023232230132012-3220202120313213-2000033133332231-1330013203001013-1031332111002030-1022331210210033"></a>

<a id="canonical-3313133112312112-3130002111111012-1301311302021132-0011101033330333-3012231021222232-3230333231111203-3212300202201332-1033021233031133"></a>

#### `oidc_auth.client_secret.clear_secret_info.url` property

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

<a id="canonical-1221330010102002-2102120122013232-1202123012303113-0312030123202321-0102003103331033-0213013330332320-1012102300010132-2232213123132012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oidc_auth.oidc_auth_params` properties

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Property reference](data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [oidc_auth](data-sources--authentication--reference--group-001.md#canonical-0031110023022323-1212131130130231-2113110302012310-3311202301313322-2312132012100201-2323221103000200-2310003300131031-0221311313011123)
- oidc_auth.oidc_auth_params

<a id="canonical-0203120310310321-3101301132110030-1110013133021100-1102332023223213-1002132201322120-2320333033012322-1121111202220322-2102003121321023"></a>

Type: `"single"`. Computed.

Configuration parameter for oidc auth params.

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

<a id="canonical-1131213231313011-2202112010031212-2123032302132001-0022202211132333-2100231031103000-3123302120212222-3222231133201032-0113300032223323"></a>

### Direct properties for `oidc_auth.oidc_auth_params`

<a id="canonical-3323223302030210-1021021200231031-3201111103302012-0211231210202313-1222023312100121-0012200031300200-1310120213121113-1001222112003313"></a>

#### `oidc_auth.oidc_auth_params.auth_endpoint_url` property

Type: `"string"`. Computed.

URL of the authorization server's authorization endpoint.

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

<a id="canonical-1100110223132031-3103001111303033-3321231300030032-1302323111213103-0130100321333313-0332320122332233-3233001233111022-1021322031221231"></a>

<a id="canonical-0030220112131333-2120333012220200-1311011011110111-1233200001300032-2103202122031231-0030132330322133-2133220131200213-1312200101012101"></a>

#### `oidc_auth.oidc_auth_params.end_session_endpoint_url` property

Type: `"string"`. Computed.

URL of the authorization server's Logout endpoint.

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

<a id="canonical-3133321330103120-3232121211021300-3022302330311303-3010232121210201-0123000222223132-2232111301001033-2112333230112120-0323332321003330"></a>

<a id="canonical-3311121312113211-2012020301303100-2033203323021121-2322323031121211-3200133212130310-2131222022323131-1322222133020113-1311001211233130"></a>

#### `oidc_auth.oidc_auth_params.token_endpoint_url` property

Type: `"string"`. Computed.

URL of the authorization server's Token endpoint.

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
