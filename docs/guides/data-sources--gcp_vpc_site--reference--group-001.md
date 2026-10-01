---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331303113032001-0110210311230030-3200112133312023-1320133203211031-3013133123100311-1203010320002033-1012013313011023-0022112121123032"></a>

## Property reference — Property reference / 202011020233 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- Property reference

<a id="canonical-3213110102332023-1310032320011012-0031211133200220-2112230211013030-1332231333330320-2021022202002223-2201302233220132-2322300112101230"></a>

## Direct properties — Property reference / 202011020233 / 3

<a id="canonical-2320113301311111-0002300000123330-0323303012120002-3220300323030100-0133323223003220-0122121112130120-1300330030313210-2112032312010032"></a>

<a id="canonical-3222302330023023-2121211002110132-3001233233321031-1013033112302132-1133121303120333-0330312201032333-2033132113233031-3123333302233000"></a>

## address property — Property reference / 202011020233 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1122011033112013-0023030003113122-3210112130312312-1032332312323121-2102122211113223-1311220103313223-0301102033030320-0121032031032011): complete subsection reference.

<a id="canonical-1123122331230020-2330310332110333-2111100201102030-2320001331303232-1323212231012302-0022212033132222-0022330323113331-3202221203221013"></a>

<a id="canonical-3220312301330132-0113103011221023-2212332233020203-0320203232010213-2303311002023212-1003313312130110-3200132331302233-0033010223133103"></a>

## annotations property — Property reference / 202011020233 / 5

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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

- [block_all_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0113100323303113-0310110023300331-3103003023231223-3321211023203020-3010203011100033-3323230220130232-0302012300333301-2211200122013012): complete subsection reference.

- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0112133220222300-0000323123001311-2213320321113330-3300321323210101-2300111003213112-2130320223132032-3002112220002022-3201011103233301): complete subsection reference.

- [cloud_credentials](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2111320020011303-2323022312120210-1223222021022201-1120210103101211-2323012203320133-0332130303010223-0211232311122123-1120013213330001): complete subsection reference.

- [coordinates](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0322002133323313-1221023113310113-1232320212221333-1132003133221100-1330011002121300-0222310031323213-1210211311101320-3212100232200210): complete subsection reference.

- [custom_dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1212103203120331-0311103230311123-2000010120211201-3231323133012121-0201022203101321-0011303231201302-1302030121312313-2030312232203020): complete subsection reference.

- [default_blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0321222301321021-1113000001310201-2120223011233031-0320301313032331-1012313112133021-3030102201310311-3331001013221102-1222010313100301): complete subsection reference.

<a id="canonical-3203311023303202-3302333021010122-3312301213121322-2230233323333110-1101023232312031-2302010233331200-3231211221320030-1100203120110111"></a>

<a id="canonical-0222230230011101-3001123322210131-0323032223211310-0322301312222132-0113021203203213-1202001223202000-3112101023002201-3033002000133112"></a>

## description property — Property reference / 202011020233 / 6

Type: `"string"`. Computed.

Description of the GCPVPCSite.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [disable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3100312200010312-1300200010110230-0322012301333101-1312123133310211-0200301121321303-0122303210002102-2013123332313322-3122211123201302): complete subsection reference.

<a id="canonical-1030213233031002-0201222213002321-0012320200012300-0330230223231330-3211330221131020-2230102212100213-2002113203322112-1310023003331332"></a>

<a id="canonical-2030221100233110-0300332201202313-0032011230330221-0002232213011031-2132101032102001-3021233220133032-3030000000132120-3132132022333103"></a>

## disk_size property — Property reference / 202011020233 / 7

Type: `"number"`. Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2230311030212003-2200003030123112-0333131120113200-1312032120021222-3102300300031323-1102000322031212-3102223002032122-2323112010030033): complete subsection reference.

<a id="canonical-2032131201031322-3311221310321223-2013120122320331-0000101332203300-0310113300033121-2123312130300233-3210030310133333-0220303200221221"></a>

<a id="canonical-2013011320321003-0301200230201321-3221132133232132-0221121200303233-0221331300123121-3213313301232331-3110330311210322-1130001220230331"></a>

## gcp_labels property — Property reference / 202011020233 / 8

Type: `["map", "string"]`. Computed.

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

Upstream description:

GCP Label is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in GCP console.

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

<a id="canonical-0333220100212220-0230010113330013-0333131113012322-0021030000031031-2133123102232203-3132023203233121-2131013203102301-0302033100023211"></a>

<a id="canonical-2101232233322102-1320020110333133-1210221212301301-0223323301311211-1200201100111232-1311311112102111-3010201311333222-2220021222210010"></a>

## gcp_region property — Property reference / 202011020233 / 9

Type: `"string"`. Computed.

GCP Region. Name for GCP Region.

Upstream description:

Name for GCP Region.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3331123333200113-0303302321332303-0211222021223030-2302020210332300-2311210201230202-0301300131003231-1213310133123300-1232312220120310"></a>

<a id="canonical-1231333033131301-1023311333032003-2222000131333323-0310022011313223-2301013330130122-0221212010301200-3133011203032011-1323031003133233"></a>

## ID property — Property reference / 202011020233 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031): complete subsection reference.

- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1203301201320331-1230203103032032-0033211230213112-1103110310132010-2311223132020020-1113110103030030-2212320302313110-3111021020113222): complete subsection reference.

<a id="canonical-2213220100123022-3112000221011111-0031333211321330-3000132101201123-1202233300100301-0112032210112012-0020002202131322-2310120110032223"></a>

<a id="canonical-0332212002111121-2031201001333330-1133002212330221-3002211303202033-2000001322333200-2302133012213033-1101103231322301-3123221232031121"></a>

## instance_type property — Property reference / 202011020233 / 11

Type: `"string"`. Computed.

Select Instance size based on performance needed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1022113113131303-2323310123132033-2333330201223220-0203120001033203-3132021212231333-3113000331201222-3002010313110322-3330113112120320): complete subsection reference.

<a id="canonical-0030113023031300-0113220000110121-2302103003120033-3213123300300223-3032222330003332-3323021022002113-1110322300122211-1302232003332132"></a>

<a id="canonical-1311123232213201-2130200303333022-2231300033322003-1033321001102302-0033322020323303-1232103232133021-3232301221222320-1311023211222110"></a>

## labels property — Property reference / 202011020233 / 12

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

- [log_receiver](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3012023010120110-0132220330001122-0031313332213232-0130033203331102-0312222230232111-3100031123133102-0313211302213130-2122312222311212): complete subsection reference.

- [logs_streaming_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2202020213103112-3302231032102222-0001333100102011-2232113332301013-0123022110213010-2310010331223201-0233011232203110-2120130220331303): complete subsection reference.

<a id="canonical-2211221221131011-0303102211331030-1202221302320132-0230110211333131-2322030221022333-1331331302133102-1020203112323131-0002230213322132"></a>

<a id="canonical-2312233220130123-2323303122223133-0012013302222220-0323202003133100-2130032002011321-0032202030132310-2013003230013231-0232031222130312"></a>

## name property — Property reference / 202011020233 / 13

Type: `"string"`. Required.

Name of the GCPVPCSite.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1230130000200211-1232231202020113-2302130211313322-1330330210030201-1211110320133122-0000333302221330-1233303223033311-3033211201223222"></a>

<a id="canonical-3031321310030003-3320223211133011-1300302301233222-2000301032113332-3233321121013120-2132123103103131-2120111103131333-2310200121002301"></a>

## namespace property — Property reference / 202011020233 / 14

Type: `"string"`. Required.

Namespace where the GCPVPCSite exists.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1110001323312320-3212002110322030-1220001101201303-1333110011111112-0302100001333323-0132330102013113-0202320003302122-1131123010110312): complete subsection reference.

- [os](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2312120333122222-1322013012302121-3031012213121300-3212110230333003-1231033133000002-3321231322032311-1300102131230223-2103010231021132): complete subsection reference.

- [private_connect_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2213201100213331-0202020301100211-1311213110003020-2103223312210301-1113131132003333-3120221123300130-2032302102110033-0222302310211223): complete subsection reference.

- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1203003321300011-0223032211022011-1000102131312201-3122010333311010-1021023223222201-1131021212120001-3233201320331210-2110010200210122): complete subsection reference.

<a id="canonical-1131200321210212-0100100122002120-2311021313321002-1333101002100203-0130032133330011-1213231202000101-3222013031313032-0011020301132002"></a>

<a id="canonical-0110122332221300-0311320122203021-3033003133321333-0032021012132020-0303302003103303-3311032302203103-1213332330202300-2332021110000321"></a>

## ssh_key property — Property reference / 202011020233 / 15

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0310330212021223-3303111001132333-2110220022330203-2030023332031212-0003311012011331-0210223302200113-1032100302211312-2001221222220020): complete subsection reference.

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220): complete subsection reference.

- [waf_signatures](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0012012033031300-2110002220223133-3303202210233302-1331310110333111-3031100320201031-2130003013122113-0123330322110322-1011031331030122): complete subsection reference.

<a id="canonical-1312330210333113-0130230133213230-1133300102120001-1000233200302313-0110231331200223-3120201333221302-0210321000110200-0101012030023233"></a>

## All schema paths — Property reference / 202011020233 / 16

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2320113301311111-0002300000123330-0323303012120002-3220300323030100-0133323223003220-0122121112130120-1300330030313210-2112032312010032) |
| `admin_password` | [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3033323011023010-0300012110310133-3322121202001221-0000212122220323-0131131032130010-3223232213310202-1213310212033301-1130221100330102) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0123203133030211-3003033001331113-0232300000013323-3230132120131300-3010123020120103-2121032111100103-1132132032311131-0311132012121023) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1133321132300131-2310133030330312-3121123030232113-0210330233331100-3330210211022320-3120103113300310-0202200322230223-2301311230000032) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0122001132313003-2011233300323210-3100013003322012-2232233133323202-0333013220201013-0213331331313233-3212231131000103-2300122131013231) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3131133132102300-3122111320110012-1011103101123021-0022111330101032-1033010112320200-0211121022333033-3220203321023102-3333113130022132) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3230313103123211-0033130310330212-3003313222123112-3120220211110332-1331102322323310-3122100122021232-2022023023011323-2113303332123311) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2303033232331212-3123101230112312-1201203033033102-3013202332000310-1002301302121030-2310302030111123-1300330000202002-2012303311200211) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1311222210002333-1212000013222133-1323200011101122-0022320020223210-3013303233020012-3223231130012001-3100220111301000-2103022031230131) |
| `annotations` | [annotations](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123122331230020-2330310332110333-2111100201102030-2320001331303232-1323212231012302-0022212033132222-0022330323113331-3202221203221013) |
| `block_all_services` | [block_all_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2321132300131021-3120103231201003-0132232310133110-3332101201330320-1311220331223123-1323203303130321-2023300210122320-1102303210032113) |
| `blocked_services` | [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2300113002021223-1232203130003210-3021221232212012-0103322233302132-2212120121101103-0312330301111031-2232120212102203-0302011032303022) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2220331212111013-2311331210033001-1121311301203031-1200133031312022-0012202221222322-0231001110111120-3320013031111302-3003330133103112) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2102333102031100-3233222311023001-0023121210011110-3211231233123223-1131300332202001-0031220323111101-2023032012321322-0021301201030320) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3221132120121110-2100001221102032-2200311203033101-1103130012233100-3002233101030110-2013331023302012-3312220131320310-2110302201232122) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3011223012123301-1122110132031111-1212223302301133-2213323122330210-1020000332300020-3320222000311230-0221200330320230-1321103213002333) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0201033311102320-0301310201030231-0022020332122000-3103130222010012-0203323332023310-2122300220211333-0031123122333022-3200320010013032) |
| `cloud_credentials` | [cloud_credentials](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1031211303000002-2120012302332133-3011302012221322-2020003301313100-0222102033023012-0302133032323321-1200033100022333-0220211331120320) |
| `cloud_credentials.name` | [cloud_credentials.name](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1010132301132200-2001010021333231-1203211322323012-2113210213002322-3201101321022322-1201221122333022-1312132212023032-2333300230111321) |
| `cloud_credentials.namespace` | [cloud_credentials.namespace](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1013313100030010-1110133200232320-0211231221212203-3212110202002303-2010031022321100-2003312310200103-2333133103202300-3301331212212022) |
| `cloud_credentials.tenant` | [cloud_credentials.tenant](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3210302021110102-0003023203013331-3130221113012003-3101132213120123-3133310030030213-1201132003111133-3330300130111123-2231210231213322) |
| `coordinates` | [coordinates](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0220231032011300-2312010103121300-2233002220222030-0312322002111203-1310332131102213-2232112010103130-3212030211101111-1110310002131303) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0320230333031323-0133123000233320-1210230001330332-0220133011200133-2310121231103012-3202011122310000-3112101200333133-0223210332203013) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1110300022130031-0000111312223101-3012032023121030-0333030200210313-2331012310100000-2011330010112012-0201210123310021-3333223201320233) |
| `custom_dns` | [custom_dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0033221203121322-0233123323010122-1003313302220131-3221003323102322-0113023302112330-1123313112033112-3111122021303230-1011332000232230) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2221321132232331-0132120011302103-0221320230031003-2112320000311031-0021001021123332-1201301321303100-1230021311120033-1303020000223201) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2012222130203300-2202332310113033-0111201002333032-3331301331002200-3212023030321200-1213211103001032-0213122323030023-2102330330123203) |
| `default_blocked_services` | [default_blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2301211332100101-0233022321011302-3101330201013312-0302321221010310-1100312212323232-0311300230010000-2022312303321300-3012222312231322) |
| `description` | [description](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3203311023303202-3302333021010122-3312301213121322-2230233323333110-1101023232312031-2302010233331200-3231211221320030-1100203120110111) |
| `disable_encryption` | [disable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1033112032333201-1033213312021303-1221112331120233-2330013211032333-3300000311123110-1331213100130132-3112011212321322-1212110220200100) |
| `disk_size` | [disk_size](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1030213233031002-0201222213002321-0012320200012300-0330230223231330-3211330221131020-2230102212100213-2002113203322112-1310023003331332) |
| `enable_encryption` | [enable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1001213133130022-3301300111031002-0300311200120333-3001120320012320-3013211321222111-0132133302231310-1201000121233101-1202131020033233) |
| `enable_encryption.kms_key_resource_id` | [enable_encryption.kms_key_resource_id](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2121001003110200-3321023112023323-2322333203122130-1122000223001333-2231313331101121-2311201230331301-0101313131022103-0322021013203313) |
| `enable_encryption.kms_key_ring_id` | [enable_encryption.kms_key_ring_id](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1010331221321311-3033332202022132-1221020001232130-2001223221322110-1311303300000332-0031302332210123-0131112233112110-1123202013103333) |
| `gcp_labels` | [gcp_labels](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2032131201031322-3311221310321223-2013120122320331-0000101332203300-0310113300033121-2123312130300233-3210030310133333-0220303200221221) |
| `gcp_region` | [gcp_region](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0333220100212220-0230010113330013-0333131113012322-0021030000031031-2133123102232203-3132023203233121-2131013203102301-0302033100023211) |
| `id` | [ID](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3331123333200113-0303302321332303-0211222021223030-2302020210332300-2311210201230202-0301300131003231-1213310133123300-1232312220120310) |
| `ingress_egress_gw` | [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0313012223320302-0022031003132120-2102001313322132-3131121230300302-2221202331000301-0133330100301331-2203003130202202-3103202120020301) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0011220003212110-2112122203023232-2022222130132320-0133233300323002-0133221111331002-3102203113032303-1212301213110033-2020200302131313) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3202013331021013-2210210233100202-2120221222333113-2031031131203300-0112222212312130-3133013103130010-2310132100123120-0202302112013323) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2012031002300210-2031303300132110-2023301013323112-2001322212010320-0131123203022131-3103012011101233-2212132323201102-3222212331330032) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2100112101233103-1323232110213301-3120332031113101-2112112123100222-2330220213032133-2003211233022220-3232101213330003-0113103330220020) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2220031031033123-2122213333313302-2133212211000201-1020233020321123-2300332201300100-3110113122303100-0320133032310110-2023033223112202) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1221110013213020-2110232331123000-2112302122103212-0312200021222233-3322022200111321-0220112103210200-0000332211101133-0030012100120023) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2310112330031102-3120302213200130-3101310311011333-3301213232212003-3013223201110223-0113102100030301-3000112102110313-2303332002322022) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3132102033133330-2002233111300020-0221112103311230-2003223321113300-0331213322022011-0331030012133123-3032032012120000-3230312312322211) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2110100003323020-3212120122002121-0330102020030021-2303211303201310-1031220233233331-0031133313021302-0012323033231111-3120130220011000) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2203131330033320-0300213023321223-0012312110331301-1121231031011003-3031121011323203-1321133230303100-2331001110021331-0320103121033330) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2231111100313323-2313020120032212-0033131221311330-0123102201101211-2313023033032222-2003333111332211-1223232103100011-3110022122121322) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1213002002220332-0131122033012131-1011123032302203-2110003020233330-1032200323313011-3320113013002023-1123030010300222-0210312032132033) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3223231131301220-2022302332131102-2322231132232110-0233030331123020-0022333033133031-2103322002001102-3103330122230223-2003010000031132) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2321022013003021-2101323321132213-3001010011302012-2211312330120020-0001010002210310-0203322321211020-3333210130021030-0213310311312132) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1022220131330112-3203132210311030-0112212100220233-1131111200301313-1013230103133302-2011000312232312-0133002330130133-1102201323221011) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0122210210301102-0310313022012330-1313320110222103-3221031213200221-0320012323023032-0332202210303322-0210202213011002-3010232230210200) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3011023202323032-0313220122123310-2123302012222100-2113103002231032-3221301322212223-0213330102123032-1313031232100203-2022023102122100) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1102201102033212-3223223012120332-1313122001222220-2103322100103202-0213123003213333-0331011211332230-3300133000002020-1102332131323333) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2333220111320323-1333101320023110-0000003300310213-0102223321320123-2132020131203013-2132022122301123-2333300003110332-1322130301022100) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2032030201233212-1133310201222023-1100310323232103-1133131022322332-1133001212121333-2212221220321111-2202310300220313-0022203101332111) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0231132111013301-2102211323133112-2132102203030021-1021130110312100-1121023203023333-3021230223311233-1012013221022310-3030210210032210) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2223212001031020-3323313102012312-0020032332331332-2132131202122102-3232311112001110-1011232312213202-3130021213221112-1210000132203202) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1332130011201200-1322110302003323-0101231311323010-3331111102103013-3302310132301112-1331322310332022-1130313002210201-0023001112032000) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1012223021333013-3033013201110301-3233122303221303-3110201233221101-1103310013333123-3323000200013031-2222311220331211-1223210300320132) |
| `ingress_egress_gw.gcp_certified_hw` | [ingress_egress_gw.gcp_certified_hw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0213113302333011-0220223322322001-3013303010131032-3002013312201012-3310100303221303-0333113330003331-1020010103213120-0031303001313023) |
| `ingress_egress_gw.gcp_zone_names` | [ingress_egress_gw.gcp_zone_names](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0103031030102332-1302223121311220-1133313131223120-1123133110020130-0333210231121002-2132020112021133-2023210231012021-2012000202100223) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0132031033010220-0022031212001320-1132002213322013-2133020323001310-3213012203222220-3133032312103301-1223321221112202-1310123112020132) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3222122122102121-2301132320202131-1011122131303123-0011210202321203-1100003110013211-1110233111320101-1100231220210311-1031102111010003) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2331003130102333-2301322331003201-0112213010332233-3203131030003111-3210232210301311-0110112112013031-0201022022112200-0021222300301221) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3302103113202030-2033003301000202-3302211232123023-2320030321330333-0311010211322322-0031111132311322-2213330100310210-2000022022212111) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0100302310013220-1011013130203202-3320300230310222-2002131121023330-2030100101131121-2310023233301112-3030000011200120-0100200112020333) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2202311201300030-0120212220201021-0313200330232003-1213332212013103-2130211303122322-1330212300101202-3223023132311010-0220020201233210) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0000030113321100-3210132101003301-2303131230303122-0203112133022032-3231111001313210-3122222132111000-3111001220122002-3110132201133031) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1321300321231221-3201210130020211-2130200002123303-1312123311123302-2001130031121322-0013212030331233-3311103331222112-0032231310100123) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1313023012101113-0011103032321200-3223022222102332-0001003100122123-1033012331130232-0120332221331123-0133323031003123-0102013320221311) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2132321002220031-0112023011110111-0133012313200333-0011003222000310-3023100222012330-3001301332023020-1013132202102121-1102021011101330) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3313230330203201-1013223120012102-2002102110010300-3101233123102331-1200032202022321-1230311203321013-0221223110301301-1210022020322110) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1133033110203021-0220123033000002-1220120222001123-3031100011302302-2302231111203112-1300121120110223-3330230003100110-2321220130321222) |
| `ingress_egress_gw.inside_network` | [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0100211132231310-3310032001321123-1013102103221111-3221220201002330-3031200123011313-3222232120302230-1303003031000232-0221020313323031) |
| `ingress_egress_gw.inside_network.existing_network` | [ingress_egress_gw.inside_network.existing_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1220121103122102-3103121301131010-3231023011230101-3303131303211232-2113213300302312-1030100112201232-2333212110212110-3130210223210203) |
| `ingress_egress_gw.inside_network.existing_network.name` | [ingress_egress_gw.inside_network.existing_network.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1100203330003302-1203213110222102-1103000211330132-3032021132033230-1312002212323231-2313222101023303-2121102233223221-0331230030213130) |
| `ingress_egress_gw.inside_network.new_network` | [ingress_egress_gw.inside_network.new_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2210300111302201-0201113302221132-0132000002112331-1311313110111102-2323132213031100-3310231002121232-1033031203310321-2022102203121210) |
| `ingress_egress_gw.inside_network.new_network.name` | [ingress_egress_gw.inside_network.new_network.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1022002303201230-0102322100031231-3011323110211212-2302133030123203-3312013333033332-3122213220313230-2101332321100020-1030222202102202) |
| `ingress_egress_gw.inside_network.new_network_autogenerate` | [ingress_egress_gw.inside_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2221101310120210-3103010101232031-1303203121200312-1131112222033310-0201330122222120-3120230200032331-3133002122132231-3000331213330220) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0330302001121010-0000023223330023-1311322132211002-3110330223133003-3210302123201210-0131230101121200-1033000331322131-2033330033032113) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1133110212121311-2212112003103102-0000312111112121-0211310132101302-3203232220132101-3030203111200122-0011110120121120-2121210301211210) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0011003110330210-0021131012120112-1221102110302032-3333000133113130-3023033303230023-2033312202122002-2221023221112123-1122130211012331) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1121211012131102-3032012301112330-0202321233002230-2232303200023021-2311020212100033-1111210032023000-3302131211312032-2211310023300300) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0111323000131300-1132031131213102-2303330201102110-1002310223103331-2120112000011031-0023031132220102-0130132120323321-2201333003010203) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2210033313010300-2111131021203311-3133311323312203-0011230211132031-3120332213003022-0302122333003031-2231203323300121-0011003132200323) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2101002203120322-0110323120300202-0103013222203002-0131301322222303-1232113130311123-2232003202021321-3020223133133311-2100231113321020) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1023313320222021-3103031321023330-0111230331223132-3132300030010002-3001100023130033-1321000321011230-3321023302300132-3120223003213100) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2020122031212133-1021010302012101-2210121120011302-2230001313031030-2320322121010233-1213310221220030-1321103300313111-0120130113311110) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0131222332131201-1003012220322001-3020110000022001-3233332212123233-1101133130011013-0133233330213102-1003121230231230-1020132023331031) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1132321012203333-0323321301313012-0020012301112222-0131123102322323-0012312322221230-1230000311233012-1021103030020131-3033311003132311) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1110121210030010-3301012032213333-2323323003211010-3222102012120100-1120111203113230-1131312110032131-1131022322120031-0312312101330220) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2101302302032302-0023333230100010-1101023031330232-3121120313210231-3123103132003103-2323310020000122-1131221220032222-1200032122110122) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1033011200001310-3333012230203311-3012010132310121-3303222131230011-3300220121032210-1033131130130103-3130001221212003-2222002032221301) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3303122120233032-3322013203330123-0323301312030022-3122001022033232-3031320121030102-0323220212030222-3301123220023020-1032312322111333) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1110022101220222-3033312102022023-2232132120332302-3323330313311111-0322102233121301-1032211213031213-3321120120001100-1313303210001003) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321030110132221-2323320002131110-2300300103222330-0223302200200111-3120232020021020-1310231301330003-0023221312302002-3203331132012322) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0123132033211033-2102213201101022-1033101021130302-0023001313132102-0212301123023323-2111100031323113-0013001001002302-1132003201331013) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2122331232333120-1301020213313032-3000312000232000-1101132132200302-3220012333223301-1122121003113303-2122122010022111-3113122203023113) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2321010222123213-3133001222120220-0113100120133012-2202010200213130-1300010031312122-1213002002233230-3032113112333300-0232130312131320) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0303023121120203-3133300221131122-1212203103003312-1231122032021012-1210311111110020-1132112313130123-3200220313120220-3300030112220132) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3320020103101312-0111112122320112-1310311213310220-1213112022132020-2232012320233101-3030210230101201-0220102103231210-3200101101211003) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0333023003123210-2133331231010203-1200111222022022-0022020033312331-1103033232100102-1302230311030130-1110201013333013-3230312102233111) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2130122032312300-2232300321103002-1021101122010022-3021330113131100-0330011210312122-1300001020312011-0313132000032310-3013302001010330) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0330322031201032-1203302302103133-3121121012303102-0022230212030033-2100333030110321-2131301123013323-2012022312202212-2202113321112022) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0320203021023111-1330322232221311-1010312023031122-0212233301321300-3323002021210030-3011330222122133-0020000320300033-3123230001323033) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2201232223233230-2022021332230231-3201123323101101-1112202003122030-1023210113120023-0132033001112320-0102120233311100-1113301131100323) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2213232323302103-0301332323231211-0301230201030020-1010000023011231-3301120231010332-3022011032020031-3032122130302021-3302313122002332) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2012322032030221-3023110311121130-2103332020202021-0102121320122110-1000112303010222-3133332113233111-2133111031123110-3103302101002210) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1133321320302013-0303231112130101-0023111002133310-3010233330130220-2130221130002322-1013110101210212-1230100322203030-2221110112332011) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1001130210203312-0101101212231011-0311111010013220-2220101302203232-2233102110230220-2232021133203332-2333231221001310-2033221111022211) |
| `ingress_egress_gw.inside_subnet` | [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3231330330333310-0233032200002020-0320132032223232-2212122102321221-0221013303310012-0220130000213120-0002033021023220-0320002213320033) |
| `ingress_egress_gw.inside_subnet.existing_subnet` | [ingress_egress_gw.inside_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2301010221110120-0200203121013312-2221131212313031-3012103021213312-1110123012220110-0011320331230232-0123110132033331-3313130101022001) |
| `ingress_egress_gw.inside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1022023302023012-2032230123103010-2012123203020331-1123022132030233-0313101232201030-1201023201320102-1131031333100321-0330322033323202) |
| `ingress_egress_gw.inside_subnet.new_subnet` | [ingress_egress_gw.inside_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0132332133032322-0321320111010312-2322110013321011-2321313320321300-2323113101133030-2230123030320302-3111321110321330-1010031031122332) |
| `ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.inside_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0133312313130312-0323101122201131-2103112122303220-1022002030201201-2002323230013320-3112323232102122-2213100113313033-1200013322331010) |
| `ingress_egress_gw.inside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.inside_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2220123323130313-2333212230331320-2021103203222331-1103033012112311-0112220123100202-2113310312013321-1012232200002013-3222030321212200) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3213110011303132-3311232203201033-2303223001021321-0101223113331223-2001002021102011-1102300113101320-0303102000033120-2311322132130203) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1030111330000031-1013102311022312-1203213012221121-3010200202220312-0003120110223320-2123000223213321-2010211330110032-1010021212130310) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3203122321121030-0331210203131120-2310201231021033-0113022301301030-3010231123122331-3323121031332222-0330301302031332-2320212003133213) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3221332002031313-2010131221011230-1230231202011312-0202013120103102-1333022113120120-0013320312002211-3102000313211302-2000331100123023) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0132122202101131-2020132123110110-2113003000123102-2223220313032212-0310133313030033-0100311313301110-0003133333001101-3133301121303302) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0333203110320330-1132120310113110-3332012300033200-1213033032112101-2211320123122222-1230301130333211-2230110310223013-1213101023032300) |
| `ingress_egress_gw.node_number` | [ingress_egress_gw.node_number](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0233223010101003-2232200201322111-2011132001312231-0231200021103302-1321111023122121-1113313300311111-3222120012313021-1030112301231220) |
| `ingress_egress_gw.outside_network` | [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0313202301100112-1001031023230231-3333101222012102-3202133211302210-3303133320333332-0112122323102331-1102101233320033-0121013330311312) |
| `ingress_egress_gw.outside_network.existing_network` | [ingress_egress_gw.outside_network.existing_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2110022110221201-2320022330033321-3110300110031332-0120013110220221-1200332320202100-0101031321000120-1211011132222203-0310120022223202) |
| `ingress_egress_gw.outside_network.existing_network.name` | [ingress_egress_gw.outside_network.existing_network.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1123202222333000-0021203322311032-3102302020102323-1030111000220003-3010302231213311-0032221010120321-2230131332010130-0201122221120112) |
| `ingress_egress_gw.outside_network.new_network` | [ingress_egress_gw.outside_network.new_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3110302032210223-3133330210133031-2202112321030301-2001010202333221-3210002033222310-3331302330303001-1103003213010012-3312223120333112) |
| `ingress_egress_gw.outside_network.new_network.name` | [ingress_egress_gw.outside_network.new_network.name](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1311211120021301-3013220230111123-0203232332213210-0012311001130001-3013032100332023-2030020310133131-2223011321203113-2332223222031120) |
| `ingress_egress_gw.outside_network.new_network_autogenerate` | [ingress_egress_gw.outside_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0031111213113013-0103010030232011-0003003202321103-3132010030011132-0231012221320112-0331300220032221-2120210112102030-0333030113032331) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0323023232303223-2121211210022131-3300033022311100-1032131332303230-3001120032220202-1321312121020112-1202102232031111-0312000212011101) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1032020001221203-3211103022330321-1111130120110231-2211301200131021-1000121011213311-3222012231020003-3103131101000310-1322333223202230) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2103200313332222-1131023303323121-2222313212320320-2303113331130031-1300031032303223-2112212011131213-3031301123323102-2200101331133010) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1113021133222132-0233222330121200-0222023222131011-0303101213230232-2112320320012000-0032320310322333-2202322002203032-2000002032301032) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2203111211123001-1302311011012203-0302303101021333-3013003113122111-0220031031331011-2333233023222203-3111332212100030-0230300130223113) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3312030212332122-0102121002311120-0301030210002220-1000213212223312-1213022022311223-1010033033313123-1122103012030111-0200222031303031) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3030101010031321-2022031021330132-1110202320301300-0332112331321312-0233220303000302-3302030233313111-2120131202212030-0222031011031303) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2200212023232010-0201322232021112-3331233202200300-1312001221323320-3232123133121122-0022233232132322-1122021230332101-3002320310300201) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0210222222111121-1221322013011031-0103300033212332-0212023222000130-1033022123232233-0230013131023021-2311000113330103-1033331232100223) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0103111321022123-0233012200222120-3020231213001212-3113201012102113-3220121121021013-2221103232230010-1033100010002021-3301113212231113) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3133220212323130-2102100200330201-1000102210331101-2021121221113100-3101321222333023-1223130032310122-1200220231313020-2120000312032100) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3300310211222002-1230221310102213-3330021321231332-1232222131130030-3310230023110122-0210330303222200-3230133011223022-2223100333211002) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2120331203231312-1203121311303012-1330100021113112-0302102300231322-0123301333002033-0311213021001102-1301320212230203-1330302102213200) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1103021013212323-0323132032002031-3232023331031133-1110233110233003-1020123330223030-0003102331203311-3033132103033100-1031121211033321) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3211211122301202-2023012300121232-3102332302310002-2201230211200302-1120221221211103-2102232331311232-2211131103302100-3023100121112311) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3301121200323022-3301011230131322-3002011203310300-2300131013002020-0013020023232200-0313221302213003-2012021000313220-0102103131013200) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0200302101313103-3000220300311123-3022130310333220-0101213303102233-1101220003112003-3031323230100132-0120001223031130-1302311332201030) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1131102112130022-3221202211010100-3020022313110102-0121310323201131-2013123220020113-0130002020001103-2212310021111133-2232212303223111) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1000003201201330-0201222133033320-3022201131222132-0321233310030301-0111312021133003-1303231031120223-1100120002200300-2013322220332002) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0133033010112322-0131023020323020-3332233223213300-1001033232010212-2302322112210321-1130223320133320-0013020323331021-3013112232233110) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1020023133220201-2211210130003301-1131212330323131-1133322013212123-0101211012102012-0233303022103233-1312032133233211-0123003232100113) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0312202201211010-3133110010132221-3321120100020023-1031111223302211-1010031030131202-2201030030223033-1133310201303211-3113023210010130) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2311112202221121-1001102313100112-1133302031103330-3211232111310122-2211201010002013-0012311231312222-1110301301013010-2332301231001303) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1023132223232001-2012320302220032-1202112001212232-0321211100223131-3012001120112300-1022021100102102-3131112201020321-3132033111220201) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1210002031020302-2110113301020213-0201033130302002-3001321231303223-2103211023200300-3313112113311302-3231200122231120-3331112023202122) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0223031232321133-1112233223220200-3120320213030122-2021231000121313-3201202232203013-2123121333213021-3321232301210331-0010222131122112) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2211301101010213-3302311222112030-2111020223230110-0313213302322002-1133112302211101-2320103333102110-2002220110222301-3311231232331122) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0033033011110301-2213212101221101-0221331312210123-2332330121213211-2132333320312321-2213321233330210-0012211211213331-1213311301003001) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2320011233233023-3011110131310001-2330011132201300-2110302113032111-1300331221113123-2100101101011223-0133202101031033-3322022100222301) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3322133231011012-3213210020330010-3032113223002112-1020301320213321-2322330120333122-3310212210312111-3002202111012000-0331332001103122) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2332030301211220-1332310011310010-0320032223000203-2121312203320303-1013300120022111-0102302210223320-1102200301222021-2100311302300133) |
| `ingress_egress_gw.outside_subnet` | [ingress_egress_gw.outside_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3320303120201113-1231311110033312-3212301012031132-3230020220323322-2321213101333201-3200000103033131-1231212230110330-3000011133011210) |
| `ingress_egress_gw.outside_subnet.existing_subnet` | [ingress_egress_gw.outside_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0232112100322310-3000232122212131-1103311002133301-0233100322212100-3331332233023003-0113321310323302-2210332123231132-3002333210032122) |
| `ingress_egress_gw.outside_subnet.existing_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2011131132111310-2230110131120313-2322113222333032-2302211103131300-3313013012213311-0133103000000033-0200023232200303-1121110022320030) |
| `ingress_egress_gw.outside_subnet.new_subnet` | [ingress_egress_gw.outside_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1203300300330222-1102111322221103-1312103303202010-1220232002312020-1123313233132301-1110312333323200-1203233122321121-0322310323010200) |
| `ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4` | [ingress_egress_gw.outside_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1101332030023102-0020210123210113-0233203032301232-2323121331002122-1310101113003110-2031200012333212-2213133102233012-3132121130202300) |
| `ingress_egress_gw.outside_subnet.new_subnet.subnet_name` | [ingress_egress_gw.outside_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0000323220233033-3111101130133013-3323010001122020-1222230000012030-3320212013112201-2231212303023211-2022212210123202-1312101232122021) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3202313120112320-0131113222022222-0301203233302211-0013123213120332-3203200322003103-0023030210311002-3033301303203311-2221313021003310) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0320113031201111-1101031211001312-3010223331011310-1301032323320313-1131311010332020-1330110230133322-2130200021321031-2210310223313322) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2112303332123102-3333022332202200-0121023012103300-0123010202323322-2331100201312131-0002310113120311-2232332233312230-3113323133312233) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3201120212323003-0200000333213013-1301101300031222-1221010102302301-0212310002133122-2103111322213220-1233031030013013-0023211103000002) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0031332120220011-3101033033123201-1223122301230221-0032332321123202-2220300200323212-3220032223110130-3211200213333132-3211133313003212) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1321033212232331-3021222033103123-3323110313110231-0321032320030022-3133213312031331-2221203312033311-1330023321210102-2213132313313120) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0110130123110130-3232203321233100-1003221301111331-2222211330010223-0110322030232132-0302220132031330-1203020032130102-0322213302223233) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1021113303131112-2020221230112303-0122211213021120-1132121123020212-0230331021202132-0102301301300333-0120210132311113-0032113332110200) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3013322112132220-3011130301222110-0102010330123213-2302231110001232-2223031030000303-3332302131213103-3230300111013003-0220223232310111) |
| `ingress_gw` | [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2222232300212112-3211332221211323-3302333303221100-3122023230323133-1112203332332221-1323033133300103-0220100330013232-3332102330003130) |
| `ingress_gw.gcp_certified_hw` | [ingress_gw.gcp_certified_hw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1030220313203311-3030010002301131-1132230332021312-2133010101320321-1333131210102122-0232103300132233-3310120211201212-2031331120303110) |
| `ingress_gw.gcp_zone_names` | [ingress_gw.gcp_zone_names](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0120231010331100-2322020001321002-3120121230330122-3012022132200203-0011210101100301-2302000113333021-2111222203323112-0001213032011322) |
| `ingress_gw.local_network` | [ingress_gw.local_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2300021130312333-1013101303013031-2123220112122302-2130312121111103-1302022200311302-1200232332021013-3002103211022212-1111331121022133) |
| `ingress_gw.local_network.existing_network` | [ingress_gw.local_network.existing_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1022232030013331-1020202030131303-2330331103112312-1202112122012012-2031023101232313-0233223110100200-2211011033232121-1333310210101022) |
| `ingress_gw.local_network.existing_network.name` | [ingress_gw.local_network.existing_network.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1103102321310203-3013322313312312-1231332031121302-1103032223020221-0101302131012133-1330020312330220-0312033121303302-0102123011230313) |
| `ingress_gw.local_network.new_network` | [ingress_gw.local_network.new_network](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1023133201302001-0210312032301033-2202210121023110-1322021133100210-2002121223132332-0331021022023311-2102111130121031-1132232222021001) |
| `ingress_gw.local_network.new_network.name` | [ingress_gw.local_network.new_network.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3011330103013110-1002231010231020-2133211133023232-0210322102301102-0203032010211101-0120101223203303-3301032020303012-1011223311101222) |
| `ingress_gw.local_network.new_network_autogenerate` | [ingress_gw.local_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1112102321303020-3020012312320212-2103312101223332-0230133032103310-2200223233011102-2300310220303301-3003110302122322-0032203031211322) |
| `ingress_gw.local_subnet` | [ingress_gw.local_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1221333212131133-2012032302302113-0012313002313001-3220012200302100-2131230123333032-2311332310131301-2323330323020312-0311221000020231) |
| `ingress_gw.local_subnet.existing_subnet` | [ingress_gw.local_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1211022333130121-1223111312022200-0121220213302002-1233312213100030-2302100201003113-1220303203201023-1001101311010103-2000320303012321) |
| `ingress_gw.local_subnet.existing_subnet.subnet_name` | [ingress_gw.local_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3002100230130223-3212223302332033-2221100110202210-3032111110302331-2001131322301221-1002012032003320-1110302333211011-1113023113030211) |
| `ingress_gw.local_subnet.new_subnet` | [ingress_gw.local_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0101210312020130-0001001013022103-3302301100223121-3223121020111201-2121031133001002-0300111000110033-3020001133031330-3321100112332233) |
| `ingress_gw.local_subnet.new_subnet.primary_ipv4` | [ingress_gw.local_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3122100331100312-2023020211222123-2003113301011111-3311320230320033-0211013301232222-0101332222010300-0212022323201231-3210321030201231) |
| `ingress_gw.local_subnet.new_subnet.subnet_name` | [ingress_gw.local_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2033133122330300-1310230130200102-2320121222021220-1100300203032212-3332023101031301-2102313103130112-0130310030310331-0300231211232022) |
| `ingress_gw.node_number` | [ingress_gw.node_number](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3021002310321202-0303133012211221-1202300012213003-2033012212203031-1111020131312221-2302333101000023-0032113332333101-0123203000113311) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1013201030102321-2021320033323230-2023001200001021-3222201202322233-2131030320033122-2032301302323232-1310103201111320-0020023212222101) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1000121230312330-1301211201301112-1132321311221111-2032311301202132-3102223302131211-1223222021312200-1222002021302200-3110303100121323) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3011002203200230-1030332000021331-0130331100120030-3023130030110233-0312013300320303-3200132031200331-0302313103210003-1013321010302201) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2102103131101303-1000113010022312-3322201001111113-0222311002131002-3300030231311122-2310002000022203-1033111002023003-2201200111323101) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0313200331222213-3011003113232123-3201220020232311-3323121002122212-2303021320013111-1010130302332133-3012311020312120-3131033030003130) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3010102002000003-0021023210003330-2100232200311313-1303002032222013-1101032023321313-1133222000031032-2123312112112021-1131010230311201) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2231301320323132-3321012201321121-3133002303212030-0003221300213332-2000333322330113-3330133113202323-0322031220213220-2203013332301030) |
| `instance_type` | [instance_type](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2213220100123022-3112000221011111-0031333211321330-3000132101201123-1202233300100301-0112032210112012-0020002202131322-2310120110032223) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0002122032131302-2200202021122223-3022132011312112-1122013331003113-3112220132232020-1103123130233212-2120200021130012-1122211210330310) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3301112121031322-3013000213113120-0001112023333211-2000132121010303-2133032232110133-3010221021111213-1023303111002332-2301103313232223) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0310023332203232-1000010221212310-1101312113030310-2033201223202301-3230211322231010-2020030031113333-1202310300020002-2123031112222011) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1131102200031101-1301220300220130-1131333003311120-3131133230113111-3012122230130330-3123123332202133-1111323133133301-3000330023322023) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2122211012012301-2203332313101330-2032302110220113-2132122202302131-0321112032313321-3103212321310211-1312132331023101-0112203312222001) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2112321323121332-2013222311322221-0302120122030030-2202310222102210-3233332131201310-1021302132111232-0022110323333023-3312103222231203) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1000312003003221-0132033211203122-3310230323121231-2100031203233033-2223331023303102-2010100133020311-2301000113300221-2101113112322312) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3232210220233102-1233233211213033-0203000032323333-0321111003322010-1221211121131302-2220301122202021-3101213230133002-3310100212022003) |
| `labels` | [labels](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0030113023031300-0113220000110121-2302103003120033-3213123300300223-3032222330003332-3323021022002113-1110322300122211-1302232003332132) |
| `log_receiver` | [log_receiver](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3331323212131010-3111112031103213-0133320332031013-0102000120300233-1330031131303001-3321111310320230-0200133201133031-2132331321100323) |
| `log_receiver.name` | [log_receiver.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3230013010132121-1100231303213101-0313333213201222-0133220332123312-2221030312321202-1031021302031113-1032113113323322-3101031221301300) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1221120031132332-2300320211211222-1101301213011303-0012111033203113-2021332222300103-2031213100233021-0103333200123213-2130020312020332) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3101011223301001-0313320010001012-0322320101332303-1131030211013000-0123132201033332-2123003022220132-0111311333312200-1210202022021322) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0111102031022101-0210132112322321-1232313313322030-0010301103033013-0113301133211123-1030233303202120-3100312000310323-2200003000303001) |
| `name` | [name](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2211221221131011-0303102211331030-1202221302320132-0230110211333131-2322030221022333-1331331302133102-1020203112323131-0002230213322132) |
| `namespace` | [namespace](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1230130000200211-1232231202020113-2302130211313322-1330330210030201-1211110320133122-0000333302221330-1233303223033311-3033211201223222) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1231320222331220-0033021033321301-2121030121313330-0322130303303010-0201031312232220-0222102023133322-0010130330111311-2023103222001031) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3233231322211120-0010020113211022-2221012133321223-2303211311221212-0230122023120212-1322020202222330-3230122113120213-0310011120231300) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3221132001320310-1200300220233120-3112123322301030-2022330301302020-1231322020021131-1223302332023321-1223210113032311-0011301221102131) |
| `os` | [os](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0002023020202321-2233101100131012-1111332130233122-0223120022103220-0010001220332111-2310202232133313-1121312202220123-2321203033310132) |
| `os.default_os_version` | [os.default_os_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1010102031122310-0302313230101210-2030223222333113-3323031023022302-1211222322112112-1120233301222010-0021201022310221-3233321032120001) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1303313020223220-3032230220201113-3200221103113020-0203311212230220-2232133323221122-1012103213203022-2103023003133000-3203220011230122) |
| `private_connect_disabled` | [private_connect_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3231202332022223-1300332201011312-3201201033313130-0122032123131012-2101122222321311-3212201111202202-3230001011202122-3103103232130333) |
| `private_connectivity` | [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3101102303013132-2123031132212110-3112212212100023-0232100020232331-0010303022333033-0330322101100330-3310201120031111-3021223203332133) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1223012231302030-1003320211330111-3102031323011222-3222333230331012-3210210100102003-2131330031213313-1200323201030010-2030111223321002) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0021323012112130-1331222203103113-1302023223201012-1330020321122301-0321131101101112-3211131110232032-3001203300122123-1131212321201101) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3220332331101121-3323222231111023-1030020130113132-2002012300221211-0211200301301322-2323120030032322-0220320231230220-2133300011321201) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0131200002333212-3020111330030133-2303320331103312-3302331202313003-2013322202120032-0312022233000212-0201303222120321-2123133312332201) |
| `private_connectivity.inside` | [private_connectivity.inside](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3001210001211311-2201220011232011-2121120310211102-2312000311030023-2213020300023123-1333322033133003-0300123310132312-1212121032132030) |
| `private_connectivity.outside` | [private_connectivity.outside](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3200032013200101-1013231120020332-0321321000230031-2313313112200113-0302202301323132-0321310202102102-0323002220330323-0203321021232012) |
| `ssh_key` | [ssh_key](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1131200321210212-0100100122002120-2311021313321002-1333101002100203-0130032133330011-1213231202000101-3222013031313032-0011020301132002) |
| `sw` | [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0231010130132301-3230223211201230-1131212221130330-0011311122031022-0120120303021111-2220120011022310-1011103233320012-0330322321331211) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0132233033323202-1310331310021111-2121232112202121-1202121333223203-2001033123100213-0200202030133023-3101210230330033-0333212311310302) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2203231232232023-1203321030310030-3103210020103020-3322022212330113-2223202122100212-2231102012032012-0203011100131001-1022200312222002) |
| `voltstack_cluster` | [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0203133203303222-3213130300030211-2023033333200210-0202320122220013-3100113312003032-1000031013121311-1032330000300302-1022103102203101) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1301310132132100-3120112033223003-2021010133032212-1232003133112021-2302100121300103-2120320013223220-2111310301313110-0130212003130212) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2333332221130203-2210321221022230-1312133312201032-0012201031123002-3223313300320233-0122120033331320-1122332113323020-1321222131101011) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1131133112333322-2331302323202033-3222322111312222-1330222120300023-1310223132103031-0302302330111113-3202210321103123-2330012120113322) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3023202303130112-2000031300213010-0120330321323132-1100020223032312-3312310022332223-2210202302322323-2100101101001133-1022331320013012) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2331021311101122-0132111120113112-0132223231100123-3230131020130221-2111211113111111-2220120333332120-1230230323022231-3311203112313120) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1000212330333001-2333210132332120-0000031021301222-0233333230010220-1100111323210213-2332322013331013-0131223201010213-3303023123220032) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2222200203033223-0121330323113201-2110031330120021-3122313312030303-0301231033011331-1021003010322200-1102013300312210-1121331301210012) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3121333122202223-1001220202013023-2313103301311020-2313323002200120-3321301001100103-3312032331102130-1132130113202133-0130100011033031) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3213021113231212-1211031013031311-3022230123003330-2221311323111130-0123213110000322-2121200102202000-1323113331230321-3311320222232130) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0012310101033132-0021120311023111-2203311310322012-0332221113132301-0003323221101011-1302122123300103-1223322133033302-3212320331102313) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1002021111311231-3032230301121310-2010032131113300-3223132213020212-2222331102232232-3331122210212332-1120223203022223-3132312031232300) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1002322110130202-3001303120223211-3112321020333103-0100031311010320-0210322031303011-3131133321202022-0210033320202230-0232112031021103) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1213130331010300-0231312320100112-0031110202212322-2121110120031012-3313101221203310-3332310131313322-3012331221322203-3132233212021011) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1102011223122230-1222011103132012-3022123033220222-1230310230210231-0222202200202311-3201001013132200-3223322113230213-2302203022332310) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2320112310120030-0201033222023213-0011221230120110-0130310030121212-2132312013112022-3212301032213000-2222233001331123-1221003011310233) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2010012001130330-3232102120221230-0030211202311023-0233202030030312-3023210133022123-0030133300012301-3130331233020110-2211132210031022) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1103322322121321-2321013010223003-2013310203213021-3123110210032121-1200130031020212-0300331302300323-0233231033011022-3230323132132033) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2300201000012210-2031001303011202-2312020110301023-2221102223301202-2321003310301133-3002030331011033-2030100312121122-1200110032221230) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0033133132321230-2203121222212002-2232203223220312-3102022331110303-1113211221003121-0211030102222123-1333310202121233-3101230330301111) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0302310021333323-1321002032210330-1200130211001000-3331230013223013-2310332323102303-0312222201112220-0121100232032101-2313331023033000) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3300230133233231-0012011302210331-1112131321010002-3030012112122320-1312121023023311-2221132212111132-1102312312321023-1101302311302100) |
| `voltstack_cluster.gcp_certified_hw` | [voltstack_cluster.gcp_certified_hw](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3330231013222232-3321313230011130-0032321332331231-1213203102220100-0101111323111112-3022123301211323-3002312003220231-0001111121302013) |
| `voltstack_cluster.gcp_zone_names` | [voltstack_cluster.gcp_zone_names](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0102121311130021-1031012132233013-1231301111110112-3312130101231322-1312320311223301-1222032223203313-0202323112103230-3020203312301021) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2323211020333101-2331101232201103-3031231202233221-0101202333110323-3311333230212303-3230032001131321-1011021300201231-1220001323300121) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2323021200122121-1112123121330131-2013232220021033-2313002100221012-1102002000122131-0231303203012030-1021012302212321-1122222002102013) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3011212102300232-3332222111311222-1331033131013103-3200000321310322-3013111221102012-2321011121210220-1020331211311320-3311201230102100) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2123123122031302-0210231213100132-3122300312232303-1111111132220001-2032021112111333-0101000111103233-3331322013001211-0210320312001232) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0113322312101302-2121111033321030-2300201101220322-2000010310303000-2033132223210322-3133030232012332-1021332110131331-2311030200113002) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2331320022022111-3231200120210022-1200313121021303-3032113013303313-2011311221003120-0203030002011121-1201311131201031-3300310223033033) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1200033210122333-1312220100201031-0132132321201131-1200100300231200-0200123110301311-0110323002331322-0231202023123202-2011210230320003) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2333012221331233-1320331201222321-2022130033110123-0213202102333102-2103001033213133-0210222023133031-0132100233300233-0322121003021103) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2102210221020223-1230200023033011-2322212130021020-0021311002323302-2111203300012233-2302333210123330-0100201201001132-2302233013230313) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2220221210120101-3321321212333001-1330232222031231-1213323023130131-3223110032202132-2002001321201122-2010311000330201-3100303012132130) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3101122013110203-2220002102000203-2031331123003300-0220310120003003-1311013303213101-1310013032323232-3110003020223010-1021110203111111) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0012123011330122-2103020030100120-3211202012123123-2121313023032121-2130022312120033-3303112110311212-2322230222010311-1330131210011312) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3012333312002112-3332331132132332-1200012230321313-0300211130331212-2232312223111312-0320101302210220-2113233222300301-0222233123133312) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1231001122333220-2130000001033320-2202100011303030-2032333323130121-0010232322200203-3013233020320013-3303013202222200-3203111210110123) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2322001012002123-0210221031113133-3002222113223311-3323323012100221-2000010113300131-2200210321121122-3330301013000021-0023021210120133) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0110323020212110-0011103230330200-0021331211131232-3131120300232313-1301012000122300-1023303111122321-0310123021031230-2332211203001030) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2331111121102012-1200111113312120-0133223021033211-2121231011121022-3123313001033000-2000300322303101-3312131230231200-0201031103102210) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0201303332132022-0202022011323033-0022321111110133-1120332310233103-2210130113113131-1213121212011000-0300130133233101-3220120211132303) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1013301121033122-0333032031223020-1003021032201301-3101203323033111-2113333301312312-1323330210323203-3031202002100222-3302031122103333) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3210331011103011-0110012321222232-1130112221121310-3011130231300032-2001010211212300-1320130031211030-3121130212120213-2102021230321321) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2111033332332101-2013233102021121-2213103220013123-0111120113331213-0000210123032012-2302221032112012-3212000303301111-3210230321212113) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2210201330222020-2301110012323202-0230230230200301-2113331313132023-0322023021321120-3031320113032011-1131323011230233-1111212313030111) |
| `voltstack_cluster.node_number` | [voltstack_cluster.node_number](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3113000302101220-1122020203210231-3103013101232202-2231131200322321-0132322031220132-1012113231130323-0332233120031001-1220311031122013) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3021201222101301-3232131222313211-1022333212333130-2313002013201203-3200211102122100-3311232311232212-1333220230021212-0000300321103203) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0132203322010303-3020013120102102-1121323033220132-2332001110202220-1213323102303231-1321131320032130-1023321031001313-0010223033230211) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2322023120212230-3131030211111322-1131300130311332-0031023233220130-3223113022023103-2022203032123023-3123002101200330-0301132330222130) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1000102113000112-3113022311332201-3000213120031101-1020010130133322-3200333220032022-0322311212121131-3001111010010031-1222121310232212) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3011231031233121-0202322231020100-1132022210300020-1031220212213102-2313120031133131-2010300333022113-1331032310132122-3201302303111032) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1202210222132031-3202202020200022-1131010102230021-0320133101030022-1312202120211120-0230000203331022-3011002032221030-1330020023300003) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3012302233112010-2133011201302032-2330202002002303-3312112320103203-1100103320231031-2212010312132331-2332232231120133-3133122031031301) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0213312003021032-1300013011312031-2231032032031003-3313013313010023-1223101101320330-1321223003331300-3103010231300111-0002131112111323) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1120332121022223-2301232110331001-0330120021202333-3332212112001010-2020013130101221-2101122231321110-1013212321020032-2100000031322131) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2112112231330103-2010002200020310-0312121323210300-2003110013200002-2301212210113211-2320002031130031-0120322101132233-0022131201013101) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1030311320300122-2330321212331002-1120332220132230-2311022002011012-3022310030030321-1130230122000221-1131013011223010-2032231333213333) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3020021023110311-3202300103003013-1022131333022230-3223122030033331-1202013202130011-1232130013010313-0002320222331011-0312122322032231) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1311213101231300-3033023211310010-0033132321301022-0322200313010330-3110100010101331-2313122123113011-3312331320132112-3212300320030301) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0031320202030113-1113121012311220-1210202222302313-1033203110200110-2111303000101133-2300313122200212-2221230303132223-2012111232220021) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1001200133321233-1100311001112331-2012000233203320-2231110220031210-1132201013003233-1300322333201301-2031211110213103-2303332131230313) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0130032113112200-2012231230221331-2013203213320130-1223303022021320-0133221200222003-2313203202220210-2211232201302033-1032201203031032) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3332220102010130-0311010310033300-0031210321200220-0123120201213012-3131210231210203-1103322020122323-2003010303130213-2031110110012220) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3232001301020001-0132133332321021-3213111122031303-0020022010321001-2033320030332001-3200000213301320-3031303210220102-2023303202200000) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0030220132002131-1331002300320020-2112131013013131-2003323331212001-1130233333010003-3032220003223212-1302222120332102-1103031003311221) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0121223312202203-1220232333321110-3331302100230300-2223101301010103-1212012312230012-3321322020010011-2032200112121333-2023233100330013) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0012321212000301-0110110202033002-3100011202320102-0030103031120210-2030210011213331-2232303310211110-0002310231002223-1213012300202111) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2133003110022331-2330202301333301-3111120203012013-1203120010232231-3020112102123233-0331122111101123-1213101130120010-1112113102221111) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3212200001332303-1130123202020010-0021212100212030-0010113123031222-0212013332103011-3221223133300331-2220022101330223-0100131103010022) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3301031311030203-0301302210223303-0123031010113012-2322211301211202-0322200132202301-1323103311230312-0120223020330331-2230201222120333) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3003203323233113-1222113223233000-0211220202012330-1232332212011102-0310321001131231-1202020001102311-3323330321311232-3103021111200311) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3311211220132211-1202221310321011-1313022332002211-0332021202132311-1311210303231113-2311120000211220-3111213133131001-0133331213233323) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1323130310030021-3323001230321031-2013232123211330-2001031211322233-3033121223133003-0112130131012233-1323222021302213-1030302222233320) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2321031022113230-0031201301200122-2203003130222101-1232323130323233-0203133000020100-3230333122223211-2001131201313211-2112231103111300) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3303232121301001-3122021130113331-2333101001220313-1111212023001312-1123103013313102-0312100212030012-1000030113112322-3232133212023003) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3210103210212333-2130122210220201-2120121132100231-3312313211213212-0230102031321332-0010023113212202-0330120212001032-2131002320131111) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2200213103200322-1020332111133313-2032222032021121-0313303300123012-3322122230330112-3230310323020112-1120133113320032-0011132132120203) |
| `voltstack_cluster.site_local_network` | [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0031131122131032-0222213322201333-1303321220010012-0003220011033221-2301010101203102-3100003023130132-2103302312333210-0303023012023020) |
| `voltstack_cluster.site_local_network.existing_network` | [voltstack_cluster.site_local_network.existing_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230233221332300-1212003311332221-0320102302000333-2022223103003020-1301101330332332-2202101121020022-2203210100312103-1003210122332020) |
| `voltstack_cluster.site_local_network.existing_network.name` | [voltstack_cluster.site_local_network.existing_network.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2320321311110020-0231112123331333-1312231033300111-1010331023122022-2013201332000023-1311312133013133-2130022333101113-0002322212212012) |
| `voltstack_cluster.site_local_network.new_network` | [voltstack_cluster.site_local_network.new_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2201203010022200-2311122333203001-3301213233232200-2202023322122112-3333220302023022-0213230330023220-2000132301313003-0131100301100321) |
| `voltstack_cluster.site_local_network.new_network.name` | [voltstack_cluster.site_local_network.new_network.name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0303330030320103-3021002020312130-0210203031022212-1300012231312233-1111322000221112-1311022203020311-1131201103033302-2221213103013221) |
| `voltstack_cluster.site_local_network.new_network_autogenerate` | [voltstack_cluster.site_local_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2022331232201000-3020001030220202-2303033011221130-3001103130222330-0322021012200201-0223101311310302-0120200332103301-1122120231211002) |
| `voltstack_cluster.site_local_subnet` | [voltstack_cluster.site_local_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2111002221121010-1221120331301001-3312232203003331-0110113130200222-2212001320120031-3231320000000000-3311223333033313-0232232132003223) |
| `voltstack_cluster.site_local_subnet.existing_subnet` | [voltstack_cluster.site_local_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1000332030302133-0320033302130330-0323032113230211-0202332321003031-0113323103031301-1230331131330331-0203133301030001-1002303222122211) |
| `voltstack_cluster.site_local_subnet.existing_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.existing_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1302022331123000-3103223333202121-0133223021121010-1303131222201303-3101332002320213-2100302201132031-1213031323030203-2222130023230111) |
| `voltstack_cluster.site_local_subnet.new_subnet` | [voltstack_cluster.site_local_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2113130123002210-0322320302231021-1333220320313321-3101332311201222-0003103031013331-2201002332121201-3022123130300000-2220131301231333) |
| `voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4` | [voltstack_cluster.site_local_subnet.new_subnet.primary_ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3220301011121031-0312113333312201-1122202100320002-3331202320033303-1112102231101030-1012213223331300-2130210021202333-3110031313101001) |
| `voltstack_cluster.site_local_subnet.new_subnet.subnet_name` | [voltstack_cluster.site_local_subnet.new_subnet.subnet_name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1002021223200332-3011123032023223-0020031111012320-1013212313201132-0033220132003102-0222032322012002-2232112313111221-0123323331223311) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0133330010111302-3010032013233030-2000121201122022-3113021012113111-1212232311002200-2231323212021002-2120213221311131-2101331320300210) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2011231121031122-2312100022223201-3012211312232211-0031320001201210-3013010102221203-2220132210011011-1030010322132032-0223120313102333) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2101000023011303-1112110022001333-2222022122330310-3130002222022303-1333300120100112-3310021222132202-2031122113112202-0330200101301201) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0311303011230133-3312312201011010-2223020012222300-0313221020120323-1311312330312113-3230100312232030-3313031020312123-2313032132202133) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0222101230201113-0121011300203232-1330310301112230-2122001123011111-0323231212020321-1311203132330003-3321310320232331-1012220222312330) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0302111001000032-1022313333131312-3321201121133102-2223321021202222-1202211110032020-0002330101323232-3212103031111000-2030030113230133) |
| `waf_signatures` | [waf_signatures](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1332200232203221-0021133013311332-0302302031310200-2231200203011332-2120020021102201-1320223013002030-1302022313101313-3023333000322303) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--gcp_vpc_site--reference--group-005.md#canonical-1133312313033201-1200003012001311-3020002011010231-1101311112301233-1330023213113131-2322023123200223-2112103201310212-1203210120220122) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--gcp_vpc_site--reference--group-005.md#canonical-2010300233313130-1023010211230311-3013001132203322-1011200020132302-0302200330310111-2123200323031102-3301301203211313-3201022110122212) |

<a id="canonical-0212123332031002-2113113112311211-3133032131130220-2201132133100332-0233123223222000-1033232333021333-0010011310313333-1113100221033322"></a>

## Next pages — Property reference / 202011020233 / 17

- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1122011033112013-0023030003113122-3210112130312312-1032332312323121-2102122211113223-1311220103313223-0301102033030320-0121032031032011)
- [block_all_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0113100323303113-0310110023300331-3103003023231223-3321211023203020-3010203011100033-3323230220130232-0302012300333301-2211200122013012)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0112133220222300-0000323123001311-2213320321113330-3300321323210101-2300111003213112-2130320223132032-3002112220002022-3201011103233301)
- [cloud_credentials](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2111320020011303-2323022312120210-1223222021022201-1120210103101211-2323012203320133-0332130303010223-0211232311122123-1120013213330001)
- [coordinates](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0322002133323313-1221023113310113-1232320212221333-1132003133221100-1330011002121300-0222310031323213-1210211311101320-3212100232200210)
- [custom_dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1212103203120331-0311103230311123-2000010120211201-3231323133012121-0201022203101321-0011303231201302-1302030121312313-2030312232203020)
- [default_blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0321222301321021-1113000001310201-2120223011233031-0320301313032331-1012313112133021-3030102201310311-3331001013221102-1222010313100301)
- [disable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3100312200010312-1300200010110230-0322012301333101-1312123133310211-0200301121321303-0122303210002102-2013123332313322-3122211123201302)
- [enable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2230311030212003-2200003030123112-0333131120113200-1312032120021222-3102300300031323-1102000322031212-3102223002032122-2323112010030033)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1203301201320331-1230203103032032-0033211230213112-1103110310132010-2311223132020020-1113110103030030-2212320302313110-3111021020113222)
- [kubernetes_upgrade_drain](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1022113113131303-2323310123132033-2333330201223220-0203120001033203-3132021212231333-3113000331201222-3002010313110322-3330113112120320)
- [log_receiver](data-sources--gcp_vpc_site--reference--group-003.md#canonical-3012023010120110-0132220330001122-0031313332213232-0130033203331102-0312222230232111-3100031123133102-0313211302213130-2122312222311212)
- [logs_streaming_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2202020213103112-3302231032102222-0001333100102011-2232113332301013-0123022110213010-2310010331223201-0233011232203110-2120130220331303)
- [offline_survivability_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1110001323312320-3212002110322030-1220001101201303-1333110011111112-0302100001333323-0132330102013113-0202320003302122-1131123010110312)
- [os](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2312120333122222-1322013012302121-3031012213121300-3212110230333003-1231033133000002-3321231322032311-1300102131230223-2103010231021132)
- [private_connect_disabled](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2213201100213331-0202020301100211-1311213110003020-2103223312210301-1113131132003333-3120221123300130-2032302102110033-0222302310211223)
- [private_connectivity](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1203003321300011-0223032211022011-1000102131312201-3122010333311010-1021023223222201-1131021212120001-3233201320331210-2110010200210122)
- [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0310330212021223-3303111001132333-2110220022330203-2030023332031212-0003311012011331-0210223302200113-1032100302211312-2001221222220020)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [waf_signatures](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0012012033031300-2110002220223133-3303202210233302-1331310110333111-3031100320201031-2130003013122113-0123330322110322-1011031331030122)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1122011033112013-0023030003113122-3210112130312312-1032332312323121-2102122211113223-1311220103313223-0301102033030320-0121032031032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330212032131112-1002231223321101-1333022030323301-0213321321322121-3023230030303202-3223022121221123-3021230233310023-3312000311013323"></a>

## admin_password — admin_password / 222323130022 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- admin_password

<a id="canonical-3033323011023010-0300012110310133-3322121202001221-0000212122220323-0131131032130010-3223232213310202-1213310212033301-1130221100330102"></a>

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

<a id="canonical-2233011123203303-3221130002020012-3321022122232013-1311111121302202-0022021230033000-3103200000011310-2102320131131010-2012200221110310"></a>

## Direct properties — admin_password / 222323130022 / 3

- [blindfold_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1113332303221111-1210232023033231-2321100131333132-0132211221323132-2211330331202001-3010031221002313-1211202320122033-3301111130333313): complete subsection reference.

- [clear_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3100213022320000-3113200020013000-3133023301300321-0321313333022233-0110333100320211-1323113201021212-3310202313020232-1021332012312102): complete subsection reference.

<a id="canonical-0113311022330231-2220300022210301-2011133203100220-1132203321232021-0232000010011222-2233213103200212-2201332310111121-3103232102203102"></a>

## Next pages — admin_password / 222323130022 / 4

- [admin_password.blindfold_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1113332303221111-1210232023033231-2321100131333132-0132211221323132-2211330331202001-3010031221002313-1211202320122033-3301111130333313)
- [admin_password.clear_secret_info](data-sources--gcp_vpc_site--reference--group-001.md#canonical-3100213022320000-3113200020013000-3133023301300321-0321313333022233-0110333100320211-1323113201021212-3310202313020232-1021332012312102)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1113332303221111-1210232023033231-2321100131333132-0132211221323132-2211330331202001-3010031221002313-1211202320122033-3301111130333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302300132113222-1220030102011011-3123300313002122-2020231133213033-2003013121222203-1033202120103210-0023100103122213-2202020333330321"></a>

## admin_password.blindfold_secret_info — blindfold_secret_info / 323201100332 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1122011033112013-0023030003113122-3210112130312312-1032332312323121-2102122211113223-1311220103313223-0301102033030320-0121032031032011)
- admin_password.blindfold_secret_info

<a id="canonical-0123203133030211-3003033001331113-0232300000013323-3230132120131300-3010123020120103-2121032111100103-1132132032311131-0311132012121023"></a>

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

<a id="canonical-1001303032223201-2103230130302320-0221300201121102-1102302103200322-1333133110223110-0021300222331222-0010222321323330-2123200221031031"></a>

## Direct properties — blindfold_secret_info / 323201100332 / 3

<a id="canonical-1133321132300131-2310133030330312-3121123030232113-0210330233331100-3330210211022320-3120103113300310-0202200322230223-2301311230000032"></a>

<a id="canonical-0113033100031000-3010333202100121-3130230103220202-1210120322033100-3130022112120001-2100311123033321-2213220303231122-3332003202233102"></a>

## decryption_provider property — blindfold_secret_info / 323201100332 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0122001132313003-2011233300323210-3100013003322012-2232233133323202-0333013220201013-0213331331313233-3212231131000103-2300122131013231"></a>

<a id="canonical-1130103203013012-0233221313332211-3011001203220131-1112130120230121-3011132201302030-0311013101300200-1232232013001322-0023100002010201"></a>

## location property — blindfold_secret_info / 323201100332 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3131133132102300-3122111320110012-1011103101123021-0022111330101032-1033010112320200-0211121022333033-3220203321023102-3333113130022132"></a>

<a id="canonical-2131202003303022-3112301221100211-3033103302321210-1220120110131330-1302201212030003-0213232301030220-1321312313121220-1320220003321232"></a>

## store_provider property — blindfold_secret_info / 323201100332 / 6

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1123200200220202-0131222302311221-1231120301331203-0321302113113010-1023120130031111-1222331001031103-3220021320323333-2002013113311333"></a>

## Next pages — blindfold_secret_info / 323201100332 / 7

- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1122011033112013-0023030003113122-3210112130312312-1032332312323121-2102122211113223-1311220103313223-0301102033030320-0121032031032011)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3100213022320000-3113200020013000-3133023301300321-0321313333022233-0110333100320211-1323113201021212-3310202313020232-1021332012312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331100101120301-3320022103100113-2210220022121032-1130210221322203-0220113133200313-0110131223011220-0012332132323133-0012322000200110"></a>

## admin_password.clear_secret_info — clear_secret_info / 132132000133 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1122011033112013-0023030003113122-3210112130312312-1032332312323121-2102122211113223-1311220103313223-0301102033030320-0121032031032011)
- admin_password.clear_secret_info

<a id="canonical-3230313103123211-0033130310330212-3003313222123112-3120220211110332-1331102322323310-3122100122021232-2022023023011323-2113303332123311"></a>

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

<a id="canonical-2203103020022311-2301322333321000-1013322133013013-2100200210123332-0232210113023333-3333020030010211-2101230133300221-1222200313333122"></a>

## Direct properties — clear_secret_info / 132132000133 / 3

<a id="canonical-2303033232331212-3123101230112312-1201203033033102-3013202332000310-1002301302121030-2310302030111123-1300330000202002-2012303311200211"></a>

<a id="canonical-0322331212230030-3003223201120131-0332120131230322-1213120200300020-1200303311202020-3002200113022111-0100302010221000-0320213121310320"></a>

## provider_ref property — clear_secret_info / 132132000133 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1311222210002333-1212000013222133-1323200011101122-0022320020223210-3013303233020012-3223231130012001-3100220111301000-2103022031230131"></a>

<a id="canonical-2033323320113033-2120013111222100-0032221133312121-3301123312130201-0230223013103112-0010011231102321-1121111123003201-1201212321010213"></a>

## URL property — clear_secret_info / 132132000133 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2323120311221210-2233133210030121-1331123322213002-0100303303302002-2323100223313201-2022020212313303-0233103001111323-3103023330311122"></a>

## Next pages — clear_secret_info / 132132000133 / 6

- [admin_password](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1122011033112013-0023030003113122-3210112130312312-1032332312323121-2102122211113223-1311220103313223-0301102033030320-0121032031032011)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0113100323303113-0310110023300331-3103003023231223-3321211023203020-3010203011100033-3323230220130232-0302012300333301-2211200122013012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030003102032223-3111020023200101-0301000031030333-2232130202212210-1303000210131330-3220222302033010-3202110003300210-1312310230321120"></a>

## block_all_services — block_all_services / 211212230310 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- block_all_services

<a id="canonical-2321132300131021-3120103231201003-0132232310133110-3332101201330320-1311220331223123-1323203303130321-2023300210122320-1102303210032113"></a>

Type: `["object", {}]`. Computed.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

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

OneOf alternatives in this subsection:

- [block_all_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2321132300131021-3120103231201003-0132232310133110-3332101201330320-1311220331223123-1323203303130321-2023300210122320-1102303210032113)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2300113002021223-1232203130003210-3021221232212012-0103322233302132-2212120121101103-0312330301111031-2232120212102203-0302011032303022)
- [default_blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2301211332100101-0233022321011302-3101330201013312-0302321221010310-1100312212323232-0311300230010000-2022312303321300-3012222312231322)

Select alternatives according to the provider validators above.

<a id="canonical-1321331321113320-1032000013020121-0332031023101212-3220233312221131-3312030011000302-1003133210202110-0111003102111011-0030112312021203"></a>

## Direct properties — block_all_services / 211212230310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131103300031122-1121011030133311-0211331113312022-0120212100210003-0303303332222121-3212221223211302-1002023301020120-2121201021203312"></a>

## Next pages — block_all_services / 211212230310 / 4

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0112133220222300-0000323123001311-2213320321113330-3300321323210101-2300111003213112-2130320223132032-3002112220002022-3201011103233301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003201011122131-3330113123303202-2302323210300011-2002210033203130-1201222333121213-0332022033101320-0121013211323021-2102221012333321"></a>

## blocked_services — blocked_services / 211330213123 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- blocked_services

<a id="canonical-2300113002021223-1232203130003210-3021221232212012-0103322233302132-2212120121101103-0312330301111031-2232120212102203-0302011032303022"></a>

Type: `"single"`. Computed.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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

<a id="canonical-2313321210111002-2311331202113102-0033112132313300-0113213111111102-3231110233131012-2300122313200103-2100113203203020-3320110023312121"></a>

## Direct properties — blocked_services / 211330213123 / 3

- [blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2202022012222332-0221002302320022-1203313311322132-1321330221302202-1023300210312100-0020022200001323-1330023011013000-2131131133010121): complete subsection reference.

<a id="canonical-1213011102323233-1332222232331112-1100203010002313-0103023200110103-0033021301210000-1032301233102121-3210113130031203-2000032020102313"></a>

## Next pages — blocked_services / 211330213123 / 4

- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2202022012222332-0221002302320022-1203313311322132-1321330221302202-1023300210312100-0020022200001323-1330023011013000-2131131133010121)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2202022012222332-0221002302320022-1203313311322132-1321330221302202-1023300210312100-0020022200001323-1330023011013000-2131131133010121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032131332213022-2123202221332032-0021312101332020-2320121130133012-3131102201323112-1121312030020122-2103220132013101-2220321332002330"></a>

## blocked_services.blocked_service — blocked_service / 232131231002 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0112133220222300-0000323123001311-2213320321113330-3300321323210101-2300111003213112-2130320223132032-3002112220002022-3201011103233301)
- blocked_services.blocked_service

<a id="canonical-2220331212111013-2311331210033001-1121311301203031-1200133031312022-0012202221222322-0231001110111120-3320013031111302-3003330133103112"></a>

Type: `"list"`. Computed.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333100300023310-3132001223111001-1102121002023303-2312200031121011-1221311200031310-0200220002010121-1133300022102133-2122323003321213"></a>

## Direct properties — blocked_service / 232131231002 / 3

- [DNS](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0010002033322101-3033312220112100-0222313233102331-0210030030100202-1230120033213221-2221003122233130-0320211032321232-0310312211022122): complete subsection reference.

<a id="canonical-3221132120121110-2100001221102032-2200311203033101-1103130012233100-3002233101030110-2013331023302012-3312220131320310-2110302201232122"></a>

<a id="canonical-1210200223221120-3022112113222112-1320203330222230-1001303312330331-1210300121103021-1002322110223232-1313222011030320-3330213320111113"></a>

## network_type property — blocked_service / 232131231002 / 4

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ssh](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1330121112223010-1312221320201203-2333203121320203-0001132331331021-2301012012030022-1322132203102012-2130322031330311-1030221231230101): complete subsection reference.

- [web_user_interface](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1013113000301201-3120113110201013-0013133202103320-1031102203010211-2030001032300200-0013223003322211-2100301202031212-2222203302133230): complete subsection reference.

<a id="canonical-0112232112323001-1100310003212310-3131121000213121-0111322120323232-1303321310120102-0223011000120203-0332230021332123-2012130321010222"></a>

## Next pages — blocked_service / 232131231002 / 5

- [blocked_services.blocked_service.dns](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0010002033322101-3033312220112100-0222313233102331-0210030030100202-1230120033213221-2221003122233130-0320211032321232-0310312211022122)
- [blocked_services.blocked_service.ssh](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1330121112223010-1312221320201203-2333203121320203-0001132331331021-2301012012030022-1322132203102012-2130322031330311-1030221231230101)
- [blocked_services.blocked_service.web_user_interface](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1013113000301201-3120113110201013-0013133202103320-1031102203010211-2030001032300200-0013223003322211-2100301202031212-2222203302133230)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0112133220222300-0000323123001311-2213320321113330-3300321323210101-2300111003213112-2130320223132032-3002112220002022-3201011103233301)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0010002033322101-3033312220112100-0222313233102331-0210030030100202-1230120033213221-2221003122233130-0320211032321232-0310312211022122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311211131203131-0033122201133200-0222321021232111-3130132202330333-3221223110310323-1131312300212201-2103211100020033-1023303303322212"></a>

## blocked_services.blocked_service.DNS — DNS / 031013200220 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0112133220222300-0000323123001311-2213320321113330-3300321323210101-2300111003213112-2130320223132032-3002112220002022-3201011103233301)
- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2202022012222332-0221002302320022-1203313311322132-1321330221302202-1023300210312100-0020022200001323-1330023011013000-2131131133010121)
- blocked_services.blocked_service.DNS

<a id="canonical-2102333102031100-3233222311023001-0023121210011110-3211231233123223-1131300332202001-0031220323111101-2023032012321322-0021301201030320"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0221123022103132-1013123020210211-1310232233101032-2211220232312021-2313222202003213-3121221330223012-3212330201310302-2020020313102130"></a>

## Direct properties — DNS / 031013200220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331101002222310-2101002120022331-1331323111102302-0133031320101301-3212013310003011-2203120313033223-2302332103012321-2300102020203031"></a>

## Next pages — DNS / 031013200220 / 4

- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2202022012222332-0221002302320022-1203313311322132-1321330221302202-1023300210312100-0020022200001323-1330023011013000-2131131133010121)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1330121112223010-1312221320201203-2333203121320203-0001132331331021-2301012012030022-1322132203102012-2130322031330311-1030221231230101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102210120002130-2233323301133230-1202002101202312-2003023333232332-2132202213212221-1120311111131230-0003202132303210-2312100322230021"></a>

## blocked_services.blocked_service.ssh — ssh / 333333212130 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0112133220222300-0000323123001311-2213320321113330-3300321323210101-2300111003213112-2130320223132032-3002112220002022-3201011103233301)
- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2202022012222332-0221002302320022-1203313311322132-1321330221302202-1023300210312100-0020022200001323-1330023011013000-2131131133010121)
- blocked_services.blocked_service.ssh

<a id="canonical-3011223012123301-1122110132031111-1212223302301133-2213323122330210-1020000332300020-3320222000311230-0221200330320230-1321103213002333"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1130300011233321-2130303331012221-2022000230221001-2011100311110221-2302331301220212-0033102113233003-2130300133000132-3223211233323330"></a>

## Direct properties — ssh / 333333212130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213301232123222-2330320310230321-1320103122333123-0103122122220313-2323223010301002-2001211222213201-1321031311331132-2201002023210011"></a>

## Next pages — ssh / 333333212130 / 4

- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2202022012222332-0221002302320022-1203313311322132-1321330221302202-1023300210312100-0020022200001323-1330023011013000-2131131133010121)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1013113000301201-3120113110201013-0013133202103320-1031102203010211-2030001032300200-0013223003322211-2100301202031212-2222203302133230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020201221000003-3232033300001030-1123321212212323-0101002033211123-3332332131200200-1313030102220012-1201231212132003-1000203322303332"></a>

## blocked_services.blocked_service.web_user_interface — web_user_interface / 021013030203 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [blocked_services](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0112133220222300-0000323123001311-2213320321113330-3300321323210101-2300111003213112-2130320223132032-3002112220002022-3201011103233301)
- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2202022012222332-0221002302320022-1203313311322132-1321330221302202-1023300210312100-0020022200001323-1330023011013000-2131131133010121)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-0201033311102320-0301310201030231-0022020332122000-3103130222010012-0203323332023310-2122300220211333-0031123122333022-3200320010013032"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3133122033112000-1333222013331223-0312003202312332-2222033303231122-0013322012300120-3000330001032233-1130121110023102-3012313111021303"></a>

## Direct properties — web_user_interface / 021013030203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011111021200312-2300203002330122-3010031203001200-1331030030222101-3331002012113330-0130030223030123-1022121021130213-0323331211112032"></a>

## Next pages — web_user_interface / 021013030203 / 4

- [blocked_services.blocked_service](data-sources--gcp_vpc_site--reference--group-001.md#canonical-2202022012222332-0221002302320022-1203313311322132-1321330221302202-1023300210312100-0020022200001323-1330023011013000-2131131133010121)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2111320020011303-2323022312120210-1223222021022201-1120210103101211-2323012203320133-0332130303010223-0211232311122123-1120013213330001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332302231320231-1010113013032223-3212010111223021-3101020012322010-1210332303120233-0030021201013121-3033011321003132-2232220320122230"></a>

## cloud_credentials — cloud_credentials / 023202322330 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- cloud_credentials

<a id="canonical-1031211303000002-2120012302332133-3011302012221322-2020003301313100-0222102033023012-0302133032323321-1200033100022333-0220211331120320"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1223020013103230-2101233033000010-3000331022022313-3300311323032013-2302130322123002-2332030323223313-3213001233000122-0212103233222303"></a>

## Direct properties — cloud_credentials / 023202322330 / 3

<a id="canonical-1010132301132200-2001010021333231-1203211322323012-2113210213002322-3201101321022322-1201221122333022-1312132212023032-2333300230111321"></a>

<a id="canonical-0332030001103300-0103022123023332-0101311220220031-0302121102121332-3123222212332230-2003211113302232-2130031222313210-1313232303211111"></a>

## name property — cloud_credentials / 023202322330 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1013313100030010-1110133200232320-0211231221212203-3212110202002303-2010031022321100-2003312310200103-2333133103202300-3301331212212022"></a>

<a id="canonical-0223113333130113-1332001021212320-1333031210011200-3023033021220101-2102011320233331-1130323002323303-1313123103130110-2221332211202130"></a>

## namespace property — cloud_credentials / 023202322330 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3210302021110102-0003023203013331-3130221113012003-3101132213120123-3133310030030213-1201132003111133-3330300130111123-2231210231213322"></a>

<a id="canonical-3222130123130033-1300121030032130-0213022020133033-2300110303320312-2113020132233010-0031302131303222-1011221300232011-3031022211312003"></a>

## tenant property — cloud_credentials / 023202322330 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2303211203321101-1200331303213013-3000111003222310-3310031301132200-2032122011112312-1333113122310300-3123031211120231-1021320120330330"></a>

## Next pages — cloud_credentials / 023202322330 / 7

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0322002133323313-1221023113310113-1232320212221333-1132003133221100-1330011002121300-0222310031323213-1210211311101320-3212100232200210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220023121123201-1212113001310102-2310012003131111-2330203302132011-0120023121001101-0211101222211322-0121122110002203-1232222012101102"></a>

## coordinates — coordinates / 213123001233 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- coordinates

<a id="canonical-0220231032011300-2312010103121300-2233002220222030-0312322002111203-1310332131102213-2232112010103130-3212030211101111-1110310002131303"></a>

Type: `"single"`. Computed.

Coordinates of the site which provides the site physical location.

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

<a id="canonical-3020203033201203-3003130023120301-1002300133303112-1331212200200300-3111021333022231-2211103001102313-1210103033302021-0200122301323103"></a>

## Direct properties — coordinates / 213123001233 / 3

<a id="canonical-0320230333031323-0133123000233320-1210230001330332-0220133011200133-2310121231103012-3202011122310000-3112101200333133-0223210332203013"></a>

<a id="canonical-2320220312012222-3023332311202010-1110313103210000-0022212102020112-2223003331121311-1300131203132322-3113211231322002-3033021301031312"></a>

## latitude property — coordinates / 213123001233 / 4

Type: `"number"`. Computed.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-1110300022130031-0000111312223101-3012032023121030-0333030200210313-2331012310100000-2011330010112012-0201210123310021-3333223201320233"></a>

<a id="canonical-1201331011122300-0200322123310230-3213113222302003-1303320133130112-2323311003032131-1300013212320203-0102230330030020-2121220132031200"></a>

## longitude property — coordinates / 213123001233 / 5

Type: `"number"`. Computed.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-3313012311113313-0021132321310302-0302330031111123-2033032110322312-1300322303000232-3312331232031230-2233211110201101-1123103130110132"></a>

## Next pages — coordinates / 213123001233 / 6

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1212103203120331-0311103230311123-2000010120211201-3231323133012121-0201022203101321-0011303231201302-1302030121312313-2030312232203020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020101123033013-3103203301231203-0023021233202033-3113301320300110-2131311010003003-2131033332121011-2003301132303212-1333301003221111"></a>

## custom_dns — custom_dns / 202123000030 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- custom_dns

<a id="canonical-0033221203121322-0233123323010122-1003313302220131-3221003323102322-0113023302112330-1123313112033112-3111122021303230-1011332000232230"></a>

Type: `"single"`. Computed.

Custom DNS is the configured for specify CE site.

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

<a id="canonical-1310300132003101-1230223312300322-1233322203322020-3333123221033332-0312031233311300-0231233010311110-1233330303220210-0300033120022111"></a>

## Direct properties — custom_dns / 202123000030 / 3

<a id="canonical-2221321132232331-0132120011302103-0221320230031003-2112320000311031-0021001021123332-1201301321303100-1230021311120033-1303020000223201"></a>

<a id="canonical-1323313203131222-1331013120112210-1233231302100302-2033020321133203-1220312301302313-2333121110011130-3103031313202100-2021300032333012"></a>

## inside_nameserver property — custom_dns / 202123000030 / 4

Type: `"string"`. Computed.

Optional DNS server IP to be used for name resolution in inside network.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2012222130203300-2202332310113033-0111201002333032-3331301331002200-3212023030321200-1213211103001032-0213122323030023-2102330330123203"></a>

<a id="canonical-1231123102321232-2123022310223000-3013223303330132-2010112323330123-3023201100012333-2101110021020301-3313323123133222-2121213132000030"></a>

## outside_nameserver property — custom_dns / 202123000030 / 5

Type: `"string"`. Computed.

Optional DNS server IP to be used for name resolution in outside network.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3132011223323330-1210103203130102-0301102213113300-2213231122100021-1223300120132303-3201010133210130-3133203222123303-1112033230332323"></a>

## Next pages — custom_dns / 202123000030 / 6

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0321222301321021-1113000001310201-2120223011233031-0320301313032331-1012313112133021-3030102201310311-3331001013221102-1222010313100301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122121030010331-3223123131030300-1302233200331312-1011203303021203-2023310202102031-0201322010303022-1333311031011010-1331221120002231"></a>

## default_blocked_services — default_blocked_services / 000201102330 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- default_blocked_services

<a id="canonical-2301211332100101-0233022321011302-3101330201013312-0302321221010310-1100312212323232-0311300230010000-2022312303321300-3012222312231322"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2213310232012123-3223021012011210-3022331201320200-1113332010322030-2310313333200031-2130330122121002-1330003220331030-1111303113012033"></a>

## Direct properties — default_blocked_services / 000201102330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111222213300002-1001322201113003-2223323023022133-0121101101113333-3032130230301320-3311332100323210-3033021213203113-0222000311222121"></a>

## Next pages — default_blocked_services / 000201102330 / 4

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3100312200010312-1300200010110230-0322012301333101-1312123133310211-0200301121321303-0122303210002102-2013123332313322-3122211123201302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012333120112303-3112103013022101-2112210030223213-2231133212113300-2333030230130031-0110210033030303-0121231311131231-0001230322130321"></a>

## disable_encryption — disable_encryption / 123120101021 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- disable_encryption

<a id="canonical-1033112032333201-1033213312021303-1221112331120233-2330013211032333-3300000311123110-1331213100130132-3112011212321322-1212110220200100"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_encryption, enable\_encryption; Default: disable\_encryption\] Configuration
parameter for disable encryption.

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

OneOf alternatives in this subsection:

- [disable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1033112032333201-1033213312021303-1221112331120233-2330013211032333-3300000311123110-1331213100130132-3112011212321322-1212110220200100)
- [enable_encryption](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1001213133130022-3301300111031002-0300311200120333-3001120320012320-3013211321222111-0132133302231310-1201000121233101-1202131020033233)

Select alternatives according to the provider validators above.

<a id="canonical-1002300032103010-0212230103221111-0000132332010312-3211201102020321-2300010002101333-2323030011301032-1121020220010023-0023122230301020"></a>

## Direct properties — disable_encryption / 123120101021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110123200111031-1121120301210120-1202113131210332-1032331110313211-2101132022200212-0133332031311311-2220000223223123-0130210000202030"></a>

## Next pages — disable_encryption / 123120101021 / 4

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2230311030212003-2200003030123112-0333131120113200-1312032120021222-3102300300031323-1102000322031212-3102223002032122-2323112010030033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020321132231232-3211021223222030-0210113220121101-0001320133211210-2130231201221233-1032110032102113-0121031230112130-2001310302303311"></a>

## enable_encryption — enable_encryption / 230020223311 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- enable_encryption

<a id="canonical-1001213133130022-3301300111031002-0300311200120333-3001120320012320-3013211321222111-0132133302231310-1201000121233101-1202131020033233"></a>

Type: `"single"`. Computed.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

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

<a id="canonical-2311312030111001-3100320101302233-0213233011131131-0010320212212230-1220300201101031-3103323131110331-3333011333333103-1212003012221121"></a>

## Direct properties — enable_encryption / 230020223311 / 3

<a id="canonical-2121001003110200-3321023112023323-2322333203122130-1122000223001333-2231313331101121-2311201230331301-0101313131022103-0322021013203313"></a>

<a id="canonical-2322210301123300-1131212300002020-3202023202203102-2210332222232331-0213123211333000-1232230331311333-0120201203311312-3101122220302322"></a>

## kms_key_resource_id property — enable_encryption / 230020223311 / 4

Type: `"string"`. Computed.

GCP KMS Key to be used to encrypt the disk attached to the VM.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1010331221321311-3033332202022132-1221020001232130-2001223221322110-1311303300000332-0031302332210123-0131112233112110-1123202013103333"></a>

<a id="canonical-2030131203020323-2123120012132032-1012301113322012-2230133013230320-2120211120010221-1131300110000000-2212113331311101-2011332220220301"></a>

## kms_key_ring_id property — enable_encryption / 230020223311 / 5

Type: `"string"`. Computed.

Key ring in which the CMK to be used to encrypt is present.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0001111103131131-3032300120321022-2233212210120030-3202133030120220-0022001303130002-0021313110112302-2123103031323300-3230222131210231"></a>

## Next pages — enable_encryption / 230020223311 / 6

- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202300130031231-1303011332231220-2011000312000322-2013121022031210-3102233010322031-3330321202102102-1203211022003213-3122322121012201"></a>

## ingress_egress_gw — ingress_egress_gw / 211132000220 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- ingress_egress_gw

<a id="canonical-0313012223320302-0022031003132120-2102001313322132-3131121230300302-2221202331000301-0133330100301331-2203003130202202-3103202120020301"></a>

Type: `"single"`. Computed.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface GCP ingress/egress site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0313012223320302-0022031003132120-2102001313322132-3131121230300302-2221202331000301-0133330100301331-2203003130202202-3103202120020301)
- [ingress_gw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2222232300212112-3211332221211323-3302333303221100-3122023230323133-1112203332332221-1323033133300103-0220100330013232-3332102330003130)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0203133203303222-3213130300030211-2023033333200210-0202320122220013-3100113312003032-1000031013121311-1032330000300302-1022103102203101)

Select alternatives according to the provider validators above.

<a id="canonical-2203010201203230-1112131313000021-1232231230200023-1020031231012001-2022032013121231-2321230031030001-2003203121333203-3300000302031101"></a>

## Direct properties — ingress_egress_gw / 211132000220 / 3

- [active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3233300222201200-1110331210220311-1213131211123213-3330101101013122-0203013123032112-3220323012000312-3023030231002231-3030233101230200): complete subsection reference.

- [active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3212012103120301-3323313231002021-0233131311130103-1120023320133210-3310332102220022-3330211100123220-3312233112213331-3030020303332303): complete subsection reference.

- [active_network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3222311311231013-0311301030103101-3030230333023110-0213023110303123-1010031333220100-3130011010011023-1103312132000003-0220313311222311): complete subsection reference.

- [dc_cluster_group_inside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2031022112303103-2210230022033131-0001223112320121-0330132121020222-2211002033113200-0020112222110023-2032232310213122-1123323301301001): complete subsection reference.

- [dc_cluster_group_outside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2303202101201031-1220332112033302-2023202030113102-0030131100101331-2311223011220100-3120013003100133-2300110232001312-1031222330302022): complete subsection reference.

- [forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0032132213332210-2032133310221201-3121310321330300-0322330120331213-3122212000320102-0310002110102331-0211331112222332-0003201333311232): complete subsection reference.

<a id="canonical-0213113302333011-0220223322322001-3013303010131032-3002013312201012-3310100303221303-0333113330003331-1020010103213120-0031303001313023"></a>

<a id="canonical-3323301310011100-2101133303033302-0211231030300111-2112012303232322-2110321222003020-1310003332130221-2221123023321101-2220220311201320"></a>

## gcp_certified_hw property — ingress_egress_gw / 211132000220 / 4

Type: `"string"`. Computed.

\[Enum: gcp-byol-multi-nic-voltmesh\] GCP Certified Hardware. Name for GCP certified hardware. The
only possible value is \`gcp-byol-multi-nic-voltmesh\`.

Upstream description:

Name for GCP certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-multi-nic-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0103031030102332-1302223121311220-1133313131223120-1123133110020130-0333210231121002-2132020112021133-2023210231012021-2012000202100223"></a>

<a id="canonical-0202333111131320-2303303001023332-2312112220000333-1312133220220201-1030020123223123-2011331210110133-3021021031100200-2103333030022311"></a>

## gcp_zone_names property — ingress_egress_gw / 211132000220 / 5

Type: `["list", "string"]`. Computed.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1130312120203102-2122110323320300-1102233103300112-2022130100300120-2110022023011312-0123323310213232-0112231123131301-3101313020332221): complete subsection reference.

- [inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1103010332011311-1122120123211002-0313120122102101-1212112111232323-2002213301120131-2031023110130320-2332020000113011-2101203131100023): complete subsection reference.

- [inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302): complete subsection reference.

- [inside_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3203003222300230-2112031310101033-3303130321221222-3231311000312033-1330123121120320-1222230233122203-2132222101023230-1023322133302111): complete subsection reference.

- [no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202032310301232-3202321000020000-0210202121112010-2120030213110333-3332322130201230-2131302013000120-3331220230110321-1232320231201231): complete subsection reference.

- [no_forward_proxy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0032301122112030-0230102021011322-1331221201233202-2310133101313001-3300103223100011-2101033003202031-3113223310023220-0023332112202012): complete subsection reference.

- [no_global_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1220311131221300-0120302031132331-1130330201130100-0013112232110113-3210123101120022-2312223312233301-1213032212023333-1001202002130112): complete subsection reference.

- [no_inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3323101233232200-1121011001111111-2122331313313332-1103001110103122-3203022213321323-3303022232101131-3320313203232111-0001201212301111): complete subsection reference.

- [no_network_policy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2321230320203023-2012002031123220-0211203302300320-0331010030233012-0113111022300021-1131233223230022-0203113021211023-0321203131330223): complete subsection reference.

- [no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2121122223003320-1111002330210202-1113133312303121-3331201332000322-1210100303320031-3012330310320210-2111211301122021-3130310200112133): complete subsection reference.

<a id="canonical-0233223010101003-2232200201322111-2011132001312231-0231200021103302-1321111023122121-1113313300311111-3222120012313021-1030112301231220"></a>
