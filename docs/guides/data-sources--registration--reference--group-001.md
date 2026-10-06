---
page_title: "xcsh_registration reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration reference."
---

# xcsh_registration reference

<a id="canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- Property reference

<a id="canonical-2312322010101012-0311122301202312-3312230101120313-2303212133123130-0301300013202130-1103120321201000-3020012021332002-3013211221121122"></a>

### Direct properties for `xcsh_registration`

<a id="canonical-0031202300031033-2001210302033001-2202113102320003-3032033310211010-2323123032202133-2123302002210001-2333002221023111-0220130003110112"></a>

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

<a id="canonical-0322122100231031-2302233132030013-0303103302322210-3113131103013213-2222333132033110-2311011200300021-2333212303222010-3320130332232102"></a>

<a id="canonical-0323233303132310-2230300002312000-0033121001313133-3212302111230330-2020233300330112-2201112122011332-3331100023121323-3210032213111221"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Registration.

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

<a id="canonical-3010331002122120-2330311222220132-0110300103301030-3203111232333221-2130030113310133-3121212130023031-2020201200303322-3210212213232320"></a>

<a id="canonical-0232332330112130-2213020023311013-1122313122112233-1300100103020310-0133003203330330-3222123300133321-3100222201013333-1022332000033031"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101): complete subsection reference.

<a id="canonical-2223020132121332-0132033301203311-3220113320101032-2022211310232003-1321321300322303-0110203310001123-0010210020211211-3301013020201323"></a>

<a id="canonical-2010310023322331-3101231103303130-0322310322022103-2330033300130200-2032300321111120-1320003220110311-0030033310001221-3133113102311312"></a>

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

<a id="canonical-2012322121012313-0001100230323323-0213212101313230-1020233222013031-3312123031213000-3233122111322013-2221210033113332-2323003132101333"></a>

<a id="canonical-0311331203333303-3032020312111122-0110030312110120-1120001120102032-3201033313000332-3202302130331210-0320220313200222-3200011010302000"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Registration.

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

<a id="canonical-0010110123221301-2000302210223212-1012132031123232-2111100032322223-0011302132131022-1133202303011020-2013230201131201-3031203030232202"></a>

<a id="canonical-0211210312320332-2033111010003033-0222311233011032-3210231231023101-0011002220303033-1032222020133310-3003201130121132-3323300113302223"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Registration exists.

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

- [passport](data-sources--registration--reference--group-001.md#canonical-0223232121310231-1203011310020223-0003303003232021-2210303010321002-0213203010232110-3211330322101103-3313301302301300-0202203223232211): complete subsection reference.

<a id="canonical-3302320203230301-1332012020212013-2202001212322020-3222010220322101-2000023331032012-1200221113022103-1331201111313102-2302013010213311"></a>

<a id="canonical-3011311121011120-0022233323002011-0130323021121021-2321331321210121-1210001122302310-0003331322003131-1033010020300303-1212212221211232"></a>

#### `token` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0302022301310333-2332132313333330-2101012213303311-2211203003023113-0301012001010333-2003121012011213-3001310011302020-1233010311223201"></a>

### All schema paths for `xcsh_registration`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--registration--reference--group-001.md#canonical-0031202300031033-2001210302033001-2202113102320003-3032033310211010-2323123032202133-2123302002210001-2333002221023111-0220130003110112) |
| `description` | [description](data-sources--registration--reference--group-001.md#canonical-0322122100231031-2302233132030013-0303103302322210-3113131103013213-2222333132033110-2311011200300021-2333212303222010-3320130332232102) |
| `id` | [ID](data-sources--registration--reference--group-001.md#canonical-3010331002122120-2330311222220132-0110300103301030-3203111232333221-2130030113310133-3121212130023031-2020201200303322-3210212213232320) |
| `infra` | [infra](data-sources--registration--reference--group-001.md#canonical-3112222200000320-1111111333003022-2331222031103212-2220131210213222-3101031232133121-2112022212001311-1230331211120220-3023021133322232) |
| `infra.availability_zone` | [infra.availability_zone](data-sources--registration--reference--group-001.md#canonical-3102030321312222-3312110223132133-2020331000033003-2112030310212002-2023313302102113-3320210220212221-3220220120023103-3301321221022233) |
| `infra.bond_config` | [infra.bond_config](data-sources--registration--reference--group-001.md#canonical-2011122311112200-2230130331013302-2331320010101333-3211031100313120-1201221311120131-1113310300132331-1223031310012130-3321133013302313) |
| `infra.bond_config.interfaces` | [infra.bond_config.interfaces](data-sources--registration--reference--group-001.md#canonical-3313102011123030-0112131103101222-2031220202121031-0100202302213102-3110001011230000-0201002002123022-1300213200031312-1211113233120102) |
| `infra.bond_config.mode` | [infra.bond_config.mode](data-sources--registration--reference--group-001.md#canonical-3201231130232022-2200003231111122-3331100300212310-0120320333300321-2012100301220313-3130202222022333-1123221003222120-1112132330130023) |
| `infra.bond_config.name` | [infra.bond_config.name](data-sources--registration--reference--group-001.md#canonical-2321122021132300-2233120222123012-1100122033103023-1220003311123203-1010313123311312-3121003033122001-1232210031221103-2303103202002313) |
| `infra.certified_hw` | [infra.certified_hw](data-sources--registration--reference--group-001.md#canonical-0132323102031312-1200022013122233-2300310301122131-1000303212113130-1210131112223133-2113230121002223-3320323301132320-1120111222011020) |
| `infra.domain` | [infra.domain](data-sources--registration--reference--group-001.md#canonical-2112113223301323-2121211013330220-2103000112113223-3301331200100032-2131010000223023-1120230010300020-1331303011223020-3131110022232231) |
| `infra.hostname` | [infra.hostname](data-sources--registration--reference--group-001.md#canonical-0211023101310320-3213322120311301-2030121013032300-0213110002313102-1021222321110002-2110322302201200-0033100220213100-3313010200012120) |
| `infra.hugepages` | [infra.hugepages](data-sources--registration--reference--group-001.md#canonical-3023201330231122-0300313230211012-3112332121023123-1310002202030323-0000032021323213-0330010101020112-1003131101330320-1113012100330102) |
| `infra.hugepages.free` | [infra.hugepages.free](data-sources--registration--reference--group-001.md#canonical-1101222220102313-3010121031020322-3313023130201123-0302120132210130-3101020100010212-3022013011122203-3210031211112300-2030201312202332) |
| `infra.hugepages.page_size` | [infra.hugepages.page_size](data-sources--registration--reference--group-001.md#canonical-0011131130011030-3113221332212113-2330013200123032-2020331030113331-3322330331300033-0010113230331332-1112211333310200-3101201322322033) |
| `infra.hugepages.total` | [infra.hugepages.total](data-sources--registration--reference--group-001.md#canonical-3330231001322203-2201132303302323-2302222003001223-0203201213201001-2103203201333012-0332112011223101-3202111102031130-0111002202332011) |
| `infra.hw_info` | [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-2302313323300202-2222112222012030-1123102301303202-1111323301031310-1100001230213121-1111202222323020-3321203310211110-2121202001012133) |
| `infra.hw_info.bios` | [infra.hw_info.bios](data-sources--registration--reference--group-001.md#canonical-0211320230322221-2222301313021322-0122303111120001-3110013013203003-1221322102203221-2102223202331202-1032233312010131-1220231131103122) |
| `infra.hw_info.bios.date` | [infra.hw_info.bios.date](data-sources--registration--reference--group-001.md#canonical-1130000033021132-3122121203211033-0221110322313100-1022210122203223-3130200301200133-1210032203133132-2333000031032002-1020102113021100) |
| `infra.hw_info.bios.vendor` | [infra.hw_info.bios.vendor](data-sources--registration--reference--group-001.md#canonical-0011300033223200-2123333213102102-0020133003331131-2002331131122103-0212131223331232-1113230321311010-2231112213033320-0113221323131103) |
| `infra.hw_info.bios.version` | [infra.hw_info.bios.version](data-sources--registration--reference--group-001.md#canonical-0303321202210203-2020112300010321-2112213233111123-1202223033213332-2200131113221022-1001130311201322-1111223202123220-2323301322310110) |
| `infra.hw_info.board` | [infra.hw_info.board](data-sources--registration--reference--group-001.md#canonical-2102201002300101-3001013221301201-2322103130022101-1011222030130002-0122133122010330-3302323223121310-2122203113222130-1101223003032031) |
| `infra.hw_info.board.asset_tag` | [infra.hw_info.board.asset_tag](data-sources--registration--reference--group-001.md#canonical-2033220010033222-2022130313200122-1121001032321322-3132311321303211-3320120013120301-0211122221312103-1303211223331323-1112002010002020) |
| `infra.hw_info.board.name` | [infra.hw_info.board.name](data-sources--registration--reference--group-001.md#canonical-2232132032332233-1010121032131031-2011302100030032-1111132120102023-2331303013010313-3321213003002312-1231323100302203-3300322131020111) |
| `infra.hw_info.board.serial` | [infra.hw_info.board.serial](data-sources--registration--reference--group-001.md#canonical-2301321200111303-0032330301110320-2003212331023002-2333031030012021-2303322303300002-3122013001221130-1110203101201300-2021130121223330) |
| `infra.hw_info.board.vendor` | [infra.hw_info.board.vendor](data-sources--registration--reference--group-001.md#canonical-0003000323231210-0223102120310121-1012033121032311-3120011311130222-0022310013220001-3312132310133031-1311032021201211-3010111002121203) |
| `infra.hw_info.board.version` | [infra.hw_info.board.version](data-sources--registration--reference--group-001.md#canonical-2113013101322000-0230223233032133-3011000103011302-1111113033001030-1102023000220131-3023223030202223-1310020002321212-3011111212030121) |
| `infra.hw_info.chassis` | [infra.hw_info.chassis](data-sources--registration--reference--group-001.md#canonical-2331023032221200-1030032231223110-1212210213011132-0333232010233220-3312202300332002-0030231221220120-1030020000131211-2010121012200101) |
| `infra.hw_info.chassis.asset_tag` | [infra.hw_info.chassis.asset_tag](data-sources--registration--reference--group-001.md#canonical-0322023101311333-2030132220110111-1313102302023112-0330222021302003-1222312123020021-3303100230133130-1212313313103211-0033223013220001) |
| `infra.hw_info.chassis.serial` | [infra.hw_info.chassis.serial](data-sources--registration--reference--group-001.md#canonical-2130120112010132-3231112132313332-0111310033033013-2231233113200232-1020320301201310-1001320132010023-1300003120221213-1120120230111220) |
| `infra.hw_info.chassis.type` | [infra.hw_info.chassis.type](data-sources--registration--reference--group-001.md#canonical-1100010232201202-2300303003020021-0123122022220133-0302231213102030-2001322103130021-0011200302131303-3232323102201211-2230300302011002) |
| `infra.hw_info.chassis.vendor` | [infra.hw_info.chassis.vendor](data-sources--registration--reference--group-001.md#canonical-2321301021133103-3021022222013132-1320313111201001-0111020330213310-2210211102002120-0001003110201211-3010303233232100-1121111012301111) |
| `infra.hw_info.chassis.version` | [infra.hw_info.chassis.version](data-sources--registration--reference--group-001.md#canonical-0203202322233113-1311020030003200-1030330111313101-0130122003200300-1032113200233103-2310213130221113-0300133003320123-1300223033013013) |
| `infra.hw_info.cpu` | [infra.hw_info.cpu](data-sources--registration--reference--group-001.md#canonical-2000002120230000-1211331133122112-0303212313032103-3230010032231103-0322222030300333-2202330210320023-0333133311233033-0230303303000013) |
| `infra.hw_info.cpu.cache` | [infra.hw_info.cpu.cache](data-sources--registration--reference--group-001.md#canonical-3330302032120011-1122302332222032-3201021103103000-3013332102011130-0102110103321333-3102201011112012-1210112233202113-0211223031022112) |
| `infra.hw_info.cpu.cores` | [infra.hw_info.cpu.cores](data-sources--registration--reference--group-001.md#canonical-0312023210322130-0332020123102333-1333231113020002-0202311002201123-2313331133213301-3210300311331111-0021333133213133-3103202211130010) |
| `infra.hw_info.cpu.cpus` | [infra.hw_info.cpu.cpus](data-sources--registration--reference--group-001.md#canonical-3333023202233031-1210321023023311-1003222332002111-2113121322212132-0023322023020013-3230102202330221-2111023202110220-0113133000000020) |
| `infra.hw_info.cpu.model` | [infra.hw_info.cpu.model](data-sources--registration--reference--group-001.md#canonical-1021313011100222-1230110201201233-2003200322311101-0122132032332002-3310013322302020-2222312011121320-2030220222023021-0322010121123121) |
| `infra.hw_info.cpu.speed` | [infra.hw_info.cpu.speed](data-sources--registration--reference--group-001.md#canonical-0302011011022313-0302123300130320-2033333032013022-1023030313121233-3322110031222323-1032000002323223-1230322303123103-3300130033011103) |
| `infra.hw_info.cpu.threads` | [infra.hw_info.cpu.threads](data-sources--registration--reference--group-001.md#canonical-0013333111120013-1311201111002112-0101131123033332-0312203011112001-3021232120233112-1010212330130233-1002301200231222-3210200202231232) |
| `infra.hw_info.cpu.vendor` | [infra.hw_info.cpu.vendor](data-sources--registration--reference--group-001.md#canonical-0311301112302113-3230220212331011-0133130000000122-3310110023001131-0003310201331111-0132020302323102-3130001203323110-1032312100023123) |
| `infra.hw_info.gpu` | [infra.hw_info.gpu](data-sources--registration--reference--group-001.md#canonical-3000231302121113-0032101332011012-3212011222003132-1231023232031103-2100001002031011-3012000333010000-1002223003303101-3320320003222130) |
| `infra.hw_info.gpu.cuda_version` | [infra.hw_info.gpu.cuda_version](data-sources--registration--reference--group-001.md#canonical-2100131331201221-3310213303132323-0123221101221212-2030220221033220-0211123213021301-2020121033323332-0003103212003330-2112332231010120) |
| `infra.hw_info.gpu.driver_version` | [infra.hw_info.gpu.driver_version](data-sources--registration--reference--group-001.md#canonical-1222133002030210-2213112100331231-0223010022332322-3021121200020330-0022120021223302-1312033303211321-0012023220331301-0330032103231011) |
| `infra.hw_info.gpu.gpu_device` | [infra.hw_info.gpu.gpu_device](data-sources--registration--reference--group-001.md#canonical-2311301030030330-1112313100231332-1202002022331213-3131232330012122-0131211100112300-2231103202031022-2220001100202032-0300032231201333) |
| `infra.hw_info.gpu.gpu_device.id` | [infra.hw_info.gpu.gpu_device.id](data-sources--registration--reference--group-001.md#canonical-2312032102310203-0020302120123131-0201220110200313-1001322110030103-1000220303311322-2010303132222311-2311121301332222-2022303020120003) |
| `infra.hw_info.gpu.gpu_device.processes` | [infra.hw_info.gpu.gpu_device.processes](data-sources--registration--reference--group-001.md#canonical-3312001203113002-2313230110032233-0020101201203321-1300102033102300-1132322233012221-3303210223330102-0023223300010121-2221100322113101) |
| `infra.hw_info.gpu.gpu_device.product_name` | [infra.hw_info.gpu.gpu_device.product_name](data-sources--registration--reference--group-001.md#canonical-3110000022133133-3110210330101000-2001001231220020-1302132011311332-0313133001122012-1211102101321210-2033020330203012-2121000022123100) |
| `infra.hw_info.kernel` | [infra.hw_info.kernel](data-sources--registration--reference--group-001.md#canonical-0001301233323231-1312032233200133-2001323103011111-0113310202333100-3122230100021122-2232130130112030-0300200001203332-2113230331333313) |
| `infra.hw_info.kernel.architecture` | [infra.hw_info.kernel.architecture](data-sources--registration--reference--group-001.md#canonical-0031132321031322-2010211012122012-3300302312330211-0102033312331332-1201322232201111-1212101031221211-0033333132032001-1002333122131311) |
| `infra.hw_info.kernel.release` | [infra.hw_info.kernel.release](data-sources--registration--reference--group-001.md#canonical-3302322301233301-2113320102213230-1213031031322210-0112301022212213-1211211223020213-1033131233231232-0203230011003013-0132110123300121) |
| `infra.hw_info.kernel.version` | [infra.hw_info.kernel.version](data-sources--registration--reference--group-001.md#canonical-2030221031010333-2311120101322312-1022330311102022-2001112203233322-3021113233322220-0211203021101322-1310132010103203-2312122233021200) |
| `infra.hw_info.memory` | [infra.hw_info.memory](data-sources--registration--reference--group-001.md#canonical-2131020102211122-0111202301322020-0230303022223132-0202313223030123-2100301031112110-2123033231322011-0012000202323233-3302303120000103) |
| `infra.hw_info.memory.size_mb` | [infra.hw_info.memory.size_mb](data-sources--registration--reference--group-001.md#canonical-0203232132310000-3321230020111331-2003000023001013-0023220031213122-2122031302011321-0002332111002023-2100111121122030-1132233313120031) |
| `infra.hw_info.memory.speed` | [infra.hw_info.memory.speed](data-sources--registration--reference--group-001.md#canonical-3130100000011322-2313121021021122-3032212231313231-2303123132011223-3301132103212300-3101301331120132-0120303002201220-1300200003223310) |
| `infra.hw_info.memory.type` | [infra.hw_info.memory.type](data-sources--registration--reference--group-001.md#canonical-2103321332233123-3112001101212132-1201031102011122-2110201032313031-2230101201110312-2000120230301210-1211121030021201-1302001002211010) |
| `infra.hw_info.network` | [infra.hw_info.network](data-sources--registration--reference--group-001.md#canonical-1122203303322210-3003200033200001-2211232023013203-0012013002001312-1102030203212210-0310110220130110-0223131122223001-0332300031111012) |
| `infra.hw_info.network.driver` | [infra.hw_info.network.driver](data-sources--registration--reference--group-001.md#canonical-0111020102000210-1200030323032122-2010023220100322-3212033132022301-0032313102222000-3313130333320130-0010332330301031-2011311300320313) |
| `infra.hw_info.network.ip_address` | [infra.hw_info.network.ip_address](data-sources--registration--reference--group-001.md#canonical-1013221310300111-1111101020202101-3231131203102220-1100312313323110-2303010103333320-1230033113231330-2030323002130000-0032012210202013) |
| `infra.hw_info.network.link_quality` | [infra.hw_info.network.link_quality](data-sources--registration--reference--group-001.md#canonical-2303000022221030-1100022100230000-3223330100203031-0103103100231202-0132300003100231-3000201011201131-0013303111021011-1100123332003303) |
| `infra.hw_info.network.link_type` | [infra.hw_info.network.link_type](data-sources--registration--reference--group-001.md#canonical-0233122033001210-1003231001332113-3322200112022233-2311103033003301-3103313233131332-1120102002300310-0032301333321012-2000213121122300) |
| `infra.hw_info.network.mac_address` | [infra.hw_info.network.mac_address](data-sources--registration--reference--group-001.md#canonical-0111013322001310-1231312210233231-1330213332230312-3112022310001302-2232213233303231-1111022111223000-1113230110123330-1012231022131221) |
| `infra.hw_info.network.name` | [infra.hw_info.network.name](data-sources--registration--reference--group-001.md#canonical-2221230021002030-3232200203333331-1113312312303210-1310033121012111-2301210232132103-3320232222332013-3022312232010330-0231101111122210) |
| `infra.hw_info.network.port` | [infra.hw_info.network.port](data-sources--registration--reference--group-001.md#canonical-2033223002311210-0131013211231032-3112011202212213-3021312333232221-2031031112001213-3321002113313032-0012200210202131-2322311011002000) |
| `infra.hw_info.network.speed` | [infra.hw_info.network.speed](data-sources--registration--reference--group-001.md#canonical-1331210331102020-3010330212223111-0300310230221120-1231112102332331-1120021311112221-2300223233121310-3031222212301001-2020123121332003) |
| `infra.hw_info.numa_nodes` | [infra.hw_info.numa_nodes](data-sources--registration--reference--group-001.md#canonical-0201002213133012-3223031321122301-3113233102202202-2123121103112211-2230331000121103-1313112233103110-2130200230331200-3302031103301332) |
| `infra.hw_info.os` | [infra.hw_info.os](data-sources--registration--reference--group-001.md#canonical-3102202100220232-0003312222031102-0222301331131100-0203111223330221-1302300200010220-3012220230331122-3023310132022002-1232021223312231) |
| `infra.hw_info.os.architecture` | [infra.hw_info.os.architecture](data-sources--registration--reference--group-001.md#canonical-1323132130310220-1333301312301321-0113321010011023-0212032023013211-0111122322110130-2013131031220330-0221131100321211-0130212002233032) |
| `infra.hw_info.os.name` | [infra.hw_info.os.name](data-sources--registration--reference--group-001.md#canonical-2231001033121103-1112130223003323-1033310302313132-1133012200023010-3333233120321032-0112013123123233-0110102013130202-2202223102330022) |
| `infra.hw_info.os.release` | [infra.hw_info.os.release](data-sources--registration--reference--group-001.md#canonical-3212321221021231-0302003021112231-2222023020033312-2131200021103003-3110212302123003-3121123231223113-3032200222100231-2320300132101212) |
| `infra.hw_info.os.vendor` | [infra.hw_info.os.vendor](data-sources--registration--reference--group-001.md#canonical-2010013012102102-2133311121122030-2120101100110031-0320030332203210-1300222002312030-1022101131211003-0112232111133333-1110103333103232) |
| `infra.hw_info.os.version` | [infra.hw_info.os.version](data-sources--registration--reference--group-001.md#canonical-3323221331112233-0032332133133112-0102232021133321-3123202101232031-0230321000123333-2132231330100231-3132032321103302-2122112302000131) |
| `infra.hw_info.product` | [infra.hw_info.product](data-sources--registration--reference--group-001.md#canonical-0001331130320000-0001012311100102-0213223112303000-2310131213331003-1013302323333012-0221322113013301-0310100031211230-3021021011121201) |
| `infra.hw_info.product.name` | [infra.hw_info.product.name](data-sources--registration--reference--group-001.md#canonical-3200010310022222-3112001121132032-3022203303233033-2133302323132303-2302333020110002-2033330332003221-1330233232012311-2032332313133023) |
| `infra.hw_info.product.serial` | [infra.hw_info.product.serial](data-sources--registration--reference--group-001.md#canonical-3033132211303101-3333000121220203-2020002123231220-0113210021202332-0120122303333110-2113332111113202-0311020122102301-3113330303332111) |
| `infra.hw_info.product.vendor` | [infra.hw_info.product.vendor](data-sources--registration--reference--group-001.md#canonical-3310210330200203-2312223102220321-0002000121132010-2020303233220222-0332120221121020-1103211002311121-2323103113031320-0122200033203220) |
| `infra.hw_info.product.version` | [infra.hw_info.product.version](data-sources--registration--reference--group-001.md#canonical-3302301223201113-2323122301000101-3223330210331310-0220212011111223-1331113112301113-2311220120301301-2221321021211002-3223032303022232) |
| `infra.hw_info.storage` | [infra.hw_info.storage](data-sources--registration--reference--group-001.md#canonical-0103223310021211-2222012313000102-2020002002330120-2100132200013013-3213130112301030-0000202301012200-0323232130013313-1021311202023311) |
| `infra.hw_info.storage.driver` | [infra.hw_info.storage.driver](data-sources--registration--reference--group-001.md#canonical-2123323113032130-3000233222010311-2013100123112102-1213203213131232-1020000212321122-2321033321010122-0000312100330323-0300030213100222) |
| `infra.hw_info.storage.model` | [infra.hw_info.storage.model](data-sources--registration--reference--group-001.md#canonical-2332212101202023-0232213202313113-0312211100230210-2002101111103222-3301031022223113-1331131233130013-3310130121231021-1021112232321231) |
| `infra.hw_info.storage.name` | [infra.hw_info.storage.name](data-sources--registration--reference--group-001.md#canonical-1311111000233323-2001323232211302-2132023022330302-0000211200331000-1020100301000300-0223233023012023-0002313322300003-2100002003221102) |
| `infra.hw_info.storage.serial` | [infra.hw_info.storage.serial](data-sources--registration--reference--group-001.md#canonical-0101221030210002-0331333110302233-2231331000303221-1211320201232230-2001321023000231-2201000331003213-1112100000020002-3100232303321002) |
| `infra.hw_info.storage.size_gb` | [infra.hw_info.storage.size_gb](data-sources--registration--reference--group-001.md#canonical-0323213230032103-3202331320000122-2223323201223021-0110233032233032-2023132130202322-2003213123123222-0300220131101023-1301313320233102) |
| `infra.hw_info.storage.vendor` | [infra.hw_info.storage.vendor](data-sources--registration--reference--group-001.md#canonical-2332221120231233-0223203123313102-2023031101023201-1022322000031122-1031203033001133-0301331103300221-1102232213220011-2230000031210200) |
| `infra.hw_info.usb` | [infra.hw_info.usb](data-sources--registration--reference--group-001.md#canonical-1201303230230100-3020203212333213-2012333220013110-3120133001201011-2033200200311202-1320111312202003-2001322002121130-1120201330033302) |
| `infra.hw_info.usb.address` | [infra.hw_info.usb.address](data-sources--registration--reference--group-001.md#canonical-1133313112212113-2302230032223012-1030000020320222-1223221012013331-1232133330210201-2123230300201103-0322131033320123-2001312201011301) |
| `infra.hw_info.usb.b_device_class` | [infra.hw_info.usb.b_device_class](data-sources--registration--reference--group-001.md#canonical-2112222202111132-1010102031233011-2201233020113013-0303230123203020-2320023231010120-1320120301203120-3331012233012301-2103112322211003) |
| `infra.hw_info.usb.b_device_protocol` | [infra.hw_info.usb.b_device_protocol](data-sources--registration--reference--group-001.md#canonical-3021010322002102-0120222212203333-0102031203222230-2303031201311010-3220020232300112-0310230320103100-1311013333001110-0333033212013323) |
| `infra.hw_info.usb.b_device_sub_class` | [infra.hw_info.usb.b_device_sub_class](data-sources--registration--reference--group-001.md#canonical-2331301221101100-3210122302123023-0221310313110330-0122122022332230-1222003230310222-2122000232120122-2130121300013102-3033203110000312) |
| `infra.hw_info.usb.b_max_packet_size` | [infra.hw_info.usb.b_max_packet_size](data-sources--registration--reference--group-001.md#canonical-0300120301012312-0103032301232131-3010211030222120-0320323110221012-0232303131222310-3021131331202230-1120030122302311-0232300110011210) |
| `infra.hw_info.usb.bcd_device` | [infra.hw_info.usb.bcd_device](data-sources--registration--reference--group-001.md#canonical-1320210121132202-0003220233003230-3213231221012110-0302333212230130-1112100232211303-3233101132011033-0132122303212130-1223231210330233) |
| `infra.hw_info.usb.bcd_usb` | [infra.hw_info.usb.bcd_usb](data-sources--registration--reference--group-001.md#canonical-3232012323102022-0310231110311203-3021302010000211-0331133110030031-3131323013330232-0130110323333012-0213022311003312-1330223231231123) |
| `infra.hw_info.usb.bus` | [infra.hw_info.usb.bus](data-sources--registration--reference--group-001.md#canonical-1210312230313233-1212231221201233-2211200122121323-0310130212330210-0330023231323123-1220120201000210-0333231110322230-2201331111300023) |
| `infra.hw_info.usb.description_spec` | [infra.hw_info.usb.description_spec](data-sources--registration--reference--group-001.md#canonical-1201210323233320-3110023201010110-1332113222300223-0320200130312103-1110113231201201-1310110110310223-1303320331031031-1023113322322212) |
| `infra.hw_info.usb.i_manufacturer` | [infra.hw_info.usb.i_manufacturer](data-sources--registration--reference--group-001.md#canonical-0032211010210132-1102202013031302-0302001002333002-3220300102330003-2103121100022330-2330222030111301-3220010103111302-1310321122213310) |
| `infra.hw_info.usb.i_product` | [infra.hw_info.usb.i_product](data-sources--registration--reference--group-001.md#canonical-3112330102200311-0011231013133311-1230032031130201-3020031130301311-3230133310131230-0232120132010320-3011312210121233-2112132310113121) |
| `infra.hw_info.usb.i_serial` | [infra.hw_info.usb.i_serial](data-sources--registration--reference--group-001.md#canonical-1133232132022012-1112222321211002-0220031232201312-1200203000110132-0120302330331300-0013200120232332-3011211210321121-0021201101030301) |
| `infra.hw_info.usb.id_product` | [infra.hw_info.usb.id_product](data-sources--registration--reference--group-001.md#canonical-1033311210133301-2300122020031131-3133321121013130-0032300232220333-0321233131300123-3302312002132230-3000101332221132-1320302220132230) |
| `infra.hw_info.usb.id_vendor` | [infra.hw_info.usb.id_vendor](data-sources--registration--reference--group-001.md#canonical-3312121101331031-0111103212003010-3111003311102220-2023011332232313-0202101130020103-1211301210113331-0330222203330001-1333130000333332) |
| `infra.hw_info.usb.port` | [infra.hw_info.usb.port](data-sources--registration--reference--group-001.md#canonical-2313131332313202-2121003310023013-1020010001310013-1232302121313013-2211021022300201-1112312311010221-1233111131323300-3123023133030321) |
| `infra.hw_info.usb.product_name` | [infra.hw_info.usb.product_name](data-sources--registration--reference--group-001.md#canonical-0021012131233021-0120311032222120-0311021221132320-0121001013210002-1322112133310033-0303102120330333-2230222221333132-2230332322311222) |
| `infra.hw_info.usb.speed` | [infra.hw_info.usb.speed](data-sources--registration--reference--group-001.md#canonical-2201101133211331-1113331113103222-2120131231203230-1233202121333301-3330211203002013-3211223120002211-0032302330112311-0310222223231010) |
| `infra.hw_info.usb.usb_type` | [infra.hw_info.usb.usb_type](data-sources--registration--reference--group-001.md#canonical-0011032100222322-3120102000200032-1322320230200221-3312213231311203-0002303201120011-3032033001130332-0031032312301121-2131223200033220) |
| `infra.hw_info.usb.vendor_name` | [infra.hw_info.usb.vendor_name](data-sources--registration--reference--group-001.md#canonical-3333202113000000-0212023111213230-2032320020102013-1303223010300202-0020121111100132-2200133030112131-2102311313020212-1122320030312233) |
| `infra.instance_id` | [infra.instance_id](data-sources--registration--reference--group-001.md#canonical-0031202120112020-2131211011003232-3120001130213310-0220033132001122-3222222031002201-2102332011210302-1112101210222002-3113111300332012) |
| `infra.interfaces` | [infra.interfaces](data-sources--registration--reference--group-001.md#canonical-3030011312300101-3100323121201120-0320212110131033-0002010300332013-0231311323320301-2033001120222211-1210120200130311-2311102230111233) |
| `infra.internet_proxy` | [infra.internet_proxy](data-sources--registration--reference--group-001.md#canonical-1300102132123302-3113312100320313-2203310020010231-1300223120023233-0210202010033122-0031020220203031-0211311300223101-2101321201313212) |
| `infra.internet_proxy.http_proxy` | [infra.internet_proxy.http_proxy](data-sources--registration--reference--group-001.md#canonical-0310300122132010-1113110201212023-2102113202332320-3202230312301233-3030002130201023-1333100013231032-2220003003302331-3133331330210201) |
| `infra.internet_proxy.https_proxy` | [infra.internet_proxy.https_proxy](data-sources--registration--reference--group-001.md#canonical-0102112121222010-3333231031233010-3302312202210012-2011301220222033-2102122021120101-0323021133103213-2003302320320212-0203320101210212) |
| `infra.internet_proxy.no_proxy` | [infra.internet_proxy.no_proxy](data-sources--registration--reference--group-001.md#canonical-0020133012101103-0020133102122002-0131232120200211-2203111231222003-1203012310321001-3031231230111301-0123230100331110-2333220303313320) |
| `infra.internet_proxy.proxy_cacert_url` | [infra.internet_proxy.proxy_cacert_url](data-sources--registration--reference--group-001.md#canonical-0031332233000331-0121112331011323-3003213012030120-2202023101310113-0331113003213121-2032002332102100-1223231332020133-0130331300230032) |
| `infra.is_slo_static` | [infra.is_slo_static](data-sources--registration--reference--group-001.md#canonical-0023101021103100-3230131103221313-2002312112223212-2131133020103131-0323203003031132-3130233323303033-0201301010331011-2221311322031201) |
| `infra.machine_id` | [infra.machine_id](data-sources--registration--reference--group-001.md#canonical-0333302002211122-1200011232222021-2203330313020333-3230010221102330-2302112030130213-0011301331103221-1020113231101032-0301013020233012) |
| `infra.provider_ref` | [infra.provider_ref](data-sources--registration--reference--group-001.md#canonical-3100311010302303-2132131103002332-1111232220121202-0102213002132332-0132222102111230-3202303310233001-0133013210020213-2111020013110321) |
| `infra.sw_info` | [infra.sw_info](data-sources--registration--reference--group-001.md#canonical-0003013103231100-0313031200312102-0230111123113122-3302100301133112-3213230120020223-1020130130121031-3320120230201031-3102111033211231) |
| `infra.sw_info.sw_version` | [infra.sw_info.sw_version](data-sources--registration--reference--group-001.md#canonical-3033303202000302-0012331232200002-3201323212132111-0210303122301012-1213232112222203-1320102121203111-2030321122102232-0130220021322110) |
| `infra.timestamp` | [infra.timestamp](data-sources--registration--reference--group-001.md#canonical-0010311311232231-0321112232113210-2011131122031331-1122321320233232-1211322011013211-1233113133131213-1112011222011311-0020011322201312) |
| `infra.zone` | [infra.zone](data-sources--registration--reference--group-001.md#canonical-0012122333032333-1203322212130000-0133230322311213-3030331131023203-2101122301200100-1302022021030100-1030110023221321-3323323012320033) |
| `labels` | [labels](data-sources--registration--reference--group-001.md#canonical-2223020132121332-0132033301203311-3220113320101032-2022211310232003-1321321300322303-0110203310001123-0010210020211211-3301013020201323) |
| `name` | [name](data-sources--registration--reference--group-001.md#canonical-2012322121012313-0001100230323323-0213212101313230-1020233222013031-3312123031213000-3233122111322013-2221210033113332-2323003132101333) |
| `namespace` | [namespace](data-sources--registration--reference--group-001.md#canonical-0010110123221301-2000302210223212-1012132031123232-2111100032322223-0011302132131022-1133202303011020-2013230201131201-3031203030232202) |
| `passport` | [passport](data-sources--registration--reference--group-001.md#canonical-2211123122231133-1322031103221030-2232331132131310-2221202133030102-0110220101320300-2132313310232222-1312002231012220-1131111000312210) |
| `passport.cluster_name` | [passport.cluster_name](data-sources--registration--reference--group-001.md#canonical-1100313202130030-3122012022313032-2122032023210012-2323301100310132-3002201020120303-0212131220223123-0311233231132303-3100211203013323) |
| `passport.cluster_size` | [passport.cluster_size](data-sources--registration--reference--group-001.md#canonical-0010022003123222-3332210222213033-0231002022130032-0233122103212123-1223012102111112-3331233201000322-2323123301203210-3010222133303210) |
| `passport.cluster_type` | [passport.cluster_type](data-sources--registration--reference--group-001.md#canonical-1230003210213110-2133201130300233-0003001020232330-0011102302322331-1002023101330131-3332333322012211-3311311333030010-0202201203232321) |
| `passport.default_os_version` | [passport.default_os_version](data-sources--registration--reference--group-001.md#canonical-2021231302100201-1122030102210232-3210202120333102-0310321031003322-3111213231023303-2003030100230232-2312011303330010-3011021220330102) |
| `passport.default_sw_version` | [passport.default_sw_version](data-sources--registration--reference--group-001.md#canonical-2003033332110300-2200301110012012-0323311023210233-1122212221203231-2120203131102320-0023311030021123-0231001210120030-3130332030012113) |
| `passport.latitude` | [passport.latitude](data-sources--registration--reference--group-001.md#canonical-1103032331313131-2002111002122112-0330300023210313-3211111221202223-3001131311030001-0002000123030212-2211021323130111-1220003030220301) |
| `passport.longitude` | [passport.longitude](data-sources--registration--reference--group-001.md#canonical-3110010212312331-3110202322013211-1103230120223222-2131111330331031-2232202311310123-0000012022001000-0021011110013332-3332230000003211) |
| `passport.operating_system_version` | [passport.operating_system_version](data-sources--registration--reference--group-001.md#canonical-2213322102331012-0032303330202211-2130211321203012-0203033003330303-0222302001113133-0103012102030322-0332033301020211-2323220123330012) |
| `passport.private_network_name` | [passport.private_network_name](data-sources--registration--reference--group-001.md#canonical-3211022121311122-1313103020323210-0331000313200113-2212213311223320-1231003332133312-0122121010312210-2232110130231012-1321213031202300) |
| `passport.volterra_software_version` | [passport.volterra_software_version](data-sources--registration--reference--group-001.md#canonical-1111303001322301-0010230002312123-3313123222111100-1302100123302232-3210021231200232-3021233032133200-0321121220002210-0032210212220221) |
| `token` | [token](data-sources--registration--reference--group-001.md#canonical-3302320203230301-1332012020212013-2202001212322020-3222010220322101-2000023331032012-1200221113022103-1331201111313102-2302013010213311) |

<a id="canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- infra

<a id="canonical-3112222200000320-1111111333003022-2331222031103212-2220131210213222-3101031232133121-2112022212001311-1230331211120220-3023021133322232"></a>

Type: `"single"`. Computed.

InfraMetadata stores information about instance infrastructure.

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

<a id="canonical-3131201010333123-1131011002103323-2312003031200111-2102312312032100-0132113131303100-2202032210302022-3311332003300022-2230131300000000"></a>

### Direct properties for `infra`

<a id="canonical-3102030321312222-3312110223132133-2020331000033003-2112030310212002-2023313302102113-3320210220212221-3220220120023103-3301321221022233"></a>

#### `infra.availability_zone` property

Type: `"string"`. Computed.

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

- [bond_config](data-sources--registration--reference--group-001.md#canonical-2210100300012100-2211112222210322-1103231110102010-2110123130310323-3310311300212203-3120001033022233-3100221203232103-1203021310230010): complete subsection reference.

<a id="canonical-0132323102031312-1200022013122233-2300310301122131-1000303212113130-1210131112223133-2113230121002223-3320323301132320-1120111222011020"></a>

<a id="canonical-2230022001312223-2223220232230320-2312121323132323-0202113230331301-0021132203000011-0222233212211033-1312313023212001-1210000110100221"></a>

#### `infra.certified_hw` property

Type: `"string"`. Computed.

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

<a id="canonical-2112113223301323-2121211013330220-2103000112113223-3301331200100032-2131010000223023-1120230010300020-1331303011223020-3131110022232231"></a>

<a id="canonical-1130231301021313-3310110002002133-2013221313320131-3303010223203321-0311003013201021-1110210210313013-0230000331322122-1133200333313022"></a>

#### `infra.domain` property

Type: `"string"`. Computed.

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

<a id="canonical-0211023101310320-3213322120311301-2030121013032300-0213110002313102-1021222321110002-2110322302201200-0033100220213100-3313010200012120"></a>

<a id="canonical-2011232032301001-2122121203302233-2211331331100102-1021012313203212-3031203100222313-1122002301120232-3232212102320002-0022233330110213"></a>

#### `infra.hostname` property

Type: `"string"`. Computed.

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

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

- [hugepages](data-sources--registration--reference--group-001.md#canonical-2120031103113301-3131220333132032-3213120323123113-0301303301112301-1202231231220101-2130323303003113-0133003202112021-0123331322323300): complete subsection reference.

- [hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110): complete subsection reference.

<a id="canonical-0031202120112020-2131211011003232-3120001130213310-0220033132001122-3222222031002201-2102332011210302-1112101210222002-3113111300332012"></a>

<a id="canonical-1300231011333100-2323011133113100-3110022011122220-2223010020010330-1113110023130111-1302201312310101-1112002323111230-2113032100001100"></a>

#### `infra.instance_id` property

Type: `"string"`. Computed.

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

- [interfaces](data-sources--registration--reference--group-001.md#canonical-1212321321230203-1333031332333111-2013211110013002-2012312130010010-3132000213031022-1302312311201033-1112302223303330-1312233300011331): complete subsection reference.

- [internet_proxy](data-sources--registration--reference--group-001.md#canonical-2022113132332131-3301201322311101-1303311110321310-3103112022020203-2132031321200102-2100333112033011-0222212203001113-2210131301030123): complete subsection reference.

<a id="canonical-0023101021103100-3230131103221313-2002312112223212-2131133020103131-0323203003031132-3130233323303033-0201301010331011-2221311322031201"></a>

<a id="canonical-1002330320201323-2333300131321111-0330233331323200-3332012220110012-3010310030313102-1031020111203031-1323103230321002-3303110300213203"></a>

#### `infra.is_slo_static` property

Type: `"bool"`. Computed.

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

<a id="canonical-0333302002211122-1200011232222021-2203330313020333-3230010221102330-2302112030130213-0011301331103221-1020113231101032-0301013020233012"></a>

<a id="canonical-2321122002210223-1002111313010000-1133010011001110-0312233020223002-3320013201200301-2100330323113102-3120021131000301-0311111131030313"></a>

#### `infra.machine_id` property

Type: `"string"`. Computed.

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

<a id="canonical-3100311010302303-2132131103002332-1111232220121202-0102213002132332-0132222102111230-3202303310233001-0133013210020213-2111020013110321"></a>

<a id="canonical-3112302310020311-2323230201331231-3211130031011100-1130000211230311-2310331133033331-1313003003300232-0333003303123202-1313101200002100"></a>

#### `infra.provider_ref` property

Type: `"string"`. Computed.

\[Enum:
UNKNOWN|AWS|GOOGLE|AZURE|VMWARE|KVM|OTHER|VOLTERRA|IBMCLOUD|UNKNOWN\_K8S|AWS\_K8S|GCP\_K8S|AZURE\_K8S|VMWARE\_K8S|KVM\_K8S|OTHER\_K8S|VOLTERRA\_K8S|IBMCLOUD\_K8S|F5OS|RSERIES|OCI|NUTANIX|OPENSTACK|EQUINIX|OPENSHIFT\_VIRTUALIZATION|KUBERNETES\]
Infrastructure provider enum for registration. It describes where is instance running. Provider was
not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other
provider, which was not identified by system. Possible values are \`UNKNOWN\`, \`AWS\`, \`GOOGLE\`,
\`AZURE\`, \`VMWARE\`, \`KVM\`, \`OTHER\`, \`VOLTERRA\`, \`IBMCLOUD\`, \`UNKNOWN\_K8S\`,
\`AWS\_K8S\`, \`GCP\_K8S\`, \`AZURE\_K8S\`, \`VMWARE\_K8S\`, \`KVM\_K8S\`, \`OTHER\_K8S\`,
\`VOLTERRA\_K8S\`, \`IBMCLOUD\_K8S\`, \`F5OS\`, \`RSERIES\`, \`OCI\`, \`NUTANIX\`, \`OPENSTACK\`,
\`EQUINIX\`, \`OPENSHIFT\_VIRTUALIZATION\`, \`KUBERNETES\`.

- [sw_info](data-sources--registration--reference--group-001.md#canonical-1231311000020020-2022312100213123-2023300303302201-1200001213120331-1211131030203221-1032233033010103-3210023002233002-0031113231233303): complete subsection reference.

<a id="canonical-0010311311232231-0321112232113210-2011131122031331-1122321320233232-1211322011013211-1233113133131213-1112011222011311-0020011322201312"></a>

<a id="canonical-0220003103101113-0201021221313111-0130111312103032-2000101011303111-2212030200302033-1323232123130321-1210130232030323-0323200102102300"></a>

#### `infra.timestamp` property

Type: `"string"`. Computed.

It's used to verify machine have acceptable time difference from server.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0012122333032333-1203322212130000-0133230322311213-3030331131023203-2101122301200100-1302022021030100-1030110023221321-3323323012320033"></a>

<a id="canonical-0021122133230332-1312010203210013-3320030211200220-2101130320312213-1020231130232002-2330330100232122-3001020322232211-2123203133201022"></a>

#### `infra.zone` property

Type: `"string"`. Computed.

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

<a id="canonical-2210100300012100-2211112222210322-1103231110102010-2110123130310323-3310311300212203-3120001033022233-3100221203232103-1203021310230010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.bond_config` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- infra.bond_config

<a id="canonical-2011122311112200-2230130331013302-2331320010101333-3211031100313120-1201221311120131-1113310300132331-1223031310012130-3321133013302313"></a>

Type: `"single"`. Computed.

Bond device configuration for VPM registration.

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

<a id="canonical-2120002003010013-3300221013023312-3010303120310000-1311320223021010-2023321023120000-0010331220221312-1321202211100122-0201300322202331"></a>

### Direct properties for `infra.bond_config`

<a id="canonical-3313102011123030-0112131103101222-2031220202121031-0100202302213102-3110001011230000-0201002002123022-1300213200031312-1211113233120102"></a>

#### `infra.bond_config.interfaces` property

Type: `["list", "string"]`. Computed.

Member Interfaces. Configuration parameter for interfaces

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

<a id="canonical-3201231130232022-2200003231111122-3331100300212310-0120320333300321-2012100301220313-3130202222022333-1123221003222120-1112132330130023"></a>

<a id="canonical-3233013020222330-2020130101322200-1300313132110110-1110111021033330-1303332103313131-3031113022233200-0313212021032222-2300032103302030"></a>

#### `infra.bond_config.mode` property

Type: `"string"`. Computed.

\[Enum: BOND\_MODE\_UNSPECIFIED|ACTIVE\_BACKUP|LACP\_802\_3AD\] Bonding mode for bond device
configuration Bond mode is not specified Active-backup bond mode (one interface active, others as
backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are
\`BOND\_MODE\_UNSPECIFIED\`, \`ACTIVE\_BACKUP\`, \`LACP\_802\_3AD\`. Defaults to
\`BOND\_MODE\_UNSPECIFIED\`.

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

<a id="canonical-2321122021132300-2233120222123012-1100122033103023-1220003311123203-1010313123311312-3121003033122001-1232210031221103-2303103202002313"></a>

<a id="canonical-0213003001312333-1033212323312000-2300121312303232-3331220020112133-3232323203210011-1001312201012103-3300320233121312-1020131023303310"></a>

#### `infra.bond_config.name` property

Type: `"string"`. Computed.

Bond Name. Human-readable name for the resource

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2120031103113301-3131220333132032-3213120323123113-0301303301112301-1202231231220101-2130323303003113-0133003202112021-0123331322323300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hugepages` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- infra.hugepages

<a id="canonical-3023201330231122-0300313230211012-3112332121023123-1310002202030323-0000032021323213-0330010101020112-1003131101330320-1113012100330102"></a>

Type: `"list"`. Computed.

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

<a id="canonical-3321221223231302-1031312101113120-0021231231231221-2333131200221012-3322210000202011-0232220203201133-1202001031220033-1121003311331200"></a>

### Direct properties for `infra.hugepages`

<a id="canonical-1101222220102313-3010121031020322-3313023130201123-0302120132210130-3101020100010212-3022013011122203-3210031211112300-2030201312202332"></a>

#### `infra.hugepages.free` property

Type: `"number"`. Computed.

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

<a id="canonical-0011131130011030-3113221332212113-2330013200123032-2020331030113331-3322330331300033-0010113230331332-1112211333310200-3101201322322033"></a>

<a id="canonical-3210120011001032-2023032331302101-3132030120131112-1123320132100113-0320113232130002-1323003013101321-2222311011230310-3120203112031201"></a>

#### `infra.hugepages.page_size` property

Type: `"number"`. Computed.

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

<a id="canonical-3330231001322203-2201132303302323-2302222003001223-0203201213201001-2103203201333012-0332112011223101-3202111102031130-0111002202332011"></a>

<a id="canonical-2113100330020120-2002011300233231-3130322030323311-3122313312213200-1323301223201000-3302131112120321-0130332013001220-1232331320230213"></a>

#### `infra.hugepages.total` property

Type: `"number"`. Computed.

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

<a id="canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- infra.hw_info

<a id="canonical-2302313323300202-2222112222012030-1123102301303202-1111323301031310-1100001230213121-1111202222323020-3321203310211110-2121202001012133"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3131030202101001-0311022223233021-2120111331330130-1310220003231203-1023011120133231-1300022101113120-0202323330130320-3303123213312233"></a>

### Direct properties for `infra.hw_info`

- [bios](data-sources--registration--reference--group-001.md#canonical-0000102130023132-1202301130022012-2132301332302323-0303131002201310-1123230100220122-2331312121301322-0003310221202221-2033231331300212): complete subsection reference.

- [board](data-sources--registration--reference--group-001.md#canonical-0300033332101323-2231303002032101-2312221131233033-2200021113223020-2310013023012212-3020301112132030-3130303021312113-2013200313132132): complete subsection reference.

- [chassis](data-sources--registration--reference--group-001.md#canonical-1001312020030130-3132122321112032-3230333320202221-2132323210011311-1223030323003231-0002333330132110-3102121021030212-0321100020112332): complete subsection reference.

- [CPU](data-sources--registration--reference--group-001.md#canonical-2310202130323300-2033301303131032-0201000311132033-1222111330021332-0312020321111232-1013202023021311-2030212020012223-3133012100322023): complete subsection reference.

- [GPU](data-sources--registration--reference--group-001.md#canonical-0120231302203333-1211031200013033-3000023223131131-3003131202220332-1101110002222320-2020023130030103-2213213010201021-2211113320220310): complete subsection reference.

- [kernel](data-sources--registration--reference--group-001.md#canonical-3221100101322111-0322230203320111-2222313030300303-2111330331331222-3133232133112212-1023103301201132-2202120013320222-3133222303011322): complete subsection reference.

- [memory](data-sources--registration--reference--group-001.md#canonical-0323000330332202-1230030330120123-2310033331331110-1303220032211210-0112023331101110-1201312233212003-3113001221310000-0022032120333011): complete subsection reference.

- [network](data-sources--registration--reference--group-001.md#canonical-2002003011233221-2101002023203311-0113332333132302-0203012031033203-1302212230320112-2233130100003303-0003203001333131-1022321123001211): complete subsection reference.

<a id="canonical-0201002213133012-3223031321122301-3113233102202202-2123121103112211-2230331000121103-1313112233103110-2130200230331200-3302031103301332"></a>

<a id="canonical-3010102210300331-3213013021213311-2120333120131212-0321123231233023-3000012133223231-3211030131103322-2111220320110113-2303013302201130"></a>

#### `infra.hw_info.numa_nodes` property

Type: `"number"`. Computed.

Non-uniform memory access (NUMA) nodes count.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [os](data-sources--registration--reference--group-001.md#canonical-2023302323122322-3003210101112121-2100312221020002-2332203332311233-0010331003002231-0120201232012120-2332013030003023-0101232230013331): complete subsection reference.

- [product](data-sources--registration--reference--group-001.md#canonical-1130310222113311-3300132232301211-2013110222110020-2030303131202120-3223322133310021-0322013302100103-2312221203132313-3000332022101223): complete subsection reference.

- [storage](data-sources--registration--reference--group-001.md#canonical-3123013201111023-2110203003311322-1023313313031222-1302230023123001-3122301203113321-1220031100211023-0012313202001122-3313013022312331): complete subsection reference.

- [usb](data-sources--registration--reference--group-001.md#canonical-1110023331232012-0110323002300313-0131123232121130-0310333130302000-0321212033001212-2330232222300223-2313213031323201-0101213101122321): complete subsection reference.

<a id="canonical-0000102130023132-1202301130022012-2132301332302323-0303131002201310-1123230100220122-2331312121301322-0003310221202221-2033231331300212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.bios` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.bios

<a id="canonical-0211320230322221-2222301313021322-0122303111120001-3110013013203003-1221322102203221-2102223202331202-1032233312010131-1220231131103122"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0030023013103301-3131133011330211-1012210103133100-3322102321221100-2033133120022303-3223300203011121-3220110110321310-2322233013232133"></a>

### Direct properties for `infra.hw_info.bios`

<a id="canonical-1130000033021132-3122121203211033-0221110322313100-1022210122203223-3130200301200133-1210032203133132-2333000031032002-1020102113021100"></a>

#### `infra.hw_info.bios.date` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/bios\_date.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0011300033223200-2123333213102102-0020133003331131-2002331131122103-0212131223331232-1113230321311010-2231112213033320-0113221323131103"></a>

<a id="canonical-1323231010232332-3130113122301012-0201002002322132-0223302110320222-3031011220021301-0101301233110002-3302312122312332-2103131303032301"></a>

#### `infra.hw_info.bios.vendor` property

Type: `"string"`. Computed.

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

<a id="canonical-0303321202210203-2020112300010321-2112213233111123-1202223033213332-2200131113221022-1001130311201322-1111223202123220-2323301322310110"></a>

<a id="canonical-2103003211321110-3102223213303331-3103112321112211-3102322300111202-2310123311303311-1220031110333311-2000012303032002-0123222011012330"></a>

#### `infra.hw_info.bios.version` property

Type: `"string"`. Computed.

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

<a id="canonical-0300033332101323-2231303002032101-2312221131233033-2200021113223020-2310013023012212-3020301112132030-3130303021312113-2013200313132132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.board` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.board

<a id="canonical-2102201002300101-3001013221301201-2322103130022101-1011222030130002-0122133122010330-3302323223121310-2122203113222130-1101223003032031"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0220203232020001-2212011130011311-0003232203323322-3212022302312030-1120033203231030-1123032322003200-2202002201000230-1030120223111130"></a>

### Direct properties for `infra.hw_info.board`

<a id="canonical-2033220010033222-2022130313200122-1121001032321322-3132311321303211-3320120013120301-0211122221312103-1303211223331323-1112002010002020"></a>

#### `infra.hw_info.board.asset_tag` property

Type: `"string"`. Computed.

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

<a id="canonical-2232132032332233-1010121032131031-2011302100030032-1111132120102023-2331303013010313-3321213003002312-1231323100302203-3300322131020111"></a>

<a id="canonical-0121000121033011-1003302032201033-1230121311331331-2113003133330301-3310132023120203-0223030211222231-1132323312211312-2032020332220210"></a>

#### `infra.hw_info.board.name` property

Type: `"string"`. Computed.

Information from /sys/class/dmi/ID/board\_name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301321200111303-0032330301110320-2003212331023002-2333031030012021-2303322303300002-3122013001221130-1110203101201300-2021130121223330"></a>

<a id="canonical-0310230121132232-3022230202012322-3023133032012233-2002122310213303-2303133101120201-3230312312212200-3332313333311303-1112322021302331"></a>

#### `infra.hw_info.board.serial` property

Type: `"string"`. Computed.

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

<a id="canonical-0003000323231210-0223102120310121-1012033121032311-3120011311130222-0022310013220001-3312132310133031-1311032021201211-3010111002121203"></a>

<a id="canonical-1300331100311330-2213110123210013-1110320321013112-1233032102203021-2300122202221330-0323100103202332-1311021111230332-1213312311231211"></a>

#### `infra.hw_info.board.vendor` property

Type: `"string"`. Computed.

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

<a id="canonical-2113013101322000-0230223233032133-3011000103011302-1111113033001030-1102023000220131-3023223030202223-1310020002321212-3011111212030121"></a>

<a id="canonical-0022030313030123-2311221021132301-3230130312230111-0130231030112031-3002213223221221-1003210312233323-2032220310013010-2221211223100013"></a>

#### `infra.hw_info.board.version` property

Type: `"string"`. Computed.

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

<a id="canonical-1001312020030130-3132122321112032-3230333320202221-2132323210011311-1223030323003231-0002333330132110-3102121021030212-0321100020112332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.chassis` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.chassis

<a id="canonical-2331023032221200-1030032231223110-1212210213011132-0333232010233220-3312202300332002-0030231221220120-1030020000131211-2010121012200101"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3310312032222031-3310200021321202-0320213110223031-1121213120103022-3122213122130033-3100100203020230-1103213033312333-2012013223303133"></a>

### Direct properties for `infra.hw_info.chassis`

<a id="canonical-0322023101311333-2030132220110111-1313102302023112-0330222021302003-1222312123020021-3303100230133130-1212313313103211-0033223013220001"></a>

#### `infra.hw_info.chassis.asset_tag` property

Type: `"string"`. Computed.

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

<a id="canonical-2130120112010132-3231112132313332-0111310033033013-2231233113200232-1020320301201310-1001320132010023-1300003120221213-1120120230111220"></a>

<a id="canonical-1200111300002022-0012102113000022-3221032231232001-2332032311231313-1032302231233003-0000202220110103-2120011003203332-0213213321131322"></a>

#### `infra.hw_info.chassis.serial` property

Type: `"string"`. Computed.

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

<a id="canonical-1100010232201202-2300303003020021-0123122022220133-0302231213102030-2001322103130021-0011200302131303-3232323102201211-2230300302011002"></a>

<a id="canonical-1130313130311000-2122013310013200-2212321211222011-3033230230003021-2033100122122203-2311223023002112-2333303211303132-2110122202302323"></a>

#### `infra.hw_info.chassis.type` property

Type: `"number"`. Computed.

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

<a id="canonical-2321301021133103-3021022222013132-1320313111201001-0111020330213310-2210211102002120-0001003110201211-3010303233232100-1121111012301111"></a>

<a id="canonical-2200132312133313-0010022330232101-1200303302200120-1212201323031112-0033330103233212-2230013312031023-1131202221110220-3133313233132031"></a>

#### `infra.hw_info.chassis.vendor` property

Type: `"string"`. Computed.

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

<a id="canonical-0203202322233113-1311020030003200-1030330111313101-0130122003200300-1032113200233103-2310213130221113-0300133003320123-1300223033013013"></a>

<a id="canonical-1303000222211310-1232030331122211-3112330300033000-3212230303002202-0212020000001323-2313023333322201-1131013323321131-0202312310130332"></a>

#### `infra.hw_info.chassis.version` property

Type: `"string"`. Computed.

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

<a id="canonical-2310202130323300-2033301303131032-0201000311132033-1222111330021332-0312020321111232-1013202023021311-2030212020012223-3133012100322023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.cpu` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.CPU

<a id="canonical-2000002120230000-1211331133122112-0303212313032103-3230010032231103-0322222030300333-2202330210320023-0333133311233033-0230303303000013"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3132210000332332-3131302122211333-2013320301033210-2303113113332033-3313102013103211-1132201023222021-2300000233132033-0122231210302322"></a>

### Direct properties for `infra.hw_info.cpu`

<a id="canonical-3330302032120011-1122302332222032-3201021103103000-3013332102011130-0102110103321333-3102201011112012-1210112233202113-0211223031022112"></a>

#### `infra.hw_info.cpu.cache` property

Type: `"number"`. Computed.

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

<a id="canonical-0312023210322130-0332020123102333-1333231113020002-0202311002201123-2313331133213301-3210300311331111-0021333133213133-3103202211130010"></a>

<a id="canonical-1000121001000231-3120001020030003-2332233201322212-0201132200333211-3033011331131210-3013203033003110-0230212230113111-1032023213331300"></a>

#### `infra.hw_info.cpu.cores` property

Type: `"number"`. Computed.

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

<a id="canonical-3333023202233031-1210321023023311-1003222332002111-2113121322212132-0023322023020013-3230102202330221-2111023202110220-0113133000000020"></a>

<a id="canonical-3223111311232122-2330122232010132-2333130330303231-3103221331010103-2323210333131032-3132323212133023-2332212300032320-1201330130203110"></a>

#### `infra.hw_info.cpu.cpus` property

Type: `"number"`. Computed.

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

<a id="canonical-1021313011100222-1230110201201233-2003200322311101-0122132032332002-3310013322302020-2222312011121320-2030220222023021-0322010121123121"></a>

<a id="canonical-1103100300212230-3132030200203322-0331212320301011-3311123330333210-0013130302011320-1033130101311020-2223031121003231-3222113233001303"></a>

#### `infra.hw_info.cpu.model` property

Type: `"string"`. Computed.

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

<a id="canonical-0302011011022313-0302123300130320-2033333032013022-1023030313121233-3322110031222323-1032000002323223-1230322303123103-3300130033011103"></a>

<a id="canonical-2030012300233211-2232220302032133-2001212103102122-2332123312030322-1330011111212201-0131103330020122-0022211021103012-1311223132102120"></a>

#### `infra.hw_info.cpu.speed` property

Type: `"number"`. Computed.

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

<a id="canonical-0013333111120013-1311201111002112-0101131123033332-0312203011112001-3021232120233112-1010212330130233-1002301200231222-3210200202231232"></a>

<a id="canonical-3112012231311321-1122231023223230-0000121302332230-2132101002213110-1200301032000303-0202131120131203-3202221013011022-2302311130200120"></a>

#### `infra.hw_info.cpu.threads` property

Type: `"number"`. Computed.

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

<a id="canonical-0311301112302113-3230220212331011-0133130000000122-3310110023001131-0003310201331111-0132020302323102-3130001203323110-1032312100023123"></a>

<a id="canonical-2203100222322303-1331210120231321-0312310312020330-3001303011131022-1201032000201213-0101313223001303-3333120010231023-2302331100121021"></a>

#### `infra.hw_info.cpu.vendor` property

Type: `"string"`. Computed.

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

<a id="canonical-0120231302203333-1211031200013033-3000023223131131-3003131202220332-1101110002222320-2020023130030103-2213213010201021-2211113320220310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.gpu` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.GPU

<a id="canonical-3000231302121113-0032101332011012-3212011222003132-1231023232031103-2100001002031011-3012000333010000-1002223003303101-3320320003222130"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1330001322131110-1032212110210101-3231011101330032-1230211202223111-3031031330230203-3201101322212221-0123303300002323-0232222231103031"></a>

### Direct properties for `infra.hw_info.gpu`

<a id="canonical-2100131331201221-3310213303132323-0123221101221212-2030220221033220-0211123213021301-2020121033323332-0003103212003330-2112332231010120"></a>

#### `infra.hw_info.gpu.cuda_version` property

Type: `"string"`. Computed.

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

<a id="canonical-1222133002030210-2213112100331231-0223010022332322-3021121200020330-0022120021223302-1312033303211321-0012023220331301-0330032103231011"></a>

<a id="canonical-0321210212312110-3023002303233313-3131102130132103-2300200011300120-0112333222220320-2300211003010333-2131322010211100-0321012030220103"></a>

#### `infra.hw_info.gpu.driver_version` property

Type: `"string"`. Computed.

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

- [gpu_device](data-sources--registration--reference--group-001.md#canonical-2020012123012313-0100232310221002-3123303132303113-0201101333322121-3312232230123200-2303102232331222-3013233030002123-1203331312133012): complete subsection reference.

<a id="canonical-2020012123012313-0100232310221002-3123303132303113-0201101333322121-3312232230123200-2303102232331222-3013233030002123-1203331312133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.gpu.gpu_device` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- [infra.hw_info.gpu](data-sources--registration--reference--group-001.md#canonical-0120231302203333-1211031200013033-3000023223131131-3003131202220332-1101110002222320-2020023130030103-2213213010201021-2211113320220310)
- infra.hw_info.GPU.gpu_device

<a id="canonical-2311301030030330-1112313100231332-1202002022331213-3131232330012122-0131211100112300-2231103202031022-2220001100202032-0300032231201333"></a>

Type: `"list"`. Computed.

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

<a id="canonical-2220301233000001-2022213033032121-3333001130323100-0113222100231201-1030133212231300-1033131022232303-2201321021011302-1100321233310011"></a>

### Direct properties for `infra.hw_info.gpu.gpu_device`

<a id="canonical-2312032102310203-0020302120123131-0201220110200313-1001322110030103-1000220303311322-2010303132222311-2311121301332222-2022303020120003"></a>

#### `infra.hw_info.gpu.gpu_device.id` property

Type: `"string"`. Computed.

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

<a id="canonical-3312001203113002-2313230110032233-0020101201203321-1300102033102300-1132322233012221-3303210223330102-0023223300010121-2221100322113101"></a>

<a id="canonical-3130002321023222-3302001230132323-1210112101303033-0013332302111020-1221332003303122-0033312220130122-0101100022132331-3300223011310320"></a>

#### `infra.hw_info.gpu.gpu_device.processes` property

Type: `"string"`. Computed.

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

<a id="canonical-3110000022133133-3110210330101000-2001001231220020-1302132011311332-0313133001122012-1211102101321210-2033020330203012-2121000022123100"></a>

<a id="canonical-2030233101312122-1201101121220222-1023330030020303-0010122301302122-0012320312200333-1010333202132123-1321230231202020-2203021030221223"></a>

#### `infra.hw_info.gpu.gpu_device.product_name` property

Type: `"string"`. Computed.

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

<a id="canonical-3221100101322111-0322230203320111-2222313030300303-2111330331331222-3133232133112212-1023103301201132-2202120013320222-3133222303011322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.kernel` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.kernel

<a id="canonical-0001301233323231-1312032233200133-2001323103011111-0113310202333100-3122230100021122-2232130130112030-0300200001203332-2113230331333313"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2201133202100113-3111121331103122-2312232301011112-0220010022203203-3310122030010100-2013201110200132-1112230130303120-1231011002320112"></a>

### Direct properties for `infra.hw_info.kernel`

<a id="canonical-0031132321031322-2010211012122012-3300302312330211-0102033312331332-1201322232201111-1212101031221211-0033333132032001-1002333122131311"></a>

#### `infra.hw_info.kernel.architecture` property

Type: `"string"`. Computed.

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

<a id="canonical-3302322301233301-2113320102213230-1213031031322210-0112301022212213-1211211223020213-1033131233231232-0203230011003013-0132110123300121"></a>

<a id="canonical-0312131312022111-2320120202332011-2101330312230323-3202300110011300-0020120030113220-0013220110130330-3221103023213323-2232103032231213"></a>

#### `infra.hw_info.kernel.release` property

Type: `"string"`. Computed.

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

<a id="canonical-2030221031010333-2311120101322312-1022330311102022-2001112203233322-3021113233322220-0211203021101322-1310132010103203-2312122233021200"></a>

<a id="canonical-0121233212223221-3020223301301130-2332323233320023-3233312311101131-3131033200233023-3323030332320331-3203233130120211-0003031100332330"></a>

#### `infra.hw_info.kernel.version` property

Type: `"string"`. Computed.

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

<a id="canonical-0323000330332202-1230030330120123-2310033331331110-1303220032211210-0112023331101110-1201312233212003-3113001221310000-0022032120333011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.memory` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.memory

<a id="canonical-2131020102211122-0111202301322020-0230303022223132-0202313223030123-2100301031112110-2123033231322011-0012000202323233-3302303120000103"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2031231103110002-0020012200331310-2303223033032022-1132103010231032-1200321320000210-1331012332213213-1321122122230203-2210202221031313"></a>

### Direct properties for `infra.hw_info.memory`

<a id="canonical-0203232132310000-3321230020111331-2003000023001013-0023220031213122-2122031302011321-0002332111002023-2100111121122030-1132233313120031"></a>

#### `infra.hw_info.memory.size_mb` property

Type: `"number"`. Computed.

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

<a id="canonical-3130100000011322-2313121021021122-3032212231313231-2303123132011223-3301132103212300-3101301331120132-0120303002201220-1300200003223310"></a>

<a id="canonical-0130320330302320-0013211321033132-2113011120230110-3130201222110210-2002020311210332-2100103221223132-3131211030311200-1302021301011100"></a>

#### `infra.hw_info.memory.speed` property

Type: `"number"`. Computed.

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

<a id="canonical-2103321332233123-3112001101212132-1201031102011122-2110201032313031-2230101201110312-2000120230301210-1211121030021201-1302001002211010"></a>

<a id="canonical-3301232212212123-0221022222033110-2003222201011301-3003203113320001-2332302322001302-1312020232231030-3223120123113300-2321101033220200"></a>

#### `infra.hw_info.memory.type` property

Type: `"string"`. Computed.

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

<a id="canonical-2002003011233221-2101002023203311-0113332333132302-0203012031033203-1302212230320112-2233130100003303-0003203001333131-1022321123001211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.network` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.network

<a id="canonical-1122203303322210-3003200033200001-2211232023013203-0012013002001312-1102030203212210-0310110220130110-0223131122223001-0332300031111012"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2232122203311123-0231112212332200-2122221210100003-3121000210012313-2233023302023221-0112213001022020-1123120103031200-0313201332301012"></a>

### Direct properties for `infra.hw_info.network`

<a id="canonical-0111020102000210-1200030323032122-2010023220100322-3212033132022301-0032313102222000-3313130333320130-0010332330301031-2011311300320313"></a>

#### `infra.hw_info.network.driver` property

Type: `"string"`. Computed.

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

<a id="canonical-1013221310300111-1111101020202101-3231131203102220-1100312313323110-2303010103333320-1230033113231330-2030323002130000-0032012210202013"></a>

<a id="canonical-0103333001132110-0032300111102333-0322303332330100-3313012323202131-1313221311023111-1133123221102202-2010012113033222-1001323213032001"></a>

#### `infra.hw_info.network.ip_address` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2303000022221030-1100022100230000-3223330100203031-0103103100231202-0132300003100231-3000201011201131-0013303111021011-1100123332003303"></a>

<a id="canonical-3300212121102331-2101210211303333-0033111211220313-1010032203100220-0221120312310230-3101301320110001-2133033130001223-2220313021031102"></a>

#### `infra.hw_info.network.link_quality` property

Type: `"string"`. Computed.

\[Enum: QUALITY\_UNKNOWN|QUALITY\_GOOD|QUALITY\_POOR|QUALITY\_DISABLED\] Link quality determined by
VER using different probes Unknown quality Link quality is good Link quality is poor Quality
disabled. Possible values are \`QUALITY\_UNKNOWN\`, \`QUALITY\_GOOD\`, \`QUALITY\_POOR\`,
\`QUALITY\_DISABLED\`. Defaults to \`QUALITY\_UNKNOWN\`.

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

<a id="canonical-0233122033001210-1003231001332113-3322200112022233-2311103033003301-3103313233131332-1120102002300310-0032301333321012-2000213121122300"></a>

<a id="canonical-2021222321210203-0133333311313301-1001100232212130-3132130131133313-0302132133301000-1232222211020212-2200112021221321-3002000032013210"></a>

#### `infra.hw_info.network.link_type` property

Type: `"string"`. Computed.

\[Enum:
LINK\_TYPE\_UNKNOWN|LINK\_TYPE\_ETHERNET|LINK\_TYPE\_WIFI\_802\_11AC|LINK\_TYPE\_WIFI\_802\_11BGN|LINK\_TYPE\_4G|LINK\_TYPE\_WIFI|LINK\_TYPE\_WAN\]
Link type of interface determined operationally Link type unknown Link type ethernet Wi-Fi link of
type 802.11ac Wi-Fi link of type 802.11bgn Link type 4G Wi-Fi link Wan link. Possible values are
\`LINK\_TYPE\_UNKNOWN\`, \`LINK\_TYPE\_ETHERNET\`, \`LINK\_TYPE\_WIFI\_802\_11AC\`,
\`LINK\_TYPE\_WIFI\_802\_11BGN\`, \`LINK\_TYPE\_4G\`, \`LINK\_TYPE\_WIFI\`, \`LINK\_TYPE\_WAN\`.
Defaults to \`LINK\_TYPE\_UNKNOWN\`.

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

<a id="canonical-0111013322001310-1231312210233231-1330213332230312-3112022310001302-2232213233303231-1111022111223000-1113230110123330-1012231022131221"></a>

<a id="canonical-3113330221321303-2322230201012330-3132310032001331-1321310010201232-3333021231230103-1302303120221022-0221111220133033-3133113230301110"></a>

#### `infra.hw_info.network.mac_address` property

Type: `"string"`. Computed.

MAC Address. MAC address on interface.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2221230021002030-3232200203333331-1113312312303210-1310033121012111-2301210232132103-3320232222332013-3022312232010330-0231101111122210"></a>

<a id="canonical-3101210131232233-1233012011223322-3323312203313022-3233133030000321-3332232103200002-3132130033001002-1020201331301030-0323131102212323"></a>

#### `infra.hw_info.network.name` property

Type: `"string"`. Computed.

Name. Name of device, eg. Eth0.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2033223002311210-0131013211231032-3112011202212213-3021312333232221-2031031112001213-3321002113313032-0012200210202131-2322311011002000"></a>

<a id="canonical-1222330220100232-2213210003333300-2323323002220310-3021313311213003-2021011113332301-0001320320113203-1122031212311132-3033112121101312"></a>

#### `infra.hw_info.network.port` property

Type: `"string"`. Computed.

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

<a id="canonical-1331210331102020-3010330212223111-0300310230221120-1231112102332331-1120021311112221-2300223233121310-3031222212301001-2020123121332003"></a>

<a id="canonical-2233210121021331-0031113023113302-0221103322333102-2232311003112320-1211312232230132-3101102310120212-2021302231202003-1021001111333130"></a>

#### `infra.hw_info.network.speed` property

Type: `"number"`. Computed.

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

<a id="canonical-2023302323122322-3003210101112121-2100312221020002-2332203332311233-0010331003002231-0120201232012120-2332013030003023-0101232230013331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.os` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.os

<a id="canonical-3102202100220232-0003312222031102-0222301331131100-0203111223330221-1302300200010220-3012220230331122-3023310132022002-1232021223312231"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2301311213020010-3231102312201221-2201212221110002-1003213323133330-2012311313100331-2033233010011203-0112100023110322-2120033110231031"></a>

### Direct properties for `infra.hw_info.os`

<a id="canonical-1323132130310220-1333301312301321-0113321010011023-0212032023013211-0111122322110130-2013131031220330-0221131100321211-0130212002233032"></a>

#### `infra.hw_info.os.architecture` property

Type: `"string"`. Computed.

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

<a id="canonical-2231001033121103-1112130223003323-1033310302313132-1133012200023010-3333233120321032-0112013123123233-0110102013130202-2202223102330022"></a>

<a id="canonical-0131202020312003-2203331133313200-0210232300011112-1132233310110211-1322220022010021-1032112111223223-1102132211331230-1323031132123132"></a>

#### `infra.hw_info.os.name` property

Type: `"string"`. Computed.

Name. Name of OS.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3212321221021231-0302003021112231-2222023020033312-2131200021103003-3110212302123003-3121123231223113-3032200222100231-2320300132101212"></a>

<a id="canonical-2232111303001021-1211131021330333-0100320232332220-2033031332033000-2011131221102330-3300311212122111-0103121010213131-3312202221020102"></a>

#### `infra.hw_info.os.release` property

Type: `"string"`. Computed.

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

<a id="canonical-2010013012102102-2133311121122030-2120101100110031-0320030332203210-1300222002312030-1022101131211003-0112232111133333-1110103333103232"></a>

<a id="canonical-2012331123110313-0322332102220122-1002201312001211-3013223120113123-3323300213021113-1203002021131233-0300222223233230-2200230211110301"></a>

#### `infra.hw_info.os.vendor` property

Type: `"string"`. Computed.

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

<a id="canonical-3323221331112233-0032332133133112-0102232021133321-3123202101232031-0230321000123333-2132231330100231-3132032321103302-2122112302000131"></a>

<a id="canonical-1113120110101322-1222312331022131-2031332332212321-0300230113223111-1023123001310100-3311101112001222-2220112202020320-2033010113010010"></a>

#### `infra.hw_info.os.version` property

Type: `"string"`. Computed.

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

<a id="canonical-1130310222113311-3300132232301211-2013110222110020-2030303131202120-3223322133310021-0322013302100103-2312221203132313-3000332022101223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.product` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.product

<a id="canonical-0001331130320000-0001012311100102-0213223112303000-2310131213331003-1013302323333012-0221322113013301-0310100031211230-3021021011121201"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1203012301003210-0212330110033201-1131102103321201-2013113012203231-1312131310103121-3310121320333130-0112100323013110-3221102333310321"></a>

### Direct properties for `infra.hw_info.product`

<a id="canonical-3200010310022222-3112001121132032-3022203303233033-2133302323132303-2302333020110002-2033330332003221-1330233232012311-2032332313133023"></a>

#### `infra.hw_info.product.name` property

Type: `"string"`. Computed.

Name. Product name, eg. For AWS m5a.xlarge. Info taken from /sys/class/dmi/ID/product\_name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3033132211303101-3333000121220203-2020002123231220-0113210021202332-0120122303333110-2113332111113202-0311020122102301-3113330303332111"></a>

<a id="canonical-1210231313102322-0003211332130211-2221030112311332-2333233221322200-3230213011000223-3230200012301001-0202330000213203-2232133312032202"></a>

#### `infra.hw_info.product.serial` property

Type: `"string"`. Computed.

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

<a id="canonical-3310210330200203-2312223102220321-0002000121132010-2020303233220222-0332120221121020-1103211002311121-2323103113031320-0122200033203220"></a>

<a id="canonical-2021033330220201-0332313221122213-0202131113123033-0312301221330223-0002330020202211-2101131021111000-0101223132221011-3023300303223202"></a>

#### `infra.hw_info.product.vendor` property

Type: `"string"`. Computed.

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

<a id="canonical-3302301223201113-2323122301000101-3223330210331310-0220212011111223-1331113112301113-2311220120301301-2221321021211002-3223032303022232"></a>

<a id="canonical-1022100300212131-2331333311232233-3332220333012312-1032001322212333-2011300313031203-0020020331320133-1320000123220112-1221001210011121"></a>

#### `infra.hw_info.product.version` property

Type: `"string"`. Computed.

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

<a id="canonical-3123013201111023-2110203003311322-1023313313031222-1302230023123001-3122301203113321-1220031100211023-0012313202001122-3313013022312331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.storage` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.storage

<a id="canonical-0103223310021211-2222012313000102-2020002002330120-2100132200013013-3213130112301030-0000202301012200-0323232130013313-1021311202023311"></a>

Type: `"list"`. Computed.

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

<a id="canonical-3010203310321331-3323222221022310-1022213200223131-3211223301331032-0211211220013202-3123233131231010-3212011031321203-1221323300322332"></a>

### Direct properties for `infra.hw_info.storage`

<a id="canonical-2123323113032130-3000233222010311-2013100123112102-1213203213131232-1020000212321122-2321033321010122-0000312100330323-0300030213100222"></a>

#### `infra.hw_info.storage.driver` property

Type: `"string"`. Computed.

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

<a id="canonical-2332212101202023-0232213202313113-0312211100230210-2002101111103222-3301031022223113-1331131233130013-3310130121231021-1021112232321231"></a>

<a id="canonical-2021010131330303-3121022222103332-2022132232101020-1103331132100030-0123320311332201-0100210020203100-3231001012333030-2201120330303211"></a>

#### `infra.hw_info.storage.model` property

Type: `"string"`. Computed.

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

<a id="canonical-1311111000233323-2001323232211302-2132023022330302-0000211200331000-1020100301000300-0223233023012023-0002313322300003-2100002003221102"></a>

<a id="canonical-0322010111113311-0321223000123221-2130332330022100-3202133130230220-3132321322001010-2131333332020033-0322012303023000-2023033120312231"></a>

#### `infra.hw_info.storage.name` property

Type: `"string"`. Computed.

Name. Name of device, eg. Nvme0n1.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0101221030210002-0331333110302233-2231331000303221-1211320201232230-2001321023000231-2201000331003213-1112100000020002-3100232303321002"></a>

<a id="canonical-2103002213223322-2020120300203110-2013311102231333-3321032120233123-3212023331111001-2020300230201130-3221022303113132-0100112032032123"></a>

#### `infra.hw_info.storage.serial` property

Type: `"string"`. Computed.

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

<a id="canonical-0323213230032103-3202331320000122-2223323201223021-0110233032233032-2023132130202322-2003213123123222-0300220131101023-1301313320233102"></a>

<a id="canonical-0010030021021333-3133323200311333-2023230003303032-0303323021313111-2032002230313313-0200310133123013-2130212001012122-0310001030313000"></a>

#### `infra.hw_info.storage.size_gb` property

Type: `"number"`. Computed.

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

<a id="canonical-2332221120231233-0223203123313102-2023031101023201-1022322000031122-1031203033001133-0301331103300221-1102232213220011-2230000031210200"></a>

<a id="canonical-2011100133310333-3013221010022010-1220330101123001-2103030231200303-0103131211330303-0120230310330010-3021122221032223-2200111002112120"></a>

#### `infra.hw_info.storage.vendor` property

Type: `"string"`. Computed.

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

<a id="canonical-1110023331232012-0110323002300313-0131123232121130-0310333130302000-0321212033001212-2330232222300223-2313213031323201-0101213101122321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.hw_info.usb` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- [infra.hw_info](data-sources--registration--reference--group-001.md#canonical-0003310021103010-3122202131113301-0331323321110213-2312121331103122-3103202101303101-2323130030130123-2303303203021203-1033300032103110)
- infra.hw_info.usb

<a id="canonical-1201303230230100-3020203212333213-2012333220013110-3120133001201011-2033200200311202-1320111312202003-2001322002121130-1120201330033302"></a>

Type: `"list"`. Computed.

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

<a id="canonical-3330301103001313-1112321312301333-0222132332223320-3210021033033002-2103333232120201-1110223310323222-2132210231031020-0312210020233302"></a>

### Direct properties for `infra.hw_info.usb`

<a id="canonical-1133313112212113-2302230032223012-1030000020320222-1223221012013331-1232133330210201-2123230300201103-0322131033320123-2001312201011301"></a>

#### `infra.hw_info.usb.address` property

Type: `"number"`. Computed.

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

<a id="canonical-2112222202111132-1010102031233011-2201233020113013-0303230123203020-2320023231010120-1320120301203120-3331012233012301-2103112322211003"></a>

<a id="canonical-2231223102103032-3032332002213321-0233002230101132-0311013021133123-2022100010110133-2120203011212132-2131313101030111-1303101311322310"></a>

#### `infra.hw_info.usb.b_device_class` property

Type: `"string"`. Computed.

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

<a id="canonical-3021010322002102-0120222212203333-0102031203222230-2303031201311010-3220020232300112-0310230320103100-1311013333001110-0333033212013323"></a>

<a id="canonical-2331232021003200-1133112000232013-1312333210102102-0031230120310313-3212211112213212-2033130031100111-2000000232123320-0223013122312032"></a>

#### `infra.hw_info.usb.b_device_protocol` property

Type: `"string"`. Computed.

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

<a id="canonical-2331301221101100-3210122302123023-0221310313110330-0122122022332230-1222003230310222-2122000232120122-2130121300013102-3033203110000312"></a>

<a id="canonical-1020220300222222-0010132330010200-3131110230132330-1021300312211032-0220123311020223-2133100120220001-2202220023022022-1132102322131211"></a>

#### `infra.hw_info.usb.b_device_sub_class` property

Type: `"string"`. Computed.

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

<a id="canonical-0300120301012312-0103032301232131-3010211030222120-0320323110221012-0232303131222310-3021131331202230-1120030122302311-0232300110011210"></a>

<a id="canonical-0032231032100213-2333211133231310-2212301301120032-1120211033233320-3112030213330100-2213012030232303-1312132021310100-3122013013220210"></a>

#### `infra.hw_info.usb.b_max_packet_size` property

Type: `"number"`. Computed.

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

<a id="canonical-1320210121132202-0003220233003230-3213231221012110-0302333212230130-1112100232211303-3233101132011033-0132122303212130-1223231210330233"></a>

<a id="canonical-2323103302030131-1312103111310300-0101320323313120-3303100013233313-0321133100213233-1202312122113300-1033013022022200-2322321311103130"></a>

#### `infra.hw_info.usb.bcd_device` property

Type: `"string"`. Computed.

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

<a id="canonical-3232012323102022-0310231110311203-3021302010000211-0331133110030031-3131323013330232-0130110323333012-0213022311003312-1330223231231123"></a>

<a id="canonical-0330023010020120-2321032212113002-0311332102113023-0013203102300110-2030130230112222-0232110312102313-0302022311310013-1323223211001123"></a>

#### `infra.hw_info.usb.bcd_usb` property

Type: `"string"`. Computed.

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

<a id="canonical-1210312230313233-1212231221201233-2211200122121323-0310130212330210-0330023231323123-1220120201000210-0333231110322230-2201331111300023"></a>

<a id="canonical-0102110310010333-2011130121331220-2211321003201002-0223202211230112-3320033003131021-0130202303301100-3303121101301330-3133221000300231"></a>

#### `infra.hw_info.usb.bus` property

Type: `"number"`. Computed.

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

<a id="canonical-1201210323233320-3110023201010110-1332113222300223-0320200130312103-1110113231201201-1310110110310223-1303320331031031-1023113322322212"></a>

<a id="canonical-1122310021202033-1023131111131321-0331213011100221-1200111332222321-3132033221233000-2320013121212001-0010300030321131-1120232013210103"></a>

#### `infra.hw_info.usb.description_spec` property

Type: `"string"`. Computed.

Description. Device description.

<a id="canonical-0032211010210132-1102202013031302-0302001002333002-3220300102330003-2103121100022330-2330222030111301-3220010103111302-1310321122213310"></a>

<a id="canonical-1101032110212100-2000301113020212-2013303201301220-3102133113221020-1321211300201200-3301011032312222-2322212021331031-1100301222213033"></a>

#### `infra.hw_info.usb.i_manufacturer` property

Type: `"string"`. Computed.

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

<a id="canonical-3112330102200311-0011231013133311-1230032031130201-3020031130301311-3230133310131230-0232120132010320-3011312210121233-2112132310113121"></a>

<a id="canonical-3100010020212130-1100210020222023-2020331321332020-1132101032003023-0100000113001310-1202100231330100-1230010002213002-3021122333122331"></a>

#### `infra.hw_info.usb.i_product` property

Type: `"string"`. Computed.

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

<a id="canonical-1133232132022012-1112222321211002-0220031232201312-1200203000110132-0120302330331300-0013200120232332-3011211210321121-0021201101030301"></a>

<a id="canonical-3232023311331120-3022333313310202-3130000002132302-0012330333010303-1330002100022303-1123121103223031-3013120220311120-0101200021111112"></a>

#### `infra.hw_info.usb.i_serial` property

Type: `"string"`. Computed.

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

<a id="canonical-1033311210133301-2300122020031131-3133321121013130-0032300232220333-0321233131300123-3302312002132230-3000101332221132-1320302220132230"></a>

<a id="canonical-1311123123323131-0130013131222320-0010302000021332-2202033230222102-3233230202013003-0200101230113321-0110023131202100-2130330332103000"></a>

#### `infra.hw_info.usb.id_product` property

Type: `"string"`. Computed.

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

<a id="canonical-3312121101331031-0111103212003010-3111003311102220-2023011332232313-0202101130020103-1211301210113331-0330222203330001-1333130000333332"></a>

<a id="canonical-1213322032221110-1001202212112002-3032013133313020-1302313131312022-3033333220102310-1203122301030312-2222320102310113-2331031103312312"></a>

#### `infra.hw_info.usb.id_vendor` property

Type: `"string"`. Computed.

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

<a id="canonical-2313131332313202-2121003310023013-1020010001310013-1232302121313013-2211021022300201-1112312311010221-1233111131323300-3123023133030321"></a>

<a id="canonical-0103220020220110-0212111013230103-1312102022012232-0110100203233122-2113032333213131-0112003332120030-2121232012213311-1003012122120020"></a>

#### `infra.hw_info.usb.port` property

Type: `"number"`. Computed.

Port on which the device was detected in decimal.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0021012131233021-0120311032222120-0311021221132320-0121001013210002-1322112133310033-0303102120330333-2230222221333132-2230332322311222"></a>

<a id="canonical-2231203133203211-1222320311023121-0233023203313303-1130122000322232-2202231200213111-0301113100002133-1001122000122330-1122012121111010"></a>

#### `infra.hw_info.usb.product_name` property

Type: `"string"`. Computed.

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

<a id="canonical-2201101133211331-1113331113103222-2120131231203230-1233202121333301-3330211203002013-3211223120002211-0032302330112311-0310222223231010"></a>

<a id="canonical-2121311223301201-3323210320211112-1231002132333321-0101211210133012-1110321230210113-0320302230213131-2112133000030012-3211332021112000"></a>

#### `infra.hw_info.usb.speed` property

Type: `"string"`. Computed.

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

<a id="canonical-0011032100222322-3120102000200032-1322320230200221-3312213231311203-0002303201120011-3032033001130332-0031032312301121-2131223200033220"></a>

<a id="canonical-3321203300013023-3321131331220022-3321112030222113-3023011033003231-1201133313120232-3211320323300232-1301213303112032-1221111101102332"></a>

#### `infra.hw_info.usb.usb_type` property

Type: `"string"`. Computed.

\[Enum: UNKNOWN\_USB|INTERNAL|REGISTERED|CONFIGURABLE\] Type of USB device Unknown USB device type
Internal USB present in Certified HW USB device present during node registration USB device that can
be matched by USB rules. Possible values are \`UNKNOWN\_USB\`, \`INTERNAL\`, \`REGISTERED\`,
\`CONFIGURABLE\`. Defaults to \`UNKNOWN\_USB\`.

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

<a id="canonical-3333202113000000-0212023111213230-2032320020102013-1303223010300202-0020121111100132-2200133030112131-2102311313020212-1122320030312233"></a>

<a id="canonical-1022021103113310-0230123233233101-2223130012000220-0030130332121233-3230233200331013-1222301100311212-0300103101333122-3330002310220320"></a>

#### `infra.hw_info.usb.vendor_name` property

Type: `"string"`. Computed.

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

<a id="canonical-1212321321230203-1333031332333111-2013211110013002-2012312130010010-3132000213031022-1302312311201033-1112302223303330-1312233300011331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.interfaces` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- infra.interfaces

<a id="canonical-3030011312300101-3100323121201120-0320212110131033-0002010300332013-0231311323320301-2033001120222211-1210120200130311-2311102230111233"></a>

Type: `"single"`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022113132332131-3301201322311101-1303311110321310-3103112022020203-2132031321200102-2100333112033011-0222212203001113-2210131301030123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.internet_proxy` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- infra.internet_proxy

<a id="canonical-1300102132123302-3113312100320313-2203310020010231-1300223120023233-0210202010033122-0031020220203031-0211311300223101-2101321201313212"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2132002230202323-3211103122022012-1311320201203202-3030011022223321-1200212121130011-2102330001122321-3112023312201131-2310101131331101"></a>

### Direct properties for `infra.internet_proxy`

<a id="canonical-0310300122132010-1113110201212023-2102113202332320-3202230312301233-3030002130201023-1333100013231032-2220003003302331-3133331330210201"></a>

#### `infra.internet_proxy.http_proxy` property

Type: `"string"`. Computed.

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

<a id="canonical-0102112121222010-3333231031233010-3302312202210012-2011301220222033-2102122021120101-0323021133103213-2003302320320212-0203320101210212"></a>

<a id="canonical-1122133311322012-2230111300120132-2312310031120210-3232131302011121-2132012323102110-0031103201122123-3212021132030200-2203211312111112"></a>

#### `infra.internet_proxy.https_proxy` property

Type: `"string"`. Computed.

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

<a id="canonical-0020133012101103-0020133102122002-0131232120200211-2203111231222003-1203012310321001-3031231230111301-0123230100331110-2333220303313320"></a>

<a id="canonical-0233003033131100-2132230331121331-3022011212323213-2202323312101001-0102022013331313-3310203031211130-3321013022300132-2202000221220120"></a>

#### `infra.internet_proxy.no_proxy` property

Type: `"string"`. Computed.

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

<a id="canonical-0031332233000331-0121112331011323-3003213012030120-2202023101310113-0331113003213121-2032002332102100-1223231332020133-0130331300230032"></a>

<a id="canonical-1113012300031012-0100221313102022-0031200022012301-1023302322222003-0012022311012201-2103003201100300-0233112210301003-3313202233333210"></a>

#### `infra.internet_proxy.proxy_cacert_url` property

Type: `"string"`. Computed.

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

<a id="canonical-1231311000020020-2022312100213123-2023300303302201-1200001213120331-1211131030203221-1032233033010103-3210023002233002-0031113231233303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `infra.sw_info` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [infra](data-sources--registration--reference--group-001.md#canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101)
- infra.sw_info

<a id="canonical-0003013103231100-0313031200312102-0230111123113122-3302100301133112-3213230120020223-1020130130121031-3320120230201031-3102111033211231"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0021102232123203-0302130320320300-1130013200322012-1000333232100320-3233123210031300-2220211132121001-3312201121210000-3031123301333102"></a>

### Direct properties for `infra.sw_info`

<a id="canonical-3033303202000302-0012331232200002-3201323212132111-0210303122301012-1213232112222203-1320102121203111-2030321122102232-0130220021322110"></a>

#### `infra.sw_info.sw_version` property

Type: `"string"`. Computed.

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

<a id="canonical-0223232121310231-1203011310020223-0003303003232021-2210303010321002-0213203010232110-3211330322101103-3313301302301300-0202203223232211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `passport` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- passport

<a id="canonical-2211123122231133-1322031103221030-2232331132131310-2221202133030102-0110220101320300-2132313310232222-1312002231012220-1131111000312210"></a>

Type: `"single"`. Computed.

Passport stores information about identification and node configuration provided by CE during
registration. It can be manually updated by user during approval.

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

<a id="canonical-1231311330311210-1011303101231123-2331203203300213-3002233301300321-3011323011301213-3321220222231013-2011202221200202-3332332311132010"></a>

### Direct properties for `passport`

<a id="canonical-1100313202130030-3122012022313032-2122032023210012-2323301100310132-3002201020120303-0212131220223123-0311233231132303-3100211203013323"></a>

#### `passport.cluster_name` property

Type: `"string"`. Computed.

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

<a id="canonical-0010022003123222-3332210222213033-0231002022130032-0233122103212123-1223012102111112-3331233201000322-2323123301203210-3010222133303210"></a>

<a id="canonical-1022222330310312-0002233322002202-0300310023100021-0010213021233231-0113210313213012-1011013110223320-3022031201311333-3132212103023223"></a>

#### `passport.cluster_size` property

Type: `"number"`. Computed.

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

<a id="canonical-1230003210213110-2133201130300233-0003001020232330-0011102302322331-1002023101330131-3332333322012211-3311311333030010-0202201203232321"></a>

<a id="canonical-0213010221221002-0103313323032001-0023220301313210-3130030023030203-1232111020020012-0231031100012001-3012301312122100-2212133332000231"></a>

#### `passport.cluster_type` property

Type: `"string"`. Computed.

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

- [default_os_version](data-sources--registration--reference--group-001.md#canonical-1301131031220001-1012233231111102-2312012311030032-2002033211200231-0300000321112312-1122110320103311-3312031211020310-2321133310203310): complete subsection reference.

- [default_sw_version](data-sources--registration--reference--group-001.md#canonical-3011233211021332-2030122230313110-1111212010123012-3330230001122300-1230222203303102-2131210311002031-3310322202301312-2331212212131331): complete subsection reference.

<a id="canonical-1103032331313131-2002111002122112-0330300023210313-3211111221202223-3001131311030001-0002000123030212-2211021323130111-1220003030220301"></a>

<a id="canonical-2033221202030312-0313200230032011-0330001113031212-1030010311320220-2223313332310130-2102312113011033-3132121101021021-2332210103203023"></a>

#### `passport.latitude` property

Type: `"number"`. Computed.

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

<a id="canonical-3110010212312331-3110202322013211-1103230120223222-2131111330331031-2232202311310123-0000012022001000-0021011110013332-3332230000003211"></a>

<a id="canonical-2312022131213012-1000112011022220-3103013222010123-1001113003010130-0203230121331001-1332322132200031-3222123133002331-3111030303012310"></a>

#### `passport.longitude` property

Type: `"number"`. Computed.

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

<a id="canonical-2213322102331012-0032303330202211-2130211321203012-0203033003330303-0222302001113133-0103012102030322-0332033301020211-2323220123330012"></a>

<a id="canonical-1013101022023113-0101331320020122-2300232230030312-3001122103231130-0321103200012330-3100112320213231-3310130231213010-2331122310031232"></a>

#### `passport.operating_system_version` property

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-3211022121311122-1313103020323210-0331000313200113-2212213311223320-1231003332133312-0122121010312210-2232110130231012-1321213031202300"></a>

<a id="canonical-2312132032331331-3003230123101332-0303123211211023-1111311003323213-2320132212312313-3233230113322213-1212011200212023-1102210031202231"></a>

#### `passport.private_network_name` property

Type: `"string"`. Computed.

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1111303001322301-0010230002312123-3313123222111100-1302100123302232-3210021231200232-3021233032133200-0321121220002210-0032210212220221"></a>

<a id="canonical-1221012332120213-3022302110033303-0112020212021023-2033230201020303-2221201310210113-2013202113333021-0321320100013022-2231210031002012"></a>

#### `passport.volterra_software_version` property

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-1301131031220001-1012233231111102-2312012311030032-2002033211200231-0300000321112312-1122110320103311-3312031211020310-2321133310203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `passport.default_os_version` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [passport](data-sources--registration--reference--group-001.md#canonical-0223232121310231-1203011310020223-0003303003232021-2210303010321002-0213203010232110-3211330322101103-3313301302301300-0202203223232211)
- passport.default_os_version

<a id="canonical-2021231302100201-1122030102210232-3210202120333102-0310321031003322-3111213231023303-2003030100230232-2312011303330010-3011021220330102"></a>

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

<a id="canonical-3011233211021332-2030122230313110-1111212010123012-3330230001122300-1230222203303102-2131210311002031-3310322202301312-2331212212131331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `passport.default_sw_version` properties

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301)
- [Property reference](data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [passport](data-sources--registration--reference--group-001.md#canonical-0223232121310231-1203011310020223-0003303003232021-2210303010321002-0213203010232110-3211330322101103-3313301302301300-0202203223232211)
- passport.default_sw_version

<a id="canonical-2003033332110300-2200301110012012-0323311023210233-1122212221203231-2120203131102320-0023311030021123-0231001210120030-3130332030012113"></a>

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
