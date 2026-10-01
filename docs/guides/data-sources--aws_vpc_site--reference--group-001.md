---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310110002222031-0200321101203030-0120122332210111-3121131021010111-2031032202210311-1223320301101311-2230100320311120-1123233233002202"></a>

## Property reference — Property reference / 033112021310 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- Property reference

<a id="canonical-2020333333033012-1331002022211220-2123312101332311-2320030033222103-3011100103200323-2003321020022100-1230112322123231-3031012223222023"></a>

## Direct properties — Property reference / 033112021310 / 3

<a id="canonical-1211323000212102-2213331213201222-3130302011013222-0210133330330122-3303130003011123-2332201121332010-3330310032113023-0012232213201203"></a>

<a id="canonical-3310010332231102-0130001030103102-2300212022333122-0122323303212131-3210023012323303-2001102211001010-0300233111000223-3323011202212331"></a>

## address property — Property reference / 033112021310 / 4

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

- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-3033120122312321-1232030211110202-1231333230331023-0220211223003123-1001301212123113-1332022300231133-1202101132000213-1323130030112100): complete subsection reference.

<a id="canonical-3332232223101130-2302310023301200-0200121322332011-1123033333332112-1321032122120302-0013212301131000-0313100023110012-0232032302232313"></a>

<a id="canonical-3313000223000211-2011313001332231-2222230203231123-0103222000121220-3222222023123022-3201013331200111-0101210300233002-0110203030233013"></a>

## annotations property — Property reference / 033112021310 / 5

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

- [aws_cred](data-sources--aws_vpc_site--reference--group-001.md#canonical-1021223323310232-3313030321212322-3100201130201122-3102101130303133-3333203221111132-1331100310300302-0012011220131003-1031021302011220): complete subsection reference.

<a id="canonical-1001003221112211-2112333332011312-3003222322023222-0323221320300010-1320001201231322-3330033030331010-2230300123331012-1132011220331320"></a>

<a id="canonical-1222003323132122-0223332212330121-0131113131233033-0123233133221100-1212210302220130-1033130020223012-2011310003133111-0200110130123222"></a>

## aws_region property — Property reference / 033112021310 / 6

Type: `"string"`. Computed.

AWS Region. Name for AWS Region.

Upstream description:

Name for AWS Region.

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

- [block_all_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-2313123122311312-3011032310230121-3320322022011110-0211200103023020-2002012333213223-1303323300222123-2110302033102331-2021103312222313): complete subsection reference.

- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100030222021212-2102031222223121-0110213010203222-3023232123211000-1230132322331222-0300220230232211-3000230113130333-1230031121313300): complete subsection reference.

- [coordinates](data-sources--aws_vpc_site--reference--group-001.md#canonical-2032033333332111-1230303113131311-1210120103333310-1210111013120310-0310221101302230-2123132000131001-3221322310230220-0013231203130212): complete subsection reference.

- [custom_dns](data-sources--aws_vpc_site--reference--group-002.md#canonical-0032303302311232-3330310111332131-0311311232001013-2021201310012332-0323301211013311-2100121202000201-3200002000310020-1011021120303300): complete subsection reference.

- [custom_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-2300130301211001-0301032331023031-0212213222101003-2022302012112121-2122001000222312-3212331321020021-3232021233203302-1103121031022303): complete subsection reference.

- [default_blocked_services](data-sources--aws_vpc_site--reference--group-002.md#canonical-2200133112232132-2102313312112301-0123021223303231-3202333311322201-1101011011131001-3323121000102020-0123222002133021-3010113133031202): complete subsection reference.

<a id="canonical-3220330310000021-2113033112022230-2100123210302011-0211211032203003-3032221132003003-1021102330023323-3301100212222220-3203202132303122"></a>

<a id="canonical-3020033112110000-3203330333021021-0110003013310120-2121300133101333-1212121132231200-0032023203200330-1110322301213303-2201121110121113"></a>

## description property — Property reference / 033112021310 / 7

Type: `"string"`. Computed.

Description of the AWSVPCSite.

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

- [direct_connect_disabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-2121300020212131-1223330033200121-2020110322021122-2022312120022130-1310122021203111-3320211121321222-1200012310123212-0130333023303301): complete subsection reference.

- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232): complete subsection reference.

- [disable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-1311011230011120-1331003133110322-3120131230223200-1013132302300320-0320113132301031-3101000301101123-3122110020101101-1231233022030212): complete subsection reference.

- [disable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-0133220213033001-3001221013333231-2010200113011123-1220313332323222-1112031312223002-3331002110130023-1320101232000122-3311101023223213): complete subsection reference.

<a id="canonical-1022233020223122-1032223111003133-0122100333002330-1312313301323302-3320211303030232-0203101231030112-3002001133031223-3100210321303013"></a>

<a id="canonical-1220013312230130-2211200022103301-2023222030311130-3203121111221211-1213321010212120-3013210102303310-3333123200311302-2013120123232012"></a>

## disk_size property — Property reference / 033112021310 / 8

Type: `"number"`. Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2048,
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
    "ves.io.schema.rules.uint32.lte": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "2048"
  }
}
```

- [egress_gateway_default](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213323122120020-1331223112003223-3322230223312131-0120033112033110-3331033111333311-3020001110302122-2313012032030122-3203230102103211): complete subsection reference.

- [egress_nat_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132102231211003-1132002300321213-0112301333121331-0112133330321102-0012020001310321-3330131223000122-2323012013130320-1132301112001223): complete subsection reference.

- [egress_virtual_private_gateway](data-sources--aws_vpc_site--reference--group-002.md#canonical-1000201021001020-0231311230020321-2302002113210031-1333233021033222-2311232002212123-1223300310301200-2023111102303320-1313212210223013): complete subsection reference.

- [enable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-3223303221330113-3211103231323113-3311331023113321-2023323121233120-2330000301131101-3121201332001302-2000032020123010-2333221101313323): complete subsection reference.

- [enable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-2302322100030210-1000301001133201-0123332222102130-1000222233020000-1321301321221202-1210323123023112-0333010121111021-2003313232323022): complete subsection reference.

- [f5_orchestrated_routing](data-sources--aws_vpc_site--reference--group-002.md#canonical-3030120030303220-2301110223320001-1122111030010130-1111021302102103-1002211331003112-2202112322210313-2221121212123311-2201321223220112): complete subsection reference.

- [f5xc_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-3211023102020300-2302030022120103-1220113121003003-2133023230212130-3221203030310131-2221131101321200-1022323100121112-3321122111312302): complete subsection reference.

<a id="canonical-0123310331310123-1220030132031300-0032003022023302-0232231023312313-2320002202203002-2210101301331330-3320231120132131-2210212202200311"></a>

<a id="canonical-0022223020310311-3321233312203300-0123312320220212-0010330030020212-1222121031221311-3132012233133001-2302010131133110-3321131000030201"></a>

## ID property — Property reference / 033112021310 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131): complete subsection reference.

- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-1302213333133332-3120121201021030-3320202212031101-3102212320020302-1021222113330310-2303011103002003-1100320223011200-3033111010200223): complete subsection reference.

<a id="canonical-0300230120121103-2101322030311003-0313122233130203-2002002220110102-3131030010200110-0103121223100120-2111021012230233-2110223301010023"></a>

<a id="canonical-0301321201011113-1011211111010332-0033031131330110-3122111131033230-1222313120003213-3202333012203000-0023022112111210-3200103011121332"></a>

## instance_type property — Property reference / 033112021310 / 10

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

- [kubernetes_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-1121100213201121-3032003230332001-2003001301202102-0313232032202222-1223201002022130-3210031302233001-2220133000321111-0003233030322221): complete subsection reference.

<a id="canonical-3311321031103103-3302233213200210-2233223002232011-0200203201013310-3001311110223003-3030321020133031-2021333211102323-0210321020312023"></a>

<a id="canonical-1321131312232031-0130210202221112-1131300030221012-3023101032300120-1330210213001213-0201132300001020-1103022020102010-1301010030301131"></a>

## labels property — Property reference / 033112021310 / 11

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

- [log_receiver](data-sources--aws_vpc_site--reference--group-004.md#canonical-0123202023221231-2213300001101331-0031223101030213-1213131032002322-3020031032020212-1310032231031131-1121211200201202-1220021303011330): complete subsection reference.

- [logs_streaming_disabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-0223201103213213-1113333103312133-3323103103220312-3100210332002331-0000332113221200-1202212311000112-0300011110303323-1211330300101003): complete subsection reference.

- [manual_routing](data-sources--aws_vpc_site--reference--group-004.md#canonical-2301211131311302-3233331132303110-2122121313023121-3103110102300023-2300121020332312-0332132310230123-1202102013030223-3332202133211112): complete subsection reference.

<a id="canonical-2222221021230112-0122023133312311-0032311110131122-2230213302023333-3203033222232310-0030021000330001-3113132301111102-2233313313202232"></a>

<a id="canonical-1312210003323103-0013331213231200-2022112021223000-3313321001223203-1221001101210330-3130213111011301-3133130203123313-3021201100230132"></a>

## name property — Property reference / 033112021310 / 12

Type: `"string"`. Required.

Name of the AWSVPCSite.

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

<a id="canonical-2321312332200033-3030022030013300-2013113320303012-0313131220113301-1312100131130223-0331203221213220-3200203000230223-1221001111221212"></a>

<a id="canonical-2302220113100000-3101321121120201-1112121201332303-1111330323203313-2222023131022003-1232200021232213-2231011122320203-0332212333110310"></a>

## namespace property — Property reference / 033112021310 / 13

Type: `"string"`. Required.

Namespace where the AWSVPCSite exists.

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

- [no_worker_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-3312013130023311-2113313033223122-2311313031121010-3032133022333132-0303122113103231-3112213113321231-3011122230033212-3033322313222000): complete subsection reference.

<a id="canonical-0210212023303330-0103223011301222-2301320222202022-1220011033230132-2031023110202211-0333211301203203-2100212130332232-3200331213103300"></a>

<a id="canonical-1231011311331221-0013033133230101-3021110110222212-2131020020301012-0030212321001222-1003331231330102-0013330131003033-2301220323012001"></a>

## nodes_per_az property — Property reference / 033112021310 / 14

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-0101112130332322-3000213012003320-1200022302232032-0133232303003102-2332110132003221-1032211132030122-0313101133233311-1231012223311010): complete subsection reference.

- [os](data-sources--aws_vpc_site--reference--group-004.md#canonical-1212232232313230-2003232012010301-2313210202110132-3231230333322011-1231323013320030-2222122122032201-3301121120121133-2012312210123102): complete subsection reference.

- [private_connectivity](data-sources--aws_vpc_site--reference--group-004.md#canonical-3222113303233102-1200033031120321-3311110103210233-3323022122133110-0110303313203022-0133010333031011-0101331200221333-0023033202333233): complete subsection reference.

<a id="canonical-1030220302100202-3301131213010322-3133211231233211-1023211131132331-2320312213302131-1011023112121223-2021333030031101-2231301123121101"></a>

<a id="canonical-3000123121031011-2302020301113012-2110122120200201-3111122323302133-2301332020320310-2132222220020333-0012030200333122-3100320200312313"></a>

## ssh_key property — Property reference / 033112021310 / 15

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

- [sw](data-sources--aws_vpc_site--reference--group-004.md#canonical-1322310132013020-1222031210323202-3030311333311032-2012310302111012-0021211033122112-1111313033132021-3123122312311203-0130021201313221): complete subsection reference.

<a id="canonical-3002312223130231-3311321030212323-1023301200232101-1100302023330000-2102002311010210-0033222222303123-1200000101002033-2222011021323320"></a>

<a id="canonical-3233230323331112-0232300202023020-1010033200231123-3120223102321103-2112110301132322-3320022100122211-1300211322200130-0220031133113131"></a>

## tags property — Property reference / 033112021310 / 16

Type: `["map", "string"]`. Computed.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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

<a id="canonical-2110322003030220-1301220323332230-1010123031210023-2323002230211030-2322133212013203-3332323130121002-1122213211022110-1231123002133212"></a>

<a id="canonical-2030110120300031-1211202330100032-0102301310130121-3001231001021310-0012310222111223-2221223330022023-3033312133301002-3231231111321002"></a>

## total_nodes property — Property reference / 033112021310 / 17

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-3032000321201331-3010332213001100-2033203210030223-0033302023003012-1131213122303130-0100122111311301-3321301033013012-0131000331232321): complete subsection reference.

- [vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-3010200233311333-3132330201203021-1310330123300331-0121322132302120-2300121303030010-1210301131010120-3130122321223033-3302121033020132): complete subsection reference.

- [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-0122223301132202-1212103311001312-2012320313202110-3130011301320211-0203021111202012-2030020330330203-2202333311031103-1310322103023103): complete subsection reference.

<a id="canonical-3311213202323100-3230303300210203-0033002033021030-1002320133223123-1102003302233202-0231113013023023-0100332312302300-2320233300000103"></a>

## All schema paths — Property reference / 033112021310 / 18

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--aws_vpc_site--reference--group-001.md#canonical-1211323000212102-2213331213201222-3130302011013222-0210133330330122-3303130003011123-2332201121332010-3330310032113023-0012232213201203) |
| `admin_password` | [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-3330022101310001-0020001312220203-1232013010132001-3332331113223203-3231021330332310-2200032310021001-0203001212033103-1223113313001300) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-1233021302022330-0010302123320321-1202031123333013-0332310201222033-2332321221331031-3313130100002130-0212103210013031-0100222120000333) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](data-sources--aws_vpc_site--reference--group-001.md#canonical-2023012233232321-0300330013321022-3221232333131101-1333302202201303-3301302331020031-0321312032330333-1222300132021312-1301211012313213) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](data-sources--aws_vpc_site--reference--group-001.md#canonical-1230103321033323-1000311230103001-3011323330020032-2112330301132213-3101001222001130-1120000033033312-0013310033123310-2200111000013330) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](data-sources--aws_vpc_site--reference--group-001.md#canonical-0312222030132233-3310301232100111-0123102310330112-2312302321111123-2122212332313313-0323312003000103-0331031002000001-2032210003302133) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-2303013020300122-0330200320101101-3133310013313313-0213301301233221-2032213103310013-2201230221023201-0321013132220310-1330220333330112) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](data-sources--aws_vpc_site--reference--group-001.md#canonical-0311012200213322-2231231322030001-3231132033012211-1310003132123202-2011031112020222-0303131233012032-1301032320130020-3030022030102320) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](data-sources--aws_vpc_site--reference--group-001.md#canonical-0312332233132230-3100112230202302-3221100130000211-3023101310011200-1003302330012203-0000022332122223-1332030332012213-0121011001011331) |
| `annotations` | [annotations](data-sources--aws_vpc_site--reference--group-001.md#canonical-3332232223101130-2302310023301200-0200121322332011-1123033333332112-1321032122120302-0013212301131000-0313100023110012-0232032302232313) |
| `aws_cred` | [aws_cred](data-sources--aws_vpc_site--reference--group-001.md#canonical-3231021020323311-3321311101312012-2030200331132323-2020331010203020-3310022220220312-1123003120033332-0021010231232030-1200113112330002) |
| `aws_cred.name` | [aws_cred.name](data-sources--aws_vpc_site--reference--group-001.md#canonical-3301323030101012-1102113133113212-3020101120130233-3201313300210322-2032213022333131-3122301011112322-2103023223213331-0011032201031302) |
| `aws_cred.namespace` | [aws_cred.namespace](data-sources--aws_vpc_site--reference--group-001.md#canonical-0200321110300301-1231113021131300-1110211012012121-3102211123120212-3222231321302202-3110323001312301-2102111200233332-1023300031131013) |
| `aws_cred.tenant` | [aws_cred.tenant](data-sources--aws_vpc_site--reference--group-001.md#canonical-2211210101303331-2111303022331133-0230113323003002-2210123110310101-3002201311110313-2333323112220012-2220231133231131-0110101131131120) |
| `aws_region` | [aws_region](data-sources--aws_vpc_site--reference--group-001.md#canonical-1001003221112211-2112333332011312-3003222322023222-0323221320300010-1320001201231322-3330033030331010-2230300123331012-1132011220331320) |
| `block_all_services` | [block_all_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-3130231112002020-2021330113332310-3331010023201200-0211022020302201-3211023133110013-0003022233113230-3220200331201220-0211232120021100) |
| `blocked_services` | [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-0001311020011020-1011301313323232-1321003233302322-3023023212303232-3000130200201222-3203311213023123-3112231300103210-3200012212130121) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-3110131312102001-2113010021211122-0220302021233203-0203332220133202-2012003300323213-3330120331320120-3100213011031133-3313233223110320) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--aws_vpc_site--reference--group-001.md#canonical-2223021121313210-0031022303333321-2022012323312323-2203231233031022-0010332222021003-0003100201210121-0012230133330110-2022132330010323) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--aws_vpc_site--reference--group-001.md#canonical-0221332310002312-1120130220000100-2031211030223013-3012030022321123-2102213332201220-2223230210220302-2111120313012220-0111033210321121) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--aws_vpc_site--reference--group-001.md#canonical-3201022000300022-1212110110320301-1123000322133032-2110001131000221-0301322112333000-0220223200121122-3013332133010211-1232031101300210) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--aws_vpc_site--reference--group-001.md#canonical-0013301123211332-0232301201203201-3213012300313222-1133233113323112-3020322310321010-3031012131211302-3320222220202000-1130021202113203) |
| `coordinates` | [coordinates](data-sources--aws_vpc_site--reference--group-001.md#canonical-0320320211033001-2200120210223123-3220230320233022-0012302013112130-2303220321111212-0332233311332102-1032133220113101-0122031331203223) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--aws_vpc_site--reference--group-001.md#canonical-1300211331110330-3200322013003001-2131203121103200-2003032310010212-0013100213213231-3323332202100233-0232000312012130-3222101011310131) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--aws_vpc_site--reference--group-001.md#canonical-3310122100011313-1213030023302103-3002112230133002-2302000233210303-0133330313301202-3320231022121221-1323100022332221-1001301001020220) |
| `custom_dns` | [custom_dns](data-sources--aws_vpc_site--reference--group-002.md#canonical-2212002130231013-1033332212111033-0211201200312220-1313230133131012-1310232000132033-3102021023002113-1003101313002010-2333112133122201) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](data-sources--aws_vpc_site--reference--group-002.md#canonical-1323301221102022-3113121202012302-0131203010003133-3032213031210201-3312031331012112-1003001200121112-1333003220202123-0231021033232313) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](data-sources--aws_vpc_site--reference--group-002.md#canonical-1000011301312023-1021000211032110-1332020211102320-1132202130213210-2310213332110213-3131333321202302-0001331031313301-1233330330303201) |
| `custom_security_group` | [custom_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-1121011102222323-3322023132122003-0013003300201310-1312120312020130-0333011110300202-1310003220203203-2200232313110112-3301122020333333) |
| `custom_security_group.inside_security_group_id` | [custom_security_group.inside_security_group_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-2033132312200213-3231302023100123-1333100301023333-1231202221033313-0330210203210303-2013221122032202-1200003323123133-3202330030210311) |
| `custom_security_group.outside_security_group_id` | [custom_security_group.outside_security_group_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-0313010001212103-0220113123311231-0120223030332113-2002211031330301-1233302122103233-0220300022202103-2201232023311030-1001202001230133) |
| `default_blocked_services` | [default_blocked_services](data-sources--aws_vpc_site--reference--group-002.md#canonical-0112331121222233-1000312311001331-0210012000332221-2210212100220022-0023032332321003-2000320301231132-1122103102220310-1021000111332010) |
| `description` | [description](data-sources--aws_vpc_site--reference--group-001.md#canonical-3220330310000021-2113033112022230-2100123210302011-0211211032203003-3032221132003003-1021102330023323-3301100212222220-3203202132303122) |
| `direct_connect_disabled` | [direct_connect_disabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0230120113020202-3002110201133130-0023100321311303-3031130012211013-1023101213112233-2001001103133321-0312300231323201-1202110111100110) |
| `direct_connect_enabled` | [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-1003222102201001-3322220302231232-3032320012233011-0310221002013003-0002232213300320-2003023331011120-0111223300102311-1211022322032201) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](data-sources--aws_vpc_site--reference--group-002.md#canonical-3333211130113111-0103322130220332-0033231113121232-2201112102123110-0110333213202223-2000102111210203-3211303031132133-0303111222123331) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](data-sources--aws_vpc_site--reference--group-002.md#canonical-2013130000322121-1002232032001220-0122100110331101-1212223232112230-3013133220313322-2032032331121201-1322131122323330-3121331103013210) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-0323011012131010-1223322301222233-3122202013332330-2303232123133123-2012211232013012-0212323122213112-2301032330211110-1112131223011001) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](data-sources--aws_vpc_site--reference--group-002.md#canonical-2320203000313101-0200211132221131-3200131330210023-2200110011022312-0303230020030112-3112223320023132-2220131312232301-0312000122303001) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](data-sources--aws_vpc_site--reference--group-002.md#canonical-0230211131223020-0320022110322101-3021213030001213-2302333002112021-1223133331211032-1120313213223303-3022031022301311-1323303100122131) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](data-sources--aws_vpc_site--reference--group-002.md#canonical-2212133331012312-2123311222310223-2000311103211300-1131010301103313-0032202233131010-2322110101110232-2110213133123120-1331300030300133) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-2001132101300200-3303210330121030-0012131003200201-3100130220120203-2300221223102022-1130211233221121-2323233201202101-3121102203321200) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](data-sources--aws_vpc_site--reference--group-002.md#canonical-3102320100102231-0101003132311131-3033030021001211-1003301200130210-3020233310123322-0033020313012210-2222031110231102-3200210030321020) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](data-sources--aws_vpc_site--reference--group-002.md#canonical-3312331110203111-3222322103313213-1022221302030022-3230011011030221-0311031122002301-0123213102023022-3101230300003332-3101033112201011) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-0131033003100310-0221112230001000-2020001131312211-3033130102132132-1332303322103311-0211030301120200-3332313233100300-2312331130100113) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-0101123103332032-1010000310331301-1121223310112020-0010032031333311-1233102000103032-0013030032211231-3303303020310213-2200003103110011) |
| `disable_encryption` | [disable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-3333332333310302-3203122103030103-1010121230131302-2021312232332001-0122100332023223-0302120122122002-2003012222033202-2301121101301303) |
| `disable_internet_vip` | [disable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-0021030323003211-3121302001232032-0230232102301131-2001112232012020-1132302110233312-3102032112301200-3131320213100232-0020132312211121) |
| `disk_size` | [disk_size](data-sources--aws_vpc_site--reference--group-001.md#canonical-1022233020223122-1032223111003133-0122100333002330-1312313301323302-3320211303030232-0203101231030112-3002001133031223-3100210321303013) |
| `egress_gateway_default` | [egress_gateway_default](data-sources--aws_vpc_site--reference--group-002.md#canonical-0010223221310313-3303122223302011-0022110130011022-2211001211220102-3130010221213111-0133321112302300-0301210232221023-3211311222220003) |
| `egress_nat_gw` | [egress_nat_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-3321232120013301-0202012303233001-1022000032100312-3231303133121103-0122112200232331-2013110113103211-3023000112301310-3213310223020211) |
| `egress_nat_gw.nat_gw_id` | [egress_nat_gw.nat_gw_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-2320223132031100-1102302010030101-2100001101133322-1220113203322332-3332331001232200-3321130320131001-1232323232200332-0023311211331000) |
| `egress_virtual_private_gateway` | [egress_virtual_private_gateway](data-sources--aws_vpc_site--reference--group-002.md#canonical-1330021311030311-0203012201110310-2321200310111012-1002130132111022-0300323113020320-0133203201113022-3222332331020013-3311133130202322) |
| `egress_virtual_private_gateway.vgw_id` | [egress_virtual_private_gateway.vgw_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-1101120211201321-2003123131131103-3031321210211103-1021303220212032-0333330030023103-1021333222113033-0110303132001320-2210232123223322) |
| `enable_encryption` | [enable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-3301022002013232-2020123112233212-0232010322323330-0023202033111233-1220132121033331-2303302023323102-0031120112211312-3203021002103132) |
| `enable_encryption.kms_key_id` | [enable_encryption.kms_key_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-1232230023300303-1321123331233301-1203212203030103-1312102220310100-1211212230323200-1322003310203303-1013000233311231-1331231021010031) |
| `enable_internet_vip` | [enable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221323001223200-0012211123003113-2321022213222002-1233322021220122-3111231100211013-3113330211333100-2201020333330213-3312100003311200) |
| `f5_orchestrated_routing` | [f5_orchestrated_routing](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302133212120110-3332133210022111-1322211233133221-0032300011331303-3200101323211310-3233200321212100-0221021012213130-2013223021332011) |
| `f5xc_security_group` | [f5xc_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-3112003020000123-2331132210333332-0230132300013222-3330011302011032-0023332023301030-2103113111223203-1232023323213100-2223022203302023) |
| `id` | [id](data-sources--aws_vpc_site--reference--group-001.md#canonical-0123310331310123-1220030132031300-0032003022023302-0232231023312313-2320002202203002-2210101301331330-3320231120132131-2210212202200311) |
| `ingress_egress_gw` | [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2011112031122032-2010321010101232-0202223030133130-2013300013202300-1102313232023013-2212030132231233-3232123330120231-0303021133200033) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-3310220132111100-2103213210221332-1220222210311011-0211120201312213-3311203212202111-2332303010213022-0322213013132321-3101121121130021) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-1220231300111323-2122122323132130-0303223221232310-3002103333210110-3223221130200333-3302320232201000-2003133131231232-0321202000020032) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-1223103302302320-3311010310133020-2103102232313001-3222231231011322-3113030323121012-2020201211133300-3300210301000330-3001223313313020) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-2210121113112212-1331232112223102-1110030132111302-1302220000331010-3122211203311203-2200031322121201-0102223100012333-2233130112301231) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-2013201221321122-1300100220322021-1122331223001301-1312202110121023-2313302012330130-2331113003200331-3223220303310111-3131003322011120) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-1322030213302301-1132220131112111-1210222302310300-3330233022312010-0330000303111110-0330313300111221-2032112323122300-0203310310300030) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-0022212303203032-1331231221331302-2302121311000333-1330222221302013-2300220102101321-1201232223212200-0130320110110122-2221231221031232) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-1003001030110310-2301013202131021-2303112313333132-1211320023310222-1001223101112032-3303331201322322-3010310202102133-1230301033311233) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-1122001012213121-0202333132120220-2312102130100130-3303010101220101-1221033001122123-0323323011000011-1231113023020203-0102103322031231) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-2231022233310302-0321211010023220-1332312021211331-1312222002002320-2303021123020123-3102210010210331-0200030232233212-3202323223102323) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-1031332322223321-2122012200301301-0102132103000030-0211332011320033-0130332311311023-3313023032131213-0003200312311022-0230332211300313) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-3100320202102303-3231101312302322-3211332321321320-2301022333121032-2123033000001202-0000300220023202-2021233121133220-1121100130223231) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-1010010232212111-3201331213021132-3122232212010322-0013133103203113-0131120322322032-0133032132301303-1321303010201101-0303220011230231) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-1021312312200020-2212000321233201-3213221102211313-1202333301331033-0313321100112122-0022232010112303-2220213103110012-3301321302101031) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-1130332003332201-1003112323032332-1023302230111300-3003312012233311-0203211331003223-2320231320030032-0012010103323112-3303300112011220) |
| `ingress_egress_gw.allowed_vip_port` | [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-3202313033200222-1003300232102131-0213200031020232-3102112223010303-2102230232101112-0111231303222020-0201300111111030-0223310201010322) |
| `ingress_egress_gw.allowed_vip_port.custom_ports` | [ingress_egress_gw.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302023102331033-0002203132200123-1113102320130212-0011023030202223-3003211331103000-0332232223200120-1013102023012310-3312312000321110) |
| `ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges](data-sources--aws_vpc_site--reference--group-002.md#canonical-1223212130331211-1011232323031221-1013021230033320-0102212102020120-2003023133311220-2032212321032121-0111120221131321-2301200330223200) |
| `ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0131332103322321-1130303230023332-0203233320222103-1313121021202003-2201200100213022-1000333032203022-0032222032320231-3100302312230032) |
| `ingress_egress_gw.allowed_vip_port.use_http_https_port` | [ingress_egress_gw.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-1330010201312331-2202331212132233-3313223222022203-0120122232013002-2122102032031131-0103332200201102-1033232113010212-3101002022132300) |
| `ingress_egress_gw.allowed_vip_port.use_http_port` | [ingress_egress_gw.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-3323010223310002-1202313113020221-2120320211231131-2113300132010112-3333033121131130-0013113001112121-2201002311333033-3123211320222312) |
| `ingress_egress_gw.allowed_vip_port.use_https_port` | [ingress_egress_gw.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-2220012303210121-3111203022031103-3301233003312303-0002233110210333-2312022330221033-3223202230102330-2230310000010032-3322012232333331) |
| `ingress_egress_gw.allowed_vip_port_sli` | [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-0033311122010213-3312110111112230-0000002112000010-3301321001313311-1233030112330001-1131020102133203-1322001210303011-2032321022131101) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-2131123211032032-2322103201300232-2112200220120113-0201302003312031-1121223222033123-2200330313103222-2310003231213023-1122221302011210) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges](data-sources--aws_vpc_site--reference--group-002.md#canonical-3322001121321231-0120320330330311-3310223303331111-1033303032132000-3030232203130021-3223323313213331-1033202133033213-1301303123023131) |
| `ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-1211133033030213-3221111222032002-2033221012331102-3100020310130133-0320321023313120-3030331313221000-2130210001021223-1100323332033021) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-1011133320030323-0300230332113022-1112123220130331-2222031333230300-1103323303211020-2202122123220132-1230300010132332-0231022121032212) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-3203000002130110-2201231022011233-3300213232121100-3233211321132211-3200321113103221-0122101033231201-2210323111023122-0011023222331023) |
| `ingress_egress_gw.allowed_vip_port_sli.use_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0022321102023312-2021323023313322-0332222202112120-3203320313302121-1300013322031113-2211201131021010-1201030223303310-0322033300000310) |
| `ingress_egress_gw.aws_certified_hw` | [ingress_egress_gw.aws_certified_hw](data-sources--aws_vpc_site--reference--group-002.md#canonical-3202220032021113-3121113100123331-2310320212022133-1122203113313333-2230111310013310-2300003313231211-1211211110001223-2101023133033021) |
| `ingress_egress_gw.az_nodes` | [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-0221320323213323-0101130220031013-2211010201000011-2030223210121012-2012012110221332-0031203102301111-1021011320301103-3022110223232130) |
| `ingress_egress_gw.az_nodes.aws_az_name` | [ingress_egress_gw.az_nodes.aws_az_name](data-sources--aws_vpc_site--reference--group-002.md#canonical-3302300101011221-1202202123312333-0321012103121203-0211323130303101-0001233323213302-2310222322133110-0133311110003212-1032021322130023) |
| `ingress_egress_gw.az_nodes.inside_subnet` | [ingress_egress_gw.az_nodes.inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-0301110103132331-2202211003132221-2032322122221232-1332200122123232-1212022130333230-1020301331213032-0003311212332121-1030332000312323) |
| `ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-3212331220131203-1232133222032000-1301330213113330-2203112003210102-1211322302222023-3321322323221123-2113320332022302-1321232133302121) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-3032331122200310-0323103122333333-2032121011031001-2201013011101230-3003030300320303-0021302123100122-2321130312113212-1130031112221111) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-002.md#canonical-0210312310202231-1223003310031321-0131212130320220-1031001301321333-3000202011120213-2020131310221320-0133002031102122-0110010213110312) |
| `ingress_egress_gw.az_nodes.outside_subnet` | [ingress_egress_gw.az_nodes.outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-2321012212231200-3332233201112312-1001322030222320-3123130011303201-0030300233102302-1000023002000002-2021131012000332-1132331002100132) |
| `ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-0210103310220213-2212301113333221-2201003113220011-1020332201000213-2100110113321132-3233023332311022-3300103203001002-0031111310032313) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-1011230000323301-1301033323301033-0303233212230113-3130101133331202-3220211121323133-0110002012110331-2023023230022321-1130031203211311) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-002.md#canonical-1233030211230113-0231020322133010-2230201013102331-3010101210221230-3333232130212003-1331132022310320-1111321220111222-1120202011310013) |
| `ingress_egress_gw.az_nodes.reserved_inside_subnet` | [ingress_egress_gw.az_nodes.reserved_inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1223100113003021-3112331210101221-3103323020200021-1020011203301211-3130221303221022-2233030131301200-3121202120212220-3322322312303322) |
| `ingress_egress_gw.az_nodes.workload_subnet` | [ingress_egress_gw.az_nodes.workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-0320121012111322-3231121212313121-2223103102130321-0131003111310302-1111301203121033-1030203222020212-1012021122023132-0001122012013331) |
| `ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-002.md#canonical-3202120122200232-2303003302031011-3113000120132310-2221212200010022-0322311233010121-2001221133231123-1201001232133111-2030013221011310) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-2330202003311321-1011010230300003-3302223332320001-0223212112123213-1313030100212300-2120023013322202-0331113002023023-1321110133101030) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-002.md#canonical-1321323022303123-2313232033133020-1103302330313023-1302102132322012-1300202303232012-2031013020133020-0330323321120321-3332112200303020) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302331131303331-2003213203230200-0302032003120111-3231001202033002-1311103302112122-2332213202100220-0312323030132112-0123101131321033) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-1201200112200110-1000133102332202-0221020033331130-3121031030023202-3213300113233211-0023220330111320-2031102123320202-2002323331002101) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-2101011020221013-1312222221103110-2130113323103232-2012222232100032-1110303020201101-1220010320102032-3031323010332002-2313030300000000) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-3321113030032013-0003122220221010-3110113221011123-3200233223322201-2222332123111113-3000111030333323-3101320202012001-0130110311003330) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-3202300011231220-1031202201123333-1221211021110122-3121333121101200-2002102011103211-1121011010232111-3201132030203031-2222110230231323) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213011232312310-2300331121103012-2333303113113211-1021020133122032-2101223313230223-2303310231333131-0022131100000231-2010213101133322) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-1133012000322023-2220321300020101-1312331022320311-2023011003031323-3203221000032230-3303221111323132-1312133202001220-2122203211120122) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-0002113223211221-3223203131112101-1213001112213101-3312223121203220-3320311321003010-3120212001131332-2302200032322110-3311030023231013) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](data-sources--aws_vpc_site--reference--group-002.md#canonical-1023330021231200-0102330000330323-2201031231021122-2232100002013320-2302312300303221-1303232310131202-1200310111313232-0231310211231110) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-1111020122200100-2003231202312120-3213312210123202-1032200103301020-2032302011322030-1112300202322231-0121121231130212-0020300302102022) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011201202332022-3301311210123331-0332113333321312-3300001233031103-0133332232032033-0202300102202200-1013111030030021-1231312100013210) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-1001202011001323-2023103110103013-1222223112113203-1232323303231022-1001022003112102-1023010303202222-1201321131212221-1331011010312213) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-0221223303232220-1301220022022311-0331122212002010-3012111100020110-3332002001022321-2203312113213222-3023213132133131-2110231002021303) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-2033113131232000-2022231331110203-2132302230212220-0020312122323303-1332231112002221-3320331222211113-3233202112121100-0233231001223202) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-0322103231002102-3321313112101333-3230111102221001-2220310303212010-0033021212021223-3321321022002131-0333113323302020-0112301320312001) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-3021122133211133-0213322101231031-3003020120032002-3231220132300020-0230011230333002-2021313202210100-1012003001211013-2232011023022210) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-1132131011000122-2332110130022132-0133210332310031-1033331013221213-1223323023203323-0320203211231102-2110312221312023-1133103122102201) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-3313320113322121-1320330312310102-3232120123013221-0123030321101221-0331312031012010-0300032013123101-3211213111013132-0130112332130202) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--aws_vpc_site--reference--group-002.md#canonical-3211233322330322-3200232330320010-3133032301223102-0302321321231223-2330231123310213-1013100132123202-2321011110220223-0111333322022211) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--aws_vpc_site--reference--group-002.md#canonical-1022300321120201-0313112112330211-1232330102201330-1030022201213112-3031112020333201-3211201323303330-0112203110300313-3002103201101233) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--aws_vpc_site--reference--group-002.md#canonical-0100320133013231-0011213330130211-3300303023313031-0121002303333110-1320321330321121-1332201101000320-1223103023302102-0121312020020232) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0303201210020022-3112000332220201-3110101202301103-3120212323223123-0320233213033023-1211321233331100-1233231023012310-1323110022223330) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3223303033002210-0022100213110110-2031211220002131-1321301310202212-2123013112230213-0111001120033312-2301023032010232-0300323010022020) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-1202023013102313-2001210030220202-1012121233031113-1311123012322123-2032010023111033-3021200120231320-1200010330202200-3030122012113130) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_vpc_site--reference--group-003.md#canonical-3202332310023032-0122320302331321-3223232200001332-0311203212222323-2223020313222132-1013200120233112-3000002122211031-2233333033021131) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--reference--group-003.md#canonical-1011332011000223-3232021330300110-2031311312101212-2231121210211113-3001003322322032-0323112013103101-0123220120032021-2111232101133010) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-2212001302131213-3301322201333302-2221031301132011-0100213023312332-2033311130322022-3023220010002110-2030133300313203-0012101333223010) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_vpc_site--reference--group-003.md#canonical-3320113303031103-3331022333010302-2131310032011133-1233221123002021-3022123031103313-2133321300010020-0213221032203101-1101020003133233) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_vpc_site--reference--group-003.md#canonical-2011310213121320-0030132300023123-1233023113303212-0022003012100121-0332211113200111-0303300030300321-3333020220230100-3012103320210232) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_vpc_site--reference--group-003.md#canonical-0223010130003011-3230132101032012-1130332121232303-1112020012302320-1223332221120303-2223333322323013-1022123312232213-3310121232000102) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_vpc_site--reference--group-003.md#canonical-3230210030232131-0100313133121230-2203231023013020-0223001020121101-0013002033020303-3303130103130212-2110220022110231-0001022110202003) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_vpc_site--reference--group-003.md#canonical-0322023013130313-0321122233302232-3323313311330101-1000303302021201-2031213231203023-0220303110331012-0322231221203323-3331021100021311) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_vpc_site--reference--group-003.md#canonical-0303012320021131-1223211200203121-3102200322323200-1323303101023303-0232333130123200-2001020121313100-3211033211211130-2130003221120103) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-1003313213312333-3203022002112212-0312203220321323-2331211202033020-0121330033213312-0302101023110001-2302233003310313-1011313102111223) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-2232302230031203-2001121332210120-0100111330030302-2130212023231120-0311023333220302-2122203120111222-0130120210323013-1211303000023231) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-2122331132310200-0311102121133210-2123310020100323-0121322103223322-3301013321301112-1310020010303200-1321321110223202-1012023002223021) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-0233303332223011-2321122113222132-0111023200202220-2022301030111133-0223120202113232-1233221203022313-2322032330322210-0231102133001103) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-2003332033112320-1133013122310032-3031331101011101-2112220313110100-2121033231022321-2022122312322031-2132222103231202-0332322332300312) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-0111031220122120-0303132311113010-2313320101300023-0103101023112322-0102103022002100-1032033221133320-0123113310102123-3211312200103223) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-3312311231300221-0231301103122221-0113100321031003-1033201203313102-2312121102023213-0000331222200312-1312303202103221-2221101311222320) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-0110312233313020-3113203331323302-0021211130012011-2013032133303312-0030011030122233-3020122222131302-3031331123031230-3321030002023302) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-2313301203310203-2200011331303120-3230030212300302-0233130300311232-2030312000313032-1331330303110002-2201332332021311-2013201210312320) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-2130203101322122-1300021311203121-1313323201221012-3100023002222012-3020200300113220-2110232023021103-2200202103301022-2221120313232211) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_vpc_site--reference--group-003.md#canonical-3211303330121003-3032323113202133-3022031021021131-0322021323013332-3210123330212012-1030020302120320-1312030023101023-2132320323321131) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-2122200210312103-0223331310120031-2121202031213023-0122203230210302-3233023233303211-3222222200103111-1220013223032000-3331301220113301) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-0312111313032003-2031011103010023-0013020023300003-2010210332131021-0313231301332123-1310001110110022-2230023221233102-2303231002301030) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_vpc_site--reference--group-003.md#canonical-1010122101031211-1021032123220222-3320302330101131-0121301130001122-2100023102233311-0021022222003000-1200100100213111-3230331132122010) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_vpc_site--reference--group-003.md#canonical-3001223213230113-2032323133301320-1012012222332022-3123330121313000-1201330210133003-3213132210030301-1231301221330210-3210222333000000) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-0232203033002011-3220010100013203-3322313131320323-3303121131121333-3010123203220202-3003010032022312-2321202130201320-1102200220233003) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_vpc_site--reference--group-003.md#canonical-3022100200030312-1333132012223332-2320023113131123-2222223303013200-1333303130032102-2221110021210332-0123033220002122-0012322213301312) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_vpc_site--reference--group-003.md#canonical-3111231331213302-0013323322130133-0022200321032213-0013320320300101-3003002221323032-3100311122212211-1020323131113221-2020022020121213) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-3331302211212012-2213121333131332-2302222120203011-3212213232112212-2130332121001202-3003322311101322-1211012020210013-0111211232100123) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](data-sources--aws_vpc_site--reference--group-003.md#canonical-2211210312030103-0200011301120030-0222121203033211-1131233312000102-2200330312200311-0022031003202322-3023022202220211-0100313201022231) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](data-sources--aws_vpc_site--reference--group-003.md#canonical-0300230323332103-2321221333323132-1303221023121013-1320233201132331-2103111302103131-2231203203310302-1010212303132101-1232001233332032) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](data-sources--aws_vpc_site--reference--group-003.md#canonical-1311103311110320-1323110302101001-2113332322101123-0300311232310233-1020313323030212-0001311230111210-1210322232323031-0022013131332023) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0232311121302013-1331020113311123-1013112310011320-0032100223000202-0132110323300013-2220101302321112-0223101031012321-3131101102121311) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](data-sources--aws_vpc_site--reference--group-003.md#canonical-2333111013203202-1013230111323233-3230112231011201-2130110312220200-1321202212323212-0303233022220332-1220000102131033-0022233132003111) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0110003023310303-3132020021102112-1132323232032132-3122322210220310-1220000333013130-3133121031203321-3200020012230131-2130231130331002) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1333002123321123-1223022100322122-1203123000033321-0103313310110323-1023012211133212-3200322100022111-1331022121132122-1323232113102203) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-003.md#canonical-3231031330121321-1220322300111003-2231002230130302-2100110021321111-2002020132220130-1301303322000303-1002031300131022-2220101221120220) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-0031213112111102-0133130003131133-3313200223113332-0102321022023232-2113101301110133-1030002220123230-1031303021021032-2113221321202213) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_vpc_site--reference--group-003.md#canonical-3120222030013233-0003031200133121-1200122320303300-0320013103333211-2322321033203010-0103100000021213-2231201212330121-1302012333332030) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--reference--group-003.md#canonical-2030221023203032-2111320232330210-1021231100132023-1013013113302012-1131221122310111-2133222010122120-2310012313232022-0121132033132101) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-003.md#canonical-0131001232003312-2101011301303310-0011333231213122-0212322303302112-3102330300031223-0101003002011033-1320010123000331-2123030232003130) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_vpc_site--reference--group-003.md#canonical-3001021122202022-2300212333100312-2332022133111121-0101031131302012-2210213313311312-3011232320220300-3213213031320022-3301231110032303) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_vpc_site--reference--group-003.md#canonical-3210303323020332-3122203122120213-3020011200330032-2223120330002003-0211310301223003-1001111103230303-0123001302112203-0300321033213001) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_vpc_site--reference--group-003.md#canonical-1001010020001221-2011102311223010-0233231202203021-0311232302033331-3300011120231023-1331111211202130-1022213013213000-1130032300030312) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_vpc_site--reference--group-003.md#canonical-3220323030013321-0322010323133010-3110333112333112-1010211022011102-1130000020120001-0110123021333013-1112332311220023-2203001132100330) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_vpc_site--reference--group-003.md#canonical-2330303113200222-3201231033000013-1120220111123200-2212010101010013-3202102012012102-1333331133312010-3111011012313212-1001311321131131) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_vpc_site--reference--group-003.md#canonical-2032320120333302-0110012123331010-2121002010233320-0211220311233012-2233023211221322-1310111110121313-0011321232001210-0002222200201321) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-003.md#canonical-3113020201132233-0001101021311220-1202311302222213-3223311123220223-0021010321232231-0203231121012230-0030203230002020-0112122233032223) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-003.md#canonical-0313132331203131-0121232001311310-1313232301210231-2011013322200003-0112112023020201-0102033210331221-1031031201323021-0303102030013332) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-2323313020123101-0121123122110321-2333330212221210-2023332011111202-1101211122330123-1200223312013111-1213012203011321-0121231031202230) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-2222222211302300-1320032002231002-3101130023211302-2300033023323203-2220003011330213-0013023231111202-2002101103113102-2310200222213310) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-3103310110011221-0220001331023321-2210032201000332-1100301010201121-1130302133331111-2312212213331230-3102121131323322-2102212202121122) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-3123311020212103-2132312301211221-2233201032130322-0332033001330023-1130113111233233-2101112132031331-1322212202003022-1200303313111021) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-1132021100000313-1320101113220233-0103103220033201-1211302001123201-1323212021001100-3330022001213203-2332213010333320-2330202221330300) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-0020323323312111-1311112211320332-0321103333023100-1233010320010120-1101021302200131-0003322101121130-3103101331333132-0203321220021013) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-2111003000031110-1231302102110033-2102113033201330-0203330003032023-0211220100313103-1023203333330031-3320330130012213-1333312330112033) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_vpc_site--reference--group-003.md#canonical-2111310121221311-1022313311313231-1210031001121312-1200312231112032-0133032232110213-2330032131320332-2002121221230122-0111120032322110) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_vpc_site--reference--group-003.md#canonical-2132211123102210-2323013223330102-3023212100223212-1000323111111111-1310332121131111-1311132312330321-3030112322121213-1020202313013303) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-003.md#canonical-1113332033200311-0121101231101103-2202131323103201-0330120331203232-0003013202312111-3223112011312102-1212332113032000-3230112321203103) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--reference--group-003.md#canonical-2230121203020333-1301032321113120-3031232023233110-0011000132131210-0022323020023202-3101111310313230-2320331001322313-3132311213030233) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_vpc_site--reference--group-003.md#canonical-0321012001211203-0020323333302000-2101221022202030-0313232200322301-0011032213310130-1113313002113022-0000320130230320-3313121332303303) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_vpc_site--reference--group-003.md#canonical-2000122210202301-2032201111012212-1130123012013210-2210201103312022-2333333020130103-2322323231332013-0102020212310002-0121011132113032) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--reference--group-003.md#canonical-3223112313021333-2132302123132202-3101200330331313-3312022102202003-1210030031202003-1230210012121320-1111301002133233-3000013132003131) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_vpc_site--reference--group-003.md#canonical-1330012003021032-1332332003201100-1301023211012011-1020200112030220-1013320022131223-2320212001311320-0321210123213113-3232301213013331) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_vpc_site--reference--group-003.md#canonical-2110333211030022-2120022223113233-1333000322211201-0130011202231102-2033122311212223-2103002013321231-1032112020111310-3313002301313113) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](data-sources--aws_vpc_site--reference--group-003.md#canonical-2100332222003021-1313101132213132-2000221233332030-3123311333211132-2111330223032011-2320102332030011-3100332112013031-0331013002220111) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-1000022222111013-1101030000222200-1330102312303203-2330113330323212-1331221321000333-1101033122221110-2130221133222011-2031332022023202) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-2020230303210112-1103012312113022-2112010030201210-3000011132023232-1322121233001120-1233333331210122-3000000111010110-0232200321231213) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--aws_vpc_site--reference--group-003.md#canonical-2131233032001322-1220313110210102-1112033102100132-1303112310111311-3202011232203201-0330323211023302-0323200023103203-0311012310130133) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--aws_vpc_site--reference--group-003.md#canonical-2013222301311230-2012221000330123-2131033113333012-2032331310321230-3133233120023033-1130233100021222-3030003011213332-2333012200030020) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-003.md#canonical-2002133213032303-1003301102230332-1313300003012022-0323333101111123-2000220313200201-3302300320123122-2310113211032203-1100301111332313) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--aws_vpc_site--reference--group-003.md#canonical-3002200332030030-0013201012203103-2302221102231130-1331021233201132-1200302211302113-3320222022202332-3331303111030301-0320001122102002) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--aws_vpc_site--reference--group-003.md#canonical-3112311003230003-0320321132131021-0110001230202100-1102003123010013-1330222111111023-0033233303012203-2210230302123013-3010023021113121) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-3320023331003320-2112032331232113-1013011122122022-0133023120012002-0212311210233022-2303030011112023-2330011121112333-1201113311110311) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-0120333321313220-1323333211030301-3200022223300310-0301211201110312-1220133101332230-0321012221333101-1230121331031031-0103121331211333) |
| `ingress_gw` | [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-0211201331003211-2132210332011032-1003100231033223-3330301033201113-3213013312201031-1131230330130022-1313121113320212-3103122022210202) |
| `ingress_gw.allowed_vip_port` | [ingress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-2202201322020111-3232021323231100-0321213223112133-3031301232111231-1220012211000223-1320012021333310-0012322212302222-0202332322232231) |
| `ingress_gw.allowed_vip_port.custom_ports` | [ingress_gw.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--reference--group-003.md#canonical-1202101003321332-1020201310011232-2203311111002123-1231031103103203-3331203102222123-1213123311001001-1223303130133312-1012210121011213) |
| `ingress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_gw.allowed_vip_port.custom_ports.port_ranges](data-sources--aws_vpc_site--reference--group-003.md#canonical-1311121133210302-3330122010123120-0301332210311002-2231303120202203-0211030012332130-3100221332120010-3300320220133210-1212203112230303) |
| `ingress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_gw.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-0210133201012121-2101321101330200-2023212223333123-2113101130011310-2233132121111231-2003021011000211-3102121323230102-1020330130301132) |
| `ingress_gw.allowed_vip_port.use_http_https_port` | [ingress_gw.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-1103011202102001-1301021132122103-2020031132231303-1233200201033330-3100330303123111-1110222202111101-0031211211132013-1202112133210012) |
| `ingress_gw.allowed_vip_port.use_http_port` | [ingress_gw.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--reference--group-003.md#canonical-1323112202301300-0330111132223322-2003132322301300-0111312100312032-3033032033210221-1132303112021101-1330333320312023-0102201003201120) |
| `ingress_gw.allowed_vip_port.use_https_port` | [ingress_gw.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-0330333301003213-0021203332300022-3300122232223202-1323131210231011-0211010032310032-1201102022312120-2021132320211031-2201202310313033) |
| `ingress_gw.aws_certified_hw` | [ingress_gw.aws_certified_hw](data-sources--aws_vpc_site--reference--group-003.md#canonical-0332130213330220-3213103201230030-2122213301102202-2212110331110030-0312330332300103-3222021232133112-0001332210031123-3122133230320001) |
| `ingress_gw.az_nodes` | [ingress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-0302020232331013-3112202021002230-0013013331221103-2010130033032001-1333200021323210-3003032233131311-0212112100333222-1310223112100203) |
| `ingress_gw.az_nodes.aws_az_name` | [ingress_gw.az_nodes.aws_az_name](data-sources--aws_vpc_site--reference--group-004.md#canonical-2103010223103013-2303332211100313-2111223222231122-1323122213202323-0100312212100233-1222333103022020-3331010313311133-1302201012203203) |
| `ingress_gw.az_nodes.local_subnet` | [ingress_gw.az_nodes.local_subnet](data-sources--aws_vpc_site--reference--group-004.md#canonical-3211112021230323-3320212330030312-3312011321323130-2112101312331300-3013121322133102-1221322322110310-0001002122122121-3110330203001021) |
| `ingress_gw.az_nodes.local_subnet.existing_subnet_id` | [ingress_gw.az_nodes.local_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-004.md#canonical-3130312321232200-0330002012320010-0222122110011131-0022003013022220-2021112102011022-2120330110302013-2112102131131030-3110133203033310) |
| `ingress_gw.az_nodes.local_subnet.subnet_param` | [ingress_gw.az_nodes.local_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-004.md#canonical-1123000202010332-2210133021102323-2202231331123211-3122232021010013-0210310131033011-3113221001303002-2203310232322233-0313131211032231) |
| `ingress_gw.az_nodes.local_subnet.subnet_param.ipv4` | [ingress_gw.az_nodes.local_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-004.md#canonical-0322300131102312-1112032122210321-1102103232313233-3011121331320311-0120200233222001-2003321221001100-0000322031200020-0122303100001200) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-3012320232200303-0011031000221220-3100110000322013-0302022303132131-3101311033233000-3333323310012123-0112123120030212-1031021123332303) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_vpc_site--reference--group-004.md#canonical-0002312303131010-3330230001223312-1122121320311131-2331302201330323-1011333012122321-0120120000223032-1010200223330010-3201112103232210) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--aws_vpc_site--reference--group-004.md#canonical-3112232101011320-2210133130330332-3110030101022003-0000133212333003-0013301100232312-2211222310020132-1020213111032320-3133300322212030) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--aws_vpc_site--reference--group-004.md#canonical-0110120302120323-3211131312331223-0130001110031002-0032010312322313-2030301102102322-1303101123301331-1020031013322211-2110110100131203) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_vpc_site--reference--group-004.md#canonical-3100022211101323-0233213113312101-3130013000121010-1303033310321122-3322122201323321-0221010303033011-2200100233300100-2213111221320231) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-2130220023301112-2102002122221111-1010312010301302-2210021202001130-2012012030232300-3333210212103321-2330211230220032-2010132120102032) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-3002321203102112-1212010231223223-2312330312223301-2302022010223202-1011012123301331-0021133132112222-0021331202102211-2301110321311201) |
| `instance_type` | [instance_type](data-sources--aws_vpc_site--reference--group-001.md#canonical-0300230120121103-2101322030311003-0313122233130203-2002002220110102-3131030010200110-0103121223100120-2111021012230233-2110223301010023) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-2130110022300212-1131322112201210-1100023022020301-2101220220022003-1111212303301003-1323232030022211-0313012011333121-2312311013220011) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-2113021310300011-0300031032013303-0202211010233231-1011031010220013-3321202310123203-2033200330301120-0322220202133313-2331003001220103) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-1002012331201030-2200023123022133-3012130312030021-3220323332222333-3122222212313201-0111031111130333-1202231013032310-1000301222022030) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-0133210323330111-1232311031121002-1231010312202001-3302133323022221-3330131210023320-0012020111033221-3321032102222203-2130232132331301) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--aws_vpc_site--reference--group-004.md#canonical-0210303301213022-2332332333120001-0310113003023333-1032312310310212-0200010230230322-2323331210303101-2232033013203002-0011120010110302) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--aws_vpc_site--reference--group-004.md#canonical-1110030223002320-3030223123021332-0031102220203012-3032312123301213-3200133311012223-2211232123011030-0231330212320030-2213113112323120) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--aws_vpc_site--reference--group-004.md#canonical-0330222123032211-3123212231220300-3000131332322303-1311210233130331-3232201020222110-3023012310100032-0201031022230331-2211023210200310) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-3323133010213221-3033022300223313-3022211110102031-0312223313220103-2013300032211223-1210210200221110-2321122131302012-0321102033322101) |
| `labels` | [labels](data-sources--aws_vpc_site--reference--group-001.md#canonical-3311321031103103-3302233213200210-2233223002232011-0200203201013310-3001311110223003-3030321020133031-2021333211102323-0210321020312023) |
| `log_receiver` | [log_receiver](data-sources--aws_vpc_site--reference--group-004.md#canonical-0022331211010330-2012012212323330-2320320211300121-0201230033102200-3301102022121310-3333320122211313-0232021211131312-0221323112300332) |
| `log_receiver.name` | [log_receiver.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-0100222213133100-0222011301301101-3103111231100130-1131101231032032-1020333212132102-0203223303222101-3120112301210132-2021000322130031) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-3030200302112300-1113202321300231-1331031033121200-2001200002111013-0110211301221202-1030011120122331-3321210313333322-2300301301031220) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-0111011033010220-0202131211020130-0220212211333201-0222202333000231-3232033222120301-0332120121203012-1322222313103013-3013122112033013) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-3123232330111120-2311110203310310-2201212220001101-1222332101320130-1021201113033000-1120231010100123-1303113032010200-2130303213122030) |
| `manual_routing` | [manual_routing](data-sources--aws_vpc_site--reference--group-004.md#canonical-0000121213033322-0302020313201122-2333221311101301-1103013210212003-0023002322232130-1211330203023202-1000313013333112-0200002200213232) |
| `name` | [name](data-sources--aws_vpc_site--reference--group-001.md#canonical-2222221021230112-0122023133312311-0032311110131122-2230213302023333-3203033222232310-0030021000330001-3113132301111102-2233313313202232) |
| `namespace` | [namespace](data-sources--aws_vpc_site--reference--group-001.md#canonical-2321312332200033-3030022030013300-2013113320303012-0313131220113301-1312100131130223-0331203221213220-3200203000230223-1221001111221212) |
| `no_worker_nodes` | [no_worker_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-1122113221103131-2230021331210223-3223012321302103-3110131131120001-0223012200120332-1212232113222132-3021103003313020-3201230023102022) |
| `nodes_per_az` | [nodes_per_az](data-sources--aws_vpc_site--reference--group-001.md#canonical-0210212023303330-0103223011301222-2301320222202022-1220011033230132-2031023110202211-0333211301203203-2100212130332232-3200331213103300) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-3323001031022102-0220233121030313-1232030130302220-0223231213033100-2123031222020202-2133202030321310-2323102133020011-0233130113213310) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-2121300221102102-3233022231323103-1101302313003120-3202303132230213-1021113001203102-2122201302130231-1101000210212322-1012222223313120) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-1320100013310322-2331200103123001-3113130001022231-0203000200320232-2122102003230010-3103111312233302-2110230303333231-1320000101320121) |
| `os` | [os](data-sources--aws_vpc_site--reference--group-004.md#canonical-2133202122330233-0111231100133130-2100212222030011-1331030132333222-0223320012332300-2311210232233112-3213010200311113-1131213121021002) |
| `os.default_os_version` | [os.default_os_version](data-sources--aws_vpc_site--reference--group-004.md#canonical-1302232102200003-0103230311023121-2100233100203230-3103221310013202-2213030033120012-1112231202222110-1131113101331230-1322211232110223) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--aws_vpc_site--reference--group-004.md#canonical-1200030000013112-1000111110231233-0103201001320001-1210110011011013-2132201201302122-3220000030101033-3100032132211301-1002033301112201) |
| `private_connectivity` | [private_connectivity](data-sources--aws_vpc_site--reference--group-004.md#canonical-0210232211322203-1122311101121013-1311303122000210-1112330120000233-1130020213201030-0203112031031330-1233221031331220-1202301003333132) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--aws_vpc_site--reference--group-004.md#canonical-2212012223323031-2101323032102121-1231102332320003-3202111021333110-3013320322203033-1022321310313330-0313213230003021-3201231213001230) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-3313112331033201-0322023302003112-2011320213200313-0021022200302003-1310033121212303-0122003123030203-1231230220210121-0201110023132301) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-3210211332102331-3132333312113112-1002231011232231-3130332012302222-1200033122302032-1033330333330013-3300331122122320-3033133130212032) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-0102020122110002-2003311030200110-0300303310110330-0013023321201302-2321303021222123-3211032302000231-0231212133012103-1303033022223301) |
| `private_connectivity.inside` | [private_connectivity.inside](data-sources--aws_vpc_site--reference--group-004.md#canonical-2332310213010322-3131220131030331-3101121333230022-0213311132200211-1131033002023101-1013023331313321-2102232111321200-1303233101231031) |
| `private_connectivity.outside` | [private_connectivity.outside](data-sources--aws_vpc_site--reference--group-004.md#canonical-0200030011221232-1300101310013303-3311223301102113-2100133111210113-2303130100310131-1220321232132003-2122203313323110-1223310300202103) |
| `ssh_key` | [ssh_key](data-sources--aws_vpc_site--reference--group-001.md#canonical-1030220302100202-3301131213010322-3133211231233211-1023211131132331-2320312213302131-1011023112121223-2021333030031101-2231301123121101) |
| `sw` | [sw](data-sources--aws_vpc_site--reference--group-004.md#canonical-3223130023000230-2133101330211111-3121213133000233-3330210032211311-0111200310000303-1320033213112013-2201321330012331-0103103010101232) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--aws_vpc_site--reference--group-004.md#canonical-1223233103001333-0031100100301320-2122022230121102-3023202312223100-2222011320322003-1031300120023330-1031003023333103-3212312131120220) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--aws_vpc_site--reference--group-004.md#canonical-0110003211021213-1223021013301203-1320222010333200-1222223102132001-0010100222320200-1320300000230012-3110002030303213-0031302303023122) |
| `tags` | [tags](data-sources--aws_vpc_site--reference--group-001.md#canonical-3002312223130231-3311321030212323-1023301200232101-1100302023330000-2102002311010210-0033222222303123-1200000101002033-2222011021323320) |
| `total_nodes` | [total_nodes](data-sources--aws_vpc_site--reference--group-001.md#canonical-2110322003030220-1301220323332230-1010123031210023-2323002230211030-2322133212013203-3332323130121002-1122213211022110-1231123002133212) |
| `voltstack_cluster` | [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-2012232113013000-2100230113230223-1220200212223012-1300311102110030-1332220233303002-1112200023110211-2211321033232213-3123211213032200) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-2131202000130311-3121313000000310-3302332221111112-3313312310123221-1310303021201123-2003101333021133-1211003103333223-1111111112131323) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-0221102303301020-3132122131101231-1002131310323120-3312311330311230-2230203203311001-3203302011231312-0131220000121220-2110303201132212) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-2223202312331022-1310200123321121-3131120310201333-0223213101000013-0220130233010232-0331101110302132-3210313120002200-0330212022033211) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-2110330220112012-3012120331231000-1312100231100131-2220123323210001-2223321122031302-3332013000310111-0101311012003030-2103230210021110) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-3200000030112332-1323233320323231-2222010302031102-0202220212031031-0302333011300323-0303132331222200-0102033003203101-1300321300130323) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-1131031232023312-2202032123333330-0012031022003203-3221123203103232-3112310133003033-1130132213130323-3312201213101010-0331020102202221) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-0032121123210300-2311113002312320-2021000300211031-3122031200020111-0121222230022201-0100213230013020-3230330101122200-3332113222310220) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-0023113000102132-3000331300121230-1121303321331311-1121102002123310-2210210231121011-3030011120033230-2120023101222211-3303332322000023) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-1023233122301001-3130130303201132-0232122023200223-0113130100132312-0033311331310330-0312231331333333-1003232313011011-3211221132122003) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-3210223023110102-3200332311203013-2122321231102112-1312331322033320-3321033332130232-1113302001313332-2320033230033210-0211021300223223) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-1331230231101202-0113203321213301-3123110323203020-1122121023113123-1101023003122220-2023200132112023-1120313113330031-3212322213110023) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](data-sources--aws_vpc_site--reference--group-004.md#canonical-3123202222213131-2310120122122331-0321300233123211-0132111022312323-1222330033113300-1203333101231130-2330130211130213-1123120111303013) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-1012330031123330-3021203230022331-2113330202033212-3320202332122300-1222303010111110-0120313031210211-3300122321000322-0203300221302021) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-2133032013210030-0012000022300320-2013011130201011-0333022020312110-1232100212220303-3003120101221331-2123030320300322-1012022213012332) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-1130331103123320-1210332322231031-1213213033333322-1302210012121012-1031102330011023-2032230223332332-1331302001121131-3301302132001212) |
| `voltstack_cluster.allowed_vip_port` | [voltstack_cluster.allowed_vip_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-2303101132100130-3013033221000211-1120011232211230-3223031203123301-2200330303201212-0110122121222221-2303002101120320-0031000213211013) |
| `voltstack_cluster.allowed_vip_port.custom_ports` | [voltstack_cluster.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--reference--group-004.md#canonical-0231113021021213-2102113302301111-1320331032133313-0222303100000322-1232000321203101-0310331300313132-0031033203310001-2001231003232123) |
| `voltstack_cluster.allowed_vip_port.custom_ports.port_ranges` | [voltstack_cluster.allowed_vip_port.custom_ports.port_ranges](data-sources--aws_vpc_site--reference--group-004.md#canonical-1312130233130133-1330012133221001-0021013003330233-3310300232030010-1220003033200030-2200220233311310-3111312001200132-2120022301200312) |
| `voltstack_cluster.allowed_vip_port.disable_allowed_vip_port` | [voltstack_cluster.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-0322320233230101-1200232133331320-0022303202010303-2221103003201322-0232031103223022-0303130300301230-0333132322213130-3211200222322232) |
| `voltstack_cluster.allowed_vip_port.use_http_https_port` | [voltstack_cluster.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-0032212202223313-3033231313323231-1322023322233223-2322303010002323-1231000211112133-0021300113111030-0132112221102122-3313323232303302) |
| `voltstack_cluster.allowed_vip_port.use_http_port` | [voltstack_cluster.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-0032030033210131-3210011011300030-1010112122210220-3332001121333003-0223202011133123-1100300230332011-1210102101203120-3203210221031111) |
| `voltstack_cluster.allowed_vip_port.use_https_port` | [voltstack_cluster.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--reference--group-004.md#canonical-3203002321110121-1212003221031230-0331133033210213-3110033201210111-0332102032120320-1100311312323211-0100020000113300-2213101202031232) |
| `voltstack_cluster.aws_certified_hw` | [voltstack_cluster.aws_certified_hw](data-sources--aws_vpc_site--reference--group-004.md#canonical-1201010122321301-3232223001232330-2030221011232302-2233103112333321-3210121331123233-1321033021211303-3031100023230321-3023011230003320) |
| `voltstack_cluster.az_nodes` | [voltstack_cluster.az_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-1111033002100333-3223201022013103-0300202132120013-0022313210122330-1022133010010002-1210013301012301-2301211303333300-0323332111003203) |
| `voltstack_cluster.az_nodes.aws_az_name` | [voltstack_cluster.az_nodes.aws_az_name](data-sources--aws_vpc_site--reference--group-004.md#canonical-0202320032010121-0312113211101033-1310032002210201-0012100113303022-3213300321100331-3011202102230202-1230322032103000-1010222203102213) |
| `voltstack_cluster.az_nodes.local_subnet` | [voltstack_cluster.az_nodes.local_subnet](data-sources--aws_vpc_site--reference--group-004.md#canonical-1031101233231222-3211212300021330-0333031023020230-1033300230212013-1132022000201002-2222311211110011-3133302130131221-1131030332310220) |
| `voltstack_cluster.az_nodes.local_subnet.existing_subnet_id` | [voltstack_cluster.az_nodes.local_subnet.existing_subnet_id](data-sources--aws_vpc_site--reference--group-004.md#canonical-1312233130330231-2203223330311200-2201322313110231-2102300102021311-3123220321330022-3112101113321200-0021033231111113-1202011321130330) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param` | [voltstack_cluster.az_nodes.local_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-004.md#canonical-2033210001121232-2210202210113022-3331121021311120-0211211002023003-2230332212002223-2212031122031303-0220320230232123-2230001121303002) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4` | [voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4](data-sources--aws_vpc_site--reference--group-004.md#canonical-2113100000320323-0011313110321031-1021230003003233-3202130102313020-2230320212032113-1110023221331010-1232000002131312-0132023020002112) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](data-sources--aws_vpc_site--reference--group-004.md#canonical-0031101110312113-0322012313233023-1300300101232012-1210000020113301-1311021013210302-3010232200301000-2211232331122120-0303112301131220) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-3212321312212332-0132331122113111-2031131112323222-1301120222331111-3122121323110020-0323002101021212-3122322112232133-3132331011021301) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-0203220200313332-0100122112122022-1032033100003011-3331122110133023-1222102023001000-0003322012232311-1121121000111220-1001212021010110) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-0212121101002310-2322002231303222-0010013000123302-0212131202101011-3100122100000131-0001323133123321-3231323223121000-1332020120230322) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](data-sources--aws_vpc_site--reference--group-004.md#canonical-3131110011203030-0202032010011200-0020020333100102-3303110221313132-1031330203112320-3132330303012320-2320132021033022-2021200020023221) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](data-sources--aws_vpc_site--reference--group-004.md#canonical-0122333212233030-0220002032010112-1212113333213213-3322113233222032-3120111132120311-3101131320223323-3203302112211310-0323122223011003) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-3111032233002111-3123131031000011-2201032221212021-1223111300222230-1010313022200330-3212030311103212-3303233021210132-2111003321032023) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-004.md#canonical-0202020011330133-2230222112131123-0012113323300332-3012210321020112-0120230002100110-1223120130010112-2032120223231033-3102133003333210) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-004.md#canonical-0232202130013313-1032311130132232-1202202201231132-1030221133333230-0010111101313132-3333321201131020-3330213123312230-1121133312201202) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-004.md#canonical-1130131111212130-1131203201002130-3110203201023020-0302333222232001-1000210113010302-2103230220303312-2213312232233220-2200033113312013) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-2303001002313320-0132133023300200-3111233220203300-2102203113200011-0122231012301103-2030010101223120-0311213111212232-2323233032233013) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-3321230011232110-3323031100333222-1111103122202030-0121310231200130-0200233011331211-2333313022030232-0120030000020233-1213123311333311) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-3130113333312212-1021130021322011-1231332311020033-2311103222323302-0112303302303130-2213331302020132-1132200212013223-3000032320122303) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-004.md#canonical-2330222300110013-0113010111221232-3303001010123323-0121011210233032-1021113023302330-3103003021123020-2020223122013132-1222030123230033) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-004.md#canonical-1010023211011210-1033200001010001-1123302233203301-1303222221030320-0100300110133001-3032221113112132-3103310122232011-0122012012003202) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-3103211320220012-3203210103002003-0210223332021302-0023010020131333-2332232111300012-1311210203003121-2221012130101300-0202001221201322) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-1131213013203102-2101133333133130-0123311111131030-1030022020221010-2000213331201101-0102123032301331-2002301013301132-3001200210130213) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-1100213130221220-3211201033303302-2131233202110000-2321213333033022-3102323111122322-1230100203030301-0132312222023003-2230001312320021) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-2020200121201332-2012110112231231-0101133320210330-0201232220000002-3003320221330020-0313101333122011-3132332323300311-2133001232320210) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](data-sources--aws_vpc_site--reference--group-004.md#canonical-3002303130030123-1232123212133132-0100201013120223-0222220310111131-0022321021230333-3320210013002212-3021211210001030-1023113020310131) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](data-sources--aws_vpc_site--reference--group-004.md#canonical-2111031121303022-2122232321123121-2320000323131013-0001103331320011-3320023111102312-0030333311111320-1311313201312212-2022233110132311) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](data-sources--aws_vpc_site--reference--group-004.md#canonical-1120321333130133-1322323331311303-2302003132323023-2220303203012000-3002200101003121-0103321302110112-1212113331030121-0322032203031120) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](data-sources--aws_vpc_site--reference--group-004.md#canonical-2110302133302022-3010301111302333-2012230232201013-3312203312312201-1121100033312321-3333223031121110-2313211100121030-3333312212330203) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](data-sources--aws_vpc_site--reference--group-004.md#canonical-0002331023130010-2210211200212312-0100321213333331-0021302103102303-3032230103201102-2011211122210230-3303021220301333-1131223123020210) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](data-sources--aws_vpc_site--reference--group-005.md#canonical-2330210202200302-1111210033031011-1023000312330010-1022302303130100-1123020230203231-1210103200323010-0330020211302231-2012032302101030) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](data-sources--aws_vpc_site--reference--group-005.md#canonical-2133003213300012-0001012012133131-3320123011112130-3120121330102203-3203101222310013-0330023022132102-0030303023122223-3033023322033312) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](data-sources--aws_vpc_site--reference--group-005.md#canonical-0033211312312310-0132132023031310-0201111220202303-0000011322203300-2013103101112012-3000202120122111-1101130302001012-3033303030233132) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](data-sources--aws_vpc_site--reference--group-005.md#canonical-2122020111123201-2133203233131011-1312003033103111-1331023000323312-1113121031133211-3123232212232333-0311321322111133-3332232222301013) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-005.md#canonical-1311232301002310-1111010323100233-1330121313012203-3220211121323220-3232131330210102-2200101023132320-0231332023123102-3330330000230133) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-005.md#canonical-0202031211013232-3322303011123111-0121220221120332-2231022023203302-1001100332220231-2220111110120200-1303032023102013-0011330111102330) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-005.md#canonical-1211130322110200-3023102210031012-2213323223011212-1220320220302223-1211311312021203-2122023002001333-1202230223332001-3022310002102021) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_vpc_site--reference--group-005.md#canonical-0122330032232223-2111003303320123-1221102333301100-1030132032103002-2322322313120113-1120212000130323-1021033200321132-0032002132122211) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--reference--group-005.md#canonical-3023100101030132-3203111321312033-2120232121210230-1230201110333111-1111132131102012-2002023130131231-3132222011101132-1003003331210031) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-005.md#canonical-0221000201222111-0011303133210012-1212202021202323-3122000211022213-1000101002320112-1312032303222320-2023203030120102-2131120231231012) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_vpc_site--reference--group-005.md#canonical-2221011101123130-2011033133231121-3300010122232323-0203121311110322-2032103110103030-1023320021211030-1323013032222020-3332332223212200) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_vpc_site--reference--group-005.md#canonical-3330210120313200-0202023011011011-2313302231112113-2022003231130001-1000031233200020-3233132000301222-2212010000311233-3331213211112312) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_vpc_site--reference--group-005.md#canonical-1033123100011103-2123010202110102-1010113310332123-1011210211321211-1001210303333030-0202213212123011-3132121013111301-0132101331330000) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_vpc_site--reference--group-005.md#canonical-0233132301113010-3223200202002003-0212321211311123-0132003020313311-0033330012213201-0000120121113001-1012230302110303-1101131321021130) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_vpc_site--reference--group-005.md#canonical-2000131133311211-0231031321103133-0022200331133123-2333112020233213-1023120223230123-3332311032310120-1211231100233302-0302303113100022) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_vpc_site--reference--group-005.md#canonical-3020000233032010-0023132013232021-3002023300223131-2313310321132020-1121333112331330-3221231101312012-3100111130203233-2013103333330321) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-3030031332022233-0213121021120231-1100320232201122-1020321233201100-3232303131121011-0022101120222121-2131321012033221-2033213032100030) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-005.md#canonical-3023102132013102-2312212333313302-0110231322101320-0110310220331311-3020010312230232-1233300332233221-3032223322220202-3102332023201233) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-2020200132100201-2220303212000210-2222321212332223-0322313310011313-0000132320320123-0023031131332123-1222230212121100-2221020100233313) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_vpc_site--reference--group-005.md#canonical-2331230322031032-0113013232222200-2311011223113213-2221230211121031-1311013010001310-1203200111112303-3303020111230301-2013133022013323) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-1113223321020333-0213001213330023-3233012312100322-0112033100103232-1321120020332220-2203233020311030-2011002122011310-3320320233112201) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_vpc_site--reference--group-005.md#canonical-2303123333030221-1312202201000030-1233213111231213-0122131010022032-0231223012110321-3311111223223032-1113032220203132-1023132203012311) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-1313020221110303-3300131331202210-2210302233232012-3320302321230122-0311231020231322-3131113310033121-2010013111113110-3310011220130002) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_vpc_site--reference--group-005.md#canonical-2120312120011032-3332132121012321-0212012103210001-3231333123301303-0130331330103213-3312013331211110-2302303023031003-1210002332333331) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-2301331211311122-3023202120220211-3320030123013303-0322120322222133-3202333113230231-3333212302111132-2202121221002013-1200230202303001) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_vpc_site--reference--group-005.md#canonical-0300303320133110-2311023002002311-2110011011310122-2201111312020332-3310102002322211-3312121021021131-3032221131120230-1133033223323221) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_vpc_site--reference--group-005.md#canonical-0332003312011111-1211111000301301-1220233232321221-3302313223003222-1231103022103010-1230321200200021-3212333322030333-3113031210013323) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-005.md#canonical-1213023303213220-1130213331330013-0203011322302203-3310002213222332-3301312122021332-0000220002112113-3211211012323312-1010211330323033) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-2230010212020012-1300220200213311-0112123032020000-0302023001123221-3112223311130320-3003001320132111-0132010110102000-3122221202123122) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_vpc_site--reference--group-005.md#canonical-1222122321202020-0222313332313012-1330003121111230-3223101200231233-3303112320022223-3030212203032001-1101210323102001-1221320301330222) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_vpc_site--reference--group-005.md#canonical-0100221310013322-3310102132102333-3323300002231331-2303131023102321-0022110031330030-1000002002023332-0202323323230200-3323021000002212) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-0031111120003110-3201300131222321-0031220002112312-1311322211310311-0232213230021203-1222213331301002-2233212211001021-3102112032322020) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_vpc_site--reference--group-005.md#canonical-2223000200333023-2232102011023021-0331223202212220-3033131021030101-3012101320212203-0132002113323320-1223001033032000-0112233031033100) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_vpc_site--reference--group-005.md#canonical-2300013311110033-1011333303112102-2331332010231110-1032200131100232-0120001101022020-1332321330021222-2002023303020002-0023111112311133) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](data-sources--aws_vpc_site--reference--group-005.md#canonical-0332300021001332-0200132133133123-3000322023301030-2002300101200001-1100022131013213-2030321203012211-2232133202210231-2131202301133113) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](data-sources--aws_vpc_site--reference--group-005.md#canonical-2113201113200021-2312321022223312-2120320222330001-2003122123023023-1301320221002213-0312301113033100-3003232102221310-0130303132323103) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](data-sources--aws_vpc_site--reference--group-005.md#canonical-2032112231001310-0103323103133212-0201202220032100-2210222101111220-1122231020230330-1122331030013120-0133120133002331-2322131330210301) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](data-sources--aws_vpc_site--reference--group-005.md#canonical-0013012320021101-1110323311333310-0232022302322300-0122221112122111-1321311010232123-0003111022331010-0333011013122120-1020302230033030) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](data-sources--aws_vpc_site--reference--group-005.md#canonical-2123220230133133-3001321232321123-3121130203320300-0003100133123112-1030131131113003-1030122330020011-2313033303110003-1312312033202012) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](data-sources--aws_vpc_site--reference--group-005.md#canonical-2212010333301000-3003011230321232-0220030320012201-2002320200220233-1332212231031333-0113301331203221-1212110123003232-2210233302210122) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](data-sources--aws_vpc_site--reference--group-005.md#canonical-3131310023032311-0022122213320331-1133123322101112-3103102120330320-3223321010212023-1023002231221013-0023301320332313-1231112313112102) |
| `vpc` | [vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-0001122022122131-3320031231112123-3320233031332010-0222332213211010-2210003221003321-3033212311212312-1023130012032303-0201100130311301) |
| `vpc.new_vpc` | [vpc.new_vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-1011203310010210-0010100002100123-3121021233210110-1130312322112130-2131230013011011-0033320331111310-0002130201133132-3111221303003112) |
| `vpc.new_vpc.autogenerate` | [vpc.new_vpc.autogenerate](data-sources--aws_vpc_site--reference--group-005.md#canonical-1033011331000001-2213313133313131-0033222333023010-3221132011201333-1012321200313110-1100203201132230-0300131011120232-0323102201021201) |
| `vpc.new_vpc.name_tag` | [vpc.new_vpc.name_tag](data-sources--aws_vpc_site--reference--group-005.md#canonical-3220320232000303-2200023000010322-0113303312223201-1202332010012001-3202100102332232-3311210111022012-0031302030133130-2321202031322033) |
| `vpc.new_vpc.primary_ipv4` | [vpc.new_vpc.primary_ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-3311101112231321-1322221022122222-0310132012320202-2222223132022033-3302202123330310-1033133021133203-0032320023220031-0130123300131230) |
| `vpc.vpc_id` | [vpc.vpc_id](data-sources--aws_vpc_site--reference--group-005.md#canonical-3102123310211022-0201320200102221-1022332232212203-3221132013300123-0130113203100131-2030321331223032-0232033311031311-3303210202311221) |
| `waf_signatures` | [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-2032233333323322-1311021321011330-0111021230030330-3131003102113331-3120022200111231-3323130000300322-2121031203223300-3230302322312011) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--aws_vpc_site--reference--group-005.md#canonical-2330130312201120-2001100002203123-3302202212031221-0132033330003210-2232313202111231-2002331302302220-1101313321222300-2313011133203012) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--aws_vpc_site--reference--group-005.md#canonical-2102033003213311-3132013000321100-1103132312123311-1220200110030100-1222011201321301-0023223100003210-3120022313220300-2013310023003112) |

<a id="canonical-0132330022321200-3111211133303322-1022001031033002-3231301312200200-1111330011213130-3123303303103311-1033310331311021-0222001123120202"></a>

## Next pages — Property reference / 033112021310 / 19

- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-3033120122312321-1232030211110202-1231333230331023-0220211223003123-1001301212123113-1332022300231133-1202101132000213-1323130030112100)
- [aws_cred](data-sources--aws_vpc_site--reference--group-001.md#canonical-1021223323310232-3313030321212322-3100201130201122-3102101130303133-3333203221111132-1331100310300302-0012011220131003-1031021302011220)
- [block_all_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-2313123122311312-3011032310230121-3320322022011110-0211200103023020-2002012333213223-1303323300222123-2110302033102331-2021103312222313)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100030222021212-2102031222223121-0110213010203222-3023232123211000-1230132322331222-0300220230232211-3000230113130333-1230031121313300)
- [coordinates](data-sources--aws_vpc_site--reference--group-001.md#canonical-2032033333332111-1230303113131311-1210120103333310-1210111013120310-0310221101302230-2123132000131001-3221322310230220-0013231203130212)
- [custom_dns](data-sources--aws_vpc_site--reference--group-002.md#canonical-0032303302311232-3330310111332131-0311311232001013-2021201310012332-0323301211013311-2100121202000201-3200002000310020-1011021120303300)
- [custom_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-2300130301211001-0301032331023031-0212213222101003-2022302012112121-2122001000222312-3212331321020021-3232021233203302-1103121031022303)
- [default_blocked_services](data-sources--aws_vpc_site--reference--group-002.md#canonical-2200133112232132-2102313312112301-0123021223303231-3202333311322201-1101011011131001-3323121000102020-0123222002133021-3010113133031202)
- [direct_connect_disabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-2121300020212131-1223330033200121-2020110322021122-2022312120022130-1310122021203111-3320211121321222-1200012310123212-0130333023303301)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- [disable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-1311011230011120-1331003133110322-3120131230223200-1013132302300320-0320113132301031-3101000301101123-3122110020101101-1231233022030212)
- [disable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-0133220213033001-3001221013333231-2010200113011123-1220313332323222-1112031312223002-3331002110130023-1320101232000122-3311101023223213)
- [egress_gateway_default](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213323122120020-1331223112003223-3322230223312131-0120033112033110-3331033111333311-3020001110302122-2313012032030122-3203230102103211)
- [egress_nat_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132102231211003-1132002300321213-0112301333121331-0112133330321102-0012020001310321-3330131223000122-2323012013130320-1132301112001223)
- [egress_virtual_private_gateway](data-sources--aws_vpc_site--reference--group-002.md#canonical-1000201021001020-0231311230020321-2302002113210031-1333233021033222-2311232002212123-1223300310301200-2023111102303320-1313212210223013)
- [enable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-3223303221330113-3211103231323113-3311331023113321-2023323121233120-2330000301131101-3121201332001302-2000032020123010-2333221101313323)
- [enable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-2302322100030210-1000301001133201-0123332222102130-1000222233020000-1321301321221202-1210323123023112-0333010121111021-2003313232323022)
- [f5_orchestrated_routing](data-sources--aws_vpc_site--reference--group-002.md#canonical-3030120030303220-2301110223320001-1122111030010130-1111021302102103-1002211331003112-2202112322210313-2221121212123311-2201321223220112)
- [f5xc_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-3211023102020300-2302030022120103-1220113121003003-2133023230212130-3221203030310131-2221131101321200-1022323100121112-3321122111312302)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-1302213333133332-3120121201021030-3320202212031101-3102212320020302-1021222113330310-2303011103002003-1100320223011200-3033111010200223)
- [kubernetes_upgrade_drain](data-sources--aws_vpc_site--reference--group-004.md#canonical-1121100213201121-3032003230332001-2003001301202102-0313232032202222-1223201002022130-3210031302233001-2220133000321111-0003233030322221)
- [log_receiver](data-sources--aws_vpc_site--reference--group-004.md#canonical-0123202023221231-2213300001101331-0031223101030213-1213131032002322-3020031032020212-1310032231031131-1121211200201202-1220021303011330)
- [logs_streaming_disabled](data-sources--aws_vpc_site--reference--group-004.md#canonical-0223201103213213-1113333103312133-3323103103220312-3100210332002331-0000332113221200-1202212311000112-0300011110303323-1211330300101003)
- [manual_routing](data-sources--aws_vpc_site--reference--group-004.md#canonical-2301211131311302-3233331132303110-2122121313023121-3103110102300023-2300121020332312-0332132310230123-1202102013030223-3332202133211112)
- [no_worker_nodes](data-sources--aws_vpc_site--reference--group-004.md#canonical-3312013130023311-2113313033223122-2311313031121010-3032133022333132-0303122113103231-3112213113321231-3011122230033212-3033322313222000)
- [offline_survivability_mode](data-sources--aws_vpc_site--reference--group-004.md#canonical-0101112130332322-3000213012003320-1200022302232032-0133232303003102-2332110132003221-1032211132030122-0313101133233311-1231012223311010)
- [os](data-sources--aws_vpc_site--reference--group-004.md#canonical-1212232232313230-2003232012010301-2313210202110132-3231230333322011-1231323013320030-2222122122032201-3301121120121133-2012312210123102)
- [private_connectivity](data-sources--aws_vpc_site--reference--group-004.md#canonical-3222113303233102-1200033031120321-3311110103210233-3323022122133110-0110303313203022-0133010333031011-0101331200221333-0023033202333233)
- [sw](data-sources--aws_vpc_site--reference--group-004.md#canonical-1322310132013020-1222031210323202-3030311333311032-2012310302111012-0021211033122112-1111313033132021-3123122312311203-0130021201313221)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-3032000321201331-3010332213001100-2033203210030223-0033302023003012-1131213122303130-0100122111311301-3321301033013012-0131000331232321)
- [vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-3010200233311333-3132330201203021-1310330123300331-0121322132302120-2300121303030010-1210301131010120-3130122321223033-3302121033020132)
- [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-0122223301132202-1212103311001312-2012320313202110-3130011301320211-0203021111202012-2030020330330203-2202333311031103-1310322103023103)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3033120122312321-1232030211110202-1231333230331023-0220211223003123-1001301212123113-1332022300231133-1202101132000213-1323130030112100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011110020131111-1233101032100131-2330022303222002-3213013233113320-1230001000333031-1123310012203203-0120022022221010-3012020212330010"></a>

## admin_password — admin_password / 100302003312 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- admin_password

<a id="canonical-3330022101310001-0020001312220203-1232013010132001-3332331113223203-3231021330332310-2200032310021001-0203001212033103-1223113313001300"></a>

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

<a id="canonical-1031210001101122-3203110223133222-0002120213123311-2130313133031213-2030121232321220-3030302013300330-0203102300232020-1211020133230232"></a>

## Direct properties — admin_password / 100302003312 / 3

- [blindfold_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-3313032310120233-2031130211031133-0231300330333211-1030213210001001-0020113330100312-2123301303213010-2003102033220102-0010031202121333): complete subsection reference.

- [clear_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-1331231323220010-0121312121321032-3322332003211211-0231331310013203-3220330003223212-2200323001301102-2012220331112223-1122302003331203): complete subsection reference.

<a id="canonical-0011033012333120-2031301011310230-3322112312130020-2120332213323102-3120003023310033-0312100012101221-2301113300112130-2123013301100130"></a>

## Next pages — admin_password / 100302003312 / 4

- [admin_password.blindfold_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-3313032310120233-2031130211031133-0231300330333211-1030213210001001-0020113330100312-2123301303213010-2003102033220102-0010031202121333)
- [admin_password.clear_secret_info](data-sources--aws_vpc_site--reference--group-001.md#canonical-1331231323220010-0121312121321032-3322332003211211-0231331310013203-3220330003223212-2200323001301102-2012220331112223-1122302003331203)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3313032310120233-2031130211031133-0231300330333211-1030213210001001-0020113330100312-2123301303213010-2003102033220102-0010031202121333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001103323302030-0122000122313000-0132211320210022-1222321313220022-0212313201001333-3300302011323032-2321021101131020-0232120020321201"></a>

## admin_password.blindfold_secret_info — blindfold_secret_info / 123211022210 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-3033120122312321-1232030211110202-1231333230331023-0220211223003123-1001301212123113-1332022300231133-1202101132000213-1323130030112100)
- admin_password.blindfold_secret_info

<a id="canonical-1233021302022330-0010302123320321-1202031123333013-0332310201222033-2332321221331031-3313130100002130-0212103210013031-0100222120000333"></a>

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

<a id="canonical-0131330323303312-3312031122223001-3111232202023021-0110221131032301-0222213123323113-1311032210202213-3100112133033231-1302133033032121"></a>

## Direct properties — blindfold_secret_info / 123211022210 / 3

<a id="canonical-2023012233232321-0300330013321022-3221232333131101-1333302202201303-3301302331020031-0321312032330333-1222300132021312-1301211012313213"></a>

<a id="canonical-2322003230210231-0033000300011101-1333131000020323-2000313203333230-3301022330101003-3131313101003012-0201201022330101-0132312333323221"></a>

## decryption_provider property — blindfold_secret_info / 123211022210 / 4

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

<a id="canonical-1230103321033323-1000311230103001-3011323330020032-2112330301132213-3101001222001130-1120000033033312-0013310033123310-2200111000013330"></a>

<a id="canonical-1332331131213200-2313122313222212-3301311320000033-0231121100000131-0201022323001020-0220023230222020-2212101202211232-1221300032000012"></a>

## location property — blindfold_secret_info / 123211022210 / 5

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

<a id="canonical-0312222030132233-3310301232100111-0123102310330112-2312302321111123-2122212332313313-0323312003000103-0331031002000001-2032210003302133"></a>

<a id="canonical-1200113301003302-0032100021001121-1222322132033323-2322331230122322-3312122232302201-0020203122231332-3020123011220333-2313221322311020"></a>

## store_provider property — blindfold_secret_info / 123211022210 / 6

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

<a id="canonical-1102033320220113-0313011213101022-0101110310011203-0032331312213030-2000200032010023-1030223133232020-1222122230033220-3313331333113201"></a>

## Next pages — blindfold_secret_info / 123211022210 / 7

- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-3033120122312321-1232030211110202-1231333230331023-0220211223003123-1001301212123113-1332022300231133-1202101132000213-1323130030112100)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1331231323220010-0121312121321032-3322332003211211-0231331310013203-3220330003223212-2200323001301102-2012220331112223-1122302003331203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122120101020301-0300211210303110-3220120200021021-0303102233300200-3233220021231003-3303312020131013-3032231023320312-0313213031231222"></a>

## admin_password.clear_secret_info — clear_secret_info / 033003231001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-3033120122312321-1232030211110202-1231333230331023-0220211223003123-1001301212123113-1332022300231133-1202101132000213-1323130030112100)
- admin_password.clear_secret_info

<a id="canonical-2303013020300122-0330200320101101-3133310013313313-0213301301233221-2032213103310013-2201230221023201-0321013132220310-1330220333330112"></a>

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

<a id="canonical-1332012310113200-2210222112223120-1223303330212112-3313233312322300-3132020202321021-1323101112320212-1031331021200101-2210310220132032"></a>

## Direct properties — clear_secret_info / 033003231001 / 3

<a id="canonical-0311012200213322-2231231322030001-3231132033012211-1310003132123202-2011031112020222-0303131233012032-1301032320130020-3030022030102320"></a>

<a id="canonical-3210032010110003-0320021121231023-2120301102301002-1101101102133101-1213111223113302-2311221113310310-3201130031310222-3100201101033301"></a>

## provider_ref property — clear_secret_info / 033003231001 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0312332233132230-3100112230202302-3221100130000211-3023101310011200-1003302330012203-0000022332122223-1332030332012213-0121011001011331"></a>

<a id="canonical-1221333032131030-2223323023312223-3123332020112011-1130331001211202-1022132330300110-0111311033213100-3311332201113321-0010032011120221"></a>

## URL property — clear_secret_info / 033003231001 / 5

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

<a id="canonical-0111213122331322-0021220230233111-2130312221020121-2020211220223021-1101132200231012-0122312333200102-1100202300232020-1300023021322311"></a>

## Next pages — clear_secret_info / 033003231001 / 6

- [admin_password](data-sources--aws_vpc_site--reference--group-001.md#canonical-3033120122312321-1232030211110202-1231333230331023-0220211223003123-1001301212123113-1332022300231133-1202101132000213-1323130030112100)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1021223323310232-3313030321212322-3100201130201122-3102101130303133-3333203221111132-1331100310300302-0012011220131003-1031021302011220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001103321313210-0031202101132122-3133213233113310-0030001301031102-3320312323001000-0111213300212312-2301320221202110-2212303202210212"></a>

## aws_cred — aws_cred / 003203322202 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- aws_cred

<a id="canonical-3231021020323311-3321311101312012-2030200331132323-2020331010203020-3310022220220312-1123003120033332-0021010231232030-1200113112330002"></a>

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

<a id="canonical-3302331310313033-0330312303113311-0132221231312210-3333122113003223-1011012233022222-2322123311120003-1313101223233032-2220212221310013"></a>

## Direct properties — aws_cred / 003203322202 / 3

<a id="canonical-3301323030101012-1102113133113212-3020101120130233-3201313300210322-2032213022333131-3122301011112322-2103023223213331-0011032201031302"></a>

<a id="canonical-0202323013121310-2213320031132300-3232302301302001-0001133201202321-2012210021232323-3032230012001303-1303102232123100-2032133212222012"></a>

## name property — aws_cred / 003203322202 / 4

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

<a id="canonical-0200321110300301-1231113021131300-1110211012012121-3102211123120212-3222231321302202-3110323001312301-2102111200233332-1023300031131013"></a>

<a id="canonical-3322210200311021-1201013103011100-1133221303210312-3023002102032311-3311133033203111-2333111110020002-0001201232122311-1223323112233230"></a>

## namespace property — aws_cred / 003203322202 / 5

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

<a id="canonical-2211210101303331-2111303022331133-0230113323003002-2210123110310101-3002201311110313-2333323112220012-2220231133231131-0110101131131120"></a>

<a id="canonical-1131012310331221-1202313122302112-2012322110010130-1330322133203001-1222123230211333-2020201211312011-0300001123221301-1233033213110203"></a>

## tenant property — aws_cred / 003203322202 / 6

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

<a id="canonical-0212001202123033-3101021330213330-0221021312203121-2300021202021223-3022130313010122-1002000202022201-2133102203010131-3212222121332012"></a>

## Next pages — aws_cred / 003203322202 / 7

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2313123122311312-3011032310230121-3320322022011110-0211200103023020-2002012333213223-1303323300222123-2110302033102331-2021103312222313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112312131033302-3322213003213232-2202123112300301-1032232231032012-1220131102001323-1223222010131203-3131113131132300-3213132110201200"></a>

## block_all_services — block_all_services / 131320221232 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- block_all_services

<a id="canonical-3130231112002020-2021330113332310-3331010023201200-0211022020302201-3211023133110013-0003022233113230-3220200331201220-0211232120021100"></a>

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

- [block_all_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-3130231112002020-2021330113332310-3331010023201200-0211022020302201-3211023133110013-0003022233113230-3220200331201220-0211232120021100)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-0001311020011020-1011301313323232-1321003233302322-3023023212303232-3000130200201222-3203311213023123-3112231300103210-3200012212130121)
- [default_blocked_services](data-sources--aws_vpc_site--reference--group-002.md#canonical-0112331121222233-1000312311001331-0210012000332221-2210212100220022-0023032332321003-2000320301231132-1122103102220310-1021000111332010)

Select alternatives according to the provider validators above.

<a id="canonical-2302120030333301-2222111021011122-0113120112222100-1011013012022001-3333131301112012-2103202132002310-3233300102110001-0133100003113111"></a>

## Direct properties — block_all_services / 131320221232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000002233231020-3310312302122311-1332010130110022-3022011202220100-1321200030032013-0113231132302222-2031111323301203-1323110320010313"></a>

## Next pages — block_all_services / 131320221232 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3100030222021212-2102031222223121-0110213010203222-3023232123211000-1230132322331222-0300220230232211-3000230113130333-1230031121313300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132120121130330-3113200201211301-2023001333122333-3003130212213210-3000331020202032-2311302120313003-2031201321201132-0112112333322203"></a>

## blocked_services — blocked_services / 232231301210 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- blocked_services

<a id="canonical-0001311020011020-1011301313323232-1321003233302322-3023023212303232-3000130200201222-3203311213023123-3112231300103210-3200012212130121"></a>

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

<a id="canonical-0030030223031101-2113323312132333-0303220202111012-0312202333012313-2122202110010231-3322001003102232-0310202132303213-3212321220132231"></a>

## Direct properties — blocked_services / 232231301210 / 3

- [blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100020001012032-2213132300121230-1330011003111130-1012302221130320-0201100032013112-0333010001201121-0211012020011102-1020322201230321): complete subsection reference.

<a id="canonical-2231322313002311-2002313123130301-0211300212023032-2213302221033000-2320132122333233-2301233112330000-2133323301121002-2110223212023010"></a>

## Next pages — blocked_services / 232231301210 / 4

- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100020001012032-2213132300121230-1330011003111130-1012302221130320-0201100032013112-0333010001201121-0211012020011102-1020322201230321)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3100020001012032-2213132300121230-1330011003111130-1012302221130320-0201100032013112-0333010001201121-0211012020011102-1020322201230321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112120212323202-2331110313033020-0322213123100120-3123223322321223-3220122112232123-3300302333101302-2301033003020201-2030123231032000"></a>

## blocked_services.blocked_service — blocked_service / 122303231313 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100030222021212-2102031222223121-0110213010203222-3023232123211000-1230132322331222-0300220230232211-3000230113130333-1230031121313300)
- blocked_services.blocked_service

<a id="canonical-3110131312102001-2113010021211122-0220302021233203-0203332220133202-2012003300323213-3330120331320120-3100213011031133-3313233223110320"></a>

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

<a id="canonical-3120103130210301-1021122202220033-0323132301331222-3212000121221020-3121103331021203-1032310132103300-1333201120110211-1232022002333120"></a>

## Direct properties — blocked_service / 122303231313 / 3

- [dns](data-sources--aws_vpc_site--reference--group-001.md#canonical-3031132113210203-2130332111102301-3131023132020312-2300303002220300-2221320300031300-3210031130320031-0331133011002003-1323120131203100): complete subsection reference.

<a id="canonical-0221332310002312-1120130220000100-2031211030223013-3012030022321123-2102213332201220-2223230210220302-2111120313012220-0111033210321121"></a>

<a id="canonical-1002301113202010-0210100112231313-1320333231111002-1033030312213020-0120102321333033-0230230233233231-3102102311030210-2122222131000313"></a>

## network_type property — blocked_service / 122303231313 / 4

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

- [ssh](data-sources--aws_vpc_site--reference--group-001.md#canonical-2322231121211021-2303121013122311-3122211232312110-0201211031121300-3332121111321110-1213110033323103-2212013232231210-1001303022213122): complete subsection reference.

- [web_user_interface](data-sources--aws_vpc_site--reference--group-001.md#canonical-0212212313003220-1001121113301312-1221222222120110-2031313213323310-3101111113312303-0303131033133333-3021221212301103-3332133230232203): complete subsection reference.

<a id="canonical-1322303301313203-3311131313133030-3121122010021033-1031122331301023-3220023321132301-0332332021123211-2320003220022321-0010020013023230"></a>

## Next pages — blocked_service / 122303231313 / 5

- [blocked_services.blocked_service.dns](data-sources--aws_vpc_site--reference--group-001.md#canonical-3031132113210203-2130332111102301-3131023132020312-2300303002220300-2221320300031300-3210031130320031-0331133011002003-1323120131203100)
- [blocked_services.blocked_service.ssh](data-sources--aws_vpc_site--reference--group-001.md#canonical-2322231121211021-2303121013122311-3122211232312110-0201211031121300-3332121111321110-1213110033323103-2212013232231210-1001303022213122)
- [blocked_services.blocked_service.web_user_interface](data-sources--aws_vpc_site--reference--group-001.md#canonical-0212212313003220-1001121113301312-1221222222120110-2031313213323310-3101111113312303-0303131033133333-3021221212301103-3332133230232203)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100030222021212-2102031222223121-0110213010203222-3023232123211000-1230132322331222-0300220230232211-3000230113130333-1230031121313300)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3031132113210203-2130332111102301-3131023132020312-2300303002220300-2221320300031300-3210031130320031-0331133011002003-1323120131203100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120211211122323-2231321123133100-0021321002020320-0013221232010211-1000103122003300-0100213023200033-2300230001110231-3312210223313031"></a>

## blocked_services.blocked_service.dns — dns / 011322323012 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100030222021212-2102031222223121-0110213010203222-3023232123211000-1230132322331222-0300220230232211-3000230113130333-1230031121313300)
- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100020001012032-2213132300121230-1330011003111130-1012302221130320-0201100032013112-0333010001201121-0211012020011102-1020322201230321)
- blocked_services.blocked_service.dns

<a id="canonical-2223021121313210-0031022303333321-2022012323312323-2203231233031022-0010332222021003-0003100201210121-0012230133330110-2022132330010323"></a>

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

<a id="canonical-3210130013321103-0120110033022133-3330033102033031-2110303201202120-2002133221232201-2101332000000310-1201103103200022-3032331033301320"></a>

## Direct properties — dns / 011322323012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123030200113230-3210200201233311-2033301110032312-0030103232003011-3210222320030020-0222023300133212-0110300132030230-3221103331033300"></a>

## Next pages — dns / 011322323012 / 4

- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100020001012032-2213132300121230-1330011003111130-1012302221130320-0201100032013112-0333010001201121-0211012020011102-1020322201230321)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2322231121211021-2303121013122311-3122211232312110-0201211031121300-3332121111321110-1213110033323103-2212013232231210-1001303022213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223322232131323-1023012222110133-0303332320301301-3113323223200021-3200010030001331-0233123212121013-2131123032322302-3233330332113100"></a>

## blocked_services.blocked_service.ssh — ssh / 002100210311 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100030222021212-2102031222223121-0110213010203222-3023232123211000-1230132322331222-0300220230232211-3000230113130333-1230031121313300)
- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100020001012032-2213132300121230-1330011003111130-1012302221130320-0201100032013112-0333010001201121-0211012020011102-1020322201230321)
- blocked_services.blocked_service.ssh

<a id="canonical-3201022000300022-1212110110320301-1123000322133032-2110001131000221-0301322112333000-0220223200121122-3013332133010211-1232031101300210"></a>

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

<a id="canonical-0303213233100013-0001333103221311-3312321102102323-3030311132013332-1121231132023221-1232210110120132-0313101002031102-1121232313001123"></a>

## Direct properties — ssh / 002100210311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122110210331330-3201311213113232-2313323233322201-1010011000001022-0213030231111021-3310022210312112-3133330121311201-2332131220101110"></a>

## Next pages — ssh / 002100210311 / 4

- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100020001012032-2213132300121230-1330011003111130-1012302221130320-0201100032013112-0333010001201121-0211012020011102-1020322201230321)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0212212313003220-1001121113301312-1221222222120110-2031313213323310-3101111113312303-0303131033133333-3021221212301103-3332133230232203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213021022131033-2130013132110011-1023020331122201-2312222032001022-0322332223023112-1213011031312231-3023130321121121-3031321232121130"></a>

## blocked_services.blocked_service.web_user_interface — web_user_interface / 133123002122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [blocked_services](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100030222021212-2102031222223121-0110213010203222-3023232123211000-1230132322331222-0300220230232211-3000230113130333-1230031121313300)
- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100020001012032-2213132300121230-1330011003111130-1012302221130320-0201100032013112-0333010001201121-0211012020011102-1020322201230321)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-0013301123211332-0232301201203201-3213012300313222-1133233113323112-3020322310321010-3031012131211302-3320222220202000-1130021202113203"></a>

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

<a id="canonical-3213003221022122-2131321131233123-2010222103312012-3211102323000001-2203133023122212-2001211011033210-0102322020033210-3132033231300211"></a>

## Direct properties — web_user_interface / 133123002122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000222021323202-2202013123123032-1200031033110202-2020003021210302-1112323201020300-0221311302131332-1233230012331303-0111210030031313"></a>

## Next pages — web_user_interface / 133123002122 / 4

- [blocked_services.blocked_service](data-sources--aws_vpc_site--reference--group-001.md#canonical-3100020001012032-2213132300121230-1330011003111130-1012302221130320-0201100032013112-0333010001201121-0211012020011102-1020322201230321)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2032033333332111-1230303113131311-1210120103333310-1210111013120310-0310221101302230-2123132000131001-3221322310230220-0013231203130212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302010210333113-2320210211123332-2003201001330230-0313131101312023-1112230003203132-1031300122232103-0220320002331013-2210203032010311"></a>

## coordinates — coordinates / 032210200303 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- coordinates

<a id="canonical-0320320211033001-2200120210223123-3220230320233022-0012302013112130-2303220321111212-0332233311332102-1032133220113101-0122031331203223"></a>

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

<a id="canonical-1001123330012310-2113222322311010-3113023321212121-3203310000132011-1211232103331303-1221222311222322-3301102100301321-1210321232013003"></a>

## Direct properties — coordinates / 032210200303 / 3

<a id="canonical-1300211331110330-3200322013003001-2131203121103200-2003032310010212-0013100213213231-3323332202100233-0232000312012130-3222101011310131"></a>

<a id="canonical-2112121330103001-0021221310232302-3000030123111300-1301222213231331-0231320333113110-2103033330333200-0123110020313320-2130202333203213"></a>

## latitude property — coordinates / 032210200303 / 4

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

<a id="canonical-3310122100011313-1213030023302103-3002112230133002-2302000233210303-0133330313301202-3320231022121221-1323100022332221-1001301001020220"></a>

<a id="canonical-3321200113031020-0223013120332122-2300200130123322-3213013120323003-3101132001120301-3220131310000211-3210203002203203-2011202030320132"></a>

## longitude property — coordinates / 032210200303 / 5

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
