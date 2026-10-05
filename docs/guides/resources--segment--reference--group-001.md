---
page_title: "xcsh_segment reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment reference."
---

# xcsh_segment reference

<a id="canonical-0333310231313321-0120303030320002-3001033000112301-0230202002331312-0213020300113131-0011002021012330-3202203300010331-3312231333130010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313013221232010-3020010031010333-0302123123031200-0203330112122031-0311001123111212-0110212012000033-2132313110132302-1113213322221230"></a>

## Property reference — Property reference / 121033312131 / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100)
- Property reference

<a id="canonical-2233313111233210-0310132130121020-1120000330021203-0012020333210223-2210202100111203-1232332331311330-0023313031211203-1032221311223032"></a>

## Direct properties — Property reference / 121033312131 / 3

<a id="canonical-1021213212200220-0000013320213102-1102313311210103-0133321020312110-2023103103223011-2032311132321122-2000102310030022-3111312113002132"></a>

<a id="canonical-3230330032323302-3331032113111031-1221023100333110-1101323323003120-2312221202120302-2112330000332012-0120020330313312-0223303322022231"></a>

## annotations property — Property reference / 121033312131 / 4

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

<a id="canonical-0311331110130212-3320120333123233-0113301333303122-0313022303320321-1321320031311202-1100023021233310-3133303221300001-1330122130131321"></a>

<a id="canonical-2110331332232010-3030212200022023-1003123103230110-2232120120120322-3200101303210032-2012101133320202-0022231033212120-3020322112113320"></a>

## description property — Property reference / 121033312131 / 5

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

- [disable_spec](resources--segment--reference--group-001.md#canonical-2230301223223231-1210200033203123-2012221332020222-0321210221322101-1031033232103110-3101300111301033-1120310213133103-2102310210032001): complete subsection reference.

- [enable](resources--segment--reference--group-001.md#canonical-3100223213111010-2103312210220322-2233203002011122-3223023302011331-2333321113330202-2322330231112103-1333211031112123-0331313121113330): complete subsection reference.

<a id="canonical-0133131112232021-2310210331210201-2303301113211021-2100102212100232-3132030233002000-3023121022320032-3323222320110212-2300003031321013"></a>

<a id="canonical-3100022332020331-1011302023102221-3223221332032021-1321312101232030-2321201203221111-0001012013311113-2221010211203011-0003313131103321"></a>

## ID property — Property reference / 121033312131 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2211210000001200-2000220302322123-2321311300313210-2331220130130031-2001313212201121-1120032113133220-0023013310301123-3103302031210120"></a>

<a id="canonical-3031302002022112-1213233131312303-2033203011331212-2321120322212222-1323320321223022-1303220330130231-0110303032223220-1321021313311311"></a>

## labels property — Property reference / 121033312131 / 7

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

<a id="canonical-1312123321003010-1133131111332001-1230212310111231-0102222200013200-2132101003313320-1311202303032123-0132132131203100-1302323213132223"></a>

<a id="canonical-0021011011213111-3301000122311112-3232321322333121-3201333302231200-1322030010010303-2110323231320300-0000311200211212-3211022001013323"></a>

## name property — Property reference / 121033312131 / 8

Type: `"string"`. Required.

Name of the Segment. Must be unique within the namespace.

Upstream description:

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

<a id="canonical-2002101313133210-2132002120322023-0133212303320103-2011232323220330-0332132012101213-2013310233132332-2310013122021023-0301132203023032"></a>

<a id="canonical-1211202022212113-0220011303111300-1001123111210013-1101012133130213-1323211121021120-1133101310123230-0120303033022022-3313200222003203"></a>

## namespace property — Property reference / 121033312131 / 9

Type: `"string"`. Optional, Computed.

Namespace for the Segment. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

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

- [timeouts](resources--segment--reference--group-001.md#canonical-1230030032300111-2111112323333230-2120202033322032-2203231122103121-0322102231213023-1200132100132013-3132003133113311-0112102302012032): complete subsection reference.

<a id="canonical-0233031020330002-2310122103010300-1033311200130103-2102203201222333-3220322013222302-3133322230000213-3113222110212010-1220133020012302"></a>

## All schema paths — Property reference / 121033312131 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--segment--reference--group-001.md#canonical-1021213212200220-0000013320213102-1102313311210103-0133321020312110-2023103103223011-2032311132321122-2000102310030022-3111312113002132) |
| `description` | [description](resources--segment--reference--group-001.md#canonical-0311331110130212-3320120333123233-0113301333303122-0313022303320321-1321320031311202-1100023021233310-3133303221300001-1330122130131321) |
| `disable_spec` | [disable_spec](resources--segment--reference--group-001.md#canonical-0312203231301201-3001332310333231-2330310111020233-2021300210013203-1301233301032313-0323232313100323-1003230333201011-0213000232310130) |
| `enable` | [enable](resources--segment--reference--group-001.md#canonical-3220231133021320-0021033303012031-3132230332002333-3331121321213023-3330230320101100-1102300221301213-3111101321221030-1300022000220021) |
| `id` | [ID](resources--segment--reference--group-001.md#canonical-0133131112232021-2310210331210201-2303301113211021-2100102212100232-3132030233002000-3023121022320032-3323222320110212-2300003031321013) |
| `labels` | [labels](resources--segment--reference--group-001.md#canonical-2211210000001200-2000220302322123-2321311300313210-2331220130130031-2001313212201121-1120032113133220-0023013310301123-3103302031210120) |
| `name` | [name](resources--segment--reference--group-001.md#canonical-1312123321003010-1133131111332001-1230212310111231-0102222200013200-2132101003313320-1311202303032123-0132132131203100-1302323213132223) |
| `namespace` | [namespace](resources--segment--reference--group-001.md#canonical-2002101313133210-2132002120322023-0133212303320103-2011232323220330-0332132012101213-2013310233132332-2310013122021023-0301132203023032) |
| `timeouts` | [timeouts](resources--segment--reference--group-001.md#canonical-2320201103103120-0032102310211302-2120122312112033-0213021022312131-0012120331101210-0132201222321202-3031111333101031-1010233310001310) |
| `timeouts.create` | [timeouts.create](resources--segment--reference--group-001.md#canonical-2331321301232330-1201123011022311-0033033310301330-2211013223331301-0232223023323330-2010223222222000-2222110220111231-0123100222121211) |
| `timeouts.delete` | [timeouts.delete](resources--segment--reference--group-001.md#canonical-3121003032322211-2200232221211131-2102133311231300-3232311201111022-0223011201001310-1030232333123232-2222302121231303-0133032122031031) |
| `timeouts.read` | [timeouts.read](resources--segment--reference--group-001.md#canonical-2110002210221220-2333112032203223-2032113203332101-1233102130212230-2213130330223030-3121100211313013-1033121031332031-2310313101033132) |
| `timeouts.update` | [timeouts.update](resources--segment--reference--group-001.md#canonical-3222320023113312-1020212322203302-2322202301010212-1100300231110022-0133032111011103-1330220220211021-0210222302223310-3100030102011222) |

<a id="canonical-2222232231203120-1230310003113111-0013032002303100-3000320130012303-1302331111323021-3321032022223333-3131123002321232-3313031220011211"></a>

## Next pages — Property reference / 121033312131 / 11

- [disable_spec](resources--segment--reference--group-001.md#canonical-2230301223223231-1210200033203123-2012221332020222-0321210221322101-1031033232103110-3101300111301033-1120310213133103-2102310210032001)
- [enable](resources--segment--reference--group-001.md#canonical-3100223213111010-2103312210220322-2233203002011122-3223023302011331-2333321113330202-2322330231112103-1333211031112123-0331313121113330)
- [timeouts](resources--segment--reference--group-001.md#canonical-1230030032300111-2111112323333230-2120202033322032-2203231122103121-0322102231213023-1200132100132013-3132003133113311-0112102302012032)
- [xcsh_segment](../resources/segment.md#canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100)

<a id="canonical-2230301223223231-1210200033203123-2012221332020222-0321210221322101-1031033232103110-3101300111301033-1120310213133103-2102310210032001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210333031212331-2011101310330210-2030102010121020-2001131113203003-1312331002123202-2330023221233000-3201231302122333-3301131031321223"></a>

## disable_spec — disable_spec / 032032030233 / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100)
- [Property reference](resources--segment--reference--group-001.md#canonical-0333310231313321-0120303030320002-3001033000112301-0230202002331312-0213020300113131-0011002021012330-3202203300010331-3312231333130010)
- disable_spec

<a id="canonical-0312203231301201-3001332310333231-2330310111020233-2021300210013203-1301233301032313-0323232313100323-1003230333201011-0213000232310130"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable, enable\] Enable this option

OneOf alternatives in this subsection:

- `disable`
- [enable](resources--segment--reference--group-001.md#canonical-3220231133021320-0021033303012031-3132230332002333-3331121321213023-3330230320101100-1102300221301213-3111101321221030-1300022000220021)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-0120021232312222-0221020213000100-2031123322002102-0032021200223233-1201022033102231-1130210313313023-0110221231322223-3223333130213302"></a>

## Direct properties — disable_spec / 032032030233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031003333030213-2211020102311011-3013310000330003-2103211022311331-2230203122100302-3022210313311221-0132333033330311-0123313110332213"></a>

## Next pages — disable_spec / 032032030233 / 4

- [Property reference](resources--segment--reference--group-001.md#canonical-0333310231313321-0120303030320002-3001033000112301-0230202002331312-0213020300113131-0011002021012330-3202203300010331-3312231333130010)
- [xcsh_segment](../resources/segment.md#canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100)

<a id="canonical-3100223213111010-2103312210220322-2233203002011122-3223023302011331-2333321113330202-2322330231112103-1333211031112123-0331313121113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321120033003102-1031313202030323-3121021203202232-2101132211322110-0313200323323312-1303313333000221-1011213221121313-2300113100123322"></a>

## enable — enable / 332233021111 / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100)
- [Property reference](resources--segment--reference--group-001.md#canonical-0333310231313321-0120303030320002-3001033000112301-0230202002331312-0213020300113131-0011002021012330-3202203300010331-3312231333130010)
- enable

<a id="canonical-3220231133021320-0021033303012031-3132230332002333-3331121321213023-3330230320101100-1102300221301213-3111101321221030-1300022000220021"></a>

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
enable = {}
```

<a id="canonical-1301301210001223-3130200230110211-3021333110312202-0033132323320023-3210222331332330-3210121012002010-1012113213110033-2211203201303030"></a>

## Direct properties — enable / 332233021111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111012101312002-2130313111230031-2303213100221333-0102322133122103-2102200221323220-3320323102103221-2202202332221020-2030223223031322"></a>

## Next pages — enable / 332233021111 / 4

- [Property reference](resources--segment--reference--group-001.md#canonical-0333310231313321-0120303030320002-3001033000112301-0230202002331312-0213020300113131-0011002021012330-3202203300010331-3312231333130010)
- [xcsh_segment](../resources/segment.md#canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100)

<a id="canonical-1230030032300111-2111112323333230-2120202033322032-2203231122103121-0322102231213023-1200132100132013-3132003133113311-0112102302012032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220212201310203-0021313301130332-3101032333030211-1201112030102130-0302023203211032-1112203121030123-0222333233103123-3010120332103310"></a>

## timeouts — timeouts / 313330031231 / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100)
- [Property reference](resources--segment--reference--group-001.md#canonical-0333310231313321-0120303030320002-3001033000112301-0230202002331312-0213020300113131-0011002021012330-3202203300010331-3312231333130010)
- timeouts

<a id="canonical-2320201103103120-0032102310211302-2120122312112033-0213021022312131-0012120331101210-0132201222321202-3031111333101031-1010233310001310"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100110303110003-2102200101320021-1002031233201002-3023322310112211-0021300322102133-2221323212120032-3013232230321230-3313022030231132"></a>

## Direct properties — timeouts / 313330031231 / 3

<a id="canonical-2331321301232330-1201123011022311-0033033310301330-2211013223331301-0232223023323330-2010223222222000-2222110220111231-0123100222121211"></a>

<a id="canonical-1201112000031232-1210023313300130-0321103031133020-2030301210113310-2203313010033003-2112111333102011-0133000012302122-1010303133331221"></a>

## create property — timeouts / 313330031231 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3121003032322211-2200232221211131-2102133311231300-3232311201111022-0223011201001310-1030232333123232-2222302121231303-0133032122031031"></a>

<a id="canonical-3231032200203233-3110222310010303-3322123021002113-2210233122122233-1332212023323032-0122322202010223-0121131132100110-1320230230000323"></a>

## delete property — timeouts / 313330031231 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2110002210221220-2333112032203223-2032113203332101-1233102130212230-2213130330223030-3121100211313013-1033121031332031-2310313101033132"></a>

<a id="canonical-3200112203323201-2203233120232223-0203330022121310-2203020022302010-0122220302221013-1121023101033310-3120322221330201-3321023321020031"></a>

## read property — timeouts / 313330031231 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3222320023113312-1020212322203302-2322202301010212-1100300231110022-0133032111011103-1330220220211021-0210222302223310-3100030102011222"></a>

<a id="canonical-1030210323120130-2203220220312310-1101031001233020-1013203221232311-1303002322210303-3333120310102321-2012132011101133-3032130203210333"></a>

## update property — timeouts / 313330031231 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2332122323102133-3331331201310322-1230110201201330-1112223122132200-0002332111230020-3221322022010212-3321221010232112-1021203313032310"></a>

## Next pages — timeouts / 313330031231 / 8

- [Property reference](resources--segment--reference--group-001.md#canonical-0333310231313321-0120303030320002-3001033000112301-0230202002331312-0213020300113131-0011002021012330-3202203300010331-3312231333130010)
- [xcsh_segment](../resources/segment.md#canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100)
