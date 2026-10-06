---
page_title: "xcsh_dns_lb_health_check reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check reference."
---

# xcsh_dns_lb_health_check reference

<a id="canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- Property reference

<a id="canonical-3002033311333003-3233233303022131-2131303032111112-0112123032222310-3133223131122212-1023213201122003-2311020122311122-1132122122001011"></a>

### Direct properties for `xcsh_dns_lb_health_check`

<a id="canonical-3213122131011223-0231002322223311-3210101112333112-2001311130202231-2103222312303200-1203132323101302-3032231110322232-3330211210110233"></a>

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

<a id="canonical-0221102330002203-2213312210212232-0033303333111012-3202230010110212-2303121101113113-3211113111121320-1332212011310030-2201210102202212"></a>

<a id="canonical-2120231222101202-1033331010201032-3221131231331330-0233202201021210-3232000110221303-1122210021132232-2312023322203302-3303001033320132"></a>

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

<a id="canonical-0333123030332121-3231111113211022-0022312022002000-3113321110032212-0322230333230311-1202011030013000-3013322323330231-1230010332311021"></a>

<a id="canonical-2010330010210011-2120222021123021-0013013233211211-2103230232120313-2101211333331112-2011133233131310-2302321200110213-0220021010212213"></a>

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

- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-3123301133033020-3001312112033103-2023010020133212-1212321113313010-2303023100323100-0031323300231202-0032201133132201-2022313223021312): complete subsection reference.

- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1013121033333001-2311212031003021-3122120000133312-0311032031230022-0031330011111330-3130133200012211-0102332022220133-2100120110113333): complete subsection reference.

- [icmp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-0022113030000323-2330222301031023-0120100012200301-2333110032132213-0003230321003203-2023312222130303-1322131221322010-0133303003303103): complete subsection reference.

<a id="canonical-0032121332010132-0130123330021333-0310201033102311-0130020320112201-2111330312031220-1110200113103013-2122213023323320-3330032321101023"></a>

<a id="canonical-3221132013203223-1131323311300021-1322200111221130-0312230011112301-1000212001102330-3210122021201023-2111320320000231-3213310311121212"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1210321202303301-1331200321120102-3221233213312211-1310021110030223-1300101100213012-0323030000021112-2222132022212130-1331220300213231"></a>

<a id="canonical-0123121232122103-1130101233232330-2300310230223221-2330030200101003-2313132011232220-1312133130223020-3220200110312012-0203202030122023"></a>

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

<a id="canonical-1132300312011302-2201300011333131-2021210313302012-0120303332111212-0130203331323021-2132123003022330-1312331211102201-3023102113210223"></a>

<a id="canonical-3021200121032133-0030300222230332-0321031221333133-3121122322333121-2101000200232130-3222333023330112-2033312101133031-2001023002322133"></a>

#### `name` property

Type: `"string"`. Required.

Name of the DNS LB Health Check. Must be unique within the namespace.

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

<a id="canonical-0000022000202101-0331300212232011-2202013220031000-3201210211302221-3032121231212322-1211201101020101-0021301331202131-2032320103031333"></a>

<a id="canonical-3120203210003321-3101033120111112-2101110133003110-0023110120222302-1202012001011111-1312103202123000-0230011202022100-0211130321113213"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the DNS LB Health Check. The F5 XC API restricts this resource to the system
namespace; it defaults to that value and may be omitted.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
EnumExtractionComplete: false
EnumValidators: [{"version":1,"validator":"OneOf","values":["system"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [tcp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1202311232132230-0212321030133021-3022310202313131-3312222212221302-0133231310311030-0303013212230320-1231023213133231-0200233332322222): complete subsection reference.

- [tcp_hex_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-3310122233311003-0331022321313010-3222232300211013-0210133110202032-0302023320232031-1333300003101322-2030330013112003-2012332003332031): complete subsection reference.

- [timeouts](resources--dns_lb_health_check--reference--group-001.md#canonical-3100003210332223-1023023222311302-1120310230130002-1321123111113000-2033201021003301-2300302321021312-2131002323220212-1303032312322311): complete subsection reference.

- [udp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1033111102330311-0210031231001012-0311000130202221-2102102013102323-2102322000022300-3023201333202310-3212111232132212-1013112131001321): complete subsection reference.

<a id="canonical-2102120201012001-2022102112002030-1130202120112132-0101303020123332-1121120020011202-3201302103003013-3301313320210032-3002330213013223"></a>

### All schema paths for `xcsh_dns_lb_health_check`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_lb_health_check--reference--group-001.md#canonical-3213122131011223-0231002322223311-3210101112333112-2001311130202231-2103222312303200-1203132323101302-3032231110322232-3330211210110233) |
| `description` | [description](resources--dns_lb_health_check--reference--group-001.md#canonical-0221102330002203-2213312210212232-0033303333111012-3202230010110212-2303121101113113-3211113111121320-1332212011310030-2201210102202212) |
| `disable` | [disable](resources--dns_lb_health_check--reference--group-001.md#canonical-0333123030332121-3231111113211022-0022312022002000-3113321110032212-0322230333230311-1202011030013000-3013322323330231-1230010332311021) |
| `http_health_check` | [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1221130202322331-2130221223120123-1211123222201231-2233103123111303-2032112002130020-2220011132023230-1313232301122121-1231133313310202) |
| `http_health_check.disable_virtual_host` | [http_health_check.disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-3130301010010113-0302111222100311-1021033230212102-2231301320120313-0321030130202233-1123003131233233-3130032002311310-3001222212331222) |
| `http_health_check.health_check_port` | [http_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-2031220000233033-2100002331130132-1322223211300203-2013330102020303-0003032302330210-2100221031212211-1113131202133300-3313011323320133) |
| `http_health_check.health_check_secondary_port` | [http_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-2213302303303101-1302302300331110-0212001011210100-1330200130300312-0322131010213030-0200230220110221-2303100323301023-2332003113223231) |
| `http_health_check.inherit_load_balancer_fqdn` | [http_health_check.inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-3220231232213020-1223011031021132-2011020223100333-2111330122132121-0313220120132111-1311203320333031-3021133221000331-2230333120212103) |
| `http_health_check.receive` | [http_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-1112003222332101-0221102131312230-1111233130331021-3213013012120323-0203313223111000-3032001121003122-0112222031002321-3231202031110000) |
| `http_health_check.send` | [http_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-2112330130302033-2322102211022112-1033233123332031-1321111223032013-2013212120110200-1300103010222232-0121221202222200-3233032100021023) |
| `http_health_check.virtual_host` | [http_health_check.virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-0320322010003211-2023133311310302-0103200100313130-2220021301023313-0303333112301302-1322100123210303-0321022020121303-0021230211212313) |
| `https_health_check` | [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-2121323320010013-3120203001021122-1001323231220330-0231100022101023-1231011100023131-0120121102300101-3212131332013100-2220103210020310) |
| `https_health_check.disable_virtual_host` | [https_health_check.disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-0201032311031110-0210300001030312-3020001332301012-3203110110201103-2331323112001223-0223221031301003-3023013311113120-3011100113322123) |
| `https_health_check.health_check_port` | [https_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-2213211221323112-0112313202121011-3233100313220310-2030212033022212-3230022111003031-0211131030311312-0023232030010221-3023312002321212) |
| `https_health_check.health_check_secondary_port` | [https_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-3020211013120202-3132303110011022-2020300302120223-1331000123102123-0001132130213112-2322223121310030-1031233111312300-3322323311000303) |
| `https_health_check.inherit_load_balancer_fqdn` | [https_health_check.inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-0231002232301122-2313322221322130-3111011013013011-3111001210302013-2321000202003130-1010023113001003-0020223110003120-2110011212221310) |
| `https_health_check.receive` | [https_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-3322021133213123-2113001112133010-2101132301133220-1221201211123330-2302203032132012-3310132203033221-3111013131202322-1220003013130002) |
| `https_health_check.send` | [https_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-2333101111010201-2022201033120131-2033011300213301-2133230201011221-0033232032300312-3320130233310211-0320032301120330-3110210012101121) |
| `https_health_check.virtual_host` | [https_health_check.virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-0113222231322330-0230331213301000-3302113003023313-1011322033000131-1322013323303333-0023100212113202-3133223100032203-3220201011312101) |
| `icmp_health_check` | [icmp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1310132303122201-1313023323230202-3112133311310331-3333021133321210-1231300020202122-3333033313020211-0031030113232033-2031112223031213) |
| `id` | [ID](resources--dns_lb_health_check--reference--group-001.md#canonical-0032121332010132-0130123330021333-0310201033102311-0130020320112201-2111330312031220-1110200113103013-2122213023323320-3330032321101023) |
| `labels` | [labels](resources--dns_lb_health_check--reference--group-001.md#canonical-1210321202303301-1331200321120102-3221233213312211-1310021110030223-1300101100213012-0323030000021112-2222132022212130-1331220300213231) |
| `name` | [name](resources--dns_lb_health_check--reference--group-001.md#canonical-1132300312011302-2201300011333131-2021210313302012-0120303332111212-0130203331323021-2132123003022330-1312331211102201-3023102113210223) |
| `namespace` | [namespace](resources--dns_lb_health_check--reference--group-001.md#canonical-0000022000202101-0331300212232011-2202013220031000-3201210211302221-3032121231212322-1211201101020101-0021301331202131-2032320103031333) |
| `tcp_health_check` | [tcp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1233320320313201-3030311113023011-3332212201011200-2121310112321230-0331302000032112-0112022020211300-2213212322021230-3221111113122022) |
| `tcp_health_check.health_check_port` | [tcp_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-1123111132201300-0221120111211331-3032230031032310-3131133113311123-3102231321220203-0102133000011220-0200130003021031-1103202131233333) |
| `tcp_health_check.health_check_secondary_port` | [tcp_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-0330330003231011-0223001120230110-2203310212303333-1300220311220320-3002300111311200-1310303010311302-3223323000021122-3333130322223102) |
| `tcp_health_check.receive` | [tcp_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-2333312110022220-2002203200210220-2320100130312021-1102311010030102-2300220201212131-1302000213022022-0133112331100011-3220122230112200) |
| `tcp_health_check.send` | [tcp_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-3033323010103030-3020210032010233-0302101222032013-1303212221233321-1000132030300311-2130121301330110-0023202111300211-3100302231201203) |
| `tcp_hex_health_check` | [tcp_hex_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-0003022320003130-1220001120201032-2130132022331110-3333202122002322-3220010301000031-3322230113013213-3112332023333031-1103131213310300) |
| `tcp_hex_health_check.health_check_port` | [tcp_hex_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-0023231131101323-1223223321001001-2310232210023313-2101123023002000-3013123321223332-3101110031030323-1132102121310303-3213310111023213) |
| `tcp_hex_health_check.health_check_secondary_port` | [tcp_hex_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-2131033322323111-3211121120022322-1300131201202302-0231033210211031-0100233002320031-1131122312320003-3122031311020331-2313323123211000) |
| `tcp_hex_health_check.receive` | [tcp_hex_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-1332220331112122-3222111003102332-0132131212203203-3201222231000312-1113331101110312-0323100022002101-1301011103102210-2123130003213101) |
| `tcp_hex_health_check.send` | [tcp_hex_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-3011010212330103-3331031113303010-2102001333033131-0011203202032121-2113313230203002-2221012233313331-3010101121003201-1123221023212123) |
| `timeouts` | [timeouts](resources--dns_lb_health_check--reference--group-001.md#canonical-2210202011320332-3213212202233330-2002202231332130-2323123021321301-0331021002332311-1203132320000312-1303320232123032-2203303120233311) |
| `timeouts.create` | [timeouts.create](resources--dns_lb_health_check--reference--group-001.md#canonical-0002120002300000-0121301012002230-3212332322222311-1121123312032210-0231310111003300-3220323110001331-3101000220331103-1312021110120000) |
| `timeouts.delete` | [timeouts.delete](resources--dns_lb_health_check--reference--group-001.md#canonical-0002010333021122-2002012123210011-2112233132011011-1202203113113222-1313301223112031-0223021110020131-0312230331010132-0023233121330301) |
| `timeouts.read` | [timeouts.read](resources--dns_lb_health_check--reference--group-001.md#canonical-0110222121031101-2301233030112123-0130033031332311-0211033120332320-3330100003032020-1201130333321333-2101113003130022-0313111000331201) |
| `timeouts.update` | [timeouts.update](resources--dns_lb_health_check--reference--group-001.md#canonical-1023032330010121-3311212131131230-2012030020132100-0301100032120013-1021133220012233-2112213002010203-1122320322333100-3110123220032130) |
| `udp_health_check` | [udp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-0130331031001020-2302301112100011-0300233312120111-1011030230021200-0312132103220031-2103111330003230-2320013022220023-2230101113020303) |
| `udp_health_check.health_check_port` | [udp_health_check.health_check_port](resources--dns_lb_health_check--reference--group-001.md#canonical-0011301320333121-3123032013203003-1020223011211311-1023302231011113-2232121213111103-0120332132013120-2022310223000001-1333231023321101) |
| `udp_health_check.health_check_secondary_port` | [udp_health_check.health_check_secondary_port](resources--dns_lb_health_check--reference--group-001.md#canonical-2232200003221112-2130010022311221-0111320100022033-3132010232131011-1103100012202100-3311112320202021-3000023203302220-3211223310020333) |
| `udp_health_check.receive` | [udp_health_check.receive](resources--dns_lb_health_check--reference--group-001.md#canonical-0203103211321100-2031030101210330-2001131231322022-1323201021312023-0022301023001320-2220000033021302-3221213213233131-1031102201301220) |
| `udp_health_check.send` | [udp_health_check.send](resources--dns_lb_health_check--reference--group-001.md#canonical-0313103223111230-2311023303332022-3232220310232020-0111303333310232-3103030232223100-1030123231321133-1333320112031002-3033300021032312) |

<a id="canonical-3123301133033020-3001312112033103-2023010020133212-1212321113313010-2303023100323100-0031323300231202-0032201133132201-2022313223021312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- http_health_check

<a id="canonical-1221130202322331-2130221223120123-1211123222201231-2233103123111303-2032112002130020-2220011132023230-1313232301122121-1231133313310202"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: http\_health\_check, https\_health\_check, icmp\_health\_check, tcp\_health\_check,
tcp\_hex\_health\_check, udp\_health\_check\] Configuration parameter for http health check.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "inherit_load_balancer_fqdn"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "virtual_host"),
  validators.ConflictingObjectAttributes("inherit_load_balancer_fqdn",
    "virtual_host")}
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
  "x-ves-oneof-field-virtual_host_choice": "[\"disable_virtual_host\",\"inherit_load_balancer_fqdn\",\"virtual_host\"]"
}
```

OneOf alternatives in this subsection:

- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1221130202322331-2130221223120123-1211123222201231-2233103123111303-2032112002130020-2220011132023230-1313232301122121-1231133313310202)
- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-2121323320010013-3120203001021122-1001323231220330-0231100022101023-1231011100023131-0120121102300101-3212131332013100-2220103210020310)
- [icmp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1310132303122201-1313023323230202-3112133311310331-3333021133321210-1231300020202122-3333033313020211-0031030113232033-2031112223031213)
- [tcp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1233320320313201-3030311113023011-3332212201011200-2121310112321230-0331302000032112-0112022020211300-2213212322021230-3221111113122022)
- [tcp_hex_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-0003022320003130-1220001120201032-2130132022331110-3333202122002322-3220010301000031-3322230113013213-3112332023333031-1103131213310300)
- [udp_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-0130331031001020-2302301112100011-0300233312120111-1011030230021200-0312132103220031-2103111330003230-2320013022220023-2230101113020303)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020211222012332-2223131033112110-1333310232131133-3003012212002233-1112211100101021-0122002313233221-2223231203121223-3001212200313121"></a>

### Direct properties for `http_health_check`

- [disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-1220313032231202-1002300301213000-1103213221323331-0322331110013232-0011321232220222-1112302020313320-0321110201220330-1101231220110331): complete subsection reference.

<a id="canonical-2031220000233033-2100002331130132-1322223211300203-2013330102020303-0003032302330210-2100221031212211-1113131202133300-3313011323320133"></a>

<a id="canonical-0302311300132032-3011122222203003-1221231321330000-0011022222230213-0300231000100323-1003132031122002-2020112300221222-3323211210022021"></a>

#### `http_health_check.health_check_port` property

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

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
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2213302303303101-1302302300331110-0212001011210100-1330200130300312-0322131010213030-0200230220110221-2303100323301023-2332003113223231"></a>

<a id="canonical-0231301010213302-3201223300110233-3101130303210120-0013001110321100-1230011001013313-3311132130110123-3100020200212322-1011113323120332"></a>

#### `http_health_check.health_check_secondary_port` property

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-3320022032300331-0121013102111103-1310323122010210-3320302330133012-1311211023011003-3123103321230002-0301303323301013-0021011000012232): complete subsection reference.

<a id="canonical-1112003222332101-0221102131312230-1111233130331021-3213013012120323-0203313223111000-3032001121003122-0112222031002321-3231202031110000"></a>

<a id="canonical-0120031301321011-0000200023121032-2013322132031123-1333121222201313-2323201233231002-0012331222033311-1332112333031322-1201030233100303"></a>

#### `http_health_check.receive` property

Type: `"string"`. Optional.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2112330130302033-2322102211022112-1033233123332031-1321111223032013-2013212120110200-1300103010222232-0121221202222200-3233032100021023"></a>

<a id="canonical-3003202133233201-0110312121102211-0223311223033020-3303102101100311-0231133222303233-0002320103200100-0303033332132220-1030223210300303"></a>

#### `http_health_check.send` property

Type: `"string"`. Optional.

Send String. HTTP payload to send to the target.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0320322010003211-2023133311310302-0103200100313130-2220021301023313-0303333112301302-1322100123210303-0321022020121303-0021230211212313"></a>

<a id="canonical-1212330103012220-0003111333113310-0201130020110320-0121011301102113-1023001231123211-2321031231030121-3320201313222112-1023213120333302"></a>

#### `http_health_check.virtual_host` property

Type: `"string"`. Optional.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-1220313032231202-1002300301213000-1103213221323331-0322331110013232-0011321232220222-1112302020313320-0321110201220330-1101231220110331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_health_check.disable_virtual_host` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-3123301133033020-3001312112033103-2023010020133212-1212321113313010-2303023100323100-0031323300231202-0032201133132201-2022313223021312)
- http_health_check.disable_virtual_host

<a id="canonical-3130301010010113-0302111222100311-1021033230212102-2231301320120313-0321030130202233-1123003131233233-3130032002311310-3001222212331222"></a>

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
disable_virtual_host = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320022032300331-0121013102111103-1310323122010210-3320302330133012-1311211023011003-3123103321230002-0301303323301013-0021011000012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_health_check.inherit_load_balancer_fqdn` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- [http_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-3123301133033020-3001312112033103-2023010020133212-1212321113313010-2303023100323100-0031323300231202-0032201133132201-2022313223021312)
- http_health_check.inherit_load_balancer_fqdn

<a id="canonical-3220231232213020-1223011031021132-2011020223100333-2111330122132121-0313220120132111-1311203320333031-3021133221000331-2230333120212103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit load balancer fqdn.

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
inherit_load_balancer_fqdn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013121033333001-2311212031003021-3122120000133312-0311032031230022-0031330011111330-3130133200012211-0102332022220133-2100120110113333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- https_health_check

<a id="canonical-2121323320010013-3120203001021122-1001323231220330-0231100022101023-1231011100023131-0120121102300101-3212131332013100-2220103210020310"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for https health check.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "inherit_load_balancer_fqdn"),
  validators.ConflictingObjectAttributes("disable_virtual_host",
    "virtual_host"),
  validators.ConflictingObjectAttributes("inherit_load_balancer_fqdn",
    "virtual_host")}
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
  "x-ves-oneof-field-virtual_host_choice": "[\"disable_virtual_host\",\"inherit_load_balancer_fqdn\",\"virtual_host\"]"
}
```

Terraform syntax:

```terraform
https_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013001131301233-3313312110021012-2131330311130132-2002012310022032-2211000302203321-3120301012132202-2203320221220122-0101312222212303"></a>

### Direct properties for `https_health_check`

- [disable_virtual_host](resources--dns_lb_health_check--reference--group-001.md#canonical-2331031220121003-3232220000000333-0002202130020323-3302122321300310-3231231201103032-0012230232011223-1221032012103211-1020222102210202): complete subsection reference.

<a id="canonical-2213211221323112-0112313202121011-3233100313220310-2030212033022212-3230022111003031-0211131030311312-0023232030010221-3023312002321212"></a>

<a id="canonical-0112123112312102-0301133211203020-1022230321031002-1331021021021020-1102112011313220-0020010321302310-2232001301312121-2310101201131201"></a>

#### `https_health_check.health_check_port` property

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

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
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3020211013120202-3132303110011022-2020300302120223-1331000123102123-0001132130213112-2322223121310030-1031233111312300-3322323311000303"></a>

<a id="canonical-0312301201113311-3313202010102013-1133133002320120-0031300123302203-2030313132322102-3213003212303001-2322212012133320-0202333132220123"></a>

#### `https_health_check.health_check_secondary_port` property

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [inherit_load_balancer_fqdn](resources--dns_lb_health_check--reference--group-001.md#canonical-1012102332121320-2033103023300022-1201003031022222-2330312202111103-3232232222321322-3100330300321323-0030330132132203-2333001313300111): complete subsection reference.

<a id="canonical-3322021133213123-2113001112133010-2101132301133220-1221201211123330-2302203032132012-3310132203033221-3111013131202322-1220003013130002"></a>

<a id="canonical-1131301232232133-3011100000003001-1320033220213121-2333320110133213-0302021202002102-1001001001012200-1330122210232133-1033111122000221"></a>

#### `https_health_check.receive` property

Type: `"string"`. Optional.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2333101111010201-2022201033120131-2033011300213301-2133230201011221-0033232032300312-3320130233310211-0320032301120330-3110210012101121"></a>

<a id="canonical-0332001132021321-1301333132010122-0030022022220003-0202311301112020-2233200102132331-1201032323321322-0031101223313301-2001231011313231"></a>

#### `https_health_check.send` property

Type: `"string"`. Optional.

Send String. HTTP payload to send to the target.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0113222231322330-0230331213301000-3302113003023313-1011322033000131-1322013323303333-0023100212113202-3133223100032203-3220201011312101"></a>

<a id="canonical-2203002221313212-2133221212320231-3101111203233311-1232000330023230-2031130310001101-3010011103223321-3233000101021210-3132012321021232"></a>

#### `https_health_check.virtual_host` property

Type: `"string"`. Optional.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-2331031220121003-3232220000000333-0002202130020323-3302122321300310-3231231201103032-0012230232011223-1221032012103211-1020222102210202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_health_check.disable_virtual_host` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1013121033333001-2311212031003021-3122120000133312-0311032031230022-0031330011111330-3130133200012211-0102332022220133-2100120110113333)
- https_health_check.disable_virtual_host

<a id="canonical-0201032311031110-0210300001030312-3020001332301012-3203110110201103-2331323112001223-0223221031301003-3023013311113120-3011100113322123"></a>

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
disable_virtual_host = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012102332121320-2033103023300022-1201003031022222-2330312202111103-3232232222321322-3100330300321323-0030330132132203-2333001313300111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_health_check.inherit_load_balancer_fqdn` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- [https_health_check](resources--dns_lb_health_check--reference--group-001.md#canonical-1013121033333001-2311212031003021-3122120000133312-0311032031230022-0031330011111330-3130133200012211-0102332022220133-2100120110113333)
- https_health_check.inherit_load_balancer_fqdn

<a id="canonical-0231002232301122-2313322221322130-3111011013013011-3111001210302013-2321000202003130-1010023113001003-0020223110003120-2110011212221310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit load balancer fqdn.

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
inherit_load_balancer_fqdn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022113030000323-2330222301031023-0120100012200301-2333110032132213-0003230321003203-2023312222130303-1322131221322010-0133303003303103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `icmp_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- icmp_health_check

<a id="canonical-1310132303122201-1313023323230202-3112133311310331-3333021133321210-1231300020202122-3333033313020211-0031030113232033-2031112223031213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for icmp health check.

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
icmp_health_check = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202311232132230-0212321030133021-3022310202313131-3312222212221302-0133231310311030-0303013212230320-1231023213133231-0200233332322222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tcp_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- tcp_health_check

<a id="canonical-1233320320313201-3030311113023011-3332212201011200-2121310112321230-0331302000032112-0112022020211300-2213212322021230-3221111113122022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp health check.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port")}
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200021023233033-0213203201111231-3112203311223033-0312312013232110-0022002301002211-1321232033110300-0220020033313330-1323300132311302"></a>

### Direct properties for `tcp_health_check`

<a id="canonical-1123111132201300-0221120111211331-3032230031032310-3131133113311123-3102231321220203-0102133000011220-0200130003021031-1103202131233333"></a>

#### `tcp_health_check.health_check_port` property

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

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
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0330330003231011-0223001120230110-2203310212303333-1300220311220320-3002300111311200-1310303010311302-3223323000021122-3333130322223102"></a>

<a id="canonical-1130230100123231-1301311031233031-3011231003222322-1111201121332332-3322001330111032-0331302020210231-2323130020001303-1211223213210231"></a>

#### `tcp_health_check.health_check_secondary_port` property

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2333312110022220-2002203200210220-2320100130312021-1102311010030102-2300220201212131-1302000213022022-0133112331100011-3220122230112200"></a>

<a id="canonical-2003212321220331-2031311231232020-2030220021120232-0012322022313130-2223313311111312-0003120113023003-3003010122333201-1211003120100221"></a>

#### `tcp_health_check.receive` property

Type: `"string"`. Optional.

Regular expression used to match against the response to the monitor's request. Mark node up upon
receipt of a successful regular expression match. Uses re2 regular expression syntax.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3033323010103030-3020210032010233-0302101222032013-1303212221233321-1000132030300311-2130121301330110-0023202111300211-3100302231201203"></a>

<a id="canonical-2330303133202020-2102223231113122-2322103311030130-1112031223300320-2130311300111331-3121113223003211-3232302330233003-2010010133301021"></a>

#### `tcp_health_check.send` property

Type: `"string"`. Optional.

Send this string to target (default empty. When send and receive are both empty, monitor just tests
3WHS).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-3310122233311003-0331022321313010-3222232300211013-0210133110202032-0302023320232031-1333300003101322-2030330013112003-2012332003332031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tcp_hex_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- tcp_hex_health_check

<a id="canonical-0003022320003130-1220001120201032-2130132022331110-3333202122002322-3220010301000031-3322230113013213-3112332023333031-1103131213310300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp hex health check.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port")}
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
tcp_hex_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010102012000000-2330212111321013-3203122200330123-1221312102222302-1203213212003213-1002230131301132-0313013231213122-3111210200121333"></a>

### Direct properties for `tcp_hex_health_check`

<a id="canonical-0023231131101323-1223223321001001-2310232210023313-2101123023002000-3013123321223332-3101110031030323-1132102121310303-3213310111023213"></a>

#### `tcp_hex_health_check.health_check_port` property

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

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
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2131033322323111-3211121120022322-1300131201202302-0231033210211031-0100233002320031-1131122312320003-3122031311020331-2313323123211000"></a>

<a id="canonical-1233031201131212-2313133303000031-3132330101133002-1032010201031222-1313331202021032-0000310222100202-3213320113132132-0202120302033100"></a>

#### `tcp_hex_health_check.health_check_secondary_port` property

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1332220331112122-3222111003102332-0132131212203203-3201222231000312-1113331101110312-0323100022002101-1301011103102210-2123130003213101"></a>

<a id="canonical-0003223010330220-2223301201023202-0112110330311133-3202130111133332-1303112021331120-1312101110301031-2221100030321330-0312003002033303"></a>

#### `tcp_hex_health_check.receive` property

Type: `"string"`. Optional.

Hex encoded raw bytes expected in the response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-3011010212330103-3331031113303010-2102001333033131-0011203202032121-2113313230203002-2221012233313331-3010101121003201-1123221023212123"></a>

<a id="canonical-1320033212233221-1332121102232313-3031021132131102-1202101032001003-1311120312011132-2120102001011121-2100021102211130-1330321223012333"></a>

#### `tcp_hex_health_check.send` property

Type: `"string"`. Optional.

Hex encoded raw bytes sent in the request. Empty payloads imply a connect-only health check.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-3100003210332223-1023023222311302-1120310230130002-1321123111113000-2033201021003301-2300302321021312-2131002323220212-1303032312322311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- timeouts

<a id="canonical-2210202011320332-3213212202233330-2002202231332130-2323123021321301-0331021002332311-1203132320000312-1303320232123032-2203303120233311"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212100010013020-0313102123030321-3203122001230130-2133220010122321-0201003222232130-2122331101132013-2100120133112020-3220003333332303"></a>

### Direct properties for `timeouts`

<a id="canonical-0002120002300000-0121301012002230-3212332322222311-1121123312032210-0231310111003300-3220323110001331-3101000220331103-1312021110120000"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0002010333021122-2002012123210011-2112233132011011-1202203113113222-1313301223112031-0223021110020131-0312230331010132-0023233121330301"></a>

<a id="canonical-1002232210023222-1312001232233132-1333312020233110-2333221320203220-2120012030312121-0000213100100131-0020023301103230-3002313021313300"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0110222121031101-2301233030112123-0130033031332311-0211033120332320-3330100003032020-1201130333321333-2101113003130022-0313111000331201"></a>

<a id="canonical-0303111300312110-3330123230223332-1222233113212113-2233133211033132-3103133220332213-3033230013021230-0230212213120001-0031200133120010"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1023032330010121-3311212131131230-2012030020132100-0301100032120013-1021133220012233-2112213002010203-1122320322333100-3110123220032130"></a>

<a id="canonical-2332023133010232-0113223000113130-3133200323300021-3300211131310332-2310110310020320-3312220132321321-1011323232121131-1010100000013131"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1033111102330311-0210031231001012-0311000130202221-2102102013102323-2102322000022300-3023201333202310-3212111232132212-1013112131001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `udp_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Property reference](resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- udp_health_check

<a id="canonical-0130331031001020-2302301112100011-0300233312120111-1011030230021200-0312132103220031-2103111330003230-2320013022220023-2230101113020303"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for udp health check.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port",
    "receive",
    "send")}
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
udp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131123032231010-0333033132002101-2332301201011122-0222032223320130-2322221220032230-3211333212001331-2331002121300103-0002210303201313"></a>

### Direct properties for `udp_health_check`

<a id="canonical-0011301320333121-3123032013203003-1020223011211311-1023302231011113-2232121213111103-0120332132013120-2022310223000001-1333231023321101"></a>

#### `udp_health_check.health_check_port` property

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

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
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2232200003221112-2130010022311221-0111320100022033-3132010232131011-1103100012202100-3311112320202021-3000023203302220-3211223310020333"></a>

<a id="canonical-1000131211313103-3021033232131030-0121012323331113-3130023330221330-3101123330233123-2303031012332130-2111322203022230-2033313121331233"></a>

#### `udp_health_check.health_check_secondary_port` property

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0203103211321100-2031030101210330-2001131231322022-1323201021312023-0022301023001320-2220000033021302-3221213213233131-1031102201301220"></a>

<a id="canonical-0333001312000100-3332301113322322-2232100002300333-0020301030322312-3330220022130112-2002122111200120-3112333302020132-1320313213323220"></a>

#### `udp_health_check.receive` property

Type: `"string"`. Optional.

UDP response to be matched. It can be a regular expression.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0313103223111230-2311023303332022-3232220310232020-0111303333310232-3103030232223100-1030123231321133-1333320112031002-3033300021032312"></a>

<a id="canonical-3102032321222020-1100003122002113-2311233030013221-2003020120023222-0213130133021332-1301202120321030-0100233130003000-3232330011300203"></a>

#### `udp_health_check.send` property

Type: `"string"`. Optional.

Send String. UDP payload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```
