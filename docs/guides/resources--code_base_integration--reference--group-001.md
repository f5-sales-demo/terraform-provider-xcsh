---
page_title: "xcsh_code_base_integration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration reference."
---

# xcsh_code_base_integration reference

<a id="canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230120132031112-2232113132133120-3230033023312331-1013213102331301-3113330301013211-2311312201331001-0311233312031211-0010022220120030"></a>

## Property reference — Property reference / 211133000232 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- Property reference

<a id="canonical-2213033100330212-2222330030313313-1220022002210002-1131112120213300-2123331313103002-3331001031020021-3332013202323023-3321130033212302"></a>

## Direct properties — Property reference / 211133000232 / 3

<a id="canonical-2302200332310022-2300320103103300-0130033200231330-2300111033033123-1101211323012021-3203302233303312-2300303120020333-2001110112001213"></a>

<a id="canonical-2300101023121110-2203313310333020-1303230230122212-3110211132212321-0303021322233001-3122012320021002-2123211303030213-0212221232021001"></a>

## annotations property — Property reference / 211133000232 / 4

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

- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001): complete subsection reference.

<a id="canonical-1321012110001012-3002330112200111-2003033230331321-1120311302211231-0322301211122332-1120123233023220-3322032222022222-1322101331102312"></a>

<a id="canonical-1001133120023323-3031022303322211-1210012233201103-1320323322301202-1011032210031303-2323333233020111-2330332033202203-1333231210320232"></a>

## description property — Property reference / 211133000232 / 5

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

<a id="canonical-1033111122000330-0133221222210223-2102031010112132-1220313021131111-1132000030001210-1103022303222023-0323101230201130-0330031011011113"></a>

<a id="canonical-2103231122133101-3232222313111321-1210200113221102-0301022210011200-3021211121031120-0033101201312030-3203031030033202-0313002202111002"></a>

## disable property — Property reference / 211133000232 / 6

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

<a id="canonical-3102031222313332-0333120012321321-2330301102110322-2330301112323201-2231221212201332-1231212102331302-1131321303123303-0113231011131100"></a>

<a id="canonical-3111022003122202-2312021030121212-2000232332323220-1300121101332001-0123213210330002-2312023323111330-1310231200320101-1113033003102203"></a>

## ID property — Property reference / 211133000232 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1202002203231023-3232233201012023-1203000303123210-3202233222302102-2120321201323221-2311133232301112-3200002230022122-1212210022030231"></a>

<a id="canonical-2333210021332313-3332033312133013-1121031321110123-1113110123211032-0033231301030130-0321023122210030-1311200012022222-3223323010012002"></a>

## labels property — Property reference / 211133000232 / 8

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

<a id="canonical-1223320001233001-3212321210131223-2023322012222221-1022032330023010-2132110231112230-1333232321102330-2113211003301021-2011012221313302"></a>

<a id="canonical-3120230330220320-2233321323102023-1312023100302133-3301121133131332-1103310210000132-1223101030233220-0003203303003022-1332322331222011"></a>

## name property — Property reference / 211133000232 / 9

Type: `"string"`. Required.

Name of the codebase Integration. Must be unique within the namespace.

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

<a id="canonical-0101112211330333-2131201022130022-0131001110220102-3330211110312013-1203211100213103-2112322001101211-1131003300021030-0322332013033012"></a>

<a id="canonical-1130030232011310-2102321322013331-1003132322133122-0002203333013020-2300133331312123-0321102221000013-2333202133111031-0101301112212223"></a>

## namespace property — Property reference / 211133000232 / 10

Type: `"string"`. Required.

Namespace where the codebase Integration is created.

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

- [timeouts](resources--code_base_integration--reference--group-001.md#canonical-1001330330201111-0023201102213331-2220222121001222-0223210221230112-1120101313111223-3303312223222112-0220130220213110-0330212303030203): complete subsection reference.

<a id="canonical-1023203010133333-3032201111003002-3131000322232323-3102112322313000-1300001231022030-1133013313113003-3020310333232323-1101101113310011"></a>

## All schema paths — Property reference / 211133000232 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--code_base_integration--reference--group-001.md#canonical-2302200332310022-2300320103103300-0130033200231330-2300111033033123-1101211323012021-3203302233303312-2300303120020333-2001110112001213) |
| `code_base_integration` | [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-1331230301021310-0033332103122132-2031232031223120-3200323001323011-0320221321000202-2230023302201300-3030100022003121-3122020222100121) |
| `code_base_integration.azure_repos` | [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-0000313003011222-1031323321331132-1100213313332301-0222100320300120-2211202022122210-2123101200013220-2033022323013121-0301313120211123) |
| `code_base_integration.azure_repos.access_token` | [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-3203011230000012-2122322100331221-3001010130301211-0233211102003332-0022122232330303-1230122000212223-0220220211223302-0323122002321010) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info` | [code_base_integration.azure_repos.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1332311113123301-0121211121120130-2202230112312333-2203111212221023-0223111221202301-2012200130000011-2332133212001132-2033033323011330) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-2130322010001113-2332120123332110-0213303332102022-3013202210001003-1012330120330002-0203131322100110-0001311201333230-2330012003221232) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.location` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-2211003100213013-1031210032112131-2312032130011213-1100333202113310-3333131302003203-0322032210230212-2120000303011112-3320012202032221) |
| `code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider` | [code_base_integration.azure_repos.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-2000002333332231-0013111133332110-0311103222030000-2120210013331101-2320311210033213-1330112000123010-2323010200230210-3311021312313333) |
| `code_base_integration.azure_repos.access_token.clear_secret_info` | [code_base_integration.azure_repos.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0303200101113003-3120333013002231-1330220330031131-2113111210310213-0320202333222200-0113131201222003-3123230120211011-0030200333303303) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref` | [code_base_integration.azure_repos.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-0312121333220323-3303231333302232-1202333230330312-1013130020121333-1300101023220232-0231230321103303-0121102230000000-1202003111230011) |
| `code_base_integration.azure_repos.access_token.clear_secret_info.url` | [code_base_integration.azure_repos.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-3130012202311333-1323122203203313-0321320131013232-2310112000121022-0103222112102221-1321303022130131-2031333012121023-2133212000122020) |
| `code_base_integration.bitbucket` | [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-0022103102333130-3020121332202311-0101331220323212-2222110013111230-1010330003213101-2221002033202200-1201312130203212-0300312201032313) |
| `code_base_integration.bitbucket.passwd` | [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-0011121130201002-0131112301202112-1022032322332001-1021103102211310-1120103123202123-3103130302002332-3303033321022001-0021100210102101) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info` | [code_base_integration.bitbucket.passwd.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3200003222201123-2223313100321310-1130011013323112-0300132131113002-0111011002313203-1120320130210010-1032100311030203-1213002212122132) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-0130330030002303-3302321130330200-3311212201302021-0303200011232203-1003331023120100-2010220102120320-0133203111323313-3013301222033010) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-2021021331213301-1203121113100333-0031000301322123-3122210132130201-2312032012130321-2211032010112211-0312110231110202-1323122332123212) |
| `code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket.passwd.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-2112020122313312-1301221320212103-0213231030223200-3012103203320332-2013031331012021-3021021321230002-0210100300310123-0133211320121011) |
| `code_base_integration.bitbucket.passwd.clear_secret_info` | [code_base_integration.bitbucket.passwd.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3132312210310300-1320122101111023-3022021202131120-2122011133031231-0110221022202011-3102023213333010-1213113332001321-2233223002133130) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket.passwd.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-1100220320011213-3113231010033001-1112210321020201-0031120100222120-1333110210131112-2122113300202302-0002132233330113-0130211221110210) |
| `code_base_integration.bitbucket.passwd.clear_secret_info.url` | [code_base_integration.bitbucket.passwd.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-2001203013223232-3023122000212300-2032023100213333-3310031111210230-0012333312023212-1013323303100122-0131133123201232-1001110030203113) |
| `code_base_integration.bitbucket.username` | [code_base_integration.bitbucket.username](resources--code_base_integration--reference--group-001.md#canonical-0311120023102130-1100212120010130-3322202213233221-0303211210221122-1011233012112222-3112112331113320-2110120321103131-0013221323311000) |
| `code_base_integration.bitbucket_server` | [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-1111222202021031-3211010021313130-3311110130300213-2100212220102221-1330211201112030-1302221033211330-1333020003230310-1313030210210001) |
| `code_base_integration.bitbucket_server.passwd` | [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-0100312011133133-3112120030321330-0203011313203332-1331033321322103-3303200022221220-0002313221002131-1311032000002203-3110130000200112) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2223021223101221-3131121123000232-0330222001210003-1320312211320313-0312113300133200-2023323332111030-3133311030030302-0032330311003301) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-1200002000122011-2103301231333201-1112030001003211-3133101331212030-2333221301011230-3223001233230333-3000321100213320-0301013203311001) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-3022101023130001-2012223321203301-1022133301001133-3013003111122100-3300230300010310-3113211023201222-2023233130311103-3113223133131213) |
| `code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider` | [code_base_integration.bitbucket_server.passwd.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-2230203322211013-2103221131233100-2022320330211303-2121010312033101-0003220103323110-2313030132333203-0130330200030112-1332311031110220) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info` | [code_base_integration.bitbucket_server.passwd.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0100311030322130-3322000123121011-2102220111102310-0111033230322023-1012012200123332-2112103310101101-0010023103010232-2013332112113303) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-2112311201312013-3122101103100031-0210210312131132-1023021120321302-3003011033310033-1332132233320132-1132012023030121-0001110330303030) |
| `code_base_integration.bitbucket_server.passwd.clear_secret_info.url` | [code_base_integration.bitbucket_server.passwd.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-2132322302221132-2310223313120311-0323031101311223-3301122210213023-2112301310001210-0301113320120320-2203131111303230-2331303311131023) |
| `code_base_integration.bitbucket_server.url` | [code_base_integration.bitbucket_server.url](resources--code_base_integration--reference--group-001.md#canonical-3111101331112130-0002023323032003-3011013212133011-3022022322222113-3211001000012331-1131311122212000-1223011031031011-2210002001122303) |
| `code_base_integration.bitbucket_server.username` | [code_base_integration.bitbucket_server.username](resources--code_base_integration--reference--group-001.md#canonical-3332100001200130-1303223103313021-2013130222110223-1230233122032223-2130011133310202-0320121202313112-3310323030111003-2111022020332211) |
| `code_base_integration.bitbucket_server.verify_ssl` | [code_base_integration.bitbucket_server.verify_ssl](resources--code_base_integration--reference--group-001.md#canonical-1210231223030022-2022002223120320-3002231001130023-3131003112202032-3321130122202210-1322233012113030-1233101101300011-3231121223301303) |
| `code_base_integration.github` | [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-0321120102332130-2200011201202203-0303022113032130-3102020321120012-2032033023110132-3003103230322213-1013311120020132-3212012012310001) |
| `code_base_integration.github.access_token` | [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-0122103013012303-0222032100110212-3102023121023200-2122023311210201-2332200013121121-0200030032113113-1022203212100121-0100203103221322) |
| `code_base_integration.github.access_token.blindfold_secret_info` | [code_base_integration.github.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0230210110201110-0323132233122222-2131330202211133-1003302130120231-1331330301310033-3320331200101323-2311100311021332-1210201233031222) |
| `code_base_integration.github.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-1213213132021221-2222113011210010-1232210132210203-3001230320012010-2212331102221122-3210232232113203-1012312323111021-1022132010230313) |
| `code_base_integration.github.access_token.blindfold_secret_info.location` | [code_base_integration.github.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-2323112030331110-2033133113112211-0321032320331020-0211312103330120-3221221020131003-2223112132322212-0033220303012202-2210133222132023) |
| `code_base_integration.github.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-2211333010110323-3130322101023200-0131033031112201-1113020310120011-2102130130003110-0321021203213002-2103223321131211-1130221022032331) |
| `code_base_integration.github.access_token.clear_secret_info` | [code_base_integration.github.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0233031221111100-1210311312322311-3122300100120320-0032222301230333-0033130103312032-2202121223122213-3102130231020131-0221220030333211) |
| `code_base_integration.github.access_token.clear_secret_info.provider_ref` | [code_base_integration.github.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-2203132333220132-1300302201023133-3002031233211131-1203113303101213-2231131013020211-0022333201222313-2030322100211101-1001100013003311) |
| `code_base_integration.github.access_token.clear_secret_info.url` | [code_base_integration.github.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-1120100303033113-1303222103310100-3001302033012302-3321021312202232-2333231302033211-1300231202333020-0023030311313123-2002321103222130) |
| `code_base_integration.github.username` | [code_base_integration.github.username](resources--code_base_integration--reference--group-001.md#canonical-0100031002230310-2300202132020013-2002013122200310-3230121030121211-2311312301120001-0302132233131312-2213203130002101-0020001022320222) |
| `code_base_integration.github.verify_ssl` | [code_base_integration.github.verify_ssl](resources--code_base_integration--reference--group-001.md#canonical-0320322210201000-0100110112101331-3333303322120202-0120031321322321-3311313003203013-2213111020213123-1230211333321231-0022212123321330) |
| `code_base_integration.github_enterprise` | [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-3313010321301003-1201110111000210-2133210332313212-1222030203320213-3202020031032232-3003222131132013-1312230032322333-2310302200303231) |
| `code_base_integration.github_enterprise.access_token` | [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-3122003233231001-2232033121033021-2000022030120202-0133210111012002-2023132311130032-0102232012012220-2230331013010233-2113221010200010) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1320131001022213-1202013132021113-0121201102031332-3132103330101323-3202030132223230-2100210022323322-1012121311320231-2301122212001121) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-2301113020212122-3122100231121133-0020111113011123-2201201010310121-0032210122021113-2011233031330132-3001213310003213-1333220123131132) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-3301233020300033-3311330201011012-2102101012212003-3320030312102201-0003002313001320-2332031010222211-3331131013203103-3013203130012220) |
| `code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.github_enterprise.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-2032210000232023-0323333023131132-1332333000321020-3221101231132133-0301210333323330-2110023132012202-3320030131320212-3032013321213321) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info` | [code_base_integration.github_enterprise.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2231332203103321-2133121221112201-1231123312131300-1200003100323212-1203310012111321-3112333003132211-2113003133310213-2021232122233022) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.github_enterprise.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-2320033123100032-3013130231332313-3020312310200233-1223211003131321-1302203323102312-3011013112332331-1023112302233101-0003101213111322) |
| `code_base_integration.github_enterprise.access_token.clear_secret_info.url` | [code_base_integration.github_enterprise.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-0100323032213020-0100123220133203-2020120300033133-2031002013112200-2223131220211322-2201322322113101-2122023300023112-2322023230030010) |
| `code_base_integration.github_enterprise.hostname` | [code_base_integration.github_enterprise.hostname](resources--code_base_integration--reference--group-001.md#canonical-2302131331022022-1033010222330232-2022101032102310-0012313230020233-3200301012312033-2132220031210230-0203203233020022-3203130211120231) |
| `code_base_integration.github_enterprise.username` | [code_base_integration.github_enterprise.username](resources--code_base_integration--reference--group-001.md#canonical-1003320322021113-2112332000301213-0213123113220023-3003101300021312-0032203320202220-1020001120222203-1303332023103320-3321203031100111) |
| `code_base_integration.gitlab` | [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-1100130301111201-2122131221233013-3221101222300121-0302020012122111-0220320013003100-1021013223112202-1120120001200131-2123123302132022) |
| `code_base_integration.gitlab.access_token` | [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-1231030203033030-2220321013321311-2211003020030103-1331213212201202-1332110323032320-3013203330321222-2203110321210223-2312211010120231) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info` | [code_base_integration.gitlab.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0303221123033031-0223021332213333-3120102322121021-3112313313331112-1103202232323003-2121203220332113-0233013301121100-3002223011303023) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-2231223101310012-2321233331020201-0222130300100031-1202103302032020-0032213130111223-1102231132320213-0231112231011020-3313313002002231) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-1001213120223333-1331113022221220-2002331103212200-3133321121200220-2220001312111200-0111330030332331-2230311013001333-1022300332322132) |
| `code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-0321110000002132-2312321012322201-2012221332232302-1301313230232201-0100221323313312-0002202002112201-3301000132313322-3310100203201022) |
| `code_base_integration.gitlab.access_token.clear_secret_info` | [code_base_integration.gitlab.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0211002233302201-1321102011321030-2131120313120133-1113223303301102-0022112022310202-0103112013233203-1032111103010303-3110300223223012) |
| `code_base_integration.gitlab.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-0210133201032200-1202312331330023-0133130030031201-0303100230131220-1321220310233322-2201113102132230-0013121302202302-1230113300002122) |
| `code_base_integration.gitlab.access_token.clear_secret_info.url` | [code_base_integration.gitlab.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-0311203013210103-3113232303333320-3130103322211233-1312231102010113-3020132323302112-2300020120221323-1123221213122301-0023303121011030) |
| `code_base_integration.gitlab_enterprise` | [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-3332012310231021-3001321133001002-0333112132013022-1210323123213333-3020121030102001-1001231112023111-3020121020001120-2233332320001331) |
| `code_base_integration.gitlab_enterprise.access_token` | [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-0233322103303010-2301032313010230-0221002010000011-1122131020300211-2220312113110320-0001102122013023-1032330013312302-0301211311120210) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0023302030303200-3202323212330221-1103130310022232-1332230320221220-3333210230333131-0120230220202212-2321111312101123-3103001333320201) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.decryption_provider](resources--code_base_integration--reference--group-001.md#canonical-0113302330133133-3033103312110320-2101231110203323-3133222230221221-1031312021223032-1223131332121000-2301030300130203-2311023033000222) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.location](resources--code_base_integration--reference--group-001.md#canonical-3030210123022133-1120231013203030-3231231111022032-2103123221021022-0203211310133021-3310103131202213-2123231011331331-3003232220023123) |
| `code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider` | [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info.store_provider](resources--code_base_integration--reference--group-001.md#canonical-0301321333230210-1030231333233012-2332011023013231-2031231001210133-3301023233011022-0330233013021112-3131222031030002-2233221030331230) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3112133332102123-1131011002210032-1212203203332212-2210011032120020-2200312032122213-0030310113102211-1222230021223231-3020333201312020) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.provider_ref](resources--code_base_integration--reference--group-001.md#canonical-1130320003210221-1111312103313021-1003211211030312-3220000101123101-3212132120230101-2111203202313010-2332222001002222-0100011002022323) |
| `code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url` | [code_base_integration.gitlab_enterprise.access_token.clear_secret_info.url](resources--code_base_integration--reference--group-001.md#canonical-0100132210013223-3320302222111322-1332223300302303-0330212130120212-1231010221322312-3132321313231232-2013200301123310-1121022003023321) |
| `code_base_integration.gitlab_enterprise.url` | [code_base_integration.gitlab_enterprise.url](resources--code_base_integration--reference--group-001.md#canonical-0233101000302110-3002312302312113-3113000302022311-3131201213203032-0003333003333232-0103310223201121-0001223023303101-0112112013232222) |
| `description` | [description](resources--code_base_integration--reference--group-001.md#canonical-1321012110001012-3002330112200111-2003033230331321-1120311302211231-0322301211122332-1120123233023220-3322032222022222-1322101331102312) |
| `disable` | [disable](resources--code_base_integration--reference--group-001.md#canonical-1033111122000330-0133221222210223-2102031010112132-1220313021131111-1132000030001210-1103022303222023-0323101230201130-0330031011011113) |
| `id` | [ID](resources--code_base_integration--reference--group-001.md#canonical-3102031222313332-0333120012321321-2330301102110322-2330301112323201-2231221212201332-1231212102331302-1131321303123303-0113231011131100) |
| `labels` | [labels](resources--code_base_integration--reference--group-001.md#canonical-1202002203231023-3232233201012023-1203000303123210-3202233222302102-2120321201323221-2311133232301112-3200002230022122-1212210022030231) |
| `name` | [name](resources--code_base_integration--reference--group-001.md#canonical-1223320001233001-3212321210131223-2023322012222221-1022032330023010-2132110231112230-1333232321102330-2113211003301021-2011012221313302) |
| `namespace` | [namespace](resources--code_base_integration--reference--group-001.md#canonical-0101112211330333-2131201022130022-0131001110220102-3330211110312013-1203211100213103-2112322001101211-1131003300021030-0322332013033012) |
| `timeouts` | [timeouts](resources--code_base_integration--reference--group-001.md#canonical-1330131233002223-2133202331322133-3113202313130220-0313213201031030-0213113232323120-2123112233003230-3220330222121313-0331221233311222) |
| `timeouts.create` | [timeouts.create](resources--code_base_integration--reference--group-001.md#canonical-3203331120121322-0213303322001333-2012221122233123-3201130330301010-0202310210230201-1333132310201333-1310320323112111-2313211020321033) |
| `timeouts.delete` | [timeouts.delete](resources--code_base_integration--reference--group-001.md#canonical-1021302301202201-1120033300130103-1321213320122310-2320121210322131-1021022131010000-2301131111133213-3010023110303110-3001331221210231) |
| `timeouts.read` | [timeouts.read](resources--code_base_integration--reference--group-001.md#canonical-0022201130231211-2112230111300021-0210113311002312-1023223110311030-0030112230330123-3130102120021011-2230132303120112-1211120301312230) |
| `timeouts.update` | [timeouts.update](resources--code_base_integration--reference--group-001.md#canonical-1303223310030233-1220212132103000-1110111300212232-3103202002232003-2030210320202222-3102310200120031-0110220003302232-1213122123231311) |

<a id="canonical-2122132133010331-0033012210032331-3222333130210112-3213032022231203-1222200330002033-3000033010231300-3332320133011133-0120202022232002"></a>

## Next pages — Property reference / 211133000232 / 12

- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [timeouts](resources--code_base_integration--reference--group-001.md#canonical-1001330330201111-0023201102213331-2220222121001222-0223210221230112-1120101313111223-3303312223222112-0220130220213110-0330212303030203)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110213130132202-1113021100011313-3230231230000213-3231001020301303-1110210101211322-3102203203310033-2111332330131102-3313211002301333"></a>

## code_base_integration — code_base_integration / 122330333303 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- code_base_integration

<a id="canonical-1331230301021310-0033332103122132-2031232031223120-3200323001323011-0320221321000202-2230023302201300-3030100022003121-3122020222100121"></a>

Type: `"object"`. single nested block, Optional.

Choose your codebase (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection
details.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("azure_repos",
    "bitbucket"),
  validators.ConflictingObjectAttributes("azure_repos",
    "bitbucket_server"),
  validators.ConflictingObjectAttributes("azure_repos",
    "github"),
  validators.ConflictingObjectAttributes("azure_repos",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("azure_repos",
    "gitlab"),
  validators.ConflictingObjectAttributes("azure_repos",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket",
    "bitbucket_server"),
  validators.ConflictingObjectAttributes("bitbucket",
    "github"),
  validators.ConflictingObjectAttributes("bitbucket",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket",
    "gitlab"),
  validators.ConflictingObjectAttributes("bitbucket",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "github"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "gitlab"),
  validators.ConflictingObjectAttributes("bitbucket_server",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("github",
    "github_enterprise"),
  validators.ConflictingObjectAttributes("github",
    "gitlab"),
  validators.ConflictingObjectAttributes("github",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("github_enterprise",
    "gitlab"),
  validators.ConflictingObjectAttributes("github_enterprise",
    "gitlab_enterprise"),
  validators.ConflictingObjectAttributes("gitlab",
    "gitlab_enterprise")}
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
  "x-ves-oneof-field-type": "[\"azure_repos\",\"bitbucket\",\"bitbucket_server\",\"github\",\"github_enterprise\",\"gitlab\",\"gitlab_enterprise\"]"
}
```

Terraform syntax:

```terraform
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233221120323002-0002112321213333-3022111001000213-1133333233210021-3110003333111221-1301113110332031-3030321201201013-1131011301211100"></a>

## Direct properties — code_base_integration / 122330333303 / 3

- [azure_repos](resources--code_base_integration--reference--group-001.md#canonical-3201011132331200-1130220300313200-1031001232220211-3122301211033132-1132033001110210-1120032113203100-0101321313202112-1220222110032100): complete subsection reference.

- [Bitbucket](resources--code_base_integration--reference--group-001.md#canonical-3331021320323120-0021121133102300-3331011033013300-0131212023232003-3131102311220212-2030030132023320-0333230023300002-0133212120120111): complete subsection reference.

- [bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-1120030323021110-3332221313312113-1230210301103113-2030303033132103-2032321200333101-0031322133320210-3300211013200001-0300233133212230): complete subsection reference.

- [GitHub](resources--code_base_integration--reference--group-001.md#canonical-0120221310020322-3030103310232300-0032130112020031-0033101303023130-0200232032210221-2113122232123220-3121000201013200-1022232122022023): complete subsection reference.

- [github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-1232323021020220-2213003210330030-2202310102321311-2102032013032021-2112021221021322-1202203303201020-1122202103122132-0312100133211130): complete subsection reference.

- [GitLab](resources--code_base_integration--reference--group-001.md#canonical-2021200022200103-1003301322023302-2022313230222201-2101232003211313-1001233301102023-1203233330221213-2310133231313303-3121331302210300): complete subsection reference.

- [gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-2300101201031033-0233121100211322-0323002120210212-3020213213320323-2323322312133130-2103132210131111-2130333232002130-1220301121211123): complete subsection reference.

<a id="canonical-2330300011112220-3113013221022013-0223130330012011-2031102210112203-0220110221303211-2021123133313113-3030030013313211-0223221220010233"></a>

## Next pages — code_base_integration / 122330333303 / 4

- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-3201011132331200-1130220300313200-1031001232220211-3122301211033132-1132033001110210-1120032113203100-0101321313202112-1220222110032100)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-3331021320323120-0021121133102300-3331011033013300-0131212023232003-3131102311220212-2030030132023320-0333230023300002-0133212120120111)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-1120030323021110-3332221313312113-1230210301103113-2030303033132103-2032321200333101-0031322133320210-3300211013200001-0300233133212230)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-0120221310020322-3030103310232300-0032130112020031-0033101303023130-0200232032210221-2113122232123220-3121000201013200-1022232122022023)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-1232323021020220-2213003210330030-2202310102321311-2102032013032021-2112021221021322-1202203303201020-1122202103122132-0312100133211130)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-2021200022200103-1003301322023302-2022313230222201-2101232003211313-1001233301102023-1203233330221213-2310133231313303-3121331302210300)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-2300101201031033-0233121100211322-0323002120210212-3020213213320323-2323322312133130-2103132210131111-2130333232002130-1220301121211123)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-3201011132331200-1130220300313200-1031001232220211-3122301211033132-1132033001110210-1120032113203100-0101321313202112-1220222110032100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331233102231302-3121013202120223-1222301011011002-0123302101133301-0132222131032012-2330133130330122-2102213031030023-2310131233313011"></a>

## code_base_integration.azure_repos — azure_repos / 120300131121 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- code_base_integration.azure_repos

<a id="canonical-0000313003011222-1031323321331132-1100213313332301-0222100320300120-2211202022122210-2123101200013220-2033022323013121-0301313120211123"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
azure_repos {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220120000000310-3322202121133200-2303111322230121-2232111001000232-1032111330031102-3321112320110132-2020021312012132-3121122020130101"></a>

## Direct properties — azure_repos / 120300131121 / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-1233003213210212-3323013131003222-3221203232232021-3210320021330021-2013102310332030-2001133031213302-0222133013312000-3100210022010210): complete subsection reference.

<a id="canonical-0321321223120330-3112330022033232-1022221322301133-3032120212332101-0211310330100330-0103203001030021-2222313122013331-0031023302132322"></a>

## Next pages — azure_repos / 120300131121 / 4

- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-1233003213210212-3323013131003222-3221203232232021-3210320021330021-2013102310332030-2001133031213302-0222133013312000-3100210022010210)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1233003213210212-3323013131003222-3221203232232021-3210320021330021-2013102310332030-2001133031213302-0222133013312000-3100210022010210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023113322103233-2303210100013212-2033300221032221-1312133301121313-3023103111202222-3102303020021323-0132331122222322-1210313220200322"></a>

## code_base_integration.azure_repos.access_token — access_token / 323321011310 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-3201011132331200-1130220300313200-1031001232220211-3122301211033132-1132033001110210-1120032113203100-0101321313202112-1220222110032100)
- code_base_integration.azure_repos.access_token

<a id="canonical-3203011230000012-2122322100331221-3001010130301211-0233211102003332-0022122232330303-1230122000212223-0220220211223302-0323122002321010"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213202023032130-1202200211011211-1302223212331020-1201321123331022-2012321320310203-0210201110303012-1020100233210032-0201320110033301"></a>

## Direct properties — access_token / 323321011310 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3002133311102322-1310110210003011-2212130313033222-1022332223312031-3323210231331233-0122103312001310-2003021202222032-1220111331333123): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3323010223310032-1331233001013122-0202123111302331-1330312320022321-3312231032331001-2101331312221322-1212130320102233-2112010031031110): complete subsection reference.

<a id="canonical-1132112300022032-2001130213021310-3012302230122302-2102003320210102-2211002130113202-1130122112231130-3212300121121112-0030121322123023"></a>

## Next pages — access_token / 323321011310 / 4

- [code_base_integration.azure_repos.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3002133311102322-1310110210003011-2212130313033222-1022332223312031-3323210231331233-0122103312001310-2003021202222032-1220111331333123)
- [code_base_integration.azure_repos.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3323010223310032-1331233001013122-0202123111302331-1330312320022321-3312231032331001-2101331312221322-1212130320102233-2112010031031110)
- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-3201011132331200-1130220300313200-1031001232220211-3122301211033132-1132033001110210-1120032113203100-0101321313202112-1220222110032100)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-3002133311102322-1310110210003011-2212130313033222-1022332223312031-3323210231331233-0122103312001310-2003021202222032-1220111331333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132113103021221-2103311003023111-1303101300122030-2311222032122312-1201020310222303-1031032001033232-1010010023133312-0010333311012002"></a>

## code_base_integration.azure_repos.access_token.blindfold_secret_info — blindfold_secret_info / 100303011110 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-3201011132331200-1130220300313200-1031001232220211-3122301211033132-1132033001110210-1120032113203100-0101321313202112-1220222110032100)
- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-1233003213210212-3323013131003222-3221203232232021-3210320021330021-2013102310332030-2001133031213302-0222133013312000-3100210022010210)
- code_base_integration.azure_repos.access_token.blindfold_secret_info

<a id="canonical-1332311113123301-0121211121120130-2202230112312333-2203111212221023-0223111221202301-2012200130000011-2332133212001132-2033033323011330"></a>

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

<a id="canonical-3113300133112113-2102110121301211-3302321032301223-0220332213202211-2101303333203030-2230021320220223-1220220323123011-0312102210233032"></a>

## Direct properties — blindfold_secret_info / 100303011110 / 3

<a id="canonical-2130322010001113-2332120123332110-0213303332102022-3013202210001003-1012330120330002-0203131322100110-0001311201333230-2330012003221232"></a>

<a id="canonical-2222031031232000-1033031011111202-3001000313333321-0210101023010110-2230301231331233-3022103303301331-1303103323013300-0322310331332003"></a>

## decryption_provider property — blindfold_secret_info / 100303011110 / 4

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

<a id="canonical-2211003100213013-1031210032112131-2312032130011213-1100333202113310-3333131302003203-0322032210230212-2120000303011112-3320012202032221"></a>

<a id="canonical-0010313033321132-2021130211112323-0112103301032203-3002313311113323-0323213201202010-1110021233320313-3123201310330022-3020320333132232"></a>

## location property — blindfold_secret_info / 100303011110 / 5

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

<a id="canonical-2000002333332231-0013111133332110-0311103222030000-2120210013331101-2320311210033213-1330112000123010-2323010200230210-3311021312313333"></a>

<a id="canonical-2212310100033121-1023231020103011-2231231333200000-1101003132131030-2122113212113223-3221120222130331-3221023113100033-3311230111020131"></a>

## store_provider property — blindfold_secret_info / 100303011110 / 6

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

<a id="canonical-3001022200231212-0333310220103220-2233312210302101-0220302331301010-3032103001212303-3110022311233310-3213123022330221-1223223031013310"></a>

## Next pages — blindfold_secret_info / 100303011110 / 7

- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-1233003213210212-3323013131003222-3221203232232021-3210320021330021-2013102310332030-2001133031213302-0222133013312000-3100210022010210)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-3323010223310032-1331233001013122-0202123111302331-1330312320022321-3312231032331001-2101331312221322-1212130320102233-2112010031031110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022233223101221-1313020312333011-0013102011010013-1123302232130002-1000233212100032-2232210202010202-1020321032111200-1102320101120201"></a>

## code_base_integration.azure_repos.access_token.clear_secret_info — clear_secret_info / 023112213103 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.azure_repos](resources--code_base_integration--reference--group-001.md#canonical-3201011132331200-1130220300313200-1031001232220211-3122301211033132-1132033001110210-1120032113203100-0101321313202112-1220222110032100)
- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-1233003213210212-3323013131003222-3221203232232021-3210320021330021-2013102310332030-2001133031213302-0222133013312000-3100210022010210)
- code_base_integration.azure_repos.access_token.clear_secret_info

<a id="canonical-0303200101113003-3120333013002231-1330220330031131-2113111210310213-0320202333222200-0113131201222003-3123230120211011-0030200333303303"></a>

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

<a id="canonical-2020210310003302-3011303301133311-0100203202323013-1223301200021032-2110302333312223-3302130210220013-1012312222213312-2023220001313322"></a>

## Direct properties — clear_secret_info / 023112213103 / 3

<a id="canonical-0312121333220323-3303231333302232-1202333230330312-1013130020121333-1300101023220232-0231230321103303-0121102230000000-1202003111230011"></a>

<a id="canonical-3030221223210031-1021223220121210-2311131312110310-0220301130232011-2311000113012221-2120200323323000-2010032301011120-3011003100000232"></a>

## provider_ref property — clear_secret_info / 023112213103 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3130012202311333-1323122203203313-0321320131013232-2310112000121022-0103222112102221-1321303022130131-2031333012121023-2133212000122020"></a>

<a id="canonical-0213310131210010-2003222030023202-2332323023112212-3030130113200212-0101033201010103-2321320200313111-3332203030313211-3330332323231232"></a>

## URL property — clear_secret_info / 023112213103 / 5

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

<a id="canonical-0112301023223203-2102121013333230-0122032102131333-1000202213102023-1132121232123331-3022232123231122-1022001301232302-0303130232210031"></a>

## Next pages — clear_secret_info / 023112213103 / 6

- [code_base_integration.azure_repos.access_token](resources--code_base_integration--reference--group-001.md#canonical-1233003213210212-3323013131003222-3221203232232021-3210320021330021-2013102310332030-2001133031213302-0222133013312000-3100210022010210)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-3331021320323120-0021121133102300-3331011033013300-0131212023232003-3131102311220212-2030030132023320-0333230023300002-0133212120120111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033232011223120-3313321212110122-3121102223311210-3200110230301222-0003022030203322-0221021010021021-2213301310022120-2022012103101002"></a>

## code_base_integration.Bitbucket — Bitbucket / 231113222210 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- code_base_integration.Bitbucket

<a id="canonical-0022103102333130-3020121332202311-0101331220323212-2222110013111230-1010330003213101-2221002033202200-1201312130203212-0300312201032313"></a>

Type: `"object"`. single nested block, Optional.

Bitbucket Cloud Integration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("username")}
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
bitbucket {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033133303001231-2131213203101102-1220100212301212-0323300002133201-1231103012233021-3013133101033230-3200211211120003-0201213313122012"></a>

## Direct properties — Bitbucket / 231113222210 / 3

- [passwd](resources--code_base_integration--reference--group-001.md#canonical-3232102310003100-2120310013221103-2132231301203010-2123133130031211-1021101010231323-1101013323031311-2320111220003000-1233311102331303): complete subsection reference.

<a id="canonical-0311120023102130-1100212120010130-3322202213233221-0303211210221122-1011233012112222-3112112331113320-2110120321103131-0013221323311000"></a>

<a id="canonical-0020102030002322-0113100131232013-1110333122321012-2310010001233223-2232331223132212-3310323133333221-1120011111232001-0002000332032222"></a>

## username property — Bitbucket / 231113222210 / 4

Type: `"string"`. Optional.

Bitbucket Username. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-0030333100202131-2031111200022011-0111033121310321-0232002232200301-3103022301103311-3021301331331123-2101301011103302-2232320232123302"></a>

## Next pages — Bitbucket / 231113222210 / 5

- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-3232102310003100-2120310013221103-2132231301203010-2123133130031211-1021101010231323-1101013323031311-2320111220003000-1233311102331303)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-3232102310003100-2120310013221103-2132231301203010-2123133130031211-1021101010231323-1101013323031311-2320111220003000-1233311102331303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232312210001332-3301201131101100-1211332031310130-3211132033303122-1120032020000101-3211102220220302-0132103122113212-1013101221312133"></a>

## code_base_integration.Bitbucket.passwd — passwd / 231203211021 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-3331021320323120-0021121133102300-3331011033013300-0131212023232003-3131102311220212-2030030132023320-0333230023300002-0133212120120111)
- code_base_integration.Bitbucket.passwd

<a id="canonical-0011121130201002-0131112301202112-1022032322332001-1021103102211310-1120103123202123-3103130302002332-3303033321022001-0021100210102101"></a>

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
passwd {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200033200311001-3131332110013023-0313330130331120-2221022100223300-0112111003132001-1222101022222110-2021020201000210-3322030300320322"></a>

## Direct properties — passwd / 231203211021 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2121221120313202-2112023210102302-3203133031200310-3321021313123010-1233222101102321-0022322300113001-0211031320323131-0022102222123300): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1300102200222013-0002221311113131-3230110012110010-3300330111111001-0200222122031020-0102130321102013-1233003231102302-0012100221002012): complete subsection reference.

<a id="canonical-2211120230030212-0131313030332021-2010220212330103-3331233300212201-1313321320320003-1220001332302023-1131011232012130-1101103131121001"></a>

## Next pages — passwd / 231203211021 / 4

- [code_base_integration.bitbucket.passwd.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2121221120313202-2112023210102302-3203133031200310-3321021313123010-1233222101102321-0022322300113001-0211031320323131-0022102222123300)
- [code_base_integration.bitbucket.passwd.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1300102200222013-0002221311113131-3230110012110010-3300330111111001-0200222122031020-0102130321102013-1233003231102302-0012100221002012)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-3331021320323120-0021121133102300-3331011033013300-0131212023232003-3131102311220212-2030030132023320-0333230023300002-0133212120120111)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-2121221120313202-2112023210102302-3203133031200310-3321021313123010-1233222101102321-0022322300113001-0211031320323131-0022102222123300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033331222111203-0331221011000133-0100231101111012-2010022112120020-0200311033021133-0313202322122220-0021121313230121-3112023201321021"></a>

## code_base_integration.Bitbucket.passwd.blindfold_secret_info — blindfold_secret_info / 311211201322 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-3331021320323120-0021121133102300-3331011033013300-0131212023232003-3131102311220212-2030030132023320-0333230023300002-0133212120120111)
- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-3232102310003100-2120310013221103-2132231301203010-2123133130031211-1021101010231323-1101013323031311-2320111220003000-1233311102331303)
- code_base_integration.Bitbucket.passwd.blindfold_secret_info

<a id="canonical-3200003222201123-2223313100321310-1130011013323112-0300132131113002-0111011002313203-1120320130210010-1032100311030203-1213002212122132"></a>

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

<a id="canonical-2220302202003203-0000033133313223-3001110333011313-1203313221210003-0223033303310200-1213021231222301-0001133012010333-2123001121130000"></a>

## Direct properties — blindfold_secret_info / 311211201322 / 3

<a id="canonical-0130330030002303-3302321130330200-3311212201302021-0303200011232203-1003331023120100-2010220102120320-0133203111323313-3013301222033010"></a>

<a id="canonical-3231010030103032-2102222113023230-3202200130012202-0101110233033303-3332201013023032-3020222301211233-3222323300111002-0313002000102022"></a>

## decryption_provider property — blindfold_secret_info / 311211201322 / 4

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

<a id="canonical-2021021331213301-1203121113100333-0031000301322123-3122210132130201-2312032012130321-2211032010112211-0312110231110202-1323122332123212"></a>

<a id="canonical-2302103313000011-3211222302033201-1031302310021101-0032120330021321-0121022012320110-3131302123123231-0301121300301310-1012121010031213"></a>

## location property — blindfold_secret_info / 311211201322 / 5

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

<a id="canonical-2112020122313312-1301221320212103-0213231030223200-3012103203320332-2013031331012021-3021021321230002-0210100300310123-0133211320121011"></a>

<a id="canonical-2213131303302201-0001322212033200-1310210021133320-1101000320101122-1232211000110320-0223311022100010-2213003030231311-2111012022221332"></a>

## store_provider property — blindfold_secret_info / 311211201322 / 6

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

<a id="canonical-3133001330221003-3022223202030113-3232022112310311-3221131012001220-1232012221110303-1220300302203301-0030110213023301-3112000101133313"></a>

## Next pages — blindfold_secret_info / 311211201322 / 7

- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-3232102310003100-2120310013221103-2132231301203010-2123133130031211-1021101010231323-1101013323031311-2320111220003000-1233311102331303)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1300102200222013-0002221311113131-3230110012110010-3300330111111001-0200222122031020-0102130321102013-1233003231102302-0012100221002012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120131133303102-0123032232303001-0103103122021113-2333320212011212-3331013203123131-2131210001131322-2330221023323021-3031312330011200"></a>

## code_base_integration.Bitbucket.passwd.clear_secret_info — clear_secret_info / 213002321123 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.bitbucket](resources--code_base_integration--reference--group-001.md#canonical-3331021320323120-0021121133102300-3331011033013300-0131212023232003-3131102311220212-2030030132023320-0333230023300002-0133212120120111)
- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-3232102310003100-2120310013221103-2132231301203010-2123133130031211-1021101010231323-1101013323031311-2320111220003000-1233311102331303)
- code_base_integration.Bitbucket.passwd.clear_secret_info

<a id="canonical-3132312210310300-1320122101111023-3022021202131120-2122011133031231-0110221022202011-3102023213333010-1213113332001321-2233223002133130"></a>

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

<a id="canonical-1131201310313330-2233010111013031-0130011220332102-2033122312121213-0110211303021230-2013102310121313-1030021112102220-1113210110312310"></a>

## Direct properties — clear_secret_info / 213002321123 / 3

<a id="canonical-1100220320011213-3113231010033001-1112210321020201-0031120100222120-1333110210131112-2122113300202302-0002132233330113-0130211221110210"></a>

<a id="canonical-0210311120133330-1232110231122011-2310300231330311-1103120202012011-3222130103102322-3302232002330033-1022322120202123-3100103131232211"></a>

## provider_ref property — clear_secret_info / 213002321123 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2001203013223232-3023122000212300-2032023100213333-3310031111210230-0012333312023212-1013323303100122-0131133123201232-1001110030203113"></a>

<a id="canonical-0002213001202323-3113033313022222-1010312003030233-0021002230010213-2011131133003102-1122223013310021-3201000003301112-1101201202320102"></a>

## URL property — clear_secret_info / 213002321123 / 5

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

<a id="canonical-2230132233021201-1111231010203023-1210111213320321-1022232211133013-0100101221333002-1022300213230001-0010031213221133-1230000322000233"></a>

## Next pages — clear_secret_info / 213002321123 / 6

- [code_base_integration.bitbucket.passwd](resources--code_base_integration--reference--group-001.md#canonical-3232102310003100-2120310013221103-2132231301203010-2123133130031211-1021101010231323-1101013323031311-2320111220003000-1233311102331303)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1120030323021110-3332221313312113-1230210301103113-2030303033132103-2032321200333101-0031322133320210-3300211013200001-0300233133212230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213332001232000-3221210330023302-0013121333013002-2331131031023222-1320200130003223-3231330001313010-1322131220312303-2122000022232102"></a>

## code_base_integration.bitbucket_server — bitbucket_server / 202113320023 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- code_base_integration.bitbucket_server

<a id="canonical-1111222202021031-3211010021313130-3311110130300213-2100212220102221-1330211201112030-1302221033211330-1333020003230310-1313030210210001"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for Bitbucket server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url",
    "username")}
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
bitbucket_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301331323110220-0310333302122330-0300201232130322-3132312120211030-0012121333100322-2120232310100110-0301223303323113-2102213302211230"></a>

## Direct properties — bitbucket_server / 202113320023 / 3

- [passwd](resources--code_base_integration--reference--group-001.md#canonical-3002010221033221-3132231131213233-2221012010023021-3311213313100101-0010330212023132-1131003000230033-3112322211213313-0032002333000300): complete subsection reference.

<a id="canonical-3111101331112130-0002023323032003-3011013212133011-3022022322222113-3211001000012331-1131311122212000-1223011031031011-2210002001122303"></a>

<a id="canonical-3302131303323010-2002122030231120-0211332133112200-3233310123223000-2232212102312312-2322230000033230-3332011133233312-0100220031310031"></a>

## URL property — bitbucket_server / 202113320023 / 4

Type: `"string"`. Optional.

Bitbucket Server URL. URL or URI reference

Upstream description:

URL or URI reference

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`),
    ""),
}
```

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
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3332100001200130-1303223103313021-2013130222110223-1230233122032223-2130011133310202-0320121202313112-3310323030111003-2111022020332211"></a>

<a id="canonical-0133311231002011-0110021013303000-2021031310113110-0103202222233110-1022221213002333-1210202033203020-3031132232320000-1132121032201000"></a>

## username property — bitbucket_server / 202113320023 / 5

Type: `"string"`. Optional.

Bitbucket Server Username. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-1210231223030022-2022002223120320-3002231001130023-3131003112202032-3321130122202210-1322233012113030-1233101101300011-3231121223301303"></a>

<a id="canonical-2033212001303232-3233001332021000-3031131013112203-3033223233310201-2013012220331303-0222122223203212-1232002202202111-1130011302010102"></a>

## verify_ssl property — bitbucket_server / 202113320023 / 6

Type: `"bool"`. Optional.

Verify SSL. Configuration parameter for verify SSL

Upstream description:

Configuration parameter for verify SSL

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

<a id="canonical-3113020030311222-1223311113223001-1012320231031002-1220333221301000-1120303102301330-0032101221131113-3201021120102322-1112022100033212"></a>

## Next pages — bitbucket_server / 202113320023 / 7

- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-3002010221033221-3132231131213233-2221012010023021-3311213313100101-0010330212023132-1131003000230033-3112322211213313-0032002333000300)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-3002010221033221-3132231131213233-2221012010023021-3311213313100101-0010330212023132-1131003000230033-3112322211213313-0032002333000300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330131130303101-3223111131121022-1212031330120130-2302112322011113-0312120322313122-3231313212011113-1332312032202100-3023103101201000"></a>

## code_base_integration.bitbucket_server.passwd — passwd / 111112311101 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-1120030323021110-3332221313312113-1230210301103113-2030303033132103-2032321200333101-0031322133320210-3300211013200001-0300233133212230)
- code_base_integration.bitbucket_server.passwd

<a id="canonical-0100312011133133-3112120030321330-0203011313203332-1331033321322103-3303200022221220-0002313221002131-1311032000002203-3110130000200112"></a>

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
passwd {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133113300223020-2102320022300011-1030000001333012-0332003123112102-2321302013200321-3120311131000303-1200003201332132-2311332311300122"></a>

## Direct properties — passwd / 111112311101 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1213022333031213-2201103002302131-1001231222122010-1131302013031122-3100032320233001-3000123202301321-3131301031220210-3023301103332031): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3331131030010030-2030230102131110-2100321000303020-2303012012022233-0020313213010032-0020233212130133-2300322002113001-2223311021321321): complete subsection reference.

<a id="canonical-0020000202323131-3130310132023333-2001030232023321-1200002132101231-1302203102221302-2321221101301133-3321202311021100-3113300032233103"></a>

## Next pages — passwd / 111112311101 / 4

- [code_base_integration.bitbucket_server.passwd.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1213022333031213-2201103002302131-1001231222122010-1131302013031122-3100032320233001-3000123202301321-3131301031220210-3023301103332031)
- [code_base_integration.bitbucket_server.passwd.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-3331131030010030-2030230102131110-2100321000303020-2303012012022233-0020313213010032-0020233212130133-2300322002113001-2223311021321321)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-1120030323021110-3332221313312113-1230210301103113-2030303033132103-2032321200333101-0031322133320210-3300211013200001-0300233133212230)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1213022333031213-2201103002302131-1001231222122010-1131302013031122-3100032320233001-3000123202301321-3131301031220210-3023301103332031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121312222202123-2211100232303020-3220002033231303-1200211211210201-3332122122301330-1231201330331213-2333310230302101-2333023022020331"></a>

## code_base_integration.bitbucket_server.passwd.blindfold_secret_info — blindfold_secret_info / 001002321300 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-1120030323021110-3332221313312113-1230210301103113-2030303033132103-2032321200333101-0031322133320210-3300211013200001-0300233133212230)
- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-3002010221033221-3132231131213233-2221012010023021-3311213313100101-0010330212023132-1131003000230033-3112322211213313-0032002333000300)
- code_base_integration.bitbucket_server.passwd.blindfold_secret_info

<a id="canonical-2223021223101221-3131121123000232-0330222001210003-1320312211320313-0312113300133200-2023323332111030-3133311030030302-0032330311003301"></a>

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

<a id="canonical-3302333202130303-1100322102331230-0111112101123303-1031033302310231-1320020112211022-1111101000122130-2331113102033301-2013010113003320"></a>

## Direct properties — blindfold_secret_info / 001002321300 / 3

<a id="canonical-1200002000122011-2103301231333201-1112030001003211-3133101331212030-2333221301011230-3223001233230333-3000321100213320-0301013203311001"></a>

<a id="canonical-2223100021302130-0233011232133321-0333220131202030-0011022113130320-0102012130221130-1132231020213330-3120331330022332-0201223323201021"></a>

## decryption_provider property — blindfold_secret_info / 001002321300 / 4

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

<a id="canonical-3022101023130001-2012223321203301-1022133301001133-3013003111122100-3300230300010310-3113211023201222-2023233130311103-3113223133131213"></a>

<a id="canonical-3133032110233312-3330000020210011-1322300211320010-0313313111331012-1120132333203133-0011122220200311-0322132201300002-1002201012203122"></a>

## location property — blindfold_secret_info / 001002321300 / 5

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

<a id="canonical-2230203322211013-2103221131233100-2022320330211303-2121010312033101-0003220103323110-2313030132333203-0130330200030112-1332311031110220"></a>

<a id="canonical-3030210032021101-2013310133102331-0111211220000310-2000233322003323-0332233211031003-2011033323332221-2021011123103121-3120301000220211"></a>

## store_provider property — blindfold_secret_info / 001002321300 / 6

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

<a id="canonical-2110101210312011-3200113010111000-1003321121201012-3201203201320001-3311020330233001-1211121322013012-0302212023201132-2233021330231221"></a>

## Next pages — blindfold_secret_info / 001002321300 / 7

- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-3002010221033221-3132231131213233-2221012010023021-3311213313100101-0010330212023132-1131003000230033-3112322211213313-0032002333000300)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-3331131030010030-2030230102131110-2100321000303020-2303012012022233-0020313213010032-0020233212130133-2300322002113001-2223311021321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011112203310122-2311320202202223-1200001231110310-0230232320123031-0003213310100030-0313033322321132-2213300010132203-0131301011231231"></a>

## code_base_integration.bitbucket_server.passwd.clear_secret_info — clear_secret_info / 311310103333 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.bitbucket_server](resources--code_base_integration--reference--group-001.md#canonical-1120030323021110-3332221313312113-1230210301103113-2030303033132103-2032321200333101-0031322133320210-3300211013200001-0300233133212230)
- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-3002010221033221-3132231131213233-2221012010023021-3311213313100101-0010330212023132-1131003000230033-3112322211213313-0032002333000300)
- code_base_integration.bitbucket_server.passwd.clear_secret_info

<a id="canonical-0100311030322130-3322000123121011-2102220111102310-0111033230322023-1012012200123332-2112103310101101-0010023103010232-2013332112113303"></a>

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

<a id="canonical-2032122201022000-2111030320311130-0131302221132331-0020311110101213-0032231210023120-1100020021110101-3133212020323201-3030222332200031"></a>

## Direct properties — clear_secret_info / 311310103333 / 3

<a id="canonical-2112311201312013-3122101103100031-0210210312131132-1023021120321302-3003011033310033-1332132233320132-1132012023030121-0001110330303030"></a>

<a id="canonical-3213221103322033-2111002101032202-0001003232030321-2313131210113222-3212032331203212-2320031201112332-1121332012320231-1303032301222113"></a>

## provider_ref property — clear_secret_info / 311310103333 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2132322302221132-2310223313120311-0323031101311223-3301122210213023-2112301310001210-0301113320120320-2203131111303230-2331303311131023"></a>

<a id="canonical-0202221320131020-3122200032112320-0121132010110113-0121122021300221-1300231310323300-2102001102102332-0102121203033230-1030100330332132"></a>

## URL property — clear_secret_info / 311310103333 / 5

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

<a id="canonical-3200301232213332-3002023113312201-0033212321301100-3223233300022031-0131010302000231-3312123332103212-1113303001001103-0031110001022023"></a>

## Next pages — clear_secret_info / 311310103333 / 6

- [code_base_integration.bitbucket_server.passwd](resources--code_base_integration--reference--group-001.md#canonical-3002010221033221-3132231131213233-2221012010023021-3311213313100101-0010330212023132-1131003000230033-3112322211213313-0032002333000300)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-0120221310020322-3030103310232300-0032130112020031-0033101303023130-0200232032210221-2113122232123220-3121000201013200-1022232122022023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222220011032120-0223002002211000-1323331232112012-2112222301102322-3332002022213020-2310013211020031-2302323021100202-0121320130111010"></a>

## code_base_integration.GitHub — GitHub / 033120021211 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- code_base_integration.GitHub

<a id="canonical-0321120102332130-2200011201202203-0303022113032130-3102020321120012-2032033023110132-3003103230322213-1013311120020132-3212012012310001"></a>

Type: `"object"`. single nested block, Optional.

GitHub Integration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("username")}
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
github {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303112330111312-1121230112323203-0303100233330302-2110021223123332-0320220230211112-2002122103222323-0221321121331133-3231021222100002"></a>

## Direct properties — GitHub / 033120021211 / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-0113123002131103-3101203021030002-1221123322121031-3332233121211221-0031222322333230-1130221000013320-0333200010220010-3223023122332120): complete subsection reference.

<a id="canonical-0100031002230310-2300202132020013-2002013122200310-3230121030121211-2311312301120001-0302132233131312-2213203130002101-0020001022320222"></a>

<a id="canonical-3012011320302200-1023030100001133-1103222303220233-3202333002023111-1120011223002210-2002121330012310-2320220203132010-3032110120320231"></a>

## username property — GitHub / 033120021211 / 4

Type: `"string"`. Optional.

GitHub Username. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-0320322210201000-0100110112101331-3333303322120202-0120031321322321-3311313003203013-2213111020213123-1230211333321231-0022212123321330"></a>

<a id="canonical-1112223120223311-0021113321211002-0001303002200321-1011021122022100-3221122010020300-1132120320201320-0321110312310330-1101122320132022"></a>

## verify_ssl property — GitHub / 033120021211 / 5

Type: `"bool"`. Optional.

GitHub Verify SSL. Configuration parameter for verify SSL

Upstream description:

Configuration parameter for verify SSL

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

<a id="canonical-3120211123322323-3222132233302323-2121000303122310-1000011132233322-0102101100102133-0202002100122001-2023131023122133-1120211213221310"></a>

## Next pages — GitHub / 033120021211 / 6

- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-0113123002131103-3101203021030002-1221123322121031-3332233121211221-0031222322333230-1130221000013320-0333200010220010-3223023122332120)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-0113123002131103-3101203021030002-1221123322121031-3332233121211221-0031222322333230-1130221000013320-0333200010220010-3223023122332120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013021033112211-2300331321112032-1332001113120232-1310003033112132-3101031031133111-0112322012300111-2031200230121012-2232222301231303"></a>

## code_base_integration.GitHub.access_token — access_token / 012313303323 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-0120221310020322-3030103310232300-0032130112020031-0033101303023130-0200232032210221-2113122232123220-3121000201013200-1022232122022023)
- code_base_integration.GitHub.access_token

<a id="canonical-0122103013012303-0222032100110212-3102023121023200-2122023311210201-2332200013121121-0200030032113113-1022203212100121-0100203103221322"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320302011010132-0311231230130001-3103002011310101-2310222221102321-1323102102023210-0010101300133001-2301032201333110-3000313330030032"></a>

## Direct properties — access_token / 012313303323 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0311120023132301-3323032100203223-0120322122003230-2323012001212002-1013112133033313-1300011332133121-2233023213330203-2311333310303312): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1313232102321211-3100022211301122-2020003032233231-2222212022310331-2122121112000013-0012110301030233-3201331300110320-0002001323131023): complete subsection reference.

<a id="canonical-3302103013232200-3001200221221010-0013030002300101-2231331132021020-1011002020311320-1221013022001133-2013021232032301-3021022112012131"></a>

## Next pages — access_token / 012313303323 / 4

- [code_base_integration.github.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-0311120023132301-3323032100203223-0120322122003230-2323012001212002-1013112133033313-1300011332133121-2233023213330203-2311333310303312)
- [code_base_integration.github.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1313232102321211-3100022211301122-2020003032233231-2222212022310331-2122121112000013-0012110301030233-3201331300110320-0002001323131023)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-0120221310020322-3030103310232300-0032130112020031-0033101303023130-0200232032210221-2113122232123220-3121000201013200-1022232122022023)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-0311120023132301-3323032100203223-0120322122003230-2323012001212002-1013112133033313-1300011332133121-2233023213330203-2311333310303312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303221303212122-2030222211203230-3013211231301123-0010003033131233-0302030003101033-3110023132300211-1111320031030120-2212033102103330"></a>

## code_base_integration.GitHub.access_token.blindfold_secret_info — blindfold_secret_info / 323002112033 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-0120221310020322-3030103310232300-0032130112020031-0033101303023130-0200232032210221-2113122232123220-3121000201013200-1022232122022023)
- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-0113123002131103-3101203021030002-1221123322121031-3332233121211221-0031222322333230-1130221000013320-0333200010220010-3223023122332120)
- code_base_integration.GitHub.access_token.blindfold_secret_info

<a id="canonical-0230210110201110-0323132233122222-2131330202211133-1003302130120231-1331330301310033-3320331200101323-2311100311021332-1210201233031222"></a>

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

<a id="canonical-3321233031321230-0111131110111230-0121033100303123-0211302000302300-3200212030303132-2123203231003311-3302121330222211-0222300233132133"></a>

## Direct properties — blindfold_secret_info / 323002112033 / 3

<a id="canonical-1213213132021221-2222113011210010-1232210132210203-3001230320012010-2212331102221122-3210232232113203-1012312323111021-1022132010230313"></a>

<a id="canonical-0310022021231202-2011010221111222-3232231203121132-1232220013110100-3232131030003113-1031332220113311-0323122221331110-0103130101111222"></a>

## decryption_provider property — blindfold_secret_info / 323002112033 / 4

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

<a id="canonical-2323112030331110-2033133113112211-0321032320331020-0211312103330120-3221221020131003-2223112132322212-0033220303012202-2210133222132023"></a>

<a id="canonical-2011120032220221-3232233112122110-1331102011230031-0303002011211233-2032322123031232-2131010111233111-0333311332203101-1033233230223030"></a>

## location property — blindfold_secret_info / 323002112033 / 5

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

<a id="canonical-2211333010110323-3130322101023200-0131033031112201-1113020310120011-2102130130003110-0321021203213002-2103223321131211-1130221022032331"></a>

<a id="canonical-3121003320123210-1332222103202101-0102032123001030-0303233132200000-1210013132311020-2201221100122103-1110302102223303-3321331121323213"></a>

## store_provider property — blindfold_secret_info / 323002112033 / 6

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

<a id="canonical-1321211033020002-3302233123032323-2011301131211323-0323020122032230-1213230203212232-1321220000133020-3203021000121320-0022032220322221"></a>

## Next pages — blindfold_secret_info / 323002112033 / 7

- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-0113123002131103-3101203021030002-1221123322121031-3332233121211221-0031222322333230-1130221000013320-0333200010220010-3223023122332120)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1313232102321211-3100022211301122-2020003032233231-2222212022310331-2122121112000013-0012110301030233-3201331300110320-0002001323131023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123223331301132-1102203022012100-2111300332031120-2013133310011233-1030333311021133-3011323232232120-3221112201212012-1013132100211312"></a>

## code_base_integration.GitHub.access_token.clear_secret_info — clear_secret_info / 122112303000 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.github](resources--code_base_integration--reference--group-001.md#canonical-0120221310020322-3030103310232300-0032130112020031-0033101303023130-0200232032210221-2113122232123220-3121000201013200-1022232122022023)
- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-0113123002131103-3101203021030002-1221123322121031-3332233121211221-0031222322333230-1130221000013320-0333200010220010-3223023122332120)
- code_base_integration.GitHub.access_token.clear_secret_info

<a id="canonical-0233031221111100-1210311312322311-3122300100120320-0032222301230333-0033130103312032-2202121223122213-3102130231020131-0221220030333211"></a>

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

<a id="canonical-1000100332323110-3121323231110022-3132231203021331-2021131322032222-1322103112102211-1022202200202002-2100011311102123-0332211000001000"></a>

## Direct properties — clear_secret_info / 122112303000 / 3

<a id="canonical-2203132333220132-1300302201023133-3002031233211131-1203113303101213-2231131013020211-0022333201222313-2030322100211101-1001100013003311"></a>

<a id="canonical-0202301120202033-1000132133231232-1011012011101021-3103130010221121-2131010301232131-1313100322031202-3220112010313303-0310302000030033"></a>

## provider_ref property — clear_secret_info / 122112303000 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1120100303033113-1303222103310100-3001302033012302-3321021312202232-2333231302033211-1300231202333020-0023030311313123-2002321103222130"></a>

<a id="canonical-0323132000211302-2320303312232002-0113332110121311-2133021210300032-1003323103131001-3222213020031003-1331310232320332-0122202300231111"></a>

## URL property — clear_secret_info / 122112303000 / 5

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

<a id="canonical-0123012012303000-0122202310210032-2233303102011230-2112203321210033-2122112221001212-2200233203130100-3022212121203210-1333332031212100"></a>

## Next pages — clear_secret_info / 122112303000 / 6

- [code_base_integration.github.access_token](resources--code_base_integration--reference--group-001.md#canonical-0113123002131103-3101203021030002-1221123322121031-3332233121211221-0031222322333230-1130221000013320-0333200010220010-3223023122332120)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1232323021020220-2213003210330030-2202310102321311-2102032013032021-2112021221021322-1202203303201020-1122202103122132-0312100133211130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110130212213222-3131010212022123-1102121323203233-2232031103120012-1032233021120321-1201300012010113-0010010110013322-0301210011012000"></a>

## code_base_integration.github_enterprise — github_enterprise / 212130132233 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- code_base_integration.github_enterprise

<a id="canonical-3313010321301003-1201110111000210-2133210332313212-1222030203320213-3202020031032232-3003222131132013-1312230032322333-2310302200303231"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for GitHub enterprise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hostname",
    "username")}
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
github_enterprise {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000202300112130-2210233222202113-2233130102123310-1310130332321232-1020003223002100-2322103330102200-1300021232133120-1123100001001023"></a>

## Direct properties — github_enterprise / 212130132233 / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-1100111033010231-2311210212133223-0101312322321320-3323203311102011-3301201030211232-3001030010111320-2332022303302311-3330120332132203): complete subsection reference.

<a id="canonical-2302131331022022-1033010222330232-2022101032102310-0012313230020233-3200301012312033-2132220031210230-0203203233020022-3203130211120231"></a>

<a id="canonical-2010100000330010-3311133021233233-2313331310112000-0013303103033303-1121133123023003-0300320312301132-3110322031002300-3333130321123222"></a>

## hostname property — github_enterprise / 212130132233 / 4

Type: `"string"`. Optional.

GitHub Hostname. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1003320322021113-2112332000301213-0213123113220023-3003101300021312-0032203320202220-1020001120222203-1303332023103320-3321203031100111"></a>

<a id="canonical-2300223021231031-3333223002233011-3232233312200320-3032000323210133-0012133031310302-2303330012111231-3202302301331233-3310300222310131"></a>

## username property — github_enterprise / 212130132233 / 5

Type: `"string"`. Optional.

GitHub Username. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-0113301123210213-3220112123111222-1321022033212103-3102300202000230-3020001333313120-0100130320331213-2323123213120100-2000021202122322"></a>

## Next pages — github_enterprise / 212130132233 / 6

- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-1100111033010231-2311210212133223-0101312322321320-3323203311102011-3301201030211232-3001030010111320-2332022303302311-3330120332132203)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1100111033010231-2311210212133223-0101312322321320-3323203311102011-3301201030211232-3001030010111320-2332022303302311-3330120332132203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200022312133331-0020032022100322-1211020232311312-0100102220120130-0130311120102122-2220300111332032-2323021220221121-2113033233123102"></a>

## code_base_integration.github_enterprise.access_token — access_token / 202122313001 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-1232323021020220-2213003210330030-2202310102321311-2102032013032021-2112021221021322-1202203303201020-1122202103122132-0312100133211130)
- code_base_integration.github_enterprise.access_token

<a id="canonical-3122003233231001-2232033121033021-2000022030120202-0133210111012002-2023132311130032-0102232012012220-2230331013010233-2113221010200010"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010120200121021-3211110012210201-2120130332010231-1022310311122213-2212201102211021-2003201302200301-1011200200021303-2233130212100113"></a>

## Direct properties — access_token / 202122313001 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1113113221100221-2213113212013130-3012001213323111-0130111110031130-2123223122231013-2120232001113201-0133011212003102-0000213321313222): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2012031101223122-1331032131303033-2101001211303010-3313012223211323-2311332013002123-3001233100010110-0300232103102200-0111332311232132): complete subsection reference.

<a id="canonical-0121133100011133-1033023313300123-0103231231233131-3102333213203213-0012012113201120-1020100210200003-2101132032200210-0110220330103102"></a>

## Next pages — access_token / 202122313001 / 4

- [code_base_integration.github_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1113113221100221-2213113212013130-3012001213323111-0130111110031130-2123223122231013-2120232001113201-0133011212003102-0000213321313222)
- [code_base_integration.github_enterprise.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2012031101223122-1331032131303033-2101001211303010-3313012223211323-2311332013002123-3001233100010110-0300232103102200-0111332311232132)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-1232323021020220-2213003210330030-2202310102321311-2102032013032021-2112021221021322-1202203303201020-1122202103122132-0312100133211130)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1113113221100221-2213113212013130-3012001213323111-0130111110031130-2123223122231013-2120232001113201-0133011212003102-0000213321313222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301112132110103-3013030201011213-2300100231221131-2021311213013331-2230303031021003-2113202311311311-3112312322202022-1223222230030030"></a>

## code_base_integration.github_enterprise.access_token.blindfold_secret_info — blindfold_secret_info / 112313103102 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-1232323021020220-2213003210330030-2202310102321311-2102032013032021-2112021221021322-1202203303201020-1122202103122132-0312100133211130)
- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-1100111033010231-2311210212133223-0101312322321320-3323203311102011-3301201030211232-3001030010111320-2332022303302311-3330120332132203)
- code_base_integration.github_enterprise.access_token.blindfold_secret_info

<a id="canonical-1320131001022213-1202013132021113-0121201102031332-3132103330101323-3202030132223230-2100210022323322-1012121311320231-2301122212001121"></a>

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

<a id="canonical-1321022202302000-3223011110122102-2120320223312010-0310000131331321-0231201222203131-3103012111113330-1302321032011033-1221123201203211"></a>

## Direct properties — blindfold_secret_info / 112313103102 / 3

<a id="canonical-2301113020212122-3122100231121133-0020111113011123-2201201010310121-0032210122021113-2011233031330132-3001213310003213-1333220123131132"></a>

<a id="canonical-1012222220210320-0223201330020202-2233313022303220-0101002230000122-1031321321232323-0332031312021131-2102122130211313-0030100102021310"></a>

## decryption_provider property — blindfold_secret_info / 112313103102 / 4

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

<a id="canonical-3301233020300033-3311330201011012-2102101012212003-3320030312102201-0003002313001320-2332031010222211-3331131013203103-3013203130012220"></a>

<a id="canonical-2031120112221031-3211113323023220-0310213111113030-2123222010013330-1223133120111220-3231011130031333-0100002113013232-2031103122001213"></a>

## location property — blindfold_secret_info / 112313103102 / 5

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

<a id="canonical-2032210000232023-0323333023131132-1332333000321020-3221101231132133-0301210333323330-2110023132012202-3320030131320212-3032013321213321"></a>

<a id="canonical-1313013111100002-0320332020021232-2300322023333010-2220032322330230-1121101112310231-2112031313232322-3233213132333333-0120110013030022"></a>

## store_provider property — blindfold_secret_info / 112313103102 / 6

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

<a id="canonical-1100231320003320-1030201130323200-2000301221103033-0233320333200211-2233132001321301-0101100120301011-0020120321132230-3212122201030212"></a>

## Next pages — blindfold_secret_info / 112313103102 / 7

- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-1100111033010231-2311210212133223-0101312322321320-3323203311102011-3301201030211232-3001030010111320-2332022303302311-3330120332132203)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-2012031101223122-1331032131303033-2101001211303010-3313012223211323-2311332013002123-3001233100010110-0300232103102200-0111332311232132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302100033101200-2103113021233103-1033103303322002-3223301131133201-3022132322333333-1322012033100022-2233210022011232-2331322312023110"></a>

## code_base_integration.github_enterprise.access_token.clear_secret_info — clear_secret_info / 303211300101 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.github_enterprise](resources--code_base_integration--reference--group-001.md#canonical-1232323021020220-2213003210330030-2202310102321311-2102032013032021-2112021221021322-1202203303201020-1122202103122132-0312100133211130)
- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-1100111033010231-2311210212133223-0101312322321320-3323203311102011-3301201030211232-3001030010111320-2332022303302311-3330120332132203)
- code_base_integration.github_enterprise.access_token.clear_secret_info

<a id="canonical-2231332203103321-2133121221112201-1231123312131300-1200003100323212-1203310012111321-3112333003132211-2113003133310213-2021232122233022"></a>

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

<a id="canonical-0030301101222121-0310110330222100-3302203122312123-3302311123232301-2312323122012111-3331011121011131-3020233210303221-3221333303030023"></a>

## Direct properties — clear_secret_info / 303211300101 / 3

<a id="canonical-2320033123100032-3013130231332313-3020312310200233-1223211003131321-1302203323102312-3011013112332331-1023112302233101-0003101213111322"></a>

<a id="canonical-0211110011121212-2003222322130222-0232323213312331-3313030003102101-3322332101130202-2110011032013020-1012213020012000-1202102100300313"></a>

## provider_ref property — clear_secret_info / 303211300101 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0100323032213020-0100123220133203-2020120300033133-2031002013112200-2223131220211322-2201322322113101-2122023300023112-2322023230030010"></a>

<a id="canonical-2203322312302121-2303322310031003-1220021003300220-3223023013032322-3230203302132211-1033223010232323-0133221103021031-3122323032132221"></a>

## URL property — clear_secret_info / 303211300101 / 5

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

<a id="canonical-3303330330203312-0303232033313202-3333213033010003-2032212310232120-1332320022310103-2232322220110010-2103310313132302-1221111032221302"></a>

## Next pages — clear_secret_info / 303211300101 / 6

- [code_base_integration.github_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-1100111033010231-2311210212133223-0101312322321320-3323203311102011-3301201030211232-3001030010111320-2332022303302311-3330120332132203)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-2021200022200103-1003301322023302-2022313230222201-2101232003211313-1001233301102023-1203233330221213-2310133231313303-3121331302210300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001310010231332-3030233002330231-1122102122031223-1231021132330233-1110311333322130-0233002100300231-2223002210013220-2133133013031302"></a>

## code_base_integration.GitLab — GitLab / 232230210201 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- code_base_integration.GitLab

<a id="canonical-1100130301111201-2122131221233013-3221101222300121-0302020012122111-0220320013003100-1021013223112202-1120120001200131-2123123302132022"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
gitlab {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223312303322223-3002013033120022-3112123302301101-3333200010200230-1113222001100231-2201123002022110-2032131323012220-3122232110122332"></a>

## Direct properties — GitLab / 232230210201 / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-0011110102122312-0110133233221113-0102333102211311-1301310303221123-1003313123003032-3012123110212312-1231113013331131-3233203333010233): complete subsection reference.

<a id="canonical-3201233012012102-1001310232033131-0312130330023013-0223100210013330-0101212030012031-2102113100032023-2032220332332223-1232223031212003"></a>

## Next pages — GitLab / 232230210201 / 4

- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-0011110102122312-0110133233221113-0102333102211311-1301310303221123-1003313123003032-3012123110212312-1231113013331131-3233203333010233)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-0011110102122312-0110133233221113-0102333102211311-1301310303221123-1003313123003032-3012123110212312-1231113013331131-3233203333010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023331210113003-2202202011233112-3120020131111232-0202222331013020-3112111203200311-0010233131122121-3210223010112000-1001030100020223"></a>

## code_base_integration.GitLab.access_token — access_token / 101120311101 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-2021200022200103-1003301322023302-2022313230222201-2101232003211313-1001233301102023-1203233330221213-2310133231313303-3121331302210300)
- code_base_integration.GitLab.access_token

<a id="canonical-1231030203033030-2220321013321311-2211003020030103-1331213212201202-1332110323032320-3013203330321222-2203110321210223-2312211010120231"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212121201131310-1213123233010010-0111012303322113-2210011202302332-1131320020230133-2302002313130320-3131232122032200-2301333121030230"></a>

## Direct properties — access_token / 101120311101 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1123300331120212-3233312211331132-2312231131023030-1321332220201022-1321130300212203-3121030311231301-1020123210131322-3213000002020002): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2223230232231002-3032022213200021-2132231111330302-0320121033112201-3033300002311300-2120331312020102-0021320010011031-2012010120001333): complete subsection reference.

<a id="canonical-2003300031211130-1322010200200013-3332231102223301-3310232202110213-0330200221002202-2031103011122210-2232112123111131-2010122123130333"></a>

## Next pages — access_token / 101120311101 / 4

- [code_base_integration.gitlab.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-1123300331120212-3233312211331132-2312231131023030-1321332220201022-1321130300212203-3121030311231301-1020123210131322-3213000002020002)
- [code_base_integration.gitlab.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2223230232231002-3032022213200021-2132231111330302-0320121033112201-3033300002311300-2120331312020102-0021320010011031-2012010120001333)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-2021200022200103-1003301322023302-2022313230222201-2101232003211313-1001233301102023-1203233330221213-2310133231313303-3121331302210300)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1123300331120212-3233312211331132-2312231131023030-1321332220201022-1321130300212203-3121030311231301-1020123210131322-3213000002020002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112101322231302-3120033033233313-2102001100122323-3111113032012301-3222330320012003-1001103231301130-1321223320003301-3330233312233022"></a>

## code_base_integration.GitLab.access_token.blindfold_secret_info — blindfold_secret_info / 122313310331 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-2021200022200103-1003301322023302-2022313230222201-2101232003211313-1001233301102023-1203233330221213-2310133231313303-3121331302210300)
- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-0011110102122312-0110133233221113-0102333102211311-1301310303221123-1003313123003032-3012123110212312-1231113013331131-3233203333010233)
- code_base_integration.GitLab.access_token.blindfold_secret_info

<a id="canonical-0303221123033031-0223021332213333-3120102322121021-3112313313331112-1103202232323003-2121203220332113-0233013301121100-3002223011303023"></a>

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

<a id="canonical-1200113222011233-3122320310121320-1230013001030331-1300233322002210-3123230210330021-3331022121312333-2023202021320223-0310331323121232"></a>

## Direct properties — blindfold_secret_info / 122313310331 / 3

<a id="canonical-2231223101310012-2321233331020201-0222130300100031-1202103302032020-0032213130111223-1102231132320213-0231112231011020-3313313002002231"></a>

<a id="canonical-1310122203021201-2013232030111233-3023321200102233-0023120321113001-3201223011122110-1022330222313101-2220101303232001-1221130002112321"></a>

## decryption_provider property — blindfold_secret_info / 122313310331 / 4

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

<a id="canonical-1001213120223333-1331113022221220-2002331103212200-3133321121200220-2220001312111200-0111330030332331-2230311013001333-1022300332322132"></a>

<a id="canonical-0301121033213003-1221312120223201-2331231012020131-3103311030013311-3321031002033131-1211002112210213-1112213200313003-2332022102321021"></a>

## location property — blindfold_secret_info / 122313310331 / 5

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

<a id="canonical-0321110000002132-2312321012322201-2012221332232302-1301313230232201-0100221323313312-0002202002112201-3301000132313322-3310100203201022"></a>

<a id="canonical-3121322301202223-0212331020011032-0033202130122030-3123301123302013-1211123101023221-0222103102123221-0200201033132101-2321030332001310"></a>

## store_provider property — blindfold_secret_info / 122313310331 / 6

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

<a id="canonical-1012203133320321-3233111033210112-2333020130001003-2221233332202010-2313301301200211-3003023110200320-3320011221233220-3010002203210333"></a>

## Next pages — blindfold_secret_info / 122313310331 / 7

- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-0011110102122312-0110133233221113-0102333102211311-1301310303221123-1003313123003032-3012123110212312-1231113013331131-3233203333010233)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-2223230232231002-3032022213200021-2132231111330302-0320121033112201-3033300002311300-2120331312020102-0021320010011031-2012010120001333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012223230231310-2020010220101201-0202221111330201-1300122322221122-0332233231001003-3023232233201003-3222321023001330-3210020213001022"></a>

## code_base_integration.GitLab.access_token.clear_secret_info — clear_secret_info / 233103210121 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.gitlab](resources--code_base_integration--reference--group-001.md#canonical-2021200022200103-1003301322023302-2022313230222201-2101232003211313-1001233301102023-1203233330221213-2310133231313303-3121331302210300)
- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-0011110102122312-0110133233221113-0102333102211311-1301310303221123-1003313123003032-3012123110212312-1231113013331131-3233203333010233)
- code_base_integration.GitLab.access_token.clear_secret_info

<a id="canonical-0211002233302201-1321102011321030-2131120313120133-1113223303301102-0022112022310202-0103112013233203-1032111103010303-3110300223223012"></a>

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

<a id="canonical-1131233023003232-0113210000210232-0221210230032032-1320123021232130-1223100133322121-2010303221202012-2200333230311313-3032101131333123"></a>

## Direct properties — clear_secret_info / 233103210121 / 3

<a id="canonical-0210133201032200-1202312331330023-0133130030031201-0303100230131220-1321220310233322-2201113102132230-0013121302202302-1230113300002122"></a>

<a id="canonical-3000222232323020-1003301233311023-2322110202300213-2110232313301000-0131121232122230-3232103132320211-1312011111231011-1131201231312013"></a>

## provider_ref property — clear_secret_info / 233103210121 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0311203013210103-3113232303333320-3130103322211233-1312231102010113-3020132323302112-2300020120221323-1123221213122301-0023303121011030"></a>

<a id="canonical-1111010011311023-0222312313323200-2213030130002110-2103122322201130-1003201232212210-3130023331111222-1112223123212112-2212312213320332"></a>

## URL property — clear_secret_info / 233103210121 / 5

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

<a id="canonical-0232003210103202-0220330013313313-1223120331203213-2122313201310201-3220132202001310-2232321000300021-2131111310331220-0020233030102323"></a>

## Next pages — clear_secret_info / 233103210121 / 6

- [code_base_integration.gitlab.access_token](resources--code_base_integration--reference--group-001.md#canonical-0011110102122312-0110133233221113-0102333102211311-1301310303221123-1003313123003032-3012123110212312-1231113013331131-3233203333010233)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-2300101201031033-0233121100211322-0323002120210212-3020213213320323-2323322312133130-2103132210131111-2130333232002130-1220301121211123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011312003031323-2201232311032103-2032122332133013-1121303330212133-2102113230331311-3323120321132023-1012233110332012-2110121100211322"></a>

## code_base_integration.gitlab_enterprise — gitlab_enterprise / 122021021022 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- code_base_integration.gitlab_enterprise

<a id="canonical-3332012310231021-3001321133001002-0333112132013022-1210323123213333-3020121030102001-1001231112023111-3020121020001120-2233332320001331"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for GitLab enterprise.

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
gitlab_enterprise {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221301232231213-1223100131310122-3210230120021113-1201010110220200-3120022000312300-2232200020030223-1211303233132221-0010010231211023"></a>

## Direct properties — gitlab_enterprise / 122021021022 / 3

- [access_token](resources--code_base_integration--reference--group-001.md#canonical-3212032121013322-0112030301122103-3330010000103200-1013100110132233-2322022303312123-1230311212332232-3222301012303332-1021320033031032): complete subsection reference.

<a id="canonical-0233101000302110-3002312302312113-3113000302022311-3131201213203032-0003333003333232-0103310223201121-0001223023303101-0112112013232222"></a>

<a id="canonical-1120232202332322-0111133331213321-0031321130223301-2110211020123121-3312012121122011-2330102100213310-3330230113003322-1122101012023132"></a>

## URL property — gitlab_enterprise / 122021021022 / 4

Type: `"string"`. Optional.

GitLab URL. URL or URI reference

Upstream description:

URL or URI reference

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`),
    ""),
}
```

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
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3103101023302103-3002332010321000-0100103220112012-0221012211300332-3021220133002123-3203331221333303-2202023030102111-1013200032102033"></a>

## Next pages — gitlab_enterprise / 122021021022 / 5

- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-3212032121013322-0112030301122103-3330010000103200-1013100110132233-2322022303312123-1230311212332232-3222301012303332-1021320033031032)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-3212032121013322-0112030301122103-3330010000103200-1013100110132233-2322022303312123-1230311212332232-3222301012303332-1021320033031032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100312221021301-0001133010231312-2300123300002213-2011213131203212-0330031121232132-3003131222321030-3232330010322011-1103122100210332"></a>

## code_base_integration.gitlab_enterprise.access_token — access_token / 212132211100 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-2300101201031033-0233121100211322-0323002120210212-3020213213320323-2323322312133130-2103132210131111-2130333232002130-1220301121211123)
- code_base_integration.gitlab_enterprise.access_token

<a id="canonical-0233322103303010-2301032313010230-0221002010000011-1122131020300211-2220312113110320-0001102122013023-1032330013312302-0301211311120210"></a>

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
access_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131330222001302-0322103120133111-2322002101132223-0100331013312123-2023010301313033-1213012110021331-1332220123032133-0112030213212120"></a>

## Direct properties — access_token / 212132211100 / 3

- [blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2032122221103120-3320113330102331-2101200303022232-1132323223101012-1002331331223032-2221321120203103-0031121033233001-0123101002231211): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2222200310110312-3132032101123132-1212021103230303-3002210311032302-3103331110202311-3002112103002221-1310212122013301-3022001003030321): complete subsection reference.

<a id="canonical-3130330121301113-0313232223001111-1320320131211032-0023320112011132-3323013220320020-2232130203012331-3112111033332131-1203331121300000"></a>

## Next pages — access_token / 212132211100 / 4

- [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2032122221103120-3320113330102331-2101200303022232-1132323223101012-1002331331223032-2221321120203103-0031121033233001-0123101002231211)
- [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](resources--code_base_integration--reference--group-001.md#canonical-2222200310110312-3132032101123132-1212021103230303-3002210311032302-3103331110202311-3002112103002221-1310212122013301-3022001003030321)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-2300101201031033-0233121100211322-0323002120210212-3020213213320323-2323322312133130-2103132210131111-2130333232002130-1220301121211123)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-2032122221103120-3320113330102331-2101200303022232-1132323223101012-1002331331223032-2221321120203103-0031121033233001-0123101002231211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212103321132102-1113031013120030-2020301213013323-1331210011222112-3331220320313332-1021200023030320-2120312231322211-0311231011130233"></a>

## code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info — blindfold_secret_info / 020201122211 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-2300101201031033-0233121100211322-0323002120210212-3020213213320323-2323322312133130-2103132210131111-2130333232002130-1220301121211123)
- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-3212032121013322-0112030301122103-3330010000103200-1013100110132233-2322022303312123-1230311212332232-3222301012303332-1021320033031032)
- code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info

<a id="canonical-0023302030303200-3202323212330221-1103130310022232-1332230320221220-3333210230333131-0120230220202212-2321111312101123-3103001333320201"></a>

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

<a id="canonical-3002103211113022-3322322012020102-0232132321312103-1121101030320123-0021030030132001-1022132010023023-0130103211222300-3202332233313313"></a>

## Direct properties — blindfold_secret_info / 020201122211 / 3

<a id="canonical-0113302330133133-3033103312110320-2101231110203323-3133222230221221-1031312021223032-1223131332121000-2301030300130203-2311023033000222"></a>

<a id="canonical-1011213102222030-2200321301112323-2312123213303020-0102332332113100-2213010002101210-0322112132131002-3331031011003133-1320312330132231"></a>

## decryption_provider property — blindfold_secret_info / 020201122211 / 4

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

<a id="canonical-3030210123022133-1120231013203030-3231231111022032-2103123221021022-0203211310133021-3310103131202213-2123231011331331-3003232220023123"></a>

<a id="canonical-3033100022033111-1212132011332110-2302203122203332-2220303332131011-2331000123111010-0301022121301302-2122320203000013-2220123123011032"></a>

## location property — blindfold_secret_info / 020201122211 / 5

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

<a id="canonical-0301321333230210-1030231333233012-2332011023013231-2031231001210133-3301023233011022-0330233013021112-3131222031030002-2233221030331230"></a>

<a id="canonical-2021031002333200-3003223202302002-3321033231203031-1111121021132030-3012103010302310-3120213112301110-3100110102202011-3211312300321110"></a>

## store_provider property — blindfold_secret_info / 020201122211 / 6

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

<a id="canonical-1122311131333303-1010110123011323-0301313032230311-2020132301122313-2111310131333111-0000300230102210-1013130003113322-0011321013122110"></a>

## Next pages — blindfold_secret_info / 020201122211 / 7

- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-3212032121013322-0112030301122103-3330010000103200-1013100110132233-2322022303312123-1230311212332232-3222301012303332-1021320033031032)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-2222200310110312-3132032101123132-1212021103230303-3002210311032302-3103331110202311-3002112103002221-1310212122013301-3022001003030321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202101302201103-3102022023313220-2123033133100103-3133011102101132-2131230200033000-3111100021330103-1232121023230000-1033021131211221"></a>

## code_base_integration.gitlab_enterprise.access_token.clear_secret_info — clear_secret_info / 022232302223 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [code_base_integration](resources--code_base_integration--reference--group-001.md#canonical-3112030122221330-0102012211000202-0010312121312200-0332112011122310-3223232312303223-2221012322320200-0333301332013311-3112322220223001)
- [code_base_integration.gitlab_enterprise](resources--code_base_integration--reference--group-001.md#canonical-2300101201031033-0233121100211322-0323002120210212-3020213213320323-2323322312133130-2103132210131111-2130333232002130-1220301121211123)
- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-3212032121013322-0112030301122103-3330010000103200-1013100110132233-2322022303312123-1230311212332232-3222301012303332-1021320033031032)
- code_base_integration.gitlab_enterprise.access_token.clear_secret_info

<a id="canonical-3112133332102123-1131011002210032-1212203203332212-2210011032120020-2200312032122213-0030310113102211-1222230021223231-3020333201312020"></a>

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

<a id="canonical-2011132313330210-0330210103022231-0302221220333232-1202101002313221-2323323330331110-3012013131321033-2110032232113013-1033321022122101"></a>

## Direct properties — clear_secret_info / 022232302223 / 3

<a id="canonical-1130320003210221-1111312103313021-1003211211030312-3220000101123101-3212132120230101-2111203202313010-2332222001002222-0100011002022323"></a>

<a id="canonical-2103120013130012-2332221332112033-3022300202013313-0320232013031031-2322020102131123-1232313330313110-0323321011310122-1113033022033120"></a>

## provider_ref property — clear_secret_info / 022232302223 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0100132210013223-3320302222111322-1332223300302303-0330212130120212-1231010221322312-3132321313231232-2013200301123310-1121022003023321"></a>

<a id="canonical-1313003033313203-0323330010131210-2222010033330331-0313211100033313-2130132133102323-2101032221301313-1310002100032210-0023313311233123"></a>

## URL property — clear_secret_info / 022232302223 / 5

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

<a id="canonical-2310031000122201-2101033113110311-2130213222023023-2132031103222010-2300303103302122-0332101322330113-0131313220322233-3313003333333001"></a>

## Next pages — clear_secret_info / 022232302223 / 6

- [code_base_integration.gitlab_enterprise.access_token](resources--code_base_integration--reference--group-001.md#canonical-3212032121013322-0112030301122103-3330010000103200-1013100110132233-2322022303312123-1230311212332232-3222301012303332-1021320033031032)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)

<a id="canonical-1001330330201111-0023201102213331-2220222121001222-0223210221230112-1120101313111223-3303312223222112-0220130220213110-0330212303030203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333101132132022-0100220223220120-2130301321000220-1031310010031301-3112131303223330-0022101233013013-3133132000132312-2031103213221130"></a>

## timeouts — timeouts / 122123210230 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- timeouts

<a id="canonical-1330131233002223-2133202331322133-3113202313130220-0313213201031030-0213113232323120-2123112233003230-3220330222121313-0331221233311222"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112302022201023-2221000221213032-3233310203122011-0003322133130301-1133202212122213-2221032332323131-1010023020122313-3031210313000221"></a>

## Direct properties — timeouts / 122123210230 / 3

<a id="canonical-3203331120121322-0213303322001333-2012221122233123-3201130330301010-0202310210230201-1333132310201333-1310320323112111-2313211020321033"></a>

<a id="canonical-3330213122013123-0303130212111131-2102020130003101-3022301303202203-0303222102101331-2101033203112230-1301022031120322-3221233223023101"></a>

## create property — timeouts / 122123210230 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1021302301202201-1120033300130103-1321213320122310-2320121210322131-1021022131010000-2301131111133213-3010023110303110-3001331221210231"></a>

<a id="canonical-3220222302100232-2001103303111131-2332013112202113-2213101103230331-2131233113111210-3102111132030300-0112013033000210-2331002321211001"></a>

## delete property — timeouts / 122123210230 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0022201130231211-2112230111300021-0210113311002312-1023223110311030-0030112230330123-3130102120021011-2230132303120112-1211120301312230"></a>

<a id="canonical-0223230110022333-3211210330123303-0013023212102023-3313222120330232-3200002320121033-0231130310333312-3130202030123333-0112201111302303"></a>

## read property — timeouts / 122123210230 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1303223310030233-1220212132103000-1110111300212232-3103202002232003-2030210320202222-3102310200120031-0110220003302232-1213122123231311"></a>

<a id="canonical-2223132103213312-2022213202033321-2123311300321330-2032312302130302-2221110230322031-0121333002032233-2011012023230030-0023030121100331"></a>

## update property — timeouts / 122123210230 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2310113123001000-1131102003330011-3103002223303113-0230010302011233-2113301132012212-1202303213300232-2200120221330101-1032302100002302"></a>

## Next pages — timeouts / 122123210230 / 8

- [Property reference](resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
