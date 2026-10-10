---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-1323200312211101-3000213103000332-3033030220220301-3321110123121210-3021212000212000-1202001200122110-1121122303001302-3302233022031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-001.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-001.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="canonical-2321101003301320-0132020133333132-0201310101022302-1303010132031330-3332032102203103-0323220200012211-2312021210220302-0020311330210000"></a>

Type: `"single"`. Computed.

CloudLink ADN Network Config.

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

<a id="canonical-3113323302000112-3023002201101322-1120120132030113-1103330020220121-2201210112102333-2310103231330012-1200003013333303-1231003123331100"></a>

### Direct properties for `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect`

<a id="canonical-2323203213210011-3013013213032101-2202113110121330-2312030232233132-3330211231013102-1303223100100012-3233311003322121-3003200011102110"></a>

#### `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` property

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0120213130102001-3131030112310103-0301302322131302-1322310201013033-1031030121332323-2001310120030023-0201320303102323-0033212103321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.hosted_vifs.site_registration_over_internet` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-001.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-001.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

<a id="canonical-0100313110132030-3220332211321231-0222031210222120-0103203001201320-0020122311031132-2110002113011101-3330201203203203-3030312030000013"></a>

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

<a id="canonical-1120303330300310-1302023120223012-0010311013220332-1012220130010311-1123313031102023-3232232001032113-0033323222012230-1233220122203211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.hosted_vifs.vif_list` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-001.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-001.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="canonical-0333121031232201-1332232200302310-1001210212202202-1123012102310122-1013223000010213-0030021221122333-2233001331333303-2131202220112232"></a>

Type: `"list"`. Computed.

List of Hosted VIF Config. List of Hosted VIF Config.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 30,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 30,
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
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0023203310323013-2033212122320330-3213301132100230-2112321021012111-1230121223201221-2112130321011300-3020221013230230-2331330103000103"></a>

### Direct properties for `direct_connect_enabled.hosted_vifs.vif_list`

<a id="canonical-2231210011203332-1313212310323222-1230100030000122-2231100312221312-2100221131033012-3300232121302330-1322103101023110-2202203222323023"></a>

#### `direct_connect_enabled.hosted_vifs.vif_list.other_region` property

Type: `"string"`. Computed.

\[Enum:
af-south-1|ap-east-1|ap-northeast-1|ap-northeast-2|ap-south-1|ap-southeast-1|ap-southeast-2|ap-southeast-3|ca-central-1|eu-central-1|eu-north-1|eu-south-1|eu-west-1|eu-west-2|eu-west-3|me-south-1|sa-east-1|us-east-1|us-east-2|us-west-1|us-west-2\]
Exclusive with \[same\_as\_site\_region\] Other Region. Possible values are \`af-south-1\`,
\`ap-east-1\`, \`ap-northeast-1\`, \`ap-northeast-2\`, \`ap-south-1\`, \`ap-southeast-1\`,
\`ap-southeast-2\`, \`ap-southeast-3\`, \`ca-central-1\`, \`eu-central-1\`, \`eu-north-1\`,
\`eu-south-1\`, \`eu-west-1\`, \`eu-west-2\`, \`eu-west-3\`, \`me-south-1\`, \`sa-east-1\`,
\`us-east-1\`, \`us-east-2\`, \`us-west-1\`, \`us-west-2\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  }
}
```

- [same_as_site_region](data-sources--aws_tgw_site--reference--group-002.md#canonical-2121021030232021-3233123102000303-3212113213202013-1003220133131200-0213323212123213-2003321223123133-3113101301110002-1123000323030020): complete subsection reference.

<a id="canonical-2002200330330311-2230213211011122-3112203120022233-1102022022011231-2202333301120221-0133133011000032-1230100222110302-3020323211313320"></a>

<a id="canonical-1002123220021230-3021003102211323-1001010231311121-0303003313203123-1302213303333001-2022232222232231-1230313323022010-3300213201110000"></a>

#### `direct_connect_enabled.hosted_vifs.vif_list.vif_id` property

Type: `"string"`. Computed.

AWS Direct Connect VIF ID that needs to be connected to the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-2121021030232021-3233123102000303-3212113213202013-1003220133131200-0213323212123213-2003321223123133-3113101301110002-1123000323030020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-001.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-001.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_tgw_site--reference--group-002.md#canonical-1120303330300310-1302023120223012-0010311013220332-1012220130010311-1123313031102023-3232232001032113-0033323222012230-1233220122203211)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="canonical-1011200330323330-2323320003233213-0131313033300302-2101223133130123-3113300002101103-1130132111212131-2022101233332032-3011110300312020"></a>

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

<a id="canonical-3223331130303203-1032102332230311-1312022330000323-1322110023132332-0130230133222002-2002122012033202-2312021211203030-2002121030100230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.standard_vifs` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-001.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- direct_connect_enabled.standard_vifs

<a id="canonical-0021132032220222-3313030201213213-2113013132001223-1322002331101133-0010201200123023-0210113210223122-0122310002230013-0201233233011023"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for standard vifs.

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

<a id="canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- kubernetes_upgrade_drain

<a id="canonical-1122122331212220-3032221130130121-0313003211221012-2230033000102231-2010301113220211-2131223323213320-1032133012300310-1211323311001222"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

<a id="canonical-1111113001223230-1102121201223223-0022330000330302-0100032203122210-3212001131012200-3131300310222100-2102231320313133-2100113103111020"></a>

### Direct properties for `kubernetes_upgrade_drain`

- [disable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1330022323302122-2313111301230123-3111011212023010-0220003010013113-0323002001011103-2223311323220103-3222001000222120-0110033132122333): complete subsection reference.

- [enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000): complete subsection reference.

<a id="canonical-1330022323302122-2313111301230123-3111011212023010-0220003010013113-0323002001011103-2223311323220103-3222001000222120-0110033132122333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.disable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-1021212301023113-1001322031023123-2331000332002122-2210103203322303-0320221312311011-3121211231300203-2313212333102212-2013311031300321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable upgrade drain.

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

<a id="canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-3303003312323121-3230301232021123-0011200130203332-2313202303331010-1221111311231320-0232233001003320-3123030210020223-2112111333111202"></a>

Type: `"single"`. Computed.

Specify batch upgrade settings for worker nodes within a site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

<a id="canonical-1012000030311223-3021213132112012-0302021132000320-2011023103311323-1310200320303302-2331301000022133-3200311110212001-3230131030022332"></a>

### Direct properties for `kubernetes_upgrade_drain.enable_upgrade_drain`

- [disable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-2000003330112002-2003330021300013-0000200211203121-3030213033233232-0303121322330021-2303201310233311-2303330002230123-0223110013320030): complete subsection reference.

<a id="canonical-1013312200111212-1302123211112313-0320023312020132-3023231031201202-3222120323002232-1202320102023033-1001211210203322-0111321032033113"></a>

<a id="canonical-2321233321301033-2323033330120123-2011212203333030-1321122130101300-1222320330232131-3212320012322212-2322231332110010-2010021312122123"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` property

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-3230133031301232-1222330323310331-1012011001303101-2030010222000001-2100301020321122-1221011021200113-0232202121233220-2220332200220200"></a>

<a id="canonical-1103103321133112-1112111031103200-3130001230213332-1030210013120211-0220001113230311-3312131102302013-3230200131230033-1220103000310231"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` property

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-0223320123313300-1203200203010303-3033231212310313-0302113102111120-2201020113220113-1211132011032313-1001010011231223-3112011211001032"></a>

<a id="canonical-0003313322300331-2102102010212033-0221110023223012-1130133110233030-0332101222033313-1230133313231232-3131001303130002-0220113033221333"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` property

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0331112212013002-2203201322101103-1011202201100112-1003001000011020-0210310013012130-2001021333323133-3131012130123301-1310021333331000): complete subsection reference.

<a id="canonical-2000003330112002-2003330021300013-0000200211203121-3030213033233232-0303121322330021-2303201310233311-2303330002230123-0223110013320030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-2012012130000103-1102333031131303-2001003112030121-2233320123033111-1202331330123223-2222101111321011-1321233100230301-1223003101022221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable vega upgrade mode.

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

<a id="canonical-0331112212013002-2203201322101103-1011202201100112-1003001000011020-0210310013012130-2001021333323133-3131012130123301-1310021333331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-1012331302012031-2212033001331030-2002301103211302-0031112311231022-2211122111112203-3321210333112302-3023013103021201-3111322130313001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable vega upgrade mode.

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

<a id="canonical-2020021301323331-3001033223203323-2332033222202220-3311033003321302-1031013201223320-2333333000113313-1000012203232223-3002210130211211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- log_receiver

<a id="canonical-3030230100032313-2200010231230310-1212303202312110-0123013031313320-2201210013021220-0221212203303210-3103200302312121-0213011232213332"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

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

- [log_receiver](data-sources--aws_tgw_site--reference--group-002.md#canonical-3030230100032313-2200010231230310-1212303202312110-0123013031313320-2201210013021220-0221212203303210-3103200302312121-0213011232213332)
- [logs_streaming_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-1311031232301110-0101211333020012-2300103001323231-0300221120113121-0111103220222123-0222120000232331-1032131220021020-1031000220022031)

Select alternatives according to the provider validators above.

<a id="canonical-2311302232223202-2002110333203101-1013023313132030-3210031112201212-1333202133210301-1122131111220333-1201212033030022-3212223002223001"></a>

### Direct properties for `log_receiver`

<a id="canonical-1102033123210212-1211003203000332-0031232321020330-3023210202122331-0122200320332220-2210100123030002-2001202231100220-0000111022122211"></a>

#### `log_receiver.name` property

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

<a id="canonical-0232000112231101-2210221322332000-3132210322200313-1233310013122120-2130311213331213-3221031302311310-0231223123302333-1322202323211011"></a>

<a id="canonical-0323220223231200-1312201120000102-1030332321221231-0033122320111010-2322222122320323-2231010321203111-2313023220001122-3132021311121001"></a>

#### `log_receiver.namespace` property

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

<a id="canonical-0020300220133220-0212313132101101-3100200031123220-1101332023221222-2231233311301010-1321012230212012-1332203133031323-1131112032232111"></a>

<a id="canonical-1200002010030032-0010131032333031-3033010333030213-2210203010213203-2003100101101231-1202101131102300-3012301102023130-0032221230322033"></a>

#### `log_receiver.tenant` property

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

<a id="canonical-1113021113220301-2010123223000213-3012113321101301-3220123032012310-0000133322303120-1033130320232331-1300213221112233-1232321020023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `logs_streaming_disabled` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- logs_streaming_disabled

<a id="canonical-1311031232301110-0101211333020012-2300103001323231-0300221120113121-0111103220222123-0222120000232331-1032131220021020-1031000220022031"></a>

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

<a id="canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `offline_survivability_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- offline_survivability_mode

<a id="canonical-0121322102202200-1331102101131233-1223212212232333-0003220302220012-3222203311222213-0012212203333031-3200021112131013-3230321023213130"></a>

Type: `"single"`. Computed.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

<a id="canonical-3211120021232130-2133210022202210-3222100213230230-0033110221333013-0323333032202031-1321313300311120-1311320132102121-3202122002111121"></a>

### Direct properties for `offline_survivability_mode`

- [enable_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1011032221023211-3332303213033000-2311323210133312-3222031002030013-3021133130302030-1320110013200131-3100232112212222-0133032303031231): complete subsection reference.

- [no_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-3300233122100233-1231203331022330-2130331002022111-2120312301121301-2100010122020002-2301223120133203-1321213031130210-3003030201201120): complete subsection reference.

<a id="canonical-1011032221023211-3332303213033000-2311323210133312-3222031002030013-3021133130302030-1320110013200131-3100232112212222-0133032303031231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `offline_survivability_mode.enable_offline_survivability_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-0102012010001030-3220201221032100-2120133313313132-0123130011110111-3101000221211001-0010312231021102-3100100130032023-0132030020030311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable offline survivability mode.

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

<a id="canonical-3300233122100233-1231203331022330-2130331002022111-2120312301121301-2100010122020002-2301223120133203-1321213031130210-3003030201201120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `offline_survivability_mode.no_offline_survivability_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-2231222223113003-3000113203203111-1012301003122221-0000113331323212-2203010210322130-2021323312022301-2103300333001023-0102222231221303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no offline survivability mode.

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

<a id="canonical-1131211032312230-3213210020110320-2302211020330112-2302302311111200-3123111312222202-0112321312131023-1133232031010020-1131203321212130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `os` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- os

<a id="canonical-1232012323003122-0120303332133123-2220313303021232-3032131333312031-2332331321012213-0130213020300022-2332023032222111-0213321231100301"></a>

Type: `"single"`. Computed.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

<a id="canonical-3110321120323232-3303111111221012-1121303201121311-0022013202301131-1220100110301000-0120001021002233-0230222020213330-2120000313110123"></a>

### Direct properties for `os`

- [default_os_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322333210121233-1023320320022331-3201323132302200-3120102220303122-2323322330133213-0220030221233021-2203002010203002-1303231022302133): complete subsection reference.

<a id="canonical-3000032112303030-1002121020003022-0033031230011321-3111330301032002-2132332130202133-3301201213132032-1012012233313102-3322322232331010"></a>

<a id="canonical-0010303313201221-0000011022231231-2313301112120312-2021000021030221-3023100231102313-1313122112002323-2202323010012211-1033202312033132"></a>

#### `os.operating_system_version` property

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

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

<a id="canonical-0322333210121233-1023320320022331-3201323132302200-3120102220303122-2323322330133213-0220030221233021-2203002010203002-1303231022302133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `os.default_os_version` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [os](data-sources--aws_tgw_site--reference--group-002.md#canonical-1131211032312230-3213210020110320-2302211020330112-2302302311111200-3123111312222202-0112321312131023-1133232031010020-1131203321212130)
- os.default_os_version

<a id="canonical-2023023303222013-3211003212023212-0030332022103301-1322023030320033-0123212232130033-2331311012021033-1132330010310023-3103310322022200"></a>

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

<a id="canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- performance_enhancement_mode

<a id="canonical-1101020121232223-3210011003013000-0223332002021320-2201131210333333-2013100221112302-2031013202200013-1321130210022200-3220232202332300"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

<a id="canonical-0332122211010102-1301032302023322-2122203301001001-2002122201230232-2000331320320231-1232233231100332-1222010331022333-2300333232033211"></a>

### Direct properties for `performance_enhancement_mode`

- [perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103): complete subsection reference.

<a id="canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-3030233020320202-2123120321021312-3301010310332203-0312012120000012-2020033002111011-3130201222210002-3122331110102120-2332221031222311"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l3 enhanced.

Additional upstream details:

L3 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

<a id="canonical-0333233301301212-3110222212101122-0003301111100213-2210200213020221-1132231333300021-3323200131312033-0220323331103132-1000331101112202"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l3_enhanced`

- [jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-1333121320231111-0230331332033100-1133312012012321-3202200102302210-2100311101213232-2013302003103201-1031331333010121-1213201300020230): complete subsection reference.

- [no_jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-3110022120220222-2223030213102012-2230110120123112-2110000002233112-2223203023002323-3332312001322101-2313202212321201-3100003211301111): complete subsection reference.

<a id="canonical-1333121320231111-0230331332033100-1133312012012321-3202200102302210-2100311101213232-2013302003103201-1031331333010121-1213201300020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-1120100000332311-0213302031131332-2130020033033212-1321032031113103-2312332111020132-1201312132100100-2131100022230023-3111333101220000"></a>

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

<a id="canonical-3110022120220222-2223030213102012-2230110120123112-2110000002233112-2223203023002323-3332312001322101-2313202212321201-3100003211301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-2303321201113123-0310112320213100-0022301322230013-2212031111120202-0222112200233211-0323210013003203-3213000123031322-1012033332102033"></a>

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

<a id="canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-0313233110130320-1012221112013120-1320201110033212-0310132123111323-2120230311233202-2201201200311223-1310202120033202-3323211112010332"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Additional upstream details:

L7 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

<a id="canonical-1310000221002212-2010301022032310-3230011213120223-1000312000300212-2201233030202123-2121310303003202-0131112020132133-0320022120033013"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l7_enhanced`

- [jumbo_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-2002033222131000-3223103011213132-3103020002031101-1232311132213111-3330212131311011-1132030000123012-0133332031312303-3132323210032210): complete subsection reference.

- [jumbo_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-3201323033210031-3213130333211311-0133122020313320-1121122310102210-3332100013303322-2123123112111313-2030013320320303-2022212231100112): complete subsection reference.

<a id="canonical-2002033222131000-3223103011213132-3103020002031101-1232311132213111-3330212131311011-1132030000123012-0133332031312303-3132323210032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3203213023101120-2213121223303013-0021110100010223-0312131212003131-2101030021100213-2331301030302033-1210122111011301-0011201003333111"></a>

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

<a id="canonical-3201323033210031-3213130333211311-0133122020313320-1121122310102210-3332100013303322-2123123112111313-2030013320320303-2022212231100112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-0021300211302203-1220210311311133-0213221020033130-0102020023033133-2131012231303302-2120021000032232-3300233020321001-1112002320220121"></a>

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

<a id="canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_connectivity` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- private_connectivity

<a id="canonical-2330311303310000-3201101003320222-1332021212103110-0001203211113110-3100222000231102-3210211331232231-0231310132333032-2201013101212203"></a>

Type: `"single"`. Computed.

Configuration parameter for private connectivity.

Additional upstream details:

Private Connect Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

<a id="canonical-2210312021230111-2330001222320121-1313201313013033-2101123201233122-1321220111001132-1101330202131112-3322111222301200-1131200013310221"></a>

### Direct properties for `private_connectivity`

- [cloud_link](data-sources--aws_tgw_site--reference--group-002.md#canonical-3321322232102332-2112201320310312-0311210203101010-0312213101303202-0232131102223122-3223222232103321-0310321102023012-2213200200022221): complete subsection reference.

- [inside](data-sources--aws_tgw_site--reference--group-002.md#canonical-2010132320123203-0002302222122002-0130011020313231-0202032223312101-1230131023013123-3313020112213310-0303101110111311-2032233030300313): complete subsection reference.

- [outside](data-sources--aws_tgw_site--reference--group-002.md#canonical-1323110210103000-3202031100211003-0332312012313311-2132130311232210-2023023030200222-0030112130333122-0131222013232010-2321112333022001): complete subsection reference.

<a id="canonical-3321322232102332-2112201320310312-0311210203101010-0312213101303202-0232131102223122-3223222232103321-0310321102023012-2213200200022221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_connectivity.cloud_link` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- private_connectivity.cloud_link

<a id="canonical-0030010303001102-3211101233101010-1031003113032013-0003010312130110-1313122011002123-0021012031012203-0322032211223232-3000130131121322"></a>

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

<a id="canonical-1321232033202202-2202232200031301-1331310222223302-3201201010010221-2130223313010222-1330333232031222-1320132122011013-2033121130000322"></a>

### Direct properties for `private_connectivity.cloud_link`

<a id="canonical-1110230310311330-3030213103330201-1132113100030031-3313120002221230-0313103030122001-3100020001330322-1101313133301031-3300231230030200"></a>

#### `private_connectivity.cloud_link.name` property

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

<a id="canonical-1130022210102202-3121323120301202-3201331203010103-1321223103002313-3110033132022301-1102300212301323-3130230302010323-2311113332301010"></a>

<a id="canonical-1223130002012111-0331323003300131-2001321103100022-3330312110230210-1301033231233122-3201022120031202-0103322112120313-3200011323211020"></a>

#### `private_connectivity.cloud_link.namespace` property

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

<a id="canonical-0303001311113201-2100003213330112-2111230133002301-1011212132331303-0131311223322113-3233310020100203-3200033221011232-3211202121021303"></a>

<a id="canonical-2003031133033133-2330222230333032-3222330001032200-3100113122320333-1130313023120231-0232201301312002-2011021123320031-3102221211213000"></a>

#### `private_connectivity.cloud_link.tenant` property

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

<a id="canonical-2010132320123203-0002302222122002-0130011020313231-0202032223312101-1230131023013123-3313020112213310-0303101110111311-2032233030300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_connectivity.inside` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- private_connectivity.inside

<a id="canonical-2113000200131023-3030031001000010-2031030333032233-1130320311133321-1001132220112220-3112020210200002-2223233211113211-0131103322131020"></a>

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

<a id="canonical-1323110210103000-3202031100211003-0332312012313311-2132130311232210-2023023030200222-0030112130333122-0131222013232010-2321112333022001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_connectivity.outside` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- private_connectivity.outside

<a id="canonical-1310323331313330-2313213231101102-1002212000230223-0302233232112023-2001011101231112-3011331220000223-1302130030333120-3200112231331203"></a>

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

<a id="canonical-0202131032330112-1333222111331323-2111010000110002-0310200320103110-0231230201000123-3131020001033213-1033001213101002-0122303031233103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sw` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- sw

<a id="canonical-3121000022210322-1121332102321311-2330332310212001-3020230120010333-3233120322221131-3221030110012133-1312011032321302-1100323303132121"></a>

Type: `"single"`. Computed.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

<a id="canonical-2312002012331002-3021133322201011-0230322221121220-2313020213321010-0112312231032330-0100231231220113-0113001001233020-2221120300033222"></a>

### Direct properties for `sw`

- [default_sw_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-3120222200220310-2320300020313321-1211200021031130-1233333330232332-0311032120221132-0311303311103321-0210021300213131-3221320102231110): complete subsection reference.

<a id="canonical-0330021302200232-2310211202123212-1331202201113120-2131102000102223-1310033221203321-1303033021202233-2230300133302213-0021203312101100"></a>

<a id="canonical-0123220333002210-0211303300222022-3121110023132101-1210320223212202-1213313212002313-0102001230230313-1133001213110303-0121113012033021"></a>

#### `sw.volterra_software_version` property

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

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

<a id="canonical-3120222200220310-2320300020313321-1211200021031130-1233333330232332-0311032120221132-0311303311103321-0210021300213131-3221320102231110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sw.default_sw_version` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [sw](data-sources--aws_tgw_site--reference--group-002.md#canonical-0202131032330112-1333222111331323-2111010000110002-0310200320103110-0231230201000123-3131020001033213-1033001213101002-0122303031233103)
- sw.default_sw_version

<a id="canonical-0213122220113013-2113220022300013-1200232313120203-2121301030033221-3020132013131221-0310011121112011-3023131312311131-0120230221100220"></a>

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

<a id="canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- tgw_security

<a id="canonical-0023100022300302-1300301103130303-3021020032020322-1333132001230213-2103322223003320-2033221013203121-1133002230031233-3102111313311311"></a>

Type: `"single"`. Computed.

Security Configuration for transit gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-east_west_service_policy_choice": "[\"active_east_west_service_policies\",\"east_west_service_policy_allow_all\",\"no_east_west_policy\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]"
}
```

<a id="canonical-2121110133011230-2130332033230011-0122103312032313-2201032331232223-1303112030210132-0330232302313111-0231100211302011-1311202103000223"></a>

### Direct properties for `tgw_security`

- [active_east_west_service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-1223111221121130-0203133113223223-0102132323221213-3231020102031321-3223203211300330-2321100312311320-1222002221033331-0311332230213302): complete subsection reference.

- [active_enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3121012103112123-2223131323010000-2321213100110321-1233031122122303-1003133030001021-2302120313122322-3112033332302000-2113011220102221): complete subsection reference.

- [active_forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0200013210003100-2111311022132200-0332003130013133-2222310131011233-0212022222023211-2011330121301302-0021202320301030-1120200003133031): complete subsection reference.

- [active_network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2130222323310320-2130103102203233-3210213102330312-1113112011302122-3313132312302301-1310322312030322-0332023123122000-2330023121102212): complete subsection reference.

- [east_west_service_policy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-2111121131102303-1012030112201202-0222131300232320-2032122121023311-0103030113032032-1313101030213320-0300300013311000-3123000002313320): complete subsection reference.

- [forward_proxy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-0200331020201031-3023221303113223-1202102322123222-0130100010331231-1200303303111211-0110321123133220-3123003002033012-1003002030310003): complete subsection reference.

- [no_east_west_policy](data-sources--aws_tgw_site--reference--group-002.md#canonical-0301100121103102-3121203222311002-1103212021001311-0102032132330311-0300303133231323-0313213123003212-0230002223210331-0211122000111303): complete subsection reference.

- [no_forward_proxy](data-sources--aws_tgw_site--reference--group-002.md#canonical-1301103203102322-0133023022222333-3010101323211220-2222002221021132-2213132332313320-1022312003231022-2300100120031033-2201010011122332): complete subsection reference.

- [no_network_policy](data-sources--aws_tgw_site--reference--group-003.md#canonical-3023032101103211-1222102132121230-0311102331000023-3123210102330121-2232212002002033-0223213203212313-1100003110201222-2220221203033003): complete subsection reference.

<a id="canonical-1223111221121130-0203133113223223-0102132323221213-3231020102031321-3223203211300330-2321100312311320-1222002221033331-0311332230213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_east_west_service_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.active_east_west_service_policies

<a id="canonical-1330121210213301-1310222323032202-3323310311102020-2200111100020032-0130333132111101-2231032131121030-1310201332001110-2121012121032111"></a>

Type: `"single"`. Computed.

Active service policies for the east-west proxy.

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

<a id="canonical-0322332030213221-3333020312102122-1212200110201130-1033113200301120-1102001301313313-1330320132122221-3320001033323333-2213023022220222"></a>

### Direct properties for `tgw_security.active_east_west_service_policies`

- [service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0210130022232311-1230113111213300-2133210300222022-1322202333000210-3121311122301110-0332200033230223-1132212012321002-0022313111322122): complete subsection reference.

<a id="canonical-0210130022232311-1230113111213300-2133210300222022-1322202333000210-3121311122301110-0332200033230223-1132212012321002-0022313111322122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_east_west_service_policies.service_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [tgw_security.active_east_west_service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-1223111221121130-0203133113223223-0102132323221213-3231020102031321-3223203211300330-2321100312311320-1222002221033331-0311332230213302)
- tgw_security.active_east_west_service_policies.service_policies

<a id="canonical-0023300202230132-1213113002112211-1202311310233132-0312223330202013-3230303223120001-3000130201211111-1130012002321020-3102311111032133"></a>

Type: `"list"`. Computed.

A list of references to service\_policy objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3010032102032032-3102230212303231-2230201332003303-2200023221003323-0023013331003130-1220020031131201-1220303111000303-1330310113023213"></a>

### Direct properties for `tgw_security.active_east_west_service_policies.service_policies`

<a id="canonical-2300210223323132-3313002130320201-1030201002213231-0303301201001001-3111300003120320-3020131111300101-0221220122100103-0232320210200113"></a>

#### `tgw_security.active_east_west_service_policies.service_policies.name` property

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

<a id="canonical-0130222230131121-1203201023221202-0010222302100210-2131231312003300-2332303113222313-1312020213310011-1031110113211013-1133032222311030"></a>

<a id="canonical-3231123231320333-2100001312032100-3102233122301321-1313303023333232-1103033203010201-2231302123020223-2023331320310121-0111131113232112"></a>

#### `tgw_security.active_east_west_service_policies.service_policies.namespace` property

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

<a id="canonical-0022003203333222-2210130213112320-1300320303103010-2021203111322003-0300211231203121-0011012331130010-2122003210222003-1113310111023330"></a>

<a id="canonical-0223212322001133-1003030222222001-3130331303233032-0322101003000110-2003131103022311-1113102223121200-1311232132022301-2302101222022330"></a>

#### `tgw_security.active_east_west_service_policies.service_policies.tenant` property

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

<a id="canonical-3121012103112123-2223131323010000-2321213100110321-1233031122122303-1003133030001021-2302120313122322-3112033332302000-2113011220102221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_enhanced_firewall_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.active_enhanced_firewall_policies

<a id="canonical-3100122300123201-0233210101101332-2032133310332201-0300320223332330-3200233303333021-1030110010222113-0321103313321202-3330202320330210"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

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

<a id="canonical-0230231232031023-2102300130332001-1113213331103203-2103013300030322-0311011022212102-0333231001003123-2313320011131112-3002303303011230"></a>

### Direct properties for `tgw_security.active_enhanced_firewall_policies`

- [enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2200310110303210-3110231030130313-3130133122320223-0212101230223100-3221100303233120-1112011311313313-3101302210101311-0011031230101131): complete subsection reference.

<a id="canonical-2200310110303210-3110231030130313-3130133122320223-0212101230223100-3221100303233120-1112011311313313-3101302210101311-0011031230101131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [tgw_security.active_enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3121012103112123-2223131323010000-2321213100110321-1233031122122303-1003133030001021-2302120313122322-3112033332302000-2113011220102221)
- tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-3311213201332131-3213123121330331-1300121311023210-1332223320132301-0222323200222112-0111220013320122-3023300131232031-1000110333232202"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1121200130332033-2111323301220202-2313012011210102-2110201323331322-2313103133123221-2103323302331102-1200010211102230-0211123223031003"></a>

### Direct properties for `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies`

<a id="canonical-3010323312003203-1230332212330132-3203200313101303-1013321303231302-1303122010103021-1131201010202322-0330311022312110-0010113221001020"></a>

#### `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name` property

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

<a id="canonical-0021201300132102-1033001333301030-2100222303012100-2012300120100133-3000300111331202-0212112023323221-2000001110213232-2130121112323313"></a>

<a id="canonical-0121311101220203-0032311332320210-0202310122223200-0222023112200222-1331031233001210-2131021330203000-2022311201120223-3201100001323200"></a>

#### `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` property

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

<a id="canonical-2331103130010323-3130111132200330-1220112112021230-1313101012330230-0322123320300202-3322200302213133-0212023012032010-2110201031210310"></a>

<a id="canonical-2303200331031011-1201122310311302-2001230030120201-0123310220310032-0002301000110312-2201002301322102-1331332112010102-0010311011031011"></a>

#### `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` property

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

<a id="canonical-0200013210003100-2111311022132200-0332003130013133-2222310131011233-0212022222023211-2011330121301302-0021202320301030-1120200003133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.active_forward_proxy_policies

<a id="canonical-0033102321332031-2332102301121223-0313232211023310-1333321222213222-2323210303021100-1232122101112302-2000133233122212-0213120201022001"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

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

<a id="canonical-3110012113111120-0132211331320201-0121230101031032-2121013300233233-1032323211230012-0031212200333310-1231131100021312-1230000210110123"></a>

### Direct properties for `tgw_security.active_forward_proxy_policies`

- [forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3233302333102031-2222230103310213-1022332121222030-2322212132002211-3110113130230121-0003333232013100-1320213103220020-0120103123210220): complete subsection reference.

<a id="canonical-3233302333102031-2222230103310213-1022332121222030-2322212132002211-3110113130230121-0003333232013100-1320213103220020-0120103123210220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_forward_proxy_policies.forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [tgw_security.active_forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0200013210003100-2111311022132200-0332003130013133-2222310131011233-0212022222023211-2011330121301302-0021202320301030-1120200003133031)
- tgw_security.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-2112321233233312-1331021033300013-0121102131110002-0002213231223000-2102111010210002-2210100012200132-0000132223211122-3311302100310103"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2123303333320232-0112010200101221-1113231213112002-0322221223233123-2021321012231302-0313120301031321-1230213223332133-1103002211131232"></a>

### Direct properties for `tgw_security.active_forward_proxy_policies.forward_proxy_policies`

<a id="canonical-0101232113031113-3303103321202312-3321311030230001-0210302133011113-3320003223332311-2021222320333133-0201200132023312-2013212133012211"></a>

#### `tgw_security.active_forward_proxy_policies.forward_proxy_policies.name` property

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

<a id="canonical-2320121310321232-1111130131203130-1122301313230332-0000310233130122-0230110132321011-0132003031322313-3012313021023200-1130001130313331"></a>

<a id="canonical-2011233130333102-1210212130230321-2110033001003302-1120013112020032-3112222101113001-2321122331201032-2000012123220002-0230020102132030"></a>

#### `tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace` property

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

<a id="canonical-2022002112133322-0001323210012012-2122131022023033-2030323322033323-3230220030202231-3131031133111101-2131031212213311-1231303211030322"></a>

<a id="canonical-3230313111330020-2331101133003130-1122301323033133-0330303332333103-1023023000122121-2310221110000010-1110121021320321-1120223302220123"></a>

#### `tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant` property

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

<a id="canonical-2130222323310320-2130103102203233-3210213102330312-1113112011302122-3313132312302301-1310322312030322-0332023123122000-2330023121102212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_network_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.active_network_policies

<a id="canonical-1222303002231310-2201021313212000-2230020330321333-0103103301323010-1311200333300131-2003333010222302-1233201021213103-0013202220231320"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Additional upstream details:

List of firewall policy views.

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

<a id="canonical-1221222331030210-0321123032023133-3310211201021010-1320312231101121-1312121032003130-2011023100313130-2111120120310033-0113233023331101"></a>

### Direct properties for `tgw_security.active_network_policies`

- [network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2330203012312330-0323122102031213-0011320200200001-3310301012103122-3021102202001312-3020333312013001-0301102100311301-1021010130300033): complete subsection reference.

<a id="canonical-2330203012312330-0323122102031213-0011320200200001-3310301012103122-3021102202001312-3020333312013001-0301102100311301-1021010130300033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_network_policies.network_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [tgw_security.active_network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2130222323310320-2130103102203233-3210213102330312-1113112011302122-3313132312302301-1310322312030322-0332023123122000-2330023121102212)
- tgw_security.active_network_policies.network_policies

<a id="canonical-1121302101300001-0202001231021112-1133232032322033-2212222312131313-0013221133130332-1202112133200120-0110000130321123-2213121112123313"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0212120230121310-2111222330313201-2222201203003023-2232220320100100-0332100300213312-2133003001223003-3220223132022301-3203113101130013"></a>

### Direct properties for `tgw_security.active_network_policies.network_policies`

<a id="canonical-2222103201031010-2021112100232323-2312023031220012-0012202330003132-3011313120333133-1311231010020103-3023303100100332-2033232032313113"></a>

#### `tgw_security.active_network_policies.network_policies.name` property

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

<a id="canonical-2130022010310232-0013213301223112-1333023200203321-3300030300102313-3331312033233303-1310202032231221-2311113013201133-0021310010322203"></a>

<a id="canonical-1133320132132003-3131220232223210-3111302231332002-3000302311111333-3322113111013001-0120320222012112-0301113010211111-0130320003123220"></a>

#### `tgw_security.active_network_policies.network_policies.namespace` property

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

<a id="canonical-0000101223311230-3112013221302121-3003231312211322-2010132220131122-2103100220123033-1130020230102113-2120333013213320-2120232003320200"></a>

<a id="canonical-2302023122020312-0223330300213200-2332102022231200-0110111222322111-0312231302231221-1001303321012032-1332232200113231-2312331111031201"></a>

#### `tgw_security.active_network_policies.network_policies.tenant` property

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

<a id="canonical-2111121131102303-1012030112201202-0222131300232320-2032122121023311-0103030113032032-1313101030213320-0300300013311000-3123000002313320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.east_west_service_policy_allow_all` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.east_west_service_policy_allow_all

<a id="canonical-3331033230110031-0230010010022130-3101101003311012-0111122010311331-2232323013303200-3321223132332001-1302302302000103-2002310323013031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for east west service policy allow all.

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

<a id="canonical-0200331020201031-3023221303113223-1202102322123222-0130100010331231-1200303303111211-0110321123133220-3123003002033012-1003002030310003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.forward_proxy_allow_all` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.forward_proxy_allow_all

<a id="canonical-0112211023013312-3330223220001323-0232131133032102-2222132133003102-2103033103232202-3003213132210012-0120320112213113-3033320133131220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for forward proxy allow all.

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

<a id="canonical-0301100121103102-3121203222311002-1103212021001311-0102032132330311-0300303133231323-0313213123003212-0230002223210331-0211122000111303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.no_east_west_policy` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.no_east_west_policy

<a id="canonical-2331002212110332-2000100311121000-3311310302011001-1221132123002103-1301013220332310-2100320330222211-3132010312103123-1200002303212020"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-1301103203102322-0133023022222333-3010101323211220-2222002221021132-2213132332313320-1022312003231022-2300100120031033-2201010011122332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.no_forward_proxy` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.no_forward_proxy

<a id="canonical-1023311231031313-0033101023212023-3332201011001102-1101033121233122-3233303123232221-0311331213021211-0210132002102121-2022313001210110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

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
