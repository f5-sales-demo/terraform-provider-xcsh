---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333010221302231-3230322113030200-3111211232322330-2333210303122221-1003231120123032-2231132231132121-0013100123030202-0000023123213002"></a>

## Property reference — Property reference / 010010133120 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- Property reference

<a id="canonical-0303332021300221-3033132030311321-0033222331123230-0120233233311011-1222102111122122-3300331221211022-2110001000323131-1131111133221212"></a>

## Direct properties — Property reference / 010010133120 / 3

<a id="canonical-3122231302100131-1000312130113302-3113131311130312-3332312011312332-3220313103122133-2031112100103000-2312010311003132-1101221021210020"></a>

<a id="canonical-2332111220212220-3312330031033133-0222233002220220-2200120132223302-0201231331303032-3331332210022221-3320013232003031-3222311011320000"></a>

## address property — Property reference / 010010133120 / 4

Type: `"string"`. Computed.

Site's geographical address that can be used to determine its latitude and longitude.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [admin_password](data-sources--azure_vnet_site--reference--group-003.md#canonical-0122202133003213-3212232123303332-1101131331110000-0311100100222030-1232023021332211-0211031030022210-2120232101012223-0302203221332132): complete subsection reference.

<a id="canonical-0102130211133211-1020210312033301-2020323102032102-2221110033132012-2332123202313130-3032012222002313-1322230302030111-1002000111333031"></a>

<a id="canonical-3303103011222121-2331302132102112-0231100111323231-0230033002331212-2103312222013202-0012130311212331-2230211202132120-1130012032222013"></a>

## alternate_region property — Property reference / 010010133120 / 5

Type: `"string"`. Computed.

\[OneOf: alternate\_region, Azure\_region\] Exclusive with \[Azure\_region\] Name of the Azure
region which does not support availability zones.

Upstream description:

Exclusive with \[Azure\_region\] Name of the Azure region which does not support availability zones.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

OneOf alternatives in this subsection:

- [alternate_region](data-sources--azure_vnet_site--reference--group-001.md#canonical-0102130211133211-1020210312033301-2020323102032102-2221110033132012-2332123202313130-3032012222002313-1322230302030111-1002000111333031)
- [azure_region](data-sources--azure_vnet_site--reference--group-001.md#canonical-0301030323211000-1100000122102113-3221323022012232-1223202033230013-0133120101120112-3230123010200223-3101202212320113-1210020000211103)

Select alternatives according to the provider validators above.

<a id="canonical-1332100120230132-2033332003000123-1303133303303233-1012332230222033-2012322331031300-0212301330303310-3111010213131231-3122333100123021"></a>

<a id="canonical-2013313011022011-3232130113122212-2300222021211030-2231200121300211-3032113200333332-3200102031232130-2331110311111221-2121120123120330"></a>

## annotations property — Property reference / 010010133120 / 6

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [azure_cred](data-sources--azure_vnet_site--reference--group-003.md#canonical-2302012332311232-2012022030130201-1222102013333110-2211320313113311-2111020032222322-2021303331010001-2013111030320100-2003003232233123): complete subsection reference.

<a id="canonical-0301030323211000-1100000122102113-3221323022012232-1223202033230013-0133120101120112-3230123010200223-3101202212320113-1210020000211103"></a>

<a id="canonical-3232301113112133-2200123023201100-0222330231222201-1233111110123321-3221133113223103-0230112302020122-0202303003123313-2001233233102123"></a>

## azure_region property — Property reference / 010010133120 / 7

Type: `"string"`. Computed.

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

Upstream description:

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [block_all_services](data-sources--azure_vnet_site--reference--group-003.md#canonical-0021230320320202-0233211131032112-2202000032321000-0321023033233331-1011031310201203-2333313322023113-2303233120211300-1021133023200100): complete subsection reference.

- [blocked_services](data-sources--azure_vnet_site--reference--group-003.md#canonical-1213030211010010-3021103100132013-1031200303212310-3001033333130031-1200132112312231-2210031103313020-3201133313220321-1023103332330012): complete subsection reference.

- [coordinates](data-sources--azure_vnet_site--reference--group-003.md#canonical-2221030113333311-3331222130323012-2030111312000201-2102031030222321-3231303332300301-1013303333131223-1101332223133203-2221111132133113): complete subsection reference.

- [custom_dns](data-sources--azure_vnet_site--reference--group-003.md#canonical-3202321321031121-0103011232030020-0122311111300122-3233321113211232-1221103233320201-2211132031111311-3120333101333300-3220332333000330): complete subsection reference.

- [default_blocked_services](data-sources--azure_vnet_site--reference--group-003.md#canonical-1331021023332011-2201103333230102-0230101333020023-3223021212331231-2031322100001121-3030031333021323-1331220301213003-0311030303302112): complete subsection reference.

<a id="canonical-1332330100132100-1013100222111201-1222223220122133-2112200310031021-1003331202122201-3230010233133210-0021131011101333-2031122300310320"></a>

<a id="canonical-0113012320103012-0110100231031133-2321300313233011-1002011211200320-2331330102320200-2232023023213031-2020322121031111-3222302010010020"></a>

## description property — Property reference / 010010133120 / 8

Type: `"string"`. Computed.

Description of the AzureVNETSite.

Upstream description:

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

- [disable_encryption](data-sources--azure_vnet_site--reference--group-003.md#canonical-1110201002133230-1100032102003223-3011233101122321-3231110100322222-3323310023231212-0300332202200313-1202011102011202-3213300100210113): complete subsection reference.

<a id="canonical-2211333330100123-1120032020303000-1022002101200320-0200312323022103-1133331231331232-0120131213000223-0111321002312313-1113123103301101"></a>

<a id="canonical-3313022120232302-2122332023131021-2233200300310123-3233201211030121-0203201302120101-2013031012213002-1102131333030321-3033232120200131"></a>

## disk_size property — Property reference / 010010133120 / 9

Type: `"number"`. Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB. Server applies default when omitted.

Upstream description:

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

- [enable_encryption](data-sources--azure_vnet_site--reference--group-003.md#canonical-3121211200212213-2123110103020203-0231002333103330-2110232121232103-0002212132210103-1123302030223112-1333232030331330-0331331211022113): complete subsection reference.

<a id="canonical-0323213010102003-0131113021202120-2202332301013120-2133023023022031-3301330213233221-3011010131120201-0233231132023321-3331231203333102"></a>

<a id="canonical-0221113220113021-0332121010231310-0022212113312030-0033311033210133-0122103003021031-0130201332203033-2010203001333302-3312003113223031"></a>

## ID property — Property reference / 010010133120 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112): complete subsection reference.

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111): complete subsection reference.

- [ingress_gw](data-sources--azure_vnet_site--reference--group-007.md#canonical-3131002031302201-3131200302211201-3321302320022222-0121110331203131-0213023331312031-1121133023012302-0231001231103032-0203312022210001): complete subsection reference.

- [ingress_gw_ar](data-sources--azure_vnet_site--reference--group-007.md#canonical-2130001100013102-1023322101312032-2231002233331330-3112120002103100-2201230013130331-3102220300103101-2010212103122113-2002003323320322): complete subsection reference.

- [kubernetes_upgrade_drain](data-sources--azure_vnet_site--reference--group-008.md#canonical-2131210123103103-0212213010122033-0313321123013002-1222200013311032-0020301031022033-2000020033022120-1321323012330221-3312333110101331): complete subsection reference.

<a id="canonical-0021031311301031-3131000223120302-1021001313313012-2133102331002231-2331120003123003-2300203120223123-1213202221112211-1001223003231333"></a>

<a id="canonical-2223012321212013-1121322130021022-0320132123033132-3102010220013131-0311321113022323-1112323112013010-3111101121000202-2213201120131012"></a>

## labels property — Property reference / 010010133120 / 11

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

- [log_receiver](data-sources--azure_vnet_site--reference--group-008.md#canonical-2321212320110210-3201220001032212-2213023103222133-1312120133323101-3013201210320312-3203323222211113-2221313111133013-1333232322131022): complete subsection reference.

- [logs_streaming_disabled](data-sources--azure_vnet_site--reference--group-008.md#canonical-1100313123320111-2200231033203030-2012210211200001-0122022101330212-1133231201110022-0320210003022230-1222111220202222-0030030111003030): complete subsection reference.

<a id="canonical-3021211331130202-2032203112121133-3301011210030020-3003321003032323-2200133003032313-0220233132030211-2331130102200332-1311302302123130"></a>

<a id="canonical-1032012002213223-2220132221030223-2001333003320321-1202322032001131-3002122323133002-1012310111013203-1302022132130002-1310330010000212"></a>

## machine_type property — Property reference / 010010133120 / 12

Type: `"string"`. Computed.

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

Upstream description:

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0202321130110022-1012120221102323-1321301231130102-3313121233012010-1211022322322110-2032203112322312-0312323030013212-2322302312322311"></a>

<a id="canonical-3212211320013130-1011321122320020-1022020130220022-3300030020332130-3100200110310131-2311302333200110-3003022120201130-2022020013303223"></a>

## name property — Property reference / 010010133120 / 13

Type: `"string"`. Required.

Name of the AzureVNETSite.

Upstream description:

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

<a id="canonical-3023022202002221-3212332121133121-1012010213132303-3023201200013012-3303331020232122-1020323231220130-3030111012121131-3122102231312032"></a>

<a id="canonical-3233021322232131-2202131232312122-1201133301222222-2212032320313333-0230023110111210-2010100202220233-0011212021113323-0131003022212212"></a>

## namespace property — Property reference / 010010133120 / 14

Type: `"string"`. Optional, Computed.

Namespace where the AzureVNETSite exists.

Upstream description:

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

- [no_worker_nodes](data-sources--azure_vnet_site--reference--group-008.md#canonical-0331313212302121-2101302223323001-0333300021312221-2321132311321023-1130102320301113-3312211233201200-1213212322132322-1011101102020213): complete subsection reference.

<a id="canonical-1332303213010023-2113222333301003-1321330002323130-2132311102233013-0221223010220020-0311212103101231-0121232302230023-3022221312301320"></a>

<a id="canonical-3312122301203011-1122333303331230-1322303322313302-0122112033031301-2033010223322220-0000332102302033-2001131212020232-1220213312322002"></a>

## nodes_per_az property — Property reference / 010010133120 / 15

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
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
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [offline_survivability_mode](data-sources--azure_vnet_site--reference--group-008.md#canonical-1100112233133012-0031212113133133-0101021312123110-1031030221212120-2333312012131233-1302130021330313-2000133020220300-3103233223203131): complete subsection reference.

- [os](data-sources--azure_vnet_site--reference--group-008.md#canonical-2231210132330122-2030111311223120-2213001010110333-2100012220101330-3022010210101232-2123132120222110-1233000001221223-1320223221331003): complete subsection reference.

<a id="canonical-3011001113102213-2030030222110130-0210110311232121-2030033010311233-3331020321113200-1330203011111230-3231331201130113-2020130202000221"></a>

<a id="canonical-0132231233300333-1012131000011300-2202103002232203-2201222122103231-1302202000232112-2031233223000010-3220013331010232-3001233330030111"></a>

## resource_group property — Property reference / 010010133120 / 16

Type: `"string"`. Computed.

Azure resource group for resources that will be created.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3102213023202211-2010230333330332-1321330320021013-3231222323321200-0002100023021201-0211300201322222-1203001132321221-2110310312220222"></a>

<a id="canonical-2110231001211331-2023123212330301-0312033011302210-3332232011332333-2012012203020223-3200220121213311-3311213013200112-3201000302020322"></a>

## ssh_key property — Property reference / 010010133120 / 17

Type: `"string"`. Computed.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [sw](data-sources--azure_vnet_site--reference--group-008.md#canonical-3113112022003021-0100233301033123-2031313113130233-2323223032333011-1302323031030331-3202220021210221-2300330020210010-3301322220010112): complete subsection reference.

<a id="canonical-1303302023111211-3231310103010133-0010302302001011-3133230002003112-0032000231121313-0321322012313313-0320333102033301-0333021311303202"></a>

<a id="canonical-0010311310203222-1013302032021310-0312321002133101-1320023300022020-2101311321021223-0323332232311232-0112230310313310-2202222001022112"></a>

## tags property — Property reference / 010010133120 / 18

Type: `["map", "string"]`. Computed.

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console.

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

<a id="canonical-1213101002121000-2002220103333111-3223213210221110-1122002020133121-2202032313033332-3322201133020333-1303102000000021-1123221013230113"></a>

<a id="canonical-0132221020021202-2101201202112203-2232021231033300-2303230322112213-2001011103112101-3311023221110300-2032303210310220-0233310112213230"></a>

## total_nodes property — Property reference / 010010133120 / 19

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
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
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

- [vnet](data-sources--azure_vnet_site--reference--group-008.md#canonical-1211302120213333-3002011311101310-0133223332111330-1313013330132133-2022210131232102-0210120323332022-0023122121132121-3100103210221122): complete subsection reference.

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000): complete subsection reference.

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113): complete subsection reference.

- [waf_signatures](data-sources--azure_vnet_site--reference--group-010.md#canonical-0112121011020013-3003121032200202-2121122103000210-3220000110023102-2211302220322002-3132201311322211-2031200202330003-1233010213132031): complete subsection reference.
