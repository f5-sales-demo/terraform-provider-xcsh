---
page_title: "xcsh_cloud_credentials reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials reference."
---

# xcsh_cloud_credentials reference

<a id="canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- Property reference

<a id="canonical-1212001323233323-1011203130323103-3330331001211222-0001032323102201-1023112133100223-2231203303223032-1313032210031201-2231032030332332"></a>

### Direct properties for `xcsh_cloud_credentials`

<a id="canonical-2012200000312213-2332302233202333-3303021323210211-2203210101000121-2021120212011120-1121012202121213-0101320012013102-3330102230301313"></a>

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

- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-3201131310302023-0122011012123001-0130131130232113-3232133100213330-3012223312001231-3021312122320320-3313012123220011-3033311203311103): complete subsection reference.

- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-3233111313023123-2112001202221123-1003132111031121-1301010333300232-2201130032133313-3201203222322230-0230201100222323-2123010303221211): complete subsection reference.

- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0031003222133102-3220220230302231-3221203120103122-2132001113201132-2100013132100213-2210320123302223-3120332020133311-3031023322021323): complete subsection reference.

- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-3121121300221320-0323123210301120-3332103121012010-0133333031111230-3231111231030100-0102030303212231-0002210222100033-0213220221123010): complete subsection reference.

<a id="canonical-2232121111000120-0013123212310232-2322133013133331-0313023310302310-1233003312023123-1303331223222013-0220132201312131-1033100322331331"></a>

<a id="canonical-2120300032212032-0320333302212021-2201002210312233-2122011021113333-1031331332220020-3322130213002022-0312100303112232-1130333231013210"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0203301030112020-0030002131022310-1132310313132101-3303121031210022-3301220022202223-1222232313121020-2020001302113323-1120322332321003"></a>

<a id="canonical-1103011122223130-2322000311313233-3332302133020230-1333121101222213-3010230231130221-2030200303010030-1111111121302203-1130321210300011"></a>

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

- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-3222330220322011-0302130220202222-3211013011122330-0030301333103222-2332221013220321-3331320310312200-2132212011130212-0002121112103213): complete subsection reference.

<a id="canonical-1210321310122133-3033031310011332-0010221030002110-1101020012213230-1230102200313212-1221022100111330-2312331233100213-1201000033330130"></a>

<a id="canonical-1211033332200011-2320113332020213-1230132030202132-2103113213022321-0000311121210110-2032231221211202-3033010130212133-2320113131210203"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0200203123003201-1333000223221011-1011013003231230-2130300033030321-2302002322130032-3232013310120131-0210123010211000-3233000333303303"></a>

<a id="canonical-1013122113023211-2201031030303033-3003033130210300-3122331320322202-2303021012012330-2330313102100101-2121102210303103-3323212333110201"></a>

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

<a id="canonical-3131110202313223-3202103021101230-1020122320103111-3003000102331013-2232212232020010-1031022023113032-2311100202211320-2130212120032323"></a>

<a id="canonical-0302033231010130-0010113203323202-2103123120110210-2001031033123120-0112301112200221-3210210301212211-1110210223020213-1021232213332003"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Cloud Credentials. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2303113230130011-1101220112230020-0232133122022213-1222313022033112-1232231313011131-3113322133223111-0021032301213000-3230122213031220"></a>

<a id="canonical-2201002130032000-0300323121221331-2331021320200333-2301213331323112-2313332013332331-1323203210020023-0331212311010100-1321002220002232"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Cloud Credentials is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [timeouts](resources--cloud_credentials--reference--group-001.md#canonical-0322303012302112-3031111001122203-3201002312132113-1301113100000020-0130001003100221-3021000310323011-1230333102123130-3302111233300121): complete subsection reference.

<a id="canonical-3022100323112121-2203111213203023-0032021330222233-1033210330313332-1231012212230102-0132332013303132-1222002110330110-1101120322312133"></a>

### All schema paths for `xcsh_cloud_credentials`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_credentials--reference--group-001.md#canonical-2012200000312213-2332302233202333-3303021323210211-2203210101000121-2021120212011120-1121012202121213-0101320012013102-3330102230301313) |
| `aws_assume_role` | [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-2010330202131221-3331033201131031-0033321330121033-3331321002230022-0130310102333320-1001313300031220-1233232231022131-0200333111010302) |
| `aws_assume_role.custom_external_id` | [aws_assume_role.custom_external_id](resources--cloud_credentials--reference--group-001.md#canonical-0102231300311031-0122201102322012-2210203132123110-1313102301303330-3033101203331110-1211212231130203-1103210112111013-1121202132223013) |
| `aws_assume_role.duration_seconds` | [aws_assume_role.duration_seconds](resources--cloud_credentials--reference--group-001.md#canonical-3200023330202312-0113030230310310-1232100021331231-3021320203101111-2110031210030211-0230030111203001-0110020130131321-2201023223203022) |
| `aws_assume_role.external_id_is_optional` | [aws_assume_role.external_id_is_optional](resources--cloud_credentials--reference--group-001.md#canonical-1310012012331200-2323001202310013-0022130002213230-3210122013300100-2012130033113103-1022113003301120-3100113230010032-3331231012022312) |
| `aws_assume_role.external_id_is_tenant_id` | [aws_assume_role.external_id_is_tenant_id](resources--cloud_credentials--reference--group-001.md#canonical-3201103211110100-3133011220303303-1301100321123220-3330100332111222-1033101210200103-1202230113131313-0323131311102002-3013323101232231) |
| `aws_assume_role.role_arn` | [aws_assume_role.role_arn](resources--cloud_credentials--reference--group-001.md#canonical-0212313213223311-2123332300103330-3132110202101203-2003300232103320-2330200133011222-3132031021101001-0032231023101123-2323222232023213) |
| `aws_assume_role.session_name` | [aws_assume_role.session_name](resources--cloud_credentials--reference--group-001.md#canonical-2020221000003010-0221123032120112-2332130221312121-1213202113311221-3131332032321102-2313301203033203-1211022003222321-3312210210313233) |
| `aws_assume_role.session_tags` | [aws_assume_role.session_tags](resources--cloud_credentials--reference--group-001.md#canonical-0233100121322230-2300302322312301-0203110101221203-3110201331110200-2132103011113023-1220310202302300-1220103200320200-2230120013113312) |
| `aws_secret_key` | [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-0023023223123233-0110233210102020-0012022122021003-1013311213210320-1032303101022312-1222331221202232-2313321000130101-3311302322131220) |
| `aws_secret_key.access_key` | [aws_secret_key.access_key](resources--cloud_credentials--reference--group-001.md#canonical-2230230020012211-3331121000210100-3021201103320110-3232201032223201-3231303302300113-1213222031132113-3333030211213121-3213202320000301) |
| `aws_secret_key.secret_key` | [aws_secret_key.secret_key](resources--cloud_credentials--reference--group-001.md#canonical-1121333312013222-3221121111300031-1212332132010231-2131232122301001-0121323001130221-0210112010123033-3330112022010011-3013000320222312) |
| `aws_secret_key.secret_key.blindfold_secret_info` | [aws_secret_key.secret_key.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-3113100211102011-1223331000230301-1200302130323133-0032200333021101-3030023213233331-2013202210330301-3311132332030030-1113130113321203) |
| `aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](resources--cloud_credentials--reference--group-001.md#canonical-3031113030302002-0112301213013322-0203232131112313-0230130000302320-0221200221212011-0023110230303011-3100133320100311-1131001310331220) |
| `aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_secret_key.secret_key.blindfold_secret_info.location](resources--cloud_credentials--reference--group-001.md#canonical-2302111031212132-0031230220013123-2132323001120101-3320023103220302-1133131110032301-1133202202333030-3223121121022030-2300312023302221) |
| `aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_secret_key.secret_key.blindfold_secret_info.store_provider](resources--cloud_credentials--reference--group-001.md#canonical-2012321021313312-3133001003222102-0032300312320013-1023210013322010-0021310211201010-0120032002013303-3311030032113101-3132110231120123) |
| `aws_secret_key.secret_key.clear_secret_info` | [aws_secret_key.secret_key.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-1111301233212100-0021202003332131-1032130021101320-0020131012000301-1122230022231303-0002320200033233-1010323001211331-1223323110230301) |
| `aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_secret_key.secret_key.clear_secret_info.provider_ref](resources--cloud_credentials--reference--group-001.md#canonical-1032222102221103-0113133123100002-3311221120230103-1032022300311211-1210221020100330-1132010012201222-2010121003200033-1232131100311312) |
| `aws_secret_key.secret_key.clear_secret_info.url` | [aws_secret_key.secret_key.clear_secret_info.url](resources--cloud_credentials--reference--group-001.md#canonical-0101003111103022-3103331331132021-2310212111232112-0023311033100201-3121113123313223-1100201223133113-3232031012221000-1220230001011133) |
| `azure_client_secret` | [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-3310333023110001-0332313323233102-3201332102213110-0102330200133020-3313031333322311-3332010103031010-2121130001302120-2002203311022112) |
| `azure_client_secret.client_id` | [azure_client_secret.client_id](resources--cloud_credentials--reference--group-001.md#canonical-3030301330202312-1331013202232012-2001121322103031-1130100300023331-2211212301123200-3101120012031032-0111021102220120-0310110231011002) |
| `azure_client_secret.client_secret` | [azure_client_secret.client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0102103103330102-0203333133232212-2201133302220200-3030003320302132-2230220033203312-3011013013312323-2112000201103302-1111001112123233) |
| `azure_client_secret.client_secret.blindfold_secret_info` | [azure_client_secret.client_secret.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-3301322210033322-2201022330333230-1303323121101113-2100010312221310-2320200100101122-1132111112223130-3122131003223030-3000012302330103) |
| `azure_client_secret.client_secret.blindfold_secret_info.decryption_provider` | [azure_client_secret.client_secret.blindfold_secret_info.decryption_provider](resources--cloud_credentials--reference--group-001.md#canonical-1103001023300112-0202210101031310-3322212030212203-1113222123313000-1113331211122220-2213100011022330-2000110020223013-3030311311101132) |
| `azure_client_secret.client_secret.blindfold_secret_info.location` | [azure_client_secret.client_secret.blindfold_secret_info.location](resources--cloud_credentials--reference--group-001.md#canonical-0002101220003210-3212111001033211-2311210011021301-0030310100011033-3223103320323221-3122110303130203-1100201311303122-2332120233012232) |
| `azure_client_secret.client_secret.blindfold_secret_info.store_provider` | [azure_client_secret.client_secret.blindfold_secret_info.store_provider](resources--cloud_credentials--reference--group-001.md#canonical-3201330230112131-1001333012303303-3032122023110311-3031100032332002-0230331312023003-0333030222310011-0033010300101202-0212321330203221) |
| `azure_client_secret.client_secret.clear_secret_info` | [azure_client_secret.client_secret.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-3331102123301303-1220210000210032-1103010122103130-1313122303011010-0110022223232112-3302011002112022-1132303122231233-1032231211013013) |
| `azure_client_secret.client_secret.clear_secret_info.provider_ref` | [azure_client_secret.client_secret.clear_secret_info.provider_ref](resources--cloud_credentials--reference--group-001.md#canonical-2113320313222131-3303013302101120-3000332330110111-1121223112002130-1121321202200331-3112031213130120-0012103013210121-0000323221022023) |
| `azure_client_secret.client_secret.clear_secret_info.url` | [azure_client_secret.client_secret.clear_secret_info.url](resources--cloud_credentials--reference--group-001.md#canonical-2220202031012302-2001321101231003-0213102333103110-2222012202133031-1231102200031321-1111230330313323-1210132131001123-3203133101120132) |
| `azure_client_secret.subscription_id` | [azure_client_secret.subscription_id](resources--cloud_credentials--reference--group-001.md#canonical-2033231202231230-2002211332232300-0122223320123130-3001232001013011-0322323320221030-3023111102112230-3120201121311131-3021231003030003) |
| `azure_client_secret.tenant_id` | [azure_client_secret.tenant_id](resources--cloud_credentials--reference--group-001.md#canonical-0321132230203232-0010233002310212-1100203330031131-2222220033123332-0120112030002212-0323220023112002-2230132333133011-3103311222223030) |
| `azure_pfx_certificate` | [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-3133023000012002-3121021330033032-0030301001130301-1331333301221311-0031100002122123-1230122323130010-1012221100310221-2320101012131111) |
| `azure_pfx_certificate.certificate_url` | [azure_pfx_certificate.certificate_url](resources--cloud_credentials--reference--group-001.md#canonical-1201133212101231-2000023020230300-2031031320001203-2203202130133323-1003210222032330-1201202330132022-1023121322023203-1123231111331202) |
| `azure_pfx_certificate.client_id` | [azure_pfx_certificate.client_id](resources--cloud_credentials--reference--group-001.md#canonical-3233232101221223-1111301213220003-1203133120333103-1032320110333132-3211233132112021-0200031002033010-3222001012312103-3131031310030230) |
| `azure_pfx_certificate.password` | [azure_pfx_certificate.password](resources--cloud_credentials--reference--group-001.md#canonical-1200010231200102-0122030023023010-0003132232120000-3320120233333233-1200013110012322-2123012303022231-2333201302000031-1033321031222331) |
| `azure_pfx_certificate.password.blindfold_secret_info` | [azure_pfx_certificate.password.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-3132220323110112-2300011013313201-1000031230303102-1303213301220110-3300323321231030-3311131021002133-3303201201112212-0103333113110002) |
| `azure_pfx_certificate.password.blindfold_secret_info.decryption_provider` | [azure_pfx_certificate.password.blindfold_secret_info.decryption_provider](resources--cloud_credentials--reference--group-001.md#canonical-0300002031010302-0023022121232121-2303232203123103-2122310131332231-0233201320313121-1212332303313033-1002113302100330-0223122032002222) |
| `azure_pfx_certificate.password.blindfold_secret_info.location` | [azure_pfx_certificate.password.blindfold_secret_info.location](resources--cloud_credentials--reference--group-001.md#canonical-0110101101011013-2130102132120201-1310201313201003-0331020001110010-2331031322230022-1201312131012201-3230331312303022-0103321331022032) |
| `azure_pfx_certificate.password.blindfold_secret_info.store_provider` | [azure_pfx_certificate.password.blindfold_secret_info.store_provider](resources--cloud_credentials--reference--group-001.md#canonical-3033320120321322-1210230313321123-2022112323033000-3003031300122002-0131202232120232-1020030303321201-0101011332333001-3102331023202321) |
| `azure_pfx_certificate.password.clear_secret_info` | [azure_pfx_certificate.password.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-2032222012111023-0231230333022322-3003100100213213-2303120323030131-2321323031013031-0033000230123322-3013212203102203-0330030121132332) |
| `azure_pfx_certificate.password.clear_secret_info.provider_ref` | [azure_pfx_certificate.password.clear_secret_info.provider_ref](resources--cloud_credentials--reference--group-001.md#canonical-2103333132101111-1302232213001031-1330031001130103-3120231110201002-0303032331032013-0123333011010321-3311102120023321-0112322032133233) |
| `azure_pfx_certificate.password.clear_secret_info.url` | [azure_pfx_certificate.password.clear_secret_info.url](resources--cloud_credentials--reference--group-001.md#canonical-1333201220233320-3303203111101011-0220331010020322-1030111123111212-2021313020223110-1212220103011310-3211323110320300-0200220121100021) |
| `azure_pfx_certificate.subscription_id` | [azure_pfx_certificate.subscription_id](resources--cloud_credentials--reference--group-001.md#canonical-1112113030232320-0101110132303021-0223220320220223-0021203330330312-1212301002020113-1333221132230212-3022333311322233-3311020022301233) |
| `azure_pfx_certificate.tenant_id` | [azure_pfx_certificate.tenant_id](resources--cloud_credentials--reference--group-001.md#canonical-2311320113033321-1123130311321211-3030000102313013-1112220012222313-0221323001132112-2103120211202130-3002310322111102-3022221032300221) |
| `description` | [description](resources--cloud_credentials--reference--group-001.md#canonical-2232121111000120-0013123212310232-2322133013133331-0313023310302310-1233003312023123-1303331223222013-0220132201312131-1033100322331331) |
| `disable` | [disable](resources--cloud_credentials--reference--group-001.md#canonical-0203301030112020-0030002131022310-1132310313132101-3303121031210022-3301220022202223-1222232313121020-2020001302113323-1120322332321003) |
| `gcp_cred_file` | [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-3012311132000010-1330230001100002-0013100310223110-3022303300301001-1031323121122101-2221321012331122-3212110231123100-0320233111322012) |
| `gcp_cred_file.credential_file` | [gcp_cred_file.credential_file](resources--cloud_credentials--reference--group-001.md#canonical-0101022232132123-3213231310331013-0223102123202333-0121120230112020-0303003001311112-3312202031303003-0022132112302012-2003303122112313) |
| `gcp_cred_file.credential_file.blindfold_secret_info` | [gcp_cred_file.credential_file.blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-2313030313013221-2103023313211220-2322212321332013-3130321320333030-1002231013102130-3300330021010012-0012110030010200-1333011002212010) |
| `gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider](resources--cloud_credentials--reference--group-001.md#canonical-3132310123313211-3311000022000023-1220332233032112-1012122012200111-1201032021223012-0020303331022103-0303202003321221-3132121032213332) |
| `gcp_cred_file.credential_file.blindfold_secret_info.location` | [gcp_cred_file.credential_file.blindfold_secret_info.location](resources--cloud_credentials--reference--group-001.md#canonical-0013100131223330-0203000013303322-2233331300332200-2203011233011033-1212023211211333-2112310323011330-3323111323121323-0000133310000200) |
| `gcp_cred_file.credential_file.blindfold_secret_info.store_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.store_provider](resources--cloud_credentials--reference--group-001.md#canonical-3133221200220013-3203012120302002-3003020230331013-2302323000111201-1030222201221330-2103131323223102-1021233222100130-0131102100213210) |
| `gcp_cred_file.credential_file.clear_secret_info` | [gcp_cred_file.credential_file.clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-3021230102003023-3331013100030230-3332103033021110-3320231230220210-3232310233110003-3101222210212122-0021000322220322-1022103330131132) |
| `gcp_cred_file.credential_file.clear_secret_info.provider_ref` | [gcp_cred_file.credential_file.clear_secret_info.provider_ref](resources--cloud_credentials--reference--group-001.md#canonical-1312302310330120-0003201022332133-1003223022111201-1011022323302120-0311212131322133-3133331313023121-0201101120310031-3322233000301131) |
| `gcp_cred_file.credential_file.clear_secret_info.url` | [gcp_cred_file.credential_file.clear_secret_info.url](resources--cloud_credentials--reference--group-001.md#canonical-0103213320320111-0032202121121013-0333221030122221-3001032233003330-1210312113012003-0032032331323010-0221001013002232-2112130210210212) |
| `id` | [ID](resources--cloud_credentials--reference--group-001.md#canonical-1210321310122133-3033031310011332-0010221030002110-1101020012213230-1230102200313212-1221022100111330-2312331233100213-1201000033330130) |
| `labels` | [labels](resources--cloud_credentials--reference--group-001.md#canonical-0200203123003201-1333000223221011-1011013003231230-2130300033030321-2302002322130032-3232013310120131-0210123010211000-3233000333303303) |
| `name` | [name](resources--cloud_credentials--reference--group-001.md#canonical-3131110202313223-3202103021101230-1020122320103111-3003000102331013-2232212232020010-1031022023113032-2311100202211320-2130212120032323) |
| `namespace` | [namespace](resources--cloud_credentials--reference--group-001.md#canonical-2303113230130011-1101220112230020-0232133122022213-1222313022033112-1232231313011131-3113322133223111-0021032301213000-3230122213031220) |
| `timeouts` | [timeouts](resources--cloud_credentials--reference--group-001.md#canonical-1010302113011230-0312131112302202-3222303202323132-0312231112310033-3220113131003111-1021322010202313-1123231122200320-3011222333201333) |
| `timeouts.create` | [timeouts.create](resources--cloud_credentials--reference--group-001.md#canonical-3123331133012301-0002330101221221-2330320220003122-1110102132100033-1010023002223111-2133021323212303-1001021332211231-3301222213033103) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_credentials--reference--group-001.md#canonical-2030103223013332-0133021003211020-1313333102233232-3121313330010103-0131112323313331-0202120311110320-2100323222310202-3022003320023113) |
| `timeouts.read` | [timeouts.read](resources--cloud_credentials--reference--group-001.md#canonical-3110202320212103-1013210320110112-1203210112231110-0020010010120331-2201103100020232-1002323120213033-3103311223310120-1130220011302301) |
| `timeouts.update` | [timeouts.update](resources--cloud_credentials--reference--group-001.md#canonical-3132023003113101-3323231033231010-0013222113201121-0133323330332030-1303313233232233-3322013030311121-0121011000112013-0132022231032232) |

<a id="canonical-3201131310302023-0122011012123001-0130131130232113-3232133100213330-3012223312001231-3021312122320320-3313012123220011-3033311203311103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_assume_role` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- aws_assume_role

<a id="canonical-2010330202131221-3331033201131031-0033321330121033-3331321002230022-0130310102333320-1001313300031220-1233232231022131-0200333111010302"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws\_assume\_role, aws\_secret\_key, Azure\_client\_secret, Azure\_pfx\_certificate,
gcp\_cred\_file\] AWS Assume Role to Handle Delegated Access.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("duration_seconds",
    "role_arn",
    "session_name"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_optional"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_tenant_id"),
  validators.ConflictingObjectAttributes("external_id_is_optional",
    "external_id_is_tenant_id")}
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
  "x-ves-oneof-field-external_id": "[\"custom_external_id\",\"external_id_is_optional\",\"external_id_is_tenant_id\"]"
}
```

OneOf alternatives in this subsection:

- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-2010330202131221-3331033201131031-0033321330121033-3331321002230022-0130310102333320-1001313300031220-1233232231022131-0200333111010302)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-0023023223123233-0110233210102020-0012022122021003-1013311213210320-1032303101022312-1222331221202232-2313321000130101-3311302322131220)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-3310333023110001-0332313323233102-3201332102213110-0102330200133020-3313031333322311-3332010103031010-2121130001302120-2002203311022112)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-3133023000012002-3121021330033032-0030301001130301-1331333301221311-0031100002122123-1230122323130010-1012221100310221-2320101012131111)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-3012311132000010-1330230001100002-0013100310223110-3022303300301001-1031323121122101-2221321012331122-3212110231123100-0320233111322012)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_assume_role {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223112130023122-1203130210233123-3302101322313223-0232030121031221-3101120133000010-3103213002331202-0323212220212332-0132322230232010"></a>

### Direct properties for `aws_assume_role`

<a id="canonical-0102231300311031-0122201102322012-2210203132123110-1313102301303330-3033101203331110-1211212231130203-1103210112111013-1121202132223013"></a>

#### `aws_assume_role.custom_external_id` property

Type: `"string"`. Optional.

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  }
}
```

<a id="canonical-3200023330202312-0113030230310310-1232100021331231-3021320203101111-2110031210030211-0230030111203001-0110020130131321-2201023223203022"></a>

<a id="canonical-2212000300233232-0230130003123210-1211212023223110-1112320112031320-0300221033320203-1001202321120102-1322211301232203-2131213003001000"></a>

#### `aws_assume_role.duration_seconds` property

Type: `"number"`. Optional.

The duration, in seconds of the role session.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(3600, 43200),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

- [external_id_is_optional](resources--cloud_credentials--reference--group-001.md#canonical-3122230332231220-2323133033302121-2100003011232100-0331002310102220-2203130002323111-3100121320132111-0113231202303222-2220133223231210): complete subsection reference.

- [external_id_is_tenant_id](resources--cloud_credentials--reference--group-001.md#canonical-2203103211303300-3032003222210131-0131212122013230-0313303303201103-1023331120233333-1313220201120323-1032330303131223-0003332300111312): complete subsection reference.

<a id="canonical-0212313213223311-2123332300103330-3132110202101203-2003300232103320-2330200133011222-3132031021101001-0032231023101123-2323222232023213"></a>

<a id="canonical-0231011231120133-1333332322201130-1301122212312122-3012033321021310-1301132302211221-0003000131230330-2203333010200111-1103230321210023"></a>

#### `aws_assume_role.role_arn` property

Type: `"string"`. Optional.

IAM Role ARN. IAM Role ARN to assume the role.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 20,
    "pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  }
}
```

<a id="canonical-2020221000003010-0221123032120112-2332130221312121-1213202113311221-3131332032321102-2313301203033203-1211022003222321-3312210210313233"></a>

<a id="canonical-2232113123210001-0033032013203311-2223010111320023-0323132210320223-1000031000121332-3110100010113000-1203021211100012-2113333013310021"></a>

#### `aws_assume_role.session_name` property

Type: `"string"`. Optional.

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 2,
    "pattern": "[\\\\w+=,.@-]*"
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
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  }
}
```

<a id="canonical-0233100121322230-2300302322312301-0203110101221203-3110201331110200-2132103011113023-1220310202302300-1220103200320200-2230120013113312"></a>

<a id="canonical-0020003110032010-3212220000002322-0302313322223032-1211212103301330-3220303331022120-2211210313103132-0102211123220103-3223120123033020"></a>

#### `aws_assume_role.session_tags` property

Type: `["map", "string"]`. Optional.

Session tags are key-value pair attributes that you pass when you assume an IAM role.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":40},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":127,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"127\",\"ves.io.schema.rules.map.max_pairs\":\"40\",\"ves.io.schema.rules.map.values.string.max_len\":\"255\"},\"values\":{\"maxLength\":255,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 40
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 127,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "127",
      "ves.io.schema.rules.map.max_pairs": "40",
      "ves.io.schema.rules.map.values.string.max_len": "255"
    },
    "values": {
      "maxLength": 255,
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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-3122230332231220-2323133033302121-2100003011232100-0331002310102220-2203130002323111-3100121320132111-0113231202303222-2220133223231210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_assume_role.external_id_is_optional` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-3201131310302023-0122011012123001-0130131130232113-3232133100213330-3012223312001231-3021312122320320-3313012123220011-3033311203311103)
- aws_assume_role.external_id_is_optional

<a id="canonical-1310012012331200-2323001202310013-0022130002213230-3210122013300100-2012130033113103-1022113003301120-3100113230010032-3331231012022312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for external ID is optional.

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
external_id_is_optional = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203103211303300-3032003222210131-0131212122013230-0313303303201103-1023331120233333-1313220201120323-1032330303131223-0003332300111312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_assume_role.external_id_is_tenant_id` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [aws_assume_role](resources--cloud_credentials--reference--group-001.md#canonical-3201131310302023-0122011012123001-0130131130232113-3232133100213330-3012223312001231-3021312122320320-3313012123220011-3033311203311103)
- aws_assume_role.external_id_is_tenant_id

<a id="canonical-3201103211110100-3133011220303303-1301100321123220-3330100332111222-1033101210200103-1202230113131313-0323131311102002-3013323101232231"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
external_id_is_tenant_id = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233111313023123-2112001202221123-1003132111031121-1301010333300232-2201130032133313-3201203222322230-0230201100222323-2123010303221211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_secret_key` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- aws_secret_key

<a id="canonical-0023023223123233-0110233210102020-0012022122021003-1013311213210320-1032303101022312-1222331221202232-2313321000130101-3311302322131220"></a>

Type: `"object"`. single nested block, Optional.

AWS Programmatic Access Credentials type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("access_key")}
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
aws_secret_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101212303020032-2210330030332321-3101102203031312-3213012212333122-0232021032001312-2000020133030021-3312303003031101-0310102123230002"></a>

### Direct properties for `aws_secret_key`

<a id="canonical-2230230020012211-3331121000210100-3021201103320110-3232201032223201-3231303302300113-1213222031132113-3333030211213121-3213202320000301"></a>

#### `aws_secret_key.access_key` property

Type: `"string"`. Optional.

Access Key ID. Access key ID for your AWS account.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [secret_key](resources--cloud_credentials--reference--group-001.md#canonical-0213131110213223-0223232101023323-2123311133211101-3201320111130300-1021022120000213-3110212022203331-2313010131003223-2002021021331003): complete subsection reference.

<a id="canonical-0213131110213223-0223232101023323-2123311133211101-3201320111130300-1021022120000213-3110212022203331-2313010131003223-2002021021331003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_secret_key.secret_key` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-3233111313023123-2112001202221123-1003132111031121-1301010333300232-2201130032133313-3201203222322230-0230201100222323-2123010303221211)
- aws_secret_key.secret_key

<a id="canonical-1121333312013222-3221121111300031-1212332132010231-2131232122301001-0121323001130221-0210112010123033-3330112022010011-3013000320222312"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
secret_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103033110203301-3211300222222110-2033120031233100-1303013300022202-0202020312301110-1320313002010103-2311111302003120-3200022331321010"></a>

### Direct properties for `aws_secret_key.secret_key`

- [blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-3333201120120231-3223100301331100-0123323312200330-3203322313201130-3232222302032021-1100011301032210-0312133112310022-2321231001010313): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-0012032032232232-2020022112230222-0333130123120132-0300101302302230-3112111221011310-0201213011013012-0202122113030311-3103322202223021): complete subsection reference.

<a id="canonical-3333201120120231-3223100301331100-0123323312200330-3203322313201130-3232222302032021-1100011301032210-0312133112310022-2321231001010313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_secret_key.secret_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-3233111313023123-2112001202221123-1003132111031121-1301010333300232-2201130032133313-3201203222322230-0230201100222323-2123010303221211)
- [aws_secret_key.secret_key](resources--cloud_credentials--reference--group-001.md#canonical-0213131110213223-0223232101023323-2123311133211101-3201320111130300-1021022120000213-3110212022203331-2313010131003223-2002021021331003)
- aws_secret_key.secret_key.blindfold_secret_info

<a id="canonical-3113100211102011-1223331000230301-1200302130323133-0032200333021101-3030023213233331-2013202210330301-3311132332030030-1113130113321203"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1330322300203331-2133000123320130-1303000213112001-2210323311131322-2120323322031122-1023322103030323-3233300231301301-1020112033230021"></a>

### Direct properties for `aws_secret_key.secret_key.blindfold_secret_info`

<a id="canonical-3031113030302002-0112301213013322-0203232131112313-0230130000302320-0221200221212011-0023110230303011-3100133320100311-1131001310331220"></a>

#### `aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2302111031212132-0031230220013123-2132323001120101-3320023103220302-1133131110032301-1133202202333030-3223121121022030-2300312023302221"></a>

<a id="canonical-3121022210232020-0213023300103213-2322222222021212-3012100103320220-3310202323133330-1132231102120302-1120232231021031-0021010211122121"></a>

#### `aws_secret_key.secret_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2012321021313312-3133001003222102-0032300312320013-1023210013322010-0021310211201010-0120032002013303-3311030032113101-3132110231120123"></a>

<a id="canonical-2311320130033002-0003020111003203-0133220222102122-1133223223301101-2213212220101202-1011103111231320-0202230320200223-3032203031100312"></a>

#### `aws_secret_key.secret_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0012032032232232-2020022112230222-0333130123120132-0300101302302230-3112111221011310-0201213011013012-0202122113030311-3103322202223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws_secret_key.secret_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [aws_secret_key](resources--cloud_credentials--reference--group-001.md#canonical-3233111313023123-2112001202221123-1003132111031121-1301010333300232-2201130032133313-3201203222322230-0230201100222323-2123010303221211)
- [aws_secret_key.secret_key](resources--cloud_credentials--reference--group-001.md#canonical-0213131110213223-0223232101023323-2123311133211101-3201320111130300-1021022120000213-3110212022203331-2313010131003223-2002021021331003)
- aws_secret_key.secret_key.clear_secret_info

<a id="canonical-1111301233212100-0021202003332131-1032130021101320-0020131012000301-1122230022231303-0002320200033233-1010323001211331-1223323110230301"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1222222103100300-2201112002322003-1000010323023113-3030302311031220-2330133303201233-0130111323120010-3303201230223020-3232302003311032"></a>

### Direct properties for `aws_secret_key.secret_key.clear_secret_info`

<a id="canonical-1032222102221103-0113133123100002-3311221120230103-1032022300311211-1210221020100330-1132010012201222-2010121003200033-1232131100311312"></a>

#### `aws_secret_key.secret_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0101003111103022-3103331331132021-2310212111232112-0023311033100201-3121113123313223-1100201223133113-3232031012221000-1220230001011133"></a>

<a id="canonical-1210231130320003-2200101232223110-2230113001101022-1032312120201103-0031313000011332-1000312112302133-0121112023001330-2120023121210211"></a>

#### `aws_secret_key.secret_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0031003222133102-3220220230302231-3221203120103122-2132001113201132-2100013132100213-2210320123302223-3120332020133311-3031023322021323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_client_secret` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- azure_client_secret

<a id="canonical-3310333023110001-0332313323233102-3201332102213110-0102330200133020-3313031333322311-3332010103031010-2121130001302120-2002203311022112"></a>

Type: `"object"`. single nested block, Optional.

Azure Client Secret. Azure Credentials Client Secret type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("client_id",
    "subscription_id",
    "tenant_id")}
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
azure_client_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001233103303303-2033323210020310-2323310103101203-1222330213210320-0002231233100132-0310022012300303-2020000013332130-0331223320233113"></a>

### Direct properties for `azure_client_secret`

<a id="canonical-3030301330202312-1331013202232012-2001121322103031-1130100300023331-2211212301123200-3101120012031032-0111021102220120-0310110231011002"></a>

#### `azure_client_secret.client_id` property

Type: `"string"`. Optional.

Client ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [client_secret](resources--cloud_credentials--reference--group-001.md#canonical-2010100203212332-1110203230303022-1312013232030032-0311031313033313-3311310202030013-2320110323223123-3220210021311103-0203100022211200): complete subsection reference.

<a id="canonical-2033231202231230-2002211332232300-0122223320123130-3001232001013011-0322323320221030-3023111102112230-3120201121311131-3021231003030003"></a>

<a id="canonical-2212100311332331-3032232210323130-3113303302312313-1323112113022011-2221133223330102-3132302233331213-3202133121203300-2103012022102101"></a>

#### `azure_client_secret.subscription_id` property

Type: `"string"`. Optional.

Subscription ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0321132230203232-0010233002310212-1100203330031131-2222220033123332-0120112030002212-0323220023112002-2230132333133011-3103311222223030"></a>

<a id="canonical-2103202330113201-3331023232312200-1330033313133333-0331230020101100-1302121000222032-3033323121100211-0023113312213033-3021022201332112"></a>

#### `azure_client_secret.tenant_id` property

Type: `"string"`. Optional.

Tenant ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2010100203212332-1110203230303022-1312013232030032-0311031313033313-3311310202030013-2320110323223123-3220210021311103-0203100022211200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_client_secret.client_secret` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0031003222133102-3220220230302231-3221203120103122-2132001113201132-2100013132100213-2210320123302223-3120332020133311-3031023322021323)
- azure_client_secret.client_secret

<a id="canonical-0102103103330102-0203333133232212-2201133302220200-3030003320302132-2230220033203312-3011013013312323-2112000201103302-1111001112123233"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0013320212113203-0102112031131123-3010033001212210-3003301133100111-3123211023213310-2201023132020222-2013132332133002-2101310020330030"></a>

### Direct properties for `azure_client_secret.client_secret`

- [blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-0130201031022210-0211013222301100-0333222111233120-1321331002131100-0033033103302122-3113302012123222-2133200113023021-1102333301021330): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-1113001302011012-3122110020131131-0321023013110311-0023333123113132-2322023320021101-2331020331232120-0000331030003000-1201201131131211): complete subsection reference.

<a id="canonical-0130201031022210-0211013222301100-0333222111233120-1321331002131100-0033033103302122-3113302012123222-2133200113023021-1102333301021330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_client_secret.client_secret.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0031003222133102-3220220230302231-3221203120103122-2132001113201132-2100013132100213-2210320123302223-3120332020133311-3031023322021323)
- [azure_client_secret.client_secret](resources--cloud_credentials--reference--group-001.md#canonical-2010100203212332-1110203230303022-1312013232030032-0311031313033313-3311310202030013-2320110323223123-3220210021311103-0203100022211200)
- azure_client_secret.client_secret.blindfold_secret_info

<a id="canonical-3301322210033322-2201022330333230-1303323121101113-2100010312221310-2320200100101122-1132111112223130-3122131003223030-3000012302330103"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2212023032010323-3213130101103002-3333011023312000-2003200301110200-2123221121132130-0120331203300203-3010330112132331-0303222031331230"></a>

### Direct properties for `azure_client_secret.client_secret.blindfold_secret_info`

<a id="canonical-1103001023300112-0202210101031310-3322212030212203-1113222123313000-1113331211122220-2213100011022330-2000110020223013-3030311311101132"></a>

#### `azure_client_secret.client_secret.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0002101220003210-3212111001033211-2311210011021301-0030310100011033-3223103320323221-3122110303130203-1100201311303122-2332120233012232"></a>

<a id="canonical-0100232001223230-2330311313201333-2120120001322302-1021210202321232-0212303121321030-1222210001133212-3032032022211101-3032002110210312"></a>

#### `azure_client_secret.client_secret.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3201330230112131-1001333012303303-3032122023110311-3031100032332002-0230331312023003-0333030222310011-0033010300101202-0212321330203221"></a>

<a id="canonical-1131031211111101-1221122200202110-1210121303221132-1132023010231301-2113110111221321-1201322123222100-2221003333211330-3130010101030123"></a>

#### `azure_client_secret.client_secret.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1113001302011012-3122110020131131-0321023013110311-0023333123113132-2322023320021101-2331020331232120-0000331030003000-1201201131131211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_client_secret.client_secret.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [azure_client_secret](resources--cloud_credentials--reference--group-001.md#canonical-0031003222133102-3220220230302231-3221203120103122-2132001113201132-2100013132100213-2210320123302223-3120332020133311-3031023322021323)
- [azure_client_secret.client_secret](resources--cloud_credentials--reference--group-001.md#canonical-2010100203212332-1110203230303022-1312013232030032-0311031313033313-3311310202030013-2320110323223123-3220210021311103-0203100022211200)
- azure_client_secret.client_secret.clear_secret_info

<a id="canonical-3331102123301303-1220210000210032-1103010122103130-1313122303011010-0110022223232112-3302011002112022-1132303122231233-1032231211013013"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0230321232111110-0103013222331211-3022113301322131-0232321331331000-0003200131131102-0001203012311113-1303200222212121-2012110323131112"></a>

### Direct properties for `azure_client_secret.client_secret.clear_secret_info`

<a id="canonical-2113320313222131-3303013302101120-3000332330110111-1121223112002130-1121321202200331-3112031213130120-0012103013210121-0000323221022023"></a>

#### `azure_client_secret.client_secret.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2220202031012302-2001321101231003-0213102333103110-2222012202133031-1231102200031321-1111230330313323-1210132131001123-3203133101120132"></a>

<a id="canonical-0221202020022013-2213012302022003-1111230333022210-0000230201112030-2202111113022311-3310320111113203-1321303202232310-0112130102111200"></a>

#### `azure_client_secret.client_secret.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3121121300221320-0323123210301120-3332103121012010-0133333031111230-3231111231030100-0102030303212231-0002210222100033-0213220221123010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_pfx_certificate` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- azure_pfx_certificate

<a id="canonical-3133023000012002-3121021330033032-0030301001130301-1331333301221311-0031100002122123-1230122323130010-1012221100310221-2320101012131111"></a>

Type: `"object"`. single nested block, Optional.

Azure Credentials Client Certificate type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("certificate_url",
    "client_id",
    "subscription_id",
    "tenant_id")}
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
azure_pfx_certificate {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011222233121122-1301313202220212-3210232031312323-1023103120122002-3221232131000232-2102221131020313-3322231321031230-1311112232030032"></a>

### Direct properties for `azure_pfx_certificate`

<a id="canonical-1201133212101231-2000023020230300-2031031320001203-2203202130133323-1003210222032330-1201202330132022-1023121322023203-1123231111331202"></a>

#### `azure_pfx_certificate.certificate_url` property

Type: `"string"`. Optional.

URL for Client Certificate in '.pfx' or '.p12' whose certificate is linked to service principal
object Certificate URL can contain client certificate in string:///&lt;base64 of certificate&gt;
format. Here &lt;base64 of certificate&gt; is base64 of '.pfx' or '.p12' binary file.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3233232101221223-1111301213220003-1203133120333103-1032320110333132-3211233132112021-0200031002033010-3222001012312103-3131031310030230"></a>

<a id="canonical-1023303233221323-0232020120013233-3302122300031222-3132330310130011-2303013213302120-3102011032323301-3030120013332231-3311022303320331"></a>

#### `azure_pfx_certificate.client_id` property

Type: `"string"`. Optional.

Client ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [password](resources--cloud_credentials--reference--group-001.md#canonical-2031130313303332-1313101222131123-2013331003223330-0210331002102030-2321321100123023-1122332123000113-0213011110303113-0100122132020231): complete subsection reference.

<a id="canonical-1112113030232320-0101110132303021-0223220320220223-0021203330330312-1212301002020113-1333221132230212-3022333311322233-3311020022301233"></a>

<a id="canonical-0002202020120223-0312310330311300-1220000021220211-2002130130113010-2023332311212231-1200012113002301-0203100101321012-1002120121330211"></a>

#### `azure_pfx_certificate.subscription_id` property

Type: `"string"`. Optional.

Subscription ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2311320113033321-1123130311321211-3030000102313013-1112220012222313-0221323001132112-2103120211202130-3002310322111102-3022221032300221"></a>

<a id="canonical-2232231013103032-1331220330212002-2301112111112312-2100233202220332-1320220301322200-2230103301001332-3230211222001102-2131113222233011"></a>

#### `azure_pfx_certificate.tenant_id` property

Type: `"string"`. Optional.

Tenant ID for your Azure service principal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2031130313303332-1313101222131123-2013331003223330-0210331002102030-2321321100123023-1122332123000113-0213011110303113-0100122132020231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_pfx_certificate.password` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-3121121300221320-0323123210301120-3332103121012010-0133333031111230-3231111231030100-0102030303212231-0002210222100033-0213220221123010)
- azure_pfx_certificate.password

<a id="canonical-1200010231200102-0122030023023010-0003132232120000-3320120233333233-1200013110012322-2123012303022231-2333201302000031-1033321031222331"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0111332320021020-3013321032121131-2321132020132303-3100012221102123-3222221320102300-1221322121331111-3221132111301113-1010111003303122"></a>

### Direct properties for `azure_pfx_certificate.password`

- [blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-1310001113023013-3001022212223113-2120100110203131-2030210221001020-3020220111302020-1331111210101031-0030122310010200-0333302212101311): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-0020203022020331-0102202103100332-3301232132332130-1331010013313123-0233022330321020-2210302021003012-3100021220023210-3002203232332002): complete subsection reference.

<a id="canonical-1310001113023013-3001022212223113-2120100110203131-2030210221001020-3020220111302020-1331111210101031-0030122310010200-0333302212101311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_pfx_certificate.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-3121121300221320-0323123210301120-3332103121012010-0133333031111230-3231111231030100-0102030303212231-0002210222100033-0213220221123010)
- [azure_pfx_certificate.password](resources--cloud_credentials--reference--group-001.md#canonical-2031130313303332-1313101222131123-2013331003223330-0210331002102030-2321321100123023-1122332123000113-0213011110303113-0100122132020231)
- azure_pfx_certificate.password.blindfold_secret_info

<a id="canonical-3132220323110112-2300011013313201-1000031230303102-1303213301220110-3300323321231030-3311131021002133-3303201201112212-0103333113110002"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3021123002132200-3101313203233012-1202032332212101-2311000211323231-1113030121012123-2012122021212133-0133313102321200-2110033012230301"></a>

### Direct properties for `azure_pfx_certificate.password.blindfold_secret_info`

<a id="canonical-0300002031010302-0023022121232121-2303232203123103-2122310131332231-0233201320313121-1212332303313033-1002113302100330-0223122032002222"></a>

#### `azure_pfx_certificate.password.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0110101101011013-2130102132120201-1310201313201003-0331020001110010-2331031322230022-1201312131012201-3230331312303022-0103321331022032"></a>

<a id="canonical-1003312210002212-0133010212033230-0111230002021113-3100102012223130-2203021220221032-2111301302120010-2212112033102003-2123113120231321"></a>

#### `azure_pfx_certificate.password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3033320120321322-1210230313321123-2022112323033000-3003031300122002-0131202232120232-1020030303321201-0101011332333001-3102331023202321"></a>

<a id="canonical-1202110302331223-1301110312330223-3122200020322320-2132203022320211-3021212133112211-3203231321001022-2103121220332130-3031211132032010"></a>

#### `azure_pfx_certificate.password.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0020203022020331-0102202103100332-3301232132332130-1331010013313123-0233022330321020-2210302021003012-3100021220023210-3002203232332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `azure_pfx_certificate.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [azure_pfx_certificate](resources--cloud_credentials--reference--group-001.md#canonical-3121121300221320-0323123210301120-3332103121012010-0133333031111230-3231111231030100-0102030303212231-0002210222100033-0213220221123010)
- [azure_pfx_certificate.password](resources--cloud_credentials--reference--group-001.md#canonical-2031130313303332-1313101222131123-2013331003223330-0210331002102030-2321321100123023-1122332123000113-0213011110303113-0100122132020231)
- azure_pfx_certificate.password.clear_secret_info

<a id="canonical-2032222012111023-0231230333022322-3003100100213213-2303120323030131-2321323031013031-0033000230123322-3013212203102203-0330030121132332"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3132223112211132-3201000110201032-1031030120223103-3333103030110033-1230203010232001-0300113000321023-1130210233210222-1310310312313221"></a>

### Direct properties for `azure_pfx_certificate.password.clear_secret_info`

<a id="canonical-2103333132101111-1302232213001031-1330031001130103-3120231110201002-0303032331032013-0123333011010321-3311102120023321-0112322032133233"></a>

#### `azure_pfx_certificate.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1333201220233320-3303203111101011-0220331010020322-1030111123111212-2021313020223110-1212220103011310-3211323110320300-0200220121100021"></a>

<a id="canonical-2123021222003013-0323303312310201-2323132102001232-2013012002230210-1303110101312011-2332211110331023-2101330310121120-0102113211033121"></a>

#### `azure_pfx_certificate.password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3222330220322011-0302130220202222-3211013011122330-0030301333103222-2332221013220321-3331320310312200-2132212011130212-0002121112103213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_cred_file` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- gcp_cred_file

<a id="canonical-3012311132000010-1330230001100002-0013100310223110-3022303300301001-1031323121122101-2221321012331122-3212110231123100-0320233111322012"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for gcp cred file.

Additional upstream details:

GCP Credentials type.

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
gcp_cred_file {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301313203001210-0313211131011320-2322013211233003-3230331331130300-2101003013023220-1113011033213001-0130322230311113-3031221230022012"></a>

### Direct properties for `gcp_cred_file`

- [credential_file](resources--cloud_credentials--reference--group-001.md#canonical-1030130012133333-0222122201313312-1201132103211113-1330031302032020-2323202310330000-3000120330331113-2003212222202203-0133001000320110): complete subsection reference.

<a id="canonical-1030130012133333-0222122201313312-1201132103211113-1330031302032020-2323202310330000-3000120330331113-2003212222202203-0133001000320110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_cred_file.credential_file` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-3222330220322011-0302130220202222-3211013011122330-0030301333103222-2332221013220321-3331320310312200-2132212011130212-0002121112103213)
- gcp_cred_file.credential_file

<a id="canonical-0101022232132123-3213231310331013-0223102123202333-0121120230112020-0303003001311112-3312202031303003-0022132112302012-2003303122112313"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
credential_file {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211213111132101-1103112022202132-2112122000332123-3211130030021312-2013002031020320-2021132333013230-3201322102121113-3120023031213111"></a>

### Direct properties for `gcp_cred_file.credential_file`

- [blindfold_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-3001331110213211-0011102121131001-3030222021121213-1211203220232221-3120303000001110-3300020303331002-1222023101003301-2113023232301120): complete subsection reference.

- [clear_secret_info](resources--cloud_credentials--reference--group-001.md#canonical-0223211232323111-1121032210201033-0130202123010220-1003132111331311-2212010330330210-0020333211101222-1322201333022123-2010112332112320): complete subsection reference.

<a id="canonical-3001331110213211-0011102121131001-3030222021121213-1211203220232221-3120303000001110-3300020303331002-1222023101003301-2113023232301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_cred_file.credential_file.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-3222330220322011-0302130220202222-3211013011122330-0030301333103222-2332221013220321-3331320310312200-2132212011130212-0002121112103213)
- [gcp_cred_file.credential_file](resources--cloud_credentials--reference--group-001.md#canonical-1030130012133333-0222122201313312-1201132103211113-1330031302032020-2323202310330000-3000120330331113-2003212222202203-0133001000320110)
- gcp_cred_file.credential_file.blindfold_secret_info

<a id="canonical-2313030313013221-2103023313211220-2322212321332013-3130321320333030-1002231013102130-3300330021010012-0012110030010200-1333011002212010"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3302021121031321-3122323100023311-1202103213013223-1233023113231333-3203133132002322-0201231012021200-0102012313200120-2213313312012211"></a>

### Direct properties for `gcp_cred_file.credential_file.blindfold_secret_info`

<a id="canonical-3132310123313211-3311000022000023-1220332233032112-1012122012200111-1201032021223012-0020303331022103-0303202003321221-3132121032213332"></a>

#### `gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0013100131223330-0203000013303322-2233331300332200-2203011233011033-1212023211211333-2112310323011330-3323111323121323-0000133310000200"></a>

<a id="canonical-3303110021332332-0311110303003132-3223000310110103-1131232111021311-1332202323033311-2230331000330201-3033211323120003-0202301311232300"></a>

#### `gcp_cred_file.credential_file.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3133221200220013-3203012120302002-3003020230331013-2302323000111201-1030222201221330-2103131323223102-1021233222100130-0131102100213210"></a>

<a id="canonical-2201201311122211-3001010033111111-1100132330112113-2112210202001313-0120213222113031-0220322002132132-0003002123320030-3310230201332223"></a>

#### `gcp_cred_file.credential_file.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0223211232323111-1121032210201033-0130202123010220-1003132111331311-2212010330330210-0020333211101222-1322201333022123-2010112332112320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_cred_file.credential_file.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [gcp_cred_file](resources--cloud_credentials--reference--group-001.md#canonical-3222330220322011-0302130220202222-3211013011122330-0030301333103222-2332221013220321-3331320310312200-2132212011130212-0002121112103213)
- [gcp_cred_file.credential_file](resources--cloud_credentials--reference--group-001.md#canonical-1030130012133333-0222122201313312-1201132103211113-1330031302032020-2323202310330000-3000120330331113-2003212222202203-0133001000320110)
- gcp_cred_file.credential_file.clear_secret_info

<a id="canonical-3021230102003023-3331013100030230-3332103033021110-3320231230220210-3232310233110003-3101222210212122-0021000322220322-1022103330131132"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3301302223303202-2331231333122232-1001222300131013-1210320312031032-0230020333113202-1130310323032130-3301322332231223-1300213012321110"></a>

### Direct properties for `gcp_cred_file.credential_file.clear_secret_info`

<a id="canonical-1312302310330120-0003201022332133-1003223022111201-1011022323302120-0311212131322133-3133331313023121-0201101120310031-3322233000301131"></a>

#### `gcp_cred_file.credential_file.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0103213320320111-0032202121121013-0333221030122221-3001032233003330-1210312113012003-0032032331323010-0221001013002232-2112130210210212"></a>

<a id="canonical-1321030311123300-0022032031132000-1120233113001111-2303322022321312-2012121011313230-2323102133123223-3112130010122000-1121013123103213"></a>

#### `gcp_cred_file.credential_file.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0322303012302112-3031111001122203-3201002312132113-1301113100000020-0130001003100221-3021000310323011-1230333102123130-3302111233300121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Property reference](resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- timeouts

<a id="canonical-1010302113011230-0312131112302202-3222303202323132-0312231112310033-3220113131003111-1021322010202313-1123231122200320-3011222333201333"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103012300111310-0303321322131233-1120322123321110-3022323020030201-2133211231333221-1312230000310133-3322300123211200-3000221020202302"></a>

### Direct properties for `timeouts`

<a id="canonical-3123331133012301-0002330101221221-2330320220003122-1110102132100033-1010023002223111-2133021323212303-1001021332211231-3301222213033103"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2030103223013332-0133021003211020-1313333102233232-3121313330010103-0131112323313331-0202120311110320-2100323222310202-3022003320023113"></a>

<a id="canonical-3320001023131122-1100023213323220-1321033303203232-3201303130100312-3212013123302111-3120133213001001-2132301223313310-1201311030332332"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3110202320212103-1013210320110112-1203210112231110-0020010010120331-2201103100020232-1002323120213033-3103311223310120-1130220011302301"></a>

<a id="canonical-3002021011332120-0211122102303212-3001212120010012-1112331210100121-2021322323030120-2211330211201312-2100322023103230-1023001201332020"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3132023003113101-3323231033231010-0013222113201121-0133323330332030-1303313233232233-3322013030311121-0121011000112013-0132022231032232"></a>

<a id="canonical-3032033320012223-3110023312203301-1323201200223333-1022033131323322-0021231110001132-3003332302211233-0320013212233201-2200000030101102"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
