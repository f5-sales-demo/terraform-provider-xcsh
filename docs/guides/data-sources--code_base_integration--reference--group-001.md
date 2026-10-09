---
page_title: "xcsh_code_base_integration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration reference."
---

# xcsh_code_base_integration reference

<a id="canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- Property reference

<a id="canonical-0121022113323310-1320301323013013-3102312010331103-1003010233032021-2130233023110021-0012101303033102-3131102003201002-0122331330233013"></a>

### Direct properties for `xcsh_code_base_integration`

<a id="canonical-3112001020303122-0310300001211213-1320120331231111-1302111312011130-0020222323001231-1000202223022001-0023320031302113-0300113110032103"></a>

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

- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122): complete subsection reference.

<a id="canonical-3033100213003330-3230013303013003-2003030102313212-1303210131002112-0331331222000213-3303000303023122-3031331213133121-3000301231311023"></a>

<a id="canonical-3331202322132320-1011223233010130-3000323320313131-1002320133000332-1320002232313022-3112111011002023-3000101203231000-1023131200303200"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the CodeBaseIntegration.

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

<a id="canonical-2330101111131302-1330110132121202-1100031030303233-3130112321000031-0322232102021321-0312232321033113-3220312303000312-3022233302231000"></a>

<a id="canonical-3113230233022203-2303213013312323-1321233112130100-0203112103313133-2321232113010311-0001213010230032-3313213202030033-3320303131331113"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1211103333122103-2221121011303331-3112101233211300-3020300232202213-0113200312103023-0222330002131200-0121003112233222-0211030221301320"></a>

<a id="canonical-1213020221001311-3023232031110231-1331312020021122-1222212210122120-2200312333232102-1203220021012222-0312011030300311-3210032132101210"></a>

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

<a id="canonical-1300221212302001-1000213011121030-0321032323132330-3122111210331313-2231231110031223-2020123003011102-0101303100232132-2023333130212331"></a>

<a id="canonical-3012002323121123-3023231132020203-3210313113131102-3011233111111003-3133321233213103-3220121323223332-3332222133230221-0033111322331210"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CodeBaseIntegration.

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

<a id="canonical-3320002100120300-2213021210221331-3121010301302212-2223101121012030-0102001000233033-3031201100031103-1120101103230013-2031022120111000"></a>

<a id="canonical-3122220303331312-2313031331013331-2110320321032213-3003023213220033-3211123010001220-2332033001313332-3311202133011200-3120003121111203"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CodeBaseIntegration exists.

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

<a id="canonical-1201331002311013-2312232111331123-3120033323300330-1233131310132011-1002121331132003-1323113020233110-3000130102223221-0023333232222010"></a>

### All schema paths for `xcsh_code_base_integration`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--code_base_integration--reference--group-001.md#canonical-3112001020303122-0310300001211213-1320120331231111-1302111312011130-0020222323001231-1000202223022001-0023320031302113-0300113110032103) |
| `code_base_integration` | [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1010013100022000-2331302120111003-3230313001232321-3331311313301223-2012313003210112-2121311111203222-2300312033130132-3101331223002302) |
| `code_base_integration.azure_repos` | [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-0202130131131210-0010233112112311-0012131033110100-3000033013130230-3322303200331101-3012011133121221-3203220301322220-0031212103121011) |
| `code_base_integration.azure_repos.access_token` | [code_base_integration.azure_repos.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-3223313110122200-2223113112320210-1132312320133201-1312011231033112-1200321003020330-0332233310302231-1013322302121320-2333202012321322) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info` | [code_base_integration.azure_repos.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-3330132102012010-3311222213333200-1112310023013112-3000002231110100-1110303310110011-3211303321000102-3322002113332033-0122213000013232) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-3033031213301131-0133132300132203-3223112201232100-1231321101301330-3132022303213130-3330010303103001-0110003202003122-2303310221312031) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.location` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-1012103013003232-3110310333000212-0210022230202020-3301021101331330-1021301033013130-2121303201313333-2321111023313333-2132312113131321) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-1112233211301030-1203320210002320-1222102121201330-1300230002103303-2331310300010232-2112231111123121-1231110210130331-1010011102301112) |
| `code_base_integration.azure_repos.access_token.clear_secret_info` | [code_base_integration.azure_repos.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-0301211332122013-1103031333112031-0111321202330211-1322332123012313-2131332111030022-0323230133120221-0230130113213332-0121003201133301) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref` | [code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-0101102222223220-0020120131100303-3223031312302111-0230101303001331-2111111122003210-0133222120020023-2232221222012302-3013112323212211) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.url` | [code_base_integration.azure_repos.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-3222020132201113-1120322013012111-0323220331230033-1100131303311100-0021030301103000-1111213232330313-3131130003033322-3032312023321211) |
| `code_base_integration.bitbucket` | [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-0211321132030210-0030302103320323-2233001201210011-3333020331033230-2123103302123311-0200231003203301-3230003133132311-2003133212102033) |
| `code_base_integration.bitbucket.passwd` | [code_base_integration.bitbucket.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-1103202030320000-3200010132112102-2001220020010231-1110211003233130-3331321010201111-2321333222200313-2323003223012311-1222212231122120) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info` | [code_base_integration.bitbucket.passwd.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-2022000113022010-0023203131323030-3331122101121230-1121032132312302-3012000200230211-1003103323332023-1010111121113001-1032333112223130) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-1030203010123020-0033102031230102-0131231211023220-0330133333030223-1320122223233121-0032211102232313-3032100321331220-1300233221211132) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-3302211001323032-3031103003332001-3320322302223002-1222332300313110-0002221031210231-0202132030113001-1210131232301211-3323202133210310) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-2001112313231232-1002121031222033-3031130133023001-0302203331212012-0201011013303212-1213003113100303-0102331211123020-1223012130133333) |
| `code_base_integration.bitbucket.passwd.clear_secret_info` | [code_base_integration.bitbucket.passwd.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1032222211201223-0102321331223113-2101212223113223-3133121103000231-3223320322201233-2022101023310032-2012122311303232-0013210231001120) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-3030121022131122-0020222313031130-3323103330302032-0201322210211132-3100202301122230-3003232130101032-1323103232301303-3222102323123020) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.url` | [code_base_integration.bitbucket.passwd.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-3223302302212002-2030312200013303-1021230313200003-3233000120001001-0022032312201132-2302103002112300-1022303131133333-1212311332131312) |
| `code_base_integration.bitbucket.username` | [code_base_integration.bitbucket.username](data-sources--code_base_integration--reference--group-001.md#canonical-0212210122321313-2222131103103131-3223201111330213-1311313032332321-1021333100012100-0032320323322020-2301110310022131-2021330223001231) |
| `code_base_integration.bitbucket_server` | [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-1213013313131331-2132223200313111-2120213010302210-1310302303301330-1103311112203111-3212131330031202-1011101232023133-1103101113311232) |
| `code_base_integration.bitbucket_server.passwd` | [code_base_integration.bitbucket_server.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-2213030111203132-3101001032211031-0220013300312032-0101001031321022-1010221322203202-3111130021222310-1130100120023020-1103222231031220) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-0210303333122013-0031011322311030-2120110112102213-2330201000210121-0210331100231120-0123213212003131-2033333133333223-2231031013121111) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-2313132130221022-1232201012110301-3012331333121322-1310203123220002-0020321330331122-0312303110313322-3223023323000231-1002200121013120) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-1332300132221310-0030001331120130-2033310013230001-1113020133003131-0232002311121312-2322132102233122-0230211103331323-2331130000203302) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-1022113122313211-0020133133023230-2000003220221130-2111303330201312-0202133311300312-0023320023031200-3101001311101322-0100003123213130) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info` | [code_base_integration.bitbucket_server.passwd.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1001030300121102-3213113203320222-0302131023023101-0033110301030033-2320123212223230-3202021211310000-3022210123101301-0111110133323201) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-2313331310011021-2330330100022321-0320202331112303-2012312223020200-3012313310311033-2231313203033111-2300232011300011-3131103231001103) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.url` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-1311120301122100-2010000101030221-2011011011300033-2233121220203033-2111022131310320-0300011132120231-0210002130302111-0100011301211203) |
| `code_base_integration.bitbucket_server.url` | [code_base_integration.bitbucket_server.url](data-sources--code_base_integration--reference--group-001.md#canonical-3313323020101033-1220131032001230-1203202201022322-3000311131321320-3000233112012023-1112130312332230-3231330213210233-2003330232103312) |
| `code_base_integration.bitbucket_server.username` | [code_base_integration.bitbucket_server.username](data-sources--code_base_integration--reference--group-001.md#canonical-0333333323011100-1223210000012210-0100122001201101-3330022233220202-0101230002120330-0011321023021232-3021103332030313-3002300100330220) |
| `code_base_integration.bitbucket_server.verify_ssl` | [code_base_integration.bitbucket_server.verify_ssl](data-sources--code_base_integration--reference--group-001.md#canonical-2211303113033033-2303220113122110-3201130212002010-3130130321021200-3023212022232203-3212133230112012-1203023003302323-1120123302110231) |
| `code_base_integration.github` | [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-3033202230113000-1113223010021122-3333310101332220-1332212013230101-0101212111133203-2010331330210201-3312001231112212-1121100022131312) |
| `code_base_integration.github.access_token` | [code_base_integration.github.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0121010303222003-0212012030131202-3200231013121223-1223302212001131-3303122300023230-0200100321002002-0200233103121001-2312310331211211) |
| `code_base_integration.github.access_token.blindfold_secret_info` | [code_base_integration.github.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-3101200132210123-0233300121033202-2321032330010210-0321022110220202-3022102032211311-1333212030103021-1311011032003222-0133113101030113) |
| `code_base_integration.github.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-2100212323031132-1313130302013123-1311130010222302-0220302032203011-0001320132011211-3322321332112113-3313103213331210-0320312032232121) |
| `code_base_integration.github.access_token.blindfold_secret_info.location` | [code_base_integration.github.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-3110321202111300-1111231130332022-1130200321100033-1013212023001103-0322312201133323-3321132310333330-3322012001132200-0302201033022013) |
| `code_base_integration.github.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-3132112131210020-2130012330120312-0132212013221201-0302211220222120-0022130232333012-1333103203320332-2002002012223113-3112021012000313) |
| `code_base_integration.github.access_token.clear_secret_info` | [code_base_integration.github.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1210102330033233-0113200023211111-0311100232030001-2111033201331131-3323311203022333-3221201002322222-1331231121132011-2001303203122010) |
| `code_base_integration.github.access_token.clear_secret_info.provider_ref` | [code_base_integration.github.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-1333111302130120-2002202022033120-3021211032013123-0231003330302320-2303120313022032-0020233310303302-1120123202322232-2333220113331112) |
| `code_base_integration.github.access_token.clear_secret_info.url` | [code_base_integration.github.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-2302120320310221-2000011002201212-3131220122101311-2120232303102202-3320221122301310-3021033311021000-0233013302223133-1120312100200202) |
| `code_base_integration.github.username` | [code_base_integration.github.username](data-sources--code_base_integration--reference--group-001.md#canonical-3101001321301121-0323013313111230-0123211220123133-0133232303112332-1223201111213012-2103131212221110-0111122301121331-1003233313131300) |
| `code_base_integration.github.verify_ssl` | [code_base_integration.github.verify_ssl](data-sources--code_base_integration--reference--group-001.md#canonical-0313132122130311-2021102120001011-2322020012212030-1103212113330231-3003300110202230-3011300300311201-1101211120230212-1323131120322003) |
| `code_base_integration.github_enterprise` | [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-0311010301031223-3203230023012031-1333331213221212-2311010121230212-0101300331233133-1021201203030001-0023102031312232-2220221010123212) |
| `code_base_integration.github_enterprise.access_token` | [code_base_integration.github_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0003121020013131-3102222300013231-3022330111312000-2013220121003030-0002020002332333-0331202301203330-2330211112021202-2202330132121223) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-2130030303023122-0100002030202013-1202230221233233-3221231131101131-0020012211031132-0100012032313232-1201231220023200-1311002311030211) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-1313002011133120-3112031131113200-1130020003230023-0230330000130011-3033310013310012-0100211300101233-0021100203230231-3322222211113302) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-1203301103310311-1012010123220111-2210222002000303-0001321112103320-1020220303322201-2102310030231010-3000023202020000-3233213303112311) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-3220123330113223-2132023312213321-2110002103313331-0012211002010231-1021213033123012-1130300232232213-1213232001100132-2211230213031131) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info` | [code_base_integration.github_enterprise.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1232130201211313-2131330220120112-1022121132231200-3203011003123311-3131023102123101-0231120332101233-1300022203203220-2010132002011102) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-2202203221103013-1022302031121111-3201130100021103-3212212200222013-0210101300331312-1120233320201011-3301132100221232-2121230221111201) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.url` | [code_base_integration.github_enterprise.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-3100103132103033-2102011231113303-2121122111122010-1333303110233032-3302302121111210-2321330232312211-1303011013232021-0033110033110223) |
| `code_base_integration.github_enterprise.hostname` | [code_base_integration.github_enterprise.hostname](data-sources--code_base_integration--reference--group-001.md#canonical-0130331330312023-2323110010331001-1103231301223131-1230232022221102-2233333131023101-1203130012131011-0210102103002200-1002333333000231) |
| `code_base_integration.github_enterprise.username` | [code_base_integration.github_enterprise.username](data-sources--code_base_integration--reference--group-001.md#canonical-0210011022302203-3132233030030033-0131131001223233-1302311102121332-2010003331101203-1103101121301221-3330310223013212-0112231232311110) |
| `code_base_integration.gitlab` | [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-2313001223332131-0033312233010000-0121123132202130-3322330013232232-2312032120203013-2320130303010123-2002201121201011-2110330231023031) |
| `code_base_integration.gitlab.access_token` | [code_base_integration.gitlab.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-2031030002313130-1111130033130321-2100332332100013-3200210023312220-2110222100002032-3212322112312002-2321311033121322-2201133103121123) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info` | [code_base_integration.gitlab.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-3000210213110323-3120210111222012-2301000010020131-2203212200010203-0313010001122031-0121032200220032-0020311233302102-3120033010012312) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-1333313211110322-1212133121033332-0300030121222001-0111321203302321-0003100131301100-2210302101231112-1331030030011200-1301200232230332) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-0101222232233032-0102110300322031-3202313312333000-1131210330020103-3311312312232300-0203222100130001-0203323310321210-2332030333232121) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-1203101033103322-2323230222021233-1133310231133311-2032011232300112-2233132133101220-0221300022002231-2132212020200021-0021031301330110) |
| `code_base_integration.gitlab.access_token.clear_secret_info` | [code_base_integration.gitlab.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1121330012130122-2132202021103200-3023130132333323-1233200232030210-0102123032133220-3132222003122202-1032020111302313-1033223033012021) |
| `code_base_integration.gitlab.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-1203231021231230-3020110233011020-3032302000302222-1132303122033300-2212112313122020-3123012333122312-3122020003103203-1231133113231202) |
| `code_base_integration.gitlab.access_token.clear_secret_info.url` | [code_base_integration.gitlab.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-0312001210133303-3202201203210022-1100003330221102-1211012201113011-2232122003230123-2130303110030210-1300113230323311-3310110323230122) |
| `code_base_integration.gitlab_enterprise` | [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-0233331330322001-1001232010212213-0131020200321101-1301321303200230-1201133032230023-2112320230102313-2221013121211232-1010013010121322) |
| `code_base_integration.gitlab_enterprise.access_token` | [code_base_integration.gitlab_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-2323312003131323-0202133000302133-0030301210220310-1210223002013112-1120002310201101-0333220331010103-2211021031331011-3200011300033322) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1320200310233203-2010210211023312-3022030120321102-3102311123130310-0023222033023010-3333313223110122-1122211120103023-0332012122221010) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider](data-sources--code_base_integration--reference--group-001.md#canonical-0113333133201200-2103033320220101-3111201301230003-3131330122100230-2213103232131113-0310213223211131-2100300130012313-3313013311120313) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location](data-sources--code_base_integration--reference--group-001.md#canonical-0000000223100223-1023011113112102-1122222333321122-2223011133132203-0331031102010221-2122121031313233-3201303332232010-0013300021102110) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider](data-sources--code_base_integration--reference--group-001.md#canonical-1313012313110313-2131300030110013-0023112102001232-2231130023310033-0111101222321110-2301033311202112-1322131020001222-1002322003013013) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1202112021311223-1232133102121031-1011020222231303-2121213121333311-0122322211031323-2102202121012112-2231121311102223-1012301133203221) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref](data-sources--code_base_integration--reference--group-001.md#canonical-3130203213113020-0232033000312203-0303113202110331-0231000320102033-3133000101133302-1212222200122021-0131011020223210-0111220223120033) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url](data-sources--code_base_integration--reference--group-001.md#canonical-2202033211301201-2332110223102201-1210233130011023-2133333133303111-1323130011012101-3030113212021121-2202230130300302-0323332110032313) |
| `code_base_integration.gitlab_enterprise.url` | [code_base_integration.gitlab_enterprise.url](data-sources--code_base_integration--reference--group-001.md#canonical-0232300011033023-2231232332000101-3032013123203111-1032132332203120-0231133102312000-0001212030112100-3231330001230323-1030020010321311) |
| `description` | [description](data-sources--code_base_integration--reference--group-001.md#canonical-3033100213003330-3230013303013003-2003030102313212-1303210131002112-0331331222000213-3303000303023122-3031331213133121-3000301231311023) |
| `id` | [ID](data-sources--code_base_integration--reference--group-001.md#canonical-2330101111131302-1330110132121202-1100031030303233-3130112321000031-0322232102021321-0312232321033113-3220312303000312-3022233302231000) |
| `labels` | [labels](data-sources--code_base_integration--reference--group-001.md#canonical-1211103333122103-2221121011303331-3112101233211300-3020300232202213-0113200312103023-0222330002131200-0121003112233222-0211030221301320) |
| `name` | [name](data-sources--code_base_integration--reference--group-001.md#canonical-1300221212302001-1000213011121030-0321032323132330-3122111210331313-2231231110031223-2020123003011102-0101303100232132-2023333130212331) |
| `namespace` | [namespace](data-sources--code_base_integration--reference--group-001.md#canonical-3320002100120300-2213021210221331-3121010301302212-2223101121012030-0102001000233033-3031201100031103-1120101103230013-2031022120111000) |

<a id="canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- code_base_integration

<a id="canonical-1010013100022000-2331302120111003-3230313001232321-3331311313301223-2012313003210112-2121311111203222-2300312033130132-3101331223002302"></a>

Type: `"single"`. Computed.

Choose your codebase (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection
details.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"azure_repos\",\"bitbucket\",\"bitbucket_server\",\"github\",\"github_enterprise\",\"gitlab\",\"gitlab_enterprise\"]"
}
```

<a id="canonical-3202133313103202-2312301213120321-0021130321222323-2113220331232311-0000032003322212-0330101131130131-0320001232331202-1332122203300132"></a>

### Direct properties for `code_base_integration`

- [azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-2312013102301110-0230231030233300-0021211301310222-3103213213210321-1312133231110012-1030310210300111-1212332313332222-3230030133322202): complete subsection reference.

- [Bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-0010220120131132-0221013133301221-2210300132010130-2301202332031301-2012233320030022-3320303231133210-2103230111313330-3231102123130123): complete subsection reference.

- [bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-2233300220000300-2201300211120013-0002332300233220-3012101120033023-0213023212111120-0302220111313222-3302330000303311-2020232020102220): complete subsection reference.

- [GitHub](data-sources--code_base_integration--reference--group-001.md#canonical-1012030010112222-2032212302001133-2002020102101203-0221201223202231-2232320300332211-2212003200120223-0013132032220323-3320210011001231): complete subsection reference.

- [github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-2132121301020103-1300212030202032-2121130231313332-2112002121001100-0333130020303211-0233223033132322-1202232133321220-3203310131312332): complete subsection reference.

- [GitLab](data-sources--code_base_integration--reference--group-001.md#canonical-3001110332103203-2321013323113130-1313101302030220-0022123212003010-1333220121113030-3023002300322231-1130100230201103-0111230333123233): complete subsection reference.

- [gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-2333101020030012-3001202030003000-0111001000010320-1311301013031322-1001100012201323-1022130330200320-1132213202121130-2232110330132123): complete subsection reference.

<a id="canonical-2312013102301110-0230231030233300-0021211301310222-3103213213210321-1312133231110012-1030310210300111-1212332313332222-3230030133322202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.azure_repos` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- code_base_integration.azure_repos

<a id="canonical-0202130131131210-0010233112112311-0012131033110100-3000033013130230-3322303200331101-3012011133121221-3203220301322220-0031212103121011"></a>

Type: `"single"`. Computed.

Configuration parameter for Azure repos.

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

<a id="canonical-0032312020131021-2132232032101032-2102011000220110-0000112023003111-3230131300131221-3310133233031123-1223332130301002-3230133331332312"></a>

### Direct properties for `code_base_integration.azure_repos`

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-2230020312020311-0122131201203010-1120011113012332-0200230003003211-1121110022220011-2313203012233310-0131100110332101-1121233112203000): complete subsection reference.

<a id="canonical-2230020312020311-0122131201203010-1120011113012332-0200230003003211-1121110022220011-2313203012233310-0131100110332101-1121233112203000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.azure_repos.access_token` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-2312013102301110-0230231030233300-0021211301310222-3103213213210321-1312133231110012-1030310210300111-1212332313332222-3230030133322202)
- code_base_integration.azure_repos.access_token

<a id="canonical-3223313110122200-2223113112320210-1132312320133201-1312011231033112-1200321003020330-0332233310302231-1013322302121320-2333202012321322"></a>

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

<a id="canonical-1123212213001121-2232223231020023-0100221233222122-2203221333123132-3010300202313230-2230211202320132-0210131012111222-3220122223322033"></a>

### Direct properties for `code_base_integration.azure_repos.access_token`

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-2020230001233321-3110111223222302-0021322111320120-1230033013031010-0033101123323132-3303121211203300-1111211011033012-3122031310032333): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-2100303031002311-1110310121220332-3031122033233012-2133302021331203-2130331021002101-0021123301121032-1220110313223100-3130030023103110): complete subsection reference.

<a id="canonical-2020230001233321-3110111223222302-0021322111320120-1230033013031010-0033101123323132-3303121211203300-1111211011033012-3122031310032333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.azure_repos.access_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-2312013102301110-0230231030233300-0021211301310222-3103213213210321-1312133231110012-1030310210300111-1212332313332222-3230030133322202)
- [code_base_integration.azure_repos.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-2230020312020311-0122131201203010-1120011113012332-0200230003003211-1121110022220011-2313203012233310-0131100110332101-1121233112203000)
- code_base_integration.azure_repos.access_token.blindfold_secret_info

<a id="canonical-3330132102012010-3311222213333200-1112310023013112-3000002231110100-1110303310110011-3211303321000102-3322002113332033-0122213000013232"></a>

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

<a id="canonical-1001222210313112-1003201103013200-1220332101001201-0012202303110231-2021313133031301-0210230011022110-2221200313213133-3303112230023123"></a>

### Direct properties for `code_base_integration.azure_repos.access_token.blindfold_secret_info`

<a id="canonical-3033031213301131-0133132300132203-3223112201232100-1231321101301330-3132022303213130-3330010303103001-0110003202003122-2303310221312031"></a>

#### `code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1012103013003232-3110310333000212-0210022230202020-3301021101331330-1021301033013130-2121303201313333-2321111023313333-2132312113131321"></a>

<a id="canonical-1300012112033101-1020033102011323-1230200003310202-3003112111110131-3003330232212203-0121002013230010-0013110301113020-1033101012230020"></a>

#### `code_base_integration.azure_repos.access_token.blindfold_secret_info.location` property

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

<a id="canonical-1112233211301030-1203320210002320-1222102121201330-1300230002103303-2331310300010232-2112231111123121-1231110210130331-1010011102301112"></a>

<a id="canonical-3110320200000331-1201212012000120-2110223110210300-1302222210103033-3110210203231312-3231233313300200-3211312303031231-3133301220132311"></a>

#### `code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider` property

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

<a id="canonical-2100303031002311-1110310121220332-3031122033233012-2133302021331203-2130331021002101-0021123301121032-1220110313223100-3130030023103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.azure_repos.access_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.azure_repos](data-sources--code_base_integration--reference--group-001.md#canonical-2312013102301110-0230231030233300-0021211301310222-3103213213210321-1312133231110012-1030310210300111-1212332313332222-3230030133322202)
- [code_base_integration.azure_repos.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-2230020312020311-0122131201203010-1120011113012332-0200230003003211-1121110022220011-2313203012233310-0131100110332101-1121233112203000)
- code_base_integration.azure_repos.access_token.clear_secret_info

<a id="canonical-0301211332122013-1103031333112031-0111321202330211-1322332123012313-2131332111030022-0323230133120221-0230130113213332-0121003201133301"></a>

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

<a id="canonical-3011232331101120-3210130123111101-1123032312123302-3133320120201331-1122020002113123-1313320033100022-2322030310202103-3033100233330301"></a>

### Direct properties for `code_base_integration.azure_repos.access_token.clear_secret_info`

<a id="canonical-0101102222223220-0020120131100303-3223031312302111-0230101303001331-2111111122003210-0133222120020023-2232221222012302-3013112323212211"></a>

#### `code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3222020132201113-1120322013012111-0323220331230033-1100131303311100-0021030301103000-1111213232330313-3131130003033322-3032312023321211"></a>

<a id="canonical-3022020001232212-3130330332002313-3123110101210221-0012111221133132-3232122333031100-1230223210321033-1302032210100301-2002302023200321"></a>

#### `code_base_integration.azure_repos.access_token.clear_secret_info.url` property

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

<a id="canonical-0010220120131132-0221013133301221-2210300132010130-2301202332031301-2012233320030022-3320303231133210-2103230111313330-3231102123130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.bitbucket` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- code_base_integration.Bitbucket

<a id="canonical-0211321132030210-0030302103320323-2233001201210011-3333020331033230-2123103302123311-0200231003203301-3230003133132311-2003133212102033"></a>

Type: `"single"`. Computed.

Bitbucket Cloud Integration.

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

<a id="canonical-0112300120122121-0011113320020132-3001033201000232-3320202230202202-3022231130032132-3332002023213301-1231022011033033-3123031323223012"></a>

### Direct properties for `code_base_integration.bitbucket`

- [passwd](data-sources--code_base_integration--reference--group-001.md#canonical-0330113222133222-3223011103200223-3011202201122021-2333211231310032-1120000100313000-0101213021013333-1200210120000310-2210132300203330): complete subsection reference.

<a id="canonical-0212210122321313-2222131103103131-3223201111330213-1311313032332321-1021333100012100-0032320323322020-2301110310022131-2021330223001231"></a>

<a id="canonical-0300003022013333-3230322113031200-0230122230231321-3212113210032133-1120033300220021-1130300203133321-0313121202033020-1120301112020222"></a>

#### `code_base_integration.bitbucket.username` property

Type: `"string"`. Computed.

Bitbucket Username. Human-readable name for the resource

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

<a id="canonical-0330113222133222-3223011103200223-3011202201122021-2333211231310032-1120000100313000-0101213021013333-1200210120000310-2210132300203330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.bitbucket.passwd` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-0010220120131132-0221013133301221-2210300132010130-2301202332031301-2012233320030022-3320303231133210-2103230111313330-3231102123130123)
- code_base_integration.Bitbucket.passwd

<a id="canonical-1103202030320000-3200010132112102-2001220020010231-1110211003233130-3331321010201111-2321333222200313-2323003223012311-1222212231122120"></a>

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

<a id="canonical-1312311012001121-0122010222010021-2330232312113200-2103032000333310-0233220330110222-1303201123121110-2300222130312100-0232122323230220"></a>

### Direct properties for `code_base_integration.bitbucket.passwd`

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-2012332213320020-0231211100002103-2012200023321123-0211323101312312-1112101121210102-3333100330113110-0133332111323033-3331033101323323): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-3130110203232303-2131233112300111-1122102323122030-1013020022132021-0102202101220023-3010010112202103-3203301333321023-2320002233002233): complete subsection reference.

<a id="canonical-2012332213320020-0231211100002103-2012200023321123-0211323101312312-1112101121210102-3333100330113110-0133332111323033-3331033101323323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.bitbucket.passwd.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-0010220120131132-0221013133301221-2210300132010130-2301202332031301-2012233320030022-3320303231133210-2103230111313330-3231102123130123)
- [code_base_integration.bitbucket.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-0330113222133222-3223011103200223-3011202201122021-2333211231310032-1120000100313000-0101213021013333-1200210120000310-2210132300203330)
- code_base_integration.Bitbucket.passwd.blindfold_secret_info

<a id="canonical-2022000113022010-0023203131323030-3331122101121230-1121032132312302-3012000200230211-1003103323332023-1010111121113001-1032333112223130"></a>

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

<a id="canonical-0301230020000231-2330230223310303-3301320020331311-0321213300103320-1032311223031300-2223101211320131-3321313310233213-1013322231010103"></a>

### Direct properties for `code_base_integration.bitbucket.passwd.blindfold_secret_info`

<a id="canonical-1030203010123020-0033102031230102-0131231211023220-0330133333030223-1320122223233121-0032211102232313-3032100321331220-1300233221211132"></a>

#### `code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3302211001323032-3031103003332001-3320322302223002-1222332300313110-0002221031210231-0202132030113001-1210131232301211-3323202133210310"></a>

<a id="canonical-2013031033232231-2331201223223033-3333323311013012-2213320110030113-3233322333001332-0003230210010122-1020013210303001-1021213130031321"></a>

#### `code_base_integration.bitbucket.passwd.blindfold_secret_info.location` property

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

<a id="canonical-2001112313231232-1002121031222033-3031130133023001-0302203331212012-0201011013303212-1213003113100303-0102331211123020-1223012130133333"></a>

<a id="canonical-2320203212111222-3031313112013120-3122210122131213-0300323001011221-2201233122312032-3021231230021132-0012301103003120-1033313023312311"></a>

#### `code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider` property

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

<a id="canonical-3130110203232303-2131233112300111-1122102323122030-1013020022132021-0102202101220023-3010010112202103-3203301333321023-2320002233002233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.bitbucket.passwd.clear_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.bitbucket](data-sources--code_base_integration--reference--group-001.md#canonical-0010220120131132-0221013133301221-2210300132010130-2301202332031301-2012233320030022-3320303231133210-2103230111313330-3231102123130123)
- [code_base_integration.bitbucket.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-0330113222133222-3223011103200223-3011202201122021-2333211231310032-1120000100313000-0101213021013333-1200210120000310-2210132300203330)
- code_base_integration.Bitbucket.passwd.clear_secret_info

<a id="canonical-1032222211201223-0102321331223113-2101212223113223-3133121103000231-3223320322201233-2022101023310032-2012122311303232-0013210231001120"></a>

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

<a id="canonical-3231012320312222-3221300101223210-2223130121213301-2211003332203123-2030203213022203-1313312220333201-3331103213331312-0033310103030322"></a>

### Direct properties for `code_base_integration.bitbucket.passwd.clear_secret_info`

<a id="canonical-3030121022131122-0020222313031130-3323103330302032-0201322210211132-3100202301122230-3003232130101032-1323103232301303-3222102323123020"></a>

#### `code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3223302302212002-2030312200013303-1021230313200003-3233000120001001-0022032312201132-2302103002112300-1022303131133333-1212311332131312"></a>

<a id="canonical-1233131030320111-3322122103202022-1223322210033011-3223323033232030-2010020113013321-0321212123310000-3131103001321121-1001032320313112"></a>

#### `code_base_integration.bitbucket.passwd.clear_secret_info.url` property

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

<a id="canonical-2233300220000300-2201300211120013-0002332300233220-3012101120033023-0213023212111120-0302220111313222-3302330000303311-2020232020102220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.bitbucket_server` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- code_base_integration.bitbucket_server

<a id="canonical-1213013313131331-2132223200313111-2120213010302210-1310302303301330-1103311112203111-3212131330031202-1011101232023133-1103101113311232"></a>

Type: `"single"`. Computed.

Configuration parameter for Bitbucket server.

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

<a id="canonical-3032123300122320-2320110202231221-1113200212202131-0002130001103112-2311011023200011-1202023101123232-1202013303300012-0322233000202031"></a>

### Direct properties for `code_base_integration.bitbucket_server`

- [passwd](data-sources--code_base_integration--reference--group-001.md#canonical-2222313232010230-1213030221100020-1001210013112011-1001123030123201-1330320330012232-1313302200101012-0232211100332002-2333200303132231): complete subsection reference.

<a id="canonical-3313323020101033-1220131032001230-1203202201022322-3000311131321320-3000233112012023-1112130312332230-3231330213210233-2003330232103312"></a>

<a id="canonical-1111232303320022-3023200230312111-0133231102002300-0132333012102101-0230200100231021-0212310031212202-2013130002312311-1223300101001202"></a>

#### `code_base_integration.bitbucket_server.url` property

Type: `"string"`. Computed.

Bitbucket Server URL. URL or URI reference

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
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
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0333333323011100-1223210000012210-0100122001201101-3330022233220202-0101230002120330-0011321023021232-3021103332030313-3002300100330220"></a>

<a id="canonical-0000032200223220-1111121030211113-0000023113121302-1121112332013002-0102331102231311-2121222313012120-2212121213001101-2113130021200222"></a>

#### `code_base_integration.bitbucket_server.username` property

Type: `"string"`. Computed.

Bitbucket Server Username. Human-readable name for the resource

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

<a id="canonical-2211303113033033-2303220113122110-3201130212002010-3130130321021200-3023212022232203-3212133230112012-1203023003302323-1120123302110231"></a>

<a id="canonical-2332302131222323-2122131013202210-0303301230011221-2033120322333030-2332110132311030-0233333323333323-0323101110222113-3223132202120320"></a>

#### `code_base_integration.bitbucket_server.verify_ssl` property

Type: `"bool"`. Computed.

Verify SSL. Configuration parameter for verify SSL

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

<a id="canonical-2222313232010230-1213030221100020-1001210013112011-1001123030123201-1330320330012232-1313302200101012-0232211100332002-2333200303132231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.bitbucket_server.passwd` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-2233300220000300-2201300211120013-0002332300233220-3012101120033023-0213023212111120-0302220111313222-3302330000303311-2020232020102220)
- code_base_integration.bitbucket_server.passwd

<a id="canonical-2213030111203132-3101001032211031-0220013300312032-0101001031321022-1010221322203202-3111130021222310-1130100120023020-1103222231031220"></a>

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

<a id="canonical-1210222122200133-3320010133001300-2311101121100100-3011221332312310-3232131021203132-3222303321303101-0232100033233320-0201032211121102"></a>

### Direct properties for `code_base_integration.bitbucket_server.passwd`

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-3332011002302033-3102231330120303-1221020001210223-2332120103021123-2232201122011230-3000121010101103-0323010102113110-0030332030012130): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1011130102101323-3103232313301131-1133010312200103-1002332001000331-2211003201332231-2200201032110123-0221003322030110-0201002313033102): complete subsection reference.

<a id="canonical-3332011002302033-3102231330120303-1221020001210223-2332120103021123-2232201122011230-3000121010101103-0323010102113110-0030332030012130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.bitbucket_server.passwd.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-2233300220000300-2201300211120013-0002332300233220-3012101120033023-0213023212111120-0302220111313222-3302330000303311-2020232020102220)
- [code_base_integration.bitbucket_server.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-2222313232010230-1213030221100020-1001210013112011-1001123030123201-1330320330012232-1313302200101012-0232211100332002-2333200303132231)
- code_base_integration.bitbucket_server.passwd.blindfold_secret_info

<a id="canonical-0210303333122013-0031011322311030-2120110112102213-2330201000210121-0210331100231120-0123213212003131-2033333133333223-2231031013121111"></a>

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

<a id="canonical-0320002113122302-3203201233312023-1010311003001131-3312200203130220-0011300330110131-0122030310001203-0311023300221022-3123101113220110"></a>

### Direct properties for `code_base_integration.bitbucket_server.passwd.blindfold_secret_info`

<a id="canonical-2313132130221022-1232201012110301-3012331333121322-1310203123220002-0020321330331122-0312303110313322-3223023323000231-1002200121013120"></a>

#### `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1332300132221310-0030001331120130-2033310013230001-1113020133003131-0232002311121312-2322132102233122-0230211103331323-2331130000203302"></a>

<a id="canonical-1111032013133223-3121333121123303-1313210311201202-0311313111233133-3302002321303103-0311232331223331-0131300100113201-1313313320131232"></a>

#### `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location` property

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

<a id="canonical-1022113122313211-0020133133023230-2000003220221130-2111303330201312-0202133311300312-0023320023031200-3101001311101322-0100003123213130"></a>

<a id="canonical-0133011120300320-3233221311033003-1323203113202120-0202103302310031-2223031230222302-1001001210221221-2131021201030322-0232000320202230"></a>

#### `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider` property

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

<a id="canonical-1011130102101323-3103232313301131-1133010312200103-1002332001000331-2211003201332231-2200201032110123-0221003322030110-0201002313033102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.bitbucket_server.passwd.clear_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.bitbucket_server](data-sources--code_base_integration--reference--group-001.md#canonical-2233300220000300-2201300211120013-0002332300233220-3012101120033023-0213023212111120-0302220111313222-3302330000303311-2020232020102220)
- [code_base_integration.bitbucket_server.passwd](data-sources--code_base_integration--reference--group-001.md#canonical-2222313232010230-1213030221100020-1001210013112011-1001123030123201-1330320330012232-1313302200101012-0232211100332002-2333200303132231)
- code_base_integration.bitbucket_server.passwd.clear_secret_info

<a id="canonical-1001030300121102-3213113203320222-0302131023023101-0033110301030033-2320123212223230-3202021211310000-3022210123101301-0111110133323201"></a>

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

<a id="canonical-3300210032020020-3023012332101202-2132003320311001-3332302221121332-1201220110123231-3213320003123200-1112222130103102-3321330101332130"></a>

### Direct properties for `code_base_integration.bitbucket_server.passwd.clear_secret_info`

<a id="canonical-2313331310011021-2330330100022321-0320202331112303-2012312223020200-3012313310311033-2231313203033111-2300232011300011-3131103231001103"></a>

#### `code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1311120301122100-2010000101030221-2011011011300033-2233121220203033-2111022131310320-0300011132120231-0210002130302111-0100011301211203"></a>

<a id="canonical-3331212020323130-1001322010300103-3011320310231221-0000020101310222-2302103222312332-3033203212021010-3302331330313312-2203313331032001"></a>

#### `code_base_integration.bitbucket_server.passwd.clear_secret_info.url` property

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

<a id="canonical-1012030010112222-2032212302001133-2002020102101203-0221201223202231-2232320300332211-2212003200120223-0013132032220323-3320210011001231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.github` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- code_base_integration.GitHub

<a id="canonical-3033202230113000-1113223010021122-3333310101332220-1332212013230101-0101212111133203-2010331330210201-3312001231112212-1121100022131312"></a>

Type: `"single"`. Computed.

GitHub Integration.

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

<a id="canonical-2020011003312312-2003102203102330-1211221200310023-1031302030011220-1303223231131213-3101201111030103-0332131103010213-0120003230021031"></a>

### Direct properties for `code_base_integration.github`

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0031203003130223-2331110102032102-0230201121010323-2313233122220111-0213333311303010-1120232310331033-3023103133102301-3113310111220102): complete subsection reference.

<a id="canonical-3101001321301121-0323013313111230-0123211220123133-0133232303112332-1223201111213012-2103131212221110-0111122301121331-1003233313131300"></a>

<a id="canonical-0332013112021001-2331020313023001-3323013232130110-2131130102212222-2302212121320132-3330110232320013-3102021023333203-0300212011122131"></a>

#### `code_base_integration.github.username` property

Type: `"string"`. Computed.

GitHub Username. Human-readable name for the resource

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

<a id="canonical-0313132122130311-2021102120001011-2322020012212030-1103212113330231-3003300110202230-3011300300311201-1101211120230212-1323131120322003"></a>

<a id="canonical-3021010110322202-2101332200210130-1103312323113030-0231312032213102-3301202132223011-1130202312203001-0112013120220033-1113221311011212"></a>

#### `code_base_integration.github.verify_ssl` property

Type: `"bool"`. Computed.

GitHub Verify SSL. Configuration parameter for verify SSL

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

<a id="canonical-0031203003130223-2331110102032102-0230201121010323-2313233122220111-0213333311303010-1120232310331033-3023103133102301-3113310111220102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.github.access_token` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-1012030010112222-2032212302001133-2002020102101203-0221201223202231-2232320300332211-2212003200120223-0013132032220323-3320210011001231)
- code_base_integration.GitHub.access_token

<a id="canonical-0121010303222003-0212012030131202-3200231013121223-1223302212001131-3303122300023230-0200100321002002-0200233103121001-2312310331211211"></a>

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

<a id="canonical-1132120120312122-2133120220031112-3211311202001123-0121111210100302-0223300320013331-2101013013112002-0310320033001021-2123213033213310"></a>

### Direct properties for `code_base_integration.github.access_token`

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1110302200110330-2331121331132002-0313330301221030-0202123021113302-1132122002233202-0323220110233232-0101000102021132-3200330031031100): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1131021320130121-2221122202222312-0111122032212202-3113321301331232-2113000011330031-0213320102303132-3310000313003221-1222101301302002): complete subsection reference.

<a id="canonical-1110302200110330-2331121331132002-0313330301221030-0202123021113302-1132122002233202-0323220110233232-0101000102021132-3200330031031100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.github.access_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-1012030010112222-2032212302001133-2002020102101203-0221201223202231-2232320300332211-2212003200120223-0013132032220323-3320210011001231)
- [code_base_integration.github.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0031203003130223-2331110102032102-0230201121010323-2313233122220111-0213333311303010-1120232310331033-3023103133102301-3113310111220102)
- code_base_integration.GitHub.access_token.blindfold_secret_info

<a id="canonical-3101200132210123-0233300121033202-2321032330010210-0321022110220202-3022102032211311-1333212030103021-1311011032003222-0133113101030113"></a>

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

<a id="canonical-3330023311122301-0330103303000133-1213020201010310-0101021300132203-0102312003233232-1200112223031220-2321010010300223-3100123331103320"></a>

### Direct properties for `code_base_integration.github.access_token.blindfold_secret_info`

<a id="canonical-2100212323031132-1313130302013123-1311130010222302-0220302032203011-0001320132011211-3322321332112113-3313103213331210-0320312032232121"></a>

#### `code_base_integration.github.access_token.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3110321202111300-1111231130332022-1130200321100033-1013212023001103-0322312201133323-3321132310333330-3322012001132200-0302201033022013"></a>

<a id="canonical-2221122110311131-3023231033301023-0301310013212110-3130101120220222-1330003302322003-0002023211213033-0131320030132322-3201033221101233"></a>

#### `code_base_integration.github.access_token.blindfold_secret_info.location` property

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

<a id="canonical-3132112131210020-2130012330120312-0132212013221201-0302211220222120-0022130232333012-1333103203320332-2002002012223113-3112021012000313"></a>

<a id="canonical-3101131011222220-1332000211030311-1010303011002032-2332321300310022-1222030112111210-3023122233002022-3131120232120121-3023323100332230"></a>

#### `code_base_integration.github.access_token.blindfold_secret_info.store_provider` property

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

<a id="canonical-1131021320130121-2221122202222312-0111122032212202-3113321301331232-2113000011330031-0213320102303132-3310000313003221-1222101301302002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.github.access_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.github](data-sources--code_base_integration--reference--group-001.md#canonical-1012030010112222-2032212302001133-2002020102101203-0221201223202231-2232320300332211-2212003200120223-0013132032220323-3320210011001231)
- [code_base_integration.github.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-0031203003130223-2331110102032102-0230201121010323-2313233122220111-0213333311303010-1120232310331033-3023103133102301-3113310111220102)
- code_base_integration.GitHub.access_token.clear_secret_info

<a id="canonical-1210102330033233-0113200023211111-0311100232030001-2111033201331131-3323311203022333-3221201002322222-1331231121132011-2001303203122010"></a>

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

<a id="canonical-3311302332300002-2012032221100021-1221313200332133-3021000123022210-1321201230212322-0322000033232313-3032321321233210-2011112022222100"></a>

### Direct properties for `code_base_integration.github.access_token.clear_secret_info`

<a id="canonical-1333111302130120-2002202022033120-3021211032013123-0231003330302320-2303120313022032-0020233310303302-1120123202322232-2333220113331112"></a>

#### `code_base_integration.github.access_token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2302120320310221-2000011002201212-3131220122101311-2120232303102202-3320221122301310-3021033311021000-0233013302223133-1120312100200202"></a>

<a id="canonical-0212300033100020-2201201002132111-0010313223120220-0022223011123211-1332211011220101-3003020203230033-1111111122002222-3320031300020113"></a>

#### `code_base_integration.github.access_token.clear_secret_info.url` property

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

<a id="canonical-2132121301020103-1300212030202032-2121130231313332-2112002121001100-0333130020303211-0233223033132322-1202232133321220-3203310131312332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.github_enterprise` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- code_base_integration.github_enterprise

<a id="canonical-0311010301031223-3203230023012031-1333331213221212-2311010121230212-0101300331233133-1021201203030001-0023102031312232-2220221010123212"></a>

Type: `"single"`. Computed.

Configuration parameter for GitHub enterprise.

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

<a id="canonical-2222133132330312-3201010231233102-2201211332012311-2213130302230232-3133121320202120-3023131122012123-0001002203320112-3033223203200331"></a>

### Direct properties for `code_base_integration.github_enterprise`

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-1303332230321330-1023001301112323-2320022023200131-3231100132330131-3210130333030330-3303321030033103-3301331233301210-1330133302131301): complete subsection reference.

<a id="canonical-0130331330312023-2323110010331001-1103231301223131-1230232022221102-2233333131023101-1203130012131011-0210102103002200-1002333333000231"></a>

<a id="canonical-2110131203101323-2303121210101033-1300020113311021-3101003001022213-0003031320300223-1130010213312233-2020013003123021-0010113230300323"></a>

#### `code_base_integration.github_enterprise.hostname` property

Type: `"string"`. Computed.

GitHub Hostname. Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0210011022302203-3132233030030033-0131131001223233-1302311102121332-2010003331101203-1103101121301221-3330310223013212-0112231232311110"></a>

<a id="canonical-0012200233313321-0032130120310211-3033030322101123-3323011303002120-3213131313233302-3312310002001101-0220320021121133-3231113100300301"></a>

#### `code_base_integration.github_enterprise.username` property

Type: `"string"`. Computed.

GitHub Username. Human-readable name for the resource

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

<a id="canonical-1303332230321330-1023001301112323-2320022023200131-3231100132330131-3210130333030330-3303321030033103-3301331233301210-1330133302131301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.github_enterprise.access_token` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-2132121301020103-1300212030202032-2121130231313332-2112002121001100-0333130020303211-0233223033132322-1202232133321220-3203310131312332)
- code_base_integration.github_enterprise.access_token

<a id="canonical-0003121020013131-3102222300013231-3022330111312000-2013220121003030-0002020002332333-0331202301203330-2330211112021202-2202330132121223"></a>

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

<a id="canonical-1221011220112210-3023002120233331-0210102132131023-0212220311112012-2213223313222220-0130020113231323-1122320331313112-2300323023322112"></a>

### Direct properties for `code_base_integration.github_enterprise.access_token`

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-2301113010320022-2002113011331132-0132311200233302-0222201010023033-3101212212221230-0213030322222000-1320323012203333-0120113330313032): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-0220001112330221-3321313200322210-0001111033330310-1120121202100310-2222202110130232-2201203223221303-0213313223113311-1023021111323210): complete subsection reference.

<a id="canonical-2301113010320022-2002113011331132-0132311200233302-0222201010023033-3101212212221230-0213030322222000-1320323012203333-0120113330313032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.github_enterprise.access_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-2132121301020103-1300212030202032-2121130231313332-2112002121001100-0333130020303211-0233223033132322-1202232133321220-3203310131312332)
- [code_base_integration.github_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-1303332230321330-1023001301112323-2320022023200131-3231100132330131-3210130333030330-3303321030033103-3301331233301210-1330133302131301)
- code_base_integration.github_enterprise.access_token.blindfold_secret_info

<a id="canonical-2130030303023122-0100002030202013-1202230221233233-3221231131101131-0020012211031132-0100012032313232-1201231220023200-1311002311030211"></a>

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

<a id="canonical-0100231132001211-3301331011233100-0131301223300112-0130100121223101-3233222220033303-1032210223111231-3300213212032322-1112312330101022"></a>

### Direct properties for `code_base_integration.github_enterprise.access_token.blindfold_secret_info`

<a id="canonical-1313002011133120-3112031131113200-1130020003230023-0230330000130011-3033310013310012-0100211300101233-0021100203230231-3322222211113302"></a>

#### `code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1203301103310311-1012010123220111-2210222002000303-0001321112103320-1020220303322201-2102310030231010-3000023202020000-3233213303112311"></a>

<a id="canonical-3102003012333002-2212232330330330-3321333230313131-1200320321013223-0012210322120330-1221213113233112-0303303330111231-2012312323003133"></a>

#### `code_base_integration.github_enterprise.access_token.blindfold_secret_info.location` property

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

<a id="canonical-3220123330113223-2132023312213321-2110002103313331-0012211002010231-1021213033123012-1130300232232213-1213232001100132-2211230213031131"></a>

<a id="canonical-0202220322311303-2112320321231112-2311211303033011-2333221001320233-1123330310233032-1003301123300122-1020133033031020-3031120000222303"></a>

#### `code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider` property

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

<a id="canonical-0220001112330221-3321313200322210-0001111033330310-1120121202100310-2222202110130232-2201203223221303-0213313223113311-1023021111323210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.github_enterprise.access_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.github_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-2132121301020103-1300212030202032-2121130231313332-2112002121001100-0333130020303211-0233223033132322-1202232133321220-3203310131312332)
- [code_base_integration.github_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-1303332230321330-1023001301112323-2320022023200131-3231100132330131-3210130333030330-3303321030033103-3301331233301210-1330133302131301)
- code_base_integration.github_enterprise.access_token.clear_secret_info

<a id="canonical-1232130201211313-2131330220120112-1022121132231200-3203011003123311-3131023102123101-0231120332101233-1300022203203220-2010132002011102"></a>

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

<a id="canonical-2232323213233200-3013002011131112-2111112020110010-2121000101132222-3313112233022102-2310022311321011-0103011203202021-3103120011231313"></a>

### Direct properties for `code_base_integration.github_enterprise.access_token.clear_secret_info`

<a id="canonical-2202203221103013-1022302031121111-3201130100021103-3212212200222013-0210101300331312-1120233320201011-3301132100221232-2121230221111201"></a>

#### `code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3100103132103033-2102011231113303-2121122111122010-1333303110233032-3302302121111210-2321330232312211-1303011013232021-0033110033110223"></a>

<a id="canonical-2121320112333001-3211222122303232-0011312003021130-1130201131302101-3002223132001000-2222112203202012-3110320323121300-0311212202320322"></a>

#### `code_base_integration.github_enterprise.access_token.clear_secret_info.url` property

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

<a id="canonical-3001110332103203-2321013323113130-1313101302030220-0022123212003010-1333220121113030-3023002300322231-1130100230201103-0111230333123233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.gitlab` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- code_base_integration.GitLab

<a id="canonical-2313001223332131-0033312233010000-0121123132202130-3322330013232232-2312032120203013-2320130303010123-2002201121201011-2110330231023031"></a>

Type: `"single"`. Computed.

GitLab Cloud Integration.

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

<a id="canonical-3323100212223003-3113120130303133-1221023213110102-0202222131110221-1023123202012132-3111031203303330-1220202321301033-2113031031231202"></a>

### Direct properties for `code_base_integration.gitlab`

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-3101321010213323-3223001313123203-1220333202331320-0230311012030220-2203332002220002-2020030130220031-3232231233003323-2303001122200021): complete subsection reference.

<a id="canonical-3101321010213323-3223001313123203-1220333202331320-0230311012030220-2203332002220002-2020030130220031-3232231233003323-2303001122200021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.gitlab.access_token` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-3001110332103203-2321013323113130-1313101302030220-0022123212003010-1333220121113030-3023002300322231-1130100230201103-0111230333123233)
- code_base_integration.GitLab.access_token

<a id="canonical-2031030002313130-1111130033130321-2100332332100013-3200210023312220-2110222100002032-3212322112312002-2321311033121322-2201133103121123"></a>

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

<a id="canonical-1302210330100321-3222230102100010-3331112210211021-3112330331312033-1002300330121123-0123113031022222-0333023122112222-2032123313303222"></a>

### Direct properties for `code_base_integration.gitlab.access_token`

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-1131322311100211-0130311213130020-1132010000210100-0003310132212321-1213211133211100-1200230312013111-1332003122033310-3311202123321330): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-3120133302321332-3312222332232332-2201333233213120-2201000102201032-0110201332132102-2212030013331210-3303222111000020-3320200013132012): complete subsection reference.

<a id="canonical-1131322311100211-0130311213130020-1132010000210100-0003310132212321-1213211133211100-1200230312013111-1332003122033310-3311202123321330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.gitlab.access_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-3001110332103203-2321013323113130-1313101302030220-0022123212003010-1333220121113030-3023002300322231-1130100230201103-0111230333123233)
- [code_base_integration.gitlab.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-3101321010213323-3223001313123203-1220333202331320-0230311012030220-2203332002220002-2020030130220031-3232231233003323-2303001122200021)
- code_base_integration.GitLab.access_token.blindfold_secret_info

<a id="canonical-3000210213110323-3120210111222012-2301000010020131-2203212200010203-0313010001122031-0121032200220032-0020311233302102-3120033010012312"></a>

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

<a id="canonical-2223333210133312-0332123001001222-3303012001023221-0110310313331122-1333012101003223-2113310332002131-2302230202233132-1223312310203003"></a>

### Direct properties for `code_base_integration.gitlab.access_token.blindfold_secret_info`

<a id="canonical-1333313211110322-1212133121033332-0300030121222001-0111321203302321-0003100131301100-2210302101231112-1331030030011200-1301200232230332"></a>

#### `code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0101222232233032-0102110300322031-3202313312333000-1131210330020103-3311312312232300-0203222100130001-0203323310321210-2332030333232121"></a>

<a id="canonical-1023011210131100-1020223032020303-1011023230103121-3023230032001022-2231002302110320-1032123111012010-1113001013300002-2310300331013323"></a>

#### `code_base_integration.gitlab.access_token.blindfold_secret_info.location` property

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

<a id="canonical-1203101033103322-2323230222021233-1133310231133311-2032011232300112-2233132133101220-0221300022002231-2132212020200021-0021031301330110"></a>

<a id="canonical-1233231002332233-3220212022112022-1333013211010133-0320010011303211-0302201210123211-0311123030312030-3012000110310230-0323110131232031"></a>

#### `code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider` property

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

<a id="canonical-3120133302321332-3312222332232332-2201333233213120-2201000102201032-0110201332132102-2212030013331210-3303222111000020-3320200013132012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.gitlab.access_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.gitlab](data-sources--code_base_integration--reference--group-001.md#canonical-3001110332103203-2321013323113130-1313101302030220-0022123212003010-1333220121113030-3023002300322231-1130100230201103-0111230333123233)
- [code_base_integration.gitlab.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-3101321010213323-3223001313123203-1220333202331320-0230311012030220-2203332002220002-2020030130220031-3232231233003323-2303001122200021)
- code_base_integration.GitLab.access_token.clear_secret_info

<a id="canonical-1121330012130122-2132202021103200-3023130132333323-1233200232030210-0102123032133220-3132222003122202-1032020111302313-1033223033012021"></a>

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

<a id="canonical-0112010010102223-2023031130333123-3200322331310101-2000202233013333-3303110301222012-1003103331003030-0300310123022312-2220331321003211"></a>

### Direct properties for `code_base_integration.gitlab.access_token.clear_secret_info`

<a id="canonical-1203231021231230-3020110233011020-3032302000302222-1132303122033300-2212112313122020-3123012333122312-3122020003103203-1231133113231202"></a>

#### `code_base_integration.gitlab.access_token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0312001210133303-3202201203210022-1100003330221102-1211012201113011-2232122003230123-2130303110030210-1300113230323311-3310110323230122"></a>

<a id="canonical-1023333032110002-0232323203303311-2303213300312202-0300113212301210-3003102031211021-1101300100312122-3032302203233230-2102121332112123"></a>

#### `code_base_integration.gitlab.access_token.clear_secret_info.url` property

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

<a id="canonical-2333101020030012-3001202030003000-0111001000010320-1311301013031322-1001100012201323-1022130330200320-1132213202121130-2232110330132123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.gitlab_enterprise` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- code_base_integration.gitlab_enterprise

<a id="canonical-0233331330322001-1001232010212213-0131020200321101-1301321303200230-1201133032230023-2112320230102313-2221013121211232-1010013010121322"></a>

Type: `"single"`. Computed.

Configuration parameter for GitLab enterprise.

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

<a id="canonical-1312013030311220-0123313210313001-1233303120300101-2022122312311013-0021203111310300-3300013332111032-3130032301010012-2113331213330322"></a>

### Direct properties for `code_base_integration.gitlab_enterprise`

- [access_token](data-sources--code_base_integration--reference--group-001.md#canonical-3011122023332211-2213323302222103-2022023313232233-3210101231330221-1301120110112000-2201322000220113-3131312223112103-0003313032000030): complete subsection reference.

<a id="canonical-0232300011033023-2231232332000101-3032013123203111-1032132332203120-0231133102312000-0001212030112100-3231330001230323-1030020010321311"></a>

<a id="canonical-3012331232132201-1001121333323002-1301033112301233-1120113110110002-3321113211203023-1020231030100101-0022301131023111-0020232230331333"></a>

#### `code_base_integration.gitlab_enterprise.url` property

Type: `"string"`. Computed.

GitLab URL. URL or URI reference

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
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
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3011122023332211-2213323302222103-2022023313232233-3210101231330221-1301120110112000-2201322000220113-3131312223112103-0003313032000030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.gitlab_enterprise.access_token` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-2333101020030012-3001202030003000-0111001000010320-1311301013031322-1001100012201323-1022130330200320-1132213202121130-2232110330132123)
- code_base_integration.gitlab_enterprise.access_token

<a id="canonical-2323312003131323-0202133000302133-0030301210220310-1210223002013112-1120002310201101-0333220331010103-2211021031331011-3200011300033322"></a>

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

<a id="canonical-0121121131133333-0012323131131011-3310030101233100-3233101030001132-3030300300103213-3203122212020101-0123212333330110-1221022230220211"></a>

### Direct properties for `code_base_integration.gitlab_enterprise.access_token`

- [blindfold_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-2230331311311202-3000231033323113-2101203233100312-0320302121013211-0322202031002303-3322202333333211-0210312000131001-2101301101030311): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--reference--group-001.md#canonical-0023332332010100-3000210103303310-2010331223030003-2010302201222111-1011103221221221-2030210200012200-1332300320313030-3021123033302011): complete subsection reference.

<a id="canonical-2230331311311202-3000231033323113-2101203233100312-0320302121013211-0322202031002303-3322202333333211-0210312000131001-2101301101030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-2333101020030012-3001202030003000-0111001000010320-1311301013031322-1001100012201323-1022130330200320-1132213202121130-2232110330132123)
- [code_base_integration.gitlab_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-3011122023332211-2213323302222103-2022023313232233-3210101231330221-1301120110112000-2201322000220113-3131312223112103-0003313032000030)
- code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info

<a id="canonical-1320200310233203-2010210211023312-3022030120321102-3102311123130310-0023222033023010-3333313223110122-1122211120103023-0332012122221010"></a>

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

<a id="canonical-2332003221302302-1302103311331100-0130023213232023-2012113003123131-3122103021321201-0301312321023231-0313132020210022-1210203223202230"></a>

### Direct properties for `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info`

<a id="canonical-0113333133201200-2103033320220101-3111201301230003-3131330122100230-2213103232131113-0310213223211131-2100300130012313-3313013311120313"></a>

#### `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0000000223100223-1023011113112102-1122222333321122-2223011133132203-0331031102010221-2122121031313233-3201303332232010-0013300021102110"></a>

<a id="canonical-0023101301000123-2231330212301033-1100302313131232-1211003130330220-0031303010102121-3232131000230332-1033011133201230-0113213033112120"></a>

#### `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location` property

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

<a id="canonical-1313012313110313-2131300030110013-0023112102001232-2231130023310033-0111101222321110-2301033311202112-1322131020001222-1002322003013013"></a>

<a id="canonical-3211212132103221-2012213220320100-2123001301132022-1002122030001103-2233320310133113-1203322020121213-2213121122230301-2123301130011113"></a>

#### `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider` property

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

<a id="canonical-0023332332010100-3000210103303310-2010331223030003-2010302201222111-1011103221221221-2030210200012200-1332300320313030-3021123033302011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `code_base_integration.gitlab_enterprise.access_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Property reference](data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [code_base_integration](data-sources--code_base_integration--reference--group-001.md#canonical-1103133121103233-3230311033010233-0022322231112001-0111010231301100-3021010021232203-3323302013130203-1003233331210001-1211302003302122)
- [code_base_integration.gitlab_enterprise](data-sources--code_base_integration--reference--group-001.md#canonical-2333101020030012-3001202030003000-0111001000010320-1311301013031322-1001100012201323-1022130330200320-1132213202121130-2232110330132123)
- [code_base_integration.gitlab_enterprise.access_token](data-sources--code_base_integration--reference--group-001.md#canonical-3011122023332211-2213323302222103-2022023313232233-3210101231330221-1301120110112000-2201322000220113-3131312223112103-0003313032000030)
- code_base_integration.gitlab_enterprise.access_token.clear_secret_info

<a id="canonical-1202112021311223-1232133102121031-1011020222231303-2121213121333311-0122322211031323-2102202121012112-2231121311102223-1012301133203221"></a>

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

<a id="canonical-3202303011010220-3221002312212333-2213220112010310-0003213303322223-2230001201310130-3122132202233232-2002101313022111-1010131030012010"></a>

### Direct properties for `code_base_integration.gitlab_enterprise.access_token.clear_secret_info`

<a id="canonical-3130203213113020-0232033000312203-0303113202110331-0231000320102033-3133000101133302-1212222200122021-0131011020223210-0111220223120033"></a>

#### `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2202033211301201-2332110223102201-1210233130011023-2133333133303111-1323130011012101-3030113212021121-2202230130300302-0323332110032313"></a>

<a id="canonical-0111111110202232-2030223003012133-0202222110313232-3313310313133303-0130213122100310-1002322032111112-1033023312111013-2131221103231311"></a>

#### `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url` property

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
