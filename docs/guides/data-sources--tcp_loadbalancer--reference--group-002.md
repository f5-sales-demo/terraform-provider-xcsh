---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-1100311020131301-3010333131132331-0023330230111332-0310000122323102-0220022111000103-0302233213301220-2130230323101221-0103320310122103"></a>

## `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
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

<a id="canonical-2100321003032002-3200312200232202-1110110001200103-2332133313223331-0331132122230121-1300002010013201-3012320100233002-0232210331102231"></a>

<a id="canonical-0223312303221233-1232211100013120-0130321201023121-0311330131102330-2221320210323311-3300010011322113-2322213213003210-3121100231103200"></a>

## `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0122113233033122-1212031131202012-1003210101121201-1330102202322232-2130220221202212-3003100212222003-3213102310020231-1101311333121000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-0310322333022123-1022002101100210-2213302011031032-1210032200203100-0001221333030220-2101333313002102-0103032202132311-1022130022102123"></a>

Type: `"single"`. Computed.

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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

<a id="canonical-0012302102232012-1321311111130230-2310303211311213-3212112222300012-3220022021222221-1303322020121120-1033003122313320-3012220032310002"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site`

<a id="canonical-2112011131311231-3030201200120022-1211300211230212-3123130323301103-2023322000103132-3032321202333030-0030033023231133-2002111120321210"></a>

#### `advertise_custom.advertise_where.virtual_site.network` property

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1023200013220021-0130132300030233-2121203320320010-3221001330211333-2033202031313101-2323022120311231-0123002022212331-0023122310310300): complete subsection reference.

<a id="canonical-1023200013220021-0130132300030233-2121203320320010-3221001330211333-2033202031313101-2323022120311231-0123002022212331-0023122310310300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.virtual_site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0122113233033122-1212031131202012-1003210101121201-1330102202322232-2130220221202212-3003100212222003-3213102310020231-1101311333121000)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-0000113103303012-3332233323003130-1210032120132333-3330301011121100-2322030311132211-3320002012201213-2022231211023120-2012130003100210"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3112130231301120-0000013030130002-0203030013030002-0223021302130323-3010031001101323-2221033123230203-2101233220221233-0232110232301221"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-2102110011023022-1231102131123101-3011023230111312-3220022310310100-2013110020231022-0200330330112330-1201333011121221-2203301010012322"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0221121001323332-1223313301121300-0010303001221332-0320201223131103-2321020130230123-0111320020032011-3321332002221101-0022230323101230"></a>

<a id="canonical-1210331011131310-1212100110112321-3033010220301013-0122011223031202-3133302103133012-0122312112331100-2233223121232200-2222100333133333"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
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

<a id="canonical-3101232330312322-0101213102333032-2021311313121302-3331312030003121-2303232110220033-2200033112312120-1103321021001310-1311303233232032"></a>

<a id="canonical-1000212020210123-0023222302202313-0323303302030011-2003323110303232-3222032221121211-3021300100113013-0211231120001201-1001000302012221"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0310233013030220-3111100223222230-0312202122200212-3113110332121233-0221220033233110-1300030222302200-1321121120220012-0030120222320100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site_with_vip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-3203033131312223-1222231300232132-3232010222121313-0220013302122103-1021201003113121-0213333300310111-1121010023232131-0200113230032131"></a>

Type: `"single"`. Computed.

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

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

<a id="canonical-3333021320011232-3103103130301310-2021231321110230-3120230320302311-3322211333232103-0123030310012303-3023021303011322-0303230030331013"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site_with_vip`

<a id="canonical-0322022110131022-1332100303321303-0331221201133010-1320023211313230-3232201100233002-3132133023211310-1221333122113121-1301010332330313"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.ip` property

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3213302020033203-2213310121300322-1310111002330112-2223103022120122-2032310203312133-1001323002312211-2211003301300120-1311033002231220"></a>

<a id="canonical-0311102231131010-1123301233333203-2203231300303032-0000102311100230-2203123200200020-2320311010101233-0132131120131331-1303300310223100"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.network` property

Type: `"string"`. Computed.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on virtual-site with specified VIP

All outside networks.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220021321132312-2103102131311312-3313100102331030-2012100113330311-2231021221021102-3111332112330010-0022202321201133-3312203223231110): complete subsection reference.

<a id="canonical-2220021321132312-2103102131311312-3313100102331030-2012100113330311-2231021221021102-3111332112330010-0022202321201133-3312203223231110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0310233013030220-3111100223222230-0312202122200212-3113110332121233-0221220033233110-1300030222302200-1321121120220012-0030120222320100)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-0313111112302020-0001333312031302-2013021231221332-3210102020233032-1233210001131330-2110133302233322-2133331000021231-3230031201333022"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2103030211231102-0212321030213233-0102330333000201-0120010011323333-0303311312102300-3000010101300210-3010313232201103-2220202302232111"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site`

<a id="canonical-1103322033230303-3020023223311303-2103130201033110-3322232232121321-1320212302013201-1323130103333033-3023310112313013-3123310330031011"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3000003332203221-0201233032132012-0302100030332111-3302003122230000-2311122301132133-0200210331233121-3012303303222023-1011200111210001"></a>

<a id="canonical-0112232223133312-0330101211103032-0022100102102221-0123310202030130-2313333230100101-3310320233223113-0313001311200300-2211200213221101"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
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

<a id="canonical-0023102112332201-1123023323201300-1322203212222332-2001003220030122-1020331102101312-2201332011000032-1330011233032001-1131220320121020"></a>

<a id="canonical-2222213010201122-1102301103000200-0331111013031233-0010201013102212-3011022102302313-0332132121332311-2232200301310130-2020132322220231"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0132122101013332-2002300333102030-3301212123122331-3230111111010301-2012311030311102-1231333202212000-0001103133211311-0101301312120333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-2311332223031132-0100011021032013-1321220203030310-2300010022212011-0130300201220121-1323321233201101-2121330012301231-0310221121202300"></a>

Type: `"single"`. Computed.

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-1132122211030032-0100012133131303-2020032120222110-2013221322333110-1310001110200203-0122003033320120-2333010312332220-1003022132310232"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service`

- [site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1232222222121221-1200133231302330-1333200113222010-1300123200230120-2102211213033011-3121303330120020-2330133032102222-2312030231302221): complete subsection reference.

- [virtual_site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2132033020221130-3230300030220011-2211302330010320-3001020213320211-1013332232120203-3230133001310302-2103222121333113-3230223331023233): complete subsection reference.

<a id="canonical-1232222222121221-1200133231302330-1333200113222010-1300123200230120-2102211213033011-3121303330120020-2330133032102222-2312030231302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.vk8s_service](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0132122101013332-2002300333102030-3301212123122331-3230111111010301-2012311030311102-1231333202212000-0001103133211311-0101301312120333)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-2330230222331212-2023220202010000-0221322222021103-3131333320123011-1230102022303102-0222113210010122-0200302202103211-1020221311311100"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3223132023220213-2100302112332101-1123131201012201-1323132333221123-1312220222330012-0222323031221232-3303130302131110-1111131132312111"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-1100201012023100-1012011000010313-0223331211021111-2001232031220201-3233023333330100-3332110013101011-1122211001321302-0011033221000002"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3133001131012320-1200303131031323-0131310020202300-0011300233321333-0332132131032231-3103003200332133-2322020212322003-3101110003002301"></a>

<a id="canonical-0322313121013212-1132121203310102-1110202112122021-1310323313020003-2030130130021222-2200200321311322-3132330100012112-0230323300201211"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
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

<a id="canonical-0120230020323010-1002222102230302-0301030031203033-0222121123023011-2223303031303223-1100212110103003-3110202212121212-2033203030201100"></a>

<a id="canonical-3130112031002202-1012211311003212-0112003020030212-1310331321120130-3001000203211013-3212313222130122-0003322203033030-2010012230230021"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2132033020221130-3230300030220011-2211302330010320-3001020213320211-1013332232120203-3230133001310302-2103222121333113-3230223331023233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service.virtual_site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.vk8s_service](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0132122101013332-2002300333102030-3301212123122331-3230111111010301-2012311030311102-1231333202212000-0001103133211311-0101301312120333)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-3111131113300323-1123031310001210-1321030323211210-1233112122310101-3331200221130221-2012310310010113-2212311012033121-0331023201232121"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3013022320222031-2331100102312322-3311333022332020-2311011210130233-3301312111003222-3333112112300323-3120313003122100-1221100211013110"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-3201122003211333-3303101211223302-0303021002222103-1110003012102120-2120113022023100-3012100331113222-2112123330201231-0313000331203231"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0012312120212210-0023331020321310-2121022210223303-2300332223321022-0213331131132112-2130122330302233-1210313020310331-0113301101023020"></a>

<a id="canonical-3310020013321301-1211120322200222-1130121222203102-3101333221311132-1132033121111231-3333320130031310-2211001200211233-3321022320012331"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
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

<a id="canonical-1202021312220001-3232013103322110-3100112103010323-1121330221122313-3330012013211010-1201031310222032-1113201303033321-3130121021311022"></a>

<a id="canonical-3122111300303020-3232110230020323-3232220312103031-3303133323310130-2302011230123230-1002201130313113-0220320221232021-2122003210103100"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1112201032021131-3221000132200333-3131001223100003-2232313233132313-2111023302121223-0310220211131233-2130013222232323-2223231123010031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- advertise_on_public

<a id="canonical-2200031202211021-0313200202312233-1013023211320030-3123000101033002-1113110302231111-2101321210002322-3331003310213213-0312330021311122"></a>

Type: `"single"`. Computed.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-1320131302312321-1020002131003220-2310020303313312-2011203311202202-1302333203001210-0223001012020003-0101300320132123-3332312022213203"></a>

### Direct properties for `advertise_on_public`

- [public_ip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2323110201301120-0232100210311033-3213000230012232-0121301023200133-3003210021330131-0331020233211323-0020203312031031-0023232033013202): complete subsection reference.

<a id="canonical-2323110201301120-0232100210311033-3213000230012232-0121301023200133-3003210021330131-0331020233211323-0020203312031031-0023232033013202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1112201032021131-3221000132200333-3131001223100003-2232313233132313-2111023302121223-0310220211131233-2130013222232323-2223231123010031)
- advertise_on_public.public_ip

<a id="canonical-2302330022113120-0132220223113030-1203002312332220-3011132222110320-0100103320332132-3231213110010122-0221301003313120-3312121011200103"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3233212233021022-1213112100213022-2323111230203103-3223033332301212-1221132200220301-1020212220312112-3102322201131101-2101032123311112"></a>

### Direct properties for `advertise_on_public.public_ip`

<a id="canonical-1222030022130330-0331013312310213-0220003120000133-2122222123123030-1223301100031132-0103131211033302-3113031021023121-3022122220220013"></a>

#### `advertise_on_public.public_ip.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0232110030303000-3323233312210002-0331002220113231-1222032131332110-0320133301131131-0232310110113333-1331330112132330-2303112220222012"></a>

<a id="canonical-1203020111330302-1123103320003330-1233333332132320-2221110120200322-0213022312111222-0211230201222232-0022303113330203-2200132033332323"></a>

#### `advertise_on_public.public_ip.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
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

<a id="canonical-2323212133012203-2231313332303232-0213313201013313-2200010311322211-1221211112100130-0112221313113110-0213322230323013-3310210023331003"></a>

<a id="canonical-0330321023210313-0113021023023330-2232233300103101-1323103110121022-0320332230100013-2220200111222222-2123321130031101-2302021210211022"></a>

#### `advertise_on_public.public_ip.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2103301002331000-0011121122110023-2213202121233302-1122011020100033-1201330300100321-3021121333233123-2002013133223103-3102121301033313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public_default_vip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- advertise_on_public_default_vip

<a id="canonical-1123330002330211-1310302201203201-1113023310221132-0301203233322302-0032103302032320-2010301030313101-2213001103002101-3232220103002213"></a>

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

<a id="canonical-1120310112113021-0011220322222102-0130122230032031-0102010320303202-0001311211211213-3321333330113130-2001230311133133-1322321301312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_lb_with_sni` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- default_lb_with_sni

<a id="canonical-3323001220100213-0332200120333110-0322222321133332-1232031323010000-3212121100010232-0301020212002130-0303222103123032-1320010121121321"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_lb\_with\_sni, no\_sni, sni; Default: default\_lb\_with\_sni\] Configuration
parameter for default lb with sni.

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

OneOf alternatives in this subsection:

- [default_lb_with_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3323001220100213-0332200120333110-0322222321133332-1232031323010000-3212121100010232-0301020212002130-0303222103123032-1320010121121321)
- [no_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2112113331223322-3231232001320233-1302030020311232-3100000332022230-1032022022211232-2010330311022330-0123212123130312-3120230231130212)
- [sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2312220212030333-0230231220111231-0012033022321102-0311002130002223-2112002002110030-0212103020302010-2021321302012013-2221101100130002)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301200001232003-2230220223123022-1031331131120332-3030120311113103-0010103000333222-3012312213013120-0313333310330302-1322122222220023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `do_not_advertise` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- do_not_advertise

<a id="canonical-0001303000210302-2210122211320302-0002200332110310-0310113320302002-1122011331112123-3223311220302121-3230321301031211-3123012131133000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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

<a id="canonical-2102221023203330-2101321113012132-3212312223313302-2333003020300013-3212302122333112-1203323112212131-0113122301023121-1123321312122331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `do_not_retract_cluster` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- do_not_retract_cluster

<a id="canonical-0021013132003110-0210102220132311-3302202220213101-2313122133232322-0211013312030111-2133331021111111-1102303210112013-1030322030021232"></a>

Type: `["object", {}]`. Computed.

\[OneOf: do\_not\_retract\_cluster, retract\_cluster\] Enable this option

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

OneOf alternatives in this subsection:

- [do_not_retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0021013132003110-0210102220132311-3302202220213101-2313122133232322-0211013312030111-2133331021111111-1102303210112013-1030322030021232)
- [retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1131023233233002-0323313303102032-2022112113312010-0331130133000123-0003323331110103-0100213211012000-1331100102121003-1302332022231233)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233323330321012-0022220130111220-0001311203331000-0032101313230230-0003312301003011-3322313221113031-0000100011331113-2312101310330012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hash_policy_choice_least_active` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- hash_policy_choice_least_active

<a id="canonical-0213123311230102-3223303113132302-0222311331313011-2233121031121311-1102323230101312-1223301003121202-1122203220021330-1112321202221300"></a>

Type: `["object", {}]`. Computed.

\[OneOf: hash\_policy\_choice\_least\_active, hash\_policy\_choice\_random,
hash\_policy\_choice\_round\_robin, hash\_policy\_choice\_source\_ip\_stickiness\] Enable this
option

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

OneOf alternatives in this subsection:

- [hash_policy_choice_least_active](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0213123311230102-3223303113132302-0222311331313011-2233121031121311-1102323230101312-1223301003121202-1122203220021330-1112321202221300)
- [hash_policy_choice_random](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3212033231123032-1011312023323013-1011123300220023-1231033203130221-1201031100111300-1010230221333212-1313301103132100-1122023101232212)
- [hash_policy_choice_round_robin](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0013003300111331-2221131030212002-0133311311011203-2121231303120023-1320011201221133-2232301003113132-2321033323301211-3331202302033223)
- [hash_policy_choice_source_ip_stickiness](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0001313102231321-2102202120232131-3313220313103221-2312113230312002-3312032203302111-0233332302110333-1313210300232111-1311102213312222)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323103221023221-2313321022023032-0102100310112221-3013030013003012-3002031302202202-0020002001010103-2300123101332113-1301322121013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hash_policy_choice_random` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- hash_policy_choice_random

<a id="canonical-3212033231123032-1011312023323013-1011123300220023-1231033203130221-1201031100111300-1010230221333212-1313301103132100-1122023101232212"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for hash policy choice random.

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

<a id="canonical-2001222213202013-2110122010013020-0222012122321312-2221120003230033-3112000110322320-0030000330021323-3103332101033100-3231323011021122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hash_policy_choice_round_robin` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- hash_policy_choice_round_robin

<a id="canonical-0013003300111331-2221131030212002-0133311311011203-2121231303120023-1320011201221133-2232301003113132-2321033323301211-3331202302033223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for hash policy choice round robin. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-2101213003031210-0111210032000102-1202211230323123-2123332011131030-3210312021210331-1330112233100331-2210212303011011-3120212322222111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `hash_policy_choice_source_ip_stickiness` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-0001313102231321-2102202120232131-3313220313103221-2312113230312002-3312032203302111-0233332302110333-1313210300232111-1311102213312222"></a>

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

<a id="canonical-3232303121303303-3310101031211332-3320310013303113-3130011003231312-0221113013330021-0103001323103000-3230002332010321-1123002233310003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_service_policies` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- no_service_policies

<a id="canonical-0121112133023001-2133103202133131-1030333112200300-0221311100103201-3330222002023012-0321002320113000-0310301333203033-2200011113112222"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no service policies.

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

<a id="canonical-3323311322022213-3030031301102132-3222213210012013-2113301030000013-1230313033002102-0010100310021223-2002303301021210-3332001213302330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_sni` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- no_sni

<a id="canonical-2112113331223322-3231232001320233-1302030020311232-3100000332022230-1032022022211232-2010330311022330-0123212123130312-3120230231130212"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools_weights` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- origin_pools_weights

<a id="canonical-3200202332003310-0211023331000103-0121311213201021-2001122321011101-1130023223300221-1121311210313223-3022112123003300-0100230133313000"></a>

Type: `"list"`. Computed.

Origin pools and weights used for this load balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0033333231322200-2212101031310100-1020023320203111-1311300203101332-3112003112010030-3131122030111331-0302223332223233-2032003103021120"></a>

### Direct properties for `origin_pools_weights`

- [cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1100323120213031-0021203033221103-0232210303113332-1232122132332320-2022312033230033-3203211222032232-2203103200232310-0120121202303230): complete subsection reference.

- [endpoint_subsets](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3233020300102002-3212200311012103-3303223032200213-0321120133210212-1310320312202012-3223222313112032-1320233023030102-3010212231102021): complete subsection reference.

- [pool](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3113021102231311-3020330031331101-3221032221322031-1022322103230303-1130012222122212-3122202323303323-3113103301320101-3130030230221032): complete subsection reference.

<a id="canonical-0103301112020110-2133210101322130-3120321133113130-1213031012300313-1233031101032300-2023001121123303-0030103230013311-1223211120321312"></a>

<a id="canonical-0212303220110033-0033203002123112-1123111031032033-2323013021320333-3220213103110033-1023011331320233-0331113002122230-3000320223213302"></a>

#### `origin_pools_weights.priority` property

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1231132003213230-2312112323103213-1303032131103311-1020022230023332-2001001012010113-3201030003313213-2101020123120220-0123312003101312"></a>

<a id="canonical-2030131001230230-2312321032310332-2023312103331003-3232122300301322-2321312103201312-2132033023122313-3121321033331333-2110110123220223"></a>

#### `origin_pools_weights.weight` property

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0,
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

<a id="canonical-1100323120213031-0021203033221103-0232210303113332-1232122132332320-2022312033230033-3203211222032232-2203103200232310-0120121202303230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools_weights.cluster` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101)
- origin_pools_weights.cluster

<a id="canonical-1113012100011030-2001322102301030-1131010011033030-1313003012021101-0110011203030111-0032230200000311-3003320330302213-1200101331100111"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1103030032332202-0212212312100131-1012331133221330-2131122221110112-0031010012312113-1112131133031210-1302333320302312-2013031000202032"></a>

### Direct properties for `origin_pools_weights.cluster`

<a id="canonical-2212013331103221-3031332033320031-1210122323123331-1031313113310310-3100223311311321-3120113311131312-3112322132002012-3220331130032303"></a>

#### `origin_pools_weights.cluster.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3200332330131231-3103203300030030-3031323302132210-3132221121211320-1230030120120212-0003023213012210-1221001331011023-1200032332022033"></a>

<a id="canonical-2023021211333022-1221132201201211-3113130301210113-3132101011001223-3000120111221103-0120221202103223-0102321031112223-1200022103032300"></a>

#### `origin_pools_weights.cluster.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
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

<a id="canonical-0300331231310020-2221010013102111-3022100132132022-3312032231320012-2101033320221132-1123310201312103-1221200033010122-2222310330220323"></a>

<a id="canonical-2201311012021312-0120330111301333-1302011102133022-0101311221030223-1131303210112322-2210203123032332-1320302232120221-1030123310210121"></a>

#### `origin_pools_weights.cluster.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-3233020300102002-3212200311012103-3303223032200213-0321120133210212-1310320312202012-3223222313112032-1320233023030102-3010212231102021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools_weights.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101)
- origin_pools_weights.endpoint_subsets

<a id="canonical-3220123132322001-0221230110222113-2301001313220031-2003111011112222-0210003310111311-0231121311123233-3110203332231311-1100133102333121"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113021102231311-3020330031331101-3221032221322031-1022322103230303-1130012222122212-3122202323303323-3113103301320101-3130030230221032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools_weights.pool` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101)
- origin_pools_weights.pool

<a id="canonical-1132213133202332-1322012003101323-3330023000200011-0331233133011332-2222300223122200-0331130322222123-2231333130011310-3012311322322233"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0120003031302120-2211013231221101-0121001101202030-2133322111230220-1012301120121331-3000100011303203-1103313022323122-0301310013232232"></a>

### Direct properties for `origin_pools_weights.pool`

<a id="canonical-1202230130301301-0102231022221021-1132312202023331-0323033331101101-2200020223321123-2133213330222033-3113001133033031-1130322302230120"></a>

#### `origin_pools_weights.pool.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3233100010300122-3230221023120200-1113012021131032-0020202101121311-3100220330130000-2333032032122032-2033230320321232-0220233100122310"></a>

<a id="canonical-0023003102301233-2002101101120201-0102020033210030-1110232032011132-0312112031332321-0202032120001201-0200200203302003-2010123202121033"></a>

#### `origin_pools_weights.pool.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
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

<a id="canonical-3321202331223033-2020301031110310-2212010110322331-2220302023003201-3012111002301110-0111013231221032-1311310333220110-1123230331332102"></a>

<a id="canonical-2200022121222103-3332222012203303-2220020202100031-0321000222202023-2102132100033231-2002223221331101-3203231301312103-1211103012133201"></a>

#### `origin_pools_weights.pool.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-3030013003230233-1020032322120203-1023112003112123-2102002023033311-2321113323313010-0110100232101332-0123011112033022-0022330033023211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `retract_cluster` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- retract_cluster

<a id="canonical-1131023233233002-0323313303102032-2022112113312010-0331130133000123-0003323331110103-0100213211012000-1331100102121003-1302332022231233"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-0123020233112322-0223121002333303-2100322322223000-3300212133321100-3320102320231010-0321033323023300-3233233123103010-2121331023133111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service_policies_from_namespace` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- service_policies_from_namespace

<a id="canonical-3301033031210010-2323133031222031-3022302003130302-0332310330103322-1101123130223311-2303021212131110-1301210012202313-1202323210102000"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-2133300001211310-3203211112301011-3012211322222112-3322000131211112-0133200313302211-3311012013010002-0330310112023122-1200113222020223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sni` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- sni

<a id="canonical-2312220212030333-0230231220111231-0012033022321102-0311002130002223-2112002002110030-0212103020302010-2021321302012013-2221101100130002"></a>

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

<a id="canonical-2032030011011133-2021231112123000-0010123000231211-1233113301302232-0220010312011312-1000210021110111-1203111330233211-3303302223213330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tcp` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- tcp

<a id="canonical-1233311120231210-2331220321101311-1220130030202313-0102210322110321-2301130120100010-1311032331121332-3330332030322100-0111113320312013"></a>

Type: `["object", {}]`. Computed.

\[OneOf: tcp, tls\_tcp, tls\_tcp\_auto\_cert\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

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

OneOf alternatives in this subsection:

- [tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1233311120231210-2331220321101311-1220130030202313-0102210322110321-2301130120100010-1311032331121332-3330332030322100-0111113320312013)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1212102131222122-0010230212231023-2032230231022002-3301321213331102-1011221101033321-2331201103200201-1332121102010012-1210120021232232)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3122010013013333-0010120031233203-2230121020320021-2303221102121132-3030013013032022-1321333312130213-1313113220130030-1002023233330000)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_tcp` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- tls_tcp

<a id="canonical-1212102131222122-0010230212231023-2032230231022002-3301321213331102-1011221101033321-2331201103200201-1332121102010012-1210120021232232"></a>

Type: `"single"`. Computed.

Choice for selecting TLS over TCP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-3323312010220012-0002122220102200-2122001023210330-1320221312023021-3332313123113202-0210000123103313-1120320200323231-2022021303032322"></a>

### Direct properties for `tls_tcp`

- [tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123): complete subsection reference.

- [tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120): complete subsection reference.
