---
page_title: "xcsh_registration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration reference."
---

# xcsh_registration reference

<a id="canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- Property reference

<a id="canonical-3312002032123113-3321021102012123-1012001202013100-1302121312031321-0131220312311112-0202123032000323-2122222010303101-0010000210002222"></a>

### Direct properties for `xcsh_registration`

<a id="canonical-0001103113220113-3123013220003031-1201132201200201-2102122032211223-3121131230222213-3221321301013312-3001031021031212-0012021201332223"></a>

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

<a id="canonical-1300000311102312-1212212333002223-1033322202001230-2323330103323100-1012020303231221-0320120312233003-1321000201020313-2332311111333320"></a>

<a id="canonical-1111222200233033-3220032301033021-2312330330231103-3112202301122123-1320213331121311-3311001320130011-3233110223222133-0300201121232002"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0213111102222033-3232110230302003-3133320301230201-0021020031110313-3011223220003122-0330002212023010-2333321330231103-2230333111121302"></a>

<a id="canonical-2001033022320131-1221110333331133-2022110121202033-0212012020113130-0111320222003202-1222303022323032-1322101321302030-1212221133013033"></a>

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

<a id="canonical-2331121301110102-0012201232112131-1103111123323130-3033013302320022-1030332031130120-2000032233131201-2332000313322313-3113121101302021"></a>

<a id="canonical-3102110102012112-1003201131023222-3131131313010321-2121331102101120-0313002033130310-0101113231021300-1301011332322033-3230030211212110"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111): complete subsection reference.

<a id="canonical-3103123323021200-0213013032103201-0002223013012121-2221020023232323-0313012113311030-1323133112001013-2122002203030120-1120211230200033"></a>

<a id="canonical-0332311010210211-3101311310133321-2030121130322010-1101200232112103-0132221030303300-3201030233012323-1100303020332221-1013130022112030"></a>

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

<a id="canonical-2320033113101223-0322300012111230-2220323322330213-2003211213210313-0030212003303313-1003132232213312-1313223022230010-2211310310123102"></a>

<a id="canonical-0330222013302211-3311100113321113-1322320113202023-1302201322212021-3331210322021022-0112012102111223-3200133113230331-0020033302120103"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Registration. Must be unique within the namespace.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2311212321311133-3131230100220120-0201110113311332-1003032333123001-2130031031211110-0031130130230111-0130031131322102-0112233312313222"></a>

<a id="canonical-1120003231120113-1301321133123111-3100033001030200-1133002033201202-0230130111333321-2213123210302023-1123213300122220-0220101230002100"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Registration is created.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [passport](resources--registration--reference--group-001.md#canonical-1120101321201223-0133323222022210-1001000011103111-2322312323001321-3201033220310002-2220113311003213-3101110103313231-1211213320021213): complete subsection reference.

- [timeouts](resources--registration--reference--group-001.md#canonical-3213031231333002-0230321022300100-0023012110121303-3101201022223132-2020331103120330-0112000313011320-2032122330213002-1312312001121102): complete subsection reference.

<a id="canonical-2233233122220320-2231011130200311-3321103022031013-2010332002221231-3222321311002110-1031012310131112-0331111300323201-1132101310121201"></a>

<a id="canonical-3102301031312122-3320001011300210-3012021312021320-0111313033013022-1033333121023330-1311131012321100-1200330113123111-1010130012021223"></a>

#### `token` property

Type: `"string"`. Required.

Token is used for machine and tenant identification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "security",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 20
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

<a id="canonical-1201210103311021-0202202030300003-2200000102210330-1010033102032022-0120113301023310-3110013312302112-1210020010032013-2223022211102213"></a>

### All schema paths for `xcsh_registration`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--registration--reference--group-001.md#canonical-0001103113220113-3123013220003031-1201132201200201-2102122032211223-3121131230222213-3221321301013312-3001031021031212-0012021201332223) |
| `description` | [description](resources--registration--reference--group-001.md#canonical-1300000311102312-1212212333002223-1033322202001230-2323330103323100-1012020303231221-0320120312233003-1321000201020313-2332311111333320) |
| `disable` | [disable](resources--registration--reference--group-001.md#canonical-0213111102222033-3232110230302003-3133320301230201-0021020031110313-3011223220003122-0330002212023010-2333321330231103-2230333111121302) |
| `id` | [ID](resources--registration--reference--group-001.md#canonical-2331121301110102-0012201232112131-1103111123323130-3033013302320022-1030332031130120-2000032233131201-2332000313322313-3113121101302021) |
| `infra` | [infra](resources--registration--reference--group-001.md#canonical-0333202120232121-0302332001312301-2302131021201221-2331300122122012-1010112021230303-0011330200002231-3021332113332312-1312121130112213) |
| `infra.availability_zone` | [infra.availability_zone](resources--registration--reference--group-001.md#canonical-1103310230111111-3000023310022230-0222303122130301-0212031323030221-1313201002310100-0203133112023211-3203332130331201-0312110233322111) |
| `infra.bond_config` | [infra.bond_config](resources--registration--reference--group-001.md#canonical-0113010113111223-0201010222323221-0202013223212030-1022221211101233-0213331301332213-1202232130321123-2012333303320301-3220202212223030) |
| `infra.bond_config.interfaces` | [infra.bond_config.interfaces](resources--registration--reference--group-001.md#canonical-0033131331300101-0103233013010210-3231303200131232-1003310230203202-3313221201013222-1311213113023323-3002110010200121-3112030323300303) |
| `infra.bond_config.mode` | [infra.bond_config.mode](resources--registration--reference--group-001.md#canonical-2012122221000220-2113222303131002-3003132203330331-3301321012103203-0113032131313022-0011121022122112-0001123211001011-0203222030003101) |
| `infra.bond_config.name` | [infra.bond_config.name](resources--registration--reference--group-001.md#canonical-3332101322323330-1012123303003013-2130001102032112-2200031332212311-2200133110201133-0132310010202010-0002023101020012-1321003023301133) |
| `infra.certified_hw` | [infra.certified_hw](resources--registration--reference--group-001.md#canonical-2132121300101012-2331201231011222-1013331311313020-0233000220100330-3202223023030302-0033022211101010-1031130030310322-1132103202103200) |
| `infra.domain` | [infra.domain](resources--registration--reference--group-001.md#canonical-2213012130320323-0013021201331100-2000012213221230-1123223001001031-0230100211210011-2211103332320232-2122221312200313-3203011221310011) |
| `infra.hostname` | [infra.hostname](resources--registration--reference--group-001.md#canonical-1233222201331213-3131013303213003-0321011221033121-2132111312113330-1010120323013110-0100130302212331-3211211333213001-2133123123002020) |
| `infra.hugepages` | [infra.hugepages](resources--registration--reference--group-001.md#canonical-2032301231212002-0130113300121021-3201303012312333-2213030310203110-3332300333011022-0003120111200223-3320322010013320-3310331320003032) |
| `infra.hugepages.free` | [infra.hugepages.free](resources--registration--reference--group-001.md#canonical-2232222223303222-1000111323010030-0303013310230132-2330202001211222-0130002133220300-0011211123212321-3131201332030121-0111321233030121) |
| `infra.hugepages.page_size` | [infra.hugepages.page_size](resources--registration--reference--group-001.md#canonical-1033221321230222-1211132231101220-0110112130320013-2211012120110303-1310113121111322-3231220220333031-1230011113332120-0001330200303300) |
| `infra.hugepages.total` | [infra.hugepages.total](resources--registration--reference--group-001.md#canonical-1012322003132200-2301121003113022-2232111331113002-2111233232111221-0321302222231233-2211312032032210-2021332303001311-1031121123332131) |
| `infra.hw_info` | [infra.hw_info](resources--registration--reference--group-001.md#canonical-2112213100301233-3101013330101230-1022101030230130-0131131332111320-0300201233212102-3311030001002311-1002122022301120-1131221130012220) |
| `infra.hw_info.bios` | [infra.hw_info.bios](resources--registration--reference--group-001.md#canonical-2212201011120332-1302010023213312-2330200031022212-2132113201301000-0311031012111232-0130011020303333-3101113321213200-1230223010233232) |
| `infra.hw_info.bios.date` | [infra.hw_info.bios.date](resources--registration--reference--group-001.md#canonical-0211031321232321-3230131100022313-3000300001030010-2210132123113303-0321231030312330-3321313311220013-1323230100313230-2103121100100110) |
| `infra.hw_info.bios.vendor` | [infra.hw_info.bios.vendor](resources--registration--reference--group-001.md#canonical-1002210201323222-1013223220222213-0103032333201030-3332300200322202-1030200303010112-1322002113311331-1203113003233232-3221310012112031) |
| `infra.hw_info.bios.version` | [infra.hw_info.bios.version](resources--registration--reference--group-001.md#canonical-2010100331021301-0233213111333303-2223323033020003-3132023000131022-2311012301321120-2100230111310303-2300200002133311-1010313301023113) |
| `infra.hw_info.board` | [infra.hw_info.board](resources--registration--reference--group-001.md#canonical-2232033112110232-1032310110200220-0230203020333312-1221113001301112-0000300222022131-3331332012332033-2312122103000213-1131012220103333) |
| `infra.hw_info.board.asset_tag` | [infra.hw_info.board.asset_tag](resources--registration--reference--group-001.md#canonical-0021103310103033-3202313020020120-2013312133121332-3203013023022321-2122022012211311-3320310322030322-1331100211100230-2330101102300333) |
| `infra.hw_info.board.name` | [infra.hw_info.board.name](resources--registration--reference--group-001.md#canonical-1022000301222100-1033211221102303-2003132013220211-1332210102302230-1201031012031323-0033130233110200-3101032201320212-1203303131021233) |
| `infra.hw_info.board.serial` | [infra.hw_info.board.serial](resources--registration--reference--group-001.md#canonical-1133002100233031-0033001331310033-0033320201232202-3122031303222121-3132221221332031-0331311332330102-2210232001010213-2231200120012121) |
| `infra.hw_info.board.vendor` | [infra.hw_info.board.vendor](resources--registration--reference--group-001.md#canonical-1132203131303230-1203031113000121-1332011022331122-2302000321000231-1320111301200031-0222130322123120-1012323212113222-3103013033213021) |
| `infra.hw_info.board.version` | [infra.hw_info.board.version](resources--registration--reference--group-001.md#canonical-3333023100133001-0322010131203313-0000002212330110-3133323301312131-3023233011231102-0120011230023000-3222131110210033-0013010033023033) |
| `infra.hw_info.chassis` | [infra.hw_info.chassis](resources--registration--reference--group-001.md#canonical-1312011201130301-1000001021130121-2110221332032123-0221321032102222-1023033333303211-0021120022122223-3210033322230332-0023203222121002) |
| `infra.hw_info.chassis.asset_tag` | [infra.hw_info.chassis.asset_tag](resources--registration--reference--group-001.md#canonical-1303322302012302-0232223330110210-2302333233302130-2202113322301120-3202301300323333-2333021121322013-3130202310021320-1022002132000102) |
| `infra.hw_info.chassis.serial` | [infra.hw_info.chassis.serial](resources--registration--reference--group-001.md#canonical-3232000011322013-0313313131201223-1232303201031322-2131330222020310-2011033112132121-0123000100112311-2212222132000000-2231102302033030) |
| `infra.hw_info.chassis.type` | [infra.hw_info.chassis.type](resources--registration--reference--group-001.md#canonical-2010031112023330-1303310110322133-0330133320011121-1023203023301021-0122330022302112-2310012300330222-2023001202130302-3200232332113112) |
| `infra.hw_info.chassis.vendor` | [infra.hw_info.chassis.vendor](resources--registration--reference--group-001.md#canonical-2012032002122121-1230013202013321-3031331332311022-3002213023221232-0020321203303313-0011333233300033-1113312033112110-3010221021302201) |
| `infra.hw_info.chassis.version` | [infra.hw_info.chassis.version](resources--registration--reference--group-001.md#canonical-3320320003100221-1302002112022322-3132130003302313-3132203000213111-2113323031232302-3333120023331120-0210100003213301-3121212132322322) |
| `infra.hw_info.cpu` | [infra.hw_info.cpu](resources--registration--reference--group-001.md#canonical-0230323221112310-0333213013033222-3133000300123023-3111312231103123-3103032322121323-0112203111031130-1100021000023313-2011000201213132) |
| `infra.hw_info.cpu.cache` | [infra.hw_info.cpu.cache](resources--registration--reference--group-001.md#canonical-3322201112220101-1212011003131022-2011100013110113-3320223000310021-3303110321123123-0132213220030331-1032333320212002-1031322120233302) |
| `infra.hw_info.cpu.cores` | [infra.hw_info.cpu.cores](resources--registration--reference--group-001.md#canonical-2031332000300222-1013230132213203-0321320211211033-1202023122222101-1012001301020312-2010113111301303-3113030020330333-2213313333021033) |
| `infra.hw_info.cpu.cpus` | [infra.hw_info.cpu.cpus](resources--registration--reference--group-001.md#canonical-3311201221223122-0201013232113221-0221230212103111-0001302333132210-2111030122201031-3231220311203111-0221213303011002-3311310323233032) |
| `infra.hw_info.cpu.model` | [infra.hw_info.cpu.model](resources--registration--reference--group-001.md#canonical-0212330332011212-3001100110202322-0221330132310203-2003301113232211-3322001321311333-3030331230101102-1002332302020232-2103323000101103) |
| `infra.hw_info.cpu.speed` | [infra.hw_info.cpu.speed](resources--registration--reference--group-001.md#canonical-0230012313033133-1003033333230213-1223030210002003-2223033132302320-1110010120313320-1220023323103233-3301233231220311-1101103013111320) |
| `infra.hw_info.cpu.threads` | [infra.hw_info.cpu.threads](resources--registration--reference--group-001.md#canonical-0013212013001310-0330112003300212-0102030320010200-1330121300103120-0323132022331213-1021022033221320-0321323000221113-2131030102123121) |
| `infra.hw_info.cpu.vendor` | [infra.hw_info.cpu.vendor](resources--registration--reference--group-001.md#canonical-1133203231332103-2002202333210203-1313332333021232-0030022102123331-3122101032001230-2300121232013132-3230322121220122-3021212003002112) |
| `infra.hw_info.gpu` | [infra.hw_info.gpu](resources--registration--reference--group-001.md#canonical-3300011012312122-1033321001030202-3201231000233211-3210213321231020-3100130212233232-3330000030321333-1120103000030312-1320111303211031) |
| `infra.hw_info.gpu.cuda_version` | [infra.hw_info.gpu.cuda_version](resources--registration--reference--group-001.md#canonical-3202032203111220-2233233011111000-0100120220311113-2202020033000030-1221122322022133-3231103211003321-1131232131110331-1101303110300120) |
| `infra.hw_info.gpu.driver_version` | [infra.hw_info.gpu.driver_version](resources--registration--reference--group-001.md#canonical-3010211223032133-1111002210301230-3222311102323033-3110222301232111-2330110320101210-2100103102301012-2211221011311121-0132330321331221) |
| `infra.hw_info.gpu.gpu_device` | [infra.hw_info.gpu.gpu_device](resources--registration--reference--group-001.md#canonical-0220322132002100-1333020113232303-3103130310221322-0112003232111110-3003321201110031-2000100013111102-1312232232221133-0100033233111011) |
| `infra.hw_info.gpu.gpu_device.id` | [infra.hw_info.gpu.gpu_device.id](resources--registration--reference--group-001.md#canonical-2212310323011023-0223013223320111-0330110321200201-2011131023201033-2220101332230310-1030121221212120-2032332033010302-0212330222301211) |
| `infra.hw_info.gpu.gpu_device.processes` | [infra.hw_info.gpu.gpu_device.processes](resources--registration--reference--group-001.md#canonical-3010031002313103-3111103100232023-3002130311303301-1123020202330033-3322022332113110-1003001130233230-1222300201020211-1223012003323133) |
| `infra.hw_info.gpu.gpu_device.product_name` | [infra.hw_info.gpu.gpu_device.product_name](resources--registration--reference--group-001.md#canonical-0133010223013201-3123123023111231-2302123202132103-3032302100020200-1001331101030011-0003013102210311-0303112023230300-2010330210013210) |
| `infra.hw_info.kernel` | [infra.hw_info.kernel](resources--registration--reference--group-001.md#canonical-1230220122001331-1220332023233010-3332302130211101-3231321330301233-2322211101123033-0020012233301120-0013223030231030-0200020122012031) |
| `infra.hw_info.kernel.architecture` | [infra.hw_info.kernel.architecture](resources--registration--reference--group-001.md#canonical-0221112033103123-2120300322322232-2322122122330310-2102213130303310-2230030032103012-3123321322122111-0132101021231103-2330323111332320) |
| `infra.hw_info.kernel.release` | [infra.hw_info.kernel.release](resources--registration--reference--group-001.md#canonical-2123003202302011-2220001331100100-2300311213311000-1121323030213130-1211010301133332-2031322233211112-0121211012013131-2332022200310113) |
| `infra.hw_info.kernel.version` | [infra.hw_info.kernel.version](resources--registration--reference--group-001.md#canonical-1003321300020132-0302102021221101-3133323333300001-2211121221131300-2020011023122221-0031210301203320-0113013320320100-2201122011003233) |
| `infra.hw_info.memory` | [infra.hw_info.memory](resources--registration--reference--group-001.md#canonical-2000310221113132-1032101333023220-2030221022013001-2332123121030023-0020013131020200-1220123210332223-2131221020331021-0020031311030330) |
| `infra.hw_info.memory.size_mb` | [infra.hw_info.memory.size_mb](resources--registration--reference--group-001.md#canonical-0023120331133221-2121321000023032-2031112331012312-3023002231332001-1020013101223223-1102303002221102-0020112320012121-0223313303220030) |
| `infra.hw_info.memory.speed` | [infra.hw_info.memory.speed](resources--registration--reference--group-001.md#canonical-3112303002301331-1121021202120221-2203000222010203-1031001021301123-0013022101220131-2310132100302201-0000130302003330-0003220331220003) |
| `infra.hw_info.memory.type` | [infra.hw_info.memory.type](resources--registration--reference--group-001.md#canonical-2130012101010320-3212222332302020-2112010103021120-3303311231013301-3011200322011200-0301012223020021-2213101023103323-3200233322110021) |
| `infra.hw_info.network` | [infra.hw_info.network](resources--registration--reference--group-001.md#canonical-0211210001022121-1230303301321310-2102233223301110-2012110130030103-1232113012213003-0321210102232210-0313120301230301-0032231222302323) |
| `infra.hw_info.network.driver` | [infra.hw_info.network.driver](resources--registration--reference--group-001.md#canonical-3203331230202030-2332023321302233-3022110012332212-2321003002332200-1030202002130213-1333210001312130-3312202133101021-0030331333312221) |
| `infra.hw_info.network.ip_address` | [infra.hw_info.network.ip_address](resources--registration--reference--group-001.md#canonical-0010323211023302-3330233211203010-3232221200231022-2003223021003332-3003012321223212-0121223103230002-2330211002131033-3113302112331030) |
| `infra.hw_info.network.link_quality` | [infra.hw_info.network.link_quality](resources--registration--reference--group-001.md#canonical-3223032320311002-0321303000032131-0032220120330302-1013032102132000-3031320300303112-2321320001022201-0233013012221202-1330222021101213) |
| `infra.hw_info.network.link_type` | [infra.hw_info.network.link_type](resources--registration--reference--group-001.md#canonical-0313132120211102-2310033011310301-0230221312013112-3023031313011003-1120300112310233-2301323120302310-1233231230013133-3211011232221101) |
| `infra.hw_info.network.mac_address` | [infra.hw_info.network.mac_address](resources--registration--reference--group-001.md#canonical-2203301131111000-1130020200233213-1200021220010211-3301301020201211-2312300102133302-1230021323122212-3113300322121023-3222001223203313) |
| `infra.hw_info.network.name` | [infra.hw_info.network.name](resources--registration--reference--group-001.md#canonical-3000323122232333-0223131333131020-1023130031311112-3032033320323101-1301232123223012-3003221130301313-1033003303101331-2220002122132023) |
| `infra.hw_info.network.port` | [infra.hw_info.network.port](resources--registration--reference--group-001.md#canonical-3221112000321330-0231011010202303-2300201232023200-2100032132320313-0030221013212222-2100023231231331-3103032213322003-1111231210001221) |
| `infra.hw_info.network.speed` | [infra.hw_info.network.speed](resources--registration--reference--group-001.md#canonical-2203120022030213-1133232112000002-0330313031231122-3220011303112303-3012223302221021-2031122113221222-2201231032300021-2313300012002310) |
| `infra.hw_info.numa_nodes` | [infra.hw_info.numa_nodes](resources--registration--reference--group-001.md#canonical-2220103100103110-0121220231332201-2212302322300201-1102311200300202-0233331102322301-1331212101100202-2211012232033031-0300202032313112) |
| `infra.hw_info.os` | [infra.hw_info.os](resources--registration--reference--group-001.md#canonical-1231002013131030-1012200120133232-1013121131001033-1330312002012230-1310121133131001-2302321110031022-3210200302322210-3131022121331302) |
| `infra.hw_info.os.architecture` | [infra.hw_info.os.architecture](resources--registration--reference--group-001.md#canonical-1033132130023221-0210011231332030-3011013301013100-3212111212020113-0121213020203100-2111113110233230-3311003232113131-2130330111333131) |
| `infra.hw_info.os.name` | [infra.hw_info.os.name](resources--registration--reference--group-001.md#canonical-3002311233011223-0001330221310211-0032002201021110-0210212102320311-1301200230012021-2102203323323231-1033012313012013-3323233223210100) |
| `infra.hw_info.os.release` | [infra.hw_info.os.release](resources--registration--reference--group-001.md#canonical-3303211031231023-1302112133010200-3101102031010200-1110333222320000-0011023300103331-0333311132330230-0201013032230133-1103301130000033) |
| `infra.hw_info.os.vendor` | [infra.hw_info.os.vendor](resources--registration--reference--group-001.md#canonical-1131110322033201-0030120010000100-0032132120020110-0031332222120210-2133223023301330-1210203113131120-1131133121111122-0113301210030031) |
| `infra.hw_info.os.version` | [infra.hw_info.os.version](resources--registration--reference--group-001.md#canonical-3033332132221230-1013231000122301-3331103201300131-1002233303133233-1312303032033333-2331311231113233-3301210132023301-1212322323233321) |
| `infra.hw_info.product` | [infra.hw_info.product](resources--registration--reference--group-001.md#canonical-0223022230303020-1030223112321332-3012233303131121-1113023212121213-2201230020002102-1231220012313210-0300132032223303-2200332121331001) |
| `infra.hw_info.product.name` | [infra.hw_info.product.name](resources--registration--reference--group-001.md#canonical-3320201133032300-2220212120221011-0212230233302223-1210130032100222-2003130202223032-0132223330213122-3030323131013312-2033113210123231) |
| `infra.hw_info.product.serial` | [infra.hw_info.product.serial](resources--registration--reference--group-001.md#canonical-3020212003232113-2032011200122001-0322011333021323-2020102032011132-2031203212123332-0111102133111001-0332332233123332-1103323032321110) |
| `infra.hw_info.product.vendor` | [infra.hw_info.product.vendor](resources--registration--reference--group-001.md#canonical-2002032011330333-0123132102111223-2030113310223021-3330031020100212-3220303212232300-3101031322001230-3231011101032101-1231033331310231) |
| `infra.hw_info.product.version` | [infra.hw_info.product.version](resources--registration--reference--group-001.md#canonical-1110330101311202-1010002111332003-2221231100123123-1312003203121110-1233321332123320-3032113133033000-2221023323030302-1232020332121310) |
| `infra.hw_info.storage` | [infra.hw_info.storage](resources--registration--reference--group-001.md#canonical-2210002233111210-2021333332211311-3230033103002002-2023333320010232-3223200201021120-3002222231113303-2232220201223331-2320301101331221) |
| `infra.hw_info.storage.driver` | [infra.hw_info.storage.driver](resources--registration--reference--group-001.md#canonical-0222212030312002-3213030333033020-1222220311211302-3223232001301113-0233303002113013-0131330202321033-3132230302130111-0022232021121033) |
| `infra.hw_info.storage.model` | [infra.hw_info.storage.model](resources--registration--reference--group-001.md#canonical-3133210030133321-3230130013211233-1001032303002010-2133302132213200-3023302113100332-2033130021022211-3112010111232302-3300131200322311) |
| `infra.hw_info.storage.name` | [infra.hw_info.storage.name](resources--registration--reference--group-001.md#canonical-3033200311311233-1322111110102233-2313231012212123-1333331210013302-1010000330302001-3201313001021121-0301113131212130-0330002203132212) |
| `infra.hw_info.storage.serial` | [infra.hw_info.storage.serial](resources--registration--reference--group-001.md#canonical-1222013101030312-3002331213123301-1030220202131112-2002311322231231-3011111210312303-0300231003000310-1103313001000333-3211002213200311) |
| `infra.hw_info.storage.size_gb` | [infra.hw_info.storage.size_gb](resources--registration--reference--group-001.md#canonical-3123232121232020-1123121002310223-3011212330033131-3220200321221013-1022321101221213-2210210221220230-0111130101010131-0001113000201011) |
| `infra.hw_info.storage.vendor` | [infra.hw_info.storage.vendor](resources--registration--reference--group-001.md#canonical-2101230313213023-0023012031212112-2333232333130212-2023023013030212-1002132330002210-0333122003003223-0122211200030232-3003123221000032) |
| `infra.hw_info.usb` | [infra.hw_info.usb](resources--registration--reference--group-001.md#canonical-3231321232012313-2312132120033121-1202131002332203-3232103331033212-1210301032021032-1201222312123202-2220313303210231-1322311111022311) |
| `infra.hw_info.usb.address` | [infra.hw_info.usb.address](resources--registration--reference--group-001.md#canonical-3113013010223302-1102301210133031-3321213132023302-3330312230110130-0111221212233210-1031313033222111-1100100222221222-1011123322201001) |
| `infra.hw_info.usb.b_device_class` | [infra.hw_info.usb.b_device_class](resources--registration--reference--group-001.md#canonical-0231322113220120-0032133223320131-2322210333332332-0221103002101320-2030312031212210-1020121010312313-0213023232311302-2331122323333212) |
| `infra.hw_info.usb.b_device_protocol` | [infra.hw_info.usb.b_device_protocol](resources--registration--reference--group-001.md#canonical-1111031102110312-3102323001211213-3222110011211311-0110113321022213-2103123103332031-0120112003203021-0103202002003011-2032310211111330) |
| `infra.hw_info.usb.b_device_sub_class` | [infra.hw_info.usb.b_device_sub_class](resources--registration--reference--group-001.md#canonical-1030112323233230-1330313132330122-1312313021331113-3001003322003133-2003010312131022-0020201103321033-1101112103323103-0111310011130302) |
| `infra.hw_info.usb.b_max_packet_size` | [infra.hw_info.usb.b_max_packet_size](resources--registration--reference--group-001.md#canonical-1103122201123220-2132323032202002-0213302120011010-2033011012311211-2122110023010132-3002020021122121-3023000203132110-1003022110202113) |
| `infra.hw_info.usb.bcd_device` | [infra.hw_info.usb.bcd_device](resources--registration--reference--group-001.md#canonical-1332330322221322-3101112130320212-3311301202320122-2323031022232030-2012303122131001-1010322332012320-3210202120211112-0331022022122310) |
| `infra.hw_info.usb.bcd_usb` | [infra.hw_info.usb.bcd_usb](resources--registration--reference--group-001.md#canonical-0002330022032202-1013010030303201-1302313010130121-2310210002103211-0133010131301122-3033003012233211-1223333002231023-3003121321233313) |
| `infra.hw_info.usb.bus` | [infra.hw_info.usb.bus](resources--registration--reference--group-001.md#canonical-0023232003222321-2022031122123102-3232103330001032-0010222023320123-2123102232310333-0002332101201010-1320211003313131-3122332001100210) |
| `infra.hw_info.usb.description_spec` | [infra.hw_info.usb.description_spec](resources--registration--reference--group-001.md#canonical-3102130331112233-2221131012230333-2012202231112122-2303112211311013-2000113301202101-1300123000033331-2033330131203002-3301110133330023) |
| `infra.hw_info.usb.i_manufacturer` | [infra.hw_info.usb.i_manufacturer](resources--registration--reference--group-001.md#canonical-1031220202230300-1102311303021023-0102211302003003-0221022311023020-0200302111331110-1301031013031212-2030101010010120-1203300121031002) |
| `infra.hw_info.usb.i_product` | [infra.hw_info.usb.i_product](resources--registration--reference--group-001.md#canonical-2200303203013221-3011223030020302-3301231013003300-2013220301131211-1031301310330231-2102033313002112-0132123310003133-3133020202211013) |
| `infra.hw_info.usb.i_serial` | [infra.hw_info.usb.i_serial](resources--registration--reference--group-001.md#canonical-1120223203232200-2011233333202030-0023033132100233-3113321330132320-1203031130212212-0113123221013201-2331123031131023-0113100010223112) |
| `infra.hw_info.usb.id_product` | [infra.hw_info.usb.id_product](resources--registration--reference--group-001.md#canonical-0222001113120202-1020230122222100-2301013223101010-1112013303021313-1300130012012232-0312103132002102-3222011010100100-0033300110022201) |
| `infra.hw_info.usb.id_vendor` | [infra.hw_info.usb.id_vendor](resources--registration--reference--group-001.md#canonical-1131020001212132-0100213323002212-3322301010233323-3233333321333101-3033110203321203-0013211231232121-0301110013322213-0300230011100113) |
| `infra.hw_info.usb.port` | [infra.hw_info.usb.port](resources--registration--reference--group-001.md#canonical-1012312020201200-2301132302113132-2032310032100000-0023020010323132-3103321020222203-0201102232231111-0312331223012130-1020312202111312) |
| `infra.hw_info.usb.product_name` | [infra.hw_info.usb.product_name](resources--registration--reference--group-001.md#canonical-2023111331223130-3213223100013232-3121032131332001-2311310102303121-1131232311002332-3233231201221023-1032022231100030-3031113133311223) |
| `infra.hw_info.usb.speed` | [infra.hw_info.usb.speed](resources--registration--reference--group-001.md#canonical-1031011333310320-1201000300302133-1213020303020010-3002103302022311-0112202300203311-1222121232111003-0323231031310210-1210331021101212) |
| `infra.hw_info.usb.usb_type` | [infra.hw_info.usb.usb_type](resources--registration--reference--group-001.md#canonical-1230302310103203-3221020230003021-3133203100313001-3213001323120210-0132302132023201-1332102013032030-0020211313320232-0000213213003212) |
| `infra.hw_info.usb.vendor_name` | [infra.hw_info.usb.vendor_name](resources--registration--reference--group-001.md#canonical-0213112132023322-0010001100110312-3313020320321123-3023131233230131-2002212131031233-3031322301102202-0323100023020223-1311223200000100) |
| `infra.instance_id` | [infra.instance_id](resources--registration--reference--group-001.md#canonical-0233233011200103-3201330302211210-0130222200331100-2213301332233002-1130321212123332-2113112222123123-0100033312010223-3133303333101211) |
| `infra.interfaces` | [infra.interfaces](resources--registration--reference--group-001.md#canonical-2331111231202111-2121032333223302-0120032023030011-3003031122111101-0302231221220001-0010212123212333-2000102111333313-2113122313001020) |
| `infra.internet_proxy` | [infra.internet_proxy](resources--registration--reference--group-001.md#canonical-0002201220031220-0311321212103020-0132032103302121-2100201333303033-0321300213220222-1020130311313233-1330200233233211-0122103033133202) |
| `infra.internet_proxy.http_proxy` | [infra.internet_proxy.http_proxy](resources--registration--reference--group-001.md#canonical-3323021301131203-2011331110200021-3223111112132032-2311002232210301-3032202103002232-3321221200110310-1011212221330103-3032310313123300) |
| `infra.internet_proxy.https_proxy` | [infra.internet_proxy.https_proxy](resources--registration--reference--group-001.md#canonical-2201213101223201-1330030120133123-3313300102031320-1213000030311011-1031220300311200-2011010313222232-1003103031100200-3021223203332231) |
| `infra.internet_proxy.no_proxy` | [infra.internet_proxy.no_proxy](resources--registration--reference--group-001.md#canonical-1223201133213223-2131023213322122-2133120300133023-3210310102132303-2023310330330210-1310202132021333-2232212112122303-3110023202201033) |
| `infra.internet_proxy.proxy_cacert_url` | [infra.internet_proxy.proxy_cacert_url](resources--registration--reference--group-001.md#canonical-0133322313233122-2120213122323200-2101210001122201-1110202101000110-0323100103122313-1002212233300231-3230212332320103-1101220233301021) |
| `infra.is_slo_static` | [infra.is_slo_static](resources--registration--reference--group-001.md#canonical-1133102112001000-1103200200202300-0330231300020230-3220232100200300-2311122230010132-1022213312000022-0022123202303100-1310303222132212) |
| `infra.machine_id` | [infra.machine_id](resources--registration--reference--group-001.md#canonical-3100333000220303-1201301031223333-0231220310312033-3213320100133021-1232111030100001-0022310033102011-1331223230233001-3300013022312223) |
| `infra.provider_ref` | [infra.provider_ref](resources--registration--reference--group-001.md#canonical-0110232231003300-3212231012321001-2123122102331130-2213213321021120-1010312112322020-2230132301103312-0333123330011102-1010012121021003) |
| `infra.sw_info` | [infra.sw_info](resources--registration--reference--group-001.md#canonical-0113320030321313-2130131001132220-1301302100322010-1022222203120121-0312000010201230-1110113230201111-0120213020102132-0302123230130110) |
| `infra.sw_info.sw_version` | [infra.sw_info.sw_version](resources--registration--reference--group-001.md#canonical-3000100311220031-0122020223302211-3231201113031221-2213023201113033-0121121213220132-1000313202312031-2000210200120101-1232110020322332) |
| `infra.timestamp` | [infra.timestamp](resources--registration--reference--group-001.md#canonical-1103111312021132-1000231000023200-2100120223212122-3032133201000001-3100331300311313-0022133333211130-1131232110223310-0133223130303200) |
| `infra.zone` | [infra.zone](resources--registration--reference--group-001.md#canonical-3002110333001101-1302302200103312-2211302311202302-0303321213332313-3332002121213103-1203020312020030-1121011101002310-0322133210011100) |
| `labels` | [labels](resources--registration--reference--group-001.md#canonical-3103123323021200-0213013032103201-0002223013012121-2221020023232323-0313012113311030-1323133112001013-2122002203030120-1120211230200033) |
| `name` | [name](resources--registration--reference--group-001.md#canonical-2320033113101223-0322300012111230-2220323322330213-2003211213210313-0030212003303313-1003132232213312-1313223022230010-2211310310123102) |
| `namespace` | [namespace](resources--registration--reference--group-001.md#canonical-2311212321311133-3131230100220120-0201110113311332-1003032333123001-2130031031211110-0031130130230111-0130031131322102-0112233312313222) |
| `passport` | [passport](resources--registration--reference--group-001.md#canonical-3300131212131333-0313002231120031-2131313131131033-0100110033123313-3112220301320333-2221002321232100-1212121011013311-1102212003030200) |
| `passport.cluster_name` | [passport.cluster_name](resources--registration--reference--group-001.md#canonical-2302100033321203-2203001330303231-1001130022213200-0310130010210003-2033020122021023-1301221312200300-3312300112031330-0200201333333002) |
| `passport.cluster_size` | [passport.cluster_size](resources--registration--reference--group-001.md#canonical-2323320022121133-2011012302120201-1321332201032002-3212101300231002-1220313020133000-3112310002330200-3313102203210103-0332311300233230) |
| `passport.cluster_type` | [passport.cluster_type](resources--registration--reference--group-001.md#canonical-3130201131232212-1013011210032112-3001322233122301-3322021302303231-0023321002020113-2023333132023313-0012231201032200-2003101100001023) |
| `passport.default_os_version` | [passport.default_os_version](resources--registration--reference--group-001.md#canonical-2130131311031330-0102232320003333-1100023330123302-3010123012002113-3231131133033100-3003021223323232-0131311112231311-3101302120000012) |
| `passport.default_sw_version` | [passport.default_sw_version](resources--registration--reference--group-001.md#canonical-0213132232332322-2130103232200210-1300321021112202-3200130020223320-0301032322001111-0231133013300322-1112333102121301-3331303320133220) |
| `passport.latitude` | [passport.latitude](resources--registration--reference--group-001.md#canonical-1022101030210021-0213322322131131-3111301332022330-0012032133332032-1101113100103203-1330220013030100-1332330112211220-2010111010111013) |
| `passport.longitude` | [passport.longitude](resources--registration--reference--group-001.md#canonical-2130222322100100-3012202313203123-0030330021333321-0231121303022010-2110230021010021-2023223211013232-0032011122231331-2231112123002311) |
| `passport.operating_system_version` | [passport.operating_system_version](resources--registration--reference--group-001.md#canonical-2133031023320301-3200102232001302-2100332232301102-2223321223222321-2102122121230123-1332030030221132-3203213033211032-3012020321302322) |
| `passport.private_network_name` | [passport.private_network_name](resources--registration--reference--group-001.md#canonical-2010230203200001-0023102122320111-1133220331103100-3130013333112102-2231221001312303-2031110120330131-2312103002103113-0310033330212010) |
| `passport.volterra_software_version` | [passport.volterra_software_version](resources--registration--reference--group-001.md#canonical-1101110132201203-2101211202231311-1212211323003313-2212323000311321-3122132123331211-2111011311123112-2033023301111020-3312312313002230) |
| `timeouts` | [timeouts](resources--registration--reference--group-001.md#canonical-0332003222321313-1030212102212103-1010221332320312-1003230100020322-2102221312100020-1231222120103011-3022030223011110-1021223023300230) |
| `timeouts.create` | [timeouts.create](resources--registration--reference--group-001.md#canonical-1321232120333203-2021202311012131-2003113313002130-2120130023312230-0112103122003322-0113032333302321-3211201332321031-0022023230020203) |
| `timeouts.delete` | [timeouts.delete](resources--registration--reference--group-001.md#canonical-3133010302200303-2130030310032013-3220233200212313-3301222221003332-1201101231222201-3231013203230121-2212302311321021-1232110122133333) |
| `timeouts.read` | [timeouts.read](resources--registration--reference--group-001.md#canonical-2033333132000101-1310312333321023-3010012200032110-3123123330233111-1122200201112003-0021001030233213-0233330110223000-3231320302322101) |
| `timeouts.update` | [timeouts.update](resources--registration--reference--group-001.md#canonical-0030003213133300-3301000131230110-0133010212302331-0013102303221233-0130133100012111-1112132203311223-1201133311302103-1220031003113210) |
| `token` | [token](resources--registration--reference--group-001.md#canonical-2233233122220320-2231011130200311-3321103022031013-2010332002221231-3222321311002110-1031012310131112-0331111300323201-1132101310121201) |

<a id="canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- infra

<a id="canonical-0333202120232121-0302332001312301-2302131021201221-2331300122122012-1010112021230303-0011330200002231-3021332113332312-1312121130112213"></a>

Type: `"object"`. single nested block, Optional.

InfraMetadata stores information about instance infrastructure.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("hostname",
    "interfaces")}
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
infra {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131113032320221-2202010120231011-1023013122121320-3101103212000032-3101102122233330-3112101112121321-2031201132013032-2331003300102002"></a>

### Direct properties for `infra`

<a id="canonical-1103310230111111-3000023310022230-0222303122130301-0212031323030221-1313201002310100-0203133112023211-3203332130331201-0312110233322111"></a>

#### `infra.availability_zone` property

Type: `"string"`. Optional.

An Availability Zone is a high-availability offering that protects your applications and data from
datacenter failures.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [bond_config](resources--registration--reference--group-001.md#canonical-0121201130321321-3231311301122033-2230210013220010-2002333120112302-0120303122023013-0132320302232111-1221102101312231-0201133202131022): complete subsection reference.

<a id="canonical-2132121300101012-2331201231011222-1013331311313020-0233000220100330-3202223023030302-0033022211101010-1031130030310322-1132103202103200"></a>

<a id="canonical-0200231303010322-0330210033301032-1123031332110023-3300021001231033-2212302033312003-2023010000102003-0011020010101322-0231030011222312"></a>

#### `infra.certified_hw` property

Type: `"string"`. Optional.

Certified HW name used to map with F5XC certified\_hardware definition.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2213012130320323-0013021201331100-2000012213221230-1123223001001031-0230100211210011-2211103332320232-2122221312200313-3203011221310011"></a>

<a id="canonical-2121112213212010-2202001211131310-0103012322012003-2120103232133132-3302323013100021-2030301303332230-1123101121101213-3302203233031200"></a>

#### `infra.domain` property

Type: `"string"`. Optional.

Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5
Distributed Cloud.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1233222201331213-3131013303213003-0321011221033121-2132111312113330-1010120323013110-0100130302212331-3211211333213001-2133123123002020"></a>

<a id="canonical-2120213002100211-3100201333013230-0003022231010012-0202020212222130-2303111213030112-0132123131311112-3023100012332200-3220221200013321"></a>

#### `infra.hostname` property

Type: `"string"`. Optional.

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [hugepages](resources--registration--reference--group-001.md#canonical-3201100030132010-0301200013320120-2133223300001321-2020020110003202-2111033011003210-3210020322302302-2121023211021101-0213121211230210): complete subsection reference.

- [hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221): complete subsection reference.

<a id="canonical-0233233011200103-3201330302211210-0130222200331100-2213301332233002-1130321212123332-2113112222123123-0100033312010223-3133303333101211"></a>

<a id="canonical-0000101323231112-1020300333210002-0113213003312103-1301233033033330-0132010023303223-0310302322000312-2002010012230001-0232031220232233"></a>

#### `infra.instance_id` property

Type: `"string"`. Optional.

Instance ID (assigned by infrastructure provider).

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [interfaces](resources--registration--reference--group-001.md#canonical-1131111001122100-2210113013321231-2011202021020021-1321111013331333-2103232202022321-0113302202013102-0200012200312100-0002212230222000): complete subsection reference.

- [internet_proxy](resources--registration--reference--group-001.md#canonical-0333232120213313-2131312201312220-0332010123213011-1100232332221223-3130030323031202-2201003302332033-3110013202300010-1101212310212223): complete subsection reference.

<a id="canonical-1133102112001000-1103200200202300-0330231300020230-3220232100200300-2311122230010132-1022213312000022-0022123202303100-1310303222132212"></a>

<a id="canonical-3323110223022232-0301002113022222-2310233012110023-3222103121232130-3123131033212020-2001310331323322-2233313310212330-1033333210103233"></a>

#### `infra.is_slo_static` property

Type: `"bool"`. Optional.

Is SLO Static. Indicates whether the SLO is static.

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

<a id="canonical-3100333000220303-1201301031223333-0231220310312033-3213320100133021-1232111030100001-0022310033102011-1331223230233001-3300013022312223"></a>

<a id="canonical-3331223121303220-2022233023103110-1211323320303131-3310323331030233-0221102103122031-3311132023211001-0231030303113333-1031113030300301"></a>

#### `infra.machine_id` property

Type: `"string"`. Optional.

Machine ID - generated by operating system.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0110232231003300-3212231012321001-2123122102331130-2213213321021120-1010312112322020-2230132301103312-0333123330011102-1010012121021003"></a>

<a id="canonical-1002030031223012-1112110001301313-2221111311312222-3222213213021201-2332331020121100-3233033232323331-2221123000013212-0030033301103121"></a>

#### `infra.provider_ref` property

Type: `"string"`. Optional.

\[Enum:
UNKNOWN|AWS|GOOGLE|AZURE|VMWARE|KVM|OTHER|VOLTERRA|IBMCLOUD|UNKNOWN\_K8S|AWS\_K8S|GCP\_K8S|AZURE\_K8S|VMWARE\_K8S|KVM\_K8S|OTHER\_K8S|VOLTERRA\_K8S|IBMCLOUD\_K8S|F5OS|RSERIES|OCI|NUTANIX|OPENSTACK|EQUINIX|OPENSHIFT\_VIRTUALIZATION|KUBERNETES\]
Infrastructure provider enum for registration. It describes where is instance running. Provider was
not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other
provider, which was not identified by system. Possible values are \`UNKNOWN\`, \`AWS\`, \`GOOGLE\`,
\`AZURE\`, \`VMWARE\`, \`KVM\`, \`OTHER\`, \`VOLTERRA\`, \`IBMCLOUD\`, \`UNKNOWN\_K8S\`,
\`AWS\_K8S\`, \`GCP\_K8S\`, \`AZURE\_K8S\`, \`VMWARE\_K8S\`, \`KVM\_K8S\`, \`OTHER\_K8S\`,
\`VOLTERRA\_K8S\`, \`IBMCLOUD\_K8S\`, \`F5OS\`, \`RSERIES\`, \`OCI\`, \`NUTANIX\`, \`OPENSTACK\`,
\`EQUINIX\`, \`OPENSHIFT\_VIRTUALIZATION\`, \`KUBERNETES\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AWS","AWS_K8S","AZURE","AZURE_K8S","EQUINIX","F5OS","GCP_K8S","GOOGLE","IBMCLOUD","IBMCLOUD_K8S","KUBERNETES","KVM","KVM_K8S","NUTANIX","OCI","OPENSHIFT_VIRTUALIZATION","OPENSTACK","OTHER","OTHER_K8S","RSERIES","UNKNOWN","UNKNOWN_K8S","VMWARE","VMWARE_K8S","VOLTERRA","VOLTERRA_K8S"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN",
    "AWS",
    "GOOGLE",
    "AZURE",
    "VMWARE",
    "KVM",
    "OTHER",
    "VOLTERRA",
    "IBMCLOUD",
    "UNKNOWN_K8S",
    "AWS_K8S",
    "GCP_K8S",
    "AZURE_K8S",
    "VMWARE_K8S",
    "KVM_K8S",
    "OTHER_K8S",
    "VOLTERRA_K8S",
    "IBMCLOUD_K8S",
    "F5OS",
    "RSERIES",
    "OCI",
    "NUTANIX",
    "OPENSTACK",
    "EQUINIX",
    "OPENSHIFT_VIRTUALIZATION",
    "KUBERNETES"),
}
```

- [sw_info](resources--registration--reference--group-001.md#canonical-0312321212221203-3210322201111303-0121010122010333-2311320312012111-2333222213001230-0231301313330001-3231102122010310-3030021133012220): complete subsection reference.

<a id="canonical-1103111312021132-1000231000023200-2100120223212122-3032133201000001-3100331300311313-0022133333211130-1131232110223310-0133223130303200"></a>

<a id="canonical-1100100101111331-2232011001111111-3103202201232031-2220101111121332-3303331111123332-1331320313222100-3122331312103011-2211133032103022"></a>

#### `infra.timestamp` property

Type: `"string"`. Optional.

It's used to verify machine have acceptable time difference from server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "temporal",
    "constraintType": "string",
    "deterministic": true,
    "format": "date-time",
    "formatDescription": "ISO 8601 date-time (e.g., 2026-01-19T12:00:00Z)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 20,
    "pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d{3})?Z?$",
    "validation": {
      "standard": "ISO 8601"
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

<a id="canonical-3002110333001101-1302302200103312-2211302311202302-0303321213332313-3332002121213103-1203020312020030-1121011101002310-0322133210011100"></a>

<a id="canonical-1121331133302222-3003122021100011-0300133032201112-3132323131111330-1230300102131323-0000122203331200-3211210120123232-1211210012023210"></a>

#### `infra.zone` property

Type: `"string"`. Optional.

Instance zone (or region), depends on provider.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0121201130321321-3231311301122033-2230210013220010-2002333120112302-0120303122023013-0132320302232111-1221102101312231-0201133202131022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.bond_config` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- infra.bond_config

<a id="canonical-0113010113111223-0201010222323221-0202013223212030-1022221211101233-0213331301332213-1202232130321123-2012333303320301-3220202212223030"></a>

Type: `"object"`. single nested block, Optional.

Bond device configuration for VPM registration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces",
    "name")}
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
bond_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013110013003310-0032322221200032-3221112321223231-2311111323110323-1201101132213323-2133231130122123-2001313212311331-2332131233322210"></a>

### Direct properties for `infra.bond_config`

<a id="canonical-0033131331300101-0103233013010210-3231303200131232-1003310230203202-3313221201013222-1311213113023323-3002110010200121-3112030323300303"></a>

#### `infra.bond_config.interfaces` property

Type: `["list", "string"]`. Optional.

Member Interfaces. Configuration parameter for interfaces

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2012122221000220-2113222303131002-3003132203330331-3301321012103203-0113032131313022-0011121022122112-0001123211001011-0203222030003101"></a>

<a id="canonical-1330312332221032-2231302323211221-0231012323223130-2133200230130012-3022203302210300-1322222210033032-1013010202332131-1130120121331201"></a>

#### `infra.bond_config.mode` property

Type: `"string"`. Optional.

\[Enum: BOND\_MODE\_UNSPECIFIED|ACTIVE\_BACKUP|LACP\_802\_3AD\] Bonding mode for bond device
configuration Bond mode is not specified Active-backup bond mode (one interface active, others as
backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are
\`BOND\_MODE\_UNSPECIFIED\`, \`ACTIVE\_BACKUP\`, \`LACP\_802\_3AD\`. Defaults to
\`BOND\_MODE\_UNSPECIFIED\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ACTIVE_BACKUP","BOND_MODE_UNSPECIFIED","LACP_802_3AD"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BOND_MODE_UNSPECIFIED",
    "ACTIVE_BACKUP",
    "LACP_802_3AD"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BOND_MODE_UNSPECIFIED",
  "enum": [
    "BOND_MODE_UNSPECIFIED",
    "ACTIVE_BACKUP",
    "LACP_802_3AD"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3332101322323330-1012123303003013-2130001102032112-2200031332212311-2200133110201133-0132310010202010-0002023101020012-1321003023301133"></a>

<a id="canonical-1333221320102230-2033322023323301-2330201012033322-3123130022123202-2222130322213312-0123223201000330-0103220001022033-2202000233102123"></a>

#### `infra.bond_config.name` property

Type: `"string"`. Optional.

Bond Name. Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3201100030132010-0301200013320120-2133223300001321-2020020110003202-2111033011003210-3210020322302302-2121023211021101-0213121211230210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hugepages` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- infra.hugepages

<a id="canonical-2032301231212002-0130113300121021-3201303012312333-2213030310203110-3332300333011022-0003120111200223-3320322010013320-3310331320003032"></a>

Type: `"object"`. list nested block, Optional.

Hugepage settings for CE on K8s SMV2 site.

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
hugepages {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322231113211320-1103220332302330-3230221200220021-1301120232203201-3210123122232202-3121012113210012-0012030230030000-3200011312023200"></a>

### Direct properties for `infra.hugepages`

<a id="canonical-2232222223303222-1000111323010030-0303013310230132-2330202001211222-0130002133220300-0011211123212321-3131201332030121-0111321233030121"></a>

#### `infra.hugepages.free` property

Type: `"number"`. Optional.

Free Hugepages. Total number of free hugepages present.

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

<a id="canonical-1033221321230222-1211132231101220-0110112130320013-2211012120110303-1310113121111322-3231220220333031-1230011113332120-0001330200303300"></a>

<a id="canonical-1330321110300100-1211121123100131-1100211232133103-2203322210022003-0213020300212023-2111010310113322-0210011222212100-3101003211220232"></a>

#### `infra.hugepages.page_size` property

Type: `"number"`. Optional.

Hugepage Size. Size of each hugepage.

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

<a id="canonical-1012322003132200-2301121003113022-2232111331113002-2111233232111221-0321302222231233-2211312032032210-2021332303001311-1031121123332131"></a>

<a id="canonical-2131223323022011-0210202222213320-0021102323213230-2122313102231032-2330311203211300-1303333202333122-2133321101222303-2302030221212332"></a>

#### `infra.hugepages.total` property

Type: `"number"`. Optional.

Total Hugepages. Total number of hugepages present.

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

<a id="canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- infra.hw_info

<a id="canonical-2112213100301233-3101013330101230-1022101030230130-0131131332111320-0300201233212102-3311030001002311-1002122022301120-1131221130012220"></a>

Type: `"object"`. single nested block, Optional.

OsInfo holds information about host OS and HW.

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
hw_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111122122202300-1101230001333312-1300310303302220-2320002102303111-0300211131011121-1023233030122121-1122202133331133-2303202221202301"></a>

### Direct properties for `infra.hw_info`

- [bios](resources--registration--reference--group-001.md#canonical-3230213321222323-1333012313102022-1113123132022220-1010222000301033-2333020321320130-0121011220223330-3213220332222320-0203330323103323): complete subsection reference.

- [board](resources--registration--reference--group-001.md#canonical-0301033122022133-3301003013021102-1202011231132113-2202333220230122-0212203132212102-0122012210202233-1012231322322220-1020001303320301): complete subsection reference.

- [chassis](resources--registration--reference--group-001.md#canonical-0323323131133211-0011220111233113-3322310102330001-3332123213202301-1222121212331031-1010111023003000-1112303301202203-3232031100203230): complete subsection reference.

- [CPU](resources--registration--reference--group-001.md#canonical-1103320311303112-1300213033101000-3022123203213331-1233010332102022-1331201213120213-3222301122300203-3200322112102333-3001102020031122): complete subsection reference.

- [GPU](resources--registration--reference--group-001.md#canonical-2011333333132333-3300333110123203-0003012331203003-2213312012303211-2201121301032100-0132123323332222-1133303302222001-3030021001011310): complete subsection reference.

- [kernel](resources--registration--reference--group-001.md#canonical-3103323320112222-1021320331211233-2330232112001003-3022233110313322-3100023302310032-2112032330023102-0221131110310001-0011212333321203): complete subsection reference.

- [memory](resources--registration--reference--group-001.md#canonical-3113330200330332-3001000333122232-3332202221130100-2030231103312000-1003212113122233-2201110030130303-0200201103002202-0021320013330100): complete subsection reference.

- [network](resources--registration--reference--group-001.md#canonical-3231003001113113-1202202113333121-2322003301233222-3222112331220003-0122022011310202-0302130032123000-1113110310112033-2122332123001001): complete subsection reference.

<a id="canonical-2220103100103110-0121220231332201-2212302322300201-1102311200300202-0233331102322301-1331212101100202-2211012232033031-0300202032313112"></a>

<a id="canonical-2223100113302030-0000011020331330-3103333211210301-0233332232103232-2100310123230020-2101110112133101-3210320002313010-1301302223101331"></a>

#### `infra.hw_info.numa_nodes` property

Type: `"number"`. Optional.

Non-uniform memory access (NUMA) nodes count.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.int32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0"
  }
}
```

- [os](resources--registration--reference--group-001.md#canonical-1231001033023312-2011113220003311-0222030233001030-3203302232031331-1331210103302203-2123123330013022-3000132230133322-3213211033011311): complete subsection reference.

- [product](resources--registration--reference--group-001.md#canonical-3111131021132120-1201132311302220-0322330011200210-0332003110023033-3210113133133231-1112310100321012-1201113230332133-0200001023220311): complete subsection reference.

- [storage](resources--registration--reference--group-001.md#canonical-3313133121200132-3301001210013020-0000211123132303-3331332012323121-0112110122022030-3033200103123122-0100332030330030-2332130132231120): complete subsection reference.

- [usb](resources--registration--reference--group-001.md#canonical-2133131333302012-2311223101000032-3011003103231233-1102312333003101-0202323211123222-0220232101330330-2012330211102002-3032003002021300): complete subsection reference.

<a id="canonical-3230213321222323-1333012313102022-1113123132022220-1010222000301033-2333020321320130-0121011220223330-3213220332222320-0203330323103323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.bios` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.bios

<a id="canonical-2212201011120332-1302010023213312-2330200031022212-2132113201301000-0311031012111232-0130011020303333-3101113321213200-1230223010233232"></a>

Type: `"object"`. single nested block, Optional.

Bios Data. BIOS information.

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
bios {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333111130200133-2000230211230323-2213221303231122-1232221122031212-2121333110330110-1210112200101331-3031331301313201-0033133230100202"></a>

### Direct properties for `infra.hw_info.bios`

<a id="canonical-0211031321232321-3230131100022313-3000300001030010-2210132123113303-0321231030312330-3321313311220013-1323230100313230-2103121100100110"></a>

#### `infra.hw_info.bios.date` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/bios\_date.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(10, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "temporal",
    "constraintType": "string",
    "deterministic": true,
    "format": "date",
    "formatDescription": "ISO 8601 date (e.g., 2026-01-19)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 10,
    "pattern": "^\\d{4}-\\d{2}-\\d{2}$",
    "validation": {
      "standard": "ISO 8601"
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

<a id="canonical-1002210201323222-1013223220222213-0103032333201030-3332300200322202-1030200303010112-1322002113311331-1203113003233232-3221310012112031"></a>

<a id="canonical-2300312201003333-2030320110022313-1331302031202113-0201301133003013-3302132111302230-2313003323301110-0300310112310012-3001131103330133"></a>

#### `infra.hw_info.bios.vendor` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/bios\_vendor.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2010100331021301-0233213111333303-2223323033020003-3132023000131022-2311012301321120-2100230111310303-2300200002133311-1010313301023113"></a>

<a id="canonical-3120131212300301-1322211131231123-2232112211112001-2233303332102021-1012331330330030-3202212333213210-1010002011200301-1011321011332232"></a>

#### `infra.hw_info.bios.version` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/bios\_version.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0301033122022133-3301003013021102-1202011231132113-2202333220230122-0212203132212102-0122012210202233-1012231322322220-1020001303320301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.board` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.board

<a id="canonical-2232033112110232-1032310110200220-0230203020333312-1221113001301112-0000300222022131-3331332012332033-2312122103000213-1131012220103333"></a>

Type: `"object"`. single nested block, Optional.

Board Details. Board information.

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
board {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021112103330111-2121132202032020-2233023100300320-0102002133032213-2212023312123113-0022031303111031-2230003330010313-0202333112323130"></a>

### Direct properties for `infra.hw_info.board`

<a id="canonical-0021103310103033-3202313020020120-2013312133121332-3203013023022321-2122022012211311-3320310322030322-1331100211100230-2330101102300333"></a>

#### `infra.hw_info.board.asset_tag` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_asset\_tag.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1022000301222100-1033211221102303-2003132013220211-1332210102302230-1201031012031323-0033130233110200-3101032201320212-1203303131021233"></a>

<a id="canonical-2331230101013020-2030112031213110-2122221213212211-3222122231013020-1001230211103122-0100030022223332-1130212222133023-1220333111103123"></a>

#### `infra.hw_info.board.name` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1133002100233031-0033001331310033-0033320201232202-3122031303222121-3132221221332031-0331311332330102-2210232001010213-2231200120012121"></a>

<a id="canonical-2123202232223022-2103220003002222-2232003101301231-1231130310202033-1010323301010030-0000301102313000-2202233332122121-0210221211122133"></a>

#### `infra.hw_info.board.serial` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_serial.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1132203131303230-1203031113000121-1332011022331122-2302000321000231-1320111301200031-0222130322123120-1012323212113222-3103013033213021"></a>

<a id="canonical-0122120202011321-1031231321202333-2023311213221103-1333103321103323-0220103230230221-0002111222201110-2323312331320031-1020311232322310"></a>

#### `infra.hw_info.board.vendor` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_vendor.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3333023100133001-0322010131203313-0000002212330110-3133323301312131-3023233011231102-0120011230023000-3222131110210033-0013010033023033"></a>

<a id="canonical-3013133311002313-0021230221201301-1330101323233303-1103111303011333-3130013221033121-0321321113233323-1320333201033130-0321101210302333"></a>

#### `infra.hw_info.board.version` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/board\_version.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0323323131133211-0011220111233113-3322310102330001-3332123213202301-1222121212331031-1010111023003000-1112303301202203-3232031100203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.chassis` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.chassis

<a id="canonical-1312011201130301-1000001021130121-2110221332032123-0221321032102222-1023033333303211-0021120022122223-3210033322230332-0023203222121002"></a>

Type: `"object"`. single nested block, Optional.

Chassis Details. Chassis information.

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
chassis {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002123233003220-2012313210322003-2310010233021312-3003313120302231-2032230232200200-2133122020130313-2222133123010002-1110021212011122"></a>

### Direct properties for `infra.hw_info.chassis`

<a id="canonical-1303322302012302-0232223330110210-2302333233302130-2202113322301120-3202301300323333-2333021121322013-3130202310021320-1022002132000102"></a>

#### `infra.hw_info.chassis.asset_tag` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_asset\_tag.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3232000011322013-0313313131201223-1232303201031322-2131330222020310-2011033112132121-0123000100112311-2212222132000000-2231102302033030"></a>

<a id="canonical-0121212122002022-3130100222010013-3121122120231032-2112212120021323-1313120211313123-0031213301333032-1211221331332003-1131000220113332"></a>

#### `infra.hw_info.chassis.serial` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_serial.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2010031112023330-1303310110322133-0330133320011121-1023203023301021-0122330022302112-2310012300330222-2023001202130302-3200232332113112"></a>

<a id="canonical-1100202131203310-1032013203032100-1232233322201011-0221013022312101-3300110120320230-2021323230300133-2013213310100212-2303312113100112"></a>

#### `infra.hw_info.chassis.type` property

Type: `"number"`. Optional.

Information from /sys/class/dmi/ID/chassis\_type.

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

<a id="canonical-2012032002122121-1230013202013321-3031331332311022-3002213023221232-0020321203303313-0011333233300033-1113312033112110-3010221021302201"></a>

<a id="canonical-1300000112313300-1203021322133000-2201212022112321-1222211003203302-0030233302230003-1021131232230322-3322012100003312-3220312032010302"></a>

#### `infra.hw_info.chassis.vendor` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_vendor.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3320320003100221-1302002112022322-3132130003302313-3132203000213111-2113323031232302-3333120023331120-0210100003213301-3121212132322322"></a>

<a id="canonical-3121110202231123-1121300201203020-3212112233332122-1001312220013121-3131122202321332-2033023202302231-2303031333210012-2210033222323123"></a>

#### `infra.hw_info.chassis.version` property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/chassis\_version.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1103320311303112-1300213033101000-3022123203213331-1233010332102022-1331201213120213-3222301122300203-3200322112102333-3001102020031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.cpu` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.CPU

<a id="canonical-0230323221112310-0333213013033222-3133000300123023-3111312231103123-3103032322121323-0112203111031130-1100021000023313-2011000201213132"></a>

Type: `"object"`. single nested block, Optional.

CPU Information. CPU information.

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
cpu {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322021130132132-3010112132002003-1111200333221101-1312320212113012-1013310332233032-2031112021311121-3012331102221130-2223022331031110"></a>

### Direct properties for `infra.hw_info.cpu`

<a id="canonical-3322201112220101-1212011003131022-2011100013110113-3320223000310021-3303110321123123-0132213220030331-1032333320212002-1031322120233302"></a>

#### `infra.hw_info.cpu.cache` property

Type: `"number"`. Optional.

Cache. CPU cache size in KB.

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

<a id="canonical-2031332000300222-1013230132213203-0321320211211033-1202023122222101-1012001301020312-2010113111301303-3113030020330333-2213313333021033"></a>

<a id="canonical-0232121030230232-3030233231220210-2021102200212131-1022120321100222-2200111211100330-0111321012002220-1310310010333322-2011222203301331"></a>

#### `infra.hw_info.cpu.cores` property

Type: `"number"`. Optional.

Cores. Number of physical CPU cores.

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

<a id="canonical-3311201221223122-0201013232113221-0221230212103111-0001302333132210-2111030122201031-3231220311203111-0221213303011002-3311310323233032"></a>

<a id="canonical-0331233233011232-0112031003023301-3030130231210030-3233320102132210-3232110031002121-3111122333031330-1003321310020012-3031233100210213"></a>

#### `infra.hw_info.cpu.cpus` property

Type: `"number"`. Optional.

CPUs. Number of physical CPUs.

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

<a id="canonical-0212330332011212-3001100110202322-0221330132310203-2003301113232211-3322001321311333-3030331230101102-1002332302020232-2103323000101103"></a>

<a id="canonical-3022202231203030-0212033321013122-0301211312302023-3213230233223113-0123023213123021-3123221103031011-0231311211320031-0310122032201031"></a>

#### `infra.hw_info.cpu.model` property

Type: `"string"`. Optional.

Model. CPU model

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0230012313033133-1003033333230213-1223030210002003-2223033132302320-1110010120313320-1220023323103233-3301233231220311-1101103013111320"></a>

<a id="canonical-3020013312321131-0332011012220003-2010032220213320-3202330012133203-1113233213122033-1031113123230012-3311120232203233-2230101222330211"></a>

#### `infra.hw_info.cpu.speed` property

Type: `"number"`. Optional.

Speed. CPU clock rate in MHz.

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

<a id="canonical-0013212013001310-0330112003300212-0102030320010200-1330121300103120-0323132022331213-1021022033221320-0321323000221113-2131030102123121"></a>

<a id="canonical-2300211122022231-0313012330002231-2301133221021233-1121010300300300-1202331301233130-1213210133022202-1113333230033332-0300303011123300"></a>

#### `infra.hw_info.cpu.threads` property

Type: `"number"`. Optional.

Threads. Number of logical (HT) CPU cores.

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

<a id="canonical-1133203231332103-2002202333210203-1313332333021232-0030022102123331-3122101032001230-2300121232013132-3230322121220122-3021212003002112"></a>

<a id="canonical-3302201220301111-1021002012202013-2103103103303303-1030211201322112-2121323300031121-0220133331002130-2231103312201333-0102203132131131"></a>

#### `infra.hw_info.cpu.vendor` property

Type: `"string"`. Optional.

Vendor. CPU vendor.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2011333333132333-3300333110123203-0003012331203003-2213312012303211-2201121301032100-0132123323332222-1133303302222001-3030021001011310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.gpu` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.GPU

<a id="canonical-3300011012312122-1033321001030202-3201231000233211-3210213321231020-3100130212233232-3330000030321333-1120103000030312-1320111303211031"></a>

Type: `"object"`. single nested block, Optional.

GPU. GPU information on server.

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
gpu {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201110222001301-2113123222330233-3032302312001131-0222020232133122-2311000312302101-2320303201133231-1001300033103303-1301020311310332"></a>

### Direct properties for `infra.hw_info.gpu`

<a id="canonical-3202032203111220-2233233011111000-0100120220311113-2202020033000030-1221122322022133-3231103211003321-1131232131110331-1101303110300120"></a>

#### `infra.hw_info.gpu.cuda_version` property

Type: `"string"`. Optional.

Cuda Version. GPU Cuda Version.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3010211223032133-1111002210301230-3222311102323033-3110222301232111-2330110320101210-2100103102301012-2211221011311121-0132330321331221"></a>

<a id="canonical-1030122223133202-3111211010002222-1321213203133120-0020023022102233-1013102131033232-1031233122013213-0013231332001300-0120011003012100"></a>

#### `infra.hw_info.gpu.driver_version` property

Type: `"string"`. Optional.

Driver Version. GPU Driver Version.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [gpu_device](resources--registration--reference--group-001.md#canonical-1331133021031001-2101301001120200-1031301200113130-3013010302203132-2302103313011011-1030203101033121-3202220230102331-2031132132102023): complete subsection reference.

<a id="canonical-1331133021031001-2101301001120200-1031301200113130-3013010302203132-2302103313011011-1030203101033121-3202220230102331-2031132132102023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.gpu.gpu_device` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- [infra.hw_info.gpu](resources--registration--reference--group-001.md#canonical-2011333333132333-3300333110123203-0003012331203003-2213312012303211-2201121301032100-0132123323332222-1133303302222001-3030021001011310)
- infra.hw_info.GPU.gpu_device

<a id="canonical-0220322132002100-1333020113232303-3103130310221322-0112003232111110-3003321201110031-2000100013111102-1312232232221133-0100033233111011"></a>

Type: `"object"`. list nested block, Optional.

GPU devices. List of GPU devices in server.

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
gpu_device {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032221111113332-2131323210223031-1031100330031001-1332322223102030-2111321320310130-2220230333020001-1131213130212032-0032022001012313"></a>

### Direct properties for `infra.hw_info.gpu.gpu_device`

<a id="canonical-2212310323011023-0223013223320111-0330110321200201-2011131023201033-2220101332230310-1030121221212120-2032332033010302-0212330222301211"></a>

#### `infra.hw_info.gpu.gpu_device.id` property

Type: `"string"`. Optional.

GPU ID. GPU ID

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3010031002313103-3111103100232023-3002130311303301-1123020202330033-3322022332113110-1003001130233230-1222300201020211-1223012003323133"></a>

<a id="canonical-1031313100302202-2223022113231131-2000232012300012-3312101001203211-3020200013032022-3022213210333201-0101121332003310-1310203331223331"></a>

#### `infra.hw_info.gpu.gpu_device.processes` property

Type: `"string"`. Optional.

Processes. GPU Processes.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0133010223013201-3123123023111231-2302123202132103-3032302100020200-1001331101030011-0003013102210311-0303112023230300-2010330210013210"></a>

<a id="canonical-0102120031020011-0113100311011122-1032133101002201-3031312221031100-2323233211122303-2233103200230022-2112232222210332-3130303022323200"></a>

#### `infra.hw_info.gpu.gpu_device.product_name` property

Type: `"string"`. Optional.

Product Name. GPU Product Name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3103323320112222-1021320331211233-2330232112001003-3022233110313322-3100023302310032-2112032330023102-0221131110310001-0011212333321203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.kernel` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.kernel

<a id="canonical-1230220122001331-1220332023233010-3332302130211101-3231321330301233-2322211101123033-0020012233301120-0013223030231030-0200020122012031"></a>

Type: `"object"`. single nested block, Optional.

Kernel. Kernel information.

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
kernel {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011003233022022-0133101003211002-3130220223233000-2212010111223211-1102233122200131-1230132201110233-2331232233012203-3232000210220310"></a>

### Direct properties for `infra.hw_info.kernel`

<a id="canonical-0221112033103123-2120300322322232-2322122122330310-2102213130303310-2230030032103012-3123321322122111-0132101021231103-2330323111332320"></a>

#### `infra.hw_info.kernel.architecture` property

Type: `"string"`. Optional.

Architecture. Kernel architecture.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2123003202302011-2220001331100100-2300311213311000-1121323030213130-1211010301133332-2031322233211112-0121211012013131-2332022200310113"></a>

<a id="canonical-0300311100033131-0010301222010330-3022223101103223-1233200303030223-3030202203102301-0133201022213210-0031011111130022-2013022011020302"></a>

#### `infra.hw_info.kernel.release` property

Type: `"string"`. Optional.

Release. Kernel release.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1003321300020132-0302102021221101-3133323333300001-2211121221131300-2020011023122221-0031210301203320-0113013320320100-2201122011003233"></a>

<a id="canonical-1331332121132213-3223133220121011-3030332212022033-3332122100212330-1101322300231102-2222231221023003-2100301330231220-3100213321311002"></a>

#### `infra.hw_info.kernel.version` property

Type: `"string"`. Optional.

Version. Kernel version.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3113330200330332-3001000333122232-3332202221130100-2030231103312000-1003212113122233-2201110030130303-0200201103002202-0021320013330100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.memory` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.memory

<a id="canonical-2000310221113132-1032101333023220-2030221022013001-2332123121030023-0020013131020200-1220123210332223-2131221020331021-0020031311030330"></a>

Type: `"object"`. single nested block, Optional.

Memory Information. Memory information.

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
memory {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023022132321220-0022121133120321-2002223000123102-2312203013331100-1011322333012223-3223223221302333-2120121120300100-2223221103133110"></a>

### Direct properties for `infra.hw_info.memory`

<a id="canonical-0023120331133221-2121321000023032-2031112331012312-3023002231332001-1020013101223223-1102303002221102-0020112320012121-0223313303220030"></a>

#### `infra.hw_info.memory.size_mb` property

Type: `"number"`. Optional.

RAM. RAM size in MB.

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

<a id="canonical-3112303002301331-1121021202120221-2203000222010203-1031001021301123-0013022101220131-2310132100302201-0000130302003330-0003220331220003"></a>

<a id="canonical-2302230130123110-3202002313112322-0001310212102211-1231312033100333-1131103033123200-3300323332000323-3002230333101003-3311202221122110"></a>

#### `infra.hw_info.memory.speed` property

Type: `"number"`. Optional.

Speed. RAM data rate in MT/s.

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

<a id="canonical-2130012101010320-3212222332302020-2112010103021120-3303311231013301-3011200322011200-0301012223020021-2213101023103323-3200233322110021"></a>

<a id="canonical-3233013113103211-1320132213020022-1232211120132001-1102100031333311-2312130120020231-0121003023321120-1302000222131011-3102111321303101"></a>

#### `infra.hw_info.memory.type` property

Type: `"string"`. Optional.

Type. Type of memory, eg. DDR4.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3231003001113113-1202202113333121-2322003301233222-3222112331220003-0122022011310202-0302130032123000-1113110310112033-2122332123001001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.network` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.network

<a id="canonical-0211210001022121-1230303301321310-2102233223301110-2012110130030103-1232113012213003-0321210102232210-0313120301230301-0032231222302323"></a>

Type: `"object"`. list nested block, Optional.

Network. List of network devices in server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "array",
    "maxItems": 32,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
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
network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003113133221232-1311220232013020-3302011122120323-2023132312010033-3312322311011212-0110013120011132-3202132033003010-1232111100211333"></a>

### Direct properties for `infra.hw_info.network`

<a id="canonical-3203331230202030-2332023321302233-3022110012332212-2321003002332200-1030202002130213-1333210001312130-3312202133101021-0030331333312221"></a>

#### `infra.hw_info.network.driver` property

Type: `"string"`. Optional.

Driver. Driver of device, eg. E1000e.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0010323211023302-3330233211203010-3232221200231022-2003223021003332-3003012321223212-0121223103230002-2330211002131033-3113302112331030"></a>

<a id="canonical-2311031321301223-3012331113032013-2003203233322300-2101233303233010-2002320220010311-2330130113313023-1323310012120131-3201211201132303"></a>

#### `infra.hw_info.network.ip_address` property

Type: `["list", "string"]`. Optional.

IP Address. IP address on interface.

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

<a id="canonical-3223032320311002-0321303000032131-0032220120330302-1013032102132000-3031320300303112-2321320001022201-0233013012221202-1330222021101213"></a>

<a id="canonical-1202010333213012-0321012313321003-1020331233333323-1220301221303013-3111302211122322-1103101102031230-1301020100322112-2313211231233223"></a>

#### `infra.hw_info.network.link_quality` property

Type: `"string"`. Optional.

\[Enum: QUALITY\_UNKNOWN|QUALITY\_GOOD|QUALITY\_POOR|QUALITY\_DISABLED\] Link quality determined by
VER using different probes Unknown quality Link quality is good Link quality is poor Quality
disabled. Possible values are \`QUALITY\_UNKNOWN\`, \`QUALITY\_GOOD\`, \`QUALITY\_POOR\`,
\`QUALITY\_DISABLED\`. Defaults to \`QUALITY\_UNKNOWN\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["QUALITY_DISABLED","QUALITY_GOOD","QUALITY_POOR","QUALITY_UNKNOWN"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("QUALITY_UNKNOWN",
    "QUALITY_GOOD",
    "QUALITY_POOR",
    "QUALITY_DISABLED"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "QUALITY_UNKNOWN",
  "enum": [
    "QUALITY_UNKNOWN",
    "QUALITY_GOOD",
    "QUALITY_POOR",
    "QUALITY_DISABLED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0313132120211102-2310033011310301-0230221312013112-3023031313011003-1120300112310233-2301323120302310-1233231230013133-3211011232221101"></a>

<a id="canonical-2110301013123030-1002013113010223-2103121201133033-1031112133021000-2131101013201012-2123020030230022-1223101301122311-2301120101011231"></a>

#### `infra.hw_info.network.link_type` property

Type: `"string"`. Optional.

\[Enum:
LINK\_TYPE\_UNKNOWN|LINK\_TYPE\_ETHERNET|LINK\_TYPE\_WIFI\_802\_11AC|LINK\_TYPE\_WIFI\_802\_11BGN|LINK\_TYPE\_4G|LINK\_TYPE\_WIFI|LINK\_TYPE\_WAN\]
Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of
type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link. Possible values are
\`LINK\_TYPE\_UNKNOWN\`, \`LINK\_TYPE\_ETHERNET\`, \`LINK\_TYPE\_WIFI\_802\_11AC\`,
\`LINK\_TYPE\_WIFI\_802\_11BGN\`, \`LINK\_TYPE\_4G\`, \`LINK\_TYPE\_WIFI\`, \`LINK\_TYPE\_WAN\`.
Defaults to \`LINK\_TYPE\_UNKNOWN\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LINK_TYPE_4G","LINK_TYPE_ETHERNET","LINK_TYPE_UNKNOWN","LINK_TYPE_WAN","LINK_TYPE_WIFI","LINK_TYPE_WIFI_802_11AC","LINK_TYPE_WIFI_802_11BGN"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("LINK_TYPE_UNKNOWN",
    "LINK_TYPE_ETHERNET",
    "LINK_TYPE_WIFI_802_11AC",
    "LINK_TYPE_WIFI_802_11BGN",
    "LINK_TYPE_4G",
    "LINK_TYPE_WIFI",
    "LINK_TYPE_WAN"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "LINK_TYPE_UNKNOWN",
  "enum": [
    "LINK_TYPE_UNKNOWN",
    "LINK_TYPE_ETHERNET",
    "LINK_TYPE_WIFI_802_11AC",
    "LINK_TYPE_WIFI_802_11BGN",
    "LINK_TYPE_4G",
    "LINK_TYPE_WIFI",
    "LINK_TYPE_WAN"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203301131111000-1130020200233213-1200021220010211-3301301020201211-2312300102133302-1230021323122212-3113300322121023-3222001223203313"></a>

<a id="canonical-2032323110211001-1313320211311222-1320233233113301-0332221323210232-1320213222033211-0302032331321202-2003002211332021-2320031233131103"></a>

#### `infra.hw_info.network.mac_address` property

Type: `"string"`. Optional.

MAC Address. MAC address on interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(17, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`),
    ""),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "formatDescription": "MAC address (e.g., 00:1A:2B:3C:4D:5E)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3000323122232333-0223131333131020-1023130031311112-3032033320323101-1301232123223012-3003221130301313-1033003303101331-2220002122132023"></a>

<a id="canonical-3231022102230010-3222021100131331-0210332312120122-0101121031302021-3012102030330021-2131003133101020-3112210012213112-3222332020001223"></a>

#### `infra.hw_info.network.name` property

Type: `"string"`. Optional.

Name. Name of device, eg. Eth0.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3221112000321330-0231011010202303-2300201232023200-2100032132320313-0030221013212222-2100023231231331-3103032213322003-1111231210001221"></a>

<a id="canonical-1003203213001210-2321323231322200-0323121030022121-1032011031101033-2022130122331230-2302110323321310-2320331321333213-2213310222320333"></a>

#### `infra.hw_info.network.port` property

Type: `"string"`. Optional.

Port. Used port, eg. Tp.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2203120022030213-1133232112000002-0330313031231122-3220011303112303-3012223302221021-2031122113221222-2201231032300021-2313300012002310"></a>

<a id="canonical-1003112213112111-1213000202131223-1211111101003301-2230322331130001-2032212302203302-0031123210220101-2300102013213111-1310033211100032"></a>

#### `infra.hw_info.network.speed` property

Type: `"number"`. Optional.

Speed. Device max supported speed in Mbps.

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

<a id="canonical-1231001033023312-2011113220003311-0222030233001030-3203302232031331-1331210103302203-2123123330013022-3000132230133322-3213211033011311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.os` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.os

<a id="canonical-1231002013131030-1012200120133232-1013121131001033-1330312002012230-1310121133131001-2302321110031022-3210200302322210-3131022121331302"></a>

Type: `"object"`. single nested block, Optional.

OS. Details of Operating System.

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
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-2022222320013313-3002331103011033-2123301123312102-3323230023221101-2012130320103110-1102021013132011-1212113313031300-0102332223030213"></a>

### Direct properties for `infra.hw_info.os`

<a id="canonical-1033132130023221-0210011231332030-3011013301013100-3212111212020113-0121213020203100-2111113110233230-3311003232113131-2130330111333131"></a>

#### `infra.hw_info.os.architecture` property

Type: `"string"`. Optional.

Architecture. Architecture of OS.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3002311233011223-0001330221310211-0032002201021110-0210212102320311-1301200230012021-2102203323323231-1033012313012013-3323233223210100"></a>

<a id="canonical-3130313213313131-3300323132312101-0232013003331300-3122120121103013-2330213222012012-0333020021230232-0033221133200030-1001310103030222"></a>

#### `infra.hw_info.os.name` property

Type: `"string"`. Optional.

Name. Name of OS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3303211031231023-1302112133010200-3101102031010200-1110333222320000-0011023300103331-0333311132330230-0201013032230133-1103301130000033"></a>

<a id="canonical-0232111202013331-2221200010230311-2310002112200021-0033130301123302-1131103122202330-2111020330201333-3123001010210310-2312220020001003"></a>

#### `infra.hw_info.os.release` property

Type: `"string"`. Optional.

Release. Release of the OS.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1131110322033201-0030120010000100-0032132120020110-0031332222120210-2133223023301330-1210203113131120-1131133121111122-0113301210030031"></a>

<a id="canonical-1220022222120032-2113332031320133-3130003331103331-1310231300331032-1232022230131033-2130200300122303-1213121022301220-2021112030321232"></a>

#### `infra.hw_info.os.vendor` property

Type: `"string"`. Optional.

Vendor. Vendor of OS.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3033332132221230-1013231000122301-3331103201300131-1002233303133233-1312303032033333-2331311231113233-3301210132023301-1212322323233321"></a>

<a id="canonical-0133111113220013-2333010301232310-0001333021221222-2331210201221220-3013011310111012-2333201310021101-0230233311022133-2220101033303313"></a>

#### `infra.hw_info.os.version` property

Type: `"string"`. Optional.

Version. Version of OS.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3111131021132120-1201132311302220-0322330011200210-0332003110023033-3210113133133231-1112310100321012-1201113230332133-0200001023220311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.product` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.product

<a id="canonical-0223022230303020-1030223112321332-3012233303131121-1113023212121213-2201230020002102-1231220012313210-0300132032223303-2200332121331001"></a>

Type: `"object"`. single nested block, Optional.

Product Information. Product information.

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
product {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011222032003012-1210030023023330-3020112310032320-3001221130131321-3133310223012122-1010233010112002-2110030333122210-3202030031000322"></a>

### Direct properties for `infra.hw_info.product`

<a id="canonical-3320201133032300-2220212120221011-0212230233302223-1210130032100222-2003130202223032-0132223330213122-3030323131013312-2033113210123231"></a>

#### `infra.hw_info.product.name` property

Type: `"string"`. Optional.

Name. Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product\_name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3020212003232113-2032011200122001-0322011333021323-2020102032011132-2031203212123332-0111102133111001-0332332233123332-1103323032321110"></a>

<a id="canonical-2310323102032201-2223320210001330-1210210200232012-1113230320103130-1030320203330200-3113012010201302-3313123200002130-3101320333313122"></a>

#### `infra.hw_info.product.serial` property

Type: `"string"`. Optional.

Serial number, eg. For AWS 00000000-0000-4000-8000-23460f645a1f. Info taken from
/sys/class/dmi/ID/product\_serial.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2002032011330333-0123132102111223-2030113310223021-3330031020100212-3220303212232300-3101031322001230-3231011101032101-1231033331310231"></a>

<a id="canonical-1301122213300323-3003013022310301-2011021011233013-1102132201121103-2120133320021020-2103211002000310-3031332021021330-2233302122111031"></a>

#### `infra.hw_info.product.vendor` property

Type: `"string"`. Optional.

Vendor. Vendor name, eg. For AWS Amazon EC2. Info taken from /sys/class/dmi/ID/product\_vendor.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1110330101311202-1010002111332003-2221231100123123-1312003203121110-1233321332123320-3032113133033000-2221023323030302-1232020332121310"></a>

<a id="canonical-0220312132220010-2130130300031030-1120132310222303-0310221302323020-2020232312030023-2132311033312132-3021321101130312-0213300013333323"></a>

#### `infra.hw_info.product.version` property

Type: `"string"`. Optional.

Version name. Info taken from /sys/class/dmi/ID/product\_version.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3313133121200132-3301001210013020-0000211123132303-3331332012323121-0112110122022030-3033200103123122-0100332030330030-2332130132231120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.storage` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.storage

<a id="canonical-2210002233111210-2021333332211311-3230033103002002-2023333320010232-3223200201021120-3002222231113303-2232220201223331-2320301101331221"></a>

Type: `"object"`. list nested block, Optional.

Storage. List of storage devices in server.

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
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033012001213230-1122023331033110-0200313033010033-2330203130313322-2322312101222100-2103110203231120-0113101020011203-2323323201000202"></a>

### Direct properties for `infra.hw_info.storage`

<a id="canonical-0222212030312002-3213030333033020-1222220311211302-3223232001301113-0233303002113013-0131330202321033-3132230302130111-0022232021121033"></a>

#### `infra.hw_info.storage.driver` property

Type: `"string"`. Optional.

Driver. Driver of device.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3133210030133321-3230130013211233-1001032303002010-2133302132213200-3023302113100332-2033130021022211-3112010111232302-3300131200322311"></a>

<a id="canonical-3133302131201200-1001323220300322-3201320133130020-1113233013011030-1012223120023033-1013112302203330-0220130333121001-1323020121231210"></a>

#### `infra.hw_info.storage.model` property

Type: `"string"`. Optional.

Model. Model of device.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3033200311311233-1322111110102233-2313231012212123-1333331210013302-1010000330302001-3201313001021121-0301113131212130-0330002203132212"></a>

<a id="canonical-2221323000100312-1133022130021230-1101023230223112-0310120030303110-2330311201203203-2323013032323221-1130312202023111-3232121103032203"></a>

#### `infra.hw_info.storage.name` property

Type: `"string"`. Optional.

Name. Name of device, eg. Nvme0n1.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1222013101030312-3002331213123301-1030220202131112-2002311322231231-3011111210312303-0300231003000310-1103313001000333-3211002213200311"></a>

<a id="canonical-0231113112113032-0200211331120213-2131311112222312-3323031123300032-1110013001033220-0013312303322203-2201213330023312-3222211312221212"></a>

#### `infra.hw_info.storage.serial` property

Type: `"string"`. Optional.

Serial Number. Serial of device.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3123232121232020-1123121002310223-3011212330033131-3220200321221013-1022321101221213-2210210221220230-0111130101010131-0001113000201011"></a>

<a id="canonical-3320100330231113-0220113002212210-3311213030221100-0013031022020301-2331030233010110-2122302011233022-0333301133331223-1310212031200033"></a>

#### `infra.hw_info.storage.size_gb` property

Type: `"number"`. Optional.

Size(GB). Device size in GB.

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

<a id="canonical-2101230313213023-0023012031212112-2333232333130212-2023023013030212-1002132330002210-0333122003003223-0122211200030232-3003123221000032"></a>

<a id="canonical-3300320022231311-0000332022132322-2110102103021231-1023033213023321-0002031120223200-3323232322121031-0002032211033002-2320002130332001"></a>

#### `infra.hw_info.storage.vendor` property

Type: `"string"`. Optional.

Vendor. Vendor of device.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2133131333302012-2311223101000032-3011003103231233-1102312333003101-0202323211123222-0220232101330330-2012330211102002-3032003002021300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.usb` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- [infra.hw_info](resources--registration--reference--group-001.md#canonical-1320103332320020-0302112103213001-2102130003310211-3213300320113023-3310101202221211-2030122310002101-0111000220232133-2120111022213221)
- infra.hw_info.usb

<a id="canonical-3231321232012313-2312132120033121-1202131002332203-3232103331033212-1210301032021032-1201222312123202-2220313303210231-1322311111022311"></a>

Type: `"object"`. list nested block, Optional.

USB devices. List of USB devices in server.

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
usb {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222031002322100-0013111103031121-3320230030100332-1201011220032330-2030220310102301-2012003112131023-2023120131202313-0100013311010233"></a>

### Direct properties for `infra.hw_info.usb`

<a id="canonical-3113013010223302-1102301210133031-3321213132023302-3330312230110130-0111221212233210-1031313033222111-1100100222221222-1011123322201001"></a>

#### `infra.hw_info.usb.address` property

Type: `"number"`. Optional.

Address of the device on the bus in decimal.

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

<a id="canonical-0231322113220120-0032133223320131-2322210333332332-0221103002101320-2030312031212210-1020121010312313-0213023232311302-2331122323333212"></a>

<a id="canonical-1301100020311200-3331212031311003-2120323103330110-2332022132201201-0033330322210031-0112111011233100-1011032203201202-0232130213311013"></a>

#### `infra.hw_info.usb.b_device_class` property

Type: `"string"`. Optional.

Class. The class of this device.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1111031102110312-3102323001211213-3222110011211311-0110113321022213-2103123103332031-0120112003203021-0103202002003011-2032310211111330"></a>

<a id="canonical-0320101131000220-2020230230232231-0330100203311331-3131132123303102-0100123301332321-1113010330313112-0131131132300221-1220222012323112"></a>

#### `infra.hw_info.usb.b_device_protocol` property

Type: `"string"`. Optional.

The protocol (within the subclass) of this device.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1030112323233230-1330313132330122-1312313021331113-3001003322003133-2003010312131022-0020201103321033-1101112103323103-0111310011130302"></a>

<a id="canonical-1130202320333123-2232323030313303-1200102222333030-3311023122223310-3000302103332303-2311110030011021-3322031021312221-2201132113313100"></a>

#### `infra.hw_info.usb.b_device_sub_class` property

Type: `"string"`. Optional.

The subclass (within the class) of this device.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1103122201123220-2132323032202002-0213302120011010-2033011012311211-2122110023010132-3002020021122121-3023000203132110-1003022110202113"></a>

<a id="canonical-0110111000232113-3131123022000003-3323332300220113-2021221301021221-2132030023320320-3010130220101021-3123232100020013-1230011303203100"></a>

#### `infra.hw_info.usb.b_max_packet_size` property

Type: `"number"`. Optional.

Max packet size. Maximum size of the control transfer.

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

<a id="canonical-1332330322221322-3101112130320212-3311301202320122-2323031022232030-2012303122131001-1010322332012320-3210202120211112-0331022022122310"></a>

<a id="canonical-3001132222301320-1201222000223333-1111330323021330-1213013000110302-2230310233223211-2330333212231122-1101000131112032-0130123000223320"></a>

#### `infra.hw_info.usb.bcd_device` property

Type: `"string"`. Optional.

BCD Device. The device version.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0002330022032202-1013010030303201-1302313010130121-2310210002103211-0133010131301122-3033003012233211-1223333002231023-3003121321233313"></a>

<a id="canonical-2301101320321220-2300203130200030-1101020111323330-1202020122012300-3230231322213131-3323020200002313-0132311231333330-1121101110231332"></a>

#### `infra.hw_info.usb.bcd_usb` property

Type: `"string"`. Optional.

BCD Spec. USB Specification Release Number.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0023232003222321-2022031122123102-3232103330001032-0010222023320123-2123102232310333-0002332101201010-1320211003313131-3122332001100210"></a>

<a id="canonical-2300220323021011-0103303300222012-3030302321033102-0330110201213203-0122313132312322-2233131130113320-1132322031133010-2303033330101110"></a>

#### `infra.hw_info.usb.bus` property

Type: `"number"`. Optional.

The bus on which the device was detected in decimal.

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

<a id="canonical-3102130331112233-2221131012230333-2012202231112122-2303112211311013-2000113301202101-1300123000033331-2033330131203002-3301110133330023"></a>

<a id="canonical-3212003321113121-2032021000212203-2131020221302021-2123230032132001-3202111221202303-0100132212132121-1122131220102213-0033003202000030"></a>

#### `infra.hw_info.usb.description_spec` property

Type: `"string"`. Optional.

Description. Device description.

<a id="canonical-1031220202230300-1102311303021023-0102211302003003-0221022311023020-0200302111331110-1301031013031212-2030101010010120-1203300121031002"></a>

<a id="canonical-3010113210103302-1212301221300212-3021012311002021-0020112321222102-0200020003332222-3132212032303111-3023121310300013-1110201311011232"></a>

#### `infra.hw_info.usb.i_manufacturer` property

Type: `"string"`. Optional.

Manufacturer. Manufacturer name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2200303203013221-3011223030020302-3301231013003300-2013220301131211-1031301310330231-2102033313002112-0132123310003133-3133020202211013"></a>

<a id="canonical-1301221122100020-0133230323221323-0031333200311021-1103332211103112-2020112202001011-1310331333123013-1002232310000230-1020201001311022"></a>

#### `infra.hw_info.usb.i_product` property

Type: `"string"`. Optional.

Device product. Product name reported by device.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1120223203232200-2011233333202030-0023033132100233-3113321330132320-1203031130212212-0113123221013201-2331123031131023-0113100010223112"></a>

<a id="canonical-0131301330301320-0101300112103312-0031230101101121-2230202332111103-0101133103110011-0233311222231333-3000221303013302-0101301021030001"></a>

#### `infra.hw_info.usb.i_serial` property

Type: `"string"`. Optional.

Index of Serial Number String Descriptor.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0222001113120202-1020230122222100-2301013223101010-1112013303021313-1300130012012232-0312103132002102-3222011010100100-0033300110022201"></a>

<a id="canonical-2133130303320022-3312023231323212-1302110200130230-1133212011330130-1110003101000323-0211120220331112-3303010313310022-1113232231111030"></a>

#### `infra.hw_info.usb.id_product` property

Type: `"string"`. Optional.

Product ID (Assigned by Manufacturer) in hex.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1131020001212132-0100213323002212-3322301010233323-3233333321333101-3033110203321203-0013211231232121-0301110013322213-0300230011100113"></a>

<a id="canonical-2230113101033221-0111200001300111-0303322032012122-0223223222100331-0210210200030113-1020303310111020-0000210332010233-3133203001023200"></a>

#### `infra.hw_info.usb.id_vendor` property

Type: `"string"`. Optional.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1012312020201200-2301132302113132-2032310032100000-0023020010323132-3103321020222203-0201102232231111-0312331223012130-1020312202111312"></a>

<a id="canonical-1003320313201200-2220300321331000-3112302031102313-1021331031201233-3323323211021323-1220033332323011-2303210102310230-2010312020001003"></a>

#### `infra.hw_info.usb.port` property

Type: `"number"`. Optional.

Port on which the device was detected in decimal.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2023111331223130-3213223100013232-3121032131332001-2311310102303121-1131232311002332-3233231201221023-1032022231100030-3031113133311223"></a>

<a id="canonical-2211320200021221-0012322330022322-3123220111220102-2332332321000120-2231101021021303-3111023001113012-1222003022013011-2100330021323101"></a>

#### `infra.hw_info.usb.product_name` property

Type: `"string"`. Optional.

Product ID translated to name (if available).

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1031011333310320-1201000300302133-1213020303020010-3002103302022311-0112202300203311-1222121232111003-0323231031310210-1210331021101212"></a>

<a id="canonical-2010213222312031-2010000302203220-2032131223310311-3323311032311200-1310131323100333-3202012302021101-2101313333113310-0100032231231122"></a>

#### `infra.hw_info.usb.speed` property

Type: `"string"`. Optional.

The negotiated operating speed for the device.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1230302310103203-3221020230003021-3133203100313001-3213001323120210-0132302132023201-1332102013032030-0020211313320232-0000213213003212"></a>

<a id="canonical-3101213010001331-3210000230000222-0230103220101001-0032003110031022-0111122130101000-3320023021312001-0310202320313111-0320320230303122"></a>

#### `infra.hw_info.usb.usb_type` property

Type: `"string"`. Optional.

\[Enum: UNKNOWN\_USB|INTERNAL|REGISTERED|CONFIGURABLE\] Type of USB device Unknown USB device type
Internal USB present in Certified HW USB device present during node registration USB device that can
be matched by USB rules. Possible values are \`UNKNOWN\_USB\`, \`INTERNAL\`, \`REGISTERED\`,
\`CONFIGURABLE\`. Defaults to \`UNKNOWN\_USB\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONFIGURABLE","INTERNAL","REGISTERED","UNKNOWN_USB"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN_USB",
    "INTERNAL",
    "REGISTERED",
    "CONFIGURABLE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNKNOWN_USB",
  "enum": [
    "UNKNOWN_USB",
    "INTERNAL",
    "REGISTERED",
    "CONFIGURABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0213112132023322-0010001100110312-3313020320321123-3023131233230131-2002212131031233-3031322301102202-0323100023020223-1311223200000100"></a>

<a id="canonical-3232113122300033-3122123012211322-1032220233032032-0203012211100130-3030211321313322-3331122023330232-0321200312321111-2201332033010111"></a>

#### `infra.hw_info.usb.vendor_name` property

Type: `"string"`. Optional.

Vendor ID translated to name (if available).

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1131111001122100-2210113013321231-2011202021020021-1321111013331333-2103232202022321-0113302202013102-0200012200312100-0002212230222000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.interfaces` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- infra.interfaces

<a id="canonical-2331111231202111-2121032333223302-0120032023030011-3003031122111101-0302231221220001-0010212123212333-2000102111333313-2113122313001020"></a>

Type: `"object"`. single nested block, Optional.

Machine interfaces present during registration time.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
interfaces {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333232120213313-2131312201312220-0332010123213011-1100232332221223-3130030323031202-2201003302332033-3110013202300010-1101212310212223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.internet_proxy` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- infra.internet_proxy

<a id="canonical-0002201220031220-0311321212103020-0132032103302121-2100201333303033-0321300213220222-1020130311313233-1330200233233211-0122103033133202"></a>

Type: `"object"`. single nested block, Optional.

Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.

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
internet_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131101200213213-3110103200010202-3111213302200332-1210032100130230-2201110011321313-1112103330113012-0033100001001330-2012120231122131"></a>

### Direct properties for `infra.internet_proxy`

<a id="canonical-3323021301131203-2011331110200021-3223111112132032-2311002232210301-3032202103002232-3321221200110310-1011212221330103-3032310313123300"></a>

#### `infra.internet_proxy.http_proxy` property

Type: `"string"`. Optional.

It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by
HTTPSProxy or NoProxy.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2201213101223201-1330030120133123-3313300102031320-1213000030311011-1031220300311200-2011010313222232-1003103031100200-3021223203332231"></a>

<a id="canonical-3213120112013231-1032311132332112-1330333020231333-0023202303022002-1101201113101113-1333002300011320-1330232110200222-3210233220103203"></a>

#### `infra.internet_proxy.https_proxy` property

Type: `"string"`. Optional.

It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1223201133213223-2131023213322122-2133120300133023-3210310102132303-2023310330330210-1310202132021333-2232212112122303-3110023202201033"></a>

<a id="canonical-1131033021123323-1013021101321130-2032131131221022-1020120123311113-2123001202211230-2202012201130333-2122212122101330-0321331013203320"></a>

#### `infra.internet_proxy.no_proxy` property

Type: `"string"`. Optional.

It specifies a string that contains comma-separated values specifying hosts that should be excluded
from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix
in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (\*). An IP address prefix
and domain name can also include a literal port number (192.0.2.103:80). A domain name matches that
name and all subdomains. A domain name with a leading "." matches subdomains only. For example
"example.com" matches "example.com" and "bar.example.com"; ".y.com" matches "x.y.com" but not
"y.com". A single asterisk (\*) indicates that no proxying should be done.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0133322313233122-2120213122323200-2101210001122201-1110202101000110-0323100103122313-1002212233300231-3230212332320103-1101220233301021"></a>

<a id="canonical-3020123112033000-2321322331223111-3022111210212003-1033212131301213-2112100311000313-0122210303303320-1221030302023021-2233230210133330"></a>

#### `infra.internet_proxy.proxy_cacert_url` property

Type: `"string"`. Optional.

Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate
value.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0312321212221203-3210322201111303-0121010122010333-2311320312012111-2333222213001230-0231301313330001-3231102122010310-3030021133012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.sw_info` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [infra](resources--registration--reference--group-001.md#canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111)
- infra.sw_info

<a id="canonical-0113320030321313-2130131001132220-1301302100322010-1022222203120121-0312000010201230-1110113230201111-0120213020102132-0302123230130110"></a>

Type: `"object"`. single nested block, Optional.

SWInfo holds information about sw version.

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
sw_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113003222232030-1213002313230232-0111230300221011-1330003210110201-1013020012331312-1110201331111013-0013100002111101-0131323330112330"></a>

### Direct properties for `infra.sw_info`

<a id="canonical-3000100311220031-0122020223302211-3231201113031221-2213023201113033-0121121213220132-1000313202312031-2000210200120101-1232110020322332"></a>

#### `infra.sw_info.sw_version` property

Type: `"string"`. Optional.

SW Version. SW Version in the site.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1120101321201223-0133323222022210-1001000011103111-2322312323001321-3201033220310002-2220113311003213-3101110103313231-1211213320021213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `passport` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- passport

<a id="canonical-3300131212131333-0313002231120031-2131313131131033-0100110033123313-3112220301320333-2221002321232100-1212121011013311-1102212003030200"></a>

Type: `"object"`. single nested block, Optional.

Passport stores information about identification and node configuration provided by CE during
registration. It can be manually updated by user during approval.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_name",
    "cluster_type",
    "latitude",
    "longitude"),
  validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version"),
  validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]",
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
passport {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301110133211020-3203133002100033-0301031210302123-2313120100122321-3231123031322232-0300101332201332-2210021331323210-2322213231211121"></a>

### Direct properties for `passport`

<a id="canonical-2302100033321203-2203001330303231-1001130022213200-0310130010210003-2033020122021023-1301221312200300-3312300112031330-0200201333333002"></a>

#### `passport.cluster_name` property

Type: `"string"`. Optional.

Cluster Name. Human-readable name for the resource

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2323320022121133-2011012302120201-1321332201032002-3212101300231002-1220313020133000-3112310002330200-3313102203210103-0332311300233230"></a>

<a id="canonical-1213032201113020-3002100210120203-0223013012310302-3232322201001033-3111323100302203-1312310113010101-1022000323010331-0123320000011300"></a>

#### `passport.cluster_size` property

Type: `"number"`. Optional.

Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single
master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time,
cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after
installation. It does not interact with auto-scaling as only pool nodes are scaled.

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
    "ves.io.schema.rules.int32.in": "[0,1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.in": "[0,1,3]"
  }
}
```

<a id="canonical-3130201131232212-1013011210032112-3001322233122301-3322021302303231-0023321002020113-2023333132023313-0012231201032200-2003101100001023"></a>

<a id="canonical-2220233112020320-0211102221123202-0021021012211031-3212013220323002-3122201123223201-2212112231120022-2003222320330131-3331003301130001"></a>

#### `passport.cluster_type` property

Type: `"string"`. Optional.

Cluster Type. Cluster or grouping configuration

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [default_os_version](resources--registration--reference--group-001.md#canonical-0121101322021220-2201332112213101-2230331311023230-1311323103322112-0302220111230211-0313313133223230-2003203200110302-0331003203032000): complete subsection reference.

- [default_sw_version](resources--registration--reference--group-001.md#canonical-1230302223233311-0121113213210021-1101003111002300-1122011123120233-2323322020031012-1222330320111013-1132013300303123-1313001033012230): complete subsection reference.

<a id="canonical-1022101030210021-0213322322131131-3111301332022330-0012032133332032-1101113100103203-1330220013030100-1332330112211220-2010111010111013"></a>

<a id="canonical-1001323202231222-3322123000222130-2133223222321101-0020003111010122-0131001033022101-3222001230202031-1030130313133313-3032321113113233"></a>

#### `passport.latitude` property

Type: `"number"`. Optional.

Latitude. Geographic location of this site.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2130222322100100-3012202313203123-0030330021333321-0231121303022010-2110230021010021-2023223211013232-0032011122231331-2231112123002311"></a>

<a id="canonical-2021103032320130-3323300012120323-1031320031211000-1332310320221210-3022123310100203-3023123323012132-1211202113013300-1113211212333212"></a>

#### `passport.longitude` property

Type: `"number"`. Optional.

Longitude. Geographic location of this site.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2133031023320301-3200102232001302-2100332232301102-2223321223222321-2102122121230123-1332030030221132-3203213033211032-3012020321302322"></a>

<a id="canonical-3300012210122000-2320111002103023-3221320031232012-1111321001000313-3320333000123230-3213333033102302-1121211023210112-3111313331213110"></a>

#### `passport.operating_system_version` property

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-2010230203200001-0023102122320111-1133220331103100-3130013333112102-2231221001312303-2031110120330131-2312103002103113-0310033330212010"></a>

<a id="canonical-3121102300122201-1210010320310202-0020112103311022-1113101130323220-0333301221002202-1101231000320222-2120110210112321-3001013010331322"></a>

#### `passport.private_network_name` property

Type: `"string"`. Optional.

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1101110132201203-2101211202231311-1212211323003313-2212323000311321-3122132123331211-2111011311123112-2033023301111020-3312312313002230"></a>

<a id="canonical-1012201201311323-1311110110223033-3003011122211310-3123322110001131-3203332031133303-3120331310233122-2022321111003213-1122232022322102"></a>

#### `passport.volterra_software_version` property

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-0121101322021220-2201332112213101-2230331311023230-1311323103322112-0302220111230211-0313313133223230-2003203200110302-0331003203032000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `passport.default_os_version` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [passport](resources--registration--reference--group-001.md#canonical-1120101321201223-0133323222022210-1001000011103111-2322312323001321-3201033220310002-2220113311003213-3101110103313231-1211213320021213)
- passport.default_os_version

<a id="canonical-2130131311031330-0102232320003333-1100023330123302-3010123012002113-3231131133033100-3003021223323232-0131311112231311-3101302120000012"></a>

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
default_os_version = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230302223233311-0121113213210021-1101003111002300-1122011123120233-2323322020031012-1222330320111013-1132013300303123-1313001033012230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `passport.default_sw_version` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [passport](resources--registration--reference--group-001.md#canonical-1120101321201223-0133323222022210-1001000011103111-2322312323001321-3201033220310002-2220113311003213-3101110103313231-1211213320021213)
- passport.default_sw_version

<a id="canonical-0213132232332322-2130103232200210-1300321021112202-3200130020223320-0301032322001111-0231133013300322-1112333102121301-3331303320133220"></a>

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
default_sw_version = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213031231333002-0230321022300100-0023012110121303-3101201022223132-2020331103120330-0112000313011320-2032122330213002-1312312001121102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Property reference](resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- timeouts

<a id="canonical-0332003222321313-1030212102212103-1010221332320312-1003230100020322-2102221312100020-1231222120103011-3022030223011110-1021223023300230"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131131112221120-1330331230203020-0320100100202212-3320312323202231-0003203311323010-3100220322320213-2110201232120011-2000322013010111"></a>

### Direct properties for `timeouts`

<a id="canonical-1321232120333203-2021202311012131-2003113313002130-2120130023312230-0112103122003322-0113032333302321-3211201332321031-0022023230020203"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3133010302200303-2130030310032013-3220233200212313-3301222221003332-1201101231222201-3231013203230121-2212302311321021-1232110122133333"></a>

<a id="canonical-3302333110013300-2120110110333122-1031302310032232-0213210023011202-2321211320130021-1331103000013020-3212002121011323-2222100111312132"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2033333132000101-1310312333321023-3010012200032110-3123123330233111-1122200201112003-0021001030233213-0233330110223000-3231320302322101"></a>

<a id="canonical-1113221311023232-1001210112200320-3233131201320200-1300100121200330-3012233120200133-3002332333122312-2211332331130120-1001320120011212"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0030003213133300-3301000131230110-0133010212302331-0013102303221233-0130133100012111-1112132203311223-1201133311302103-1220031003113210"></a>

<a id="canonical-1011021102020233-3000002321222301-2103002312222313-2333331021011303-3132002201232113-0323210021131103-2221120023203210-1132010220003211"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
